package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	array  [][]float64
	rows   int
	colums int
}

func square_matrix_mutiply_recursive(a Matrix, b Matrix) [][]float64 {
	n := a.rows
	c := Matrix{make([][]float64, n), n, n}
	if n == 1 {
		c.array[0] = append(c.array[0], a.array[0][0]*b.array[0][0])
	} else {
		m1 := Matrix{[][]float64{{a.array[0][0]}}, 1, 1}
		m2 := Matrix{[][]float64{{b.array[0][0]}}, 1, 1}
		m3 := Matrix{[][]float64{{a.array[0][1]}}, 1, 1}
		m4 := Matrix{[][]float64{{b.array[1][0]}}, 1, 1}
		//c[0][0] = a[0][0]* b[0][0] + a[0][1]* b[1][0]
		c.array[0] = append(c.array[0], square_matrix_mutiply_recursive(m1, m2)[0][0]+square_matrix_mutiply_recursive(m3, m4)[0][0])
		m1 = Matrix{[][]float64{{a.array[0][0]}}, 1, 1}
		m2 = Matrix{[][]float64{{b.array[0][1]}}, 1, 1}
		m3 = Matrix{[][]float64{{a.array[0][1]}}, 1, 1}
		m4 = Matrix{[][]float64{{b.array[1][1]}}, 1, 1}
		//c[0][1] = a[0][0]* b[0][1] + a[0][1]* b[1][1]
		c.array[0] = append(c.array[0], square_matrix_mutiply_recursive(m1, m2)[0][0]+square_matrix_mutiply_recursive(m3, m4)[0][0])
		m1 = Matrix{[][]float64{{a.array[1][0]}}, 1, 1}
		m2 = Matrix{[][]float64{{b.array[0][0]}}, 1, 1}
		m3 = Matrix{[][]float64{{a.array[1][1]}}, 1, 1}
		m4 = Matrix{[][]float64{{b.array[1][0]}}, 1, 1}
		c.array[1] = append(c.array[1], square_matrix_mutiply_recursive(m1, m2)[0][0]+square_matrix_mutiply_recursive(m3, m4)[0][0])
		m1 = Matrix{[][]float64{{a.array[1][0]}}, 1, 1}
		m2 = Matrix{[][]float64{{b.array[0][1]}}, 1, 1}
		m3 = Matrix{[][]float64{{a.array[1][1]}}, 1, 1}
		m4 = Matrix{[][]float64{{b.array[1][1]}}, 1, 1}
		c.array[1] = append(c.array[1], square_matrix_mutiply_recursive(m1, m2)[0][0]+square_matrix_mutiply_recursive(m3, m4)[0][0])
	}
	return c.array
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Square Matrix Mutiply 2x2")
	fmt.Println("Enter the number of square matrix separated by spaces\n* type quit to quit")
	fmt.Println("---------------------")
	first := []float64{}
	second := []float64{}
	a_matrix := Matrix{}
	b_matrix := Matrix{}
loop: //label loop
	for {
	first_loop:
		for {
			fmt.Print("First Matrix => ")
			text, _ := reader.ReadString('\n') // reading the entries
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			if text == "quit" {
				break loop
			}
			s := strings.Fields(text)
			for i := range s {
				if len(first) >= 4 {
					break
				}
				j, err := strconv.ParseFloat(s[i], 64) //converting to float64
				if err != nil {
					continue
				}
				first = append(first, j)
			}
			sqrt := math.Sqrt(float64(len(first)))
			trunc := math.Trunc(sqrt)
			if sqrt == trunc {
				a_matrix.array = make([][]float64, int(sqrt))
				a_matrix.rows = int(sqrt)
				a_matrix.colums = int(sqrt)
				count := 0
				it := 0
				for i := 0; i < len(first); i++ {
					if count > int(sqrt)-1 {
						it++
						count = 0
					}
					a_matrix.array[it] = append(a_matrix.array[it], first[i])
					count++
				}
				fmt.Println("First Matrix", a_matrix.array)
				fmt.Println("----------------------")
				if len(a_matrix.array) > 0 {
					break first_loop
				}
			} else {
				fmt.Println("Square matrix is not a square matrix", first)
			}
		}
	second_loop:
		for {
			fmt.Print("Second Matrix => ")
			text, _ := reader.ReadString('\n') // reading the entries
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)
			if text == "quit" {
				break loop
			}
			s := strings.Fields(text)
			for i := range s {
				if len(second) >= 4 {
					break
				}
				j, err := strconv.ParseFloat(s[i], 64) //converting to float64
				if err != nil {
					continue
				}
				second = append(second, j)
			}
			sqrt := math.Sqrt(float64(len(second)))
			trunc := math.Trunc(sqrt)
			if sqrt == trunc {
				b_matrix.array = make([][]float64, int(sqrt))
				b_matrix.rows = int(sqrt)
				b_matrix.colums = int(sqrt)
				count := 0
				it := 0
				for i := 0; i < len(second); i++ {
					if count > int(sqrt)-1 {
						it++
						count = 0
					}
					b_matrix.array[it] = append(b_matrix.array[it], second[i])
					count++
				}
				fmt.Println("Second Matrix", b_matrix.array)
				fmt.Println("----------------------")
				if len(b_matrix.array) > 0 {
					break second_loop
				}
			} else {
				fmt.Println("Square matrix is not a square matrix", second)
			}
		}
		if a_matrix.colums != b_matrix.colums {
			fmt.Println("Square matrices with different number of columns")
		} else if a_matrix.rows != b_matrix.rows {
			fmt.Println("Square matrices with different number of rows")
		} else {
			fmt.Println("Result Matrix", square_matrix_mutiply_recursive(a_matrix, b_matrix))
		}
		fmt.Println("----------------------")
		first = first[:0] //clear the matrix
		second = second[:0]
	}
}
