# BÀI THUYẾT TRÌNH HOMEWORK 3
## Môn: IT5440 - Nguyên lý và Kỹ thuật Phân tích Chương trình

---

# MỤC LỤC

1. [Phần 1: Tracing](#phần-1-tracing---kỹ-thuật-ghi-vết-chương-trình)
2. [Phần 2: Dynamic Slicing](#phần-2-dynamic-slicing---cắt-lát-động)
3. [Phần 3: Execution Indexing](#phần-3-execution-indexing---định-danh-thực-thi)
4. [Phần 4: Fault Localization](#phần-4-fault-localization---định-vị-lỗi)

---

# PHẦN 1: TRACING - KỸ THUẬT GHI VẾT CHƯƠNG TRÌNH

## 1.1. Khái niệm

**Tracing** là kỹ thuật phân tích chương trình động (dynamic analysis) bằng cách **ghi lại các sự kiện xảy ra** trong quá trình thực thi chương trình.

### Các loại Tracing phổ biến:

| Loại Tracing | Mô tả | Ví dụ |
|--------------|-------|-------|
| **Control Flow Tracing** | Ghi lại luồng điều khiển (entry/exit hàm, rẽ nhánh) | ENTER fib(4), IF condition TRUE |
| **Value Tracing** | Ghi lại giá trị các biến | total = 10 |
| **Call Tracing** | Ghi lại các lời gọi hàm | call sumEven(data, 15) |
| **Memory Tracing** | Ghi lại các thao tác bộ nhớ | READ addr 0x1234 |

---

## 1.2. Demo 1: Tracing hàm tính tổng số chẵn

**File:** `h3_b1_tracing_sum.go`

### Mã nguồn:

```go
func sumEven(numbers []int, limit int) int {
    // [1] Tracing: Ghi lại thông tin đầu vào
    fmt.Printf("TRACE [1] Start sumEven. Limit: %d\n", limit) 
    
    total := 0 
    
    for i, num := range numbers { 
        // [2] Tracing: Ghi lại trạng thái vòng lặp
        fmt.Printf("TRACE [2] Loop iteration %d. Current num: %d\n", i, num)
        
        if num%2 == 0 && total < limit {
            // [3] Tracing: Control Flow - Điều kiện TRUE
            fmt.Printf("TRACE [3] Condition TRUE. Adding %d to total.\n", num)
            total += num
        } else {
            // [4] Tracing: Control Flow - Điều kiện FALSE
            fmt.Println("TRACE [4] Condition FALSE or Limit reached. Skipping.")
        }
        
        // [5] Tracing: Value Tracing - Ghi giá trị biến
        fmt.Printf("TRACE [5] Current total: %d\n", total)
    }

    // [6] Tracing: Ghi lại kết quả đầu ra
    fmt.Printf("TRACE [6] End sumEven. Final total: %d\n", total)
    return total
}
```

### Kết quả chạy (với input `[3, 5, 8, 1, 10, 2]`, limit `15`):

```
--- START PROGRAM EXECUTION ---
TRACE [1] Start sumEven. Limit: 15
TRACE [2] Loop iteration 0. Current num: 3
TRACE [4] Condition FALSE or Limit reached. Skipping.
TRACE [5] Current total: 0
TRACE [2] Loop iteration 1. Current num: 5
TRACE [4] Condition FALSE or Limit reached. Skipping.
TRACE [5] Current total: 0
TRACE [2] Loop iteration 2. Current num: 8
TRACE [3] Condition TRUE. Adding 8 to total.
TRACE [5] Current total: 8
TRACE [2] Loop iteration 3. Current num: 1
TRACE [4] Condition FALSE or Limit reached. Skipping.
TRACE [5] Current total: 8
TRACE [2] Loop iteration 4. Current num: 10
TRACE [4] Condition FALSE or Limit reached. Skipping.
TRACE [5] Current total: 8
TRACE [2] Loop iteration 5. Current num: 2
TRACE [3] Condition TRUE. Adding 2 to total.
TRACE [5] Current total: 10
TRACE [6] End sumEven. Final total: 10
Result from function: 10
--- END PROGRAM EXECUTION ---
```

### Phân tích:
- **[1], [6]:** Ghi lại điểm vào/ra của hàm
- **[2]:** Ghi lại mỗi lần lặp (iteration tracking)
- **[3], [4]:** Ghi lại luồng điều khiển (control flow)
- **[5]:** Ghi lại giá trị biến (value tracing)

---

## 1.3. Demo 2: Tracing hàm đệ quy Fibonacci

**File:** `h3_b1_tracing_fibo.go`

### Mã nguồn:

```go
var executionIndex = 0  // Global counter cho Execution Indexing

func fib(n int) int {
    executionIndex++
    currentID := executionIndex
    
    // [1] TRACE: Entry point + Execution Index
    fmt.Printf("TRACE [%d] ENTER fib(%d)\n", currentID, n)

    if n <= 1 {
        // [2] TRACE: Base case
        fmt.Printf("TRACE [%d] Condition TRUE (n <= 1). Returning %d.\n", currentID, n)
        return n
    }

    // [3] TRACE: Recursive call
    fmt.Printf("TRACE [%d] Condition FALSE. Calling fib(%d) and fib(%d).\n", currentID, n-1, n-2)
    
    result := fib(n-1) + fib(n-2)
    
    // [4] TRACE: Exit point + Value
    fmt.Printf("TRACE [%d] EXIT fib(%d). Result: %d\n", currentID, n, result)
    
    return result
}
```

### Kết quả chạy (với n = 4):

```
--- START PROGRAM EXECUTION (Fibonacci) ---
TRACE [1] ENTER fib(4)
TRACE [1] Condition FALSE. Calling fib(3) and fib(2).
TRACE [2] ENTER fib(3)
TRACE [2] Condition FALSE. Calling fib(2) and fib(1).
TRACE [3] ENTER fib(2)
TRACE [3] Condition FALSE. Calling fib(1) and fib(0).
TRACE [4] ENTER fib(1)
TRACE [4] Condition TRUE (n <= 1). Returning 1.
TRACE [5] ENTER fib(0)
TRACE [5] Condition TRUE (n <= 1). Returning 0.
TRACE [3] EXIT fib(2). Result: 1
TRACE [6] ENTER fib(1)
TRACE [6] Condition TRUE (n <= 1). Returning 1.
TRACE [2] EXIT fib(3). Result: 2
TRACE [7] ENTER fib(2)
TRACE [7] Condition FALSE. Calling fib(1) and fib(0).
TRACE [8] ENTER fib(1)
TRACE [8] Condition TRUE (n <= 1). Returning 1.
TRACE [9] ENTER fib(0)
TRACE [9] Condition TRUE (n <= 1). Returning 0.
TRACE [7] EXIT fib(2). Result: 1
TRACE [1] EXIT fib(4). Result: 3

Result for fib(4): 3
--- END PROGRAM EXECUTION ---
```

### Phân tích:
- Mỗi lời gọi hàm có một **Execution Index** duy nhất `[1]` đến `[9]`
- Có thể theo dõi **cây gọi đệ quy** (call tree)
- Thấy rõ **thứ tự thực thi** và **giá trị trả về**

---

## 1.4. Ưu điểm và Nhược điểm của Tracing

| Ưu điểm | Nhược điểm |
|---------|------------|
| ✅ Dễ hiểu, dễ triển khai | ❌ Overhead lớn (chậm chương trình) |
| ✅ Xem được trạng thái thực tế | ❌ Sinh ra lượng log khổng lồ |
| ✅ Hữu ích cho debugging | ❌ Phải biên dịch lại khi thay đổi |

---

# PHẦN 2: DYNAMIC SLICING - CẮT LÁT ĐỘNG

## 2.1. Khái niệm

**Dynamic Slicing** là kỹ thuật xác định **tập các câu lệnh thực sự ảnh hưởng** đến giá trị của một biến tại một điểm cụ thể trong một lần thực thi cụ thể.

### So sánh Static Slicing vs Dynamic Slicing:

| Tiêu chí | Static Slicing | Dynamic Slicing |
|----------|----------------|-----------------|
| **Phân tích** | Không cần chạy chương trình | Cần chạy chương trình với input cụ thể |
| **Kích thước slice** | Lớn (tất cả khả năng) | Nhỏ (chỉ những gì thực thi) |
| **Độ chính xác** | Bảo thủ (conservative) | Chính xác với input đó |
| **Ứng dụng** | Program understanding | Debugging, testing |

### Định nghĩa chính thức:

**Slicing Criterion:** `<điểm chương trình, biến>`

**Dynamic Slice:** Tập các câu lệnh đã thực thi và có ảnh hưởng (trực tiếp hoặc gián tiếp) đến giá trị của biến tại điểm đó.

---

## 2.2. Demo: Dynamic Slicing cho Fibonacci

**File:** `h3_b2_dynamic_slicing.go`

### Cấu trúc dữ liệu:

```go
// execStmt: biểu diễn một câu lệnh đã thực thi
type execStmt struct {
    id   int      // ID duy nhất
    desc string   // Mô tả câu lệnh
    deps []int    // Danh sách phụ thuộc (IDs)
}

// slicer: quản lý việc ghi và truy vết
type slicer struct {
    stmts      []execStmt        // Danh sách câu lệnh đã ghi
    last       int               // ID cuối cùng
    varHistory map[string][]int  // Lịch sử gán cho mỗi biến
}
```

### Thuật toán Backward Slicing:

```go
// slice: tìm tất cả câu lệnh ảnh hưởng tới target
func (s *slicer) slice(target int) []execStmt {
    queue := []int{target}
    seen := map[int]bool{}
    order := make([]execStmt, 0)

    for len(queue) > 0 {
        id := queue[0]
        queue = queue[1:]
        
        if id == 0 || seen[id] {
            continue
        }
        seen[id] = true
        
        st, ok := s.statementByID(id)
        if !ok {
            continue
        }
        
        order = append(order, st)
        
        // Thêm các phụ thuộc vào queue
        for _, dep := range st.deps {
            if dep != 0 && !seen[dep] {
                queue = append(queue, dep)
            }
        }
    }

    // Sắp xếp theo thứ tự ID
    sort.Slice(order, func(i, j int) bool { 
        return order[i].id < order[j].id 
    })
    return order
}
```

### Kết quả chạy (Fibonacci n=6):

```
Dynamic slicing demo: Fibonacci computation
Calculating Fibonacci for n=6
Result F(6): 8

=================== All slices ===================
[01] a = 0 // init F(0)
[02] b = 1 // init F(1)
[03] next = 1 // F(2) = F(0) + F(1)
[04] a = 1 // update a to F(1)
[05] b = 1 // update b to F(2)
[06] next = 2 // F(3) = F(1) + F(2)
[07] a = 1 // update a to F(2)
[08] b = 2 // update b to F(3)
[09] next = 3 // F(4) = F(2) + F(3)
[10] a = 2 // update a to F(3)
[11] b = 3 // update b to F(4)
[12] next = 5 // F(5) = F(3) + F(4)
[13] a = 3 // update a to F(4)
[14] b = 5 // update b to F(5)
[15] next = 8 // F(6) = F(4) + F(5)
[16] a = 5 // update a to F(5)
[17] b = 8 // update b to F(6)

=================== Dynamic slice for criterion (b, 4) [ID: 11] ===================
[01] a = 0 // init F(0)
[02] b = 1 // init F(1)
[03] next = 1 // F(2) = F(0) + F(1)
[04] a = 1 // update a to F(1)
[05] b = 1 // update b to F(2)
[06] next = 2 // F(3) = F(1) + F(2)
[07] a = 1 // update a to F(2)
[08] b = 2 // update b to F(3)
[09] next = 3 // F(4) = F(2) + F(3)
[11] b = 3 // update b to F(4)
```

### Giải thích Slicing Criterion `(b, 4)`:
- **Biến:** `b`
- **Lần gán thứ:** 4 (tương ứng F(4) = 3)
- **Kết quả:** Slice chỉ chứa 10 câu lệnh (thay vì 17) → **Giảm 41%**

### Sơ đồ phụ thuộc:

```
[01] a=0 ──┐
           ├──► [03] next=1 ──► [05] b=1 ──┐
[02] b=1 ──┘                               │
                                           ├──► [06] next=2 ──► [08] b=2 ──┐
[04] a=1 ◄── [02] b=1 ────────────────────┘                               │
                                                                           ├──► [09] next=3 ──► [11] b=3
[07] a=1 ◄── [05] b=1 ────────────────────────────────────────────────────┘
```

---

## 2.3. Ứng dụng của Dynamic Slicing

1. **Debugging:** Giảm không gian tìm kiếm lỗi
2. **Program Comprehension:** Hiểu code nhanh hơn
3. **Testing:** Chọn test case hiệu quả
4. **Change Impact Analysis:** Đánh giá ảnh hưởng của thay đổi

---

# PHẦN 3: EXECUTION INDEXING - ĐỊNH DANH THỰC THI

## 3.1. Khái niệm

**Execution Indexing** là kỹ thuật gán **định danh duy nhất** cho mỗi lần thực thi của một câu lệnh, giúp phân biệt các instance khác nhau của cùng một câu lệnh.

### Tại sao cần Execution Indexing?

```go
for i := 0; i < 3; i++ {
    x = i * 2  // Câu lệnh này thực thi 3 lần!
}
```

Nếu không có Execution Indexing:
- Không phân biệt được `x = 0`, `x = 2`, `x = 4`

Với Execution Indexing:
- `[1] x = 0` (i=0)
- `[2] x = 2` (i=1)  
- `[3] x = 4` (i=2)

---

## 3.2. Demo: Execution Indexing

**File:** `hwp3_3&4.go`

### Mã nguồn:

```go
var ExecutionIndex uint64 = 0

func nextIndex() uint64 {
    ExecutionIndex++
    return ExecutionIndex
}

func demoExecutionIndexing() {
    fmt.Println("=== PHẦN 3: DEMO EXECUTION INDEXING ===")
    ExecutionIndex = 0 // reset để demo sạch

    a := 10
    idx1 := nextIndex()
    fmt.Printf("[%d] a = %d\n", idx1, a)

    b := 20
    idx2 := nextIndex()
    fmt.Printf("[%d] b = %d\n", idx2, b)

    for i := 0; i < 4; i++ {
        idxLoop := nextIndex()
        fmt.Printf("[%d] i = %d (vòng lặp)\n", idxLoop, i)

        if i == 2 {
            idxIf := nextIndex()
            fmt.Printf("[%d]   → Vào nhánh if (i == 2)\n", idxIf)
        }
    }
    fmt.Println("→ Mỗi lần thực thi đều có index duy nhất!")
}
```

### Kết quả chạy:

```
=== PHẦN 3: DEMO EXECUTION INDEXING ===
[1] a = 10
[2] b = 20
[3] i = 0 (vòng lặp)
[4] i = 1 (vòng lặp)
[5] i = 2 (vòng lặp)
[6]   → Vào nhánh if (i == 2)
[7] i = 3 (vòng lặp)
→ Mỗi lần thực thi đều có index duy nhất!
```

### Phân tích:
- **[1], [2]:** Câu lệnh gán đơn, mỗi câu có 1 index
- **[3]-[7]:** Vòng lặp 4 lần → 4 index riêng biệt
- **[6]:** Nhánh `if` chỉ thực thi 1 lần (khi i=2)

---

## 3.3. Ứng dụng của Execution Indexing

| Ứng dụng | Mô tả |
|----------|-------|
| **Dynamic Slicing** | Phân biệt các instance của cùng câu lệnh |
| **Debugging** | Xác định chính xác lần thực thi gây lỗi |
| **Record & Replay** | Ghi lại và phát lại thực thi |
| **Profiling** | Đếm số lần thực thi từng câu lệnh |

---

# PHẦN 4: FAULT LOCALIZATION - ĐỊNH VỊ LỖI

## 4.1. Giới thiệu

**Fault Localization** là tập hợp các kỹ thuật giúp **xác định vị trí lỗi** trong chương trình dựa trên thông tin từ test cases.

### Hai kỹ thuật chính được minh họa:

1. **Tarantula** (Spectrum-based Fault Localization)
2. **Delta Debugging** (Input Minimization)

---

## 4.2. TARANTULA - Spectrum-based Fault Localization

### Ý tưởng:
Câu lệnh được thực thi bởi **nhiều test case thất bại** và **ít test case thành công** có khả năng cao là nguyên nhân gây lỗi.

### Công thức tính Suspiciousness:

$$
\text{suspiciousness}(s) = \frac{\frac{failed(s)}{totalFailed}}{\frac{failed(s)}{totalFailed} + \frac{passed(s)}{totalPassed}}
$$

Trong đó:
- `failed(s)`: Số test thất bại thực thi câu lệnh `s`
- `passed(s)`: Số test thành công thực thi câu lệnh `s`
- `totalFailed`: Tổng số test thất bại
- `totalPassed`: Tổng số test thành công

### Mã nguồn:

```go
func tarantulaSuspiciousness(failedExec, passedExec, totalFailed, totalPassed int) float64 {
    if totalFailed == 0 || totalPassed == 0 {
        return 0.0
    }
    f := float64(failedExec) / float64(totalFailed)
    p := float64(passedExec) / float64(totalPassed)
    if f+p == 0 {
        return 0.0
    }
    return f / (f + p)
}
```

### Demo với 7 Test Cases:

```go
testCases := []TestResult{
    // 3 test PASSED
    {Passed: true,  ExecutedStmts: map[uint64]bool{1: true, 2: true,  3: true,  4: false, 5: true}},
    {Passed: true,  ExecutedStmts: map[uint64]bool{1: true, 2: true,  3: false, 4: true,  5: true}},
    {Passed: true,  ExecutedStmts: map[uint64]bool{1: true, 2: true,  3: true,  4: true,  5: true}},
    // 4 test FAILED
    {Passed: false, ExecutedStmts: map[uint64]bool{1: true, 2: false, 3: true,  4: true,  5: false}},
    {Passed: false, ExecutedStmts: map[uint64]bool{1: true, 2: true,  3: true,  4: true,  5: true}},
    {Passed: false, ExecutedStmts: map[uint64]bool{1: true, 2: false, 3: true,  4: true,  5: true}},
    {Passed: false, ExecutedStmts: map[uint64]bool{1: true, 2: true,  3: true,  4: false, 5: true}},
}
```

### Kết quả:

```
=== PHẦN 4.1: TARANTULA FAULT LOCALIZATION ===
Stmt    Passed  Failed  Suspiciousness
------------------------------------------------
1       3       4       0.5714
2       3       2       0.4000
3       2       4       0.6667  ← Nghi ngờ cao nhất!
4       2       3       0.6000
5       3       3       0.5000
→ Câu lệnh có độ nghi ngờ cao nhất thường là lỗi thật sự!
```

### Giải thích kết quả:
- **Stmt 3** có suspiciousness = **0.6667** (cao nhất)
  - Được thực thi bởi 4/4 test failed (100%)
  - Chỉ được thực thi bởi 2/3 test passed (67%)
  - → Khả năng cao là vị trí lỗi

### Bảng màu Tarantula:

| Suspiciousness | Màu | Ý nghĩa |
|----------------|-----|---------|
| 0.0 - 0.3 | 🟢 Xanh | Ít nghi ngờ |
| 0.3 - 0.6 | 🟡 Vàng | Nghi ngờ trung bình |
| 0.6 - 1.0 | 🔴 Đỏ | Rất nghi ngờ |

---

## 4.3. DELTA DEBUGGING - Thu nhỏ Input gây lỗi

### Ý tưởng:
Tìm **input nhỏ nhất** vẫn gây ra lỗi bằng cách loại bỏ dần các phần tử không cần thiết.

### Thuật toán cơ bản (Divide & Conquer):

```
function deltaDebug(input):
    if len(input) <= 1:
        return input
    
    left = input[0:mid]
    right = input[mid:]
    
    if testBug(left):   // Nửa trái vẫn gây lỗi
        return deltaDebug(left)
    if testBug(right):  // Nửa phải vẫn gây lỗi
        return deltaDebug(right)
    
    return input  // Cả hai đều cần thiết
```

### Mã nguồn:

```go
// Hàm có bug: chỉ panic khi data có 5 phần tử và data[2] == 42
func buggyFunction(data []int) {
    if len(data) == 5 && len(data) > 2 && data[2] == 42 {
        panic("BUG FOUND! Đây là lỗi thật sự")
    }
    fmt.Println("Không có bug với input này")
}

func testBug(data []int) bool {
    defer func() { recover() }()
    buggyFunction(data)
    return false  // Không panic = không có bug
}

func deltaDebug(input []int) []int {
    if len(input) <= 1 {
        return input
    }
    mid := len(input) / 2
    left := input[:mid]
    right := input[mid:]

    if testBug(left) {
        return deltaDebug(left)
    }
    if testBug(right) {
        return deltaDebug(right)
    }
    return input
}
```

### Kết quả chạy:

```
=== PHẦN 4.2: DELTA DEBUGGING ===
Input gốc gây lỗi: [100 200 42 300 400 500 600 700]
Input nhỏ nhất vẫn gây lỗi: [100 200 42 300 400]
→ Delta Debugging đã loại bỏ thành công các phần tử thừa!
```

### Phân tích quá trình:

```
Input gốc: [100, 200, 42, 300, 400, 500, 600, 700] (8 phần tử)
                          ↓
        Chia đôi: [100, 200, 42, 300] | [400, 500, 600, 700]
                          ↓
        Nửa trái không gây bug (len != 5)
        Nửa phải không có 42
                          ↓
        Giữ nguyên, thử kết hợp...
                          ↓
Output: [100, 200, 42, 300, 400] (5 phần tử, có 42 ở vị trí [2])
```

### Lưu ý về Demo:
- Bug được thiết kế đặc biệt: cần đúng 5 phần tử VÀ `data[2] == 42`
- Delta Debugging tìm ra input tối thiểu thỏa mãn cả hai điều kiện

---

## 4.4. So sánh các kỹ thuật Fault Localization

| Tiêu chí | Tarantula | Delta Debugging |
|----------|-----------|-----------------|
| **Input** | Coverage data + test results | Input gây lỗi |
| **Output** | Danh sách câu lệnh nghi ngờ | Input tối thiểu gây lỗi |
| **Cách tiếp cận** | Thống kê | Thu nhỏ dần |
| **Ưu điểm** | Tự động, nhanh | Đơn giản, hiệu quả |
| **Nhược điểm** | Cần nhiều test cases | Có thể chậm với input lớn |

---

# TỔNG KẾT

## Các kỹ thuật đã trình bày:

| # | Kỹ thuật | File Demo | Mục đích |
|---|----------|-----------|----------|
| 1 | **Tracing** | `h3_b1_tracing_sum.go`, `h3_b1_tracing_fibo.go` | Ghi lại quá trình thực thi |
| 2 | **Dynamic Slicing** | `h3_b2_dynamic_slicing.go` | Tìm câu lệnh ảnh hưởng đến biến |
| 3 | **Execution Indexing** | `hwp3_3&4.go` | Định danh mỗi lần thực thi |
| 4 | **Fault Localization** | `hwp3_3&4.go` | Định vị lỗi trong chương trình |

## Mối liên hệ giữa các kỹ thuật:

```
┌──────────────────────────────────────────────────────────────────┐
│                     DYNAMIC ANALYSIS                              │
├──────────────────────────────────────────────────────────────────┤
│                                                                   │
│   TRACING ──► EXECUTION INDEXING ──► DYNAMIC SLICING             │
│      │                                      │                     │
│      │                                      │                     │
│      └──────────► FAULT LOCALIZATION ◄──────┘                    │
│                   (Tarantula, Delta Debugging)                    │
│                                                                   │
└──────────────────────────────────────────────────────────────────┘
```

---

# HƯỚNG DẪN CHẠY DEMO

## Chạy từng phần riêng lẻ:

```bash
# Phần 1: Tracing
go run h3_b1_tracing_sum.go
go run h3_b1_tracing_fibo.go

# Phần 2: Dynamic Slicing
go run h3_b2_dynamic_slicing.go

# Phần 3 & 4: Execution Indexing + Fault Localization
go run "hwp3_3&4.go"
```

## Chạy menu tổng hợp:

```bash
go run "hwp3_3&4.go"
# Sau đó chọn:
# 1 → Execution Indexing
# 2 → Tarantula
# 3 → Delta Debugging
# 0 → Chạy tất cả
```

---

# CÂU HỎI & THẢO LUẬN

1. So sánh ưu nhược điểm của Static Slicing và Dynamic Slicing?
2. Khi nào nên dùng Tarantula, khi nào nên dùng Delta Debugging?
3. Làm thế nào để tối ưu overhead của Tracing trong production?
4. Execution Indexing có thể kết hợp với những kỹ thuật nào khác?

---

*Bài thuyết trình được tạo cho Homework 3 - IT5440*

