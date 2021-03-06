package model

type Role struct {
	Rolename string `gorm:"size:32;primary_key;"       json:"role_name"     form:"role_name"`
}

type User struct {
	ID          uint   `gorm:"primary_key"                json:"id"           form:"id"`             //编号
	Username    string `gorm:"size:64;unique;not null;"   json:"username"     form:"username"`       //用户名
	Email       string `gorm:"size:128;unique;not null;"  json:"email"        form:"email"`          //邮箱地址
	Password    string `gorm:"size:512;not null;"         json:"password"     form:"password"`       //密码
	Nickname    string `gorm:"size:16;"                   json:"nickname"     form:"nickname"`       //昵称
	Description string `gorm:"type:text"                  json:"description"  form:"description"`    //描述
	Status      string `gorm:"size:32;not null;"          json:"status"       form:"status"`         //状态
	Role        string `gorm:"size:32;not null;"          json:"role"         form:"role"`           //角色
	CreatedTime int64  `gorm:"not null"                   json:"created_time"   form:"created_time"` //创建于
}

type Token struct {
	Token       string `gorm:"size:64;primary_key;"       json:"token"       form:"token"`          //Token串
	UserID      uint   `gorm:"size:64;"                   json:"-"           form:"-"`              //用户ID
	Status      string `gorm:"size:32"                    json:"status"      form:"status"`         //状态
	CreatedTime int64  `gorm:"not null"                   json:"created_time"  form:"created_time"` //创建于
	ExpiredTime int64  `gorm:"not null"                   json:"expired_time"  form:"expired_time"` //过期于
}

type Config struct {
	ID      uint   `gorm:"primary_key"     json:"id"       form:"id"`      //编号
	Name    string `gorm:"size:32;"        json:"name"     form:"name"`    //名称
	Value   string `gorm:"size:64;"        json:"value"    form:"value"`   //值
	Section string `gorm:"size:64;"        json:"section"  form:"section"` //段
}

type RuningInfo struct {
	ID      uint   `gorm:"primary_key"     json:"id"       form:"id"`      //编号
	Name    string `gorm:"size:32;"        json:"name"     form:"name"`    //名称
	Value   string `gorm:"size:64;"        json:"value"    form:"value"`   //值
	Section string `gorm:"size:64;"        json:"section"  form:"section"` //段
}

type Article struct {
	ID          uint   `gorm:"primary_key"        json:"id"            form:"id"`           //编号
	Title       string `gorm:"size:256"           json:"title"         form:"title"`        //标题
	Content     string `gorm:"type:text"          json:"content"       form:"content"`      //内容
	Author      string `gorm:"size:64;"           json:"author"        form:"author"`       //作者
	AuthorID    uint   `                          json:"author_id"     form:"author_id"`    //作者
	Status      string `gorm:"size:32"            json:"status"        form:"status"`       //状态
	Type        string `gorm:"size:32"            json:"type"          form:"type"`         //文章分类
	CreatedTime int64  `gorm:"not null"           json:"created_time"  form:"created_time"` //创建于
	UpdatedTime int64  `gorm:"not null"           json:"updated_time"  form:"updated_time"` //编辑于
}

type Comment struct {
	ID          uint   `gorm:"primary_key"         json:"id"            form:"id"`           //编号
	Author      string `gorm:"size:64;"            json:"author"        form:"author"`       //作者
	ArticleID   uint   `gorm:"not null"            json:"article_id"    form:"article_id"`   //评论的文章编号
	Email       string `gorm:"size:128;not null;"  json:"email"         form:"email"`        //评论者邮箱
	Site        string `gorm:"size:128;not null;"  json:"email"         form:"email"`        //评论者网址
	Status      string `gorm:"size:32"             json:"status"        form:"status"`       //评论状态
	Content     string `gorm:"type:text"           json:"content"       form:"content"`      //评论内容
	CreatedTime int64  `gorm:"not null"            json:"created_time"  form:"created_time"` //评论时间
	UpdatedTime int64  `gorm:"not null"            json:"updated_time"  form:"updated_time"` //编辑于
	ParentID    uint   `gorm:"not null"            json:"parent_id"     form:"parent_id"`    //父评论编号
}
