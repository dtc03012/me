CREATE TABLE IF NOT EXISTS board_like (
    pid INT NOT NULL,
    uuid VARCHAR(50) NOT NULL,

    PRIMARY KEY (pid, uuid),
    FOREIGN KEY (pid) REFERENCES board_post(pid) on UPDATE CASCADE on DELETE RESTRICT
);