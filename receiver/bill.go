package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func biller(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (lol *bill) updateTip(tip float64) {
	lol.tip = tip
}

func (lol *bill) updateItems(pname string, pvalue float64) {
	lol.items[pname] = pvalue
}

func (lol *bill) showBill() string {
	fs := "My product Bill \n"
	var totalPrice float64 = 0

	for k, v := range lol.items {
		fs += fmt.Sprintf("%-5v ... %v \n", k, v)
		totalPrice += v
	}
	tip := lol.tip
	fs += fmt.Sprintf("%-5v ... %v \n", "Tip", tip)
	fs += fmt.Sprintf("%-5v ... %v", "Total", totalPrice+tip)

	return fs
}
