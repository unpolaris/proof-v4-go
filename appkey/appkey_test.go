package appkey

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGenSign(t *testing.T) {
	appId := "fe2458ae000b4253a49bcb196ebf9fa2"
	appKey := "467F30291CF59D900E03533293A9AA65"
	appSecret := "7+3B1XaQF+o1MqkL3Pn4IO/0tlVed9rDYIfUNvXc2Og="
	signType := "other"

	timestamp := "1631084416"
	nonceStr := "rSb9j9e7"

	p := make(url.Values)
	p.Set("app_id", appId)
	p.Set("app_key", appKey)
	p.Set("nonce_str", nonceStr)
	p.Set("timestamp", timestamp)
	sign := GenSign(p, appSecret, signType)
	fmt.Println(sign)

	// app_id=fe2458ae000b4253a49bcb196ebf9fa2&app_key=467F30291CF59D900E03533293A9AA65&nonce_str=rSb9j9e7&timestamp=1631084416&key=7+3B1XaQF+o1MqkL3Pn4IO/0tlVed9rDYIfUNvXc2Og=
	// app_id=fe2458ae000b4253a49bcb196ebf9fa2&app_key=467F30291CF59D900E03533293A9AA65&nonce_str=rSb9j9e7&timestamp=1631084416&key=7+3B1XaQF+o1MqkL3Pn4IO/0tlVed9rDYIfUNvXc2Og=

	// 52EB95E88D3A528F02972E15D804498A5529A06CEFCFF19475E165698050EAFB
	// 52EB95E88D3A528F02972E15D804498A5529A06CEFCFF19475E165698050EAFB
}
