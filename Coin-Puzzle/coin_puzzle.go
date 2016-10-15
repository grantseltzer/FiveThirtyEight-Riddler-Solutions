package main

import "fmt"

type coin struct {
	numberCoin *int
	headsUp    *bool
}

func flipCoins(valueToFlip int, coins *[]coin) {
	for _, coin := range *coins {
		if *coin.numberCoin%valueToFlip == 0 {
			*coin.headsUp = !*coin.headsUp
		}
	}
}

func main() {
	numberCoins := []int{}
	for i := 0; i < 100; i++ {
		numberCoins = append(numberCoins, i+1)
	}

	coinsBool := []bool{}
	for i := 0; i < 100; i++ {
		coinsBool = append(coinsBool, true)
	}

	coins := []coin{}
	for i := 0; i < 100; i++ {
		shinyNewCoin := coin{
			numberCoin: &numberCoins[i],
			headsUp:    &coinsBool[i],
		}
		coins = append(coins, shinyNewCoin)
	}

	for x := 1; x < 101; x++ {
		flipCoins(x, &coins)
	}

	for x := range coins {
		if *coins[x].headsUp == false {
			fmt.Println(*coins[x].numberCoin)
		}
	}
}
