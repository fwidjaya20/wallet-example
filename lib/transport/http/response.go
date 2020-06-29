package http

import (
	"context"
	"github.com/fwidjaya20/wallet-example/lib/transport/http/chi"
	"time"
)

type ResponseStructure struct {
	Data     interface{}            `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}

func Response(ctx context.Context, data interface{}, metadata map[string]interface{}) interface{} {
	meta := make(map[string]interface{})

	for k, v := range metadata {
		meta[k] = v
	}

	if ctx.Value(chi.KEY_REQUEST_TIME) != nil {
		meta["request_time"] = time.Since(ctx.Value(chi.KEY_REQUEST_TIME).(time.Time)).Seconds()
	}

	return ResponseStructure{
		Data:     data,
		Metadata: meta,
	}
}