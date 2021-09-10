package stats

import "github.com/me0888/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var i types.Money

	for _, v := range payments {
		if v.Status != types.StatusFail {
			sum += v.Amount
			i += 1
		}
	}

	return sum / i
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

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	catCol := map[types.Category]types.Money{}

	var payment types.Payment
	for _, payment = range payments {
		categories[payment.Category] += payment.Amount
		catCol[payment.Category]++
	}

	for i := range categories {
		categories[i] /= catCol[i]
	}

	return categories
}
