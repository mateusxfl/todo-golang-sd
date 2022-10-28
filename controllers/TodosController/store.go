package todosController

import (
	"encoding/json"
	"log"
	"net/http"
	"tarefas/entities"
	todoServices "tarefas/services/TodoServices"
)

// Rota de cadastro.
func Store(w http.ResponseWriter, r *http.Request) {
	// Declarando variável referente a entidade, para depois fazer o Parse da request para dentro dela.
	var todo entities.Todo

	// Aqui receberemos o payload da request json, passando para a variavel da entidade.
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Realizamos o insert, retornando um ID.
	id, err := todoServices.StoreTodoService(todo)
	if err != nil {
		log.Printf("Ocorreu um erro ao tentar inserir: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo.ID = int64(id)

	// Fara o encode da nossa resposta.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
