// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package proto

// ProtoId is a numeric identifier for the protocol message type.
type ProtoId uint16

// Unknown message type.
const PIDUnknown ProtoId = 0

// Known protocol message types according to IRC spec.
const (
	PIDWelcome           ProtoId = 1
	PIDTraceLink         ProtoId = 200
	PIDTraceConnecting   ProtoId = 201
	PIDTraceHandshake    ProtoId = 202
	PIDTraceUnknown      ProtoId = 203
	PIDTraceOperator     ProtoId = 204
	PIDTraceUser         ProtoId = 205
	PIDTraceServer       ProtoId = 206
	PIDTraceNewtype      ProtoId = 208
	PIDStatsLinkinfo     ProtoId = 211
	PIDStatsCommands     ProtoId = 212
	PIDStatsCLine        ProtoId = 213
	PIDStatsNLine        ProtoId = 214
	PIDStatsILine        ProtoId = 215
	PIDStatsKLine        ProtoId = 216
	PIDStatsYLine        ProtoId = 218
	PIDEndOfStats        ProtoId = 219
	PIDUModeIs           ProtoId = 221
	PIDStatsLLine        ProtoId = 241
	PIDtatsUptime        ProtoId = 242
	PIDStatsOLine        ProtoId = 243
	PIDStatsHLine        ProtoId = 244
	PIDConnectionCount   ProtoId = 250
	PIDLUserClient       ProtoId = 251
	PIDLUserOp           ProtoId = 252
	PIDLUserUnknown      ProtoId = 253
	PIDLUserChannels     ProtoId = 254
	PIDLUserMe           ProtoId = 255
	PIDAdminMe           ProtoId = 256
	PIDAdminLoc1         ProtoId = 257
	PIDAdminLoc2         ProtoId = 258
	PIDAdminEmail        ProtoId = 259
	PIDTraceLog          ProtoId = 261
	PIDNone              ProtoId = 300
	PIDAway              ProtoId = 301
	PIDUserhost          ProtoId = 302
	PIDJson              ProtoId = 303
	PIDUnaway            ProtoId = 305
	PIDNoaway            ProtoId = 306
	PIDWhoIsUser         ProtoId = 311
	PIDWhoIsServer       ProtoId = 312
	PIDWhoIsOperator     ProtoId = 313
	PIDWhoWasUser        ProtoId = 314
	PIDEndOfWho          ProtoId = 315
	PIDWhoIsIdle         ProtoId = 317
	PIDEndOfWhoIs        ProtoId = 318
	PIDWhoisChannels     ProtoId = 319
	PIDListStart         ProtoId = 321
	PIDList              ProtoId = 322
	PIDListEnd           ProtoId = 323
	PIDChannelModeIs     ProtoId = 324
	PIDNoTopic           ProtoId = 331
	PIDTopic             ProtoId = 332
	PIDNameListBegin     ProtoId = 333
	PIDInviting          ProtoId = 341
	PIDSummoning         ProtoId = 342
	PIDVersion           ProtoId = 351
	PIDWhoReply          ProtoId = 353
	PIDNameReply         ProtoId = 353
	PIDLinks             ProtoId = 364
	PIDEndOfLinks        ProtoId = 365
	PIDEndOfNames        ProtoId = 366
	PIDBanList           ProtoId = 367
	PIDEndOfBanList      ProtoId = 368
	PIDEndOfWhoWas       ProtoId = 369
	PIDInfo              ProtoId = 371
	PIDMOTD              ProtoId = 372
	PIDEndOfInfo         ProtoId = 374
	PIDMOTDStart         ProtoId = 375
	PIDEndOfMOTD         ProtoId = 376
	PIDYoureOper         ProtoId = 381
	PIDRehashing         ProtoId = 382
	PIDTime              ProtoId = 391
	PIDUserStart         ProtoId = 392
	PIDUsers             ProtoId = 393
	PIDEndOfUsers        ProtoId = 394
	PIDNoUsers           ProtoId = 395
	PIDNoSuchNick        ProtoId = 401
	PIDNoSuchServer      ProtoId = 402
	PIDNoSuchChannel     ProtoId = 403
	PIDCannotSendToChan  ProtoId = 404
	PIDTooManyChannels   ProtoId = 405
	PIDWasNoSuchNick     ProtoId = 406
	PIDTooManytargets    ProtoId = 407
	PIDNoOrigin          ProtoId = 409
	PIDNoRecipient       ProtoId = 411
	PIDNoTextToSend      ProtoId = 412
	PIDNoTopLevel        ProtoId = 413
	PIDWildeTopLevel     ProtoId = 414
	PIDUnknownCommand    ProtoId = 421
	PIDNoMOTD            ProtoId = 422
	PIDNoAdminInfo       ProtoId = 423
	PIDFileError         ProtoId = 424
	PIDNoNicknameGiven   ProtoId = 431
	PIDErroneousNickname ProtoId = 432
	PIDNickInUse         ProtoId = 433
	PIDNickCollision     ProtoId = 436
	PIDUserNotInChannel  ProtoId = 441
	PIDNotOnChannel      ProtoId = 442
	PIDUserOnChannel     ProtoId = 443
	PIDNoLogin           ProtoId = 444
	PIDSummonDisabled    ProtoId = 445
	PIDUserDisabled      ProtoId = 446
	PIDNotRegistered     ProtoId = 451
	PIDNeedMoreParams    ProtoId = 461
	PIDAlreadyRegistered ProtoId = 462
	PIDNoPermForHost     ProtoId = 463
	PIDPasswordMismatch  ProtoId = 464
	PIDYouAreBannedCreep ProtoId = 465
	PIDKeyset            ProtoId = 467
	PIDChannelIsFull     ProtoId = 471
	PIDUnknownMode       ProtoId = 472
	PIDInviteOnlyChannel ProtoId = 473
	PIDBannedFromChannel ProtoId = 474
	PIDBadChannelKey     ProtoId = 475
	PIDNoPrivileges      ProtoId = 481
	PIDChanOPrivsNeeded  ProtoId = 482
	PIDCannotKillServer  ProtoId = 483
	PIDNoOperHost        ProtoId = 491
	PIDUModeUnknownFlag  ProtoId = 501
	PIDUsersDontMatch    ProtoId = 502
)

// Custom protocol Ids: These have no direct equivalent numerical value in the
// protocol spec, so we define them ourselves for easier message handling.
const (
	PIDNotice ProtoId = 0xffff - iota
	PIDPrivMsg
	PIDQuit
	PIDJoin
	PIDPart
	PIDKick
	PIDNick
	PIDPing
	PIDError
	PIDMode
)

func (p ProtoId) String() string {
	switch p {
	case PIDNotice:
		return "Notice"
	case PIDPrivMsg:
		return "Private message"
	case PIDQuit:
		return "Quit"
	case PIDJoin:
		return "Join"
	case PIDPart:
		return "Part"
	case PIDKick:
		return "Kick"
	case PIDNick:
		return "Nick"
	case PIDPing:
		return "Ping"
	case PIDError:
		return "Error"
	case PIDMode:
		return "Mode"

	case 1:
		return "Welcome"

	case 200:
		return "Trace link"
	case 201:
		return "Trace connecting"
	case 202:
		return "Trace handshake"
	case 203:
		return "Trace unknown"
	case 204:
		return "Trace operator"
	case 205:
		return "Trace user"
	case 206:
		return "Trace server"
	case 208:
		return "Trace new type"
	case 211:
		return "Stats link info"
	case 212:
		return "Stats commands"
	case 213:
		return "Stats c-line"
	case 214:
		return "Stats n-line"
	case 215:
		return "Stats i-line"
	case 216:
		return "Stats k-line"
	case 218:
		return "Stats y-line"
	case 219:
		return "End of stats"
	case 221:
		return "Umode is"
	case 241:
		return "Stats l-line"
	case 242:
		return "Stats uptime"
	case 243:
		return "Stats o-line"
	case 244:
		return "Stats h-line"
	case 250:
		return "Connection count"
	case 251:
		return "l-user client"
	case 252:
		return "l-user op"
	case 253:
		return "l-user unknown"
	case 254:
		return "l-user channels"
	case 255:
		return "l-user me"
	case 256:
		return "Admin me"
	case 257:
		return "Admin loc1"
	case 258:
		return "Admin loc2"
	case 259:
		return "Admin e-mail"
	case 261:
		return "Trace log"

	case 300:
		return "None"
	case 301:
		return "Away"
	case 302:
		return "Userhost"
	case 303:
		return "JSON"
	case 305:
		return "Un-away"
	case 306:
		return "No away"
	case 311:
		return "Whois user"
	case 312:
		return "Whois server"
	case 313:
		return "Whois operator"
	case 314:
		return "Whowas user"
	case 315:
		return "End of who"
	case 317:
		return "Whois idle"
	case 318:
		return "End of whois"
	case 319:
		return "Whois channels"
	case 321:
		return "List start"
	case 322:
		return "List"
	case 323:
		return "List end"
	case 324:
		return "Channel mode is"
	case 331:
		return "No topic"
	case 332:
		return "Topic"
	case 333:
		return "Name list begin"
	case 341:
		return "Inviting"
	case 342:
		return "Summoning"
	case 351:
		return "Version"
	case 352:
		return "Who reply"
	case 353:
		return "Name reply"
	case 364:
		return "Links"
	case 365:
		return "End of links"
	case 366:
		return "End of names"
	case 367:
		return "Banlist"
	case 368:
		return "End of ban list"
	case 369:
		return "End of whowas"
	case 371:
		return "Info"
	case 374:
		return "End of info"
	case 372:
		return "MOTD"
	case 375:
		return "MOTD start"
	case 376:
		return "End of MOTD"
	case 381:
		return "You are oper"
	case 382:
		return "Rehashing"
	case 391:
		return "Time"
	case 392:
		return "Users start"
	case 393:
		return "Users"
	case 394:
		return "End of users"
	case 395:
		return "No users"

	case 401:
		return "No such nick"
	case 402:
		return "No such server"
	case 403:
		return "No such channel"
	case 404:
		return "Cannot send to chan"
	case 405:
		return "Too many channels"
	case 406:
		return "Was no such nick"
	case 407:
		return "Too many targets"
	case 409:
		return "No origin"
	case 411:
		return "No recipient"
	case 412:
		return "No text to send"
	case 413:
		return "No top-level"
	case 414:
		return "Wild top-level"
	case 421:
		return "Unknown command"
	case 422:
		return "No MOTD"
	case 423:
		return "No admin info"
	case 424:
		return "File error"
	case 431:
		return "No nickname given"
	case 432:
		return "Erroneous nickname"
	case 433:
		return "Nick in use"
	case 436:
		return "Nick collision"
	case 441:
		return "User not in channel"
	case 442:
		return "Not in channel"
	case 443:
		return "User in channel"
	case 444:
		return "No login"
	case 445:
		return "Summon disabled"
	case 446:
		return "User disabled"
	case 451:
		return "Not registered"
	case 461:
		return "Need more paramaters"
	case 462:
		return "Already registered"
	case 463:
		return "no perm for host"
	case 464:
		return "Password mismatch"
	case 465:
		return "You're banned, creep!"
	case 467:
		return "Keyset"
	case 471:
		return "Channel is full"
	case 472:
		return "Unknown mode"
	case 473:
		return "Invite-only channel"
	case 474:
		return "Banned from channel"
	case 475:
		return "Bad channel key"
	case 481:
		return "No privileges"
	case 482:
		return "Channel ooperator privileges needed"
	case 483:
		return "Can't kill server"
	case 491:
		return "No oper host"

	case 501:
		return "Umode unknown flag"
	case 502:
		return "Users don't match"
	}

	return "Unknown"
}
