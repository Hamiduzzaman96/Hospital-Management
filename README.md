# Hospital-Management


A simple RESTful API built with Go for managing hospitals, doctors, and their relationships.  
It supports two main user roles:

- **Super Admin** — can create, list, and delete hospitals
- **Hospital Admin** — can manage doctors (create, update, delete, assign/remove from their hospital)

## Features

- JWT-based authentication & authorization
- Role-based access control
- Hospital CRUD operations (Super Admin only)
- Doctor CRUD operations (Hospital Admin only)
- Assign / remove / list doctors for a hospital
- Pagination + search support on lists
- PostgreSQL database with automatic migrations
- Clean architecture structure
---------------------------------------------------------------------------------------------------------------------------------------
**Requirements**

- **Language**: Go 1.25.5
- **HTTP Router**: Standard library `net/http` + `http.ServeMux`
- **Database**: PostgreSQL
- **Migrations**: golang-migrate
- **Authentication**: JWT (`golang-jwt/jwt/v5`)
- **Password Hashing**: bcrypt
- **Environment**: godotenv
- **Architecture**: Domain → Usecase → Repository → Handler
--------------------------------------------------------------------------------------------------------------------------------------
## Project Structure

**A. internal ---->** 
        1.Domain
            a.doctor.go
            b.hospital_doctor_relationship.go
            c.hospital.go
            d.user.go

        2.Repository 
            a. doctor_repo.go
            b.hospital_doctor_ralationship_repo.go
            c.hospital_repo.go
            d.user_repository.go

        3.Usecase
            a.doctor_usecase.go
            b.hospital_doctor_usecase.go
            c.hospital_usecase.go
            d.user_usecase.go

        4.Infrastructure
            a.config\config.go
            b.databse\database_connection.go

        5.Middleware
            a.role_middleware.go
            b.user_middleware.go

        6.Handler
            a.doctor
                i.create.go
                ii.delete.go
                iii.get_by_id.go
                iv.list.go
                v.update.go

            b.hospital
                i.create.go
                ii.delete.go
                iii.get_by_id.go
                iv.list.go

            c. hospital_doctor
                i.hospital_doctor_handler.go

            d.user
                i.login.go
                ii.register.go
                iii.user_handler.go

        7.Router
            a.router.go 

**B. pkg\ helper.go**

**C. main.go**

**D. .env.example**

**E. README.md**
    
----------------------------------------------------------------------------------------------------------------------  
## Setup Instructions (Local Development)

### Prerequisites

- Go 1.25.5
- require (
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/golang-migrate/migrate/v4 v4.19.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.47.0
)

--------------------------------------------------------------------------------------------------------------------------
**Clone the repository**
   git clone https://github.com/Hamiduzzaman96/Hospital-Management.git

**update: go mod tidy**

**Run the Server**
go run main.go or go run .

-------------------------------------------------------------------------------------------------------------------------
**API Endpoints**
**Method,                Endpoint,                 Description,                           Required Role**
POST,               /users/register,           Register new user,                            Public
POST,               /users/login,             Login & get JWT token,                          Public

POST,               /hospitals,               Create new hospital,                          super_admin
GET,                /hospitals,                 List hospitals (search + pagination),       authenticated

POST,               /doctors,                  Create new doctor,                           hospital_admin
PUT,                /doctors,                   Update doctor,                              hospital_admin
DELETE,             /doctors,                    Delete doctor,                              hospital_admin
GET,                /doctors,                    List doctors (search + pagination),         authenticated

POST,               /hospitals/doctors/assign,    Assign doctor to hospital,                 hospital_admin
DELETE,             /hospitals/doctors,           Remove doctor from hospital,               hospital_admin
GET,                /hospitals/doctors,             List doctors of the hospital,              hospital_admin


