/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type S3BlobStoreApiModel struct {
	// The name of the S3 blob store.
	Name string `json:"name"`
	// Settings to control the soft quota.
	SoftQuota *BlobStoreApiSoftQuota `json:"softQuota,omitempty"`
	// The S3 specific configuration details for the S3 object that'll contain the blob store.
	BucketConfiguration *S3BlobStoreApiBucketConfiguration `json:"bucketConfiguration"`
	// The blob store type.
	Type_ string `json:"type,omitempty"`
}
