# SpotSync API 🚗💨

SpotSync API is a high-performance smart parking reservation backend built with **Go (Golang)**. It uses **Clean Architecture** and **Row-Level Locking (`FOR UPDATE`)** inside database transactions to handle high concurrency and prevent over-booking.

🔗 **Live URL:** [https://spotsync-api-ivgz.onrender.com](https://spotsync-api-ivgz.onrender.com)

---

## 🛠️ Tech Stack
- **Language:** Go (Golang 1.22+)
- **Framework:** Echo v4
- **ORM:** GORM
- **Database:** NeonDB (Cloud PostgreSQL)
- **Token:** JWT (JSON Web Tokens)
- **Deployment:** Render

---

## 📐 Architecture

The project follows a strict layered structure for separation of concerns:

```text
HTTP Request ➔ Handlers (DTO Validation) ➔ Services (Business Logic) ➔ Repositories (GORM DB Access) ➔ NeonDB
Handlers: Validate input payloads and send JSON responses.

Services: Orchestrate business rules and map models to clean DTOs.

Repositories: Execute raw database operations and handle atomic row-level locks (FOR UPDATE).

🚀 Local Setup
1. Clone & Navigate
Bash
git clone https://github.com/YOUR_USERNAME/spotsync-api.git
cd spotsync-api
2. Configure Environment (.env)
Create a .env file in the root folder:

Code snippet
PORT=8080
DB_HOST=your-neondb-cluster.neon.tech
DB_USER=your_neon_username
DB_PASSWORD=your_neon_password
DB_NAME=spotsync
DB_PORT=5432
JWT_SECRET=supersecretspotsyncjwtkey2026
3. Install & Run
Bash
go mod tidy
go run main.go
The server will start at http://localhost:8080.

📬 API Endpoints
🔐 Authentication
POST /api/v1/auth/register - Register a user (driver or admin)

POST /api/v1/auth/login - Login user & get JWT Token

🚗 Parking Zones
POST /api/v1/zones - Create a new zone (Admin Only)

GET /api/v1/zones - Get all zones with dynamic available slots

GET /api/v1/zones/:id - Get a specific zone by ID

📝 Reservations (Concurrency Protected)
POST /api/v1/reservations - Create a secure booking (Enforces Capacity)

GET /api/v1/reservations/my-reservations - Get logged-in user's bookings

DELETE /api/v1/reservations/:id - Cancel a reservation (Ownership Enforced)

GET /api/v1/reservations - View all historical reservations (Admin Only)