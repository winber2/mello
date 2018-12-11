package mongo

import (
	"github.com/globalsign/mgo"
)

type Collection struct {
	Name           string
	CollectionInfo *mgo.CollectionInfo
	Indexes        []mgo.Index
}

func (c *Collection) Create() error {
	collection := mgo.Collection{
		Database: Database,
		Name:     c.Name,
		FullName: "mello." + c.Name,
	}

	err := collection.Create(c.CollectionInfo)

	if err != nil {
		return err
	}

	for _, index := range c.Indexes {
		if err := collection.EnsureIndex(index); err != nil {
			return err
		}
	}

	return nil
}
