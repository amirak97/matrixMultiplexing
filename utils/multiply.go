// 🔧 فایل multiply.go بهینه‌شده با sync.Pool و عملیات in-place
package utils

import (
	"errors"
	"sync"
)

type MatrixInt = [][]int64

const Threshold = 128

var matPool = sync.Pool{
	New: func() interface{} {
		return MakeMatrixInt(0, 0)
	},
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nextPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

func CrossInt(A, B MatrixInt) (MatrixInt, error) {
	if len(A) == 0 || len(B) == 0 || len(A[0]) != len(B) {
		return nil, errors.New("matrix sizes are incompatible for multiplication")
	}
	x1, y1 := len(A), len(A[0])
	_, y2 := len(B), len(B[0])

	C := MakeMatrixInt(x1, y2)
	for i := 0; i < x1; i++ {
		for j := 0; j < y2; j++ {
			for k := 0; k < y1; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C, nil
}

func NormalizeCopy(A MatrixInt) MatrixInt {
	x, y := len(A), len(A[0])
	n := nextPowerOfTwo(max(x, y))
	if x == n && y == n {
		return A
	}

	newA := MakeMatrixInt(n, n)
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			newA[i][j] = A[i][j]
		}
	}
	return newA
}

func TrimMatrix(A MatrixInt, rows, cols int) MatrixInt {
	result := MakeMatrixInt(rows, cols)
	for i := 0; i < rows; i++ {
		copy(result[i], A[i][:cols])
	}
	return result
}

func StrassenTop(A, B MatrixInt) MatrixInt {
	if len(A) == 0 || len(B) == 0 || len(A[0]) != len(B) {
		return nil
	}
	An := NormalizeCopy(A)
	Bn := NormalizeCopy(B)
	Cn := Strassen(An, Bn)
	return TrimMatrix(Cn, len(A), len(B[0]))
}

func addInPlace(C, A, B MatrixInt) {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
}

func subInPlace(C, A, B MatrixInt) {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
}

func splitMat(A MatrixInt) (MatrixInt, MatrixInt, MatrixInt, MatrixInt) {
	n := len(A) / 2
	a11 := MakeMatrixInt(n, n)
	a12 := MakeMatrixInt(n, n)
	a21 := MakeMatrixInt(n, n)
	a22 := MakeMatrixInt(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a11[i][j] = A[i][j]
			a12[i][j] = A[i][j+n]
			a21[i][j] = A[i+n][j]
			a22[i][j] = A[i+n][j+n]
		}
	}
	return a11, a12, a21, a22
}

func Strassen(A, B MatrixInt) MatrixInt {
	n := len(A)
	if n <= Threshold {
		C, err := CrossInt(A, B)
		if err != nil {
			panic(err)
		}
		return C
	}

	a11, a12, a21, a22 := splitMat(A)
	b11, b12, b21, b22 := splitMat(B)

	p1 := Strassen(add(a11, a22), add(b11, b22))
	p2 := Strassen(add(a21, a22), b11)
	p3 := Strassen(a11, sub(b12, b22))
	p4 := Strassen(a22, sub(b21, b11))
	p5 := Strassen(add(a11, a12), b22)
	p6 := Strassen(sub(a21, a11), add(b11, b12))
	p7 := Strassen(sub(a12, a22), add(b21, b22))

	c11 := add(sub(add(p1, p4), p5), p7)
	c12 := add(p3, p5)
	c21 := add(p2, p4)
	c22 := sub(sub(add(p1, p3), p2), p6)

	C := MakeMatrixInt(n, n)
	half := n / 2
	for i := 0; i < half; i++ {
		for j := 0; j < half; j++ {
			C[i][j] = c11[i][j]
			C[i][j+half] = c12[i][j]
			C[i+half][j] = c21[i][j]
			C[i+half][j+half] = c22[i][j]
		}
	}
	return C
}

func add(A, B MatrixInt) MatrixInt {
	n := len(A)
	C := MakeMatrixInt(n, n)
	addInPlace(C, A, B)
	return C
}

func sub(A, B MatrixInt) MatrixInt {
	n := len(A)
	C := MakeMatrixInt(n, n)
	subInPlace(C, A, B)
	return C
}
