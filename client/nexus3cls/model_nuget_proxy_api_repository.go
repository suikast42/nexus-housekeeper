/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NugetProxyApiRepository struct {
	// A unique identifier for this repository
	Name string `json:"name,omitempty"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage *StorageAttributes `json:"storage"`
	Cleanup *CleanupPolicyAttributes `json:"cleanup,omitempty"`
	Proxy *ProxyAttributes `json:"proxy"`
	NegativeCache *NegativeCacheAttributes `json:"negativeCache"`
	HttpClient *HttpClientAttributes `json:"httpClient"`
	// The name of the routing rule assigned to this repository
	RoutingRuleName string `json:"routingRuleName,omitempty"`
	Replication *ReplicationAttributes `json:"replication,omitempty"`
	NugetProxy *NugetAttributes `json:"nugetProxy"`
}
