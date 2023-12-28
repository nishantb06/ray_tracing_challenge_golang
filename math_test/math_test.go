package math

import (
	"fmt"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	fmt.Println("Hello, World!")
	if math.Abs(-1) != 1 {
		t.Error("Abs(-1) should be 1")
	}
}
