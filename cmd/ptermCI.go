package cmd

import (
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"

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

		pterm.Fatal.PrintOnError(os.RemoveAll(getPathTo("/docs/commands")))
		pterm.Fatal.PrintOnError(os.MkdirAll(getPathTo("/docs/commands"), 0600))

		for _, doc := range markdownDocs {
			docName := strings.ReplaceAll(strings.TrimSpace(strings.TrimLeft(strings.Split(doc, "\n")[0], "# ")), " ", "_")
			pterm.Fatal.PrintOnError(ioutil.WriteFile(getPathTo(pterm.Sprintf("/docs/commands/%s.md", docName)), []byte(doc), 0600))
		}

		project := struct {
			ProjectPath string
			Name        string
			RepoName    string
			UserName    string
			URL         string
			Short       string
			Long        string
		}{}

		originURL := detectOriginURL()
		projectParts := strings.Split(strings.TrimPrefix(originURL, "https://github.com/"), "/")

		project.UserName = projectParts[0]
		project.RepoName = projectParts[1]
		project.ProjectPath = pterm.Sprintf("%s/%s", project.UserName, project.RepoName)
		project.Name = rootCmd.Name()
		project.URL = pterm.Sprintf("https://github.com/%s", project.ProjectPath)
		project.Short = rootCmd.Short
		project.Long = rootCmd.Long

		walkOverExt("./docs", ".template.md", func(path string) {
			contentBytes, err := ioutil.ReadFile(path)
			content := string(contentBytes)
			tmpl, err := template.New(filepath.Base(path)).Parse(content)
			pterm.Fatal.PrintOnError(err)
			file, err := os.OpenFile(strings.ReplaceAll(path, ".template", ""), os.O_RDWR, 0600)
			pterm.Fatal.PrintOnError(err)
			pterm.Fatal.PrintOnError(tmpl.Execute(file, project))
			pterm.Fatal.PrintOnError(file.Close())
			pterm.Fatal.PrintOnError(ioutil.WriteFile(path, []byte(content), 0600))
		})

		updateSidebar(markdownDocs)
		input, err := ioutil.ReadFile(getPathTo("/README.md"))
		pterm.Fatal.PrintOnError(err)

		err = ioutil.WriteFile(getPathTo("/docs/README.md"), input, 0600)
		pterm.Fatal.PrintOnError(err)
	},
	Hidden: true,
}

func init() {
	rootCmd.AddCommand(ptermCICmd)
}

func detectOriginURL() (url string) {
	out, err := exec.Command("git", "remote", "-v").Output()
	pterm.Fatal.PrintOnError(err)
	pterm.Debug.Printfln("Git output:\n%s", string(out))

	output := string(out)

	for _, s := range strings.Split(output, "\n") {
		s = strings.TrimSpace(strings.TrimLeft(s, "origin"))
		if strings.HasPrefix(s, "https://github.com/") && strings.Contains(s, "push") {
			pterm.Debug.Printfln("Detected GitHub Repo: %s", s)
			url = strings.TrimSpace(strings.TrimRight(s, "(push)"))

			return
		}
	}

	return
}

func walkOverExt(path, exts string, f func(path string)) {
	_ = filepath.Walk(getPathTo(path), func(path string, info fs.FileInfo, err error) error {

		for _, ext := range strings.Split(exts, ",") {
			if strings.HasSuffix(path, ext) {
				f(path)
			}
		}
		return nil
	})
}

func getPathTo(file string) string {
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "../../", file)
}

func updateSidebar(docs []string) {
	sidebarPath := getPathTo("./docs/_sidebar.md")
	sidebarContentByte, err := ioutil.ReadFile(sidebarPath)
	pterm.Fatal.PrintOnError(err)

	sidebarContent := string(sidebarContentByte)

	beforeRegex := regexp.MustCompile(`(?ms).*<!-- <<<PTERM-CI-COMMANDS-START>>> -->`)
	afterRegex := regexp.MustCompile(`(?ms)<!-- <<<PTERM-CI-COMMANDS-END>>> -->.*`)

	before := beforeRegex.FindAllString(sidebarContent, 1)[0]
	after := afterRegex.FindAllString(sidebarContent, 1)[0]

	var newSidebarContent string

	newSidebarContent += before + "\n"

	for _, doc := range docs {
		docName := strings.ReplaceAll(strings.TrimSpace(strings.TrimLeft(strings.Split(doc, "\n")[0], "# ")), " ", "_")
		newSidebarContent += "  - [" + docName + "](commands/" + docName + ".md)\n"
	}

	newSidebarContent += after

	pterm.Fatal.PrintOnError(ioutil.WriteFile(sidebarPath, []byte(newSidebarContent), 0600))
}
