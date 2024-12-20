/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type StackTraceElement struct {
	MethodName string `json:"methodName,omitempty"`
	FileName string `json:"fileName,omitempty"`
	LineNumber int32 `json:"lineNumber,omitempty"`
	ClassName string `json:"className,omitempty"`
	NativeMethod bool `json:"nativeMethod,omitempty"`
}
