package main

import "testing"

func BenchmarkSerializeToJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeJSON(metadata)
	}
}

func BenchmarkSerializeToXml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeXML(metadata)
	}
}

func BenchmarkSerializeProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		serializeProtoBuf(genMetadata)
	}
}
