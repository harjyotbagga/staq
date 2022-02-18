package stack_exchange

import (
	"encoding/json"
	"fmt"

	API_Types "staq/pkg/stack_exchange/types"
	API "staq/pkg/stack_exchange/utils"
)

func GetStackExchangeSites() {
	url := "https://api.stackexchange.com/2.3/sites"

	request := API.APIRequest{
		URL:         url,
		Method:      API.GET,
		ContentType: API.JSON,
		Body:        nil,
	}
	response_byte, err := API.CallAPI(request)
	if err != nil {
		panic(err)
	}
	var response API_Types.API_SitesResponse
	err = json.Unmarshal(response_byte, &response)
	if err != nil {
		panic(err)
	}

	for _, site := range response.Items {
		fmt.Printf("%s: %s (%s)\n", site.Slug, site.Name, site.SiteURL)
	}
}
