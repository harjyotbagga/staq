package stack_exchange

import (
	"encoding/json"
	"fmt"

	API_Types "staq/pkg/stack_exchange/types"
	API "staq/pkg/stack_exchange/utils"
)

func GetStackExchangeSites() {
	queries := map[string]string{}
	url := API.GenerateURL("sites", queries)

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
	var sites []API_Types.SiteInfo
	response := API_Types.APIResponse{Items: &sites}
	err = json.Unmarshal(response_byte, &response)
	if err != nil {
		panic(err)
	}

	for _, site := range sites {
		fmt.Printf("%s: %s (%s)\n", site.Slug, site.Name, site.SiteURL)
	}
}
