layui.use(['form', 'layer', 'laydate', 'table', 'laytpl'], function () {
    var form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer,
        $ = layui.jquery,
        laydate = layui.laydate,
        laytpl = layui.laytpl,
        table = layui.table;

    //新闻列表
    var tableIns = table.render({
        elem: '#booksList',
        url: '/preg/index',
        cellMinWidth: 95,
        page: true,
        height: "full-125",
        limit: 10,
        limits: [10, 15, 20, 25],
        id: "booksListTable",
        cols: [[
            {type: "checkbox", fixed: "left", width: 50},
            {field: 'id', title: 'ID', width: 60, align: "center"},
            {field: 'name', title: '规则名称'},
            {field: 'list_msg_block', title: '列表描述块', align: 'center'},
            {field: 'list_describe', title: '描述', align: 'center'},
            {field: 'list_a_block', title: '列表链接', align: 'center'},
            {field: 'content_block', title: '内容块', align: 'center'},
            {field: 'content_title', title: '内容标题', align: 'center'},
            {field: 'content_text', title: '内容文本', align: 'center'},
            {title: '操作', width: 170, templet: '#booksListBar', fixed: "right", align: "center"}
        ]]
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

    //列表操作
    table.on('tool(booksList)', function (obj) {
        var layEvent = obj.event,
            data = obj.data;

        if (layEvent === 'edit') { //编辑
            addBooks(data, "edit");
        } else if (layEvent === 'del') { //删除
            layer.confirm('确定删除此文章？', {icon: 3, title: '提示信息'}, function (index) {
                $.get("del", {
                    id: data.Id  //将需要删除的newsId作为参数传入
                }, function (data) {
                    setTimeout(function () {
                        layer.msg(data.msg)
                        tableIns.reload();
                        layer.close(index);
                    }, 500);
                })
            });
        } else if (layEvent === 'update') { //预览
            $.get("update", {
                id: data.Id  //将需要删除的newsId作为参数传入
            }, function (data) {
                setTimeout(function () {
                    layer.msg(data.msg)
                    tableIns.reload();
                    layer.close(index);
                }, 500);
            })

        }
    });

    //搜索【此功能需要后台配合，所以暂时没有动态效果演示】
    $(".search_btn").on("click", function () {
        if ($(".searchVal").val() != '') {
            table.reload("booksListTable", {
                page: {
                    curr: 1 //重新从第 1 页开始
                },
                where: {
                    key: $(".searchVal").val()  //搜索的关键字
                }
            })
        } else {
            layer.msg("请输入搜索的内容");
        }
    });

    //编辑文章
    function addBooks(edit, action) {
        var index = layui.layer.open({
            title: "小说",
            type: 2,
            content: action,
            success: function (layero, index) {
                var body = layui.layer.getChildFrame('body', index);
                if (edit) {
                    body.find(".layui-form").attr("action", "edit?id=" + edit.Id)
                    body.find(".name").val(edit.name);
                    body.find(".list_a_block").val(edit.list_a_block);
                    body.find(".content_block").val(edit.content_block);
                    body.find(".content_title").val(edit.content_title);
                    body.find(".content_text").val(edit.content_text);
                    form.render();
                } else {
                    body.find(".layui-form").attr("action", "add")
                }
                setTimeout(function () {
                    layui.layer.tips('点击此处返回文章列表', '.layui-layer-setwin .layui-layer-close', {
                        tips: 3
                    });
                }, 500)
            }
        })
        layui.layer.full(index);
        //改变窗口大小时，重置弹窗的宽高，防止超出可视区域（如F12调出debug的操作）
        $(window).on("resize", function () {
            layui.layer.full(index);
        })
    }

    $(".addBooks_btn").click(function () {
        addBooks("", "add");
    })

    //批量删除
    $(".delAll_btn").click(function () {
        var checkStatus = table.checkStatus('booksListTable'),
            data = checkStatus.data,
            Id = [];
        if (data.length > 0) {
            for (var i in data) {
                Id.push(data[i].Id);
            }
            layer.confirm('确定删除选中的文章？', {icon: 3, title: '提示信息'}, function (index) {
                $.get("batch-del", {
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

})