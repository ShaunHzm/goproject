package test

import "fmt"

// 自定义类型
type Player struct {
	Name       string
	profession string
	Level      int
}

func NewPlayer(name string, prof string, level int) *Player {
	return &Player{
		name,
		prof,
		level,
	}
}

func (p Player) attack() {
	fmt.Printf("%s发起攻击\n", p.Name)
}
