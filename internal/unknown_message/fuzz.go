package unknown_message

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jhump/protoreflect/dynamic"
)

func Fuzz(data []byte) int {
	dyn, _ := dynamic.AsDynamicMessage(&empty.Empty{})
	err := proto.Unmarshal(data, dyn)
	if err != nil {
		return 0
	}

	unknownMessage, err := GenerateDescriptorForUnknownMessage(dyn).Build()
	if err != nil {
		return 0
	}
	dyn = dynamic.NewMessage(unknownMessage)
	// now unmarshal again using the new generated message type
	err = proto.Unmarshal(data, dyn)
	if err != nil {
		panic(fmt.Sprint("failed to unmarshal data the second time", err))
	}

	return 1
}
