package definition

import "fmt"

type EngineOneC struct {
	IsStarting bool
}

type EngineTwoC struct {
	IsStarting bool
}

func (ec *EngineOneC) StartEngineC() {
	ec.IsStarting = true
	fmt.Println("Engine one been started")
}

func (ec *EngineOneC) StopEngineC() {
	ec.IsStarting = false
	fmt.Println("Engine one been stopped")
}

func (ec *EngineTwoC) StartEngineC() {
	ec.IsStarting = true
	fmt.Println("Engine two been started")
}

func (ec *EngineTwoC) StopEngineC() {
	ec.IsStarting = false
	fmt.Println("Engine two been stopped")
}
