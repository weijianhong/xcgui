package wutil_test

import (
	"fmt"
	"github.com/weijianhong/xcgui/app"
	"github.com/weijianhong/xcgui/tf"
	"github.com/weijianhong/xcgui/wapi/wutil"
	"github.com/weijianhong/xcgui/window"
	"github.com/weijianhong/xcgui/xc"
	"testing"
)

func TestGetDropFiles(t *testing.T) {
	tf.TFunc(func(a *app.App, w *window.Window) {
		w.EnableDragFiles(true)
		w.Event_DROPFILES(func(hDropInfo uintptr, pbHandled *bool) int {
			fmt.Println(wutil.GetDropFiles(hDropInfo))
			return 0
		})
	})
}

func TestOpenDir(t *testing.T) {
	fmt.Println(wutil.OpenDir(0))
}

func TestOpenFile(t *testing.T) {
	fmt.Println(wutil.OpenFile(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, ""))
}

func TestOpenFiles(t *testing.T) {
	for _, s := range wutil.OpenFiles(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "") {
		fmt.Println(s)
	}
}

func TestSaveFile(t *testing.T) {
	fmt.Println(wutil.SaveFile(0, []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, "", "默认文件名.txt"))
}

func TestChooseColor(t *testing.T) {
	rgb := wutil.ChooseColor(0)
	fmt.Println("rgb颜色", rgb)
	fmt.Println("ARGB颜色", xc.RGB2ARGB(rgb, 255))
}
