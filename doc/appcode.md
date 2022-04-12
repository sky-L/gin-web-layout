## Api自定义code

### 接口自定义 code，msg

一般情况下，我们和前端联调接口的时候，前端只需要判断成功状态 code == successCode 走正向流程, 异常流程只判断 code != successCode, 把后端返回的 msg 直接提示给用户。
这样做的好处就是用户就能知道接下来怎么做，是根据提示操作，完成正向流程，还是直接退出放弃，用户都可以决定，而不用再反馈给客服-运营-产品-技术， 技术再找日志，麻烦的一批。

当只有正向、异常流程情况下，一般只需要 successCode、failCode 就够了。

但是有时候会有多分支流程，满足A条件执行A1流程，满足B条件执行B1流程... , 这个时候就需要自定义code了，前端也就需要用特殊的 code 写不同的逻辑了


#### 定义
新建 pkg/appcode/error.go, 标准库 Error 是个接口，且只有一个方法

```php
type error interface {
	Error() string
}

```

那我们自定义 error 只要实现这一个方式就可以

```php

type AppError struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func New(code int, msg string) *AppError {
	return &AppError{code, msg}
}

func (a AppError) Error () string  {
	return fmt.Sprintf("code %d msg:%s", a.Code, a.Msg)
}

```
#### 使用

原来的 api.NewResponse:

```php
func NewResponse(c *gin.Context, err error, data interface{}) {
	res := &Response{Data: data}
    if err != nil {
		res.Code = -1
		res.Message = err.Error()
	} else {
		res.Code = SuccessCode
		res.Message = "success"
	}
	c.JSON(http.StatusOK, res)
}

```

改造一下:

```php

func NewResponse(c *gin.Context, err error, data interface{}) {
	res := &Response{Data: data}

	switch typed := err.(type) {
	case appcode.AppError:
		res.Code = typed.Code
		res.Message = typed.Msg
	default:
		if typed != nil {
			res.Code = -1
			res.Message = err.Error()
		} else {
			res.Code = SuccessCode
			res.Message = "success"
		}
	}
	c.JSON(http.StatusOK, res)
}

```