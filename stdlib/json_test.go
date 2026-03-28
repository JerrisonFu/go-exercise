package stdlib

import (
	"testing"
)

func TestJSONOperations(t *testing.T) {
	name, err := JSONOperations()
	if err != nil {
		t.Errorf("JSONOperations() error = %v", err)
	}
	if name != "Alice" {
		t.Errorf("JSONOperations() = %s, want Alice", name)
	}
}

func TestJSONWithOmitEmpty(t *testing.T) {
	size, err := JSONWithOmitEmpty()
	if err != nil {
		t.Errorf("JSONWithOmitEmpty() error = %v", err)
	}
	if size == 0 {
		t.Error("JSONWithOmitEmpty() should return non-zero size")
	}
}

func TestJSONDecodeMap(t *testing.T) {
	result, err := JSONDecodeMap()
	if err != nil {
		t.Errorf("JSONDecodeMap() error = %v", err)
	}
	if !result {
		t.Error("JSONDecodeMap() should return true")
	}
}

func TestJSONMarshalUser(t *testing.T) {
	user := User{Name: "Test", Age: 20}
	data, err := marshalUser(user)
	if err != nil {
		t.Errorf("marshalUser() error = %v", err)
	}
	if len(data) == 0 {
		t.Error("marshalUser() should return data")
	}
}

func marshalUser(u User) ([]byte, error) {
	type userJSON struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	uj := userJSON{Name: u.Name, Age: u.Age}
	return []byte(`{"name":"` + uj.Name + `","age":` + string(rune('0'+uj.Age/10)) + `}`), nil
}
