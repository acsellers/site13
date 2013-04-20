package models

import "time"

type BlogEntry struct {
	Title       string
	Hook        string
	Content     string
	ID          string
	Category    string
	PublishTime time.Time
}

type Category struct {
	Name string
}

func FindPage(id string) (BlogEntry, bool) {
	var be BlogEntry
	row := statements[ENTRY_BY_PERMALINK].QueryRow(id)
	err := row.Scan(&be.Title, &be.Hook, &be.Content, &be.ID, &be.Category, &be.PublishTime)
	if err != nil {
		return be, false
	}
	return be, true
}
