package reader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadJSONHTTP(request *http.Request, dto interface{}) (err error) {
	if request.Header.Get("Content-Type") != "application/json" {
		return errors.New(fmt.Sprintf("unsupported type for json: %v", err))
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("can't read request: %v", err))
	}
	defer func() {
		err := request.Body.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return errors.New(fmt.Sprintf("can't unmarshal json: %v", err))
	}
	return nil
}

func ReadJSON(body[]byte, dataStruct interface{}) error {
	err := json.Unmarshal(body, &dataStruct)
	if err != nil {
		return errors.New(fmt.Sprintf("can't unmarshal json: %v", err))
	}
	return nil
}