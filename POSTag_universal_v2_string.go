// +build universaltagsv2

// Code generated by "stringer -tags universaltagsv2 -type=POSTag -output=POSTag_universal_v2_string.go"; DO NOT EDIT.

package lingo

import "strconv"

const _POSTag_name = "XUNKNOWN_TAGROOT_TAGADJADPADVAUXCONJDETINTJNOUNNUMPARTPRONPROPNPUNCTSCONJSYMVERBMAXTAG"

var _POSTag_index = [...]uint8{0, 1, 12, 20, 23, 26, 29, 32, 36, 39, 43, 47, 50, 54, 58, 63, 68, 73, 76, 80, 86}

func (i POSTag) String() string {
	if i >= POSTag(len(_POSTag_index)-1) {
		return "POSTag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _POSTag_name[_POSTag_index[i]:_POSTag_index[i+1]]
}
