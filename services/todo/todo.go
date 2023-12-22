package todo

import (
	"github.com/C3LOUD/the-heart-stack/db"
	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) ([]db.Todo, error) {
	queries := db.UseQuery()

	todos, err := queries.GetTodos(c.Request().Context())
	if err != nil {
		return nil, err
	}

	return todos, err
}

func CreateTodo(c echo.Context, content string) (db.Todo, error) {
	todo, err := db.UseQuery().CreateTodo(c.Request().Context(), content)
	if err != nil {
		return db.Todo{}, err
	}
	return todo, nil
}

func DeleteTodo(c echo.Context, id int64) error {
	if err := db.UseQuery().DeleteTodo(c.Request().Context(), id); err != nil {
		return err
	}
	return nil
}
