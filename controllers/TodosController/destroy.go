package todosController

import (
	"log"
	"net/http"
	"strconv"
	todoServices "tarefas/services/TodoServices"

	"github.com/go-chi/chi/v5"
)

// Rota de deleção.
func Destroy(w http.ResponseWriter, r *http.Request) {
	// Usaremos o CHI para capturar os parâmetros de rota, e o Atoi para convertê-lo para inteiro.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Realizamos o delete, retornando a quantidade de linhas afetadas.
	rows, err := todoServices.DestroyTodoService(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Verificamos se a quantidade de registros atualizados foi maior que um.
	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Retornamos um NoContent.
	w.WriteHeader(http.StatusNoContent)
}
