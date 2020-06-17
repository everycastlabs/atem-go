// Code generated by "stringer -type=ExternalPortType,PortType -linecomment -output=inputs_string.go"; DO NOT EDIT.

package models

import "strconv"

const _ExternalPortType_name = "InternalSDIHDMICompositeComponentSVideo"

var _ExternalPortType_index = [...]uint8{0, 8, 11, 15, 24, 33, 39}

func (i ExternalPortType) String() string {
	if i >= ExternalPortType(len(_ExternalPortType_index)-1) {
		return "ExternalPortType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExternalPortType_name[_ExternalPortType_index[i]:_ExternalPortType_index[i+1]]
}

const (
	_PortType_name_0 = "ExternalBlackColor BarsColor GeneratorMedia Player FillMedia Player KeySuperSource"
	_PortType_name_1 = "ME OutputAuxilaryMask"
)

var (
	_PortType_index_0 = [...]uint8{0, 8, 13, 23, 38, 55, 71, 82}
	_PortType_index_1 = [...]uint8{0, 9, 17, 21}
)

func (i PortType) String() string {
	switch {
	case 0 <= i && i <= 6:
		return _PortType_name_0[_PortType_index_0[i]:_PortType_index_0[i+1]]
	case 128 <= i && i <= 130:
		i -= 128
		return _PortType_name_1[_PortType_index_1[i]:_PortType_index_1[i+1]]
	default:
		return "PortType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
