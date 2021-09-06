package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "protector",
		Short: "A CLI tool for protecting the image",
		Long:  `This tool provides an easy and extensible way to protect the image.`,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}
