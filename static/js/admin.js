function parseDate(time) {
    time = time.replace(/-/g, "/").replace(/T/g, " ");
    var date = new Date(Date.parse(time));
    return (date.getMonth() + 1) + "-" + date.getDate()
}

function parseContent(str) {
    var len = 0;
    if (str == null)
        len = 0;
    else if (typeof str != "string") {
        str += "";
    } else
        len = str.replace(/[^\x00-\xff]/g, "01").length;
    if (len <= 10) {
        return str;
    } else
        return str.substring(0, 10) + "...";
}

function parseName(str) {
    var len = 0;
    if (str == null)
        len = 0;
    else if (typeof str != "string") {
        str += "";
    } else
        len = str.replace(/[^\x00-\xff]/g, "01").length;
    if (len <= 6) {
        return str;
    } else
        return str.substring(0, 6) + "...";
}

function parsePermission(p) {
    if (p === 0) {
        return "仅可评论"
    } else if (p === 1) {
        return "上传文件"
    } else if (p >= 5) {
        return "管理员"
    }
}

function parseId(id) {
    return id.split("-")[2];
}

function getComments() {
    $.ajax({
        type: "get",
        url: "/comment/getAllComments",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                msg = msg + "<li>";
                msg = msg + "<span>日期</span>";
                msg = msg + "<label>用户</label>";
                msg = msg + "<s>内容</s>";
                msg = msg + "<em>操作</em>";
                msg = msg + "</li>";

                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<span>" + parseDate(payload[i].comment_time) + "</span>";
                    msg = msg + "<label>" + parseName(payload[i].user) + "</label>";
                    msg = msg + "<s>" + parseContent(payload[i].content) + "</s>";
                    msg = msg + "<em id='comments-update-" + payload[i].id + "'>&olarr;</em>";
                    msg = msg + "<em id='comments-del-" + payload[i].id + "'>&times;&nbsp;&nbsp;</em>";
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

function getMoves() {
    $.ajax({
        type: "get",
        url: "/move/getMoves",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                msg = msg + "<li>";
                msg = msg + "<span>日期</span>";
                msg = msg + "<label>用户</label>";
                msg = msg + "<s>内容</s>";
                msg = msg + "<em>操作</em>";
                msg = msg + "</li>";
                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<span>" + parseDate(payload[i].time) + "</span>";
                    msg = msg + "<label>" + parseName(payload[i].user) + "</label>";
                    msg = msg + "<s>" + parseContent(payload[i].content) + "</s>";
                    msg = msg + "<em id='moves-update-" + payload[i].id + "'>&olarr;</em>";
                    msg = msg + "<em id='moves-del-" + payload[i].id + "'>&times;&nbsp;&nbsp;</em>";
                    msg = msg + "</li>";
                }
                $("#moves").empty();
                $("#moves").append(msg);
            } else {
                alert("系统内部错误");
            }
        }
    });


}

function getUsers() {
    $.ajax({
        type: "get",
        url: "/user/getUsers",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                msg = msg + "<li>";
                msg = msg + "<span>日期</span>";
                msg = msg + "<label>用户</label>";
                msg = msg + "<s>权限</s>";
                msg = msg + "<em>操作</em>";
                msg = msg + "</li>";
                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<span>" + parseDate(payload[i].register_time) + "</span>";
                    msg = msg + "<label>" + parseName(payload[i].name) + "</label>";
                    msg = msg + "<s>" + parsePermission(payload[i].permission) + "</s>";
                    msg = msg + "<em id='users-update-" + payload[i].id + "'>&olarr;</em>";
                    msg = msg + "<em id='users-del-" + payload[i].id + "'>&times;&nbsp;&nbsp;</em>";
                    msg = msg + "</li>";
                }
                $("#users").empty();
                $("#users").append(msg);
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function getAnnounces() {
    $.ajax({
        type: "get",
        url: "/announce/getAnnounces",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                msg = "";
                msg = msg + "<li>";
                msg = msg + "<span>日期</span>";
                msg = msg + "<label>用户</label>";
                msg = msg + "<s>内容</s>";
                msg = msg + "<em>操作</em>";
                msg = msg + "</li>";
                for (i = payload.length - 1, end = 0; i >= end; i--) {
                    msg = msg + "<li>";
                    msg = msg + "<span>" + parseDate(payload[i].time) + "</span>";
                    msg = msg + "<label>" + parseName(payload[i].user) + "</label>";
                    msg = msg + "<s>" + parseContent(payload[i].content) + "</s>";
                    msg = msg + "<em id='announces-update-" + payload[i].id + "'>&olarr;</em>";
                    msg = msg + "<em id='announces-del-" + payload[i].id + "'>&times;&nbsp;&nbsp;</em>";
                    msg = msg + "</li>";
                }
                $("#announces").empty();
                $("#announces").append(msg);
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delComment(id) {
    $.ajax({
        type: "get",
        url: "/admin/delComment",
        data: {
            'commentId': id
        },
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert(result.payload)
            } else if (result.status === 0) {
                alert("删除成功");
                getComments();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delMove(id) {
    $.ajax({
        type: "get",
        url: "/admin/delMove",
        data: {
            'moveId': id
        },
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert(result.payload)
            } else if (result.status === 0) {
                alert("删除成功");
                getMoves();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delUser(id) {
    $.ajax({
        type: "get",
        url: "/admin/delUser",
        data: {
            'userId': id
        },
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert(result.payload)
            } else if (result.status === 0) {
                alert("删除成功");
                getUsers();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

function delAnnounce(id) {
    $.ajax({
        type: "get",
        url: "/admin/delAnnounce",
        data: {
            'announceId': id
        },
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert(result.payload)
            } else if (result.status === 0) {
                alert("删除成功");
                getAnnounces();
            } else {
                alert("系统内部错误");
            }
        }
    });
}

$(function () {
    getComments();
    getMoves();
    getUsers();
    getAnnounces();

    $('#CommentModal').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    $('#MoveModal').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    $('#UserModal').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    $('#AnnounceModal').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    $('#NewAnnounceModal').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });


    $('#loginForm').on('submit', function () {

        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/user/login', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    location.href = "/";
                } else if (result.status === 0) {
                    location.href = "/admin"
                } else {
                    location.href = "/";
                }
            }
        });
        $("#myModal").modal("hide");
        return false; // 阻止表单自动提交事件
    });

    $('#commentForm').on('submit', function () {

        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/admin/updateComment', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload);
                } else if (result.status === 0) {

                } else {
                    alert("更新失败")
                }
            }
        });
        $("#CommentModal").modal("hide");
        getComments();
        return false; // 阻止表单自动提交事件
    });

    $('#moveForm').on('submit', function () {

        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/admin/updateMove', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload);
                } else if (result.status === 0) {
                    getMoves();
                } else {
                    alert("更新失败")
                }
            }
        });
        $("#MoveModal").modal("hide");
        return false; // 阻止表单自动提交事件
    });

    $('#userForm').on('submit', function () {

        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/admin/updateUser', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload);
                } else if (result.status === 0) {
                    getUsers();
                } else {
                    alert("更新失败")
                }
            }
        });
        $("#UserModal").modal("hide");

        return false; // 阻止表单自动提交事件
    });

    $('#announceForm').on('submit', function () {
        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/admin/updateAnnounce', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload);
                } else if (result.status === 0) {
                    getAnnounces();
                } else {
                    alert("更新失败")
                }
            },
        });
        $("#AnnounceModal").modal("hide");

        return false;
    });

    $('#newAnnounceForm').on('submit', function () {

        $(this).ajaxSubmit({
            type: 'post', // 提交方式 get/post
            url: '/admin/addAnnounce', // 需要提交的 url
            success: function (data) { // data 保存提交后返回的数据，一般为 json 数据
                // 此处可对 data 作相关处理
                result = $.parseJSON(data);
                if (result.status === 1) {
                    alert(result.payload);
                } else if (result.status === 0) {
                    getAnnounces();
                } else {
                    alert("更新失败")
                }
            }
        });
        $("#NewAnnounceModal").modal("hide");

        return false; // 阻止表单自动提交事件
    });


    $("#comments").on('click', '[id^=comments-del-]', function () {
        var id = $(this).attr("id");
        var commentId = parseId(id);
        var msg = confirm("确定要删除本信息？")
        if (msg == true) {
            delComment(commentId);
        }
    });

    $("#moves").on('click', '[id^=moves-del-]', function () {
        var id = $(this).attr("id");
        var moveId = parseId(id);
        var msg = confirm("确定要删除本信息？")
        if (msg == true) {
            delMove(moveId);
        }
    });

    $("#users").on('click', '[id^=users-del-]', function () {
        var id = $(this).attr("id");
        var userId = parseId(id);
        var msg = confirm("确定要删除本信息？")
        if (msg == true) {
            delUser(userId);
        }
    });

    $("#announces").on('click', '[id^=announces-del-]', function () {
        var id = $(this).attr("id");
        var announceId = parseId(id);
        var msg = confirm("确定要删除本信息？")
        if (msg == true) {
            delAnnounce(announceId);
        }
    });

    $("#moves").on('click', '[id^=moves-update-]', function () {
        var id = $(this).attr("id");
        var moveId = parseId(id);
        $.ajax({
            type: "get",
            url: "/move/getMove",
            data: {
                "mid": moveId
            },
            success: function (rr) {
                result = $.parseJSON(rr);
                if (result.status === 1) {
                    alert("failed")
                } else if (result.status === 0) {
                    $("#moveId").attr("value", moveId);
                    $("#moveUser").attr("value", result.payload.user);
                    $("#moveContent").val(result.payload.content);
                    $("#MoveModal").modal("show");
                } else {
                    alert("系统内部错误");
                }
            }
        });


    });

    $("#comments").on('click', '[id^=comments-update-]', function () {
        var id = $(this).attr("id");
        var commentId = parseId(id);
        $.ajax({
            type: "get",
            url: "/comment/getComment",
            data: {
                "commentId": commentId
            },
            success: function (rr) {
                result = $.parseJSON(rr);
                if (result.status === 1) {
                    alert("failed")
                } else if (result.status === 0) {
                    $("#commentId").attr("value", commentId);
                    $("#commentUser").attr("value", result.payload.user);
                    $("#commentContent").val(result.payload.content);
                    $("#CommentModal").modal("show");
                } else {
                    alert("系统内部错误");
                }
            }
        });


    });

    $("#users").on('click', '[id^=users-update-]', function () {
        var id = $(this).attr("id");
        var userId = parseId(id);
        $.ajax({
            type: "get",
            url: "/user/getUser",
            data: {
                "userId": userId
            },
            success: function (rr) {
                result = $.parseJSON(rr);
                if (result.status === 1) {
                    alert("failed")
                } else if (result.status === 0) {
                    $("#userId").attr("value", userId);
                    $("#username").attr("value", result.payload.name);
                    $("#password").attr("value", result.payload.password);
                    $("#UserModal").modal("show");
                } else {
                    alert("系统内部错误");
                }
            }
        });


    });

    $("#announces").on('click', '[id^=announces-update-]', function () {
        var id = $(this).attr("id");
        var announceId = parseId(id);
        $.ajax({
            type: "get",
            url: "/announce/getAnnounce",
            data: {
                "announceId": announceId
            },
            success: function (rr) {
                result = $.parseJSON(rr);
                if (result.status === 1) {
                    alert("failed")
                } else if (result.status === 0) {
                    $("#announceId").attr("value", result.payload.id);
                    $("#announceUser").attr("value", result.payload.user);
                    $("#announceContent").val(result.payload.content);
                    $("#announceDisplay").attr("checked", result.payload.display);
                    $("#AnnounceModal").modal("show");
                } else {
                    alert("系统内部错误");
                }
            }
        });


    });
});