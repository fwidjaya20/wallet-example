package server

import (
	"github.com/fwidjaya20/wallet-example/lib/server/http"
	"github.com/go-kit/kit/endpoint"
	kitHttp "github.com/go-kit/kit/transport/http"
)

type HTTPOption struct {
	DecodeModel interface{}
	Encoder http.EncodeFunc
	Decoder http.DecodeFunc
}

func NewHTTPServer(
	endpoint endpoint.Endpoint,
	httpOption HTTPOption,
	serverOption []kitHttp.ServerOption) *kitHttp.Server {
	if httpOption.Encoder == nil {
		httpOption.Encoder = http.Encode
	}
	if httpOption.Decoder == nil {
		httpOption.Decoder = http.Decode
	}

	return kitHttp.NewServer(endpoint, httpOption.Decoder(httpOption.DecodeModel), httpOption.Encoder(), serverOption...)
}
