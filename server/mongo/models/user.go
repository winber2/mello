package models

import (
	"mello/server/mongo"
	"time"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
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
	Username string        `json:"username" bson:"username"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

func (user *User) Save() error {
	now := time.Now()
	session := mongo.Database.Session.Copy()
	defer session.Close()

	c := mongo.Database.C("users").With(session)
	result, _ := FindUserByName(user.Username)

	if result != nil {
		return mongo.Error("Username already exists")
	}

	password, _ := getPasswordHash(user.Password)

	err := c.Insert(bson.M{
		"name":     user.Username,
		"email":    user.Email,
		"password": password,
		"created":  now,
		"updated":  now,
	})

	return err
}

// FindUserByName is a function to help find users in the database
func FindUserByName(name string) (*User, error) {
	var result *User
	err := mongo.Database.C("users").Find(name).One(&result)

	return result, err
}

func getPasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

// func (user *User) GetPID() string {
// 	return user.Name
// }
//
// func (user *User) PutPID(pid string) {
// 	mongo.Database.C("users").Update(
// 		bson.M{"_id": user.ID},
// 		bson.M{"$set": bson.M{"name": pid}},
// 	)
// }

// type UserStorer struct {
// 	// Users  map[string]User
// 	// Tokens map[string][]string
// }
//
// func (s *UserStorer) Load(ctx context.Context, key string) (authboss.User, error) {
// 	var result authboss.User
// 	objectID := bson.ObjectId(key)
// 	err := mongo.Database.C("users").Find(bson.M{"_id": objectID}).One(&result)
//
// 	return result, err
// }
//
// func (s *UserStorer) Save(ctx context.Context, u authboss.User) error {
// 	var result User
// 	user := u.(*User)
//
// 	c := mongo.Database.C("users")
// 	err := c.Find(bson.M{"_id": user.ID}).One(&result)
//
// 	if err != nil {
// 		return authboss.ErrUserNotFound
// 	}
//
// 	c.Update(
// 		bson.M{"_id": user.ID},
// 		result,
// 	)
//
// 	return nil
// }
//
// func (s *UserStorer) New(ctx context.Context) authboss.User {
// 	return &User{}
// }
//
// // Create the user
// func (s *UserStorer) Create(ctx context.Context, u authboss.User) error {
// 	var result *User
// 	user := u.(*User)
//
// 	c := mongo.Database.C("users")
// 	err := c.FindId(user.ID).One(&result)
//
// 	if err == nil {
// 		return authboss.ErrUserFound
// 	}
//
// 	// debugln("Created new user:", u.Name)
// 	c.Insert(*user)
//
// 	return nil
// }
