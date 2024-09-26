package main

// Insertion Sort
//
// Insertion Sort is a simple and intuitive sorting algorithm that builds
// the final sorted array one item at a time. It is much like sorting
// playing cards in your hands. Here’s a step-by-step explanation:
//
// Start with the first element: Consider the first element to be sorted.
// Pick the next element: Take the next element in the list.
// Compare with sorted elements: Compare the picked element with the elements in the sorted part of the list.
// Shift elements: Shift all the elements in the sorted part that are greater than the picked element to the right.
// Insert the picked element: Insert the picked element into its correct position.
// Repeat: Repeat steps 2-5 until all elements are sorted.
//
// Given the list [5, 3, 8, 4, 2]:
//
// Start with the first element [5].
// Pick the next element 3, compare with 5, and insert before 5: [3, 5].
// Pick the next element 8, compare with 5, and insert after 5: [3, 5, 8].
// Pick the next element 4, compare with 8 and 5, and insert before 5: [3, 4, 5, 8].
// Pick the next element 2, compare with 8, 5, 4, and 3, and insert before 3: [2, 3, 4, 5, 8].
//
// Time Complexity:
//
// Best Case: (O(n)) -> this happens when the array is already sorted and the array only needs to compare each element to each other.
// Average Case: (O(n^2))
// Worst Case: (O(n^2))

// my solution of InsertionSort
// Benchmark: 6.484 ns/op
// func InsertionSort(slice []int) []int {
// 	n := len(slice)
// 	for i := 1; i < n; i++ {
// 		if slice[i-1] > slice[i] {
// 			tmp := slice[i]
// 			index := i
// 			for j := i-1; j >= 0; j-- {
// 				if slice[j] > tmp {
// 					slice = append(slice[:index], slice[index+1:]...)
// 					slice = append(slice[:j+1], slice[j:]...)
// 					slice[j] = tmp
// 					index--
// 				} 
// 			}
// 		}
// 	}
// 	return slice
// }

// Optimal implementation of Insertion Sort
// Benchmark: 9.747 ns/op
// huh mine is better?
// but this one is more elegant :)
func InsertionSort(arr []int) []int {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1
        // Move elements of arr[0..i-1], that are greater than key, to one position ahead of their current position
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
    return arr
}
