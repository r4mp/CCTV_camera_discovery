package communication

import (
//"log"
)

type Auth struct{}

type Header struct {
	Security //`xml:"s:mustUnderstand=1, attr"`
}

type Security struct {
	UsernameToken
}

type UsernameToken struct {
	Username string
	Password string
	Nonce    string
	Created  string
}
