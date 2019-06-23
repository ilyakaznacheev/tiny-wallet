package model

import "errors"

// ErrRowExists means that the row with the key provided already exists
var ErrRowExists = errors.New("row exists")
