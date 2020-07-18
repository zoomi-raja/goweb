package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	cwd, _ := os.Getwd()
	tpl, _ = template.ParseGlob(cwd + "/src/gotut/templates/practice/*.html")
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	//newsPage() main function to show newss on final template
	// templateFunc()
	//parse()
	//refrence()
}
