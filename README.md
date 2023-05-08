# slack-spec

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/shadowcodex-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"slack-spec"
	"slack-spec/pkg/models/operations"
)

func main() {
    s := slackspec.New()

    ctx := context.Background()
    res, err := s.Admin.AdminAppsApprove(ctx, operations.AdminAppsApproveRequest{
        RequestBody: &operations.AdminAppsApproveApplicationJSON{
            AppID: slackspec.String("corrupti"),
            RequestID: slackspec.String("provident"),
            TeamID: slackspec.String("distinctio"),
        },
        Token: "quibusdam",
    }, operations.AdminAppsApproveSecurity{
        SlackAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DefaultSuccessTemplate != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Admin](docs/admin/README.md)

* [AdminAppsApprove](docs/admin/README.md#adminappsapprove) - Approve an app for installation on a workspace.
* [AdminAppsApprovedList](docs/admin/README.md#adminappsapprovedlist) - List approved apps for an org or workspace.
* [AdminAppsRequestsList](docs/admin/README.md#adminappsrequestslist) - List app requests for a team/workspace.
* [AdminAppsRestrict](docs/admin/README.md#adminappsrestrict) - Restrict an app for installation on a workspace.
* [AdminAppsRestrictedList](docs/admin/README.md#adminappsrestrictedlist) - List restricted apps for an org or workspace.
* [AdminConversationsArchive](docs/admin/README.md#adminconversationsarchive) - Archive a public or private channel.
* [AdminConversationsConvertToPrivate](docs/admin/README.md#adminconversationsconverttoprivate) - Convert a public channel to a private channel.
* [AdminConversationsCreate](docs/admin/README.md#adminconversationscreate) - Create a public or private channel-based conversation.
* [AdminConversationsDelete](docs/admin/README.md#adminconversationsdelete) - Delete a public or private channel.
* [AdminConversationsDisconnectShared](docs/admin/README.md#adminconversationsdisconnectshared) - Disconnect a connected channel from one or more workspaces.
* [AdminConversationsEkmListOriginalConnectedChannelInfo](docs/admin/README.md#adminconversationsekmlistoriginalconnectedchannelinfo) - List all disconnected channels—i.e., channels that were once connected to other workspaces and then disconnected—and the corresponding original channel IDs for key revocation with EKM.
* [AdminConversationsGetConversationPrefs](docs/admin/README.md#adminconversationsgetconversationprefs) - Get conversation preferences for a public or private channel.
* [AdminConversationsGetTeams](docs/admin/README.md#adminconversationsgetteams) - Get all the workspaces a given public or private channel is connected to within this Enterprise org.
* [AdminConversationsInvite](docs/admin/README.md#adminconversationsinvite) - Invite a user to a public or private channel.
* [AdminConversationsRename](docs/admin/README.md#adminconversationsrename) - Rename a public or private channel.
* [AdminConversationsRestrictAccessAddGroup](docs/admin/README.md#adminconversationsrestrictaccessaddgroup) - Add an allowlist of IDP groups for accessing a channel
* [AdminConversationsRestrictAccessListGroups](docs/admin/README.md#adminconversationsrestrictaccesslistgroups) - List all IDP Groups linked to a channel
* [AdminConversationsRestrictAccessRemoveGroup](docs/admin/README.md#adminconversationsrestrictaccessremovegroup) - Remove a linked IDP group linked from a private channel
* [AdminConversationsSearch](docs/admin/README.md#adminconversationssearch) - Search for public or private channels in an Enterprise organization.
* [AdminConversationsSetConversationPrefs](docs/admin/README.md#adminconversationssetconversationprefs) - Set the posting permissions for a public or private channel.
* [AdminConversationsSetTeams](docs/admin/README.md#adminconversationssetteams) - Set the workspaces in an Enterprise grid org that connect to a public or private channel.
* [AdminConversationsUnarchive](docs/admin/README.md#adminconversationsunarchive) - Unarchive a public or private channel.
* [AdminEmojiAdd](docs/admin/README.md#adminemojiadd) - Add an emoji.
* [AdminEmojiAddAlias](docs/admin/README.md#adminemojiaddalias) - Add an emoji alias.
* [AdminEmojiList](docs/admin/README.md#adminemojilist) - List emoji for an Enterprise Grid organization.
* [AdminEmojiRemove](docs/admin/README.md#adminemojiremove) - Remove an emoji across an Enterprise Grid organization
* [AdminEmojiRename](docs/admin/README.md#adminemojirename) - Rename an emoji.
* [AdminInviteRequestsApprove](docs/admin/README.md#admininviterequestsapprove) - Approve a workspace invite request.
* [AdminInviteRequestsApprovedList](docs/admin/README.md#admininviterequestsapprovedlist) - List all approved workspace invite requests.
* [AdminInviteRequestsDeniedList](docs/admin/README.md#admininviterequestsdeniedlist) - List all denied workspace invite requests.
* [AdminInviteRequestsDeny](docs/admin/README.md#admininviterequestsdeny) - Deny a workspace invite request.
* [AdminInviteRequestsList](docs/admin/README.md#admininviterequestslist) - List all pending workspace invite requests.
* [AdminTeamsAdminsList](docs/admin/README.md#adminteamsadminslist) - List all of the admins on a given workspace.
* [AdminTeamsCreate](docs/admin/README.md#adminteamscreate) - Create an Enterprise team.
* [AdminTeamsList](docs/admin/README.md#adminteamslist) - List all teams on an Enterprise organization
* [AdminTeamsOwnersList](docs/admin/README.md#adminteamsownerslist) - List all of the owners on a given workspace.
* [AdminTeamsSettingsInfo](docs/admin/README.md#adminteamssettingsinfo) - Fetch information about settings in a workspace
* [AdminTeamsSettingsSetDefaultChannels](docs/admin/README.md#adminteamssettingssetdefaultchannels) - Set the default channels of a workspace.
* [AdminTeamsSettingsSetDescription](docs/admin/README.md#adminteamssettingssetdescription) - Set the description of a given workspace.
* [AdminTeamsSettingsSetDiscoverability](docs/admin/README.md#adminteamssettingssetdiscoverability) - An API method that allows admins to set the discoverability of a given workspace
* [AdminTeamsSettingsSetIcon](docs/admin/README.md#adminteamssettingsseticon) - Sets the icon of a workspace.
* [AdminTeamsSettingsSetName](docs/admin/README.md#adminteamssettingssetname) - Set the name of a given workspace.
* [AdminUsergroupsAddChannels](docs/admin/README.md#adminusergroupsaddchannels) - Add one or more default channels to an IDP group.
* [AdminUsergroupsAddTeams](docs/admin/README.md#adminusergroupsaddteams) - Associate one or more default workspaces with an organization-wide IDP group.
* [AdminUsergroupsListChannels](docs/admin/README.md#adminusergroupslistchannels) - List the channels linked to an org-level IDP group (user group).
* [AdminUsergroupsRemoveChannels](docs/admin/README.md#adminusergroupsremovechannels) - Remove one or more default channels from an org-level IDP group (user group).
* [AdminUsersAssign](docs/admin/README.md#adminusersassign) - Add an Enterprise user to a workspace.
* [AdminUsersInvite](docs/admin/README.md#adminusersinvite) - Invite a user to a workspace.
* [AdminUsersList](docs/admin/README.md#adminuserslist) - List users on a workspace
* [AdminUsersRemove](docs/admin/README.md#adminusersremove) - Remove a user from a workspace.
* [AdminUsersSessionInvalidate](docs/admin/README.md#adminuserssessioninvalidate) - Invalidate a single session for a user by session_id
* [AdminUsersSessionReset](docs/admin/README.md#adminuserssessionreset) - Wipes all valid sessions on all devices for a given user
* [AdminUsersSetAdmin](docs/admin/README.md#adminuserssetadmin) - Set an existing guest, regular user, or owner to be an admin user.
* [AdminUsersSetExpiration](docs/admin/README.md#adminuserssetexpiration) - Set an expiration for a guest user
* [AdminUsersSetOwner](docs/admin/README.md#adminuserssetowner) - Set an existing guest, regular user, or admin user to be a workspace owner.
* [AdminUsersSetRegular](docs/admin/README.md#adminuserssetregular) - Set an existing guest user, admin user, or owner to be a regular user.

### [AdminApps](docs/adminapps/README.md)

* [AdminAppsApprove](docs/adminapps/README.md#adminappsapprove) - Approve an app for installation on a workspace.
* [AdminAppsRestrict](docs/adminapps/README.md#adminappsrestrict) - Restrict an app for installation on a workspace.

### [AdminAppsApproved](docs/adminappsapproved/README.md)

* [AdminAppsApprovedList](docs/adminappsapproved/README.md#adminappsapprovedlist) - List approved apps for an org or workspace.

### [AdminAppsRequests](docs/adminappsrequests/README.md)

* [AdminAppsRequestsList](docs/adminappsrequests/README.md#adminappsrequestslist) - List app requests for a team/workspace.

### [AdminAppsRestricted](docs/adminappsrestricted/README.md)

* [AdminAppsRestrictedList](docs/adminappsrestricted/README.md#adminappsrestrictedlist) - List restricted apps for an org or workspace.

### [AdminConversations](docs/adminconversations/README.md)

* [AdminConversationsArchive](docs/adminconversations/README.md#adminconversationsarchive) - Archive a public or private channel.
* [AdminConversationsConvertToPrivate](docs/adminconversations/README.md#adminconversationsconverttoprivate) - Convert a public channel to a private channel.
* [AdminConversationsCreate](docs/adminconversations/README.md#adminconversationscreate) - Create a public or private channel-based conversation.
* [AdminConversationsDelete](docs/adminconversations/README.md#adminconversationsdelete) - Delete a public or private channel.
* [AdminConversationsDisconnectShared](docs/adminconversations/README.md#adminconversationsdisconnectshared) - Disconnect a connected channel from one or more workspaces.
* [AdminConversationsGetConversationPrefs](docs/adminconversations/README.md#adminconversationsgetconversationprefs) - Get conversation preferences for a public or private channel.
* [AdminConversationsGetTeams](docs/adminconversations/README.md#adminconversationsgetteams) - Get all the workspaces a given public or private channel is connected to within this Enterprise org.
* [AdminConversationsInvite](docs/adminconversations/README.md#adminconversationsinvite) - Invite a user to a public or private channel.
* [AdminConversationsRename](docs/adminconversations/README.md#adminconversationsrename) - Rename a public or private channel.
* [AdminConversationsSearch](docs/adminconversations/README.md#adminconversationssearch) - Search for public or private channels in an Enterprise organization.
* [AdminConversationsSetConversationPrefs](docs/adminconversations/README.md#adminconversationssetconversationprefs) - Set the posting permissions for a public or private channel.
* [AdminConversationsSetTeams](docs/adminconversations/README.md#adminconversationssetteams) - Set the workspaces in an Enterprise grid org that connect to a public or private channel.
* [AdminConversationsUnarchive](docs/adminconversations/README.md#adminconversationsunarchive) - Unarchive a public or private channel.

### [AdminConversationsEkm](docs/adminconversationsekm/README.md)

* [AdminConversationsEkmListOriginalConnectedChannelInfo](docs/adminconversationsekm/README.md#adminconversationsekmlistoriginalconnectedchannelinfo) - List all disconnected channels—i.e., channels that were once connected to other workspaces and then disconnected—and the corresponding original channel IDs for key revocation with EKM.

### [AdminConversationsRestrictAccess](docs/adminconversationsrestrictaccess/README.md)

* [AdminConversationsRestrictAccessAddGroup](docs/adminconversationsrestrictaccess/README.md#adminconversationsrestrictaccessaddgroup) - Add an allowlist of IDP groups for accessing a channel
* [AdminConversationsRestrictAccessListGroups](docs/adminconversationsrestrictaccess/README.md#adminconversationsrestrictaccesslistgroups) - List all IDP Groups linked to a channel
* [AdminConversationsRestrictAccessRemoveGroup](docs/adminconversationsrestrictaccess/README.md#adminconversationsrestrictaccessremovegroup) - Remove a linked IDP group linked from a private channel

### [AdminEmoji](docs/adminemoji/README.md)

* [AdminEmojiAdd](docs/adminemoji/README.md#adminemojiadd) - Add an emoji.
* [AdminEmojiAddAlias](docs/adminemoji/README.md#adminemojiaddalias) - Add an emoji alias.
* [AdminEmojiList](docs/adminemoji/README.md#adminemojilist) - List emoji for an Enterprise Grid organization.
* [AdminEmojiRemove](docs/adminemoji/README.md#adminemojiremove) - Remove an emoji across an Enterprise Grid organization
* [AdminEmojiRename](docs/adminemoji/README.md#adminemojirename) - Rename an emoji.

### [AdminInviteRequests](docs/admininviterequests/README.md)

* [AdminInviteRequestsApprove](docs/admininviterequests/README.md#admininviterequestsapprove) - Approve a workspace invite request.
* [AdminInviteRequestsDeny](docs/admininviterequests/README.md#admininviterequestsdeny) - Deny a workspace invite request.
* [AdminInviteRequestsList](docs/admininviterequests/README.md#admininviterequestslist) - List all pending workspace invite requests.

### [AdminInviteRequestsApproved](docs/admininviterequestsapproved/README.md)

* [AdminInviteRequestsApprovedList](docs/admininviterequestsapproved/README.md#admininviterequestsapprovedlist) - List all approved workspace invite requests.

### [AdminInviteRequestsDenied](docs/admininviterequestsdenied/README.md)

* [AdminInviteRequestsDeniedList](docs/admininviterequestsdenied/README.md#admininviterequestsdeniedlist) - List all denied workspace invite requests.

### [AdminTeams](docs/adminteams/README.md)

* [AdminTeamsCreate](docs/adminteams/README.md#adminteamscreate) - Create an Enterprise team.
* [AdminTeamsList](docs/adminteams/README.md#adminteamslist) - List all teams on an Enterprise organization

### [AdminTeamsAdmins](docs/adminteamsadmins/README.md)

* [AdminTeamsAdminsList](docs/adminteamsadmins/README.md#adminteamsadminslist) - List all of the admins on a given workspace.

### [AdminTeamsOwners](docs/adminteamsowners/README.md)

* [AdminTeamsOwnersList](docs/adminteamsowners/README.md#adminteamsownerslist) - List all of the owners on a given workspace.

### [AdminTeamsSettings](docs/adminteamssettings/README.md)

* [AdminTeamsSettingsInfo](docs/adminteamssettings/README.md#adminteamssettingsinfo) - Fetch information about settings in a workspace
* [AdminTeamsSettingsSetDefaultChannels](docs/adminteamssettings/README.md#adminteamssettingssetdefaultchannels) - Set the default channels of a workspace.
* [AdminTeamsSettingsSetDescription](docs/adminteamssettings/README.md#adminteamssettingssetdescription) - Set the description of a given workspace.
* [AdminTeamsSettingsSetDiscoverability](docs/adminteamssettings/README.md#adminteamssettingssetdiscoverability) - An API method that allows admins to set the discoverability of a given workspace
* [AdminTeamsSettingsSetIcon](docs/adminteamssettings/README.md#adminteamssettingsseticon) - Sets the icon of a workspace.
* [AdminTeamsSettingsSetName](docs/adminteamssettings/README.md#adminteamssettingssetname) - Set the name of a given workspace.

### [AdminUsergroups](docs/adminusergroups/README.md)

* [AdminUsergroupsAddChannels](docs/adminusergroups/README.md#adminusergroupsaddchannels) - Add one or more default channels to an IDP group.
* [AdminUsergroupsAddTeams](docs/adminusergroups/README.md#adminusergroupsaddteams) - Associate one or more default workspaces with an organization-wide IDP group.
* [AdminUsergroupsListChannels](docs/adminusergroups/README.md#adminusergroupslistchannels) - List the channels linked to an org-level IDP group (user group).
* [AdminUsergroupsRemoveChannels](docs/adminusergroups/README.md#adminusergroupsremovechannels) - Remove one or more default channels from an org-level IDP group (user group).

### [AdminUsers](docs/adminusers/README.md)

* [AdminUsersAssign](docs/adminusers/README.md#adminusersassign) - Add an Enterprise user to a workspace.
* [AdminUsersInvite](docs/adminusers/README.md#adminusersinvite) - Invite a user to a workspace.
* [AdminUsersList](docs/adminusers/README.md#adminuserslist) - List users on a workspace
* [AdminUsersRemove](docs/adminusers/README.md#adminusersremove) - Remove a user from a workspace.
* [AdminUsersSetAdmin](docs/adminusers/README.md#adminuserssetadmin) - Set an existing guest, regular user, or owner to be an admin user.
* [AdminUsersSetExpiration](docs/adminusers/README.md#adminuserssetexpiration) - Set an expiration for a guest user
* [AdminUsersSetOwner](docs/adminusers/README.md#adminuserssetowner) - Set an existing guest, regular user, or admin user to be a workspace owner.
* [AdminUsersSetRegular](docs/adminusers/README.md#adminuserssetregular) - Set an existing guest user, admin user, or owner to be a regular user.

### [AdminUsersSession](docs/adminuserssession/README.md)

* [AdminUsersSessionInvalidate](docs/adminuserssession/README.md#adminuserssessioninvalidate) - Invalidate a single session for a user by session_id
* [AdminUsersSessionReset](docs/adminuserssession/README.md#adminuserssessionreset) - Wipes all valid sessions on all devices for a given user

### [API](docs/api/README.md)

* [APITest](docs/api/README.md#apitest) - Checks API calling code.

### [Apps](docs/apps/README.md)

* [AppsEventAuthorizationsList](docs/apps/README.md#appseventauthorizationslist) - Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to.
* [AppsPermissionsInfo](docs/apps/README.md#appspermissionsinfo) - Returns list of permissions this app has on a team.
* [AppsPermissionsRequest](docs/apps/README.md#appspermissionsrequest) - Allows an app to request additional scopes
* [AppsPermissionsResourcesList](docs/apps/README.md#appspermissionsresourceslist) - Returns list of resource grants this app has on a team.
* [AppsPermissionsScopesList](docs/apps/README.md#appspermissionsscopeslist) - Returns list of scopes this app has on a team.
* [AppsUninstall](docs/apps/README.md#appsuninstall) - Uninstalls your app from a workspace.

### [AppsEventAuthorizations](docs/appseventauthorizations/README.md)

* [AppsEventAuthorizationsList](docs/appseventauthorizations/README.md#appseventauthorizationslist) - Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to.

### [AppsPermissions](docs/appspermissions/README.md)

* [AppsPermissionsInfo](docs/appspermissions/README.md#appspermissionsinfo) - Returns list of permissions this app has on a team.
* [AppsPermissionsRequest](docs/appspermissions/README.md#appspermissionsrequest) - Allows an app to request additional scopes

### [AppsPermissionsResources](docs/appspermissionsresources/README.md)

* [AppsPermissionsResourcesList](docs/appspermissionsresources/README.md#appspermissionsresourceslist) - Returns list of resource grants this app has on a team.

### [AppsPermissionsScopes](docs/appspermissionsscopes/README.md)

* [AppsPermissionsScopesList](docs/appspermissionsscopes/README.md#appspermissionsscopeslist) - Returns list of scopes this app has on a team.

### [Auth](docs/auth/README.md)

* [AuthRevoke](docs/auth/README.md#authrevoke) - Revokes a token.
* [AuthTest](docs/auth/README.md#authtest) - Checks authentication & identity.

### [Bots](docs/bots/README.md)

* [BotsInfo](docs/bots/README.md#botsinfo) - Gets information about a bot user.

### [Calls](docs/calls/README.md)

* [CallsAdd](docs/calls/README.md#callsadd) - Registers a new Call.
* [CallsEnd](docs/calls/README.md#callsend) - Ends a Call.
* [CallsInfo](docs/calls/README.md#callsinfo) - Returns information about a Call.
* [CallsParticipantsAdd](docs/calls/README.md#callsparticipantsadd) - Registers new participants added to a Call.
* [CallsParticipantsRemove](docs/calls/README.md#callsparticipantsremove) - Registers participants removed from a Call.
* [CallsUpdate](docs/calls/README.md#callsupdate) - Updates information about a Call.

### [CallsParticipants](docs/callsparticipants/README.md)

* [CallsParticipantsAdd](docs/callsparticipants/README.md#callsparticipantsadd) - Registers new participants added to a Call.
* [CallsParticipantsRemove](docs/callsparticipants/README.md#callsparticipantsremove) - Registers participants removed from a Call.

### [Chat](docs/chat/README.md)

* [ChatDelete](docs/chat/README.md#chatdelete) - Deletes a message.
* [ChatDeleteScheduledMessage](docs/chat/README.md#chatdeletescheduledmessage) - Deletes a pending scheduled message from the queue.
* [ChatGetPermalink](docs/chat/README.md#chatgetpermalink) - Retrieve a permalink URL for a specific extant message
* [ChatMeMessage](docs/chat/README.md#chatmemessage) - Share a me message into a channel.
* [ChatPostEphemeral](docs/chat/README.md#chatpostephemeral) - Sends an ephemeral message to a user in a channel.
* [ChatPostMessage](docs/chat/README.md#chatpostmessage) - Sends a message to a channel.
* [ChatScheduleMessage](docs/chat/README.md#chatschedulemessage) - Schedules a message to be sent to a channel.
* [ChatScheduledMessagesList](docs/chat/README.md#chatscheduledmessageslist) - Returns a list of scheduled messages.
* [ChatUnfurl](docs/chat/README.md#chatunfurl) - Provide custom unfurl behavior for user-posted URLs
* [ChatUpdate](docs/chat/README.md#chatupdate) - Updates a message.

### [ChatScheduledMessages](docs/chatscheduledmessages/README.md)

* [ChatScheduledMessagesList](docs/chatscheduledmessages/README.md#chatscheduledmessageslist) - Returns a list of scheduled messages.

### [Conversations](docs/conversations/README.md)

* [ConversationsArchive](docs/conversations/README.md#conversationsarchive) - Archives a conversation.
* [ConversationsClose](docs/conversations/README.md#conversationsclose) - Closes a direct message or multi-person direct message.
* [ConversationsCreate](docs/conversations/README.md#conversationscreate) - Initiates a public or private channel-based conversation
* [ConversationsHistory](docs/conversations/README.md#conversationshistory) - Fetches a conversation's history of messages and events.
* [ConversationsInfo](docs/conversations/README.md#conversationsinfo) - Retrieve information about a conversation.
* [ConversationsInvite](docs/conversations/README.md#conversationsinvite) - Invites users to a channel.
* [ConversationsJoin](docs/conversations/README.md#conversationsjoin) - Joins an existing conversation.
* [ConversationsKick](docs/conversations/README.md#conversationskick) - Removes a user from a conversation.
* [ConversationsLeave](docs/conversations/README.md#conversationsleave) - Leaves a conversation.
* [ConversationsList](docs/conversations/README.md#conversationslist) - Lists all channels in a Slack team.
* [ConversationsMark](docs/conversations/README.md#conversationsmark) - Sets the read cursor in a channel.
* [ConversationsMembers](docs/conversations/README.md#conversationsmembers) - Retrieve members of a conversation.
* [ConversationsOpen](docs/conversations/README.md#conversationsopen) - Opens or resumes a direct message or multi-person direct message.
* [ConversationsRename](docs/conversations/README.md#conversationsrename) - Renames a conversation.
* [ConversationsReplies](docs/conversations/README.md#conversationsreplies) - Retrieve a thread of messages posted to a conversation
* [ConversationsSetPurpose](docs/conversations/README.md#conversationssetpurpose) - Sets the purpose for a conversation.
* [ConversationsSetTopic](docs/conversations/README.md#conversationssettopic) - Sets the topic for a conversation.
* [ConversationsUnarchive](docs/conversations/README.md#conversationsunarchive) - Reverses conversation archival.

### [Dialog](docs/dialog/README.md)

* [DialogOpen](docs/dialog/README.md#dialogopen) - Open a dialog with a user

### [Dnd](docs/dnd/README.md)

* [DndEndDnd](docs/dnd/README.md#dndenddnd) - Ends the current user's Do Not Disturb session immediately.
* [DndEndSnooze](docs/dnd/README.md#dndendsnooze) - Ends the current user's snooze mode immediately.
* [DndInfo](docs/dnd/README.md#dndinfo) - Retrieves a user's current Do Not Disturb status.
* [DndSetSnooze](docs/dnd/README.md#dndsetsnooze) - Turns on Do Not Disturb mode for the current user, or changes its duration.
* [DndTeamInfo](docs/dnd/README.md#dndteaminfo) - Retrieves the Do Not Disturb status for up to 50 users on a team.

### [Emoji](docs/emoji/README.md)

* [EmojiList](docs/emoji/README.md#emojilist) - Lists custom emoji for a team.

### [Files](docs/files/README.md)

* [FilesCommentsDelete](docs/files/README.md#filescommentsdelete) - Deletes an existing comment on a file.
* [FilesDelete](docs/files/README.md#filesdelete) - Deletes a file.
* [FilesInfo](docs/files/README.md#filesinfo) - Gets information about a file.
* [FilesList](docs/files/README.md#fileslist) - List for a team, in a channel, or from a user with applied filters.
* [FilesRemoteAdd](docs/files/README.md#filesremoteadd) - Adds a file from a remote service
* [FilesRemoteInfo](docs/files/README.md#filesremoteinfo) - Retrieve information about a remote file added to Slack
* [FilesRemoteList](docs/files/README.md#filesremotelist) - Retrieve information about a remote file added to Slack
* [FilesRemoteRemove](docs/files/README.md#filesremoteremove) - Remove a remote file.
* [FilesRemoteShare](docs/files/README.md#filesremoteshare) - Share a remote file into a channel.
* [FilesRemoteUpdate](docs/files/README.md#filesremoteupdate) - Updates an existing remote file.
* [FilesRevokePublicURL](docs/files/README.md#filesrevokepublicurl) - Revokes public/external sharing access for a file
* [FilesSharedPublicURL](docs/files/README.md#filessharedpublicurl) - Enables a file for public/external sharing.
* [FilesUpload](docs/files/README.md#filesupload) - Uploads or creates a file.

### [FilesComments](docs/filescomments/README.md)

* [FilesCommentsDelete](docs/filescomments/README.md#filescommentsdelete) - Deletes an existing comment on a file.

### [FilesRemote](docs/filesremote/README.md)

* [FilesRemoteAdd](docs/filesremote/README.md#filesremoteadd) - Adds a file from a remote service
* [FilesRemoteInfo](docs/filesremote/README.md#filesremoteinfo) - Retrieve information about a remote file added to Slack
* [FilesRemoteList](docs/filesremote/README.md#filesremotelist) - Retrieve information about a remote file added to Slack
* [FilesRemoteRemove](docs/filesremote/README.md#filesremoteremove) - Remove a remote file.
* [FilesRemoteShare](docs/filesremote/README.md#filesremoteshare) - Share a remote file into a channel.
* [FilesRemoteUpdate](docs/filesremote/README.md#filesremoteupdate) - Updates an existing remote file.

### [Migration](docs/migration/README.md)

* [MigrationExchange](docs/migration/README.md#migrationexchange) - For Enterprise Grid workspaces, map local user IDs to global user IDs

### [Oauth](docs/oauth/README.md)

* [OauthAccess](docs/oauth/README.md#oauthaccess) - Exchanges a temporary OAuth verifier code for an access token.
* [OauthToken](docs/oauth/README.md#oauthtoken) - Exchanges a temporary OAuth verifier code for a workspace token.
* [OauthV2Access](docs/oauth/README.md#oauthv2access) - Exchanges a temporary OAuth verifier code for an access token.

### [OauthV2](docs/oauthv2/README.md)

* [OauthV2Access](docs/oauthv2/README.md#oauthv2access) - Exchanges a temporary OAuth verifier code for an access token.

### [Pins](docs/pins/README.md)

* [PinsAdd](docs/pins/README.md#pinsadd) - Pins an item to a channel.
* [PinsList](docs/pins/README.md#pinslist) - Lists items pinned to a channel.
* [PinsRemove](docs/pins/README.md#pinsremove) - Un-pins an item from a channel.

### [Reactions](docs/reactions/README.md)

* [ReactionsAdd](docs/reactions/README.md#reactionsadd) - Adds a reaction to an item.
* [ReactionsGet](docs/reactions/README.md#reactionsget) - Gets reactions for an item.
* [ReactionsList](docs/reactions/README.md#reactionslist) - Lists reactions made by a user.
* [ReactionsRemove](docs/reactions/README.md#reactionsremove) - Removes a reaction from an item.

### [Reminders](docs/reminders/README.md)

* [RemindersAdd](docs/reminders/README.md#remindersadd) - Creates a reminder.
* [RemindersComplete](docs/reminders/README.md#reminderscomplete) - Marks a reminder as complete.
* [RemindersDelete](docs/reminders/README.md#remindersdelete) - Deletes a reminder.
* [RemindersInfo](docs/reminders/README.md#remindersinfo) - Gets information about a reminder.
* [RemindersList](docs/reminders/README.md#reminderslist) - Lists all reminders created by or for a given user.

### [Rtm](docs/rtm/README.md)

* [RtmConnect](docs/rtm/README.md#rtmconnect) - Starts a Real Time Messaging session.

### [Search](docs/search/README.md)

* [SearchMessages](docs/search/README.md#searchmessages) - Searches for messages matching a query.

### [Stars](docs/stars/README.md)

* [StarsAdd](docs/stars/README.md#starsadd) - Adds a star to an item.
* [StarsList](docs/stars/README.md#starslist) - Lists stars for a user.
* [StarsRemove](docs/stars/README.md#starsremove) - Removes a star from an item.

### [Team](docs/team/README.md)

* [TeamAccessLogs](docs/team/README.md#teamaccesslogs) - Gets the access logs for the current team.
* [TeamBillableInfo](docs/team/README.md#teambillableinfo) - Gets billable users information for the current team.
* [TeamInfo](docs/team/README.md#teaminfo) - Gets information about the current team.
* [TeamIntegrationLogs](docs/team/README.md#teamintegrationlogs) - Gets the integration logs for the current team.
* [TeamProfileGet](docs/team/README.md#teamprofileget) - Retrieve a team's profile.

### [TeamProfile](docs/teamprofile/README.md)

* [TeamProfileGet](docs/teamprofile/README.md#teamprofileget) - Retrieve a team's profile.

### [Usergroups](docs/usergroups/README.md)

* [UsergroupsCreate](docs/usergroups/README.md#usergroupscreate) - Create a User Group
* [UsergroupsDisable](docs/usergroups/README.md#usergroupsdisable) - Disable an existing User Group
* [UsergroupsEnable](docs/usergroups/README.md#usergroupsenable) - Enable a User Group
* [UsergroupsList](docs/usergroups/README.md#usergroupslist) - List all User Groups for a team
* [UsergroupsUpdate](docs/usergroups/README.md#usergroupsupdate) - Update an existing User Group
* [UsergroupsUsersList](docs/usergroups/README.md#usergroupsuserslist) - List all users in a User Group
* [UsergroupsUsersUpdate](docs/usergroups/README.md#usergroupsusersupdate) - Update the list of users for a User Group

### [UsergroupsUsers](docs/usergroupsusers/README.md)

* [UsergroupsUsersList](docs/usergroupsusers/README.md#usergroupsuserslist) - List all users in a User Group
* [UsergroupsUsersUpdate](docs/usergroupsusers/README.md#usergroupsusersupdate) - Update the list of users for a User Group

### [Users](docs/users/README.md)

* [UsersConversations](docs/users/README.md#usersconversations) - List conversations the calling user may access.
* [UsersDeletePhoto](docs/users/README.md#usersdeletephoto) - Delete the user profile photo
* [UsersGetPresence](docs/users/README.md#usersgetpresence) - Gets user presence information.
* [UsersIdentity](docs/users/README.md#usersidentity) - Get a user's identity.
* [UsersInfo](docs/users/README.md#usersinfo) - Gets information about a user.
* [UsersList](docs/users/README.md#userslist) - Lists all users in a Slack team.
* [UsersLookupByEmail](docs/users/README.md#userslookupbyemail) - Find a user with an email address.
* [UsersProfileGet](docs/users/README.md#usersprofileget) - Retrieves a user's profile information.
* [UsersProfileSet](docs/users/README.md#usersprofileset) - Set the profile information for a user.
* [UsersSetActive](docs/users/README.md#userssetactive) - Marked a user as active. Deprecated and non-functional.
* [UsersSetPhoto](docs/users/README.md#userssetphoto) - Set the user profile photo
* [UsersSetPresence](docs/users/README.md#userssetpresence) - Manually sets user presence.

### [UsersProfile](docs/usersprofile/README.md)

* [UsersProfileGet](docs/usersprofile/README.md#usersprofileget) - Retrieves a user's profile information.
* [UsersProfileSet](docs/usersprofile/README.md#usersprofileset) - Set the profile information for a user.

### [Views](docs/views/README.md)

* [ViewsOpen](docs/views/README.md#viewsopen) - Open a view for a user.
* [ViewsPublish](docs/views/README.md#viewspublish) - Publish a static view for a User.
* [ViewsPush](docs/views/README.md#viewspush) - Push a view onto the stack of a root view.
* [ViewsUpdate](docs/views/README.md#viewsupdate) - Update an existing view.

### [Workflows](docs/workflows/README.md)

* [WorkflowsStepCompleted](docs/workflows/README.md#workflowsstepcompleted) - Indicate that an app's step in a workflow completed execution.
* [WorkflowsStepFailed](docs/workflows/README.md#workflowsstepfailed) - Indicate that an app's step in a workflow failed to execute.
* [WorkflowsUpdateStep](docs/workflows/README.md#workflowsupdatestep) - Update the configuration for a workflow extension step.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
