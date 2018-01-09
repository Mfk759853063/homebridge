package httpUtil

import (
	"net/url"
	"strings"
	"io/ioutil"
	"net/http"
)

func DoPost(url string,parms *url.Values, contentType string) ([]byte, error) {
	parmsStr :=  parms.Encode()
	resp,err := http.Post(url,contentType,strings.NewReader(parmsStr))
	defer resp.Body.Close()
	if err != nil {
		return nil,err
	}
	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}

	return body,nil
}

func DoGet(url string, contentType string) ([]byte, error) {

	responseData,error:= http.Get(url)
	if error != nil {
		return nil,error
	}
	defer responseData.Body.Close()
	body,error:= ioutil.ReadAll(responseData.Body)
	if error != nil {
		return nil,error
	}

	return body,nil
}
