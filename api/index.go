package api

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
	"net/http"
	"time-clock/controllers"
	"time-clock/external"
	timeclockgateway "time-clock/gateways/db/timeclock"
	usergateway "time-clock/gateways/db/user"
	"time-clock/usecases/user"
)

func SetupDB() *gorm.DB {
	dialector := external.GetPostgresDialector()
	db := external.NewORM(dialector)

	return db
}

func SetupRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections

	// Gateways
	userGateway := usergateway.NewGateway(orm)
	timeClockGateway := timeclockgateway.NewGateway(orm)
	// Use cases
	userUseCase := user.NewUseCase(userGateway, timeClockGateway)
	// Handlers
	_ = controllers.NewUserController(userUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
