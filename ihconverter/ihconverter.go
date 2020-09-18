package ihconverter

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Publication struct {
	Id       string   `json:"id"`
	Created  string   `json:"created"`
	Updated  string   `json:"updated"`
	Uuid     string   `json:"uuid"`
	Metadata Metadata `json:"metadata"`
	Links    Links    `json:"links"`
}

type Links struct {
	Bibtex    string `json:"bibtex"`
	Citations string `json:"citations"`
	Json      string `json:"json"`
	LatexEu   string `json:"latext-eu"`
	LatexUs   string `json:"latext-us"`
}

type Metadata struct {
	Authors         []Author         `json:"authors"`
	PublicationInfo []PublicationInfo  `json:"publication_info"`
	Titles          []Title          `json:"titles"`
}

type Author struct {
	Emails    []string   `json:"emails"`
	FullName  string     `json:"full_name"`
	AuthorIds []AuthorId `json:"ids"`
}

type AuthorId struct {
	Schema string `json:"schema"`
	Value string `json:"value"`
}

type Title struct {
	Source string `json:"source"`
	Title  string `json:"title"`
}

type PublicationInfo struct {
	Artid         string `json:"artid"`
	JournalIssue  string `json:"journal_issue"`
	JournalTitle  string `json:"journal_title"`
	JournalVolume string `json:"journal_volume"`
	Material      string `json:"material"`
	Freetext      string `json:"pubinfo_freetext"`
	Year          uint64 `json:"year"`
}

func ConvertJsonToMarkdown(jsonData string) string {
	publication := &Publication{}
	inputByteArray := []byte(jsonData)
	if err := json.Unmarshal(inputByteArray, &publication); err != nil {
		log.Fatalln("error:", err)
	}

	result := fmt.Sprintln(publication.Metadata.Titles[0].Title)

	// Append author names
	authorsList := ""
	for i, entry := range publication.Metadata.Authors {
		// remove the '.1' ending from name
		if len(entry.AuthorIds) > 0 {
			authorName := strings.TrimSuffix(entry.AuthorIds[0].Value, ".1")
			if i == 0 {
				authorsList += fmt.Sprint(authorName)
			} else {
				authorsList += fmt.Sprint(",", authorName)
			}
		}
	}
	result += fmt.Sprintln("_", authorsList, "_")
	// Append Publication journal
	if len(publication.Metadata.PublicationInfo) > 0 {
		result += fmt.Sprint("**", publication.Metadata.PublicationInfo[0].Freetext ,"**")
	}
	// Append InspireHep Bibtex Link
	result += fmt.Sprintf(",[inspireHep](" + publication.Links.Bibtex + ")")
	return result
}