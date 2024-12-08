//go:build windows

package openuem_utils

import "golang.org/x/sys/windows/registry"

func OpenRegistryForQuery(key registry.Key, path string) (registry.Key, error) {
	return registry.OpenKey(key, path, registry.QUERY_VALUE)
}

func GetValueFromRegistry(k registry.Key, key string) (string, error) {
	s, _, err := k.GetStringValue(key)
	if err != nil {
		return "", err
	}

	return s, nil
}
