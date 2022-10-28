package todoServices

import "tarefas/db"

// Service reponsável por realizar o delete de uma entidade.
func DestroyTodoService(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close() // Fecha a conexão após o delete.

	result, err := connection.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	// Retorna o número de linhas que foram afetadas com esse delete.
	return result.RowsAffected()
}
