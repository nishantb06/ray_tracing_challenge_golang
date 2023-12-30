package features

import (
	"fmt"
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
	pixel := c.Pixels[2][3]
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

