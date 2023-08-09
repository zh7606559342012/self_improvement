package sorting

//mao pao pai xu
func BubbleSort(tmp []int) []int {
	tmpLen := len(tmp)
	for i := 0; i < tmpLen; i++ {
		for j := 0; j < tmpLen-1; j++ {
			if tmp[j] > tmp[j+1] {
				tmp[j], tmp[j+1] = tmp[j+1], tmp[j]
			}
		}
	}
	return tmp
}

//xuan ze pai xu
func SelectionSort(tmp []int) []int {
	tmpLen := len(tmp)
	for i := 0; i < tmpLen; i++ {
		swapLocation := i
		for j := i; j < tmpLen-1; j++ {
			if tmp[j] > tmp[j+1] {
				swapLocation = j + 1
			}
		}
		tmp[i], tmp[swapLocation] = tmp[swapLocation], tmp[i]
	}
	return tmp
}

//cha ru pai xu
func ChaRuSort(tmp []int) []int {
	tmpLen := len(tmp)
	for i := 1; i < tmpLen; i++ {
		for j := 0; j < i; j++ {
			if tmp[j] > tmp[i] {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
	}
	return tmp
}

//test
func TestSum(t *testing.T) {
	tmp := []int{4, 3, 2, 1}
	result := BubbleSort(tmp)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
