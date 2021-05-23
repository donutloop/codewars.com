package Twice_as_old

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	if sonYearsOld == 0 {
		return dadYearsOld
	}

	diff := dadYearsOld - sonYearsOld
	if diff < sonYearsOld {
		return sonYearsOld - diff
	}

	return dadYearsOld - (2 * sonYearsOld)
}

