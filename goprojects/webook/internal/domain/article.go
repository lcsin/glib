package domain

type Article struct {
	ID      int64         `json:"id"`
	Title   string        `json:"title"`
	Content string        `json:"content"`
	Status  ArticleStatus `json:"status"`

	CreatedTime int64 `json:"created_time"`
	UpdatedTime int64 `json:"updated_time"`
	PublishTime int64 `json:"publish_time"`

	Author Author `json:"author"`
}

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

const (
	ArticleStatusUnknown ArticleStatus = iota
	ArticleStatusUnpublished
	ArticleStatusPublished
	ArticleStatusPrivate
)

type ArticleStatus uint8

func (as ArticleStatus) Published() bool {
	return as == ArticleStatusPublished
}

func (as ArticleStatus) ToInt8() int8 {
	return int8(as)
}

func (as ArticleStatus) String() string {
	switch as {
	case ArticleStatusPublished:
		return "published"
	case ArticleStatusUnpublished:
		return "unpublished"
	case ArticleStatusPrivate:
		return "private"
	default:
		return "unknown"
	}
}
