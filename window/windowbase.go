package window

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/xc"
)

// 窗口基类
type windowBase struct {
	objectbase.UI
}

// 炫彩_消息框
// pText: 内容文本
// pCaption: 标题
// nFlags: 标识, MessageBox_Flag_
func (w *windowBase) MessageBox(pText string, pCaption string, nFlags int) int {
	return xc.XC_MessageBox(w.Handle, pText, pCaption, nFlags)
}

// 炫彩_发送窗口消息
// msg:
// wParam:
// lParam:
func (w *windowBase) SendMessage(msg int, wParam int, lParam int) int {
	return xc.XC_SendMessage(w.Handle, msg, wParam, lParam)
}

// 炫彩_投递窗口消息
// msg:
// wParam:
// lParam:
func (w *windowBase) PostMessage(msg int, wParam int, lParam int) bool {
	return xc.XC_PostMessage(w.Handle, msg, wParam, lParam)
}

// 炫彩_判断窗口, 判断是否为窗口句柄
func (w *windowBase) IsHWINDOW() bool {
	return xc.XC_IsHWINDOW(w.Handle)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象
// nID: ID值.
func (w *windowBase) GetObjectByID(nID int) int {
	return xc.XC_GetObjectByID(w.Handle, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄
// pName: ID名称.
func (w *windowBase) GetObjectByIDName(pName string) int {
	return xc.XC_GetObjectByIDName(w.Handle, pName)
}

// 窗口_显示
// nCmdShow: 显示方式, SW_
func (w *windowBase) ShowWindow(nCmdShow int) int {
	return xc.XWnd_ShowWindow(w.Handle, nCmdShow)
}

// 窗口_置顶
func (w *windowBase) SetTop() int {
	return xc.XWnd_SetTop(w.Handle)
}

// 窗口_注册事件C
// nEvent: 事件类型.
// pFun: 事件函数.
func (w *windowBase) RegEventC(nEvent int, pFun interface{}) bool {
	return xc.XWnd_RegEventC(w.Handle, nEvent, pFun)
}

// 窗口_注册事件C1
// nEvent: 事件类型.
// pFun: 事件函数.
func (w *windowBase) RegEventC1(nEvent int, pFun interface{}) bool {
	return xc.XWnd_RegEventC1(w.Handle, nEvent, pFun)
}

// 窗口_移除事件C
// nEvent: 事件类型.
// pFun: 事件函数.
func (w *windowBase) RemoveEventC(nEvent int, pFun interface{}) bool {
	return xc.XWnd_RemoveEventC(w.Handle, nEvent, pFun)
}

// 窗口_添加子对象
// hChild: 要添加的对象句柄.
func (w *windowBase) AddChild(hChild int) bool {
	return xc.XWnd_AddChild(w.Handle, hChild)
}

// 窗口_插入子对象
// hChild: 要插入的对象句柄.
// index: 插入位置索引.
func (w *windowBase) InsertChild(hChild int, index int) bool {
	return xc.XWnd_InsertChild(w.Handle, hChild, index)
}

// 窗口_取HWND
func (w *windowBase) GetHWND() int {
	return xc.XWnd_GetHWND(w.Handle)
}

// 窗口_重绘
// bImmediate: 是否立即重绘, 默认为否.
func (w *windowBase) Redraw(bImmediate bool) int {
	return xc.XWnd_Redraw(w.Handle, bImmediate)
}

// 窗口_重绘指定区域
// pRect: 需要重绘的区域坐标.
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func (w *windowBase) RedrawRect(pRect *xc.RECT, bImmediate bool) int {
	return xc.XWnd_RedrawRect(w.Handle, pRect, bImmediate)
}

// 窗口_置坐标
// pRect: 坐标
func (w *windowBase) SetRect(pRect *xc.RECT) bool {
	return xc.XWnd_SetRect(w.Handle, pRect)
}

// 窗口_置焦点
// hFocusEle: 将获得焦点的元素.
func (w *windowBase) SetFocusEle(hFocusEle int) bool {
	return xc.XWnd_SetFocusEle(w.Handle, hFocusEle)
}

// 窗口_取焦点
func (w *windowBase) GetFocusEle() int {
	return xc.XWnd_GetFocusEle(w.Handle)
}

// 窗口_取鼠标停留元素
func (w *windowBase) GetStayEle() int {
	return xc.XWnd_GetStayEle(w.Handle)
}

// 窗口_绘制
// hDraw: 图形绘制句柄.
func (w *windowBase) DrawWindow(hDraw int) int {
	return xc.XWnd_DrawWindow(w.Handle, hDraw)
}

// 窗口_居中
func (w *windowBase) Center() int {
	return xc.XWnd_Center(w.Handle)
}

// 窗口_居中扩展
// width: 窗口宽度
// height: 窗口高度
func (w *windowBase) CenterEx(width, height int) int {
	return xc.XWnd_CenterEx(w.Handle, width, height)
}

// 窗口_置光标
// hCursor: 鼠标光标句柄.
func (w *windowBase) SetCursor(hCursor int) int {
	return xc.XWnd_SetCursor(w.Handle, hCursor)
}

// 窗口_取光标
func (w *windowBase) GetCursor() int {
	return xc.XWnd_GetCursor(w.Handle)
}

// 窗口_启用拖动边框
// bEnable: 是否启用.
func (w *windowBase) EnableDragBorder(bEnable bool) int {
	return xc.XWnd_EnableDragBorder(w.Handle, bEnable)
}

// 窗口_启用拖动窗口
// bEnable: 是否启用.
func (w *windowBase) EnableDragWindow(bEnable bool) int {
	return xc.XWnd_EnableDragWindow(w.Handle, bEnable)
}

// 窗口_启用拖动标题栏
// bEnable: 是否启用.
func (w *windowBase) EnableDragCaption(bEnable bool) int {
	return xc.XWnd_EnableDragCaption(w.Handle, bEnable)
}

// 窗口_启用绘制背景
// bEnable: 是否启用.
func (w *windowBase) EnableDrawBk(bEnable bool) int {
	return xc.XWnd_EnableDrawBk(w.Handle, bEnable)
}

// 窗口_启用自动焦点
// bEnable: 是否启用.
func (w *windowBase) EnableAutoFocus(bEnable bool) int {
	return xc.XWnd_EnableAutoFocus(w.Handle, bEnable)
}

// 窗口_启用允许最大化
// bEnable: 是否启用.
func (w *windowBase) EnableMaxWindow(bEnable bool) int {
	return xc.XWnd_EnableMaxWindow(w.Handle, bEnable)
}

// 窗口_启用限制窗口大小
// bEnable: 是否启用
func (w *windowBase) EnablemLimitWindowSize(bEnable bool) int {
	return xc.XWnd_EnablemLimitWindowSize(w.Handle, bEnable)
}

// 窗口_启用布局
// bEnable: 是否启用.
func (w *windowBase) EnableLayout(bEnable bool) int {
	return xc.XWnd_EnableLayout(w.Handle, bEnable)
}

// 窗口_启用布局覆盖边框
// bEnable: 是否启用
func (w *windowBase) EnableLayoutOverlayBorder(bEnable bool) int {
	return xc.XWnd_EnableLayoutOverlayBorder(w.Handle, bEnable)
}

// 窗口_显示布局边界
// bEnable: 是否启用.
func (w *windowBase) ShowLayoutFrame(bEnable bool) int {
	return xc.XWnd_ShowLayoutFrame(w.Handle, bEnable)
}

// 窗口_判断启用布局
func (w *windowBase) IsEnableLayout() bool {
	return xc.XWnd_IsEnableLayout(w.Handle)
}

// 窗口_是否最大化
func (w *windowBase) IsMaxWindow() bool {
	return xc.XWnd_IsMaxWindow(w.Handle)
}

// 窗口_置鼠标捕获元素
// hEle: 元素句柄.
func (w *windowBase) SetCaptureEle(hEle int) int {
	return xc.XWnd_SetCaptureEle(w.Handle, hEle)
}

// 窗口_取鼠标捕获元素
func (w *windowBase) GetCaptureEle() int {
	return xc.XWnd_GetCaptureEle(w.Handle)
}

// 窗口_取绘制矩形
// pRcPaint: 重绘区域坐标.
func (w *windowBase) GetDrawRect(pRcPaint *xc.RECT) int {
	return xc.XWnd_GetDrawRect(w.Handle, pRcPaint)
}

// 窗口_置系统光标
// hCursor: 光标句柄.
func (w *windowBase) SetCursorSys(hCursor int) int {
	return xc.XWnd_SetCursorSys(w.Handle, hCursor)
}

// 窗口_置字体
// hFontx: 炫彩字体句柄.
func (w *windowBase) SetFont(hFontx int) int {
	return xc.XWnd_SetFont(w.Handle, hFontx)
}

// 窗口_置文本颜色
// color: RGB颜色值.
// alpha: 透明度.
func (w *windowBase) SetTextColor(color int, alpha uint8) int {
	return xc.XWnd_SetTextColor(w.Handle, color, alpha)
}

// 窗口_取文本颜色
func (w *windowBase) GetTextColor() int {
	return xc.XWnd_GetTextColor(w.Handle)
}

// 窗口_取文本颜色扩展
func (w *windowBase) GetTextColorEx() int {
	return xc.XWnd_GetTextColorEx(w.Handle)
}

// 窗口_置ID
// nID: ID值.
func (w *windowBase) SetID(nID int) int {
	return xc.XWnd_SetID(w.Handle, nID)
}

// 窗口_取ID
func (w *windowBase) GetID() int {
	return xc.XWnd_GetID(w.Handle)
}

// 窗口_置名称
// pName: name值, 字符串.
func (w *windowBase) SetName(pName string) int {
	return xc.XWnd_SetName(w.Handle, pName)
}

// 窗口_取名称
func (w *windowBase) GetName() string {
	return xc.XWnd_GetName(w.Handle)
}

// 窗口_置边大小
// left: 窗口左边大小.
// top: 窗口上边大小.
// right: 窗口右边大小.
// bottom: 窗口底部大小.
func (w *windowBase) SetBorderSize(left, top, right, bottom int) int {
	return xc.XWnd_SetBorderSize(w.Handle, left, top, right, bottom)
}

// 窗口_取边大小
func (w *windowBase) GetBorderSize(pBorder *xc.RECT) int {
	return xc.XWnd_GetBorderSize(w.Handle, pBorder)
}

// 窗口_置布局内填充大小
// left: 左边大小.
// top: 上边大小.
// right: 右边大小.
// bottom: 下边大小.
func (w *windowBase) SetPadding(left, top, right, bottom int) int {
	return xc.XWnd_SetPadding(w.Handle, left, top, right, bottom)
}

// 窗口_置拖动边框大小
// left: 窗口左边大小.
// top: 窗口上边大小.
// right: 窗口右边大小.
// bottom: 窗口底边大小.
func (w *windowBase) SetDragBorderSize(left, top, right, bottom int) int {
	return xc.XWnd_SetDragBorderSize(w.Handle, left, top, right, bottom)
}

// 窗口_取拖动边框大小
// pSize: 拖动边框大小.
func (w *windowBase) GetDragBorderSize(pBorder *xc.RECT) int {
	return xc.XWnd_GetDragBorderSize(w.Handle, pBorder)
}

// 窗口_置最小大小
// width: 最小宽度.
// height: 最小高度.
func (w *windowBase) SetMinimumSize(width, height int) int {
	return xc.XWnd_SetMinimumSize(w.Handle, width, height)
}

// 窗口_测试点击子元素
// pPt: 左边点.
func (w *windowBase) HitChildEle(pPt *xc.POINT) int {
	return xc.XWnd_HitChildEle(w.Handle, pPt)
}

// 窗口_取子对象数量
func (w *windowBase) GetChildCount() int {
	return xc.XWnd_GetChildCount(w.Handle)
}

// 窗口_取子对象从索引
// index: 元素索引.
func (w *windowBase) GetChildByIndex(index int) int {
	return xc.XWnd_GetChildByIndex(w.Handle, index)
}

// 窗口_取子对象从ID
// nID: 元素ID, ID必须大于0.
func (w *windowBase) GetChildByID(nID int) int {
	return xc.XWnd_GetChildByID(w.Handle, nID)
}

// 窗口_取子对象
// nID: 对象ID,ID必须大于0.
func (w *windowBase) GetChild(nID int) int {
	return xc.XWnd_GetChild(w.Handle, nID)
}

// 窗口_关闭
func (w *windowBase) CloseWindow() int {
	return xc.XWnd_CloseWindow(w.Handle)
}

// 窗口_调整布局
func (w *windowBase) AdjustLayout() int {
	return xc.XWnd_AdjustLayout(w.Handle)
}

// 窗口_调整布局扩展
// nFlags: 调整标识, AdjustLayout_
func (w *windowBase) AdjustLayoutEx(nFlags int) int {
	return xc.XWnd_AdjustLayoutEx(w.Handle, nFlags)
}

// 窗口_创建插入符
// hEle: 元素句柄.
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func (w *windowBase) CreateCaret(hEle, x, y, width, height int) int {
	return xc.XWnd_CreateCaret(w.Handle, hEle, x, y, width, height)
}

// 窗口_置插入符大小
// width: 宽度.
// height: 高度.
func (w *windowBase) SetCaretSize(width, height int) int {
	return xc.XWnd_SetCaretSize(w.Handle, width, height)
}

// 窗口_置插入符位置
// x: x坐标.
// y: y坐标.
func (w *windowBase) SetCaretPos(x, y int) int {
	return xc.XWnd_SetCaretPos(w.Handle, x, y)
}

// 窗口_置插入符位置扩展
// x: x坐标.
// y: y坐标.
// width: 宽度.
// height: 高度.
func (w *windowBase) SetCaretPosEx(x, y, width, height int) int {
	return xc.XWnd_SetCaretPosEx(w.Handle, x, y, width, height)
}

// 窗口_置插入符颜色
// color: 颜色值.
func (w *windowBase) SetCaretColor(color int) int {
	return xc.XWnd_SetCaretColor(w.Handle, color)
}

// 窗口_显示插入符
// bShow: 是否显示.
func (w *windowBase) ShowCaret(bShow bool) int {
	return xc.XWnd_ShowCaret(w.Handle, bShow)
}

// 窗口_销毁插入符
func (w *windowBase) DestroyCaret() int {
	return xc.XWnd_DestroyCaret(w.Handle)
}

// 窗口_取插入符元素
func (w *windowBase) GetCaretHELE() int {
	return xc.XWnd_GetCaretHELE(w.Handle)
}

// 窗口_取客户区坐标
// pRect: 坐标.
func (w *windowBase) GetClientRect(pRect *xc.RECT) int {
	return xc.XWnd_GetClientRect(w.Handle, pRect)
}

// 窗口_取Body坐标
// pRect: 坐标.
func (w *windowBase) GetBodyRect(pRect *xc.RECT) int {
	return xc.XWnd_GetBodyRect(w.Handle, pRect)
}

// 窗口_取布局坐标
// pRect: 接收返回坐标
func (w *windowBase) GetLayoutRect(pRect *xc.RECT) int {
	return xc.XWnd_GetLayoutRect(w.Handle, pRect)
}

// 窗口_移动窗口
// x: X坐标
// y: Y坐标
func (w *windowBase) Move(x, y int) int {
	return xc.XWnd_Move(w.Handle, x, y)
}

// 窗口_取坐标
// pRect: 坐标
func (w *windowBase) GetRect(pRect *xc.RECT) int {
	return xc.XWnd_GetRect(w.Handle, pRect)
}

// 窗口_最大化
// bMaximize: 是否最大化
func (w *windowBase) MaxWindow(bMaximize bool) int {
	return xc.XWnd_MaxWindow(w.Handle, bMaximize)
}

// 窗口_置定时器
// nIDEvent: 定时器ID.
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) SetTimer(nIDEvent, uElapse int) int {
	return xc.XWnd_SetTimer(w.Handle, nIDEvent, uElapse)
}

// 窗口_关闭定时器
// nIDEvent: 定时器ID.
func (w *windowBase) KillTimer(nIDEvent int) int {
	return xc.XWnd_KillTimer(w.Handle, nIDEvent)
}

// 窗口_置炫彩定时器
// nIDEvent: 定时器ID.
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) SetXCTimer(nIDEvent, uElapse int) int {
	return xc.XWnd_SetXCTimer(w.Handle, nIDEvent, uElapse)
}

// 窗口_关闭炫彩定时器
// nIDEvent: 定时器ID.
func (w *windowBase) KillXCTimer(nIDEvent int) int {
	return xc.XWnd_KillXCTimer(w.Handle, nIDEvent)
}

// 窗口_取背景管理器
func (w *windowBase) GetBkManager() int {
	return xc.XWnd_GetBkManager(w.Handle)
}

// 窗口_取背景管理器扩展
func (w *windowBase) GetBkManagerEx() int {
	return xc.XWnd_GetBkManagerEx(w.Handle)
}

// 窗口_置背景管理器
// hBkInfoM: 背景管理器
func (w *windowBase) SetBkMagager(hBkInfoM int) int {
	return xc.XWnd_SetBkMagager(w.Handle, hBkInfoM)
}

// 窗口_置透明类型
// nType: 窗口透明类型.
func (w *windowBase) SetTransparentType(nType int) int {
	return xc.XWnd_SetTransparentType(w.Handle, nType)
}

// 窗口_置透明度
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func (w *windowBase) SetTransparentAlpha(alpha uint8) int {
	return xc.XWnd_SetTransparentAlpha(w.Handle, alpha)
}

// 窗口_置透明色
// color: 窗口透明色.
func (w *windowBase) SetTransparentColor(color int) int {
	return xc.XWnd_SetTransparentColor(w.Handle, color)
}

// 窗口_置阴影信息
// nSize: 阴影大小
// nDepth: 阴影深度, 0-255.
// nAngeleSize: 圆角阴影内收大小.
// bRightAngle: 是否强制直角.
// color: 阴影颜色.
func (w *windowBase) SetShadowInfo(nSize int, nDepth uint8, nAngeleSize int, bRightAngle bool, color int) int {
	return xc.XWnd_SetShadowInfo(w.Handle, nSize, nDepth, nAngeleSize, bRightAngle, color)
}

// 窗口_取阴影信息
// pnSize: 阴影大小.
// pnDepth: 阴影深度.
// pnAngeleSize: 圆角阴影内收大小 .
// pbRightAngle: 是否强制直角.
// pColor: 阴影颜色.
func (w *windowBase) GetShadowInfo(nSize *int, nDepth *uint8, nAngeleSize *int, bRightAngle *bool, color *int) int {
	return xc.XWnd_GetShadowInfo(w.Handle, nSize, nDepth, nAngeleSize, bRightAngle, color)
}

// 窗口_取透明类型
func (w *windowBase) GetTransparentType() int {
	return xc.XWnd_GetTransparentType(w.Handle)
}
