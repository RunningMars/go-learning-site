package fixtures

import (
	"database/sql"
)

func CreateTestArticles(db *sql.DB) error {
	_, err := db.Exec(`
        INSERT INTO articles (user_id, title, content) VALUES 
        (1, '测试文章1', '测试内容1'),
        (1, '测试文章2', '测试内容2')
    `)
	return err
}
