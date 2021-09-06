package image

import (
	"fmt"
	"image"
	"math/rand"
	"protector/pkg"
)

type Pixels struct {
	OriginPixels  []*Pixel
	ShufflePixels []*Pixel
	Bounds        image.Rectangle
}

// GetPixels - returns origin pixels and shuffle pixels
func GetPixels(img image.Image) (*Pixels, error) {
	var pixels []*Pixel
	var hidePixels []*Pixel
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			pixels = append(pixels, NewPixel(x, y).WithColor(color))
			hidePixels = append(hidePixels, NewPixel(x, y).WithColor(color))
		}
	}

	// input the seed
	fmt.Printf("\n\nPlease input the seed: ")
	var seed string
	fmt.Scanf("%s", &seed)
	if seed == "" {
		return nil, fmt.Errorf("invalid seed: %s", seed)
	}

	// shuffle
	s, err := pkg.GenerateSeed(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate seed: %v", err)
	}
	rand.Seed(s)
	rand.Shuffle(len(hidePixels), func(i, j int) {
		hidePixels[i], hidePixels[j] = hidePixels[j], hidePixels[i]
	})
	return &Pixels{
		OriginPixels:  pixels,
		ShufflePixels: hidePixels,
		Bounds:        bounds,
	}, nil
}
