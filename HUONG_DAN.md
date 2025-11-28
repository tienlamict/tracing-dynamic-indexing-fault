# HƯỚNG DẪN SỬ DỤNG CHI TIẾT

## Giới thiệu

Project này minh họa 4 kỹ thuật quan trọng trong phân tích chương trình:

1. **Tracing**: Theo dấu mọi bước thực thi
2. **Dynamic Slicing**: Cắt chương trình theo biến quan tâm
3. **Execution Indexing**: Đánh số thứ tự thực thi
4. **Fault Localization**: Tự động tìm vị trí lỗi

## Cách chạy chương trình

### Bước 1: Mở terminal/command prompt

```bash
cd D:\Project\Golang\principle_and_technique_of_program_analysis\tracing-dynamic-indexing-fault
```

### Bước 2: Chạy chương trình

```bash
go run main.go
```

Hoặc nếu đã build:

```bash
.\program-analysis-demo.exe
```

### Bước 3: Chọn menu

Nhập số từ 1-5 để chọn demo tương ứng.

## Chi tiết từng kỹ thuật

### 1. TRACING - THEO DẤU THỰC THI

#### Mục đích
Ghi lại toàn bộ quá trình thực thi chương trình, bao gồm:
- Dòng code nào được thực thi
- Thứ tự thực thi
- Giá trị của các biến tại mỗi bước

#### Ví dụ minh họa
Program sẽ chạy 2 ví dụ:
1. **Tính giai thừa**: factorial(5) = 120
2. **Tìm số lớn nhất**: findMax([3, 7, 2, 9, 5, 1]) = 9

#### Output mẫu
```
+------+---------------------------------------------+------------------------+
| Dòng | Câu lệnh                                    | Biến                   |
+------+---------------------------------------------+------------------------+
|    1 | result := 1                                 | n=5, result=1          |
|    2 | for i := 1; i <= 5                         | i=1, n=5, result=1     |
|    3 | result = 1 * 1 = 1                         | i=1, result=1          |
|    2 | for i := 2; i <= 5                         | i=2, n=5, result=1     |
|    3 | result = 1 * 2 = 2                         | i=2, result=2          |
...
```

#### Ứng dụng thực tế
- Debug: Hiểu chương trình chạy như thế nào
- Performance: Phát hiện bottleneck
- Teaching: Giảng dạy thuật toán

---

### 2. DYNAMIC SLICING - CẮT ĐỘNG

#### Mục đích
Tìm tất cả các dòng code ảnh hưởng đến một biến cụ thể tại một điểm cụ thể.

#### Slicing Criterion
Format: `<line_number, variable_name>`

Ví dụ: `<7, sum>` = "Tìm tất cả dòng ảnh hưởng đến biến `sum` tại dòng 7"

#### Ví dụ minh họa
Program có 2 biến: `sum` và `product`

```go
1: sum := 0
2: product := 1
3: for i := 1; i <= 5; i++ {
4:     sum = sum + i
5:     product = product * i
6: }
7: result := sum + product
```

**Case 1**: Slice cho `<7, sum>`
- Kết quả: Dòng 1, 3, 4
- Giải thích: Chỉ các dòng này ảnh hưởng đến `sum`

**Case 2**: Slice cho `<7, product>`
- Kết quả: Dòng 2, 3, 5
- Giải thích: Dòng 4 không ảnh hưởng đến `product`

#### Output mẫu
```
🔪 DYNAMIC SLICE:
Slicing Criterion: <7, sum>

+------+--------------------------------------------------+
| Dòng | Câu lệnh                                         |
+------+--------------------------------------------------+
|    1 | sum := 0                                         |
|    3 | for i := 1; i <= 5; i++                         |
|    4 | sum = sum + i                                    |
+------+--------------------------------------------------+

Số dòng trong slice: 3
Tổng số dòng thực thi: 12
Tỷ lệ rút gọn: 25.0%
```

#### Ứng dụng
- Debugging: Tập trung vào code liên quan
- Code understanding: Hiểu dependencies
- Testing: Tạo test cases hiệu quả

---

### 3. EXECUTION INDEXING - ĐÁNH CHỈ MỤC

#### Mục đích
Gán một số thứ tự duy nhất (index) cho mỗi lần thực thi câu lệnh.

#### Tại sao cần?
- Cùng một dòng code có thể chạy nhiều lần (loop)
- Cần phân biệt lần 1, lần 2, lần 3...
- Hỗ trợ time-travel debugging

#### Ví dụ minh họa

**Demo 1**: Vòng lặp đơn giản
```go
1: n := 5
2: sum := 0
3: for i := 1; i <= n; i++ {
4:     sum = sum + i
5: }
```

**Demo 2**: Vòng lặp lồng nhau
```go
1: for i := 1; i <= 3; i++ {
2:     for j := 1; j <= 2; j++ {
3:         print(i, j)
4:     }
5: }
```

**Demo 3**: Đa luồng (multi-threading)
```
Thread 1: x = 0 → x = x + 1 → x = x + 2
Thread 2: y = 0 → y = y + 10 → y = y + 20
```

#### Output mẫu
```
+-------+------+------+----------------------------------------+-----------------+
| Index | Dòng | TID  | Câu lệnh                               | Biến            |
+-------+------+------+----------------------------------------+-----------------+
|     1 |    1 | T1   | n := 5                                 | n=5             |
|     2 |    2 | T1   | sum := 0                               | sum=0           |
|     3 |    3 | T1   | for i := 1                            | i=1, n=5, sum=0 |
|     4 |    4 | T1   | sum = 0 + 1                           | i=1, sum=1      |
|     5 |    3 | T1   | for i := 2                            | i=2, n=5, sum=1 |
...
```

#### Tính năng query
- `GetExecutionAtIndex(7)`: Lấy thực thi tại index 7
- `GetExecutionsBetween(5, 10)`: Lấy từ index 5 đến 10
- `FindExecutionsByLine(3)`: Tìm tất cả lần chạy dòng 3

#### Ứng dụng
- Time-travel debugging: Quay lại thời điểm trước
- Record & Replay: Tái hiện execution
- Concurrency analysis: Phân tích đa luồng

---

### 4. FAULT LOCALIZATION - ĐỊNH VỊ LỖI

#### Mục đích
Tự động tìm vị trí có khả năng chứa lỗi dựa trên kết quả test cases.

#### Nguyên lý
1. Chạy nhiều test cases (passed và failed)
2. Ghi lại execution trace của mỗi test
3. Tính "suspicious score" cho mỗi dòng code
4. Dòng nào có điểm cao nhất = nghi ngờ nhất

#### Công thức

**Tarantula:**
```
suspiciousness = (failed/totalFailed) / 
                 ((failed/totalFailed) + (passed/totalPassed))
```

**Ochiai:**
```
suspiciousness = failed / sqrt(totalFailed * (failed + passed))
```

#### Ví dụ minh họa

**Demo 1**: Lỗi chia integer
```go
func average(arr []int) float64 {
    sum := 0
    for i := 0; i < len(arr); i++ {
        sum = sum + arr[i]
    }
    avg := sum / len(arr)  // BUG: Chia integer!
    return avg
}
```

Test cases:
- `[10, 20, 30]` → Expected 20.0, Got 20 ✓ (ngẫu nhiên đúng)
- `[1, 2]` → Expected 1.5, Got 1 ❌ (lỗi!)
- `[3, 4]` → Expected 3.5, Got 3 ❌ (lỗi!)

**Demo 2**: Lỗi khởi tạo
```go
func findMax(arr []int) int {
    max := 0  // BUG: Nếu tất cả số âm?
    for i := 0; i < len(arr); i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}
```

Test cases:
- `[1, 5, 3]` → Expected 5, Got 5 ✓
- `[-5, -2, -8]` → Expected -2, Got 0 ❌
- `[-10, -20, -5]` → Expected -5, Got 0 ❌

#### Output mẫu
```
🔍 SUSPICIOUS SCORES (Phương pháp: Tarantula):
+------+---------------------------------------+-------+--------+--------+
| Dòng | Statement                             | Score | Failed | Passed |
+------+---------------------------------------+-------+--------+--------+
|    6 | avg := sum / len(arr)                 | 0.800 |      2 |      4 |
|    7 | return avg                            | 0.800 |      2 |      4 |
|    4 | sum = sum + arr[i]                    | 0.444 |      2 |      4 |
|    3 | for i := 0; i < len(arr); i++        | 0.444 |      2 |      4 |
...
```

⚠️ Dòng nghi ngờ nhất: Dòng 6 (Score: 0.800)

#### Cách sử dụng kết quả
1. Sắp xếp dòng theo điểm giảm dần
2. Kiểm tra từ dòng có điểm cao nhất
3. Tiết kiệm thời gian debug đáng kể

#### Xử lý khi nhiều dòng có cùng điểm (Tie-Breaking)

**Vấn đề**: Nhiều dòng code có thể có cùng suspicious score!

**Ví dụ**: Tất cả các dòng trong một function đều được thực thi trong mọi test case
→ Chúng sẽ có cùng tỷ lệ failed/passed → Cùng điểm!

**Giải pháp - Tie-Breaking Strategy**:
1. **Ưu tiên 1**: Dòng có điểm suspicious cao nhất
2. **Ưu tiên 2**: Nếu điểm bằng nhau → Dòng xuất hiện nhiều hơn trong failed tests
3. **Ưu tiên 3**: Nếu vẫn bằng nhau → Dòng có line number nhỏ hơn

**Ví dụ cụ thể**:
```
Dòng 6: Score=0.800, Failed=2, Passed=4
Dòng 7: Score=0.800, Failed=2, Passed=4
Dòng 4: Score=0.800, Failed=2, Passed=4
```

**Kết quả**: Cả 3 dòng đều nghi ngờ như nhau!
- Nhưng phải chọn 1 dòng để kiểm tra trước
- Chọn dòng 4 vì line number nhỏ nhất
- Developer nên kiểm tra cả 3 dòng theo thứ tự

**Lời khuyên thực tế**:
- Fault localization chỉ là **gợi ý**, không phải kết luận chắc chắn
- Nên kiểm tra TOP 3-5 dòng có điểm cao nhất
- Kết hợp với kinh nghiệm và logic để tìm lỗi

#### Ứng dụng
- Automated debugging
- Regression testing
- Code quality analysis

---

## Kết hợp các kỹ thuật

Trong thực tế, các kỹ thuật thường được kết hợp:

1. **Tracing** → Ghi lại execution
2. **Execution Indexing** → Đánh số thứ tự
3. **Dynamic Slicing** → Tìm code liên quan
4. **Fault Localization** → Xác định vị trí lỗi

## Tips sử dụng

### Khi nào dùng Tracing?
- Muốn hiểu chương trình chạy như thế nào
- Debug thuật toán phức tạp
- Giảng dạy/học tập

### Khi nào dùng Dynamic Slicing?
- Chương trình quá lớn, khó hiểu
- Muốn tìm code ảnh hưởng đến một biến
- Chuẩn bị refactor code

### Khi nào dùng Execution Indexing?
- Debug vòng lặp
- Phân tích đa luồng
- Cần time-travel debugging

### Khi nào dùng Fault Localization?
- Có nhiều test cases failed
- Không biết lỗi ở đâu
- Muốn ưu tiên kiểm tra

## Tham khảo thêm

### Papers
- "Efficient Program Slicing Algorithms" - Mark Weiser
- "The Tarantula Algorithm" - Jones & Harrold
- "Ochiai: A Family of Fault Localization" - Abreu et al.

### Tools thực tế
- **GDB**: GNU Debugger (có tracing)
- **rr**: Record and Replay debugger
- **Frama-C**: Static & Dynamic Analysis
- **Pinpoint**: Fault Localization tool

## Liên hệ

Nếu có thắc mắc về project, vui lòng tạo issue trên repository.

---

**Chúc bạn học tốt! 🎓**

