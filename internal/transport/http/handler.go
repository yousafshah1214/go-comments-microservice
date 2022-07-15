package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) SetupRoutes() error {
	fmt.Println("Setting up routes")
	handler.Router = mux.NewRouter()
	handler.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
