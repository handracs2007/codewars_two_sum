package main

func TwoSum(numbers []int, target int) [2]int {
	for index, val := range numbers {
		for j := index + 1; j < len(numbers); j++ {
			var val2 = numbers[j]

			if val+val2 == target {
				return [2]int{index, j}
			}
		}
	}

	return [2]int{}
}

func main() {

}
