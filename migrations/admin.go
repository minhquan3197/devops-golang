package migrations

import (
	"fmt"
	"project-golang/internal/interfaces"
	"project-golang/internal/services"
)

// Admin func seed admin info
func Admin() {
	adminAccount := "admin"
	_, err := services.FindUserByUsername(adminAccount)
	if err != nil {
		newAdmin := interfaces.CreateUser{
			Username: adminAccount,
			Password: adminAccount,
		}
		services.Register(newAdmin)
		fmt.Println("Seed admin complete")
	}
}
