// 🔧 فایل main.go بهینه‌شده
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"matrixMultiplexing/utils"
	"os"
	"strconv"
	"time"
)

const (
	logFile  = "results.csv"
	maxLimit = 5000
)

func operation(n int) {
	matrixA := utils.RandomMatrix(n, n)
	matrixB := utils.RandomMatrix(n, n)

	start := time.Now()
	_, err := utils.CrossInt(matrixA, matrixB)
	if err != nil {
		log.Fatalf("خطا در ضرب معمولی: %v", err)
	}
	crossTime := time.Since(start).Seconds()

	start = time.Now()
	_ = utils.StrassenTop(matrixA, matrixB)
	strassenTime := time.Since(start).Seconds()

	fmt.Printf("n=%d | Threshold: %d | Cross: %.6f s | Strassen: %.6f s\n", n, utils.Threshold, crossTime, strassenTime)
	appendToCSV(logFile, n, utils.Threshold, crossTime, strassenTime)
}

func appendToCSV(filename string, n, Threshold int, cross, strassen float64) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("خطا در باز کردن فایل لاگ: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	info, _ := file.Stat()
	if info.Size() == 0 {
		_ = writer.Write([]string{"n", "threshold", "cross_time_s", "strassen_time_s"})
	}

	_ = writer.Write([]string{
		strconv.Itoa(n),
		strconv.Itoa(Threshold),
		fmt.Sprintf("%.6f", cross),
		fmt.Sprintf("%.6f", strassen),
	})

	writer.Flush()
}

func getNextNFromCSV(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 1
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil || len(records) < 2 {
		return 1
	}

	last := records[len(records)-1]
	n, err := strconv.Atoi(last[0])
	if err != nil {
		return 1
	}
	return n + 1
}

func main() {
	startN := getNextNFromCSV(logFile)
	for n := startN; n <= maxLimit; n += 100 {
		operation(n)
	}
}
