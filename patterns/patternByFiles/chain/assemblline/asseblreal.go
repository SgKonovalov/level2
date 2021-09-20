package assemblline

import (
	"chain/definition"
	"fmt"
)

type BodyShopCR struct {
	Next AssemblingLineCR
}

func (bscr *BodyShopCR) Execute(ccr *definition.CarCR) {
	if ccr.BodyShopIsDone {
		fmt.Println("Took body from stock")
		bscr.Next.Execute(ccr)
		return
	}
	fmt.Println("Welding body by robots")
	ccr.BodyShopIsDone = true
	bscr.Next.Execute(ccr)
}

func (bscr *BodyShopCR) SetNext(next AssemblingLineCR) {
	bscr.Next = next
}

type EngineInstallationCR struct {
	Next AssemblingLineCR
}

func (eicr *EngineInstallationCR) Execute(ccr *definition.CarCR) {
	if ccr.EngineInstalled {
		fmt.Println("Engine already in body")
		eicr.Next.Execute(ccr)
		return
	}
	fmt.Println("Installation engine by workers")
	ccr.EngineInstalled = true
	eicr.Next.Execute(ccr)
}

func (eicr *EngineInstallationCR) SetNext(next AssemblingLineCR) {
	eicr.Next = next
}

type InteriorCR struct {
	Next AssemblingLineCR
}

func (icr *InteriorCR) Execute(ccr *definition.CarCR) {
	if ccr.InteriorInstalled {
		fmt.Println("Intrerior already installed")
		icr.Next.Execute(ccr)
		return
	}
	fmt.Println("Installation interior by workers")
	ccr.InteriorInstalled = true
	icr.Next.Execute(ccr)
}

func (icr *InteriorCR) SetNext(next AssemblingLineCR) {
	icr.Next = next
}

type WheelsCR struct {
	Next AssemblingLineCR
}

func (wcr *WheelsCR) Execute(ccr *definition.CarCR) {
	if ccr.WheelsInstaled {
		fmt.Println("Wheels already installed")
		wcr.Next.Execute(ccr)
	}
	fmt.Println("Installation wheels by workers")
}

func (wcr *WheelsCR) SetNext(next AssemblingLineCR) {
	wcr.Next = next
}
