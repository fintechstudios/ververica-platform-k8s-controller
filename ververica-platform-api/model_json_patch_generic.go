/*
 * AppManager API
 *
 * HTTP REST API to connect to the AppManager
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ververicaplatformapi

// Since Swagger 2 doesn't have oneOf support, this definition should fit all JSON Patch objects.
type JsonPatchGeneric struct {
	Op string `json:"op"`
	Path string `json:"path"`
	Value *Any `json:"value,omitempty"`
	From string `json:"from,omitempty"`
}
