package aria2

import (
	"github.com/jeffguorg/jsonrpc"
	"time"
)

type Client struct {
	Endpoint jsonrpc.Endpoint
	Secret   string
}

func (client Client) sendRequest(method string, params ...interface{}) (*jsonrpc.Response, error) {
	req := jsonrpc.Request{
		Method: method,
		ID:     time.Now().Unix()*1000 + time.Now().UnixNano(),
	}
	if client.Secret != "" {
		req.Parameters = []interface{}{client.Secret}
	}
	req.Parameters = append(req.Parameters, params...)
	return client.Endpoint.Call(req)
}
