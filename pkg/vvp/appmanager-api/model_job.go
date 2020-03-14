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

type Job struct {
	ApiVersion string       `json:"apiVersion,omitempty"`
	Kind       string       `json:"kind,omitempty"`
	Metadata   *JobMetadata `json:"metadata,omitempty"`
	Spec       *JobSpec     `json:"spec,omitempty"`
	Status     *JobStatus   `json:"status,omitempty"`
}
