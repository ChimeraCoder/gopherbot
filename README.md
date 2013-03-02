## ircb

ircb is an IRC bot. In its bare form, it does not do much besides 
connecting to a server and joining the appropriate channels.
In order to make it useful, commands have to be registered to allow
users to interact with it.

A bot can be configured through an external .ini file. For an example
of one, refer to the `config.example.ini` file in the root of this repo.


### Dependencies

    go get github.com/jteeuwen/ini


### Usage

    go get github.com/jteeuwen/ircb

The bot is launched with the `-p` flag. This flag expects an existing directory
path, contianing the bot profile, as well as any optional plugin configurations.
Its directory structure is as follows:

	[$path]
	   |
	   |- [plugins]
	   |    |
	   |    |- [foo]
	   |    |    |- config.ini
	   |    |
	   |    |- [bar]
	   |    |    |- config.ini
	   |
	   |- config.ini

For example:

	$ ircb -p ~/.ircb/someprofile


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

