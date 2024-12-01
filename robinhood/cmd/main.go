package main

import (
	"fmt"

	"github.com/nadedan/cryptGo/keys"
	"github.com/nadedan/cryptGo/robinhood"
)

func main() {
	k, err := keys.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", k)
	api, err := robinhood.NewHandler(robinhood.Keys{
		Api:     k.Robinhood.Api,
		Private: k.Robinhood.Private,
	})
	if err != nil {
		panic(err)
	}

	res, err := api.Get(robinhood.Account())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
