CREATE TABLE IF NOT EXISTS hosts
(
    id         VARCHAR(255) NOT NULL
        PRIMARY KEY,
    created_at TIMESTAMP(0) NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS checks_id_seq;
CREATE TABLE IF NOT EXISTS checks
(
    id            INTEGER DEFAULT nextval('checks_id_seq'::regclass) NOT NULL
        CONSTRAINT checks_pkey
            PRIMARY KEY,
    created_at    TIMESTAMP(0)                                       NOT NULL,
    host          VARCHAR(255)                                       NOT NULL
        CONSTRAINT checks_hosts_id_fk
            REFERENCES hosts,
    state         SMALLINT                                           NOT NULL,
    response_time FLOAT                                              NOT NULL
);
CREATE INDEX checks_host_index
    ON checks (host);

INSERT INTO hosts (id, created_at)
VALUES ('https://nonexisteddomain.com', now()),
       ('https://google.com', now()),
       ('https://youtube.com', now()),
       ('https://facebook.com', now()),
       ('https://baidu.com', now()),
       ('https://wikipedia.org', now()),
       ('https://yahoo.com', now()),
       ('https://tmall.com', now()),
       ('https://amazon.com', now()),
       ('https://twitter.com', now()),
       ('https://live.com', now()),
       ('https://instagram.com', now());
