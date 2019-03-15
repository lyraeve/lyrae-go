package finder

import (
	"github.com/lyraeve/lyrae-go/contracts"
)

func FindByNumber(s contracts.Finder, number string) (lyr contracts.Lyr, err error) {
	return s.FindByNumber(number)
}
