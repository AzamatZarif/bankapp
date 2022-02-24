package types

//Money представлет собой денежную сумму.
type Money int64

//Currency представляет собой код валюти.
type Currency string

//Коды валют.
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN Представляет собой номер карты.
type PAN string

//Card представляет информацию о платежной карте.
type Card struct {
	ID       int
	PAN      PAN
	Balance  Money //использовали тип Money.
	Currency Currency
	Color    string
	Name     string
	Active   bool
	MinBalance Money
	Bonus Money
}

//Payment представляет информация о платеже.
type Payment struct {
	ID int
	Amount Money
}

type PaymentSource struct {
	Type string
	Number string
	Balance Money //дирам
}