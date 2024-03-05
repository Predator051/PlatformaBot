CREATE TABLE group_lists
(
    id   BIGSERIAL PRIMARY KEY,
    name text NOT NULL
);


CREATE TABLE admins_of_group_list_request (
                             id   BIGSERIAL PRIMARY KEY,
                             chat_id bigint,
                             group_list_id integer ,
                             username text,
                             first_name text,
                             second_name text,
                             CONSTRAINT fk_customer
                                 FOREIGN KEY(group_list_id)
                                     REFERENCES group_lists(id)
);

CREATE TABLE group_lists_admins
(
    id   BIGSERIAL PRIMARY KEY,
    chat_Id bigint,
    group_list_id integer,
    CONSTRAINT fk_group_lists
        FOREIGN KEY(group_list_id)
        REFERENCES group_lists(id)
);

CREATE TABLE subscription_to_group_lists
(
    id   BIGSERIAL PRIMARY KEY,
    chat_Id bigint,
    group_list_id integer,
    username text,
    title text,
    chat_type text,
    CONSTRAINT fk_group_lists
        FOREIGN KEY(group_list_id)
            REFERENCES group_lists(id)
);
