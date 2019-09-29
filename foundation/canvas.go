package foundation

import "image"
import "image/png"
import "image/color"
import "os"


type Canvas struct {
	Width int
	Height int
	Pixels []Color
}

func NewCanvas(width, height int) Canvas {
	pixels := make([]Color, width * height)
	return Canvas { width, height, pixels }
}

func (c *Canvas) Write(x, y int, clr Color) {
	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		panic("Attempt to write outside canvas")
	}

	c.Pixels[y * c.Width + x] = clr
}

func (c *Canvas) Read(x, y int) Color {
	if x < 0 || y < 0 || x >= c.Width || y >= c.Height {
		panic("Attempt to read outside canvas")
	}

	return c.Pixels[y * c.Width + x]
}

func clamp(f float64) uint8 {
	r := int(0.5 + 255*f)
	if r < 0 {
		r = 0
	}
	if r > 255 {
		r = 255
	}

	return uint8(r)
}

func convert(c Color) color.Color {
	r := clamp(c.Red)
	g := clamp(c.Green)
	b := clamp(c.Blue)

	return color.RGBA {r, g, b, 255 }
}

func (c *Canvas) Save(filename string) error {
	rgba := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for i := 0; i<c.Width; i++ {
		for j := 0; j<c.Height; j++ {
			clr := convert(c.Read(i, j))
			rgba.Set(i, c.Height - j, clr)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, rgba)
	if err != nil {
		return err
	}

	return nil
}
