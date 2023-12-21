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
