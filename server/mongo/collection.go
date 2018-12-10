package mongo

import (
	"github.com/globalsign/mgo"
)

type Collection struct {
	Name           string
	CollectionInfo *mgo.CollectionInfo
}

func (c *Collection) Create() error {
	collection := mgo.Collection{
		Database: Database,
		Name:     c.Name,
		FullName: "db." + c.Name,
	}

	return collection.Create(c.CollectionInfo)
}
