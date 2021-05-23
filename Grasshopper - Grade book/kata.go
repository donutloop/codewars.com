package Grasshopper___Grade_book

func GetGrade(a,b,c int) rune {
	mean := (a+b+c)/3

	var score rune
	if 90 <= mean && mean <= 100 {
		score = 'A'
	} else if 80 <= mean && mean <= 90 {
		score = 'B'
	} else if 70 <= mean && mean <= 80 {
		score = 'C'
	} else if 60 <= mean && mean <= 70 {
		score = 'D'
	} else if 0 <= mean && mean <= 60 {
		score = 'F'
	}

	return score
}