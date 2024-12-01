package keys

import (
	_ "embed"
	"encoding/json"
	"strings"
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

//go:embed secrets.json
var secrets string

func Get() (Keys, error) {
	k := Keys{}

	dec := json.NewDecoder(strings.NewReader(secrets))
	dec.Decode(&k)

	return k, nil
}
