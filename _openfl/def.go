package _openfl
type XArray uintptr // TClassdecl
type XArray_T uintptr // TClassdecl
func (x XArray) Xfilter(Xf interface{}) XArray { return 0; }
func (x XArray) Xmap(Xf interface{}) XArray { return 0; }
func (x XArray) Xiterator() uintptr { return 0; }
func (x XArray) Xcopy() XArray { return 0; }
func (x XArray) Xremove(Xx XArray_T) bool { return false; }
func (x XArray) Xinsert(Xpos int,Xx XArray_T) {}
func (x XArray) Xunshift(Xx XArray_T) {}
func (x XArray) XtoString() string { return ""; }
func (x XArray) Xsplice(Xpos int,Xlen int) XArray { return 0; }
func (x XArray) Xsort(Xf interface{}) {}
func (x XArray) Xslice(Xpos int,Xend int) XArray { return 0; }
func (x XArray) Xshift() uintptr { return 0; }
func (x XArray) Xreverse() {}
func (x XArray) Xpush(Xx XArray_T) int { return 0; }
func (x XArray) Xpop() uintptr { return 0; }
func (x XArray) Xjoin(Xsep string) string { return ""; }
func (x XArray) Xconcat(Xa XArray) XArray { return 0; }
func (x XArray) Xlength_goget()  int { return 0; }
func (x XArray) Xlength_goset(int){}
func XArray_new() XArray { return XArray(0); }
type XClass uintptr // TAbstractdecl
type XEnum uintptr // TAbstractdecl
type XEnumValue uintptr // TAbstractdecl
type XMap uintptr // TAbstractdecl
type XIMap uintptr // TClassdecl
type XIMap_K uintptr // TClassdecl
type XIMap_V uintptr // TClassdecl
func XMath_PI_goget()  float64 { return 0; }
func XMath_PI_goset(float64){}
func XMath_NEGATIVE___INFINITY_goget()  float64 { return 0; }
func XMath_NEGATIVE___INFINITY_goset(float64){}
func XMath_POSITIVE___INFINITY_goget()  float64 { return 0; }
func XMath_POSITIVE___INFINITY_goset(float64){}
func XMath_NaN_goget()  float64 { return 0; }
func XMath_NaN_goset(float64){}
func XMath_abs(Xv float64) float64 { return 0; }
func XMath_min(Xa float64,Xb float64) float64 { return 0; }
func XMath_max(Xa float64,Xb float64) float64 { return 0; }
func XMath_sin(Xv float64) float64 { return 0; }
func XMath_cos(Xv float64) float64 { return 0; }
func XMath_tan(Xv float64) float64 { return 0; }
func XMath_asin(Xv float64) float64 { return 0; }
func XMath_acos(Xv float64) float64 { return 0; }
func XMath_atan(Xv float64) float64 { return 0; }
func XMath_atan2(Xy float64,Xx float64) float64 { return 0; }
func XMath_exp(Xv float64) float64 { return 0; }
func XMath_log(Xv float64) float64 { return 0; }
func XMath_pow(Xv float64,Xexp float64) float64 { return 0; }
func XMath_sqrt(Xv float64) float64 { return 0; }
func XMath_round(Xv float64) int { return 0; }
func XMath_floor(Xv float64) int { return 0; }
func XMath_ceil(Xv float64) int { return 0; }
func XMath_random() float64 { return 0; }
func XMath_ffloor(Xv float64) float64 { return 0; }
func XMath_fceil(Xv float64) float64 { return 0; }
func XMath_fround(Xv float64) float64 { return 0; }
func XMath_isFinite(Xf float64) bool { return false; }
func XMath_isNaN(Xf float64) bool { return false; }
type XMath uintptr // TClassdecl
func XStd_string(Xs interface{}) string { return ""; }
type XStd uintptr // TClassdecl
type XVoid uintptr // TAbstractdecl
type XFloat uintptr // TAbstractdecl
type XInt uintptr // TAbstractdecl
type XNull uintptr // TTypedecl
type XBool uintptr // TAbstractdecl
type XDynamic uintptr // TAbstractdecl
type XIterator uintptr // TTypedecl
type XIterable uintptr // TTypedecl
type XArrayAccess uintptr // TClassdecl
type XArrayAccess_T uintptr // TClassdecl
func XString_fromCharCode(Xcode int) string { return ""; }
type XString uintptr // TClassdecl
func (x XString) XtoString() string { return ""; }
func (x XString) Xsubstring(XstartIndex int,XendIndex int) string { return ""; }
func (x XString) Xsubstr(Xpos int,Xlen int) string { return ""; }
func (x XString) Xsplit(Xdelimiter string) XArray { return 0; }
func (x XString) XlastIndexOf(Xstr string,XstartIndex int) int { return 0; }
func (x XString) XindexOf(Xstr string,XstartIndex int) int { return 0; }
func (x XString) XcharCodeAt(Xindex int) uintptr { return 0; }
func (x XString) XcharAt(Xindex int) string { return ""; }
func (x XString) XtoLowerCase() string { return ""; }
func (x XString) XtoUpperCase() string { return ""; }
func (x XString) Xlength_goget()  int { return 0; }
func (x XString) Xlength_goset(int){}
func XString_new(Xstring string) XString { return XString(0); }
type XUInt int // TTypedecl
func Xflash_Lib_current_goget()  Xflash_display_MovieClip { return 0; }
func Xflash_Lib_current_goset(Xflash_display_MovieClip){}
func Xflash_Lib_getTimer() int { return 0; }
func Xflash_Lib_eval(Xpath string) interface{} { return nil; }
func Xflash_Lib_getURL(Xurl Xflash_net_URLRequest,Xtarget string) {}
func Xflash_Lib_fscommand(Xcmd string,Xparam string) {}
func Xflash_Lib_trace(Xarg interface{}) {}
func Xflash_Lib_attach(Xname string) Xflash_display_MovieClip { return 0; }
func Xflash_Lib_as(Xv interface{},Xc XClass) uintptr { return 0; }
func Xflash_Lib_redirectTraces() {}
type Xflash_Lib uintptr // TClassdecl
func Xflash_Memory_select(XinBytes Xflash_utils_ByteArray) {}
func Xflash_Memory_getByte(Xaddr int) int { return 0; }
func Xflash_Memory_getDouble(Xaddr int) float64 { return 0; }
func Xflash_Memory_getFloat(Xaddr int) float64 { return 0; }
func Xflash_Memory_getI32(Xaddr int) int { return 0; }
func Xflash_Memory_getUI16(Xaddr int) int { return 0; }
func Xflash_Memory_setByte(Xaddr int,Xv int) {}
func Xflash_Memory_setDouble(Xaddr int,Xv float64) {}
func Xflash_Memory_setFloat(Xaddr int,Xv float64) {}
func Xflash_Memory_setI16(Xaddr int,Xv int) {}
func Xflash_Memory_setI32(Xaddr int,Xv int) {}
type Xflash_Memory uintptr // TClassdecl
type Xflash_Vector uintptr // TClassdecl
type Xflash_Vector_T uintptr // TClassdecl
func (x Xflash_Vector) XlastIndexOf(Xx Xflash_Vector_T,Xfrom int) int { return 0; }
func (x Xflash_Vector) XindexOf(Xx Xflash_Vector_T,Xfrom int) int { return 0; }
func (x Xflash_Vector) XtoString() string { return ""; }
func (x Xflash_Vector) Xsplice(Xpos int,Xlen int) Xflash_Vector { return 0; }
func (x Xflash_Vector) Xsort(Xf interface{}) {}
func (x Xflash_Vector) Xslice(Xpos int,Xend int) Xflash_Vector { return 0; }
func (x Xflash_Vector) Xunshift(Xx Xflash_Vector_T) {}
func (x Xflash_Vector) Xshift() uintptr { return 0; }
func (x Xflash_Vector) Xreverse() {}
func (x Xflash_Vector) Xpush(Xx Xflash_Vector_T) int { return 0; }
func (x Xflash_Vector) Xpop() uintptr { return 0; }
func (x Xflash_Vector) Xjoin(Xsep string) string { return ""; }
func (x Xflash_Vector) Xconcat(Xa Xflash_Vector) Xflash_Vector { return 0; }
func (x Xflash_Vector) Xfixed_goget()  bool { return false; }
func (x Xflash_Vector) Xfixed_goset(bool){}
func (x Xflash_Vector) Xlength_goget()  int { return 0; }
func (x Xflash_Vector) Xlength_goset(int){}
func Xflash_Vector_new(Xlength uintptr,Xfixed bool) Xflash_Vector { return Xflash_Vector(0); }
type Xflash_events_IEventDispatcher uintptr // TClassdecl
func (x Xflash_events_IEventDispatcher) XwillTrigger(Xtype string) bool { return false; }
func (x Xflash_events_IEventDispatcher) XremoveEventListener(Xtype string,Xlistener interface{},XuseCapture bool) {}
func (x Xflash_events_IEventDispatcher) XhasEventListener(Xtype string) bool { return false; }
func (x Xflash_events_IEventDispatcher) XdispatchEvent(Xevent Xflash_events_Event) bool { return false; }
func (x Xflash_events_IEventDispatcher) XaddEventListener(Xtype string,Xlistener interface{},XuseCapture bool,Xpriority int,XuseWeakReference bool) {}
type Xflash_events_EventDispatcher uintptr // TClassdecl
func (x Xflash_events_EventDispatcher) XwillTrigger(Xtype string) bool { return false; }
func (x Xflash_events_EventDispatcher) XtoString() string { return ""; }
func (x Xflash_events_EventDispatcher) XremoveEventListener(Xtype string,Xlistener interface{},XuseCapture bool) {}
func (x Xflash_events_EventDispatcher) XhasEventListener(Xtype string) bool { return false; }
func (x Xflash_events_EventDispatcher) XdispatchEvent(Xevent Xflash_events_Event) bool { return false; }
func (x Xflash_events_EventDispatcher) XaddEventListener(Xtype string,Xlistener interface{},XuseCapture bool,Xpriority int,XuseWeakReference bool) {}
func Xflash_events_EventDispatcher_new(Xtarget Xflash_events_IEventDispatcher) Xflash_events_EventDispatcher { return Xflash_events_EventDispatcher(0); }
func Xflash_events_Event_ACTIVATE_goget()  string { return ""; }
func Xflash_events_Event_ACTIVATE_goset(string){}
func Xflash_events_Event_ADDED_goget()  string { return ""; }
func Xflash_events_Event_ADDED_goset(string){}
func Xflash_events_Event_ADDED___TO___STAGE_goget()  string { return ""; }
func Xflash_events_Event_ADDED___TO___STAGE_goset(string){}
func Xflash_events_Event_CANCEL_goget()  string { return ""; }
func Xflash_events_Event_CANCEL_goset(string){}
func Xflash_events_Event_CHANNEL___MESSAGE_goget()  string { return ""; }
func Xflash_events_Event_CHANNEL___MESSAGE_goset(string){}
func Xflash_events_Event_CHANNEL___STATE_goget()  string { return ""; }
func Xflash_events_Event_CHANNEL___STATE_goset(string){}
func Xflash_events_Event_CHANGE_goget()  string { return ""; }
func Xflash_events_Event_CHANGE_goset(string){}
func Xflash_events_Event_CLOSE_goget()  string { return ""; }
func Xflash_events_Event_CLOSE_goset(string){}
func Xflash_events_Event_COMPLETE_goget()  string { return ""; }
func Xflash_events_Event_COMPLETE_goset(string){}
func Xflash_events_Event_CONNECT_goget()  string { return ""; }
func Xflash_events_Event_CONNECT_goset(string){}
func Xflash_events_Event_DEACTIVATE_goget()  string { return ""; }
func Xflash_events_Event_DEACTIVATE_goset(string){}
func Xflash_events_Event_ENTER___FRAME_goget()  string { return ""; }
func Xflash_events_Event_ENTER___FRAME_goset(string){}
func Xflash_events_Event_FULLSCREEN_goget()  string { return ""; }
func Xflash_events_Event_FULLSCREEN_goset(string){}
func Xflash_events_Event_ID3_goget()  string { return ""; }
func Xflash_events_Event_ID3_goset(string){}
func Xflash_events_Event_INIT_goget()  string { return ""; }
func Xflash_events_Event_INIT_goset(string){}
func Xflash_events_Event_MOUSE___LEAVE_goget()  string { return ""; }
func Xflash_events_Event_MOUSE___LEAVE_goset(string){}
func Xflash_events_Event_OPEN_goget()  string { return ""; }
func Xflash_events_Event_OPEN_goset(string){}
func Xflash_events_Event_REMOVED_goget()  string { return ""; }
func Xflash_events_Event_REMOVED_goset(string){}
func Xflash_events_Event_REMOVED___FROM___STAGE_goget()  string { return ""; }
func Xflash_events_Event_REMOVED___FROM___STAGE_goset(string){}
func Xflash_events_Event_RENDER_goget()  string { return ""; }
func Xflash_events_Event_RENDER_goset(string){}
func Xflash_events_Event_RESIZE_goget()  string { return ""; }
func Xflash_events_Event_RESIZE_goset(string){}
func Xflash_events_Event_SCROLL_goget()  string { return ""; }
func Xflash_events_Event_SCROLL_goset(string){}
func Xflash_events_Event_SELECT_goget()  string { return ""; }
func Xflash_events_Event_SELECT_goset(string){}
func Xflash_events_Event_SOUND___COMPLETE_goget()  string { return ""; }
func Xflash_events_Event_SOUND___COMPLETE_goset(string){}
func Xflash_events_Event_TAB___CHILDREN___CHANGE_goget()  string { return ""; }
func Xflash_events_Event_TAB___CHILDREN___CHANGE_goset(string){}
func Xflash_events_Event_TAB___ENABLED___CHANGE_goget()  string { return ""; }
func Xflash_events_Event_TAB___ENABLED___CHANGE_goset(string){}
func Xflash_events_Event_TAB___INDEX___CHANGE_goget()  string { return ""; }
func Xflash_events_Event_TAB___INDEX___CHANGE_goset(string){}
func Xflash_events_Event_UNLOAD_goget()  string { return ""; }
func Xflash_events_Event_UNLOAD_goset(string){}
func Xflash_events_Event_VIDEO___FRAME_goget()  string { return ""; }
func Xflash_events_Event_VIDEO___FRAME_goset(string){}
func Xflash_events_Event_WORKER___STATE_goget()  string { return ""; }
func Xflash_events_Event_WORKER___STATE_goset(string){}
type Xflash_events_Event uintptr // TClassdecl
func (x Xflash_events_Event) XtoString() string { return ""; }
func (x Xflash_events_Event) XstopPropagation() {}
func (x Xflash_events_Event) XstopImmediatePropagation() {}
func (x Xflash_events_Event) XpreventDefault() {}
func (x Xflash_events_Event) XisDefaultPrevented() bool { return false; }
func (x Xflash_events_Event) XformatToString(XclassName string,Xp1 interface{},Xp2 interface{},Xp3 interface{},Xp4 interface{},Xp5 interface{}) string { return ""; }
func (x Xflash_events_Event) Xclone() Xflash_events_Event { return 0; }
func (x Xflash_events_Event) Xtype_goget()  string { return ""; }
func (x Xflash_events_Event) Xtype_goset(string){}
func (x Xflash_events_Event) Xtarget_goget()  interface{} { return nil; }
func (x Xflash_events_Event) Xtarget_goset(interface{}){}
func (x Xflash_events_Event) XeventPhase_goget()  Xflash_events_EventPhase { return 0; }
func (x Xflash_events_Event) XeventPhase_goset(Xflash_events_EventPhase){}
func (x Xflash_events_Event) XcurrentTarget_goget()  interface{} { return nil; }
func (x Xflash_events_Event) XcurrentTarget_goset(interface{}){}
func (x Xflash_events_Event) Xcancelable_goget()  bool { return false; }
func (x Xflash_events_Event) Xcancelable_goset(bool){}
func (x Xflash_events_Event) Xbubbles_goget()  bool { return false; }
func (x Xflash_events_Event) Xbubbles_goset(bool){}
func Xflash_events_Event_new(Xtype string,Xbubbles bool,Xcancelable bool) Xflash_events_Event { return Xflash_events_Event(0); }
type Xflash_events_AccelerometerEvent uintptr // TClassdecl
func (x Xflash_events_AccelerometerEvent) Xtimestamp_goget()  float64 { return 0; }
func (x Xflash_events_AccelerometerEvent) Xtimestamp_goset(float64){}
func (x Xflash_events_AccelerometerEvent) XaccelerationZ_goget()  float64 { return 0; }
func (x Xflash_events_AccelerometerEvent) XaccelerationZ_goset(float64){}
func (x Xflash_events_AccelerometerEvent) XaccelerationY_goget()  float64 { return 0; }
func (x Xflash_events_AccelerometerEvent) XaccelerationY_goset(float64){}
func (x Xflash_events_AccelerometerEvent) XaccelerationX_goget()  float64 { return 0; }
func (x Xflash_events_AccelerometerEvent) XaccelerationX_goset(float64){}
func Xflash_events_AccelerometerEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xtimestamp float64,XaccelerationX float64,XaccelerationY float64,XaccelerationZ float64) Xflash_events_AccelerometerEvent { return Xflash_events_AccelerometerEvent(0); }
func Xflash_events_TextEvent_LINK_goget()  string { return ""; }
func Xflash_events_TextEvent_LINK_goset(string){}
func Xflash_events_TextEvent_TEXT___INPUT_goget()  string { return ""; }
func Xflash_events_TextEvent_TEXT___INPUT_goset(string){}
type Xflash_events_TextEvent uintptr // TClassdecl
func (x Xflash_events_TextEvent) Xtext_goget()  string { return ""; }
func (x Xflash_events_TextEvent) Xtext_goset(string){}
func Xflash_events_TextEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xtext string) Xflash_events_TextEvent { return Xflash_events_TextEvent(0); }
func Xflash_events_ErrorEvent_ERROR_goget()  string { return ""; }
func Xflash_events_ErrorEvent_ERROR_goset(string){}
type Xflash_events_ErrorEvent uintptr // TClassdecl
func Xflash_events_ErrorEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xtext string,Xid int) Xflash_events_ErrorEvent { return Xflash_events_ErrorEvent(0); }
type Xflash_events_EventPhase uintptr // TEnumdecl
func Xflash_events_FocusEvent_FOCUS___IN_goget()  string { return ""; }
func Xflash_events_FocusEvent_FOCUS___IN_goset(string){}
func Xflash_events_FocusEvent_FOCUS___OUT_goget()  string { return ""; }
func Xflash_events_FocusEvent_FOCUS___OUT_goset(string){}
func Xflash_events_FocusEvent_KEY___FOCUS___CHANGE_goget()  string { return ""; }
func Xflash_events_FocusEvent_KEY___FOCUS___CHANGE_goset(string){}
func Xflash_events_FocusEvent_MOUSE___FOCUS___CHANGE_goget()  string { return ""; }
func Xflash_events_FocusEvent_MOUSE___FOCUS___CHANGE_goset(string){}
type Xflash_events_FocusEvent uintptr // TClassdecl
func (x Xflash_events_FocusEvent) XshiftKey_goget()  bool { return false; }
func (x Xflash_events_FocusEvent) XshiftKey_goset(bool){}
func (x Xflash_events_FocusEvent) XrelatedObject_goget()  Xflash_display_InteractiveObject { return 0; }
func (x Xflash_events_FocusEvent) XrelatedObject_goset(Xflash_display_InteractiveObject){}
func (x Xflash_events_FocusEvent) XkeyCode_goget()  int { return 0; }
func (x Xflash_events_FocusEvent) XkeyCode_goset(int){}
func Xflash_events_FocusEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,XrelatedObject Xflash_display_InteractiveObject,XshiftKey bool,XkeyCode int) Xflash_events_FocusEvent { return Xflash_events_FocusEvent(0); }
func Xflash_events_HTTPStatusEvent_HTTP___STATUS_goget()  string { return ""; }
func Xflash_events_HTTPStatusEvent_HTTP___STATUS_goset(string){}
type Xflash_events_HTTPStatusEvent uintptr // TClassdecl
func (x Xflash_events_HTTPStatusEvent) Xstatus_goget()  int { return 0; }
func (x Xflash_events_HTTPStatusEvent) Xstatus_goset(int){}
func Xflash_events_HTTPStatusEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xstatus int) Xflash_events_HTTPStatusEvent { return Xflash_events_HTTPStatusEvent(0); }
func Xflash_events_IOErrorEvent_DISK___ERROR_goget()  string { return ""; }
func Xflash_events_IOErrorEvent_DISK___ERROR_goset(string){}
func Xflash_events_IOErrorEvent_IO___ERROR_goget()  string { return ""; }
func Xflash_events_IOErrorEvent_IO___ERROR_goset(string){}
func Xflash_events_IOErrorEvent_NETWORK___ERROR_goget()  string { return ""; }
func Xflash_events_IOErrorEvent_NETWORK___ERROR_goset(string){}
func Xflash_events_IOErrorEvent_VERIFY___ERROR_goget()  string { return ""; }
func Xflash_events_IOErrorEvent_VERIFY___ERROR_goset(string){}
type Xflash_events_IOErrorEvent uintptr // TClassdecl
func Xflash_events_IOErrorEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xtext string,Xid int) Xflash_events_IOErrorEvent { return Xflash_events_IOErrorEvent(0); }
func Xflash_events_KeyboardEvent_KEY___DOWN_goget()  string { return ""; }
func Xflash_events_KeyboardEvent_KEY___DOWN_goset(string){}
func Xflash_events_KeyboardEvent_KEY___UP_goget()  string { return ""; }
func Xflash_events_KeyboardEvent_KEY___UP_goset(string){}
type Xflash_events_KeyboardEvent uintptr // TClassdecl
func (x Xflash_events_KeyboardEvent) XupdateAfterEvent() {}
func (x Xflash_events_KeyboardEvent) XshiftKey_goget()  bool { return false; }
func (x Xflash_events_KeyboardEvent) XshiftKey_goset(bool){}
func (x Xflash_events_KeyboardEvent) XkeyCode_goget()  int { return 0; }
func (x Xflash_events_KeyboardEvent) XkeyCode_goset(int){}
func (x Xflash_events_KeyboardEvent) XctrlKey_goget()  bool { return false; }
func (x Xflash_events_KeyboardEvent) XctrlKey_goset(bool){}
func (x Xflash_events_KeyboardEvent) XcharCode_goget()  int { return 0; }
func (x Xflash_events_KeyboardEvent) XcharCode_goset(int){}
func (x Xflash_events_KeyboardEvent) XaltKey_goget()  bool { return false; }
func (x Xflash_events_KeyboardEvent) XaltKey_goset(bool){}
func Xflash_events_KeyboardEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,XcharCodeValue int,XkeyCodeValue int) Xflash_events_KeyboardEvent { return Xflash_events_KeyboardEvent(0); }
func Xflash_events_MouseEvent_CLICK_goget()  string { return ""; }
func Xflash_events_MouseEvent_CLICK_goset(string){}
func Xflash_events_MouseEvent_DOUBLE___CLICK_goget()  string { return ""; }
func Xflash_events_MouseEvent_DOUBLE___CLICK_goset(string){}
func Xflash_events_MouseEvent_MOUSE___DOWN_goget()  string { return ""; }
func Xflash_events_MouseEvent_MOUSE___DOWN_goset(string){}
func Xflash_events_MouseEvent_MOUSE___MOVE_goget()  string { return ""; }
func Xflash_events_MouseEvent_MOUSE___MOVE_goset(string){}
func Xflash_events_MouseEvent_MOUSE___OUT_goget()  string { return ""; }
func Xflash_events_MouseEvent_MOUSE___OUT_goset(string){}
func Xflash_events_MouseEvent_MOUSE___OVER_goget()  string { return ""; }
func Xflash_events_MouseEvent_MOUSE___OVER_goset(string){}
func Xflash_events_MouseEvent_MOUSE___UP_goget()  string { return ""; }
func Xflash_events_MouseEvent_MOUSE___UP_goset(string){}
func Xflash_events_MouseEvent_MOUSE___WHEEL_goget()  string { return ""; }
func Xflash_events_MouseEvent_MOUSE___WHEEL_goset(string){}
func Xflash_events_MouseEvent_ROLL___OUT_goget()  string { return ""; }
func Xflash_events_MouseEvent_ROLL___OUT_goset(string){}
func Xflash_events_MouseEvent_ROLL___OVER_goget()  string { return ""; }
func Xflash_events_MouseEvent_ROLL___OVER_goset(string){}
type Xflash_events_MouseEvent uintptr // TClassdecl
func (x Xflash_events_MouseEvent) XupdateAfterEvent() {}
func (x Xflash_events_MouseEvent) XstageY_goget()  float64 { return 0; }
func (x Xflash_events_MouseEvent) XstageY_goset(float64){}
func (x Xflash_events_MouseEvent) XstageX_goget()  float64 { return 0; }
func (x Xflash_events_MouseEvent) XstageX_goset(float64){}
func (x Xflash_events_MouseEvent) XshiftKey_goget()  bool { return false; }
func (x Xflash_events_MouseEvent) XshiftKey_goset(bool){}
func (x Xflash_events_MouseEvent) XrelatedObject_goget()  Xflash_display_InteractiveObject { return 0; }
func (x Xflash_events_MouseEvent) XrelatedObject_goset(Xflash_display_InteractiveObject){}
func (x Xflash_events_MouseEvent) XlocalY_goget()  float64 { return 0; }
func (x Xflash_events_MouseEvent) XlocalY_goset(float64){}
func (x Xflash_events_MouseEvent) XlocalX_goget()  float64 { return 0; }
func (x Xflash_events_MouseEvent) XlocalX_goset(float64){}
func (x Xflash_events_MouseEvent) Xdelta_goget()  int { return 0; }
func (x Xflash_events_MouseEvent) Xdelta_goset(int){}
func (x Xflash_events_MouseEvent) XctrlKey_goget()  bool { return false; }
func (x Xflash_events_MouseEvent) XctrlKey_goset(bool){}
func (x Xflash_events_MouseEvent) XbuttonDown_goget()  bool { return false; }
func (x Xflash_events_MouseEvent) XbuttonDown_goset(bool){}
func (x Xflash_events_MouseEvent) XaltKey_goget()  bool { return false; }
func (x Xflash_events_MouseEvent) XaltKey_goset(bool){}
func Xflash_events_MouseEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,XlocalX float64,XlocalY float64,XrelatedObject Xflash_display_InteractiveObject,XctrlKey bool,XaltKey bool,XshiftKey bool,XbuttonDown bool,Xdelta int) Xflash_events_MouseEvent { return Xflash_events_MouseEvent(0); }
func Xflash_events_ProgressEvent_PROGRESS_goget()  string { return ""; }
func Xflash_events_ProgressEvent_PROGRESS_goset(string){}
func Xflash_events_ProgressEvent_SOCKET___DATA_goget()  string { return ""; }
func Xflash_events_ProgressEvent_SOCKET___DATA_goset(string){}
type Xflash_events_ProgressEvent uintptr // TClassdecl
func (x Xflash_events_ProgressEvent) XbytesTotal_goget()  float64 { return 0; }
func (x Xflash_events_ProgressEvent) XbytesTotal_goset(float64){}
func (x Xflash_events_ProgressEvent) XbytesLoaded_goget()  float64 { return 0; }
func (x Xflash_events_ProgressEvent) XbytesLoaded_goset(float64){}
func Xflash_events_ProgressEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,XbytesLoaded float64,XbytesTotal float64) Xflash_events_ProgressEvent { return Xflash_events_ProgressEvent(0); }
func Xflash_events_SampleDataEvent_SAMPLE___DATA_goget()  string { return ""; }
func Xflash_events_SampleDataEvent_SAMPLE___DATA_goset(string){}
type Xflash_events_SampleDataEvent uintptr // TClassdecl
func (x Xflash_events_SampleDataEvent) Xposition_goget()  float64 { return 0; }
func (x Xflash_events_SampleDataEvent) Xposition_goset(float64){}
func (x Xflash_events_SampleDataEvent) Xdata_goget()  Xflash_utils_ByteArray { return 0; }
func (x Xflash_events_SampleDataEvent) Xdata_goset(Xflash_utils_ByteArray){}
func Xflash_events_SampleDataEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xtheposition float64,Xthedata Xflash_utils_ByteArray) Xflash_events_SampleDataEvent { return Xflash_events_SampleDataEvent(0); }
func Xflash_events_SecurityErrorEvent_SECURITY___ERROR_goget()  string { return ""; }
func Xflash_events_SecurityErrorEvent_SECURITY___ERROR_goset(string){}
type Xflash_events_SecurityErrorEvent uintptr // TClassdecl
func Xflash_events_SecurityErrorEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xtext string,Xid int) Xflash_events_SecurityErrorEvent { return Xflash_events_SecurityErrorEvent(0); }
func Xflash_events_TimerEvent_TIMER_goget()  string { return ""; }
func Xflash_events_TimerEvent_TIMER_goset(string){}
func Xflash_events_TimerEvent_TIMER___COMPLETE_goget()  string { return ""; }
func Xflash_events_TimerEvent_TIMER___COMPLETE_goset(string){}
type Xflash_events_TimerEvent uintptr // TClassdecl
func (x Xflash_events_TimerEvent) XupdateAfterEvent() {}
func Xflash_events_TimerEvent_new(Xtype string,Xbubbles bool,Xcancelable bool) Xflash_events_TimerEvent { return Xflash_events_TimerEvent(0); }
type Xflash_events_TouchEvent uintptr // TClassdecl
func (x Xflash_events_TouchEvent) XupdateAfterEvent() {}
func (x Xflash_events_TouchEvent) XtouchPointID_goget()  int { return 0; }
func (x Xflash_events_TouchEvent) XtouchPointID_goset(int){}
func (x Xflash_events_TouchEvent) XstageY_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) XstageY_goset(float64){}
func (x Xflash_events_TouchEvent) XstageX_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) XstageX_goset(float64){}
func (x Xflash_events_TouchEvent) XsizeY_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) XsizeY_goset(float64){}
func (x Xflash_events_TouchEvent) XsizeX_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) XsizeX_goset(float64){}
func (x Xflash_events_TouchEvent) XshiftKey_goget()  bool { return false; }
func (x Xflash_events_TouchEvent) XshiftKey_goset(bool){}
func (x Xflash_events_TouchEvent) XrelatedObject_goget()  Xflash_display_InteractiveObject { return 0; }
func (x Xflash_events_TouchEvent) XrelatedObject_goset(Xflash_display_InteractiveObject){}
func (x Xflash_events_TouchEvent) Xpressure_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) Xpressure_goset(float64){}
func (x Xflash_events_TouchEvent) XlocalY_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) XlocalY_goset(float64){}
func (x Xflash_events_TouchEvent) XlocalX_goget()  float64 { return 0; }
func (x Xflash_events_TouchEvent) XlocalX_goset(float64){}
func (x Xflash_events_TouchEvent) XisRelatedObjectInaccessible_goget()  bool { return false; }
func (x Xflash_events_TouchEvent) XisRelatedObjectInaccessible_goset(bool){}
func (x Xflash_events_TouchEvent) XisPrimaryTouchPoint_goget()  bool { return false; }
func (x Xflash_events_TouchEvent) XisPrimaryTouchPoint_goset(bool){}
func (x Xflash_events_TouchEvent) XctrlKey_goget()  bool { return false; }
func (x Xflash_events_TouchEvent) XctrlKey_goset(bool){}
func (x Xflash_events_TouchEvent) XaltKey_goget()  bool { return false; }
func (x Xflash_events_TouchEvent) XaltKey_goset(bool){}
func Xflash_events_TouchEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,XtouchPointID int,XisPrimaryTouchPoint bool,XlocalX float64,XlocalY float64,XsizeX float64,XsizeY float64,Xpressure float64,XrelatedObject Xflash_display_InteractiveObject,XctrlKey bool,XaltKey bool,XshiftKey bool) Xflash_events_TouchEvent { return Xflash_events_TouchEvent(0); }
type Xflash_display_IBitmapDrawable uintptr // TClassdecl
type Xflash_display_DisplayObject uintptr // TClassdecl
func (x Xflash_display_DisplayObject) XlocalToGlobal(Xpoint Xflash_geom_Point) Xflash_geom_Point { return 0; }
func (x Xflash_display_DisplayObject) XhitTestPoint(Xx float64,Xy float64,XshapeFlag bool) bool { return false; }
func (x Xflash_display_DisplayObject) XhitTestObject(Xobj Xflash_display_DisplayObject) bool { return false; }
func (x Xflash_display_DisplayObject) XglobalToLocal(Xpoint Xflash_geom_Point) Xflash_geom_Point { return 0; }
func (x Xflash_display_DisplayObject) XgetRect(XtargetCoordinateSpace Xflash_display_DisplayObject) Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_DisplayObject) XgetBounds(XtargetCoordinateSpace Xflash_display_DisplayObject) Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_DisplayObject) Xy_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) Xy_goset(float64){}
func (x Xflash_display_DisplayObject) Xx_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) Xx_goset(float64){}
func (x Xflash_display_DisplayObject) Xwidth_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) Xwidth_goset(float64){}
func (x Xflash_display_DisplayObject) Xvisible_goget()  bool { return false; }
func (x Xflash_display_DisplayObject) Xvisible_goset(bool){}
func (x Xflash_display_DisplayObject) Xtransform_goget()  Xflash_geom_Transform { return 0; }
func (x Xflash_display_DisplayObject) Xtransform_goset(Xflash_geom_Transform){}
func (x Xflash_display_DisplayObject) Xstage_goget()  Xflash_display_Stage { return 0; }
func (x Xflash_display_DisplayObject) Xstage_goset(Xflash_display_Stage){}
func (x Xflash_display_DisplayObject) XscrollRect_goget()  Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_DisplayObject) XscrollRect_goset(Xflash_geom_Rectangle){}
func (x Xflash_display_DisplayObject) XscaleY_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) XscaleY_goset(float64){}
func (x Xflash_display_DisplayObject) XscaleX_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) XscaleX_goset(float64){}
func (x Xflash_display_DisplayObject) Xscale9Grid_goget()  Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_DisplayObject) Xscale9Grid_goset(Xflash_geom_Rectangle){}
func (x Xflash_display_DisplayObject) Xrotation_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) Xrotation_goset(float64){}
func (x Xflash_display_DisplayObject) Xroot_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObject) Xroot_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_DisplayObject) Xparent_goget()  Xflash_display_DisplayObjectContainer { return 0; }
func (x Xflash_display_DisplayObject) Xparent_goset(Xflash_display_DisplayObjectContainer){}
func (x Xflash_display_DisplayObject) XopaqueBackground_goget()  uintptr { return 0; }
func (x Xflash_display_DisplayObject) XopaqueBackground_goset(uintptr){}
func (x Xflash_display_DisplayObject) Xname_goget()  string { return ""; }
func (x Xflash_display_DisplayObject) Xname_goset(string){}
func (x Xflash_display_DisplayObject) XmouseY_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) XmouseY_goset(float64){}
func (x Xflash_display_DisplayObject) XmouseX_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) XmouseX_goset(float64){}
func (x Xflash_display_DisplayObject) Xmask_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObject) Xmask_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_DisplayObject) XloaderInfo_goget()  Xflash_display_LoaderInfo { return 0; }
func (x Xflash_display_DisplayObject) XloaderInfo_goset(Xflash_display_LoaderInfo){}
func (x Xflash_display_DisplayObject) Xheight_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) Xheight_goset(float64){}
func (x Xflash_display_DisplayObject) Xfilters_goget()  XArray { return 0; }
func (x Xflash_display_DisplayObject) Xfilters_goset(XArray){}
func (x Xflash_display_DisplayObject) XcacheAsBitmap_goget()  bool { return false; }
func (x Xflash_display_DisplayObject) XcacheAsBitmap_goset(bool){}
func (x Xflash_display_DisplayObject) XblendMode_goget()  Xflash_display_BlendMode { return 0; }
func (x Xflash_display_DisplayObject) XblendMode_goset(Xflash_display_BlendMode){}
func (x Xflash_display_DisplayObject) Xalpha_goget()  float64 { return 0; }
func (x Xflash_display_DisplayObject) Xalpha_goset(float64){}
type Xflash_display_Bitmap uintptr // TClassdecl
func (x Xflash_display_Bitmap) Xsmoothing_goget()  bool { return false; }
func (x Xflash_display_Bitmap) Xsmoothing_goset(bool){}
func (x Xflash_display_Bitmap) XpixelSnapping_goget()  Xflash_display_PixelSnapping { return 0; }
func (x Xflash_display_Bitmap) XpixelSnapping_goset(Xflash_display_PixelSnapping){}
func (x Xflash_display_Bitmap) XbitmapData_goget()  Xflash_display_BitmapData { return 0; }
func (x Xflash_display_Bitmap) XbitmapData_goset(Xflash_display_BitmapData){}
func Xflash_display_Bitmap_new(XbitmapData Xflash_display_BitmapData,XpixelSnapping Xflash_display_PixelSnapping,Xsmoothing bool) Xflash_display_Bitmap { return Xflash_display_Bitmap(0); }
type Xflash_display_BitmapData uintptr // TClassdecl
func (x Xflash_display_BitmapData) Xunlock(XchangeRect Xflash_geom_Rectangle) {}
func (x Xflash_display_BitmapData) Xthreshold(XsourceBitmapData Xflash_display_BitmapData,XsourceRect Xflash_geom_Rectangle,XdestPoint Xflash_geom_Point,Xoperation string,Xthreshold int,Xcolor uintptr,Xmask uintptr,XcopySource bool) uintptr { return 0; }
func (x Xflash_display_BitmapData) XsetVector(Xrect Xflash_geom_Rectangle,XinputVector Xflash_Vector) {}
func (x Xflash_display_BitmapData) XsetPixels(Xrect Xflash_geom_Rectangle,XinputByteArray Xflash_utils_ByteArray) {}
func (x Xflash_display_BitmapData) XsetPixel32(Xx int,Xy int,Xcolor int) {}
func (x Xflash_display_BitmapData) XsetPixel(Xx int,Xy int,Xcolor int) {}
func (x Xflash_display_BitmapData) Xscroll(Xx int,Xy int) {}
func (x Xflash_display_BitmapData) XperlinNoise(XbaseX float64,XbaseY float64,XnumOctaves uintptr,XrandomSeed int,Xstitch bool,XfractalNoise bool,XchannelOptions uintptr,XgrayScale bool,Xoffsets XArray) {}
func (x Xflash_display_BitmapData) Xnoise(XrandomSeed int,Xlow uintptr,Xhigh uintptr,XchannelOptions uintptr,XgrayScale bool) {}
func (x Xflash_display_BitmapData) Xlock() {}
func (x Xflash_display_BitmapData) XgetVector(Xrect Xflash_geom_Rectangle) Xflash_Vector { return 0; }
func (x Xflash_display_BitmapData) XgetPixels(Xrect Xflash_geom_Rectangle) Xflash_utils_ByteArray { return 0; }
func (x Xflash_display_BitmapData) XgetPixel32(Xx int,Xy int) uintptr { return 0; }
func (x Xflash_display_BitmapData) XgetPixel(Xx int,Xy int) uintptr { return 0; }
func (x Xflash_display_BitmapData) XgetColorBoundsRect(Xmask uintptr,Xcolor uintptr,XfindColor bool) Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_BitmapData) XgenerateFilterRect(XsourceRect Xflash_geom_Rectangle,Xfilter Xflash_filters_BitmapFilter) Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_BitmapData) XfloodFill(Xx int,Xy int,Xcolor uintptr) {}
func (x Xflash_display_BitmapData) XfillRect(Xrect Xflash_geom_Rectangle,Xcolor uintptr) {}
func (x Xflash_display_BitmapData) Xdraw(Xsource Xflash_display_IBitmapDrawable,Xmatrix Xflash_geom_Matrix,XcolorTransform Xflash_geom_ColorTransform,XblendMode Xflash_display_BlendMode,XclipRect Xflash_geom_Rectangle,Xsmoothing bool) {}
func (x Xflash_display_BitmapData) Xdispose() {}
func (x Xflash_display_BitmapData) XcopyPixels(XsourceBitmapData Xflash_display_BitmapData,XsourceRect Xflash_geom_Rectangle,XdestPoint Xflash_geom_Point,XalphaBitmapData Xflash_display_BitmapData,XalphaPoint Xflash_geom_Point,XmergeAlpha bool) {}
func (x Xflash_display_BitmapData) XcopyChannel(XsourceBitmapData Xflash_display_BitmapData,XsourceRect Xflash_geom_Rectangle,XdestPoint Xflash_geom_Point,XsourceChannel uintptr,XdestChannel uintptr) {}
func (x Xflash_display_BitmapData) XcolorTransform(Xrect Xflash_geom_Rectangle,XcolorTransform Xflash_geom_ColorTransform) {}
func (x Xflash_display_BitmapData) Xclone() Xflash_display_BitmapData { return 0; }
func (x Xflash_display_BitmapData) XapplyFilter(XsourceBitmapData Xflash_display_BitmapData,XsourceRect Xflash_geom_Rectangle,XdestPoint Xflash_geom_Point,Xfilter Xflash_filters_BitmapFilter) {}
func (x Xflash_display_BitmapData) Xwidth_goget()  int { return 0; }
func (x Xflash_display_BitmapData) Xwidth_goset(int){}
func (x Xflash_display_BitmapData) Xtransparent_goget()  bool { return false; }
func (x Xflash_display_BitmapData) Xtransparent_goset(bool){}
func (x Xflash_display_BitmapData) Xrect_goget()  Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_BitmapData) Xrect_goset(Xflash_geom_Rectangle){}
func (x Xflash_display_BitmapData) Xheight_goget()  int { return 0; }
func (x Xflash_display_BitmapData) Xheight_goset(int){}
func Xflash_display_BitmapData_new(Xwidth int,Xheight int,Xtransparent bool,XfillColor uintptr) Xflash_display_BitmapData { return Xflash_display_BitmapData(0); }
func Xflash_display_BitmapDataChannel_ALPHA_goget()  int { return 0; }
func Xflash_display_BitmapDataChannel_ALPHA_goset(int){}
func Xflash_display_BitmapDataChannel_BLUE_goget()  int { return 0; }
func Xflash_display_BitmapDataChannel_BLUE_goset(int){}
func Xflash_display_BitmapDataChannel_GREEN_goget()  int { return 0; }
func Xflash_display_BitmapDataChannel_GREEN_goset(int){}
func Xflash_display_BitmapDataChannel_RED_goget()  int { return 0; }
func Xflash_display_BitmapDataChannel_RED_goset(int){}
type Xflash_display_BitmapDataChannel uintptr // TClassdecl
type Xflash_display_BlendMode uintptr // TEnumdecl
type Xflash_display_CapsStyle uintptr // TEnumdecl
type Xflash_display_InteractiveObject uintptr // TClassdecl
func (x Xflash_display_InteractiveObject) XtabIndex_goget()  int { return 0; }
func (x Xflash_display_InteractiveObject) XtabIndex_goset(int){}
func (x Xflash_display_InteractiveObject) XtabEnabled_goget()  bool { return false; }
func (x Xflash_display_InteractiveObject) XtabEnabled_goset(bool){}
func (x Xflash_display_InteractiveObject) XmouseEnabled_goget()  bool { return false; }
func (x Xflash_display_InteractiveObject) XmouseEnabled_goset(bool){}
func (x Xflash_display_InteractiveObject) XfocusRect_goget()  interface{} { return nil; }
func (x Xflash_display_InteractiveObject) XfocusRect_goset(interface{}){}
func (x Xflash_display_InteractiveObject) XdoubleClickEnabled_goget()  bool { return false; }
func (x Xflash_display_InteractiveObject) XdoubleClickEnabled_goset(bool){}
func Xflash_display_InteractiveObject_new() Xflash_display_InteractiveObject { return Xflash_display_InteractiveObject(0); }
type Xflash_display_DisplayObjectContainer uintptr // TClassdecl
func (x Xflash_display_DisplayObjectContainer) XswapChildrenAt(Xindex1 int,Xindex2 int) {}
func (x Xflash_display_DisplayObjectContainer) XswapChildren(Xchild1 Xflash_display_DisplayObject,Xchild2 Xflash_display_DisplayObject) {}
func (x Xflash_display_DisplayObjectContainer) XsetChildIndex(Xchild Xflash_display_DisplayObject,Xindex int) {}
func (x Xflash_display_DisplayObjectContainer) XremoveChildAt(Xindex int) Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObjectContainer) XremoveChild(Xchild Xflash_display_DisplayObject) Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObjectContainer) XgetObjectsUnderPoint(Xpoint Xflash_geom_Point) XArray { return 0; }
func (x Xflash_display_DisplayObjectContainer) XgetChildIndex(Xchild Xflash_display_DisplayObject) int { return 0; }
func (x Xflash_display_DisplayObjectContainer) XgetChildByName(Xname string) Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObjectContainer) XgetChildAt(Xindex int) Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObjectContainer) Xcontains(Xchild Xflash_display_DisplayObject) bool { return false; }
func (x Xflash_display_DisplayObjectContainer) XareInaccessibleObjectsUnderPoint(Xpoint Xflash_geom_Point) bool { return false; }
func (x Xflash_display_DisplayObjectContainer) XaddChildAt(Xchild Xflash_display_DisplayObject,Xindex int) Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObjectContainer) XaddChild(Xchild Xflash_display_DisplayObject) Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_DisplayObjectContainer) XtabChildren_goget()  bool { return false; }
func (x Xflash_display_DisplayObjectContainer) XtabChildren_goset(bool){}
func (x Xflash_display_DisplayObjectContainer) XnumChildren_goget()  int { return 0; }
func (x Xflash_display_DisplayObjectContainer) XnumChildren_goset(int){}
func (x Xflash_display_DisplayObjectContainer) XmouseChildren_goget()  bool { return false; }
func (x Xflash_display_DisplayObjectContainer) XmouseChildren_goset(bool){}
type Xflash_display_GradientType uintptr // TEnumdecl
type Xflash_display_Graphics uintptr // TClassdecl
func (x Xflash_display_Graphics) XmoveTo(Xx float64,Xy float64) {}
func (x Xflash_display_Graphics) XlineTo(Xx float64,Xy float64) {}
func (x Xflash_display_Graphics) XlineStyle(Xthickness float64,Xcolor int,Xalpha float64,XpixelHinting bool,XscaleMode Xflash_display_LineScaleMode,Xcaps Xflash_display_CapsStyle,Xjoints Xflash_display_JointStyle,XmiterLimit float64) {}
func (x Xflash_display_Graphics) XlineGradientStyle(Xtype Xflash_display_GradientType,Xcolors XArray,Xalphas XArray,Xratios XArray,Xmatrix Xflash_geom_Matrix,XspreadMethod Xflash_display_SpreadMethod,XinterpolationMethod Xflash_display_InterpolationMethod,XfocalPointRatio float64) {}
func (x Xflash_display_Graphics) XlineBitmapStyle(Xbitmap Xflash_display_BitmapData,Xmatrix Xflash_geom_Matrix,Xrepeat bool,Xsmooth bool) {}
func (x Xflash_display_Graphics) XendFill() {}
func (x Xflash_display_Graphics) XdrawTriangles(Xvertices Xflash_Vector,Xindices Xflash_Vector,XuvtData Xflash_Vector,Xculling Xflash_display_TriangleCulling) {}
func (x Xflash_display_Graphics) XdrawRoundRectComplex(Xx float64,Xy float64,Xwidth float64,Xheight float64,XtopLeftRadius float64,XtopRightRadius float64,XbottomLeftRadius float64,XbottomRightRadius float64) {}
func (x Xflash_display_Graphics) XdrawRoundRect(Xx float64,Xy float64,Xwidth float64,Xheight float64,XellipseWidth float64,XellipseHeight float64) {}
func (x Xflash_display_Graphics) XdrawRect(Xx float64,Xy float64,Xwidth float64,Xheight float64) {}
func (x Xflash_display_Graphics) XdrawPath(Xcommands Xflash_Vector,Xdata Xflash_Vector,Xwinding Xflash_display_GraphicsPathWinding) {}
func (x Xflash_display_Graphics) XdrawGraphicsData(XgraphicsData Xflash_Vector) {}
func (x Xflash_display_Graphics) XdrawEllipse(Xx float64,Xy float64,Xwidth float64,Xheight float64) {}
func (x Xflash_display_Graphics) XdrawCircle(Xx float64,Xy float64,Xradius float64) {}
func (x Xflash_display_Graphics) XcurveTo(XcontrolX float64,XcontrolY float64,XanchorX float64,XanchorY float64) {}
func (x Xflash_display_Graphics) Xclear() {}
func (x Xflash_display_Graphics) XbeginGradientFill(Xtype Xflash_display_GradientType,Xcolors XArray,Xalphas XArray,Xratios XArray,Xmatrix Xflash_geom_Matrix,XspreadMethod Xflash_display_SpreadMethod,XinterpolationMethod Xflash_display_InterpolationMethod,XfocalPointRatio float64) {}
func (x Xflash_display_Graphics) XbeginFill(Xcolor int,Xalpha float64) {}
func (x Xflash_display_Graphics) XbeginBitmapFill(Xbitmap Xflash_display_BitmapData,Xmatrix Xflash_geom_Matrix,Xrepeat bool,Xsmooth bool) {}
func Xflash_display_Graphics_new() Xflash_display_Graphics { return Xflash_display_Graphics(0); }
type Xflash_display_IGraphicsFill uintptr // TClassdecl
type Xflash_display_IGraphicsData uintptr // TClassdecl
type Xflash_display_GraphicsBitmapFill uintptr // TClassdecl
func (x Xflash_display_GraphicsBitmapFill) Xsmooth_goget()  bool { return false; }
func (x Xflash_display_GraphicsBitmapFill) Xsmooth_goset(bool){}
func (x Xflash_display_GraphicsBitmapFill) Xrepeat_goget()  bool { return false; }
func (x Xflash_display_GraphicsBitmapFill) Xrepeat_goset(bool){}
func (x Xflash_display_GraphicsBitmapFill) Xmatrix_goget()  Xflash_geom_Matrix { return 0; }
func (x Xflash_display_GraphicsBitmapFill) Xmatrix_goset(Xflash_geom_Matrix){}
func (x Xflash_display_GraphicsBitmapFill) XbitmapData_goget()  Xflash_display_BitmapData { return 0; }
func (x Xflash_display_GraphicsBitmapFill) XbitmapData_goset(Xflash_display_BitmapData){}
func Xflash_display_GraphicsBitmapFill_new(XbitmapData Xflash_display_BitmapData,Xmatrix Xflash_geom_Matrix,Xrepeat bool,Xsmooth bool) Xflash_display_GraphicsBitmapFill { return Xflash_display_GraphicsBitmapFill(0); }
type Xflash_display_GraphicsEndFill uintptr // TClassdecl
func Xflash_display_GraphicsEndFill_new() Xflash_display_GraphicsEndFill { return Xflash_display_GraphicsEndFill(0); }
type Xflash_display_GraphicsGradientFill uintptr // TClassdecl
func (x Xflash_display_GraphicsGradientFill) Xtype_goget()  Xflash_display_GradientType { return 0; }
func (x Xflash_display_GraphicsGradientFill) Xtype_goset(Xflash_display_GradientType){}
func (x Xflash_display_GraphicsGradientFill) XspreadMethod_goget()  Xflash_display_SpreadMethod { return 0; }
func (x Xflash_display_GraphicsGradientFill) XspreadMethod_goset(Xflash_display_SpreadMethod){}
func (x Xflash_display_GraphicsGradientFill) Xratios_goget()  XArray { return 0; }
func (x Xflash_display_GraphicsGradientFill) Xratios_goset(XArray){}
func (x Xflash_display_GraphicsGradientFill) Xmatrix_goget()  Xflash_geom_Matrix { return 0; }
func (x Xflash_display_GraphicsGradientFill) Xmatrix_goset(Xflash_geom_Matrix){}
func (x Xflash_display_GraphicsGradientFill) XinterpolationMethod_goget()  Xflash_display_InterpolationMethod { return 0; }
func (x Xflash_display_GraphicsGradientFill) XinterpolationMethod_goset(Xflash_display_InterpolationMethod){}
func (x Xflash_display_GraphicsGradientFill) XfocalPointRatio_goget()  float64 { return 0; }
func (x Xflash_display_GraphicsGradientFill) XfocalPointRatio_goset(float64){}
func (x Xflash_display_GraphicsGradientFill) Xcolors_goget()  XArray { return 0; }
func (x Xflash_display_GraphicsGradientFill) Xcolors_goset(XArray){}
func (x Xflash_display_GraphicsGradientFill) Xalphas_goget()  XArray { return 0; }
func (x Xflash_display_GraphicsGradientFill) Xalphas_goset(XArray){}
func Xflash_display_GraphicsGradientFill_new(Xtype Xflash_display_GradientType,Xcolors XArray,Xalphas XArray,Xratios XArray,Xmatrix Xflash_geom_Matrix,XspreadMethod Xflash_display_SpreadMethod,XinterpolationMethod Xflash_display_InterpolationMethod,XfocalPointRatio float64) Xflash_display_GraphicsGradientFill { return Xflash_display_GraphicsGradientFill(0); }
type Xflash_display_GraphicsPath uintptr // TClassdecl
func (x Xflash_display_GraphicsPath) XwideMoveTo(Xx float64,Xy float64) {}
func (x Xflash_display_GraphicsPath) XwideLineTo(Xx float64,Xy float64) {}
func (x Xflash_display_GraphicsPath) XmoveTo(Xx float64,Xy float64) {}
func (x Xflash_display_GraphicsPath) XlineTo(Xx float64,Xy float64) {}
func (x Xflash_display_GraphicsPath) XcurveTo(XcontrolX float64,XcontrolY float64,XanchorX float64,XanchorY float64) {}
func (x Xflash_display_GraphicsPath) Xwinding_goget()  Xflash_display_GraphicsPathWinding { return 0; }
func (x Xflash_display_GraphicsPath) Xwinding_goset(Xflash_display_GraphicsPathWinding){}
func (x Xflash_display_GraphicsPath) Xdata_goget()  Xflash_Vector { return 0; }
func (x Xflash_display_GraphicsPath) Xdata_goset(Xflash_Vector){}
func (x Xflash_display_GraphicsPath) Xcommands_goget()  Xflash_Vector { return 0; }
func (x Xflash_display_GraphicsPath) Xcommands_goset(Xflash_Vector){}
func Xflash_display_GraphicsPath_new(Xcommands Xflash_Vector,Xdata Xflash_Vector,Xwinding Xflash_display_GraphicsPathWinding) Xflash_display_GraphicsPath { return Xflash_display_GraphicsPath(0); }
type Xflash_display_GraphicsPathCommand uintptr // TEnumdecl
type Xflash_display_GraphicsPathWinding uintptr // TEnumdecl
type Xflash_display_GraphicsSolidFill uintptr // TClassdecl
func (x Xflash_display_GraphicsSolidFill) Xcolor_goget()  int { return 0; }
func (x Xflash_display_GraphicsSolidFill) Xcolor_goset(int){}
func (x Xflash_display_GraphicsSolidFill) Xalpha_goget()  float64 { return 0; }
func (x Xflash_display_GraphicsSolidFill) Xalpha_goset(float64){}
func Xflash_display_GraphicsSolidFill_new(Xcolor int,Xalpha float64) Xflash_display_GraphicsSolidFill { return Xflash_display_GraphicsSolidFill(0); }
type Xflash_display_GraphicsStroke uintptr // TClassdecl
func (x Xflash_display_GraphicsStroke) Xthickness_goget()  float64 { return 0; }
func (x Xflash_display_GraphicsStroke) Xthickness_goset(float64){}
func (x Xflash_display_GraphicsStroke) XscaleMode_goget()  Xflash_display_LineScaleMode { return 0; }
func (x Xflash_display_GraphicsStroke) XscaleMode_goset(Xflash_display_LineScaleMode){}
func (x Xflash_display_GraphicsStroke) XpixelHinting_goget()  bool { return false; }
func (x Xflash_display_GraphicsStroke) XpixelHinting_goset(bool){}
func (x Xflash_display_GraphicsStroke) XmiterLimit_goget()  float64 { return 0; }
func (x Xflash_display_GraphicsStroke) XmiterLimit_goset(float64){}
func (x Xflash_display_GraphicsStroke) Xjoints_goget()  Xflash_display_JointStyle { return 0; }
func (x Xflash_display_GraphicsStroke) Xjoints_goset(Xflash_display_JointStyle){}
func (x Xflash_display_GraphicsStroke) Xfill_goget()  Xflash_display_IGraphicsFill { return 0; }
func (x Xflash_display_GraphicsStroke) Xfill_goset(Xflash_display_IGraphicsFill){}
func (x Xflash_display_GraphicsStroke) Xcaps_goget()  Xflash_display_CapsStyle { return 0; }
func (x Xflash_display_GraphicsStroke) Xcaps_goset(Xflash_display_CapsStyle){}
func Xflash_display_GraphicsStroke_new(Xthickness float64,XpixelHinting bool,XscaleMode Xflash_display_LineScaleMode,Xcaps Xflash_display_CapsStyle,Xjoints Xflash_display_JointStyle,XmiterLimit float64,Xfill Xflash_display_IGraphicsFill) Xflash_display_GraphicsStroke { return Xflash_display_GraphicsStroke(0); }
type Xflash_display_InterpolationMethod uintptr // TEnumdecl
type Xflash_display_JointStyle uintptr // TEnumdecl
type Xflash_display_LineScaleMode uintptr // TEnumdecl
type Xflash_display_Loader uintptr // TClassdecl
func (x Xflash_display_Loader) Xunload() {}
func (x Xflash_display_Loader) XloadBytes(Xbytes Xflash_utils_ByteArray,Xcontext Xflash_system_LoaderContext) {}
func (x Xflash_display_Loader) Xload(Xrequest Xflash_net_URLRequest,Xcontext Xflash_system_LoaderContext) {}
func (x Xflash_display_Loader) Xclose() {}
func (x Xflash_display_Loader) XcontentLoaderInfo_goget()  Xflash_display_LoaderInfo { return 0; }
func (x Xflash_display_Loader) XcontentLoaderInfo_goset(Xflash_display_LoaderInfo){}
func (x Xflash_display_Loader) Xcontent_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_Loader) Xcontent_goset(Xflash_display_DisplayObject){}
func Xflash_display_Loader_new() Xflash_display_Loader { return Xflash_display_Loader(0); }
type Xflash_display_LoaderInfo uintptr // TClassdecl
func (x Xflash_display_LoaderInfo) Xwidth_goget()  int { return 0; }
func (x Xflash_display_LoaderInfo) Xwidth_goset(int){}
func (x Xflash_display_LoaderInfo) Xurl_goget()  string { return ""; }
func (x Xflash_display_LoaderInfo) Xurl_goset(string){}
func (x Xflash_display_LoaderInfo) XsharedEvents_goget()  Xflash_events_EventDispatcher { return 0; }
func (x Xflash_display_LoaderInfo) XsharedEvents_goset(Xflash_events_EventDispatcher){}
func (x Xflash_display_LoaderInfo) XsameDomain_goget()  bool { return false; }
func (x Xflash_display_LoaderInfo) XsameDomain_goset(bool){}
func (x Xflash_display_LoaderInfo) XparentAllowsChild_goget()  bool { return false; }
func (x Xflash_display_LoaderInfo) XparentAllowsChild_goset(bool){}
func (x Xflash_display_LoaderInfo) Xparameters_goget()  interface{} { return nil; }
func (x Xflash_display_LoaderInfo) Xparameters_goset(interface{}){}
func (x Xflash_display_LoaderInfo) XloaderURL_goget()  string { return ""; }
func (x Xflash_display_LoaderInfo) XloaderURL_goset(string){}
func (x Xflash_display_LoaderInfo) Xloader_goget()  Xflash_display_Loader { return 0; }
func (x Xflash_display_LoaderInfo) Xloader_goset(Xflash_display_Loader){}
func (x Xflash_display_LoaderInfo) Xheight_goget()  int { return 0; }
func (x Xflash_display_LoaderInfo) Xheight_goset(int){}
func (x Xflash_display_LoaderInfo) XframeRate_goget()  float64 { return 0; }
func (x Xflash_display_LoaderInfo) XframeRate_goset(float64){}
func (x Xflash_display_LoaderInfo) XcontentType_goget()  string { return ""; }
func (x Xflash_display_LoaderInfo) XcontentType_goset(string){}
func (x Xflash_display_LoaderInfo) Xcontent_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_LoaderInfo) Xcontent_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_LoaderInfo) XbytesTotal_goget()  int { return 0; }
func (x Xflash_display_LoaderInfo) XbytesTotal_goset(int){}
func (x Xflash_display_LoaderInfo) XbytesLoaded_goget()  int { return 0; }
func (x Xflash_display_LoaderInfo) XbytesLoaded_goset(int){}
func (x Xflash_display_LoaderInfo) Xbytes_goget()  Xflash_utils_ByteArray { return 0; }
func (x Xflash_display_LoaderInfo) Xbytes_goset(Xflash_utils_ByteArray){}
func (x Xflash_display_LoaderInfo) XapplicationDomain_goget()  Xflash_system_ApplicationDomain { return 0; }
func (x Xflash_display_LoaderInfo) XapplicationDomain_goset(Xflash_system_ApplicationDomain){}
type Xflash_display_Sprite uintptr // TClassdecl
func (x Xflash_display_Sprite) XstopDrag() {}
func (x Xflash_display_Sprite) XstartDrag(XlockCenter bool,Xbounds Xflash_geom_Rectangle) {}
func (x Xflash_display_Sprite) XuseHandCursor_goget()  bool { return false; }
func (x Xflash_display_Sprite) XuseHandCursor_goset(bool){}
func (x Xflash_display_Sprite) XsoundTransform_goget()  Xflash_media_SoundTransform { return 0; }
func (x Xflash_display_Sprite) XsoundTransform_goset(Xflash_media_SoundTransform){}
func (x Xflash_display_Sprite) XhitArea_goget()  Xflash_display_Sprite { return 0; }
func (x Xflash_display_Sprite) XhitArea_goset(Xflash_display_Sprite){}
func (x Xflash_display_Sprite) Xgraphics_goget()  Xflash_display_Graphics { return 0; }
func (x Xflash_display_Sprite) Xgraphics_goset(Xflash_display_Graphics){}
func (x Xflash_display_Sprite) XdropTarget_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_Sprite) XdropTarget_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_Sprite) XbuttonMode_goget()  bool { return false; }
func (x Xflash_display_Sprite) XbuttonMode_goset(bool){}
func Xflash_display_Sprite_new() Xflash_display_Sprite { return Xflash_display_Sprite(0); }
type Xflash_display_MovieClip uintptr // TClassdecl
func (x Xflash_display_MovieClip) Xstop() {}
func (x Xflash_display_MovieClip) XprevScene() {}
func (x Xflash_display_MovieClip) XprevFrame() {}
func (x Xflash_display_MovieClip) Xplay() {}
func (x Xflash_display_MovieClip) XnextScene() {}
func (x Xflash_display_MovieClip) XnextFrame() {}
func (x Xflash_display_MovieClip) XgotoAndStop(Xframe uintptr,Xscene string) {}
func (x Xflash_display_MovieClip) XgotoAndPlay(Xframe uintptr,Xscene string) {}
func (x Xflash_display_MovieClip) XaddFrameScript(Xp1 interface{},Xp2 interface{},Xp3 interface{},Xp4 interface{},Xp5 interface{}) {}
func (x Xflash_display_MovieClip) XtrackAsMenu_goget()  bool { return false; }
func (x Xflash_display_MovieClip) XtrackAsMenu_goset(bool){}
func (x Xflash_display_MovieClip) XtotalFrames_goget()  int { return 0; }
func (x Xflash_display_MovieClip) XtotalFrames_goset(int){}
func (x Xflash_display_MovieClip) XframesLoaded_goget()  int { return 0; }
func (x Xflash_display_MovieClip) XframesLoaded_goset(int){}
func (x Xflash_display_MovieClip) Xenabled_goget()  bool { return false; }
func (x Xflash_display_MovieClip) Xenabled_goset(bool){}
func (x Xflash_display_MovieClip) XcurrentLabel_goget()  string { return ""; }
func (x Xflash_display_MovieClip) XcurrentLabel_goset(string){}
func (x Xflash_display_MovieClip) XcurrentFrame_goget()  int { return 0; }
func (x Xflash_display_MovieClip) XcurrentFrame_goset(int){}
func Xflash_display_MovieClip_new() Xflash_display_MovieClip { return Xflash_display_MovieClip(0); }
type Xflash_display_PixelSnapping uintptr // TEnumdecl
type Xflash_display_Shape uintptr // TClassdecl
func (x Xflash_display_Shape) Xgraphics_goget()  Xflash_display_Graphics { return 0; }
func (x Xflash_display_Shape) Xgraphics_goset(Xflash_display_Graphics){}
func Xflash_display_Shape_new() Xflash_display_Shape { return Xflash_display_Shape(0); }
type Xflash_display_SimpleButton uintptr // TClassdecl
func (x Xflash_display_SimpleButton) XuseHandCursor_goget()  bool { return false; }
func (x Xflash_display_SimpleButton) XuseHandCursor_goset(bool){}
func (x Xflash_display_SimpleButton) XupState_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_SimpleButton) XupState_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_SimpleButton) XtrackAsMenu_goget()  bool { return false; }
func (x Xflash_display_SimpleButton) XtrackAsMenu_goset(bool){}
func (x Xflash_display_SimpleButton) XsoundTransform_goget()  Xflash_media_SoundTransform { return 0; }
func (x Xflash_display_SimpleButton) XsoundTransform_goset(Xflash_media_SoundTransform){}
func (x Xflash_display_SimpleButton) XoverState_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_SimpleButton) XoverState_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_SimpleButton) XhitTestState_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_SimpleButton) XhitTestState_goset(Xflash_display_DisplayObject){}
func (x Xflash_display_SimpleButton) Xenabled_goget()  bool { return false; }
func (x Xflash_display_SimpleButton) Xenabled_goset(bool){}
func (x Xflash_display_SimpleButton) XdownState_goget()  Xflash_display_DisplayObject { return 0; }
func (x Xflash_display_SimpleButton) XdownState_goset(Xflash_display_DisplayObject){}
func Xflash_display_SimpleButton_new(XupState Xflash_display_DisplayObject,XoverState Xflash_display_DisplayObject,XdownState Xflash_display_DisplayObject,XhitTestState Xflash_display_DisplayObject) Xflash_display_SimpleButton { return Xflash_display_SimpleButton(0); }
type Xflash_display_SpreadMethod uintptr // TEnumdecl
type Xflash_display_Stage uintptr // TClassdecl
func (x Xflash_display_Stage) XisFocusInaccessible() bool { return false; }
func (x Xflash_display_Stage) Xinvalidate() {}
func (x Xflash_display_Stage) XstageWidth_goget()  int { return 0; }
func (x Xflash_display_Stage) XstageWidth_goset(int){}
func (x Xflash_display_Stage) XstageHeight_goget()  int { return 0; }
func (x Xflash_display_Stage) XstageHeight_goset(int){}
func (x Xflash_display_Stage) XstageFocusRect_goget()  bool { return false; }
func (x Xflash_display_Stage) XstageFocusRect_goset(bool){}
func (x Xflash_display_Stage) XshowDefaultContextMenu_goget()  bool { return false; }
func (x Xflash_display_Stage) XshowDefaultContextMenu_goset(bool){}
func (x Xflash_display_Stage) XscaleMode_goget()  Xflash_display_StageScaleMode { return 0; }
func (x Xflash_display_Stage) XscaleMode_goset(Xflash_display_StageScaleMode){}
func (x Xflash_display_Stage) Xquality_goget()  Xflash_display_StageQuality { return 0; }
func (x Xflash_display_Stage) Xquality_goset(Xflash_display_StageQuality){}
func (x Xflash_display_Stage) XfullScreenWidth_goget()  int { return 0; }
func (x Xflash_display_Stage) XfullScreenWidth_goset(int){}
func (x Xflash_display_Stage) XfullScreenSourceRect_goget()  Xflash_geom_Rectangle { return 0; }
func (x Xflash_display_Stage) XfullScreenSourceRect_goset(Xflash_geom_Rectangle){}
func (x Xflash_display_Stage) XfullScreenHeight_goget()  int { return 0; }
func (x Xflash_display_Stage) XfullScreenHeight_goset(int){}
func (x Xflash_display_Stage) XframeRate_goget()  float64 { return 0; }
func (x Xflash_display_Stage) XframeRate_goset(float64){}
func (x Xflash_display_Stage) Xfocus_goget()  Xflash_display_InteractiveObject { return 0; }
func (x Xflash_display_Stage) Xfocus_goset(Xflash_display_InteractiveObject){}
func (x Xflash_display_Stage) XdisplayState_goget()  Xflash_display_StageDisplayState { return 0; }
func (x Xflash_display_Stage) XdisplayState_goset(Xflash_display_StageDisplayState){}
func (x Xflash_display_Stage) Xalign_goget()  Xflash_display_StageAlign { return 0; }
func (x Xflash_display_Stage) Xalign_goset(Xflash_display_StageAlign){}
type Xflash_display_StageAlign uintptr // TEnumdecl
type Xflash_display_StageDisplayState uintptr // TEnumdecl
type Xflash_display_StageQuality uintptr // TEnumdecl
type Xflash_display_StageScaleMode uintptr // TEnumdecl
type Xflash_display_TriangleCulling uintptr // TEnumdecl
func Xflash_errors_Error_length_goget()  int { return 0; }
func Xflash_errors_Error_length_goset(int){}
func Xflash_errors_Error_getErrorMessage(Xindex int) string { return ""; }
func Xflash_errors_Error_throwError(Xtype XClass,Xindex int,Xp1 interface{},Xp2 interface{},Xp3 interface{},Xp4 interface{},Xp5 interface{}) interface{} { return nil; }
type Xflash_errors_Error uintptr // TClassdecl
func (x Xflash_errors_Error) XgetStackTrace() string { return ""; }
func (x Xflash_errors_Error) Xname_goget()  interface{} { return nil; }
func (x Xflash_errors_Error) Xname_goset(interface{}){}
func (x Xflash_errors_Error) Xmessage_goget()  interface{} { return nil; }
func (x Xflash_errors_Error) Xmessage_goset(interface{}){}
func (x Xflash_errors_Error) XerrorID_goget()  int { return 0; }
func (x Xflash_errors_Error) XerrorID_goset(int){}
func Xflash_errors_Error_new(Xmessage interface{},Xid interface{}) Xflash_errors_Error { return Xflash_errors_Error(0); }
type Xflash_errors_ArgumentError uintptr // TClassdecl
type Xflash_errors_EOFError uintptr // TClassdecl
func Xflash_errors_EOFError_new(Xmessage string,Xid int) Xflash_errors_EOFError { return Xflash_errors_EOFError(0); }
type Xflash_errors_IllegalOperationError uintptr // TClassdecl
func Xflash_errors_IllegalOperationError_new(Xmessage string,Xid int) Xflash_errors_IllegalOperationError { return Xflash_errors_IllegalOperationError(0); }
type Xflash_errors_RangeError uintptr // TClassdecl
type Xflash_errors_SecurityError uintptr // TClassdecl
type Xflash_errors_TypeError uintptr // TClassdecl
func Xflash_external_ExternalInterface_available_goget()  bool { return false; }
func Xflash_external_ExternalInterface_available_goset(bool){}
func Xflash_external_ExternalInterface_marshallExceptions_goget()  bool { return false; }
func Xflash_external_ExternalInterface_marshallExceptions_goset(bool){}
func Xflash_external_ExternalInterface_objectID_goget()  string { return ""; }
func Xflash_external_ExternalInterface_objectID_goset(string){}
func Xflash_external_ExternalInterface_addCallback(XfunctionName string,Xclosure interface{}) {}
func Xflash_external_ExternalInterface_call(XfunctionName string,Xp1 interface{},Xp2 interface{},Xp3 interface{},Xp4 interface{},Xp5 interface{}) interface{} { return nil; }
type Xflash_external_ExternalInterface uintptr // TClassdecl
type Xflash_filters_BitmapFilter uintptr // TClassdecl
func (x Xflash_filters_BitmapFilter) Xclone() Xflash_filters_BitmapFilter { return 0; }
func Xflash_filters_BitmapFilter_new() Xflash_filters_BitmapFilter { return Xflash_filters_BitmapFilter(0); }
func Xflash_filters_BitmapFilterQuality_HIGH_goget()  int { return 0; }
func Xflash_filters_BitmapFilterQuality_HIGH_goset(int){}
func Xflash_filters_BitmapFilterQuality_LOW_goget()  int { return 0; }
func Xflash_filters_BitmapFilterQuality_LOW_goset(int){}
func Xflash_filters_BitmapFilterQuality_MEDIUM_goget()  int { return 0; }
func Xflash_filters_BitmapFilterQuality_MEDIUM_goset(int){}
type Xflash_filters_BitmapFilterQuality uintptr // TClassdecl
type Xflash_filters_BitmapFilterType uintptr // TEnumdecl
type Xflash_filters_BlurFilter uintptr // TClassdecl
func (x Xflash_filters_BlurFilter) Xquality_goget()  int { return 0; }
func (x Xflash_filters_BlurFilter) Xquality_goset(int){}
func (x Xflash_filters_BlurFilter) XblurY_goget()  float64 { return 0; }
func (x Xflash_filters_BlurFilter) XblurY_goset(float64){}
func (x Xflash_filters_BlurFilter) XblurX_goget()  float64 { return 0; }
func (x Xflash_filters_BlurFilter) XblurX_goset(float64){}
func Xflash_filters_BlurFilter_new(XblurX float64,XblurY float64,Xquality int) Xflash_filters_BlurFilter { return Xflash_filters_BlurFilter(0); }
type Xflash_filters_ColorMatrixFilter uintptr // TClassdecl
func (x Xflash_filters_ColorMatrixFilter) Xmatrix_goget()  XArray { return 0; }
func (x Xflash_filters_ColorMatrixFilter) Xmatrix_goset(XArray){}
func Xflash_filters_ColorMatrixFilter_new(Xmatrix XArray) Xflash_filters_ColorMatrixFilter { return Xflash_filters_ColorMatrixFilter(0); }
type Xflash_filters_DropShadowFilter uintptr // TClassdecl
func (x Xflash_filters_DropShadowFilter) Xstrength_goget()  float64 { return 0; }
func (x Xflash_filters_DropShadowFilter) Xstrength_goset(float64){}
func (x Xflash_filters_DropShadowFilter) Xquality_goget()  int { return 0; }
func (x Xflash_filters_DropShadowFilter) Xquality_goset(int){}
func (x Xflash_filters_DropShadowFilter) Xknockout_goget()  bool { return false; }
func (x Xflash_filters_DropShadowFilter) Xknockout_goset(bool){}
func (x Xflash_filters_DropShadowFilter) Xinner_goget()  bool { return false; }
func (x Xflash_filters_DropShadowFilter) Xinner_goset(bool){}
func (x Xflash_filters_DropShadowFilter) XhideObject_goget()  bool { return false; }
func (x Xflash_filters_DropShadowFilter) XhideObject_goset(bool){}
func (x Xflash_filters_DropShadowFilter) Xdistance_goget()  float64 { return 0; }
func (x Xflash_filters_DropShadowFilter) Xdistance_goset(float64){}
func (x Xflash_filters_DropShadowFilter) Xcolor_goget()  int { return 0; }
func (x Xflash_filters_DropShadowFilter) Xcolor_goset(int){}
func (x Xflash_filters_DropShadowFilter) XblurY_goget()  float64 { return 0; }
func (x Xflash_filters_DropShadowFilter) XblurY_goset(float64){}
func (x Xflash_filters_DropShadowFilter) XblurX_goget()  float64 { return 0; }
func (x Xflash_filters_DropShadowFilter) XblurX_goset(float64){}
func (x Xflash_filters_DropShadowFilter) Xangle_goget()  float64 { return 0; }
func (x Xflash_filters_DropShadowFilter) Xangle_goset(float64){}
func (x Xflash_filters_DropShadowFilter) Xalpha_goget()  float64 { return 0; }
func (x Xflash_filters_DropShadowFilter) Xalpha_goset(float64){}
func Xflash_filters_DropShadowFilter_new(Xdistance float64,Xangle float64,Xcolor int,Xalpha float64,XblurX float64,XblurY float64,Xstrength float64,Xquality int,Xinner bool,Xknockout bool,XhideObject bool) Xflash_filters_DropShadowFilter { return Xflash_filters_DropShadowFilter(0); }
type Xflash_filters_GlowFilter uintptr // TClassdecl
func (x Xflash_filters_GlowFilter) Xstrength_goget()  float64 { return 0; }
func (x Xflash_filters_GlowFilter) Xstrength_goset(float64){}
func (x Xflash_filters_GlowFilter) Xquality_goget()  int { return 0; }
func (x Xflash_filters_GlowFilter) Xquality_goset(int){}
func (x Xflash_filters_GlowFilter) Xknockout_goget()  bool { return false; }
func (x Xflash_filters_GlowFilter) Xknockout_goset(bool){}
func (x Xflash_filters_GlowFilter) Xinner_goget()  bool { return false; }
func (x Xflash_filters_GlowFilter) Xinner_goset(bool){}
func (x Xflash_filters_GlowFilter) Xcolor_goget()  int { return 0; }
func (x Xflash_filters_GlowFilter) Xcolor_goset(int){}
func (x Xflash_filters_GlowFilter) XblurY_goget()  float64 { return 0; }
func (x Xflash_filters_GlowFilter) XblurY_goset(float64){}
func (x Xflash_filters_GlowFilter) XblurX_goget()  float64 { return 0; }
func (x Xflash_filters_GlowFilter) XblurX_goset(float64){}
func (x Xflash_filters_GlowFilter) Xalpha_goget()  float64 { return 0; }
func (x Xflash_filters_GlowFilter) Xalpha_goset(float64){}
func Xflash_filters_GlowFilter_new(Xcolor int,Xalpha float64,XblurX float64,XblurY float64,Xstrength float64,Xquality int,Xinner bool,Xknockout bool) Xflash_filters_GlowFilter { return Xflash_filters_GlowFilter(0); }
type Xflash_geom_ColorTransform uintptr // TClassdecl
func (x Xflash_geom_ColorTransform) XtoString() string { return ""; }
func (x Xflash_geom_ColorTransform) Xconcat(Xsecond Xflash_geom_ColorTransform) {}
func (x Xflash_geom_ColorTransform) XredOffset_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XredOffset_goset(float64){}
func (x Xflash_geom_ColorTransform) XredMultiplier_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XredMultiplier_goset(float64){}
func (x Xflash_geom_ColorTransform) XgreenOffset_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XgreenOffset_goset(float64){}
func (x Xflash_geom_ColorTransform) XgreenMultiplier_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XgreenMultiplier_goset(float64){}
func (x Xflash_geom_ColorTransform) Xcolor_goget()  int { return 0; }
func (x Xflash_geom_ColorTransform) Xcolor_goset(int){}
func (x Xflash_geom_ColorTransform) XblueOffset_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XblueOffset_goset(float64){}
func (x Xflash_geom_ColorTransform) XblueMultiplier_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XblueMultiplier_goset(float64){}
func (x Xflash_geom_ColorTransform) XalphaOffset_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XalphaOffset_goset(float64){}
func (x Xflash_geom_ColorTransform) XalphaMultiplier_goget()  float64 { return 0; }
func (x Xflash_geom_ColorTransform) XalphaMultiplier_goset(float64){}
func Xflash_geom_ColorTransform_new(XredMultiplier float64,XgreenMultiplier float64,XblueMultiplier float64,XalphaMultiplier float64,XredOffset float64,XgreenOffset float64,XblueOffset float64,XalphaOffset float64) Xflash_geom_ColorTransform { return Xflash_geom_ColorTransform(0); }
type Xflash_geom_Matrix uintptr // TClassdecl
func (x Xflash_geom_Matrix) Xtranslate(Xdx float64,Xdy float64) {}
func (x Xflash_geom_Matrix) XtransformPoint(Xpoint Xflash_geom_Point) Xflash_geom_Point { return 0; }
func (x Xflash_geom_Matrix) XtoString() string { return ""; }
func (x Xflash_geom_Matrix) Xscale(Xsx float64,Xsy float64) {}
func (x Xflash_geom_Matrix) Xrotate(Xangle float64) {}
func (x Xflash_geom_Matrix) Xinvert() {}
func (x Xflash_geom_Matrix) Xidentity() {}
func (x Xflash_geom_Matrix) XdeltaTransformPoint(Xpoint Xflash_geom_Point) Xflash_geom_Point { return 0; }
func (x Xflash_geom_Matrix) XcreateGradientBox(Xwidth float64,Xheight float64,Xrotation float64,Xtx float64,Xty float64) {}
func (x Xflash_geom_Matrix) XcreateBox(XscaleX float64,XscaleY float64,Xrotation float64,Xtx float64,Xty float64) {}
func (x Xflash_geom_Matrix) Xconcat(Xm Xflash_geom_Matrix) {}
func (x Xflash_geom_Matrix) Xclone() Xflash_geom_Matrix { return 0; }
func (x Xflash_geom_Matrix) Xty_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix) Xty_goset(float64){}
func (x Xflash_geom_Matrix) Xtx_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix) Xtx_goset(float64){}
func (x Xflash_geom_Matrix) Xd_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix) Xd_goset(float64){}
func (x Xflash_geom_Matrix) Xc_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix) Xc_goset(float64){}
func (x Xflash_geom_Matrix) Xb_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix) Xb_goset(float64){}
func (x Xflash_geom_Matrix) Xa_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix) Xa_goset(float64){}
func Xflash_geom_Matrix_new(Xa float64,Xb float64,Xc float64,Xd float64,Xtx float64,Xty float64) Xflash_geom_Matrix { return Xflash_geom_Matrix(0); }
type Xflash_geom_Matrix3D uintptr // TClassdecl
func (x Xflash_geom_Matrix3D) Xtranspose() {}
func (x Xflash_geom_Matrix3D) XtransformVectors(Xvin Xflash_Vector,Xvout Xflash_Vector) {}
func (x Xflash_geom_Matrix3D) XtransformVector(Xv Xflash_geom_Vector3D) Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Matrix3D) XprependTranslation(Xx float64,Xy float64,Xz float64) {}
func (x Xflash_geom_Matrix3D) XprependScale(XxScale float64,XyScale float64,XzScale float64) {}
func (x Xflash_geom_Matrix3D) XprependRotation(Xdegrees float64,Xaxis Xflash_geom_Vector3D,XpivotPoint Xflash_geom_Vector3D) {}
func (x Xflash_geom_Matrix3D) Xprepend(Xrhs Xflash_geom_Matrix3D) {}
func (x Xflash_geom_Matrix3D) XpointAt(Xpos Xflash_geom_Vector3D,Xat Xflash_geom_Vector3D,Xup Xflash_geom_Vector3D) {}
func (x Xflash_geom_Matrix3D) Xinvert() bool { return false; }
func (x Xflash_geom_Matrix3D) XinterpolateTo(XtoMat Xflash_geom_Matrix3D,Xpercent float64) {}
func (x Xflash_geom_Matrix3D) Xidentity() {}
func (x Xflash_geom_Matrix3D) XdeltaTransformVector(Xv Xflash_geom_Vector3D) Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Matrix3D) Xclone() Xflash_geom_Matrix3D { return 0; }
func (x Xflash_geom_Matrix3D) XappendTranslation(Xx float64,Xy float64,Xz float64) {}
func (x Xflash_geom_Matrix3D) XappendScale(XxScale float64,XyScale float64,XzScale float64) {}
func (x Xflash_geom_Matrix3D) XappendRotation(Xdegrees float64,Xaxis Xflash_geom_Vector3D,XpivotPoint Xflash_geom_Vector3D) {}
func (x Xflash_geom_Matrix3D) Xappend(Xlhs Xflash_geom_Matrix3D) {}
func (x Xflash_geom_Matrix3D) XrawData_goget()  Xflash_Vector { return 0; }
func (x Xflash_geom_Matrix3D) XrawData_goset(Xflash_Vector){}
func (x Xflash_geom_Matrix3D) Xposition_goget()  Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Matrix3D) Xposition_goset(Xflash_geom_Vector3D){}
func (x Xflash_geom_Matrix3D) Xdeterminant_goget()  float64 { return 0; }
func (x Xflash_geom_Matrix3D) Xdeterminant_goset(float64){}
func Xflash_geom_Matrix3D_new(Xv Xflash_Vector) Xflash_geom_Matrix3D { return Xflash_geom_Matrix3D(0); }
func Xflash_geom_Point_distance(Xpt1 Xflash_geom_Point,Xpt2 Xflash_geom_Point) float64 { return 0; }
func Xflash_geom_Point_interpolate(Xpt1 Xflash_geom_Point,Xpt2 Xflash_geom_Point,Xf float64) Xflash_geom_Point { return 0; }
func Xflash_geom_Point_polar(Xlen float64,Xangle float64) Xflash_geom_Point { return 0; }
type Xflash_geom_Point uintptr // TClassdecl
func (x Xflash_geom_Point) XtoString() string { return ""; }
func (x Xflash_geom_Point) Xsubtract(Xv Xflash_geom_Point) Xflash_geom_Point { return 0; }
func (x Xflash_geom_Point) Xoffset(Xdx float64,Xdy float64) {}
func (x Xflash_geom_Point) Xnormalize(Xthickness float64) {}
func (x Xflash_geom_Point) Xequals(XtoCompare Xflash_geom_Point) bool { return false; }
func (x Xflash_geom_Point) Xclone() Xflash_geom_Point { return 0; }
func (x Xflash_geom_Point) Xadd(Xv Xflash_geom_Point) Xflash_geom_Point { return 0; }
func (x Xflash_geom_Point) Xy_goget()  float64 { return 0; }
func (x Xflash_geom_Point) Xy_goset(float64){}
func (x Xflash_geom_Point) Xx_goget()  float64 { return 0; }
func (x Xflash_geom_Point) Xx_goset(float64){}
func (x Xflash_geom_Point) Xlength_goget()  float64 { return 0; }
func (x Xflash_geom_Point) Xlength_goset(float64){}
func Xflash_geom_Point_new(Xx float64,Xy float64) Xflash_geom_Point { return Xflash_geom_Point(0); }
type Xflash_geom_Rectangle uintptr // TClassdecl
func (x Xflash_geom_Rectangle) Xunion(XtoUnion Xflash_geom_Rectangle) Xflash_geom_Rectangle { return 0; }
func (x Xflash_geom_Rectangle) XtoString() string { return ""; }
func (x Xflash_geom_Rectangle) XsetEmpty() {}
func (x Xflash_geom_Rectangle) XoffsetPoint(Xpoint Xflash_geom_Point) {}
func (x Xflash_geom_Rectangle) Xoffset(Xdx float64,Xdy float64) {}
func (x Xflash_geom_Rectangle) XisEmpty() bool { return false; }
func (x Xflash_geom_Rectangle) Xintersects(XtoIntersect Xflash_geom_Rectangle) bool { return false; }
func (x Xflash_geom_Rectangle) Xintersection(XtoIntersect Xflash_geom_Rectangle) Xflash_geom_Rectangle { return 0; }
func (x Xflash_geom_Rectangle) XinflatePoint(Xpoint Xflash_geom_Point) {}
func (x Xflash_geom_Rectangle) Xinflate(Xdx float64,Xdy float64) {}
func (x Xflash_geom_Rectangle) Xequals(XtoCompare Xflash_geom_Rectangle) bool { return false; }
func (x Xflash_geom_Rectangle) XcontainsRect(Xrect Xflash_geom_Rectangle) bool { return false; }
func (x Xflash_geom_Rectangle) XcontainsPoint(Xpoint Xflash_geom_Point) bool { return false; }
func (x Xflash_geom_Rectangle) Xcontains(Xx float64,Xy float64) bool { return false; }
func (x Xflash_geom_Rectangle) Xclone() Xflash_geom_Rectangle { return 0; }
func (x Xflash_geom_Rectangle) Xy_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xy_goset(float64){}
func (x Xflash_geom_Rectangle) Xx_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xx_goset(float64){}
func (x Xflash_geom_Rectangle) Xwidth_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xwidth_goset(float64){}
func (x Xflash_geom_Rectangle) XtopLeft_goget()  Xflash_geom_Point { return 0; }
func (x Xflash_geom_Rectangle) XtopLeft_goset(Xflash_geom_Point){}
func (x Xflash_geom_Rectangle) Xtop_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xtop_goset(float64){}
func (x Xflash_geom_Rectangle) Xsize_goget()  Xflash_geom_Point { return 0; }
func (x Xflash_geom_Rectangle) Xsize_goset(Xflash_geom_Point){}
func (x Xflash_geom_Rectangle) Xright_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xright_goset(float64){}
func (x Xflash_geom_Rectangle) Xleft_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xleft_goset(float64){}
func (x Xflash_geom_Rectangle) Xheight_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xheight_goset(float64){}
func (x Xflash_geom_Rectangle) XbottomRight_goget()  Xflash_geom_Point { return 0; }
func (x Xflash_geom_Rectangle) XbottomRight_goset(Xflash_geom_Point){}
func (x Xflash_geom_Rectangle) Xbottom_goget()  float64 { return 0; }
func (x Xflash_geom_Rectangle) Xbottom_goset(float64){}
func Xflash_geom_Rectangle_new(Xx float64,Xy float64,Xwidth float64,Xheight float64) Xflash_geom_Rectangle { return Xflash_geom_Rectangle(0); }
type Xflash_geom_Transform uintptr // TClassdecl
func (x Xflash_geom_Transform) XpixelBounds_goget()  Xflash_geom_Rectangle { return 0; }
func (x Xflash_geom_Transform) XpixelBounds_goset(Xflash_geom_Rectangle){}
func (x Xflash_geom_Transform) Xmatrix_goget()  Xflash_geom_Matrix { return 0; }
func (x Xflash_geom_Transform) Xmatrix_goset(Xflash_geom_Matrix){}
func (x Xflash_geom_Transform) XconcatenatedMatrix_goget()  Xflash_geom_Matrix { return 0; }
func (x Xflash_geom_Transform) XconcatenatedMatrix_goset(Xflash_geom_Matrix){}
func (x Xflash_geom_Transform) XconcatenatedColorTransform_goget()  Xflash_geom_ColorTransform { return 0; }
func (x Xflash_geom_Transform) XconcatenatedColorTransform_goset(Xflash_geom_ColorTransform){}
func (x Xflash_geom_Transform) XcolorTransform_goget()  Xflash_geom_ColorTransform { return 0; }
func (x Xflash_geom_Transform) XcolorTransform_goset(Xflash_geom_ColorTransform){}
func Xflash_geom_Transform_new(XdisplayObject Xflash_display_DisplayObject) Xflash_geom_Transform { return Xflash_geom_Transform(0); }
type Xflash_geom_Vector3D uintptr // TClassdecl
func (x Xflash_geom_Vector3D) XtoString() string { return ""; }
func (x Xflash_geom_Vector3D) Xsubtract(Xa Xflash_geom_Vector3D) Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Vector3D) XscaleBy(Xs float64) {}
func (x Xflash_geom_Vector3D) Xproject() {}
func (x Xflash_geom_Vector3D) Xnormalize() float64 { return 0; }
func (x Xflash_geom_Vector3D) Xnegate() {}
func (x Xflash_geom_Vector3D) XnearEquals(XtoCompare Xflash_geom_Vector3D,Xtolerance float64,XallFour bool) bool { return false; }
func (x Xflash_geom_Vector3D) XincrementBy(Xa Xflash_geom_Vector3D) {}
func (x Xflash_geom_Vector3D) Xequals(XtoCompare Xflash_geom_Vector3D,XallFour bool) bool { return false; }
func (x Xflash_geom_Vector3D) XdotProduct(Xa Xflash_geom_Vector3D) float64 { return 0; }
func (x Xflash_geom_Vector3D) XdecrementBy(Xa Xflash_geom_Vector3D) {}
func (x Xflash_geom_Vector3D) XcrossProduct(Xa Xflash_geom_Vector3D) Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Vector3D) Xclone() Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Vector3D) Xadd(Xa Xflash_geom_Vector3D) Xflash_geom_Vector3D { return 0; }
func (x Xflash_geom_Vector3D) Xz_goget()  float64 { return 0; }
func (x Xflash_geom_Vector3D) Xz_goset(float64){}
func (x Xflash_geom_Vector3D) Xy_goget()  float64 { return 0; }
func (x Xflash_geom_Vector3D) Xy_goset(float64){}
func (x Xflash_geom_Vector3D) Xx_goget()  float64 { return 0; }
func (x Xflash_geom_Vector3D) Xx_goset(float64){}
func (x Xflash_geom_Vector3D) Xw_goget()  float64 { return 0; }
func (x Xflash_geom_Vector3D) Xw_goset(float64){}
func (x Xflash_geom_Vector3D) XlengthSquared_goget()  float64 { return 0; }
func (x Xflash_geom_Vector3D) XlengthSquared_goset(float64){}
func (x Xflash_geom_Vector3D) Xlength_goget()  float64 { return 0; }
func (x Xflash_geom_Vector3D) Xlength_goset(float64){}
func Xflash_geom_Vector3D_new(Xx float64,Xy float64,Xz float64,Xw float64) Xflash_geom_Vector3D { return Xflash_geom_Vector3D(0); }
type Xflash_media_ID3Info uintptr // TClassdecl
func (x Xflash_media_ID3Info) Xyear_goget()  string { return ""; }
func (x Xflash_media_ID3Info) Xyear_goset(string){}
func (x Xflash_media_ID3Info) Xtrack_goget()  string { return ""; }
func (x Xflash_media_ID3Info) Xtrack_goset(string){}
func (x Xflash_media_ID3Info) XsongName_goget()  string { return ""; }
func (x Xflash_media_ID3Info) XsongName_goset(string){}
func (x Xflash_media_ID3Info) Xgenre_goget()  string { return ""; }
func (x Xflash_media_ID3Info) Xgenre_goset(string){}
func (x Xflash_media_ID3Info) Xcomment_goget()  string { return ""; }
func (x Xflash_media_ID3Info) Xcomment_goset(string){}
func (x Xflash_media_ID3Info) Xartist_goget()  string { return ""; }
func (x Xflash_media_ID3Info) Xartist_goset(string){}
func (x Xflash_media_ID3Info) Xalbum_goget()  string { return ""; }
func (x Xflash_media_ID3Info) Xalbum_goset(string){}
func Xflash_media_ID3Info_new() Xflash_media_ID3Info { return Xflash_media_ID3Info(0); }
type Xflash_media_Sound uintptr // TClassdecl
func (x Xflash_media_Sound) Xplay(XstartTime float64,Xloops int,XsndTransform Xflash_media_SoundTransform) Xflash_media_SoundChannel { return 0; }
func (x Xflash_media_Sound) Xload(Xstream Xflash_net_URLRequest,Xcontext Xflash_media_SoundLoaderContext) {}
func (x Xflash_media_Sound) Xclose() {}
func (x Xflash_media_Sound) Xurl_goget()  string { return ""; }
func (x Xflash_media_Sound) Xurl_goset(string){}
func (x Xflash_media_Sound) Xlength_goget()  float64 { return 0; }
func (x Xflash_media_Sound) Xlength_goset(float64){}
func (x Xflash_media_Sound) XisBuffering_goget()  bool { return false; }
func (x Xflash_media_Sound) XisBuffering_goset(bool){}
func (x Xflash_media_Sound) Xid3_goget()  Xflash_media_ID3Info { return 0; }
func (x Xflash_media_Sound) Xid3_goset(Xflash_media_ID3Info){}
func (x Xflash_media_Sound) XbytesTotal_goget()  int { return 0; }
func (x Xflash_media_Sound) XbytesTotal_goset(int){}
func (x Xflash_media_Sound) XbytesLoaded_goget()  int { return 0; }
func (x Xflash_media_Sound) XbytesLoaded_goset(int){}
func Xflash_media_Sound_new(Xstream Xflash_net_URLRequest,Xcontext Xflash_media_SoundLoaderContext) Xflash_media_Sound { return Xflash_media_Sound(0); }
type Xflash_media_SoundChannel uintptr // TClassdecl
func (x Xflash_media_SoundChannel) Xstop() {}
func (x Xflash_media_SoundChannel) XsoundTransform_goget()  Xflash_media_SoundTransform { return 0; }
func (x Xflash_media_SoundChannel) XsoundTransform_goset(Xflash_media_SoundTransform){}
func (x Xflash_media_SoundChannel) XrightPeak_goget()  float64 { return 0; }
func (x Xflash_media_SoundChannel) XrightPeak_goset(float64){}
func (x Xflash_media_SoundChannel) Xposition_goget()  float64 { return 0; }
func (x Xflash_media_SoundChannel) Xposition_goset(float64){}
func (x Xflash_media_SoundChannel) XleftPeak_goget()  float64 { return 0; }
func (x Xflash_media_SoundChannel) XleftPeak_goset(float64){}
func Xflash_media_SoundChannel_new() Xflash_media_SoundChannel { return Xflash_media_SoundChannel(0); }
type Xflash_media_SoundLoaderContext uintptr // TClassdecl
func (x Xflash_media_SoundLoaderContext) XcheckPolicyFile_goget()  bool { return false; }
func (x Xflash_media_SoundLoaderContext) XcheckPolicyFile_goset(bool){}
func (x Xflash_media_SoundLoaderContext) XbufferTime_goget()  float64 { return 0; }
func (x Xflash_media_SoundLoaderContext) XbufferTime_goset(float64){}
func Xflash_media_SoundLoaderContext_new(XbufferTime float64,XcheckPolicyFile bool) Xflash_media_SoundLoaderContext { return Xflash_media_SoundLoaderContext(0); }
type Xflash_media_SoundTransform uintptr // TClassdecl
func (x Xflash_media_SoundTransform) Xvolume_goget()  float64 { return 0; }
func (x Xflash_media_SoundTransform) Xvolume_goset(float64){}
func (x Xflash_media_SoundTransform) XrightToRight_goget()  float64 { return 0; }
func (x Xflash_media_SoundTransform) XrightToRight_goset(float64){}
func (x Xflash_media_SoundTransform) XrightToLeft_goget()  float64 { return 0; }
func (x Xflash_media_SoundTransform) XrightToLeft_goset(float64){}
func (x Xflash_media_SoundTransform) Xpan_goget()  float64 { return 0; }
func (x Xflash_media_SoundTransform) Xpan_goset(float64){}
func (x Xflash_media_SoundTransform) XleftToRight_goget()  float64 { return 0; }
func (x Xflash_media_SoundTransform) XleftToRight_goset(float64){}
func (x Xflash_media_SoundTransform) XleftToLeft_goget()  float64 { return 0; }
func (x Xflash_media_SoundTransform) XleftToLeft_goset(float64){}
func Xflash_media_SoundTransform_new(Xvol float64,Xpanning float64) Xflash_media_SoundTransform { return Xflash_media_SoundTransform(0); }
func Xflash_net_SharedObject_defaultObjectEncoding_goget()  int { return 0; }
func Xflash_net_SharedObject_defaultObjectEncoding_goset(int){}
func Xflash_net_SharedObject_deleteAll(Xurl string) int { return 0; }
func Xflash_net_SharedObject_getDiskUsage(Xurl string) int { return 0; }
func Xflash_net_SharedObject_getLocal(Xname string,XlocalPath string,Xsecure bool) Xflash_net_SharedObject { return 0; }
func Xflash_net_SharedObject_getRemote(Xname string,XremotePath string,Xpersistence interface{},Xsecure bool) Xflash_net_SharedObject { return 0; }
type Xflash_net_SharedObject uintptr // TClassdecl
func (x Xflash_net_SharedObject) XsetProperty(XpropertyName string,Xvalue uintptr) {}
func (x Xflash_net_SharedObject) XsetDirty(XpropertyName string) {}
func (x Xflash_net_SharedObject) Xflush(XminDiskSpace int) string { return ""; }
func (x Xflash_net_SharedObject) Xclose() {}
func (x Xflash_net_SharedObject) Xclear() {}
func (x Xflash_net_SharedObject) Xsize_goget()  int { return 0; }
func (x Xflash_net_SharedObject) Xsize_goset(int){}
func (x Xflash_net_SharedObject) XobjectEncoding_goget()  int { return 0; }
func (x Xflash_net_SharedObject) XobjectEncoding_goset(int){}
func (x Xflash_net_SharedObject) Xfps_goget()  float64 { return 0; }
func (x Xflash_net_SharedObject) Xfps_goset(float64){}
func (x Xflash_net_SharedObject) Xdata_goget()  interface{} { return nil; }
func (x Xflash_net_SharedObject) Xdata_goset(interface{}){}
func (x Xflash_net_SharedObject) Xclient_goget()  interface{} { return nil; }
func (x Xflash_net_SharedObject) Xclient_goset(interface{}){}
func Xflash_net_SharedObject_new() Xflash_net_SharedObject { return Xflash_net_SharedObject(0); }
func Xflash_net_SharedObjectFlushStatus_FLUSHED_goget()  string { return ""; }
func Xflash_net_SharedObjectFlushStatus_FLUSHED_goset(string){}
func Xflash_net_SharedObjectFlushStatus_PENDING_goget()  string { return ""; }
func Xflash_net_SharedObjectFlushStatus_PENDING_goset(string){}
type Xflash_net_SharedObjectFlushStatus uintptr // TClassdecl
func Xflash_net_SharedObjectFlushStatus_new() Xflash_net_SharedObjectFlushStatus { return Xflash_net_SharedObjectFlushStatus(0); }
type Xflash_net_URLLoader uintptr // TClassdecl
func (x Xflash_net_URLLoader) Xload(Xrequest Xflash_net_URLRequest) {}
func (x Xflash_net_URLLoader) Xclose() {}
func (x Xflash_net_URLLoader) XdataFormat_goget()  Xflash_net_URLLoaderDataFormat { return 0; }
func (x Xflash_net_URLLoader) XdataFormat_goset(Xflash_net_URLLoaderDataFormat){}
func (x Xflash_net_URLLoader) Xdata_goget()  interface{} { return nil; }
func (x Xflash_net_URLLoader) Xdata_goset(interface{}){}
func (x Xflash_net_URLLoader) XbytesTotal_goget()  int { return 0; }
func (x Xflash_net_URLLoader) XbytesTotal_goset(int){}
func (x Xflash_net_URLLoader) XbytesLoaded_goget()  int { return 0; }
func (x Xflash_net_URLLoader) XbytesLoaded_goset(int){}
func Xflash_net_URLLoader_new(Xrequest Xflash_net_URLRequest) Xflash_net_URLLoader { return Xflash_net_URLLoader(0); }
type Xflash_net_URLLoaderDataFormat uintptr // TEnumdecl
type Xflash_net_URLRequest uintptr // TClassdecl
func (x Xflash_net_URLRequest) Xurl_goget()  string { return ""; }
func (x Xflash_net_URLRequest) Xurl_goset(string){}
func (x Xflash_net_URLRequest) XrequestHeaders_goget()  XArray { return 0; }
func (x Xflash_net_URLRequest) XrequestHeaders_goset(XArray){}
func (x Xflash_net_URLRequest) Xmethod_goget()  string { return ""; }
func (x Xflash_net_URLRequest) Xmethod_goset(string){}
func (x Xflash_net_URLRequest) Xdigest_goget()  string { return ""; }
func (x Xflash_net_URLRequest) Xdigest_goset(string){}
func (x Xflash_net_URLRequest) Xdata_goget()  interface{} { return nil; }
func (x Xflash_net_URLRequest) Xdata_goset(interface{}){}
func (x Xflash_net_URLRequest) XcontentType_goget()  string { return ""; }
func (x Xflash_net_URLRequest) XcontentType_goset(string){}
func Xflash_net_URLRequest_new(Xurl string) Xflash_net_URLRequest { return Xflash_net_URLRequest(0); }
type Xflash_net_URLRequestHeader uintptr // TClassdecl
func (x Xflash_net_URLRequestHeader) Xvalue_goget()  string { return ""; }
func (x Xflash_net_URLRequestHeader) Xvalue_goset(string){}
func (x Xflash_net_URLRequestHeader) Xname_goget()  string { return ""; }
func (x Xflash_net_URLRequestHeader) Xname_goset(string){}
func Xflash_net_URLRequestHeader_new(Xname string,Xvalue string) Xflash_net_URLRequestHeader { return Xflash_net_URLRequestHeader(0); }
func Xflash_net_URLRequestMethod_GET_goget()  string { return ""; }
func Xflash_net_URLRequestMethod_GET_goset(string){}
func Xflash_net_URLRequestMethod_POST_goget()  string { return ""; }
func Xflash_net_URLRequestMethod_POST_goset(string){}
type Xflash_net_URLRequestMethod uintptr // TClassdecl
type Xflash_net_URLVariables uintptr // TClassdecl
func (x Xflash_net_URLVariables) XtoString() string { return ""; }
func (x Xflash_net_URLVariables) Xdecode(Xsource string) {}
func Xflash_net_URLVariables_new(Xsource string) Xflash_net_URLVariables { return Xflash_net_URLVariables(0); }
type Xflash_sensors_Accelerometer uintptr // TClassdecl
func (x Xflash_sensors_Accelerometer) XsetRequestedUpdateInterval(Xinterval float64) {}
func (x Xflash_sensors_Accelerometer) Xmuted_goget()  bool { return false; }
func (x Xflash_sensors_Accelerometer) Xmuted_goset(bool){}
func Xflash_sensors_Accelerometer_new() Xflash_sensors_Accelerometer { return Xflash_sensors_Accelerometer(0); }
func Xflash_system_ApplicationDomain_currentDomain_goget()  Xflash_system_ApplicationDomain { return 0; }
func Xflash_system_ApplicationDomain_currentDomain_goset(Xflash_system_ApplicationDomain){}
type Xflash_system_ApplicationDomain uintptr // TClassdecl
func (x Xflash_system_ApplicationDomain) XhasDefinition(Xname string) bool { return false; }
func (x Xflash_system_ApplicationDomain) XgetDefinition(Xname string) uintptr { return 0; }
func (x Xflash_system_ApplicationDomain) XparentDomain_goget()  Xflash_system_ApplicationDomain { return 0; }
func (x Xflash_system_ApplicationDomain) XparentDomain_goset(Xflash_system_ApplicationDomain){}
func Xflash_system_ApplicationDomain_new(XparentDomain Xflash_system_ApplicationDomain) Xflash_system_ApplicationDomain { return Xflash_system_ApplicationDomain(0); }
func Xflash_system_Capabilities_language_goget()  string { return ""; }
func Xflash_system_Capabilities_language_goset(string){}
func Xflash_system_Capabilities_screenDPI_goget()  float64 { return 0; }
func Xflash_system_Capabilities_screenDPI_goset(float64){}
func Xflash_system_Capabilities_screenResolutionX_goget()  float64 { return 0; }
func Xflash_system_Capabilities_screenResolutionX_goset(float64){}
func Xflash_system_Capabilities_screenResolutionY_goget()  float64 { return 0; }
func Xflash_system_Capabilities_screenResolutionY_goset(float64){}
type Xflash_system_Capabilities uintptr // TClassdecl
type Xflash_system_ImageDecodingPolicy uintptr // TEnumdecl
type Xflash_system_LoaderContext uintptr // TClassdecl
func (x Xflash_system_LoaderContext) XsecurityDomain_goget()  Xflash_system_SecurityDomain { return 0; }
func (x Xflash_system_LoaderContext) XsecurityDomain_goset(Xflash_system_SecurityDomain){}
func (x Xflash_system_LoaderContext) XcheckPolicyFile_goget()  bool { return false; }
func (x Xflash_system_LoaderContext) XcheckPolicyFile_goset(bool){}
func (x Xflash_system_LoaderContext) XapplicationDomain_goget()  Xflash_system_ApplicationDomain { return 0; }
func (x Xflash_system_LoaderContext) XapplicationDomain_goset(Xflash_system_ApplicationDomain){}
func Xflash_system_LoaderContext_new(XcheckPolicyFile bool,XapplicationDomain Xflash_system_ApplicationDomain,XsecurityDomain Xflash_system_SecurityDomain) Xflash_system_LoaderContext { return Xflash_system_LoaderContext(0); }
func Xflash_system_SecurityDomain_currentDomain_goget()  Xflash_system_SecurityDomain { return 0; }
func Xflash_system_SecurityDomain_currentDomain_goset(Xflash_system_SecurityDomain){}
type Xflash_system_SecurityDomain uintptr // TClassdecl
func Xflash_system_System_deviceID_goget()  string { return ""; }
func Xflash_system_System_deviceID_goset(string){}
func Xflash_system_System_totalMemory_goget()  int { return 0; }
func Xflash_system_System_totalMemory_goset(int){}
func Xflash_system_System_exit(Xcode int) {}
func Xflash_system_System_gc() {}
type Xflash_system_System uintptr // TClassdecl
type Xflash_text_AntiAliasType uintptr // TEnumdecl
func Xflash_text_Font_enumerateFonts(XenumerateDeviceFonts bool) XArray { return 0; }
func Xflash_text_Font_registerFont(Xfont XClass) {}
type Xflash_text_Font uintptr // TClassdecl
func (x Xflash_text_Font) XhasGlyphs(Xstr string) bool { return false; }
func (x Xflash_text_Font) XfontType_goget()  Xflash_text_FontType { return 0; }
func (x Xflash_text_Font) XfontType_goset(Xflash_text_FontType){}
func (x Xflash_text_Font) XfontStyle_goget()  Xflash_text_FontStyle { return 0; }
func (x Xflash_text_Font) XfontStyle_goset(Xflash_text_FontStyle){}
func (x Xflash_text_Font) XfontName_goget()  string { return ""; }
func (x Xflash_text_Font) XfontName_goset(string){}
func Xflash_text_Font_new() Xflash_text_Font { return Xflash_text_Font(0); }
type Xflash_text_FontStyle uintptr // TEnumdecl
type Xflash_text_FontType uintptr // TEnumdecl
type Xflash_text_GridFitType uintptr // TEnumdecl
type Xflash_text_TextField uintptr // TClassdecl
func (x Xflash_text_TextField) XsetTextFormat(Xformat Xflash_text_TextFormat,XbeginIndex int,XendIndex int) {}
func (x Xflash_text_TextField) XsetSelection(XbeginIndex int,XendIndex int) {}
func (x Xflash_text_TextField) XreplaceText(XbeginIndex int,XendIndex int,XnewText string) {}
func (x Xflash_text_TextField) XreplaceSelectedText(Xvalue string) {}
func (x Xflash_text_TextField) XpasteRichText(XrichText string) bool { return false; }
func (x Xflash_text_TextField) XinsertXMLText(XbeginIndex int,XendIndex int,XrichText string,Xpasting bool) {}
func (x Xflash_text_TextField) XgetXMLText(XbeginIndex int,XendIndex int) string { return ""; }
func (x Xflash_text_TextField) XgetTextRuns(XbeginIndex int,XendIndex int) XArray { return 0; }
func (x Xflash_text_TextField) XgetTextFormat(XbeginIndex int,XendIndex int) Xflash_text_TextFormat { return 0; }
func (x Xflash_text_TextField) XgetRawText() string { return ""; }
func (x Xflash_text_TextField) XgetParagraphLength(XcharIndex int) int { return 0; }
func (x Xflash_text_TextField) XgetLineText(XlineIndex int) string { return ""; }
func (x Xflash_text_TextField) XgetLineOffset(XlineIndex int) int { return 0; }
func (x Xflash_text_TextField) XgetLineMetrics(XlineIndex int) Xflash_text_TextLineMetrics { return 0; }
func (x Xflash_text_TextField) XgetLineLength(XlineIndex int) int { return 0; }
func (x Xflash_text_TextField) XgetLineIndexOfChar(XcharIndex int) int { return 0; }
func (x Xflash_text_TextField) XgetLineIndexAtPoint(Xx float64,Xy float64) int { return 0; }
func (x Xflash_text_TextField) XgetImageReference(Xid string) Xflash_display_DisplayObject { return 0; }
func (x Xflash_text_TextField) XgetFirstCharInParagraph(XcharIndex int) int { return 0; }
func (x Xflash_text_TextField) XgetCharIndexAtPoint(Xx float64,Xy float64) int { return 0; }
func (x Xflash_text_TextField) XgetCharBoundaries(XcharIndex int) Xflash_geom_Rectangle { return 0; }
func (x Xflash_text_TextField) XcopyRichText() string { return ""; }
func (x Xflash_text_TextField) XappendText(XnewText string) {}
func (x Xflash_text_TextField) XwordWrap_goget()  bool { return false; }
func (x Xflash_text_TextField) XwordWrap_goset(bool){}
func (x Xflash_text_TextField) XuseRichTextClipboard_goget()  bool { return false; }
func (x Xflash_text_TextField) XuseRichTextClipboard_goset(bool){}
func (x Xflash_text_TextField) Xtype_goget()  Xflash_text_TextFieldType { return 0; }
func (x Xflash_text_TextField) Xtype_goset(Xflash_text_TextFieldType){}
func (x Xflash_text_TextField) Xthickness_goget()  float64 { return 0; }
func (x Xflash_text_TextField) Xthickness_goset(float64){}
func (x Xflash_text_TextField) XtextWidth_goget()  float64 { return 0; }
func (x Xflash_text_TextField) XtextWidth_goset(float64){}
func (x Xflash_text_TextField) XtextHeight_goget()  float64 { return 0; }
func (x Xflash_text_TextField) XtextHeight_goset(float64){}
func (x Xflash_text_TextField) XtextColor_goget()  int { return 0; }
func (x Xflash_text_TextField) XtextColor_goset(int){}
func (x Xflash_text_TextField) Xtext_goget()  string { return ""; }
func (x Xflash_text_TextField) Xtext_goset(string){}
func (x Xflash_text_TextField) Xsharpness_goget()  float64 { return 0; }
func (x Xflash_text_TextField) Xsharpness_goset(float64){}
func (x Xflash_text_TextField) XselectionEndIndex_goget()  int { return 0; }
func (x Xflash_text_TextField) XselectionEndIndex_goset(int){}
func (x Xflash_text_TextField) XselectionBeginIndex_goget()  int { return 0; }
func (x Xflash_text_TextField) XselectionBeginIndex_goset(int){}
func (x Xflash_text_TextField) XselectedText_goget()  string { return ""; }
func (x Xflash_text_TextField) XselectedText_goset(string){}
func (x Xflash_text_TextField) Xselectable_goget()  bool { return false; }
func (x Xflash_text_TextField) Xselectable_goset(bool){}
func (x Xflash_text_TextField) XscrollV_goget()  int { return 0; }
func (x Xflash_text_TextField) XscrollV_goset(int){}
func (x Xflash_text_TextField) XscrollH_goget()  int { return 0; }
func (x Xflash_text_TextField) XscrollH_goset(int){}
func (x Xflash_text_TextField) Xrestrict_goget()  string { return ""; }
func (x Xflash_text_TextField) Xrestrict_goset(string){}
func (x Xflash_text_TextField) XnumLines_goget()  int { return 0; }
func (x Xflash_text_TextField) XnumLines_goset(int){}
func (x Xflash_text_TextField) Xmultiline_goget()  bool { return false; }
func (x Xflash_text_TextField) Xmultiline_goset(bool){}
func (x Xflash_text_TextField) XmouseWheelEnabled_goget()  bool { return false; }
func (x Xflash_text_TextField) XmouseWheelEnabled_goset(bool){}
func (x Xflash_text_TextField) XmaxScrollV_goget()  int { return 0; }
func (x Xflash_text_TextField) XmaxScrollV_goset(int){}
func (x Xflash_text_TextField) XmaxScrollH_goget()  int { return 0; }
func (x Xflash_text_TextField) XmaxScrollH_goset(int){}
func (x Xflash_text_TextField) XmaxChars_goget()  int { return 0; }
func (x Xflash_text_TextField) XmaxChars_goset(int){}
func (x Xflash_text_TextField) Xlength_goget()  int { return 0; }
func (x Xflash_text_TextField) Xlength_goset(int){}
func (x Xflash_text_TextField) XhtmlText_goget()  string { return ""; }
func (x Xflash_text_TextField) XhtmlText_goset(string){}
func (x Xflash_text_TextField) XgridFitType_goget()  Xflash_text_GridFitType { return 0; }
func (x Xflash_text_TextField) XgridFitType_goset(Xflash_text_GridFitType){}
func (x Xflash_text_TextField) XembedFonts_goget()  bool { return false; }
func (x Xflash_text_TextField) XembedFonts_goset(bool){}
func (x Xflash_text_TextField) XdisplayAsPassword_goget()  bool { return false; }
func (x Xflash_text_TextField) XdisplayAsPassword_goset(bool){}
func (x Xflash_text_TextField) XdefaultTextFormat_goget()  Xflash_text_TextFormat { return 0; }
func (x Xflash_text_TextField) XdefaultTextFormat_goset(Xflash_text_TextFormat){}
func (x Xflash_text_TextField) XcondenseWhite_goget()  bool { return false; }
func (x Xflash_text_TextField) XcondenseWhite_goset(bool){}
func (x Xflash_text_TextField) XcaretIndex_goget()  int { return 0; }
func (x Xflash_text_TextField) XcaretIndex_goset(int){}
func (x Xflash_text_TextField) XbottomScrollV_goget()  int { return 0; }
func (x Xflash_text_TextField) XbottomScrollV_goset(int){}
func (x Xflash_text_TextField) XborderColor_goget()  int { return 0; }
func (x Xflash_text_TextField) XborderColor_goset(int){}
func (x Xflash_text_TextField) Xborder_goget()  bool { return false; }
func (x Xflash_text_TextField) Xborder_goset(bool){}
func (x Xflash_text_TextField) XbackgroundColor_goget()  int { return 0; }
func (x Xflash_text_TextField) XbackgroundColor_goset(int){}
func (x Xflash_text_TextField) Xbackground_goget()  bool { return false; }
func (x Xflash_text_TextField) Xbackground_goset(bool){}
func (x Xflash_text_TextField) XautoSize_goget()  Xflash_text_TextFieldAutoSize { return 0; }
func (x Xflash_text_TextField) XautoSize_goset(Xflash_text_TextFieldAutoSize){}
func (x Xflash_text_TextField) XantiAliasType_goget()  Xflash_text_AntiAliasType { return 0; }
func (x Xflash_text_TextField) XantiAliasType_goset(Xflash_text_AntiAliasType){}
func (x Xflash_text_TextField) XalwaysShowSelection_goget()  bool { return false; }
func (x Xflash_text_TextField) XalwaysShowSelection_goset(bool){}
func Xflash_text_TextField_new() Xflash_text_TextField { return Xflash_text_TextField(0); }
type Xflash_text_TextFieldAutoSize uintptr // TEnumdecl
type Xflash_text_TextFieldType uintptr // TEnumdecl
type Xflash_text_TextFormat uintptr // TClassdecl
func (x Xflash_text_TextFormat) Xurl_goget()  string { return ""; }
func (x Xflash_text_TextFormat) Xurl_goset(string){}
func (x Xflash_text_TextFormat) Xunderline_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xunderline_goset(uintptr){}
func (x Xflash_text_TextFormat) Xtarget_goget()  string { return ""; }
func (x Xflash_text_TextFormat) Xtarget_goset(string){}
func (x Xflash_text_TextFormat) XtabStops_goget()  XArray { return 0; }
func (x Xflash_text_TextFormat) XtabStops_goset(XArray){}
func (x Xflash_text_TextFormat) Xsize_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xsize_goset(uintptr){}
func (x Xflash_text_TextFormat) XrightMargin_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) XrightMargin_goset(uintptr){}
func (x Xflash_text_TextFormat) XletterSpacing_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) XletterSpacing_goset(uintptr){}
func (x Xflash_text_TextFormat) XleftMargin_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) XleftMargin_goset(uintptr){}
func (x Xflash_text_TextFormat) Xleading_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xleading_goset(uintptr){}
func (x Xflash_text_TextFormat) Xkerning_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xkerning_goset(uintptr){}
func (x Xflash_text_TextFormat) Xitalic_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xitalic_goset(uintptr){}
func (x Xflash_text_TextFormat) Xindent_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xindent_goset(uintptr){}
func (x Xflash_text_TextFormat) Xfont_goget()  string { return ""; }
func (x Xflash_text_TextFormat) Xfont_goset(string){}
func (x Xflash_text_TextFormat) Xcolor_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xcolor_goset(uintptr){}
func (x Xflash_text_TextFormat) Xbullet_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xbullet_goset(uintptr){}
func (x Xflash_text_TextFormat) Xbold_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) Xbold_goset(uintptr){}
func (x Xflash_text_TextFormat) XblockIndent_goget()  uintptr { return 0; }
func (x Xflash_text_TextFormat) XblockIndent_goset(uintptr){}
func (x Xflash_text_TextFormat) Xalign_goget()  Xflash_text_TextFormatAlign { return 0; }
func (x Xflash_text_TextFormat) Xalign_goset(Xflash_text_TextFormatAlign){}
func Xflash_text_TextFormat_new(Xfont string,Xsize float64,Xcolor int,Xbold bool,Xitalic bool,Xunderline bool,Xurl string,Xtarget string,Xalign Xflash_text_TextFormatAlign,XleftMargin float64,XrightMargin float64,Xindent float64,Xleading float64) Xflash_text_TextFormat { return Xflash_text_TextFormat(0); }
type Xflash_text_TextFormatAlign uintptr // TEnumdecl
type Xflash_text_TextLineMetrics uintptr // TClassdecl
func (x Xflash_text_TextLineMetrics) Xx_goget()  float64 { return 0; }
func (x Xflash_text_TextLineMetrics) Xx_goset(float64){}
func (x Xflash_text_TextLineMetrics) Xwidth_goget()  float64 { return 0; }
func (x Xflash_text_TextLineMetrics) Xwidth_goset(float64){}
func (x Xflash_text_TextLineMetrics) Xleading_goget()  float64 { return 0; }
func (x Xflash_text_TextLineMetrics) Xleading_goset(float64){}
func (x Xflash_text_TextLineMetrics) Xheight_goget()  float64 { return 0; }
func (x Xflash_text_TextLineMetrics) Xheight_goset(float64){}
func (x Xflash_text_TextLineMetrics) Xdescent_goget()  float64 { return 0; }
func (x Xflash_text_TextLineMetrics) Xdescent_goset(float64){}
func (x Xflash_text_TextLineMetrics) Xascent_goget()  float64 { return 0; }
func (x Xflash_text_TextLineMetrics) Xascent_goset(float64){}
func Xflash_text_TextLineMetrics_new(Xx float64,Xwidth float64,Xheight float64,Xascent float64,Xdescent float64,Xleading float64) Xflash_text_TextLineMetrics { return Xflash_text_TextLineMetrics(0); }
func Xflash_ui_Keyboard_BACKSPACE_goget()  int { return 0; }
func Xflash_ui_Keyboard_BACKSPACE_goset(int){}
func Xflash_ui_Keyboard_CAPS___LOCK_goget()  int { return 0; }
func Xflash_ui_Keyboard_CAPS___LOCK_goset(int){}
func Xflash_ui_Keyboard_CONTROL_goget()  int { return 0; }
func Xflash_ui_Keyboard_CONTROL_goset(int){}
func Xflash_ui_Keyboard_DELETE_goget()  int { return 0; }
func Xflash_ui_Keyboard_DELETE_goset(int){}
func Xflash_ui_Keyboard_DOWN_goget()  int { return 0; }
func Xflash_ui_Keyboard_DOWN_goset(int){}
func Xflash_ui_Keyboard_END_goget()  int { return 0; }
func Xflash_ui_Keyboard_END_goset(int){}
func Xflash_ui_Keyboard_ENTER_goget()  int { return 0; }
func Xflash_ui_Keyboard_ENTER_goset(int){}
func Xflash_ui_Keyboard_ESCAPE_goget()  int { return 0; }
func Xflash_ui_Keyboard_ESCAPE_goset(int){}
func Xflash_ui_Keyboard_F1_goget()  int { return 0; }
func Xflash_ui_Keyboard_F1_goset(int){}
func Xflash_ui_Keyboard_F10_goget()  int { return 0; }
func Xflash_ui_Keyboard_F10_goset(int){}
func Xflash_ui_Keyboard_F11_goget()  int { return 0; }
func Xflash_ui_Keyboard_F11_goset(int){}
func Xflash_ui_Keyboard_F12_goget()  int { return 0; }
func Xflash_ui_Keyboard_F12_goset(int){}
func Xflash_ui_Keyboard_F13_goget()  int { return 0; }
func Xflash_ui_Keyboard_F13_goset(int){}
func Xflash_ui_Keyboard_F14_goget()  int { return 0; }
func Xflash_ui_Keyboard_F14_goset(int){}
func Xflash_ui_Keyboard_F15_goget()  int { return 0; }
func Xflash_ui_Keyboard_F15_goset(int){}
func Xflash_ui_Keyboard_F2_goget()  int { return 0; }
func Xflash_ui_Keyboard_F2_goset(int){}
func Xflash_ui_Keyboard_F3_goget()  int { return 0; }
func Xflash_ui_Keyboard_F3_goset(int){}
func Xflash_ui_Keyboard_F4_goget()  int { return 0; }
func Xflash_ui_Keyboard_F4_goset(int){}
func Xflash_ui_Keyboard_F5_goget()  int { return 0; }
func Xflash_ui_Keyboard_F5_goset(int){}
func Xflash_ui_Keyboard_F6_goget()  int { return 0; }
func Xflash_ui_Keyboard_F6_goset(int){}
func Xflash_ui_Keyboard_F7_goget()  int { return 0; }
func Xflash_ui_Keyboard_F7_goset(int){}
func Xflash_ui_Keyboard_F8_goget()  int { return 0; }
func Xflash_ui_Keyboard_F8_goset(int){}
func Xflash_ui_Keyboard_F9_goget()  int { return 0; }
func Xflash_ui_Keyboard_F9_goset(int){}
func Xflash_ui_Keyboard_HOME_goget()  int { return 0; }
func Xflash_ui_Keyboard_HOME_goset(int){}
func Xflash_ui_Keyboard_INSERT_goget()  int { return 0; }
func Xflash_ui_Keyboard_INSERT_goset(int){}
func Xflash_ui_Keyboard_LEFT_goget()  int { return 0; }
func Xflash_ui_Keyboard_LEFT_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___0_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___0_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___1_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___1_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___2_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___2_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___3_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___3_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___4_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___4_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___5_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___5_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___6_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___6_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___7_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___7_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___8_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___8_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___9_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___9_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___ADD_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___ADD_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___DECIMAL_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___DECIMAL_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___DIVIDE_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___DIVIDE_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___ENTER_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___ENTER_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___MULTIPLY_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___MULTIPLY_goset(int){}
func Xflash_ui_Keyboard_NUMPAD___SUBTRACT_goget()  int { return 0; }
func Xflash_ui_Keyboard_NUMPAD___SUBTRACT_goset(int){}
func Xflash_ui_Keyboard_PAGE___DOWN_goget()  int { return 0; }
func Xflash_ui_Keyboard_PAGE___DOWN_goset(int){}
func Xflash_ui_Keyboard_PAGE___UP_goget()  int { return 0; }
func Xflash_ui_Keyboard_PAGE___UP_goset(int){}
func Xflash_ui_Keyboard_RIGHT_goget()  int { return 0; }
func Xflash_ui_Keyboard_RIGHT_goset(int){}
func Xflash_ui_Keyboard_SHIFT_goget()  int { return 0; }
func Xflash_ui_Keyboard_SHIFT_goset(int){}
func Xflash_ui_Keyboard_SPACE_goget()  int { return 0; }
func Xflash_ui_Keyboard_SPACE_goset(int){}
func Xflash_ui_Keyboard_TAB_goget()  int { return 0; }
func Xflash_ui_Keyboard_TAB_goset(int){}
func Xflash_ui_Keyboard_UP_goget()  int { return 0; }
func Xflash_ui_Keyboard_UP_goset(int){}
type Xflash_ui_Keyboard uintptr // TClassdecl
func Xflash_ui_Mouse_hide() {}
func Xflash_ui_Mouse_show() {}
type Xflash_ui_Mouse uintptr // TClassdecl
type Xflash_ui_Multitouch uintptr // TClassdecl
type Xflash_ui_MultitouchInputMode uintptr // TEnumdecl
type Xflash_utils_IDataInput uintptr // TClassdecl
func (x Xflash_utils_IDataInput) XreadUnsignedShort() int { return 0; }
func (x Xflash_utils_IDataInput) XreadUnsignedInt() int { return 0; }
func (x Xflash_utils_IDataInput) XreadUnsignedByte() int { return 0; }
func (x Xflash_utils_IDataInput) XreadUTFBytes(Xlength int) string { return ""; }
func (x Xflash_utils_IDataInput) XreadUTF() string { return ""; }
func (x Xflash_utils_IDataInput) XreadShort() int { return 0; }
func (x Xflash_utils_IDataInput) XreadObject() interface{} { return nil; }
func (x Xflash_utils_IDataInput) XreadMultiByte(Xlength int,XcharSet string) string { return ""; }
func (x Xflash_utils_IDataInput) XreadInt() int { return 0; }
func (x Xflash_utils_IDataInput) XreadFloat() float64 { return 0; }
func (x Xflash_utils_IDataInput) XreadDouble() float64 { return 0; }
func (x Xflash_utils_IDataInput) XreadBytes(Xbytes Xflash_utils_ByteArray,Xoffset int,Xlength int) {}
func (x Xflash_utils_IDataInput) XreadByte() int { return 0; }
func (x Xflash_utils_IDataInput) XreadBoolean() bool { return false; }
func (x Xflash_utils_IDataInput) XobjectEncoding_goget()  int { return 0; }
func (x Xflash_utils_IDataInput) XobjectEncoding_goset(int){}
func (x Xflash_utils_IDataInput) Xendian_goget()  Xflash_utils_Endian { return 0; }
func (x Xflash_utils_IDataInput) Xendian_goset(Xflash_utils_Endian){}
func (x Xflash_utils_IDataInput) XbytesAvailable_goget()  int { return 0; }
func (x Xflash_utils_IDataInput) XbytesAvailable_goset(int){}
func Xflash_utils_ByteArray_defaultObjectEncoding_goget()  int { return 0; }
func Xflash_utils_ByteArray_defaultObjectEncoding_goset(int){}
type Xflash_utils_ByteArray uintptr // TClassdecl
func (x Xflash_utils_ByteArray) XwriteUnsignedInt(Xvalue int) {}
func (x Xflash_utils_ByteArray) XwriteUTFBytes(Xvalue string) {}
func (x Xflash_utils_ByteArray) XwriteUTF(Xvalue string) {}
func (x Xflash_utils_ByteArray) XwriteShort(Xvalue int) {}
func (x Xflash_utils_ByteArray) XwriteObject(Xobject interface{}) {}
func (x Xflash_utils_ByteArray) XwriteMultiByte(Xvalue string,XcharSet string) {}
func (x Xflash_utils_ByteArray) XwriteInt(Xvalue int) {}
func (x Xflash_utils_ByteArray) XwriteFloat(Xvalue float64) {}
func (x Xflash_utils_ByteArray) XwriteDouble(Xvalue float64) {}
func (x Xflash_utils_ByteArray) XwriteBytes(Xbytes Xflash_utils_ByteArray,Xoffset int,Xlength int) {}
func (x Xflash_utils_ByteArray) XwriteByte(Xvalue int) {}
func (x Xflash_utils_ByteArray) XwriteBoolean(Xvalue bool) {}
func (x Xflash_utils_ByteArray) Xuncompress() {}
func (x Xflash_utils_ByteArray) XtoString() string { return ""; }
func (x Xflash_utils_ByteArray) XreadUnsignedShort() int { return 0; }
func (x Xflash_utils_ByteArray) XreadUnsignedInt() int { return 0; }
func (x Xflash_utils_ByteArray) XreadUnsignedByte() int { return 0; }
func (x Xflash_utils_ByteArray) XreadUTFBytes(Xlength int) string { return ""; }
func (x Xflash_utils_ByteArray) XreadUTF() string { return ""; }
func (x Xflash_utils_ByteArray) XreadShort() int { return 0; }
func (x Xflash_utils_ByteArray) XreadObject() interface{} { return nil; }
func (x Xflash_utils_ByteArray) XreadMultiByte(Xlength int,XcharSet string) string { return ""; }
func (x Xflash_utils_ByteArray) XreadInt() int { return 0; }
func (x Xflash_utils_ByteArray) XreadFloat() float64 { return 0; }
func (x Xflash_utils_ByteArray) XreadDouble() float64 { return 0; }
func (x Xflash_utils_ByteArray) XreadBytes(Xbytes Xflash_utils_ByteArray,Xoffset int,Xlength int) {}
func (x Xflash_utils_ByteArray) XreadByte() int { return 0; }
func (x Xflash_utils_ByteArray) XreadBoolean() bool { return false; }
func (x Xflash_utils_ByteArray) Xcompress() {}
func (x Xflash_utils_ByteArray) Xposition_goget()  int { return 0; }
func (x Xflash_utils_ByteArray) Xposition_goset(int){}
func (x Xflash_utils_ByteArray) XobjectEncoding_goget()  int { return 0; }
func (x Xflash_utils_ByteArray) XobjectEncoding_goset(int){}
func (x Xflash_utils_ByteArray) Xlength_goget()  int { return 0; }
func (x Xflash_utils_ByteArray) Xlength_goset(int){}
func (x Xflash_utils_ByteArray) Xendian_goget()  Xflash_utils_Endian { return 0; }
func (x Xflash_utils_ByteArray) Xendian_goset(Xflash_utils_Endian){}
func (x Xflash_utils_ByteArray) XbytesAvailable_goget()  int { return 0; }
func (x Xflash_utils_ByteArray) XbytesAvailable_goset(int){}
func Xflash_utils_ByteArray_new() Xflash_utils_ByteArray { return Xflash_utils_ByteArray(0); }
type Xflash_utils_CompressionAlgorithm uintptr // TEnumdecl
type Xflash_utils_Endian uintptr // TEnumdecl
type Xflash_utils_Object interface{} // TTypedecl
type Xflash_utils_Timer uintptr // TClassdecl
func (x Xflash_utils_Timer) Xstop() {}
func (x Xflash_utils_Timer) Xstart() {}
func (x Xflash_utils_Timer) Xreset() {}
func (x Xflash_utils_Timer) Xrunning_goget()  bool { return false; }
func (x Xflash_utils_Timer) Xrunning_goset(bool){}
func (x Xflash_utils_Timer) XrepeatCount_goget()  int { return 0; }
func (x Xflash_utils_Timer) XrepeatCount_goset(int){}
func (x Xflash_utils_Timer) Xdelay_goget()  float64 { return 0; }
func (x Xflash_utils_Timer) Xdelay_goset(float64){}
func (x Xflash_utils_Timer) XcurrentCount_goget()  int { return 0; }
func (x Xflash_utils_Timer) XcurrentCount_goset(int){}
func Xflash_utils_Timer_new(Xdelay float64,XrepeatCount int) Xflash_utils_Timer { return Xflash_utils_Timer(0); }
func Xhaxe_EnumTools_getName(Xe XEnum) string { return ""; }
func Xhaxe_EnumTools_createByName(Xe XEnum,Xconstr string,Xparams XArray) uintptr { return 0; }
func Xhaxe_EnumTools_createByIndex(Xe XEnum,Xindex int,Xparams XArray) uintptr { return 0; }
func Xhaxe_EnumTools_createAll(Xe XEnum) XArray { return 0; }
func Xhaxe_EnumTools_getConstructors(Xe XEnum) XArray { return 0; }
type Xhaxe_EnumTools uintptr // TClassdecl
func Xhaxe_EnumValueTools_equals(Xa uintptr,Xb uintptr) bool { return false; }
func Xhaxe_EnumValueTools_getName(Xe XEnumValue) string { return ""; }
func Xhaxe_EnumValueTools_getParameters(Xe XEnumValue) XArray { return 0; }
func Xhaxe_EnumValueTools_getIndex(Xe XEnumValue) int { return 0; }
type Xhaxe_EnumValueTools uintptr // TClassdecl
func Xhaxe_Log_trace(Xv interface{},Xinfos uintptr) {}
type Xhaxe_Log uintptr // TClassdecl
type Xhaxe_PosInfos uintptr // TTypedecl
func Xhaxe_Timer_stamp() float64 { return 0; }
type Xhaxe_Timer uintptr // TClassdecl
type Xhaxe_TypeResolver uintptr // TTypedecl
type Xhaxe_ds_HashMap uintptr // TAbstractdecl
type Xhaxe_ds_StringMap uintptr // TClassdecl
type Xhaxe_ds_StringMap_T uintptr // TClassdecl
func (x Xhaxe_ds_StringMap) Xexists(Xkey string) bool { return false; }
func (x Xhaxe_ds_StringMap) Xget(Xkey string) uintptr { return 0; }
func (x Xhaxe_ds_StringMap) Xset(Xkey string,Xvalue Xhaxe_ds_StringMap_T) {}
func Xhaxe_ds_StringMap_new() Xhaxe_ds_StringMap { return Xhaxe_ds_StringMap(0); }
type Xhaxe_io_Unsigned___char______ uintptr // TClassdecl
type Xhaxe_io_BytesData XArray // TTypedecl
type Xopenfl_AssetCache uintptr // TClassdecl
func (x Xopenfl_AssetCache) Xclear() {}
func (x Xopenfl_AssetCache) XbitmapData_goget()  XMap { return 0; }
func (x Xopenfl_AssetCache) XbitmapData_goset(XMap){}
func Xopenfl_AssetCache_new() Xopenfl_AssetCache { return Xopenfl_AssetCache(0); }
func Xopenfl_Assets_cache_goget()  Xopenfl_AssetCache { return 0; }
func Xopenfl_Assets_cache_goset(Xopenfl_AssetCache){}
func Xopenfl_Assets_libraries_goget()  XMap { return 0; }
func Xopenfl_Assets_libraries_goset(XMap){}
func Xopenfl_Assets_exists(Xid string,Xtype Xopenfl_AssetType) bool { return false; }
func Xopenfl_Assets_getBitmapData(Xid string,XuseCache bool) Xflash_display_BitmapData { return 0; }
func Xopenfl_Assets_getBytes(Xid string) Xflash_utils_ByteArray { return 0; }
func Xopenfl_Assets_getFont(Xid string) Xflash_text_Font { return 0; }
func Xopenfl_Assets_getMovieClip(Xid string) Xflash_display_MovieClip { return 0; }
func Xopenfl_Assets_getSound(Xid string) Xflash_media_Sound { return 0; }
func Xopenfl_Assets_getText(Xid string) string { return ""; }
func Xopenfl_Assets_loadBitmapData(Xid string,Xhandler interface{}) {}
func Xopenfl_Assets_loadBytes(Xid string,Xhandler interface{}) {}
func Xopenfl_Assets_loadFont(Xid string,Xhandler interface{}) {}
func Xopenfl_Assets_loadLibrary(Xname string,Xhandler interface{}) {}
func Xopenfl_Assets_loadMovieClip(Xid string,Xhandler interface{}) {}
func Xopenfl_Assets_loadSound(Xid string,Xhandler interface{}) {}
func Xopenfl_Assets_loadText(Xid string,Xhandler interface{}) {}
func Xopenfl_Assets_registerLibrary(Xname string,Xlibrary Xopenfl_AssetLibrary) {}
func Xopenfl_Assets_unloadLibrary(Xname string) {}
type Xopenfl_Assets uintptr // TClassdecl
type Xopenfl_AssetLibrary uintptr // TClassdecl
func (x Xopenfl_AssetLibrary) XloadSound(Xid string,Xhandler interface{}) {}
func (x Xopenfl_AssetLibrary) XloadMovieClip(Xid string,Xhandler interface{}) {}
func (x Xopenfl_AssetLibrary) XloadFont(Xid string,Xhandler interface{}) {}
func (x Xopenfl_AssetLibrary) XloadBytes(Xid string,Xhandler interface{}) {}
func (x Xopenfl_AssetLibrary) XloadBitmapData(Xid string,Xhandler interface{}) {}
func (x Xopenfl_AssetLibrary) XgetSound(Xid string) Xflash_media_Sound { return 0; }
func (x Xopenfl_AssetLibrary) XgetMovieClip(Xid string) Xflash_display_MovieClip { return 0; }
func (x Xopenfl_AssetLibrary) XgetFont(Xid string) Xflash_text_Font { return 0; }
func (x Xopenfl_AssetLibrary) XgetBytes(Xid string) Xflash_utils_ByteArray { return 0; }
func (x Xopenfl_AssetLibrary) XgetBitmapData(Xid string) Xflash_display_BitmapData { return 0; }
func (x Xopenfl_AssetLibrary) Xexists(Xid string,Xtype Xopenfl_AssetType) bool { return false; }
func Xopenfl_AssetLibrary_new() Xopenfl_AssetLibrary { return Xopenfl_AssetLibrary(0); }
type Xopenfl_AssetData uintptr // TClassdecl
func (x Xopenfl_AssetData) Xtype_goget()  Xopenfl_AssetType { return 0; }
func (x Xopenfl_AssetData) Xtype_goset(Xopenfl_AssetType){}
func (x Xopenfl_AssetData) Xpath_goget()  string { return ""; }
func (x Xopenfl_AssetData) Xpath_goset(string){}
func (x Xopenfl_AssetData) Xid_goget()  string { return ""; }
func (x Xopenfl_AssetData) Xid_goset(string){}
func Xopenfl_AssetData_new() Xopenfl_AssetData { return Xopenfl_AssetData(0); }
type Xopenfl_AssetType uintptr // TEnumdecl
type Xopenfl_display_DirectRenderer uintptr // TClassdecl
func (x Xopenfl_display_DirectRenderer) Xrender(XinRect Xflash_geom_Rectangle) {}
func Xopenfl_display_DirectRenderer_new(XinType string) Xopenfl_display_DirectRenderer { return Xopenfl_display_DirectRenderer(0); }
type Xopenfl_display_FPS uintptr // TClassdecl
func Xopenfl_display_FPS_new(Xx float64,Xy float64,Xcolor int) Xopenfl_display_FPS { return Xopenfl_display_FPS(0); }
func Xopenfl_display_OpenGLView_CONTEXT___LOST_goget()  string { return ""; }
func Xopenfl_display_OpenGLView_CONTEXT___LOST_goset(string){}
func Xopenfl_display_OpenGLView_CONTEXT___RESTORED_goget()  string { return ""; }
func Xopenfl_display_OpenGLView_CONTEXT___RESTORED_goset(string){}
func Xopenfl_display_OpenGLView_isSupported_goget()  bool { return false; }
func Xopenfl_display_OpenGLView_isSupported_goset(bool){}
type Xopenfl_display_OpenGLView uintptr // TClassdecl
func Xopenfl_display_OpenGLView_new() Xopenfl_display_OpenGLView { return Xopenfl_display_OpenGLView(0); }
func Xopenfl_display_Tilesheet_TILE___SCALE_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___SCALE_goset(int){}
func Xopenfl_display_Tilesheet_TILE___ROTATION_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___ROTATION_goset(int){}
func Xopenfl_display_Tilesheet_TILE___RGB_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___RGB_goset(int){}
func Xopenfl_display_Tilesheet_TILE___ALPHA_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___ALPHA_goset(int){}
func Xopenfl_display_Tilesheet_TILE___TRANS___2x2_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___TRANS___2x2_goset(int){}
func Xopenfl_display_Tilesheet_TILE___BLEND___NORMAL_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___BLEND___NORMAL_goset(int){}
func Xopenfl_display_Tilesheet_TILE___BLEND___ADD_goget()  int { return 0; }
func Xopenfl_display_Tilesheet_TILE___BLEND___ADD_goset(int){}
type Xopenfl_display_Tilesheet uintptr // TClassdecl
func (x Xopenfl_display_Tilesheet) XdrawTiles(Xgraphics Xflash_display_Graphics,XtileData XArray,Xsmooth bool,Xflags int) {}
func (x Xopenfl_display_Tilesheet) XaddTileRect(Xrectangle Xflash_geom_Rectangle,XcenterPoint Xflash_geom_Point) int { return 0; }
func Xopenfl_display_Tilesheet_new(XinImage Xflash_display_BitmapData) Xopenfl_display_Tilesheet { return Xopenfl_display_Tilesheet(0); }
func Xopenfl_events_JoystickEvent_AXIS___MOVE_goget()  string { return ""; }
func Xopenfl_events_JoystickEvent_AXIS___MOVE_goset(string){}
func Xopenfl_events_JoystickEvent_BALL___MOVE_goget()  string { return ""; }
func Xopenfl_events_JoystickEvent_BALL___MOVE_goset(string){}
func Xopenfl_events_JoystickEvent_BUTTON___DOWN_goget()  string { return ""; }
func Xopenfl_events_JoystickEvent_BUTTON___DOWN_goset(string){}
func Xopenfl_events_JoystickEvent_BUTTON___UP_goget()  string { return ""; }
func Xopenfl_events_JoystickEvent_BUTTON___UP_goset(string){}
func Xopenfl_events_JoystickEvent_HAT___MOVE_goget()  string { return ""; }
func Xopenfl_events_JoystickEvent_HAT___MOVE_goset(string){}
type Xopenfl_events_JoystickEvent uintptr // TClassdecl
func (x Xopenfl_events_JoystickEvent) Xz_goget()  float64 { return 0; }
func (x Xopenfl_events_JoystickEvent) Xz_goset(float64){}
func (x Xopenfl_events_JoystickEvent) Xy_goget()  float64 { return 0; }
func (x Xopenfl_events_JoystickEvent) Xy_goset(float64){}
func (x Xopenfl_events_JoystickEvent) Xx_goget()  float64 { return 0; }
func (x Xopenfl_events_JoystickEvent) Xx_goset(float64){}
func (x Xopenfl_events_JoystickEvent) Xid_goget()  int { return 0; }
func (x Xopenfl_events_JoystickEvent) Xid_goset(int){}
func (x Xopenfl_events_JoystickEvent) Xdevice_goget()  int { return 0; }
func (x Xopenfl_events_JoystickEvent) Xdevice_goset(int){}
func (x Xopenfl_events_JoystickEvent) Xaxis_goget()  XArray { return 0; }
func (x Xopenfl_events_JoystickEvent) Xaxis_goset(XArray){}
func Xopenfl_events_JoystickEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xdevice int,Xid int,Xx float64,Xy float64,Xz float64) Xopenfl_events_JoystickEvent { return Xopenfl_events_JoystickEvent(0); }
func Xopenfl_events_SystemEvent_SYSTEM_goget()  string { return ""; }
func Xopenfl_events_SystemEvent_SYSTEM_goset(string){}
type Xopenfl_events_SystemEvent uintptr // TClassdecl
func (x Xopenfl_events_SystemEvent) Xdata_goget()  int { return 0; }
func (x Xopenfl_events_SystemEvent) Xdata_goset(int){}
func Xopenfl_events_SystemEvent_new(Xtype string,Xbubbles bool,Xcancelable bool,Xdata int) Xopenfl_events_SystemEvent { return Xopenfl_events_SystemEvent(0); }
func Xopenfl_gl_GL_DEPTH___BUFFER___BIT_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___BUFFER___BIT_goset(int){}
func Xopenfl_gl_GL_STENCIL___BUFFER___BIT_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BUFFER___BIT_goset(int){}
func Xopenfl_gl_GL_COLOR___BUFFER___BIT_goget()  int { return 0; }
func Xopenfl_gl_GL_COLOR___BUFFER___BIT_goset(int){}
func Xopenfl_gl_GL_POINTS_goget()  int { return 0; }
func Xopenfl_gl_GL_POINTS_goset(int){}
func Xopenfl_gl_GL_LINES_goget()  int { return 0; }
func Xopenfl_gl_GL_LINES_goset(int){}
func Xopenfl_gl_GL_LINE___LOOP_goget()  int { return 0; }
func Xopenfl_gl_GL_LINE___LOOP_goset(int){}
func Xopenfl_gl_GL_LINE___STRIP_goget()  int { return 0; }
func Xopenfl_gl_GL_LINE___STRIP_goset(int){}
func Xopenfl_gl_GL_TRIANGLES_goget()  int { return 0; }
func Xopenfl_gl_GL_TRIANGLES_goset(int){}
func Xopenfl_gl_GL_TRIANGLE___STRIP_goget()  int { return 0; }
func Xopenfl_gl_GL_TRIANGLE___STRIP_goset(int){}
func Xopenfl_gl_GL_TRIANGLE___FAN_goget()  int { return 0; }
func Xopenfl_gl_GL_TRIANGLE___FAN_goset(int){}
func Xopenfl_gl_GL_ZERO_goget()  int { return 0; }
func Xopenfl_gl_GL_ZERO_goset(int){}
func Xopenfl_gl_GL_ONE_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE_goset(int){}
func Xopenfl_gl_GL_SRC___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_SRC___COLOR_goset(int){}
func Xopenfl_gl_GL_ONE___MINUS___SRC___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE___MINUS___SRC___COLOR_goset(int){}
func Xopenfl_gl_GL_SRC___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_SRC___ALPHA_goset(int){}
func Xopenfl_gl_GL_ONE___MINUS___SRC___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE___MINUS___SRC___ALPHA_goset(int){}
func Xopenfl_gl_GL_DST___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_DST___ALPHA_goset(int){}
func Xopenfl_gl_GL_ONE___MINUS___DST___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE___MINUS___DST___ALPHA_goset(int){}
func Xopenfl_gl_GL_DST___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_DST___COLOR_goset(int){}
func Xopenfl_gl_GL_ONE___MINUS___DST___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE___MINUS___DST___COLOR_goset(int){}
func Xopenfl_gl_GL_SRC___ALPHA___SATURATE_goget()  int { return 0; }
func Xopenfl_gl_GL_SRC___ALPHA___SATURATE_goset(int){}
func Xopenfl_gl_GL_FUNC___ADD_goget()  int { return 0; }
func Xopenfl_gl_GL_FUNC___ADD_goset(int){}
func Xopenfl_gl_GL_BLEND___EQUATION_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___EQUATION_goset(int){}
func Xopenfl_gl_GL_BLEND___EQUATION___RGB_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___EQUATION___RGB_goset(int){}
func Xopenfl_gl_GL_BLEND___EQUATION___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___EQUATION___ALPHA_goset(int){}
func Xopenfl_gl_GL_FUNC___SUBTRACT_goget()  int { return 0; }
func Xopenfl_gl_GL_FUNC___SUBTRACT_goset(int){}
func Xopenfl_gl_GL_FUNC___REVERSE___SUBTRACT_goget()  int { return 0; }
func Xopenfl_gl_GL_FUNC___REVERSE___SUBTRACT_goset(int){}
func Xopenfl_gl_GL_BLEND___DST___RGB_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___DST___RGB_goset(int){}
func Xopenfl_gl_GL_BLEND___SRC___RGB_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___SRC___RGB_goset(int){}
func Xopenfl_gl_GL_BLEND___DST___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___DST___ALPHA_goset(int){}
func Xopenfl_gl_GL_BLEND___SRC___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___SRC___ALPHA_goset(int){}
func Xopenfl_gl_GL_CONSTANT___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_CONSTANT___COLOR_goset(int){}
func Xopenfl_gl_GL_ONE___MINUS___CONSTANT___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE___MINUS___CONSTANT___COLOR_goset(int){}
func Xopenfl_gl_GL_CONSTANT___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_CONSTANT___ALPHA_goset(int){}
func Xopenfl_gl_GL_ONE___MINUS___CONSTANT___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_ONE___MINUS___CONSTANT___ALPHA_goset(int){}
func Xopenfl_gl_GL_BLEND___COLOR_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND___COLOR_goset(int){}
func Xopenfl_gl_GL_ARRAY___BUFFER_goget()  int { return 0; }
func Xopenfl_gl_GL_ARRAY___BUFFER_goset(int){}
func Xopenfl_gl_GL_ELEMENT___ARRAY___BUFFER_goget()  int { return 0; }
func Xopenfl_gl_GL_ELEMENT___ARRAY___BUFFER_goset(int){}
func Xopenfl_gl_GL_ARRAY___BUFFER___BINDING_goget()  int { return 0; }
func Xopenfl_gl_GL_ARRAY___BUFFER___BINDING_goset(int){}
func Xopenfl_gl_GL_ELEMENT___ARRAY___BUFFER___BINDING_goget()  int { return 0; }
func Xopenfl_gl_GL_ELEMENT___ARRAY___BUFFER___BINDING_goset(int){}
func Xopenfl_gl_GL_STREAM___DRAW_goget()  int { return 0; }
func Xopenfl_gl_GL_STREAM___DRAW_goset(int){}
func Xopenfl_gl_GL_STATIC___DRAW_goget()  int { return 0; }
func Xopenfl_gl_GL_STATIC___DRAW_goset(int){}
func Xopenfl_gl_GL_DYNAMIC___DRAW_goget()  int { return 0; }
func Xopenfl_gl_GL_DYNAMIC___DRAW_goset(int){}
func Xopenfl_gl_GL_BUFFER___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_BUFFER___SIZE_goset(int){}
func Xopenfl_gl_GL_BUFFER___USAGE_goget()  int { return 0; }
func Xopenfl_gl_GL_BUFFER___USAGE_goset(int){}
func Xopenfl_gl_GL_CURRENT___VERTEX___ATTRIB_goget()  int { return 0; }
func Xopenfl_gl_GL_CURRENT___VERTEX___ATTRIB_goset(int){}
func Xopenfl_gl_GL_FRONT_goget()  int { return 0; }
func Xopenfl_gl_GL_FRONT_goset(int){}
func Xopenfl_gl_GL_BACK_goget()  int { return 0; }
func Xopenfl_gl_GL_BACK_goset(int){}
func Xopenfl_gl_GL_FRONT___AND___BACK_goget()  int { return 0; }
func Xopenfl_gl_GL_FRONT___AND___BACK_goset(int){}
func Xopenfl_gl_GL_CULL___FACE_goget()  int { return 0; }
func Xopenfl_gl_GL_CULL___FACE_goset(int){}
func Xopenfl_gl_GL_BLEND_goget()  int { return 0; }
func Xopenfl_gl_GL_BLEND_goset(int){}
func Xopenfl_gl_GL_DITHER_goget()  int { return 0; }
func Xopenfl_gl_GL_DITHER_goset(int){}
func Xopenfl_gl_GL_STENCIL___TEST_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___TEST_goset(int){}
func Xopenfl_gl_GL_DEPTH___TEST_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___TEST_goset(int){}
func Xopenfl_gl_GL_SCISSOR___TEST_goget()  int { return 0; }
func Xopenfl_gl_GL_SCISSOR___TEST_goset(int){}
func Xopenfl_gl_GL_POLYGON___OFFSET___FILL_goget()  int { return 0; }
func Xopenfl_gl_GL_POLYGON___OFFSET___FILL_goset(int){}
func Xopenfl_gl_GL_SAMPLE___ALPHA___TO___COVERAGE_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLE___ALPHA___TO___COVERAGE_goset(int){}
func Xopenfl_gl_GL_SAMPLE___COVERAGE_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLE___COVERAGE_goset(int){}
func Xopenfl_gl_GL_NO___ERROR_goget()  int { return 0; }
func Xopenfl_gl_GL_NO___ERROR_goset(int){}
func Xopenfl_gl_GL_INVALID___ENUM_goget()  int { return 0; }
func Xopenfl_gl_GL_INVALID___ENUM_goset(int){}
func Xopenfl_gl_GL_INVALID___VALUE_goget()  int { return 0; }
func Xopenfl_gl_GL_INVALID___VALUE_goset(int){}
func Xopenfl_gl_GL_INVALID___OPERATION_goget()  int { return 0; }
func Xopenfl_gl_GL_INVALID___OPERATION_goset(int){}
func Xopenfl_gl_GL_OUT___OF___MEMORY_goget()  int { return 0; }
func Xopenfl_gl_GL_OUT___OF___MEMORY_goset(int){}
func Xopenfl_gl_GL_CW_goget()  int { return 0; }
func Xopenfl_gl_GL_CW_goset(int){}
func Xopenfl_gl_GL_CCW_goget()  int { return 0; }
func Xopenfl_gl_GL_CCW_goset(int){}
func Xopenfl_gl_GL_LINE___WIDTH_goget()  int { return 0; }
func Xopenfl_gl_GL_LINE___WIDTH_goset(int){}
func Xopenfl_gl_GL_ALIASED___POINT___SIZE___RANGE_goget()  int { return 0; }
func Xopenfl_gl_GL_ALIASED___POINT___SIZE___RANGE_goset(int){}
func Xopenfl_gl_GL_ALIASED___LINE___WIDTH___RANGE_goget()  int { return 0; }
func Xopenfl_gl_GL_ALIASED___LINE___WIDTH___RANGE_goset(int){}
func Xopenfl_gl_GL_CULL___FACE___MODE_goget()  int { return 0; }
func Xopenfl_gl_GL_CULL___FACE___MODE_goset(int){}
func Xopenfl_gl_GL_FRONT___FACE_goget()  int { return 0; }
func Xopenfl_gl_GL_FRONT___FACE_goset(int){}
func Xopenfl_gl_GL_DEPTH___RANGE_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___RANGE_goset(int){}
func Xopenfl_gl_GL_DEPTH___WRITEMASK_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___WRITEMASK_goset(int){}
func Xopenfl_gl_GL_DEPTH___CLEAR___VALUE_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___CLEAR___VALUE_goset(int){}
func Xopenfl_gl_GL_DEPTH___FUNC_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___FUNC_goset(int){}
func Xopenfl_gl_GL_STENCIL___CLEAR___VALUE_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___CLEAR___VALUE_goset(int){}
func Xopenfl_gl_GL_STENCIL___FUNC_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___FUNC_goset(int){}
func Xopenfl_gl_GL_STENCIL___FAIL_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___FAIL_goset(int){}
func Xopenfl_gl_GL_STENCIL___PASS___DEPTH___FAIL_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___PASS___DEPTH___FAIL_goset(int){}
func Xopenfl_gl_GL_STENCIL___PASS___DEPTH___PASS_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___PASS___DEPTH___PASS_goset(int){}
func Xopenfl_gl_GL_STENCIL___REF_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___REF_goset(int){}
func Xopenfl_gl_GL_STENCIL___VALUE___MASK_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___VALUE___MASK_goset(int){}
func Xopenfl_gl_GL_STENCIL___WRITEMASK_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___WRITEMASK_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___FUNC_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___FUNC_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___FAIL_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___FAIL_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___PASS___DEPTH___FAIL_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___PASS___DEPTH___FAIL_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___PASS___DEPTH___PASS_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___PASS___DEPTH___PASS_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___REF_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___REF_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___VALUE___MASK_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___VALUE___MASK_goset(int){}
func Xopenfl_gl_GL_STENCIL___BACK___WRITEMASK_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BACK___WRITEMASK_goset(int){}
func Xopenfl_gl_GL_VIEWPORT_goget()  int { return 0; }
func Xopenfl_gl_GL_VIEWPORT_goset(int){}
func Xopenfl_gl_GL_SCISSOR___BOX_goget()  int { return 0; }
func Xopenfl_gl_GL_SCISSOR___BOX_goset(int){}
func Xopenfl_gl_GL_COLOR___CLEAR___VALUE_goget()  int { return 0; }
func Xopenfl_gl_GL_COLOR___CLEAR___VALUE_goset(int){}
func Xopenfl_gl_GL_COLOR___WRITEMASK_goget()  int { return 0; }
func Xopenfl_gl_GL_COLOR___WRITEMASK_goset(int){}
func Xopenfl_gl_GL_UNPACK___ALIGNMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_UNPACK___ALIGNMENT_goset(int){}
func Xopenfl_gl_GL_PACK___ALIGNMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_PACK___ALIGNMENT_goset(int){}
func Xopenfl_gl_GL_MAX___TEXTURE___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___TEXTURE___SIZE_goset(int){}
func Xopenfl_gl_GL_MAX___VIEWPORT___DIMS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___VIEWPORT___DIMS_goset(int){}
func Xopenfl_gl_GL_SUBPIXEL___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_SUBPIXEL___BITS_goset(int){}
func Xopenfl_gl_GL_RED___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_RED___BITS_goset(int){}
func Xopenfl_gl_GL_GREEN___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_GREEN___BITS_goset(int){}
func Xopenfl_gl_GL_BLUE___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_BLUE___BITS_goset(int){}
func Xopenfl_gl_GL_ALPHA___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_ALPHA___BITS_goset(int){}
func Xopenfl_gl_GL_DEPTH___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___BITS_goset(int){}
func Xopenfl_gl_GL_STENCIL___BITS_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___BITS_goset(int){}
func Xopenfl_gl_GL_POLYGON___OFFSET___UNITS_goget()  int { return 0; }
func Xopenfl_gl_GL_POLYGON___OFFSET___UNITS_goset(int){}
func Xopenfl_gl_GL_POLYGON___OFFSET___FACTOR_goget()  int { return 0; }
func Xopenfl_gl_GL_POLYGON___OFFSET___FACTOR_goset(int){}
func Xopenfl_gl_GL_TEXTURE___BINDING___2D_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___BINDING___2D_goset(int){}
func Xopenfl_gl_GL_SAMPLE___BUFFERS_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLE___BUFFERS_goset(int){}
func Xopenfl_gl_GL_SAMPLES_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLES_goset(int){}
func Xopenfl_gl_GL_SAMPLE___COVERAGE___VALUE_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLE___COVERAGE___VALUE_goset(int){}
func Xopenfl_gl_GL_SAMPLE___COVERAGE___INVERT_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLE___COVERAGE___INVERT_goset(int){}
func Xopenfl_gl_GL_COMPRESSED___TEXTURE___FORMATS_goget()  int { return 0; }
func Xopenfl_gl_GL_COMPRESSED___TEXTURE___FORMATS_goset(int){}
func Xopenfl_gl_GL_DONT___CARE_goget()  int { return 0; }
func Xopenfl_gl_GL_DONT___CARE_goset(int){}
func Xopenfl_gl_GL_FASTEST_goget()  int { return 0; }
func Xopenfl_gl_GL_FASTEST_goset(int){}
func Xopenfl_gl_GL_NICEST_goget()  int { return 0; }
func Xopenfl_gl_GL_NICEST_goset(int){}
func Xopenfl_gl_GL_GENERATE___MIPMAP___HINT_goget()  int { return 0; }
func Xopenfl_gl_GL_GENERATE___MIPMAP___HINT_goset(int){}
func Xopenfl_gl_GL_BYTE_goget()  int { return 0; }
func Xopenfl_gl_GL_BYTE_goset(int){}
func Xopenfl_gl_GL_UNSIGNED___BYTE_goget()  int { return 0; }
func Xopenfl_gl_GL_UNSIGNED___BYTE_goset(int){}
func Xopenfl_gl_GL_SHORT_goget()  int { return 0; }
func Xopenfl_gl_GL_SHORT_goset(int){}
func Xopenfl_gl_GL_UNSIGNED___SHORT_goget()  int { return 0; }
func Xopenfl_gl_GL_UNSIGNED___SHORT_goset(int){}
func Xopenfl_gl_GL_INT_goget()  int { return 0; }
func Xopenfl_gl_GL_INT_goset(int){}
func Xopenfl_gl_GL_UNSIGNED___INT_goget()  int { return 0; }
func Xopenfl_gl_GL_UNSIGNED___INT_goset(int){}
func Xopenfl_gl_GL_FLOAT_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT_goset(int){}
func Xopenfl_gl_GL_DEPTH___COMPONENT_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___COMPONENT_goset(int){}
func Xopenfl_gl_GL_ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_ALPHA_goset(int){}
func Xopenfl_gl_GL_RGB_goget()  int { return 0; }
func Xopenfl_gl_GL_RGB_goset(int){}
func Xopenfl_gl_GL_RGBA_goget()  int { return 0; }
func Xopenfl_gl_GL_RGBA_goset(int){}
func Xopenfl_gl_GL_LUMINANCE_goget()  int { return 0; }
func Xopenfl_gl_GL_LUMINANCE_goset(int){}
func Xopenfl_gl_GL_LUMINANCE___ALPHA_goget()  int { return 0; }
func Xopenfl_gl_GL_LUMINANCE___ALPHA_goset(int){}
func Xopenfl_gl_GL_UNSIGNED___SHORT___4___4___4___4_goget()  int { return 0; }
func Xopenfl_gl_GL_UNSIGNED___SHORT___4___4___4___4_goset(int){}
func Xopenfl_gl_GL_UNSIGNED___SHORT___5___5___5___1_goget()  int { return 0; }
func Xopenfl_gl_GL_UNSIGNED___SHORT___5___5___5___1_goset(int){}
func Xopenfl_gl_GL_UNSIGNED___SHORT___5___6___5_goget()  int { return 0; }
func Xopenfl_gl_GL_UNSIGNED___SHORT___5___6___5_goset(int){}
func Xopenfl_gl_GL_FRAGMENT___SHADER_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAGMENT___SHADER_goset(int){}
func Xopenfl_gl_GL_VERTEX___SHADER_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___SHADER_goset(int){}
func Xopenfl_gl_GL_MAX___VERTEX___ATTRIBS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___VERTEX___ATTRIBS_goset(int){}
func Xopenfl_gl_GL_MAX___VERTEX___UNIFORM___VECTORS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___VERTEX___UNIFORM___VECTORS_goset(int){}
func Xopenfl_gl_GL_MAX___VARYING___VECTORS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___VARYING___VECTORS_goset(int){}
func Xopenfl_gl_GL_MAX___COMBINED___TEXTURE___IMAGE___UNITS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___COMBINED___TEXTURE___IMAGE___UNITS_goset(int){}
func Xopenfl_gl_GL_MAX___VERTEX___TEXTURE___IMAGE___UNITS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___VERTEX___TEXTURE___IMAGE___UNITS_goset(int){}
func Xopenfl_gl_GL_MAX___TEXTURE___IMAGE___UNITS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___TEXTURE___IMAGE___UNITS_goset(int){}
func Xopenfl_gl_GL_MAX___FRAGMENT___UNIFORM___VECTORS_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___FRAGMENT___UNIFORM___VECTORS_goset(int){}
func Xopenfl_gl_GL_SHADER___TYPE_goget()  int { return 0; }
func Xopenfl_gl_GL_SHADER___TYPE_goset(int){}
func Xopenfl_gl_GL_DELETE___STATUS_goget()  int { return 0; }
func Xopenfl_gl_GL_DELETE___STATUS_goset(int){}
func Xopenfl_gl_GL_LINK___STATUS_goget()  int { return 0; }
func Xopenfl_gl_GL_LINK___STATUS_goset(int){}
func Xopenfl_gl_GL_VALIDATE___STATUS_goget()  int { return 0; }
func Xopenfl_gl_GL_VALIDATE___STATUS_goset(int){}
func Xopenfl_gl_GL_ATTACHED___SHADERS_goget()  int { return 0; }
func Xopenfl_gl_GL_ATTACHED___SHADERS_goset(int){}
func Xopenfl_gl_GL_ACTIVE___UNIFORMS_goget()  int { return 0; }
func Xopenfl_gl_GL_ACTIVE___UNIFORMS_goset(int){}
func Xopenfl_gl_GL_ACTIVE___ATTRIBUTES_goget()  int { return 0; }
func Xopenfl_gl_GL_ACTIVE___ATTRIBUTES_goset(int){}
func Xopenfl_gl_GL_SHADING___LANGUAGE___VERSION_goget()  int { return 0; }
func Xopenfl_gl_GL_SHADING___LANGUAGE___VERSION_goset(int){}
func Xopenfl_gl_GL_CURRENT___PROGRAM_goget()  int { return 0; }
func Xopenfl_gl_GL_CURRENT___PROGRAM_goset(int){}
func Xopenfl_gl_GL_NEVER_goget()  int { return 0; }
func Xopenfl_gl_GL_NEVER_goset(int){}
func Xopenfl_gl_GL_LESS_goget()  int { return 0; }
func Xopenfl_gl_GL_LESS_goset(int){}
func Xopenfl_gl_GL_EQUAL_goget()  int { return 0; }
func Xopenfl_gl_GL_EQUAL_goset(int){}
func Xopenfl_gl_GL_LEQUAL_goget()  int { return 0; }
func Xopenfl_gl_GL_LEQUAL_goset(int){}
func Xopenfl_gl_GL_GREATER_goget()  int { return 0; }
func Xopenfl_gl_GL_GREATER_goset(int){}
func Xopenfl_gl_GL_NOTEQUAL_goget()  int { return 0; }
func Xopenfl_gl_GL_NOTEQUAL_goset(int){}
func Xopenfl_gl_GL_GEQUAL_goget()  int { return 0; }
func Xopenfl_gl_GL_GEQUAL_goset(int){}
func Xopenfl_gl_GL_ALWAYS_goget()  int { return 0; }
func Xopenfl_gl_GL_ALWAYS_goset(int){}
func Xopenfl_gl_GL_KEEP_goget()  int { return 0; }
func Xopenfl_gl_GL_KEEP_goset(int){}
func Xopenfl_gl_GL_REPLACE_goget()  int { return 0; }
func Xopenfl_gl_GL_REPLACE_goset(int){}
func Xopenfl_gl_GL_INCR_goget()  int { return 0; }
func Xopenfl_gl_GL_INCR_goset(int){}
func Xopenfl_gl_GL_DECR_goget()  int { return 0; }
func Xopenfl_gl_GL_DECR_goset(int){}
func Xopenfl_gl_GL_INVERT_goget()  int { return 0; }
func Xopenfl_gl_GL_INVERT_goset(int){}
func Xopenfl_gl_GL_INCR___WRAP_goget()  int { return 0; }
func Xopenfl_gl_GL_INCR___WRAP_goset(int){}
func Xopenfl_gl_GL_DECR___WRAP_goget()  int { return 0; }
func Xopenfl_gl_GL_DECR___WRAP_goset(int){}
func Xopenfl_gl_GL_VENDOR_goget()  int { return 0; }
func Xopenfl_gl_GL_VENDOR_goset(int){}
func Xopenfl_gl_GL_RENDERER_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERER_goset(int){}
func Xopenfl_gl_GL_VERSION_goget()  int { return 0; }
func Xopenfl_gl_GL_VERSION_goset(int){}
func Xopenfl_gl_GL_NEAREST_goget()  int { return 0; }
func Xopenfl_gl_GL_NEAREST_goset(int){}
func Xopenfl_gl_GL_LINEAR_goget()  int { return 0; }
func Xopenfl_gl_GL_LINEAR_goset(int){}
func Xopenfl_gl_GL_NEAREST___MIPMAP___NEAREST_goget()  int { return 0; }
func Xopenfl_gl_GL_NEAREST___MIPMAP___NEAREST_goset(int){}
func Xopenfl_gl_GL_LINEAR___MIPMAP___NEAREST_goget()  int { return 0; }
func Xopenfl_gl_GL_LINEAR___MIPMAP___NEAREST_goset(int){}
func Xopenfl_gl_GL_NEAREST___MIPMAP___LINEAR_goget()  int { return 0; }
func Xopenfl_gl_GL_NEAREST___MIPMAP___LINEAR_goset(int){}
func Xopenfl_gl_GL_LINEAR___MIPMAP___LINEAR_goget()  int { return 0; }
func Xopenfl_gl_GL_LINEAR___MIPMAP___LINEAR_goset(int){}
func Xopenfl_gl_GL_TEXTURE___MAG___FILTER_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___MAG___FILTER_goset(int){}
func Xopenfl_gl_GL_TEXTURE___MIN___FILTER_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___MIN___FILTER_goset(int){}
func Xopenfl_gl_GL_TEXTURE___WRAP___S_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___WRAP___S_goset(int){}
func Xopenfl_gl_GL_TEXTURE___WRAP___T_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___WRAP___T_goset(int){}
func Xopenfl_gl_GL_TEXTURE___2D_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___2D_goset(int){}
func Xopenfl_gl_GL_TEXTURE_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP_goset(int){}
func Xopenfl_gl_GL_TEXTURE___BINDING___CUBE___MAP_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___BINDING___CUBE___MAP_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___POSITIVE___X_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___POSITIVE___X_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___NEGATIVE___X_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___NEGATIVE___X_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___POSITIVE___Y_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___POSITIVE___Y_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___NEGATIVE___Y_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___NEGATIVE___Y_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___POSITIVE___Z_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___POSITIVE___Z_goset(int){}
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___NEGATIVE___Z_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE___CUBE___MAP___NEGATIVE___Z_goset(int){}
func Xopenfl_gl_GL_MAX___CUBE___MAP___TEXTURE___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___CUBE___MAP___TEXTURE___SIZE_goset(int){}
func Xopenfl_gl_GL_TEXTURE0_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE0_goset(int){}
func Xopenfl_gl_GL_TEXTURE1_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE1_goset(int){}
func Xopenfl_gl_GL_TEXTURE2_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE2_goset(int){}
func Xopenfl_gl_GL_TEXTURE3_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE3_goset(int){}
func Xopenfl_gl_GL_TEXTURE4_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE4_goset(int){}
func Xopenfl_gl_GL_TEXTURE5_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE5_goset(int){}
func Xopenfl_gl_GL_TEXTURE6_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE6_goset(int){}
func Xopenfl_gl_GL_TEXTURE7_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE7_goset(int){}
func Xopenfl_gl_GL_TEXTURE8_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE8_goset(int){}
func Xopenfl_gl_GL_TEXTURE9_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE9_goset(int){}
func Xopenfl_gl_GL_TEXTURE10_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE10_goset(int){}
func Xopenfl_gl_GL_TEXTURE11_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE11_goset(int){}
func Xopenfl_gl_GL_TEXTURE12_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE12_goset(int){}
func Xopenfl_gl_GL_TEXTURE13_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE13_goset(int){}
func Xopenfl_gl_GL_TEXTURE14_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE14_goset(int){}
func Xopenfl_gl_GL_TEXTURE15_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE15_goset(int){}
func Xopenfl_gl_GL_TEXTURE16_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE16_goset(int){}
func Xopenfl_gl_GL_TEXTURE17_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE17_goset(int){}
func Xopenfl_gl_GL_TEXTURE18_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE18_goset(int){}
func Xopenfl_gl_GL_TEXTURE19_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE19_goset(int){}
func Xopenfl_gl_GL_TEXTURE20_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE20_goset(int){}
func Xopenfl_gl_GL_TEXTURE21_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE21_goset(int){}
func Xopenfl_gl_GL_TEXTURE22_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE22_goset(int){}
func Xopenfl_gl_GL_TEXTURE23_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE23_goset(int){}
func Xopenfl_gl_GL_TEXTURE24_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE24_goset(int){}
func Xopenfl_gl_GL_TEXTURE25_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE25_goset(int){}
func Xopenfl_gl_GL_TEXTURE26_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE26_goset(int){}
func Xopenfl_gl_GL_TEXTURE27_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE27_goset(int){}
func Xopenfl_gl_GL_TEXTURE28_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE28_goset(int){}
func Xopenfl_gl_GL_TEXTURE29_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE29_goset(int){}
func Xopenfl_gl_GL_TEXTURE30_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE30_goset(int){}
func Xopenfl_gl_GL_TEXTURE31_goget()  int { return 0; }
func Xopenfl_gl_GL_TEXTURE31_goset(int){}
func Xopenfl_gl_GL_ACTIVE___TEXTURE_goget()  int { return 0; }
func Xopenfl_gl_GL_ACTIVE___TEXTURE_goset(int){}
func Xopenfl_gl_GL_REPEAT_goget()  int { return 0; }
func Xopenfl_gl_GL_REPEAT_goset(int){}
func Xopenfl_gl_GL_CLAMP___TO___EDGE_goget()  int { return 0; }
func Xopenfl_gl_GL_CLAMP___TO___EDGE_goset(int){}
func Xopenfl_gl_GL_MIRRORED___REPEAT_goget()  int { return 0; }
func Xopenfl_gl_GL_MIRRORED___REPEAT_goset(int){}
func Xopenfl_gl_GL_FLOAT___VEC2_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT___VEC2_goset(int){}
func Xopenfl_gl_GL_FLOAT___VEC3_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT___VEC3_goset(int){}
func Xopenfl_gl_GL_FLOAT___VEC4_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT___VEC4_goset(int){}
func Xopenfl_gl_GL_INT___VEC2_goget()  int { return 0; }
func Xopenfl_gl_GL_INT___VEC2_goset(int){}
func Xopenfl_gl_GL_INT___VEC3_goget()  int { return 0; }
func Xopenfl_gl_GL_INT___VEC3_goset(int){}
func Xopenfl_gl_GL_INT___VEC4_goget()  int { return 0; }
func Xopenfl_gl_GL_INT___VEC4_goset(int){}
func Xopenfl_gl_GL_BOOL_goget()  int { return 0; }
func Xopenfl_gl_GL_BOOL_goset(int){}
func Xopenfl_gl_GL_BOOL___VEC2_goget()  int { return 0; }
func Xopenfl_gl_GL_BOOL___VEC2_goset(int){}
func Xopenfl_gl_GL_BOOL___VEC3_goget()  int { return 0; }
func Xopenfl_gl_GL_BOOL___VEC3_goset(int){}
func Xopenfl_gl_GL_BOOL___VEC4_goget()  int { return 0; }
func Xopenfl_gl_GL_BOOL___VEC4_goset(int){}
func Xopenfl_gl_GL_FLOAT___MAT2_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT___MAT2_goset(int){}
func Xopenfl_gl_GL_FLOAT___MAT3_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT___MAT3_goset(int){}
func Xopenfl_gl_GL_FLOAT___MAT4_goget()  int { return 0; }
func Xopenfl_gl_GL_FLOAT___MAT4_goset(int){}
func Xopenfl_gl_GL_SAMPLER___2D_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLER___2D_goset(int){}
func Xopenfl_gl_GL_SAMPLER___CUBE_goget()  int { return 0; }
func Xopenfl_gl_GL_SAMPLER___CUBE_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___ENABLED_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___ENABLED_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___SIZE_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___STRIDE_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___STRIDE_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___TYPE_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___TYPE_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___NORMALIZED_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___NORMALIZED_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___POINTER_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___POINTER_goset(int){}
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___BUFFER___BINDING_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___ATTRIB___ARRAY___BUFFER___BINDING_goset(int){}
func Xopenfl_gl_GL_VERTEX___PROGRAM___POINT___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_VERTEX___PROGRAM___POINT___SIZE_goset(int){}
func Xopenfl_gl_GL_POINT___SPRITE_goget()  int { return 0; }
func Xopenfl_gl_GL_POINT___SPRITE_goset(int){}
func Xopenfl_gl_GL_COMPILE___STATUS_goget()  int { return 0; }
func Xopenfl_gl_GL_COMPILE___STATUS_goset(int){}
func Xopenfl_gl_GL_LOW___FLOAT_goget()  int { return 0; }
func Xopenfl_gl_GL_LOW___FLOAT_goset(int){}
func Xopenfl_gl_GL_MEDIUM___FLOAT_goget()  int { return 0; }
func Xopenfl_gl_GL_MEDIUM___FLOAT_goset(int){}
func Xopenfl_gl_GL_HIGH___FLOAT_goget()  int { return 0; }
func Xopenfl_gl_GL_HIGH___FLOAT_goset(int){}
func Xopenfl_gl_GL_LOW___INT_goget()  int { return 0; }
func Xopenfl_gl_GL_LOW___INT_goset(int){}
func Xopenfl_gl_GL_MEDIUM___INT_goget()  int { return 0; }
func Xopenfl_gl_GL_MEDIUM___INT_goset(int){}
func Xopenfl_gl_GL_HIGH___INT_goget()  int { return 0; }
func Xopenfl_gl_GL_HIGH___INT_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER_goset(int){}
func Xopenfl_gl_GL_RGBA4_goget()  int { return 0; }
func Xopenfl_gl_GL_RGBA4_goset(int){}
func Xopenfl_gl_GL_RGB5___A1_goget()  int { return 0; }
func Xopenfl_gl_GL_RGB5___A1_goset(int){}
func Xopenfl_gl_GL_RGB565_goget()  int { return 0; }
func Xopenfl_gl_GL_RGB565_goset(int){}
func Xopenfl_gl_GL_DEPTH___COMPONENT16_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___COMPONENT16_goset(int){}
func Xopenfl_gl_GL_STENCIL___INDEX_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___INDEX_goset(int){}
func Xopenfl_gl_GL_STENCIL___INDEX8_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___INDEX8_goset(int){}
func Xopenfl_gl_GL_DEPTH___STENCIL_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___STENCIL_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___WIDTH_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___WIDTH_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___HEIGHT_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___HEIGHT_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___INTERNAL___FORMAT_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___INTERNAL___FORMAT_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___RED___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___RED___SIZE_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___GREEN___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___GREEN___SIZE_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___BLUE___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___BLUE___SIZE_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___ALPHA___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___ALPHA___SIZE_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___DEPTH___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___DEPTH___SIZE_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___STENCIL___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___STENCIL___SIZE_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___OBJECT___TYPE_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___OBJECT___TYPE_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___OBJECT___NAME_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___OBJECT___NAME_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___TEXTURE___LEVEL_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___TEXTURE___LEVEL_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___TEXTURE___CUBE___MAP___FACE_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___ATTACHMENT___TEXTURE___CUBE___MAP___FACE_goset(int){}
func Xopenfl_gl_GL_COLOR___ATTACHMENT0_goget()  int { return 0; }
func Xopenfl_gl_GL_COLOR___ATTACHMENT0_goset(int){}
func Xopenfl_gl_GL_DEPTH___ATTACHMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___ATTACHMENT_goset(int){}
func Xopenfl_gl_GL_STENCIL___ATTACHMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_STENCIL___ATTACHMENT_goset(int){}
func Xopenfl_gl_GL_DEPTH___STENCIL___ATTACHMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_DEPTH___STENCIL___ATTACHMENT_goset(int){}
func Xopenfl_gl_GL_NONE_goget()  int { return 0; }
func Xopenfl_gl_GL_NONE_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___COMPLETE_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___COMPLETE_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___INCOMPLETE___ATTACHMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___INCOMPLETE___ATTACHMENT_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___INCOMPLETE___MISSING___ATTACHMENT_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___INCOMPLETE___MISSING___ATTACHMENT_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___INCOMPLETE___DIMENSIONS_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___INCOMPLETE___DIMENSIONS_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___UNSUPPORTED_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___UNSUPPORTED_goset(int){}
func Xopenfl_gl_GL_FRAMEBUFFER___BINDING_goget()  int { return 0; }
func Xopenfl_gl_GL_FRAMEBUFFER___BINDING_goset(int){}
func Xopenfl_gl_GL_RENDERBUFFER___BINDING_goget()  int { return 0; }
func Xopenfl_gl_GL_RENDERBUFFER___BINDING_goset(int){}
func Xopenfl_gl_GL_MAX___RENDERBUFFER___SIZE_goget()  int { return 0; }
func Xopenfl_gl_GL_MAX___RENDERBUFFER___SIZE_goset(int){}
func Xopenfl_gl_GL_INVALID___FRAMEBUFFER___OPERATION_goget()  int { return 0; }
func Xopenfl_gl_GL_INVALID___FRAMEBUFFER___OPERATION_goset(int){}
func Xopenfl_gl_GL_UNPACK___FLIP___Y___WEBGL_goget()  int { return 0; }
func Xopenfl_gl_GL_UNPACK___FLIP___Y___WEBGL_goset(int){}
func Xopenfl_gl_GL_UNPACK___PREMULTIPLY___ALPHA___WEBGL_goget()  int { return 0; }
func Xopenfl_gl_GL_UNPACK___PREMULTIPLY___ALPHA___WEBGL_goset(int){}
func Xopenfl_gl_GL_CONTEXT___LOST___WEBGL_goget()  int { return 0; }
func Xopenfl_gl_GL_CONTEXT___LOST___WEBGL_goset(int){}
func Xopenfl_gl_GL_UNPACK___COLORSPACE___CONVERSION___WEBGL_goget()  int { return 0; }
func Xopenfl_gl_GL_UNPACK___COLORSPACE___CONVERSION___WEBGL_goset(int){}
func Xopenfl_gl_GL_BROWSER___DEFAULT___WEBGL_goget()  int { return 0; }
func Xopenfl_gl_GL_BROWSER___DEFAULT___WEBGL_goset(int){}
func Xopenfl_gl_GL_drawingBufferHeight_goget()  int { return 0; }
func Xopenfl_gl_GL_drawingBufferHeight_goset(int){}
func Xopenfl_gl_GL_drawingBufferWidth_goget()  int { return 0; }
func Xopenfl_gl_GL_drawingBufferWidth_goset(int){}
func Xopenfl_gl_GL_version_goget()  int { return 0; }
func Xopenfl_gl_GL_version_goset(int){}
func Xopenfl_gl_GL_activeTexture(Xtexture int) {}
func Xopenfl_gl_GL_attachShader(Xprogram Xopenfl_gl_GLProgram,Xshader Xopenfl_gl_GLShader) {}
func Xopenfl_gl_GL_bindAttribLocation(Xprogram Xopenfl_gl_GLProgram,Xindex int,Xname string) {}
func Xopenfl_gl_GL_bindBitmapDataTexture(Xtexture Xflash_display_BitmapData) {}
func Xopenfl_gl_GL_bindBuffer(Xtarget int,Xbuffer Xopenfl_gl_GLBuffer) {}
func Xopenfl_gl_GL_bindFramebuffer(Xtarget int,Xframebuffer Xopenfl_gl_GLFramebuffer) {}
func Xopenfl_gl_GL_bindRenderbuffer(Xtarget int,Xrenderbuffer Xopenfl_gl_GLRenderbuffer) {}
func Xopenfl_gl_GL_bindTexture(Xtarget int,Xtexture Xopenfl_gl_GLTexture) {}
func Xopenfl_gl_GL_blendColor(Xred float64,Xgreen float64,Xblue float64,Xalpha float64) {}
func Xopenfl_gl_GL_blendEquation(Xmode int) {}
func Xopenfl_gl_GL_blendEquationSeparate(XmodeRGB int,XmodeAlpha int) {}
func Xopenfl_gl_GL_blendFunc(Xsfactor int,Xdfactor int) {}
func Xopenfl_gl_GL_blendFuncSeparate(XsrcRGB int,XdstRGB int,XsrcAlpha int,XdstAlpha int) {}
func Xopenfl_gl_GL_bufferData(Xtarget int,Xdata Xopenfl_utils_ArrayBufferView,Xusage int) {}
func Xopenfl_gl_GL_bufferSubData(Xtarget int,Xoffset int,Xdata Xopenfl_utils_ArrayBufferView) {}
func Xopenfl_gl_GL_checkFramebufferStatus(Xtarget int) int { return 0; }
func Xopenfl_gl_GL_clear(Xmask int) {}
func Xopenfl_gl_GL_clearColor(Xred float64,Xgreen float64,Xblue float64,Xalpha float64) {}
func Xopenfl_gl_GL_clearDepth(Xdepth float64) {}
func Xopenfl_gl_GL_clearStencil(Xs int) {}
func Xopenfl_gl_GL_colorMask(Xred bool,Xgreen bool,Xblue bool,Xalpha bool) {}
func Xopenfl_gl_GL_compileShader(Xshader Xopenfl_gl_GLShader) {}
func Xopenfl_gl_GL_compressedTexImage2D(Xtarget int,Xlevel int,Xinternalformat int,Xwidth int,Xheight int,Xborder int,Xdata Xopenfl_utils_ArrayBufferView) {}
func Xopenfl_gl_GL_compressedTexSubImage2D(Xtarget int,Xlevel int,Xxoffset int,Xyoffset int,Xwidth int,Xheight int,Xformat int,Xdata Xopenfl_utils_ArrayBufferView) {}
func Xopenfl_gl_GL_copyTexImage2D(Xtarget int,Xlevel int,Xinternalformat int,Xx int,Xy int,Xwidth int,Xheight int,Xborder int) {}
func Xopenfl_gl_GL_copyTexSubImage2D(Xtarget int,Xlevel int,Xxoffset int,Xyoffset int,Xx int,Xy int,Xwidth int,Xheight int) {}
func Xopenfl_gl_GL_createBuffer() Xopenfl_gl_GLBuffer { return 0; }
func Xopenfl_gl_GL_createFramebuffer() Xopenfl_gl_GLFramebuffer { return 0; }
func Xopenfl_gl_GL_createProgram() Xopenfl_gl_GLProgram { return 0; }
func Xopenfl_gl_GL_createRenderbuffer() Xopenfl_gl_GLRenderbuffer { return 0; }
func Xopenfl_gl_GL_createShader(Xtype int) Xopenfl_gl_GLShader { return 0; }
func Xopenfl_gl_GL_createTexture() Xopenfl_gl_GLTexture { return 0; }
func Xopenfl_gl_GL_cullFace(Xmode int) {}
func Xopenfl_gl_GL_deleteBuffer(Xbuffer Xopenfl_gl_GLBuffer) {}
func Xopenfl_gl_GL_deleteFramebuffer(Xframebuffer Xopenfl_gl_GLFramebuffer) {}
func Xopenfl_gl_GL_deleteProgram(Xprogram Xopenfl_gl_GLProgram) {}
func Xopenfl_gl_GL_deleteRenderbuffer(Xrenderbuffer Xopenfl_gl_GLRenderbuffer) {}
func Xopenfl_gl_GL_deleteShader(Xshader Xopenfl_gl_GLShader) {}
func Xopenfl_gl_GL_deleteTexture(Xtexture Xopenfl_gl_GLTexture) {}
func Xopenfl_gl_GL_depthFunc(Xfunc int) {}
func Xopenfl_gl_GL_depthMask(Xflag bool) {}
func Xopenfl_gl_GL_depthRange(XzNear float64,XzFar float64) {}
func Xopenfl_gl_GL_detachShader(Xprogram Xopenfl_gl_GLProgram,Xshader Xopenfl_gl_GLShader) {}
func Xopenfl_gl_GL_disable(Xcap int) {}
func Xopenfl_gl_GL_disableVertexAttribArray(Xindex int) {}
func Xopenfl_gl_GL_drawArrays(Xmode int,Xfirst int,Xcount int) {}
func Xopenfl_gl_GL_drawElements(Xmode int,Xcount int,Xtype int,Xoffset int) {}
func Xopenfl_gl_GL_enable(Xcap int) {}
func Xopenfl_gl_GL_enableVertexAttribArray(Xindex int) {}
func Xopenfl_gl_GL_finish() {}
func Xopenfl_gl_GL_flush() {}
func Xopenfl_gl_GL_framebufferRenderbuffer(Xtarget int,Xattachment int,Xrenderbuffertarget int,Xrenderbuffer Xopenfl_gl_GLRenderbuffer) {}
func Xopenfl_gl_GL_framebufferTexture2D(Xtarget int,Xattachment int,Xtextarget int,Xtexture Xopenfl_gl_GLTexture,Xlevel int) {}
func Xopenfl_gl_GL_frontFace(Xmode int) {}
func Xopenfl_gl_GL_generateMipmap(Xtarget int) {}
func Xopenfl_gl_GL_getActiveAttrib(Xprogram Xopenfl_gl_GLProgram,Xindex int) uintptr { return 0; }
func Xopenfl_gl_GL_getActiveUniform(Xprogram Xopenfl_gl_GLProgram,Xindex int) uintptr { return 0; }
func Xopenfl_gl_GL_getAttachedShaders(Xprogram Xopenfl_gl_GLProgram) XArray { return 0; }
func Xopenfl_gl_GL_getAttribLocation(Xprogram Xopenfl_gl_GLProgram,Xname string) int { return 0; }
func Xopenfl_gl_GL_getBufferParameter(Xtarget int,Xpname int) interface{} { return nil; }
func Xopenfl_gl_GL_getContextAttributes() uintptr { return 0; }
func Xopenfl_gl_GL_getError() int { return 0; }
func Xopenfl_gl_GL_getExtension(Xname string) interface{} { return nil; }
func Xopenfl_gl_GL_getFramebufferAttachmentParameter(Xtarget int,Xattachment int,Xpname int) interface{} { return nil; }
func Xopenfl_gl_GL_getParameter(Xpname int) interface{} { return nil; }
func Xopenfl_gl_GL_getProgramInfoLog(Xprogram Xopenfl_gl_GLProgram) string { return ""; }
func Xopenfl_gl_GL_getProgramParameter(Xprogram Xopenfl_gl_GLProgram,Xpname int) int { return 0; }
func Xopenfl_gl_GL_getRenderbufferParameter(Xtarget int,Xpname int) interface{} { return nil; }
func Xopenfl_gl_GL_getShaderInfoLog(Xshader Xopenfl_gl_GLShader) string { return ""; }
func Xopenfl_gl_GL_getShaderParameter(Xshader Xopenfl_gl_GLShader,Xpname int) int { return 0; }
func Xopenfl_gl_GL_getShaderPrecisionFormat(Xshadertype int,Xprecisiontype int) uintptr { return 0; }
func Xopenfl_gl_GL_getShaderSource(Xshader Xopenfl_gl_GLShader) string { return ""; }
func Xopenfl_gl_GL_getSupportedExtensions() XArray { return 0; }
func Xopenfl_gl_GL_getTexParameter(Xtarget int,Xpname int) interface{} { return nil; }
func Xopenfl_gl_GL_getUniform(Xprogram Xopenfl_gl_GLProgram,Xlocation uintptr) interface{} { return nil; }
func Xopenfl_gl_GL_getUniformLocation(Xprogram Xopenfl_gl_GLProgram,Xname string) uintptr { return 0; }
func Xopenfl_gl_GL_getVertexAttrib(Xindex int,Xpname int) interface{} { return nil; }
func Xopenfl_gl_GL_getVertexAttribOffset(Xindex int,Xpname int) int { return 0; }
func Xopenfl_gl_GL_hint(Xtarget int,Xmode int) {}
func Xopenfl_gl_GL_isBuffer(Xbuffer Xopenfl_gl_GLBuffer) bool { return false; }
func Xopenfl_gl_GL_isEnabled(Xcap int) bool { return false; }
func Xopenfl_gl_GL_isFramebuffer(Xframebuffer Xopenfl_gl_GLFramebuffer) bool { return false; }
func Xopenfl_gl_GL_isProgram(Xprogram Xopenfl_gl_GLProgram) bool { return false; }
func Xopenfl_gl_GL_isRenderbuffer(Xrenderbuffer Xopenfl_gl_GLRenderbuffer) bool { return false; }
func Xopenfl_gl_GL_isShader(Xshader Xopenfl_gl_GLShader) bool { return false; }
func Xopenfl_gl_GL_isTexture(Xtexture Xopenfl_gl_GLTexture) bool { return false; }
func Xopenfl_gl_GL_lineWidth(Xwidth float64) {}
func Xopenfl_gl_GL_linkProgram(Xprogram Xopenfl_gl_GLProgram) {}
func Xopenfl_gl_GL_pixelStorei(Xpname int,Xparam int) {}
func Xopenfl_gl_GL_polygonOffset(Xfactor float64,Xunits float64) {}
func Xopenfl_gl_GL_readPixels(Xx int,Xy int,Xwidth int,Xheight int,Xformat int,Xtype int,Xpixels Xflash_utils_ByteArray) {}
func Xopenfl_gl_GL_renderbufferStorage(Xtarget int,Xinternalformat int,Xwidth int,Xheight int) {}
func Xopenfl_gl_GL_sampleCoverage(Xvalue float64,Xinvert bool) {}
func Xopenfl_gl_GL_scissor(Xx int,Xy int,Xwidth int,Xheight int) {}
func Xopenfl_gl_GL_shaderSource(Xshader Xopenfl_gl_GLShader,Xsource string) {}
func Xopenfl_gl_GL_stencilFunc(Xfunc int,Xref int,Xmask int) {}
func Xopenfl_gl_GL_stencilFuncSeparate(Xface int,Xfunc int,Xref int,Xmask int) {}
func Xopenfl_gl_GL_stencilMask(Xmask int) {}
func Xopenfl_gl_GL_stencilMaskSeparate(Xface int,Xmask int) {}
func Xopenfl_gl_GL_stencilOp(Xfail int,Xzfail int,Xzpass int) {}
func Xopenfl_gl_GL_stencilOpSeparate(Xface int,Xfail int,Xzfail int,Xzpass int) {}
func Xopenfl_gl_GL_texImage2D(Xtarget int,Xlevel int,Xinternalformat int,Xwidth int,Xheight int,Xborder int,Xformat int,Xtype int,Xpixels Xopenfl_utils_ArrayBufferView) {}
func Xopenfl_gl_GL_texParameterf(Xtarget int,Xpname int,Xparam float64) {}
func Xopenfl_gl_GL_texParameteri(Xtarget int,Xpname int,Xparam int) {}
func Xopenfl_gl_GL_texSubImage2D(Xtarget int,Xlevel int,Xxoffset int,Xyoffset int,Xwidth int,Xheight int,Xformat int,Xtype int,Xpixels Xflash_utils_ByteArray) {}
func Xopenfl_gl_GL_uniform1f(Xlocation uintptr,Xx float64) {}
func Xopenfl_gl_GL_uniform1fv(Xlocation uintptr,Xx Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniform1i(Xlocation uintptr,Xx int) {}
func Xopenfl_gl_GL_uniform1iv(Xlocation uintptr,Xv Xopenfl_utils_Int32Array) {}
func Xopenfl_gl_GL_uniform2f(Xlocation uintptr,Xx float64,Xy float64) {}
func Xopenfl_gl_GL_uniform2fv(Xlocation uintptr,Xv Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniform2i(Xlocation uintptr,Xx int,Xy int) {}
func Xopenfl_gl_GL_uniform2iv(Xlocation uintptr,Xv Xopenfl_utils_Int32Array) {}
func Xopenfl_gl_GL_uniform3f(Xlocation uintptr,Xx float64,Xy float64,Xz float64) {}
func Xopenfl_gl_GL_uniform3fv(Xlocation uintptr,Xv Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniform3i(Xlocation uintptr,Xx int,Xy int,Xz int) {}
func Xopenfl_gl_GL_uniform3iv(Xlocation uintptr,Xv Xopenfl_utils_Int32Array) {}
func Xopenfl_gl_GL_uniform4f(Xlocation uintptr,Xx float64,Xy float64,Xz float64,Xw float64) {}
func Xopenfl_gl_GL_uniform4fv(Xlocation uintptr,Xv Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniform4i(Xlocation uintptr,Xx int,Xy int,Xz int,Xw int) {}
func Xopenfl_gl_GL_uniform4iv(Xlocation uintptr,Xv Xopenfl_utils_Int32Array) {}
func Xopenfl_gl_GL_uniformMatrix2fv(Xlocation uintptr,Xtranspose bool,Xv Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniformMatrix3fv(Xlocation uintptr,Xtranspose bool,Xv Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniformMatrix4fv(Xlocation uintptr,Xtranspose bool,Xv Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_uniformMatrix3D(Xlocation uintptr,Xtranspose bool,Xmatrix Xflash_geom_Matrix3D) {}
func Xopenfl_gl_GL_useProgram(Xprogram Xopenfl_gl_GLProgram) {}
func Xopenfl_gl_GL_validateProgram(Xprogram Xopenfl_gl_GLProgram) {}
func Xopenfl_gl_GL_vertexAttrib1f(Xindx int,Xx float64) {}
func Xopenfl_gl_GL_vertexAttrib1fv(Xindx int,Xvalues Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_vertexAttrib2f(Xindx int,Xx float64,Xy float64) {}
func Xopenfl_gl_GL_vertexAttrib2fv(Xindx int,Xvalues Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_vertexAttrib3f(Xindx int,Xx float64,Xy float64,Xz float64) {}
func Xopenfl_gl_GL_vertexAttrib3fv(Xindx int,Xvalues Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_vertexAttrib4f(Xindx int,Xx float64,Xy float64,Xz float64,Xw float64) {}
func Xopenfl_gl_GL_vertexAttrib4fv(Xindx int,Xvalues Xopenfl_utils_Float32Array) {}
func Xopenfl_gl_GL_vertexAttribPointer(Xindx int,Xsize int,Xtype int,Xnormalized bool,Xstride int,Xoffset int) {}
func Xopenfl_gl_GL_viewport(Xx int,Xy int,Xwidth int,Xheight int) {}
type Xopenfl_gl_GL uintptr // TClassdecl
type Xopenfl_gl_ShaderPrecisionFormat uintptr // TTypedecl
type Xopenfl_gl_GLActiveInfo uintptr // TTypedecl
type Xopenfl_gl_GLObject uintptr // TClassdecl
func (x Xopenfl_gl_GLObject) XisInvalid() bool { return false; }
func (x Xopenfl_gl_GLObject) XisValid() bool { return false; }
func (x Xopenfl_gl_GLObject) XtoString() string { return ""; }
func (x Xopenfl_gl_GLObject) Xinvalidate() {}
func (x Xopenfl_gl_GLObject) XgetType() string { return ""; }
func (x Xopenfl_gl_GLObject) Xvalid_goget()  bool { return false; }
func (x Xopenfl_gl_GLObject) Xvalid_goset(bool){}
func (x Xopenfl_gl_GLObject) Xinvalidated_goget()  bool { return false; }
func (x Xopenfl_gl_GLObject) Xinvalidated_goset(bool){}
func (x Xopenfl_gl_GLObject) Xid_goget()  interface{} { return nil; }
func (x Xopenfl_gl_GLObject) Xid_goset(interface{}){}
func Xopenfl_gl_GLObject_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLObject { return Xopenfl_gl_GLObject(0); }
type Xopenfl_gl_GLBuffer uintptr // TClassdecl
func Xopenfl_gl_GLBuffer_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLBuffer { return Xopenfl_gl_GLBuffer(0); }
type Xopenfl_gl_GLContextAttributes uintptr // TTypedecl
type Xopenfl_gl_GLFramebuffer uintptr // TClassdecl
func Xopenfl_gl_GLFramebuffer_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLFramebuffer { return Xopenfl_gl_GLFramebuffer(0); }
type Xopenfl_gl_GLProgram uintptr // TClassdecl
func (x Xopenfl_gl_GLProgram) XgetShaders() XArray { return 0; }
func (x Xopenfl_gl_GLProgram) Xattach(Xs Xopenfl_gl_GLShader) {}
func Xopenfl_gl_GLProgram_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLProgram { return Xopenfl_gl_GLProgram(0); }
type Xopenfl_gl_GLRenderbuffer uintptr // TClassdecl
func Xopenfl_gl_GLRenderbuffer_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLRenderbuffer { return Xopenfl_gl_GLRenderbuffer(0); }
type Xopenfl_gl_GLShader uintptr // TClassdecl
func Xopenfl_gl_GLShader_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLShader { return Xopenfl_gl_GLShader(0); }
type Xopenfl_gl_GLTexture uintptr // TClassdecl
func Xopenfl_gl_GLTexture_new(XinVersion int,XinId interface{}) Xopenfl_gl_GLTexture { return Xopenfl_gl_GLTexture(0); }
type Xopenfl_gl_GLUniformLocation interface{} // TTypedecl
type Xopenfl_utils_ArrayBuffer Xflash_utils_ByteArray // TTypedecl
type Xopenfl_utils_ArrayBufferView uintptr // TClassdecl
type Xopenfl_utils_Float32Array uintptr // TClassdecl
type Xopenfl_utils_Int16Array uintptr // TClassdecl
type Xopenfl_utils_Int32Array uintptr // TClassdecl
type Xopenfl_utils_UInt8Array uintptr // TClassdecl
