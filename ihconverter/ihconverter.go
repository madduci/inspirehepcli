package ihconverter

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type OutputType int

const (
	htmlType OutputType = iota
	markdownType
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
	Authors         []Author          `json:"authors"`
	PublicationInfo []PublicationInfo `json:"publication_info"`
	Titles          []Title           `json:"titles"`
	ReportNumbers   []ReportNumber    `json:"report_numbers"`
}

type Author struct {
	Emails    []string   `json:"emails"`
	FullName  string     `json:"full_name"`
	AuthorIds []AuthorId `json:"ids"`
}

type AuthorId struct {
	Schema string `json:"schema"`
	Value  string `json:"value"`
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

type ReportNumber struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

func getPublicationFromJson(jsonData string) *Publication {
	publication := &Publication{}
	inputByteArray := []byte(jsonData)
	if err := json.Unmarshal(inputByteArray, &publication); err != nil {
		log.Fatalln("error:", err)
	}
	return publication
}

// Converts to Lastname, F. (firstname truncated)
func getShortAuthorNameFromFullName(fullName string) string {
	name := strings.Split(fullName, ",")
	name[1] = strings.TrimLeft(name[1], " ")
	firstNameLetter := name[1][0]

	return fmt.Sprintf("%s, %c.", name[0], firstNameLetter)
}

func parsePublication(publication *Publication, outputType OutputType) string {
	// Append author names
	authorsList := ""
	for i, entry := range publication.Metadata.Authors {
		currentAuthorName := getShortAuthorNameFromFullName(entry.FullName)
		if len(entry.AuthorIds) > 0 {
			if i == 0 {
				authorsList += fmt.Sprint(currentAuthorName)
			} else {
				authorsList += fmt.Sprintf(", %s", currentAuthorName)
			}
		}
	}

	// Extract Publication journal - if empty, get the report number
	fulltextName := ""
	if len(publication.Metadata.PublicationInfo) > 0 && len(publication.Metadata.PublicationInfo[0].Freetext) > 0 {
		fulltextName = publication.Metadata.PublicationInfo[0].Freetext
	} else {
		for i, reportNumber := range publication.Metadata.ReportNumbers {
			if i == 0 {
				fulltextName += fmt.Sprint(reportNumber.Value)
			} else {
				fulltextName += fmt.Sprintf(", %s", reportNumber.Value)
			}
		}
	}

	// Append Title, Authors list, Fulltext name and InspireHep Bibtex Link
	result := fmt.Sprintln(publication.Metadata.Titles[0].Title)
	switch outputType {
	case htmlType:
		result += fmt.Sprintf("<i>%s</i>\n", authorsList)
		result += fmt.Sprintf("%s, <a href='https://www.inspirehep.net/literature?q=%s' title='InspireHep Link' target='_blank' rel='noopener>%s</a>\n", fulltextName, publication.Id, "InspireHep")
	case markdownType:
		result += fmt.Sprintf("_%s_\n", authorsList)
		result += fmt.Sprintf("%s, [InspireHep](https://www.inspirehep.net/literature?q=%s)\n", fulltextName, publication.Id)
	}

	return result
}

func ConvertJsonToMarkdown(jsonData string) string {
	publication := getPublicationFromJson(jsonData)
	return parsePublication(publication, markdownType)
}

func ConvertJsonToHtml(jsonData string) string {
	publication := getPublicationFromJson(jsonData)
	return parsePublication(publication, htmlType)
}
