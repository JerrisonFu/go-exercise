package stdlib

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

func StrconvExamples() (string, error) {
	i, err := strconv.Atoi("123")
	if err != nil {
		return "", err
	}

	s := strconv.Itoa(456)

	b, err := strconv.ParseBool("true")
	if err != nil {
		return "", err
	}

	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("int: %d, str: %s, bool: %v, float: %.2f", i, s, b, f), nil
}

func BytesOperations() bool {
	b1 := []byte("hello")
	b2 := []byte("hello")
	b3 := []byte("world")

	return bytes.Equal(b1, b2) && !bytes.Equal(b1, b3)
}

func RegexpExamples() bool {
	matched, _ := regexp.MatchString(`^\d{3}-\d{4}$`, "123-4567")
	return matched
}

func BufferExample() string {
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteByte(' ')
	buf.Write([]byte("World"))

	return buf.String()
}
