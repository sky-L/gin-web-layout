package test

import (
	"fmt"
	"testing"
)

/**
模拟组装2台电脑
--- 抽象层 ---
有显卡Card 方法display 有内存Memory 方法storage 有处理器CPU 方法calculate
--- 实现层层 ---
有 Intel因特尔公司 、产品有(显卡、内存、CPU) 有 Kingston 公司， 产品有(内存3)
有 NVIDIA 公司， 产品有(显卡)
--- 逻辑层 ---
1. 组装一台Intel系列的电脑，并运行
2. 组装一台 Intel CPU Kingston内存 NVIDIA显卡的电脑，并运行
*/

type Card interface {
	Display()
}

type Memory interface {
	Storage()
}
type Cpu interface {
	Calculate()
}

type IntelCard struct {
	Card
}

func (IntelCard) Display() {
}

type IntelMemory struct {
	Memory
}

func (IntelMemory) Storage() {
	fmt.Println("intel")
}

type InterCpu struct {
	Cpu
}

func (InterCpu) Calculate() {

}

type Computer struct {
	Cpu
	Memory
	Card
}

func NewComputer(cpu Cpu, m Memory, c Card) *Computer {
	return &Computer{
		cpu, m, c,
	}
}

func (c *Computer) Run() {
	fmt.Println("run")
}

func TestInterface(t *testing.T) {
	// 依赖倒置
	// 这玩意核心就是 拆、又名抽象
	c1 := NewComputer(InterCpu{}, IntelMemory{}, IntelCard{})
	c1.Run()
}
