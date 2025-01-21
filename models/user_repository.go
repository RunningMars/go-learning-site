package models

import (
    "database/sql"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) CreateUser(user *User, passwordHash string) (*User, error) {
    query := `
        INSERT INTO users (username, phone, email, password_hash, age, gender)
        VALUES (?, ?, ?, ?, ?, ?)
    `
    result, err := r.DB.Exec(query, user.Username, user.Phone, user.Email, passwordHash, user.Age, user.Gender)
    if err != nil {
        return nil, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }
    user.ID = id

    return user, nil
}

func (r *UserRepository) GetUserByPhone(phone string) (*User, error) {
    var user User
    query := `SELECT * FROM users WHERE phone = ?`
    err := r.DB.QueryRow(query, phone).Scan(
        &user.ID, &user.Username, &user.Phone, &user.Email,
        &user.Age, &user.Gender, &user.Avatar, &user.CreatedAt, &user.UpdatedAt,
    )
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
    var count int
    query := `SELECT COUNT(*) FROM users WHERE username = ?`
    err := r.DB.QueryRow(query, username).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *UserRepository) CheckPhoneExists(phone string) (bool, error) {
    var count int
    query := `SELECT COUNT(*) FROM users WHERE phone = ?`
    err := r.DB.QueryRow(query, phone).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}