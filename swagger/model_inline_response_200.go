/*
 * 存证溯源接口
 *
 * 尽量通用的存证溯源平台，目前是简单用户体系。
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse200 struct {
	Data *ModelResAutoProof `json:"data,omitempty"`
	ModelResponse *ModelResponse `json:"model.Response,omitempty"`
}
