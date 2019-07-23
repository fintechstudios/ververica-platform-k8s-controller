/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

type ResourceQuota struct {
	Kind string `json:"kind,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Spec *ResourceQuotaSpec `json:"spec,omitempty"`
	Status *ResourceQuotaStatus `json:"status,omitempty"`
}
