package ihconverter

import (
	"encoding/json"
	"log"
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

func getPublicationFromJson(jsonData string) *Publication {
	publication := &Publication{}
	inputByteArray := []byte(jsonData)
	if err := json.Unmarshal(inputByteArray, &publication); err != nil {
		log.Fatalln("error:", err)
	}
	return publication
}
