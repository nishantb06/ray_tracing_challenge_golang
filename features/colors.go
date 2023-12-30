package features

type Color struct {
	Red, Green, Blue float64
}

// constructor for Color, Any name could have been chosen for this function
// but the convention is to use the struct name as the function name followed
// by an underscore
func Color_(red, green, blue float64) Color {
	return Color{Red: red, Green: green, Blue: blue}
}

func (c Color) Add(c2 Color) Color {
	return Color{c.Red + c2.Red, c.Green + c2.Green, c.Blue + c2.Blue}
}

func AddColor(c1, c2 Color) Color {
	return Color{c1.Red + c2.Red, c1.Green + c2.Green, c1.Blue + c2.Blue}
}

func (c Color) Subtract(c2 Color) Color {
	return Color{c.Red - c2.Red, c.Green - c2.Green, c.Blue - c2.Blue}
}

func SubtractColor(c1, c2 Color) Color {
	return Color{c1.Red - c2.Red, c1.Green - c2.Green, c1.Blue - c2.Blue}
}

func (c Color) Multiply(scalar float64) Color {
	return Color{c.Red * scalar, c.Green * scalar, c.Blue * scalar}
}

func MultiplyColor(c Color, scalar float64) Color {
	return Color{c.Red * scalar, c.Green * scalar, c.Blue * scalar}
}

func (c Color) HadamardProduct(c2 Color) Color {
	return Color{c.Red * c2.Red, c.Green * c2.Green, c.Blue * c2.Blue}
}

func HadamardProductColor(c1, c2 Color) Color {
	return Color{c1.Red * c2.Red, c1.Green * c2.Green, c1.Blue * c2.Blue}
}

type Canvas struct {
	Width, Height int
	Pixels        [][]Color
}

// constructor for Canvas struct
func Canvas_(width, height int) Canvas {
	pixels := make([][]Color, height) // lenth of the inner slice can vary
	for i := range pixels {
		pixels[i] = make([]Color, width) // allocate each inner slice
	}
	// initialize the canvas with black pixels
	for i := range pixels {
		for j := range pixels[i] {
			pixels[i][j] = Color_(0, 0, 0)
		}
	}
	return Canvas{Width: width, Height: height, Pixels: pixels}
}

func (c Canvas) WritePixel(x, y int, color Color) {
	// check if the pixel is within the canvas, raise error if not
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		panic("pixel out of bounds")
	}
	c.Pixels[x][y] = color
}

