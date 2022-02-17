package protocol

type BlogListReq struct {
	Id int `json:"id" form:"id"`
}

type BlogResp struct {
	Name string `json:"name"`
}
