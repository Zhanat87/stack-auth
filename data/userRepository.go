package data

import (
	"golang.org/x/crypto/bcrypt"
	// @link https://godoc.org/gopkg.in/mgo.v2
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/Zhanat87/stack-auth/models"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) Create(user *models.User) error {
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	//clear the incoming text password
	user.Password = ""
	err = r.C.Insert(&user)
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	// partial update on MongoDB
	err := r.C.Update(bson.M{"_id": user.Id},
		bson.M{"$set": bson.M{
			"firstname": user.FirstName,
			"lastname": user.LastName,
			"email": user.Email,
			"chatroom": user.ChatRoom,
		}})
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {

	err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	iter := r.C.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *UserRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *UserRepository) GetById(id string) (user models.User, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}