package internal

import (
	"encoding/json"
	"log"
)

func Marshall(error *string, data interface{}, success bool) []byte {
	jdata, err := json.Marshal(&ApiResponse{
		Error:   error,
		Data:    data,
		Success: success,
	})

	if err != nil {
		log.Fatal("Error marshalling api response ", err.Error())
	}

	return jdata
}
