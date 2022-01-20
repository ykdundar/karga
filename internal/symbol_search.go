package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Search struct {
	BestMatches []BestMatches `json:"bestMatches"`
}

type BestMatches struct {
	Symbol      string `json:"1. symbol"`
	Name        string `json:"2. name"`
	Type        string `json:"3. type"`
	Region      string `json:"4. region"`
	MarketOpen  string `json:"5. marketOpen"`
	MarketClose string `json:"6. marketClose"`
	Timezone    string `json:"7. timezone"`
	Currency    string `json:"8. currency"`
	MatchScore  string `json:"9. matchScore"`
}

func SearchRequest(keyword string) (Search, error) {
	const ENDPOINT_URL string = "SYMBOL_SEARCH"

	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("keywords", keyword)

	baseUrl.RawQuery = values.Encode()

	response, err := http.Get(baseUrl.String())

	if err != nil {
		return Search{}, errors.New("the HTTP request is failed with an error")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		search := Search{}

		json.Unmarshal(data, &search)

		return search, nil
	}

}
