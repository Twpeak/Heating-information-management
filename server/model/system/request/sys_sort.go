package request

//添加分类请求所需要结构
type AddSort struct {
	SortName	string		`json:"sortName" gorm:"comment:分类名称;column:sort_name"`
	SortComment string		`json:"sortComment" gorm:"comment:分类描述;column:sort_comment"`
}

//修改文章的请求体
type UpdateSortByArt struct {
	ArticleId     uint		`json:"articleId"`//当前文章id
	Id        uint			`json:"Id"`	//标签id或分类id
}
