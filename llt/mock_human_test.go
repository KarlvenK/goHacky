package llt

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/KarlvenK/goHacky/llt/mock"
	"github.com/golang/mock/gomock"
)

func Test_Mock_Human(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHuman := mock.NewMockHuman(ctrl)
	mockHuman.EXPECT().Speak().DoAndReturn(func() string {
		return "ohhhhhhhhhhhh"
	}).Times(2)

	output := mockHuman.Speak()
	t.Logf("speak: %s", output)
	output2 := mockHuman.Speak()
	t.Logf("speak2: %s", output2)

	mockHuman.EXPECT().Get(gomock.Eq("shit"), gomock.Any()).Return("you fucking idiot")
	mockHuman.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(_ string, _ int) string {
		fmt.Println("==========!!!!+=======")
		return "ass"
	})
	resp := mockHuman.Get("shit", 0)
	t.Log(resp)
	resp = mockHuman.Get("shits", 1)
	t.Log(resp)
}

func TestSplit(t *testing.T) {
	Convey("basic", t, func() {
		var (
			s      = "foobarfoofoobar"
			sep    = "bar"
			expect = []string{"foo", "foofoo", ""}
		)
		got := Split(s, sep)
		So(got, ShouldResemble, expect)
	})
}
