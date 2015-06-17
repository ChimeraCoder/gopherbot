// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package proto

import (
	"bytes"
	"github.com/chimeracoder/gopherbot/irc"
	"testing"
)

func TestRaw(t *testing.T) {
	const want = "abc def: 123\n"
	var have bytes.Buffer

	a := "def"
	b := 123
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Raw("abc %s: %d", a, b); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestPrivMsg(t *testing.T) {
	const want = "PRIVMSG bob :I like cake\n"
	var have bytes.Buffer

	a := "bob"
	b := "I like cake"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.PrivMsg(a, b); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestNotice(t *testing.T) {
	const want = "NOTICE bob :I like cake\n"
	var have bytes.Buffer

	a := "bob"
	b := "I like cake"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Notice(a, b); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestQuit1(t *testing.T) {
	const want = "QUIT I like cake\n"
	var have bytes.Buffer

	a := "I like cake"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Quit(a); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestQuit2(t *testing.T) {
	const want = "QUIT\n"
	var have bytes.Buffer

	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Quit(""); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestPong(t *testing.T) {
	const want = "PONG 1234567890\n"
	var have bytes.Buffer

	a := "1234567890"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Pong(a); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestNick1(t *testing.T) {
	const want = "NICK bob\n"
	var have bytes.Buffer

	a := "bob"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Nick(a, ""); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestNick2(t *testing.T) {
	const want = "NICK bob\nPRIVMSG nickserv :IDENTIFY 12345\n"
	var have bytes.Buffer

	a := "bob"
	b := "12345"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Nick(a, b); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestJoin(t *testing.T) {
	const want = `CS INVITE #test1
JOIN #test1
CS INVITE #test2
JOIN #test2 abc
CS INVITE #test3
JOIN #test3
PRIVMSG chanserv :IDENTIFY def
CS INVITE #test4
JOIN #test4 abc
PRIVMSG chanserv :IDENTIFY def
`
	var have bytes.Buffer

	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})
	chans := []*irc.Channel{
		&irc.Channel{"#test1", "", ""},
		&irc.Channel{"#test2", "abc", ""},
		&irc.Channel{"#test3", "", "def"},
		&irc.Channel{"#test4", "abc", "def"},
	}

	if err := c.Join(chans); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestPart(t *testing.T) {
	const want = `PART #test1
PART #test2
PART #test3
PART #test4
`
	var have bytes.Buffer

	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})
	chans := []*irc.Channel{
		&irc.Channel{"#test1", "", ""},
		&irc.Channel{"#test2", "abc", ""},
		&irc.Channel{"#test3", "", "def"},
		&irc.Channel{"#test4", "abc", "def"},
	}

	if err := c.Part(chans); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestMode(t *testing.T) {
	const want = `MODE #test1 +m
MODE bob +b for no reason.
`
	var have bytes.Buffer

	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})
	list := [][3]string{
		{"#test1", "+m", ""},
		{"bob", "+b", "for no reason."},
	}

	for _, m := range list {
		if err := c.Mode(m[0], m[1], m[2]); err != nil {
			t.Fatal(err)
		}
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestTopic(t *testing.T) {
	const want = `TOPIC #test1 :Topic all up in this
`
	var have bytes.Buffer

	a := "#test1"
	b := "Topic all up in this"
	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})

	if err := c.Topic(a, b); err != nil {
		t.Fatal(err)
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestInvite(t *testing.T) {
	const want = `:bob INVITE #test1 :
:bob INVITE #test2 :Just because.
`
	var have bytes.Buffer

	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})
	list := [][3]string{
		{"bob", "#test1", ""},
		{"bob", "#test2", "Just because."},
	}

	for _, i := range list {
		if err := c.Invite(i[0], i[1], i[2]); err != nil {
			t.Fatal(err)
		}
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}

func TestKick(t *testing.T) {
	const want = `KICK #test1 bob
KICK #test2 bob :Just because.
`
	var have bytes.Buffer

	c := NewClient(func(d []byte) error {
		_, err := have.Write(d)
		return err
	})
	list := [][3]string{
		{"#test1", "bob", ""},
		{"#test2", "bob", "Just because."},
	}

	for _, i := range list {
		if err := c.Kick(i[0], i[1], i[2]); err != nil {
			t.Fatal(err)
		}
	}

	if have.String() != want {
		t.Fatalf("Want: %q\nHave: %q", want, have.String())
	}
}
