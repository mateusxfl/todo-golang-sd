package todosController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	todoServices "tarefas/services/TodoServices"

	"github.com/go-chi/chi/v5"
)

// Rota de listagem individual.
func Show(w http.ResponseWriter, r *http.Request) {
	// Usaremos o CHI para capturar os parâmetros de rota, e o Atoi para convertê-lo para inteiro.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Realizamos um read, retornando uma única entidade.
	todo, err := todoServices.ListSingleTodoService(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Fara o encode da nossa resposta.
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
