// Code generated by "stringer -type=UnOp"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[UnOpUnknown-0]
}

const _UnOp_name = "UnOpUnknown"

var _UnOp_index = [...]uint8{0, 11}

func (i UnOp) String() string {
	if i >= UnOp(len(_UnOp_index)-1) {
		return "UnOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _UnOp_name[_UnOp_index[i]:_UnOp_index[i+1]]
}