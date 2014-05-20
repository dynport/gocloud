package main

import (
	"net/http"

	"github.com/dynport/martini"
)

func listVolumes(w http.ResponseWriter) (int, []byte) {
	volumes, e := describeVolumes()
	if e != nil {
		return returnError(e)
	}
	return renderJson(w, volumes)
}

func showVolume(w http.ResponseWriter, params martini.Params) (int, []byte) {
	volumes, e := describeVolumes()
	if e != nil {
		return returnError(e)
	}
	for _, v := range volumes {
		if v.VolumeId == params["id"] {
			return renderJson(w, v)
		}
	}
	return 404, nil
}
