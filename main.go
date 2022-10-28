package main

import (
	"fmt"
	"net/http"
	"tarefas/configs"
	todosController "tarefas/controllers/TodosController"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Tentamos carregar nossas configurações, caso não consigamos, encerraremos a aplicação.
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	// Iremos declarar nossas rotas.
	router := chi.NewRouter()

	router.Get("/", todosController.Index)

	router.Get("/{id}", todosController.Show)

	router.Post("/", todosController.Store)

	router.Put("/{id}", todosController.Update)

	router.Delete("/{id}", todosController.Destroy)

	// Subindo nosso servidor, de acordo com a porta da config.
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
