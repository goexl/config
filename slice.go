package config

import (
	"encoding/json"
	"encoding/xml"

	"github.com/goexl/config/internal/constant"
	"github.com/naoina/toml"
	"gopkg.in/yaml.v3"
)

// Slice 切片，既可以兼容单个值也可以兼容数组
type Slice[T any] []T

// NewSlice 快速创建数组
func NewSlice[T any](items ...T) Slice[T] {
	return items
}

func (s *Slice[T]) Length() int {
	return len(*s)
}

func (s *Slice[T]) Capacity() int {
	return cap(*s)
}

func (s *Slice[T]) Clone() (t Slice[T]) {
	t = make(Slice[T], len(*s))
	copy(t, *s)

	return
}

func (s *Slice[T]) UnmarshalJSON(bytes []byte) (err error) {
	t := new(T)
	ts := make([]T, 0)
	start := bytes[0]
	if constant.JsonArrayStart == start {
		err = json.Unmarshal(bytes, &ts)
	} else if constant.JsonObjectStart == start {
		err = json.Unmarshal(bytes, t)
	}
	if constant.JsonObjectStart == start && nil == err {
		*s = []T{*t}
	} else if constant.JsonArrayStart == start && nil == err {
		*s = ts
	}

	return
}

func (s *Slice[T]) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	t := new(T)
	if ue := decoder.DecodeElement(t, &start); nil != ue {
		err = decoder.DecodeElement(s, &start)
	} else {
		*s = []T{*t}
	}

	return
}

func (s *Slice[T]) UnmarshalYAML(node *yaml.Node) (err error) {
	t := new(T)
	ts := make([]T, 0)
	tag := node.Tag
	if constant.YamlTagSeq == tag {
		err = node.Decode(&ts)
	} else if constant.YamlTagMap == tag {
		err = node.Decode(t)
	}
	if constant.YamlTagMap == tag && nil == err {
		*s = []T{*t}
	} else if constant.YamlTagSeq == tag && nil == err {
		*s = ts
	}

	return
}

func (s *Slice[T]) UnmarshalTOML(bytes []byte) (err error) {
	t := new(T)
	ts := make([]T, 0)
	first := bytes[0]
	start := string(bytes[0:1])
	if constant.TomlArrayStart == start {
		err = toml.Unmarshal(bytes, &ts)
	} else if constant.TomlObjectStart == first {
		err = toml.Unmarshal(bytes, t)
	}
	if constant.TomlObjectStart == first && nil == err {
		*s = []T{*t}
	} else if constant.TomlArrayStart == start && nil == err {
		*s = ts
	}

	return
}
