package stack_exchange

import (
	"encoding/json"
	"fmt"

	API_Types "staq/pkg/stack_exchange/types"
	API "staq/pkg/stack_exchange/utils"
)

func GetStackExchangeArticles() {
	queries := map[string]string{}
	queries["site"] = "stackoverflow"
	queries["sort"] = "activity"
	queries["order"] = "desc"
	url := API.GenerateURL("articles", queries)

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

	var articles []API_Types.ArticleInfo
	response := API_Types.APIResponse{Items: &articles}
	err = json.Unmarshal(response_byte, &response)
	if err != nil {
		panic(err)
	}

	for _, article := range articles {
		fmt.Printf("%s: %s\n", article.Title, article.Tags)
	}
}
