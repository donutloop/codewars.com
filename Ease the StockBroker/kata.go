package Ease_the_StockBroker

import "fmt"
import "strconv"
import "strings"

const (
	quoteType int = iota
	quantiyType
	priceType
	statusType
)

const whitespace byte = ' '
const comma byte = ','

func BalanceStatement(lst string) string {
	var quantity int
	var price float64
	var status string
	var quoteRaw string
	var quantityRaw string
	var priceRaw    string
	var statusRaw  string

	value := make([]byte, 0)
	var valueType int
	var counter int

	var buy float64
	var sell float64
	var badOrders string
	var goodFormat = "Buy: %.0f Sell: %.0f"
	var badFormat = "Buy: %.0f Sell: %.0f; Badly formed %v: %s"
	var badOrdersCount int
	for i := 0; i < len(lst); i++ {
		if i > 0 && comma == lst[i-1] && whitespace == lst[i] {
			continue
		}

		if whitespace == lst[i] || comma == lst[i] || i == len(lst)-1 {
			if i == len(lst)-1  {
				value = append(value, lst[i])
			}


			switch valueType {
			case quoteType:
				if len(value) != 0 {
					counter++
				}
				quoteRaw = string(value)
			case quantiyType:

				num, err := strconv.Atoi(string(value))
				if err == nil {
					quantity = num
					counter++
				}
				quantityRaw = string(value)
			case priceType:
				if strings.Contains(string(value), ".") {
					num, err := strconv.ParseFloat(string(value), 64)
					if err == nil {
						price = num
						counter++
					}
				}
				priceRaw = string(value)
			case statusType:
				if string(value) == "S" || string(value) == "B" {
					status = string(value)
					counter++
				}
				statusRaw = string(value)
			}
			value = make([]byte, 0)
			valueType++
		} else {
			value = append(value, lst[i])
		}

		if comma == lst[i] || i == len(lst)-1 {
			if counter != 4 {
				badOrdersCount++
				badOrders += quoteRaw + " " + quantityRaw + " " + priceRaw + " "
				if statusRaw != "" {
					badOrders += statusRaw + " ;"
				} else {
					badOrders += ";"
				}

			} else {
				if status == "B" {
					buy += float64(quantity) * price
				} else if status == "S" {
					sell += float64(quantity) * price
				}
			}
			quantity = 0
			price = 0
			status = ""
			valueType = 0
			counter=0
		}
	}

	if badOrders != "" {
		return fmt.Sprintf(badFormat, buy, sell, badOrdersCount, badOrders)
	}


	return fmt.Sprintf(goodFormat, buy, sell)
}