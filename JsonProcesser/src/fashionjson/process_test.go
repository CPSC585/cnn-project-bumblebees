package fashionjson

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	const jsonStream = `{
	"info": {
		"url": "https://www.wish.com",
		"dateCreated": "2-27-2018",
		"version": "2",
		"description": "Train Set for FGVC5 CVPR 2018 by https://www.wish.com",
		"year": "2018"
		},
	} `
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	token, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%T: %v\n", token, token)

	token, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%T: %v\n", token, token)

	// while the array contains values
	var info TrainInfo
	// decode an array value (Message)
	err = dec.Decode(&info)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%T: %+v\n", info, info)
	log.Fatal("Debug")
}
