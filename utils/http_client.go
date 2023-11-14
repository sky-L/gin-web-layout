package utils

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func Post() {

	_, r2, err := gorequest.New().Post("https://glados.one/api/user/checkin").
		SendMap(map[string]string{"token": "glados.network"}).
		AppendHeader("Cookie", "_ga=GA1.2.1798683538.1662461089; koa:sess=eyJjb2RlIjoiTTlPSEgtUTg4SlEtRFg3MkQtUjA0Uk4iLCJ1c2VySWQiOjIwMDg5MywiX2V4cGlyZSI6MTY4ODM4MjQzNTc3NywiX21heEFnZSI6MjU5MjAwMDAwMDB9; koa:sess.sig=FeO-epWUGVtC-neGrco9HwDq1I4; _gid=GA1.2.2072231038.1663213848; _gat_gtag_UA_104464600_2=1").
		AppendHeader("authorization", "63205931141406492246251978760129-1080-1920").
		End()

	fmt.Println(r2, err)

}
