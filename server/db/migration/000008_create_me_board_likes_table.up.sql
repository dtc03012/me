CREATE TABLE IF NOT EXISTS board_likes (
    pid INT NOT NULL,
    uuid VARCHAR(50) NOT NULL,

    PRIMARY KEY (pid, uuid),
    FOREIGN KEY (pid) REFERENCES board_post(pid) on UPDATE CASCADE on DELETE CASCADE
);