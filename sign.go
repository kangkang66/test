package main

import (
	"log"
	"net/http"
)

func Sign()  {
	path := "https://mini.tuodan.tech/jstd-doger/app/sign/v1/signIn"
	content,err := sendRequest(http.MethodGet,path,"")
	log.Println(string(content),err)
}
