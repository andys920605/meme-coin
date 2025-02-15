CREATE TABLE meme_coin (
    id BIGINT UNSIGNED NOT NULL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL UNIQUE,       
    description TEXT NOT NULL,               
    popularity_score INT NOT NULL, 
    created_at TIMESTAMP NOT NULL,          
    updated_at TIMESTAMP NOT NULL            
);
