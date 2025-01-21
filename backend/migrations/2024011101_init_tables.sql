-- 用户表
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 文章表
CREATE TABLE IF NOT EXISTS articles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    cover_image VARCHAR(255),
    category VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 视频表
CREATE TABLE IF NOT EXISTS videos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    cover_image VARCHAR(255),
    category VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 电子书表
CREATE TABLE IF NOT EXISTS ebooks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    file_url VARCHAR(255) NOT NULL,
    cover_image VARCHAR(255),
    category VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 评论表
CREATE TABLE IF NOT EXISTS comments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    content TEXT NOT NULL,
    article_id INT,
    video_id INT,
    ebook_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (video_id) REFERENCES videos(id),
    FOREIGN KEY (ebook_id) REFERENCES ebooks(id)
);

-- 插入测试数据
INSERT INTO users (username, email, password_hash) VALUES
('testuser1', 'test1@example.com', 'hashed_password_1'),
('testuser2', 'test2@example.com', 'hashed_password_2');

INSERT INTO articles (user_id, title, content, cover_image, category) VALUES
(1, 'First Article', 'This is the content of the first article', 'cover1.jpg', 'Technology'),
(2, 'Second Article', 'This is the content of the second article', 'cover2.jpg', 'Science');

INSERT INTO videos (user_id, title, url, cover_image, category) VALUES
(1, 'First Video', 'https://example.com/video1.mp4', 'video1.jpg', 'Education'),
(2, 'Second Video', 'https://example.com/video2.mp4', 'video2.jpg', 'Entertainment');

INSERT INTO ebooks (user_id, title, file_url, cover_image, category) VALUES
(1, 'First Ebook', 'https://example.com/ebook1.pdf', 'ebook1.jpg', 'Programming'),
(2, 'Second Ebook', 'https://example.com/ebook2.pdf', 'ebook2.jpg', 'Design');

INSERT INTO comments (user_id, content, article_id, video_id, ebook_id) VALUES
(1, 'Great article!', 1, NULL, NULL),
(2, 'Nice video!', NULL, 1, NULL),
(1, 'Helpful ebook!', NULL, NULL, 1);
