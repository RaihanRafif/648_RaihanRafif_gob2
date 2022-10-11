package controllers

import (
	"ass-2/database"
	"ass-2/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func createOrder(CustomerName string) {
	db := database.GetDB()
	Order := models.Order{
		CustomerName: CustomerName,
	}

	err := db.Create(&Order).Error

	if err != nil {
		fmt.Println("Error creating order data:", err)
		return
	}
	fmt.Println("New Order Data :", Order)
}

func getOrderById(id uint) {
	db := database.GetDB()

	order := models.Order{}

	err := db.First(&order, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Order data not found")
			return
		}
		print("error finding order:", err)
	}
	fmt.Printf("Order Data: %+v \n", order)
}

func updateOrderById(id uint, CustomerName string) {
	db := database.GetDB()

	order := models.Order{}

	err := db.Model(&order).Where("id = ?", id).Updates(models.Order{CustomerName: CustomerName}).Error

	if err != nil {
		fmt.Println("Error updating order data:", err)
		return
	}
	fmt.Printf("Update order's email: %+v \n", order.CustomerName)
}
