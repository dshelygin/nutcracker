package characters

import (
	"context"
	"nutcracker/domain/data"
	"sync"
	"testing"
	"time"
)

func TestMaryGetState(t *testing.T) {
	var mary Mary
	var maryState HappinessState

	states := map[HappinessIndex]HappinessState{
		-1000: IsCrying,
		-100: IsCrying,
		-99: IsSad,
		-50: IsSad,
		-10: IsSad,
		0: IsNeutral,
		10: IsFine,
		50: IsFine,
		99: IsFine,
		100: IsHappy,
		1000: IsHappy,
	}

	for index, expectedState := range states {
		mary = Mary{happinessIndex: index}
		maryState = mary.GetState()
		if maryState != expectedState {
			t.Errorf("expected: %d, actual: %d, index: %d", expectedState, maryState, index)
		}
	}
}

func TestMaryChangeFaceColor(t *testing.T) {
	var mary Mary
	var maryFaceColor HumanFaceColorVariation

	states := map[HappinessIndex]HumanFaceColorVariation{
		-1000: RedHumanFaceColor,
		-99: PaleHumanFaceColor,
		0: FleshHumanFaceColor,
		10: RosyHumanFaceColor,
		100: RosyHumanFaceColor,
	}

	for index, expectedFaceColor := range states {
		mary = Mary{happinessIndex: index}
		mary.changeFaceColor()
		maryFaceColor = mary.GetFaceColor()

		if maryFaceColor != expectedFaceColor {
			t.Errorf("expected: %d, actual: %d, index: %d", expectedFaceColor, maryFaceColor, index)
		}
	}
}

func TestGetBadNewsForMary(t *testing.T) {
	incoming := make(chan data.Message)
	ctx, isOver := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	mary := NewMary(ctx, incoming, &wg)

	incoming <- data.MessageMouseStillUnCatch

	time.Sleep(time.Second * 2)

	maryState := mary.GetState()
	if  maryState != IsSad {
		t.Errorf("expected state: %d, actual: %d", IsSad, maryState)
	}

	maryFaceColor := mary.GetFaceColor()
	if  maryFaceColor != PaleHumanFaceColor {
		t.Errorf("expected face color: %d, actual: %d", PaleHumanFaceColor, maryFaceColor)
	}

	isOver()
	wg.Wait()
}
