package utils

// DefaultNamespace is the default namespace of the Ververica Platform
const DefaultNamespace = "default"

// GetNamespaceOrDefault will return either the given namespace, if valid, or the default Ververica Platform namespace
func GetNamespaceOrDefault(namespace string) string {
	if len(namespace) == 0 {
		return DefaultNamespace
	}
	return namespace
}