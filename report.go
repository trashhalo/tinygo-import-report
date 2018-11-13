package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const mainTemplate = `
package main

import (
	_ "{{.}}"
)

func main() {
}

`

const readmeTemplate = `
# Tinygo Import Report
This project imports each package in the stdlib and reports if it imports cleanly in tinygo.

| Package | Imported? |{{ range $key, $value := .}}
| {{$value.Name}} | {{if $value.Imported}} :heavy_check_mark: {{else}} [:x:](#{{$value.Name}}) {{end}} |{{ end }}


{{ range $key, $value := .}}
## {{$value.Name}}
BTBTBT
{{$value.Output}}
BTBTBT

{{ end }}
`

type Result struct {
	Name     string
	Imported bool
	Output   string
}

func main() {
	content, err := ioutil.ReadFile("imports")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	maint := template.Must(template.New("main").Parse(mainTemplate))
	readmet := template.Must(template.New("readme").Parse(strings.Replace(readmeTemplate, "BTBTBT", "```", -1)))
	var results []Result
	for _, line := range lines {
		println(line)
		noslash := strings.Replace(line, "/", "_", -1)
		dirsafe := fmt.Sprintf("tests/%v", noslash)
		os.Mkdir(dirsafe, 0755)
		f, err := os.Create(fmt.Sprintf("%v/main.go", dirsafe))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		err = maint.Execute(f, line)
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("make", "build", fmt.Sprintf("target=%v", noslash))
		stdoutStderr, err := cmd.CombinedOutput()
		results = append(results, Result{line, err == nil, string(stdoutStderr)})
	}

	f, err := os.Create("Readme.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = readmet.Execute(f, results)
	if err != nil {
		panic(err)
	}
}
