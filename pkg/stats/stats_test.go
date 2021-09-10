package stats

import (
	"reflect"
	"testing"

	"github.com/me0888/bank/v2/pkg/types"
)



func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 1_000_00, Category: "auto", Status: types.StatusOK},
		{ID: 2, Amount: 1_000_00, Category: "auto", Status: types.StatusOK},
		{ID: 3, Amount: 2_000_00, Category: "mobile", Status: types.StatusOK},
		{ID: 4, Amount: 2_000_00, Category: "mobile", Status: types.StatusOK},
		{ID: 5, Amount: 3_000_00, Category: "it", Status: types.StatusOK},
		{ID: 6, Amount: 3_000_00, Category: "it", Status: types.StatusOK},
		{ID: 7, Amount: 5_000_00, Category: "fun", Status: types.StatusOK},
		{ID: 8, Amount: 5_000_00, Category: "fun", Status: types.StatusOK},
	}

	expected := map[types.Category]types.Money{
		"auto":   1_000_00,
		"mobile": 2_000_00,
		"it":     3_000_00,
		"fun":    5_000_00,
	}

	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\ninvalid result, \ngot:  %v, \nwant: %v", expected, result)
	}

}


