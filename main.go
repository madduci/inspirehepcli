package main

import (
	"fmt"
	"github.com/madduci/inspirehepcli/ihclient"
	"github.com/madduci/inspirehepcli/ihconverter"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	outputType      = kingpin.Flag("output", "The desired output type. Defaults to html.").Default("html").Short('o').String()
	useArxivId      = kingpin.Flag("arxiv", "Uses the Arxiv ID for the search (default).").Short('a').Default("true").Bool()
	useLiteratureId = kingpin.Flag("literature", "Uses the Literature ID for the search.").Short('l').Default("false").Bool()
	publicationId   = kingpin.Arg("id", "Publication ID to look for in InspireHEP.").Required().String()
)

func saveJsonToFile(inputData string, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalln("error:", err)
	}

	bytes, err := f.Write([]byte(inputData))
	if err != nil {
		log.Fatalln("error:", err)
	}
	_ = f.Sync()
	log.Printf("Data written successfully: %d bytes\n", bytes)
}

func getApplicationPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func main() {
	kingpin.Version("1.2.0")
	kingpin.Parse()

	log.Println("### InspireHep CLI")
	outputFolder := getApplicationPath()

	// Prepare required information
	outputFile := ""
	publicationInformation := ""

	log.Printf("Retrieving information for ID %s\n", *publicationId)
	publicationData := ""
	if *useArxivId {
		publicationData = ihclient.GetLiteratureInfoByArxiv(*publicationId)
	} else if *useLiteratureId {
		publicationData = ihclient.GetLiteratureInfoById(*publicationId)
	}

	log.Printf("Extracting data\n")
	if strings.Contains(*outputType, "md") {
		outputFile = fmt.Sprintf("%s/%s.md", outputFolder, *publicationId)
		publicationInformation = ihconverter.ConvertJsonToMarkdown(publicationData)
	} else {
		outputFile = fmt.Sprintf("%s/%s.html", outputFolder, *publicationId)
		publicationInformation = ihconverter.ConvertJsonToHtml(publicationData)
	}

	log.Printf("Saving data in %s\n", outputFile)
	saveJsonToFile(publicationInformation, outputFile)
}
