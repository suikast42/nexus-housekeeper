/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type PageComponentXo struct {
	Items []ComponentXo `json:"items,omitempty"`
	ContinuationToken string `json:"continuationToken,omitempty"`
}