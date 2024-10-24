/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type ReplicationAttributes struct {
	// Whether pre-emptive pull is enabled
	PreemptivePullEnabled bool `json:"preemptivePullEnabled"`
	// Regular Expression of Asset Paths to pull pre-emptively pull
	AssetPathRegex string `json:"assetPathRegex,omitempty"`
}