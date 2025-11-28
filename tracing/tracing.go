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
	fmt.Println("\nTRACE - THEO DAU THUC THI:")
	fmt.Println("+------+---------------------------------------------+------------------------+")
	fmt.Println("| Dong | Cau lenh                                    | Bien                   |")
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
	fmt.Printf("\nTong so cau lenh da thuc thi: %d\n", len(t.Entries))
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
	fmt.Println("\nVi du: Tinh giai thua cua mot so")
	fmt.Println("\nChuong trinh mau:")
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
	fmt.Printf("\nChay: factorial(%d)\n", n)

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

	fmt.Println("\nPhan tich:")
	fmt.Println("- Tracing cho phep theo doi toan bo qua trinh thuc thi")
	fmt.Println("- Moi buoc thuc thi duoc ghi lai voi gia tri bien tai thoi diem do")
	fmt.Println("- Huu ich cho viec debug va hieu luong thuc thi")

	fmt.Println("\nUng dung cua Tracing:")
	fmt.Println("+ Hieu ro luong thuc thi chuong trinh")
	fmt.Println("+ Phat hien loi logic")
	fmt.Println("+ Kiem tra gia tri bien qua tung buoc")
	fmt.Println("+ Phan tich hieu nang")
}
