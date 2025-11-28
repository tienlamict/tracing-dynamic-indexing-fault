# CHƯƠNG TRÌNH TRÌNH BÀY NHANH
## 4 Kỹ thuật Phân tích chương trình (45 phút)

---

## 1. GIỚI THIỆU (5 phút) [Slides 1-5]

**Điểm chính:**
- Program Analysis giúp tự động phân tích code
- 70% thời gian dev là debug
- 4 kỹ thuật: Tracing, Slicing, Indexing, Localization

**Quote mở đầu:**
> "Debugging là quá trình loại bỏ bugs. Programming là quá trình đưa bugs vào."

---

## 2. TRACING (8 phút) [Slides 6-11]

**Key points:**
- Ghi lại MỌI câu lệnh + giá trị biến
- Ví dụ: factorial(5) → 4 dòng code, 12 lần thực thi
- Ưu: Hiểu rõ 100% | Nhược: Overhead cao, data nhiều

**Demo:** Option 1

**Câu nói nhớ:**
> "Biết chương trình ĐÃ chạy gì, không chỉ sẽ chạy gì"

---

## 3. DYNAMIC SLICING (10 phút) [Slides 12-21]

**Key points:**
- Tìm code ảnh hưởng đến biến X tại dòng Y
- Slicing Criterion: `<line, variable>`
- Ví dụ: sum vs product → 2 slices khác nhau
- Rút gọn 75% code!

**Demo:** Option 2

**Công thức:**
```
Backward Slicing:
1. Bắt đầu từ criterion
2. Tìm dòng định nghĩa biến
3. Thêm variables used
4. Lặp lại
```

**Câu nói nhớ:**
> "Không cần xem cả rừng, chỉ cần xem cây có vấn đề"

---

## 4. EXECUTION INDEXING (8 phút) [Slides 22-29]

**Key points:**
- Gán index duy nhất cho MỖI lần thực thi
- Phân biệt lần 1, lần 2, lần 3... của cùng 1 dòng
- Query: GetAt(7), GetBetween(5,10), FindByLine(4)
- Ứng dụng: Multi-threading, time-travel debug

**Demo:** Option 3

**Ví dụ quan trọng:**
- Dòng 4 chạy 5 lần → Index: 4, 6, 8, 10, 12
- Có thể quay lại index 7 để xem state

**Câu nói nhớ:**
> "Mỗi lần chạy là một câu chuyện riêng"

---

## 5. FAULT LOCALIZATION (10 phút) [Slides 30-42]

**Key points:**
- TỰ ĐỘNG tìm vị trí lỗi từ test results
- So sánh passed vs failed tests
- Suspicious score: 0.0 → 1.0
- Tarantula vs Ochiai

**Demo:** Option 4

**Công thức Tarantula:**
```
score = (failed/totalFailed) / 
        ((failed/totalFailed) + (passed/totalPassed))
```

**Ví dụ:**
- 6 tests: 4 passed, 2 failed
- Dòng 6: score = 0.800 (cao nhất)
- Đúng! Dòng 6 có bug

**Tie-breaking:**
1. Score cao nhất
2. Failed count nhiều nhất  
3. Line number nhỏ nhất

**Câu nói nhớ:**
> "Thống kê biết bạn không biết"

---

## 6. SO SÁNH (3 phút) [Slides 43-47]

**Bảng nhanh:**
| Kỹ thuật | Mục đích | Khi nào dùng |
|----------|----------|--------------|
| Tracing | Hiểu flow | Học thuật toán mới |
| Slicing | Tập trung | Bug ở 1 biến |
| Indexing | Time-travel | Debug loop, thread |
| Localization | Tìm lỗi | Nhiều test failed |

**Kết hợp:**
```
Tests failed → Localization → Slicing → Indexing → Fix!
```

**Tools thực tế:**
- GDB, rr, Frama-C, Pinpoint

---

## 7. KẾT LUẬN (1 phút) [Slides 48-52]

**Take-away messages:**
1. Dynamic Analysis giúp debug hiệu quả
2. Mỗi kỹ thuật giải quyết vấn đề khác nhau
3. Kết hợp nhiều kỹ thuật cho kết quả tốt nhất
4. Tools chỉ hỗ trợ, con người quyết định

**Call to action:**
- Thử code: `go run main.go`
- Áp dụng vào project thực
- Tìm hiểu thêm qua papers

---

## Q&A CHEAT SHEET (10 phút)

**Q: Overhead cao, dùng production được không?**
> A: Chỉ dùng dev/test. Production dùng sampling (1-5%).

**Q: Áp dụng cho ngôn ngữ nào?**
> A: Tất cả! Chỉ cần instrument code.

**Q: Tại sao không dùng debugger?**
> A: Debugger interactive, phải manual. Tracing tự động, analyze sau.

**Q: FL có 100% accurate không?**
> A: Không. Chỉ 70-80% top-5. Nhưng tốt hơn random!

**Q: Code này dùng thực tế được không?**
> A: Demo cho học tập. Thực tế cần optimize thêm.

---

## CHECKLIST NGAY TRƯỚC KHI LÊN

- [ ] Terminal mở sẵn tại project folder
- [ ] Slides mở sẵn
- [ ] Test: `go run main.go` chạy được
- [ ] Tắt notifications
- [ ] Uống nước, thở sâu
- [ ] Chuẩn bị mental: "Tôi đã chuẩn bị tốt, tôi làm được!"

---

## TIMING (Xem đồng hồ!)

```
00:00 - 05:00  Giới thiệu
05:00 - 13:00  Tracing
13:00 - 23:00  Dynamic Slicing
23:00 - 31:00  Execution Indexing
31:00 - 41:00  Fault Localization
41:00 - 44:00  So sánh & Ứng dụng
44:00 - 45:00  Kết luận
45:00 - 55:00  Q&A
```

---

## NẾU HẾT THỜI GIAN

**Có thể skip:**
- Slides về thuật toán chi tiết
- Multi-threading example
- Benchmark numbers
- Phụ lục

**KHÔNG skip:**
- 4 demos chính
- Ví dụ chính của mỗi kỹ thuật
- Comparison table
- Q&A

---

## LỜI KHUYÊN CUỐI

1. **BREATH** - Thở sâu khi căng thẳng
2. **PAUSE** - Dừng lại sau mỗi slide quan trọng
3. **EYE CONTACT** - Nhìn audience, không chỉ slides
4. **SMILE** - Tự tin và nhiệt tình
5. **TIME** - Xem đồng hồ mỗi 10 phút

**Remember:**
> "They want you to succeed. They're here to learn from you."

---

## EMERGENCY CONTACTS

**Nếu demo fail:**
- Có screenshot backup trong slides
- Có video demo (nếu đã quay)
- Giải thích bằng lời

**Nếu quên nội dung:**
- Xem notes ở chân slides
- Nhìn vào code → sẽ nhớ lại
- Honest: "Xin phép xem lại notes một chút"

**Nếu câu hỏi khó:**
- "Câu hỏi hay! Tôi sẽ research thêm"
- "Có thể thảo luận sau buổi này"
- "Theo hiểu biết của tôi..." (không chắc 100%)

---

## MENTAL NOTES

**Trước khi bắt đầu:**
```
Tôi đã code này.
Tôi hiểu nó.
Tôi có thể giải thích nó.
Let's do this!
```

**Trong khi present:**
```
Speak clearly.
Look confident.
Enjoy the moment.
```

**Sau khi xong:**
```
I did my best.
Learn from feedback.
Celebrate!
```

---

## 🎯 YOU GOT THIS! 🎯

**Good luck!** 🍀

