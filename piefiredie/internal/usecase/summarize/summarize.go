package summarize

import (
	"piefiredie/internal/repository"
	"piefiredie/internal/usecase/stringcutter"
	"strings"
)

func SummarizeBeefStock(input string) (map[string]int32, error) {
	stock := repository.CreateBeefStock()

	err := stringcutter.StringCutter(input, func(v rune) bool {
		return !strings.Contains(",. \n", string(v))
	}, func(v string) error {
		_, err := stock.AddToStock(v)
		return err
	})
	if err != nil {
		return nil, err
	}

	return stock.GetStock()
}
