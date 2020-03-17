package writer

import (
	"encoding/json"
	"errors"
	"github.com/burhon94/json/cmd/dirFileReader"
	"net/http"
	"path/filepath"
)

type filePath struct {
	Path string `json:"path"`
	FileName string `json:"fileName"`
}

func JsonWrite(dataStruct interface{}) (encoded []byte, err error) {
	encoded, err = json.Marshal(dataStruct)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func WriteJSONHTTP(response http.ResponseWriter, dto interface{}) (err error) {
	response.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(dto)
	if err != nil {
		return errors.New("can't encode to json")
	}

	_, err = response.Write(bytes)
	if err != nil {
		return errors.New("can't response to request")
	}

	return nil
}

func JsonFileUpload(path string) (encoded string, err error) {
	fileStruct := make([]filePath, 0)

	files, err := dirFileReader.DirFileReader(path)
	if err != nil {
		return "Error while reading Path directory", err
	}

	for _, file := range files {
		PathFile := filepath.Dir(file)
		fileName := filepath.Base(file)
		fileStruct = append(fileStruct, filePath{
			Path:     PathFile,
			FileName: fileName,
		})
	}

	marshal, err := json.Marshal(fileStruct)
	if err != nil {
		return "", err
	}
	encoded = string(marshal)
	return encoded, nil
}
