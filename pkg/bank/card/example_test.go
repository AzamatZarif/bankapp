package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 1000, Active: true}, 100)
	fmt.Println(result.Balance)
	//Output: 900
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 100, Active: true}, 1_000)
	fmt.Println(result.Balance)
	//Output: 100
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{}, 1)
	fmt.Println(result.Balance)
	//Output: 0
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 10_000_000}, 50_000_00)
	fmt.Println(result.Balance)
	//Output: 10000000
}

//Deposit examples.

func ExampleDeposit_normal() {
	card := types.Card{Balance: 100, Active: true}
	Deposit(&card, 100)
	fmt.Println(card.Balance)
	//Output: 200
}

func ExampleDeposit_inactive() {
	card := types.Card{}
	Deposit(&card, 100)
	fmt.Println(card.Balance)
	//Output: 0
}

func ExampleDeposit_aboveLimit() {
	card := types.Card{Balance: 100, Active: true}
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)
	//Output: 100
}


//AddBonus examples.

func ExampleAddBonus() {
	card := types.Card{Balance: 20_000_00,MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 2002465
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000_00,MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 2000000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -20_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: -2000000
	
}

func ExampleAddBonus_aboveLimit() {
	card := types.Card{Balance: 20_000_000_00, MinBalance: 500_000_00, Active: true}
	AddBonus(&card, 30, 30, 365)
	fmt.Println(card.Balance)
	//Output: 2000000000
}

// Examples для Total.
func ExampleTotal_normal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	//Output: 2000000	
}

func ExampleTotal_negativeBalance() {
	cards := []types.Card{
		{
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	//Output: 0
}

func ExampleTotal_inactive() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
		},
		{
			Balance: 10_000_00,
		},
	}
	fmt.Println(Total(cards))
	//Output: 0
}