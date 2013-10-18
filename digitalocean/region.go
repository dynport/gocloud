package digitalocean

type RegionResponse struct {
	Status  string    `json:"status"`
	Regions []*Region `json:"regions"`
}

type Region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (self *Account) Regions() (regions []*Region, e error) {
	regionResponse := &RegionResponse{}
	e = self.loadResource("/regions", regionResponse)
	regions = regionResponse.Regions
	return
}
