package cmd

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	pkgImage "protector/pkg/lib/image"

	"github.com/spf13/cobra"
)

var (
	repairCommand = &cobra.Command{
		Use:   "repair",
		Short: "Repair the damaged image",
		Long:  "Repair the damaged image",
		RunE:  runRepairCommand,
	}
	repairFile    string
	newRepairFile string
)

func init() {
	rootCmd.AddCommand(repairCommand)
	repairCommand.Flags().StringVarP(&repairFile, "file", "f", "", "specify a file")
	repairCommand.Flags().StringVarP(&newRepairFile, "output", "o", "", "export the image")
}

func runRepairCommand(cmd *cobra.Command, args []string) error {

	// opeb file
	f, err := os.Open(repairFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("failed to decode image: %v", err)
	}

	// convert pixels
	p, err := pkgImage.GetPixels(img)
	if err != nil {
		return fmt.Errorf("failed to get pixels: %v", err)
	}
	pixels := p.OriginPixels
	hidePixels := p.ShufflePixels

	// repair color
	for i, hidePixel := range hidePixels {
		repairPixel := pixels[i]
		if !hidePixel.Color.Check(repairPixel.Color) {
			var valid bool
			for j, h := range hidePixels {
				authPixel := pixels[j]
				if h.X == repairPixel.X && h.Y == repairPixel.Y && h.Color.Check(authPixel.Color) {
					valid = true
					break
				}
			}
			if valid {
				hidePixel.Color.Repair(repairPixel.Color)
			}
		}
	}

	// generate a new image
	newImg := image.NewRGBA(p.Bounds)
	for _, pixel := range hidePixels {
		r, g, b, a := pixel.Color.RGBA()
		newImg.Set(pixel.X, pixel.Y, color.RGBA{
			R: r,
			G: g,
			B: b,
			A: a,
		})
	}
	outFile, err := os.Create(newRepairFile)
	if err != nil {
		return fmt.Errorf("failed to create image: %v", err)
	}
	defer outFile.Close()
	if err := png.Encode(outFile, newImg); err != nil {
		return fmt.Errorf("failed to encode png file: %v", err)
	}

	return nil
}
