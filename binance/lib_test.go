package binance

import (
	"encoding/json"
	"github.com/fatih/structs"
	"testing"
)

type Foo struct {
	Bar1 int               `json:"bar_1"`
	Bar2 string            `json:"bar_2"`
	Bar3 []string          `json:"bar_3"`
	Bar4 bool              `json:"bar_4"`
	Bar5 map[string]string `json:"bar_5"`
	bar6 string
	Bar7 string `json:"-"`
}

func TestStructToMap(t *testing.T) {
	f := &Foo{
		Bar1: 1,
		Bar2: "abc",
		Bar3: nil,
		Bar4: true,
		Bar5: nil,
		bar6: "aaa",
		Bar7: "bbb",
	}

	m, err := StructToMap(f, TagJson)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", m)
}

func BenchmarkStructToMap(b *testing.B) {
	f := &Foo{
		Bar1: 1,
		Bar2: "abc",
		Bar3: nil,
		Bar4: true,
		Bar5: nil,
		bar6: "aaa",
		Bar7: "bbb",
	}
	for i := 0; i < b.N; i++ {
		//545490	      2148 ns/op	     448 B/op	      13 allocs/op
		StructToMap(f, TagJson)
	}
}

func TestStructToMap2(t *testing.T) {
	f := &Foo{
		Bar1: 1,
		Bar2: "abc",
		Bar3: nil,
		Bar4: true,
		Bar5: nil,
		bar6: "aaa",
		Bar7: "bbb",
	}

	m := structs.Map(f)
	t.Logf("%+v", m)
}

func BenchmarkStructToMap2(b *testing.B) {
	f := &Foo{
		Bar1: 1,
		Bar2: "abc",
		Bar3: nil,
		Bar4: true,
		Bar5: nil,
		bar6: "aaa",
		Bar7: "bbb",
	}
	for i := 0; i < b.N; i++ {
		//297464	      3619 ns/op	    2344 B/op	      34 allocs/op
		structs.Map(f) //本质仍然是使用反射
	}
}

func TestStructToMap3(t *testing.T) {
	f := &Foo{
		Bar1: 1,
		Bar2: "abc",
		Bar3: nil,
		Bar4: true,
		Bar5: nil,
		bar6: "aaa",
		Bar7: "bbb",
	}

	m := make(map[string]any)

	marshal, err := json.Marshal(f)
	if err != nil {
		t.Error(err)
		return
	}

	err = json.Unmarshal(marshal, &m)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", m)
}

func BenchmarkStructToMap3(b *testing.B) {

	f := &Foo{
		Bar1: 1,
		Bar2: "abc",
		Bar3: nil,
		Bar4: true,
		Bar5: nil,
		bar6: "aaa",
		Bar7: "bbb",
	}

	m := make(map[string]any)

	for i := 0; i < b.N; i++ {
		//311366	      3782 ns/op	     368 B/op	      17 allocs/op
		marshal, _ := json.Marshal(f)
		json.Unmarshal(marshal, &m) //依然不快，且有功能问题
	}
}
