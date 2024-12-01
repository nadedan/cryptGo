package robinhood

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/nacl/sign"
)

type Handler struct {
	keys    Keys
	baseUrl string
}

type Keys struct {
	Api     string
	Private string

	seed    string
	private *[64]byte
	public  *[32]byte
}

func timestamp() int {
	return int(time.Now().UTC().UnixMilli())
}

func NewHandler(keys Keys) (*Handler, error) {
	h := &Handler{
		keys: Keys{
			Api:  keys.Api,
			seed: keys.Private,
		},
		baseUrl: "https://trading.robinhood.com",
	}

	pk, sk, err := sign.GenerateKey(strings.NewReader(h.keys.seed))
	if err != nil {
		return nil, fmt.Errorf("NewHandler: could not generate key: %w", err)
	}

	h.keys.private = sk
	h.keys.public = pk

	return h, nil
}

func (h *Handler) Get(ep Endpoint) ([]byte, error) {
	req, err := http.NewRequest("GET", h.baseUrl+ep.Url, bytes.NewReader(ep.Body))
	if err != nil {
		return nil, fmt.Errorf("%T.Get: %s: could not make request: %w", h, ep, err)
	}

	req.Header = h.header("GET", ep.Url, bytes.NewReader(ep.Body))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%T.Get: %s: client couldn't do: %w", h, ep, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%T.Get: %s: could not read response: %w", h, ep, err)
	}

	return body, nil
}

func (h *Handler) header(method string, path string, body io.Reader) http.Header {

	hdr := http.Header{}

	ts := timestamp()
	signature := sign.Sign([]byte{}, []byte(fmt.Sprintf("%s%s%s%d", method, path, body, ts)), h.keys.private)

	hdr.Add("x-api-key", h.keys.Api)
	hdr.Add("x-signature", base64.StdEncoding.EncodeToString(signature))
	hdr.Add("x-timestamp", strconv.Itoa(ts))

	return hdr
}
