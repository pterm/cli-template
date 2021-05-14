package cmd

import (
	"io/ioutil"
	"strings"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// ptermCICmd represents the ptermCi command
// ! Do not delete this file. It it used inside the CI system.
var ptermCICmd = &cobra.Command{
	Use:   "ptermCI",
	Short: "Run internal CI-System to update documentation.",
	Long: `This command is used in the CI-System to generate new documentation of the CLI tool.
It should not be used outside the development of this tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Info.Printfln("Running PtermCI for %s", rootCmd.Name())
		markdownDocs := pcli.GenerateMarkdownDocs(rootCmd, rootCmd)
		pterm.Fatal.PrintOnError(ioutil.WriteFile("./doc.md", []byte(strings.Join(markdownDocs, "\n\n")), 0600))
	},
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(ptermCICmd)
}
