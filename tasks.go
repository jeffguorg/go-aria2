package aria2

func (client Client) AddUris(uri ...string) (string, error) {
	res, err := client.sendRequest("aria2.addUri", uri)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}

func (client Client) Pause(gid string, force bool) (string, error) {
	method := "aria2.pause"
	if force {
		method = "aria2.forcePause"
	}
	res, err := client.sendRequest(method, gid)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}

func (client Client) Unpause(gid string, force bool) (string, error) {
	method := "aria2.unpause"
	if force {
		method = "aria2.forceUnpause"
	}
	res, err := client.sendRequest(method, gid)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}

func (client Client) Remove(gid string, force bool) (string, error) {
	method := "aria2.remove"
	if force {
		method = "aria2.forceRemove"
	}
	res, err := client.sendRequest(method, gid)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}

func (client Client) PauseAll(force bool) (string, error) {
	method := "aria2.pauseAll"
	if force {
		method = "aria2.forcePauseAll"
	}
	res, err := client.sendRequest(method)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}

func (client Client) UnpauseAll(force bool) (string, error) {
	method := "aria2.unpauseAll"
	if force {
		method = "aria2.forceUnpauseAll"
	}
	res, err := client.sendRequest(method)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}

func (client Client) RemoveAll(force bool) (string, error) {
	method := "aria2.removeAll"
	if force {
		method = "aria2.forceRemoveAll"
	}
	res, err := client.sendRequest(method)
	if err != nil {
		return "", err
	}
	if res.Error != nil {
		return "", res.Error
	}
	return res.Result.(string), nil
}
