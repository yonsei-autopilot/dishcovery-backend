package domain

import "google.golang.org/genproto/googleapis/type/datetime"

type User struct {
	Id           string            `firestore:"id"`
	Name         string            `firestore:"name"`
	DislikeFoods string            `firestore:"dislikeFoods"`
	CreatedAt    datetime.DateTime `firestore:"createdAt"`
	LastLogin    datetime.DateTime `firestore:"lastLogin"`
}
