CREATE TABLE group_lists
(
    id   BIGSERIAL PRIMARY KEY,
    name text NOT NULL
);


CREATE TABLE admins_of_group_list_request (
                             id   BIGSERIAL PRIMARY KEY,
                             chat_id bigserial,
                             group_list_id bigserial ,
                             username text,
                             first_name text,
                             second_name text,
                             CONSTRAINT fk_customer
                                 FOREIGN KEY(group_list_id)
                                     REFERENCES group_lists(id)
);
