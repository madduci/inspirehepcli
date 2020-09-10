package ihclient

import (
	"github.com/go-resty/resty/v2"
)

const (
	inspireHepUrl string = "https://inspirehep.net"
	apiURL string = "/api/literature/"
)

func GetLiteratureInfoById(literatureId string) string {
	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

	resp, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		Get(inspireHepUrl + apiURL + literatureId)

	if err != nil {
		panic(err)
	}
	return string(resp.Body()[:])
}
