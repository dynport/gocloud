package digitalocean

type ImagesResponse struct {
	Images []*Image `json:"images,omitempty"`
	Meta   *Meta    `json:"meta,omitempty"`
}

type ImageResponse struct {
	Image *Image `json:"image,omitempty"`
}

func (client *Client) Images() (*ImagesResponse, error) {
	rsp := &ImagesResponse{}
	e := client.loadResponse("/v2/images", rsp)
	return rsp, e
}

func (client *Client) Image(idOrSlug string) (*ImageResponse, error) {
	rsp := &ImageResponse{}
	e := client.loadResponse("/v2/images/"+idOrSlug, rsp)
	return rsp, e
}
