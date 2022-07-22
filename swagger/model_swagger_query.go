/*
 * 存证溯源接口
 *
 * 尽量通用的存证溯源平台，目前是简单用户体系。
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SwaggerQuery struct {
	Fetch *SwaggerQFetch `json:"fetch,omitempty"`
	// 过滤
	Filter []SwaggerQMatch `json:"filter,omitempty"`
	// 且匹配
	Match []SwaggerQMatch `json:"match,omitempty"`
	// 或匹配
	MatchOne []SwaggerQMatch `json:"match_one,omitempty"`
	// 多字段匹配
	MultiMatch []SwaggerQMultiMatch `json:"multi_match,omitempty"`
	// 非匹配
	Not []SwaggerQMatch `json:"not,omitempty"`
	Page *SwaggerQPage `json:"page,omitempty"`
	// 范围
	Range_ []SwaggerQRange `json:"range,omitempty"`
	Size *SwaggerQSize `json:"size,omitempty"`
	// 排序
	Sort []SwaggerQSort `json:"sort,omitempty"`
}
