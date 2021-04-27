package ims_common

import "testing"

func TestNewClient(t *testing.T) {
	c:=NewClient("","org")
	c.NoticeApi.GetUrl()
	c.OrgApi.GetUrl()

}
