# event-tracking

## Hướng dẫn cài đặt:

### Bước 1: 
Chạy lệnh sau để khởi động các dịch vụ Docker:

```bash
docker compose up
```

### Bước 2: 
Truy cập vào địa chỉ http://localhost:19000/, sau đó tạo topic mong muốn.

### Bước 3: 
Thay đổi các biến môi trường của database và Kafka trong file dev.yaml cho phù hợp với cấu hình của bạn.

### Bước 4:
Để sử dụng được kafka ở local, cần comment out **Dialer** của kafkaReader và **Transport**  của KafkaWriter

### Bước 5: 
Run 

```bash
go run cmd/server/main.go
```

