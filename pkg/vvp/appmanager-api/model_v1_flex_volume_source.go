/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.4.3
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
type V1FlexVolumeSource struct {
	// Driver is the name of the driver to use for this volume.
	Driver string `json:"driver"`
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". The default filesystem depends on FlexVolume script.
	FsType string `json:"fsType,omitempty"`
	// Optional: Extra command options if any.
	Options map[string]string `json:"options,omitempty"`
	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly  bool                    `json:"readOnly,omitempty"`
	SecretRef *V1LocalObjectReference `json:"secretRef,omitempty"`
}