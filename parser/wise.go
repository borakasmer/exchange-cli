package parser

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Exchange struct {
	Source string
	Target string
	Value  float32
	Time   int64
}

func ParseWise(exchange string) Exchange {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Get("https://wise.com/rates/live?target=TRY&source=" + exchange)
	if err != nil {
		log.Fatal(err)
	}
	var exchangeData Exchange
	defer res.Body.Close()
	if res.StatusCode == 200 {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&exchangeData)
		if err != nil {
			log.Fatal(err)
		}
	}
	return exchangeData
}
