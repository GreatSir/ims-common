package org

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Api struct {
	Url string
}
func (a *Api)GetUrl() string {
	return a.Url
}
func (a *Api)Req(api,method string ,params map[string]interface{}) ([]byte,error) {
	body,err:=json.Marshal(params)
	if err!=nil{
		return nil, err
	}
	request,err:=http.NewRequest(method,api,bytes.NewReader(body))
	if err!=nil{
		return nil,err
	}
	client:=&http.Client{}
	response,err:=client.Do(request)
	if err!=nil{
		return nil, err
	}
	defer response.Body.Close()
	responseBody,err:=io.ReadAll(response.Body)

	return responseBody, err

}
