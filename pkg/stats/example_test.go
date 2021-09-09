package stats

import (
	"fmt"

	"github.com/me0888/bank/pkg/bank/types"
)

func ExampleAvg() {
	pays := []types.Payment{
		{
			ID:     0,
			Amount: 100,
		},
		{
			ID:     1,
			Amount: 100,
		},
		{
			ID:     2,
			Amount: 100,
		},
	}
	fmt.Println(Avg(pays))
	//Output:100

}

func ExampleTotalInCategory() {

	pays := []types.Payment{
		{
			ID:       0,
			Amount:   100,
			Category: "0",
		},
		{
			ID:       1,
			Amount:   100,
			Category: "0",
		},
		{
			ID:     2,
			Amount: 100,
		},
	}
	fmt.Println(TotalInCategory(pays, "0"))
	//Output:200

}
