## blah

Blah is an IRC bot.


### TODO

* Verify that `proto.Client.Join` works correctly.
* Properly implement us of TLS connection.
  * Do we need a certificate and key file?
  * How do we set up the TLS.Config struct?
  * Dig into `net/http.ListenAndServeTLS` to find out.
* Determine correct use of `atomic.CompareAndSwapPointer`.
  Notably in `main.Config.SetNickname`.
* Implement command parser and handler.
* Implement user tracking through sessions.


### Dependencies

    go get github.com/jteeuwen/ini


### Usage

    go get github.com/jteeuwen/blah


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

