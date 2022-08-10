package option

import "errors"

type RangeOption struct {
	Row  int
	Size int
}

func CalculateDBRange(option *RangeOption) (int, int, error) {
	if option == nil || option.Row <= 0 || option.Size <= 0 {
		return 0, 0, errors.New("range option error: option is out of range")
	}

	return (option.Row - 1) * option.Size, option.Row * option.Size, nil
}
