package main

import (
	"html/template"
	"log"
	"os"
)

type TableFieldFilter struct {
	FieldName  string
	FieldValue string
	Operator   string
}

type TableQuery struct {
	Name    string
	Fields  []string
	Filters []*TableFieldFilter
}

func tmplFile() {
	t, err := template.ParseFiles("./template-demo.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	qStr := &TableQuery{
		Name:   "activity",
		Fields: []string{"category", "style"},
		Filters: []*TableFieldFilter{
			{
				FieldName:  "category",
				Operator:   ">=",
				FieldValue: "edu",
			},
			{
				FieldName:  "style",
				Operator:   "<",
				FieldValue: "inside",
			},
		},
	}

	err = t.Execute(os.Stdout, qStr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmplFile()
}
