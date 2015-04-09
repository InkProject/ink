package main

import (
    "html/template"
    "io/ioutil"
    "os"
)

func CompileTpl(tplPath string, name string) template.Template {
    // Read template data from file
    html, err := ioutil.ReadFile(tplPath)
    if err != nil {
        Fatal(err.Error())
    }
    // Generate html content
    tpl, err := template.New(name).Parse(string(html))
    if err != nil {
        Fatal(err.Error())
    }
    return *tpl
}

func RenderPage(tpl template.Template, tplData interface{}, outPath string) {
    defer wg.Done()
    // Create file
    outFile, err := os.Create(outPath)
    if err != nil {
        Fatal(err.Error())
    }
    defer func() {
        outFile.Close()
    }()
    // Template render
    err = tpl.Execute(outFile, tplData)
    if err != nil {
        Fatal(err.Error())
    }
}
