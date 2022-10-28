package todosController

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tarefas/entities"
	todoServices "tarefas/services/TodoServices"

	"github.com/go-chi/chi/v5"
)

// Rota de atualização.
func Update(w http.ResponseWriter, r *http.Request) {
	// Usaremos o CHI para capturar os parâmetros de rota, e o Atoi para convertê-lo para inteiro.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Declarando variável referente a entidade, para depois fazer o Parse da request para dentro dela.
	var todo entities.Todo

	// Aqui receberemos o payload da request json, passando para a variavel da entidade.
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Realizamos o update, retornando a quantidade de linhas afetadas.
	rows, err := todoServices.UpdateTodoService(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Verificamos se a quantidade de registros atualizados foi maior que um.
	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo.ID = int64(id)

	// Fara o encode da nossa resposta.
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
