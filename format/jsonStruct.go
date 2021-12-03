package format

import (
	"bytes"
	"encoding/json"
	"io"
	"regexp"
)

// PrettyStruct : JSON pretty print by marshaling value
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {

		return "", err
	}

	return string(val), nil
}

// PrettyEncode : JSON pretty print by encoding value
func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {

		return err
	}

	return nil
}

// PrettyString : Pretty print JSON string
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}

// TrimSpaceDashInString : Trim all dash in string
func TrimSpaceDashInString(s string) string {
	re := regexp.MustCompile(`\s+`)

	return re.ReplaceAllString(s, "-")
}
