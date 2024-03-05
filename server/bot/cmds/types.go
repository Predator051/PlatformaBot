package cmds

type CmdName string

var (
	RequestForAdminsPrivilegesName CmdName = "request_for_admins_privileges"
	SubscribeToGroupNewsName       CmdName = "subscribe_to_channel"
	SendMsgToChannelName           CmdName = "send_msg_to_channel"
)
