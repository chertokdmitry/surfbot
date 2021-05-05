package serializer_test

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	pb "gitlab.com/chertokdmitry/pcbook/pb"
	"gitlab.com/chertokdmitry/pcbook/sample"
	"gitlab.com/chertokdmitry/pcbook/serializer"
	"testing"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}