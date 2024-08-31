package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type DataSourceJSON[M any] struct {
	path string
}

func NewDataSourceJSON[M any](path string) *DataSourceJSON[M] {
	return &DataSourceJSON[M]{
		path: string(path),
	}
}

func (dsj *DataSourceJSON[M]) Read(m *M) error {
	bytes, err := os.ReadFile(dsj.path)

	if err != nil {
		return fmt.Errorf("error reading data source at path %s", dsj.path)
	}

	unmarshalErr := json.Unmarshal(bytes, m)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	return nil
}

func (dsj *DataSourceJSON[M]) Write(m *M) (bool, error) {
	bytes, marshalErr := json.MarshalIndent(m, "", "\t")

	if marshalErr != nil {
		return false, marshalErr
	}

	writeErr := os.WriteFile(dsj.path, bytes, 0644)

	if writeErr != nil {
		return false, writeErr
	}

	return true, nil
}
