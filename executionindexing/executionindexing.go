package executionindexing

import (
	"fmt"
)

// ExecutionIndex đại diện cho một index thực thi
type ExecutionIndex struct {
	Index      int                    // Thứ tự thực thi
	LineNumber int                    // Số dòng
	Statement  string                 // Câu lệnh
	Variables  map[string]interface{} // Giá trị biến
	ThreadID   int                    // Thread ID (cho đa luồng)
}

// ExecutionIndexer quản lý việc đánh index
type ExecutionIndexer struct {
	Indices      []ExecutionIndex
	CurrentIndex int
}

// NewExecutionIndexer tạo indexer mới
func NewExecutionIndexer() *ExecutionIndexer {
	return &ExecutionIndexer{
		Indices:      make([]ExecutionIndex, 0),
		CurrentIndex: 0,
	}
}

// RecordExecution ghi lại thực thi với index
func (ei *ExecutionIndexer) RecordExecution(lineNumber int, statement string, variables map[string]interface{}) int {
	ei.CurrentIndex++

	vars := make(map[string]interface{})
	for k, v := range variables {
		vars[k] = v
	}

	index := ExecutionIndex{
		Index:      ei.CurrentIndex,
		LineNumber: lineNumber,
		Statement:  statement,
		Variables:  vars,
		ThreadID:   1, // Single thread mặc định
	}

	ei.Indices = append(ei.Indices, index)
	return ei.CurrentIndex
}

// RecordExecutionWithThread ghi lại với thread ID
func (ei *ExecutionIndexer) RecordExecutionWithThread(lineNumber int, statement string, variables map[string]interface{}, threadID int) int {
	ei.CurrentIndex++

	vars := make(map[string]interface{})
	for k, v := range variables {
		vars[k] = v
	}

	index := ExecutionIndex{
		Index:      ei.CurrentIndex,
		LineNumber: lineNumber,
		Statement:  statement,
		Variables:  vars,
		ThreadID:   threadID,
	}

	ei.Indices = append(ei.Indices, index)
	return ei.CurrentIndex
}

// GetExecutionAtIndex lấy thực thi tại index cụ thể
func (ei *ExecutionIndexer) GetExecutionAtIndex(index int) *ExecutionIndex {
	for i := range ei.Indices {
		if ei.Indices[i].Index == index {
			return &ei.Indices[i]
		}
	}
	return nil
}

// GetExecutionsBetween lấy các thực thi trong khoảng index
func (ei *ExecutionIndexer) GetExecutionsBetween(startIndex, endIndex int) []ExecutionIndex {
	result := make([]ExecutionIndex, 0)
	for _, idx := range ei.Indices {
		if idx.Index >= startIndex && idx.Index <= endIndex {
			result = append(result, idx)
		}
	}
	return result
}

// FindExecutionsByLine tìm các thực thi theo dòng
func (ei *ExecutionIndexer) FindExecutionsByLine(lineNumber int) []ExecutionIndex {
	result := make([]ExecutionIndex, 0)
	for _, idx := range ei.Indices {
		if idx.LineNumber == lineNumber {
			result = append(result, idx)
		}
	}
	return result
}

// PrintExecutionIndex in execution index
func (ei *ExecutionIndexer) PrintExecutionIndex() {
	fmt.Println("\nEXECUTION INDEX - DANH CHI MUC THUC THI:")
	fmt.Println("+-------+------+------+----------------------------------------+-----------------+")
	fmt.Println("| Index | Dong | TID  | Cau lenh                               | Bien            |")
	fmt.Println("+-------+------+------+----------------------------------------+-----------------+")

	for _, idx := range ei.Indices {
		varsStr := ""
		for k, v := range idx.Variables {
			if varsStr != "" {
				varsStr += ", "
			}
			varsStr += fmt.Sprintf("%s=%v", k, v)
		}

		fmt.Printf("| %-5d | %-4d | T%-3d | %-38s | %-15s |\n",
			idx.Index,
			idx.LineNumber,
			idx.ThreadID,
			truncate(idx.Statement, 38),
			truncate(varsStr, 15))
	}

	fmt.Println("+-------+------+------+----------------------------------------+-----------------+")
	fmt.Printf("\nTong so execution index: %d\n", len(ei.Indices))
}

// PrintExecutionsByLine in các thực thi theo dòng
func (ei *ExecutionIndexer) PrintExecutionsByLine(lineNumber int) {
	executions := ei.FindExecutionsByLine(lineNumber)

	fmt.Printf("\nTim cac lan thuc thi dong %d:\n", lineNumber)
	fmt.Println("+-------+----------------------------------------+-----------------+")
	fmt.Println("| Index | Cau lenh                               | Bien            |")
	fmt.Println("+-------+----------------------------------------+-----------------+")

	for _, idx := range executions {
		varsStr := ""
		for k, v := range idx.Variables {
			if varsStr != "" {
				varsStr += ", "
			}
			varsStr += fmt.Sprintf("%s=%v", k, v)
		}

		fmt.Printf("| %-5d | %-38s | %-15s |\n",
			idx.Index,
			truncate(idx.Statement, 38),
			truncate(varsStr, 15))
	}

	fmt.Println("+-------+----------------------------------------+-----------------+")
	fmt.Printf("Tong: %d lan thuc thi\n", len(executions))
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// RunExecutionIndexingDemo chạy demo
func RunExecutionIndexingDemo() {
	fmt.Println("\nVi du: Tinh tong cac so tu 1 den n")
	fmt.Println("\nChuong trinh mau:")
	fmt.Println("```")
	fmt.Println("1: n := 5")
	fmt.Println("2: sum := 0")
	fmt.Println("3: for i := 1; i <= n; i++ {")
	fmt.Println("4:     sum = sum + i")
	fmt.Println("5: }")
	fmt.Println("6: print(sum)")
	fmt.Println("```")

	indexer := NewExecutionIndexer()

	// Simulate execution
	n := 5
	sum := 0

	indexer.RecordExecution(1, "n := 5", map[string]interface{}{"n": n})
	indexer.RecordExecution(2, "sum := 0", map[string]interface{}{"sum": sum})

	for i := 1; i <= n; i++ {
		indexer.RecordExecution(3, fmt.Sprintf("for i := %d", i), map[string]interface{}{
			"i": i, "n": n, "sum": sum,
		})

		sum = sum + i
		indexer.RecordExecution(4, fmt.Sprintf("sum = %d + %d", sum-i, i), map[string]interface{}{
			"i": i, "sum": sum,
		})
	}

	indexer.RecordExecution(6, fmt.Sprintf("print(%d)", sum), map[string]interface{}{"sum": sum})

	// In execution index
	indexer.PrintExecutionIndex()

	// Phân tích
	fmt.Println("\nPhan tich:")
	fmt.Println("- Moi lan thuc thi duoc gan mot index duy nhat")
	fmt.Println("- Index tang dan theo thu tu thoi gian")
	fmt.Println("- Co the tra cuu trang thai tai bat ky thoi diem nao")

	// Tìm các lần thực thi dòng 4
	fmt.Println("\n==============================================")
	indexer.PrintExecutionsByLine(4)

	// Demo query
	fmt.Println("\n==============================================")
	fmt.Println("\nQuery: Lay thuc thi tai index 7")
	exec := indexer.GetExecutionAtIndex(7)
	if exec != nil {
		fmt.Printf("  Dong %d: %s\n", exec.LineNumber, exec.Statement)
		fmt.Printf("  Bien: %v\n", exec.Variables)
	}

	fmt.Println("\nQuery: Lay thuc thi tu index 5 den 10")
	execs := indexer.GetExecutionsBetween(5, 10)
	fmt.Printf("  Tim thay %d thuc thi\n", len(execs))
	for _, e := range execs {
		fmt.Printf("  [%d] Dong %d: %s\n", e.Index, e.LineNumber, e.Statement)
	}

	fmt.Println("\nUng dung Execution Indexing:")
	fmt.Println("+ Debugging: Xac dinh chinh xac thoi diem loi xay ra")
	fmt.Println("+ Time-travel debugging: Quay lai trang thai truoc do")
	fmt.Println("+ Multi-threading: Phan tich interleaving")
	fmt.Println("+ Record & Replay: Tai hien lai execution")
}

func demo2() {
	// Removed - chỉ giữ 1 ví dụ
}

func demo3() {
	// Removed - chỉ giữ 1 ví dụ
}

