$(function () {
    $('#commentForm').on('submit', function () {
        var moveId = $('input[name=moveId]').val();
        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/move/addMoveComment', // 需要提交的 url
            data: {
                "moveId": moveId
            },
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload)
                } else if (result.status === 0) {
                    getMoveComment(moveId);
                } else {
                    alert("系统内部错误");
                }
            }
        });
        $("#upComment").modal("hide");
        return false; // 阻止表单自动提交事件
    });

    $('#upComment').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    var moveOptions = {
        success: showResponse,
        url: "/move/addMove",
        type: "post",
        clearForm: true,
        resetForm: true,
        timeout: 3000
    };
    $('#move-form').ajaxForm(moveOptions);

});

//  提交后
function showResponse(responseText, statusText) {
    result = $.parseJSON(responseText);
    if (result.status === 0) {
        alert("已提交")
    }
    location.reload();

}

function move(mid) {
    getMoveComment(mid);
}

function getMoveComment(mid) {
    $.ajax({
        type: "get",
        url: "/move/getMoveComments?moveId=" + mid,
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<pre>" + payload[i].content + "</pre>";
                    msg = msg + "<div>";
                    msg = msg + "<span id='user'>用户:" + payload[i].user + "</span>";
                    msg = msg + "<span id='time'>发表时间:" + parseDate(payload[i].comment_time) + "</span>";
                    msg = msg + "</div>";
                    msg=msg+"</li>"
                }
                $("#comments").empty();
                $("#comments").append(msg);
            } else {
                alert("系统内部错误");
            }
        }
    });


}


function parseDate(time) {
    time = time.replace(/-/g, "/").replace(/T/g, " ");
    var date = new Date(Date.parse(time));
    return (date.getMonth() + 1) + "-" + date.getDate() + " " + date.getHours() + ":" + date.getMinutes()
}