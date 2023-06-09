package bootstrap

import (
	"be-user-scheme/pkg/logruslogger"
	api "be-user-scheme/server/handler"
	"be-user-scheme/server/middleware"

	chimiddleware "github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// RegisterRoutes ...
func (boot *Bootup) RegisterRoutes() {
	handlerType := api.Handler{
		DB:         boot.DB,
		EnvConfig:  boot.EnvConfig,
		Validate:   boot.Validator,
		Translator: boot.Translator,
		ContractUC: &boot.ContractUC,
		Jwe:        boot.Jwe,
		Jwt:        boot.Jwt,
	}
	mJwt := middleware.VerifyMiddlewareInit{
		ContractUC: &boot.ContractUC,
	}

	boot.R.Route("/v1", func(r chi.Router) {
		// Define a limit rate to 1000 requests per IP per request.
		rate, _ := limiter.NewRateFromFormatted("1000-S")
		store, _ := sredis.NewStoreWithOptions(boot.ContractUC.Redis, limiter.StoreOptions{
			Prefix:   "limiter_global",
			MaxRetry: 3,
		})
		rateMiddleware := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))
		r.Use(rateMiddleware.Handler)

		// Logging setup
		r.Use(chimiddleware.RequestID)
		r.Use(logruslogger.NewStructuredLogger(boot.EnvConfig["LOG_FILE_PATH"], boot.EnvConfig["LOG_DEFAULT"]))
		r.Use(chimiddleware.Recoverer)

		// API
		r.Route("/api", func(r chi.Router) {
			userHandler := api.UserHandler{Handler: handlerType}
			r.Route("/auth", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Post("/login", userHandler.LoginHandler)
				})
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Post("/logout", userHandler.LogoutHandler)
				})
			})

			r.Route("/user", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Get("/", userHandler.TokenHandler)
					r.Get("/all", userHandler.FindAllHandler)
				})
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenAdminCredential)
					r.Post("/create", userHandler.CreateHandler)
					r.Put("/update", userHandler.UpdateHandler)
					r.Delete("/delete", userHandler.DeleteHandler)
				})
			})

			jobHandler := api.JobHandler{Handler: handlerType}
			r.Route("/jobs", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(mJwt.VerifyJwtTokenCredential)
					r.Get("/", jobHandler.FindAllHandler)
					r.Get("/{id}", jobHandler.FindByIDHandler)
				})
			})
		})
	})
}
