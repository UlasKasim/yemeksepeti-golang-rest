package key_value

import "regexp"

//ModelName Model name for KeyValue
var ModelName = "KeyValue"

//keyExp path key finder for get endpoint
var keyExp = regexp.MustCompile(`/get/(?P<key>\w+)`)
