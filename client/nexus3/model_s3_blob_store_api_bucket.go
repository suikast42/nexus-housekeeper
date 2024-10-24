/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type S3BlobStoreApiBucket struct {
	// The AWS region to create a new S3 bucket in or an existing S3 bucket's region
	Region string `json:"region"`
	// The name of the S3 bucket
	Name string `json:"name"`
	// The S3 blob store (i.e S3 object) key prefix
	Prefix string `json:"prefix,omitempty"`
	// How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable)
	Expiration int32 `json:"expiration"`
}
