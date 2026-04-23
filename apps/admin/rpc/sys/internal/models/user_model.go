package models

import "time"

//UserModel
/**
用户表
@Author :  liushuiyuan
@Date 2026/4/22 21:34
*/
type UserModel struct {
	Model
	UserName string `gorm:"size:64" json:"username"`
	NickName string `gorm:"size:64" json:"nickname"`
	Avatar   string `gorm:"size:512" json:"avatar"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

//UserConfModel  用户配置表
/**
一对一关系
用户配置表和用户表
@Author :  liushuiyuan
@Date 2026/4/22 21:34
*/
type UserConfModel struct {
	UserID             string     `gorm:"unique" json:"userID"`
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"` //兴趣标签 多个 json类型 	//第三方登录的openId
	UpdateUserNameDate *time.Time `json:"updateUserNameDate"`                            //上次修改用户名时间
	OpenCollect        bool       `json:"openCollect"`                                   //是否开启收藏功能
	OpenLike           bool       `json:"openLike"`                                      //是否开启点赞功能
	OpenFans           bool       `json:"openFans"`                                      //是否开启粉丝功能
	OpenFollow         bool       `json:"openFollow"`                                    //是否开启关注功能
	OpenComment        bool       `json:"openComment"`                                   //是否开启评论功能
	HomeStyleID        int        `json:"homeStyleID"`                                   //主页风格
}
