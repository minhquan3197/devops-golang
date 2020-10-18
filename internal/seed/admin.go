package seed

import (
	"fmt"
	service "project-golang/pkg/users"
	structs "project-golang/pkg/users/structs"

	"go.mongodb.org/mongo-driver/bson"
)

// Admin func seed admin info
func Admin() {
	adminAccount := "admin"
	condition := bson.M{"username": adminAccount}
	_, err := service.FindOne(condition)
	if err != nil {
		newAdmin := structs.CreateUser{
			Username: adminAccount,
			Password: adminAccount,
		}
		service.Create(newAdmin)
		fmt.Println("Seed admin complete")
	}
}
