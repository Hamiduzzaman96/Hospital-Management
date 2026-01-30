package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/infrastructure/config"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/infrastructure/database"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/repository"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/router"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
	"github.com/Hamiduzzaman96/Hospital-Management.git/pkg/helper"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Overload()
	if err != nil {
		log.Println(".env not found")
	}

	cfg := config.LoadConfig()
	db := database.NewDatabaseConnection(cfg)

	database.NewMigrations(db, cfg)

	userRepo := repository.NewUserRepository(db)
	hospitalRepo := repository.NewHospitalRepository(db)
	doctorRepo := repository.NewDoctorRepository(db)
	hosDocRepo := repository.NewHospitalDoctorRelationship(db)

	userUC := usecase.NewUserUsecase(userRepo, cfg.JWT_Secret)
	hospitalUC := usecase.NewHospitalUsecase(hospitalRepo)
	doctorUC := usecase.NewDoctorUsecase(doctorRepo)
	hosDocUC := usecase.NewHospitalDoctorUsecase(*hosDocRepo)

	router := router.NewRouter(cfg, userUC, hospitalUC, doctorUC, hosDocUC)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: helper.EnableCORS(router),
	}
	go func() {
		log.Println("server starting on :", cfg.AppPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server connection not working", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server shutdown", err)
	}

	db.Close()
	log.Println("Server stopped")

}
