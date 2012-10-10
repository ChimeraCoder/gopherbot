// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package proto

// Unknown message type.
const Unknown = 0

// Replies in the range from 001 to 099 are used for client-server connections
// only and should never travel between servers.
const (
	Welcome   = 1 // Welcome to the Internet Relay Network <nick>!<user>@<host>
	YourHost  = 2 // Your host is <servername>, running version <ver>
	Cmdreated = 3 // This server was created <date>
	MyInfo    = 4 // <servername> <version> <available user modes> <available channel modes>
	Bound     = 5 // Try server <server name>, port <port number>
)

// Replies generated in the response to commands are found in the
// range from 200 to 399.
const (
	TraceLink       = 200 // Link <version & debug level> <destination> <next server> V<protocol version> <link uptime in seconds> <backstream sendq> <upstream sendq>
	TraceConnecting = 201 // Try. <class> <server>
	TraceHandshake  = 202 // H.S. <class> <server>
	TraceUnknown    = 203 // ???? <class> [<client IP address in dot form>]
	TraceOperator   = 204 // Oper <class> <nick>
	TraceUser       = 205 // User <class> <nick>
	TraceServer     = 206 // Serv <class> <int>S <int>C <server> <nick!user|*!*>@<host|server> V<protocol version>
	TraceService    = 207 // Service <class> <name> <type> <active type>
	TraceNewType    = 208 // <newtype> 0 <client name>
	TraceClass      = 209 // Class <class> <count>
	TraceReconnect  = 210 // Unused.
	StatsLinkInfo   = 211 // <linkname> <sendq> <sent messages> <sent Kbytes> <received messages> <received Kbytes> <time open>
	StatsCommands   = 212 // <command> <count> <byte count> <remote count>
	EndOfStats      = 219 // <stats letter> :End of STATS report
	UModeIs         = 221 // <user mode string>
	ServList        = 234 // <name> <server> <mask> <type> <hopcount> <info>
	ServListEnd     = 235 // <mask> <type> :End of service listing
	StatsUptime     = 242 // :Server Up %d days %d:%02d:%02d
	StatsOLine      = 243 // O <hostmask> * <name>
	LUserClient     = 251 // :There are <integer> users and <integer> services on <integer> servers
	LUserOp         = 252 // <integer> :operator(s) online
	LUserUnknown    = 253 // <integer> :unknown connection(s)
	LUserChannels   = 254 // <integer> :channels formed
	LUserMe         = 255 // :I have <integer> clients and <integer> servers
	AdminMe         = 256 // <server> :Administrative info
	AdminLoc1       = 257 // :<admin info>
	AdminLoc2       = 258 // :<admin info>
	AdminEmail      = 259 // :<admin info>
	TraceLog        = 261 // File <logfile> <debug level>
	TraceEnd        = 262 // <server name> <version & debug level> :End of TRACE
	TryAgain        = 263 // <command> :Please wait a while and try again.
	Away            = 301 // <nick> :<away message>
	UserHost        = 302 // :*1<reply> *(   <reply> )
	IsOn            = 303 // :*1<nick> *(   <nick> )
	UnAway          = 305 // :You are no longer marked as being away
	NowAway         = 306 // :You have been marked as being away
	WhoIsUser       = 311 // <nick> <user> <host> * :<real name>
	WhoIsServer     = 312 // <nick> <server> :<server info>
	WhoIsOperator   = 313 // <nick> :is an IRC operator
	WhoWasUser      = 314 // <nick> <user> <host> * :<real name>
	EndOfWho        = 315 // <name> :End of WHO list
	WhoIsIdle       = 317 // <nick> <integer> :seconds idle
	EndOfWhoIs      = 318 // <nick> :End of WHOIS list
	WhoIsChannels   = 319 // <nick> :*( ( @ / + ) <channel>   )
	ListStart       = 321 // Obsolete.
	List            = 322 // <channel> <# visible> :<topic>
	ListEnd         = 323 // :End of LIST
	CmdhannelModeIs = 324 // <channel> <mode> <mode params>
	UniqOpIs        = 325 // <channel> <nickname>
	NoTopic         = 331 // <channel> :No topic is set
	Topic           = 332 // <channel> :<topic>
	Inviting        = 341 // <channel> <nick>
	Summoning       = 342 // <user> :Summoning user to IRC
	InviteList      = 346 // <channel> <invitemask>
	EndOfInviteList = 347 // <channel> :End of channel invite list
	ExceptList      = 348 // <channel> <exceptionmask>
	EndOfExceptList = 349 // <channel> :End of channel exception list
	Version         = 351 // <version>.<debuglevel> <server> :<comments>
	WhoReply        = 352 // <channel> <user> <host> <server> <nick> ( H / G > [*] [ ( @ / + ) ] :<hopcount> <real name>
	NameReply       = 353 // ( = / * / @ ) <channel> :[ @ / + ] <nick> *(   [ @ / + ] <nick> )
	Links           = 364 // <mask> <server> :<hopcount> <server info>
	EndOfLinks      = 365 // <mask> :End of LINKS list
	EndOfNames      = 366 // <channel> :End of NAMES list
	BanList         = 367 // <channel> <banmask>
	EndOfBanList    = 368 // <channel> :End of channel ban list
	EndOfWhoWas     = 369 // <nick> :End of WHOWAS
	Info            = 371 // :<string>
	MOTD            = 372 // :- <text>
	EndofInfo       = 374 // :End of INFO list
	MOTDStart       = 375 // :- <server> Message of the day - 
	EndOfMOTD       = 376 // :End of MOTD command
	YouAreOper      = 381 // :You are now an IRC operator
	Rehasing        = 382 // <config file> :Rehashing
	YouAreService   = 383 // You are service <servicename>
	Time            = 391 // <server> :<string showing server's local time>
	UserStart       = 392 // :UserID Terminal Host
	Users           = 393 // :<username> <ttyline> <hostname>
	EndOfUsers      = 394 // :End of users
	Nousers         = 395 // :Nobody logged in
)

// Error replies are found in the range from 400 to 599.
const (
	ErrNoSuchNick          = 401 // <nickname> :No such nick/channel
	ErrNoSuchServer        = 402 // <server name> :No such server
	ErrNoSuchChannel       = 403 // <channel name> :No such channel
	ErrCannotSendToChannel = 404 // <channel name> :Cannot send to channel
	ErrTooManyChannels     = 405 // <channel name> :You have joined too many channels
	ErrWasNoSuchNick       = 406 // <nickname> :There was no such nickname
	ErrTooManyTargets      = 407 // <target> :<error code> recipients. <abort message>
	ErrNoSuchService       = 408 // <service name> :No such service
	ErrNoOrigin            = 409 // :No origin specified
	ErrNoRecipient         = 411 // :No recipient given (<command>)
	ErrNoTextToSend        = 412 // :No text to send
	ErrNoTopLevel          = 413 // <mask> :No toplevel domain specified
	ErrWildTopLevel        = 414 // <mask> :Wildcard in toplevel domain
	ErrBadMask             = 415 // <mask> :Bad Server/host mask
	ErrUnknownCommand      = 421 // <command> :Unknown command
	ErrNoMOTD              = 422 // :MOTD File is missing
	ErrNoAdminInfo         = 423 // <server> :No administrative info available
	ErrFileError           = 424 // :File error doing <file op> on <file>
	ErrNoNicknameGiven     = 431 // :No nickname given
	ErrErroneusNickname    = 432 // <nick> :Erroneous nickname
	ErrNicknameInUse       = 433 // <nick> :Nickname is already in use
	ErrErrNickCollision    = 436 // <nick> :Nickname collision KILL from <user>@<host>
	ErrUnavailableResource = 437 // <nick/channel> :Nick/channel is temporarily unavailable
	ErrUserNotInChannel    = 441 // <nick> <channel> :They aren't on that channel
	ErrNotOnChannel        = 442 // <channel> :You're not on that channel
	ErrUserOnChannel       = 443 // <user> <channel> :is already on channel
	ErrNoLogin             = 444 // <user> :User not logged in
	ErrSummonDisabled      = 445 // :SUMMON has been disabled
	ErrUserDisabled        = 446 // :USERS has been disabled
	ErrNotRegistered       = 451 // :You have not registered
	ErrNeedMoreParams      = 461 // <command> :Not enough parameters
	ErrAlreadyRegistered   = 462 // :Unauthorized command (already registered)
	ErrNoPermForHost       = 463 // :Your host isn't among the privileged
	ErrPasswordMismatch    = 464 // :Password incorrect
	ErrYouAreBanned        = 465 // :You are banned from this server
	ErrYouWillBeBanned     = 466 //  
	ErrKeySet              = 467 // <channel> :Channel key already set
	ErrChannelIsFull       = 471 // <channel> :Cannot join channel (+l)
	ErrUnknownMode         = 472 // <char> :is unknown mode char to me for <channel>
	ErrInviteOnlyChannel   = 473 // <channel> :Cannot join channel (+i)
	ErrBannedFromChannel   = 474 // <channel> :Cannot join channel (+b)
	ErrBadChannelKey       = 475 // <channel> :Cannot join channel (+k)
	ErrBadChannelMask      = 476 // <channel> :Bad Channel Mask
	ErrNoChannelModes      = 477 // <channel> :Channel doesn't support modes
	ErrBanListFull         = 478 // <channel> <char> :Channel list is full
	ErrNoPrivileges        = 481 // :Permission Denied- You're not an IRC operator
	ErrChannelOPrivsNeeded = 482 // <channel> :You're not channel operator
	ErrCannotKillServer    = 483 // :You can't kill a server!
	ErrRestricted          = 484 // :Your connection is restricted!
	ErrUniqOPrivsNeeded    = 485 // :You're not the original channel operator
	ErrNoOperHost          = 491 // :No O-lines for your host
	ErrUModeUnknownFlag    = 501 // :Unknown MODE flag
	ErrUsersDoNotMatch     = 502 // :Cannot change mode for other users
)

// Command identifiers.
const (
	CmdAdmin    = 801 // Get information about the administrator of a server.	FC 2812
	CmdAway     = 802 // Set an automatic reply string for any PRIVMSG commands.	FC 2812
	CmdConnect  = 803 // Request a new connection to another server immediately.	FC 2812
	CmdDie      = 804 // Shutdown the server.	FC 2812
	CmdError    = 805 // Report a serious or fatal error to a peer.	FC 2812
	CmdInfo     = 806 // Get information describing a server.	FC 2812
	CmdInvite   = 807 // Invite a user to a channel.	FC 2812
	CmdIsOn     = 808 // Determine if a nickname is currently on IRC.	FC 2812
	CmdJoin     = 809 // Join a channel.	FC 2812, RFC 2813
	CmdKick     = 810 // Request the forced removal of a user from a channel.	FC 2812
	CmdKill     = 811 // Close a client-server connection by the server which has the actual connection.	FC 2812
	CmdLinks    = 812 // List all servernames which are known by the server answering the query.	FC 2812
	CmdList     = 813 // List channels and their topics.	FC 2812
	CmdLUsers   = 814 // Get statistics about the size of the IRC network.	FC 2812
	CmdMode     = 815 // User mode.	FC 2812
	CmdMOTD     = 816 // Get the Message of the Day.	FC 2812
	CmdNames    = 817 // List all visible nicknames.	FC 2812
	CmdNick     = 818 // Define a nickname.	FC 2812, RFC 2813
	CmdNJoin    = 819 // Exchange the list of channel members for each channel between servers.	FC 2813
	CmdNotice   = 820 // RFC 2812
	CmdOper     = 821 // Obtain operator privileges.	FC 2812
	CmdPart     = 822 // Leave a channel.	FC 2812
	CmdPass     = 823 // Set a connection password.	FC 2812, RFC 2813
	CmdPing     = 824 // Test for the presence of an active client or server.	FC 2812
	CmdPong     = 825 // Reply to a PING message.	FC 2812
	CmdPrivMsg  = 826 // Send private messages between users, as well as to send messages to channels.	FC 2812
	CmdQuit     = 827 // Terminate the client session.	FC 2812, RFC 2813
	CmdRehash   = 828 // Force the server to re-read and process its configuration file.	FC 2812
	CmdRestart  = 829 // Force the server to restart itself.	FC 2812
	CmdServer   = 830 // Register a new server.	FC 2813
	CmdService  = 831 // Register a new service.	FC 2812, RFC 2813
	CmdServList = 832 // List services currently connected to the network.	FC 2812
	CmdSQuery   = 833 // RFC 2812
	CmdSquirt   = 834 // Disconnect a server link.	FC 2812
	CmdSQuit    = 835 // Break a local or remote server link.	FC 2813
	CmdStats    = 836 // Get server statistics.	FC 2812
	CmdSummon   = 837 // Ask a user to join IRC.	FC 2812
	CmdTime     = 838 // Get the local time from the specified server.	FC 2812
	CmdTopic    = 839 // Change or view the topic of a channel.	FC 2812
	CmdTrace    = 840 // Find the route to a server and information about it's peers.	FC 2812
	CmdUser     = 841 // Specify the username, hostname and realname of a new user.	FC 2812
	CmdUserHost = 842 // Get a list of information about upto 5 nicknames.	FC 2812
	CmdUsers    = 843 // Get a list of users logged into the server.	FC 2812
	CmdVersion  = 844 // Get the version of the server program.	FC 2812
	CmdWAllOps  = 845 // Send a message to all currently connected users who have set the 'w' user mode.	FC 2812
	CmdWho      = 846 // List a set of users.	FC 2812
	CmdWhoIs    = 847 // Get information about a specific user.	FC 2812
	CmdWhoWas   = 848 // Get information about a nickname which no longer exists.	FC 2812
)
