/*
 * Nexus Repository Manager REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.68.1-02
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nexus3

type ApiCreateUser struct {
	// The userid which is required for login. This value cannot be changed.
	UserId string `json:"userId,omitempty"`
	// The first name of the user.
	FirstName string `json:"firstName,omitempty"`
	// The last name of the user.
	LastName string `json:"lastName,omitempty"`
	// The email address associated with the user.
	EmailAddress string `json:"emailAddress,omitempty"`
	// The password for the new user.
	Password string `json:"password,omitempty"`
	// The user's status, e.g. active or disabled.
	Status string `json:"status"`
	// The roles which the user has been assigned within Nexus.
	Roles []string `json:"roles,omitempty"`
}