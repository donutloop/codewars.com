package Surface_Area_and_Volume_of_a_Box

//  2ab + 2bc + 2ac (surface area)
//  V = L * H * W (volume)

func GetSize(w,h,d int) [2]int {
	return [2]int{CalSurfaceArea(w,h,d), CalVolume(w,h,d) }
}

func CalSurfaceArea(w,h,d int) int {
	return (2 * w * h) + (2 * h * d) + (2 * w * d)
}

func CalVolume(w,h,d int) int {
	return w * h * d
}