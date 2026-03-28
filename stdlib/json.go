package stdlib

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func JSONOperations() (string, error) {
	user := User{Name: "Alice", Age: 30}

	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	var decoded User
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		return "", err
	}

	return decoded.Name, nil
}

func JSONWithOmitEmpty() (int, error) {
	user := User{Name: "Bob"}

	data, err := json.Marshal(user)
	if err != nil {
		return len(data), err
	}

	return len(data), nil
}

func JSONDecodeMap() (bool, error) {
	jsonStr := `{"name":"Charlie","score":95}`
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return false, err
	}

	return result["name"] == "Charlie", nil
}

func fmtExample() string {
	name := "Go"
	version := 1.21
	return fmt.Sprintf("Language: %s, Version: %.2f", name, version)
}
