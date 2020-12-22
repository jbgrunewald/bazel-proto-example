package config

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestJSONSerialization(t *testing.T) {
	configJson := `{"id": "something", "type": "major"}`
	configInBytes := []byte(configJson)
	carouselConfig := &Config{}

	if err := json.Unmarshal(configInBytes, carouselConfig); err != nil {
		t.Errorf("Unable to unmarshal the example json: %s", err)
	}

	result, err := json.Marshal(carouselConfig)
	if err != nil {
		t.Errorf("Unable to marshal example protobuf type")
	}

	if !bytes.Equal(result, configInBytes) {
		t.Errorf("Marshalled result %s not equal to expected %s", result, configInBytes)
	}

}
