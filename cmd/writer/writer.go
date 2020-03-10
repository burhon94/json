package writer

import "encoding/json"

func JsonWrite(dataStruct interface{}) (encoded[]byte, err error) {
	encoded, err = json.Marshal(dataStruct)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func JsonFileUpload(path, fileName string) (encoded string, err error) {
	marshal, err := json.Marshal(path+fileName)
	if err != nil {
		return "", err
	}
	encoded = string(marshal)
	return encoded, nil
}