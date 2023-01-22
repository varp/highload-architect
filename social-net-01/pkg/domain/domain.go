package domain

import "rsc.io/quote"

func GetDomainName() string {
	return quote.Go()
}
