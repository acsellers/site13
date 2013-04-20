package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	ENTRY_BY_PERMALINK = iota
	ENTRIES_BY_DATE
	ENTRIES_BY_CATEGORY
	CATEGORIES
)

const (
	SQL_ENTRY_BY_PERMALINK = `
    SELECT title, hook, content, permalink, category, publish_time
    FROM entries
    WHERE permalink = $1
    LIMIT 1
  `
	SQL_ENTRIES_BY_DATE = `
    SELECT title, hook, content, permalink, category, publish_time 
    FROM entries 
    ORDER BY publish_time
    LIMIT $1
  `
	SQL_ENTRIES_BY_CATEGORY = `
    SELECT title, hook, content, permalink, category, publish_time 
    FROM entries
    WHERE category = $1
    ORDER BY publish_time
    LIMIT $2
  `
	SQL_CATEGORIES = `
    SELECT DISTINCT category
    FROM entries 
    ORDER BY category
  `
)

var (
	db         *sql.DB
	statements map[int]*sql.Stmt
)

func init() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=site")
	if err != nil {
		log.Fatal(err)
	}
	statements = make(map[int]*sql.Stmt)
	statements[ENTRY_BY_PERMALINK], err = db.Prepare(SQL_ENTRY_BY_PERMALINK)
	if err != nil {
		log.Fatal(err)
	}
	statements[ENTRIES_BY_DATE], err = db.Prepare(SQL_ENTRIES_BY_DATE)
	if err != nil {
		log.Fatal(err)
	}
	statements[CATEGORIES], err = db.Prepare(SQL_CATEGORIES)
	if err != nil {
		log.Fatal(err)
	}
	statements[ENTRIES_BY_CATEGORY], err = db.Prepare(SQL_ENTRIES_BY_CATEGORY)
	if err != nil {
		log.Fatal(err)
	}

}
