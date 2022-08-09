CREATE TABLE IF NOT EXISTS board_comment (
    cid INT NOT NULL AUTO_INCREMENT,
    pid INT NOT NULL,
    writer VARCHAR(30) NOT NULL,
    password VARCHAR(30),
    comment TEXT,
    like_cnt INT,
    create_at DATETIME DEFAULT current_timestamp,

    PRIMARY KEY (cid),
    FOREIGN KEY (pid) REFERENCES board_post(pid) on UPDATE CASCADE on DELETE RESTRICT
);