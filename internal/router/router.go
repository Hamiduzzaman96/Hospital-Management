package router

import (
	"net/http"

	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/domain"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/handler/doctor"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/handler/hospital"
	hospitaldoctor "github.com/Hamiduzzaman96/Hospital-Management.git/internal/handler/hospital_doctor"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/handler/user"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/infrastructure/config"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/middleware"
	"github.com/Hamiduzzaman96/Hospital-Management.git/internal/usecase"
)

type Router struct {
	mux *http.ServeMux

	userHandler           *user.UserHandler
	hospitalHandler       *hospital.HospitalHandler
	doctorHandler         *doctor.DoctorHandler
	hospitalDoctorHandler *hospitaldoctor.HospitalDoctorHandler

	authMiddleware func(http.Handler) http.Handler
}

func NewRouter(
	cfg *config.Config,
	userUC *usecase.UserUsecase,
	hospitalUC *usecase.HospitalUsecase,
	doctorUC *usecase.DoctorUsecase,
	hosDocUC *usecase.HospitalDoctorUsecase,
) *Router {
	r := &Router{
		mux:                   http.NewServeMux(),
		userHandler:           user.NewUserHandler(userUC),
		hospitalHandler:       hospital.NewHospitalHandler(hospitalUC),
		doctorHandler:         doctor.NewDoctorHandler(doctorUC),
		hospitalDoctorHandler: hospitaldoctor.NewHospitalDoctorHandler(*hosDocUC),
		authMiddleware:        middleware.NewAuthMiddleware(cfg.JWT_Secret),
	}

	r.mux.HandleFunc("POST /api/v1/auth/register", r.userHandler.Register)
	r.mux.HandleFunc("POST /api/v1/auth/login", r.userHandler.Login)

	//Hospital Routes - SuperAdmin
	r.mux.Handle("POST /api/v1/hospitals",
		r.authMiddleware(
			middleware.NewRoleMiddleware(domain.SuperAdmin)(
				http.HandlerFunc(r.hospitalHandler.Create),
			),
		),
	)

	r.mux.Handle("DELETE /api/v1/hospitals",
		r.authMiddleware(
			middleware.NewRoleMiddleware(domain.SuperAdmin)(
				http.HandlerFunc(r.hospitalHandler.Delete),
			),
		),
	)
	r.mux.Handle("GET /api/v1/hospitals/{id}",
		r.authMiddleware(
			http.HandlerFunc(r.hospitalHandler.GetByID),
		),
	)
	r.mux.Handle("GET /api/v1/hospitals",
		r.authMiddleware(
			http.HandlerFunc(r.hospitalHandler.List),
		),
	)

	// Doctor Routers - HospitalAdmin
	r.mux.Handle("POST /api/v1/doctors",
		r.authMiddleware(
			middleware.NewAuthMiddleware(domain.HospitalAdmin)(
				http.HandlerFunc(r.doctorHandler.Create),
			),
		),
	)

	r.mux.Handle("PUT /api/v1/doctors",
		r.authMiddleware(
			middleware.NewRoleMiddleware(domain.HospitalAdmin)(
				http.HandlerFunc(r.doctorHandler.Update),
			),
		),
	)

	r.mux.Handle("DELETE /api/v1/doctors",
		r.authMiddleware(
			middleware.NewRoleMiddleware(domain.HospitalAdmin)(
				http.HandlerFunc(r.doctorHandler.Delete),
			),
		),
	)

	r.mux.Handle("GET /api/v1/doctors",
		r.authMiddleware(
			http.HandlerFunc(r.doctorHandler.List),
		),
	)

	r.mux.Handle("GET /api/v1/doctors",
		r.authMiddleware(
			http.HandlerFunc(r.doctorHandler.List),
		),
	)

	// Hospital - Doctor - Relationship  Routes - HospitalAdmin
	r.mux.Handle("POST /api/v1/hospitals/doctors/assign",
		r.authMiddleware(
			middleware.NewRoleMiddleware(domain.HospitalAdmin)(
				http.HandlerFunc(r.hospitalDoctorHandler.AssignDoctor),
			),
		),
	)

	r.mux.Handle("DELETE /api/v1/hospitals/doctors",
		r.authMiddleware(
			middleware.NewRoleMiddleware(domain.HospitalAdmin)(
				http.HandlerFunc(r.hospitalDoctorHandler.ListByDoctor),
			),
		),
	)
	return r

}

func (rvr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rvr.mux.ServeHTTP(w, r)
}
