package main

import "fmt"

// Command 命令接口
type Command interface {
	Execute()
	Undo()
}

// Light 电灯（接收者）
type Light struct {
	location string
	isOn     bool
}

func NewLight(location string) *Light {
	return &Light{location: location, isOn: false}
}

func (l *Light) On() {
	l.isOn = true
	fmt.Printf("%s 的灯已打开\n", l.location)
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Printf("%s 的灯已关闭\n", l.location)
}

// LightOnCommand 开灯命令
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

// LightOffCommand 关灯命令
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

// AirConditioner 空调（接收者）
type AirConditioner struct {
	isOn        bool
	temperature int
}

func NewAirConditioner() *AirConditioner {
	return &AirConditioner{isOn: false, temperature: 26}
}

func (a *AirConditioner) On() {
	a.isOn = true
	fmt.Println("空调已打开")
}

func (a *AirConditioner) Off() {
	a.isOn = false
	fmt.Println("空调已关闭")
}

func (a *AirConditioner) SetTemperature(temp int) {
	a.temperature = temp
	fmt.Printf("空调温度设置为: %d°C\n", temp)
}

// ACOnCommand 空调开启命令
type ACOnCommand struct {
	ac *AirConditioner
}

func NewACOnCommand(ac *AirConditioner) *ACOnCommand {
	return &ACOnCommand{ac: ac}
}

func (c *ACOnCommand) Execute() {
	c.ac.On()
	c.ac.SetTemperature(24)
}

func (c *ACOnCommand) Undo() {
	c.ac.Off()
}

// RemoteControl 遥控器（调用者）
type RemoteControl struct {
	commands    []Command
	undoCommand Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		commands: make([]Command, 0),
	}
}

func (r *RemoteControl) SetCommand(command Command) {
	r.commands = append(r.commands, command)
}

func (r *RemoteControl) PressButton(slot int) {
	if slot >= 0 && slot < len(r.commands) {
		r.commands[slot].Execute()
		r.undoCommand = r.commands[slot]
	}
}

func (r *RemoteControl) PressUndo() {
	if r.undoCommand != nil {
		fmt.Println("\n执行撤销操作:")
		r.undoCommand.Undo()
	}
}

func main() {
	fmt.Println("=== 命令模式 - 智能家居控制系统 ===\n")
	
	// 创建接收者
	livingRoomLight := NewLight("客厅")
	bedroomLight := NewLight("卧室")
	ac := NewAirConditioner()
	
	// 创建命令
	livingRoomLightOn := NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := NewLightOffCommand(livingRoomLight)
	bedroomLightOn := NewLightOnCommand(bedroomLight)
	acOn := NewACOnCommand(ac)
	
	// 创建遥控器
	remote := NewRemoteControl()
	remote.SetCommand(livingRoomLightOn)  // 按钮0
	remote.SetCommand(livingRoomLightOff) // 按钮1
	remote.SetCommand(bedroomLightOn)     // 按钮2
	remote.SetCommand(acOn)               // 按钮3
	
	// 测试命令
	fmt.Println("按下按钮0:")
	remote.PressButton(0)
	
	fmt.Println("\n按下按钮2:")
	remote.PressButton(2)
	
	fmt.Println("\n按下按钮3:")
	remote.PressButton(3)
	
	// 撤销最后一个命令
	remote.PressUndo()
	
	fmt.Println("\n按下按钮1:")
	remote.PressButton(1)
	
	// 再次撤销
	remote.PressUndo()
}
