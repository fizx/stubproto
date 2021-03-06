package proto

const Marshal = 0
const GoGoProtoPackageIsVersion3 = 0

type Message interface {
}

type InternalMessageInfo interface {
	DiscardUnknown(interface{})
	Merge(interface{}, interface{})
	Marshal(interface{}, interface{}, interface{}) ([]byte, error)
}

func RegisterFile(string, []byte) {}

func RegisterType(interface{}, string) {}

func CompactTextString(interface{}) string {
	return "nope"
}
