package imgo

import (
	"encoding/xml"
)

const (
	UploadURL = "http://api.imgur.com/2/upload.xml"
)

type ImgurImage struct {
	XMLName    xml.Name `xml:"image"`
	Name       string   `xml:"name"`
	Title      string   `xml:"title"`
	Caption    string   `xml:"caption"`
	Hash       string   `xml:"hash"`
	DeleteHash string   `xml:"deletehash"`
	DateTime   string   `xml:"datetime"`
	Type       string   `xml:"type"`
	Animated   bool     `xml:"animated"`
	Width      int64    `xml:"width"`
	Height     int64    `xml:"height"`
	Size       int64    `xml:"size"`
	Views      int64    `xml:"views"`
	Bandwidth  int64    `xml:"bandwidth"`
}

type ImgurLinks struct {
	XMLName     xml.Name `xml:"links"`
	Original    string   `xml:"original"`
	Page        string   `xml:"imgur_page"`
	Delete      string   `xml:"delete_page"`
	SmallSquare string   `xml:"small_square"`
	Thumbnail   string   `xml:"large_thumbnail"`
}

type Stats struct {
	XMLName        xml.Name `xml:"stats"`
	MostPopular    []string `xml:"most_popular_images>image_hash"`
	ImagesUploaded int64    `xml:"images_uploaded"`
	ImagesViewed   int64    `xml:"images_viewed"`
	BandwidthUsed  string   `xml:"bandwidth_used"`
	AvgImgSize     string   `xml:"average_image_size"`
}

type Upload struct {
	XMLName xml.Name `xml:"upload"`
	Image   ImgurImage
	Links   ImgurLinks
}
