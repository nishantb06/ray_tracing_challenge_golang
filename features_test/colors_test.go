package features

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nishantb06/ray_tracing_challenge_golang/features"
)

// const epsilon = 0.00001

func TestColor(t *testing.T) {
	c := features.Color_(1, 2, 3)
	if c.Red != 1 {
		t.Error("Red should be 1")
	}
	if c.Green != 2 {
		t.Error("Green should be 2")
	}
	if c.Blue != 3 {
		t.Error("Blue should be 3")
	}
}

func TestAddColor(t *testing.T) {
	c1 := features.Color_(1, 2, 3)
	c2 := features.Color_(2, 3, 4)
	c3 := features.AddColor(c1, c2)
	c4 := c1.Add(c2)
	if c3.Red != 3 {
		t.Error("Red should be 3")
	}
	if c3.Green != 5 {
		t.Error("Green should be 5")
	}
	if c3.Blue != 7 {
		t.Error("Blue should be 7")
	}
	if c4.Red != 3 {
		t.Error("Red should be 3")
	}
	if c4.Green != 5 {
		t.Error("Green should be 5")
	}
	if c4.Blue != 7 {
		t.Error("Blue should be 7")
	}
}

func TestSubtractColor(t *testing.T) {
	c1 := features.Color_(1, 2, 3)
	c2 := features.Color_(2, 3, 4)
	c3 := features.SubtractColor(c1, c2)
	c4 := c1.Subtract(c2)
	if c3.Red != -1 {
		t.Error("Red should be -1")
	}
	if c3.Green != -1 {
		t.Error("Green should be -1")
	}
	if c3.Blue != -1 {
		t.Error("Blue should be -1")
	}
	if c4.Red != -1 {
		t.Error("Red should be -1")
	}
	if c4.Green != -1 {
		t.Error("Green should be -1")
	}
	if c4.Blue != -1 {
		t.Error("Blue should be -1")
	}
}

func TestMultiplyColor(t *testing.T) {
	c1 := features.Color_(1, 2, 3)
	c3 := features.MultiplyColor(c1, 2)
	c4 := c1.Multiply(2)
	if c3.Red != 2 {
		t.Error("Red should be 2")
	}
	if c3.Green != 4 {
		t.Error("Green should be 4")
	}
	if c3.Blue != 6 {
		t.Error("Blue should be 6")
	}
	if c4.Red != 2 {
		t.Error("Red should be 2")
	}
	if c4.Green != 4 {
		t.Error("Green should be 4")
	}
	if c4.Blue != 6 {
		t.Error("Blue should be 6")
	}
}

func TestHadamardProductColor(t *testing.T) {
	c1 := features.Color_(1, 2, 3)
	c2 := features.Color_(2, 3, 4)
	c3 := features.HadamardProductColor(c1, c2)
	c4 := c1.HadamardProduct(c2)
	if c3.Red != 2 {
		t.Error("Red should be 2")
	}
	if c3.Green != 6 {
		t.Error("Green should be 6")
	}
	if c3.Blue != 12 {
		t.Error("Blue should be 12")
	}
	if c4.Red != 2 {
		t.Error("Red should be 2")
	}
	if c4.Green != 6 {
		t.Error("Green should be 6")
	}
	if c4.Blue != 12 {
		t.Error("Blue should be 12")
	}
}

func TestCanvas(t *testing.T) {
	c := features.Canvas_(10, 20)
	if c.Width != 10 {
		t.Error("Width should be 10")
	}
	if c.Height != 20 {
		t.Error("Height should be 20")
	}
	for i := range c.Pixels {
		for j := range c.Pixels[i] {
			if c.Pixels[i][j] != features.Color_(0, 0, 0) {
				t.Error("Color should be 0,0,0")
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := features.Canvas_(10, 20)
	red := features.Color_(1, 0, 0)
	c.WritePixel(2, 3, red)
	pixel := c.Pixels[3][2]
	if pixel.Red != 1 {
		fmt.Println(pixel.Red)
		t.Error("Red should be 1")
	}
	if pixel.Green != 0 {
		t.Error("Green should be 0")
	}
	if pixel.Blue != 0 {
		t.Error("Blue should be 0")
	}
}

func TestCanvasToPPM(t *testing.T) {
	c := features.Canvas_(5, 3)
	ppm := features.CanvasToPPM(c)
	// first 3 lines of ppm are​
	lines := strings.Split(ppm, "\n")
	fmt.Println(len(lines))
	if lines[0] != "P3" {
		t.Error("first line of ppm should be P3")
	}
	if lines[1] != "5 3" {
		t.Error("second line of ppm should be 5 3")
	}
	if lines[2] != "255" {
		t.Error("third line of ppm should be 255")
	}
	c.WritePixel(0, 0, features.Color_(1.5, 0, 0))
	c.WritePixel(2, 1, features.Color_(0, 0.5, 0))
	c.WritePixel(4, 2, features.Color_(-0.5, 0, 1))

	// checking the fourth, fifth and sixth line of ppm
	ppm = features.CanvasToPPM(c)
	fmt.Println(ppm)
	lines = strings.Split(ppm, "\n")
	fmt.Println(len(lines))
	if lines[3] != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0" {
		t.Error("fourth line of ppm should be 255 0 0 0 0 0 0 0 0 0 0 0 0 0 0")
	}
	if lines[4] != "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0" {
		fmt.Println(lines[4])
		t.Error("fifth line of ppm should be 0 0 0 0 0 0 0 128 0 0 0 0 0 0 0")
	}
	if lines[5] != "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255" {
		t.Error("sixth line of ppm should be 0 0 0 0 0 0 0 0 0 0 0 0 0 0 255")
	}
}
