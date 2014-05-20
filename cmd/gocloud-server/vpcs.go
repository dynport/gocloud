package main

import (
	"fmt"
	"net/http"

	"github.com/dynport/martini"
)

func vpcsListSubnets(w http.ResponseWriter, params martini.Params) (int, []byte) {
	nets, e := describeSubnets()
	if e != nil {
		return returnError(e)
	}
	filtered := nets.Query(params["id"])
	return renderJson(w, filtered)
}

func vpcsShow(w http.ResponseWriter, params martini.Params) (int, []byte) {
	all, e := describeVpcs()
	if e != nil {
		return returnError(e)
	}
	filtered := all.Search(params["id"])
	if len(filtered) == 0 {
		return 404, nil
	}
	if len(filtered) > 1 {
		return returnError(fmt.Errorf("expected to find 1 vpc, found %d", len(filtered)))
	}
	return renderJson(w, filtered.First())
}

type VpcList []*Vpc

func (list VpcList) Search(q string) VpcList {
	filtered := VpcList{}
	for _, v := range list {
		if v.Id == q {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (list VpcList) First() *Vpc {
	if len(list) > 0 {
		return list[0]
	}
	return nil
}

func vpcsList(w http.ResponseWriter) (int, []byte) {
	vpcs, e := describeVpcs()
	if e != nil {
		return returnError(e)
	}
	return renderJson(w, vpcs)
}
