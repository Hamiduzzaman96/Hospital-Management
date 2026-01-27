package main

import (
	"log"
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/infrastructure/config"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/infrastructure/database"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/router"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Overload()
	if err != nil {
		log.Println(".env not found")
	}

	cfg := config.LoadConfig()
	db := database.NewDatabaseConnection(cfg)

	userRepo := repository.NewUserRepository(db)
	hospitalRepo := repository.NewHospitalRepository(db)
	doctorRepo := repository.NewDoctorRepository(db)
	hosDocRepo := repository.NewHospitalDoctorRelationship(db)

	userUC := usecase.NewJwtUsecase(userRepo, cfg.JWT_Secret)
	hospitalUC := usecase.NewHospitalUsecase(hospitalRepo)
	doctorUC := usecase.NewDoctorUsecase(doctorRepo)
	hosDocUC := usecase.NewHospitalDoctorUsecase(*hosDocRepo)

	router := router.NewRouter(cfg, userUC, hospitalUC, doctorUC, hosDocUC)
	log.Printf("Server starting on :%s", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, router))

}
