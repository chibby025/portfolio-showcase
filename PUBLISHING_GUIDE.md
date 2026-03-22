# Publishing Your Portfolio to GitHub

## Step 1: Initialize Git Repository

```bash
cd ~/portfolio-showcase

# Initialize git
git init

# Add all files
git add .

# First commit
git commit -m "Initial portfolio commit - WeWatch showcase"
```

---

## Step 2: Create GitHub Repository

1. Go to: https://github.com/new
2. Repository name: `portfolio-showcase` or your preferred name
3. **Make it PUBLIC** ✅ (this is for Antler to see)
4. Description: "Full-Stack Engineer Portfolio - WeWatch Platform"
5. **DON'T initialize with README** (you already have one)
6. Click "Create repository"

---

## Step 3: Connect & Push to GitHub

```bash
# Replace 'yourusername' with your actual GitHub username
git remote add origin https://github.com/yourusername/portfolio-showcase.git

# Push to GitHub
git branch -M main
git push -u origin main
```

---

## Step 4: Personalize README.md

**Before publishing, update these placeholders in README.md:**

1. **Contact Information**
   ```markdown
   - Email: your.email@example.com  → your.real.email@gmail.com
   - LinkedIn: [linkedin.com/in/yourprofile]  → your real profile
   - GitHub: [@yourusername]  → your real username
   ```

2. **GitHub Stats**
   ```markdown
   Replace 'yourusername' with your actual GitHub username in the badge URLs
   ```

3. **Links**
   - Demo video link (record after crypto implementation)
   - Portfolio website (if you have one)
   - Social media handles

---

## Step 5: Add Screenshots

**Required screenshots** (take these from your WeWatch platform):

```bash
# Navigate to screenshots folder
cd screenshots/

# Add your actual screenshots here
# Rename them to match README.md references:
# - cinema-hall.png
# - payment-modal.png
# - admin-dashboard.png
# - chat-reactions.png
```

**How to take good screenshots:**
1. Use full HD resolution (1920x1080)
2. Use light theme if possible
3. Show real-looking data (not empty)
4. Blur any sensitive information
5. Use browser dev tools to simulate mobile if needed

**After adding screenshots:**
```bash
git add screenshots/
git commit -m "Add platform screenshots"
git push
```

---

## Step 6: Verify on GitHub

1. Go to: `https://github.com/yourusername/portfolio-showcase`
2. Check that README renders correctly
3. Verify all images load (screenshots)
4. Test all links work
5. Check mobile view (GitHub has responsive preview)

---

## Step 7: Share with Antler

**In your Antler application, use:**

```
GitHub Portfolio: https://github.com/yourusername/portfolio-showcase

Note: Full WeWatch codebase is private (pre-launch startup).
This portfolio showcases technical approach, architecture, and code quality.
Demo available upon request.
```

---

## Step 8: Optional Enhancements

### Add GitHub Pages (Free Hosting)

Want a nice URL like `yourusername.github.io/portfolio-showcase`?

1. Go to repository Settings → Pages
2. Source: Deploy from branch `main`
3. Folder: `/ (root)`
4. Save
5. Wait 2-3 minutes
6. Visit: `https://yourusername.github.io/portfolio-showcase`

### Add Demo Video

Record a 3-5 minute demo showing:
1. 3D Cinema experience
2. Payment flow (crypto after implementation)
3. Admin dashboard
4. Live chat/reactions

Upload to:
- YouTube (unlisted or public)
- Loom (free screen recording)
- Drive (shareable link)

Then update README.md with link.

### Add More Code Samples

If you have other interesting implementations:
- Authentication middleware
- Rate limiting
- Caching strategies
- Testing examples

Add them to `code-samples/` folder.

---

## Checklist Before Sharing

- [ ] Personal information updated (email, LinkedIn, etc.)
- [ ] GitHub username replaced in badge URLs
- [ ] Screenshots added (at least 4 key ones)
- [ ] All links tested and working
- [ ] Code samples reviewed (no secrets exposed)
- [ ] README grammar/spelling checked
- [ ] Repository is PUBLIC
- [ ] Committed and pushed to GitHub
- [ ] Verified rendering on GitHub.com
- [ ] URL copied for Antler application

---

## 🚨 Security Reminders

**NEVER commit these to ANY public repo:**
- Database credentials
- API keys (Paystack, Stripe, Coinbase, etc.)
- JWT secrets
- `.env` files
- User data (emails, passwords, transactions)
- Business logic (pricing algorithms, commission rates)

**This portfolio repo contains ONLY:**
- Sanitized code samples
- Architecture descriptions
- Screenshots (with blurred sensitive data)
- Your technical skills showcase

---

## Need Help?

**Common Issues:**

**"Permission denied (publickey)"**
```bash
# Add SSH key to GitHub
ssh-keygen -t ed25519 -C "your.email@example.com"
cat ~/.ssh/id_ed25519.pub
# Copy output and add to GitHub: Settings → SSH Keys
```

**"Screenshots not showing"**
- Check file paths in README.md match actual filenames
- Ensure files are committed: `git add screenshots/`
- Push changes: `git push`

**"README looks broken on GitHub"**
- Preview locally first: Install VS Code extension "Markdown Preview"
- Or use online tool: https://dillinger.io/

---

**Once published, send me the link and I'll review it before you submit to Antler!**
