package tracing

import (
	"fmt"
	"time"
)

// TraceEntry đại diện cho một mục trong trace
type TraceEntry struct {
	LineNumber int
	Statement  string
	Variables  map[string]interface{}
	Timestamp  time.Time
}

// Tracer quản lý việc ghi lại trace
type Tracer struct {
	Entries []TraceEntry
	enabled bool
}

// NewTracer tạo một tracer mới
func NewTracer() *Tracer {
	return &Tracer{
		Entries: make([]TraceEntry, 0),
		enabled: true,
	}
}

// Trace ghi lại một câu lệnh
func (t *Tracer) Trace(lineNumber int, statement string, variables map[string]interface{}) {
	if !t.enabled {
		return
	}

	entry := TraceEntry{
		LineNumber: lineNumber,
		Statement:  statement,
		Variables:  make(map[string]interface{}),
		Timestamp:  time.Now(),
	}

	// Sao chép variables
	for k, v := range variables {
		entry.Variables[k] = v
	}

	t.Entries = append(t.Entries, entry)
}

// PrintTrace in ra trace đã ghi lại
func (t *Tracer) PrintTrace() {
	fmt.Println("\n📝 TRACE - THEO DẤU THỰC THI:")
	fmt.Println("+------+---------------------------------------------+------------------------+")
	fmt.Println("| Dòng | Câu lệnh                                    | Biến                   |")
	fmt.Println("+------+---------------------------------------------+------------------------+")

	for _, entry := range t.Entries {
		varsStr := ""
		for k, v := range entry.Variables {
			if varsStr != "" {
				varsStr += ", "
			}
			varsStr += fmt.Sprintf("%s=%v", k, v)
		}

		fmt.Printf("| %-4d | %-43s | %-22s |\n",
			entry.LineNumber,
			truncate(entry.Statement, 43),
			truncate(varsStr, 22))
	}

	fmt.Println("+------+---------------------------------------------+------------------------+")
	fmt.Printf("\nTổng số câu lệnh đã thực thi: %d\n", len(t.Entries))
}

// truncate cắt chuỗi nếu quá dài
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// RunTracingDemo chạy demo tracing
func RunTracingDemo() {
	fmt.Println("\n🎯 Ví dụ: Tính giai thừa của một số")
	fmt.Println("\nChương trình mẫu:")
	fmt.Println("```")
	fmt.Println("func factorial(n int) int {")
	fmt.Println("    result := 1              // Line 1")
	fmt.Println("    for i := 1; i <= n; i++ {  // Line 2")
	fmt.Println("        result = result * i    // Line 3")
	fmt.Println("    }")
	fmt.Println("    return result            // Line 4")
	fmt.Println("}")
	fmt.Println("```")

	// Tạo tracer
	tracer := NewTracer()

	// Simulate factorial(5)
	n := 5
	fmt.Printf("\n▶ Chạy: factorial(%d)\n", n)

	result := 1
	tracer.Trace(1, "result := 1", map[string]interface{}{
		"n":      n,
		"result": result,
	})

	for i := 1; i <= n; i++ {
		tracer.Trace(2, fmt.Sprintf("for i := %d; i <= %d", i, n), map[string]interface{}{
			"i":      i,
			"n":      n,
			"result": result,
		})

		result = result * i
		tracer.Trace(3, fmt.Sprintf("result = %d * %d = %d", result/i, i, result), map[string]interface{}{
			"i":      i,
			"result": result,
		})
	}

	tracer.Trace(4, fmt.Sprintf("return %d", result), map[string]interface{}{
		"result": result,
	})

	// In trace
	tracer.PrintTrace()

	fmt.Println("\n📊 Phân tích:")
	fmt.Println("- Tracing cho phép theo dõi toàn bộ quá trình thực thi")
	fmt.Println("- Mỗi bước thực thi được ghi lại với giá trị biến tại thời điểm đó")
	fmt.Println("- Hữu ích cho việc debug và hiểu luồng thực thi")

	// Demo 2: Tìm số lớn nhất trong mảng
	fmt.Println("\n" + "==============================================")
	fmt.Println("\n🎯 Ví dụ 2: Tìm số lớn nhất trong mảng")
	fmt.Println("\nChương trình mẫu:")
	fmt.Println("```")
	fmt.Println("func findMax(arr []int) int {")
	fmt.Println("    max := arr[0]           // Line 1")
	fmt.Println("    for i := 1; i < len(arr); i++ {  // Line 2")
	fmt.Println("        if arr[i] > max {   // Line 3")
	fmt.Println("            max = arr[i]    // Line 4")
	fmt.Println("        }")
	fmt.Println("    }")
	fmt.Println("    return max              // Line 5")
	fmt.Println("}")
	fmt.Println("```")

	tracer2 := NewTracer()
	arr := []int{3, 7, 2, 9, 5, 1}
	fmt.Printf("\n▶ Chạy: findMax(%v)\n", arr)

	max := arr[0]
	tracer2.Trace(1, "max := arr[0]", map[string]interface{}{
		"arr": arr,
		"max": max,
	})

	for i := 1; i < len(arr); i++ {
		tracer2.Trace(2, fmt.Sprintf("for i := %d", i), map[string]interface{}{
			"i":   i,
			"max": max,
		})

		tracer2.Trace(3, fmt.Sprintf("if arr[%d]=%d > max=%d", i, arr[i], max), map[string]interface{}{
			"i":      i,
			"arr[i]": arr[i],
			"max":    max,
		})

		if arr[i] > max {
			max = arr[i]
			tracer2.Trace(4, fmt.Sprintf("max = arr[%d] = %d", i, max), map[string]interface{}{
				"i":   i,
				"max": max,
			})
		}
	}

	tracer2.Trace(5, fmt.Sprintf("return %d", max), map[string]interface{}{
		"max": max,
	})

	tracer2.PrintTrace()

	fmt.Println("\n📊 Lợi ích của Tracing:")
	fmt.Println("✓ Hiểu rõ luồng thực thi chương trình")
	fmt.Println("✓ Phát hiện lỗi logic")
	fmt.Println("✓ Kiểm tra giá trị biến qua từng bước")
	fmt.Println("✓ Phân tích hiệu năng")
}
