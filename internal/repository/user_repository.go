package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/firebase"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
)

func AddUser(ctx context.Context, user *domain.User) (string, error) {
	client := firebase.GetClient()
	ref, _, err := client.Collection("users").Add(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to add user: %w", err)
	}
	return ref.ID, nil
}

func GetAllUsers(ctx context.Context) ([]domain.User, error) {
	client := firebase.GetClient()
	iter := client.Collection("users").Documents(ctx)

	var users []domain.User
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var u domain.User
		err = doc.DataTo(&u)
		if err != nil {
			log.Printf("Skipping invalid user doc: %v", err)
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

func GetUserById(ctx context.Context, id string) (*domain.User, error) {
	client := firebase.GetClient()
	doc, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID %s: %w", id, err)
	}
	var user domain.User
	err = doc.DataTo(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user data: %w", err)
	}
	return &user, nil
}

func GetUserByName(ctx context.Context, name string) (*domain.User, error) {
	client := firebase.GetClient()

	iter := client.Collection("users").Where("name", "==", name).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return nil, fmt.Errorf("failed to find user with name %q: %w", name, err)
	}

	var user domain.User
	if err := doc.DataTo(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user data: %w", err)
	}
	return &user, nil
}
