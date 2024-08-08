package main

import (
	"html/template"
	"log"

	"github.com/go-playground/form/v4"
	"snippetbox.ngina.com/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}
