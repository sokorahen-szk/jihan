package main

import (
	"fmt"
	"strconv"
)

type PriceItem struct {
	Drink Drink
	Stock int
}

type Drink struct {
	Id    int
	Name  string
	Price int
}

type Jihan struct {
	Money               int
	PurchasedDrinks     []Drink
	AvailablePriceItems []*PriceItem
}

type Coins struct {
	oneThousand     int
	fiveHundredCoin int
	oneHundredCoin  int
	fiftyCoin       int
	tenCoin         int
}

var formatMoney = []int{1000, 500, 100, 50, 10}

func NewJihan() *Jihan {
	availablePriceItems := []*PriceItem{
		{
			Drink: Drink{
				Id:    1,
				Name:  "オレンジジュース",
				Price: 120,
			},
			Stock: 2,
		},
		{
			Drink: Drink{
				Id:    2,
				Name:  "トマトジュース",
				Price: 110,
			},
			Stock: 3,
		},
		{
			Drink: Drink{
				Id:    3,
				Name:  "モンスター",
				Price: 220,
			},
			Stock: 4,
		},
	}

	return &Jihan{
		AvailablePriceItems: availablePriceItems,
	}
}

func (j *Jihan) Exec() {
	j.inMoney()

	running := true
	for running {
		outputln("m = メニュー開く、 q = 終了する")
		output("> ")
		in := inStr()
		switch in {
		case "m":
			j.drinkMenu()
			break
		case "q":
			running = false
			break
		}
	}

	j.exchangeMoney()
}

func (j *Jihan) inMoney() {
	output("お金を入力> ")

	inMoney := inInt()

	j.SetMoney(inMoney)
}

func (j *Jihan) drinkMenu() {
	outputln("ドリンクを選択する。番号を入力してね。")
	for _, priceItem := range j.AvailablePriceItems {
		outputln(fmt.Sprintf(
			"[%d] %s %d円",
			priceItem.Drink.Id,
			priceItem.Drink.Name,
			priceItem.Drink.Price,
		))
	}
	output("ドリンク番号> ")
	j.purchase(inInt())
}

func (j *Jihan) purchase(drinkId int) {
	priceItem := j.getPriceItemById(drinkId)
	if priceItem == nil {
		outputln("選択されたドリンクが見つかりません。")
		return
	}

	if priceItem.Stock == 0 {
		outputln("選択されたドリンクの在庫がありません。")
		return
	}

	if j.GetMoney() < priceItem.Drink.Price {
		outputln("所持金が足りません。")
		outputln(fmt.Sprintf("所持金：%d円", j.GetMoney()))
		return
	}

	j.addPurchasedDrink(priceItem.Drink)
	priceItem.Stock--
	j.SetMoney(-priceItem.Drink.Price)
}

func (j *Jihan) getPriceItemById(drinkId int) *PriceItem {
	for _, priceItem := range j.AvailablePriceItems {
		if priceItem.Drink.Id == drinkId {
			return priceItem
		}
	}
	return nil
}

func (j *Jihan) exchangeMoney() int {
	exchangeMoney := j.GetMoney()

	outputln(fmt.Sprintf("おつり: %d円", exchangeMoney))

	// おつりを返すと同時に自販機の残金は0にする
	j.SetMoney(0)

	coins := j.calculationCoins(exchangeMoney)
	outputln(fmt.Sprintf(`1000 => %d
500 => %d
100 => %d
50 => %d
10 => %d`,
		coins.oneThousand,
		coins.fiveHundredCoin,
		coins.oneHundredCoin,
		coins.fiftyCoin,
		coins.tenCoin,
	))

	return exchangeMoney
}

func (j *Jihan) calculationCoins(money int) Coins {
	result := Coins{}
	in := money
	for i, coin := range formatMoney {
		var count int
		for true {
			if in < coin {
				break
			}
			in -= coin
			count++
		}

		switch i {
		case 0:
			result.oneThousand = count
		case 1:
			result.fiveHundredCoin = count
		case 2:
			result.oneHundredCoin = count
		case 3:
			result.fiftyCoin = count
		case 4:
			result.tenCoin = count
		}
	}

	return result
}
func (j *Jihan) GetMoney() int {
	return j.Money
}

func (j *Jihan) SetMoney(money int) {
	if money == 0 {
		j.Money = 0
		return
	}
	j.Money += money
}

func (j *Jihan) addPurchasedDrink(drink Drink) {
	j.PurchasedDrinks = append(j.PurchasedDrinks, drink)
}

func (j *Jihan) subStock(drinkId int) {
	priceItem := j.getPriceItemById(drinkId)
	if priceItem == nil {
		return
	}

	priceItem.Stock--
}

func main() {
	jihan := NewJihan()
	jihan.Exec()
}

func inInt() int {
	var s string
	fmt.Scan(&s)

	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func inStr() string {
	var s string
	fmt.Scan(&s)

	return s
}

func output(context string) {
	fmt.Print(context)
}

func outputln(context string) {
	fmt.Println(context)
}
