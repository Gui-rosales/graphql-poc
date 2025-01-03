package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db *sql.DB
	ID string
	Name string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO category (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{
		ID: id,
		Name: name,
		Description: description,
	}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM category")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT category.id, category.name, category.description FROM category JOIN course ON category.id = course.category_id WHERE course.id = $1", courseId).Scan(&id, &name, &description)
	if err != nil {
		return Category{}, err
	}

	return Category{
		ID: id,
		Name: name,
		Description: description,
	}, nil
}
		