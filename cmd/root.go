package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pcli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "cli-template",
	Short:   "A brief description of your application",
	Long:    `A longer description that spans multiple lines and likely contains\nexamples and usage of using your application.`,
	Version: "v0.0.1", // <---VERSION---> This comment enables auto-releases on version change!
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Use https://github.com/pterm/pcli style the output of cobra.
	rootCmd.SetVersionTemplate(pcli.GenerateVersionString(rootCmd.Name(), rootCmd.Version))
	rootCmd.SetHelpFunc(pcli.Spf13CobraHelpFunc(rootCmd))
	rootCmd.SetFlagErrorFunc(pcli.Spf13CobraFlagErrorFunc(rootCmd))

	// Change global PTerm theme
	// pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
