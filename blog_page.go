package main

import (
	"github.com/acsellers/site13/models"
	"net/http"
)

func RenderBlogPage(bp models.BlogEntry, w http.ResponseWriter) {
	entryTemplate.Execute(w, struct {
		Entry      models.BlogEntry
		NewEntries []models.BlogEntry
		Categories []models.Category
	}{
		bp,
		models.NewEntries(3),
		models.ActiveCategories(),
	})
}
