package helper

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const taxScale int64 = 10000
const moneyScale int64 = 100

type InputData struct{ Prices []int64 }

func ApplyTax(price int64, taxRate int64) int64 {
	return price * (taxScale + taxRate) / taxScale
}

func FormatRupiah(amount int64) string {
	p := message.NewPrinter(language.Indonesian)

	rupiah := amount / moneyScale
	sen := amount % moneyScale

	return p.Sprintf("Rp %d,%02d", rupiah, sen)
}

func BpsFormat(taxRate int64) string {
	integer := taxRate / 100
	decimal := taxRate % 100

	return fmt.Sprintf("%d.%.2d%%", integer, decimal)
}
