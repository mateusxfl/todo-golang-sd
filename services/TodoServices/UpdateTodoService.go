package todoServices

import (
	"tarefas/db"
	"tarefas/entities"
)

// Service reponsável por realizar o update de uma entidade.
func UpdateTodoService(id int64, todo entities.Todo) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close() // Fecha a conexão após o update.

	result, err := connection.Exec(`UPDATE todos SET title = $2, description = $3, done = $4 WHERE id = $1`, id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	// Retorna o número de linhas que foram afetadas com esse update.
	return result.RowsAffected()
}
