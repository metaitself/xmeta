package json

import "github.com/goccy/go-json"

var (
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)

func MarshalToString(v any) string {
	marshal, err := Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(marshal)
}

func MarshalToStringE(v any) (string, error) {
	marshal, err := Marshal(v)
	return string(marshal), err
}

func MarshalToByte(v any) []byte {
	marshal, err := Marshal(v)
	if err != nil {
		return []byte("{}")
	}
	return marshal
}

func UnmarshalFromStringE(buf string, v any) error {
	return Unmarshal([]byte(buf), v)
}
