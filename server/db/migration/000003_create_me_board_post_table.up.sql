CREATE TABLE IF NOT EXISTS board_post(
    pid INT NOT NULL AUTO_INCREMENT,
    password VARCHAR(50),
    writer VARCHAR(50) NOT NULL,
    title VARCHAR(100) NOT NULL,
    content MEDIUMTEXT NOT NULL,
    is_notice boolean NOT NULL,
    time_to_read_minute INT,
    create_at DATETIME DEFAULT current_timestamp,

    PRIMARY KEY (pid)
);