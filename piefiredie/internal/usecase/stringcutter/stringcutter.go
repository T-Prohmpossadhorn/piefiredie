package stringcutter

func StringCutter(input string, cutter func(rune) bool, processor func(string) error) error {
	output := ""
	for _, v := range input {
		if cutter(v) {
			output += string(v)
		} else {
			if len(output) > 0 {
				err := processor(output)
				if err != nil {
					return err
				}
			}
			output = ""
		}
	}
	return nil
}
