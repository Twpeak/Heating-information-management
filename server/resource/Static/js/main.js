$(document).ready(function() {
   $.post("http://localhost:8888/base/captcha", function (result) {     //需要解决跨域问题
      upcaptcha(result)
   });
})

//动画
$('.message a').click(function(){
   $('form').animate({height: "toggle", opacity: "toggle"}, "slow");
});

//获取验证码提取方法
var upcaptcha = function (result){
   console.log(result)
   $(".captcha").attr("src", result.data.picPath )
   $(".captchaId").attr("value",result.data.captchaId)
}

//点击获取验证码
$(".captcha").click(function (){
   $.post("http://localhost:8888/base/captcha",function(result){     //需要解决跨域问题
      upcaptcha(result)
   });
})

//提交注册
$(".createBtn").click(function (){
   $(".register-form").ajaxForm(function (request){
      //校验。。。
      // alert(request.msg)
      console.log(request)
      $('form').animate({height: "toggle", opacity: "toggle"}, "slow");
   });
})


$(".loginBtn").click(function (){
   $(".login-form").ajaxForm(function (request){
      // 校验。。。。
      // alert(request.msg)
      console.log(request)
      var reqData = request.data
      let userData = reqData.user;
      if (request.code == 200){
         //设置cookie     :这里token过期设置可能是毫秒
         $.cookie('Authorization',reqData.token,{expires:7,path:'/GOGOGO'});
         $.cookie('ExpiresAt',reqData.expiresAt,{expires:7,path:'/GOGOGO'});
         //保存用户信息
         let User = {
            CreatedAt:userData.CreatedAt,
            ID: userData.ID,
            Roel: userData.Roel,
            UpdatedAt:userData.UpdatedAt,
            articles: userData.articles,
            authorityId: userData.authorityId,
            headerImg: userData.headerImg,
            nickName:userData.nickName,
            userName: userData.userName,
            uuid: userData.uuid,
         }
         $.cookie("user", JSON.stringify(User),{expires:7,path:'/GOGOGO'});
         //发送请求信息
         var setting = {
            url: "/index",//请求路径
            type: "GET",//请求方式
            // dataType: "json",//请求的数据格式
            // timeout:4000,//超时设置
            headers:{'Authorization':reqData.token},//请求头
            success: function (data,status) {//请求成功的回调
               console.dir(data)
               alert("GOOD!!!")
               window.location.href = "/index";
               // self.location = "/index";
               // top.location = "/index";
            },
            error: function (data,status) {
               console.log(data,status)
               alert("Failed!!!")
            }
         }
         $.ajax(setting)
      }else {
         //更新验证码
         $(".captcha").click()
      }
   })
})






