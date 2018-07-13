package index

import (
	"net/http"
	"text/template"
	"handler"
	"log"
)

type IndexHandler struct{}

const INDEX = "index"

func (i *IndexHandler) Index(w http.ResponseWriter, req *http.Request) {

	t, err := template.ParseFiles(handler.GetView("index.html"))
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Execute(w, nil)

}
