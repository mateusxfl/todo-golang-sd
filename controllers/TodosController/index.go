package todosController

import (
	"encoding/json"
	"log"
	"net/http"
	todoServices "tarefas/services/TodoServices"
)

// Rota de listagem.
func Index(w http.ResponseWriter, r *http.Request) {
	// Realizamos um read, retornando X entidades.
	todos, err := todoServices.ListTodosService()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	// Fara o encode da nossa resposta.
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
