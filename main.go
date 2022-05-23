package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetRequest() {
	resp, _ := http.Get("http://127.0.0.1:8000/register")
	body, _ := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	html := string(body)
	tokenIndex := strings.Index(html, "name=\"_token\" value=\"")
	fmt.Println(html[tokenIndex+21 : tokenIndex+61])
}
func PostRequest() {
	resp, _ := http.PostForm("http://127.0.0.1:8000/register", url.Values{"_token": {}, "name": {"角田篤dsdsa思"}, "email": {"aaoaoao23a@a.com"}, "password": {"ajajajajdda"}, "password_confirmation": {"ajajajajdda"}})
	body, _ := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	println(string(body))
	fmt.Println(strings.Index(string(body), ""))
}

func main() {
	PostRequest()
}
