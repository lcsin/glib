package model

type ArticleWriter struct {
	ID       int64  `gorm:"comment:文章id;primaryKey;autoIncrement"`
	AuthorID int64  `gorm:"comment:作者id;index"`
	Title    string `gorm:"comment:标题;size:256"`
	Content  string `gorm:"comment:内容;type:text"`
	Status   int8   `gorm:"comment:状态;default:1"`

	CreatedTime int64 `gorm:"comment:创建时间"`
	UpdatedTime int64 `gorm:"comment:更新时间"`
	PublishTime int64 `gorm:"comment:发布时间"`
}

func (a ArticleWriter) TableName() string {
	return "article_writer_tbl"
}

type ArticleReader struct {
	ID       int64  `gorm:"comment:文章id;primaryKey;autoIncrement"`
	AuthorID int64  `gorm:"comment:作者id;index"`
	Title    string `gorm:"comment:标题;size:256"`
	Content  string `gorm:"comment:内容;type:text"`
	Status   int8   `gorm:"comment:状态;default:1"`

	CreatedTime int64 `gorm:"comment:创建时间"`
	UpdatedTime int64 `gorm:"comment:更新时间"`
	PublishTime int64 `gorm:"comment:发布时间"`
}

func (a ArticleReader) TableName() string {
	return "article_reader_tbl"
}
