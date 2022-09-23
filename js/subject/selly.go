package jssubject

import (
	"fmt"
)

// getSellyValue ...
func getSellyValue(val string) string {
	return fmt.Sprintf("%s.%s.%s", root, prefixes.Selly, val)
}

// Selly ...
var Selly = struct {
	Stream           string
	PushNotification string
}{
	Stream:           prefixes.Selly,
	PushNotification: getSellyValue("push_notifications"),
}
