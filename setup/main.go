package main

import (
	"flag"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
)

func main() {
	flag.BoolVar(&pterm.PrintDebugMessages, "debug", false, "show debug output")
	flag.Parse()

	pterm.DefaultHeader.Println("PTerm CLI Template Setup")

	originURL := detectOriginURL()

	projectParts := strings.Split(strings.TrimPrefix(originURL, "https://github.com/"), "/")
	repoUser := projectParts[0]
	repoName := projectParts[1]

	project := struct {
		Username    string
		Reponame    string
		ProjectURL  string
		ProjectName string
	}{}

	project.Username = repoUser
	project.Reponame = repoName
	project.ProjectURL = pterm.Sprintf("github.com/%s/%s", project.Username, project.Reponame)
	project.ProjectName = pterm.Sprintf("%s/%s", project.Username, project.Reponame)

	const cliTemplatePath = "pterm/cli-template"
	pterm.Info.Printfln("Replacing all '%s' with %s", pterm.Magenta(cliTemplatePath), pterm.Magenta(project.ProjectName))
	walkOverExt("go,mod", func(path string) {
		pterm.Debug.Printfln("Replacing '%s' in %s with %s", pterm.Magenta(cliTemplatePath), path, pterm.Magenta(project.ProjectName))
		replaceAllInFile(path, cliTemplatePath, project.ProjectName)
	})

	replaceAllInFile("./cmd/root.go", `Use:   "cli-template",`, pterm.Sprintf(`Use:   "%s",`, project.Reponame))

	pterm.Fatal.PrintOnError(os.Remove(getPathTo("./README.md")))
	pterm.Fatal.PrintOnError(os.Rename(getPathTo("./README.template-setup.md"), getPathTo("./README.template.md")))
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
	_, scriptPath, _, _ := runtime.Caller(1)
	return filepath.Join(scriptPath, "../../", file)
}
