package routes

import (
	"net/http"
	"url-shortener/1.0/controllers"
	"url-shortener/1.0/database"
	"url-shortener/1.0/middlewares"
)

func CreateAppRouter(repository database.Repository) http.Handler {
	appRouter := http.NewServeMux()

	appRouter.Handle("/shortener", middlewares.DecodeJsonResponseMiddleware(
		middlewares.ValidateRequestMethodMiddleware(controllers.ShortenerUrl(repository), http.MethodPost),
	))

	appRouter.Handle("/",
		middlewares.ValidateRequestMethodMiddleware(controllers.RedirectUrl(repository), http.MethodGet),
	)

	return appRouter
}
