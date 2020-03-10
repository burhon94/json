package writer

import (
	"encoding/json"
	"io"
)

func JsonWrite(dataStruct interface{}) (encoded[]byte, err error) {
	encoded, err = json.Marshal(dataStruct)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func JsonFileUpload(fileName io.Reader) (encoded string, err error) {
	marshal, err := json.Marshal(fileName)
	if err != nil {
		return "", err
	}
	encoded = string(marshal)
	return encoded, nil
}