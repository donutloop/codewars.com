package Do_I_get_a_bonus_

import "fmt"

func BonusTime(salary int, bonus bool) string {
	if bonus {
		return fmt.Sprintf("\u00A3%d", salary*10)
	}
	return fmt.Sprintf("\u00A3%d", salary)
}