//go:build !windows

package swiftypes

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type LUID struct {
	LowPart  uint32
	HighPart int32
}
