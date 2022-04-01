package test

import (
	"fmt"
	"testing"
)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func TestJson(t *testing.T)  {

	l := []string{"l", "l"}
	fmt.Println(k(l))


	//jsonStr := "{\"srcType\":\"web\",\"pageUrl\":\"https://tf.weimiaotj.cn/?id=672&site_id=457894#/landPage?from=wm&adid=28495568&requestid=1640225060865_4ace77a0ff5043f0b58ea5709684a803\",\"srcId\":\"ds-202112208914\",\"advertiser_id\":\"fa9985da4eabcdc2c548\",\"dataList\":[{\"cvType\":\"PAY\",\"cvTime\":1640225421000,\"userIdType\":\"other\",\"userId\":\"8002420211223101010\",\"requestId\":\"1640225060865_4ace77a0ff5043f0b58ea5709684a803\",\"creativeId\":\"28495568\"}]}"
	//r := gjson.Parse(jsonStr)
	//fmt.Println(r.Get("dataList").Array())

}
