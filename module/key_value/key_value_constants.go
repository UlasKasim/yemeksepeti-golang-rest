package key_value

import "regexp"

//ModelName Model name for KeyValue
var ModelName = "KeyValue"

var keyExp = regexp.MustCompile(`/get/(?P<key>\w+)`)
