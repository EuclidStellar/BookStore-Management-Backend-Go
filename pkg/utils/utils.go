package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Parsebody(r *http.Request, data interface{}) {
	if body , err := ioutil.ReadAll(r.Body); err != nil {
		log.Fatal(err)
	} else {
		if err := json.Unmarshal(body, data); err != nil {
			log.Fatal(err)
		}
	}
}