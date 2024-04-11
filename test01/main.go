package main

import (
	"fmt"
)

type Option func(application string)

type AllConfig struct {
	Duansg *Duansg `json:"duansg" mapstructure:"duansg" yaml:"duansg"`
}
type Duansg struct {
	Name string `json:"name" mapstructure:"name" yaml:"name"`
}
type Xd interface {
	Duansg() Duansg
}
type Caller[T Xd] func(p T) error
type CallFn[T Xd] func(fn Caller[T]) error

func MakePlugin[T Xd](super bool) CallFn[T] {
	call := func(fn Caller[T]) error {
		fmt.Println("啊？？？？？")
		return nil
	}
	return call
}

var (
	CallXd = MakePlugin[Xd](true)
)

func main() {
	fmt.Println("hello world")
	//cmd.Execute()
	//config, _ := viper.NewWithPath("config/config.yaml")
	//fmt.Println(config)
	//fmt.Println(config.Get("duansg.name"))
	//c := &AllConfig{}
	//err := config.Parse(&c)
	//if err != nil {
	//	fmt.Println("")
	//	return
	//}
	//fmt.Println(c.Duansg.Name)

	CallXd(func(base Xd) error {
		fmt.Println("plugin 晓断")
		fmt.Println(base.Duansg())
		return nil
	})

}
