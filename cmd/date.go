package cmd

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Prints the current date.",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		letters := pterm.NewLettersFromString(time.Now().Format(format))
		pterm.DefaultBigText.WithLetters(letters).Render()
	},
}

func init() {
	rootCmd.AddCommand(dateCmd)

	dateCmd.Flags().StringP("format", "f", "02 Jan 06", "specify a custom date format")
}
