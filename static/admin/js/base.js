layui.use(['form', 'layer', 'laydate', 'table', 'upload'], function () {
    var form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer,
        $ = layui.jquery,
        laydate = layui.laydate,
        upload = layui.upload,
        table = layui.table;


//列表操作
    table.on('tool(jsonTable)', function (obj) {
        var layEvent = obj.event,
            data = obj.data;
        if (layEvent === 'update') { //编辑
            add(data, "update");
        } else if (layEvent === 'delete') { //删除
            layer.confirm('确定删除此文章？', {icon: 3, title: '提示信息'}, function (index) {
                $.get("delete", {
                    id: data.id  //将需要删除的newsId作为参数传入
                }, function (data) {
                    setTimeout(function () {
                        layer.msg(data.msg)
                        tableIns.reload();
                        layer.close(index);
                    }, 500);
                })
            });
        } else if (layEvent === 'view') { //预览
            $.get("view", {
                id: data.id  //将需要删除的newsId作为参数传入
            }, function (data) {
                setTimeout(function () {
                    layer.msg(data.msg)
                    tableIns.reload();
                    layer.close(index);
                }, 500);
            })

        }
    });

    //编辑文章
    function add(edit, action) {
        var index = layui.layer.open({
            title: "小说",
            type: 2,
            maxmin: true,
            area: ['50%', '80%'],
            content: action,
            success: function (layero, index) {
                var body = layui.layer.getChildFrame('body', index);

                Edit(body, edit, form);

/*                setTimeout(function () {
                    layui.layer.tips('点击此处返回文章列表', '.layui-layer-setwin .layui-layer-close', {
                        tips: 3
                    });
                }, 500)*/
            }
        })
        //layui.layer.full(index);
        //改变窗口大小时，重置弹窗的宽高，防止超出可视区域（如F12调出debug的操作）
        $(window).on("resize", function () {
            layui.layer.full(index);
        })
    }

    //批量删除
    $(".delAll_btn").click(function () {
        var checkStatus = table.checkStatus('jsonTable'),
            data = checkStatus.data,
            Id = [];
        if (data.length > 0) {
            for (var i in data) {
                Id.push(data[i].id);
            }
            layer.confirm('确定删除选中的文章？', {icon: 3, title: '提示信息'}, function (index) {
                $.get("batch-delete", {
                    ids: Id  //将需要删除的newsId作为参数传入
                }, function (data) {
                    layer.msg(data.msg)
                    tableIns.reload();
                    layer.close(index);
                })
            })
        } else {
            layer.msg("请选择需要删除的文章");
        }
    })

    //添加
    $(".create_btn").click(function () {
        add("", "create");
    })
});


function formatUnixtimestamp(unixtimestamp) {
    var unixtimestamp = new Date(unixtimestamp * 1000);
    var year = 1900 + unixtimestamp.getYear();
    var month = "0" + (unixtimestamp.getMonth() + 1);
    var date = "0" + unixtimestamp.getDate();
    var hour = "0" + unixtimestamp.getHours();
    var minute = "0" + unixtimestamp.getMinutes();
    var second = "0" + unixtimestamp.getSeconds();
    return year + "-" + month.substring(month.length - 2, month.length) + "-" + date.substring(date.length - 2, date.length) + " " + hour.substring(hour.length - 2, hour.length) + ":" + minute.substring(minute.length - 2, minute.length) + ":" + second.substring(second.length - 2, second.length);
}