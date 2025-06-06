// 🔧 فایل generator.go بهینه‌شده
package utils

import (
	"math/rand"
	"time"
)

type Matrix [][]int64

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MakeMatrixInt(rows, cols int) Matrix {
	result := make(Matrix, rows)
	for i := range result {
		result[i] = make([]int64, cols)
	}
	return result
}

func RandomMatrix(rows, cols int) Matrix {
	result := MakeMatrixInt(rows, cols)
	for i := range result {
		for j := range result[i] {
			result[i][j] = rand.Int63n(32000)
		}
	}
	return result
}
