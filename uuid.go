package uuid

import uid "github.com/satori/go.uuid"

func Uuid() string {
	return uid.NewV4().String()
}
