-- name: ListGroupList :many
select * from group_lists order by name;

-- name: GroupListById :one
select * from group_lists where id = $1;

-- name: InsertNewGroupList :exec
insert into group_lists (id, name) VALUES (default, $1);

-- name: DeleteGroupList :exec
delete from group_lists where id = $1;

-- name: ListAdminsGroupListRequest :many
select * from admins_of_group_list_request order by id;

-- name: InsertListAdminsGroupListRequest :exec
insert into admins_of_group_list_request (id, chat_id, group_list_id, first_name, second_name, username) VALUES (default, $1, $2, $3, $4, $5);

-- name: DeleteListAdminsGroupListRequest :exec
delete from admins_of_group_list_request where id = $1;
