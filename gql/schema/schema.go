package schema

import (
	"bytes"
	"fmt"
)

//GetRootSchema get schema definition from bindata
func GetRootSchema() (string, error) {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b, err := Asset(name)
		if err != nil {
			return "", fmt.Errorf("Asset: %s : %s", name, err.Error())
		}
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String(), nil
}
