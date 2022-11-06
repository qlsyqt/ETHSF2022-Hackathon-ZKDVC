CREATE TABLE IF NOT EXISTS t_did (
    did VARCHAR(255) PRIMARY KEY,
    auth_private_key VARCHAR(255) NOT NULL,
    clt_id int(15) NOT NULL,
    rot_id int(15) NOT NULL,
    tor_id int(15) NOT NULL,
    created_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)