package stats

import "github.com/me0888/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money

	for _, v := range payments {
		if v.Status != types.StatusFail {
			sum += v.Amount
		}
	}

	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money

	for _, v := range payments {
		if category == v.Category {
			if v.Status != types.StatusFail {
				sum += v.Amount
			}
		}
	}

	return sum
}
