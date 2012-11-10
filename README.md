## ircb

ircb is an IRC bot. In its current form, it does not do much besides 
connecting to a server and joining the appropriate channels.
In order to make it useful, commands have to be registered to allow
users to interact with it.

A bot can be configured through an external .ini file. For an example
of one, refer to the `config.example.ini` file in the root of this repo.


### TODO

* Implement user tracking through sessions.
* Implement user login.
  * Deny command execution if `Command.Restricted == true` and current
    user is not authorized.


### Dependencies

    go get github.com/jteeuwen/ini


### Usage

    go get github.com/jteeuwen/ircb


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

