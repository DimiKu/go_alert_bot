package db_actions

const (
	// init actions
	createUsersTable = `CREATE TABLE users (user_id integer PRIMARY KEY, chat_id integer)`

	createChannelsTable = `CREATE TABLE channels 
						   (chat_uuid uuid PRIMARY KEY, 
						   user_id INTEGER, channel_type VARCHAR, 
						   channel_link BIGINT)`

	createTelegramChatsTable = `CREATE TABLE telegram_chats 
							   (channel_link bigint, 
							   chat_uuid uuid PRIMARY KEY, 
							   user_id integer, 
							   telegram_chat_id bigint[], 
							   format_string varchar)`

	createStdoutChatsTable = `CREATE TABLE stdout_chats 
							  (chat_uuid uuid PRIMARY KEY, 
							  user_id integer, 
							  format_string varchar, 
							  channel_link bigint)`

	// chat actions
	insertTelegramChat = `INSERT INTO telegram_chats 
    					  (user_id, chat_uuid, 
    					   telegram_chat_id, 
    					   format_string, 
    					   channel_link) 
						   values ($1, $2, $3, $4, $5)`

	insertStdoutChat = `INSERT INTO stdout_chats 
    					(user_id, chat_uuid, 
    					 format_string, 
    					 channel_link) 
						 values ($1, $2, $3, $4)`

	selectChatUuid = `SELECT chat_uuid FROM stdout_chats WHERE channel_link=$1`

	selectTelegramChat = `SELECT telegram_chat_id, 
       					  format_string 
						  FROM telegram_chats WHERE chat_uuid=$1`

	updateTelegramChat = `UPDATE telegram_chats set telegram_chat_id=$1 where chat_uuid=$2`

	selectFormatStringByStdoutChat = `SELECT format_string FROM stdout_chats WHERE chat_uuid=$1`

	selectChatsByChatUUID = `SELECT telegram_chat_id FROM telegram_chats WHERE chat_uuid=$1`

	// channels actions
	isExistChannelByChannelLink = `SELECT user_id, chat_uuid, channel_link FROM channels where channel_link=$1`

	selectChannelByChannelLink = `SELECT user_id, chat_uuid, channel_type, channel_link 
								  FROM channels where channel_link=$1`

	insertTelegramChannel = `INSERT INTO channels 
    						 (user_id, 
    						  chat_uuid, 
    						  channel_type, 
    						  channel_link) 
							  values ($1, $2, $3, $4)`

	insertStdoutChannel = `INSERT INTO channels 
    					   (user_id, 
    					    chat_uuid, 
    					    channel_type, 
    					    channel_link) 
						   values ($1, $2, $3, $4)`

	// users actions
	insertUser = `INSERT INTO users (user_id, chat_id) values ($1, $2)`

	isExistUserByUserId = `SELECT user_id FROM users where user_id=$1`
)
