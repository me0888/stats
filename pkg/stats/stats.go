package stats

import "github.com/me0888/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money

	for _, v := range payments {
		sum += v.Amount
	}

	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money

	for _, v := range payments {
		if category == v.Category {
			sum += v.Amount
		}
	}

	return sum
}
