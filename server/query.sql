-- name: ListChannels :many
select * from channels order by name;

-- name: ChannelById :one
select * from channels where id = $1;

-- name: InsertNewChannel :exec
insert into channels (id, name) VALUES (default, $1);

-- name: DeleteChannel :exec
delete from channels where id = $1;

----------REQUEST----------

-- name: ListAdminsChannelRequest :many
select * from admins_of_channels_request order by id;

-- name: InsertListAdminsChannelRequest :exec
insert into admins_of_channels_request (id, chat_id, channels_id, first_name, second_name, username) VALUES (default, $1, $2, $3, $4, $5);

-- name: DeleteListAdminsChannelRequest :exec
delete from admins_of_channels_request where id = $1;

-- name: DeleteListAdminsChannelRequestByChannelAndChatId :exec
delete from admins_of_channels_request where channels_id = $1 and chat_id = $2;

-- name: DeleteListAdminsChannelRequestByChannelId :exec
delete from admins_of_channels_request where channels_id = $1;

----------GROUP LIST ADMINS----------

-- name: InsertChannelAdmins :exec
insert into channels_admins (id, chat_id, channels_id) VALUES (default, $1, $2);

-- name: DeleteChannelAdmins :exec
delete from channels_admins where id = $1;

-- name: DeleteChannelAdminsByChannelId :exec
delete from channels_admins where channels_id = $1;

-- name: ChannelsByAdmin :many
select * from channels_admins where chat_id = $1;

----------SUBSCRIPTIONS TO GROUP LIST----------

-- name: InsertSubscriptionToChannel :exec
insert into subscription_to_channels (id, chat_id, channels_id, username, title, chat_type) VALUES (default, $1, $2, $3, $4, $5);

-- name: SubscriptionToChannelsByChatId :many
select * from subscription_to_channels where chat_id = $1;

-- name: SubscriptionToChannelsByChannelId :many
select * from subscription_to_channels where channels_id = $1;

-- name: DeleteSubscriptionToChannel :exec
delete from subscription_to_channels where id = $1;

-- name: DeleteSubscriptionToChannelByChannelId :exec
delete from subscription_to_channels where channels_id = $1;

-- name: SubscriptionToChannels :many
select * from subscription_to_channels;
