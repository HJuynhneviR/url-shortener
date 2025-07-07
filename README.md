# Week 1 - URL Shortener (Go)

Ứng dụng rút gọn URL đơn giản bằng ngôn ngữ Go thuần.

## ✅ Tính năng
- Tạo URL rút gọn từ URL gốc
- Truy cập short URL sẽ redirect về link gốc
- Lưu dữ liệu trong RAM (map)

## 🧪 Test API
- Tạo short URL:
```bash
curl.exe -X POST -d "url=https://youtube.com" http://localhost:8080/shorten
