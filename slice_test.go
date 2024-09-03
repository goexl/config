package config_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/goexl/config"
	"github.com/naoina/toml"
	"gopkg.in/yaml.v3"
)

type UserSlice struct {
	User     config.Slice[User]  `json:"user,omitempty" yaml:"user,omitempty" toml:"user,omitempty"`
	Users    config.Slice[User]  `json:"users,omitempty" yaml:"users,omitempty" toml:"users,omitempty"`
	UserPtr  config.Slice[*User] `json:"user_ptr,omitempty" yaml:"user_ptr,omitempty" toml:"user_ptr,omitempty"`
	UsersPtr config.Slice[*User] `json:"users_ptr,omitempty" yaml:"users_ptr,omitempty" toml:"users_ptr,omitempty"`
}

func receiveSlice[T any](data config.Slice[T]) {
	fmt.Println(data)
}

func receiveIntSlice(data []int) {
	fmt.Println(data)
}

func TestNewSlice(t *testing.T) {
	receiveSlice(config.NewSlice(1))
	receiveIntSlice(config.NewSlice(2))
}

func TestSliceJSON(t *testing.T) {
	slice := new(UserSlice)
	if bytes, rfe := os.ReadFile("testdata/user_slice.json"); nil != rfe {
		t.Errorf("读取文件内容出错，%v", rfe)
	} else if ue := json.Unmarshal(bytes, slice); nil != ue {
		t.Errorf("反序列化YAML出错，%v", ue)
	} else {
		checkUserSlice(t, slice)
	}
}

func TestSliceYAML(t *testing.T) {
	slice := new(UserSlice)
	if bytes, rfe := os.ReadFile("testdata/user_slice.yml"); nil != rfe {
		t.Errorf("读取文件内容出错，%v", rfe)
	} else if ue := yaml.Unmarshal(bytes, slice); nil != ue {
		t.Errorf("反序列化YAML出错，%v", ue)
	} else {
		checkUserSlice(t, slice)
	}
}

func TestSliceTOML(t *testing.T) {
	slice := new(UserSlice)
	if bytes, rfe := os.ReadFile("testdata/user_slice.toml"); nil != rfe {
		t.Errorf("读取文件内容出错，%v", rfe)
	} else if ue := toml.Unmarshal(bytes, slice); nil != ue {
		t.Errorf("反序列化TOML出错，%v", ue)
	} else {
		checkUserSlice(t, slice)
	}
}

func checkUserSlice(t *testing.T, slice *UserSlice) {
	if 1 != slice.User.Length() {
		t.Error("User字段反序列化后不是只有一个元素")
	} else if "storezhang" != slice.User[0].Name && 39 != slice.User[0].Age {
		t.Error("User字段反序列化后字段值不正确")
	} else if 1 != slice.UserPtr.Length() {
		t.Error("User指针字段反序列化后不是只有一个元素")
	} else if "store" != slice.UserPtr[0].Name && 19 != slice.UserPtr[0].Age {
		t.Error("User指针字段反序列化后字段值不正确")
	} else if 2 != slice.Users.Length() {
		t.Error("Users字段反序列化后字段长度正确")
	} else if 3 != slice.UsersPtr.Length() {
		t.Error("Users指针字段反序列化后字段长度正确")
	}
}
