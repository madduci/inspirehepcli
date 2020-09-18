package main

import (
	"fmt"
	"github.com/madduci/inspirehepcli/ihclient"
	"github.com/madduci/inspirehepcli/ihconverter"
	"log"
	"os"
	"path/filepath"
)

func SaveJsonToFile(inputData string, outputFile string) {
	f, err := os.Create(outputFile)
	if err != nil{
		log.Fatalln("error:", err)
	}

	bytes, err := f.Write([]byte(inputData))
	if err != nil{
		log.Fatalln("error:", err)
	}
	_ = f.Sync()
	log.Printf("Data written successfully: %d bytes\n", bytes)
}

func GetApplicationPath() string{
	ex, err := os.Executable()
	if err != nil {
	panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func GetApplicationName() string{
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Base(ex)
	return exPath
}

func PrintUsage(){
	log.Println("Usage: %s [LiteratureId] [Options]", GetApplicationName())
	log.Println("Options:")
	log.Println("- md        output the result as markdown")
	log.Println("- html      output the result as html (default)")
}

func main() {
	log.Println("### InspireHep CLI")
	outputFolder := GetApplicationPath()
	if len(os.Args) == 1 {
		log.Fatalln("No argument supplied. Please provide the Publication ID and output type")
	}
	// Prepare required information
	literatureId := os.Args[1]
	outputFile := ""
	publicationInformation := ""

	log.Printf("Retrieving information for ID %s\n", literatureId)
	publicationData := ihclient.GetLiteratureInfoById(literatureId)

	log.Printf("Extracting data\n" )
	if len(os.Args) > 2 && os.Args[2] == "md" {
		outputFile = fmt.Sprintf("%s/%s.md", outputFolder, literatureId)
		publicationInformation = ihconverter.ConvertJsonToMarkdown(publicationData)
	} else {
		outputFile = fmt.Sprintf("%s/%s.html", outputFolder, literatureId)
		publicationInformation = ihconverter.ConvertJsonToHtml(publicationData)
	}

	log.Printf("Saving data in %s\n", outputFile)
	SaveJsonToFile(publicationInformation, outputFile)
}
