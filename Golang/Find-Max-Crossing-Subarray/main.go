package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Tupla struct {
	low   int
	high  int
	value float64
}

func find_maximum_crossing_subarray(a []float64, low int, high int) Tupla {
	if low == high {
		return Tupla{low, high, a[low]}
	} else {
		mid := int(math.Floor(float64(low+high) / 2))
		left := find_maximum_crossing_subarray(a, low, mid)
		right := find_maximum_crossing_subarray(a, mid+1, high)
		cross := find_max_crossing_subarray(a, low, mid, high)
		if left.value >= right.value && left.value >= cross.value {
			return left
		} else if right.value >= left.value && right.value >= cross.value {
			return right
		} else {
			return cross
		}
	}
}

func find_max_crossing_subarray(a []float64, low int, mid int, high int) Tupla { //defining return type
	left_sum := -math.MaxFloat64
	var sum float64 = 0
	max_left := -math.MaxInt64
	max_right := -math.MaxInt64
	for i := mid; i >= low; i-- {
		sum = sum + a[i]
		if sum > left_sum {
			left_sum = sum
			max_left = i
		}
	}
	right_sum := -math.MaxFloat64
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum = sum + a[j]
		if sum > right_sum {
			right_sum = sum
			max_right = j
		}
	}
	sum = left_sum + right_sum
	return Tupla{max_left, max_right, sum}
}

func generate_changes_array(a []float64) []float64 {
	changes := []float64{}
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
	array := []float64{}
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
			j, err := strconv.ParseFloat(s[i], 64) //converting to float64
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
			fmt.Println("Buy in ", result.low, " price: ", array[result.low])
			fmt.Println("Sale in ", result.high+1, " price: ", array[result.high+1])
			fmt.Println("Earnings per share: ", result.value)
		}
		fmt.Println("----------------------")
	}
}
