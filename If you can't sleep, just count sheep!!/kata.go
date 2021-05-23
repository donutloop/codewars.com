package If_you_can_t_sleep__just_count_sheep__

import "fmt"

func countSheep(num int) string {
	var text string
	for i := 1; i < num+1; i++ {
		text += fmt.Sprintf("%d sheep...", i)
	}
	return text
}