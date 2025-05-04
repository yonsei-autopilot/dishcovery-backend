package domain

import "time"

type User struct {
	Id           string    `firestore:"id"`
	Name         string    `firestore:"name"`
	DislikeFoods string    `firestore:"dislikeFoods"`
	CreatedAt    time.Time `firestore:"createdAt"`
	LastLogin    time.Time `firestore:"lastLogin"`
}
