package models

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Connection struct {
	DB *gorm.DB
}

func (c *Connection) CreateRec(newRec *Order) error {
	res := c.DB.Create(&newRec)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *Connection) GetAllRec(Recs []Order) error {
	res := c.DB.Preload("Items").Find(&Recs)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *Connection) SearchRecByID(Rec *Order, id int) error {
	res := c.DB.Preload("Items").Where("order_id = ?", id).First(&Rec)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func (c *Connection) UpdateRecByID(old *Order, new *Order, id int) error {
	res := c.DB.Preload("Items").Where("order_id = ?", id).First(&old)
	if res.RowsAffected == 0 {
		return res.Error
	}
	// update data
	// delete existing items with related order_id
	var deletedItem []Item
	res = c.DB.Clauses(clause.Returning{}).Where("order_id = ?", old.OrderID).Delete(&deletedItem)
	if res.Error != nil {
		return res.Error
	}

	old.CustomerName = new.CustomerName
	old.OrderedAt = new.OrderedAt
	old.Items = new.Items
	res = c.DB.Save(&old)
	if res.Error != nil {
		return res.Error
	}
	return nil

}

func (c *Connection) DeleteRecByID(Recs *[]Order, id int) error {
	res := c.DB.Clauses(clause.Returning{}).Where("order_id = ?", id).Delete(&Recs)
	if res.RowsAffected == 0 {
		return errors.New("Order ID not found!")
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}
