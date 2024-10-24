/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Throwable struct {
	Cause *Throwable `json:"cause,omitempty"`
	StackTrace []StackTraceElement `json:"stackTrace,omitempty"`
	Message string `json:"message,omitempty"`
	LocalizedMessage string `json:"localizedMessage,omitempty"`
	Suppressed []Throwable `json:"suppressed,omitempty"`
}
