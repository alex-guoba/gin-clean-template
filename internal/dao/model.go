package dao

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	IsDel      uint8  `json:"is_del"`
}

// Model: Tag
type TagModel struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t TagModel) TableName() string {
	return "blog_tag"
}

// Model: Artile
type ArticleModel struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a ArticleModel) TableName() string {
	return "blog_article"
}

// Model: Artile-Tag
type ArticleTagModel struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTagModel) TableName() string {
	return "blog_article_tag"
}
