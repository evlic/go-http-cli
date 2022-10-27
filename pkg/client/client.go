package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BaseClient http.Client

func NewBaseClient(header http.Header) *BaseClient {
	return (*BaseClient)(http.DefaultClient)
}

func (c *BaseClient) DoReq(r *http.Request) (resp *http.Response) {
	return c.DoReq(r)
}

func handlerReqError(err error) (*http.Request, error) {
	if err != nil {
		return nil, fmt.Errorf("new request err: %w", err)
	}
	return nil, err
}
func NewGetReq(backend string, param url.Values) (r *http.Request, err error) {
	realUrl, err := Mapping(backend, param)
	if err != nil {
		return handlerReqError(err)
	}

	req, err := http.NewRequest(http.MethodGet, realUrl, nil)
	if err != nil {
		return handlerReqError(err)
	}

	return req, nil
}

func NewPostReq(backend string, body any, param url.Values) (r *http.Request, err error) {
	realUrl, err := Mapping(backend, param)
	if err != nil {
		return handlerReqError(err)
	}

	data, err := json.Marshal(body)
	if err != nil {
		return handlerReqError(err)
	}

	req, err := http.NewRequest(http.MethodPost, realUrl, bytes.NewBuffer(data))
	if err != nil {
		return handlerReqError(err)
	}

	return req, nil
}

func Mapping(backend string, param url.Values) (string, error) {
	parse, err := url.Parse(backend)
	if err != nil {
		return "", err
	}
	parse.RawQuery += param.Encode()
	return parse.String(), nil
}
