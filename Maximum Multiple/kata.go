package Maximum_Multiple

func MaxMultiple(d, b int) (r int) {
	for i := b; i > 0; i-- {
		if i % d == 0 {
			r = i
			return
		}
	}
	return
}