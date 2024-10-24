/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type SupportZipGeneratorRequest struct {
	SystemInformation bool `json:"systemInformation,omitempty"`
	ThreadDump bool `json:"threadDump,omitempty"`
	Metrics bool `json:"metrics,omitempty"`
	Configuration bool `json:"configuration,omitempty"`
	Security bool `json:"security,omitempty"`
	Log bool `json:"log,omitempty"`
	TaskLog bool `json:"taskLog,omitempty"`
	AuditLog bool `json:"auditLog,omitempty"`
	Jmx bool `json:"jmx,omitempty"`
	Replication bool `json:"replication,omitempty"`
	LimitFileSizes bool `json:"limitFileSizes,omitempty"`
	LimitZipSize bool `json:"limitZipSize,omitempty"`
	Hostname string `json:"hostname,omitempty"`
}
