package ims_common

import (
	"github.com/GreatSir/ims-common/notice"
	"github.com/GreatSir/ims-common/org"
	"github.com/facebookgo/inject"
)

type Client struct {
	NoticeApi *notice.Api `inject:""`
	OrgApi *org.Api `inject:""`
}

func NewClient(noticeUrl,orgUrl string) Client {
	var g inject.Graph
	var c Client
	noticeApi:=notice.Api{Url:noticeUrl}
	orgApi:=org.Api{Url:orgUrl}
	g.Provide(&inject.Object{Value:&c},&inject.Object{Value:&noticeApi},&inject.Object{Value:&orgApi})
	g.Populate()
	return c
}