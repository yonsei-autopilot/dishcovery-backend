package domain

import "time"

type User struct {
	Password *string `firestore:"password,omitempty"` // only for simple login

	Name         string  `firestore:"name"`
	Language     *string `firestore:"language"`
	DislikeFoods *string `firestore:"dislikeFoods,omitempty"`
	AuthProvider string  `firestore:"authProvider"` // "google", "github", "simple"

	CreatedAt time.Time  `firestore:"createdAt"`
	LastLogin *time.Time `firestore:"lastLogin,omitempty"`
}
