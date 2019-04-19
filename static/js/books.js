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
        url: '/books/index',
        cellMinWidth: 95,
        page: true,
        height: "full-125",
        limit: 10,
        limits: [20, 50,100],
        id: "booksListTable",
        cols: [[
            {type: "checkbox", fixed: "left", width: 50},
            {field: 'id', title: 'ID', width: 60, align: "center"},
            {field: 'book_name', title: '小说名称'},
            {field: 'book_type_text', title: '类型', align: 'center', width: 70},
            {field: 'book_author', title: '作者', align: 'center'},
            {field: 'book_new_chapter', title: '最新章节', align: 'center'},
            {
                field: 'bookLast_at_text', title: '最后更新', align: 'center', templet: function (data) {
                    return formatUnixtimestamp(data.book_last_at_text);
                }
            },
            {
                field: 'created_at_text', title: '创建时间', align: 'center', minWidth: 110, templet: function (data) {
                    return formatUnixtimestamp(data.created_at_text);
                }
            },
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

        if (layEvent === 'update') { //编辑
            addBooks(data, "update");
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
        } else if (layEvent === 'update') { //预览
            $.get("update", {
                id: data.Id  //将需要删除的newsId作为参数传入
            }, function (data) {
                setTimeout(function () {
                    layer.msg(data.msg)
                    tableIns.reload();
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
                console.log(edit);
                if (edit) {
                    body.find(".layui-form").attr("action", "/books/update?id=" + edit.id)
                    body.find(".book_name").val(edit.book_name);
                    body.find(".list_url").val(edit.list_url);
                    body.find(".book_author").val(edit.book_author);
                    body.find(".book_describe").val(edit.book_describe);
                    body.find(".book_status select").val(edit.book_status);
                    body.find("input[name='book_type'][value='" + edit.book_type + "']").prop("checked", "checked");
                    body.find("input[name='preg_id'][value='"+edit.preg_id+"']").prop("checked", true)
                    body.find("input[name='book_type'][value='"+edit.book_type+"']").prop("checked", true)
                    body.find(".book_status").val(edit.book_status)
                    if (edit.is_top) {
                        body.find("input[name='is_top']").prop("checked", "checked");
                    }
                    form.render();
                } else {
                    body.find(".layui-form").attr("action", "/books/create")
                    body.find("input[name='preg_id']").eq(0).attr("checked", true)
                    body.find("input[name='book_type']").eq(0).attr("checked", true)
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
        addBooks("", "create");
    })

    //批量删除
    $(".delAll_btn").click(function () {
        var checkStatus = table.checkStatus('booksListTable'),
            data = checkStatus.data,
            Id = [];
        if (data.length > 0) {
            for (var i in data) {
                Id.push(data[i].id);
            }
            layer.confirm('确定删除选中的文章？', {icon: 3, title: '提示信息'}, function (index) {
                $.get("batch-del", {
                    ids: id  //将需要删除的newsId作为参数传入
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