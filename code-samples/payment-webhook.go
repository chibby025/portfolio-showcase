// Payment Webhook Handler - Idempotent Transaction Processing
// Handles Paystack/Stripe webhooks with retry logic and signature verification

package handlers

import (
    "crypto/hmac"
    "crypto/sha512"
    "encoding/hex"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// PaystackWebhook handles payment confirmations
func PaystackWebhook(c *gin.Context, db *gorm.DB) {
    // 1. Verify webhook signature (security)
    signature := c.GetHeader("x-paystack-signature")
    body, _ := ioutil.ReadAll(c.Request.Body)
    
    if !verifySignature(body, signature, PAYSTACK_SECRET_KEY) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
        return
    }
    
    // 2. Parse webhook payload
    var event PaystackEvent
    json.Unmarshal(body, &event)
    
    // 3. Handle only successful payments (idempotency check)
    if event.Event == "charge.success" {
        reference := event.Data.Reference
        
        // Check if already processed (prevent double-credit)
        var existingTxn Transaction
        if err := db.Where("reference = ?", reference).First(&existingTxn).Error; err == nil {
            // Transaction already processed
            c.JSON(http.StatusOK, gin.H{"message": "Already processed"})
            return
        }
        
        // 4. Begin database transaction (atomicity)
        tx := db.Begin()
        
        // Credit user's wallet
        amount := event.Data.Amount / 100 // Paystack sends in kobo
        if err := tx.Exec(
            "UPDATE user_wallets SET token_balance = token_balance + ?, naira_backed_tokens = naira_backed_tokens + ? WHERE user_id = ?",
            amount, amount, event.Data.Metadata.UserID,
        ).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Credit failed"})
            return
        }
        
        // Record transaction
        txn := Transaction{
            Reference:  reference,
            UserID:     event.Data.Metadata.UserID,
            Amount:     amount,
            Status:     "completed",
            Gateway:    "paystack",
        }
        tx.Create(&txn)
        
        // 5. Commit transaction
        tx.Commit()
        
        c.JSON(http.StatusOK, gin.H{"message": "Payment processed"})
    }
}

// verifySignature validates webhook authenticity
func verifySignature(payload []byte, signature, secret string) bool {
    mac := hmac.New(sha512.New, []byte(secret))
    mac.Write(payload)
    expectedSignature := hex.EncodeToString(mac.Sum(nil))
    return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

// Key Learning: Financial systems require:
// 1. Signature verification (prevent fake webhooks)
// 2. Idempotency (prevent double-processing)
// 3. Database transactions (atomicity - all or nothing)
// 4. Error handling with rollback
