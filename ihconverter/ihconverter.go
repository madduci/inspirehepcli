package ihconverter

import (
	"encoding/json"
	"fmt"
	"strings"
)

type publication struct {
	id              string
	bibtextLink     string
	authors         []author        `json:metadata:authors`
	publicationInfo publicationInfo `json:metadata:publication_info`
	titles          []title         `json:metadata:titles`
}

type author struct {
	emails    []string
	full_name string
	ids       []authorId
}

type authorId struct {
	schema string
	value  string
}

type title struct {
	source string
	title  string
}

type publicationInfo struct {
	artid            string
	journal_issue    string
	journal_title    string
	journal_volume   string
	material         string
	pubinfo_freetext string
	year             string
}

func ConvertJsonToMarkdown(jsonData string) string {
	publication := publication{}
	if err := json.Unmarshal([]byte(jsonData), &publication); err != nil {
		panic(err)
	}

	// Append title
	result := fmt.Sprintln(publication.titles[0].title)

	// Append author names
	authorsList := ""
	for i, entry := range publication.authors {
		// remove the '.1' ending
		authorName := strings.TrimSuffix(entry.ids[0].value, ".1")
		if i == 0 {
			authorsList += fmt.Sprint(authorName)
		} else {
			authorsList += fmt.Sprint(",", authorName)
		}
	}
	result += fmt.Sprintln("_", authorsList, "_")

	// Append Publication journal
	result += fmt.Sprint("**", publication.publicationInfo.pubinfo_freetext ,"**")

	fmt.Println(result)

	return result
}
