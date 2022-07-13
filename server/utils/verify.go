package utils

//规则表，主要用来校验数据是否符合其中指定的规则，配合校验器一起使用
//一个字段可以指定多个规则：如同数据库的 非空，唯一，指定长度一样
//现在约定的规则仅有：1.NotEmpty(非空)，2.前缀：regexp= 后面是正则表达式。3.前缀是lt等判断类型。来限制字符长度
//我们规定规则一有NotEmpty方法(返回NotEmpty字段)，规则二有：RegexpMatch方法，方便填入前缀，也可以再写一个方法去验证其正则表达式合法性。规则三有lt等各类方法，

var(
	LoginVerify            = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()},"Captcha": {NotEmpty()},"CaptchaId": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "Name": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	//AddArticleVerify	   = Rules{"UserId":{NotEmpty()},"Title":{NotEmpty()},"SortId":{NotEmpty()}}
	AddArticleVerify	   = Rules{"UserId":{NotEmpty()},"Title":{NotEmpty()},"SortId":{NotEmpty()},"Labels":{NotEmpty()}}
	LabelVerify			   = Rules{"LabelName":{NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty(),Gt("0")}, "PageSize": {NotEmpty(),Gt("0")}}
	UpdateArticleVerify    = Rules{"Id":{NotEmpty()},"UserId":{NotEmpty()},"Title":{NotEmpty()},"SortId":{NotEmpty()},"Labels":{NotEmpty()}}
	LikeVerify		   	   = Rules{"Id":{NotEmpty()},"UserId":{NotEmpty()}}
	AddSortVerify		   = Rules{"SortName":{NotEmpty()}}
	UpdateSortVerify	   = Rules{"ArticleId":{NotEmpty()},"Id":{NotEmpty()}}
)

