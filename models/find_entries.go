package models

func NewEntries(num int) []BlogEntry {
	entries := make([]BlogEntry, num)
	rows, err := statements[ENTRIES_BY_DATE].Query(num)
	if err != nil {
		return []BlogEntry{}
	}

	i := 0
	for rows.Next() {
		err := rows.Scan(&entries[i].Title,
			&entries[i].Hook,
			&entries[i].Content,
			&entries[i].ID,
			&entries[i].Category,
			&entries[i].PublishTime)
		if err == nil {
			i++
		}
	}
	return entries[0:i]
}

func CategoryEntries(category string, limit int) []BlogEntry {
	entries := make([]BlogEntry, limit)
	rows, err := statements[ENTRIES_BY_CATEGORY].Query(category, limit)
	if err != nil {
		return []BlogEntry{}
	}

	i := 0
	for rows.Next() {
		err := rows.Scan(&entries[i].Title,
			&entries[i].Hook,
			&entries[i].Content,
			&entries[i].ID,
			&entries[i].Category,
			&entries[i].PublishTime)
		if err == nil {
			i++
		}
	}
	return entries[0:i]
}

func ActiveCategories() []Category {
	categories := make([]Category, 0)
	rows, err := statements[CATEGORIES].Query()
	if err != nil {
		return []Category{}
	}

	for rows.Next() {
		var Name string
		err := rows.Scan(&Name)
		if err == nil {
			categories = append(categories, Category{Name: Name})
		}
	}
	return categories
}
