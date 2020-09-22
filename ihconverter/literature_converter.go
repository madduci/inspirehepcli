package ihconverter

import (
	"fmt"
)

func parseLiteraturePublication(publication *Publication, outputType OutputType) string {
	// Append author names
	authorsList := ""
	for i, entry := range publication.Metadata.Authors {
		if len(entry.AuthorIds) > 0 {
			if i == 0 {
				authorsList += fmt.Sprint(entry.FullName)
			} else {
				authorsList += fmt.Sprintf(" - %s", entry.FullName)
			}
		}
	}

	// Extract Publication journal
	fulltextName := ""
	if len(publication.Metadata.PublicationInfo) > 0 && len(publication.Metadata.PublicationInfo[0].Freetext) > 0 {
		fulltextName = publication.Metadata.PublicationInfo[0].Freetext
	} else {
		fulltextName = "-"
	}

	// Append Title, Authors list, Fulltext name and InspireHep Bibtex Link
	result := fmt.Sprintln(publication.Metadata.Titles[0].Title)
	switch outputType {
	case htmlType:
		result += fmt.Sprintf("<i>%s</i>\n", authorsList)
		result += fmt.Sprintf("%s, <a href='https://www.inspirehep.net/%s' title='InspireHep Link' target='_blank' rel='noopener>%s</a>\n", fulltextName, publication.Id, "InspireHep")
	case markdownType:
		result += fmt.Sprintf("_%s_\n", authorsList)
		result += fmt.Sprintf("%s, [InspireHep](https://www.inspirehep.net/%s)\n", fulltextName, publication.Id)
	}

	return result
}

func ConvertLiteratureJsonToMarkdown(jsonData string) string {
	publication := getPublicationFromJson(jsonData)
	return parseLiteraturePublication(publication, markdownType)
}

func ConvertLiteratureJsonToHtml(jsonData string) string {
	publication := getPublicationFromJson(jsonData)
	return parseLiteraturePublication(publication, htmlType)
}
