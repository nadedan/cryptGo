package robinhood

type Endpoint struct {
	desc string
	Url  string
	Body []byte
}

func (e Endpoint) String() string {
	return e.desc
}

func Account() Endpoint {
	return Endpoint{
		desc: "Account",
		Url:  "/api/v1/crypto/trading/accounts/",
	}
}
