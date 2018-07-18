$(function () {

    $('#upComment').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });


    $('#upMove').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    $('#commentForm').on('submit', function () {

        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/comment/addComment', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload)
                } else if (result.status === 0) {
                    getComments();
                } else {
                    alert("系统内部错误");
                }
            }
        });
        $("#upComment").modal("hide");
        return false; // 阻止表单自动提交事件
    });







});

function index() {
    getAnnounce();
    getComments();
}


function getComments() {
    $.ajax({
        type: "get",
        url: "/comment/getComments",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<i class=\"glyphicon glyphicon-flag\"></i>";
                    msg = msg + "<label>" + payload[i].user + "</label>";
                    msg = msg + "<s>" + payload[i].content + "</s>";
                    msg = msg + "</li>";
                }
                $("#comments").empty();
                $("#comments").append(msg);
            } else {
                alert("系统内部错误");
            }
        }
    });

}

function getAnnounce() {
    $.ajax({
        type: "get",
        url: "/announce/getLatestAnnounce",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 0) {
                payload = result.payload;
                $("#announceText").empty();
                $("#announceText").append(payload.content);
                $("#announce").show();
            } else {
                $("#announce").hide();
            }
        }
    });
}