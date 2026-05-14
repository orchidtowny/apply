package util

import (
	"apply/definition"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var ctx = context.Background()

// CreateApplication inserts application into database and pushes notifications if configured
func CreateApplication(application definition.Application) {
	if Database == nil {
		return
	}

	err := gorm.G[definition.Application](Database).Create(ctx, &application)
	if err != nil {
		fmt.Println("Error creating application", err)
	}
}

// GetApplication gets an application by the ID
func GetApplication(id string) *definition.Application {
	if Database == nil {
		return nil
	}

	application, err := gorm.G[definition.Application](Database).Where("id = ?", id).First(ctx)
	if err != nil {
		fmt.Println("Error getting application", err)
		return nil
	}

	return &application
}

// AcceptApplication will update the status of an application to be accepted
func AcceptApplication(id string) {
	if Database == nil {
		return
	}

	_, err := gorm.G[definition.Application](Database).Where("id = ?", id).Update(ctx, "status", 1)
	if err != nil {
		fmt.Println("Error rejecting application", err)
	}
}

// RejectApplication will update the status of an application to be rejected
func RejectApplication(id string) {
	if Database == nil {
		return
	}

	_, err := gorm.G[definition.Application](Database).Where("id = ?", id).Update(ctx, "status", 1)
	if err != nil {
		fmt.Println("Error rejecting application", err)
	}
}
