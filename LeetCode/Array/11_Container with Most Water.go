package main

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(height)
}

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height) - 1

	for start < end {
		width := end - start
		area := 0

		if height[start] < height[end] {
			area := height[start] * width
			start++
		} else {
			area := height[end] * width
			end--
		}

		if area > max {
			max = area
		}
	}
	return max
}
