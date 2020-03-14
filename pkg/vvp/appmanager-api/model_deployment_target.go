/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type DeploymentTarget struct {
	ApiVersion string                    `json:"apiVersion,omitempty"`
	Kind       string                    `json:"kind,omitempty"`
	Metadata   *DeploymentTargetMetadata `json:"metadata,omitempty"`
	Spec       *DeploymentTargetSpec     `json:"spec,omitempty"`
}
