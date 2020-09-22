package ihclient

import (
	"github.com/go-resty/resty/v2"
)

const (
	inspireHepUrl string = "https://inspirehep.net"
)

func getDataFromInspireHep(searchItem string) string {
	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get(searchItem)

	if err != nil {
		panic(err)
	}
	return string(resp.Body()[:])
}

func GetLiteratureInfoById(literatureId string) string {
	literatureApiURL := "/api/literature/"
	return getDataFromInspireHep(inspireHepUrl + literatureApiURL + literatureId)
}

func GetLiteratureInfoByArxiv(literatureId string) string {
	arxivApiUrl := "/api/arxiv/"
	return getDataFromInspireHep(inspireHepUrl + arxivApiUrl + literatureId)
}
