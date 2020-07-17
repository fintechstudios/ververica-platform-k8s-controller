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

import core "k8s.io/api/core/v1"

type EnvVar struct {
	Name      string    `json:"name,omitempty"`
	Value     string    `json:"value,omitempty"`
	ValueFrom *core.EnvVarSource `json:"valueFrom,omitempty"`
}
