// utils/convert_to_png.go

package utils

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

func ConvertPPMtoPNG(ppmPath, pngPath string) error {
	// Open the PPM file
	ppmFile, err := os.Open(ppmPath)
	if err != nil {
		return err
	}
	defer ppmFile.Close()

	// Read the PPM header
	var width, height, maxColorValue int
	scanner := bufio.NewScanner(ppmFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			// Skip comments
			continue
		}
		if _, err := fmt.Sscanf(line, "P3 %d %d %d", &width, &height, &maxColorValue); err == nil {
			break
		}
	}

	// Check for valid image dimensions
	if width <= 0 || height <= 0 {
		return fmt.Errorf("invalid image dimensions: %d x %d", width, height)
	}

	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Read and set pixel values
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b int
			if _, err := fmt.Fscanf(ppmFile, "%d %d %d", &r, &g, &b); err != nil {
				return err
			}
			img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}

	// Create the PNG file
	pngFile, err := os.Create(pngPath)
	if err != nil {
		return err
	}
	defer pngFile.Close()

	// Encode the image as PNG and write to file
	if err := png.Encode(pngFile, img); err != nil {
		return err
	}

	return nil
}
