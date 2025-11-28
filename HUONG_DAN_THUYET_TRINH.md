# HƯỚNG DẪN THUYẾT TRÌNH

## Chuẩn bị trước khi thuyết trình

### 1. Môi trường
- ✅ Máy tính có Go đã cài đặt
- ✅ Project đã build thành công
- ✅ Terminal/CMD sẵn sàng
- ✅ Slides hoặc file BAI_THUYET_TRINH.md mở sẵn

### 2. Demo
```bash
# Test chạy trước
go run main.go

# Thử từng option 1-4
# Đảm bảo output hiển thị đúng
```

### 3. Backup
- In slides ra giấy phòng lỗi máy chiếu
- Copy project vào USB
- Chuẩn bị video demo (nếu có)

---

## Cấu trúc thuyết trình (45 phút)

### Phần 1: Giới thiệu (5 phút)
**Nội dung:**
- Chào mừng và giới thiệu đề tài
- Tại sao cần Program Analysis?
- Overview 4 kỹ thuật

**Script:**
```
"Xin chào các bạn, hôm nay tôi sẽ trình bày về 4 kỹ thuật 
phân tích chương trình: Tracing, Dynamic Slicing, Execution 
Indexing và Fault Localization.

Trong thực tế, khi code ngày càng phức tạp với hàng triệu 
dòng, việc tìm lỗi và hiểu code trở nên rất khó khăn. 
70% thời gian phát triển phần mềm là debug! 

Các kỹ thuật này giúp chúng ta tự động hóa việc phân tích, 
tìm lỗi nhanh hơn và hiểu code tốt hơn."
```

**Slides:** 1-5

---

### Phần 2: Tracing (8 phút)

**Nội dung:**
- Giải thích khái niệm Tracing
- Ví dụ factorial(5)
- Show execution trace
- Ưu nhược điểm

**Script:**
```
"Kỹ thuật đầu tiên là Tracing - Theo dấu thực thi.

Ý tưởng rất đơn giản: Ghi lại MỌI câu lệnh được thực thi 
cùng với giá trị biến tại thời điểm đó.

Hãy xem ví dụ tính giai thừa của 5..."
[Show code và trace table]

"Như các bạn thấy, function chỉ có 4 dòng code nhưng 
thực thi 12 lần! Đây chính là sức mạnh của Tracing - 
giúp ta hiểu CHÍNH XÁC chương trình chạy như thế nào."
```

**Demo:**
```bash
# Chạy option 1
go run main.go
> 1
```

**Slides:** 6-11

---

### Phần 3: Dynamic Slicing (10 phút)

**Nội dung:**
- Vấn đề: Code quá lớn, cần tập trung
- Slicing Criterion
- Ví dụ sum vs product
- So sánh 2 slices khác nhau
- Thuật toán backward slicing

**Script:**
```
"Tracing cho ta mọi thông tin, nhưng đôi khi quá nhiều!
Ví dụ chương trình 1000 dòng, lỗi ở biến 'sum' tại dòng 500.
Ta không muốn xem cả 1000 dòng!

Dynamic Slicing giải quyết vấn đề này. Nó TÌM chỉ những 
dòng code ảnh hưởng đến biến ta quan tâm.

Hãy xem ví dụ này..."
[Show code sum và product]

"Khi slice cho 'sum', ta chỉ thấy dòng 1, 3, 4.
Dòng 2 và 5 về 'product' biến mất!

Khi slice cho 'product', ngược lại!

Từ 12 dòng thực thi → chỉ còn 3 dòng. 
Rút gọn 75%!"
```

**Demo:**
```bash
# Chạy option 2
go run main.go
> 2
```

**Slides:** 12-21

---

### Phần 4: Execution Indexing (8 phút)

**Nội dung:**
- Vấn đề: Cùng dòng chạy nhiều lần
- Giải pháp: Gán index duy nhất
- Ví dụ với loop
- Query operations
- Ứng dụng multi-threading

**Script:**
```
"Một vấn đề với Tracing và Slicing là: Cùng một dòng code 
trong loop sẽ chạy nhiều lần. Làm sao phân biệt?

Execution Indexing giải quyết bằng cách gán một SỐ THỨ TỰ 
duy nhất cho MỖI LẦN thực thi.

[Show table với index]

Bây giờ ta có thể:
- Lấy thực thi tại index 7
- Lấy tất cả từ index 5 đến 10
- Tìm dòng 4 được chạy bao nhiêu lần

Đặc biệt hữu ích cho multi-threading! 
Ta có thể biết chính xác Thread nào chạy trước."
```

**Demo:**
```bash
# Chạy option 3
go run main.go
> 3
```

**Slides:** 22-29

---

### Phần 5: Fault Localization (10 phút)

**Nội dung:**
- Vấn đề: Có bug nhưng không biết ở đâu
- Nguyên lý: So sánh passed vs failed tests
- Công thức Tarantula
- Ví dụ lỗi chia integer
- Suspicious scores
- Tie-breaking strategy

**Script:**
```
"Đây là kỹ thuật thú vị nhất: TỰ ĐỘNG TÌM LỖI!

Ý tưởng: Nếu có nhiều test cases, một số passed, 
một số failed, ta có thể phân tích:

'Dòng nào xuất hiện NHIỀU trong failed tests nhưng 
ÍT trong passed tests? → Đó có thể là lỗi!'

[Show công thức Tarantula]

Hãy xem ví dụ..."
[Show code average với bug]

"Dòng 6 có suspicious score cao nhất: 0.800
Đúng vậy! Đây chính là dòng có lỗi.

Với Fault Localization, thay vì kiểm tra 8 dòng,
ta CHỈ cần kiểm tra top 2-3 dòng!

Tiết kiệm 60-70% thời gian debug!"
```

**Demo:**
```bash
# Chạy option 4
go run main.go
> 4
```

**Điểm nhấn:**
- Giải thích tại sao dòng 6 và 7 cùng điểm
- Tie-breaking strategy
- Không phải 100% chính xác

**Slides:** 30-42

---

### Phần 6: So sánh & Ứng dụng (3 phút)

**Nội dung:**
- Bảng so sánh 4 kỹ thuật
- Khi nào dùng kỹ thuật nào
- Kết hợp các kỹ thuật
- Công cụ thực tế

**Script:**
```
"Vậy khi nào dùng kỹ thuật nào?

[Show comparison table]

Và thú vị hơn: Ta có thể KẾT HỢP chúng!

Workflow thực tế:
1. Chạy tests → có failed
2. Fault Localization → tìm top 5 dòng nghi ngờ
3. Dynamic Slicing → tìm code liên quan
4. Tracing/Indexing → phân tích chi tiết
5. Fix bug!

Các công cụ như GDB, rr, Frama-C đã áp dụng 
những kỹ thuật này."
```

**Slides:** 43-47

---

### Phần 7: Kết luận (1 phút)

**Nội dung:**
- Tóm tắt 4 kỹ thuật
- Bài học kinh nghiệm
- Hướng phát triển

**Script:**
```
"Tóm lại, chúng ta đã học 4 kỹ thuật:

1. Tracing - Hiểu luồng thực thi
2. Dynamic Slicing - Tập trung vào code liên quan
3. Execution Indexing - Phân biệt các lần chạy
4. Fault Localization - Tự động tìm lỗi

Tất cả đều giúp chúng ta debug và hiểu code 
hiệu quả hơn rất nhiều!

Cảm ơn các bạn đã lắng nghe!"
```

**Slides:** 48-52

---

### Q&A (10 phút dự phòng)

**Câu hỏi có thể có:**

**Q1: "Overhead cao như vậy, dùng trong production được không?"**
```
A: Thường chỉ dùng trong development và testing.
Production có thể dùng sampling (chỉ trace 1-5% requests).
Hoặc chỉ bật khi phát hiện lỗi.
```

**Q2: "Dynamic Slicing có thể áp dụng cho JavaScript không?"**
```
A: Có! Bất kỳ ngôn ngữ nào. Chỉ cần instrument code 
để ghi lại execution trace. Đã có tools cho JS, Python, Java...
```

**Q3: "Tại sao không dùng debugger thay vì Tracing?"**
```
A: Debugger là interactive, phải dừng và step qua từng dòng.
Tracing tự động, ghi lại toàn bộ, có thể phân tích sau.
Mỗi cái có use case riêng.
```

**Q4: "Fault Localization có thể 100% chính xác không?"**
```
A: Không. Nó chỉ là "gợi ý thông minh" dựa trên thống kê.
Nhưng đã tốt hơn random rất nhiều! 
Research shows: Top-5 accuracy ~80%.
```

**Q5: "Code của bạn có áp dụng được thực tế không?"**
```
A: Code này là demo cho mục đích học tập.
Thực tế cần optimize hơn: 
- Giảm overhead
- Handle code phức tạp hơn
- Scale cho chương trình lớn
```

---

## Tips thuyết trình

### Giọng điệu
- 🎤 Nói rõ ràng, không quá nhanh
- 🎤 Nhiệt tình, tương tác với audience
- 🎤 Dừng sau mỗi phần quan trọng

### Body language
- 👀 Nhìn vào audience, không chỉ nhìn slides
- 🖐️ Dùng tay để nhấn mạnh
- 🚶 Di chuyển nhẹ, không đứng cứng

### Tương tác
- ❓ Hỏi audience: "Các bạn đã gặp vấn đề debug lâu chưa?"
- ❓ "Ai đã dùng debugger? Tốn bao nhiêu thời gian?"
- ❓ "Có bạn nào biết GDB không?"

### Xử lý sự cố
- **Lỗi kỹ thuật:** Giữ bình tĩnh, dùng backup plan
- **Câu hỏi khó:** "Câu hỏi hay! Tôi sẽ tìm hiểu thêm sau buổi này"
- **Hết thời gian:** Có slide summary để skip nếu cần

---

## Checklist trước khi present

### 1 tuần trước:
- [ ] Hoàn thiện slides
- [ ] Test code kỹ
- [ ] Luyện tập presentation 2-3 lần
- [ ] Xin feedback từ bạn bè

### 1 ngày trước:
- [ ] Review slides lần cuối
- [ ] Chuẩn bị backup (USB, print)
- [ ] Test máy chiếu/projector
- [ ] Ngủ đủ giấc

### Trước khi lên:
- [ ] Uống nước
- [ ] Thở sâu, thư giãn
- [ ] Mở sẵn terminal, slides
- [ ] Test audio (nếu có)
- [ ] Tắt notifications

---

## Phân công nếu nhóm

**Người 1: Giới thiệu + Tracing (13 phút)**
- Slides 1-11
- Demo Tracing

**Người 2: Dynamic Slicing + Execution Indexing (18 phút)**
- Slides 12-29
- Demo cả 2 kỹ thuật

**Người 3: Fault Localization + Kết luận (14 phút)**
- Slides 30-52
- Demo Fault Localization
- Tổng kết và Q&A

---

## Resources bổ sung

### Để hiểu sâu hơn:
1. Video về Program Analysis trên YouTube
2. Papers gốc (có trong slides)
3. Documentation của GDB, rr

### Để trả lời câu hỏi:
1. Đọc kỹ code đã viết
2. Chạy thử nhiều test cases
3. Research về tools thực tế

### Backup slides:
- Có thể skip phần "Phụ lục" nếu không đủ thời gian
- Phần benchmark có thể bỏ qua
- Focus vào 4 kỹ thuật chính

---

## Đánh giá sau presentation

### Tự đánh giá:
- Thời gian: Đúng 45 phút không?
- Nội dung: Đủ rõ ràng không?
- Demo: Chạy smooth không?
- Q&A: Trả lời tốt không?

### Xin feedback:
- "Phần nào chưa rõ?"
- "Demo có dễ hiểu không?"
- "Tốc độ nói có vừa không?"

---

## Good luck! 🍀

**Remember:**
- Confidence is key
- Know your material
- Practice makes perfect
- Have fun!

**You got this! 💪**

