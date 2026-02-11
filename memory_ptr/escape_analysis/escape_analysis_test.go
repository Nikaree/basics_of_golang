package escape_analysis

import "testing"

func BenchmarkCreateObjects(b *testing.B) {
	b.Run("OnStack", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = CreateObjectOnStack()
		}
	})

	b.Run("OnHeap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = CreateObjectOnHeap()
		}
	})
}

func BenchmarkProcessMethods(b *testing.B) {
	objOnStack := CreateObjectOnStack()
	objOnHeap := CreateObjectOnHeap()

	b.Run("ByValue", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ProcessByValue(objOnStack)
		}
	})

	b.Run("ByPointer", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ProcessByPointer(objOnHeap)
		}
	})
}
