// Code generated by "stringer -type ExprSourceType"; DO NOT EDIT.

package hclpack

import "strconv"

const (
	_ExprSourceType_name_0 = "ExprNative"
	_ExprSourceType_name_1 = "ExprTemplate"
)

func (i ExprSourceType) String() string {
	switch {
	case i == 78:
		return _ExprSourceType_name_0
	case i == 84:
		return _ExprSourceType_name_1
	default:
		return "ExprSourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}