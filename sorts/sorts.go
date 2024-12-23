package sorts

import (
	"log"
	"math"
	"math/rand"
)

func RunSortFunctions() {
	sliceToSort := make([]int, 0, 20)

	for range 15 {
		sliceToSort = append(sliceToSort, rand.Intn(100))
	}

	log.Println("Bubble Sort:", BubbleSort(sliceToSort))
	log.Println("Selection Sort:", SelectionSort(sliceToSort))
	log.Println("Insertion Sort:", InsertionSort(sliceToSort))
	log.Println("Merge Sort:", MergeSort(sliceToSort))
	log.Println("Quick Sort:", QuickSort(sliceToSort))
}

func BubbleSort(arr []int) []int {
	arr = append([]int{}, arr...)

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}

	return arr
}

func SelectionSort(arr []int) []int {
	arr = append([]int{}, arr...)

	size := len(arr)
	minIndex := 0

	for i := 0; i < size-1; i++ {
		minIndex = i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	}

	return arr
}

func InsertionSort(arr []int) []int {
	arr = append([]int{}, arr...)

	temp := 0
	size := len(arr)

	for i := 1; i < size; i++ {
		temp = arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}

	return arr
}

func MergeSort(arr []int) []int {
	arr = append([]int{}, arr...)
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	midIndex := int(math.Floor(float64(len(arr)) / 2))
	left := mergeSort(arr[:midIndex])
	right := mergeSort(arr[midIndex:])

	return merge(left, right)
}

func merge(arr1, arr2 []int) []int {
	var combined []int

	i := 0
	j := 0

	size1 := len(arr1)
	size2 := len(arr2)

	for i < size1 && j < size2 {
		if arr1[i] < arr2[j] {
			combined = append(combined, arr1[i])
			i++
		} else {
			combined = append(combined, arr2[j])
			j++
		}
	}

	for i < size1 {
		combined = append(combined, arr1[i])
		i++
	}

	for j < size2 {
		combined = append(combined, arr2[j])
		j++
	}

	return combined
}

func QuickSort(arr []int) []int {
	arr = append([]int{}, arr...)
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		pivotIndex := pivot(arr, left, right)
		quickSort(arr, left, pivotIndex-1)
		quickSort(arr, pivotIndex+1, right)
	}

	return arr
}

func pivot(arr []int, pIndex, endIndex int) int {
	swapIndex := pIndex

	for i := pIndex + 1; i <= endIndex; i++ {
		if arr[i] < arr[pIndex] {
			swapIndex++
			swap(arr, swapIndex, i)
		}
	}

	swap(arr, pIndex, swapIndex)
	return swapIndex
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
