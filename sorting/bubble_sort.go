package main

// Bubble Sort
//
// Bubble Sort is a simple sorting algorithm that repeatedly steps through the list,
// compares adjacent elements, and swaps them if they are in the wrong order. The
// process is repeated until the list is sorted.
//
// Given the list len(5), cap(5), [5, 3, 8, 4, 2]:
// 
// Compare 5 and 3, swap them:                  [3, 5, 8, 4, 2]
// Compare 5 and 8, no swap needed:             [3, 5, 8, 4, 2]
// Compare 8 and 4, swap them:                  [3, 5, 4, 8, 2]
// Compare 8 and 2, swap them:                  [3, 5, 4, 2, 8]
// Repeat the process until the list is sorted: [2, 3, 4, 5, 8]
//
// Time complexity of (O(n^2)) in the average and worst cases.
//
// my solution
// func BubbleSort(slice []int) []int {
// 	for j := 0; j < len(slice); j++ {
// 		for i := 0; i < len(slice)-1; i++ {
// 			if slice[i] > slice[i+1] {
// 				tmp := slice[i]
// 				slice[i] = slice[i+1]
// 				slice[i+1] = tmp
// 			} 
// 		}
// 	}
// 	return slice
// }

// optimal solution
func BubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}
