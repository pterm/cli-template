package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-template",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
	// ! This template features automatic releases.
	// ! When you have set a REPO_ACCESS_TOKEN secret in GitHub, increasing this version will push a new release automatically.
	Version: "v0.0.1", // <---VERSION---> This comment enables auto-releases on version change!
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// 	Run: func(cmd *cobra.Command, args []string) {   },
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
	// Adds global flags for PTerm settings.
	// Fill the empry strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRootCmd(rootCmd)
	rootCmd.AddCommand(pcli.GetCiCommand())
	rootCmd.SetFlagErrorFunc(pcli.FlagErrorFunc())
	rootCmd.SetGlobalNormalizationFunc(pcli.GlobalNormalizationFunc())
	rootCmd.SetHelpFunc(pcli.HelpFunc())
	rootCmd.SetHelpTemplate(pcli.HelpTemplate())
	rootCmd.SetUsageFunc(pcli.UsageFunc())
	rootCmd.SetUsageTemplate(pcli.UsageTemplate())
	rootCmd.SetVersionTemplate(pcli.VersionTemplate())

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
