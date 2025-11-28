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
	fmt.Println("\nDYNAMIC SLICE:")
	fmt.Printf("Slicing Criterion: <%d, %s>\n\n",
		ds.SlicingCriterion.LineNumber,
		ds.SlicingCriterion.Variable)

	fmt.Println("Cac dong anh huong den bien '" + ds.SlicingCriterion.Variable + "':")
	fmt.Println("+------+--------------------------------------------------+")
	fmt.Println("| Dong | Cau lenh                                         |")
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
	fmt.Printf("\nSo dong trong slice: %d\n", len(slice))
	fmt.Printf("Tong so dong thuc thi: %d\n", len(ds.ExecutionTrace))
	fmt.Printf("Ti le rut gon: %.1f%%\n", float64(len(slice))*100/float64(len(ds.ExecutionTrace)))
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// RunDynamicSlicingDemo chạy demo
func RunDynamicSlicingDemo() {
	fmt.Println("\nVi du: Tinh tong va tich cua day so")
	fmt.Println("\nChuong trinh mau:")
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
	fmt.Println("\nThuc thi chuong trinh:")
	slicer.RecordExecution(1)
	fmt.Println("  Dong 1: sum := 0")

	slicer.RecordExecution(2)
	fmt.Println("  Dong 2: product := 1")

	for i := 1; i <= 5; i++ {
		slicer.RecordExecution(3)
		fmt.Printf("  Dong 3: i = %d\n", i)

		slicer.RecordExecution(4)
		fmt.Printf("  Dong 4: sum = sum + %d\n", i)

		slicer.RecordExecution(5)
		fmt.Printf("  Dong 5: product = product * %d\n", i)
	}

	slicer.RecordExecution(7)
	fmt.Println("  Dong 7: result := sum + product")

	slicer.RecordExecution(8)
	fmt.Println("  Dong 8: print(result)")

	// Tính slice cho biến 'sum' tại dòng 7
	fmt.Println("\n==============================================")
	fmt.Println("\nSlice cho bien 'sum' tai dong 7:")
	slice1 := slicer.ComputeSlice(7, "sum")
	slicer.PrintSlice(slice1)

	fmt.Println("\nGiai thich:")
	fmt.Println("- Cac dong 1, 3, 4 anh huong den gia tri cua 'sum'")
	fmt.Println("- Dong 5 (lien quan den product) KHONG co trong slice")
	fmt.Println("- Dynamic slicing giup loai bo code khong lien quan")

	fmt.Println("\nUng dung cua Dynamic Slicing:")
	fmt.Println("+ Debugging: Tim code anh huong den bug")
	fmt.Println("+ Program comprehension: Hieu code de hon")
	fmt.Println("+ Testing: Tap trung test case vao phan lien quan")
	fmt.Println("+ Maintenance: Danh gia impact khi thay doi code")
}

func demo2() {
	// Removed - chỉ giữ 1 ví dụ
}

