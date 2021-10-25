package metadata

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// metadata houses metadata to be written out to the additional-metadata.json
type metadata struct {
	AddonInstalled bool `json:"addon-installed,string"`
}

// Instance is the singleton instance of metadata.
var Instance = metadata{}

// WriteToJSON will marshall the metadata struct and write it into the given file.
func (m *metadata) WriteToJSON(outputFilename string) (err error) {
	var data []byte
	if data, err = json.Marshal(m); err != nil {
		return err
	}

	if err = ioutil.WriteFile(outputFilename, data, os.FileMode(0644)); err != nil {
		return err
	}

	return nil
}
