package main

import (
	"html/template"
	"log"

	"snippetbox.ngina.com/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}
