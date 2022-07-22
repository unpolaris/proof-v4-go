/*
 * 存证溯源接口
 *
 * 尽量通用的存证溯源平台，目前是简单用户体系。
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModelProofExtInfo struct {
	Ext string `json:"ext,omitempty"`
	ExtDetail *ModelProofExtDetail `json:"ext_detail,omitempty"`
	IsIncrement bool `json:"is_increment,omitempty"`
	Note string `json:"note,omitempty"`
	NoteDetail *ModelProofNoteDetail `json:"note_detail,omitempty"`
	// 自定义私钥
	Prikey string `json:"prikey,omitempty"`
	TemplateId int32 `json:"template_id,omitempty"`
	Version string `json:"version,omitempty"`
}