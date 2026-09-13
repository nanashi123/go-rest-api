CREATE TABLE IF NOT EXISTS articles (
    article_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    contents TEXT NOT NULL,
    username VARCHAR(100) NOT NULL,
    nice INTEGER NOT NULL,
    created_at DATETIME
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id INTEGER UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    article_id INTEGER UNSIGNED NOT NULL,
    message TEXT NOT NULL,
    created_at DATETIME,
    FOREIGN KEY (article_id) REFERENCES articles (article_id)
);

INSERT INTO articles (
    title,
    contents,
    username,
    nice,
    created_at
) VALUES (
    'firstPost',
    'This is my first blog',
    'saki',
    3,
    NOW()
);

INSERT INTO articles (
    title,
    contents,
    username,
    nice
) VALUES (
    '2nd',
    'Second blog post',
    'saki',
    4
);

INSERT INTO comments (
    article_id,
    message,
    created_at
) VALUES (
    1,
    '1stcommentyeah',
    NOW()
);

INSERT INTO comments (
    article_id,
    message
) VALUES (
    1,
    'welcome'
);