-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE  banner_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE banner (
    id int DEFAULT nextval('banner_seq'::regclass) NOT NULL,
    title text NOT NULL,
    text text NOT NULL,
    url text NOT NULL,
    active bool DEFAULT true,
    updated_at timestamp with time zone DEFAULT NOW() NOT NULL,
    created_at timestamp with time zone DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE chains (
    banner_id int,
    tags_id int NOT NULL,
    feature_id int NOT NULL,
    PRIMARY KEY (feature_id, tags_id),
    FOREIGN KEY(banner_id) REFERENCES banner(id)
);

CREATE INDEX banner_id_idx ON banner(id);
CREATE INDEX chains_feature_idx ON chains(feature_id);
CREATE INDEX chains_tags_id_idx ON chains(tags_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chains;
DROP TABLE banner;

DROP INDEX IF EXISTS banner_id_idx;
DROP INDEX IF EXISTS banner_feature_idx;
DROP INDEX IF EXISTS chains_tags_id_idx;

DROP SEQUENCE banner_seq;
-- +goose StatementEnd