# AutoPaster

Phần mềm giao diện dòng lệnh giúp lấy dữ liệu từ Excel và nhập vào 1 phần mềm khác.

## Hướng dẫn sử dụng

1. Tải về autopaster.exe.
2. Để cùng thư mục với file excel cần nhập.
3. Tạo file .env với nội dung như sau.
4. Chọn vị trí cần nhập dữ liệu và nhấn ctrl + i.
5. Chờ nhập xong.

```.env
FILENAME=tenfile.xlsx
SHEETNAME=tensheet
COL_INDEX=[vị trí cột sử dụng, bắt đầu đếm từ 0. ví dụ cột A là 0, cột B là 1,...]
ROW_START_INDEX=[vị trí hàng bắt đầu đọc]
```

## Tech

Sử dụng ngôn ngữ Go với 2 thư viện [robotgo](https://github.com/go-vgo/robotgo) và [excelize](https://github.com/qax-os/excelize)