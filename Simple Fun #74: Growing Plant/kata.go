package Simple_Fun__74__Growing_Plant

func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	var currentHeight int
	var day int
	for {
		day++
		currentHeight += upSpeed
		if currentHeight >= desiredHeight {
			break
		}
		currentHeight -= downSpeed
	}
	return day
}