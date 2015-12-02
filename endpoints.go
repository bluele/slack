package slack

// ApiBaseUrl can be changed for private Slack solutions
var ApiBaseUrl = "https://slack.com/api/"

const (
	channelsListApiEndpoint    = "channels.list"
	channelsJoinApiEndpoint    = "channels.join"
	channelsHistoryApiEndpoint = "channels.history"

	chatPostMessageApiEndpoint = "chat.postMessage"

	groupsInviteApiEndpoint = "groups.invite"
	groupsListApiEndpoint   = "groups.list"
	groupsCreateApiEndpoint = "groups.create"

	filesUploadApiEndpoint = "files.upload"

	usersInfoApiEndpoint = "users.info"
	usersListApiEndpoint = "users.list"
)
