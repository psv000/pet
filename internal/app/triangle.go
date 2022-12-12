package app

//
//type triangle struct {
//	graphics.Primitive
//}
//
//func NewTriangle() graphics.Primitive {
//	t := &triangle{
//		Primitive: graphics.NewPrimitive(resources.DotShader, []float32{
//			-0.5, -0.5, 0.0,
//			0.5, -0.5, 0.0,
//			0., 0.5, 0.0,
//		},
//			[]uint32{0, 1, 2},
//			assets.BasicVertexShader,
//			assets.AlbedoFragmentShader),
//	}
//	t.color = [4]float32{1., 1., 1., 1.}
//	return t
//}
//
//func (t *triangle) SetSize(size mgl32.Vec3) {
//
//}
//
//func (t *triangle) Size() mgl32.Vec3 {
//	return mgl32.Vec3{}
//}
//
//func (t *triangle) Render(dt float64, project, camera mgl32.Mat4) {
//	t.primitive.render(dt, project, camera)
//}
