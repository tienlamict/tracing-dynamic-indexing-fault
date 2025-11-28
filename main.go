package main

import (
	"bufio"
	"fmt"
	"os"
	"program-analysis-demo/dynamicslicing"
	"program-analysis-demo/executionindexing"
	"program-analysis-demo/faultlocalization"
	"program-analysis-demo/tracing"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()
		fmt.Print("Nhập lựa chọn của bạn (1-5): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("\n========== TRACING - THEO DẤU THỰC THI ==========")
			tracing.RunTracingDemo()
		case "2":
			fmt.Println("\n========== DYNAMIC SLICING - CẮT ĐỘNG ==========")
			dynamicslicing.RunDynamicSlicingDemo()
		case "3":
			fmt.Println("\n========== EXECUTION INDEXING - ĐÁNH CHỈ MỤC ==========")
			executionindexing.RunExecutionIndexingDemo()
		case "4":
			fmt.Println("\n========== FAULT LOCALIZATION - ĐỊNH VỊ LỖI ==========")
			faultlocalization.RunFaultLocalizationDemo()
		case "5":
			fmt.Println("\nCảm ơn bạn đã sử dụng chương trình!")
			return
		default:
			fmt.Println("\n❌ Lựa chọn không hợp lệ! Vui lòng chọn từ 1-5.")
		}

		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Print("\nNhấn Enter để tiếp tục...")
		reader.ReadString('\n')
		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║        PHÂN TÍCH CHƯƠNG TRÌNH - PROGRAM ANALYSIS          ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. Tracing - Theo dấu thực thi chương trình              ║")
	fmt.Println("║  2. Dynamic Slicing - Cắt động chương trình               ║")
	fmt.Println("║  3. Execution Indexing - Đánh chỉ mục thực thi            ║")
	fmt.Println("║  4. Fault Localization - Định vị lỗi                      ║")
	fmt.Println("║  5. Thoát                                                  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

