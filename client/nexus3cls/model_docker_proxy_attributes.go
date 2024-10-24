/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DockerProxyAttributes struct {
	// Type of Docker Index
	IndexType string `json:"indexType,omitempty"`
	// Url of Docker Index to use
	IndexUrl string `json:"indexUrl,omitempty"`
	// Allow Nexus Repository Manager to download and cache foreign layers
	CacheForeignLayers bool `json:"cacheForeignLayers,omitempty"`
	// Regular expressions used to identify URLs that are allowed for foreign layer requests
	ForeignLayerUrlWhitelist []string `json:"foreignLayerUrlWhitelist,omitempty"`
}
