package writer

import "encoding/json"

func JsonWrite(dataStruct interface{}) (encoded[]byte, err error) {
	encoded, err = json.Marshal(dataStruct)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}