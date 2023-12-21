package handler

import (
	"github.com/C3LOUD/the-heart-stack/components/page"
	"github.com/C3LOUD/the-heart-stack/util"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

func (h IndexHandler) HandleIndex(c echo.Context) error {
	return util.Render(c, page.Index())
}
