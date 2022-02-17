package default_db

type Blog struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (Blog) TableName() string {
	return "blog"
}
