package aria2_test

import (
	"github.com/jeffguorg/go-aria2"
	"testing"
)

var (
	client = aria2.Client{
		Endpoint: "http://127.0.0.1:6800/jsonrpc",
	}
)

func TestClient_TellActive(t *testing.T) {
	t.Log(client.TellActive(nil))
}
