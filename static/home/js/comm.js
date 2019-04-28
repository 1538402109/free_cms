
//搜索
function search_bak(){
document.writeln("<div class=\"search\">");
document.writeln("<form name=\"articlesearch\" action=\"/modules/article/waps.php\">");
document.writeln("<table cellpadding=\"0\" cellspacing=\"0\" style=\"width:100%;\"><tr>");
document.writeln("<td style=\"background-color:#fff; border:1px solid #CCC;\"><input id=\"s_key\" name=\"searchkey\" type=\"text\" class=\"key\" value=\"输入书名后搜索，宁可少字不要错字\" onFocus=\"this.value=''\" /></td>");
document.writeln("<td style=\"width:35px; background-color:#0080C0; background-image:url('/images/search.png'); background-repeat:no-repeat; background-position:center\"><input name=\"submit\" type=\"submit\" value=\"\" class=\"go\"></td></tr>");
document.writeln("</table>");
document.writeln("</form>");
document.writeln("</div>");
}
function search(){
document.writeln("<div class=\"search\">");
document.writeln("<form name=\"articlesearch\" action=\"http://zhannei.baidu.com/cse/search\">");
document.writeln("<input type=\"hidden\" name=\"s\" value=\"15629547599739040164\">");
document.writeln("<table cellpadding=\"0\" cellspacing=\"0\" style=\"width:100%;\"><tr>");
document.writeln("<td style=\"background-color:#fff; border:1px solid #CCC;\"><input id=\"s_key\" name=\"q\" type=\"text\" class=\"key\" value=\"输入书名后搜索，宁可少字不要错字\" onFocus=\"this.value=''\" /></td>");
document.writeln("<td style=\"width:35px; background-color:#0080C0; background-image:url('/images/search.png'); background-repeat:no-repeat; background-position:center\"><input name=\"submit\" type=\"submit\" value=\"\" class=\"go\"></td></tr>");
document.writeln("</table>");
document.writeln("</form>");
document.writeln("</div>");
}

function getCookie(name)
{
    var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
    if(arr=document.cookie.match(reg))
        return unescape(arr[2]);
    else
        return null;
}
function setCookieWithTime(name, value, exp_time) {
    var exp = new Date();
    exp.setTime(exp.getTime() + exp_time);
    document.cookie = name + "= " + escape(value) + ";expires= " + exp.toGMTString()+";path=/";
}

//加入书架
function putbookcase(bid)
{
    $.ajax({
        cache:false,
        url:'/addbookcase/'+bid+'.php',
        success:function(data){
            if('-1'==data){
                alert('先登录再收藏！');
            }else if('-2'==data){
                alert('已收藏好多书了！');
            }else{
                alert('加入书架成功！');
            }
        },
        error:function(){
            alert('加入书架失败！');
        }
    });
}

//加入书签
function putbookmark(bid, cid)
{
    $.ajax({
        cache:false,
        url:'/addbookcase/'+bid+'/'+cid+'.php',
        success:function(data){
            if('-1'==data){
                alert('先登录再收藏！');
            }else if('-2'==data){
                alert('已收藏好多书了！');
            }else{
                alert('加入书签成功！');
            }
        },
        error:function(){
            alert('加入书签失败！');
        }
    });
}

//推荐
function vote(bid)
{
    $.ajax({
        cache:false,
        url:'/recommend/'+bid+'.php',
        success:function(data){
            alert('推荐成功！');
        },
        error:function(){
            alert('推荐失败！');
        }
    });
}

//记录点击数
function recordedclick(bid)
{
    if(check_bid_by_cookie(bid)){
        return ;
    }
    $.ajax({
        cache:false,
        url:'/bookclick/'+bid+'.php'
    });
    set_bid_in_cookie(bid);
}

function check_bid_by_cookie(bid){
    var clickbids = getCookie('clickbids');
    if(null == clickbids){
        return false;
    }
    var arr_bid = clickbids.split(',');
    for (var i = arr_bid.length - 1; i >= 0; i--) {
        if( parseInt(bid) == parseInt(arr_bid[i])){
            return true;
        }
    }
    return false;
}

function set_bid_in_cookie(bid){
    var clickbids = getCookie('clickbids');
    if(null == clickbids){
        clickbids = bid;
    }else{
        clickbids = clickbids + "," +bid;
    }
    var now_date = new Date();
    var tonight_date = new Date();
    tonight_date.setHours(23);
    var now_time = now_date.getTime();
    var tonight_time = tonight_date.getTime();
    var gap_time = tonight_time - now_time;
    setCookieWithTime('clickbids', clickbids, gap_time);
}

var checkbg = "#A7A7A7";
//内容页用户设置
function nr_setbg(intype){
    var huyandiv = document.getElementById("huyandiv");
    var light = document.getElementById("lightdiv");
    if(intype == "huyan"){
        if(huyandiv.style.backgroundColor == ""){
            set("light","huyan");
            document.cookie="light=huyan;path=/"; 
        }
        else{
            set("light","no");
            document.cookie="light=no;path=/"; 
        }
    }
    if(intype == "light"){
        if(light.innerHTML == "关灯"){
            set("light","yes");
            document.cookie="light=yes;path=/"; 
        }
        else{
            set("light","no");
            document.cookie="light=no;path=/"; 
        }
    }
    if(intype == "big"){
        set("font","big");
        document.cookie="font=big;path=/"; 
    }
    if(intype == "middle"){
        set("font","middle");
        document.cookie="font=middle;path=/"; 
    }
    if(intype == "small"){
        set("font","small");
        document.cookie="font=small;path=/"; 
    }           
}

//内容页读取设置
function getset(){ 
    var strCookie=document.cookie; 
    var arrCookie=strCookie.split("; ");  
    var light;
    var font;

    for(var i=0;i<arrCookie.length;i++){ 
        var arr=arrCookie[i].split("="); 
        if("light"==arr[0]){ 
            light=arr[1]; 
            break; 
        } 
    } 
    for(var i=0;i<arrCookie.length;i++){ 
        var arr=arrCookie[i].split("="); 
        if("font"==arr[0]){ 
            font=arr[1]; 
            break; 
        } 
    } 
    
    //light
    if(light == "yes"){
        set("light","yes");
    }
    else if(light == "no"){
        set("light","no");
    }
    else if(light == "huyan"){
        set("light","huyan");
    }
    //font
    if(font == "big"){
        set("font","big");
    }
    else if(font == "middle"){
        set("font","middle");
    }
    else if(font == "small"){
        set("font","small");
    }
    else{
        set("",""); 
    }
}

//内容页应用设置
function set(intype,p){
    var nr_body = document.getElementById("nr_body");//页面body
    var huyandiv = document.getElementById("huyandiv");//护眼div
    var lightdiv = document.getElementById("lightdiv");//灯光div
    var fontfont = document.getElementById("fontfont");//字体div
    var fontbig = document.getElementById("fontbig");//大字体div
    var fontmiddle = document.getElementById("fontmiddle");//中字体div
    var fontsmall = document.getElementById("fontsmall");//小字体div
    var nr1 =  document.getElementById("nr1");//内容div
    var nr_title =  document.getElementById("nr_title");//文章标题
    var nr_title =  document.getElementById("nr_title");//文章标题
    
    var pt_prev =  document.getElementById("pt_prev");
    var pt_mulu =  document.getElementById("pt_mulu");
    var pt_next =  document.getElementById("pt_next");
    var pb_prev =  document.getElementById("pb_prev");
    var pb_mulu =  document.getElementById("pb_mulu");
    var pb_next =  document.getElementById("pb_next");
    
    //灯光
    if(intype == "light"){
        if(p == "yes"){ 
            //关灯
            lightdiv.innerHTML = "开灯";
            nr_body.style.backgroundColor = "#32373B";
            huyandiv.style.backgroundColor = "";
            if(null != nr_title)
            nr_title.style.color = "#ccc";
            nr1.style.color = "#999";
            var pagebutton = "background-color:#3e4245;color:#ccc;border:1px solid #313538";            
            pt_prev.style.cssText = pagebutton;
            pt_mulu.style.cssText = pagebutton;
            pt_next.style.cssText = pagebutton
            if(null != pb_prev)
            pb_prev.style.cssText = pagebutton;
            pb_mulu.style.cssText = pagebutton;
            if(null != pb_next)
            pb_next.style.cssText = pagebutton;
            if(null != pt_shouye)
            pt_shouye.style.cssText = pagebutton;
            if(null != pb_shouye)
            pb_shouye.style.cssText = pagebutton;
        }
        else if(p == "no"){
            //开灯
            lightdiv.innerHTML = "关灯";
            nr_body.style.backgroundColor = "#fbf6ec";
            nr1.style.color = "#000";
            if(null != nr_title)
            nr_title.style.color = "#000";
            huyandiv.style.backgroundColor = "";
            var pagebutton = "background-color:#f4f0e9;color:green;border:1px solid #ece6da";           
            pt_prev.style.cssText = pagebutton;
            pt_mulu.style.cssText = pagebutton;
            pt_next.style.cssText = pagebutton
            if(null != pb_prev)
            pb_prev.style.cssText = pagebutton;
            pb_mulu.style.cssText = pagebutton;
            if(null != pb_next)
            pb_next.style.cssText = pagebutton;
            if(null != pt_shouye)
            pt_shouye.style.cssText = pagebutton;
            if(null != pb_shouye)
            pb_shouye.style.cssText = pagebutton;
        }
        else if(p == "huyan"){
            //护眼
            lightdiv.innerHTML = "关灯";
            huyandiv.style.backgroundColor = checkbg;
            nr_body.style.backgroundColor = "#DCECD2";
            nr1.style.color = "#000";
            var pagebutton = "background-color:#CCE2BF;color:green;border:1px solid #bbd6aa";           
            pt_prev.style.cssText = pagebutton;
            pt_mulu.style.cssText = pagebutton;
            pt_next.style.cssText = pagebutton
            if(null != pb_prev)
            pb_prev.style.cssText = pagebutton;
            pb_mulu.style.cssText = pagebutton;
            if(null != pb_next)
            pb_next.style.cssText = pagebutton;
            if(null != pt_shouye)
            pt_shouye.style.cssText = pagebutton;
            if(null != pb_shouye)
            pb_shouye.style.cssText = pagebutton;
        }
    }
    //字体
    if(intype == "font"){
        //alert(p);
        fontbig.style.backgroundColor = "";
        fontmiddle.style.backgroundColor = "";
        fontsmall.style.backgroundColor = "";
        if(p == "big"){
            fontbig.style.backgroundColor = checkbg;
            nr1.style.fontSize = "24px";
        }
        if(p == "middle"){
            fontmiddle.style.backgroundColor = checkbg;
            nr1.style.fontSize = "20px";
        }
        if(p == "small"){
            fontsmall.style.backgroundColor = checkbg;
            nr1.style.fontSize = "16px";
        }
    }
}

function getQuery()
{
    var url = window.document.location.href.toString();
    var u = url.split("/");
    return u;
}