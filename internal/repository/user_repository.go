package repository

import (
	"context"
	"fmt"

	"github.com/yonsei-autopilot/smart-menu-backend/internal/common/firebase"
	"github.com/yonsei-autopilot/smart-menu-backend/internal/domain"
)

func AddUser(ctx context.Context, id string, user *domain.User) error {
	client := firebase.GetClient()

	docRef := client.Collection("users").Doc(id)

	_, err := docRef.Set(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to add user %s: %w", id, err)
	}

	return nil
}

func UpdateUser(ctx context.Context, id string, user *domain.User) error {
	client := firebase.GetClient()

	_, err := client.Collection("users").Doc(id).Set(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to set user %s: %w", id, err)
	}

	return nil
}

func GetUserById(ctx context.Context, id string) (*domain.User, error) {
	client := firebase.GetClient()

	doc, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID %s: %w", id, err)
	}

	var user domain.User
	if err := doc.DataTo(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user data: %w", err)
	}

	return &user, nil
}
