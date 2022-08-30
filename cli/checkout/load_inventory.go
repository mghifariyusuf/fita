package checkout

import (
	"encoding/json"
	"log"
	"net/http"
)

func LoadInventory() (*Inventory, error) {
	var payload *Inventory

	content, err := http.Get(jsonFile)
	if err != nil {
		log.Println(err)
		return payload, err
	}

	err = json.NewDecoder(content.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		return payload, err
	}

	return payload, nil
}
