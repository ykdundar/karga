package internal

import (
	"fmt"
	"net/url"
	"os"
)

func base_url() *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "alphavantage.co",
		Path:       "query",
		ForceQuery: false,
		RawQuery:   fmt.Sprintf("apikey=%s", os.Getenv("KARGA_TOKEN")),
	}
}
