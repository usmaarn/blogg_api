package model

type Order struct {
	id     string
	amount int
}

func (order *Order) GetId() string {
	return order.id
}

func (order *Order) GetAmount() int {
	return order.amount
}

func (order *Order) SetId(id string) {
	order.id = id
}

func (order *Order) SetAmount(amount int) {
	order.amount = amount
}
