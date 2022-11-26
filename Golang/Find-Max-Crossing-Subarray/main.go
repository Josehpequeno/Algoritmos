package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func find_max_crossing_subarray(a []int, low int, mid int, high int) []int { //defining return type
	left_sum := -math.MaxInt64
	sum := 0
	max_left := -math.MaxInt64
	max_right := -math.MaxInt64
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if sum > left_sum {
			left_sum = sum
			max_left = i
		}
	}
	right_sum := -math.MaxInt64
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum = sum + a[j]
		if sum > right_sum {
			right_sum = sum
			max_right = j
		}
	}
	sum = left_sum + right_sum
	return []int{max_left, max_right, sum}
}

func find_maximum_crossing_subarray(a []int, low int, high int) []int {
	if low == high {
		return []int{low, high, a[low]}
	} else {
		mid := int(math.Floor(float64(low+high) / 2))
		left := find_maximum_crossing_subarray(a, low, mid)
		right := find_maximum_crossing_subarray(a, mid+1, high)
		cross := find_max_crossing_subarray(a, low, mid, high)
		left_low := left[0]
		left_high := left[1]
		left_sum := left[2]
		right_low := right[0]
		right_high := right[1]
		right_sum := right[2]
		cross_low := cross[0]
		cross_high := cross[1]
		cross_sum := cross[2]
		if left_sum >= right_sum && left_sum >= cross_sum {
			return []int{left_low, left_high, left_sum}
		} else if right_sum >= left_sum && right_sum >= cross_sum {
			return []int{right_low, right_high, right_sum}
		} else {
			return []int{cross_low, cross_high, cross_sum}
		}
	}
}

func generate_changes_array(a []int) []int {
	changes := []int{}
	for i := 1; i < len(a); i++ {
		val := a[i] - a[i-1]
		changes = append(changes, val)
	}
	return changes
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Find Max Crossing Subarray")
	fmt.Println("Enter the prices separated by spaces\n* type clear to clear array, quit to quit")
	fmt.Println("---------------------")
	array := []int{}
loop: //label loop
	for {
		fmt.Print("=> ")
		text, _ := reader.ReadString('\n') // reading the entries
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "clear" {
			array = array[:0] //clear the array
			continue loop
		}
		if text == "quit" {
			break loop
		}
		s := strings.Fields(text)
		for i := range s {
			j, err := strconv.Atoi(s[i]) //converting to int
			if err != nil {
				continue
			}
			array = append(array, j)
		}
		fmt.Println("----------------------")
		fmt.Println("Array to Find Max Crossing Subarray: ", array)
		changes := generate_changes_array(array)
		fmt.Println("Array of changes: ", changes)
		if len(array) > 1 {
			result := find_maximum_crossing_subarray(changes, 0, len(changes)-1)
			fmt.Println("Buy in ", result[0], " price: ", array[result[0]])
			fmt.Println("Sale in ", result[1]+1, " price: ", array[result[1]+1])
			fmt.Println("Earnings per share: ", result[2])
		}
		fmt.Println("----------------------")
	}
}
