// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package proto

// http://www.networksorcery.com/enp/protocol/irc.htm

// Unknown message type.
const Unknown = 0

// Replies in the range from 001 to 099 are used for client-server connections
// only and should never travel between servers.
const (
	RWelcome  = 1 // Welcome to the Internet Relay Network <nick>!<user>@<host>
	RYourHost = 2 // Your host is <servername>, running version <ver>
	RCreated  = 3 // This server was created <date>
	RMyInfo   = 4 // <servername> <version> <available user modes> <available channel modes>
	RBound    = 5 // Try server <server name>, port <port number>
)

// Replies generated in the response to commands are found in the
// range from 200 to 399.
const (
	RTraceLink       = 200 // Link <version & debug level> <destination> <next server> V<protocol version> <link uptime in seconds> <backstream sendq> <upstream sendq>
	RTraceConnecting = 201 // Try. <class> <server>
	RTraceHandshake  = 202 // H.S. <class> <server>
	RTraceUnknown    = 203 // ???? <class> [<client IP address in dot form>]
	RTraceOperator   = 204 // Oper <class> <nick>
	RTraceUser       = 205 // User <class> <nick>
	RTraceServer     = 206 // Serv <class> <int>S <int>C <server> <nick!user|*!*>@<host|server> V<protocol version>
	RTraceService    = 207 // Service <class> <name> <type> <active type>
	RTraceNewType    = 208 // <newtype> 0 <client name>
	RTraceClass      = 209 // Class <class> <count>
	RTraceReconnect  = 210 // Unused.
	RStatsLinkInfo   = 211 // <linkname> <sendq> <sent messages> <sent Kbytes> <received messages> <received Kbytes> <time open>
	RStatsCommands   = 212 // <command> <count> <byte count> <remote count>
	REndOfStats      = 219 // <stats letter> :End of STATS report
	RUModeIs         = 221 // <user mode string>
	RServList        = 234 // <name> <server> <mask> <type> <hopcount> <info>
	RServListEnd     = 235 // <mask> <type> :End of service listing
	RStatsUptime     = 242 // :Server Up %d days %d:%02d:%02d
	RStatsOLine      = 243 // O <hostmask> * <name>
	RLUserClient     = 251 // :There are <integer> users and <integer> services on <integer> servers
	RLUserOp         = 252 // <integer> :operator(s) online
	RLUserUnknown    = 253 // <integer> :unknown connection(s)
	RLUserChannels   = 254 // <integer> :channels formed
	RLUserMe         = 255 // :I have <integer> clients and <integer> servers
	RAdminMe         = 256 // <server> :Administrative info
	RAdminLoc1       = 257 // :<admin info>
	RAdminLoc2       = 258 // :<admin info>
	RAdminEmail      = 259 // :<admin info>
	RTraceLog        = 261 // File <logfile> <debug level>
	RTraceEnd        = 262 // <server name> <version & debug level> :End of TRACE
	RTryAgain        = 263 // <command> :Please wait a while and try again.
	RAway            = 301 // <nick> :<away message>
	RUserHost        = 302 // :*1<reply> *(   <reply> )
	RIsOn            = 303 // :*1<nick> *(   <nick> )
	RUnAway          = 305 // :You are no longer marked as being away
	RNowAway         = 306 // :You have been marked as being away
	RWhoIsUser       = 311 // <nick> <user> <host> * :<real name>
	RWhoIsServer     = 312 // <nick> <server> :<server info>
	RWhoIsOperator   = 313 // <nick> :is an IRC operator
	RWhoWasUser      = 314 // <nick> <user> <host> * :<real name>
	REndOfWho        = 315 // <name> :End of WHO list
	RWhoIsIdle       = 317 // <nick> <integer> :seconds idle
	REndOfWhoIs      = 318 // <nick> :End of WHOIS list
	RWhoIsChannels   = 319 // <nick> :*( ( @ / + ) <channel>   )
	RListStart       = 321 // Obsolete.
	RList            = 322 // <channel> <# visible> :<topic>
	RListEnd         = 323 // :End of LIST
	RChannelModeIs   = 324 // <channel> <mode> <mode params>
	RUniqOpIs        = 325 // <channel> <nickname>
	RNoTopic         = 331 // <channel> :No topic is set
	RTopic           = 332 // <channel> :<topic>
	RInviting        = 341 // <channel> <nick>
	RSummoning       = 342 // <user> :Summoning user to IRC
	RInviteList      = 346 // <channel> <invitemask>
	REndOfInviteList = 347 // <channel> :End of channel invite list
	RExceptList      = 348 // <channel> <exceptionmask>
	REndOfExceptList = 349 // <channel> :End of channel exception list
	RVersion         = 351 // <version>.<debuglevel> <server> :<comments>
	RWhoReply        = 352 // <channel> <user> <host> <server> <nick> ( H / G > [*] [ ( @ / + ) ] :<hopcount> <real name>
	RNameReply       = 353 // ( = / * / @ ) <channel> :[ @ / + ] <nick> *(   [ @ / + ] <nick> )
	RLinks           = 364 // <mask> <server> :<hopcount> <server info>
	REndOfLinks      = 365 // <mask> :End of LINKS list
	REndOfNames      = 366 // <channel> :End of NAMES list
	RBanList         = 367 // <channel> <banmask>
	REndOfBanList    = 368 // <channel> :End of channel ban list
	REndOfWhoWas     = 369 // <nick> :End of WHOWAS
	RInfo            = 371 // :<string>
	RMOTD            = 372 // :- <text>
	REndofInfo       = 374 // :End of INFO list
	RMOTDStart       = 375 // :- <server> Message of the day - 
	REndOfMOTD       = 376 // :End of MOTD command
	RYouAreOper      = 381 // :You are now an IRC operator
	RRehasing        = 382 // <config file> :Rehashing
	RYouAreService   = 383 // You are service <servicename>
	RTime            = 391 // <server> :<string showing server's local time>
	RUserStart       = 392 // :UserID Terminal Host
	RUsers           = 393 // :<username> <ttyline> <hostname>
	REndOfUsers      = 394 // :End of users
	RNousers         = 395 // :Nobody logged in
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
	CAdmin    = 801 // Get information about the administrator of a server.	RFC 2812
	CAway     = 802 // Set an automatic reply string for any PRIVMSG commands.	RFC 2812
	CConnect  = 803 // Request a new connection to another server immediately.	RFC 2812
	CDie      = 804 // Shutdown the server.	RFC 2812
	CError    = 805 // Report a serious or fatal error to a peer.	RFC 2812
	CInfo     = 806 // Get information describing a server.	RFC 2812
	CInvite   = 807 // Invite a user to a channel.	RFC 2812
	CIsOn     = 808 // Determine if a nickname is currently on IRC.	RFC 2812
	CJoin     = 809 // Join a channel.	RFC 2812, RFC 2813
	CKick     = 810 // Request the forced removal of a user from a channel.	RFC 2812
	CKill     = 811 // Close a client-server connection by the server which has the actual connection.	RFC 2812
	CLinks    = 812 // List all servernames which are known by the server answering the query.	RFC 2812
	CList     = 813 // List channels and their topics.	RFC 2812
	CLUsers   = 814 // Get statistics about the size of the IRC network.	RFC 2812
	CMode     = 815 // User mode.	RFC 2812
	CMOTD     = 816 // Get the Message of the Day.	RFC 2812
	CNames    = 817 // List all visible nicknames.	RFC 2812
	CNick     = 818 // Define a nickname.	RFC 2812, RFC 2813
	CNJoin    = 819 // Exchange the list of channel members for each channel between servers.	RFC 2813
	CNotice   = 820 // RFC 2812
	COper     = 821 // Obtain operator privileges.	RFC 2812
	CPart     = 822 // Leave a channel.	RFC 2812
	CPass     = 823 // Set a connection password.	RFC 2812, RFC 2813
	CPing     = 824 // Test for the presence of an active client or server.	RFC 2812
	CPong     = 825 // Reply to a PING message.	RFC 2812
	CPrivMsg  = 826 // Send private messages between users, as well as to send messages to channels.	RFC 2812
	CQuit     = 827 // Terminate the client session.	RFC 2812, RFC 2813
	CRehash   = 828 // Force the server to re-read and process its configuration file.	RFC 2812
	CRestart  = 829 // Force the server to restart itself.	RFC 2812
	CServer   = 830 // Register a new server.	RFC 2813
	CService  = 831 // Register a new service.	RFC 2812, RFC 2813
	CServList = 832 // List services currently connected to the network.	RFC 2812
	CSQuery   = 833 // RFC 2812
	CSquirt   = 834 // Disconnect a server link.	RFC 2812
	CSQuit    = 835 // Break a local or remote server link.	RFC 2813
	CStats    = 836 // Get server statistics.	RFC 2812
	CSummon   = 837 // Ask a user to join IRC.	RFC 2812
	CTime     = 838 // Get the local time from the specified server.	RFC 2812
	CTopic    = 839 // Change or view the topic of a channel.	RFC 2812
	CTrace    = 840 // Find the route to a server and information about it's peers.	RFC 2812
	CUser     = 841 // Specify the username, hostname and realname of a new user.	RFC 2812
	CUserHost = 842 // Get a list of information about upto 5 nicknames.	RFC 2812
	CUsers    = 843 // Get a list of users logged into the server.	RFC 2812
	CVersion  = 844 // Get the version of the server program.	RFC 2812
	CWAllOps  = 845 // Send a message to all currently connected users who have set the 'w' user mode.	RFC 2812
	CWho      = 846 // List a set of users.	RFC 2812
	CWhoIs    = 847 // Get information about a specific user.	RFC 2812
	CWhoWas   = 848 // Get information about a nickname which no longer exists.	RFC 2812
)

// Custom command identifiers.
const (
	CCtcpPing    = 901 // Received a CTCP ping request.
	CCtcpVersion = 902 // Received a CTCP version request.
)
