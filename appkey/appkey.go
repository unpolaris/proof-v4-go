package appkey

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

const (
	// NeedDebugSign 需要测试签名
	NeedDebugSign = "1"
	// PassDebugSign 跳过验签,用于压测
	PassDebugSign = "9999"

	// SignTypeHmacSha256 签名方式：HMAC-SHA256
	SignTypeHmacSha256 = "HMAC-SHA256"
	// SignTypeMD5 签名方式：MD5
	SignTypeMD5 = "MD5"

	// HeaderAppId 签名请求头：应用id
	HeaderAppId = "FZM-Ca-App-Id"
	// HeaderAppKey 签名请求头：应用公钥
	HeaderAppKey = "FZM-Ca-App-Key"
	// HeaderSign 签名请求头：签名
	HeaderSign = "FZM-Ca-Sign"
	// HeaderNonce 签名请求头：随机字符串
	HeaderNonce = "FZM-Ca-Nonce"
	// HeaderTimestamp 签名请求头：时间戳
	HeaderTimestamp = "FZM-Ca-Timestamp"
	// HeaderSignType 签名请求头：签名方式
	HeaderSignType = "FZM-Ca-Sign-Type"
	// HeaderDebugSign 签名请求头：签名测试参数
	HeaderDebugSign = "FZM-Ca-Debug-Sign"
	// HeaderAppSecret 签名请求头：应用私钥
	HeaderAppSecret = "FZM-Ca-App-Secret"
)

// Sign 签名参数
type Sign struct {
	AppId     string `json:"FZM-Ca-App-Id,optional" header:"FZM-Ca-App-Id,optional" extensions:"x-order=000"`         // 应用id
	AppKey    string `json:"FZM-Ca-App-Key,optional" header:"FZM-Ca-App-Key,optional" extensions:"x-order=001"`       // 应用公钥
	Sign      string `json:"FZM-Ca-Sign,optional" header:"FZM-Ca-Sign,optional" extensions:"x-order=002"`             // 签名
	NonceStr  string `json:"FZM-Ca-Nonce,optional" header:"FZM-Ca-Nonce,optional" extensions:"x-order=003"`           // 随机字符串
	TimeStamp string `json:"FZM-Ca-Timestamp,optional" header:"FZM-Ca-Timestamp,optional" extensions:"x-order=004"`   // 时间戳
	SignType  string `json:"FZM-Ca-Sign-Type,optional" header:"FZM-Ca-Sign-Type,optional" extensions:"x-order=005"`   // 签名方式，HMAC-SHA256（默认）或MD5
}

// NewSign return a new Sign
func NewSign(appid, appkey, signType string) *Sign {
	sign := &Sign{
		AppId: appid,
		AppKey: appkey,
		SignType: signType,
	}
	return sign
}

// GenSign 生成签名
func (s *Sign) GenSign(appSecret, nonce, timestamp string) string {
	values := make(url.Values)
	values.Set(HeaderAppId, s.AppId)
	values.Set(HeaderAppKey, s.AppKey)
	values.Set(HeaderNonce, nonce)
	values.Set(HeaderTimestamp, timestamp)

	return GenSign(values, appSecret, s.SignType)
}


// GenSign 生成签名
func GenSign(values url.Values, secret, signType string) (sign string) {
	params := make([]string, 0, len(values))
	for k := range values {
		if v := values.Get(k); len(v) > 0 {
			params = append(params, k+"="+v)
		}
	}
	sort.Strings(params)
	if secret != "" {
		params = append(params, HeaderAppSecret+"="+secret)
	}

	src := strings.Join(params, "&")
	if strings.ToUpper(signType) == SignTypeMD5 {
		return SignMD5(src)
	}

	return SignSHA256(src, secret)
}

// SignMD5 MD5签名
func SignMD5(data string) (sign string) {
	m := md5.New()
	m.Write([]byte(data))

	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}

// SignSHA256 HMAC-SHA256签名
func SignSHA256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))

	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

