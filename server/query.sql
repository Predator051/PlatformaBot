-- name: ListGroupList :many
select * from group_lists order by name;

-- name: InsertNewGroupList :exec
insert into group_lists (id, name) VALUES (default, $1);
