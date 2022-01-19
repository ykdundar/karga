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
	OneSymbol      string `json:"1. symbol"`
	TwoName        string `json:"2. name"`
	ThreeType      string `json:"3. type"`
	FourRegion     string `json:"4. region"`
	FiveMarketOpen string `json:"5. marketOpen"`
	SixMarketClose string `json:"6. marketClose"`
	SevenTimezone  string `json:"7. timezone"`
	EightCurrency  string `json:"8. currency"`
	NineMatchScore string `json:"9. matchScore"`
}

func SearchRequest(symbol, keywords string) (Search, error) {
	const ENDPOINT_URL string = "SYMBOL_SEARCH"
	baseUrl := base_url()
	values := baseUrl.Query()

	values.Add("function", ENDPOINT_URL)
	values.Add("symbol", symbol)
	values.Add("keywords", keywords)

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
