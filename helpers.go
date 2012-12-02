package imgo

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func UploadFile(path string) (ul Upload, err error) {

	f, err := os.Open(path)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	form := url.Values{}
	form.Add("key", APIKey)
	form.Add("image", base64.StdEncoding.EncodeToString(b))
	resp, err := http.PostForm(UploadURL, form)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	v := Upload{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
	}
	return v, nil
}
