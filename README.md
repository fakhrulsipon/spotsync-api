# SpotSync API - Smart Parking & EV Charging Reservation

## Live Deployment URL
- Backend Live URL: https://spotsync-api.onrender.com

## Tech Stack
- Go (Golang v1.22+)
- Echo Framework v4
- GORM with PostgreSQL (NeonDB Pool Configured)
- JWT & Bcrypt Architecture Security
- Validator v10

## Architecture Details
This project strictly follows **Clean Architecture Layout Rules**:
- `Models`: GORM database schema.
- `Repository`: Direct atomic DB calculations using transaction-level row updates (`FOR UPDATE`).
- `Service`: Business definitions.
- `Handler`: Network serialization & bindings.