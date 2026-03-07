package main

import "fmt"

// 子系统1: 投影仪
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("投影仪开启")
}

func (p *Projector) Off() {
	fmt.Println("投影仪关闭")
}

func (p *Projector) SetInput(input string) {
	fmt.Printf("投影仪输入源设置为: %s\n", input)
}

// 子系统2: 音响
type SoundSystem struct{}

func (s *SoundSystem) On() {
	fmt.Println("音响开启")
}

func (s *SoundSystem) Off() {
	fmt.Println("音响关闭")
}

func (s *SoundSystem) SetVolume(level int) {
	fmt.Printf("音响音量设置为: %d\n", level)
}

// 子系统3: 播放器
type Player struct{}

func (p *Player) On() {
	fmt.Println("播放器开启")
}

func (p *Player) Off() {
	fmt.Println("播放器关闭")
}

func (p *Player) Play(movie string) {
	fmt.Printf("播放电影: %s\n", movie)
}

func (p *Player) Stop() {
	fmt.Println("停止播放")
}

// 子系统4: 灯光
type Lights struct{}

func (l *Lights) Dim(level int) {
	fmt.Printf("灯光调暗至: %d%%\n", level)
}

func (l *Lights) On() {
	fmt.Println("灯光开启")
}

// 外观类: 家庭影院
type HomeTheaterFacade struct {
	projector   *Projector
	soundSystem *SoundSystem
	player      *Player
	lights      *Lights
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		projector:   &Projector{},
		soundSystem: &SoundSystem{},
		player:      &Player{},
		lights:      &Lights{},
	}
}

// WatchMovie 观看电影（简化的高层接口）
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("\n=== 准备观看电影 ===")
	h.lights.Dim(10)
	h.projector.On()
	h.projector.SetInput("HDMI")
	h.soundSystem.On()
	h.soundSystem.SetVolume(20)
	h.player.On()
	h.player.Play(movie)
	fmt.Println("=== 开始享受电影 ===\n")
}

// EndMovie 结束电影
func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("\n=== 结束观影 ===")
	h.player.Stop()
	h.player.Off()
	h.soundSystem.Off()
	h.projector.Off()
	h.lights.On()
	fmt.Println("=== 影院系统已关闭 ===\n")
}

// ListenToMusic 听音乐
func (h *HomeTheaterFacade) ListenToMusic() {
	fmt.Println("\n=== 准备听音乐 ===")
	h.lights.Dim(30)
	h.soundSystem.On()
	h.soundSystem.SetVolume(15)
	h.player.On()
	h.player.Play("音乐播放列表")
	fmt.Println("=== 开始享受音乐 ===\n")
}

func main() {
	fmt.Println("=== 外观模式 - 家庭影院系统 ===")
	
	// 创建家庭影院外观
	homeTheater := NewHomeTheaterFacade()
	
	// 观看电影（一键操作）
	homeTheater.WatchMovie("《星际穿越》")
	
	// 模拟观影中...
	fmt.Println("... 正在观影中 ...\n")
	
	// 结束电影
	homeTheater.EndMovie()
	
	// 听音乐
	homeTheater.ListenToMusic()
	
	fmt.Println("... 正在享受音乐 ...\n")
}
