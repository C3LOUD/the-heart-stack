package handler

import (
	"github.com/C3LOUD/the-heart-stack/view/index"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func (h IndexHandler) HandleIndex(c echo.Context) error {
	return render(c, index.Index())
}
