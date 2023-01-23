package domain

import (
	"fmt"

	"rsc.io/quote"
)

func GetDomainName() string {
	return quote.Go()
}

func DoAction() {
	fmt.Sprintln(quote.Glass())
}
