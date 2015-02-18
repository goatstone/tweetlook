package data

import (
	"log"
	"appengine"
	"appengine/datastore"
	"math/rand"
	"time"
	"appengine/user"

)

type appLog struct {
	Name      string
	Id        int
	TimeStamp time.Time
	Account   string
}
type TemplateData struct {
	Title      string
	Legend     string
	Inputs     []SiteProp
	Message    string
	AHref      string
}
type SiteProp struct {
	Name          string
	Value         string
	FormLabel     string
	InputType     string
	Disabled      string
}

func StoreLog(ctx appengine.Context, name string) {
	var entityKind = "appLog"
	dataKey := datastore.NewIncompleteKey(ctx, entityKind, nil)
	alog := &appLog{}
	alog.Name = name
	alog.Id = rand.Intn(1000)
	alog.TimeStamp = time.Now()
	alog.Account = user.Current(ctx).String()
	if _, err := datastore.Put(ctx, dataKey, alog); err != nil {
		log.Print("err:  ", err)
		return
	}
}
func UpdateSiteProp(ctx appengine.Context, siteProp SiteProp) (err error) {
	log.Print("updateSiteProp  :  ")
	dataKey := datastore.NewKey(ctx, "SiteProperties", siteProp.Name, 0, nil)
	if _, err := datastore.Put(ctx, dataKey, &siteProp); err != nil {
		log.Print("err::  ", err)
	}
	return
}
func AddSiteProp(ctx appengine.Context, prop map[string]string) (err error) {
	keyName := prop["Name"]
	sp := &SiteProp{Disabled:""}
	sp.Name = prop[ "Name"]
	sp.Value = prop["Value"]
	sp.FormLabel = prop["Name"]
	dataKey := datastore.NewKey(ctx, "SiteProperties", keyName, 0, nil)
	if _, err := datastore.Put(ctx, dataKey, sp); err != nil {
		log.Print("err::  ", err)
	}
	return
}
func GetSiteProps(ctx appengine.Context) (siteProps []SiteProp , err error) {
	siteProps = []SiteProp{}
	q := datastore.NewQuery("SiteProperties")
	_, err = q.GetAll(ctx, &siteProps)
	if err != nil {
		log.Print("ERROR :  ", err)
		return
	}
	return
}
