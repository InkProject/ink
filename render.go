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
    // Produce html content
    // .Funcs(funcMap)
    tpl, err := template.New(name).Parse(string(html))
    if err != nil {
        Fatal(err.Error())
    }
    return *tpl
}

func RenderPage(tpl template.Template, tplData interface{}, outPath string) {
    // Create file
    outFile, err := os.Create(outPath)
    if err != nil {
        Fatal(err.Error())
    }
    defer func() {
        outFile.Close()
    }()
    // Define data preprocess
    // funcMap := template.FuncMap{
    //     "content": func(content []byte) string {
    //         return string(content)
    //     },
    //     "date": func(date string) string {
    //         return parseDate(date).Format(STD_FORMAT)
    //     },
    // }
    // Render
    err = tpl.Execute(outFile, tplData)
    if err != nil {
        Fatal(err.Error())
    }
}
