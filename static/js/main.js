$(function () {

    $('#comment-submit').on('click', function () {

        $('form').on('submit', function () {
            var username = $('input[name=name]').val(),
                password = $('input[name=password]').val(),
                content = $('textarea').val();

            $(this).ajaxSubmit({
                type: 'post', // 提交方式 get/post
                url: '/comment/addComment', // 需要提交的 url
                data: {
                    'name': username,
                    'password': password,
                    'content': content
                },
                success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                    // 此处可对 data 作相关处理
                    alert('提交成功！');
                }
            });
            $("#myModal").model("hide");
            return false; // 阻止表单自动提交事件
        });
    });

    $('#move-comment-submit').on('click', function () {

        $('form').on('submit', function () {
            var username = $('input[name=username]').val(),
                password = $('input[name=password]').val(),
                content = $('textarea').val(),
                mid = $('input[name=mid]').val();

            $(this).ajaxSubmit({
                type: 'post', // 提交方式 get/post
                url: '/move/addMoveComment', // 需要提交的 url
                data: {
                    'username': username,
                    'password': password,
                    'content': content,
                    'mid': mid
                },
                success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                    // 此处可对 data 作相关处理
                    alert('提交成功！');
                }
            });
            $("#myModal").model("hide");
            return false; // 阻止表单自动提交事件
        });
    });


    var portraitOptions = {
        success: showResponse,
        url: "/portrait/addPortrait",
        type: "post",
        clearForm: true,
        resetForm: true,
        timeout: 3000
    };
    $('#portrait-form').ajaxForm(portraitOptions);


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

function updateComment() {
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
                    msg = msg + "<label>" + payload[i].u_name + "</label>";
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

function updateMoveComment(mid) {
    $.ajax({
        type: "get",
        url: "/move/getMoveComments?mid="+mid,
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<label>" + payload[i].u_name + "</label>";
                    msg = msg + "<pre>" + payload[i].content + "</pre>";
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