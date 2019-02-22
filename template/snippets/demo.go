// A tiny cli tool wrapping the ICNDB api.
// http://api.icndb.com/
// See also: https://www.youtube.com/watch?v=TA_hMq-JUOE
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const APITmpl = `http://api.icndb.com/jokes/random/%d`

type JokeResponse struct {
	Value []struct {
		Joke string
	}
}

func main() {
	nJokes := flag.Int("n", 1, "number of jokes to request")
	flag.Parse()

	url := fmt.Sprintf(APITmpl, *nJokes)
	res, err := http.Get(url)
	chk(err)
	defer res.Body.Close()

	var jr JokeResponse
	chk(json.NewDecoder(res.Body).Decode(&jr))
	for _, v := range jr.Value {
		fmt.Println(v.Joke)
	}
}

func chk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
