-- name: ListGroupList :many
select * from group_lists order by name;

-- name: GroupListById :one
select * from group_lists where id = $1;

-- name: InsertNewGroupList :exec
insert into group_lists (id, name) VALUES (default, $1);

-- name: DeleteGroupList :exec
delete from group_lists where id = $1;

----------REQUEST----------

-- name: ListAdminsGroupListRequest :many
select * from admins_of_group_list_request order by id;

-- name: InsertListAdminsGroupListRequest :exec
insert into admins_of_group_list_request (id, chat_id, group_list_id, first_name, second_name, username) VALUES (default, $1, $2, $3, $4, $5);

-- name: DeleteListAdminsGroupListRequest :exec
delete from admins_of_group_list_request where id = $1;

-- name: DeleteListAdminsGroupListRequestByGroupAndChatId :exec
delete from admins_of_group_list_request where group_list_id = $1 and chat_id = $2;

----------GROUP LIST ADMINS----------

-- name: InsertGroupListAdmins :exec
insert into group_lists_admins (id, chat_id, group_list_id) VALUES (default, $1, $2);

-- name: DeleteGroupListAdmins :exec
delete from group_lists_admins where id = $1;

-- name: GroupListsByAdmin :many
select * from group_lists_admins where chat_id = $1;

----------SUBSCRIPTIONS TO GROUP LIST----------

-- name: InsertSubscriptionToGroupList :exec
insert into subscription_to_group_lists (id, chat_id, group_list_id, username, title, chat_type) VALUES (default, $1, $2, $3, $4, $5);

-- name: SubscriptionToGroupListsByChatId :many
select * from subscription_to_group_lists where chat_id = $1;

-- name: SubscriptionToGroupListsByGroupListId :many
select * from subscription_to_group_lists where group_list_id = $1;

-- name: DeleteSubscriptionToGroupList :exec
delete from subscription_to_group_lists where id = $1;

-- name: SubscriptionToGroupLists :many
select * from subscription_to_group_lists;
