package cmd

import (
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Prints the current time",
	Long:  `You can print a live clock with the '--live' flag!`,
	Run: func(cmd *cobra.Command, args []string) {
		live, _ := cmd.Flags().GetBool("live")

		area, _ := pterm.DefaultArea.WithRemoveWhenDone().Start()
		if live {
			for {
				letters := pterm.NewLettersFromString(time.Now().Format("15:01:05"))
				txt, _ := pterm.DefaultBigText.WithLetters(letters).Srender()
				area.Update(txt)
				time.Sleep(time.Second)
			}
		} else {
			letters := pterm.NewLettersFromString(time.Now().Format("15:01:05"))
			pterm.DefaultBigText.WithLetters(letters).Render()
		}

		area.Stop()
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	timeCmd.Flags().BoolP("live", "l", false, "live output")
}
