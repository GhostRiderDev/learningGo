package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	proto "github.com/golang/protobuf/proto"

	model "github.com/ghostriderdev/movies/metadata/pkg"
	"github.com/ghostriderdev/movies/src/gen"
)

var metadata = &model.Metadata{
	ID:          "123abcd",
	Title:       "The Movie 2",
	Description: "Sequel of the legendary The Movie",
	Director:    "Foo Bars",
}

var genMetadata = &gen.Metadata{
	Id:          "123",
	Title:       "The Movie 2",
	Description: "Sequel of the legendary The Movie",
	Director:    "Foo Bars",
}

func main() {
	jsonBytes, err := serializeJSON(metadata)

	if err != nil {
		panic(err)
	}
	xmlBytes, err := serializeXML(metadata)
	if err != nil {
		panic(err)
	}
	protoBytes, err := serializeProtoBuf(genMetadata)
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON size:\t%dB\n", len(jsonBytes))
	fmt.Printf("XML size:\t%dB\n", len(xmlBytes))
	fmt.Printf("Proto size:\t%dB\n", len(protoBytes))
}

func serializeJSON(m *model.Metadata) ([]byte, error) {
	return json.Marshal(m)
}

func serializeXML(m *model.Metadata) ([]byte, error) {
	return xml.Marshal(m)
}

func serializeProtoBuf(m *gen.Metadata) ([]byte, error) {
	return proto.Marshal(m)
}
