package app
import (
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

type App struct {
	server *http.Server
}

func New() *App {
	router := httptransport.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return &App{
		server: server,
	}
}

func (a *App) Run() error {
	fmt.Println("Приложение запущено на порту 8080")
	return a.server.ListenAndServe()
}