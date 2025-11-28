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
	fmt.Println("\n📑 EXECUTION INDEX - ĐÁNH CHỈ MỤC THỰC THI:")
	fmt.Println("+-------+------+------+----------------------------------------+-----------------+")
	fmt.Println("| Index | Dòng | TID  | Câu lệnh                               | Biến            |")
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
	fmt.Printf("\nTổng số execution index: %d\n", len(ei.Indices))
}

// PrintExecutionsByLine in các thực thi theo dòng
func (ei *ExecutionIndexer) PrintExecutionsByLine(lineNumber int) {
	executions := ei.FindExecutionsByLine(lineNumber)

	fmt.Printf("\n🔍 Tìm các lần thực thi dòng %d:\n", lineNumber)
	fmt.Println("+-------+----------------------------------------+-----------------+")
	fmt.Println("| Index | Câu lệnh                               | Biến            |")
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
	fmt.Printf("Tổng: %d lần thực thi\n", len(executions))
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// RunExecutionIndexingDemo chạy demo
func RunExecutionIndexingDemo() {
	fmt.Println("\n🎯 Ví dụ 1: Tính tổng các số từ 1 đến n")
	fmt.Println("\nChương trình mẫu:")
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
	fmt.Println("\n📊 Phân tích:")
	fmt.Println("- Mỗi lần thực thi được gán một index duy nhất")
	fmt.Println("- Index tăng dần theo thứ tự thời gian")
	fmt.Println("- Có thể tra cứu trạng thái tại bất kỳ thời điểm nào")

	// Tìm các lần thực thi dòng 4
	fmt.Println("\n" + "==============================================")
	indexer.PrintExecutionsByLine(4)

	// Demo query
	fmt.Println("\n" + "==============================================")
	fmt.Println("\n🔎 Query: Lấy thực thi tại index 7")
	exec := indexer.GetExecutionAtIndex(7)
	if exec != nil {
		fmt.Printf("  Dòng %d: %s\n", exec.LineNumber, exec.Statement)
		fmt.Printf("  Biến: %v\n", exec.Variables)
	}

	fmt.Println("\n🔎 Query: Lấy thực thi từ index 5 đến 10")
	execs := indexer.GetExecutionsBetween(5, 10)
	fmt.Printf("  Tìm thấy %d thực thi\n", len(execs))
	for _, e := range execs {
		fmt.Printf("  [%d] Dòng %d: %s\n", e.Index, e.LineNumber, e.Statement)
	}

	// Demo 2: Nested loops
	fmt.Println("\n" + "==============================================")
	demo2()

	// Demo 3: Multi-threading simulation
	fmt.Println("\n" + "==============================================")
	demo3()
}

func demo2() {
	fmt.Println("\n🎯 Ví dụ 2: Vòng lặp lồng nhau (nested loops)")
	fmt.Println("\nChương trình mẫu:")
	fmt.Println("```")
	fmt.Println("1: for i := 1; i <= 3; i++ {")
	fmt.Println("2:     for j := 1; j <= 2; j++ {")
	fmt.Println("3:         print(i, j)")
	fmt.Println("4:     }")
	fmt.Println("5: }")
	fmt.Println("```")

	indexer := NewExecutionIndexer()

	for i := 1; i <= 3; i++ {
		indexer.RecordExecution(1, fmt.Sprintf("for i := %d", i), map[string]interface{}{"i": i})

		for j := 1; j <= 2; j++ {
			indexer.RecordExecution(2, fmt.Sprintf("for j := %d", j), map[string]interface{}{
				"i": i, "j": j,
			})
			indexer.RecordExecution(3, fmt.Sprintf("print(%d, %d)", i, j), map[string]interface{}{
				"i": i, "j": j,
			})
		}
	}

	indexer.PrintExecutionIndex()

	fmt.Println("\n💡 Lợi ích:")
	fmt.Println("- Execution Index giúp phân biệt các lần thực thi khác nhau")
	fmt.Println("- Hữu ích với loops: cùng dòng code nhưng giá trị khác nhau")
	fmt.Println("- Có thể replay chương trình từ bất kỳ index nào")
}

func demo3() {
	fmt.Println("\n🎯 Ví dụ 3: Mô phỏng đa luồng (Multi-threading)")
	fmt.Println("\nChương trình mẫu (2 threads):")
	fmt.Println("```")
	fmt.Println("Thread 1:")
	fmt.Println("1: x := 0")
	fmt.Println("2: x = x + 1")
	fmt.Println("3: x = x + 2")
	fmt.Println("")
	fmt.Println("Thread 2:")
	fmt.Println("4: y := 0")
	fmt.Println("5: y = y + 10")
	fmt.Println("6: y = y + 20")
	fmt.Println("```")

	indexer := NewExecutionIndexer()

	// Mô phỏng interleaving execution
	x := 0
	y := 0

	// Thread 1 bắt đầu
	indexer.RecordExecutionWithThread(1, "x := 0", map[string]interface{}{"x": x}, 1)

	// Thread 2 bắt đầu
	indexer.RecordExecutionWithThread(4, "y := 0", map[string]interface{}{"y": y}, 2)

	// Thread 1 tiếp tục
	x = x + 1
	indexer.RecordExecutionWithThread(2, "x = x + 1", map[string]interface{}{"x": x}, 1)

	// Thread 2 tiếp tục
	y = y + 10
	indexer.RecordExecutionWithThread(5, "y = y + 10", map[string]interface{}{"y": y}, 2)

	// Thread 1 kết thúc
	x = x + 2
	indexer.RecordExecutionWithThread(3, "x = x + 2", map[string]interface{}{"x": x}, 1)

	// Thread 2 kết thúc
	y = y + 20
	indexer.RecordExecutionWithThread(6, "y = y + 20", map[string]interface{}{"y": y}, 2)

	indexer.PrintExecutionIndex()

	fmt.Println("\n💡 Ứng dụng Execution Indexing:")
	fmt.Println("✓ Debugging: Xác định chính xác thời điểm lỗi xảy ra")
	fmt.Println("✓ Time-travel debugging: Quay lại trạng thái trước đó")
	fmt.Println("✓ Multi-threading: Phân tích interleaving")
	fmt.Println("✓ Record & Replay: Tái hiện lại execution")
	fmt.Println("✓ Distributed systems: Ordering events")
}

