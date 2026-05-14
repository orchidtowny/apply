package util

import (
	"apply/definition"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var ctx = context.Background()

/*  CREATE  */

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

/*  GET  */

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

// GetApplicationByUsername gets an application by the username
func GetApplicationByUsername(username string) *definition.Application {
	if Database == nil {
		return nil
	}

	application, err := gorm.G[definition.Application](Database).Where("username = ?", username).First(ctx)
	if err != nil {
		fmt.Println("Error getting application", err)
		return nil
	}

	return &application
}

/*  GET MANY  */

// GetPendingApplications gets all pending applications
func GetPendingApplications() []definition.Application {
	if Database == nil {
		return nil
	}

	applications, err := gorm.G[definition.Application](Database).Where("status = ?", 0).Find(ctx)
	if err != nil {
		fmt.Println("Error getting applications", err)
		return nil
	}

	return applications
}

// GetApprovedApplications gets all approved applications
func GetApprovedApplications() []definition.Application {
	if Database == nil {
		return nil
	}

	applications, err := gorm.G[definition.Application](Database).Where("status = ?", 2).Find(ctx)
	if err != nil {
		fmt.Println("Error getting applications", err)
		return nil
	}

	return applications
}

// GetRejectedApplications gets all rejected applications
func GetRejectedApplications() []definition.Application {
	if Database == nil {
		return nil
	}

	applications, err := gorm.G[definition.Application](Database).Where("status = ?", 1).Find(ctx)
	if err != nil {
		fmt.Println("Error getting applications", err)
		return nil
	}

	return applications
}

/*  ACTIONS  */

// ApproveApplication will update the status of an application to be accepted
func ApproveApplication(id string) {
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
