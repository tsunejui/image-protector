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
	hideCommand = &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt the repairing information on photo",
		Long:  "Hide the repairing information in each pixel on specified image",
		RunE:  runHideCommand,
	}
	hideFile    string
	newHideFile string
)

func init() {
	rootCmd.AddCommand(hideCommand)
	hideCommand.Flags().StringVarP(&hideFile, "file", "f", "", "specify a file")
	hideCommand.Flags().StringVarP(&newHideFile, "output", "o", "", "export the image")
}

func runHideCommand(cmd *cobra.Command, args []string) error {

	// ope file
	f, err := os.Open(hideFile)
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

	// hide the repair data
	for i, x := range hidePixels {
		pixels[i].Color.Hide(x.Color)
	}

	// generate a new image
	newImg := image.NewRGBA(p.Bounds)
	for _, pixel := range pixels {
		r, g, b, a := pixel.Color.RGBA()
		newImg.Set(pixel.X, pixel.Y, color.RGBA{
			R: r,
			G: g,
			B: b,
			A: a,
		})
	}

	outFile, err := os.Create(newHideFile)
	if err != nil {
		return fmt.Errorf("failed to create image: %v", err)
	}
	defer outFile.Close()
	if err := png.Encode(outFile, newImg); err != nil {
		return fmt.Errorf("failed to encode png file: %v", err)
	}
	return nil
}
