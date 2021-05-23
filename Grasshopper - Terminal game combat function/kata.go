package Grasshopper___Terminal_game_combat_function

func Combat(health, damage float64) float64 {
	health = health - damage
	if health < 0 {
		health = 0
	}
	return health
}
