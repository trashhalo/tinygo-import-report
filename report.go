package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
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
A package with a check may work a package with a x will definately not currently work.
If a package has a x in imported? but check in importedWithInit? it means you need to build your project
with -initinterp.

| Package | Imported? | ImportedWithInit? |
| --- | --- | --- |{{ range $key, $value := .}}
| {{$value.Name}} | {{if $value.Imported}} :heavy_check_mark: {{else}} [:x:](#{{$value.Link}}) {{end}} | {{if $value.ImportedWithInit}} :heavy_check_mark: {{else}} [:x:](#{{$value.Link}}) {{end}} | {{ end }}


{{ range $key, $value := .}}
## {{$value.Name}}
Without Init Interp

BTBTBT
{{$value.Output}}
BTBTBT

With Init Interp

BTBTBT
{{$value.OutputWithInit}}
BTBTBT
{{ end }}
`

type Result struct {
	Name             string
	Imported         bool
	ImportedWithInit bool
	Output           string
	OutputWithInit   string
	Link             string
}

func main() {
	content, err := ioutil.ReadFile("imports")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	maint := template.Must(template.New("main").Parse(mainTemplate))
	readmet := template.Must(template.New("readme").Parse(strings.Replace(readmeTemplate, "BTBTBT", "```", -1)))
	results := make(chan Result)
	var wg sync.WaitGroup
	var resultsArr []Result
	go func() {
		for r := range results {
			resultsArr = append(resultsArr, r)
			println(r.Name)
			wg.Done()
		}
	}()
	for _, line := range lines {
		wg.Add(1)
		go func(line string) {
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
			cmdI := exec.Command("make", "build", fmt.Sprintf("target=%v", noslash))
			stdoutStderrI, errI := cmdI.CombinedOutput()
			results <- Result{
				line,
				err == nil,
				errI == nil,
				string(stdoutStderr),
				string(stdoutStderrI),
				strings.Replace(line, "/", "", -1),
			}
		}(line)
	}
	wg.Wait()
	close(results)
	sort.Slice(resultsArr, func(i, j int) bool {
		return resultsArr[i].Name < resultsArr[j].Name
	})

	f, err := os.Create("Readme.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = readmet.Execute(f, resultsArr)
	if err != nil {
		panic(err)
	}
}
