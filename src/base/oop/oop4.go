package main

import "fmt"

type LuxuryItem struct {
	Item
	markup float64
}
func (item * LuxuryItem) Cost() float64{
	return item.Item.price * item.markup
}
type Item struct{
	id 			string
	price 		float64
	quantity	int
}

func (item *Item) Cost() float64  {
	return item.price * float64(item.quantity)
}

type SpecialItem struct {
	Item
	catalogId int
}

func main()  {
	specialItem := SpecialItem{Item{"Green",3,5}, 207}
	fmt.Println(specialItem.id,specialItem.price,specialItem.quantity)
	fmt.Println(specialItem.Cost())

	luxuryItem := LuxuryItem{Item{"Green",3, 5},10}
	fmt.Println(luxuryItem.Cost())
}
