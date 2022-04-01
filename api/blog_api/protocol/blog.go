package protocol

type BlogListReq struct {
	Id int `json:"id" form:"id" validate:"required" binding:"required"`
}

type BlogResp struct {
	Name string `json:"name"`
}
