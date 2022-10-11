package controllers

import (
	"ass-2/database"
	"ass-2/models"
	"fmt"
)

func CreateItem(itemCode uint, orderId string, description string, quantity uint) {
	db := database.GetDB()

	Item := models.Item{
		ItemCode:    itemCode,
		Description: description,
		Quantity:    quantity,
		OrderID:     orderId,
	}

	err := db.Create(&Item).Error

	if err != nil {
		fmt.Println("Error creating item data: ", err.Error())
		return
	}

	fmt.Println("New Item Data:", Item)
}

func getOrdersWithItems() {
	db := database.GetDB()

	orders := models.Order{}
	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas with items", err.Error())
		return
	}

	fmt.Println("Order datas with items")
	fmt.Printf("%+v", orders)
}

func deleteItemById(id uint) {
	db := database.GetDB()

	item := models.Item{}
	err := db.Where("id = ? ", id).Delete(&item).Error

	if err != nil {
		fmt.Println("Error deleting item", err.Error())
		return
	}
}
