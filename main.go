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
		fmt.Print("Nhap lua chon cua ban (1-5): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Println("\n========== TRACING - THEO DAU THUC THI ==========")
			tracing.RunTracingDemo()
		case "2":
			fmt.Println("\n========== DYNAMIC SLICING - CAT DONG ==========")
			dynamicslicing.RunDynamicSlicingDemo()
		case "3":
			fmt.Println("\n========== EXECUTION INDEXING - DANH CHI MUC ==========")
			executionindexing.RunExecutionIndexingDemo()
		case "4":
			fmt.Println("\n========== FAULT LOCALIZATION - DINH VI LOI ==========")
			faultlocalization.RunFaultLocalizationDemo()
		case "5":
			fmt.Println("\nCam on ban da su dung chuong trinh!")
			return
		default:
			fmt.Println("\nLua chon khong hop le! Vui long chon tu 1-5.")
		}

		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Print("\nNhan Enter de tiep tuc...")
		reader.ReadString('\n')
		fmt.Println()
	}
}

func printMenu() {
	fmt.Println("============================================================")
	fmt.Println("        PHAN TICH CHUONG TRINH - PROGRAM ANALYSIS")
	fmt.Println("============================================================")
	fmt.Println("  1. Tracing - Theo dau thuc thi chuong trinh")
	fmt.Println("  2. Dynamic Slicing - Cat dong chuong trinh")
	fmt.Println("  3. Execution Indexing - Danh chi muc thuc thi")
	fmt.Println("  4. Fault Localization - Dinh vi loi")
	fmt.Println("  5. Thoat")
	fmt.Println("============================================================")
}

