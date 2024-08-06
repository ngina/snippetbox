package main

import (
	"log"

	"snippetbox.ngina.com/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}
