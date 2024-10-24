/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type ApiPrivilege struct {
	// The type of privilege, each type covers different portions of the system. External values supplied to this will be ignored by the system.
	Type_ string `json:"type,omitempty"`
	// The name of the privilege.  This value cannot be changed.
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// Indicates whether the privilege can be changed. External values supplied to this will be ignored by the system.
	ReadOnly bool `json:"readOnly,omitempty"`
}