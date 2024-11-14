package web

import (
	"net/http"

	"github.com/Besufikad17/YATT-server/cmd/web/handlers"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()
	handlers := handlers.NewHandler(app.Ctx, app.DBConn)

	router.GET("/", handlers.Ping)

	// auth routes
	router.POST("/auth/signup", handlers.Signup)

	return router
}
