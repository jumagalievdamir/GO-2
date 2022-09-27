package main

import "fmt"

type bahandi interface {
	getPrice() int
}

type Burger struct {
}

func (b *Burger) getPrice() int {
	return 1200
}

type CheeseBurger struct {
}

func (b *CheeseBurger) getPrice() int {
	return 1400
}

type ketchup struct {
	bahandi bahandi
}

func (c *ketchup) getPrice() int {
	bahandiPrice := c.bahandi.getPrice()
	return bahandiPrice + 100
}

type cheese struct {
	bahandi bahandi
}

func (c *cheese) getPrice() int {
	bahandiPrice := c.bahandi.getPrice()
	return bahandiPrice + 150
}

func main() {

	BurgerBahandi := &Burger{}

	//Add cheese topping
	BurgerBahandiWithKetchup := &ketchup{
		bahandi: BurgerBahandi,
	}

	//Add tomato topping
	BurgerBahandiWithKetchupAndCheese := &cheese{
		bahandi: BurgerBahandiWithKetchup,
	}

	fmt.Printf("Price of Burger with Ketchup and Cheese is %d\n", BurgerBahandiWithKetchupAndCheese.getPrice())

	CheeseBurgerBahandi := &CheeseBurger{}

	//Add cheese topping
	CheeseBurgerBahandiWithKetchup := &ketchup{
		bahandi: CheeseBurgerBahandi,
	}

	fmt.Printf("Price of CheeseBurger with ketchup is %d\n", CheeseBurgerBahandiWithKetchup.getPrice())
}
