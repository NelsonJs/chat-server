package test

import (
	"chat/test/animal"
	"fmt"
)

type Animal struct {
	Num int64
}

type Pig struct {
	Num int64
}


func (an Animal) Fly()  {
	fmt.Println("动物飞",an.Num)
}

func (an Pig) Fly()  {
	fmt.Println("猪飞",an.Num)
}

func Test(bird,bird1 animal.Bird) {
	bird.Fly()
	bird1.Fly()
}
