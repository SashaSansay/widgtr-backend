package user_model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
	"widgtr-backend/internal/models"
)

const CollectionName = "users"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Password  string             `json:"-"`
	Email     string             `json:"email"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	CreatedAt primitive.DateTime `json:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt"`
}

func NewUser(userIn User) *User {
	id := userIn.ID
	if userIn.ID.IsZero() {
		id = primitive.NewObjectID()
	}
	createdAt := primitive.NewDateTimeFromTime(time.Now())
	updatedAt := primitive.NewDateTimeFromTime(time.Now())
	return &User{
		ID:        id,
		Email:     userIn.Email,
		FirstName: userIn.FirstName,
		LastName:  userIn.LastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Password:  userIn.Password,
	}
}

func (u *User) InsertIntoDB() (*User, error) {
	_, err := models.DB.Collection(CollectionName).InsertOne(context.TODO(), u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetFromDB(id string) (*User, error) {
	var user User

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = models.DB.Collection(CollectionName).FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetByEmailFromDB(email string) (*User, error) {
	var user User
	err := models.DB.Collection(CollectionName).FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func HashPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 8)

	return string(hashedPassword), err
}

func CheckPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
