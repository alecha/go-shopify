package shopify

import "encoding/json"

// getJSONBytesFromMap Extracts Json Bytes from map[string]interface
func getJSONBytesFromMap(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func unmarshal(responseData []byte, responseErrors []error, output interface{}) []error {
	if len(responseErrors) > 0 {
		return responseErrors
	}
	if err := json.Unmarshal(responseData, output); err != nil {
		return []error{err}
	}
	return nil
}
