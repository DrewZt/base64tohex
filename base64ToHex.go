package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
)

//decode base64 to hex
func decodeInput(decode string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decode)
	if err != nil {
		return "", err
	}
	output := hex.EncodeToString(decodeBytes)
	return output, nil
}

//encode hex to base64
func encodeInput(encode string) (string, error) {
	input, err := hex.DecodeString(encode)
	if err != nil {
		return "", err
	}
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString, nil
}

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	str := req.RequestURI
	if str == "/facicon.ico" {
		return
	}
	var content string
	if str[1:3] == "0x" {
		content = str[3:]
	} else {
		content = str[1:]
	}
	var length int
	if len(content) < 32 {
		length = len(content)
	} else {
		length = 32
	}
	var hexType bool
	for i := 0; i < length; i++ {
		// if !hex use decodeInput
		if (content[i] < 97 || content[i] > 102) && (content[i] < 48 || content[i] > 57) {
			hexType = false
			break
		}
		hexType = true
	}
	var a string
	var err error
	if hexType == true {
		a, err = encodeInput(content)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		a, err = decodeInput(content)
		if err != nil {
			fmt.Println(err)
		}
	}
	_, err = fmt.Fprint(rw, a)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	http.HandleFunc("/", IndexHandler)
	_ = http.ListenAndServe("0.0.0.0:8000", nil)

}
