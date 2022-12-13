package models

import "github.com/speps/go-hashids"

var HashID *hashids.HashID

// init the hashid
func init() {
	hd := hashids.NewData()
	hd.Salt = "a5fjk3954djkd6v"
	HashID, _ = hashids.NewWithData(hd)
}
