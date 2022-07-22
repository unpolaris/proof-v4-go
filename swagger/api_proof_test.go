package swagger

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/antihax/optional"
	"proof-v4-go/appkey"
)

func TestProofApiService_ApiAutoProofsPost(t *testing.T) {
	config := NewConfiguration()
	config.BasePath = "http://183.134.99.140:46789"

	client := NewAPIClient(config)

	data := []ModelReqAutoProof{
		{
			Id: "1",
			Ext: &ModelProofExtInfo{
				TemplateId: 76,
				Version:    "v4",

			},
			Detail: "{\"加工信息\":\"{\\\"切割时间\\\":1638436740,\\\"雕刻时间 \\\":1638436741,\\\"雕刻人\\\":\\\"高大师\\\"}\",\"玉石信息\":{\"产品名称\":\"牡丹纹双 龙耳盖瓶(白玉)\",\"玉质特征描述\":[\"玉质为和田籽料，温和坚韧。\",\"镂雕牡丹为钮，一点留皮 俏色，侧面镂琢双凤，精美绝伦。\",\"瓶身以高浮雕雕琢绽放牡丹，花枝叶茂，花上方一对喜鹊嬉戏而飞， 主次呼应，意寓“花开富贵”\"],\"鉴定机构官网\":\"https://www.33.cn/home\",\"鉴定证书\": [\"http://img.mp.itc.cn/upload/20161023/3ea477df1002455c94d8244443e0eb9a_th.jpeg \"]}}",
		},
	}
	appid := "fe2458ae000b4253a49bcb196ebf9fa2"
	appKey := "24BF2405A5CB70605EFB516281E27058"
	signType := "HMAC-SHA256"
	appSecret := "cFZJHII4AxOv7/rugXeYp7biSUlnJ0VeFuMAbYfBkUc="
	nonce := "111001"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	t.Log("timestamp=", timestamp)

	sign := appkey.NewSign(appid, appKey, signType)
	signStr := sign.GenSign(appSecret, nonce, timestamp)
	t.Log("signStr=", signStr)

	signOpt := &ProofApiApiAutoProofsPostOpts{
		FZMCaAppId:     optional.NewString(appid),
		FZMCaAppKey:    optional.NewString(appKey),
		FZMCaSign:      optional.NewString(signStr),
		FZMCaNonce:     optional.NewString(nonce),
		FZMCaTimestamp: optional.NewString(timestamp),
		FZMCaSignType:  optional.NewString(signType),
	}
	got, got1, err := client.ProofApi.ApiAutoProofsPost(context.Background(), data, signOpt)
	if err != nil {
		t.Log("err=", err)
		return
	}
	t.Log("got1=", got1)
	t.Log("got", "suc", got.Data.Suc, "fail", got.Data.Fail) // got suc [{1 0xd87a8b21072bdf721b83ec6f860167cb336c38283bff8ce20f173124d4163b90}] fail []

}

func TestProofApiService_V1ProofGetProofsPost(t *testing.T) {
	config := NewConfiguration()
	config.BasePath = "http://183.134.99.140:9992"

	client := NewAPIClient(config)

	data := ProofGetProofsBody{
		Id:     1,
		Method: "string",
		Params: []RpcutilsHashes{
			{
				Hash: []string{
					"0xd87a8b21072bdf721b83ec6f860167cb336c38283bff8ce20f173124d4163b90",
				},
			},
		},
	}
	got, got1, err := client.ProofApi.V1ProofGetProofsPost(context.Background(), data)
	if err != nil {
		t.Log("err=", err)
		return
	}
	t.Log("got1=", got1)
	t.Log("got=", got)
}



func TestProofApiService_V1ProofListPost(t *testing.T) {
	config := NewConfiguration()
	config.BasePath = "http://183.134.99.140:9992"

	client := NewAPIClient(config)

	productName := getInterfaceString("牡丹纹双龙耳盖瓶（白玉）")
	data := ProofListBody{
		Id: 1,
		Method: "string",
		Params: []SwaggerQuery{
			{
				Match: []SwaggerQMatch{
					{
						Key: "产品名称",
						Value: &productName,
					},
				},
			},
		},
	}
	got, got1, err := client.ProofApi.V1ProofListPost(context.Background(), data)
	if err != nil {
		t.Log("err=", err)
		return
	}
	t.Log("got1=", got1)
	t.Log("got=", got)
}