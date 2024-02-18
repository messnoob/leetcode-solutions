package main

func removeElement(nums []int, val int) int {
	var k = 0

	for _, value := range nums {
		if value != val {
			nums[k] = value
			k++
		}
	}

	return k
}

func main() {

}
