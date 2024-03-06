package handlers

var claim MyCustomClaims

func NewIdentity(c MyCustomClaims) {
	claim = c
}

func Identity() int64 {
	return claim.UserID
}