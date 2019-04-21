layui.use(['form', 'layer', 'laydate', 'table', 'laytpl'], function () {
    var form = layui.form,
        layer = parent.layer === undefined ? layui.layer : top.layer,
        $ = layui.jquery,
        laydate = layui.laydate,
        laytpl = layui.laytpl,
        table = layui.table;

    //新闻列表
    var tableIns = table.render({
        elem: '#jsonTable',
        url: '/books-type/index',
        cellMinWidth: 95,
        page: true,
        height: "full-125",
        limit: 10,
        limits: [10, 15, 20, 25],
        //id: "booksListTable",
        cols: [[
            {type: "checkbox", fixed: "left", width: 50},
            {field: 'id', title: 'ID', width: 60, align: "center"},
            {field: 'name', title: '规则名称'},
            {title: '操作', width: 170, templet: '#toolBar', fixed: "right", align: "center"}
        ]]
    });

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

    //搜索【此功能需要后台配合，所以暂时没有动态效果演示】
    $(".search_btn").on("click", function () {
        if ($(".searchVal").val() != '') {
            table.reload("jsonTable", {
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
    function add(edit, action) {
        var index = layui.layer.open({
            title: "小说",
            type: 2,
            content: action,
            success: function (layero, index) {
                var body = layui.layer.getChildFrame('body', index);
                if (edit) {
                    body.find(".layui-form").attr("action", "update?id=" + edit.id)
                    body.find(".name").val(edit.name);
                    form.render();
                } else {
                    body.find(".layui-form").attr("action", "create")
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

    //添加
    $(".create_btn").click(function () {
        add("", "create");
    })

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

})