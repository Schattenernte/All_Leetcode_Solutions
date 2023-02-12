package main

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	old_value := 0
	num := 0
	table := make(map[rune]int)
	table['I'] = 1
	table['V'] = 5
	table['X'] = 10
	table['L'] = 50
	table['C'] = 100
	table['D'] = 500
	table['M'] = 1000
	for _, content := range s {
		if value, ok := table[content]; ok {
			num = num + value
			if old_value < value {
				num = num - 2*old_value
			}
			old_value = value
		}
	}
	return num
}
