package tj_test

import (
	"encoding/json"
	"fmt"

	"github.com/dottedmag/tj"
)

func Example() {
	j := tj.O{
		"hello": tj.A{
			"world",
			42,
			tj.O{
				"go": nil,
			},
		},
	}
	bytes, _ := json.MarshalIndent(j, "", "  ")
	fmt.Printf("%s\n", bytes)
	// Output:
	// {
	//   "hello": [
	//     "world",
	//     42,
	//     {
	//       "go": null
	//     }
	//   ]
	// }
}
