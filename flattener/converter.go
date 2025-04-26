package flattener

import "strings"

func Flatten(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	flattenRecursive(data, []string{}, result)
	return result
}

func flattenRecursive(data interface{}, currentKeys []string, result map[string]interface{}) {
	value, ok := data.(map[string]interface{})

	if ok {
		for k, v := range value {
			keys := append(append([]string(nil), currentKeys...), k)
			flattenRecursive(v, keys, result)
		}
		return
	}

	resultKey := createKey(currentKeys)
	result[resultKey] = data
}

func createKey(keys []string) string {
	return strings.Join(keys, "__")
}
