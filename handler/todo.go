package handler

import (
	"net/http"

	"github.com/C3LOUD/the-heart-stack/components/page"
	"github.com/C3LOUD/the-heart-stack/services/todo"
	"github.com/C3LOUD/the-heart-stack/util"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct{}

func (h TodoHandler) HandleListTodos(c echo.Context) error {
	todos, err := todo.GetTodos(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return util.Render(c, page.TodoPage(todos))
}
