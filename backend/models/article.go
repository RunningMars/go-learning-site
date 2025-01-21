package models

import (
	"database/sql"
)

type ArticleRepository struct {
	DB *sql.DB
}

func (r *ArticleRepository) GetArticles(page, limit int, search string) ([]Article, error) {
	offset := (page - 1) * limit
	query := `
		SELECT id, user_id, title, content, cover_image, category, created_at, updated_at 
		FROM articles
		WHERE (? = '' OR title LIKE ? OR content LIKE ?)
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`
	
	rows, err := r.DB.Query(query, "%"+search+"%", "%"+search+"%", "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(
			&article.ID,
			&article.UserID,
			&article.Title,
			&article.Content,
			&article.CoverImage,
			&article.Category,
			&article.CreatedAt,
			&article.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *ArticleRepository) GetArticle(id int) (*Article, error) {
	query := `
		SELECT id, user_id, title, content, cover_image, category, created_at, updated_at 
		FROM articles
		WHERE id = ?
	`
	
	var article Article
	err := r.DB.QueryRow(query, id).Scan(
		&article.ID,
		&article.UserID,
		&article.Title,
		&article.Content,
		&article.CoverImage,
		&article.Category,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &article, nil
}

func (r *ArticleRepository) GetArticleComments(articleID int) ([]Comment, error) {
	query := `
		SELECT id, user_id, content, article_id, video_id, ebook_id, created_at, updated_at
		FROM comments
		WHERE article_id = ?
		ORDER BY created_at DESC
	`
	
	rows, err := r.DB.Query(query, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.Content,
			&comment.ArticleID,
			&comment.VideoID,
			&comment.EbookID,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *ArticleRepository) SearchArticles(keyword string) ([]Article, error) {
	// 复用GetArticles方法，page=1, limit=100
	return r.GetArticles(1, 100, keyword)
}
