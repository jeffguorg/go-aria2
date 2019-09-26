package aria2

func interfacesToStrings(in []interface{}) []string {
	var arr []string
	for _, item := range in {
		switch item.(type) {
		case string:
			arr = append(arr, item.(string))
		}
	}
	return arr
}

func (client Client) GetVersion() (string, []string, error) {
	res, err := client.sendRequest("aria2.getVersion")
	if err != nil {
		return "", nil, err
	}
	if res.Error != nil {
		return "", nil, err
	}
	result := res.Result.(map[string]interface{})
	return result["version"].(string), interfacesToStrings(result["enabledFeatures"].([]interface{})), nil
}

func (client Client) Shutdown(force bool) error {
	method := "aria2.shutdown"
	if force {
		method = "aria2.forceShutdown"
	}
	_, err := client.sendRequest(method)
	return err
}
func (client Client) SaveSession() error {
	_, err := client.sendRequest("aria2.saveSession")
	return err
}
