package common

import "fmt"

func GetString(data map[string]interface{}, key string) string {
	if value, ok := data[key].(string); ok {
		return value
	}
	return ""
}

func ProcessAttributes(data map[string]interface{}, keyPrefix, valuePrefix, typePrefix string) map[string]interface{} {
	result := make(map[string]interface{})

	for i := 1; ; i++ {
		key := fmt.Sprintf("%s%d", keyPrefix, i)
		valueKey := fmt.Sprintf("%s%d", valuePrefix, i)
		typeKey := fmt.Sprintf("%s%d", typePrefix, i)

		if attrKey, ok := data[key].(string); ok {
			attrVal := GetString(data, valueKey)
			attrType := GetString(data, typeKey)

			result[attrKey] = map[string]interface{}{
				"value": attrVal,
				"type":  attrType,
			}
		} else {
			break
		}

	}
	return result
}

// convertData converts the received data into the desired format
func ConvertData(data map[string]interface{}) map[string]interface{} {
	converted := map[string]interface{}{
		"event":            GetString(data, "ev"),
		"event_type":       GetString(data, "et"),
		"app_id":           GetString(data, "id"),
		"user_id":          GetString(data, "uid"),
		"message_id":       GetString(data, "mid"),
		"page_title":       GetString(data, "t"),
		"page_url":         GetString(data, "p"),
		"browser_language": GetString(data, "l"),
		"screen_size":      GetString(data, "sc"),
		"attributes":       ProcessAttributes(data, "atrk", "atrv", "atrt"),
		"traits":           ProcessAttributes(data, "uatrk", "uatrv", "uatrt"),
	}
	return converted
}
