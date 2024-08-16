package data

import (
	"encoding/json"
	"os"
)

type PathLike string

type DataSourcer interface {
	Read() ([]byte, error)
	Write(any) (bool, error)
}

type JSONDataSource struct {
	dataSourcePath string
}

func NewJSONDataSource(dataSourcePath PathLike) *JSONDataSource {
	return &JSONDataSource{
		dataSourcePath: string(dataSourcePath),
	}
}

func (jds *JSONDataSource) Read() ([]byte, error) {
	bytes, err := os.ReadFile(jds.dataSourcePath)

	return bytes, err
}

func (jds *JSONDataSource) Write(data any) (bool, error) {
	bytes, marshalErr := json.MarshalIndent(data, "", "\t")

	if marshalErr != nil {
		return false, marshalErr
	}

	writeErr := os.WriteFile(jds.dataSourcePath, bytes, 0644)

	if writeErr != nil {
		return false, writeErr
	}

	return true, nil
}
