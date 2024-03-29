CREATE TABLE IF NOT EXISTS board_post_tag(
    tid INT NOT NULL,
    pid INT NOT NULL,

    PRIMARY KEY (tid, pid),
    FOREIGN KEY (tid) REFERENCES board_tag(tid) on UPDATE CASCADE on DELETE CASCADE,
    FOREIGN KEY (pid) REFERENCES board_post(pid) on UPDATE CASCADE on DELETE CASCADE
);