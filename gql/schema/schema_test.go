package schema

import (
	"testing"
)

func TestGetRootSchema(t *testing.T) {
	schema, err := GetRootSchema()
	if err != nil {
		t.Fatalf("failed when get schema: %s", err.Error())
	}

	if schema == "" {
		t.Fatal("schema should not empty string")
	}
}
