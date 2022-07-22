/*
 * 存证溯源接口
 *
 * 尽量通用的存证溯源平台，目前是简单用户体系。
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModelResponse struct {
	// 状态码
	Code int32 `json:"code,omitempty"`
	// 消息
	Msg string `json:"msg,omitempty"`
	// 数据
	Data *interface{} `json:"data,omitempty"`
}
