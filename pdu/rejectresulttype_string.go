// Code generated by "stringer -type RejectResultType"; DO NOT EDIT.

package pdu

import "fmt"

const _RejectResultType_name = "ResultRejectedPermanentResultRejectedTransient"

var _RejectResultType_index = [...]uint8{0, 23, 46}

func (i RejectResultType) String() string {
	i -= 1
	if i >= RejectResultType(len(_RejectResultType_index)-1) {
		return fmt.Sprintf("RejectResultType(%d)", i+1)
	}
	return _RejectResultType_name[_RejectResultType_index[i]:_RejectResultType_index[i+1]]
}
