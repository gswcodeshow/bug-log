// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package test_ns

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Container struct {
	_tab flatbuffers.Table
}

func GetRootAsContainer(buf []byte, offset flatbuffers.UOffsetT) *Container {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Container{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsContainer(buf []byte, offset flatbuffers.UOffsetT) *Container {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Container{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Container) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Container) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Container) My(obj *MyData) *MyData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MyData)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ContainerStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ContainerAddMy(builder *flatbuffers.Builder, my flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(my), 0)
}
func ContainerEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}