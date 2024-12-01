package keys

import (
	"encoding/json"
	"os"
)

type Keys struct {
	Kraken struct {
		Api     string
		Private string
	}
	Robinhood struct {
		Public  string
		Private string
		Api     string
	}
}

func Get() (Keys, error) {
	k := Keys{}

	f, err := os.Open("./secrets.json")
	if err != nil {
		return k, nil
	}

	dec := json.NewDecoder(f)
	dec.Decode(&k)

	return k, nil
}
