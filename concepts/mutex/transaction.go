package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup
var balanceMutex sync.Mutex

type Income struct {
	source string
	amount int
}

func transaction() {
	var bankBalance int
	fmt.Printf("Initial account balance:$%d.00", bankBalance)
	fmt.Println()
	incomes := []Income{
		{source: "Main job", amount: 5000},
		{source: "Gifts", amount: 100},
		{source: "part-time job", amount: 2500},
		{source: "Investments", amount: 100},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balanceMutex.Lock()
				temp := bankBalance
				temp += income.amount
				bankBalance = temp
				balanceMutex.Unlock()
				fmt.Printf("On week %d, you have earned %d.00 from %s\n", week, income.amount, income.source)
			}
		}(i, income)
	}
	wg.Wait()
	fmt.Println("Final:$", bankBalance)
}
