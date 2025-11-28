# QUICKSTART - BẮT ĐẦU NHANH

## Cách chạy nhanh nhất

### Windows:
```bash
# Cách 1: Double-click file run.bat
run.bat

# Cách 2: Chạy trực tiếp
go run main.go

# Cách 3: Build rồi chạy
go build -o program-analysis-demo.exe .
.\program-analysis-demo.exe
```

### Linux/Mac:
```bash
# Chạy trực tiếp
go run main.go

# Hoặc build rồi chạy
go build -o program-analysis-demo .
./program-analysis-demo
```

## Menu chương trình

```
╔════════════════════════════════════════════════════════════╗
║        PHÂN TÍCH CHƯƠNG TRÌNH - PROGRAM ANALYSIS          ║
╠════════════════════════════════════════════════════════════╣
║  1. Tracing - Theo dấu thực thi chương trình              ║
║  2. Dynamic Slicing - Cắt động chương trình               ║
║  3. Execution Indexing - Đánh chỉ mục thực thi            ║
║  4. Fault Localization - Định vị lỗi                      ║
║  5. Thoát                                                  ║
╚════════════════════════════════════════════════════════════╝
```

## Chọn demo

Nhập số từ **1-5** và nhấn Enter:

- **1**: Xem demo Tracing (2 ví dụ: factorial và findMax)
- **2**: Xem demo Dynamic Slicing (2 ví dụ: sum/product và findMax)
- **3**: Xem demo Execution Indexing (3 ví dụ: loop, nested loop, multi-thread)
- **4**: Xem demo Fault Localization (2 ví dụ: average và findMax)
- **5**: Thoát chương trình

## Mỗi demo sẽ hiển thị

1. 📝 **Code mẫu**: Chương trình được phân tích
2. ▶️ **Thực thi**: Chạy chương trình và ghi lại trace
3. 📊 **Kết quả**: Hiển thị phân tích dưới dạng bảng
4. 💡 **Giải thích**: Ý nghĩa và ứng dụng thực tế

## Ví dụ output

### Tracing:
```
+------+---------------------------------------------+------------------------+
| Dòng | Câu lệnh                                    | Biến                   |
+------+---------------------------------------------+------------------------+
|    1 | result := 1                                 | n=5, result=1          |
|    2 | for i := 1; i <= 5                         | i=1, n=5, result=1     |
|    3 | result = 1 * 1 = 1                         | i=1, result=1          |
```

### Dynamic Slicing:
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
```

### Execution Indexing:
```
+-------+------+------+----------------------------------------+-----------------+
| Index | Dòng | TID  | Câu lệnh                               | Biến            |
+-------+------+------+----------------------------------------+-----------------+
|     1 |    1 | T1   | n := 5                                 | n=5             |
|     2 |    2 | T1   | sum := 0                               | sum=0           |
```

### Fault Localization:
```
🔍 SUSPICIOUS SCORES (Phương pháp: Tarantula):
+------+---------------------------------------+-------+--------+--------+
| Dòng | Statement                             | Score | Failed | Passed |
+------+---------------------------------------+-------+--------+--------+
|    6 | avg := sum / len(arr)                 | 0.800 |      2 |      4 |
|    7 | return avg                            | 0.800 |      2 |      4 |
```

## Tips

- Mỗi demo chạy độc lập, có thể chọn bất kỳ thứ tự nào
- Sau mỗi demo, nhấn Enter để quay lại menu
- Đọc file `HUONG_DAN.md` để hiểu chi tiết hơn
- Xem `README.md` cho thông tin tổng quan

## Yêu cầu hệ thống

- Go 1.21 trở lên
- Windows 10/11, Linux, hoặc macOS
- Terminal/Command Prompt hỗ trợ UTF-8

## Troubleshooting

**Lỗi: "go: command not found"**
→ Cài đặt Go từ https://golang.org/dl/

**Lỗi: "cannot find package"**
→ Chạy `go mod tidy` trước

**Hiển thị ký tự lạ trên Windows**
→ Chạy `chcp 65001` để đổi sang UTF-8

---

**Bắt đầu ngay: `go run main.go`** 🚀

