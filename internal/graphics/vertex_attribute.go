package graphics

type VertexAttributes struct {
	List   []VertexAttribute
	Stride int32
}

type VertexAttribute struct {
	Size       int32
	Type       uint32
	Normalized bool
	Offset     uintptr
}
