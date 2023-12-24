package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/C3LOUD/the-heart-stack/components/page"
	t "github.com/C3LOUD/the-heart-stack/components/todo"
	"github.com/C3LOUD/the-heart-stack/components/ui"
	"github.com/C3LOUD/the-heart-stack/db"
	"github.com/C3LOUD/the-heart-stack/services/todo"
	"github.com/C3LOUD/the-heart-stack/util"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct{}

type ResponseMessage struct {
	Message string `json:"message" xml:"message"`
}

func (h TodoHandler) HandleListTodos(c echo.Context) error {
	todos, err := todo.GetTodos(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return util.Render(c, t.TodoList(todos))
}

func (h TodoHandler) HandleTodoPage(c echo.Context) error {
	return util.Render(c, page.TodoPage())
}

func (h TodoHandler) HandleCreateTodo(c echo.Context) error {
	_, err := todo.CreateTodo(c, c.Request().FormValue("todo"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Response().Header().Set("HX-Trigger", "rsync")
	return c.JSON(http.StatusOK, ResponseMessage{
		Message: "Todo Updated",
	},
	)
}

func (h TodoHandler) HandleConfirmDeleteTodo(c echo.Context) error {
	return util.Render(c, ui.ExpandedButton(c.Param("id")))
}

func (h TodoHandler) HandleDeleteTodo(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := todo.DeleteTodo(c, id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Response().Header().Set("HX-Trigger", "rsync")
	return c.JSON(http.StatusOK, ResponseMessage{
		Message: "Todo Updated",
	},
	)
}

func (h TodoHandler) HandleFinishTodo(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	args := db.ToggleTodoParams{
		IsFinished: sql.NullBool{Bool: true, Valid: true},
		ID:         id,
	}

	err = db.UseQuery().ToggleTodo(c.Request().Context(), args)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Response().Header().Set("HX-Trigger", "rsync")
	return c.JSON(http.StatusOK, ResponseMessage{
		Message: "Todo Updated",
	},
	)
}
