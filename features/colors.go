package features

import (
	"fmt"
	"math"
	"os"
	"strings"
)

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
	c.Pixels[y][x] = color
}

func CanvasToPPM(c Canvas) string {
	// get height and width of the canvas
	height := c.Height
	width := c.Width
	// PPM header
	header := fmt.Sprintf("P3\n%d %d\n255\n", width, height)
	pixels := ""
	for i := range c.Pixels {
		line := ""
		for j := range c.Pixels[i] {
			pixel := c.Pixels[i][j]
			// convert each color value to a string
			red := fmt.Sprintf("%d", clamp(int(math.Ceil(pixel.Red*255)), 0, 255))
			green := fmt.Sprintf("%d", clamp(int(math.Ceil(pixel.Green*255)), 0, 255))
			blue := fmt.Sprintf("%d", clamp(int(math.Ceil(pixel.Blue*255)), 0, 255))
			// append the color values to the line string
			line += fmt.Sprintf("%s %s %s ", red, green, blue)
			// check if the line exceeds 70 characters
			if len(line) > 70 {
				// append the line to the pixels string and start a new line
				pixels += strings.TrimRight(line, " ") + "\n"
				line = ""
			}
		}
		// append the remaining line to the pixels string
		pixels += strings.TrimRight(line, " ") + "\n"
	}
	return header + pixels
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func PPMToFile(ppm string, filename string) {
	// write the ppm string to a .ppm file
	// create a file with the given filename
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// write the ppm string to the file
	_, err = f.WriteString(ppm)
	if err != nil {
		panic(err)
	}
	// close the file
	err = f.Close()
	if err != nil {
		panic(err)
	}

}
