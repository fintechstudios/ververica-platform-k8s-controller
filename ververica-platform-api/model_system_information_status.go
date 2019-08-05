/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

type SystemInformationStatus struct {
	JvmVersion string `json:"jvmVersion,omitempty"`
	RevisionInformation *RevisionInformation `json:"revisionInformation,omitempty"`
	ResourceQuota *ResourceQuota `json:"resourceQuota,omitempty"`
	License *License `json:"license,omitempty"`
}
