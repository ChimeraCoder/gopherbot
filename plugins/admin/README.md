## Admin

This plugin offers a rudimentary set of administrative commands for the bot.

### Commands

* `quit`: Unconditionally quits the bot,
* `join <channel> [<key> [<chanservpass>]]`: Unconditionally makes the bot
  join the given channel. The channel key and ChanServ password are optional.
* `leave [<channel>]`: Unconditionally makes the bot leave the given channel.
  The channel parameter is optional. When omitted, it refers to the channel
  from which the command was issued. If the command has no channel parameter and
  it was issued from outside a channel, the command is ignored.

