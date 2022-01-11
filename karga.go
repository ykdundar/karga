package main

import (
	"fmt"

	"github.com/ykdundar/karga/internal"
)

func main() {
	response, err := internal.OverviewRequest("IBM")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
