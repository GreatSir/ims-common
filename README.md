# 公共服务

安装

    go get -u github.com/GreatSir/ims-common
使用

    package main
    import "github.com/GreatSir/ims-common"
    func main(){
    c:=ims_common.NewClient("","")
    params:=make(map[string]interface{})
    params["toAll"] = 1
    params["content"] = ""
    c.NoticeApi.SendWxMsg(1,params)
    c.OrgApi.Req("/employee/list/findAll","GET",nil)
    }
    
