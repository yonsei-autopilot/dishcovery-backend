package domain

import "time"

type User struct {
	Password string `firestore:"password"` // only for simple login

	Name         string `firestore:"name"`
	Language     string `firestore:"language"`
	DislikeFoods string `firestore:"dislikeFoods"`
	AuthProvider string `firestore:"authProvider"` // "google", "github", "simple"
	RefreshToken string `firestore:"refreshToken"`

	CreatedAt time.Time `firestore:"createdAt"`
	LastLogin time.Time `firestore:"lastLogin"`
}
