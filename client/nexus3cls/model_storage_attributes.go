/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StorageAttributes struct {
	// Blob store used to store repository contents
	BlobStoreName string `json:"blobStoreName"`
	// Whether to validate uploaded content's MIME type appropriate for the repository format
	StrictContentTypeValidation bool `json:"strictContentTypeValidation"`
}
