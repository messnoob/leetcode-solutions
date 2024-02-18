package main

func removeDuplicates(nums []int) int {
	var k int
	k = 0
	for key, value := range nums {
		if key == 0 {
			continue
		}

		if nums[k] == value {
			continue
		} else {
			nums[k+1] = value
			k++
		}
	}

	return k + 1
}

func main() {

}
