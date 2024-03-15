CREATE TABLE channels
(
    id   BIGSERIAL PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE admins_of_channels_request (
                             id   BIGSERIAL PRIMARY KEY,
                             chat_id bigint,
                             channels_id integer ,
                             username text,
                             first_name text,
                             second_name text,
                             CONSTRAINT fk_customer
                                 FOREIGN KEY(channels_id)
                                     REFERENCES channels(id)
);

CREATE TABLE channels_admins
(
    id   BIGSERIAL PRIMARY KEY,
    chat_Id bigint,
    channels_id integer,
    CONSTRAINT fk_channels
        FOREIGN KEY(channels_id)
        REFERENCES channels(id)
);

CREATE TABLE subscription_to_channels
(
    id   BIGSERIAL PRIMARY KEY,
    chat_Id bigint,
    channels_id integer,
    username text,
    title text,
    chat_type text,
    CONSTRAINT fk_channels
        FOREIGN KEY(channels_id)
            REFERENCES channels(id)
);
