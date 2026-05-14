# Chibuzor's Software Engineering Portfolio

> **From Fired to Founder in 6 Months** | Full-Stack Engineer | Building Africa's First Web3-Enabled Social Streaming Platform

[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue)](https://www.linkedin.com/in/chibuzor-chinweokwu-5338bb31b/)
[![Email](https://img.shields.io/badge/Email-Contact-red)](mailto:chinweokwuchibuzor@gmail.com)
[![GitHub](https://img.shields.io/badge/GitHub-chibby025-black)](https://github.com/chibby025)

---

## 🎯 About Me

**Self-taught full-stack engineer** to building a complete social streaming platform in 6 months.

**Background:**



**What I Built While Learning:**
- Real-time 3D cinema with spatial audio
- Payment system processing real transactions (Paystack, Stripe, Coinbase Commerce)
- WebSocket-based live interactions (chat, reactions, presence)
- Admin dashboard with analytics
- Automated testing & deployment pipeline

---

## 🏆 Featured Project: WeWatch

**Africa's First Web3-Enabled Social Streaming Platform**

A platform where users watch movies together in immersive 3D environments, with real-time interactions, ticketing, and cryptocurrency payments.

### 🎬 Live Demo
- **Platform:** [Coming Soon - In Private Beta]
- **Demo Video:** [Watch 3-Minute Demo](https://youtube.com/placeholder)
- **Screenshots:** See below

### 📊 Key Metrics
- **Lines of Code:** ~50,000+ (Backend: Go, Frontend: React)
- **Development Time:** 6 months (solo)
- **Tech Stack:** 15+ technologies integrated
- **Payment Processing:** Live transactions (Paystack, Stripe, Coinbase Commerce)
- **Real-Time Users:** Tested with 50+ concurrent users

---

## 🛠️ Tech Stack

### Backend
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![WebSocket](https://img.shields.io/badge/WebSocket-010101?style=for-the-badge&logo=socket.io&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)

- **Framework:** Gin (Go web framework)
- **Database:** PostgreSQL with GORM ORM
- **Real-Time:** WebSocket (native Go)
- **Video Streaming:** LiveKit SDK
- **Authentication:** JWT with bcrypt
- **Payments:** Paystack, Stripe, Coinbase Commerce APIs
- **Deployment:** Railway (auto-deploy from Git)

### Frontend
![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![TypeScript](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Three.js](https://img.shields.io/badge/Three.js-000000?style=for-the-badge&logo=three.js&logoColor=white)

- **Library:** React 18 with Hooks
- **State Management:** Context API + Custom Hooks
- **3D Graphics:** Three.js + React Three Fiber
- **Styling:** Tailwind CSS + Custom Components
- **Build Tool:** Vite
- **Deployment:** Vercel

### DevOps & Tools
![Git](https://img.shields.io/badge/Git-F05032?style=for-the-badge&logo=git&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

- **Version Control:** Git + GitHub
- **CI/CD:** GitHub Actions (auto-deploy)
- **Database Migrations:** Custom SQL migration system
- **Testing:** Go testing package + React Testing Library
- **Monitoring:** Custom logging + error tracking

---

## 🎨 Platform Features

### 1. Immersive 3D Environments
- **Cinema Hall:** Movie theater with realistic seating (100+ seats)
- **Lecture Hall:** Educational setting with tiered seating
- **Video Lounge:** Casual watch party environment
- **Spatial Audio:** Position-based sound (hear users near you)

### 2. Real-Time Interactions
- **Live Chat:** WebSocket-based instant messaging
- **Reactions:** Emoji reactions with animations
- **User Presence:** See who's in the room, their position
- **Host Controls:** Mute, kick, pause/play for all

### 3. Payment System
- **Multi-Currency:** NGN, USD, GHS, KES, EUR, GBP
- **Payment Gateways:** Paystack (Africa), Stripe (International), Coinbase Commerce (Crypto)
- **Token Economy:** Buy/sell spread model (26% margin)
- **Withdrawals:** Automated payout system
- **KYC Integration:** Identity verification for hosts

### 4. Event Management
- **Scheduled Events:** Create events with start times
- **Ticketing:** Free, paid, early-bird pricing
- **RSVPs:** Attendance tracking
- **Donations:** Live tipping during sessions
- **Analytics:** Revenue tracking, user engagement

### 5. Admin Dashboard
- **User Management:** Search, suspend, role management
- **Revenue Analytics:** Real-time earnings breakdown
- **Session Monitoring:** Active sessions, user counts
- **Content Moderation:** Flagged content review
- **Platform Health:** System metrics, error logs

---

## 📸 Screenshots

### 🎥 3D Cinema Experience
![Cinema Hall](./screenshots/cinema-hall.png)
*Immersive 3D movie theater with 100+ seats and spatial audio*

### 💳 Payment Flow
![Token Purchase](./screenshots/payment-modal.png)
*Multi-gateway payment system (Paystack, Stripe, Coinbase Commerce)*

### 📊 Admin Dashboard
![Analytics](./screenshots/admin-dashboard.png)
*Real-time revenue analytics and user management*

### 💬 Live Interactions
![Chat System](./screenshots/chat-reactions.png)
*WebSocket-based chat with reactions and presence*

---

## 🏗️ Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                        FRONTEND                              │
│  React + Three.js + WebSocket + Tailwind                    │
│  Deployed on Vercel (Auto-deploy from Git)                  │
└────────────────────────┬────────────────────────────────────┘
                         │
                         │ HTTPS/WSS
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                     API GATEWAY (Go)                         │
│  - REST API (Gin framework)                                 │
│  - WebSocket Server (native Go)                             │
│  - JWT Authentication                                        │
│  - Rate Limiting                                             │
└────────────────────────┬────────────────────────────────────┘
                         │
          ┌──────────────┼──────────────┐
          ▼              ▼              ▼
    ┌─────────┐   ┌──────────┐   ┌──────────┐
    │PostgreSQL│   │ LiveKit  │   │ Payment  │
    │ Database │   │  Video   │   │ Gateways │
    │         │   │ Streaming │   │Paystack  │
    │- Users  │   │          │   │Stripe    │
    │- Wallets│   │          │   │Coinbase  │
    │- Sessions│   │          │   │          │
    └─────────┘   └──────────┘   └──────────┘
```

### Key Technical Decisions

**Why Go for Backend?**
- Native concurrency (goroutines) for WebSocket handling
- Fast compilation and execution
- Built-in HTTP/WebSocket support
- Excellent for high-performance APIs

**Why Three.js for 3D?**
- WebGL performance in browser
- React Three Fiber integration
- Rich ecosystem for 3D web apps
- No plugin installation required

**Why PostgreSQL?**
- ACID compliance for financial transactions
- JSON support for flexible metadata
- Excellent indexing for analytics queries
- Battle-tested at scale

---

## 💡 Key Technical Challenges Solved

### 1. Real-Time Synchronization
**Problem:** Keep 50+ users synchronized in 3D space  
**Solution:** 
- WebSocket broadcast with delta updates
- Position throttling (10 updates/sec max)
- Client-side interpolation for smooth movement

### 2. Payment Reliability
**Problem:** Handle payment gateway failures gracefully  
**Solution:**
- Webhook signature verification
- Idempotent transaction processing
- Retry logic with exponential backoff
- Database transactions for atomicity

### 3. Two-Account Problem (Crypto + Fiat)
**Problem:** Users buy with USDC but withdraw Naira - account runs dry  
**Solution:**
```sql
user_wallets:
  token_balance: 1000 (total)
  naira_backed_tokens: 600 (Paystack reserves)
  crypto_backed_tokens: 400 (Coinbase reserves)
```
Ensures each gateway only pays what it received.

### 4. 3D Performance Optimization
**Problem:** Browser crashes with 100+ 3D avatars  
**Solution:**
- Level of Detail (LOD) - simple models for distant users
- Frustum culling - don't render offscreen objects
- Instance rendering for duplicate objects
- Lazy loading of 3D assets

### 5. Concurrent WebSocket Connections
**Problem:** Race conditions with 1000+ concurrent messages  
**Solution:**
- Go channels for message queuing
- Read/Write mutexes on shared state
- Connection pooling
- Graceful connection cleanup

---

## 📈 What I Learned

### Technical Skills
- ✅ Go concurrency patterns (goroutines, channels, mutexes)
- ✅ WebSocket architecture at scale
- ✅ 3D web graphics (Three.js, React Three Fiber)
- ✅ Payment gateway integration (webhooks, idempotency)
- ✅ Database optimization (indexing, query performance)
- ✅ JWT authentication & authorization
- ✅ Real-time system design
- ✅ API design (RESTful + WebSocket)

### Soft Skills
- ✅ Self-directed learning (6 months, zero to production)
- ✅ Problem-solving under pressure
- ✅ Full-stack ownership (design → deploy)
- ✅ Technical documentation
- ✅ Resilience (fired → founder)

---

## 🎯 Current Focus

**Phase 1A: Cryptocurrency Integration** (In Progress)
- Integrating Coinbase Commerce for USDC/USDT payments
- Building dual-currency token economy
- Target: Web3 narrative for funding applications

**Phase 1B: Crypto Withdrawals** (Next)
- Circle API integration
- EVM-compatible wallet support
- Full crypto payment lifecycle

**Phase 2: NFT Event Tickets** (Planned Q2 2026)
- ERC-721 smart contracts on Polygon
- Proof of attendance collectibles
- Secondary marketplace

---

## 🚀 Open to Opportunities

**What I'm Looking For:**
- 💼 **Full-Stack/Backend roles** at startups or Web3 companies
- 🤝 **Freelance/Contract work** (Go, React, WebSocket, Payments)
- 💰 **Preseed/Seed funding** for WeWatch platform
- 🌍 **Remote opportunities** (based in Nigeria, open to relocation)

**What I Bring:**
- Proven ability to build production systems solo
- 3.8x efficiency (documented at previous role)
- Fast learner (zero → full-stack in 6 months)
- Resilience and grit
- Real-world payment processing experience
- Crypto/Web3 integration skills

---

## 📫 Contact

- **Email:** chinweokwuchibuzor@gmail.com
- **Phone:** +234 708 175 5601
- **LinkedIn:** [linkedin.com/in/chibuzor-chinweokwu](https://www.linkedin.com/in/chibuzor-chinweokwu-5338bb31b/)
- **GitHub:** [@chibby025](https://github.com/chibby025)

**Response Time:** Usually within 24 hours

---

## 📊 GitHub Stats

![GitHub Stats](https://github-readme-stats.vercel.app/api?username=chibby025&show_icons=true&theme=radical)

![Top Languages](https://github-readme-stats.vercel.app/api/top-langs/?username=chibby025&layout=compact&theme=radical)

---

## 🙏 Acknowledgments

Built with determination, coffee, and a lot of late nights.

Special thanks to:
- The Go community for excellent documentation
- React Three Fiber creators for making 3D accessible
- Everyone who believed in me when I was fired

---

**Last Updated:** March 2026  
**Status:** Actively seeking opportunities & funding

---

> "I don't just write code. I solve problems that matter."  
> – Chibuzor, Full-Stack Engineer
