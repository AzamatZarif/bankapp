package payment

import "bank/pkg/bank/types"

//Max возвращает максимальный платеж.
func Max(payments []types.Payment) types.Payment {
	var maxPayment types.Payment
	for _, payment := range payments {
		if payment.Amount >= maxPayment.Amount {
			maxPayment = payment
		}
	}
	return maxPayment
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSource []types.PaymentSource
	for _, card := range cards {
		if card.Balance <= 0 || !card.Active {
			break
		}
		paymentSource = append(paymentSource, types.PaymentSource{
			Type: "card",
			Number: string(card.PAN),
			Balance: card.Balance,
		})
	}
	return paymentSource
}