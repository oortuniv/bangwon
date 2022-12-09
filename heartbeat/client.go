package heartbeat

import (
	"bangwon/api/response"
	_type "bangwon/type"
	"github.com/go-resty/resty/v2"
)

type ApiClient struct {
	client *resty.Client
}

func NewApiClient() ApiClient {
	return ApiClient{
		client: resty.New(),
	}
}

func (a ApiClient) GetStatus(bangwon string) (_type.Status, error) {
	res := response.Status{}
	get, err := a.client.R().
		SetResult(&res).
		Get(bangwon + "/status")
	if get.IsSuccess() {
		return res.To(), err
	}
	return _type.Status{}, nil
}

func (a ApiClient) LockStatus(bangwon string) error {
	post, err := a.client.R().
		Post(bangwon + "/status/lock")
	if post.IsSuccess() {
		return nil
	}
	return err
}

func (a ApiClient) UnlockStatus(bangwon string) error {
	post, err := a.client.R().
		Post(bangwon + "/status/unlock")
	if post.IsSuccess() {
		return nil
	}
	return err
}
