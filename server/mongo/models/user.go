package models

import (
	"context"
	"mello/server/mongo"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/volatiletech/authboss"
)

/*
	AuthBoss ServerStorer and User interface from the docs that must be implemented
	type ServerStorer interface {
	    // Load will look up the user based on the passed the PrimaryID
	    Load(ctx context.Context, key string) (User, error)

	    // Save persists the user in the database, this should never
	    // create a user and instead return ErrUserNotFound if the user
	    // does not exist.
	    Save(ctx context.Context, user User) error
	}

	type User interface {
    GetPID() (pid string)
    PutPID(pid string)
	}
*/

// User table contains the information for each user
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

var (
	assertUser = &User{}

	_ authboss.User = assertUser
)

func (user *User) Save() error {
	now := time.Now()
	session := mongo.Database.Session.Copy()
	defer session.Close()

	c := mongo.Database.C("users").With(session)

	err := c.Insert(bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"created":  now,
		"updated":  now,
	})

	return err
}

func (user *User) GetPID() string {
	return user.Name
}

func (user *User) PutPID(pid string) {
	mongo.Database.C("users").Update(
		bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{"name": pid}},
	)
}

type UserStorer struct {
	// Users  map[string]User
	// Tokens map[string][]string
}

func (s *UserStorer) Load(ctx context.Context, key string) (authboss.User, error) {
	var result authboss.User
	objectID := bson.ObjectId(key)
	err := mongo.Database.C("users").Find(bson.M{"_id": objectID}).One(&result)

	return result, err
}

func (s *UserStorer) Save(ctx context.Context, u authboss.User) error {
	var result User
	user := u.(*User)

	c := mongo.Database.C("users")
	err := c.Find(bson.M{"_id": user.ID}).One(&result)

	if err != nil {
		return authboss.ErrUserNotFound
	}

	c.Update(
		bson.M{"_id": user.ID},
		result,
	)

	return nil
}
