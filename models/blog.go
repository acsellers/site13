package models

import "time"

type BlogEntry struct {
	Title       string
	Hook        string
	Content     string
	ID          string
	PublishTime time.Time
}

type Category struct {
	Name string
	ID   string
}

func MustFindPage(id string) BlogEntry {
	return BlogEntry{Title: "Test Entry",
		Hook:        "Blah blah, hook here",
		Content:     "Lost more content here, blah blah Bioshock",
		ID:          "test_blog_entry",
		PublishTime: time.Now()}
}
func FindPage(id string) (BlogEntry, bool) {
	testEntry := BlogEntry{Title: "Test Entry",
		Hook:        "Blah blah, hook here",
		Content:     "Lost more content here, blah blah Bioshock",
		ID:          "test_blog_entry",
		PublishTime: time.Now()}

	return testEntry, true
}

func NewEntries(num int) []BlogEntry {
	// TODO: make this do something
	return []BlogEntry{MustFindPage(""), MustFindPage(""), MustFindPage("")}
}

func TopEntries(num int) []BlogEntry {
	// TODO: make this do something
	return []BlogEntry{MustFindPage(""), MustFindPage(""), MustFindPage("")}
}

func ActiveCategories() []Category {
	return []Category{Category{Name: "Golang", ID: "golang"}, Category{Name: "Programming", ID: "programming"}}
}
