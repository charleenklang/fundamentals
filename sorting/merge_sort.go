package main

// Merge Sort
//
// Merge Sort is a divide-and-conquer algorithm that sorts an array by recursively dividing it into two halves,
// sorting each half, and then merging the sorted halves back together. Here’s a step-by-step explanation:
//
// Divide: Split the array into two halves. If the array has only one element or is empty, it is already sorted.
// Conquer: Recursively apply the Merge Sort algorithm to each half to sort them.
// Merge: Combine the two sorted halves into a single sorted array.
//
// Detailed Steps
// Splitting:
// Continuously divide the array into two halves until each sub-array contains only one element.
// This is the base case of the recursion, as a single-element array is inherently sorted.
//
// Sorting:
// Recursively sort each half. This involves further splitting each half until the base case is reached.
//
// Merging:
// Merge the sorted halves back together. This is done by comparing the smallest elements of each half and placing the smaller element into the new array. Repeat this process until all elements from both halves are merged into a single sorted array.
// Example
// Given the array [5, 3, 8, 4, 2]:
//
// Divide:
// Split into [5, 3, 8] and [4, 2].
// Further split [5, 3, 8] into [5] and [3, 8].
// Split [3, 8] into [3] and [8].
// Split [4, 2] into [4] and [2].
//
// Conquer:
// [5], [3], [8], [4], and [2] are already sorted.
//
// Merge:
// Merge [3] and [8] to get [3, 8].
// Merge [5] and [3, 8] to get [3, 5, 8].
// Merge [4] and [2] to get [2, 4].
// Finally, merge [3, 5, 8] and [2, 4] to get [2, 3, 4, 5, 8].
