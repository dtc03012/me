package option

import "errors"

type RangeOption struct {
	Row  int32
	Size int32
}

func CalculateDBRange(option *RangeOption) (int32, int32, error) {
	if option.Row <= 0 || option.Size <= 0 {
		return 0, 0, errors.New("range option error: option is out of range")
	}

	return (option.Row - 1) * option.Size, option.Row * option.Size, nil
}
