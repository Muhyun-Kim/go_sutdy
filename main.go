package main

import (
	"fmt"
	"os"

	"github.com/Muhyun-Kim/gocsvmapper"
)

func main() {
	// 사전에 정의된 컬럼 매핑
	columnMap := map[string]string{
		"학생": "student",
		"나이": "age",
	}

	// CSV 파일 열기
	file, err := os.Open("./example.csv") // gocsvmapper 폴더에 있는 example.csv 경로
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// CSV 파일 처리
	processor := gocsvmapper.NewCSVColumnProcessor(columnMap)
	records, err := processor.MapCSVColumns(file)
	if err != nil {
		fmt.Println("Error processing CSV:", err)
		return
	}

	// 처리된 CSV를 문자열로 출력
	csvString := gocsvmapper.CSVToString(records)
	fmt.Println(csvString)
}
