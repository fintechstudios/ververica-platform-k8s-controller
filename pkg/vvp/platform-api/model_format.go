/*
 * Ververica Platform API
 *
 * The Ververica Platform APIs, excluding Application Manager.
 *
 * API version: 2.4.3
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package platformapi

type Format struct {
	Dependencies []string   `json:"dependencies,omitempty"`
	Name         string     `json:"name,omitempty"`
	Packaged     bool       `json:"packaged,omitempty"`
	Properties   []Property `json:"properties,omitempty"`
	Sink         bool       `json:"sink,omitempty"`
	Source       bool       `json:"source,omitempty"`
	Type_        string     `json:"type,omitempty"`
}