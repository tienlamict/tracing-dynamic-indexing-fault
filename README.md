# Program Analysis Demo - Minh họa Phân tích Chương trình

Project này minh họa các kỹ thuật phân tích chương trình (Program Analysis) bằng Go.

## Các kỹ thuật được minh họa

### 1. 🔍 Tracing - Theo dấu thực thi
- Ghi lại toàn bộ quá trình thực thi chương trình
- Theo dõi giá trị biến qua từng bước
- Hữu ích cho debugging và hiểu luồng thực thi

### 2. 🔪 Dynamic Slicing - Cắt động chương trình
- Tìm các câu lệnh ảnh hưởng đến một biến tại một điểm cụ thể
- Loại bỏ code không liên quan
- Giúp hiểu dependencies giữa các statements

### 3. 📑 Execution Indexing - Đánh chỉ mục thực thi
- Gán index duy nhất cho mỗi lần thực thi
- Hỗ trợ time-travel debugging
- Phân tích multi-threading và interleaving

### 4. 🎯 Fault Localization - Định vị lỗi
- Tự động tìm vị trí có khả năng chứa lỗi
- Sử dụng các công thức: Tarantula, Ochiai
- Dựa trên kết quả test cases (passed/failed)

## Cài đặt và Chạy

### Yêu cầu
- Go 1.21 trở lên

### Cách chạy

```bash
# Clone hoặc tải project về
cd tracing-dynamic-indexing-fault

# Chạy chương trình
go run main.go
```

## Cấu trúc Project

```
tracing-dynamic-indexing-fault/
├── main.go                           # Entry point với menu lựa chọn
├── go.mod                            # Go module file
├── README.md                         # File này
├── tracing/
│   └── tracing.go                    # Package tracing
├── dynamicslicing/
│   └── dynamicslicing.go             # Package dynamic slicing
├── executionindexing/
│   └── executionindexing.go          # Package execution indexing
└── faultlocalization/
    └── faultlocalization.go          # Package fault localization
```

## Ví dụ sử dụng

### Menu chính

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

### Chọn từng demo

Mỗi demo sẽ:
1. Hiển thị chương trình mẫu
2. Thực thi và minh họa kỹ thuật
3. In kết quả phân tích
4. Giải thích ứng dụng thực tế

## Chi tiết kỹ thuật

### Tracing
- Ghi lại mỗi statement được thực thi
- Lưu trữ line number, code, và giá trị biến
- Hữu ích cho post-mortem debugging

### Dynamic Slicing
- Slicing Criterion: `<line, variable>`
- Tính backward slice từ execution trace
- Xác định data dependencies

### Execution Indexing
- Index tăng dần theo thời gian thực thi
- Hỗ trợ query theo index hoặc range
- Phân biệt các lần thực thi khác nhau

### Fault Localization

#### Công thức Tarantula
```
suspiciousness = (failed/totalFailed) / 
                 ((failed/totalFailed) + (passed/totalPassed))
```

#### Công thức Ochiai
```
suspiciousness = failed / 
                 sqrt(totalFailed * (failed + passed))
```

## Ứng dụng thực tế

- **Debugging**: Tìm và sửa lỗi nhanh hơn
- **Program Comprehension**: Hiểu code phức tạp
- **Testing**: Tối ưu test cases
- **Maintenance**: Đánh giá impact của thay đổi
- **Code Review**: Phân tích luồng thực thi

## Tác giả

Project minh họa cho môn học Principle and Technique of Program Analysis

## License

MIT License - Tự do sử dụng cho mục đích học tập

