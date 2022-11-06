package common

// Visit https://github.com/iden3/go-merkletree-sql/blob/master/db/sql/schema.sql
const CREATE_MTROOTS_TABLE = "CREATE TABLE IF NOT EXISTS mt_roots (\n    mt_id BIGINT PRIMARY KEY,\n    key BYTEA,\n    created_at BIGINT,\n    deleted_at BIGINT\n);"
const CREATE_MTNODES_TABLE = "CREATE TABLE IF NOT EXISTS mt_nodes (\n    mt_id BIGINT,\n    key BYTEA,\n    type SMALLINT NOT NULL,\n    child_l BYTEA,\n    child_r BYTEA,\n    entry BYTEA,\n    created_at BIGINT,\n    deleted_at BIGINT,\n    PRIMARY KEY(mt_id, key)\n);"

const CLAIM_TYPE_AUTH = "auth"
const CLAIM_TYPE_DCP = "dcp"
