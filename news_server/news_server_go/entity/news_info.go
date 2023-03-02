package entity

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
type NewsInfo struct {
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title         string `gorm:"type:varchar(200)" json:"title"`
	Img_url       string `gorm:"type:varchar(200)" json:"imgUrl"`
	Author        string `gorm:"type:varchar(500)" json:"author"`
	Link          string `gorm:"type:varchar(500)" json:"link"`
	Comment_count uint64 `gorm:"type:varchar(50)" json:"commentCount"`
}

func (NewsInfo) TableName() string {
	return "news_news"
}
