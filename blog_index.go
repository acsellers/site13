package main

import (
	"github.com/acsellers/site13/models"
	"net/http"
)

func RenderBlogIndex(w http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(w, struct {
		Entries []models.BlogEntry
		Title   string
	}{
		models.NewEntries(12),
		"",
	})
}

func RenderCategory(category string, w http.ResponseWriter) {
	indexTemplate.Execute(w, struct {
		Entries []models.BlogEntry
		Title   string
	}{
		models.CategoryEntries(category, 25),
		category,
	})
}
