package domain

type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int8   `json:"status"`

	CreatedTime int64 `json:"created_time"`
	UpdatedTime int64 `json:"updated_time"`
	PublishTime int64 `json:"publish_time"`

	Author Author `json:"author"`
}

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
