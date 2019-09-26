package aria2

import (
	"github.com/jeffguorg/jsonrpc"
)

func (client Client) TellActive(keys []string) (*jsonrpc.Response, error) {
	var params []interface{}
	if keys != nil {
		params = []interface{}{keys}
	}
	return client.sendRequest("aria2.tellActive", params...)
}

func (client Client) TellStatus(gid string, keys ...string) (interface{}, error) {
	var (
		res *jsonrpc.Response
		err error
	)
	if len(keys) > 0 {
		res, err = client.sendRequest("aria2.tellStatus", gid, keys)
	} else {
		res, err = client.sendRequest("aria2.tellStatus", gid)
	}
	if err != nil {
		return nil, err
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return res.Result, nil
}
