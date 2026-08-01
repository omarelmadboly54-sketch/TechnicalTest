# Go E-Commerce Technical Test

A lightweight backend application built with **Go (Gin Framework)** featuring user authentication and protected routing.

## Tech Stack
* **Language:** Go
* **Framework:** Gin (`github.com/gin-gonic/gin`)
* **Storage:** Cookies for session management

## Setup & Run

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/omarelmadboly54-sketch/TechnicalTest.git](https://github.com/omarelmadboly54-sketch/TechnicalTest.git)
   cd TechnicalTest

2.Install dependencies:
go mod tidy


3.Run the server:
go run main.go

The application will run at http://127.0.0.1:8080.


Endpoints
GET / - Product catalog (Public)

GET /signin - Sign-in page

POST /api/signin - Authentication endpoint

GET /signup - Sign-up page

POST /api/signup - Registration endpoint

GET /checkout - Protected checkout page (Requires login cookie)
