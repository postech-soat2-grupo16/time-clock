package api

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
	"net/http"
	"time-clock/controllers"
	"time-clock/external"
	timeclockgateway "time-clock/gateways/db/timeclock"
	usergateway "time-clock/gateways/db/user"
	"time-clock/gateways/notification"
	"time-clock/usecases/user"
)

func SetupDB() *gorm.DB {
	dialector := external.GetPostgresDialector()
	db := external.NewORM(dialector)

	return db
}

func SetupNotification() *sns.SNS {
	return external.GetSnsClient()
}

func SetupRouter(db *gorm.DB, sns *sns.SNS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db, sns)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB, sns *sns.SNS) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections

	// Gateways
	userGateway := usergateway.NewGateway(orm)
	timeClockGateway := timeclockgateway.NewGateway(orm)
	notificationGateway := notification.NewGateway(sns)
	// Use cases
	userUseCase := user.NewUseCase(userGateway, timeClockGateway, notificationGateway)
	// Handlers
	_ = controllers.NewUserController(userUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
