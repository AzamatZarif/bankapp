package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{ID: 1, Amount: 10_000_00},
		{ID: 2, Amount: 20_000_00},
		{ID: 3, Amount: 15_000_00},
	}

	max := Max(payments)
	fmt.Println(max.Amount)
	//Output: 2000000
}

//Examples для PaymentSources

func ExamplePaymentSource_normal() {
	cards := []types.Card{
		{PAN:"5058 xxxx xxxx 3031" , Balance: 100_000_00, Active: true},
		{PAN: "5058 xxxx xxxx 3331", Balance: 100_000_00, Active: true},
	}
	paymentSources := PaymentSources(cards)
	for _, paymenSource := range paymentSources {
		fmt.Println(paymenSource.Number)
	}
	//Output: 
	// 5058 xxxx xxxx 3031
	// 5058 xxxx xxxx 3331
}