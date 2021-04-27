package notice

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"github.com/tidwall/gjson"
)

type Api struct {
	Url string
}

func (a *Api)GetUrl() string {
	return a.Url
}
//发送微信消息
func (a *Api)SendWxMsg(msgType uint,params map[string]interface{}) error {
	url:=a.Url
	switch msgType {
	case 1:
		//单条发送 文本消息
		url+="/notify/sendTextMsg"
		break
	case 2:
		//批量发送 文本消息
		url+="notify/batchSendTextMsg"
		break
	case 3:
		//单条发送文本卡片消息
		url+="notify/sendTextCardMsg"
		break
	case 4:
		//批量发送文本卡片消息
		url+="notify/batchSendTextCardMsg"
		break
	case 5:
		//单条发送markdown消息
		url+="notify/sendMarkdownMsg"
		break
	case 6:
		//批量发送markdown消息
		url+="notify/batchSendMarkdownMsg"
		break
	case 7:
		//单条发送图文消息
		url+="notify/sendPictureTextMsg"
		break
	case 8:
		//批量发送图文消息
		url+="notify/batchSendPictureTextMsg"
		break
	default:

	}
	body,_:=json.Marshal(params)
	client:=&http.Client{}
	request,err:=http.NewRequest("POST",url,bytes.NewReader(body))
	if err!=nil{
		return err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	response,err:=client.Do(request)
	if err!=nil{
		return err
	}
	defer response.Body.Close()
	responseBody,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		return err
	}
	result:=gjson.GetBytes(responseBody,"code")

	if result.Exists()&&result.String()=="000"{
		return nil
	}
	if result.Exists()&&result.String()!="000"{
		msgRes:=gjson.GetBytes(responseBody,"msg")
		if msgRes.Exists(){
			return errors.New(msgRes.String())
		}

	}
	return errors.New("请求异常")
	
}
