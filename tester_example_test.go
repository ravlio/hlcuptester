package hlcuptester_test

import (
	"fmt"
	"github.com/ravlio/hlcuptester"
	"log"
)

func ExampleUsage() {
	ch, err := hlcuptester.Load("../../../data/", 2, hlcuptester.InsidePath("accounts/new"))

	if err != nil {
		log.Fatal(err)
	}

	for a := range ch {
		if a.Err != nil {
			log.Fatal(err)
		}

		fmt.Printf("URI: %s\nRequestBody: %s\nResponseStatus: %d\nResponseBody:%s\n\n", a.URI, a.RequestBody, a.ResponseStatus, a.ResponseBody)
	}
}
