package main

import (
	"github.com/zoomi-raja/goweb/api"
	"github.com/zoomi-raja/goweb/config"
)

func main() {
	config.Load()
	api.Run()
}

/*var tpl *template.Template

func init() {
	cwd, _ := os.Getwd()
	tpl = template.Must(template.ParseGlob(cwd + "/src/gotut/templates/practice/*.html"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	newsPage() main function to show newss on final template
	templateFunc()
	refrence()
}*/
