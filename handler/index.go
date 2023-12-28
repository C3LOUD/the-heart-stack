package handler

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/C3LOUD/the-heart-stack/components/page"
	"github.com/C3LOUD/the-heart-stack/util"
	"github.com/labstack/echo/v4"
)

type IndexHandler struct{}

type Color struct {
	R int
	G int
	B int
}

func newColor() Color {
	return Color{
		R: rand.Intn(256),
		G: rand.Intn(256),
		B: rand.Intn(256),
	}
}

func (c Color) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func (c Color) HSL() string {
	R := float64(c.R) / 255
	G := float64(c.G) / 255
	B := float64(c.B) / 255
	max := max(R, G, B)
	min := min(R, G, B)
	chroma := max - min

	var sat float64
	var hue float64
	lum := (max + min) / 2
	if chroma == 0 {
		hue = 0
		sat = 0
	} else {
		sat = chroma / (1 - math.Abs(2*lum-1))
		switch max {
		case R:
			segment := (G - B) / chroma
			if segment < 0 {
				hue = segment + 360/60
			} else {
				hue = segment
			}
		case G:
			hue = (B-R)/chroma + 120/60
		case B:
			hue = (R-G)/chroma + 240/60
		}
	}

	return fmt.Sprintf("hsl(%v, %v%%, %v%%)", math.Round(hue*60), math.Round(sat*100), math.Round(lum*100))
}

func (h IndexHandler) HandleIndex(c echo.Context) error {
	color := newColor()
	hex := color.Hex()
	fmt.Println(hex)
	hsl := color.HSL()
	fmt.Println(hsl)
	return util.Render(c, page.Index(hex))
}
