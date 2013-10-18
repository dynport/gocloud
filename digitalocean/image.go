package digitalocean

import (
	"strconv"
)

type Image struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
}

type ImagesReponse struct {
	Status string   `json:"status"`
	Images []*Image `json:"images"`
}

type ImageResponse struct {
	Status string `json:"status"`
	Image  *Image `json:"image"`
}

func (self *Account) Images() (images []*Image, e error) {
	imagesReponse := &ImagesReponse{}
	e = self.loadResource("/images", imagesReponse)
	if e != nil {
		return
	}
	images = imagesReponse.Images
	return
}

func (self *Account) GetImage(id int) (image *Image, e error) {
	imageReponse := &ImageResponse{}
	e = self.loadResource("/images/"+strconv.Itoa(id), imageReponse)
	image = imageReponse.Image
	return
}
