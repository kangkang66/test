package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetVisitors()  {
	path := "https://mini.tuodan.tech/jstd-doger/app/user/visitors"
	data := `{"before":1602149710438,"size":10}`
	content,err := sendRequest(http.MethodPost, path, data)
	if err != nil {
		return
	}
	resp := new(ResponseVistors)
	err = json.Unmarshal(content, resp)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(resp)
}
