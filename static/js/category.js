$(function () {
    $("#commit").click(function (e) {
        if ($("#name").val().length == 0) {
            alert("请输入分类名称")
            e.preventDefault()
        }
    })
})