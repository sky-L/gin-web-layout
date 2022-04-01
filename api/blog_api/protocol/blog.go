package protocol

type BlogListReq struct {
	Id   int `form:"id" binding:"required"`
}

type BlogResp struct {
	Name string `json:"name"`
}
