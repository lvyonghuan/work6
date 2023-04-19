package dao

type Store struct {
	Name  string `gorm:"primaryKey"`
	Price string
	Num   int
}
