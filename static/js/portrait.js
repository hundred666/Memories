$(function () {
    $('#upPortrait').on('shown.bs.modal', function () {
        $(this).draggable({
            handle: ".modal-header"   // 只能点击头部拖动
        });
        $(this).css("overflow", "hidden"); // 防止出现滚动条，出现的话，你会把滚动条一起拖着走的
    });

    var portraitOptions = {
        success: showResponse,
        url: "/portrait/addPortrait",
        type: "post",
        clearForm: true,
        resetForm: true,
        timeout: 10000
    };
    $('#portrait-form').ajaxForm(portraitOptions);

    $('#portraits').on('click', 'img', function () {
        var src = $(this).attr("src");
        $("#bigPortrait").attr("src", src);
        $(".container").hide();
        $("#portraitBlock").show();
    });
    $("#portraitBlock").click(function () {
        $(this).hide();
        $(".container").show();
    });

    getPortraits();

});

function showResponse(responseText) {
    result = $.parseJSON(responseText);
    if (result.status === 0) {
        alert("提交成功");
    }
    location.reload();

}

function getPortraits() {
    $.ajax({
        type: "get",
        url: "/portrait/getPortraits",
        success: function (rr) {
            result = $.parseJSON(rr);
            if (result.status === 1) {
                alert("failed")
            } else if (result.status === 0) {
                payload = result.payload;
                var msg = "";
                for (var i = 0; i < payload.length; i++) {
                    if (i % 3 === 0) {
                        msg += "<div class='row'>"
                    }
                    msg += "<div class=' portrait col-md-4 col-sm-6 col-xs-12'>";
                    msg = msg + "<span class='portrait-title'>" + payload[i].name + "</span>";
                    msg = msg + "<img src='" + payload[i].path + "' alt='" + payload[i].name + "' class='portrait-img'>";
                    msg += "</div>";
                    if (i % 3 === 2) {
                        msg += "</div>"
                    }
                }
                $("#portraits").empty();
                $("#portraits").append(msg);
            } else {
                alert("系统内部错误");
            }
        }
    });
}