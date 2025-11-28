package dynamicslicing

import (
	"fmt"
)

// Statement đại diện cho một câu lệnh
type Statement struct {
	LineNumber int
	Code       string
	Defines    []string // Biến được định nghĩa
	Uses       []string // Biến được sử dụng
	Executed   bool     // Có được thực thi không
}

// DynamicSlicer thực hiện dynamic slicing
type DynamicSlicer struct {
	Statements       []Statement
	ExecutionTrace   []int // Danh sách line number đã thực thi
	SlicingCriterion struct {
		LineNumber int
		Variable   string
	}
}

// NewDynamicSlicer tạo slicer mới
func NewDynamicSlicer() *DynamicSlicer {
	return &DynamicSlicer{
		Statements:     make([]Statement, 0),
		ExecutionTrace: make([]int, 0),
	}
}

// AddStatement thêm một câu lệnh
func (ds *DynamicSlicer) AddStatement(lineNumber int, code string, defines []string, uses []string) {
	ds.Statements = append(ds.Statements, Statement{
		LineNumber: lineNumber,
		Code:       code,
		Defines:    defines,
		Uses:       uses,
		Executed:   false,
	})
}

// RecordExecution ghi lại thực thi
func (ds *DynamicSlicer) RecordExecution(lineNumber int) {
	ds.ExecutionTrace = append(ds.ExecutionTrace, lineNumber)
	for i := range ds.Statements {
		if ds.Statements[i].LineNumber == lineNumber {
			ds.Statements[i].Executed = true
			break
		}
	}
}

// ComputeSlice tính dynamic slice
func (ds *DynamicSlicer) ComputeSlice(lineNumber int, variable string) []int {
	ds.SlicingCriterion.LineNumber = lineNumber
	ds.SlicingCriterion.Variable = variable

	slice := make([]int, 0)
	relevantVars := make(map[string]bool)
	relevantVars[variable] = true

	// Duyệt ngược execution trace
	for i := len(ds.ExecutionTrace) - 1; i >= 0; i-- {
		lineNum := ds.ExecutionTrace[i]

		// Tìm statement tương ứng
		var stmt *Statement
		for j := range ds.Statements {
			if ds.Statements[j].LineNumber == lineNum {
				stmt = &ds.Statements[j]
				break
			}
		}

		if stmt == nil {
			continue
		}

		// Nếu dòng này sau slicing criterion, bỏ qua
		if i > findLastOccurrence(ds.ExecutionTrace, lineNumber) {
			continue
		}

		// Kiểm tra xem statement có định nghĩa biến liên quan không
		hasRelevantDef := false
		for _, def := range stmt.Defines {
			if relevantVars[def] {
				hasRelevantDef = true
				break
			}
		}

		if hasRelevantDef {
			// Thêm vào slice
			if !contains(slice, lineNum) {
				slice = append([]int{lineNum}, slice...)
			}

			// Thêm các biến được sử dụng vào tập relevant
			for _, use := range stmt.Uses {
				relevantVars[use] = true
			}
		}
	}

	return slice
}

// findLastOccurrence tìm vị trí cuối cùng của line trong trace
func findLastOccurrence(trace []int, line int) int {
	for i := len(trace) - 1; i >= 0; i-- {
		if trace[i] == line {
			return i
		}
	}
	return -1
}

// contains kiểm tra xem slice có chứa giá trị không
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// PrintSlice in dynamic slice
func (ds *DynamicSlicer) PrintSlice(slice []int) {
	fmt.Println("\n🔪 DYNAMIC SLICE:")
	fmt.Printf("Slicing Criterion: <%d, %s>\n\n",
		ds.SlicingCriterion.LineNumber,
		ds.SlicingCriterion.Variable)

	fmt.Println("Các dòng ảnh hưởng đến biến '" + ds.SlicingCriterion.Variable + "':")
	fmt.Println("+------+--------------------------------------------------+")
	fmt.Println("| Dòng | Câu lệnh                                         |")
	fmt.Println("+------+--------------------------------------------------+")

	for _, lineNum := range slice {
		for _, stmt := range ds.Statements {
			if stmt.LineNumber == lineNum {
				fmt.Printf("| %-4d | %-48s |\n", lineNum, truncate(stmt.Code, 48))
				break
			}
		}
	}

	fmt.Println("+------+--------------------------------------------------+")
	fmt.Printf("\nSố dòng trong slice: %d\n", len(slice))
	fmt.Printf("Tổng số dòng thực thi: %d\n", len(ds.ExecutionTrace))
	fmt.Printf("Tỷ lệ rút gọn: %.1f%%\n", float64(len(slice))*100/float64(len(ds.ExecutionTrace)))
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// RunDynamicSlicingDemo chạy demo
func RunDynamicSlicingDemo() {
	fmt.Println("\n🎯 Ví dụ: Tính tổng và tích của dãy số")
	fmt.Println("\nChương trình mẫu:")
	fmt.Println("```")
	fmt.Println(" 1: sum := 0")
	fmt.Println(" 2: product := 1")
	fmt.Println(" 3: for i := 1; i <= 5; i++ {")
	fmt.Println(" 4:     sum = sum + i")
	fmt.Println(" 5:     product = product * i")
	fmt.Println(" 6: }")
	fmt.Println(" 7: result := sum + product")
	fmt.Println(" 8: print(result)")
	fmt.Println("```")

	// Tạo slicer
	slicer := NewDynamicSlicer()

	// Thêm các statements
	slicer.AddStatement(1, "sum := 0", []string{"sum"}, []string{})
	slicer.AddStatement(2, "product := 1", []string{"product"}, []string{})
	slicer.AddStatement(3, "for i := 1; i <= 5; i++", []string{"i"}, []string{})
	slicer.AddStatement(4, "sum = sum + i", []string{"sum"}, []string{"sum", "i"})
	slicer.AddStatement(5, "product = product * i", []string{"product"}, []string{"product", "i"})
	slicer.AddStatement(6, "}", []string{}, []string{})
	slicer.AddStatement(7, "result := sum + product", []string{"result"}, []string{"sum", "product"})
	slicer.AddStatement(8, "print(result)", []string{}, []string{"result"})

	// Ghi lại execution trace
	fmt.Println("\n▶ Thực thi chương trình:")
	slicer.RecordExecution(1)
	fmt.Println("  Dòng 1: sum := 0")

	slicer.RecordExecution(2)
	fmt.Println("  Dòng 2: product := 1")

	for i := 1; i <= 5; i++ {
		slicer.RecordExecution(3)
		fmt.Printf("  Dòng 3: i = %d\n", i)

		slicer.RecordExecution(4)
		fmt.Printf("  Dòng 4: sum = sum + %d\n", i)

		slicer.RecordExecution(5)
		fmt.Printf("  Dòng 5: product = product * %d\n", i)
	}

	slicer.RecordExecution(7)
	fmt.Println("  Dòng 7: result := sum + product")

	slicer.RecordExecution(8)
	fmt.Println("  Dòng 8: print(result)")

	// Tính slice cho biến 'sum' tại dòng 7
	fmt.Println("\n" + "==============================================")
	fmt.Println("\n📊 Case 1: Slice cho biến 'sum' tại dòng 7")
	slice1 := slicer.ComputeSlice(7, "sum")
	slicer.PrintSlice(slice1)

	fmt.Println("\n💡 Giải thích:")
	fmt.Println("- Các dòng 1, 3, 4 ảnh hưởng đến giá trị của 'sum'")
	fmt.Println("- Dòng 5 (liên quan đến product) KHÔNG có trong slice")
	fmt.Println("- Dynamic slicing giúp loại bỏ code không liên quan")

	// Tính slice cho biến 'product' tại dòng 7
	fmt.Println("\n" + "==============================================")
	fmt.Println("\n📊 Case 2: Slice cho biến 'product' tại dòng 7")
	slice2 := slicer.ComputeSlice(7, "product")
	slicer.PrintSlice(slice2)

	fmt.Println("\n💡 Giải thích:")
	fmt.Println("- Các dòng 2, 3, 5 ảnh hưởng đến giá trị của 'product'")
	fmt.Println("- Dòng 4 (liên quan đến sum) KHÔNG có trong slice")
	fmt.Println("- Slice khác nhau tùy thuộc vào biến quan tâm")

	// Ví dụ 2: Tìm max với điều kiện
	fmt.Println("\n" + "==============================================")
	demo2()
}

func demo2() {
	fmt.Println("\n🎯 Ví dụ 2: Tìm số chẵn lớn nhất")
	fmt.Println("\nChương trình mẫu:")
	fmt.Println("```")
	fmt.Println(" 1: arr := [3, 8, 2, 9, 4]")
	fmt.Println(" 2: max := -1")
	fmt.Println(" 3: found := false")
	fmt.Println(" 4: for i := 0; i < len(arr); i++ {")
	fmt.Println(" 5:     if arr[i] % 2 == 0 {")
	fmt.Println(" 6:         if arr[i] > max {")
	fmt.Println(" 7:             max = arr[i]")
	fmt.Println(" 8:             found = true")
	fmt.Println(" 9:         }")
	fmt.Println("10:     }")
	fmt.Println("11: }")
	fmt.Println("12: print(max)")
	fmt.Println("```")

	slicer := NewDynamicSlicer()

	// Thêm statements
	slicer.AddStatement(1, "arr := [3, 8, 2, 9, 4]", []string{"arr"}, []string{})
	slicer.AddStatement(2, "max := -1", []string{"max"}, []string{})
	slicer.AddStatement(3, "found := false", []string{"found"}, []string{})
	slicer.AddStatement(4, "for i := 0; i < len(arr); i++", []string{"i"}, []string{"arr"})
	slicer.AddStatement(5, "if arr[i] % 2 == 0", []string{}, []string{"arr", "i"})
	slicer.AddStatement(6, "if arr[i] > max", []string{}, []string{"arr", "i", "max"})
	slicer.AddStatement(7, "max = arr[i]", []string{"max"}, []string{"arr", "i"})
	slicer.AddStatement(8, "found = true", []string{"found"}, []string{})
	slicer.AddStatement(12, "print(max)", []string{}, []string{"max"})

	// Simulate execution
	fmt.Println("\n▶ Thực thi:")
	arr := []int{3, 8, 2, 9, 4}
	slicer.RecordExecution(1)
	slicer.RecordExecution(2)
	slicer.RecordExecution(3)

	for i := 0; i < len(arr); i++ {
		slicer.RecordExecution(4)
		slicer.RecordExecution(5)

		if arr[i]%2 == 0 {
			fmt.Printf("  i=%d, arr[%d]=%d (chẵn)\n", i, i, arr[i])
			slicer.RecordExecution(6)
			if arr[i] > -1 {
				slicer.RecordExecution(7)
				slicer.RecordExecution(8)
			}
		}
	}
	slicer.RecordExecution(12)

	fmt.Println("\n📊 Slice cho biến 'max' tại dòng 12:")
	slice := slicer.ComputeSlice(12, "max")
	slicer.PrintSlice(slice)

	fmt.Println("\n💡 Ứng dụng của Dynamic Slicing:")
	fmt.Println("✓ Debugging: Tìm code ảnh hưởng đến bug")
	fmt.Println("✓ Program comprehension: Hiểu code dễ hơn")
	fmt.Println("✓ Testing: Tập trung test case vào phần liên quan")
	fmt.Println("✓ Maintenance: Đánh giá impact khi thay đổi code")
}

