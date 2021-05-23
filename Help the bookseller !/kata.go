package Help_the_bookseller__

import "strconv"

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 {
		return ""
	}

	if len(listCat) == 0 {
		return ""
	}

	sortiment := make([]int, 26)

	for _, value := range listArt {

		category := value[0]
		numRaw := make([]byte, 0)
		for i := 0; i < len(value); i++ {
			if value[i] >= 48 && value[i] <= 57 {
				numRaw = append(numRaw, value[i])
			}
		}

		num, _ := strconv.Atoi(string(numRaw))
		sortiment[int(category)-65] += num
	}

	var tuble string
	for _, cat := range listCat {
		tuble +=  "(" + cat + " : " + strconv.Itoa(sortiment[int(cat[0])-65]) + ") - "
	}

	return tuble[:len(tuble)-3]
}