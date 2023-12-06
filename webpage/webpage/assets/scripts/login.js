function submitForm() {
    var fullName = document.getElementById("fullName").value;
    var phoneNumber = document.getElementById("phoneNumber").value;
    var dob = document.getElementById("dob").value;
    var address = document.getElementById("address").value;

    // Kiểm tra xem dữ liệu có hợp lệ không
    if (fullName && phoneNumber && dob && address) {
        // Ở đây bạn có thể thực hiện các thao tác khác như gửi dữ liệu đăng ký lên server
        alert("Đăng ký thành công!");
    } else {
        alert("Vui lòng điền đầy đủ thông tin.");
    }
}