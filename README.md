## ircb

ircb is an IRC bot. In its current form, it does not do much, besides 
connecting to a server, join the appropriate channels. In order to make it
useful, commands have to be registered to allow users to interact with it.


### TODO

* Verify that `proto.Client.Join` works correctly.
  Specifically the part where we identify ourselves with chanserv.
* Implement user tracking through sessions.


### Dependencies

    go get github.com/jteeuwen/ini


### Usage

    go get github.com/jteeuwen/ircb


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

