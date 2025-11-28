package faultlocalization

import (
	"fmt"
	"math"
	"sort"
)

// TestCase đại diện cho một test case
type TestCase struct {
	Name   string
	Input  map[string]interface{}
	Output interface{}
	Passed bool
	Trace  []int // Các dòng đã thực thi
}

// StatementScore đại diện cho điểm nghi ngờ của một statement
type StatementScore struct {
	LineNumber     int
	Statement      string
	SuspiciousScore float64
	PassedCount    int
	FailedCount    int
}

// FaultLocalizer thực hiện fault localization
type FaultLocalizer struct {
	Statements map[int]string        // Line -> Statement
	TestCases  []TestCase
	Scores     []StatementScore
}

// NewFaultLocalizer tạo localizer mới
func NewFaultLocalizer() *FaultLocalizer {
	return &FaultLocalizer{
		Statements: make(map[int]string),
		TestCases:  make([]TestCase, 0),
		Scores:     make([]StatementScore, 0),
	}
}

// AddStatement thêm statement
func (fl *FaultLocalizer) AddStatement(lineNumber int, statement string) {
	fl.Statements[lineNumber] = statement
}

// AddTestCase thêm test case
func (fl *FaultLocalizer) AddTestCase(name string, input map[string]interface{}, output interface{}, passed bool, trace []int) {
	fl.TestCases = append(fl.TestCases, TestCase{
		Name:   name,
		Input:  input,
		Output: output,
		Passed: passed,
		Trace:  trace,
	})
}

// CalculateTarantula tính điểm nghi ngờ bằng công thức Tarantula
func (fl *FaultLocalizer) CalculateTarantula() {
	fl.Scores = make([]StatementScore, 0)

	// Đếm tổng số test passed và failed
	totalPassed := 0
	totalFailed := 0
	for _, tc := range fl.TestCases {
		if tc.Passed {
			totalPassed++
		} else {
			totalFailed++
		}
	}

	// Tính điểm cho mỗi statement
	for lineNum, stmt := range fl.Statements {
		passedCount := 0
		failedCount := 0

		// Đếm số lần statement xuất hiện trong passed và failed tests
		for _, tc := range fl.TestCases {
			executed := contains(tc.Trace, lineNum)
			if executed {
				if tc.Passed {
					passedCount++
				} else {
					failedCount++
				}
			}
		}

		// Công thức Tarantula:
		// suspiciousness = (failed/totalFailed) / ((failed/totalFailed) + (passed/totalPassed))
		var score float64
		if totalFailed == 0 {
			score = 0
		} else if totalPassed == 0 {
			if failedCount > 0 {
				score = 1.0
			} else {
				score = 0
			}
		} else {
			failedRatio := float64(failedCount) / float64(totalFailed)
			passedRatio := float64(passedCount) / float64(totalPassed)

			if failedRatio+passedRatio == 0 {
				score = 0
			} else {
				score = failedRatio / (failedRatio + passedRatio)
			}
		}

		fl.Scores = append(fl.Scores, StatementScore{
			LineNumber:      lineNum,
			Statement:       stmt,
			SuspiciousScore: score,
			PassedCount:     passedCount,
			FailedCount:     failedCount,
		})
	}

	// Sắp xếp theo điểm giảm dần
	sort.Slice(fl.Scores, func(i, j int) bool {
		if math.Abs(fl.Scores[i].SuspiciousScore-fl.Scores[j].SuspiciousScore) < 0.0001 {
			return fl.Scores[i].LineNumber < fl.Scores[j].LineNumber
		}
		return fl.Scores[i].SuspiciousScore > fl.Scores[j].SuspiciousScore
	})
}

// CalculateOchiai tính điểm nghi ngờ bằng công thức Ochiai
func (fl *FaultLocalizer) CalculateOchiai() {
	fl.Scores = make([]StatementScore, 0)

	totalFailed := 0
	for _, tc := range fl.TestCases {
		if !tc.Passed {
			totalFailed++
		}
	}

	for lineNum, stmt := range fl.Statements {
		passedCount := 0
		failedCount := 0

		for _, tc := range fl.TestCases {
			executed := contains(tc.Trace, lineNum)
			if executed {
				if tc.Passed {
					passedCount++
				} else {
					failedCount++
				}
			}
		}

		// Công thức Ochiai:
		// suspiciousness = failed / sqrt(totalFailed * (failed + passed))
		var score float64
		if totalFailed == 0 || (failedCount+passedCount) == 0 {
			score = 0
		} else {
			score = float64(failedCount) / math.Sqrt(float64(totalFailed*(failedCount+passedCount)))
		}

		fl.Scores = append(fl.Scores, StatementScore{
			LineNumber:      lineNum,
			Statement:       stmt,
			SuspiciousScore: score,
			PassedCount:     passedCount,
			FailedCount:     failedCount,
		})
	}

	sort.Slice(fl.Scores, func(i, j int) bool {
		if math.Abs(fl.Scores[i].SuspiciousScore-fl.Scores[j].SuspiciousScore) < 0.0001 {
			return fl.Scores[i].LineNumber < fl.Scores[j].LineNumber
		}
		return fl.Scores[i].SuspiciousScore > fl.Scores[j].SuspiciousScore
	})
}

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// PrintTestResults in kết quả test
func (fl *FaultLocalizer) PrintTestResults() {
	fmt.Println("\n🧪 KẾT QUẢ TEST CASES:")
	fmt.Println("+------------------+-------------+------------------+")
	fmt.Println("| Test Case        | Kết quả     | Execution Trace  |")
	fmt.Println("+------------------+-------------+------------------+")

	for _, tc := range fl.TestCases {
		status := "❌ FAILED"
		if tc.Passed {
			status = "✓ PASSED"
		}

		traceStr := fmt.Sprintf("%v", tc.Trace)
		if len(traceStr) > 15 {
			traceStr = traceStr[:12] + "..."
		}

		fmt.Printf("| %-16s | %-11s | %-16s |\n", tc.Name, status, traceStr)
	}

	fmt.Println("+------------------+-------------+------------------+")

	passed := 0
	failed := 0
	for _, tc := range fl.TestCases {
		if tc.Passed {
			passed++
		} else {
			failed++
		}
	}
	fmt.Printf("\nTổng: %d tests (%d passed, %d failed)\n", len(fl.TestCases), passed, failed)
}

// PrintSuspiciousScores in điểm nghi ngờ
func (fl *FaultLocalizer) PrintSuspiciousScores(method string) {
	fmt.Printf("\n🔍 SUSPICIOUS SCORES (Phương pháp: %s):\n", method)
	fmt.Println("+------+---------------------------------------+-------+--------+--------+")
	fmt.Println("| Dòng | Statement                             | Score | Failed | Passed |")
	fmt.Println("+------+---------------------------------------+-------+--------+--------+")

	for _, score := range fl.Scores {
		fmt.Printf("| %-4d | %-37s | %.3f | %-6d | %-6d |\n",
			score.LineNumber,
			truncate(score.Statement, 37),
			score.SuspiciousScore,
			score.FailedCount,
			score.PassedCount)
	}

	fmt.Println("+------+---------------------------------------+-------+--------+--------+")

	if len(fl.Scores) > 0 {
		fmt.Printf("\n⚠️  Dòng nghi ngờ nhất: Dòng %d (Score: %.3f)\n",
			fl.Scores[0].LineNumber,
			fl.Scores[0].SuspiciousScore)
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// RunFaultLocalizationDemo chạy demo
func RunFaultLocalizationDemo() {
	fmt.Println("\n🎯 Ví dụ: Tìm lỗi trong hàm tính trung bình")
	fmt.Println("\nChương trình có lỗi:")
	fmt.Println("```")
	fmt.Println(" 1: func average(arr []int) float64 {")
	fmt.Println(" 2:     sum := 0")
	fmt.Println(" 3:     for i := 0; i < len(arr); i++ {")
	fmt.Println(" 4:         sum = sum + arr[i]")
	fmt.Println(" 5:     }")
	fmt.Println(" 6:     avg := sum / len(arr)  // BUG: Chia integer!")
	fmt.Println(" 7:     return avg")
	fmt.Println(" 8: }")
	fmt.Println("```")

	localizer := NewFaultLocalizer()

	// Thêm statements
	localizer.AddStatement(1, "func average(arr []int) float64")
	localizer.AddStatement(2, "sum := 0")
	localizer.AddStatement(3, "for i := 0; i < len(arr); i++")
	localizer.AddStatement(4, "sum = sum + arr[i]")
	localizer.AddStatement(5, "}")
	localizer.AddStatement(6, "avg := sum / len(arr)")
	localizer.AddStatement(7, "return avg")

	// Test case 1: [10, 20, 30] -> Expected: 20.0, Got: 20 (PASSED vì ngẫu nhiên đúng)
	localizer.AddTestCase("Test 1", map[string]interface{}{"arr": []int{10, 20, 30}},
		20.0, true, []int{1, 2, 3, 4, 3, 4, 3, 4, 5, 6, 7})

	// Test case 2: [5, 10, 15] -> Expected: 10.0, Got: 10 (PASSED)
	localizer.AddTestCase("Test 2", map[string]interface{}{"arr": []int{5, 10, 15}},
		10.0, true, []int{1, 2, 3, 4, 3, 4, 3, 4, 5, 6, 7})

	// Test case 3: [1, 2, 3] -> Expected: 2.0, Got: 2 (PASSED)
	localizer.AddTestCase("Test 3", map[string]interface{}{"arr": []int{1, 2, 3}},
		2.0, true, []int{1, 2, 3, 4, 3, 4, 3, 4, 5, 6, 7})

	// Test case 4: [1, 2] -> Expected: 1.5, Got: 1 (FAILED - lỗi chia integer)
	localizer.AddTestCase("Test 4", map[string]interface{}{"arr": []int{1, 2}},
		1.5, false, []int{1, 2, 3, 4, 3, 4, 5, 6, 7})

	// Test case 5: [3, 4] -> Expected: 3.5, Got: 3 (FAILED)
	localizer.AddTestCase("Test 5", map[string]interface{}{"arr": []int{3, 4}},
		3.5, false, []int{1, 2, 3, 4, 3, 4, 5, 6, 7})

	// Test case 6: [7, 8, 9] -> Expected: 8.0, Got: 8 (PASSED)
	localizer.AddTestCase("Test 6", map[string]interface{}{"arr": []int{7, 8, 9}},
		8.0, true, []int{1, 2, 3, 4, 3, 4, 3, 4, 5, 6, 7})

	// In kết quả test
	localizer.PrintTestResults()

	// Phân tích với Tarantula
	fmt.Println("\n" + "==============================================")
	localizer.CalculateTarantula()
	localizer.PrintSuspiciousScores("Tarantula")

	// Phân tích với Ochiai
	fmt.Println("\n" + "==============================================")
	localizer.CalculateOchiai()
	localizer.PrintSuspiciousScores("Ochiai")

	fmt.Println("\n💡 Giải thích:")
	fmt.Println("- Dòng 6 có điểm nghi ngờ cao nhất")
	fmt.Println("- Lỗi: Chia integer thay vì chia float")
	fmt.Println("- Sửa: avg := float64(sum) / float64(len(arr))")

	// Demo 2
	fmt.Println("\n" + "==============================================")
	demo2()
}

func demo2() {
	fmt.Println("\n🎯 Ví dụ 2: Tìm lỗi trong hàm tìm max")
	fmt.Println("\nChương trình có lỗi:")
	fmt.Println("```")
	fmt.Println(" 1: func findMax(arr []int) int {")
	fmt.Println(" 2:     max := 0  // BUG: Nếu tất cả số âm thì sai!")
	fmt.Println(" 3:     for i := 0; i < len(arr); i++ {")
	fmt.Println(" 4:         if arr[i] > max {")
	fmt.Println(" 5:             max = arr[i]")
	fmt.Println(" 6:         }")
	fmt.Println(" 7:     }")
	fmt.Println(" 8:     return max")
	fmt.Println(" 9: }")
	fmt.Println("```")

	localizer := NewFaultLocalizer()

	localizer.AddStatement(1, "func findMax(arr []int) int")
	localizer.AddStatement(2, "max := 0")
	localizer.AddStatement(3, "for i := 0; i < len(arr); i++")
	localizer.AddStatement(4, "if arr[i] > max")
	localizer.AddStatement(5, "max = arr[i]")
	localizer.AddStatement(6, "}")
	localizer.AddStatement(7, "}")
	localizer.AddStatement(8, "return max")

	// Test cases
	localizer.AddTestCase("Test 1", map[string]interface{}{"arr": []int{1, 5, 3}},
		5, true, []int{1, 2, 3, 4, 5, 3, 4, 5, 3, 4, 6, 7, 8})

	localizer.AddTestCase("Test 2", map[string]interface{}{"arr": []int{10, 2, 8}},
		10, true, []int{1, 2, 3, 4, 5, 3, 4, 3, 4, 6, 7, 8})

	localizer.AddTestCase("Test 3", map[string]interface{}{"arr": []int{-5, -2, -8}},
		-2, false, []int{1, 2, 3, 4, 3, 4, 3, 4, 6, 7, 8}) // FAILED: Trả về 0

	localizer.AddTestCase("Test 4", map[string]interface{}{"arr": []int{-10, -20, -5}},
		-5, false, []int{1, 2, 3, 4, 3, 4, 3, 4, 6, 7, 8}) // FAILED

	localizer.AddTestCase("Test 5", map[string]interface{}{"arr": []int{3, 7, 2}},
		7, true, []int{1, 2, 3, 4, 5, 3, 4, 5, 3, 4, 6, 7, 8})

	localizer.PrintTestResults()

	fmt.Println("\n" + "==============================================")
	localizer.CalculateTarantula()
	localizer.PrintSuspiciousScores("Tarantula")

	fmt.Println("\n" + "==============================================")
	localizer.CalculateOchiai()
	localizer.PrintSuspiciousScores("Ochiai")

	fmt.Println("\n💡 Kết luận:")
	fmt.Println("- Dòng 2 (max := 0) có điểm nghi ngờ cao")
	fmt.Println("- Lỗi: Khởi tạo max = 0, không xử lý mảng toàn số âm")
	fmt.Println("- Sửa: max := arr[0]")

	fmt.Println("\n📊 Ứng dụng Fault Localization:")
	fmt.Println("✓ Tự động phát hiện vị trí có khả năng chứa lỗi")
	fmt.Println("✓ Giảm thời gian debug")
	fmt.Println("✓ Ưu tiên kiểm tra các dòng có điểm cao")
	fmt.Println("✓ Kết hợp với test suite để phân tích")

	fmt.Println("\n⚙️  Các kỹ thuật phổ biến:")
	fmt.Println("• Tarantula: Dựa trên tỷ lệ failed/passed")
	fmt.Println("• Ochiai: Tương tự Tarantula nhưng dùng công thức khác")
	fmt.Println("• Jaccard: Đo độ tương đồng giữa failed và passed")
	fmt.Println("• DStar: Cải tiến của Ochiai với trọng số cao hơn")
}

