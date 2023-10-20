package helpers

import "godev/go-rest-api/models"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ZipItems(items1 []models.Item, items2 []models.Item) [][]models.Item {
	if len(items1) != len(items2) {
		return nil
	}
	result := [][]models.Item{}
	for i, itm1 := range items1 {
		result = append(result, []models.Item{itm1, items2[i]})
	}
	return result
}

func RemoveIndex(slc []models.Order, orderID int) []models.Order {
	var iStop int = -1
	newOrderData := []models.Order{}
	for i, o := range slc {
		if o.OrderID == orderID {
			iStop = i
			break
		}
		newOrderData = append(newOrderData, o)
	}
	if iStop+1 <= len(slc) {
		return append(newOrderData, slc[iStop+1:]...)
	}
	return newOrderData
}
