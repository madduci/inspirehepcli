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

func main() {
	log.Println("### InspireHep CLI")
	outputFolder := GetApplicationPath()
	if len(os.Args) == 1 {
		log.Fatalln("No argument supplied. Please provide the Publication ID")
	}
	// Prepare required information
	literatureId := os.Args[1]
	outputFile := fmt.Sprintf("%s/%s.md", outputFolder, literatureId)

	log.Printf("Retrieving information for ID %s\n", literatureId)
	publicationData := ihclient.GetLiteratureInfoById(literatureId)

	log.Printf("Extracting data\n" )
	publicationInformation := ihconverter.ConvertJsonToMarkdown(publicationData)

	log.Printf("Saving data in %s\n", outputFile)
	SaveJsonToFile(publicationInformation, outputFile)
}
