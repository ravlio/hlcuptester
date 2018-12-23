
Highloadcup 2018 datafile loader
================================

Usage
-----

Use `Load(path_to_data_folder, phase_number [,optional_array_with_allowed_uri_filters])`

```go
import (
	"fmt"
	"github.com/ravlio/hlcuptester"
	"log"
)

func main() {
	ch, err := hlcuptester.Load("../../../data/", 2, "accounts/new")

	if err != nil {
		log.Fatal(err)
	}

	for a := range ch {
		if a.Err != nil {
			log.Fatal(err)
		}

		fmt.Printf("URI: %s\nRequestBody: %s\nResponseCode: %d\nResponseBody:%s\n\n", a.URI, a.RequestBody, a.ResponseCode, a.ResponseBody)
	}
}```
