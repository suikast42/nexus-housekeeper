/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NpmGroupRepositoryApiRequest struct {
	// A unique identifier for this repository
	Name string `json:"name"`
	// Whether this repository accepts incoming requests
	Online bool `json:"online"`
	Storage *StorageAttributes `json:"storage"`
	Group *GroupDeployAttributes `json:"group"`
}
