package main

import (
	"flag"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
)

func main() {
	flag.BoolVar(&pterm.PrintDebugMessages, "debug", false, "show debug output")
	flag.Parse()

	pterm.DefaultHeader.Println("PTerm CLI Template Setup")
	pterm.DefaultParagraph.Println("This setup will ask you a couple questions about your CLI tool and get the template ready to start!")

	originURL := detectOriginURL()

	projectParts := strings.Split(strings.TrimPrefix(originURL, "https://github.com/"), "/")
	repoUser := projectParts[0]
	repoName := projectParts[1]

	var qs = []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message: "GitHub host username",
				Default: repoUser,
			},
			Validate: survey.Required,
		},
		{
			Name: "reponame",
			Prompt: &survey.Input{
				Message: "GitHub repository name",
				Default: repoName,
			},
			Validate: survey.Required,
		},
	}

	project := struct {
		Username    string
		Reponame    string
		ProjectURL  string
		ProjectName string
	}{}

	pterm.Fatal.PrintOnError(survey.Ask(qs, &project))
	project.ProjectURL = pterm.Sprintf("github.com/%s/%s", project.Username, project.Reponame)
	project.ProjectName = pterm.Sprintf("%s/%s", project.Username, project.Reponame)

	const cliTemplateURL = "pterm/cli-template"

	pterm.Info.Printfln("Replacing all '%s' with %s", pterm.Magenta(cliTemplateURL), pterm.Magenta(project.ProjectName))
	walkOverExt("go,mod", func(path string) {
		pterm.Debug.Printfln("Replacing '%s' in %s with %s", pterm.Magenta(cliTemplateURL), path, pterm.Magenta(project.ProjectName))
		replaceAllInFile(path, cliTemplateURL, project.ProjectName)
	})
}

func walkOverExt(exts string, f func(path string)) {
	_ = filepath.Walk(getPathTo(""), func(path string, info fs.FileInfo, err error) error {

		for _, ext := range strings.Split(exts, ",") {
			if filepath.Ext(path) == "."+ext {
				f(path)
			}
		}
		return nil
	})
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

func replaceAllInFile(filepath, search, replace string) {
	fileBytes, err := ioutil.ReadFile(filepath)
	pterm.Warning.PrintOnError(err)
	content := string(fileBytes)

	content = strings.ReplaceAll(content, search, replace)

	pterm.Warning.PrintOnError(ioutil.WriteFile(filepath, []byte(content), 0600))
}

func getPathTo(file string) string {
	// NOTE: Replace the 1 with a 0 if you use this code directly, instead of wrapping it in a function.
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "../../", file)
}
