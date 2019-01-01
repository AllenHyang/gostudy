package main

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for key, value := range nums {

		complement := target - value

		if v, ok := m[complement]; ok == true {
			return []int{key, v}
		}
		m[value] = key
	}

	panic("Not two sum solution")
}
