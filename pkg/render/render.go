package render

import (
	"bookings/pkg/config"
	"bookings/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//renderTemplate irá fazer a criação da página a partir do Template.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	//Criar vários templates - chamando a função abaixo.
	tc := app.TemplateCache

	//Pegar um único template que está dentro do TC
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Não foi possível")
	}

	//Renderizar o template.
	buf := new(bytes.Buffer)

	//Chama a função para adicionar qualquer conteúdo padrão no Td
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal("Não foi possível")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//Pegar todos os arquivos terminados em pages.tmpl do diretório ./templates
	//Pega todo o endereço dos arquivos terminados em page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//Aqui preciso lembrar exatamente como funciona a função loop com range.
	// for index, arquivo - não me importo com o index.
	for _, page := range pages {

		//Aqui não queremos trabalhar com o arquivo inteiro, queremos apenas a parte final do arquivo.
		name := filepath.Base(page) // vai pegar apenas o nome final do arquivo page.tmpl

		//Aqui vou fazer a transferência do arquivo para o template Set.
		//Crio um novo template com o nome pego na página e passo o endereço pelo page.
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}

		//Aqui eu acrescento elemento no Mapa.
		//Pego o nome do elemento e acrescento o valor que é o tamplate set
		myCache[name] = ts

	}

	return myCache, nil

}
