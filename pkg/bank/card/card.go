package card

import "bank/pkg/bank/types"

const (
	withdrawLimit = 20_000_00 //TJS,USD,RUB
	depositLimit  = 50_000_00
	bonusLimit    = 5_000_00
)

//IssueCard выпускает новую карту.
func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

//Withdraw снимает деньги с карты.
func Withdraw(card types.Card, amount types.Money) types.Card {
	//Случаи когда баланс не должен измениться.
	if !card.Active || amount > card.Balance || amount <= 0 || amount > withdrawLimit {
		return card
	}
	card.Balance -= types.Money(amount)
	return card
}

// Deposit пополняет счеть карты.
func Deposit(card *types.Card, amount types.Money) {
	if !card.Active || amount > depositLimit {
		return
	}
	card.Balance += amount
}

//AddBonus вычисляет бонус на карту.
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active || card.MinBalance <= 0 || card.Balance <= 0 || card.Bonus > bonusLimit {
		return
	}
	bonus := card.MinBalance * types.Money(percent) / 100 * types.Money(daysInMonth) / types.Money(daysInYear)
	if bonus > bonusLimit {
		return
	}
	card.Bonus += bonus
	card.Balance += card.Bonus
}

//Total вычисляет общую сумму на всех картах.
func Total(cards []types.Card) types.Money {
	var sum types.Money
	for _, card := range cards {
		if card.Balance <=0 || !card.Active {
			break
		}
		sum += card.Balance
	}
	return sum
}

