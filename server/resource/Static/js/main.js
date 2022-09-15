$(document).ready(function() {
   $.post("http://localhost:8888/base/captcha", function (result) {     //��Ҫ�����������
      upcaptcha(result)
   });
})

//����
$('.message a').click(function(){
   $('form').animate({height: "toggle", opacity: "toggle"}, "slow");
});

//��ȡ��֤����ȡ����
var upcaptcha = function (result){
   console.log(result)
   $(".captcha").attr("src", result.data.picPath )
   $(".captchaId").attr("value",result.data.captchaId)
}

//�����ȡ��֤��
$(".captcha").click(function (){
   $.post("http://localhost:8888/base/captcha",function(result){     //��Ҫ�����������
      upcaptcha(result)
   });
})

//�ύע��
$(".createBtn").click(function (){
   $(".register-form").ajaxForm(function (request){
      //У�顣����
      // alert(request.msg)
      console.log(request)
      $('form').animate({height: "toggle", opacity: "toggle"}, "slow");
   });
})


$(".loginBtn").click(function (){
   $(".login-form").ajaxForm(function (request){
      // У�顣������
      // alert(request.msg)
      console.log(request)
      var reqData = request.data
      let userData = reqData.user;
      if (request.code == 200){
         //����cookie     :����token�������ÿ����Ǻ���
         $.cookie('Authorization',reqData.token,{expires:7,path:'/GOGOGO'});
         $.cookie('ExpiresAt',reqData.expiresAt,{expires:7,path:'/GOGOGO'});
         //�����û���Ϣ
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
         //����������Ϣ
         var setting = {
            url: "/index",//����·��
            type: "GET",//����ʽ
            // dataType: "json",//��������ݸ�ʽ
            // timeout:4000,//��ʱ����
            headers:{'Authorization':reqData.token},//����ͷ
            success: function (data,status) {//����ɹ��Ļص�
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
         //������֤��
         $(".captcha").click()
      }
   })
})






