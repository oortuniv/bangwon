package global

import (
	_type "bangwon/type"
	"sync"
)

var MeId string
var bangwons *sync.Map

func init() {
	bangwons = &sync.Map{}
}

func Bangwons() *sync.Map {
	return bangwons
}

func GetMe() *_type.Status {
	load, ok := bangwons.Load(MeId)
	if !ok {
		return nil
	}
	me := load.(_type.Status)
	return &me
}

func SetMe(status _type.Status) *_type.Status {
	bangwons.Store(MeId, status)
	load, _ := bangwons.Load(MeId)
	me := load.(_type.Status)
	return &me
}

func GetBangwon(id string) *_type.Status {
	load, ok := bangwons.Load(id)
	if !ok {
		return nil
	}
	bangwon := load.(_type.Status)
	return &bangwon
}

func UpdateBangwon(id string, status _type.Status) *_type.Status {
	bangwons.Store(id, status)
	load, _ := bangwons.Load(id)
	bangwon := load.(_type.Status)
	return &bangwon
}
