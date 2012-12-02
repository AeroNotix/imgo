package imgo

import (
	"encoding/base64"
	"encoding/xml"
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
		return
	}
	return v, nil
}

func GetStats() (s Stats, err error) {
	body, err := GetResponse("http://api.imgur.com/2/stats")
	if err != nil {
		return
	}
	err = xml.Unmarshal(body, &s)
	if err != nil {
		return
	}
	return s, nil
}

func GetAlbum(AlbumID string) (a Album, err error) {
	body, err := GetResponse("http://api.imgur.com/2/album/" + AlbumID)
	if err != nil {
		return
	}
	err = xml.Unmarshal(body, &a)
	if err != nil {
		return
	}
	return a, nil
}

func GetInfoHash(Hash string) (i Image, err error) {
	body, err := GetResponse("http://api.imgur.com/2/image/" + Hash)
	if err != nil {
		return
	}
	err = xml.Unmarshal(body, &i)
	if err != nil {
		return
	}
	return i, nil
}

func GetResponse(URL string) (b []byte, err error) {
	resp, err := http.Get(URL)
	if err != nil {
		return
	}
	return ioutil.ReadAll(resp.Body)
}
