package models

type Category struct {
	ID   int
	Name string
}

func GetOrCreateCategory(name string) (category *Category, err error) {
	category = &Category{}
	tx := DB.FirstOrCreate(category, Category{Name: name})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}
