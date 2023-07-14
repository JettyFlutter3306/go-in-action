package generic

func Sum[T int | string](arr ...T) (ans T) {
	for _, t := range arr {
		ans += t
	}
	return
}
