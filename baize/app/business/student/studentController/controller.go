package studentController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUser, NewRole, NewPost, NewMenu, NewLogin, NewDictType, NewConfig, NewFile, NewNotice,
	NewProfile, NewDictData, NewDept, NewSse, wire.Struct(new(System), "*"))

type Student struct {
	Login *Login
}
