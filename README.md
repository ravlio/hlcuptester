
Highloadcup 2018 datafile loader
================================

Usage
-----

```go
import (
	"fmt"
	"github.com/ravlio/hlcuptester"
	"log"
)

func main() {
    // Load args: path_to_data_folder, phase_number [,optional_array_with_allowed_uri_filters])
	ch, err := hlcuptester.Load("path/to/highloadcup2018/data/", 2, "accounts/new")

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
```

Other Projects
--------------
[atercattus/highloadcup_tester](https://github.com/atercattus/highloadcup_tester) - standalone validate and benchmark tool
