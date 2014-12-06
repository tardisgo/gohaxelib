package _haxeapi

//import "github.com/tardisgo/tardisgo/tardisgolib/hx"
type XArray uintptr /* should be: hx.Dynamic */ // TClassdecl

type XArray_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XArray) Xlength_goget()

func (x XArray) Xlength_goset(int)

func (x XArray) Xconcat(Xa XArray)

func (x XArray) Xjoin(Xsep string)

func (x XArray) Xpop()

func (x XArray) Xpush(Xx XArray_T)

func (x XArray) Xreverse()

func (x XArray) Xshift()

func (x XArray) Xslice(Xpos int, Xend int)

func (x XArray) Xsort(Xf func())

func (x XArray) Xsplice(Xpos int, Xlen int)

func (x XArray) XtoString()

func (x XArray) Xunshift(Xx XArray_T)

func (x XArray) Xinsert(Xpos int, Xx XArray_T)

func (x XArray) Xremove(Xx XArray_T)

func (x XArray) XindexOf(Xx XArray_T, XfromIndex int)

func (x XArray) XlastIndexOf(Xx XArray_T, XfromIndex int)

func (x XArray) Xcopy()

func (x XArray) Xiterator()

func (x XArray) Xmap(Xf func())

func (x XArray) Xfilter(Xf func())

func XArray_new() XArray

type XClass uintptr /* TAbstractdecl - CAbstract(Class,{CClass(Class.T,{})}) */
func XDate_now()

func XDate_fromTime(Xt float64)

func XDate_fromString(Xs string)

type XDate uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XDate) XgetTime()

func (x XDate) XgetHours()

func (x XDate) XgetMinutes()

func (x XDate) XgetSeconds()

func (x XDate) XgetFullYear()

func (x XDate) XgetMonth()

func (x XDate) XgetDate()

func (x XDate) XgetDay()

func (x XDate) XtoString()

func XDate_new(Xyear int, Xmonth int, Xday int, Xhour int, Xmin int, Xsec int) XDate

func XDateTools_format(Xd XDate, Xf string)

func XDateTools_delta(Xd XDate, Xt float64)

func XDateTools_getMonthDays(Xd XDate)

func XDateTools_seconds(Xn float64)

func XDateTools_minutes(Xn float64)

func XDateTools_hours(Xn float64)

func XDateTools_days(Xn float64)

func XDateTools_parse(Xt float64)

func XDateTools_make(Xo struct {
	_seconds int
	_ms      float64
	_minutes int
	_hours   int
	_days    int
})

type XDateTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type XEReg uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XEReg) Xmatch(Xs string)

func (x XEReg) Xmatched(Xn int)

func (x XEReg) XmatchedLeft()

func (x XEReg) XmatchedRight()

func (x XEReg) XmatchedPos()

func (x XEReg) XmatchSub(Xs string, Xpos int, Xlen int)

func (x XEReg) Xsplit(Xs string)

func (x XEReg) Xreplace(Xs string, Xby string)

func (x XEReg) Xmap(Xs string, Xf func())

func XEReg_new(Xr string, Xopt string) XEReg

type XEnum uintptr        /* TAbstractdecl - CAbstract(Enum,{CClass(Enum.T,{})}) */
type XEnumValue uintptr   /* TAbstractdecl - CAbstract(EnumValue,{}) */
type XIntIterator uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XIntIterator) XhasNext()

func (x XIntIterator) Xnext()

func XIntIterator_new(Xmin int, Xmax int) XIntIterator

func XLambda_array(Xit XIterable)

func XLambda_list(Xit XIterable)

func XLambda_map(Xit XIterable, Xf func())

func XLambda_mapi(Xit XIterable, Xf func())

func XLambda_has(Xit XIterable, Xelt uintptr)

func XLambda_exists(Xit XIterable, Xf func())

func XLambda_foreach(Xit XIterable, Xf func())

func XLambda_iter(Xit XIterable, Xf func())

func XLambda_filter(Xit XIterable, Xf func())

func XLambda_fold(Xit XIterable, Xf func(), Xfirst uintptr)

func XLambda_count(Xit XIterable, Xpred func())

func XLambda_empty(Xit XIterable)

func XLambda_indexOf(Xit XIterable, Xv uintptr)

func XLambda_find(Xit XIterable, Xf func())

func XLambda_concat(Xa XIterable, Xb XIterable)

type XLambda uintptr /* should be: hx.Dynamic */ // TClassdecl

type XList uintptr /* should be: hx.Dynamic */ // TClassdecl

type XList_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XList) Xlength_goget()

func (x XList) Xlength_goset(int)

func (x XList) Xadd(Xitem XList_T)

func (x XList) Xpush(Xitem XList_T)

func (x XList) Xfirst()

func (x XList) Xlast()

func (x XList) Xpop()

func (x XList) XisEmpty()

func (x XList) Xclear()

func (x XList) Xremove(Xv XList_T)

func (x XList) Xiterator()

func (x XList) XtoString()

func (x XList) Xjoin(Xsep string)

func (x XList) Xfilter(Xf func())

func (x XList) Xmap(Xf func())

func XList_new() XList

type XMap uintptr             /* TAbstractdecl - CClass(IMap,{CClass(Map.K,{}), CClass(Map.V,{})}) */
type X___Map_Hashable uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => hashCode, overloads => null, type => CFunction({},CAbstract(Int,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
type XIMap uintptr            /* should be: hx.Dynamic */ // TClassdecl

type XIMap_K uintptr /* should be: hx.Dynamic */ // TClassdecl

type XIMap_V uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XIMap) Xget(Xk XIMap_K)

func (x XIMap) Xset(Xk XIMap_K, Xv XIMap_V)

func (x XIMap) Xexists(Xk XIMap_K)

func (x XIMap) Xremove(Xk XIMap_K)

func (x XIMap) Xkeys()

func (x XIMap) Xiterator()

func (x XIMap) XtoString()

func XMath_PI_goget()

func XMath_PI_goset(float64)

func XMath_NEGATIVE___INFINITY_goget()

func XMath_NEGATIVE___INFINITY_goset(float64)

func XMath_POSITIVE___INFINITY_goget()

func XMath_POSITIVE___INFINITY_goset(float64)

func XMath_NaN_goget()

func XMath_NaN_goset(float64)

func XMath_abs(Xv float64)

func XMath_min(Xa float64, Xb float64)

func XMath_max(Xa float64, Xb float64)

func XMath_sin(Xv float64)

func XMath_cos(Xv float64)

func XMath_tan(Xv float64)

func XMath_asin(Xv float64)

func XMath_acos(Xv float64)

func XMath_atan(Xv float64)

func XMath_atan2(Xy float64, Xx float64)

func XMath_exp(Xv float64)

func XMath_log(Xv float64)

func XMath_pow(Xv float64, Xexp float64)

func XMath_sqrt(Xv float64)

func XMath_round(Xv float64)

func XMath_floor(Xv float64)

func XMath_ceil(Xv float64)

func XMath_random()

func XMath_ffloor(Xv float64)

func XMath_fceil(Xv float64)

func XMath_fround(Xv float64)

func XMath_isFinite(Xf float64)

func XMath_isNaN(Xf float64)

type XMath uintptr /* should be: hx.Dynamic */ // TClassdecl

func XReflect_hasField(Xo uintptr, Xfield string)

func XReflect_field(Xo uintptr, Xfield string)

func XReflect_setField(Xo uintptr, Xfield string, Xvalue uintptr)

func XReflect_getProperty(Xo uintptr, Xfield string)

func XReflect_setProperty(Xo uintptr, Xfield string, Xvalue uintptr)

func XReflect_callMethod(Xo uintptr, Xfunc uintptr, Xargs XArray)

func XReflect_fields(Xo uintptr)

func XReflect_isFunction(Xf uintptr)

func XReflect_compare(Xa uintptr, Xb uintptr)

func XReflect_compareMethods(Xf1 uintptr, Xf2 uintptr)

func XReflect_isObject(Xv uintptr)

func XReflect_isEnumValue(Xv uintptr)

func XReflect_deleteField(Xo uintptr, Xfield string)

func XReflect_copy(Xo uintptr)

func XReflect_makeVarArgs(Xf func())

type XReflect uintptr /* should be: hx.Dynamic */ // TClassdecl

func XStd_is(Xv uintptr, Xt uintptr)

func XStd_instance(Xvalue uintptr, Xc XClass)

func XStd_string(Xs uintptr)

func XStd_int(Xx float64)

func XStd_parseInt(Xx string)

func XStd_parseFloat(Xx string)

func XStd_random(Xx int)

type XStd uintptr /* should be: hx.Dynamic */ // TClassdecl

type XVoid uintptr     /* TAbstractdecl - CAbstract(Void,{}) */
type XFloat uintptr    /* TAbstractdecl - CAbstract(Float,{}) */
type XInt uintptr      /* TAbstractdecl - CAbstract(Int,{}) */
type XNull uintptr     /* TTypedecl -  CClass(Null.T,{}) */
type XBool uintptr     /* TAbstractdecl - CAbstract(Bool,{}) */
type XDynamic uintptr  /* TAbstractdecl - CAbstract(Dynamic,{CClass(Dynamic.T,{})}) */
type XIterator uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => next, overloads => null, type => CFunction({},CClass(Iterator.T,{})), params => [], doc => Returns the current item of the Iterator and advances to the next one.

This method is not required to check hasNext() first. A call to this
method while hasNext() is false yields unspecified behavior., get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => hasNext, overloads => null, type => CFunction({},CAbstract(Bool,{})), params => [], doc => Returns false if the iteration is complete, true otherwise.

Usually iteration is considered to be complete if all elements of the
underlying data structure were handled through calls to next(). However,
in custom iterators any logic may be used to determine the completion
state., get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
type XIterable uintptr    /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => iterator, overloads => null, type => CFunction({},CTypedef(Iterator,{CClass(Iterable.T,{})})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
type XArrayAccess uintptr /* should be: hx.Dynamic */ // TClassdecl

type XArrayAccess_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func XString_fromCharCode(Xcode int)

type XString uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XString) Xlength_goget()

func (x XString) Xlength_goset(int)

func (x XString) XtoUpperCase()

func (x XString) XtoLowerCase()

func (x XString) XcharAt(Xindex int)

func (x XString) XcharCodeAt(Xindex int)

func (x XString) XindexOf(Xstr string, XstartIndex int)

func (x XString) XlastIndexOf(Xstr string, XstartIndex int)

func (x XString) Xsplit(Xdelimiter string)

func (x XString) Xsubstr(Xpos int, Xlen int)

func (x XString) Xsubstring(XstartIndex int, XendIndex int)

func (x XString) XtoString()

func XString_new(Xstring string) XString

type XStringBuf uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XStringBuf) Xlength_goget()

func (x XStringBuf) Xlength_goset(int)

func (x XStringBuf) Xadd(Xx uintptr)

func (x XStringBuf) XaddChar(Xc int)

func (x XStringBuf) XaddSub(Xs string, Xpos int, Xlen int)

func (x XStringBuf) XtoString()

func XStringBuf_new() XStringBuf

func XStringTools_urlEncode(Xs string)

func XStringTools_urlDecode(Xs string)

func XStringTools_htmlEscape(Xs string, Xquotes bool)

func XStringTools_htmlUnescape(Xs string)

func XStringTools_startsWith(Xs string, Xstart string)

func XStringTools_endsWith(Xs string, Xend string)

func XStringTools_isSpace(Xs string, Xpos int)

func XStringTools_ltrim(Xs string)

func XStringTools_rtrim(Xs string)

func XStringTools_trim(Xs string)

func XStringTools_lpad(Xs string, Xc string, Xl int)

func XStringTools_rpad(Xs string, Xc string, Xl int)

func XStringTools_replace(Xs string, Xsub string, Xby string)

func XStringTools_hex(Xn int, Xdigits int)

func XStringTools_fastCodeAt(Xs string, Xindex int)

func XStringTools_isEof(Xc int)

type XStringTools uintptr /* should be: hx.Dynamic */ // TClassdecl

func XSys_print(Xv uintptr)

func XSys_println(Xv uintptr)

func XSys_args()

func XSys_getEnv(Xs string)

func XSys_putEnv(Xs string, Xv string)

func XSys_environment()

func XSys_sleep(Xseconds float64)

func XSys_setTimeLocale(Xloc string)

func XSys_getCwd()

func XSys_setCwd(Xs string)

func XSys_systemName()

func XSys_command(Xcmd string, Xargs XArray)

func XSys_exit(Xcode int)

func XSys_time()

func XSys_cpuTime()

func XSys_executablePath()

func XSys_getChar(Xecho bool)

func XSys_stdin()

func XSys_stdout()

func XSys_stderr()

type XSys uintptr /* should be: hx.Dynamic */ // TClassdecl

type XValueType uintptr /* Enumdecl - {{ args => null, meta => [], name => TNull, doc => null, platforms => {cross} }, { args => null, meta => [], name => TInt, doc => null, platforms => {cross} }, { args => null, meta => [], name => TFloat, doc => null, platforms => {cross} }, { args => null, meta => [], name => TBool, doc => null, platforms => {cross} }, { args => null, meta => [], name => TObject, doc => null, platforms => {cross} }, { args => null, meta => [], name => TFunction, doc => null, platforms => {cross} }, { args => {{ name => c, t => CAbstract(Class,{CDynamic(null)}), opt => false }}, meta => [], name => TClass, doc => null, platforms => {cross} }, { args => {{ name => e, t => CAbstract(Enum,{CDynamic(null)}), opt => false }}, meta => [], name => TEnum, doc => null, platforms => {cross} }, { args => null, meta => [], name => TUnknown, doc => null, platforms => {cross} }} */
func XType_getClass(Xo uintptr)

func XType_getEnum(Xo XEnumValue)

func XType_getSuperClass(Xc XClass)

func XType_getClassName(Xc XClass)

func XType_getEnumName(Xe XEnum)

func XType_resolveClass(Xname string)

func XType_resolveEnum(Xname string)

func XType_createInstance(Xcl XClass, Xargs XArray)

func XType_createEmptyInstance(Xcl XClass)

func XType_createEnum(Xe XEnum, Xconstr string, Xparams XArray)

func XType_createEnumIndex(Xe XEnum, Xindex int, Xparams XArray)

func XType_getInstanceFields(Xc XClass)

func XType_getClassFields(Xc XClass)

func XType_getEnumConstructs(Xe XEnum)

func XType_typeof(Xv uintptr)

func XType_enumEq(Xa uintptr, Xb uintptr)

func XType_enumConstructor(Xe XEnumValue)

func XType_enumParameters(Xe XEnumValue)

func XType_enumIndex(Xe XEnumValue)

func XType_allEnums(Xe XEnum)

type XType uintptr /* should be: hx.Dynamic */ // TClassdecl

type XUInt uintptr    /* TAbstractdecl - CAbstract(Int,{}) */
type XXmlType uintptr /* Enumdecl - {} */
func XXml_Element_goget()

func XXml_Element_goset(XXmlType)

func XXml_PCData_goget()

func XXml_PCData_goset(XXmlType)

func XXml_CData_goget()

func XXml_CData_goset(XXmlType)

func XXml_Comment_goget()

func XXml_Comment_goset(XXmlType)

func XXml_DocType_goget()

func XXml_DocType_goset(XXmlType)

func XXml_ProcessingInstruction_goget()

func XXml_ProcessingInstruction_goset(XXmlType)

func XXml_Document_goget()

func XXml_Document_goset(XXmlType)

func XXml_parse(Xstr string)

func XXml_createElement(Xname string)

func XXml_createPCData(Xdata string)

func XXml_createCData(Xdata string)

func XXml_createComment(Xdata string)

func XXml_createDocType(Xdata string)

func XXml_createProcessingInstruction(Xdata string)

func XXml_createDocument()

type XXml uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x XXml) XnodeType_goget()

func (x XXml) XnodeType_goset(XXmlType)

func (x XXml) XnodeName_goget()

func (x XXml) XnodeName_goset(string)

func (x XXml) XnodeValue_goget()

func (x XXml) XnodeValue_goset(string)

func (x XXml) Xget(Xatt string)

func (x XXml) Xset(Xatt string, Xvalue string)

func (x XXml) Xremove(Xatt string)

func (x XXml) Xexists(Xatt string)

func (x XXml) Xattributes()

func (x XXml) Xparent_goget()

func (x XXml) Xparent_goset(XXml)

func (x XXml) Xiterator()

func (x XXml) Xelements()

func (x XXml) XelementsNamed(Xname string)

func (x XXml) XfirstChild()

func (x XXml) XfirstElement()

func (x XXml) XaddChild(Xx XXml)

func (x XXml) XremoveChild(Xx XXml)

func (x XXml) XinsertChild(Xx XXml, Xpos int)

func (x XXml) XtoString()

type Xhaxe_StackItem uintptr /* Enumdecl - {{ args => null, meta => [], name => CFunction, doc => null, platforms => {cross} }, { args => {{ name => m, t => CClass(String,{}), opt => false }}, meta => [], name => Module, doc => null, platforms => {cross} }, { args => {{ name => s, t => CTypedef(Null,{CEnum(haxe.StackItem,{})}), opt => false }, { name => file, t => CClass(String,{}), opt => false }, { name => line, t => CAbstract(Int,{}), opt => false }}, meta => [], name => FilePos, doc => null, platforms => {cross} }, { args => {{ name => classname, t => CClass(String,{}), opt => false }, { name => method, t => CClass(String,{}), opt => false }}, meta => [], name => Method, doc => null, platforms => {cross} }, { args => {{ name => v, t => CAbstract(Int,{}), opt => false }}, meta => [], name => LocalFunction, doc => null, platforms => {cross} }} */
func Xhaxe_CallStack_callStack()

func Xhaxe_CallStack_exceptionStack()

func Xhaxe_CallStack_toString(Xstack XArray)

type Xhaxe_CallStack uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_Function uintptr  /* TAbstractdecl - CDynamic(null) */
type Xhaxe_FlatEnum uintptr  /* TAbstractdecl - CDynamic(null) */
type Xhaxe_EnumFlags uintptr /* TAbstractdecl - CAbstract(Int,{}) */
func Xhaxe_EnumTools_getName(Xe XEnum)

func Xhaxe_EnumTools_createByName(Xe XEnum, Xconstr string, Xparams XArray)

func Xhaxe_EnumTools_createByIndex(Xe XEnum, Xindex int, Xparams XArray)

func Xhaxe_EnumTools_createAll(Xe XEnum)

func Xhaxe_EnumTools_getConstructors(Xe XEnum)

type Xhaxe_EnumTools uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_EnumValueTools_equals(Xa uintptr, Xb uintptr)

func Xhaxe_EnumValueTools_getName(Xe XEnumValue)

func Xhaxe_EnumValueTools_getParameters(Xe XEnumValue)

func Xhaxe_EnumValueTools_getIndex(Xe XEnumValue)

type Xhaxe_EnumValueTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe____Http_AbstractSocket uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => write, overloads => null, type => CFunction({{ name => str, t => CClass(String,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => shutdown, overloads => null, type => CFunction({{ name => read, t => CAbstract(Bool,{}), opt => false, value => null }, { name => write, t => CAbstract(Bool,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => setTimeout, overloads => null, type => CFunction({{ name => t, t => CAbstract(Float,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => output, overloads => null, type => CClass(haxe.io.Output,{}), params => [], doc => null, get => RNormal, set => RNo, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => input, overloads => null, type => CClass(haxe.io.Input,{}), params => [], doc => null, get => RNormal, set => RNo, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => connect, overloads => null, type => CFunction({{ name => host, t => CClass(sys.net.Host,{}), opt => false, value => null }, { name => port, t => CAbstract(Int,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => close, overloads => null, type => CFunction({},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
func Xhaxe_Http_PROXY_goget()

func Xhaxe_Http_PROXY_goset(struct {
	_port int
	_host string
	_auth struct {
		_user string
		_pass string
	}
})

func Xhaxe_Http_requestUrl(Xurl string)

type Xhaxe_Http uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_Http) Xurl_goget()

func (x Xhaxe_Http) Xurl_goset(string)

func (x Xhaxe_Http) XresponseData_goget()

func (x Xhaxe_Http) XresponseData_goset(uintptr)

func (x Xhaxe_Http) XnoShutdown_goget()

func (x Xhaxe_Http) XnoShutdown_goset(bool)

func (x Xhaxe_Http) XcnxTimeout_goget()

func (x Xhaxe_Http) XcnxTimeout_goset(float64)

func (x Xhaxe_Http) XresponseHeaders_goget()

func (x Xhaxe_Http) XresponseHeaders_goset(Xhaxe_ds_StringMap)

func (x Xhaxe_Http) XsetHeader(Xheader string, Xvalue string)

func (x Xhaxe_Http) XaddHeader(Xheader string, Xvalue string)

func (x Xhaxe_Http) XsetParameter(Xparam string, Xvalue string)

func (x Xhaxe_Http) XaddParameter(Xparam string, Xvalue string)

func (x Xhaxe_Http) XsetPostData(Xdata string)

func (x Xhaxe_Http) Xrequest(Xpost bool)

func (x Xhaxe_Http) XfileTransfert(Xargname string, Xfilename string, Xfile Xhaxe_io_Input, Xsize int, XmimeType string)

func (x Xhaxe_Http) XfileTransfer(Xargname string, Xfilename string, Xfile Xhaxe_io_Input, Xsize int, XmimeType string)

func (x Xhaxe_Http) XcustomRequest(Xpost bool, Xapi Xhaxe_io_Output, Xsock uintptr, Xmethod string)

func (x Xhaxe_Http) XonData(Xdata string)

func (x Xhaxe_Http) XonError(Xmsg string)

func (x Xhaxe_Http) XonStatus(Xstatus int)

func Xhaxe_Http_new(Xurl string) Xhaxe_Http

type Xhaxe_Int32 uintptr /* TAbstractdecl - CAbstract(Int,{}) */
func Xhaxe_Int64_make(Xhigh int, Xlow int)

func Xhaxe_Int64_ofInt(Xx int)

func Xhaxe_Int64_toInt(Xx Xhaxe_Int64)

func Xhaxe_Int64_getLow(Xx Xhaxe_Int64)

func Xhaxe_Int64_getHigh(Xx Xhaxe_Int64)

func Xhaxe_Int64_add(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_sub(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_mul(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_div(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_mod(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_shl(Xa Xhaxe_Int64, Xb int)

func Xhaxe_Int64_shr(Xa Xhaxe_Int64, Xb int)

func Xhaxe_Int64_ushr(Xa Xhaxe_Int64, Xb int)

func Xhaxe_Int64_and(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_or(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_xor(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_neg(Xa Xhaxe_Int64)

func Xhaxe_Int64_isNeg(Xa Xhaxe_Int64)

func Xhaxe_Int64_isZero(Xa Xhaxe_Int64)

func Xhaxe_Int64_compare(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_ucompare(Xa Xhaxe_Int64, Xb Xhaxe_Int64)

func Xhaxe_Int64_toStr(Xa Xhaxe_Int64)

type Xhaxe_Int64 uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_Json_parse(Xtext string)

func Xhaxe_Json_stringify(Xvalue uintptr, Xreplacer func(), Xspace string)

type Xhaxe_Json uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_Log_trace(Xv uintptr, Xinfos Xhaxe_PosInfos)

type Xhaxe_Log uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_PosInfos uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => methodName, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => lineNumber, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fileName, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => customParams, overloads => null, type => CClass(Array,{CDynamic(null)}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => className, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
func Xhaxe_Resource_listNames()

func Xhaxe_Resource_getString(Xname string)

func Xhaxe_Resource_getBytes(Xname string)

type Xhaxe_Resource uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_Serializer_USE___CACHE_goget()

func Xhaxe_Serializer_USE___CACHE_goset(bool)

func Xhaxe_Serializer_USE___ENUM___INDEX_goget()

func Xhaxe_Serializer_USE___ENUM___INDEX_goset(bool)

func Xhaxe_Serializer_run(Xv uintptr)

type Xhaxe_Serializer uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_Serializer) XuseCache_goget()

func (x Xhaxe_Serializer) XuseCache_goset(bool)

func (x Xhaxe_Serializer) XuseEnumIndex_goget()

func (x Xhaxe_Serializer) XuseEnumIndex_goset(bool)

func (x Xhaxe_Serializer) XtoString()

func (x Xhaxe_Serializer) Xserialize(Xv uintptr)

func (x Xhaxe_Serializer) XserializeException(Xe uintptr)

func Xhaxe_Serializer_new() Xhaxe_Serializer

type Xhaxe____Template_TemplateExpr uintptr /* Enumdecl - {{ args => {{ name => v, t => CClass(String,{}), opt => false }}, meta => [], name => OpVar, doc => null, platforms => {cross} }, { args => {{ name => expr, t => CFunction({},CDynamic(null)), opt => false }}, meta => [], name => OpExpr, doc => null, platforms => {cross} }, { args => {{ name => expr, t => CFunction({},CDynamic(null)), opt => false }, { name => eif, t => CEnum(haxe._Template.TemplateExpr,{}), opt => false }, { name => eelse, t => CEnum(haxe._Template.TemplateExpr,{}), opt => false }}, meta => [], name => OpIf, doc => null, platforms => {cross} }, { args => {{ name => str, t => CClass(String,{}), opt => false }}, meta => [], name => OpStr, doc => null, platforms => {cross} }, { args => {{ name => l, t => CClass(List,{CEnum(haxe._Template.TemplateExpr,{})}), opt => false }}, meta => [], name => OpBlock, doc => null, platforms => {cross} }, { args => {{ name => expr, t => CFunction({},CDynamic(null)), opt => false }, { name => loop, t => CEnum(haxe._Template.TemplateExpr,{}), opt => false }}, meta => [], name => OpForeach, doc => null, platforms => {cross} }, { args => {{ name => name, t => CClass(String,{}), opt => false }, { name => params, t => CClass(List,{CEnum(haxe._Template.TemplateExpr,{})}), opt => false }}, meta => [], name => OpMacro, doc => null, platforms => {cross} }} */
type Xhaxe____Template_Token uintptr        /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => s, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => p, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => l, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe____Template_ExprToken uintptr    /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => s, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => p, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
func Xhaxe_Template_globals_goget()

func Xhaxe_Template_globals_goset(uintptr)

type Xhaxe_Template uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_Template) Xexecute(Xcontext uintptr, Xmacros uintptr)

func Xhaxe_Template_new(Xstr string) Xhaxe_Template

func Xhaxe_Timer_measure(Xf func(), Xpos Xhaxe_PosInfos)

func Xhaxe_Timer_stamp()

type Xhaxe_Timer uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_Ucs2 uintptr /* TAbstractdecl - CClass(String,{}) */
func Xhaxe_Unserializer_DEFAULT___RESOLVER_goget()

func Xhaxe_Unserializer_DEFAULT___RESOLVER_goset(uintptr)

func Xhaxe_Unserializer_run(Xv string)

type Xhaxe_Unserializer uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_Unserializer) XsetResolver(Xr uintptr)

func (x Xhaxe_Unserializer) XgetResolver()

func (x Xhaxe_Unserializer) Xunserialize()

func Xhaxe_Unserializer_new(Xbuf string) Xhaxe_Unserializer

func Xhaxe_Utf8_iter(Xs string, Xchars func())

func Xhaxe_Utf8_encode(Xs string)

func Xhaxe_Utf8_decode(Xs string)

func Xhaxe_Utf8_charCodeAt(Xs string, Xindex int)

func Xhaxe_Utf8_validate(Xs string)

func Xhaxe_Utf8_length(Xs string)

func Xhaxe_Utf8_compare(Xa string, Xb string)

func Xhaxe_Utf8_sub(Xs string, Xpos int, Xlen int)

type Xhaxe_Utf8 uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_Utf8) XaddChar(Xc int)

func (x Xhaxe_Utf8) XtoString()

func Xhaxe_Utf8_new(Xsize int) Xhaxe_Utf8

func Xhaxe_crypto_Adler32_read(Xi Xhaxe_io_Input)

func Xhaxe_crypto_Adler32_make(Xb Xhaxe_io_Bytes)

type Xhaxe_crypto_Adler32 uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_crypto_Adler32) Xget()

func (x Xhaxe_crypto_Adler32) Xupdate(Xb Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_crypto_Adler32) Xequals(Xa Xhaxe_crypto_Adler32)

func (x Xhaxe_crypto_Adler32) XtoString()

func Xhaxe_crypto_Adler32_new() Xhaxe_crypto_Adler32

func Xhaxe_crypto_Base64_CHARS_goget()

func Xhaxe_crypto_Base64_CHARS_goset(string)

func Xhaxe_crypto_Base64_BYTES_goget()

func Xhaxe_crypto_Base64_BYTES_goset(Xhaxe_io_Bytes)

func Xhaxe_crypto_Base64_encode(Xbytes Xhaxe_io_Bytes, Xcomplement bool)

func Xhaxe_crypto_Base64_decode(Xstr string, Xcomplement bool)

type Xhaxe_crypto_Base64 uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_crypto_BaseCode_encode(Xs string, Xbase string)

func Xhaxe_crypto_BaseCode_decode(Xs string, Xbase string)

type Xhaxe_crypto_BaseCode uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_crypto_BaseCode) XencodeBytes(Xb Xhaxe_io_Bytes)

func (x Xhaxe_crypto_BaseCode) XdecodeBytes(Xb Xhaxe_io_Bytes)

func (x Xhaxe_crypto_BaseCode) XencodeString(Xs string)

func (x Xhaxe_crypto_BaseCode) XdecodeString(Xs string)

func Xhaxe_crypto_BaseCode_new(Xbase Xhaxe_io_Bytes) Xhaxe_crypto_BaseCode

func Xhaxe_crypto_Crc32_make(Xdata Xhaxe_io_Bytes)

type Xhaxe_crypto_Crc32 uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_crypto_Crc32) Xbyte(Xb int)

func (x Xhaxe_crypto_Crc32) Xupdate(Xb Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_crypto_Crc32) Xget()

func Xhaxe_crypto_Crc32_new() Xhaxe_crypto_Crc32

func Xhaxe_crypto_Md5_encode(Xs string)

func Xhaxe_crypto_Md5_make(Xb Xhaxe_io_Bytes)

type Xhaxe_crypto_Md5 uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_crypto_Sha1_encode(Xs string)

func Xhaxe_crypto_Sha1_make(Xb Xhaxe_io_Bytes)

type Xhaxe_crypto_Sha1 uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_io_Bytes_alloc(Xlength int)

func Xhaxe_io_Bytes_ofString(Xs string)

func Xhaxe_io_Bytes_ofData(Xb Xhaxe_io_BytesData)

func Xhaxe_io_Bytes_fastGet(Xb Xhaxe_io_BytesData, Xpos int)

type Xhaxe_io_Bytes uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_io_Bytes) Xlength_goget()

func (x Xhaxe_io_Bytes) Xlength_goset(int)

func (x Xhaxe_io_Bytes) Xget(Xpos int)

func (x Xhaxe_io_Bytes) Xset(Xpos int, Xv int)

func (x Xhaxe_io_Bytes) Xblit(Xpos int, Xsrc Xhaxe_io_Bytes, Xsrcpos int, Xlen int)

func (x Xhaxe_io_Bytes) Xfill(Xpos int, Xlen int, Xvalue int)

func (x Xhaxe_io_Bytes) Xsub(Xpos int, Xlen int)

func (x Xhaxe_io_Bytes) Xcompare(Xother Xhaxe_io_Bytes)

func (x Xhaxe_io_Bytes) XgetDouble(Xpos int)

func (x Xhaxe_io_Bytes) XgetFloat(Xpos int)

func (x Xhaxe_io_Bytes) XsetDouble(Xpos int, Xv float64)

func (x Xhaxe_io_Bytes) XsetFloat(Xpos int, Xv float64)

func (x Xhaxe_io_Bytes) XgetString(Xpos int, Xlen int)

func (x Xhaxe_io_Bytes) XreadString(Xpos int, Xlen int)

func (x Xhaxe_io_Bytes) XtoString()

func (x Xhaxe_io_Bytes) XtoHex()

func (x Xhaxe_io_Bytes) XgetData()

type Xhaxe_io_Input uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_io_Input) XbigEndian_goget()

func (x Xhaxe_io_Input) XbigEndian_goset(bool)

func (x Xhaxe_io_Input) XreadByte()

func (x Xhaxe_io_Input) XreadBytes(Xs Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_io_Input) Xclose()

func (x Xhaxe_io_Input) XreadAll(Xbufsize int)

func (x Xhaxe_io_Input) XreadFullBytes(Xs Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_io_Input) Xread(Xnbytes int)

func (x Xhaxe_io_Input) XreadUntil(Xend int)

func (x Xhaxe_io_Input) XreadLine()

func (x Xhaxe_io_Input) XreadFloat()

func (x Xhaxe_io_Input) XreadDouble()

func (x Xhaxe_io_Input) XreadInt8()

func (x Xhaxe_io_Input) XreadInt16()

func (x Xhaxe_io_Input) XreadUInt16()

func (x Xhaxe_io_Input) XreadInt24()

func (x Xhaxe_io_Input) XreadUInt24()

func (x Xhaxe_io_Input) XreadInt32()

func (x Xhaxe_io_Input) XreadString(Xlen int)

type Xhaxe_io_BufferInput Xhaxe_io_Input // TClassdecl

func (x Xhaxe_io_BufferInput) Xi_goget()

func (x Xhaxe_io_BufferInput) Xi_goset(Xhaxe_io_Input)

func (x Xhaxe_io_BufferInput) Xbuf_goget()

func (x Xhaxe_io_BufferInput) Xbuf_goset(Xhaxe_io_Bytes)

func (x Xhaxe_io_BufferInput) Xavailable_goget()

func (x Xhaxe_io_BufferInput) Xavailable_goset(int)

func (x Xhaxe_io_BufferInput) Xpos_goget()

func (x Xhaxe_io_BufferInput) Xpos_goset(int)

func (x Xhaxe_io_BufferInput) Xrefill()

func (x Xhaxe_io_BufferInput) XreadByte()

func (x Xhaxe_io_BufferInput) XreadBytes(Xbuf Xhaxe_io_Bytes, Xpos int, Xlen int)

func Xhaxe_io_BufferInput_new(Xi Xhaxe_io_Input, Xbuf Xhaxe_io_Bytes, Xpos int, Xavailable int) Xhaxe_io_BufferInput

type Xhaxe_io_BytesBuffer uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_io_BytesBuffer) Xlength_goget()

func (x Xhaxe_io_BytesBuffer) Xlength_goset(int)

func (x Xhaxe_io_BytesBuffer) XaddByte(Xbyte int)

func (x Xhaxe_io_BytesBuffer) Xadd(Xsrc Xhaxe_io_Bytes)

func (x Xhaxe_io_BytesBuffer) XaddString(Xv string)

func (x Xhaxe_io_BytesBuffer) XaddFloat(Xv float64)

func (x Xhaxe_io_BytesBuffer) XaddDouble(Xv float64)

func (x Xhaxe_io_BytesBuffer) XaddBytes(Xsrc Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_io_BytesBuffer) XgetBytes()

func Xhaxe_io_BytesBuffer_new() Xhaxe_io_BytesBuffer

type Xhaxe_io_BytesData uintptr         /* TTypedecl -  CClass(Array,{CAbstract(Int,{})}) */
type Xhaxe_io_BytesInput Xhaxe_io_Input // TClassdecl

func (x Xhaxe_io_BytesInput) Xposition_goget()

func (x Xhaxe_io_BytesInput) Xposition_goset(int)

func (x Xhaxe_io_BytesInput) Xlength_goget()

func (x Xhaxe_io_BytesInput) Xlength_goset(int)

func (x Xhaxe_io_BytesInput) XreadByte()

func (x Xhaxe_io_BytesInput) XreadBytes(Xbuf Xhaxe_io_Bytes, Xpos int, Xlen int)

func Xhaxe_io_BytesInput_new(Xb Xhaxe_io_Bytes, Xpos int, Xlen int) Xhaxe_io_BytesInput

type Xhaxe_io_Output uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_io_Output) XbigEndian_goget()

func (x Xhaxe_io_Output) XbigEndian_goset(bool)

func (x Xhaxe_io_Output) XwriteByte(Xc int)

func (x Xhaxe_io_Output) XwriteBytes(Xs Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_io_Output) Xflush()

func (x Xhaxe_io_Output) Xclose()

func (x Xhaxe_io_Output) Xwrite(Xs Xhaxe_io_Bytes)

func (x Xhaxe_io_Output) XwriteFullBytes(Xs Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_io_Output) XwriteFloat(Xx float64)

func (x Xhaxe_io_Output) XwriteDouble(Xx float64)

func (x Xhaxe_io_Output) XwriteInt8(Xx int)

func (x Xhaxe_io_Output) XwriteInt16(Xx int)

func (x Xhaxe_io_Output) XwriteUInt16(Xx int)

func (x Xhaxe_io_Output) XwriteInt24(Xx int)

func (x Xhaxe_io_Output) XwriteUInt24(Xx int)

func (x Xhaxe_io_Output) XwriteInt32(Xx int)

func (x Xhaxe_io_Output) Xprepare(Xnbytes int)

func (x Xhaxe_io_Output) XwriteInput(Xi Xhaxe_io_Input, Xbufsize int)

func (x Xhaxe_io_Output) XwriteString(Xs string)

type Xhaxe_io_BytesOutput Xhaxe_io_Output // TClassdecl

func (x Xhaxe_io_BytesOutput) Xlength_goget()

func (x Xhaxe_io_BytesOutput) Xlength_goset(int)

func (x Xhaxe_io_BytesOutput) XwriteByte(Xc int)

func (x Xhaxe_io_BytesOutput) XwriteBytes(Xbuf Xhaxe_io_Bytes, Xpos int, Xlen int)

func (x Xhaxe_io_BytesOutput) XgetBytes()

func Xhaxe_io_BytesOutput_new() Xhaxe_io_BytesOutput

type Xhaxe_io_Eof uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_io_Eof_new() Xhaxe_io_Eof

type Xhaxe_io_Error uintptr /* Enumdecl - {{ args => null, meta => [], name => Blocked, doc => The IO is set into nonblocking mode and some data cannot be read or written, platforms => {cross} }, { args => null, meta => [], name => Overflow, doc => An integer value is outside its allowed range, platforms => {cross} }, { args => null, meta => [], name => OutsideBounds, doc => An operation on Bytes is outside of its valid range, platforms => {cross} }, { args => {{ name => e, t => CDynamic(null), opt => false }}, meta => [], name => Custom, doc => Other errors, platforms => {cross} }} */
func Xhaxe_io_Path_withoutExtension(Xpath string)

func Xhaxe_io_Path_withoutDirectory(Xpath string)

func Xhaxe_io_Path_directory(Xpath string)

func Xhaxe_io_Path_extension(Xpath string)

func Xhaxe_io_Path_withExtension(Xpath string, Xext string)

func Xhaxe_io_Path_join(Xpaths XArray)

func Xhaxe_io_Path_normalize(Xpath string)

func Xhaxe_io_Path_addTrailingSlash(Xpath string)

func Xhaxe_io_Path_removeTrailingSlashes(Xpath string)

type Xhaxe_io_Path uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_io_Path) Xdir_goget()

func (x Xhaxe_io_Path) Xdir_goset(string)

func (x Xhaxe_io_Path) Xfile_goget()

func (x Xhaxe_io_Path) Xfile_goset(string)

func (x Xhaxe_io_Path) Xext_goget()

func (x Xhaxe_io_Path) Xext_goset(string)

func (x Xhaxe_io_Path) Xbackslash_goget()

func (x Xhaxe_io_Path) Xbackslash_goset(bool)

func (x Xhaxe_io_Path) XtoString()

func Xhaxe_io_Path_new(Xpath string) Xhaxe_io_Path

type Xhaxe_io_StringInput Xhaxe_io_BytesInput // TClassdecl

func Xhaxe_io_StringInput_new(Xs string) Xhaxe_io_StringInput

func Xhaxe_ds_ArraySort_sort(Xa XArray, Xcmp func())

type Xhaxe_ds_ArraySort uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_BalancedTree uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_BalancedTree_K uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_BalancedTree_V uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_BalancedTree) Xset(Xkey Xhaxe_ds_BalancedTree_K, Xvalue Xhaxe_ds_BalancedTree_V)

func (x Xhaxe_ds_BalancedTree) Xget(Xkey Xhaxe_ds_BalancedTree_K)

func (x Xhaxe_ds_BalancedTree) Xremove(Xkey Xhaxe_ds_BalancedTree_K)

func (x Xhaxe_ds_BalancedTree) Xexists(Xkey Xhaxe_ds_BalancedTree_K)

func (x Xhaxe_ds_BalancedTree) Xiterator()

func (x Xhaxe_ds_BalancedTree) Xkeys()

func (x Xhaxe_ds_BalancedTree) XtoString()

func Xhaxe_ds_BalancedTree_new() Xhaxe_ds_BalancedTree

type Xhaxe_ds_TreeNode uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_TreeNode_K uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_TreeNode_V uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_TreeNode) Xleft_goget()

func (x Xhaxe_ds_TreeNode) Xleft_goset(Xhaxe_ds_TreeNode)

func (x Xhaxe_ds_TreeNode) Xright_goget()

func (x Xhaxe_ds_TreeNode) Xright_goset(Xhaxe_ds_TreeNode)

func (x Xhaxe_ds_TreeNode) Xkey_goget()

func (x Xhaxe_ds_TreeNode) Xkey_goset(Xhaxe_ds_TreeNode_K)

func (x Xhaxe_ds_TreeNode) Xvalue_goget()

func (x Xhaxe_ds_TreeNode) Xvalue_goset(Xhaxe_ds_TreeNode_V)

func (x Xhaxe_ds_TreeNode) Xget___height()

func (x Xhaxe_ds_TreeNode) XtoString()

func Xhaxe_ds_TreeNode_new(Xl Xhaxe_ds_TreeNode, Xk Xhaxe_ds_TreeNode_K, Xv Xhaxe_ds_TreeNode_V, Xr Xhaxe_ds_TreeNode, Xh int) Xhaxe_ds_TreeNode

type Xhaxe_ds_EnumValueMap Xhaxe_ds_BalancedTree // TClassdecl

type Xhaxe_ds_EnumValueMap_K Xhaxe_ds_BalancedTree // TClassdecl

type Xhaxe_ds_EnumValueMap_V Xhaxe_ds_BalancedTree // TClassdecl

func Xhaxe_ds_EnumValueMap_new() Xhaxe_ds_EnumValueMap

type Xhaxe_ds_GenericCell uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_GenericCell_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_GenericCell) Xelt_goget()

func (x Xhaxe_ds_GenericCell) Xelt_goset(Xhaxe_ds_GenericCell_T)

func (x Xhaxe_ds_GenericCell) Xnext_goget()

func (x Xhaxe_ds_GenericCell) Xnext_goset(Xhaxe_ds_GenericCell)

func Xhaxe_ds_GenericCell_new(Xelt Xhaxe_ds_GenericCell_T, Xnext Xhaxe_ds_GenericCell) Xhaxe_ds_GenericCell

type Xhaxe_ds_GenericStack uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_GenericStack_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_GenericStack) Xhead_goget()

func (x Xhaxe_ds_GenericStack) Xhead_goset(Xhaxe_ds_GenericCell)

func (x Xhaxe_ds_GenericStack) Xadd(Xitem Xhaxe_ds_GenericStack_T)

func (x Xhaxe_ds_GenericStack) Xfirst()

func (x Xhaxe_ds_GenericStack) Xpop()

func (x Xhaxe_ds_GenericStack) XisEmpty()

func (x Xhaxe_ds_GenericStack) Xremove(Xv Xhaxe_ds_GenericStack_T)

func (x Xhaxe_ds_GenericStack) Xiterator()

func (x Xhaxe_ds_GenericStack) XtoString()

func Xhaxe_ds_GenericStack_new() Xhaxe_ds_GenericStack

type Xhaxe_ds_HashMap uintptr /* TAbstractdecl - CAnonymous({{ isOverride => false, line => null, meta => [], name => values, overloads => null, type => CClass(haxe.ds.IntMap,{CClass(haxe.ds.HashMap.V,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => keys, overloads => null, type => CClass(haxe.ds.IntMap,{CClass(haxe.ds.HashMap.K,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_ds_IntMap uintptr  /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_IntMap_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_IntMap) Xset(Xkey int, Xvalue Xhaxe_ds_IntMap_T)

func (x Xhaxe_ds_IntMap) Xget(Xkey int)

func (x Xhaxe_ds_IntMap) Xexists(Xkey int)

func (x Xhaxe_ds_IntMap) Xremove(Xkey int)

func (x Xhaxe_ds_IntMap) Xkeys()

func (x Xhaxe_ds_IntMap) Xiterator()

func (x Xhaxe_ds_IntMap) XtoString()

func Xhaxe_ds_IntMap_new() Xhaxe_ds_IntMap

func Xhaxe_ds_ListSort_sort(Xlist uintptr, Xcmp func())

func Xhaxe_ds_ListSort_sortSingleLinked(Xlist uintptr, Xcmp func())

type Xhaxe_ds_ListSort uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_ObjectMap uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_ObjectMap_K uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_ObjectMap_V uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_ObjectMap) Xset(Xkey Xhaxe_ds_ObjectMap_K, Xvalue Xhaxe_ds_ObjectMap_V)

func (x Xhaxe_ds_ObjectMap) Xget(Xkey Xhaxe_ds_ObjectMap_K)

func (x Xhaxe_ds_ObjectMap) Xexists(Xkey Xhaxe_ds_ObjectMap_K)

func (x Xhaxe_ds_ObjectMap) Xremove(Xkey Xhaxe_ds_ObjectMap_K)

func (x Xhaxe_ds_ObjectMap) Xkeys()

func (x Xhaxe_ds_ObjectMap) Xiterator()

func (x Xhaxe_ds_ObjectMap) XtoString()

func Xhaxe_ds_ObjectMap_new() Xhaxe_ds_ObjectMap

type Xhaxe_ds_Option uintptr    /* Enumdecl - {{ args => {{ name => v, t => CClass(haxe.ds.Option.T,{}), opt => false }}, meta => [], name => Some, doc => null, platforms => {cross} }, { args => null, meta => [], name => None, doc => null, platforms => {cross} }} */
type Xhaxe_ds_StringMap uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_StringMap_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_StringMap) Xset(Xkey string, Xvalue Xhaxe_ds_StringMap_T)

func (x Xhaxe_ds_StringMap) Xget(Xkey string)

func (x Xhaxe_ds_StringMap) Xexists(Xkey string)

func (x Xhaxe_ds_StringMap) Xremove(Xkey string)

func (x Xhaxe_ds_StringMap) Xkeys()

func (x Xhaxe_ds_StringMap) Xiterator()

func (x Xhaxe_ds_StringMap) XtoString()

func Xhaxe_ds_StringMap_new() Xhaxe_ds_StringMap

type Xhaxe_ds____Vector_VectorData uintptr /* TTypedecl -  CClass(Array,{CClass(haxe.ds._Vector.VectorData.T,{})}) */
type Xhaxe_ds_Vector uintptr               /* TAbstractdecl - CTypedef(haxe.ds._Vector.VectorData,{CClass(haxe.ds.Vector.T,{})}) */
type Xhaxe_ds_WeakMap uintptr              /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_WeakMap_K uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_ds_WeakMap_V uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_ds_WeakMap) Xset(Xkey Xhaxe_ds_WeakMap_K, Xvalue Xhaxe_ds_WeakMap_V)

func (x Xhaxe_ds_WeakMap) Xget(Xkey Xhaxe_ds_WeakMap_K)

func (x Xhaxe_ds_WeakMap) Xexists(Xkey Xhaxe_ds_WeakMap_K)

func (x Xhaxe_ds_WeakMap) Xremove(Xkey Xhaxe_ds_WeakMap_K)

func (x Xhaxe_ds_WeakMap) Xkeys()

func (x Xhaxe_ds_WeakMap) Xiterator()

func (x Xhaxe_ds_WeakMap) XtoString()

func Xhaxe_ds_WeakMap_new() Xhaxe_ds_WeakMap

func Xhaxe_format_JsonParser_parse(Xstr string)

type Xhaxe_format_JsonParser uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_format_JsonPrinter_print(Xo uintptr, Xreplacer func(), Xspace string)

type Xhaxe_format_JsonPrinter uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_macro_ComplexTypeTools_toString(Xc Xhaxe_macro_ComplexType)

type Xhaxe_macro_ComplexTypeTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_Position uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => min, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => max, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => file, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Constant uintptr /* Enumdecl - {{ args => {{ name => v, t => CClass(String,{}), opt => false }}, meta => [], name => CInt, doc => null, platforms => {cross} }, { args => {{ name => f, t => CClass(String,{}), opt => false }}, meta => [], name => CFloat, doc => null, platforms => {cross} }, { args => {{ name => s, t => CClass(String,{}), opt => false }}, meta => [], name => CString, doc => null, platforms => {cross} }, { args => {{ name => s, t => CClass(String,{}), opt => false }}, meta => [], name => CIdent, doc => null, platforms => {cross} }, { args => {{ name => r, t => CClass(String,{}), opt => false }, { name => opt, t => CClass(String,{}), opt => false }}, meta => [], name => CRegexp, doc => null, platforms => {cross} }} */
type Xhaxe_macro_Binop uintptr    /* Enumdecl - {{ args => null, meta => [], name => OpAdd, doc => `+`, platforms => {cross} }, { args => null, meta => [], name => OpMult, doc => `*`, platforms => {cross} }, { args => null, meta => [], name => OpDiv, doc => `/`, platforms => {cross} }, { args => null, meta => [], name => OpSub, doc => `-`, platforms => {cross} }, { args => null, meta => [], name => OpAssign, doc => `=`, platforms => {cross} }, { args => null, meta => [], name => OpEq, doc => `==`, platforms => {cross} }, { args => null, meta => [], name => OpNotEq, doc => `!=`, platforms => {cross} }, { args => null, meta => [], name => OpGt, doc => `>`, platforms => {cross} }, { args => null, meta => [], name => OpGte, doc => `>=`, platforms => {cross} }, { args => null, meta => [], name => OpLt, doc => `<`, platforms => {cross} }, { args => null, meta => [], name => OpLte, doc => `<=`, platforms => {cross} }, { args => null, meta => [], name => OpAnd, doc => `&`, platforms => {cross} }, { args => null, meta => [], name => OpOr, doc => `|`, platforms => {cross} }, { args => null, meta => [], name => OpXor, doc => `^`, platforms => {cross} }, { args => null, meta => [], name => OpBoolAnd, doc => `&&`, platforms => {cross} }, { args => null, meta => [], name => OpBoolOr, doc => `||`, platforms => {cross} }, { args => null, meta => [], name => OpShl, doc => `<<`, platforms => {cross} }, { args => null, meta => [], name => OpShr, doc => `>>`, platforms => {cross} }, { args => null, meta => [], name => OpUShr, doc => `>>>`, platforms => {cross} }, { args => null, meta => [], name => OpMod, doc => `%`, platforms => {cross} }, { args => {{ name => op, t => CEnum(haxe.macro.Binop,{}), opt => false }}, meta => [], name => OpAssignOp, doc => `+=`
`-=`
`/=`
`*=`
`<<=`
`>>=`
`>>>=`
`|=`
`&=`
`^=`
`%=`, platforms => {cross} }, { args => null, meta => [], name => OpInterval, doc => `...`, platforms => {cross} }, { args => null, meta => [], name => OpArrow, doc => `=>`, platforms => {cross} }} */
type Xhaxe_macro_Unop uintptr           /* Enumdecl - {{ args => null, meta => [], name => OpIncrement, doc => `++`, platforms => {cross} }, { args => null, meta => [], name => OpDecrement, doc => `--`, platforms => {cross} }, { args => null, meta => [], name => OpNot, doc => `!`, platforms => {cross} }, { args => null, meta => [], name => OpNeg, doc => `-`, platforms => {cross} }, { args => null, meta => [], name => OpNegBits, doc => `~`, platforms => {cross} }} */
type Xhaxe_macro_Expr uintptr           /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CEnum(haxe.macro.ExprDef,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_ExprOf uintptr         /* TTypedecl -  CTypedef(haxe.macro.Expr,{}) */
type Xhaxe_macro_Case uintptr           /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => values, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => guard, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Var uintptr            /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => type, overloads => null, type => CTypedef(Null,{CEnum(haxe.macro.ComplexType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Catch uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.macro.ComplexType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(haxe.macro.Expr,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_ExprDef uintptr        /* Enumdecl - {{ args => {{ name => c, t => CEnum(haxe.macro.Constant,{}), opt => false }}, meta => [], name => EConst, doc => null, platforms => {cross} }, { args => {{ name => e1, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => e2, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EArray, doc => null, platforms => {cross} }, { args => {{ name => op, t => CEnum(haxe.macro.Binop,{}), opt => false }, { name => e1, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => e2, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EBinop, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => field, t => CClass(String,{}), opt => false }}, meta => [], name => EField, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EParenthesis, doc => null, platforms => {cross} }, { args => {{ name => fields, t => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => field, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(haxe.macro.Expr,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), opt => false }}, meta => [], name => EObjectDecl, doc => null, platforms => {cross} }, { args => {{ name => values, t => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), opt => false }}, meta => [], name => EArrayDecl, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => params, t => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), opt => false }}, meta => [], name => ECall, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.TypePath,{}), opt => false }, { name => params, t => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), opt => false }}, meta => [], name => ENew, doc => null, platforms => {cross} }, { args => {{ name => op, t => CEnum(haxe.macro.Unop,{}), opt => false }, { name => postFix, t => CAbstract(Bool,{}), opt => false }, { name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EUnop, doc => null, platforms => {cross} }, { args => {{ name => vars, t => CClass(Array,{CTypedef(haxe.macro.Var,{})}), opt => false }}, meta => [], name => EVars, doc => null, platforms => {cross} }, { args => {{ name => name, t => CTypedef(Null,{CClass(String,{})}), opt => false }, { name => f, t => CTypedef(haxe.macro.Function,{}), opt => false }}, meta => [], name => EFunction, doc => null, platforms => {cross} }, { args => {{ name => exprs, t => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), opt => false }}, meta => [], name => EBlock, doc => null, platforms => {cross} }, { args => {{ name => it, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => expr, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EFor, doc => null, platforms => {cross} }, { args => {{ name => e1, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => e2, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EIn, doc => null, platforms => {cross} }, { args => {{ name => econd, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => eif, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => eelse, t => CTypedef(Null,{CTypedef(haxe.macro.Expr,{})}), opt => false }}, meta => [], name => EIf, doc => null, platforms => {cross} }, { args => {{ name => econd, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => normalWhile, t => CAbstract(Bool,{}), opt => false }}, meta => [], name => EWhile, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => cases, t => CClass(Array,{CTypedef(haxe.macro.Case,{})}), opt => false }, { name => edef, t => CTypedef(Null,{CTypedef(Null,{CTypedef(haxe.macro.Expr,{})})}), opt => false }}, meta => [], name => ESwitch, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => catches, t => CClass(Array,{CTypedef(haxe.macro.Catch,{})}), opt => false }}, meta => [], name => ETry, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => true }}, meta => [], name => EReturn, doc => null, platforms => {cross} }, { args => null, meta => [], name => EBreak, doc => null, platforms => {cross} }, { args => null, meta => [], name => EContinue, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EUntyped, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EThrow, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => t, t => CTypedef(Null,{CEnum(haxe.macro.ComplexType,{})}), opt => false }}, meta => [], name => ECast, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => isCall, t => CAbstract(Bool,{}), opt => false }}, meta => [], name => EDisplay, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.TypePath,{}), opt => false }}, meta => [], name => EDisplayNew, doc => null, platforms => {cross} }, { args => {{ name => econd, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => eif, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => eelse, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => ETernary, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }, { name => t, t => CEnum(haxe.macro.ComplexType,{}), opt => false }}, meta => [], name => ECheckType, doc => null, platforms => {cross} }, { args => {{ name => s, t => CTypedef(haxe.macro.MetadataEntry,{}), opt => false }, { name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => EMeta, doc => null, platforms => {cross} }} */
type Xhaxe_macro_ComplexType uintptr    /* Enumdecl - {{ args => {{ name => p, t => CTypedef(haxe.macro.TypePath,{}), opt => false }}, meta => [], name => TPath, doc => null, platforms => {cross} }, { args => {{ name => args, t => CClass(Array,{CEnum(haxe.macro.ComplexType,{})}), opt => false }, { name => ret, t => CEnum(haxe.macro.ComplexType,{}), opt => false }}, meta => [], name => TFunction, doc => null, platforms => {cross} }, { args => {{ name => fields, t => CClass(Array,{CTypedef(haxe.macro.Field,{})}), opt => false }}, meta => [], name => TAnonymous, doc => null, platforms => {cross} }, { args => {{ name => t, t => CEnum(haxe.macro.ComplexType,{}), opt => false }}, meta => [], name => TParent, doc => null, platforms => {cross} }, { args => {{ name => p, t => CClass(Array,{CTypedef(haxe.macro.TypePath,{})}), opt => false }, { name => fields, t => CClass(Array,{CTypedef(haxe.macro.Field,{})}), opt => false }}, meta => [], name => TExtend, doc => null, platforms => {cross} }, { args => {{ name => t, t => CEnum(haxe.macro.ComplexType,{}), opt => false }}, meta => [], name => TOptional, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TypePath uintptr       /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => sub, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => params, overloads => null, type => CClass(Array,{CEnum(haxe.macro.TypeParam,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_TypeParam uintptr      /* Enumdecl - {{ args => {{ name => t, t => CEnum(haxe.macro.ComplexType,{}), opt => false }}, meta => [], name => TPType, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => TPExpr, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TypeParamDecl uintptr  /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParamDecl,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => constraints, overloads => null, type => CClass(Array,{CEnum(haxe.macro.ComplexType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Function uintptr       /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => ret, overloads => null, type => CTypedef(Null,{CEnum(haxe.macro.ComplexType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParamDecl,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => args, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.FunctionArg,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_FunctionArg uintptr    /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => value, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => type, overloads => null, type => CTypedef(Null,{CEnum(haxe.macro.ComplexType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => opt, overloads => null, type => CTypedef(Null,{CAbstract(Bool,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_MetadataEntry uintptr  /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Metadata uintptr       /* TTypedecl -  CClass(Array,{CTypedef(haxe.macro.MetadataEntry,{})}) */
type Xhaxe_macro_Field uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => meta, overloads => null, type => CTypedef(haxe.macro.Metadata,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => kind, overloads => null, type => CEnum(haxe.macro.FieldType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => access, overloads => null, type => CClass(Array,{CEnum(haxe.macro.Access,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Access uintptr         /* Enumdecl - {{ args => null, meta => [], name => APublic, doc => null, platforms => {cross} }, { args => null, meta => [], name => APrivate, doc => null, platforms => {cross} }, { args => null, meta => [], name => AStatic, doc => null, platforms => {cross} }, { args => null, meta => [], name => AOverride, doc => null, platforms => {cross} }, { args => null, meta => [], name => ADynamic, doc => null, platforms => {cross} }, { args => null, meta => [], name => AInline, doc => null, platforms => {cross} }, { args => null, meta => [], name => AMacro, doc => null, platforms => {cross} }} */
type Xhaxe_macro_FieldType uintptr      /* Enumdecl - {{ args => {{ name => t, t => CTypedef(Null,{CEnum(haxe.macro.ComplexType,{})}), opt => false }, { name => e, t => CTypedef(haxe.macro.Expr,{}), opt => true }}, meta => [], name => FVar, doc => null, platforms => {cross} }, { args => {{ name => f, t => CTypedef(haxe.macro.Function,{}), opt => false }}, meta => [], name => FFun, doc => null, platforms => {cross} }, { args => {{ name => get, t => CClass(String,{}), opt => false }, { name => set, t => CClass(String,{}), opt => false }, { name => t, t => CEnum(haxe.macro.ComplexType,{}), opt => true }, { name => e, t => CTypedef(haxe.macro.Expr,{}), opt => true }}, meta => [], name => FProp, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TypeDefinition uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParamDecl,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => meta, overloads => null, type => CTypedef(haxe.macro.Metadata,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => kind, overloads => null, type => CEnum(haxe.macro.TypeDefKind,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => isExtern, overloads => null, type => CTypedef(Null,{CAbstract(Bool,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fields, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.Field,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_TypeDefKind uintptr    /* Enumdecl - {{ args => null, meta => [], name => TDEnum, doc => null, platforms => {cross} }, { args => null, meta => [], name => TDStructure, doc => null, platforms => {cross} }, { args => {{ name => superClass, t => CTypedef(haxe.macro.TypePath,{}), opt => true }, { name => interfaces, t => CClass(Array,{CTypedef(haxe.macro.TypePath,{})}), opt => true }, { name => isInterface, t => CAbstract(Bool,{}), opt => true }}, meta => [], name => TDClass, doc => null, platforms => {cross} }, { args => {{ name => t, t => CEnum(haxe.macro.ComplexType,{}), opt => false }}, meta => [], name => TDAlias, doc => null, platforms => {cross} }, { args => {{ name => tthis, t => CTypedef(Null,{CEnum(haxe.macro.ComplexType,{})}), opt => false }, { name => from, t => CClass(Array,{CEnum(haxe.macro.ComplexType,{})}), opt => true }, { name => to, t => CClass(Array,{CEnum(haxe.macro.ComplexType,{})}), opt => true }}, meta => [], name => TDAbstract, doc => null, platforms => {cross} }} */
type Xhaxe_macro_Error uintptr          /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_macro_Error) Xmessage_goget()

func (x Xhaxe_macro_Error) Xmessage_goset(string)

func (x Xhaxe_macro_Error) Xpos_goget()

func (x Xhaxe_macro_Error) Xpos_goset(Xhaxe_macro_Position)

func Xhaxe_macro_Error_new(Xm string, Xp Xhaxe_macro_Position) Xhaxe_macro_Error

func Xhaxe_macro_ExprTools_toString(Xe Xhaxe_macro_Expr)

func Xhaxe_macro_ExprTools_iter(Xe Xhaxe_macro_Expr, Xf func())

func Xhaxe_macro_ExprTools_map(Xe Xhaxe_macro_Expr, Xf func())

func Xhaxe_macro_ExprTools_getValue(Xe Xhaxe_macro_Expr)

type Xhaxe_macro_ExprTools uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_macro_ExprArrayTools_map(Xel XArray, Xf func())

func Xhaxe_macro_ExprArrayTools_iter(Xel XArray, Xf func())

type Xhaxe_macro_ExprArrayTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_Format uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_JSGenApi uintptr         /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => types, overloads => null, type => CClass(Array,{CEnum(haxe.macro.Type,{})}), params => [], doc => all the types that were compiled by Haxe, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => setTypeAccessor, overloads => null, type => CFunction({{ name => callb, t => CFunction({{ name => , t => CEnum(haxe.macro.Type,{}), opt => false, value => null }},CClass(String,{})), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => define the JS code that gets generated when a class or enum is accessed in a typed expression, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => setCurrentClass, overloads => null, type => CFunction({{ name => c, t => CTypedef(haxe.macro.ClassType,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => select the current classe, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => quoteString, overloads => null, type => CFunction({{ name => s, t => CClass(String,{}), opt => false, value => null }},CClass(String,{})), params => [], doc => quote and escape the given string constant, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => outputFile, overloads => null, type => CClass(String,{}), params => [], doc => the file in which the JS code can be generated, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => main, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), params => [], doc => the main call expression, if a -main class is defined, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isKeyword, overloads => null, type => CFunction({{ name => ident, t => CClass(String,{}), opt => false, value => null }},CAbstract(Bool,{})), params => [], doc => tells if the given identifier is a JS keyword, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => generateValue, overloads => null, type => CFunction({{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false, value => null }},CClass(String,{})), params => [], doc => generate the JS code for a given typed expression-value, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => generateStatement, overloads => null, type => CFunction({{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false, value => null }},CClass(String,{})), params => [], doc => generate the JS code for any given typed expression, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => buildMetaData, overloads => null, type => CFunction({{ name => t, t => CTypedef(haxe.macro.BaseType,{}), opt => false, value => null }},CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})})), params => [], doc => create the metadata expression for the given type, get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_MacroStringTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_MacroType uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_MacroType_Const uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_Printer uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_macro_Printer) XprintUnop(Xop Xhaxe_macro_Unop)

func (x Xhaxe_macro_Printer) XprintBinop(Xop Xhaxe_macro_Binop)

func (x Xhaxe_macro_Printer) XprintString(Xs string)

func (x Xhaxe_macro_Printer) XprintConstant(Xc Xhaxe_macro_Constant)

func (x Xhaxe_macro_Printer) XprintTypeParam(Xparam Xhaxe_macro_TypeParam)

func (x Xhaxe_macro_Printer) XprintTypePath(Xtp Xhaxe_macro_TypePath)

func (x Xhaxe_macro_Printer) XprintComplexType(Xct Xhaxe_macro_ComplexType)

func (x Xhaxe_macro_Printer) XprintMetadata(Xmeta Xhaxe_macro_MetadataEntry)

func (x Xhaxe_macro_Printer) XprintAccess(Xaccess Xhaxe_macro_Access)

func (x Xhaxe_macro_Printer) XprintField(Xfield Xhaxe_macro_Field)

func (x Xhaxe_macro_Printer) XprintTypeParamDecl(Xtpd Xhaxe_macro_TypeParamDecl)

func (x Xhaxe_macro_Printer) XprintFunctionArg(Xarg Xhaxe_macro_FunctionArg)

func (x Xhaxe_macro_Printer) XprintFunction(Xfunc Xhaxe_macro_Function)

func (x Xhaxe_macro_Printer) XprintVar(Xv Xhaxe_macro_Var)

func (x Xhaxe_macro_Printer) XprintExpr(Xe Xhaxe_macro_Expr)

func (x Xhaxe_macro_Printer) XprintExprs(Xel XArray, Xsep string)

func (x Xhaxe_macro_Printer) XprintTypeDefinition(Xt Xhaxe_macro_TypeDefinition, XprintPackage bool)

func Xhaxe_macro_Printer_new(XtabString string) Xhaxe_macro_Printer

type Xhaxe_macro_TExprTools uintptr        /* TTypedecl -  CClass(haxe.macro.ExprTools,{}) */
type Xhaxe_macro_TComplexTypeTools uintptr /* TTypedecl -  CClass(haxe.macro.ComplexTypeTools,{}) */
type Xhaxe_macro_TTypeTools uintptr        /* TTypedecl -  CClass(haxe.macro.TypeTools,{}) */
type Xhaxe_macro_TMacroStringTools uintptr /* TTypedecl -  CClass(haxe.macro.MacroStringTools,{}) */
type Xhaxe_macro_TTypedExprTools uintptr   /* TTypedecl -  CClass(haxe.macro.TypedExprTools,{}) */
type Xhaxe_macro_Ref uintptr               /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => toString, overloads => null, type => CFunction({},CClass(String,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => get, overloads => null, type => CFunction({},CClass(haxe.macro.Ref.T,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_Type uintptr              /* Enumdecl - {{ args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(Null,{CEnum(haxe.macro.Type,{})})}), opt => false }}, meta => [], name => TMono, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.EnumType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => TEnum, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => TInst, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.DefType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => TType, doc => null, platforms => {cross} }, { args => {{ name => args, t => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => opt, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), opt => false }, { name => ret, t => CEnum(haxe.macro.Type,{}), opt => false }}, meta => [], name => TFun, doc => null, platforms => {cross} }, { args => {{ name => a, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.AnonType,{})}), opt => false }}, meta => [], name => TAnonymous, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(Null,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => TDynamic, doc => null, platforms => {cross} }, { args => {{ name => f, t => CFunction({},CEnum(haxe.macro.Type,{})), opt => false }}, meta => [], name => TLazy, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.AbstractType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => TAbstract, doc => null, platforms => {cross} }} */
type Xhaxe_macro_AnonType uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => status, overloads => null, type => CEnum(haxe.macro.AnonStatus,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fields, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.ClassField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_AnonStatus uintptr        /* Enumdecl - {{ args => null, meta => [], name => AClosed, doc => null, platforms => {cross} }, { args => null, meta => [], name => AOpened, doc => null, platforms => {cross} }, { args => null, meta => [], name => AConst, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }}, meta => [], name => AClassStatics, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.EnumType,{})}), opt => false }}, meta => [], name => AEnumStatics, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.AbstractType,{})}), opt => false }}, meta => [], name => AAbstractStatics, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TypeParameter uintptr     /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_BaseType uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => exclude, overloads => null, type => CFunction({},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_ClassField uintptr        /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => kind, overloads => null, type => CEnum(haxe.macro.FieldKind,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPublic, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CFunction({},CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_ClassKind uintptr         /* Enumdecl - {{ args => null, meta => [], name => KNormal, doc => null, platforms => {cross} }, { args => {{ name => constraints, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => KTypeParameter, doc => null, platforms => {cross} }, { args => {{ name => cl, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => KExtension, doc => null, platforms => {cross} }, { args => {{ name => expr, t => CTypedef(haxe.macro.Expr,{}), opt => false }}, meta => [], name => KExpr, doc => null, platforms => {cross} }, { args => null, meta => [], name => KGeneric, doc => null, platforms => {cross} }, { args => {{ name => cl, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }}, meta => [], name => KGenericInstance, doc => null, platforms => {cross} }, { args => null, meta => [], name => KMacroType, doc => null, platforms => {cross} }, { args => {{ name => a, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.AbstractType,{})}), opt => false }}, meta => [], name => KAbstractImpl, doc => null, platforms => {cross} }, { args => null, meta => [], name => KGenericBuild, doc => null, platforms => {cross} }} */
type Xhaxe_macro_ClassType uintptr         /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => superClass, overloads => null, type => CTypedef(Null,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CEnum(haxe.macro.Type,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => statics, overloads => null, type => CTypedef(haxe.macro.Ref,{CClass(Array,{CTypedef(haxe.macro.ClassField,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => overrides, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassField,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => kind, overloads => null, type => CEnum(haxe.macro.ClassKind,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isInterface, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => interfaces, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CEnum(haxe.macro.Type,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => init, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fields, overloads => null, type => CTypedef(haxe.macro.Ref,{CClass(Array,{CTypedef(haxe.macro.ClassField,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => exclude, overloads => null, type => CFunction({},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => constructor, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassField,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_EnumField uintptr         /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => index, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_EnumType uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => names, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => exclude, overloads => null, type => CFunction({},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => constructs, overloads => null, type => CClass(haxe.ds.StringMap,{CTypedef(haxe.macro.EnumField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_DefType uintptr           /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => exclude, overloads => null, type => CFunction({},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_AbstractType uintptr      /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => unops, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => postFix, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => op, overloads => null, type => CEnum(haxe.macro.Unop,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => field, overloads => null, type => CTypedef(haxe.macro.ClassField,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => to, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => field, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.ClassField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pack, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.macro.MetaAccess,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => impl, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => from, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => field, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.ClassField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => exclude, overloads => null, type => CFunction({},CAbstract(Void,{})), params => [], doc => null, get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => binops, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => op, overloads => null, type => CEnum(haxe.macro.Binop,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => field, overloads => null, type => CTypedef(haxe.macro.ClassField,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => array, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.ClassField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_MetaAccess uintptr        /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => remove, overloads => null, type => CFunction({{ name => name, t => CClass(String,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => Removes all `name` metadata entries from the origin of `this`
MetaAccess.

This method might clear several metadata entries of the same name.

If a `Metadata` array is obtained through a call to `get`, a subsequent
call to `remove` has no effect on that array.

If `name` is null, compilation fails with an error., get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => has, overloads => null, type => CFunction({{ name => name, t => CClass(String,{}), opt => false, value => null }},CAbstract(Bool,{})), params => [], doc => Tells if the origin of `this` MetaAccess has a `name` metadata entry.

If `name` is null, compilation fails with an error., get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => get, overloads => null, type => CFunction({},CTypedef(haxe.macro.Metadata,{})), params => [], doc => Return the wrapped `Metadata` array.

Modifying this array has no effect on the origin of `this` MetaAccess.
The `add` and `remove` methods can be used for that., get => RNormal, set => RMethod, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => add, overloads => null, type => CFunction({{ name => name, t => CClass(String,{}), opt => false, value => null }, { name => params, t => CClass(Array,{CTypedef(haxe.macro.Expr,{})}), opt => false, value => null }, { name => pos, t => CTypedef(haxe.macro.Position,{}), opt => false, value => null }},CAbstract(Void,{})), params => [], doc => Adds the metadata specified by `name`, `params` and `pos` to the origin
of `this` MetaAccess.

Metadata names are not unique during compilation, so this method never
overwrites a previous metadata.

If a `Metadata` array is obtained through a call to `get`, a subsequent
call to `add` has no effect on that array.

If any argument is null, compilation fails with an error., get => RNormal, set => RMethod, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_FieldKind uintptr    /* Enumdecl - {{ args => {{ name => read, t => CEnum(haxe.macro.VarAccess,{}), opt => false }, { name => write, t => CEnum(haxe.macro.VarAccess,{}), opt => false }}, meta => [], name => FVar, doc => null, platforms => {cross} }, { args => {{ name => k, t => CEnum(haxe.macro.MethodKind,{}), opt => false }}, meta => [], name => FMethod, doc => null, platforms => {cross} }} */
type Xhaxe_macro_VarAccess uintptr    /* Enumdecl - {{ args => null, meta => [], name => AccNormal, doc => null, platforms => {cross} }, { args => null, meta => [], name => AccNo, doc => null, platforms => {cross} }, { args => null, meta => [], name => AccNever, doc => null, platforms => {cross} }, { args => null, meta => [], name => AccResolve, doc => null, platforms => {cross} }, { args => null, meta => [], name => AccCall, doc => null, platforms => {cross} }, { args => null, meta => [], name => AccInline, doc => null, platforms => {cross} }, { args => {{ name => r, t => CClass(String,{}), opt => false }, { name => msg, t => CClass(String,{}), opt => true }}, meta => [], name => AccRequire, doc => null, platforms => {cross} }} */
type Xhaxe_macro_MethodKind uintptr   /* Enumdecl - {{ args => null, meta => [], name => MethNormal, doc => null, platforms => {cross} }, { args => null, meta => [], name => MethInline, doc => null, platforms => {cross} }, { args => null, meta => [], name => MethDynamic, doc => null, platforms => {cross} }, { args => null, meta => [], name => MethMacro, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TConstant uintptr    /* Enumdecl - {{ args => {{ name => i, t => CAbstract(Int,{}), opt => false }}, meta => [], name => TInt, doc => null, platforms => {cross} }, { args => {{ name => s, t => CClass(String,{}), opt => false }}, meta => [], name => TFloat, doc => null, platforms => {cross} }, { args => {{ name => s, t => CClass(String,{}), opt => false }}, meta => [], name => TString, doc => null, platforms => {cross} }, { args => {{ name => b, t => CAbstract(Bool,{}), opt => false }}, meta => [], name => TBool, doc => null, platforms => {cross} }, { args => null, meta => [], name => TNull, doc => null, platforms => {cross} }, { args => null, meta => [], name => TThis, doc => null, platforms => {cross} }, { args => null, meta => [], name => TSuper, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TVar uintptr         /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => id, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => extra, overloads => null, type => CTypedef(Null,{CAnonymous({{ isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypeParameter,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => capture, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_ModuleType uintptr   /* Enumdecl - {{ args => {{ name => c, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }}, meta => [], name => TClassDecl, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.EnumType,{})}), opt => false }}, meta => [], name => TEnumDecl, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.DefType,{})}), opt => false }}, meta => [], name => TTypeDecl, doc => null, platforms => {cross} }, { args => {{ name => a, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.AbstractType,{})}), opt => false }}, meta => [], name => TAbstract, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TFunc uintptr        /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(haxe.macro.TypedExpr,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => args, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => value, overloads => null, type => CTypedef(Null,{CEnum(haxe.macro.TConstant,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => v, overloads => null, type => CTypedef(haxe.macro.TVar,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_macro_FieldAccess uintptr  /* Enumdecl - {{ args => {{ name => c, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }, { name => cf, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassField,{})}), opt => false }}, meta => [], name => FInstance, doc => null, platforms => {cross} }, { args => {{ name => c, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }, { name => cf, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassField,{})}), opt => false }}, meta => [], name => FStatic, doc => null, platforms => {cross} }, { args => {{ name => cf, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassField,{})}), opt => false }}, meta => [], name => FAnon, doc => null, platforms => {cross} }, { args => {{ name => s, t => CClass(String,{}), opt => false }}, meta => [], name => FDynamic, doc => null, platforms => {cross} }, { args => {{ name => c, t => CTypedef(Null,{CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})})}), opt => false }, { name => cf, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassField,{})}), opt => false }}, meta => [], name => FClosure, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.EnumType,{})}), opt => false }, { name => ef, t => CTypedef(haxe.macro.EnumField,{}), opt => false }}, meta => [], name => FEnum, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TypedExprDef uintptr /* Enumdecl - {{ args => {{ name => c, t => CEnum(haxe.macro.TConstant,{}), opt => false }}, meta => [], name => TConst, doc => null, platforms => {cross} }, { args => {{ name => v, t => CTypedef(haxe.macro.TVar,{}), opt => false }}, meta => [], name => TLocal, doc => null, platforms => {cross} }, { args => {{ name => e1, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => e2, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TArray, doc => null, platforms => {cross} }, { args => {{ name => op, t => CEnum(haxe.macro.Binop,{}), opt => false }, { name => e1, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => e2, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TBinop, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => fa, t => CEnum(haxe.macro.FieldAccess,{}), opt => false }}, meta => [], name => TField, doc => null, platforms => {cross} }, { args => {{ name => m, t => CEnum(haxe.macro.ModuleType,{}), opt => false }}, meta => [], name => TTypeExpr, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TParenthesis, doc => null, platforms => {cross} }, { args => {{ name => fields, t => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(haxe.macro.TypedExpr,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), opt => false }}, meta => [], name => TObjectDecl, doc => null, platforms => {cross} }, { args => {{ name => el, t => CClass(Array,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TArrayDecl, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => el, t => CClass(Array,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TCall, doc => null, platforms => {cross} }, { args => {{ name => c, t => CTypedef(haxe.macro.Ref,{CTypedef(haxe.macro.ClassType,{})}), opt => false }, { name => params, t => CClass(Array,{CEnum(haxe.macro.Type,{})}), opt => false }, { name => el, t => CClass(Array,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TNew, doc => null, platforms => {cross} }, { args => {{ name => op, t => CEnum(haxe.macro.Unop,{}), opt => false }, { name => postFix, t => CAbstract(Bool,{}), opt => false }, { name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TUnop, doc => null, platforms => {cross} }, { args => {{ name => tfunc, t => CTypedef(haxe.macro.TFunc,{}), opt => false }}, meta => [], name => TFunction, doc => null, platforms => {cross} }, { args => {{ name => v, t => CTypedef(haxe.macro.TVar,{}), opt => false }, { name => expr, t => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TVar, doc => null, platforms => {cross} }, { args => {{ name => el, t => CClass(Array,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TBlock, doc => null, platforms => {cross} }, { args => {{ name => v, t => CTypedef(haxe.macro.TVar,{}), opt => false }, { name => e1, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => e2, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TFor, doc => null, platforms => {cross} }, { args => {{ name => econd, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => eif, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => eelse, t => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TIf, doc => null, platforms => {cross} }, { args => {{ name => econd, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => normalWhile, t => CAbstract(Bool,{}), opt => false }}, meta => [], name => TWhile, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => cases, t => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => values, overloads => null, type => CClass(Array,{CTypedef(haxe.macro.TypedExpr,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(haxe.macro.TypedExpr,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), opt => false }, { name => edef, t => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TSwitch, doc => null, platforms => {cross} }, { args => null, meta => [], name => TPatMatch, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => catches, t => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => v, overloads => null, type => CTypedef(haxe.macro.TVar,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CTypedef(haxe.macro.TypedExpr,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), opt => false }}, meta => [], name => TTry, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(Null,{CTypedef(haxe.macro.TypedExpr,{})}), opt => false }}, meta => [], name => TReturn, doc => null, platforms => {cross} }, { args => null, meta => [], name => TBreak, doc => null, platforms => {cross} }, { args => null, meta => [], name => TContinue, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TThrow, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => m, t => CTypedef(Null,{CEnum(haxe.macro.ModuleType,{})}), opt => false }}, meta => [], name => TCast, doc => null, platforms => {cross} }, { args => {{ name => m, t => CTypedef(haxe.macro.MetadataEntry,{}), opt => false }, { name => e1, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }}, meta => [], name => TMeta, doc => null, platforms => {cross} }, { args => {{ name => e1, t => CTypedef(haxe.macro.TypedExpr,{}), opt => false }, { name => ef, t => CTypedef(haxe.macro.EnumField,{}), opt => false }, { name => index, t => CAbstract(Int,{}), opt => false }}, meta => [], name => TEnumParameter, doc => null, platforms => {cross} }} */
type Xhaxe_macro_TypedExpr uintptr    /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.macro.Type,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => pos, overloads => null, type => CTypedef(haxe.macro.Position,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => expr, overloads => null, type => CEnum(haxe.macro.TypedExprDef,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
func Xhaxe_macro_TypeTools_findField(Xc Xhaxe_macro_ClassType, Xname string, XisStatic bool)

type Xhaxe_macro_TypeTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_macro_TypedExprTools uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_remoting_AsyncConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_AsyncConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_AsyncConnection) Xcall(Xparams XArray, Xresult func())

func (x Xhaxe_remoting_AsyncConnection) XsetErrorHandler(Xerror func())

type Xhaxe_remoting_AMFConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_AMFConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_AMFConnection) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_AMFConnection) Xclose()

func (x Xhaxe_remoting_AMFConnection) Xcall(Xparams XArray, XonResult func())

func Xhaxe_remoting_AsyncAdapter_create(Xcnx Xhaxe_remoting_Connection)

type Xhaxe_remoting_AsyncAdapter uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_AsyncAdapter) Xresolve(Xname string)

func (x Xhaxe_remoting_AsyncAdapter) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_AsyncAdapter) Xcall(Xparams XArray, XonResult func())

func Xhaxe_remoting_AsyncDebugConnection_create(Xcnx Xhaxe_remoting_AsyncConnection)

type Xhaxe_remoting_AsyncDebugConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_AsyncDebugConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_AsyncDebugConnection) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_AsyncDebugConnection) XsetErrorDebug(Xh func())

func (x Xhaxe_remoting_AsyncDebugConnection) XsetResultDebug(Xh func())

func (x Xhaxe_remoting_AsyncDebugConnection) XsetCallDebug(Xh func())

func (x Xhaxe_remoting_AsyncDebugConnection) Xcall(Xparams XArray, XonResult func())

type Xhaxe_remoting_AsyncProxy uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_remoting_AsyncProxy_T uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_remoting_Connection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_Connection) Xresolve(Xname string)

func (x Xhaxe_remoting_Connection) Xcall(Xparams XArray)

func Xhaxe_remoting_Context_share(Xname string, Xobj struct{})

type Xhaxe_remoting_Context uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_Context) XaddObject(Xname string, Xobj struct{}, Xrecursive bool)

func (x Xhaxe_remoting_Context) Xcall(Xpath XArray, Xparams XArray)

func Xhaxe_remoting_Context_new() Xhaxe_remoting_Context

type Xhaxe_remoting_ContextAll Xhaxe_remoting_Context // TClassdecl

func (x Xhaxe_remoting_ContextAll) Xcall(Xpath XArray, Xparams XArray)

func Xhaxe_remoting_ContextAll_new() Xhaxe_remoting_ContextAll

func Xhaxe_remoting_DelayedConnection_create()

type Xhaxe_remoting_DelayedConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_DelayedConnection) Xconnection_goget()

func (x Xhaxe_remoting_DelayedConnection) Xconnection_goset(Xhaxe_remoting_AsyncConnection)

func (x Xhaxe_remoting_DelayedConnection) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_DelayedConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_DelayedConnection) Xcall(Xparams XArray, XonResult func())

type Xhaxe_remoting_ExternalConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_ExternalConnection) Xresolve(Xfield string)

func (x Xhaxe_remoting_ExternalConnection) Xclose()

func (x Xhaxe_remoting_ExternalConnection) Xcall(Xparams XArray)

type Xhaxe_remoting_FlashJsConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_remoting_HttpAsyncConnection_urlConnect(Xurl string)

type Xhaxe_remoting_HttpAsyncConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_HttpAsyncConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_HttpAsyncConnection) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_HttpAsyncConnection) Xcall(Xparams XArray, XonResult func())

func Xhaxe_remoting_HttpConnection_TIMEOUT_goget()

func Xhaxe_remoting_HttpConnection_TIMEOUT_goset(float64)

func Xhaxe_remoting_HttpConnection_processRequest(XrequestData string, Xctx Xhaxe_remoting_Context)

type Xhaxe_remoting_HttpConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_HttpConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_HttpConnection) Xcall(Xparams XArray)

type Xhaxe_remoting_LocalConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_LocalConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_LocalConnection) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_LocalConnection) Xcall(Xparams XArray, XonResult func())

func (x Xhaxe_remoting_LocalConnection) Xclose()

type Xhaxe_remoting_Proxy uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_remoting_Proxy_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_remoting_SocketConnection_create(Xs Xhaxe_remoting_Socket, Xctx Xhaxe_remoting_Context)

type Xhaxe_remoting_SocketConnection uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_SocketConnection) Xresolve(Xname string)

func (x Xhaxe_remoting_SocketConnection) Xcall(Xparams XArray, XonResult func())

func (x Xhaxe_remoting_SocketConnection) XsetErrorHandler(Xh func())

func (x Xhaxe_remoting_SocketConnection) XsetErrorLogger(Xh func())

func (x Xhaxe_remoting_SocketConnection) XsetProtocol(Xp Xhaxe_remoting_SocketProtocol)

func (x Xhaxe_remoting_SocketConnection) XgetProtocol()

func (x Xhaxe_remoting_SocketConnection) Xclose()

func (x Xhaxe_remoting_SocketConnection) XprocessMessage(Xdata string)

type Xhaxe_remoting_Socket uintptr         /* TTypedecl -  CClass(sys.net.Socket,{}) */
type Xhaxe_remoting_SocketProtocol uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_remoting_SocketProtocol) Xsocket_goget()

func (x Xhaxe_remoting_SocketProtocol) Xsocket_goset(Xhaxe_remoting_Socket)

func (x Xhaxe_remoting_SocketProtocol) Xcontext_goget()

func (x Xhaxe_remoting_SocketProtocol) Xcontext_goset(Xhaxe_remoting_Context)

func (x Xhaxe_remoting_SocketProtocol) XmessageLength(Xc1 int, Xc2 int)

func (x Xhaxe_remoting_SocketProtocol) XencodeMessageLength(Xlen int)

func (x Xhaxe_remoting_SocketProtocol) XsendRequest(Xpath XArray, Xparams XArray)

func (x Xhaxe_remoting_SocketProtocol) XsendAnswer(Xanswer uintptr, XisException bool)

func (x Xhaxe_remoting_SocketProtocol) XsendMessage(Xmsg string)

func (x Xhaxe_remoting_SocketProtocol) XdecodeData(Xdata string)

func (x Xhaxe_remoting_SocketProtocol) XisRequest(Xdata string)

func (x Xhaxe_remoting_SocketProtocol) XprocessRequest(Xdata string, XonError func())

func (x Xhaxe_remoting_SocketProtocol) XprocessAnswer(Xdata string)

func (x Xhaxe_remoting_SocketProtocol) XreadMessage()

func Xhaxe_remoting_SocketProtocol_new(Xsock Xhaxe_remoting_Socket, Xctx Xhaxe_remoting_Context) Xhaxe_remoting_SocketProtocol

type Xhaxe_rtti_Path uintptr             /* TTypedecl -  CClass(String,{}) */
type Xhaxe_rtti_Platforms uintptr        /* TTypedecl -  CClass(List,{CClass(String,{})}) */
type Xhaxe_rtti_FunctionArgument uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => value, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => opt, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_CType uintptr            /* Enumdecl - {{ args => null, meta => [], name => CUnknown, doc => null, platforms => {cross} }, { args => {{ name => name, t => CTypedef(haxe.rtti.Path,{}), opt => false }, { name => params, t => CClass(List,{CEnum(haxe.rtti.CType,{})}), opt => false }}, meta => [], name => CEnum, doc => null, platforms => {cross} }, { args => {{ name => name, t => CTypedef(haxe.rtti.Path,{}), opt => false }, { name => params, t => CClass(List,{CEnum(haxe.rtti.CType,{})}), opt => false }}, meta => [], name => CClass, doc => null, platforms => {cross} }, { args => {{ name => name, t => CTypedef(haxe.rtti.Path,{}), opt => false }, { name => params, t => CClass(List,{CEnum(haxe.rtti.CType,{})}), opt => false }}, meta => [], name => CTypedef, doc => null, platforms => {cross} }, { args => {{ name => args, t => CClass(List,{CTypedef(haxe.rtti.FunctionArgument,{})}), opt => false }, { name => ret, t => CEnum(haxe.rtti.CType,{}), opt => false }}, meta => [], name => CFunction, doc => null, platforms => {cross} }, { args => {{ name => fields, t => CClass(List,{CTypedef(haxe.rtti.ClassField,{})}), opt => false }}, meta => [], name => CAnonymous, doc => null, platforms => {cross} }, { args => {{ name => t, t => CEnum(haxe.rtti.CType,{}), opt => true }}, meta => [], name => CDynamic, doc => null, platforms => {cross} }, { args => {{ name => name, t => CTypedef(haxe.rtti.Path,{}), opt => false }, { name => params, t => CClass(List,{CEnum(haxe.rtti.CType,{})}), opt => false }}, meta => [], name => CAbstract, doc => null, platforms => {cross} }} */
type Xhaxe_rtti_PathParams uintptr       /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => path, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(List,{CEnum(haxe.rtti.CType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_TypeParams uintptr       /* TTypedecl -  CClass(Array,{CClass(String,{})}) */
type Xhaxe_rtti_Rights uintptr           /* Enumdecl - {{ args => null, meta => [], name => RNormal, doc => null, platforms => {cross} }, { args => null, meta => [], name => RNo, doc => null, platforms => {cross} }, { args => {{ name => m, t => CClass(String,{}), opt => false }}, meta => [], name => RCall, doc => null, platforms => {cross} }, { args => null, meta => [], name => RMethod, doc => null, platforms => {cross} }, { args => null, meta => [], name => RDynamic, doc => null, platforms => {cross} }, { args => null, meta => [], name => RInline, doc => null, platforms => {cross} }} */
type Xhaxe_rtti_MetaData uintptr         /* TTypedecl -  CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => params, overloads => null, type => CClass(Array,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}) */
type Xhaxe_rtti_ClassField uintptr       /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => set, overloads => null, type => CEnum(haxe.rtti.Rights,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CTypedef(haxe.rtti.TypeParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => overloads, overloads => null, type => CTypedef(Null,{CClass(List,{CTypedef(haxe.rtti.ClassField,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => line, overloads => null, type => CTypedef(Null,{CAbstract(Int,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPublic, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isOverride, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => get, overloads => null, type => CEnum(haxe.rtti.Rights,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_TypeInfos uintptr        /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => path, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CTypedef(haxe.rtti.TypeParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => file, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_Classdef uintptr         /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => tdynamic, overloads => null, type => CTypedef(Null,{CEnum(haxe.rtti.CType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => superClass, overloads => null, type => CTypedef(haxe.rtti.PathParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => statics, overloads => null, type => CClass(List,{CTypedef(haxe.rtti.ClassField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => path, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CTypedef(haxe.rtti.TypeParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isInterface, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => interfaces, overloads => null, type => CClass(List,{CTypedef(haxe.rtti.PathParams,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => file, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fields, overloads => null, type => CClass(List,{CTypedef(haxe.rtti.ClassField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_EnumField uintptr        /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => args, overloads => null, type => CTypedef(Null,{CClass(List,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => opt, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_Enumdef uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => path, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CTypedef(haxe.rtti.TypeParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isExtern, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => file, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => constructors, overloads => null, type => CClass(List,{CTypedef(haxe.rtti.EnumField,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_Typedef uintptr          /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => types, overloads => null, type => CClass(haxe.ds.StringMap,{CEnum(haxe.rtti.CType,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => type, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => path, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CTypedef(haxe.rtti.TypeParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => file, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_Abstractdef uintptr      /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => to, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => field, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => platforms, overloads => null, type => CTypedef(haxe.rtti.Platforms,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => path, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => params, overloads => null, type => CTypedef(haxe.rtti.TypeParams,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => module, overloads => null, type => CTypedef(haxe.rtti.Path,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => meta, overloads => null, type => CTypedef(haxe.rtti.MetaData,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => isPrivate, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => impl, overloads => null, type => CTypedef(haxe.rtti.Classdef,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => from, overloads => null, type => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => t, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => field, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => file, overloads => null, type => CTypedef(Null,{CClass(String,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => doc, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => athis, overloads => null, type => CEnum(haxe.rtti.CType,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_rtti_TypeTree uintptr         /* Enumdecl - {{ args => {{ name => name, t => CClass(String,{}), opt => false }, { name => full, t => CClass(String,{}), opt => false }, { name => subs, t => CClass(Array,{CEnum(haxe.rtti.TypeTree,{})}), opt => false }}, meta => [], name => TPackage, doc => null, platforms => {cross} }, { args => {{ name => c, t => CTypedef(haxe.rtti.Classdef,{}), opt => false }}, meta => [], name => TClassdecl, doc => null, platforms => {cross} }, { args => {{ name => e, t => CTypedef(haxe.rtti.Enumdef,{}), opt => false }}, meta => [], name => TEnumdecl, doc => null, platforms => {cross} }, { args => {{ name => t, t => CTypedef(haxe.rtti.Typedef,{}), opt => false }}, meta => [], name => TTypedecl, doc => null, platforms => {cross} }, { args => {{ name => a, t => CTypedef(haxe.rtti.Abstractdef,{}), opt => false }}, meta => [], name => TAbstractdecl, doc => null, platforms => {cross} }} */
type Xhaxe_rtti_TypeRoot uintptr         /* TTypedecl -  CClass(Array,{CEnum(haxe.rtti.TypeTree,{})}) */
func Xhaxe_rtti_TypeApi_typeInfos(Xt Xhaxe_rtti_TypeTree)

func Xhaxe_rtti_TypeApi_isVar(Xt Xhaxe_rtti_CType)

func Xhaxe_rtti_TypeApi_rightsEq(Xr1 Xhaxe_rtti_Rights, Xr2 Xhaxe_rtti_Rights)

func Xhaxe_rtti_TypeApi_typeEq(Xt1 Xhaxe_rtti_CType, Xt2 Xhaxe_rtti_CType)

func Xhaxe_rtti_TypeApi_fieldEq(Xf1 Xhaxe_rtti_ClassField, Xf2 Xhaxe_rtti_ClassField)

func Xhaxe_rtti_TypeApi_constructorEq(Xc1 Xhaxe_rtti_EnumField, Xc2 Xhaxe_rtti_EnumField)

type Xhaxe_rtti_TypeApi uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_rtti_Meta_getType(Xt uintptr)

func Xhaxe_rtti_Meta_getStatics(Xt uintptr)

func Xhaxe_rtti_Meta_getFields(Xt uintptr)

type Xhaxe_rtti_Meta uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_rtti_XmlParser uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_rtti_XmlParser) Xroot_goget()

func (x Xhaxe_rtti_XmlParser) Xroot_goset(Xhaxe_rtti_TypeRoot)

func (x Xhaxe_rtti_XmlParser) Xsort(Xl XArray)

func (x Xhaxe_rtti_XmlParser) Xprocess(Xx XXml, Xplatform string)

func (x Xhaxe_rtti_XmlParser) XnewField(Xc Xhaxe_rtti_Classdef, Xf Xhaxe_rtti_ClassField)

func (x Xhaxe_rtti_XmlParser) XprocessElement(Xx XXml)

func Xhaxe_rtti_XmlParser_new() Xhaxe_rtti_XmlParser

type Xhaxe_unit_TestCase uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_unit_TestCase) XcurrentTest_goget()

func (x Xhaxe_unit_TestCase) XcurrentTest_goset(Xhaxe_unit_TestStatus)

func (x Xhaxe_unit_TestCase) Xsetup()

func (x Xhaxe_unit_TestCase) XtearDown()

func (x Xhaxe_unit_TestCase) Xprint(Xv uintptr)

func (x Xhaxe_unit_TestCase) XassertTrue(Xb bool, Xc Xhaxe_PosInfos)

func (x Xhaxe_unit_TestCase) XassertFalse(Xb bool, Xc Xhaxe_PosInfos)

func (x Xhaxe_unit_TestCase) XassertEquals(Xexpected uintptr, Xactual uintptr, Xc Xhaxe_PosInfos)

func Xhaxe_unit_TestCase_new() Xhaxe_unit_TestCase

type Xhaxe_unit_TestResult uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_unit_TestResult) Xsuccess_goget()

func (x Xhaxe_unit_TestResult) Xsuccess_goset(bool)

func (x Xhaxe_unit_TestResult) Xadd(Xt Xhaxe_unit_TestStatus)

func (x Xhaxe_unit_TestResult) XtoString()

func Xhaxe_unit_TestResult_new() Xhaxe_unit_TestResult

func Xhaxe_unit_TestRunner_print(Xv uintptr)

type Xhaxe_unit_TestRunner uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_unit_TestRunner) Xresult_goget()

func (x Xhaxe_unit_TestRunner) Xresult_goset(Xhaxe_unit_TestResult)

func (x Xhaxe_unit_TestRunner) Xadd(Xc Xhaxe_unit_TestCase)

func (x Xhaxe_unit_TestRunner) Xrun()

func Xhaxe_unit_TestRunner_new() Xhaxe_unit_TestRunner

type Xhaxe_unit_TestStatus uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_unit_TestStatus) Xdone_goget()

func (x Xhaxe_unit_TestStatus) Xdone_goset(bool)

func (x Xhaxe_unit_TestStatus) Xsuccess_goget()

func (x Xhaxe_unit_TestStatus) Xsuccess_goset(bool)

func (x Xhaxe_unit_TestStatus) Xerror_goget()

func (x Xhaxe_unit_TestStatus) Xerror_goset(string)

func (x Xhaxe_unit_TestStatus) Xmethod_goget()

func (x Xhaxe_unit_TestStatus) Xmethod_goset(string)

func (x Xhaxe_unit_TestStatus) Xclassname_goget()

func (x Xhaxe_unit_TestStatus) Xclassname_goset(string)

func (x Xhaxe_unit_TestStatus) XposInfos_goget()

func (x Xhaxe_unit_TestStatus) XposInfos_goset(Xhaxe_PosInfos)

func (x Xhaxe_unit_TestStatus) Xbacktrace_goget()

func (x Xhaxe_unit_TestStatus) Xbacktrace_goset(string)

func Xhaxe_unit_TestStatus_new() Xhaxe_unit_TestStatus

type Xhaxe_web_DispatchConfig uintptr /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => rules, overloads => null, type => CDynamic(null), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => obj, overloads => null, type => CDynamic(null), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_web_Lock uintptr           /* TTypedecl -  CClass(haxe.web.Lock.T,{}) */
type Xhaxe_web_MatchRule uintptr      /* Enumdecl - {{ args => null, meta => [], name => MRInt, doc => null, platforms => {cross} }, { args => null, meta => [], name => MRBool, doc => null, platforms => {cross} }, { args => null, meta => [], name => MRFloat, doc => null, platforms => {cross} }, { args => null, meta => [], name => MRString, doc => null, platforms => {cross} }, { args => {{ name => e, t => CClass(String,{}), opt => false }}, meta => [], name => MREnum, doc => null, platforms => {cross} }, { args => null, meta => [], name => MRDispatch, doc => null, platforms => {cross} }, { args => {{ name => c, t => CClass(String,{}), opt => false }, { name => lock, t => CAbstract(Bool,{}), opt => false }}, meta => [], name => MRSpod, doc => null, platforms => {cross} }, { args => {{ name => r, t => CEnum(haxe.web.MatchRule,{}), opt => false }}, meta => [], name => MROpt, doc => null, platforms => {cross} }} */
type Xhaxe_web_DispatchRule uintptr   /* Enumdecl - {{ args => {{ name => r, t => CEnum(haxe.web.MatchRule,{}), opt => false }}, meta => [], name => DRMatch, doc => null, platforms => {cross} }, { args => {{ name => r, t => CClass(Array,{CEnum(haxe.web.MatchRule,{})}), opt => false }}, meta => [], name => DRMult, doc => null, platforms => {cross} }, { args => {{ name => r, t => CEnum(haxe.web.DispatchRule,{}), opt => false }, { name => args, t => CClass(Array,{CAnonymous({{ isOverride => false, line => null, meta => [], name => rule, overloads => null, type => CEnum(haxe.web.MatchRule,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => opt, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => name, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }})}), opt => false }, { name => opt, t => CAbstract(Bool,{}), opt => false }}, meta => [], name => DRArgs, doc => null, platforms => {cross} }, { args => {{ name => r, t => CEnum(haxe.web.DispatchRule,{}), opt => false }}, meta => [], name => DRMeta, doc => null, platforms => {cross} }} */
type Xhaxe_web_DispatchError uintptr  /* Enumdecl - {{ args => {{ name => part, t => CClass(String,{}), opt => false }}, meta => [], name => DENotFound, doc => null, platforms => {cross} }, { args => null, meta => [], name => DEInvalidValue, doc => null, platforms => {cross} }, { args => null, meta => [], name => DEMissing, doc => null, platforms => {cross} }, { args => {{ name => p, t => CClass(String,{}), opt => false }}, meta => [], name => DEMissingParam, doc => null, platforms => {cross} }, { args => null, meta => [], name => DETooManyValues, doc => null, platforms => {cross} }} */
type Xhaxe_web_Redirect uintptr       /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_web_Redirect_new() Xhaxe_web_Redirect

func Xhaxe_web_Dispatch_make(Xobj struct{})

func Xhaxe_web_Dispatch_run(Xurl string, Xparams Xhaxe_ds_StringMap, Xobj struct{})

type Xhaxe_web_Dispatch uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_web_Dispatch) Xparts_goget()

func (x Xhaxe_web_Dispatch) Xparts_goset(XArray)

func (x Xhaxe_web_Dispatch) Xparams_goget()

func (x Xhaxe_web_Dispatch) Xparams_goset(Xhaxe_ds_StringMap)

func (x Xhaxe_web_Dispatch) Xname_goget()

func (x Xhaxe_web_Dispatch) Xname_goset(string)

func (x Xhaxe_web_Dispatch) Xcfg_goget()

func (x Xhaxe_web_Dispatch) Xcfg_goset(Xhaxe_web_DispatchConfig)

func (x Xhaxe_web_Dispatch) Xdispatch(Xobj struct{})

func (x Xhaxe_web_Dispatch) XgetParams()

func (x Xhaxe_web_Dispatch) XonMeta(Xv string, Xargs uintptr)

func (x Xhaxe_web_Dispatch) XruntimeDispatch(Xcfg Xhaxe_web_DispatchConfig)

func (x Xhaxe_web_Dispatch) Xredirect(Xurl string, Xparams Xhaxe_ds_StringMap)

func (x Xhaxe_web_Dispatch) XruntimeGetParams(XcfgIndex int)

func Xhaxe_web_Dispatch_new(Xurl string, Xparams Xhaxe_ds_StringMap) Xhaxe_web_Dispatch

type Xhaxe_xml_Filter uintptr               /* Enumdecl - {{ args => null, meta => [], name => FInt, doc => null, platforms => {cross} }, { args => null, meta => [], name => FBool, doc => null, platforms => {cross} }, { args => {{ name => values, t => CClass(Array,{CClass(String,{})}), opt => false }}, meta => [], name => FEnum, doc => null, platforms => {cross} }, { args => {{ name => matcher, t => CClass(EReg,{}), opt => false }}, meta => [], name => FReg, doc => null, platforms => {cross} }} */
type Xhaxe_xml_Attrib uintptr               /* Enumdecl - {{ args => {{ name => name, t => CClass(String,{}), opt => false }, { name => filter, t => CEnum(haxe.xml.Filter,{}), opt => true }, { name => defvalue, t => CClass(String,{}), opt => true }}, meta => [], name => Att, doc => null, platforms => {cross} }} */
type Xhaxe_xml_Rule uintptr                 /* Enumdecl - {{ args => {{ name => name, t => CClass(String,{}), opt => false }, { name => attribs, t => CClass(Array,{CEnum(haxe.xml.Attrib,{})}), opt => true }, { name => childs, t => CEnum(haxe.xml.Rule,{}), opt => true }}, meta => [], name => RNode, doc => null, platforms => {cross} }, { args => {{ name => filter, t => CEnum(haxe.xml.Filter,{}), opt => true }}, meta => [], name => RData, doc => null, platforms => {cross} }, { args => {{ name => rule, t => CEnum(haxe.xml.Rule,{}), opt => false }, { name => atLeastOne, t => CAbstract(Bool,{}), opt => true }}, meta => [], name => RMulti, doc => null, platforms => {cross} }, { args => {{ name => rules, t => CClass(Array,{CEnum(haxe.xml.Rule,{})}), opt => false }, { name => ordered, t => CAbstract(Bool,{}), opt => true }}, meta => [], name => RList, doc => null, platforms => {cross} }, { args => {{ name => choices, t => CClass(Array,{CEnum(haxe.xml.Rule,{})}), opt => false }}, meta => [], name => RChoice, doc => null, platforms => {cross} }, { args => {{ name => rule, t => CEnum(haxe.xml.Rule,{}), opt => false }}, meta => [], name => ROptional, doc => null, platforms => {cross} }} */
type Xhaxe_xml____Check_CheckResult uintptr /* Enumdecl - {{ args => null, meta => [], name => CMatch, doc => null, platforms => {cross} }, { args => {{ name => r, t => CEnum(haxe.xml.Rule,{}), opt => false }}, meta => [], name => CMissing, doc => null, platforms => {cross} }, { args => {{ name => x, t => CClass(Xml,{}), opt => false }}, meta => [], name => CExtra, doc => null, platforms => {cross} }, { args => {{ name => name, t => CClass(String,{}), opt => false }, { name => x, t => CClass(Xml,{}), opt => false }}, meta => [], name => CElementExpected, doc => null, platforms => {cross} }, { args => {{ name => x, t => CClass(Xml,{}), opt => false }}, meta => [], name => CDataExpected, doc => null, platforms => {cross} }, { args => {{ name => att, t => CClass(String,{}), opt => false }, { name => x, t => CClass(Xml,{}), opt => false }}, meta => [], name => CExtraAttrib, doc => null, platforms => {cross} }, { args => {{ name => att, t => CClass(String,{}), opt => false }, { name => x, t => CClass(Xml,{}), opt => false }}, meta => [], name => CMissingAttrib, doc => null, platforms => {cross} }, { args => {{ name => att, t => CClass(String,{}), opt => false }, { name => x, t => CClass(Xml,{}), opt => false }, { name => f, t => CEnum(haxe.xml.Filter,{}), opt => false }}, meta => [], name => CInvalidAttrib, doc => null, platforms => {cross} }, { args => {{ name => x, t => CClass(Xml,{}), opt => false }, { name => f, t => CEnum(haxe.xml.Filter,{}), opt => false }}, meta => [], name => CInvalidData, doc => null, platforms => {cross} }, { args => {{ name => x, t => CClass(Xml,{}), opt => false }, { name => r, t => CEnum(haxe.xml._Check.CheckResult,{}), opt => false }}, meta => [], name => CInElement, doc => null, platforms => {cross} }} */
func Xhaxe_xml_Check_checkNode(Xx XXml, Xr Xhaxe_xml_Rule)

func Xhaxe_xml_Check_checkDocument(Xx XXml, Xr Xhaxe_xml_Rule)

type Xhaxe_xml_Check uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_xml_Fast uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_xml_Fast) Xx_goget()

func (x Xhaxe_xml_Fast) Xx_goset(XXml)

func (x Xhaxe_xml_Fast) Xname_goget()

func (x Xhaxe_xml_Fast) Xname_goset(string)

func (x Xhaxe_xml_Fast) XinnerData_goget()

func (x Xhaxe_xml_Fast) XinnerData_goset(string)

func (x Xhaxe_xml_Fast) XinnerHTML_goget()

func (x Xhaxe_xml_Fast) XinnerHTML_goset(string)

func (x Xhaxe_xml_Fast) Xnode_goget()

func (x Xhaxe_xml_Fast) Xnode_goset(uintptr)

func (x Xhaxe_xml_Fast) Xnodes_goget()

func (x Xhaxe_xml_Fast) Xnodes_goset(uintptr)

func (x Xhaxe_xml_Fast) Xatt_goget()

func (x Xhaxe_xml_Fast) Xatt_goset(uintptr)

func (x Xhaxe_xml_Fast) Xhas_goget()

func (x Xhaxe_xml_Fast) Xhas_goset(uintptr)

func (x Xhaxe_xml_Fast) XhasNode_goget()

func (x Xhaxe_xml_Fast) XhasNode_goset(uintptr)

func (x Xhaxe_xml_Fast) Xelements_goget()

func (x Xhaxe_xml_Fast) Xelements_goset(XIterator)

func Xhaxe_xml_Fast_new(Xx XXml) Xhaxe_xml_Fast

func Xhaxe_xml_Parser_parse(Xstr string)

type Xhaxe_xml_Parser uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_xml_Printer_print(Xxml XXml)

type Xhaxe_xml_Printer uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_xml_Proxy uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_xml_Proxy_Const uintptr /* should be: hx.Dynamic */ // TClassdecl

type Xhaxe_xml_Proxy_T uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_xml_Proxy) Xresolve(Xk string)

func Xhaxe_xml_Proxy_new(Xf func()) Xhaxe_xml_Proxy

func Xhaxe_zip_Compress_run(Xs Xhaxe_io_Bytes, Xlevel int)

type Xhaxe_zip_Compress uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_zip_Compress) Xexecute(Xsrc Xhaxe_io_Bytes, XsrcPos int, Xdst Xhaxe_io_Bytes, XdstPos int)

func (x Xhaxe_zip_Compress) XsetFlushMode(Xf Xhaxe_zip_FlushMode)

func (x Xhaxe_zip_Compress) Xclose()

func Xhaxe_zip_Compress_new(Xlevel int) Xhaxe_zip_Compress

type Xhaxe_zip_ExtraField uintptr /* Enumdecl - {{ args => {{ name => tag, t => CAbstract(Int,{}), opt => false }, { name => bytes, t => CClass(haxe.io.Bytes,{}), opt => false }}, meta => [], name => FUnknown, doc => null, platforms => {cross} }, { args => {{ name => name, t => CClass(String,{}), opt => false }, { name => crc, t => CAbstract(Int,{}), opt => false }}, meta => [], name => FInfoZipUnicodePath, doc => null, platforms => {cross} }, { args => null, meta => [], name => FUtf8, doc => null, platforms => {cross} }} */
type Xhaxe_zip_Entry uintptr      /* TTypedecl -  CAnonymous({{ isOverride => false, line => null, meta => [], name => fileTime, overloads => null, type => CClass(Date,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fileSize, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => fileName, overloads => null, type => CClass(String,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [{ name => :optional, params => [] }], name => extraFields, overloads => null, type => CTypedef(Null,{CClass(List,{CEnum(haxe.zip.ExtraField,{})})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => dataSize, overloads => null, type => CAbstract(Int,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => data, overloads => null, type => CTypedef(Null,{CClass(haxe.io.Bytes,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => crc32, overloads => null, type => CTypedef(Null,{CAbstract(Int,{})}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }, { isOverride => false, line => null, meta => [], name => compressed, overloads => null, type => CAbstract(Bool,{}), params => [], doc => null, get => RNormal, set => RNormal, platforms => {}, isPublic => true }}) */
type Xhaxe_zip_FlushMode uintptr  /* Enumdecl - {{ args => null, meta => [], name => NO, doc => null, platforms => {cross} }, { args => null, meta => [], name => SYNC, doc => null, platforms => {cross} }, { args => null, meta => [], name => FULL, doc => null, platforms => {cross} }, { args => null, meta => [], name => FINISH, doc => null, platforms => {cross} }, { args => null, meta => [], name => BLOCK, doc => null, platforms => {cross} }} */
type Xhaxe_zip_Huffman uintptr    /* Enumdecl - {{ args => {{ name => i, t => CAbstract(Int,{}), opt => false }}, meta => [], name => Found, doc => null, platforms => {cross} }, { args => {{ name => left, t => CEnum(haxe.zip.Huffman,{}), opt => false }, { name => right, t => CEnum(haxe.zip.Huffman,{}), opt => false }}, meta => [], name => NeedBit, doc => null, platforms => {cross} }, { args => {{ name => n, t => CAbstract(Int,{}), opt => false }, { name => table, t => CClass(Array,{CEnum(haxe.zip.Huffman,{})}), opt => false }}, meta => [], name => NeedBits, doc => null, platforms => {cross} }} */
type Xhaxe_zip_HuffTools uintptr  /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_zip_HuffTools) Xmake(Xlengths XArray, Xpos int, Xnlengths int, Xmaxbits int)

func Xhaxe_zip_HuffTools_new() Xhaxe_zip_HuffTools

type Xhaxe_zip____InflateImpl_State uintptr /* Enumdecl - {{ args => null, meta => [], name => Head, doc => null, platforms => {cross} }, { args => null, meta => [], name => Block, doc => null, platforms => {cross} }, { args => null, meta => [], name => CData, doc => null, platforms => {cross} }, { args => null, meta => [], name => Flat, doc => null, platforms => {cross} }, { args => null, meta => [], name => Crc, doc => null, platforms => {cross} }, { args => null, meta => [], name => Dist, doc => null, platforms => {cross} }, { args => null, meta => [], name => DistOne, doc => null, platforms => {cross} }, { args => null, meta => [], name => Done, doc => null, platforms => {cross} }} */
func Xhaxe_zip_InflateImpl_run(Xi Xhaxe_io_Input, Xbufsize int)

type Xhaxe_zip_InflateImpl uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_zip_InflateImpl) XreadBytes(Xb Xhaxe_io_Bytes, Xpos int, Xlen int)

func Xhaxe_zip_InflateImpl_new(Xi Xhaxe_io_Input, Xheader bool, Xcrc bool) Xhaxe_zip_InflateImpl

func Xhaxe_zip_Reader_readZip(Xi Xhaxe_io_Input)

func Xhaxe_zip_Reader_unzip(Xf Xhaxe_zip_Entry)

type Xhaxe_zip_Reader uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_zip_Reader) XreadEntryHeader()

func (x Xhaxe_zip_Reader) Xread()

func Xhaxe_zip_Reader_new(Xi Xhaxe_io_Input) Xhaxe_zip_Reader

func Xhaxe_zip_Tools_compress(Xf Xhaxe_zip_Entry, Xlevel int)

type Xhaxe_zip_Tools uintptr /* should be: hx.Dynamic */ // TClassdecl

func Xhaxe_zip_Uncompress_run(Xsrc Xhaxe_io_Bytes, Xbufsize int)

type Xhaxe_zip_Uncompress uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_zip_Uncompress) Xexecute(Xsrc Xhaxe_io_Bytes, XsrcPos int, Xdst Xhaxe_io_Bytes, XdstPos int)

func (x Xhaxe_zip_Uncompress) XsetFlushMode(Xf Xhaxe_zip_FlushMode)

func (x Xhaxe_zip_Uncompress) Xclose()

func Xhaxe_zip_Uncompress_new(XwindowBits int) Xhaxe_zip_Uncompress

type Xhaxe_zip_Writer uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xhaxe_zip_Writer) XwriteEntryHeader(Xf Xhaxe_zip_Entry)

func (x Xhaxe_zip_Writer) Xwrite(Xfiles XList)

func (x Xhaxe_zip_Writer) XwriteCDR()

func Xhaxe_zip_Writer_new(Xo Xhaxe_io_Output) Xhaxe_zip_Writer

func Xsys_net_Host_localhost()

type Xsys_net_Host uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xsys_net_Host) Xip_goget()

func (x Xsys_net_Host) Xip_goset(int)

func (x Xsys_net_Host) XtoString()

func (x Xsys_net_Host) Xreverse()

func Xsys_net_Host_new(Xname string) Xsys_net_Host

func Xsys_net_Socket_select(Xread XArray, Xwrite XArray, Xothers XArray, Xtimeout float64)

type Xsys_net_Socket uintptr /* should be: hx.Dynamic */ // TClassdecl

func (x Xsys_net_Socket) Xinput_goget()

func (x Xsys_net_Socket) Xinput_goset(Xhaxe_io_Input)

func (x Xsys_net_Socket) Xoutput_goget()

func (x Xsys_net_Socket) Xoutput_goset(Xhaxe_io_Output)

func (x Xsys_net_Socket) Xcustom_goget()

func (x Xsys_net_Socket) Xcustom_goset(uintptr)

func (x Xsys_net_Socket) Xclose()

func (x Xsys_net_Socket) Xread()

func (x Xsys_net_Socket) Xwrite(Xcontent string)

func (x Xsys_net_Socket) Xconnect(Xhost Xsys_net_Host, Xport int)

func (x Xsys_net_Socket) Xlisten(Xconnections int)

func (x Xsys_net_Socket) Xshutdown(Xread bool, Xwrite bool)

func (x Xsys_net_Socket) Xbind(Xhost Xsys_net_Host, Xport int)

func (x Xsys_net_Socket) Xaccept()

func (x Xsys_net_Socket) Xpeer()

func (x Xsys_net_Socket) Xhost()

func (x Xsys_net_Socket) XsetTimeout(Xtimeout float64)

func (x Xsys_net_Socket) XwaitForRead()

func (x Xsys_net_Socket) XsetBlocking(Xb bool)

func (x Xsys_net_Socket) XsetFastSend(Xb bool)

func Xsys_net_Socket_new() Xsys_net_Socket
