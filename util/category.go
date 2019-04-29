package util

func GetCategory(category string) string {

	categories := map[string]string{
		"science":  "C29",
		"hardware": "C6",
		"software": "C4",
		"vhiecle":  "C30",
		"security": "C14",
	}

	id := categories[category]
	return id
}
