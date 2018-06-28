//login
$(function () {
    $("#commit").click(function (e) {
        if ($("#username").val().length == 0) {
            alert("请输入账号")
            e.preventDefault();
        }
        if ($("#passwd").val().length == 0) {
            alert("请输入密码")
            e.preventDefault();
        }
    });
})