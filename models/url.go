package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Url struct {
	Id        int    `bson:"_id"`
	SourceUrl string `bson:"SourceUrl"`
	ShortUrl  string `bson:"-"`
}

var (
	URL_COLLECTION = "Url"
)

func (url *Url) Insert() error {
	GetDB().C(URL_COLLECTION).Insert(url)
}
func (url *Url) FindById() error {
	return GetDB().C(URL_COLLECTION).FindId(url.Id).One(url)
}
func (url *Url) Update() error {
	GetDB().C(URL_COLLECTION).Update(bson.M{"_id": url.Id}, url)
}
func (url *Url) Upsert() error {
	GetDB().C(URL_COLLECTION).Upsert(bson.M{"_id": url.Id}, url)
}
