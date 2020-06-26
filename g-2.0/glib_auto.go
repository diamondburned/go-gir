package g

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
extern void myGLibChildWatchFunc(gint32 pid, gint32 status, gpointer user_data);
static void* getPointer_myGLibChildWatchFunc() {
return (void*)(myGLibChildWatchFunc);
}
extern void myGLibClearHandleFunc(guint32 handle_id);
static void* getPointer_myGLibClearHandleFunc() {
return (void*)(myGLibClearHandleFunc);
}
extern void myGLibCompareDataFunc(gpointer a, gpointer b, gpointer user_data);
static void* getPointer_myGLibCompareDataFunc() {
return (void*)(myGLibCompareDataFunc);
}
extern void myGLibCompareFunc(gpointer a, gpointer b);
static void* getPointer_myGLibCompareFunc() {
return (void*)(myGLibCompareFunc);
}
extern void myGLibCopyFunc(gpointer src, gpointer data);
static void* getPointer_myGLibCopyFunc() {
return (void*)(myGLibCopyFunc);
}
extern void myGLibDataForeachFunc(guint32 key_id, gpointer data, gpointer user_data);
static void* getPointer_myGLibDataForeachFunc() {
return (void*)(myGLibDataForeachFunc);
}
extern void myGLibDestroyNotify(gpointer data);
static void* getPointer_myGLibDestroyNotify() {
return (void*)(myGLibDestroyNotify);
}
extern void myGLibDuplicateFunc(gpointer data, gpointer user_data);
static void* getPointer_myGLibDuplicateFunc() {
return (void*)(myGLibDuplicateFunc);
}
extern void myGLibEqualFunc(gpointer a, gpointer b);
static void* getPointer_myGLibEqualFunc() {
return (void*)(myGLibEqualFunc);
}
extern void myGLibFreeFunc(gpointer data);
static void* getPointer_myGLibFreeFunc() {
return (void*)(myGLibFreeFunc);
}
extern void myGLibFunc(gpointer data, gpointer user_data);
static void* getPointer_myGLibFunc() {
return (void*)(myGLibFunc);
}
extern void myGLibHFunc(gpointer key, gpointer value, gpointer user_data);
static void* getPointer_myGLibHFunc() {
return (void*)(myGLibHFunc);
}
extern void myGLibHRFunc(gpointer key, gpointer value, gpointer user_data);
static void* getPointer_myGLibHRFunc() {
return (void*)(myGLibHRFunc);
}
extern void myGLibHashFunc(gpointer key);
static void* getPointer_myGLibHashFunc() {
return (void*)(myGLibHashFunc);
}
extern void myGLibHookCheckFunc(gpointer data);
static void* getPointer_myGLibHookCheckFunc() {
return (void*)(myGLibHookCheckFunc);
}
extern void myGLibHookCheckMarshaller(GHook* hook, gpointer marshal_data);
static void* getPointer_myGLibHookCheckMarshaller() {
return (void*)(myGLibHookCheckMarshaller);
}
extern void myGLibHookCompareFunc(GHook* new_hook, GHook* sibling);
static void* getPointer_myGLibHookCompareFunc() {
return (void*)(myGLibHookCompareFunc);
}
extern void myGLibHookFinalizeFunc(GHookList* hook_list, GHook* hook);
static void* getPointer_myGLibHookFinalizeFunc() {
return (void*)(myGLibHookFinalizeFunc);
}
extern void myGLibHookFindFunc(GHook* hook, gpointer data);
static void* getPointer_myGLibHookFindFunc() {
return (void*)(myGLibHookFindFunc);
}
extern void myGLibHookFunc(gpointer data);
static void* getPointer_myGLibHookFunc() {
return (void*)(myGLibHookFunc);
}
extern void myGLibHookMarshaller(GHook* hook, gpointer marshal_data);
static void* getPointer_myGLibHookMarshaller() {
return (void*)(myGLibHookMarshaller);
}
extern void myGLibIOFunc(GIOChannel* source, GIOCondition condition, gpointer data);
static void* getPointer_myGLibIOFunc() {
return (void*)(myGLibIOFunc);
}
extern void myGLibLogFunc(gchar* log_domain, GLogLevelFlags log_level, gchar* message, gpointer user_data);
static void* getPointer_myGLibLogFunc() {
return (void*)(myGLibLogFunc);
}
extern void myGLibLogWriterFunc(GLogLevelFlags log_level, gpointer fields, guint64 n_fields, gpointer user_data);
static void* getPointer_myGLibLogWriterFunc() {
return (void*)(myGLibLogWriterFunc);
}
extern void myGLibNodeForeachFunc(GNode* node, gpointer data);
static void* getPointer_myGLibNodeForeachFunc() {
return (void*)(myGLibNodeForeachFunc);
}
extern void myGLibNodeTraverseFunc(GNode* node, gpointer data);
static void* getPointer_myGLibNodeTraverseFunc() {
return (void*)(myGLibNodeTraverseFunc);
}
extern void myGLibOptionArgFunc(gchar* option_name, gchar* value, gpointer data);
static void* getPointer_myGLibOptionArgFunc() {
return (void*)(myGLibOptionArgFunc);
}
extern void myGLibOptionErrorFunc(GOptionContext* context, GOptionGroup* group, gpointer data);
static void* getPointer_myGLibOptionErrorFunc() {
return (void*)(myGLibOptionErrorFunc);
}
extern void myGLibOptionParseFunc(GOptionContext* context, GOptionGroup* group, gpointer data);
static void* getPointer_myGLibOptionParseFunc() {
return (void*)(myGLibOptionParseFunc);
}
extern void myGLibPollFunc(GPollFD* ufds, guint32 nfsd, gint32 timeout_);
static void* getPointer_myGLibPollFunc() {
return (void*)(myGLibPollFunc);
}
extern void myGLibPrintFunc(gchar* string);
static void* getPointer_myGLibPrintFunc() {
return (void*)(myGLibPrintFunc);
}
extern void myGLibRegexEvalCallback(GMatchInfo* match_info, GString* result, gpointer user_data);
static void* getPointer_myGLibRegexEvalCallback() {
return (void*)(myGLibRegexEvalCallback);
}
extern void myGLibScannerMsgFunc(GScanner* scanner, gchar* message, gboolean error);
static void* getPointer_myGLibScannerMsgFunc() {
return (void*)(myGLibScannerMsgFunc);
}
extern void myGLibSequenceIterCompareFunc(GSequenceIter* a, GSequenceIter* b, gpointer data);
static void* getPointer_myGLibSequenceIterCompareFunc() {
return (void*)(myGLibSequenceIterCompareFunc);
}
extern void myGLibSourceDummyMarshal();
static void* getPointer_myGLibSourceDummyMarshal() {
return (void*)(myGLibSourceDummyMarshal);
}
extern void myGLibSourceFunc(gpointer user_data);
static void* getPointer_myGLibSourceFunc() {
return (void*)(myGLibSourceFunc);
}
extern void myGLibSpawnChildSetupFunc(gpointer user_data);
static void* getPointer_myGLibSpawnChildSetupFunc() {
return (void*)(myGLibSpawnChildSetupFunc);
}
extern void myGLibTestDataFunc(gpointer user_data);
static void* getPointer_myGLibTestDataFunc() {
return (void*)(myGLibTestDataFunc);
}
extern void myGLibTestFixtureFunc(gpointer fixture, gpointer user_data);
static void* getPointer_myGLibTestFixtureFunc() {
return (void*)(myGLibTestFixtureFunc);
}
extern void myGLibTestFunc();
static void* getPointer_myGLibTestFunc() {
return (void*)(myGLibTestFunc);
}
extern void myGLibTestLogFatalFunc(gchar* log_domain, GLogLevelFlags log_level, gchar* message, gpointer user_data);
static void* getPointer_myGLibTestLogFatalFunc() {
return (void*)(myGLibTestLogFatalFunc);
}
extern void myGLibThreadFunc(gpointer data);
static void* getPointer_myGLibThreadFunc() {
return (void*)(myGLibThreadFunc);
}
extern void myGLibTranslateFunc(gchar* str, gpointer data);
static void* getPointer_myGLibTranslateFunc() {
return (void*)(myGLibTranslateFunc);
}
extern void myGLibTraverseFunc(gpointer key, gpointer value, gpointer data);
static void* getPointer_myGLibTraverseFunc() {
return (void*)(myGLibTraverseFunc);
}
extern void myGLibUnixFDSourceFunc(gint32 fd, GIOCondition condition, gpointer user_data);
static void* getPointer_myGLibUnixFDSourceFunc() {
return (void*)(myGLibUnixFDSourceFunc);
}
extern void myGLibVoidFunc();
static void* getPointer_myGLibVoidFunc() {
return (void*)(myGLibVoidFunc);
}
*/
import "C"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GLib")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GLib", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct Array
type Array struct {
	P unsafe.Pointer
}

const SizeOfStructArray = 16

func ArrayGetType() gi.GType {
	ret := _I.GetGType(0, "Array")
	return ret
}

// Flags AsciiType
type AsciiTypeFlags int

const (
	AsciiTypeAlnum  AsciiTypeFlags = 1
	AsciiTypeAlpha  AsciiTypeFlags = 2
	AsciiTypeCntrl  AsciiTypeFlags = 4
	AsciiTypeDigit  AsciiTypeFlags = 8
	AsciiTypeGraph  AsciiTypeFlags = 16
	AsciiTypeLower  AsciiTypeFlags = 32
	AsciiTypePrint  AsciiTypeFlags = 64
	AsciiTypePunct  AsciiTypeFlags = 128
	AsciiTypeSpace  AsciiTypeFlags = 256
	AsciiTypeUpper  AsciiTypeFlags = 512
	AsciiTypeXdigit AsciiTypeFlags = 1024
)

func AsciiTypeGetType() gi.GType {
	ret := _I.GetGType(1, "AsciiType")
	return ret
}

// Struct AsyncQueue
type AsyncQueue struct {
	P unsafe.Pointer
}

func AsyncQueueGetType() gi.GType {
	ret := _I.GetGType(2, "AsyncQueue")
	return ret
}

// g_async_queue_length
//
// [ result ] trans: nothing
//
func (v AsyncQueue) Length() (result int32) {
	iv, err := _I.Get(0, "AsyncQueue", "length", 4, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_async_queue_length_unlocked
//
// [ result ] trans: nothing
//
func (v AsyncQueue) LengthUnlocked() (result int32) {
	iv, err := _I.Get(1, "AsyncQueue", "length_unlocked", 4, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_async_queue_lock
//
func (v AsyncQueue) Lock() {
	iv, err := _I.Get(2, "AsyncQueue", "lock", 4, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_pop
//
// [ result ] trans: nothing
//
func (v AsyncQueue) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(3, "AsyncQueue", "pop", 4, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_pop_unlocked
//
// [ result ] trans: nothing
//
func (v AsyncQueue) PopUnlocked() (result unsafe.Pointer) {
	iv, err := _I.Get(4, "AsyncQueue", "pop_unlocked", 4, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_push
//
// [ data ] trans: nothing
//
func (v AsyncQueue) Push(data unsafe.Pointer) {
	iv, err := _I.Get(5, "AsyncQueue", "push", 4, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_async_queue_push_front
//
// [ item ] trans: nothing
//
func (v AsyncQueue) PushFront(item unsafe.Pointer) {
	iv, err := _I.Get(6, "AsyncQueue", "push_front", 4, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	iv.Call(args, nil, nil)
}

// g_async_queue_push_front_unlocked
//
// [ item ] trans: nothing
//
func (v AsyncQueue) PushFrontUnlocked(item unsafe.Pointer) {
	iv, err := _I.Get(7, "AsyncQueue", "push_front_unlocked", 4, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	iv.Call(args, nil, nil)
}

// g_async_queue_push_unlocked
//
// [ data ] trans: nothing
//
func (v AsyncQueue) PushUnlocked(data unsafe.Pointer) {
	iv, err := _I.Get(8, "AsyncQueue", "push_unlocked", 4, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_async_queue_ref_unlocked
//
func (v AsyncQueue) RefUnlocked() {
	iv, err := _I.Get(9, "AsyncQueue", "ref_unlocked", 4, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_remove
//
// [ item ] trans: nothing
//
// [ result ] trans: nothing
//
func (v AsyncQueue) Remove(item unsafe.Pointer) (result bool) {
	iv, err := _I.Get(10, "AsyncQueue", "remove", 4, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_async_queue_remove_unlocked
//
// [ item ] trans: nothing
//
// [ result ] trans: nothing
//
func (v AsyncQueue) RemoveUnlocked(item unsafe.Pointer) (result bool) {
	iv, err := _I.Get(11, "AsyncQueue", "remove_unlocked", 4, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// g_async_queue_timed_pop
//
// [ end_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v AsyncQueue) TimedPop(end_time TimeVal) (result unsafe.Pointer) {
	iv, err := _I.Get(12, "AsyncQueue", "timed_pop", 4, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_end_time := gi.NewPointerArgument(end_time.P)
	args := []gi.Argument{arg_v, arg_end_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Deprecated
//
// g_async_queue_timed_pop_unlocked
//
// [ end_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v AsyncQueue) TimedPopUnlocked(end_time TimeVal) (result unsafe.Pointer) {
	iv, err := _I.Get(13, "AsyncQueue", "timed_pop_unlocked", 4, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_end_time := gi.NewPointerArgument(end_time.P)
	args := []gi.Argument{arg_v, arg_end_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_timeout_pop
//
// [ timeout ] trans: nothing
//
// [ result ] trans: nothing
//
func (v AsyncQueue) TimeoutPop(timeout uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(14, "AsyncQueue", "timeout_pop", 4, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_timeout_pop_unlocked
//
// [ timeout ] trans: nothing
//
// [ result ] trans: nothing
//
func (v AsyncQueue) TimeoutPopUnlocked(timeout uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(15, "AsyncQueue", "timeout_pop_unlocked", 4, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_try_pop
//
// [ result ] trans: nothing
//
func (v AsyncQueue) TryPop() (result unsafe.Pointer) {
	iv, err := _I.Get(16, "AsyncQueue", "try_pop", 4, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_try_pop_unlocked
//
// [ result ] trans: nothing
//
func (v AsyncQueue) TryPopUnlocked() (result unsafe.Pointer) {
	iv, err := _I.Get(17, "AsyncQueue", "try_pop_unlocked", 4, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_unlock
//
func (v AsyncQueue) Unlock() {
	iv, err := _I.Get(18, "AsyncQueue", "unlock", 4, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_unref
//
func (v AsyncQueue) Unref() {
	iv, err := _I.Get(19, "AsyncQueue", "unref", 4, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_async_queue_unref_and_unlock
//
func (v AsyncQueue) UnrefAndUnlock() {
	iv, err := _I.Get(20, "AsyncQueue", "unref_and_unlock", 4, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct BookmarkFile
type BookmarkFile struct {
	P unsafe.Pointer
}

func BookmarkFileGetType() gi.GType {
	ret := _I.GetGType(3, "BookmarkFile")
	return ret
}

// g_bookmark_file_add_application
//
// [ uri ] trans: nothing
//
// [ name ] trans: nothing
//
// [ exec ] trans: nothing
//
func (v BookmarkFile) AddApplication(uri string, name string, exec string) {
	iv, err := _I.Get(21, "BookmarkFile", "add_application", 6, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	c_exec := gi.CString(exec)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_exec := gi.NewStringArgument(c_exec)
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_exec}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_name)
	gi.Free(c_exec)
}

// g_bookmark_file_add_group
//
// [ uri ] trans: nothing
//
// [ group ] trans: nothing
//
func (v BookmarkFile) AddGroup(uri string, group string) {
	iv, err := _I.Get(22, "BookmarkFile", "add_group", 6, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_group := gi.CString(group)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_group := gi.NewStringArgument(c_group)
	args := []gi.Argument{arg_v, arg_uri, arg_group}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_group)
}

// g_bookmark_file_free
//
func (v BookmarkFile) Free() {
	iv, err := _I.Get(23, "BookmarkFile", "free", 6, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_bookmark_file_get_added
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetAdded(uri string) (result int64, err error) {
	iv, err := _I.Get(24, "BookmarkFile", "get_added", 6, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_bookmark_file_get_app_info
//
// [ uri ] trans: nothing
//
// [ name ] trans: nothing
//
// [ exec ] trans: everything, dir: out
//
// [ count ] trans: everything, dir: out
//
// [ stamp ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetAppInfo(uri string, name string) (result bool, exec string, count uint32, stamp int64, err error) {
	iv, err := _I.Get(25, "BookmarkFile", "get_app_info", 6, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_exec := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_count := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_stamp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_exec, arg_count, arg_stamp, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	err = gi.ToError(outArgs[3].Pointer())
	exec = outArgs[0].String().Take()
	count = outArgs[1].Uint32()
	stamp = outArgs[2].Int64()
	result = ret.Bool()
	return
}

// g_bookmark_file_get_applications
//
// [ uri ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v BookmarkFile) GetApplications(uri string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(26, "BookmarkFile", "get_applications", 6, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_uri, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_get_description
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func (v BookmarkFile) GetDescription(uri string) (result string, err error) {
	iv, err := _I.Get(27, "BookmarkFile", "get_description", 6, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_bookmark_file_get_groups
//
// [ uri ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v BookmarkFile) GetGroups(uri string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(28, "BookmarkFile", "get_groups", 6, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_uri, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_get_icon
//
// [ uri ] trans: nothing
//
// [ href ] trans: everything, dir: out
//
// [ mime_type ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetIcon(uri string) (result bool, href string, mime_type string, err error) {
	iv, err := _I.Get(29, "BookmarkFile", "get_icon", 6, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_href := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_mime_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_uri, arg_href, arg_mime_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[2].Pointer())
	href = outArgs[0].String().Take()
	mime_type = outArgs[1].String().Take()
	result = ret.Bool()
	return
}

// g_bookmark_file_get_is_private
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetIsPrivate(uri string) (result bool, err error) {
	iv, err := _I.Get(30, "BookmarkFile", "get_is_private", 6, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_get_mime_type
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func (v BookmarkFile) GetMimeType(uri string) (result string, err error) {
	iv, err := _I.Get(31, "BookmarkFile", "get_mime_type", 6, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_bookmark_file_get_modified
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetModified(uri string) (result int64, err error) {
	iv, err := _I.Get(32, "BookmarkFile", "get_modified", 6, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_bookmark_file_get_size
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetSize() (result int32) {
	iv, err := _I.Get(33, "BookmarkFile", "get_size", 6, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bookmark_file_get_title
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func (v BookmarkFile) GetTitle(uri string) (result string, err error) {
	iv, err := _I.Get(34, "BookmarkFile", "get_title", 6, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_bookmark_file_get_uris
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v BookmarkFile) GetUris() (result gi.CStrArray) {
	iv, err := _I.Get(35, "BookmarkFile", "get_uris", 6, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_get_visited
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) GetVisited(uri string) (result int64, err error) {
	iv, err := _I.Get(36, "BookmarkFile", "get_visited", 6, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_bookmark_file_has_application
//
// [ uri ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) HasApplication(uri string, name string) (result bool, err error) {
	iv, err := _I.Get(37, "BookmarkFile", "has_application", 6, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_has_group
//
// [ uri ] trans: nothing
//
// [ group ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) HasGroup(uri string, group string) (result bool, err error) {
	iv, err := _I.Get(38, "BookmarkFile", "has_group", 6, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_group := gi.CString(group)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_group := gi.NewStringArgument(c_group)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_group, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_group)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_has_item
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) HasItem(uri string) (result bool) {
	iv, err := _I.Get(39, "BookmarkFile", "has_item", 6, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	args := []gi.Argument{arg_v, arg_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri)
	result = ret.Bool()
	return
}

// g_bookmark_file_load_from_data
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) LoadFromData(data gi.Uint8Array, length uint64) (result bool, err error) {
	iv, err := _I.Get(40, "BookmarkFile", "load_from_data", 6, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_load_from_data_dirs
//
// [ file ] trans: nothing
//
// [ full_path ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BookmarkFile) LoadFromDataDirs(file string) (result bool, full_path string, err error) {
	iv, err := _I.Get(41, "BookmarkFile", "load_from_data_dirs", 6, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_full_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_file, arg_full_path, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[1].Pointer())
	full_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_bookmark_file_load_from_file
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) LoadFromFile(filename string) (result bool, err error) {
	iv, err := _I.Get(42, "BookmarkFile", "load_from_file", 6, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_move_item
//
// [ old_uri ] trans: nothing
//
// [ new_uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) MoveItem(old_uri string, new_uri string) (result bool, err error) {
	iv, err := _I.Get(43, "BookmarkFile", "move_item", 6, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_old_uri := gi.CString(old_uri)
	c_new_uri := gi.CString(new_uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_old_uri := gi.NewStringArgument(c_old_uri)
	arg_new_uri := gi.NewStringArgument(c_new_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_old_uri, arg_new_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_old_uri)
	gi.Free(c_new_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_remove_application
//
// [ uri ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) RemoveApplication(uri string, name string) (result bool, err error) {
	iv, err := _I.Get(44, "BookmarkFile", "remove_application", 6, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_remove_group
//
// [ uri ] trans: nothing
//
// [ group ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) RemoveGroup(uri string, group string) (result bool, err error) {
	iv, err := _I.Get(45, "BookmarkFile", "remove_group", 6, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_group := gi.CString(group)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_group := gi.NewStringArgument(c_group)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_group, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_group)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_remove_item
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) RemoveItem(uri string) (result bool, err error) {
	iv, err := _I.Get(46, "BookmarkFile", "remove_item", 6, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_set_added
//
// [ uri ] trans: nothing
//
// [ added ] trans: nothing
//
func (v BookmarkFile) SetAdded(uri string, added int64) {
	iv, err := _I.Get(47, "BookmarkFile", "set_added", 6, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_added := gi.NewInt64Argument(added)
	args := []gi.Argument{arg_v, arg_uri, arg_added}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_app_info
//
// [ uri ] trans: nothing
//
// [ name ] trans: nothing
//
// [ exec ] trans: nothing
//
// [ count ] trans: nothing
//
// [ stamp ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) SetAppInfo(uri string, name string, exec string, count int32, stamp int64) (result bool, err error) {
	iv, err := _I.Get(48, "BookmarkFile", "set_app_info", 6, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	c_exec := gi.CString(exec)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_exec := gi.NewStringArgument(c_exec)
	arg_count := gi.NewInt32Argument(count)
	arg_stamp := gi.NewInt64Argument(stamp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_exec, arg_count, arg_stamp, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	gi.Free(c_exec)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_set_description
//
// [ uri ] trans: nothing
//
// [ description ] trans: nothing
//
func (v BookmarkFile) SetDescription(uri string, description string) {
	iv, err := _I.Get(49, "BookmarkFile", "set_description", 6, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_uri, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_description)
}

// g_bookmark_file_set_groups
//
// [ uri ] trans: nothing
//
// [ groups ] trans: nothing
//
// [ length ] trans: nothing
//
func (v BookmarkFile) SetGroups(uri string, groups gi.CStrArray, length uint64) {
	iv, err := _I.Get(50, "BookmarkFile", "set_groups", 6, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_groups := gi.NewPointerArgument(groups.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_uri, arg_groups, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_icon
//
// [ uri ] trans: nothing
//
// [ href ] trans: nothing
//
// [ mime_type ] trans: nothing
//
func (v BookmarkFile) SetIcon(uri string, href string, mime_type string) {
	iv, err := _I.Get(51, "BookmarkFile", "set_icon", 6, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_href := gi.CString(href)
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_href := gi.NewStringArgument(c_href)
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_uri, arg_href, arg_mime_type}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_href)
	gi.Free(c_mime_type)
}

// g_bookmark_file_set_is_private
//
// [ uri ] trans: nothing
//
// [ is_private ] trans: nothing
//
func (v BookmarkFile) SetIsPrivate(uri string, is_private bool) {
	iv, err := _I.Get(52, "BookmarkFile", "set_is_private", 6, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_is_private := gi.NewBoolArgument(is_private)
	args := []gi.Argument{arg_v, arg_uri, arg_is_private}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_mime_type
//
// [ uri ] trans: nothing
//
// [ mime_type ] trans: nothing
//
func (v BookmarkFile) SetMimeType(uri string, mime_type string) {
	iv, err := _I.Get(53, "BookmarkFile", "set_mime_type", 6, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_uri, arg_mime_type}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_mime_type)
}

// g_bookmark_file_set_modified
//
// [ uri ] trans: nothing
//
// [ modified ] trans: nothing
//
func (v BookmarkFile) SetModified(uri string, modified int64) {
	iv, err := _I.Get(54, "BookmarkFile", "set_modified", 6, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_modified := gi.NewInt64Argument(modified)
	args := []gi.Argument{arg_v, arg_uri, arg_modified}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_title
//
// [ uri ] trans: nothing
//
// [ title ] trans: nothing
//
func (v BookmarkFile) SetTitle(uri string, title string) {
	iv, err := _I.Get(55, "BookmarkFile", "set_title", 6, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_title := gi.CString(title)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_title := gi.NewStringArgument(c_title)
	args := []gi.Argument{arg_v, arg_uri, arg_title}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_title)
}

// g_bookmark_file_set_visited
//
// [ uri ] trans: nothing
//
// [ visited ] trans: nothing
//
func (v BookmarkFile) SetVisited(uri string, visited int64) {
	iv, err := _I.Get(56, "BookmarkFile", "set_visited", 6, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_visited := gi.NewInt64Argument(visited)
	args := []gi.Argument{arg_v, arg_uri, arg_visited}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_to_data
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v BookmarkFile) ToData() (result gi.Uint8Array, err error) {
	iv, err := _I.Get(57, "BookmarkFile", "to_data", 6, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_to_file
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BookmarkFile) ToFile(filename string) (result bool, err error) {
	iv, err := _I.Get(58, "BookmarkFile", "to_file", 6, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Enum BookmarkFileError
type BookmarkFileErrorEnum int

const (
	BookmarkFileErrorInvalidUri       BookmarkFileErrorEnum = 0
	BookmarkFileErrorInvalidValue     BookmarkFileErrorEnum = 1
	BookmarkFileErrorAppNotRegistered BookmarkFileErrorEnum = 2
	BookmarkFileErrorUriNotFound      BookmarkFileErrorEnum = 3
	BookmarkFileErrorRead             BookmarkFileErrorEnum = 4
	BookmarkFileErrorUnknownEncoding  BookmarkFileErrorEnum = 5
	BookmarkFileErrorWrite            BookmarkFileErrorEnum = 6
	BookmarkFileErrorFileNotFound     BookmarkFileErrorEnum = 7
)

func BookmarkFileErrorGetType() gi.GType {
	ret := _I.GetGType(4, "BookmarkFileError")
	return ret
}

// Struct ByteArray
type ByteArray struct {
	P unsafe.Pointer
}

const SizeOfStructByteArray = 16

func ByteArrayGetType() gi.GType {
	ret := _I.GetGType(5, "ByteArray")
	return ret
}

// g_byte_array_free
//
// [ array ] trans: nothing
//
// [ free_segment ] trans: nothing
//
// [ result ] trans: nothing
//
func ByteArrayFree1(array ByteArray, free_segment bool) (result uint8) {
	iv, err := _I.Get(60, "ByteArray", "free", 8, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	arg_free_segment := gi.NewBoolArgument(free_segment)
	args := []gi.Argument{arg_array, arg_free_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_byte_array_free_to_bytes
//
// [ array ] trans: everything
//
// [ result ] trans: everything
//
func ByteArrayFreeToBytes1(array ByteArray) (result Bytes) {
	iv, err := _I.Get(61, "ByteArray", "free_to_bytes", 8, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_new_take
//
// [ data ] trans: everything
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func ByteArrayNewTake1(data gi.Uint8Array, len1 uint64) (result ByteArray) {
	iv, err := _I.Get(63, "ByteArray", "new_take", 8, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_data, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_unref
//
// [ array ] trans: nothing
//
func ByteArrayUnref1(array ByteArray) {
	iv, err := _I.Get(64, "ByteArray", "unref", 8, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	iv.Call(args, nil, nil)
}

// Struct Bytes
type Bytes struct {
	P unsafe.Pointer
}

func BytesGetType() gi.GType {
	ret := _I.GetGType(6, "Bytes")
	return ret
}

// g_bytes_new
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func NewBytes(data gi.Uint8Array, size uint64) (result Bytes) {
	iv, err := _I.Get(65, "Bytes", "new", 9, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_new_take
//
// [ data ] trans: everything
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func NewBytesTake(data gi.Uint8Array, size uint64) (result Bytes) {
	iv, err := _I.Get(66, "Bytes", "new_take", 9, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_compare
//
// [ bytes2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bytes) Compare(bytes2 Bytes) (result int32) {
	iv, err := _I.Get(67, "Bytes", "compare", 9, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes2 := gi.NewPointerArgument(bytes2.P)
	args := []gi.Argument{arg_v, arg_bytes2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bytes_equal
//
// [ bytes2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bytes) Equal(bytes2 Bytes) (result bool) {
	iv, err := _I.Get(68, "Bytes", "equal", 9, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes2 := gi.NewPointerArgument(bytes2.P)
	args := []gi.Argument{arg_v, arg_bytes2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_bytes_get_data
//
// [ size ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Bytes) GetData() (result gi.Uint8Array) {
	iv, err := _I.Get(69, "Bytes", "get_data", 9, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint64
	_ = size
	size = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(size)}
	return
}

// g_bytes_get_size
//
// [ result ] trans: nothing
//
func (v Bytes) GetSize() (result uint64) {
	iv, err := _I.Get(70, "Bytes", "get_size", 9, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_bytes_hash
//
// [ result ] trans: nothing
//
func (v Bytes) Hash() (result uint32) {
	iv, err := _I.Get(71, "Bytes", "hash", 9, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_bytes_new_from_bytes
//
// [ offset ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bytes) NewFromBytes(offset uint64, length uint64) (result Bytes) {
	iv, err := _I.Get(72, "Bytes", "new_from_bytes", 9, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_offset, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_ref
//
// [ result ] trans: everything
//
func (v Bytes) Ref() (result Bytes) {
	iv, err := _I.Get(73, "Bytes", "ref", 9, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_unref
//
func (v Bytes) Unref() {
	iv, err := _I.Get(74, "Bytes", "unref", 9, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_bytes_unref_to_array
//
// [ result ] trans: everything
//
func (v Bytes) UnrefToArray() (result ByteArray) {
	iv, err := _I.Get(75, "Bytes", "unref_to_array", 9, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_unref_to_data
//
// [ size ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Bytes) UnrefToData() (result gi.Uint8Array) {
	iv, err := _I.Get(76, "Bytes", "unref_to_data", 9, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint64
	_ = size
	size = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(size)}
	return
}

// Struct Checksum
type Checksum struct {
	P unsafe.Pointer
}

func ChecksumGetType() gi.GType {
	ret := _I.GetGType(7, "Checksum")
	return ret
}

// g_checksum_new
//
// [ checksum_type ] trans: nothing
//
// [ result ] trans: everything
//
func NewChecksum(checksum_type ChecksumTypeEnum) (result Checksum) {
	iv, err := _I.Get(77, "Checksum", "new", 13, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	args := []gi.Argument{arg_checksum_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_checksum_copy
//
// [ result ] trans: everything
//
func (v Checksum) Copy() (result Checksum) {
	iv, err := _I.Get(78, "Checksum", "copy", 13, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_checksum_free
//
func (v Checksum) Free() {
	iv, err := _I.Get(79, "Checksum", "free", 13, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_checksum_get_string
//
// [ result ] trans: nothing
//
func (v Checksum) GetString() (result string) {
	iv, err := _I.Get(80, "Checksum", "get_string", 13, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_checksum_reset
//
func (v Checksum) Reset() {
	iv, err := _I.Get(81, "Checksum", "reset", 13, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_checksum_update
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
func (v Checksum) Update(data gi.Uint8Array, length int64) {
	iv, err := _I.Get(82, "Checksum", "update", 13, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_v, arg_data, arg_length}
	iv.Call(args, nil, nil)
}

// g_checksum_type_get_length
//
// [ checksum_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ChecksumTypeGetLength1(checksum_type ChecksumTypeEnum) (result int64) {
	iv, err := _I.Get(83, "Checksum", "type_get_length", 13, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	args := []gi.Argument{arg_checksum_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// Enum ChecksumType
type ChecksumTypeEnum int

const (
	ChecksumTypeMd5    ChecksumTypeEnum = 0
	ChecksumTypeSha1   ChecksumTypeEnum = 1
	ChecksumTypeSha256 ChecksumTypeEnum = 2
	ChecksumTypeSha512 ChecksumTypeEnum = 3
	ChecksumTypeSha384 ChecksumTypeEnum = 4
)

func ChecksumTypeGetType() gi.GType {
	ret := _I.GetGType(8, "ChecksumType")
	return ret
}

type ChildWatchFuncStruct struct {
	F_pid    int32
	F_status int32
}

func GetPointer_myChildWatchFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibChildWatchFunc())
}

//export myGLibChildWatchFunc
func myGLibChildWatchFunc(pid C.gint32, status C.gint32, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &ChildWatchFuncStruct{
		F_pid:    int32(pid),
		F_status: int32(status),
	}
	fn(args)
}

type ClearHandleFuncStruct struct {
	F_handle_id uint32
}

func GetPointer_myClearHandleFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibClearHandleFunc())
}

//export myGLibClearHandleFunc
func myGLibClearHandleFunc(handle_id C.guint32) {
	// TODO: not found user_data
}

type CompareDataFuncStruct struct {
	F_a unsafe.Pointer
	F_b unsafe.Pointer
}

func GetPointer_myCompareDataFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibCompareDataFunc())
}

//export myGLibCompareDataFunc
func myGLibCompareDataFunc(a C.gpointer, b C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CompareDataFuncStruct{
		F_a: unsafe.Pointer(a),
		F_b: unsafe.Pointer(b),
	}
	fn(args)
}

type CompareFuncStruct struct {
	F_a unsafe.Pointer
	F_b unsafe.Pointer
}

func GetPointer_myCompareFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibCompareFunc())
}

//export myGLibCompareFunc
func myGLibCompareFunc(a C.gpointer, b C.gpointer) {
	// TODO: not found user_data
}

// Struct Cond
type Cond struct {
	P unsafe.Pointer
}

const SizeOfStructCond = 16

func CondGetType() gi.GType {
	ret := _I.GetGType(9, "Cond")
	return ret
}

// g_cond_broadcast
//
func (v Cond) Broadcast() {
	iv, err := _I.Get(84, "Cond", "broadcast", 19, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_clear
//
func (v Cond) Clear() {
	iv, err := _I.Get(85, "Cond", "clear", 19, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_init
//
func (v Cond) Init() {
	iv, err := _I.Get(86, "Cond", "init", 19, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_signal
//
func (v Cond) Signal() {
	iv, err := _I.Get(87, "Cond", "signal", 19, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_wait
//
// [ mutex ] trans: nothing
//
func (v Cond) Wait(mutex Mutex) {
	iv, err := _I.Get(88, "Cond", "wait", 19, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	args := []gi.Argument{arg_v, arg_mutex}
	iv.Call(args, nil, nil)
}

// g_cond_wait_until
//
// [ mutex ] trans: nothing
//
// [ end_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Cond) WaitUntil(mutex Mutex, end_time int64) (result bool) {
	iv, err := _I.Get(89, "Cond", "wait_until", 19, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	arg_end_time := gi.NewInt64Argument(end_time)
	args := []gi.Argument{arg_v, arg_mutex, arg_end_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum ConvertError
type ConvertErrorEnum int

const (
	ConvertErrorNoConversion    ConvertErrorEnum = 0
	ConvertErrorIllegalSequence ConvertErrorEnum = 1
	ConvertErrorFailed          ConvertErrorEnum = 2
	ConvertErrorPartialInput    ConvertErrorEnum = 3
	ConvertErrorBadUri          ConvertErrorEnum = 4
	ConvertErrorNotAbsolutePath ConvertErrorEnum = 5
	ConvertErrorNoMemory        ConvertErrorEnum = 6
	ConvertErrorEmbeddedNul     ConvertErrorEnum = 7
)

func ConvertErrorGetType() gi.GType {
	ret := _I.GetGType(10, "ConvertError")
	return ret
}

type CopyFuncStruct struct {
	F_src  unsafe.Pointer
	F_data unsafe.Pointer
}

func GetPointer_myCopyFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibCopyFunc())
}

//export myGLibCopyFunc
func myGLibCopyFunc(src C.gpointer, data C.gpointer) {
	// TODO: not found user_data
}

// Struct Data
type Data struct {
	P unsafe.Pointer
}

func DataGetType() gi.GType {
	ret := _I.GetGType(11, "Data")
	return ret
}

type DataForeachFuncStruct struct {
	F_key_id uint32
	F_data   unsafe.Pointer
}

func GetPointer_myDataForeachFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibDataForeachFunc())
}

//export myGLibDataForeachFunc
func myGLibDataForeachFunc(key_id C.guint32, data C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &DataForeachFuncStruct{
		F_key_id: uint32(key_id),
		F_data:   unsafe.Pointer(data),
	}
	fn(args)
}

// Struct Date
type Date struct {
	P unsafe.Pointer
}

const SizeOfStructDate = 24

func DateGetType() gi.GType {
	ret := _I.GetGType(12, "Date")
	return ret
}

// g_date_new
//
// [ result ] trans: everything
//
func NewDate() (result Date) {
	iv, err := _I.Get(90, "Date", "new", 30, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_new_dmy
//
// [ day ] trans: nothing
//
// [ month ] trans: nothing
//
// [ year ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateDmy(day uint8, month DateMonthEnum, year uint16) (result Date) {
	iv, err := _I.Get(91, "Date", "new_dmy", 30, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_day, arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_new_julian
//
// [ julian_day ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateJulian(julian_day uint32) (result Date) {
	iv, err := _I.Get(92, "Date", "new_julian", 30, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_julian_day := gi.NewUint32Argument(julian_day)
	args := []gi.Argument{arg_julian_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_add_days
//
// [ n_days ] trans: nothing
//
func (v Date) AddDays(n_days uint32) {
	iv, err := _I.Get(93, "Date", "add_days", 30, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_days := gi.NewUint32Argument(n_days)
	args := []gi.Argument{arg_v, arg_n_days}
	iv.Call(args, nil, nil)
}

// g_date_add_months
//
// [ n_months ] trans: nothing
//
func (v Date) AddMonths(n_months uint32) {
	iv, err := _I.Get(94, "Date", "add_months", 30, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_months := gi.NewUint32Argument(n_months)
	args := []gi.Argument{arg_v, arg_n_months}
	iv.Call(args, nil, nil)
}

// g_date_add_years
//
// [ n_years ] trans: nothing
//
func (v Date) AddYears(n_years uint32) {
	iv, err := _I.Get(95, "Date", "add_years", 30, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_years := gi.NewUint32Argument(n_years)
	args := []gi.Argument{arg_v, arg_n_years}
	iv.Call(args, nil, nil)
}

// g_date_clamp
//
// [ min_date ] trans: nothing
//
// [ max_date ] trans: nothing
//
func (v Date) Clamp(min_date Date, max_date Date) {
	iv, err := _I.Get(96, "Date", "clamp", 30, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_min_date := gi.NewPointerArgument(min_date.P)
	arg_max_date := gi.NewPointerArgument(max_date.P)
	args := []gi.Argument{arg_v, arg_min_date, arg_max_date}
	iv.Call(args, nil, nil)
}

// g_date_clear
//
// [ n_dates ] trans: nothing
//
func (v Date) Clear(n_dates uint32) {
	iv, err := _I.Get(97, "Date", "clear", 30, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_dates := gi.NewUint32Argument(n_dates)
	args := []gi.Argument{arg_v, arg_n_dates}
	iv.Call(args, nil, nil)
}

// g_date_compare
//
// [ rhs ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Date) Compare(rhs Date) (result int32) {
	iv, err := _I.Get(98, "Date", "compare", 30, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rhs := gi.NewPointerArgument(rhs.P)
	args := []gi.Argument{arg_v, arg_rhs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_copy
//
// [ result ] trans: everything
//
func (v Date) Copy() (result Date) {
	iv, err := _I.Get(99, "Date", "copy", 30, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_days_between
//
// [ date2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Date) DaysBetween(date2 Date) (result int32) {
	iv, err := _I.Get(100, "Date", "days_between", 30, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_date2 := gi.NewPointerArgument(date2.P)
	args := []gi.Argument{arg_v, arg_date2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_free
//
func (v Date) Free() {
	iv, err := _I.Get(101, "Date", "free", 30, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_date_get_day
//
// [ result ] trans: nothing
//
func (v Date) GetDay() (result uint8) {
	iv, err := _I.Get(102, "Date", "get_day", 30, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_day_of_year
//
// [ result ] trans: nothing
//
func (v Date) GetDayOfYear() (result uint32) {
	iv, err := _I.Get(103, "Date", "get_day_of_year", 30, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_get_iso8601_week_of_year
//
// [ result ] trans: nothing
//
func (v Date) GetIso8601WeekOfYear() (result uint32) {
	iv, err := _I.Get(104, "Date", "get_iso8601_week_of_year", 30, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_get_julian
//
// [ result ] trans: nothing
//
func (v Date) GetJulian() (result uint32) {
	iv, err := _I.Get(105, "Date", "get_julian", 30, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_get_monday_week_of_year
//
// [ result ] trans: nothing
//
func (v Date) GetMondayWeekOfYear() (result uint32) {
	iv, err := _I.Get(106, "Date", "get_monday_week_of_year", 30, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_get_month
//
// [ result ] trans: nothing
//
func (v Date) GetMonth() (result DateMonthEnum) {
	iv, err := _I.Get(107, "Date", "get_month", 30, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DateMonthEnum(ret.Int())
	return
}

// g_date_get_sunday_week_of_year
//
// [ result ] trans: nothing
//
func (v Date) GetSundayWeekOfYear() (result uint32) {
	iv, err := _I.Get(108, "Date", "get_sunday_week_of_year", 30, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_get_weekday
//
// [ result ] trans: nothing
//
func (v Date) GetWeekday() (result DateWeekdayEnum) {
	iv, err := _I.Get(109, "Date", "get_weekday", 30, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DateWeekdayEnum(ret.Int())
	return
}

// g_date_get_year
//
// [ result ] trans: nothing
//
func (v Date) GetYear() (result uint16) {
	iv, err := _I.Get(110, "Date", "get_year", 30, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// g_date_is_first_of_month
//
// [ result ] trans: nothing
//
func (v Date) IsFirstOfMonth() (result bool) {
	iv, err := _I.Get(111, "Date", "is_first_of_month", 30, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_is_last_of_month
//
// [ result ] trans: nothing
//
func (v Date) IsLastOfMonth() (result bool) {
	iv, err := _I.Get(112, "Date", "is_last_of_month", 30, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_order
//
// [ date2 ] trans: nothing
//
func (v Date) Order(date2 Date) {
	iv, err := _I.Get(113, "Date", "order", 30, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_date2 := gi.NewPointerArgument(date2.P)
	args := []gi.Argument{arg_v, arg_date2}
	iv.Call(args, nil, nil)
}

// g_date_set_day
//
// [ day ] trans: nothing
//
func (v Date) SetDay(day uint8) {
	iv, err := _I.Get(114, "Date", "set_day", 30, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_day := gi.NewUint8Argument(day)
	args := []gi.Argument{arg_v, arg_day}
	iv.Call(args, nil, nil)
}

// g_date_set_dmy
//
// [ day ] trans: nothing
//
// [ month ] trans: nothing
//
// [ y ] trans: nothing
//
func (v Date) SetDmy(day uint8, month DateMonthEnum, y uint16) {
	iv, err := _I.Get(115, "Date", "set_dmy", 30, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_y := gi.NewUint16Argument(y)
	args := []gi.Argument{arg_v, arg_day, arg_month, arg_y}
	iv.Call(args, nil, nil)
}

// g_date_set_julian
//
// [ julian_date ] trans: nothing
//
func (v Date) SetJulian(julian_date uint32) {
	iv, err := _I.Get(116, "Date", "set_julian", 30, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_julian_date := gi.NewUint32Argument(julian_date)
	args := []gi.Argument{arg_v, arg_julian_date}
	iv.Call(args, nil, nil)
}

// g_date_set_month
//
// [ month ] trans: nothing
//
func (v Date) SetMonth(month DateMonthEnum) {
	iv, err := _I.Get(117, "Date", "set_month", 30, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_month := gi.NewIntArgument(int(month))
	args := []gi.Argument{arg_v, arg_month}
	iv.Call(args, nil, nil)
}

// g_date_set_parse
//
// [ str ] trans: nothing
//
func (v Date) SetParse(str string) {
	iv, err := _I.Get(118, "Date", "set_parse", 30, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_v, arg_str}
	iv.Call(args, nil, nil)
	gi.Free(c_str)
}

// Deprecated
//
// g_date_set_time
//
// [ time_ ] trans: nothing
//
func (v Date) SetTime(time_ int32) {
	iv, err := _I.Get(119, "Date", "set_time", 30, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time_ := gi.NewInt32Argument(time_)
	args := []gi.Argument{arg_v, arg_time_}
	iv.Call(args, nil, nil)
}

// g_date_set_time_t
//
// [ timet ] trans: nothing
//
func (v Date) SetTimeT(timet int64) {
	iv, err := _I.Get(120, "Date", "set_time_t", 30, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timet := gi.NewInt64Argument(timet)
	args := []gi.Argument{arg_v, arg_timet}
	iv.Call(args, nil, nil)
}

// g_date_set_time_val
//
// [ timeval ] trans: nothing
//
func (v Date) SetTimeVal(timeval TimeVal) {
	iv, err := _I.Get(121, "Date", "set_time_val", 30, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeval := gi.NewPointerArgument(timeval.P)
	args := []gi.Argument{arg_v, arg_timeval}
	iv.Call(args, nil, nil)
}

// g_date_set_year
//
// [ year ] trans: nothing
//
func (v Date) SetYear(year uint16) {
	iv, err := _I.Get(122, "Date", "set_year", 30, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_v, arg_year}
	iv.Call(args, nil, nil)
}

// g_date_subtract_days
//
// [ n_days ] trans: nothing
//
func (v Date) SubtractDays(n_days uint32) {
	iv, err := _I.Get(123, "Date", "subtract_days", 30, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_days := gi.NewUint32Argument(n_days)
	args := []gi.Argument{arg_v, arg_n_days}
	iv.Call(args, nil, nil)
}

// g_date_subtract_months
//
// [ n_months ] trans: nothing
//
func (v Date) SubtractMonths(n_months uint32) {
	iv, err := _I.Get(124, "Date", "subtract_months", 30, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_months := gi.NewUint32Argument(n_months)
	args := []gi.Argument{arg_v, arg_n_months}
	iv.Call(args, nil, nil)
}

// g_date_subtract_years
//
// [ n_years ] trans: nothing
//
func (v Date) SubtractYears(n_years uint32) {
	iv, err := _I.Get(125, "Date", "subtract_years", 30, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_years := gi.NewUint32Argument(n_years)
	args := []gi.Argument{arg_v, arg_n_years}
	iv.Call(args, nil, nil)
}

// g_date_to_struct_tm
//
// [ tm ] trans: nothing
//
func (v Date) ToStructTm(tm unsafe.Pointer) {
	iv, err := _I.Get(126, "Date", "to_struct_tm", 30, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tm := gi.NewPointerArgument(tm)
	args := []gi.Argument{arg_v, arg_tm}
	iv.Call(args, nil, nil)
}

// g_date_valid
//
// [ result ] trans: nothing
//
func (v Date) Valid() (result bool) {
	iv, err := _I.Get(127, "Date", "valid", 30, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_get_days_in_month
//
// [ month ] trans: nothing
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateGetDaysInMonth1(month DateMonthEnum, year uint16) (result uint8) {
	iv, err := _I.Get(128, "Date", "get_days_in_month", 30, 38, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_monday_weeks_in_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateGetMondayWeeksInYear1(year uint16) (result uint8) {
	iv, err := _I.Get(129, "Date", "get_monday_weeks_in_year", 30, 39, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_sunday_weeks_in_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateGetSundayWeeksInYear1(year uint16) (result uint8) {
	iv, err := _I.Get(130, "Date", "get_sunday_weeks_in_year", 30, 40, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_is_leap_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateIsLeapYear1(year uint16) (result bool) {
	iv, err := _I.Get(131, "Date", "is_leap_year", 30, 41, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_strftime
//
// [ s ] trans: nothing
//
// [ slen ] trans: nothing
//
// [ format ] trans: nothing
//
// [ date ] trans: nothing
//
// [ result ] trans: nothing
//
func DateStrftime1(s string, slen uint64, format string, date Date) (result uint64) {
	iv, err := _I.Get(132, "Date", "strftime", 30, 42, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s := gi.CString(s)
	c_format := gi.CString(format)
	arg_s := gi.NewStringArgument(c_s)
	arg_slen := gi.NewUint64Argument(slen)
	arg_format := gi.NewStringArgument(c_format)
	arg_date := gi.NewPointerArgument(date.P)
	args := []gi.Argument{arg_s, arg_slen, arg_format, arg_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s)
	gi.Free(c_format)
	result = ret.Uint64()
	return
}

// g_date_valid_day
//
// [ day ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidDay1(day uint8) (result bool) {
	iv, err := _I.Get(133, "Date", "valid_day", 30, 43, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	args := []gi.Argument{arg_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_dmy
//
// [ day ] trans: nothing
//
// [ month ] trans: nothing
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidDmy1(day uint8, month DateMonthEnum, year uint16) (result bool) {
	iv, err := _I.Get(134, "Date", "valid_dmy", 30, 44, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_day, arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_julian
//
// [ julian_date ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidJulian1(julian_date uint32) (result bool) {
	iv, err := _I.Get(135, "Date", "valid_julian", 30, 45, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_julian_date := gi.NewUint32Argument(julian_date)
	args := []gi.Argument{arg_julian_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_month
//
// [ month ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidMonth1(month DateMonthEnum) (result bool) {
	iv, err := _I.Get(136, "Date", "valid_month", 30, 46, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	args := []gi.Argument{arg_month}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_weekday
//
// [ weekday ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidWeekday1(weekday DateWeekdayEnum) (result bool) {
	iv, err := _I.Get(137, "Date", "valid_weekday", 30, 47, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_weekday := gi.NewIntArgument(int(weekday))
	args := []gi.Argument{arg_weekday}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidYear1(year uint16) (result bool) {
	iv, err := _I.Get(138, "Date", "valid_year", 30, 48, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum DateDMY
type DateDMYEnum int

const (
	DateDMYDay   DateDMYEnum = 0
	DateDMYMonth DateDMYEnum = 1
	DateDMYYear  DateDMYEnum = 2
)

func DateDMYGetType() gi.GType {
	ret := _I.GetGType(13, "DateDMY")
	return ret
}

// Enum DateMonth
type DateMonthEnum int

const (
	DateMonthBadMonth  DateMonthEnum = 0
	DateMonthJanuary   DateMonthEnum = 1
	DateMonthFebruary  DateMonthEnum = 2
	DateMonthMarch     DateMonthEnum = 3
	DateMonthApril     DateMonthEnum = 4
	DateMonthMay       DateMonthEnum = 5
	DateMonthJune      DateMonthEnum = 6
	DateMonthJuly      DateMonthEnum = 7
	DateMonthAugust    DateMonthEnum = 8
	DateMonthSeptember DateMonthEnum = 9
	DateMonthOctober   DateMonthEnum = 10
	DateMonthNovember  DateMonthEnum = 11
	DateMonthDecember  DateMonthEnum = 12
)

func DateMonthGetType() gi.GType {
	ret := _I.GetGType(14, "DateMonth")
	return ret
}

// Struct DateTime
type DateTime struct {
	P unsafe.Pointer
}

func DateTimeGetType() gi.GType {
	ret := _I.GetGType(15, "DateTime")
	return ret
}

// g_date_time_new
//
// [ tz ] trans: nothing
//
// [ year ] trans: nothing
//
// [ month ] trans: nothing
//
// [ day ] trans: nothing
//
// [ hour ] trans: nothing
//
// [ minute ] trans: nothing
//
// [ seconds ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTime(tz TimeZone, year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(139, "DateTime", "new", 33, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tz := gi.NewPointerArgument(tz.P)
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_tz, arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_iso8601
//
// [ text ] trans: nothing
//
// [ default_tz ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromIso8601(text string, default_tz TimeZone) (result DateTime) {
	iv, err := _I.Get(140, "DateTime", "new_from_iso8601", 33, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_default_tz := gi.NewPointerArgument(default_tz.P)
	args := []gi.Argument{arg_text, arg_default_tz}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_timeval_local
//
// [ tv ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromTimevalLocal(tv TimeVal) (result DateTime) {
	iv, err := _I.Get(141, "DateTime", "new_from_timeval_local", 33, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tv := gi.NewPointerArgument(tv.P)
	args := []gi.Argument{arg_tv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_timeval_utc
//
// [ tv ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromTimevalUtc(tv TimeVal) (result DateTime) {
	iv, err := _I.Get(142, "DateTime", "new_from_timeval_utc", 33, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tv := gi.NewPointerArgument(tv.P)
	args := []gi.Argument{arg_tv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_unix_local
//
// [ t ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromUnixLocal(t int64) (result DateTime) {
	iv, err := _I.Get(143, "DateTime", "new_from_unix_local", 33, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_t := gi.NewInt64Argument(t)
	args := []gi.Argument{arg_t}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_unix_utc
//
// [ t ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromUnixUtc(t int64) (result DateTime) {
	iv, err := _I.Get(144, "DateTime", "new_from_unix_utc", 33, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_t := gi.NewInt64Argument(t)
	args := []gi.Argument{arg_t}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_local
//
// [ year ] trans: nothing
//
// [ month ] trans: nothing
//
// [ day ] trans: nothing
//
// [ hour ] trans: nothing
//
// [ minute ] trans: nothing
//
// [ seconds ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeLocal(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(145, "DateTime", "new_local", 33, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_now
//
// [ tz ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeNow(tz TimeZone) (result DateTime) {
	iv, err := _I.Get(146, "DateTime", "new_now", 33, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tz := gi.NewPointerArgument(tz.P)
	args := []gi.Argument{arg_tz}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_now_local
//
// [ result ] trans: everything
//
func NewDateTimeNowLocal() (result DateTime) {
	iv, err := _I.Get(147, "DateTime", "new_now_local", 33, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_now_utc
//
// [ result ] trans: everything
//
func NewDateTimeNowUtc() (result DateTime) {
	iv, err := _I.Get(148, "DateTime", "new_now_utc", 33, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_utc
//
// [ year ] trans: nothing
//
// [ month ] trans: nothing
//
// [ day ] trans: nothing
//
// [ hour ] trans: nothing
//
// [ minute ] trans: nothing
//
// [ seconds ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeUtc(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(149, "DateTime", "new_utc", 33, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add
//
// [ timespan ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) Add(timespan int64) (result DateTime) {
	iv, err := _I.Get(150, "DateTime", "add", 33, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timespan := gi.NewInt64Argument(timespan)
	args := []gi.Argument{arg_v, arg_timespan}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_days
//
// [ days ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddDays(days int32) (result DateTime) {
	iv, err := _I.Get(151, "DateTime", "add_days", 33, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_days := gi.NewInt32Argument(days)
	args := []gi.Argument{arg_v, arg_days}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_full
//
// [ years ] trans: nothing
//
// [ months ] trans: nothing
//
// [ days ] trans: nothing
//
// [ hours ] trans: nothing
//
// [ minutes ] trans: nothing
//
// [ seconds ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(152, "DateTime", "add_full", 33, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_years := gi.NewInt32Argument(years)
	arg_months := gi.NewInt32Argument(months)
	arg_days := gi.NewInt32Argument(days)
	arg_hours := gi.NewInt32Argument(hours)
	arg_minutes := gi.NewInt32Argument(minutes)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_v, arg_years, arg_months, arg_days, arg_hours, arg_minutes, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_hours
//
// [ hours ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddHours(hours int32) (result DateTime) {
	iv, err := _I.Get(153, "DateTime", "add_hours", 33, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hours := gi.NewInt32Argument(hours)
	args := []gi.Argument{arg_v, arg_hours}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_minutes
//
// [ minutes ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddMinutes(minutes int32) (result DateTime) {
	iv, err := _I.Get(154, "DateTime", "add_minutes", 33, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_minutes := gi.NewInt32Argument(minutes)
	args := []gi.Argument{arg_v, arg_minutes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_months
//
// [ months ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddMonths(months int32) (result DateTime) {
	iv, err := _I.Get(155, "DateTime", "add_months", 33, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_months := gi.NewInt32Argument(months)
	args := []gi.Argument{arg_v, arg_months}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_seconds
//
// [ seconds ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddSeconds(seconds float64) (result DateTime) {
	iv, err := _I.Get(156, "DateTime", "add_seconds", 33, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_v, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_weeks
//
// [ weeks ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddWeeks(weeks int32) (result DateTime) {
	iv, err := _I.Get(157, "DateTime", "add_weeks", 33, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_weeks := gi.NewInt32Argument(weeks)
	args := []gi.Argument{arg_v, arg_weeks}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_years
//
// [ years ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) AddYears(years int32) (result DateTime) {
	iv, err := _I.Get(158, "DateTime", "add_years", 33, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_years := gi.NewInt32Argument(years)
	args := []gi.Argument{arg_v, arg_years}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_difference
//
// [ begin ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DateTime) Difference(begin DateTime) (result int64) {
	iv, err := _I.Get(159, "DateTime", "difference", 33, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_begin := gi.NewPointerArgument(begin.P)
	args := []gi.Argument{arg_v, arg_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_date_time_format
//
// [ format ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) Format(format string) (result string) {
	iv, err := _I.Get(160, "DateTime", "format", 33, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_format := gi.CString(format)
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewStringArgument(c_format)
	args := []gi.Argument{arg_v, arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_format)
	result = ret.String().Take()
	return
}

// g_date_time_get_day_of_month
//
// [ result ] trans: nothing
//
func (v DateTime) GetDayOfMonth() (result int32) {
	iv, err := _I.Get(161, "DateTime", "get_day_of_month", 33, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_day_of_week
//
// [ result ] trans: nothing
//
func (v DateTime) GetDayOfWeek() (result int32) {
	iv, err := _I.Get(162, "DateTime", "get_day_of_week", 33, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_day_of_year
//
// [ result ] trans: nothing
//
func (v DateTime) GetDayOfYear() (result int32) {
	iv, err := _I.Get(163, "DateTime", "get_day_of_year", 33, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_hour
//
// [ result ] trans: nothing
//
func (v DateTime) GetHour() (result int32) {
	iv, err := _I.Get(164, "DateTime", "get_hour", 33, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_microsecond
//
// [ result ] trans: nothing
//
func (v DateTime) GetMicrosecond() (result int32) {
	iv, err := _I.Get(165, "DateTime", "get_microsecond", 33, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_minute
//
// [ result ] trans: nothing
//
func (v DateTime) GetMinute() (result int32) {
	iv, err := _I.Get(166, "DateTime", "get_minute", 33, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_month
//
// [ result ] trans: nothing
//
func (v DateTime) GetMonth() (result int32) {
	iv, err := _I.Get(167, "DateTime", "get_month", 33, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_second
//
// [ result ] trans: nothing
//
func (v DateTime) GetSecond() (result int32) {
	iv, err := _I.Get(168, "DateTime", "get_second", 33, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_seconds
//
// [ result ] trans: nothing
//
func (v DateTime) GetSeconds() (result float64) {
	iv, err := _I.Get(169, "DateTime", "get_seconds", 33, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_date_time_get_timezone
//
// [ result ] trans: nothing
//
func (v DateTime) GetTimezone() (result TimeZone) {
	iv, err := _I.Get(170, "DateTime", "get_timezone", 33, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_get_timezone_abbreviation
//
// [ result ] trans: nothing
//
func (v DateTime) GetTimezoneAbbreviation() (result string) {
	iv, err := _I.Get(171, "DateTime", "get_timezone_abbreviation", 33, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_date_time_get_utc_offset
//
// [ result ] trans: nothing
//
func (v DateTime) GetUtcOffset() (result int64) {
	iv, err := _I.Get(172, "DateTime", "get_utc_offset", 33, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_date_time_get_week_numbering_year
//
// [ result ] trans: nothing
//
func (v DateTime) GetWeekNumberingYear() (result int32) {
	iv, err := _I.Get(173, "DateTime", "get_week_numbering_year", 33, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_week_of_year
//
// [ result ] trans: nothing
//
func (v DateTime) GetWeekOfYear() (result int32) {
	iv, err := _I.Get(174, "DateTime", "get_week_of_year", 33, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_year
//
// [ result ] trans: nothing
//
func (v DateTime) GetYear() (result int32) {
	iv, err := _I.Get(175, "DateTime", "get_year", 33, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_get_ymd
//
// [ year ] trans: everything, dir: out
//
// [ month ] trans: everything, dir: out
//
// [ day ] trans: everything, dir: out
//
func (v DateTime) GetYmd() (year int32, month int32, day int32) {
	iv, err := _I.Get(176, "DateTime", "get_ymd", 33, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_year := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_month := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_day := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_year, arg_month, arg_day}
	iv.Call(args, nil, &outArgs[0])
	year = outArgs[0].Int32()
	month = outArgs[1].Int32()
	day = outArgs[2].Int32()
	return
}

// g_date_time_is_daylight_savings
//
// [ result ] trans: nothing
//
func (v DateTime) IsDaylightSavings() (result bool) {
	iv, err := _I.Get(177, "DateTime", "is_daylight_savings", 33, 38, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_ref
//
// [ result ] trans: everything
//
func (v DateTime) Ref() (result DateTime) {
	iv, err := _I.Get(178, "DateTime", "ref", 33, 39, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_to_local
//
// [ result ] trans: everything
//
func (v DateTime) ToLocal() (result DateTime) {
	iv, err := _I.Get(179, "DateTime", "to_local", 33, 40, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_to_timeval
//
// [ tv ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DateTime) ToTimeval(tv TimeVal) (result bool) {
	iv, err := _I.Get(180, "DateTime", "to_timeval", 33, 41, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tv := gi.NewPointerArgument(tv.P)
	args := []gi.Argument{arg_v, arg_tv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_to_timezone
//
// [ tz ] trans: nothing
//
// [ result ] trans: everything
//
func (v DateTime) ToTimezone(tz TimeZone) (result DateTime) {
	iv, err := _I.Get(181, "DateTime", "to_timezone", 33, 42, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tz := gi.NewPointerArgument(tz.P)
	args := []gi.Argument{arg_v, arg_tz}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_to_unix
//
// [ result ] trans: nothing
//
func (v DateTime) ToUnix() (result int64) {
	iv, err := _I.Get(182, "DateTime", "to_unix", 33, 43, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_date_time_to_utc
//
// [ result ] trans: everything
//
func (v DateTime) ToUtc() (result DateTime) {
	iv, err := _I.Get(183, "DateTime", "to_utc", 33, 44, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_unref
//
func (v DateTime) Unref() {
	iv, err := _I.Get(184, "DateTime", "unref", 33, 45, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_date_time_compare
//
// [ dt1 ] trans: nothing
//
// [ dt2 ] trans: nothing
//
// [ result ] trans: nothing
//
func DateTimeCompare1(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result int32) {
	iv, err := _I.Get(185, "DateTime", "compare", 33, 46, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_equal
//
// [ dt1 ] trans: nothing
//
// [ dt2 ] trans: nothing
//
// [ result ] trans: nothing
//
func DateTimeEqual1(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(186, "DateTime", "equal", 33, 47, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_hash
//
// [ datetime ] trans: nothing
//
// [ result ] trans: nothing
//
func DateTimeHash1(datetime unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(187, "DateTime", "hash", 33, 48, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datetime := gi.NewPointerArgument(datetime)
	args := []gi.Argument{arg_datetime}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Enum DateWeekday
type DateWeekdayEnum int

const (
	DateWeekdayBadWeekday DateWeekdayEnum = 0
	DateWeekdayMonday     DateWeekdayEnum = 1
	DateWeekdayTuesday    DateWeekdayEnum = 2
	DateWeekdayWednesday  DateWeekdayEnum = 3
	DateWeekdayThursday   DateWeekdayEnum = 4
	DateWeekdayFriday     DateWeekdayEnum = 5
	DateWeekdaySaturday   DateWeekdayEnum = 6
	DateWeekdaySunday     DateWeekdayEnum = 7
)

func DateWeekdayGetType() gi.GType {
	ret := _I.GetGType(16, "DateWeekday")
	return ret
}

// Struct DebugKey
type DebugKey struct {
	P unsafe.Pointer
}

const SizeOfStructDebugKey = 16

func DebugKeyGetType() gi.GType {
	ret := _I.GetGType(17, "DebugKey")
	return ret
}

type DestroyNotifyStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myDestroyNotify() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibDestroyNotify())
}

//export myGLibDestroyNotify
func myGLibDestroyNotify(data C.gpointer) {
	// TODO: not found user_data
}

// Struct Dir
type Dir struct {
	P unsafe.Pointer
}

func DirGetType() gi.GType {
	ret := _I.GetGType(18, "Dir")
	return ret
}

// g_dir_close
//
func (v Dir) Close() {
	iv, err := _I.Get(188, "Dir", "close", 37, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_dir_read_name
//
// [ result ] trans: nothing
//
func (v Dir) ReadName() (result string) {
	iv, err := _I.Get(189, "Dir", "read_name", 37, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_dir_rewind
//
func (v Dir) Rewind() {
	iv, err := _I.Get(190, "Dir", "rewind", 37, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_dir_make_tmp
//
// [ tmpl ] trans: nothing
//
// [ result ] trans: everything
//
func DirMakeTmp1(tmpl string) (result string, err error) {
	iv, err := _I.Get(191, "Dir", "make_tmp", 37, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_tmpl := gi.CString(tmpl)
	arg_tmpl := gi.NewStringArgument(c_tmpl)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_tmpl, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tmpl)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// Union DoubleIEEE754
type DoubleIEEE754 struct {
	P unsafe.Pointer
}

const SizeOfUnionDoubleIEEE754 = 8

func DoubleIEEE754GetType() gi.GType {
	ret := _I.GetGType(19, "DoubleIEEE754")
	return ret
}

type DuplicateFuncStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myDuplicateFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibDuplicateFunc())
}

//export myGLibDuplicateFunc
func myGLibDuplicateFunc(data C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &DuplicateFuncStruct{
		F_data: unsafe.Pointer(data),
	}
	fn(args)
}

type EqualFuncStruct struct {
	F_a unsafe.Pointer
	F_b unsafe.Pointer
}

func GetPointer_myEqualFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibEqualFunc())
}

//export myGLibEqualFunc
func myGLibEqualFunc(a C.gpointer, b C.gpointer) {
	// TODO: not found user_data
}

// Struct Error
type Error struct {
	P unsafe.Pointer
}

const SizeOfStructError = 16

func ErrorGetType() gi.GType {
	ret := _I.GetGType(20, "Error")
	return ret
}

// g_error_new_literal
//
// [ domain ] trans: nothing
//
// [ code ] trans: nothing
//
// [ message ] trans: nothing
//
// [ result ] trans: everything
//
func NewErrorLiteral(domain uint32, code int32, message string) (result Error) {
	iv, err := _I.Get(192, "Error", "new_literal", 42, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_message := gi.CString(message)
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	arg_message := gi.NewStringArgument(c_message)
	args := []gi.Argument{arg_domain, arg_code, arg_message}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_message)
	result.P = ret.Pointer()
	return
}

// g_error_copy
//
// [ result ] trans: everything
//
func (v Error) Copy() (result Error) {
	iv, err := _I.Get(193, "Error", "copy", 42, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_error_free
//
func (v Error) Free() {
	iv, err := _I.Get(194, "Error", "free", 42, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_error_matches
//
// [ domain ] trans: nothing
//
// [ code ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Error) Matches(domain uint32, code int32) (result bool) {
	iv, err := _I.Get(195, "Error", "matches", 42, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	args := []gi.Argument{arg_v, arg_domain, arg_code}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum ErrorType
type ErrorTypeEnum int

const (
	ErrorTypeUnknown           ErrorTypeEnum = 0
	ErrorTypeUnexpEof          ErrorTypeEnum = 1
	ErrorTypeUnexpEofInString  ErrorTypeEnum = 2
	ErrorTypeUnexpEofInComment ErrorTypeEnum = 3
	ErrorTypeNonDigitInConst   ErrorTypeEnum = 4
	ErrorTypeDigitRadix        ErrorTypeEnum = 5
	ErrorTypeFloatRadix        ErrorTypeEnum = 6
	ErrorTypeFloatMalformed    ErrorTypeEnum = 7
)

func ErrorTypeGetType() gi.GType {
	ret := _I.GetGType(21, "ErrorType")
	return ret
}

// Enum FileError
type FileErrorEnum int

const (
	FileErrorExist       FileErrorEnum = 0
	FileErrorIsdir       FileErrorEnum = 1
	FileErrorAcces       FileErrorEnum = 2
	FileErrorNametoolong FileErrorEnum = 3
	FileErrorNoent       FileErrorEnum = 4
	FileErrorNotdir      FileErrorEnum = 5
	FileErrorNxio        FileErrorEnum = 6
	FileErrorNodev       FileErrorEnum = 7
	FileErrorRofs        FileErrorEnum = 8
	FileErrorTxtbsy      FileErrorEnum = 9
	FileErrorFault       FileErrorEnum = 10
	FileErrorLoop        FileErrorEnum = 11
	FileErrorNospc       FileErrorEnum = 12
	FileErrorNomem       FileErrorEnum = 13
	FileErrorMfile       FileErrorEnum = 14
	FileErrorNfile       FileErrorEnum = 15
	FileErrorBadf        FileErrorEnum = 16
	FileErrorInval       FileErrorEnum = 17
	FileErrorPipe        FileErrorEnum = 18
	FileErrorAgain       FileErrorEnum = 19
	FileErrorIntr        FileErrorEnum = 20
	FileErrorIo          FileErrorEnum = 21
	FileErrorPerm        FileErrorEnum = 22
	FileErrorNosys       FileErrorEnum = 23
	FileErrorFailed      FileErrorEnum = 24
)

func FileErrorGetType() gi.GType {
	ret := _I.GetGType(22, "FileError")
	return ret
}

// Flags FileTest
type FileTestFlags int

const (
	FileTestIsRegular    FileTestFlags = 1
	FileTestIsSymlink    FileTestFlags = 2
	FileTestIsDir        FileTestFlags = 4
	FileTestIsExecutable FileTestFlags = 8
	FileTestExists       FileTestFlags = 16
)

func FileTestGetType() gi.GType {
	ret := _I.GetGType(23, "FileTest")
	return ret
}

// Union FloatIEEE754
type FloatIEEE754 struct {
	P unsafe.Pointer
}

const SizeOfUnionFloatIEEE754 = 4

func FloatIEEE754GetType() gi.GType {
	ret := _I.GetGType(24, "FloatIEEE754")
	return ret
}

// Flags FormatSizeFlags
type FormatSizeFlags int

const (
	FormatSizeFlagsDefault    FormatSizeFlags = 0
	FormatSizeFlagsLongFormat FormatSizeFlags = 1
	FormatSizeFlagsIecUnits   FormatSizeFlags = 2
	FormatSizeFlagsBits       FormatSizeFlags = 4
)

func FormatSizeFlagsGetType() gi.GType {
	ret := _I.GetGType(25, "FormatSizeFlags")
	return ret
}

type FreeFuncStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myFreeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibFreeFunc())
}

//export myGLibFreeFunc
func myGLibFreeFunc(data C.gpointer) {
	// TODO: not found user_data
}

type FuncStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibFunc())
}

//export myGLibFunc
func myGLibFunc(data C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &FuncStruct{
		F_data: unsafe.Pointer(data),
	}
	fn(args)
}

type HFuncStruct struct {
	F_key   unsafe.Pointer
	F_value unsafe.Pointer
}

func GetPointer_myHFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHFunc())
}

//export myGLibHFunc
func myGLibHFunc(key C.gpointer, value C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &HFuncStruct{
		F_key:   unsafe.Pointer(key),
		F_value: unsafe.Pointer(value),
	}
	fn(args)
}

type HRFuncStruct struct {
	F_key   unsafe.Pointer
	F_value unsafe.Pointer
}

func GetPointer_myHRFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHRFunc())
}

//export myGLibHRFunc
func myGLibHRFunc(key C.gpointer, value C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &HRFuncStruct{
		F_key:   unsafe.Pointer(key),
		F_value: unsafe.Pointer(value),
	}
	fn(args)
}

type HashFuncStruct struct {
	F_key unsafe.Pointer
}

func GetPointer_myHashFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHashFunc())
}

//export myGLibHashFunc
func myGLibHashFunc(key C.gpointer) {
	// TODO: not found user_data
}

// Struct HashTable
type HashTable struct {
	P unsafe.Pointer
}

func HashTableGetType() gi.GType {
	ret := _I.GetGType(26, "HashTable")
	return ret
}

// g_hash_table_add
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableAdd1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(196, "HashTable", "add", 77, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_contains
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableContains1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(197, "HashTable", "contains", 77, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_destroy
//
// [ hash_table ] trans: nothing
//
func HashTableDestroy1(hash_table HashTable) {
	iv, err := _I.Get(198, "HashTable", "destroy", 77, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_insert
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableInsert1(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(199, "HashTable", "insert", 77, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_lookup
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableLookup1(hash_table HashTable, key unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(200, "HashTable", "lookup", 77, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_hash_table_lookup_extended
//
// [ hash_table ] trans: nothing
//
// [ lookup_key ] trans: nothing
//
// [ orig_key ] trans: everything, dir: out
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func HashTableLookupExtended1(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, orig_key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(201, "HashTable", "lookup_extended", 77, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_orig_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_orig_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	orig_key = outArgs[0].Pointer()
	value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_remove
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableRemove1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(202, "HashTable", "remove", 77, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_remove_all
//
// [ hash_table ] trans: nothing
//
func HashTableRemoveAll1(hash_table HashTable) {
	iv, err := _I.Get(203, "HashTable", "remove_all", 77, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_replace
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableReplace1(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(204, "HashTable", "replace", 77, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_size
//
// [ hash_table ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableSize1(hash_table HashTable) (result uint32) {
	iv, err := _I.Get(205, "HashTable", "size", 77, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_hash_table_steal
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableSteal1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(206, "HashTable", "steal", 77, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_steal_all
//
// [ hash_table ] trans: nothing
//
func HashTableStealAll1(hash_table HashTable) {
	iv, err := _I.Get(207, "HashTable", "steal_all", 77, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_steal_extended
//
// [ hash_table ] trans: nothing
//
// [ lookup_key ] trans: nothing
//
// [ stolen_key ] trans: everything, dir: out
//
// [ stolen_value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func HashTableStealExtended1(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, stolen_key unsafe.Pointer, stolen_value unsafe.Pointer) {
	iv, err := _I.Get(208, "HashTable", "steal_extended", 77, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_stolen_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stolen_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_stolen_key, arg_stolen_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	stolen_key = outArgs[0].Pointer()
	stolen_value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_unref
//
// [ hash_table ] trans: nothing
//
func HashTableUnref1(hash_table HashTable) {
	iv, err := _I.Get(209, "HashTable", "unref", 77, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// Struct HashTableIter
type HashTableIter struct {
	P unsafe.Pointer
}

const SizeOfStructHashTableIter = 40

func HashTableIterGetType() gi.GType {
	ret := _I.GetGType(27, "HashTableIter")
	return ret
}

// g_hash_table_iter_init
//
// [ hash_table ] trans: nothing
//
func (v HashTableIter) Init(hash_table HashTable) {
	iv, err := _I.Get(210, "HashTableIter", "init", 78, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_v, arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_iter_next
//
// [ key ] trans: everything, dir: out
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v HashTableIter) Next() (result bool, key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(211, "HashTableIter", "next", 78, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	key = outArgs[0].Pointer()
	value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_iter_remove
//
func (v HashTableIter) Remove() {
	iv, err := _I.Get(212, "HashTableIter", "remove", 78, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_hash_table_iter_replace
//
// [ value ] trans: nothing
//
func (v HashTableIter) Replace(value unsafe.Pointer) {
	iv, err := _I.Get(213, "HashTableIter", "replace", 78, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// g_hash_table_iter_steal
//
func (v HashTableIter) Steal() {
	iv, err := _I.Get(214, "HashTableIter", "steal", 78, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Hmac
type Hmac struct {
	P unsafe.Pointer
}

func HmacGetType() gi.GType {
	ret := _I.GetGType(28, "Hmac")
	return ret
}

// g_hmac_get_digest
//
// [ buffer ] trans: nothing
//
// [ digest_len ] trans: everything, dir: inout
//
func (v Hmac) GetDigest(buffer gi.Uint8Array, digest_len int /*TODO:TYPE*/) {
	iv, err := _I.Get(215, "Hmac", "get_digest", 79, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_buffer}
	iv.Call(args, nil, &outArgs[0])
}

// g_hmac_get_string
//
// [ result ] trans: nothing
//
func (v Hmac) GetString() (result string) {
	iv, err := _I.Get(216, "Hmac", "get_string", 79, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_hmac_unref
//
func (v Hmac) Unref() {
	iv, err := _I.Get(217, "Hmac", "unref", 79, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_hmac_update
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
func (v Hmac) Update(data gi.Uint8Array, length int64) {
	iv, err := _I.Get(218, "Hmac", "update", 79, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_v, arg_data, arg_length}
	iv.Call(args, nil, nil)
}

// Struct Hook
type Hook struct {
	P unsafe.Pointer
}

const SizeOfStructHook = 64

func HookGetType() gi.GType {
	ret := _I.GetGType(29, "Hook")
	return ret
}

// g_hook_compare_ids
//
// [ sibling ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Hook) CompareIds(sibling Hook) (result int32) {
	iv, err := _I.Get(219, "Hook", "compare_ids", 80, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sibling := gi.NewPointerArgument(sibling.P)
	args := []gi.Argument{arg_v, arg_sibling}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_hook_destroy
//
// [ hook_list ] trans: nothing
//
// [ hook_id ] trans: nothing
//
// [ result ] trans: nothing
//
func HookDestroy1(hook_list HookList, hook_id uint64) (result bool) {
	iv, err := _I.Get(220, "Hook", "destroy", 80, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook_id := gi.NewUint64Argument(hook_id)
	args := []gi.Argument{arg_hook_list, arg_hook_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hook_destroy_link
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookDestroyLink1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(221, "Hook", "destroy_link", 80, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_free
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookFree1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(222, "Hook", "free", 80, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_insert_before
//
// [ hook_list ] trans: nothing
//
// [ sibling ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookInsertBefore1(hook_list HookList, sibling Hook, hook Hook) {
	iv, err := _I.Get(223, "Hook", "insert_before", 80, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_sibling := gi.NewPointerArgument(sibling.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_sibling, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_prepend
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookPrepend1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(224, "Hook", "prepend", 80, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_unref
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookUnref1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(225, "Hook", "unref", 80, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

type HookCheckFuncStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myHookCheckFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookCheckFunc())
}

//export myGLibHookCheckFunc
func myGLibHookCheckFunc(data C.gpointer) {
	// TODO: not found user_data
}

type HookCheckMarshallerStruct struct {
	F_hook         Hook
	F_marshal_data unsafe.Pointer
}

func GetPointer_myHookCheckMarshaller() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookCheckMarshaller())
}

//export myGLibHookCheckMarshaller
func myGLibHookCheckMarshaller(hook *C.GHook, marshal_data C.gpointer) {
	// TODO: not found user_data
}

type HookCompareFuncStruct struct {
	F_new_hook Hook
	F_sibling  Hook
}

func GetPointer_myHookCompareFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookCompareFunc())
}

//export myGLibHookCompareFunc
func myGLibHookCompareFunc(new_hook *C.GHook, sibling *C.GHook) {
	// TODO: not found user_data
}

type HookFinalizeFuncStruct struct {
	F_hook_list HookList
	F_hook      Hook
}

func GetPointer_myHookFinalizeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookFinalizeFunc())
}

//export myGLibHookFinalizeFunc
func myGLibHookFinalizeFunc(hook_list *C.GHookList, hook *C.GHook) {
	// TODO: not found user_data
}

type HookFindFuncStruct struct {
	F_hook Hook
	F_data unsafe.Pointer
}

func GetPointer_myHookFindFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookFindFunc())
}

//export myGLibHookFindFunc
func myGLibHookFindFunc(hook *C.GHook, data C.gpointer) {
	// TODO: not found user_data
}

// Flags HookFlagMask
type HookFlagMaskFlags int

const (
	HookFlagMaskActive HookFlagMaskFlags = 1
	HookFlagMaskInCall HookFlagMaskFlags = 2
	HookFlagMaskMask   HookFlagMaskFlags = 15
)

func HookFlagMaskGetType() gi.GType {
	ret := _I.GetGType(30, "HookFlagMask")
	return ret
}

type HookFuncStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myHookFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookFunc())
}

//export myGLibHookFunc
func myGLibHookFunc(data C.gpointer) {
	// TODO: not found user_data
}

// Struct HookList
type HookList struct {
	P unsafe.Pointer
}

const SizeOfStructHookList = 56

func HookListGetType() gi.GType {
	ret := _I.GetGType(31, "HookList")
	return ret
}

// g_hook_list_clear
//
func (v HookList) Clear() {
	iv, err := _I.Get(226, "HookList", "clear", 88, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_hook_list_init
//
// [ hook_size ] trans: nothing
//
func (v HookList) Init(hook_size uint32) {
	iv, err := _I.Get(227, "HookList", "init", 88, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hook_size := gi.NewUint32Argument(hook_size)
	args := []gi.Argument{arg_v, arg_hook_size}
	iv.Call(args, nil, nil)
}

// g_hook_list_invoke
//
// [ may_recurse ] trans: nothing
//
func (v HookList) Invoke(may_recurse bool) {
	iv, err := _I.Get(228, "HookList", "invoke", 88, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_may_recurse := gi.NewBoolArgument(may_recurse)
	args := []gi.Argument{arg_v, arg_may_recurse}
	iv.Call(args, nil, nil)
}

// g_hook_list_invoke_check
//
// [ may_recurse ] trans: nothing
//
func (v HookList) InvokeCheck(may_recurse bool) {
	iv, err := _I.Get(229, "HookList", "invoke_check", 88, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_may_recurse := gi.NewBoolArgument(may_recurse)
	args := []gi.Argument{arg_v, arg_may_recurse}
	iv.Call(args, nil, nil)
}

type HookMarshallerStruct struct {
	F_hook         Hook
	F_marshal_data unsafe.Pointer
}

func GetPointer_myHookMarshaller() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibHookMarshaller())
}

//export myGLibHookMarshaller
func myGLibHookMarshaller(hook *C.GHook, marshal_data C.gpointer) {
	// TODO: not found user_data
}

// Struct IOChannel
type IOChannel struct {
	P unsafe.Pointer
}

const SizeOfStructIOChannel = 136

func IOChannelGetType() gi.GType {
	ret := _I.GetGType(32, "IOChannel")
	return ret
}

// g_io_channel_new_file
//
// [ filename ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ result ] trans: everything
//
func NewIOChannelFile(filename string, mode string) (result IOChannel, err error) {
	iv, err := _I.Get(230, "IOChannel", "new_file", 92, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	c_mode := gi.CString(mode)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_mode := gi.NewStringArgument(c_mode)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_mode, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	gi.Free(c_mode)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_io_channel_unix_new
//
// [ fd ] trans: nothing
//
// [ result ] trans: everything
//
func IOChannelUnixNew(fd int32) (result IOChannel) {
	iv, err := _I.Get(231, "IOChannel", "unix_new", 92, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fd := gi.NewInt32Argument(fd)
	args := []gi.Argument{arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_io_channel_close
//
func (v IOChannel) Close() {
	iv, err := _I.Get(232, "IOChannel", "close", 92, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_io_channel_flush
//
// [ result ] trans: nothing
//
func (v IOChannel) Flush() (result IOStatusEnum, err error) {
	iv, err := _I.Get(233, "IOChannel", "flush", 92, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_get_buffer_condition
//
// [ result ] trans: nothing
//
func (v IOChannel) GetBufferCondition() (result IOConditionFlags) {
	iv, err := _I.Get(234, "IOChannel", "get_buffer_condition", 92, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOConditionFlags(ret.Int())
	return
}

// g_io_channel_get_buffer_size
//
// [ result ] trans: nothing
//
func (v IOChannel) GetBufferSize() (result uint64) {
	iv, err := _I.Get(235, "IOChannel", "get_buffer_size", 92, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_io_channel_get_buffered
//
// [ result ] trans: nothing
//
func (v IOChannel) GetBuffered() (result bool) {
	iv, err := _I.Get(236, "IOChannel", "get_buffered", 92, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_io_channel_get_close_on_unref
//
// [ result ] trans: nothing
//
func (v IOChannel) GetCloseOnUnref() (result bool) {
	iv, err := _I.Get(237, "IOChannel", "get_close_on_unref", 92, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_io_channel_get_encoding
//
// [ result ] trans: nothing
//
func (v IOChannel) GetEncoding() (result string) {
	iv, err := _I.Get(238, "IOChannel", "get_encoding", 92, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_io_channel_get_flags
//
// [ result ] trans: nothing
//
func (v IOChannel) GetFlags() (result IOFlags) {
	iv, err := _I.Get(239, "IOChannel", "get_flags", 92, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOFlags(ret.Int())
	return
}

// g_io_channel_get_line_term
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) GetLineTerm(length int32) (result string) {
	iv, err := _I.Get(240, "IOChannel", "get_line_term", 92, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_io_channel_init
//
func (v IOChannel) Init() {
	iv, err := _I.Get(241, "IOChannel", "init", 92, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_io_channel_read
//
// [ buf ] trans: nothing
//
// [ count ] trans: nothing
//
// [ bytes_read ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) Read(buf string, count uint64, bytes_read uint64) (result IOErrorEnum) {
	iv, err := _I.Get(242, "IOChannel", "read", 92, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buf := gi.CString(buf)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewStringArgument(c_buf)
	arg_count := gi.NewUint64Argument(count)
	arg_bytes_read := gi.NewUint64Argument(bytes_read)
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_read}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buf)
	result = IOErrorEnum(ret.Int())
	return
}

// g_io_channel_read_chars
//
// [ buf ] trans: nothing, dir: out
//
// [ count ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v IOChannel) ReadChars(buf gi.Uint8Array, count uint64) (result IOStatusEnum, bytes_read uint64, err error) {
	iv, err := _I.Get(243, "IOChannel", "read_chars", 92, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_count := gi.NewUint64Argument(count)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_read, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	bytes_read = outArgs[0].Uint64()
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_read_line
//
// [ str_return ] trans: everything, dir: out
//
// [ length ] trans: everything, dir: out
//
// [ terminator_pos ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v IOChannel) ReadLine() (result IOStatusEnum, str_return string, length uint64, terminator_pos uint64, err error) {
	iv, err := _I.Get(244, "IOChannel", "read_line", 92, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str_return := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_terminator_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_str_return, arg_length, arg_terminator_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[3].Pointer())
	str_return = outArgs[0].String().Take()
	length = outArgs[1].Uint64()
	terminator_pos = outArgs[2].Uint64()
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_read_line_string
//
// [ buffer ] trans: nothing
//
// [ terminator_pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) ReadLineString(buffer String, terminator_pos uint64) (result IOStatusEnum, err error) {
	iv, err := _I.Get(245, "IOChannel", "read_line_string", 92, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_terminator_pos := gi.NewUint64Argument(terminator_pos)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_buffer, arg_terminator_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_read_to_end
//
// [ str_return ] trans: everything, dir: out
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v IOChannel) ReadToEnd() (result IOStatusEnum, str_return gi.Uint8Array, err error) {
	iv, err := _I.Get(246, "IOChannel", "read_to_end", 92, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str_return := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_str_return, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	err = gi.ToError(outArgs[2].Pointer())
	str_return.P = outArgs[0].Pointer()
	length = outArgs[1].Uint64()
	result = IOStatusEnum(ret.Int())
	str_return.Len = int(length)
	return
}

// g_io_channel_read_unichar
//
// [ thechar ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v IOChannel) ReadUnichar() (result IOStatusEnum, thechar rune, err error) {
	iv, err := _I.Get(247, "IOChannel", "read_unichar", 92, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_thechar := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_thechar, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	thechar = rune(outArgs[0].Uint32())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_ref
//
// [ result ] trans: everything
//
func (v IOChannel) Ref() (result IOChannel) {
	iv, err := _I.Get(248, "IOChannel", "ref", 92, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_io_channel_seek
//
// [ offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) Seek(offset int64, type1 SeekTypeEnum) (result IOErrorEnum) {
	iv, err := _I.Get(249, "IOChannel", "seek", 92, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_offset, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOErrorEnum(ret.Int())
	return
}

// g_io_channel_seek_position
//
// [ offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) SeekPosition(offset int64, type1 SeekTypeEnum) (result IOStatusEnum, err error) {
	iv, err := _I.Get(250, "IOChannel", "seek_position", 92, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_set_buffer_size
//
// [ size ] trans: nothing
//
func (v IOChannel) SetBufferSize(size uint64) {
	iv, err := _I.Get(251, "IOChannel", "set_buffer_size", 92, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_size}
	iv.Call(args, nil, nil)
}

// g_io_channel_set_buffered
//
// [ buffered ] trans: nothing
//
func (v IOChannel) SetBuffered(buffered bool) {
	iv, err := _I.Get(252, "IOChannel", "set_buffered", 92, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffered := gi.NewBoolArgument(buffered)
	args := []gi.Argument{arg_v, arg_buffered}
	iv.Call(args, nil, nil)
}

// g_io_channel_set_close_on_unref
//
// [ do_close ] trans: nothing
//
func (v IOChannel) SetCloseOnUnref(do_close bool) {
	iv, err := _I.Get(253, "IOChannel", "set_close_on_unref", 92, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_do_close := gi.NewBoolArgument(do_close)
	args := []gi.Argument{arg_v, arg_do_close}
	iv.Call(args, nil, nil)
}

// g_io_channel_set_encoding
//
// [ encoding ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) SetEncoding(encoding string) (result IOStatusEnum, err error) {
	iv, err := _I.Get(254, "IOChannel", "set_encoding", 92, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_encoding := gi.CString(encoding)
	arg_v := gi.NewPointerArgument(v.P)
	arg_encoding := gi.NewStringArgument(c_encoding)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_encoding, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_encoding)
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_set_flags
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) SetFlags(flags IOFlags) (result IOStatusEnum, err error) {
	iv, err := _I.Get(255, "IOChannel", "set_flags", 92, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_set_line_term
//
// [ line_term ] trans: nothing
//
// [ length ] trans: nothing
//
func (v IOChannel) SetLineTerm(line_term string, length int32) {
	iv, err := _I.Get(256, "IOChannel", "set_line_term", 92, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_line_term := gi.CString(line_term)
	arg_v := gi.NewPointerArgument(v.P)
	arg_line_term := gi.NewStringArgument(c_line_term)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_line_term, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_line_term)
}

// g_io_channel_shutdown
//
// [ flush ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) Shutdown(flush bool) (result IOStatusEnum, err error) {
	iv, err := _I.Get(257, "IOChannel", "shutdown", 92, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flush := gi.NewBoolArgument(flush)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flush, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_unix_get_fd
//
// [ result ] trans: nothing
//
func (v IOChannel) UnixGetFd() (result int32) {
	iv, err := _I.Get(258, "IOChannel", "unix_get_fd", 92, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_io_channel_unref
//
func (v IOChannel) Unref() {
	iv, err := _I.Get(259, "IOChannel", "unref", 92, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_io_channel_write
//
// [ buf ] trans: nothing
//
// [ count ] trans: nothing
//
// [ bytes_written ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) Write(buf string, count uint64, bytes_written uint64) (result IOErrorEnum) {
	iv, err := _I.Get(260, "IOChannel", "write", 92, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buf := gi.CString(buf)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewStringArgument(c_buf)
	arg_count := gi.NewUint64Argument(count)
	arg_bytes_written := gi.NewUint64Argument(bytes_written)
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_written}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buf)
	result = IOErrorEnum(ret.Int())
	return
}

// g_io_channel_write_chars
//
// [ buf ] trans: nothing
//
// [ count ] trans: nothing
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v IOChannel) WriteChars(buf gi.Uint8Array, count int64) (result IOStatusEnum, bytes_written uint64, err error) {
	iv, err := _I.Get(261, "IOChannel", "write_chars", 92, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_count := gi.NewInt64Argument(count)
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	bytes_written = outArgs[0].Uint64()
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_write_unichar
//
// [ thechar ] trans: nothing
//
// [ result ] trans: nothing
//
func (v IOChannel) WriteUnichar(thechar rune) (result IOStatusEnum, err error) {
	iv, err := _I.Get(262, "IOChannel", "write_unichar", 92, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_thechar := gi.NewUint32Argument(uint32(thechar))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_thechar, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_error_from_errno
//
// [ en ] trans: nothing
//
// [ result ] trans: nothing
//
func IOChannelErrorFromErrno1(en int32) (result IOChannelErrorEnum) {
	iv, err := _I.Get(263, "IOChannel", "error_from_errno", 92, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_en := gi.NewInt32Argument(en)
	args := []gi.Argument{arg_en}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOChannelErrorEnum(ret.Int())
	return
}

// Enum IOChannelError
type IOChannelErrorEnum int

const (
	IOChannelErrorFbig     IOChannelErrorEnum = 0
	IOChannelErrorInval    IOChannelErrorEnum = 1
	IOChannelErrorIo       IOChannelErrorEnum = 2
	IOChannelErrorIsdir    IOChannelErrorEnum = 3
	IOChannelErrorNospc    IOChannelErrorEnum = 4
	IOChannelErrorNxio     IOChannelErrorEnum = 5
	IOChannelErrorOverflow IOChannelErrorEnum = 6
	IOChannelErrorPipe     IOChannelErrorEnum = 7
	IOChannelErrorFailed   IOChannelErrorEnum = 8
)

func IOChannelErrorGetType() gi.GType {
	ret := _I.GetGType(33, "IOChannelError")
	return ret
}

// Flags IOCondition
type IOConditionFlags int

const (
	IOConditionIn   IOConditionFlags = 1
	IOConditionOut  IOConditionFlags = 4
	IOConditionPri  IOConditionFlags = 2
	IOConditionErr  IOConditionFlags = 8
	IOConditionHup  IOConditionFlags = 16
	IOConditionNval IOConditionFlags = 32
)

func IOConditionGetType() gi.GType {
	ret := _I.GetGType(34, "IOCondition")
	return ret
}

// Enum IOError
type IOErrorEnum int

const (
	IOErrorNone    IOErrorEnum = 0
	IOErrorAgain   IOErrorEnum = 1
	IOErrorInval   IOErrorEnum = 2
	IOErrorUnknown IOErrorEnum = 3
)

func IOErrorGetType() gi.GType {
	ret := _I.GetGType(35, "IOError")
	return ret
}

// Flags IOFlags
type IOFlags int

const (
	IOFlagsAppend      IOFlags = 1
	IOFlagsNonblock    IOFlags = 2
	IOFlagsIsReadable  IOFlags = 4
	IOFlagsIsWritable  IOFlags = 8
	IOFlagsIsWriteable IOFlags = 8
	IOFlagsIsSeekable  IOFlags = 16
	IOFlagsMask        IOFlags = 31
	IOFlagsGetMask     IOFlags = 31
	IOFlagsSetMask     IOFlags = 3
)

func IOFlagsGetType() gi.GType {
	ret := _I.GetGType(36, "IOFlags")
	return ret
}

type IOFuncStruct struct {
	F_source    IOChannel
	F_condition IOConditionFlags
	F_data      unsafe.Pointer
}

func GetPointer_myIOFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibIOFunc())
}

//export myGLibIOFunc
func myGLibIOFunc(source *C.GIOChannel, condition C.GIOCondition, data C.gpointer) {
	// TODO: not found user_data
}

// Struct IOFuncs
type IOFuncs struct {
	P unsafe.Pointer
}

const SizeOfStructIOFuncs = 64

func IOFuncsGetType() gi.GType {
	ret := _I.GetGType(37, "IOFuncs")
	return ret
}

// Enum IOStatus
type IOStatusEnum int

const (
	IOStatusError  IOStatusEnum = 0
	IOStatusNormal IOStatusEnum = 1
	IOStatusEof    IOStatusEnum = 2
	IOStatusAgain  IOStatusEnum = 3
)

func IOStatusGetType() gi.GType {
	ret := _I.GetGType(38, "IOStatus")
	return ret
}

// Struct KeyFile
type KeyFile struct {
	P unsafe.Pointer
}

func KeyFileGetType() gi.GType {
	ret := _I.GetGType(39, "KeyFile")
	return ret
}

// g_key_file_new
//
// [ result ] trans: everything
//
func NewKeyFile() (result KeyFile) {
	iv, err := _I.Get(265, "KeyFile", "new", 129, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_key_file_get_boolean
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) GetBoolean(group_name string, key string) (result bool, err error) {
	iv, err := _I.Get(266, "KeyFile", "get_boolean", 129, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_get_boolean_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v KeyFile) GetBooleanList(group_name string, key string) (result gi.BoolArray, err error) {
	iv, err := _I.Get(267, "KeyFile", "get_boolean_list", 129, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.BoolArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_key_file_get_comment
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: everything
//
func (v KeyFile) GetComment(group_name string, key string) (result string, err error) {
	iv, err := _I.Get(268, "KeyFile", "get_comment", 129, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_get_double
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) GetDouble(group_name string, key string) (result float64, err error) {
	iv, err := _I.Get(269, "KeyFile", "get_double", 129, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// g_key_file_get_double_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v KeyFile) GetDoubleList(group_name string, key string) (result gi.DoubleArray, err error) {
	iv, err := _I.Get(270, "KeyFile", "get_double_list", 129, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.DoubleArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_key_file_get_groups
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v KeyFile) GetGroups() (result gi.CStrArray, length uint64) {
	iv, err := _I.Get(271, "KeyFile", "get_groups", 129, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_int64
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) GetInt64(group_name string, key string) (result int64, err error) {
	iv, err := _I.Get(272, "KeyFile", "get_int64", 129, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_key_file_get_integer
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) GetInteger(group_name string, key string) (result int32, err error) {
	iv, err := _I.Get(273, "KeyFile", "get_integer", 129, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// g_key_file_get_integer_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v KeyFile) GetIntegerList(group_name string, key string) (result gi.Int32Array, err error) {
	iv, err := _I.Get(274, "KeyFile", "get_integer_list", 129, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.Int32Array{P: ret.Pointer(), Len: int(length)}
	return
}

// g_key_file_get_keys
//
// [ group_name ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v KeyFile) GetKeys(group_name string) (result gi.CStrArray, length uint64, err error) {
	iv, err := _I.Get(275, "KeyFile", "get_keys", 129, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_locale_for_key
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ locale ] trans: nothing
//
// [ result ] trans: everything
//
func (v KeyFile) GetLocaleForKey(group_name string, key string, locale string) (result string) {
	iv, err := _I.Get(276, "KeyFile", "get_locale_for_key", 129, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	result = ret.String().Take()
	return
}

// g_key_file_get_locale_string
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ locale ] trans: nothing
//
// [ result ] trans: everything
//
func (v KeyFile) GetLocaleString(group_name string, key string, locale string) (result string, err error) {
	iv, err := _I.Get(277, "KeyFile", "get_locale_string", 129, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_get_locale_string_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ locale ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v KeyFile) GetLocaleStringList(group_name string, key string, locale string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(278, "KeyFile", "get_locale_string_list", 129, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_start_group
//
// [ result ] trans: everything
//
func (v KeyFile) GetStartGroup() (result string) {
	iv, err := _I.Get(279, "KeyFile", "get_start_group", 129, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_key_file_get_string
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: everything
//
func (v KeyFile) GetString(group_name string, key string) (result string, err error) {
	iv, err := _I.Get(280, "KeyFile", "get_string", 129, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_get_string_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v KeyFile) GetStringList(group_name string, key string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(281, "KeyFile", "get_string_list", 129, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_uint64
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) GetUint64(group_name string, key string) (result uint64, err error) {
	iv, err := _I.Get(282, "KeyFile", "get_uint64", 129, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Uint64()
	return
}

// g_key_file_get_value
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: everything
//
func (v KeyFile) GetValue(group_name string, key string) (result string, err error) {
	iv, err := _I.Get(283, "KeyFile", "get_value", 129, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_has_group
//
// [ group_name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) HasGroup(group_name string) (result bool) {
	iv, err := _I.Get(284, "KeyFile", "has_group", 129, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	args := []gi.Argument{arg_v, arg_group_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_group_name)
	result = ret.Bool()
	return
}

// g_key_file_load_from_bytes
//
// [ bytes ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) LoadFromBytes(bytes Bytes, flags KeyFileFlags) (result bool, err error) {
	iv, err := _I.Get(285, "KeyFile", "load_from_bytes", 129, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes := gi.NewPointerArgument(bytes.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_bytes, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_load_from_data
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) LoadFromData(data string, length uint64, flags KeyFileFlags) (result bool, err error) {
	iv, err := _I.Get(286, "KeyFile", "load_from_data", 129, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_data := gi.CString(data)
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewStringArgument(c_data)
	arg_length := gi.NewUint64Argument(length)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data, arg_length, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_data)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_load_from_data_dirs
//
// [ file ] trans: nothing
//
// [ full_path ] trans: everything, dir: out
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (result bool, full_path string, err error) {
	iv, err := _I.Get(287, "KeyFile", "load_from_data_dirs", 129, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_full_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_file, arg_full_path, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[1].Pointer())
	full_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_key_file_load_from_dirs
//
// [ file ] trans: nothing
//
// [ search_dirs ] trans: nothing
//
// [ full_path ] trans: everything, dir: out
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) LoadFromDirs(file string, search_dirs gi.CStrArray, flags KeyFileFlags) (result bool, full_path string, err error) {
	iv, err := _I.Get(288, "KeyFile", "load_from_dirs", 129, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_search_dirs := gi.NewPointerArgument(search_dirs.P)
	arg_full_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_file, arg_search_dirs, arg_full_path, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[1].Pointer())
	full_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_key_file_load_from_file
//
// [ file ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) LoadFromFile(file string, flags KeyFileFlags) (result bool, err error) {
	iv, err := _I.Get(289, "KeyFile", "load_from_file", 129, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_file, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_remove_comment
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) RemoveComment(group_name string, key string) (result bool, err error) {
	iv, err := _I.Get(290, "KeyFile", "remove_comment", 129, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_remove_group
//
// [ group_name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) RemoveGroup(group_name string) (result bool, err error) {
	iv, err := _I.Get(291, "KeyFile", "remove_group", 129, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_remove_key
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) RemoveKey(group_name string, key string) (result bool, err error) {
	iv, err := _I.Get(292, "KeyFile", "remove_key", 129, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_save_to_file
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) SaveToFile(filename string) (result bool, err error) {
	iv, err := _I.Get(293, "KeyFile", "save_to_file", 129, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_set_boolean
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v KeyFile) SetBoolean(group_name string, key string, value bool) {
	iv, err := _I.Get(294, "KeyFile", "set_boolean", 129, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewBoolArgument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_boolean_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ list ] trans: nothing
//
// [ length ] trans: nothing
//
func (v KeyFile) SetBooleanList(group_name string, key string, list gi.BoolArray, length uint64) {
	iv, err := _I.Get(295, "KeyFile", "set_boolean_list", 129, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_comment
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ comment ] trans: nothing
//
// [ result ] trans: nothing
//
func (v KeyFile) SetComment(group_name string, key string, comment string) (result bool, err error) {
	iv, err := _I.Get(296, "KeyFile", "set_comment", 129, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_comment := gi.CString(comment)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_comment := gi.NewStringArgument(c_comment)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_comment, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_comment)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_set_double
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v KeyFile) SetDouble(group_name string, key string, value float64) {
	iv, err := _I.Get(297, "KeyFile", "set_double", 129, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewDoubleArgument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_double_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ list ] trans: nothing
//
// [ length ] trans: nothing
//
func (v KeyFile) SetDoubleList(group_name string, key string, list gi.DoubleArray, length uint64) {
	iv, err := _I.Get(298, "KeyFile", "set_double_list", 129, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_int64
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v KeyFile) SetInt64(group_name string, key string, value int64) {
	iv, err := _I.Get(299, "KeyFile", "set_int64", 129, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewInt64Argument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_integer
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v KeyFile) SetInteger(group_name string, key string, value int32) {
	iv, err := _I.Get(300, "KeyFile", "set_integer", 129, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_integer_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ list ] trans: nothing
//
// [ length ] trans: nothing
//
func (v KeyFile) SetIntegerList(group_name string, key string, list gi.Int32Array, length uint64) {
	iv, err := _I.Get(301, "KeyFile", "set_integer_list", 129, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_list_separator
//
// [ separator ] trans: nothing
//
func (v KeyFile) SetListSeparator(separator int8) {
	iv, err := _I.Get(302, "KeyFile", "set_list_separator", 129, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_separator := gi.NewInt8Argument(separator)
	args := []gi.Argument{arg_v, arg_separator}
	iv.Call(args, nil, nil)
}

// g_key_file_set_locale_string
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ locale ] trans: nothing
//
// [ string ] trans: nothing
//
func (v KeyFile) SetLocaleString(group_name string, key string, locale string, string string) {
	iv, err := _I.Get(303, "KeyFile", "set_locale_string", 129, 38, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_string}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	gi.Free(c_string)
}

// g_key_file_set_locale_string_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ locale ] trans: nothing
//
// [ list ] trans: nothing
//
// [ length ] trans: nothing
//
func (v KeyFile) SetLocaleStringList(group_name string, key string, locale string, list gi.CStrArray, length uint64) {
	iv, err := _I.Get(304, "KeyFile", "set_locale_string_list", 129, 39, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
}

// g_key_file_set_string
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ string ] trans: nothing
//
func (v KeyFile) SetString(group_name string, key string, string string) {
	iv, err := _I.Get(305, "KeyFile", "set_string", 129, 40, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_string}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_string)
}

// g_key_file_set_string_list
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ list ] trans: nothing
//
// [ length ] trans: nothing
//
func (v KeyFile) SetStringList(group_name string, key string, list gi.CStrArray, length uint64) {
	iv, err := _I.Get(306, "KeyFile", "set_string_list", 129, 41, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_uint64
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v KeyFile) SetUint64(group_name string, key string, value uint64) {
	iv, err := _I.Get(307, "KeyFile", "set_uint64", 129, 42, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewUint64Argument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_value
//
// [ group_name ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v KeyFile) SetValue(group_name string, key string, value string) {
	iv, err := _I.Get(308, "KeyFile", "set_value", 129, 43, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_value)
}

// g_key_file_to_data
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v KeyFile) ToData() (result string, length uint64, err error) {
	iv, err := _I.Get(309, "KeyFile", "to_data", 129, 44, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = ret.String().Take()
	return
}

// g_key_file_unref
//
func (v KeyFile) Unref() {
	iv, err := _I.Get(310, "KeyFile", "unref", 129, 45, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum KeyFileError
type KeyFileErrorEnum int

const (
	KeyFileErrorUnknownEncoding KeyFileErrorEnum = 0
	KeyFileErrorParse           KeyFileErrorEnum = 1
	KeyFileErrorNotFound        KeyFileErrorEnum = 2
	KeyFileErrorKeyNotFound     KeyFileErrorEnum = 3
	KeyFileErrorGroupNotFound   KeyFileErrorEnum = 4
	KeyFileErrorInvalidValue    KeyFileErrorEnum = 5
)

func KeyFileErrorGetType() gi.GType {
	ret := _I.GetGType(40, "KeyFileError")
	return ret
}

// Flags KeyFileFlags
type KeyFileFlags int

const (
	KeyFileFlagsNone             KeyFileFlags = 0
	KeyFileFlagsKeepComments     KeyFileFlags = 1
	KeyFileFlagsKeepTranslations KeyFileFlags = 2
)

func KeyFileFlagsGetType() gi.GType {
	ret := _I.GetGType(41, "KeyFileFlags")
	return ret
}

// Struct List
type List struct {
	P unsafe.Pointer
}

const SizeOfStructList = 24

func ListGetType() gi.GType {
	ret := _I.GetGType(42, "List")
	return ret
}

// Struct LogField
type LogField struct {
	P unsafe.Pointer
}

const SizeOfStructLogField = 24

func LogFieldGetType() gi.GType {
	ret := _I.GetGType(43, "LogField")
	return ret
}

type LogFuncStruct struct {
	F_log_domain string
	F_log_level  LogLevelFlags
	F_message    string
}

func GetPointer_myLogFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibLogFunc())
}

//export myGLibLogFunc
func myGLibLogFunc(log_domain *C.gchar, log_level C.GLogLevelFlags, message *C.gchar, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &LogFuncStruct{
		F_log_domain: gi.GoString(unsafe.Pointer(log_domain)),
		F_log_level:  LogLevelFlags(log_level),
		F_message:    gi.GoString(unsafe.Pointer(message)),
	}
	fn(args)
}

// Flags LogLevelFlags
type LogLevelFlags int

const (
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	LogLevelFlagsFlagFatal     LogLevelFlags = 2
	LogLevelFlagsLevelError    LogLevelFlags = 4
	LogLevelFlagsLevelCritical LogLevelFlags = 8
	LogLevelFlagsLevelWarning  LogLevelFlags = 16
	LogLevelFlagsLevelMessage  LogLevelFlags = 32
	LogLevelFlagsLevelInfo     LogLevelFlags = 64
	LogLevelFlagsLevelDebug    LogLevelFlags = 128
	LogLevelFlagsLevelMask     LogLevelFlags = -4
)

func LogLevelFlagsGetType() gi.GType {
	ret := _I.GetGType(44, "LogLevelFlags")
	return ret
}

type LogWriterFuncStruct struct {
	F_log_level LogLevelFlags
	F_fields    unsafe.Pointer
	F_n_fields  uint64
}

func GetPointer_myLogWriterFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibLogWriterFunc())
}

//export myGLibLogWriterFunc
func myGLibLogWriterFunc(log_level C.GLogLevelFlags, fields C.gpointer, n_fields C.guint64, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &LogWriterFuncStruct{
		F_log_level: LogLevelFlags(log_level),
		F_fields:    unsafe.Pointer(fields),
		F_n_fields:  uint64(n_fields),
	}
	fn(args)
}

// Enum LogWriterOutput
type LogWriterOutputEnum int

const (
	LogWriterOutputHandled   LogWriterOutputEnum = 1
	LogWriterOutputUnhandled LogWriterOutputEnum = 0
)

func LogWriterOutputGetType() gi.GType {
	ret := _I.GetGType(45, "LogWriterOutput")
	return ret
}

// Struct MainContext
type MainContext struct {
	P unsafe.Pointer
}

func MainContextGetType() gi.GType {
	ret := _I.GetGType(46, "MainContext")
	return ret
}

// g_main_context_new
//
// [ result ] trans: everything
//
func NewMainContext() (result MainContext) {
	iv, err := _I.Get(312, "MainContext", "new", 161, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_acquire
//
// [ result ] trans: nothing
//
func (v MainContext) Acquire() (result bool) {
	iv, err := _I.Get(313, "MainContext", "acquire", 161, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_add_poll
//
// [ fd ] trans: nothing
//
// [ priority ] trans: nothing
//
func (v MainContext) AddPoll(fd PollFD, priority int32) {
	iv, err := _I.Get(314, "MainContext", "add_poll", 161, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_fd, arg_priority}
	iv.Call(args, nil, nil)
}

// g_main_context_check
//
// [ max_priority ] trans: nothing
//
// [ fds ] trans: nothing
//
// [ n_fds ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) Check(max_priority int32, fds unsafe.Pointer, n_fds int32) (result bool) {
	iv, err := _I.Get(315, "MainContext", "check", 161, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_priority := gi.NewInt32Argument(max_priority)
	arg_fds := gi.NewPointerArgument(fds)
	arg_n_fds := gi.NewInt32Argument(n_fds)
	args := []gi.Argument{arg_v, arg_max_priority, arg_fds, arg_n_fds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_dispatch
//
func (v MainContext) Dispatch() {
	iv, err := _I.Get(316, "MainContext", "dispatch", 161, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_find_source_by_funcs_user_data
//
// [ funcs ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) FindSourceByFuncsUserData(funcs SourceFuncs, user_data unsafe.Pointer) (result Source) {
	iv, err := _I.Get(317, "MainContext", "find_source_by_funcs_user_data", 161, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_funcs := gi.NewPointerArgument(funcs.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_funcs, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_find_source_by_id
//
// [ source_id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) FindSourceById(source_id uint32) (result Source) {
	iv, err := _I.Get(318, "MainContext", "find_source_by_id", 161, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_source_id := gi.NewUint32Argument(source_id)
	args := []gi.Argument{arg_v, arg_source_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_find_source_by_user_data
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) FindSourceByUserData(user_data unsafe.Pointer) (result Source) {
	iv, err := _I.Get(319, "MainContext", "find_source_by_user_data", 161, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_invoke_full
//
// [ priority ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v MainContext) InvokeFull(priority int32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(320, "MainContext", "invoke_full", 161, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_priority, arg_function, arg_data, arg_notify}
	iv.Call(args, nil, nil)
}

// g_main_context_is_owner
//
// [ result ] trans: nothing
//
func (v MainContext) IsOwner() (result bool) {
	iv, err := _I.Get(321, "MainContext", "is_owner", 161, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_iteration
//
// [ may_block ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) Iteration(may_block bool) (result bool) {
	iv, err := _I.Get(322, "MainContext", "iteration", 161, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_may_block := gi.NewBoolArgument(may_block)
	args := []gi.Argument{arg_v, arg_may_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_pending
//
// [ result ] trans: nothing
//
func (v MainContext) Pending() (result bool) {
	iv, err := _I.Get(323, "MainContext", "pending", 161, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_pop_thread_default
//
func (v MainContext) PopThreadDefault() {
	iv, err := _I.Get(324, "MainContext", "pop_thread_default", 161, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_prepare
//
// [ priority ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) Prepare(priority int32) (result bool) {
	iv, err := _I.Get(325, "MainContext", "prepare", 161, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_priority}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_push_thread_default
//
func (v MainContext) PushThreadDefault() {
	iv, err := _I.Get(326, "MainContext", "push_thread_default", 161, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_query
//
// [ max_priority ] trans: nothing
//
// [ timeout_ ] trans: everything, dir: out
//
// [ fds ] trans: nothing, dir: out
//
// [ n_fds ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) Query(max_priority int32, fds unsafe.Pointer, n_fds int32) (result int32, timeout_ int32) {
	iv, err := _I.Get(327, "MainContext", "query", 161, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_priority := gi.NewInt32Argument(max_priority)
	arg_timeout_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_fds := gi.NewPointerArgument(fds)
	arg_n_fds := gi.NewInt32Argument(n_fds)
	args := []gi.Argument{arg_v, arg_max_priority, arg_timeout_, arg_fds, arg_n_fds}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	timeout_ = outArgs[0].Int32()
	result = ret.Int32()
	return
}

// g_main_context_ref
//
// [ result ] trans: everything
//
func (v MainContext) Ref() (result MainContext) {
	iv, err := _I.Get(328, "MainContext", "ref", 161, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_release
//
func (v MainContext) Release() {
	iv, err := _I.Get(329, "MainContext", "release", 161, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_remove_poll
//
// [ fd ] trans: nothing
//
func (v MainContext) RemovePoll(fd PollFD) {
	iv, err := _I.Get(330, "MainContext", "remove_poll", 161, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// g_main_context_unref
//
func (v MainContext) Unref() {
	iv, err := _I.Get(331, "MainContext", "unref", 161, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_main_context_wait
//
// [ cond ] trans: nothing
//
// [ mutex ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MainContext) Wait(cond Cond, mutex Mutex) (result bool) {
	iv, err := _I.Get(332, "MainContext", "wait", 161, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cond := gi.NewPointerArgument(cond.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	args := []gi.Argument{arg_v, arg_cond, arg_mutex}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_wakeup
//
func (v MainContext) Wakeup() {
	iv, err := _I.Get(333, "MainContext", "wakeup", 161, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct MainLoop
type MainLoop struct {
	P unsafe.Pointer
}

func MainLoopGetType() gi.GType {
	ret := _I.GetGType(47, "MainLoop")
	return ret
}

// g_main_loop_new
//
// [ context ] trans: nothing
//
// [ is_running ] trans: nothing
//
// [ result ] trans: everything
//
func NewMainLoop(context MainContext, is_running bool) (result MainLoop) {
	iv, err := _I.Get(337, "MainLoop", "new", 162, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_context := gi.NewPointerArgument(context.P)
	arg_is_running := gi.NewBoolArgument(is_running)
	args := []gi.Argument{arg_context, arg_is_running}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_loop_get_context
//
// [ result ] trans: nothing
//
func (v MainLoop) GetContext() (result MainContext) {
	iv, err := _I.Get(338, "MainLoop", "get_context", 162, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_loop_is_running
//
// [ result ] trans: nothing
//
func (v MainLoop) IsRunning() (result bool) {
	iv, err := _I.Get(339, "MainLoop", "is_running", 162, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_loop_quit
//
func (v MainLoop) Quit() {
	iv, err := _I.Get(340, "MainLoop", "quit", 162, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_loop_ref
//
// [ result ] trans: everything
//
func (v MainLoop) Ref() (result MainLoop) {
	iv, err := _I.Get(341, "MainLoop", "ref", 162, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_loop_run
//
func (v MainLoop) Run() {
	iv, err := _I.Get(342, "MainLoop", "run", 162, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_loop_unref
//
func (v MainLoop) Unref() {
	iv, err := _I.Get(343, "MainLoop", "unref", 162, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct MappedFile
type MappedFile struct {
	P unsafe.Pointer
}

func MappedFileGetType() gi.GType {
	ret := _I.GetGType(48, "MappedFile")
	return ret
}

// g_mapped_file_new
//
// [ filename ] trans: nothing
//
// [ writable ] trans: nothing
//
// [ result ] trans: everything
//
func NewMappedFile(filename string, writable bool) (result MappedFile, err error) {
	iv, err := _I.Get(344, "MappedFile", "new", 163, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_writable := gi.NewBoolArgument(writable)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_writable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_mapped_file_new_from_fd
//
// [ fd ] trans: nothing
//
// [ writable ] trans: nothing
//
// [ result ] trans: everything
//
func NewMappedFileFromFd(fd int32, writable bool) (result MappedFile, err error) {
	iv, err := _I.Get(345, "MappedFile", "new_from_fd", 163, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fd := gi.NewInt32Argument(fd)
	arg_writable := gi.NewBoolArgument(writable)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_writable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_mapped_file_free
//
func (v MappedFile) Free() {
	iv, err := _I.Get(346, "MappedFile", "free", 163, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mapped_file_get_bytes
//
// [ result ] trans: everything
//
func (v MappedFile) GetBytes() (result Bytes) {
	iv, err := _I.Get(347, "MappedFile", "get_bytes", 163, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_mapped_file_get_contents
//
// [ result ] trans: everything
//
func (v MappedFile) GetContents() (result string) {
	iv, err := _I.Get(348, "MappedFile", "get_contents", 163, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_mapped_file_get_length
//
// [ result ] trans: nothing
//
func (v MappedFile) GetLength() (result uint64) {
	iv, err := _I.Get(349, "MappedFile", "get_length", 163, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_mapped_file_ref
//
// [ result ] trans: everything
//
func (v MappedFile) Ref() (result MappedFile) {
	iv, err := _I.Get(350, "MappedFile", "ref", 163, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_mapped_file_unref
//
func (v MappedFile) Unref() {
	iv, err := _I.Get(351, "MappedFile", "unref", 163, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags MarkupCollectType
type MarkupCollectTypeFlags int

const (
	MarkupCollectTypeInvalid  MarkupCollectTypeFlags = 0
	MarkupCollectTypeString   MarkupCollectTypeFlags = 1
	MarkupCollectTypeStrdup   MarkupCollectTypeFlags = 2
	MarkupCollectTypeBoolean  MarkupCollectTypeFlags = 3
	MarkupCollectTypeTristate MarkupCollectTypeFlags = 4
	MarkupCollectTypeOptional MarkupCollectTypeFlags = 65536
)

func MarkupCollectTypeGetType() gi.GType {
	ret := _I.GetGType(49, "MarkupCollectType")
	return ret
}

// Enum MarkupError
type MarkupErrorEnum int

const (
	MarkupErrorBadUtf8          MarkupErrorEnum = 0
	MarkupErrorEmpty            MarkupErrorEnum = 1
	MarkupErrorParse            MarkupErrorEnum = 2
	MarkupErrorUnknownElement   MarkupErrorEnum = 3
	MarkupErrorUnknownAttribute MarkupErrorEnum = 4
	MarkupErrorInvalidContent   MarkupErrorEnum = 5
	MarkupErrorMissingAttribute MarkupErrorEnum = 6
)

func MarkupErrorGetType() gi.GType {
	ret := _I.GetGType(50, "MarkupError")
	return ret
}

// Struct MarkupParseContext
type MarkupParseContext struct {
	P unsafe.Pointer
}

func MarkupParseContextGetType() gi.GType {
	ret := _I.GetGType(51, "MarkupParseContext")
	return ret
}

// g_markup_parse_context_new
//
// [ parser ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ user_data_dnotify ] trans: nothing
//
// [ result ] trans: everything
//
func NewMarkupParseContext(parser MarkupParser, flags MarkupParseFlags, user_data unsafe.Pointer, user_data_dnotify int /*TODO_TYPE CALLBACK*/) (result MarkupParseContext) {
	iv, err := _I.Get(352, "MarkupParseContext", "new", 166, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_parser := gi.NewPointerArgument(parser.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_user_data_dnotify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_parser, arg_flags, arg_user_data, arg_user_data_dnotify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_markup_parse_context_end_parse
//
// [ result ] trans: nothing
//
func (v MarkupParseContext) EndParse() (result bool, err error) {
	iv, err := _I.Get(353, "MarkupParseContext", "end_parse", 166, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_markup_parse_context_free
//
func (v MarkupParseContext) Free() {
	iv, err := _I.Get(354, "MarkupParseContext", "free", 166, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_markup_parse_context_get_element
//
// [ result ] trans: nothing
//
func (v MarkupParseContext) GetElement() (result string) {
	iv, err := _I.Get(355, "MarkupParseContext", "get_element", 166, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_markup_parse_context_get_position
//
// [ line_number ] trans: nothing
//
// [ char_number ] trans: nothing
//
func (v MarkupParseContext) GetPosition(line_number int32, char_number int32) {
	iv, err := _I.Get(356, "MarkupParseContext", "get_position", 166, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_line_number := gi.NewInt32Argument(line_number)
	arg_char_number := gi.NewInt32Argument(char_number)
	args := []gi.Argument{arg_v, arg_line_number, arg_char_number}
	iv.Call(args, nil, nil)
}

// g_markup_parse_context_get_user_data
//
// [ result ] trans: nothing
//
func (v MarkupParseContext) GetUserData() (result unsafe.Pointer) {
	iv, err := _I.Get(357, "MarkupParseContext", "get_user_data", 166, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_markup_parse_context_parse
//
// [ text ] trans: nothing
//
// [ text_len ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MarkupParseContext) Parse(text string, text_len int64) (result bool, err error) {
	iv, err := _I.Get(358, "MarkupParseContext", "parse", 166, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_text_len := gi.NewInt64Argument(text_len)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_text, arg_text_len, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_markup_parse_context_pop
//
// [ result ] trans: nothing
//
func (v MarkupParseContext) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(359, "MarkupParseContext", "pop", 166, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_markup_parse_context_push
//
// [ parser ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v MarkupParseContext) Push(parser MarkupParser, user_data unsafe.Pointer) {
	iv, err := _I.Get(360, "MarkupParseContext", "push", 166, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parser := gi.NewPointerArgument(parser.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_parser, arg_user_data}
	iv.Call(args, nil, nil)
}

// g_markup_parse_context_ref
//
// [ result ] trans: everything
//
func (v MarkupParseContext) Ref() (result MarkupParseContext) {
	iv, err := _I.Get(361, "MarkupParseContext", "ref", 166, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_markup_parse_context_unref
//
func (v MarkupParseContext) Unref() {
	iv, err := _I.Get(362, "MarkupParseContext", "unref", 166, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags MarkupParseFlags
type MarkupParseFlags int

const (
	MarkupParseFlagsDoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	MarkupParseFlagsTreatCdataAsText            MarkupParseFlags = 2
	MarkupParseFlagsPrefixErrorPosition         MarkupParseFlags = 4
	MarkupParseFlagsIgnoreQualified             MarkupParseFlags = 8
)

func MarkupParseFlagsGetType() gi.GType {
	ret := _I.GetGType(52, "MarkupParseFlags")
	return ret
}

// Struct MarkupParser
type MarkupParser struct {
	P unsafe.Pointer
}

const SizeOfStructMarkupParser = 40

func MarkupParserGetType() gi.GType {
	ret := _I.GetGType(53, "MarkupParser")
	return ret
}

// Struct MatchInfo
type MatchInfo struct {
	P unsafe.Pointer
}

func MatchInfoGetType() gi.GType {
	ret := _I.GetGType(54, "MatchInfo")
	return ret
}

// g_match_info_expand_references
//
// [ string_to_expand ] trans: nothing
//
// [ result ] trans: everything
//
func (v MatchInfo) ExpandReferences(string_to_expand string) (result string, err error) {
	iv, err := _I.Get(363, "MatchInfo", "expand_references", 169, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_string_to_expand := gi.CString(string_to_expand)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string_to_expand := gi.NewStringArgument(c_string_to_expand)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string_to_expand, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string_to_expand)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_match_info_fetch
//
// [ match_num ] trans: nothing
//
// [ result ] trans: everything
//
func (v MatchInfo) Fetch(match_num int32) (result string) {
	iv, err := _I.Get(364, "MatchInfo", "fetch", 169, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_num := gi.NewInt32Argument(match_num)
	args := []gi.Argument{arg_v, arg_match_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_match_info_fetch_all
//
// [ result ] trans: everything
//
func (v MatchInfo) FetchAll() (result gi.CStrArray) {
	iv, err := _I.Get(365, "MatchInfo", "fetch_all", 169, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_match_info_fetch_named
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v MatchInfo) FetchNamed(name string) (result string) {
	iv, err := _I.Get(366, "MatchInfo", "fetch_named", 169, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.String().Take()
	return
}

// g_match_info_fetch_named_pos
//
// [ name ] trans: nothing
//
// [ start_pos ] trans: everything, dir: out
//
// [ end_pos ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v MatchInfo) FetchNamedPos(name string) (result bool, start_pos int32, end_pos int32) {
	iv, err := _I.Get(367, "MatchInfo", "fetch_named_pos", 169, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_start_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_name, arg_start_pos, arg_end_pos}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	start_pos = outArgs[0].Int32()
	end_pos = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// g_match_info_fetch_pos
//
// [ match_num ] trans: nothing
//
// [ start_pos ] trans: everything, dir: out
//
// [ end_pos ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v MatchInfo) FetchPos(match_num int32) (result bool, start_pos int32, end_pos int32) {
	iv, err := _I.Get(368, "MatchInfo", "fetch_pos", 169, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_num := gi.NewInt32Argument(match_num)
	arg_start_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_match_num, arg_start_pos, arg_end_pos}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_pos = outArgs[0].Int32()
	end_pos = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// g_match_info_free
//
func (v MatchInfo) Free() {
	iv, err := _I.Get(369, "MatchInfo", "free", 169, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_match_info_get_match_count
//
// [ result ] trans: nothing
//
func (v MatchInfo) GetMatchCount() (result int32) {
	iv, err := _I.Get(370, "MatchInfo", "get_match_count", 169, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_match_info_get_regex
//
// [ result ] trans: everything
//
func (v MatchInfo) GetRegex() (result Regex) {
	iv, err := _I.Get(371, "MatchInfo", "get_regex", 169, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_match_info_get_string
//
// [ result ] trans: nothing
//
func (v MatchInfo) GetString() (result string) {
	iv, err := _I.Get(372, "MatchInfo", "get_string", 169, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_match_info_is_partial_match
//
// [ result ] trans: nothing
//
func (v MatchInfo) IsPartialMatch() (result bool) {
	iv, err := _I.Get(373, "MatchInfo", "is_partial_match", 169, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_match_info_matches
//
// [ result ] trans: nothing
//
func (v MatchInfo) Matches() (result bool) {
	iv, err := _I.Get(374, "MatchInfo", "matches", 169, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_match_info_next
//
// [ result ] trans: nothing
//
func (v MatchInfo) Next() (result bool, err error) {
	iv, err := _I.Get(375, "MatchInfo", "next", 169, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_match_info_ref
//
// [ result ] trans: everything
//
func (v MatchInfo) Ref() (result MatchInfo) {
	iv, err := _I.Get(376, "MatchInfo", "ref", 169, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_match_info_unref
//
func (v MatchInfo) Unref() {
	iv, err := _I.Get(377, "MatchInfo", "unref", 169, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct MemVTable
type MemVTable struct {
	P unsafe.Pointer
}

const SizeOfStructMemVTable = 48

func MemVTableGetType() gi.GType {
	ret := _I.GetGType(55, "MemVTable")
	return ret
}

// Union Mutex
type Mutex struct {
	P unsafe.Pointer
}

const SizeOfUnionMutex = 8

func MutexGetType() gi.GType {
	ret := _I.GetGType(56, "Mutex")
	return ret
}

// g_mutex_clear
//
func (v Mutex) Clear() {
	iv, err := _I.Get(378, "Mutex", "clear", 171, 0, gi.INFO_TYPE_UNION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mutex_init
//
func (v Mutex) Init() {
	iv, err := _I.Get(379, "Mutex", "init", 171, 1, gi.INFO_TYPE_UNION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mutex_lock
//
func (v Mutex) Lock() {
	iv, err := _I.Get(380, "Mutex", "lock", 171, 2, gi.INFO_TYPE_UNION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mutex_trylock
//
// [ result ] trans: nothing
//
func (v Mutex) Trylock() (result bool) {
	iv, err := _I.Get(381, "Mutex", "trylock", 171, 3, gi.INFO_TYPE_UNION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_mutex_unlock
//
func (v Mutex) Unlock() {
	iv, err := _I.Get(382, "Mutex", "unlock", 171, 4, gi.INFO_TYPE_UNION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Node
type Node struct {
	P unsafe.Pointer
}

const SizeOfStructNode = 40

func NodeGetType() gi.GType {
	ret := _I.GetGType(57, "Node")
	return ret
}

// g_node_child_index
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Node) ChildIndex(data unsafe.Pointer) (result int32) {
	iv, err := _I.Get(383, "Node", "child_index", 172, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_node_child_position
//
// [ child ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Node) ChildPosition(child Node) (result int32) {
	iv, err := _I.Get(384, "Node", "child_position", 172, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child := gi.NewPointerArgument(child.P)
	args := []gi.Argument{arg_v, arg_child}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_node_depth
//
// [ result ] trans: nothing
//
func (v Node) Depth() (result uint32) {
	iv, err := _I.Get(385, "Node", "depth", 172, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_node_destroy
//
func (v Node) Destroy() {
	iv, err := _I.Get(386, "Node", "destroy", 172, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_node_is_ancestor
//
// [ descendant ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Node) IsAncestor(descendant Node) (result bool) {
	iv, err := _I.Get(387, "Node", "is_ancestor", 172, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_descendant := gi.NewPointerArgument(descendant.P)
	args := []gi.Argument{arg_v, arg_descendant}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_node_max_height
//
// [ result ] trans: nothing
//
func (v Node) MaxHeight() (result uint32) {
	iv, err := _I.Get(388, "Node", "max_height", 172, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_node_n_children
//
// [ result ] trans: nothing
//
func (v Node) NChildren() (result uint32) {
	iv, err := _I.Get(389, "Node", "n_children", 172, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_node_n_nodes
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Node) NNodes(flags TraverseFlags) (result uint32) {
	iv, err := _I.Get(390, "Node", "n_nodes", 172, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_node_reverse_children
//
func (v Node) ReverseChildren() {
	iv, err := _I.Get(391, "Node", "reverse_children", 172, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_node_unlink
//
func (v Node) Unlink() {
	iv, err := _I.Get(392, "Node", "unlink", 172, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

type NodeForeachFuncStruct struct {
	F_node Node
	F_data unsafe.Pointer
}

func GetPointer_myNodeForeachFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibNodeForeachFunc())
}

//export myGLibNodeForeachFunc
func myGLibNodeForeachFunc(node *C.GNode, data C.gpointer) {
	// TODO: not found user_data
}

type NodeTraverseFuncStruct struct {
	F_node Node
	F_data unsafe.Pointer
}

func GetPointer_myNodeTraverseFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibNodeTraverseFunc())
}

//export myGLibNodeTraverseFunc
func myGLibNodeTraverseFunc(node *C.GNode, data C.gpointer) {
	// TODO: not found user_data
}

// Enum NormalizeMode
type NormalizeModeEnum int

const (
	NormalizeModeDefault        NormalizeModeEnum = 0
	NormalizeModeNfd            NormalizeModeEnum = 0
	NormalizeModeDefaultCompose NormalizeModeEnum = 1
	NormalizeModeNfc            NormalizeModeEnum = 1
	NormalizeModeAll            NormalizeModeEnum = 2
	NormalizeModeNfkd           NormalizeModeEnum = 2
	NormalizeModeAllCompose     NormalizeModeEnum = 3
	NormalizeModeNfkc           NormalizeModeEnum = 3
)

func NormalizeModeGetType() gi.GType {
	ret := _I.GetGType(58, "NormalizeMode")
	return ret
}

// Enum NumberParserError
type NumberParserErrorEnum int

const (
	NumberParserErrorInvalid     NumberParserErrorEnum = 0
	NumberParserErrorOutOfBounds NumberParserErrorEnum = 1
)

func NumberParserErrorGetType() gi.GType {
	ret := _I.GetGType(59, "NumberParserError")
	return ret
}

// Struct Once
type Once struct {
	P unsafe.Pointer
}

const SizeOfStructOnce = 16

func OnceGetType() gi.GType {
	ret := _I.GetGType(60, "Once")
	return ret
}

// g_once_init_enter
//
// [ location ] trans: nothing
//
// [ result ] trans: nothing
//
func OnceInitEnter1(location unsafe.Pointer) (result bool) {
	iv, err := _I.Get(393, "Once", "init_enter", 178, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	args := []gi.Argument{arg_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_once_init_leave
//
// [ location ] trans: nothing
//
// [ result ] trans: nothing
//
func OnceInitLeave1(location unsafe.Pointer, result uint64) {
	iv, err := _I.Get(394, "Once", "init_leave", 178, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	arg_result := gi.NewUint64Argument(result)
	args := []gi.Argument{arg_location, arg_result}
	iv.Call(args, nil, nil)
}

// Enum OnceStatus
type OnceStatusEnum int

const (
	OnceStatusNotcalled OnceStatusEnum = 0
	OnceStatusProgress  OnceStatusEnum = 1
	OnceStatusReady     OnceStatusEnum = 2
)

func OnceStatusGetType() gi.GType {
	ret := _I.GetGType(61, "OnceStatus")
	return ret
}

// Enum OptionArg
type OptionArgEnum int

const (
	OptionArgNone          OptionArgEnum = 0
	OptionArgString        OptionArgEnum = 1
	OptionArgInt           OptionArgEnum = 2
	OptionArgCallback      OptionArgEnum = 3
	OptionArgFilename      OptionArgEnum = 4
	OptionArgStringArray   OptionArgEnum = 5
	OptionArgFilenameArray OptionArgEnum = 6
	OptionArgDouble        OptionArgEnum = 7
	OptionArgInt64         OptionArgEnum = 8
)

func OptionArgGetType() gi.GType {
	ret := _I.GetGType(62, "OptionArg")
	return ret
}

type OptionArgFuncStruct struct {
	F_option_name string
	F_value       string
	F_data        unsafe.Pointer
}

func GetPointer_myOptionArgFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibOptionArgFunc())
}

//export myGLibOptionArgFunc
func myGLibOptionArgFunc(option_name *C.gchar, value *C.gchar, data C.gpointer) {
	// TODO: not found user_data
}

// Struct OptionContext
type OptionContext struct {
	P unsafe.Pointer
}

func OptionContextGetType() gi.GType {
	ret := _I.GetGType(63, "OptionContext")
	return ret
}

// g_option_context_add_group
//
// [ group ] trans: everything
//
func (v OptionContext) AddGroup(group OptionGroup) {
	iv, err := _I.Get(395, "OptionContext", "add_group", 182, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_v, arg_group}
	iv.Call(args, nil, nil)
}

// g_option_context_add_main_entries
//
// [ entries ] trans: nothing
//
// [ translation_domain ] trans: nothing
//
func (v OptionContext) AddMainEntries(entries OptionEntry, translation_domain string) {
	iv, err := _I.Get(396, "OptionContext", "add_main_entries", 182, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_translation_domain := gi.CString(translation_domain)
	arg_v := gi.NewPointerArgument(v.P)
	arg_entries := gi.NewPointerArgument(entries.P)
	arg_translation_domain := gi.NewStringArgument(c_translation_domain)
	args := []gi.Argument{arg_v, arg_entries, arg_translation_domain}
	iv.Call(args, nil, nil)
	gi.Free(c_translation_domain)
}

// g_option_context_free
//
func (v OptionContext) Free() {
	iv, err := _I.Get(397, "OptionContext", "free", 182, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_option_context_get_description
//
// [ result ] trans: nothing
//
func (v OptionContext) GetDescription() (result string) {
	iv, err := _I.Get(398, "OptionContext", "get_description", 182, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_option_context_get_help
//
// [ main_help ] trans: nothing
//
// [ group ] trans: nothing
//
// [ result ] trans: everything
//
func (v OptionContext) GetHelp(main_help bool, group OptionGroup) (result string) {
	iv, err := _I.Get(399, "OptionContext", "get_help", 182, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_main_help := gi.NewBoolArgument(main_help)
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_v, arg_main_help, arg_group}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_option_context_get_help_enabled
//
// [ result ] trans: nothing
//
func (v OptionContext) GetHelpEnabled() (result bool) {
	iv, err := _I.Get(400, "OptionContext", "get_help_enabled", 182, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_option_context_get_ignore_unknown_options
//
// [ result ] trans: nothing
//
func (v OptionContext) GetIgnoreUnknownOptions() (result bool) {
	iv, err := _I.Get(401, "OptionContext", "get_ignore_unknown_options", 182, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_option_context_get_main_group
//
// [ result ] trans: nothing
//
func (v OptionContext) GetMainGroup() (result OptionGroup) {
	iv, err := _I.Get(402, "OptionContext", "get_main_group", 182, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_option_context_get_strict_posix
//
// [ result ] trans: nothing
//
func (v OptionContext) GetStrictPosix() (result bool) {
	iv, err := _I.Get(403, "OptionContext", "get_strict_posix", 182, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_option_context_get_summary
//
// [ result ] trans: nothing
//
func (v OptionContext) GetSummary() (result string) {
	iv, err := _I.Get(404, "OptionContext", "get_summary", 182, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_option_context_parse
//
// [ argc ] trans: everything, dir: inout
//
// [ argv ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func (v OptionContext) Parse(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) (result bool, err error) {
	iv, err := _I.Get(405, "OptionContext", "parse", 182, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_option_context_parse_strv
//
// [ arguments ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func (v OptionContext) ParseStrv(arguments int /*TODO:TYPE*/) (result bool, err error) {
	iv, err := _I.Get(406, "OptionContext", "parse_strv", 182, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_option_context_set_description
//
// [ description ] trans: nothing
//
func (v OptionContext) SetDescription(description string) {
	iv, err := _I.Get(407, "OptionContext", "set_description", 182, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(v.P)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// g_option_context_set_help_enabled
//
// [ help_enabled ] trans: nothing
//
func (v OptionContext) SetHelpEnabled(help_enabled bool) {
	iv, err := _I.Get(408, "OptionContext", "set_help_enabled", 182, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_help_enabled := gi.NewBoolArgument(help_enabled)
	args := []gi.Argument{arg_v, arg_help_enabled}
	iv.Call(args, nil, nil)
}

// g_option_context_set_ignore_unknown_options
//
// [ ignore_unknown ] trans: nothing
//
func (v OptionContext) SetIgnoreUnknownOptions(ignore_unknown bool) {
	iv, err := _I.Get(409, "OptionContext", "set_ignore_unknown_options", 182, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ignore_unknown := gi.NewBoolArgument(ignore_unknown)
	args := []gi.Argument{arg_v, arg_ignore_unknown}
	iv.Call(args, nil, nil)
}

// g_option_context_set_main_group
//
// [ group ] trans: everything
//
func (v OptionContext) SetMainGroup(group OptionGroup) {
	iv, err := _I.Get(410, "OptionContext", "set_main_group", 182, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_v, arg_group}
	iv.Call(args, nil, nil)
}

// g_option_context_set_strict_posix
//
// [ strict_posix ] trans: nothing
//
func (v OptionContext) SetStrictPosix(strict_posix bool) {
	iv, err := _I.Get(411, "OptionContext", "set_strict_posix", 182, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_strict_posix := gi.NewBoolArgument(strict_posix)
	args := []gi.Argument{arg_v, arg_strict_posix}
	iv.Call(args, nil, nil)
}

// g_option_context_set_summary
//
// [ summary ] trans: nothing
//
func (v OptionContext) SetSummary(summary string) {
	iv, err := _I.Get(412, "OptionContext", "set_summary", 182, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_summary := gi.CString(summary)
	arg_v := gi.NewPointerArgument(v.P)
	arg_summary := gi.NewStringArgument(c_summary)
	args := []gi.Argument{arg_v, arg_summary}
	iv.Call(args, nil, nil)
	gi.Free(c_summary)
}

// g_option_context_set_translate_func
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ destroy_notify ] trans: nothing
//
func (v OptionContext) SetTranslateFunc(func1 int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, destroy_notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(413, "OptionContext", "set_translate_func", 182, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTranslateFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_destroy_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_func1, arg_data, arg_destroy_notify}
	iv.Call(args, nil, nil)
}

// g_option_context_set_translation_domain
//
// [ domain ] trans: nothing
//
func (v OptionContext) SetTranslationDomain(domain string) {
	iv, err := _I.Get(414, "OptionContext", "set_translation_domain", 182, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewStringArgument(c_domain)
	args := []gi.Argument{arg_v, arg_domain}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
}

// Struct OptionEntry
type OptionEntry struct {
	P unsafe.Pointer
}

const SizeOfStructOptionEntry = 48

func OptionEntryGetType() gi.GType {
	ret := _I.GetGType(64, "OptionEntry")
	return ret
}

// Enum OptionError
type OptionErrorEnum int

const (
	OptionErrorUnknownOption OptionErrorEnum = 0
	OptionErrorBadValue      OptionErrorEnum = 1
	OptionErrorFailed        OptionErrorEnum = 2
)

func OptionErrorGetType() gi.GType {
	ret := _I.GetGType(65, "OptionError")
	return ret
}

type OptionErrorFuncStruct struct {
	F_context OptionContext
	F_group   OptionGroup
	F_data    unsafe.Pointer
}

func GetPointer_myOptionErrorFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibOptionErrorFunc())
}

//export myGLibOptionErrorFunc
func myGLibOptionErrorFunc(context *C.GOptionContext, group *C.GOptionGroup, data C.gpointer) {
	// TODO: not found user_data
}

// Flags OptionFlags
type OptionFlags int

const (
	OptionFlagsNone        OptionFlags = 0
	OptionFlagsHidden      OptionFlags = 1
	OptionFlagsInMain      OptionFlags = 2
	OptionFlagsReverse     OptionFlags = 4
	OptionFlagsNoArg       OptionFlags = 8
	OptionFlagsFilename    OptionFlags = 16
	OptionFlagsOptionalArg OptionFlags = 32
	OptionFlagsNoalias     OptionFlags = 64
)

func OptionFlagsGetType() gi.GType {
	ret := _I.GetGType(66, "OptionFlags")
	return ret
}

// Struct OptionGroup
type OptionGroup struct {
	P unsafe.Pointer
}

func OptionGroupGetType() gi.GType {
	ret := _I.GetGType(67, "OptionGroup")
	return ret
}

// g_option_group_new
//
// [ name ] trans: nothing
//
// [ description ] trans: nothing
//
// [ help_description ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ destroy ] trans: nothing
//
// [ result ] trans: everything
//
func NewOptionGroup(name string, description string, help_description string, user_data unsafe.Pointer, destroy int /*TODO_TYPE CALLBACK*/) (result OptionGroup) {
	iv, err := _I.Get(415, "OptionGroup", "new", 187, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_description := gi.CString(description)
	c_help_description := gi.CString(help_description)
	arg_name := gi.NewStringArgument(c_name)
	arg_description := gi.NewStringArgument(c_description)
	arg_help_description := gi.NewStringArgument(c_help_description)
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_name, arg_description, arg_help_description, arg_user_data, arg_destroy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_description)
	gi.Free(c_help_description)
	result.P = ret.Pointer()
	return
}

// g_option_group_add_entries
//
// [ entries ] trans: nothing
//
func (v OptionGroup) AddEntries(entries OptionEntry) {
	iv, err := _I.Get(416, "OptionGroup", "add_entries", 187, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_entries := gi.NewPointerArgument(entries.P)
	args := []gi.Argument{arg_v, arg_entries}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_option_group_free
//
func (v OptionGroup) Free() {
	iv, err := _I.Get(417, "OptionGroup", "free", 187, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_option_group_ref
//
// [ result ] trans: everything
//
func (v OptionGroup) Ref() (result OptionGroup) {
	iv, err := _I.Get(418, "OptionGroup", "ref", 187, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_option_group_set_translate_func
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ destroy_notify ] trans: nothing
//
func (v OptionGroup) SetTranslateFunc(func1 int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, destroy_notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(419, "OptionGroup", "set_translate_func", 187, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTranslateFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_destroy_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_func1, arg_data, arg_destroy_notify}
	iv.Call(args, nil, nil)
}

// g_option_group_set_translation_domain
//
// [ domain ] trans: nothing
//
func (v OptionGroup) SetTranslationDomain(domain string) {
	iv, err := _I.Get(420, "OptionGroup", "set_translation_domain", 187, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewStringArgument(c_domain)
	args := []gi.Argument{arg_v, arg_domain}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
}

// g_option_group_unref
//
func (v OptionGroup) Unref() {
	iv, err := _I.Get(421, "OptionGroup", "unref", 187, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

type OptionParseFuncStruct struct {
	F_context OptionContext
	F_group   OptionGroup
	F_data    unsafe.Pointer
}

func GetPointer_myOptionParseFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibOptionParseFunc())
}

//export myGLibOptionParseFunc
func myGLibOptionParseFunc(context *C.GOptionContext, group *C.GOptionGroup, data C.gpointer) {
	// TODO: not found user_data
}

// Struct PatternSpec
type PatternSpec struct {
	P unsafe.Pointer
}

func PatternSpecGetType() gi.GType {
	ret := _I.GetGType(68, "PatternSpec")
	return ret
}

// g_pattern_spec_equal
//
// [ pspec2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v PatternSpec) Equal(pspec2 PatternSpec) (result bool) {
	iv, err := _I.Get(422, "PatternSpec", "equal", 200, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pspec2 := gi.NewPointerArgument(pspec2.P)
	args := []gi.Argument{arg_v, arg_pspec2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_pattern_spec_free
//
func (v PatternSpec) Free() {
	iv, err := _I.Get(423, "PatternSpec", "free", 200, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct PollFD
type PollFD struct {
	P unsafe.Pointer
}

const SizeOfStructPollFD = 8

func PollFDGetType() gi.GType {
	ret := _I.GetGType(69, "PollFD")
	return ret
}

type PollFuncStruct struct {
	F_ufds     PollFD
	F_nfsd     uint32
	F_timeout_ int32
}

func GetPointer_myPollFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibPollFunc())
}

//export myGLibPollFunc
func myGLibPollFunc(ufds *C.GPollFD, nfsd C.guint32, timeout_ C.gint32) {
	// TODO: not found user_data
}

type PrintFuncStruct struct {
	F_string string
}

func GetPointer_myPrintFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibPrintFunc())
}

//export myGLibPrintFunc
func myGLibPrintFunc(string *C.gchar) {
	// TODO: not found user_data
}

// Struct Private
type Private struct {
	P unsafe.Pointer
}

const SizeOfStructPrivate = 32

func PrivateGetType() gi.GType {
	ret := _I.GetGType(70, "Private")
	return ret
}

// g_private_get
//
// [ result ] trans: nothing
//
func (v Private) Get() (result unsafe.Pointer) {
	iv, err := _I.Get(424, "Private", "get", 204, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_private_replace
//
// [ value ] trans: nothing
//
func (v Private) Replace(value unsafe.Pointer) {
	iv, err := _I.Get(425, "Private", "replace", 204, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// g_private_set
//
// [ value ] trans: nothing
//
func (v Private) Set(value unsafe.Pointer) {
	iv, err := _I.Get(426, "Private", "set", 204, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// Struct PtrArray
type PtrArray struct {
	P unsafe.Pointer
}

const SizeOfStructPtrArray = 16

func PtrArrayGetType() gi.GType {
	ret := _I.GetGType(71, "PtrArray")
	return ret
}

// Struct Queue
type Queue struct {
	P unsafe.Pointer
}

const SizeOfStructQueue = 24

func QueueGetType() gi.GType {
	ret := _I.GetGType(72, "Queue")
	return ret
}

// g_queue_clear
//
func (v Queue) Clear() {
	iv, err := _I.Get(427, "Queue", "clear", 206, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_queue_free
//
func (v Queue) Free() {
	iv, err := _I.Get(428, "Queue", "free", 206, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_queue_free_full
//
// [ free_func ] trans: nothing
//
func (v Queue) FreeFull(free_func int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(429, "Queue", "free_full", 206, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_free_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_free_func}
	iv.Call(args, nil, nil)
}

// g_queue_get_length
//
// [ result ] trans: nothing
//
func (v Queue) GetLength() (result uint32) {
	iv, err := _I.Get(430, "Queue", "get_length", 206, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_queue_index
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Queue) Index(data unsafe.Pointer) (result int32) {
	iv, err := _I.Get(431, "Queue", "index", 206, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_queue_init
//
func (v Queue) Init() {
	iv, err := _I.Get(432, "Queue", "init", 206, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_queue_is_empty
//
// [ result ] trans: nothing
//
func (v Queue) IsEmpty() (result bool) {
	iv, err := _I.Get(433, "Queue", "is_empty", 206, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_queue_peek_head
//
// [ result ] trans: nothing
//
func (v Queue) PeekHead() (result unsafe.Pointer) {
	iv, err := _I.Get(434, "Queue", "peek_head", 206, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_peek_nth
//
// [ n ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Queue) PeekNth(n uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(435, "Queue", "peek_nth", 206, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n := gi.NewUint32Argument(n)
	args := []gi.Argument{arg_v, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_peek_tail
//
// [ result ] trans: nothing
//
func (v Queue) PeekTail() (result unsafe.Pointer) {
	iv, err := _I.Get(436, "Queue", "peek_tail", 206, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_pop_head
//
// [ result ] trans: nothing
//
func (v Queue) PopHead() (result unsafe.Pointer) {
	iv, err := _I.Get(437, "Queue", "pop_head", 206, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_pop_nth
//
// [ n ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Queue) PopNth(n uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(438, "Queue", "pop_nth", 206, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n := gi.NewUint32Argument(n)
	args := []gi.Argument{arg_v, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_pop_tail
//
// [ result ] trans: nothing
//
func (v Queue) PopTail() (result unsafe.Pointer) {
	iv, err := _I.Get(439, "Queue", "pop_tail", 206, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_push_head
//
// [ data ] trans: nothing
//
func (v Queue) PushHead(data unsafe.Pointer) {
	iv, err := _I.Get(440, "Queue", "push_head", 206, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_queue_push_nth
//
// [ data ] trans: nothing
//
// [ n ] trans: nothing
//
func (v Queue) PushNth(data unsafe.Pointer, n int32) {
	iv, err := _I.Get(441, "Queue", "push_nth", 206, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_v, arg_data, arg_n}
	iv.Call(args, nil, nil)
}

// g_queue_push_tail
//
// [ data ] trans: nothing
//
func (v Queue) PushTail(data unsafe.Pointer) {
	iv, err := _I.Get(442, "Queue", "push_tail", 206, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_queue_remove
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Queue) Remove(data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(443, "Queue", "remove", 206, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_queue_remove_all
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Queue) RemoveAll(data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(444, "Queue", "remove_all", 206, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_queue_reverse
//
func (v Queue) Reverse() {
	iv, err := _I.Get(445, "Queue", "reverse", 206, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct RWLock
type RWLock struct {
	P unsafe.Pointer
}

const SizeOfStructRWLock = 16

func RWLockGetType() gi.GType {
	ret := _I.GetGType(73, "RWLock")
	return ret
}

// g_rw_lock_clear
//
func (v RWLock) Clear() {
	iv, err := _I.Get(446, "RWLock", "clear", 207, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_init
//
func (v RWLock) Init() {
	iv, err := _I.Get(447, "RWLock", "init", 207, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_reader_lock
//
func (v RWLock) ReaderLock() {
	iv, err := _I.Get(448, "RWLock", "reader_lock", 207, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_reader_trylock
//
// [ result ] trans: nothing
//
func (v RWLock) ReaderTrylock() (result bool) {
	iv, err := _I.Get(449, "RWLock", "reader_trylock", 207, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_rw_lock_reader_unlock
//
func (v RWLock) ReaderUnlock() {
	iv, err := _I.Get(450, "RWLock", "reader_unlock", 207, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_writer_lock
//
func (v RWLock) WriterLock() {
	iv, err := _I.Get(451, "RWLock", "writer_lock", 207, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_writer_trylock
//
// [ result ] trans: nothing
//
func (v RWLock) WriterTrylock() (result bool) {
	iv, err := _I.Get(452, "RWLock", "writer_trylock", 207, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_rw_lock_writer_unlock
//
func (v RWLock) WriterUnlock() {
	iv, err := _I.Get(453, "RWLock", "writer_unlock", 207, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Rand
type Rand struct {
	P unsafe.Pointer
}

func RandGetType() gi.GType {
	ret := _I.GetGType(74, "Rand")
	return ret
}

// g_rand_double
//
// [ result ] trans: nothing
//
func (v Rand) Double() (result float64) {
	iv, err := _I.Get(454, "Rand", "double", 208, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_rand_double_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Rand) DoubleRange(begin float64, end float64) (result float64) {
	iv, err := _I.Get(455, "Rand", "double_range", 208, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_begin := gi.NewDoubleArgument(begin)
	arg_end := gi.NewDoubleArgument(end)
	args := []gi.Argument{arg_v, arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_rand_free
//
func (v Rand) Free() {
	iv, err := _I.Get(456, "Rand", "free", 208, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rand_int
//
// [ result ] trans: nothing
//
func (v Rand) Int() (result uint32) {
	iv, err := _I.Get(457, "Rand", "int", 208, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_rand_int_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Rand) IntRange(begin int32, end int32) (result int32) {
	iv, err := _I.Get(458, "Rand", "int_range", 208, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_begin := gi.NewInt32Argument(begin)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_v, arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_rand_set_seed
//
// [ seed ] trans: nothing
//
func (v Rand) SetSeed(seed uint32) {
	iv, err := _I.Get(459, "Rand", "set_seed", 208, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seed := gi.NewUint32Argument(seed)
	args := []gi.Argument{arg_v, arg_seed}
	iv.Call(args, nil, nil)
}

// g_rand_set_seed_array
//
// [ seed ] trans: nothing
//
// [ seed_length ] trans: nothing
//
func (v Rand) SetSeedArray(seed uint32, seed_length uint32) {
	iv, err := _I.Get(460, "Rand", "set_seed_array", 208, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seed := gi.NewUint32Argument(seed)
	arg_seed_length := gi.NewUint32Argument(seed_length)
	args := []gi.Argument{arg_v, arg_seed, arg_seed_length}
	iv.Call(args, nil, nil)
}

// Struct RecMutex
type RecMutex struct {
	P unsafe.Pointer
}

const SizeOfStructRecMutex = 16

func RecMutexGetType() gi.GType {
	ret := _I.GetGType(75, "RecMutex")
	return ret
}

// g_rec_mutex_clear
//
func (v RecMutex) Clear() {
	iv, err := _I.Get(461, "RecMutex", "clear", 209, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rec_mutex_init
//
func (v RecMutex) Init() {
	iv, err := _I.Get(462, "RecMutex", "init", 209, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rec_mutex_lock
//
func (v RecMutex) Lock() {
	iv, err := _I.Get(463, "RecMutex", "lock", 209, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rec_mutex_trylock
//
// [ result ] trans: nothing
//
func (v RecMutex) Trylock() (result bool) {
	iv, err := _I.Get(464, "RecMutex", "trylock", 209, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_rec_mutex_unlock
//
func (v RecMutex) Unlock() {
	iv, err := _I.Get(465, "RecMutex", "unlock", 209, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Regex
type Regex struct {
	P unsafe.Pointer
}

func RegexGetType() gi.GType {
	ret := _I.GetGType(76, "Regex")
	return ret
}

// g_regex_new
//
// [ pattern ] trans: nothing
//
// [ compile_options ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: everything
//
func NewRegex(pattern string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result Regex, err error) {
	iv, err := _I.Get(466, "Regex", "new", 210, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_pattern := gi.CString(pattern)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pattern, arg_compile_options, arg_match_options, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_pattern)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_regex_get_capture_count
//
// [ result ] trans: nothing
//
func (v Regex) GetCaptureCount() (result int32) {
	iv, err := _I.Get(467, "Regex", "get_capture_count", 210, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_regex_get_compile_flags
//
// [ result ] trans: nothing
//
func (v Regex) GetCompileFlags() (result RegexCompileFlags) {
	iv, err := _I.Get(468, "Regex", "get_compile_flags", 210, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = RegexCompileFlags(ret.Int())
	return
}

// g_regex_get_has_cr_or_lf
//
// [ result ] trans: nothing
//
func (v Regex) GetHasCrOrLf() (result bool) {
	iv, err := _I.Get(469, "Regex", "get_has_cr_or_lf", 210, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_regex_get_match_flags
//
// [ result ] trans: nothing
//
func (v Regex) GetMatchFlags() (result RegexMatchFlags) {
	iv, err := _I.Get(470, "Regex", "get_match_flags", 210, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = RegexMatchFlags(ret.Int())
	return
}

// g_regex_get_max_backref
//
// [ result ] trans: nothing
//
func (v Regex) GetMaxBackref() (result int32) {
	iv, err := _I.Get(471, "Regex", "get_max_backref", 210, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_regex_get_max_lookbehind
//
// [ result ] trans: nothing
//
func (v Regex) GetMaxLookbehind() (result int32) {
	iv, err := _I.Get(472, "Regex", "get_max_lookbehind", 210, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_regex_get_pattern
//
// [ result ] trans: nothing
//
func (v Regex) GetPattern() (result string) {
	iv, err := _I.Get(473, "Regex", "get_pattern", 210, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_regex_get_string_number
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Regex) GetStringNumber(name string) (result int32) {
	iv, err := _I.Get(474, "Regex", "get_string_number", 210, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Int32()
	return
}

// g_regex_match
//
// [ string ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ match_info ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Regex) Match(string string, match_options RegexMatchFlags) (result bool, match_info MatchInfo) {
	iv, err := _I.Get(475, "Regex", "match", 210, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_match_options, arg_match_info}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_match_all
//
// [ string ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ match_info ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Regex) MatchAll(string string, match_options RegexMatchFlags) (result bool, match_info MatchInfo) {
	iv, err := _I.Get(476, "Regex", "match_all", 210, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_match_options, arg_match_info}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_match_all_full
//
// [ string ] trans: nothing
//
// [ string_len ] trans: nothing
//
// [ start_position ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ match_info ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Regex) MatchAllFull(string gi.CStrArray, string_len int64, start_position int32, match_options RegexMatchFlags) (result bool, match_info MatchInfo, err error) {
	iv, err := _I.Get(477, "Regex", "match_all_full", 210, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_match_options, arg_match_info, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_match_full
//
// [ string ] trans: nothing
//
// [ string_len ] trans: nothing
//
// [ start_position ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ match_info ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Regex) MatchFull(string gi.CStrArray, string_len int64, start_position int32, match_options RegexMatchFlags) (result bool, match_info MatchInfo, err error) {
	iv, err := _I.Get(478, "Regex", "match_full", 210, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_match_options, arg_match_info, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_ref
//
// [ result ] trans: everything
//
func (v Regex) Ref() (result Regex) {
	iv, err := _I.Get(479, "Regex", "ref", 210, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_regex_replace
//
// [ string ] trans: nothing
//
// [ string_len ] trans: nothing
//
// [ start_position ] trans: nothing
//
// [ replacement ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: everything
//
func (v Regex) Replace(string gi.CStrArray, string_len int64, start_position int32, replacement string, match_options RegexMatchFlags) (result string, err error) {
	iv, err := _I.Get(480, "Regex", "replace", 210, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_replacement, arg_match_options, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_regex_replace_literal
//
// [ string ] trans: nothing
//
// [ string_len ] trans: nothing
//
// [ start_position ] trans: nothing
//
// [ replacement ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: everything
//
func (v Regex) ReplaceLiteral(string gi.CStrArray, string_len int64, start_position int32, replacement string, match_options RegexMatchFlags) (result string, err error) {
	iv, err := _I.Get(481, "Regex", "replace_literal", 210, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_replacement, arg_match_options, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_regex_split
//
// [ string ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: everything
//
func (v Regex) Split(string string, match_options RegexMatchFlags) (result gi.CStrArray) {
	iv, err := _I.Get(482, "Regex", "split", 210, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_v, arg_string, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_regex_split_full
//
// [ string ] trans: nothing
//
// [ string_len ] trans: nothing
//
// [ start_position ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ max_tokens ] trans: nothing
//
// [ result ] trans: everything
//
func (v Regex) SplitFull(string gi.CStrArray, string_len int64, start_position int32, match_options RegexMatchFlags, max_tokens int32) (result gi.CStrArray, err error) {
	iv, err := _I.Get(483, "Regex", "split_full", 210, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_max_tokens := gi.NewInt32Argument(max_tokens)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_match_options, arg_max_tokens, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_regex_unref
//
func (v Regex) Unref() {
	iv, err := _I.Get(484, "Regex", "unref", 210, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_regex_check_replacement
//
// [ replacement ] trans: nothing
//
// [ has_references ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func RegexCheckReplacement1(replacement string) (result bool, has_references bool, err error) {
	iv, err := _I.Get(485, "Regex", "check_replacement", 210, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_has_references := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_replacement, arg_has_references, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[1].Pointer())
	has_references = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// g_regex_escape_nul
//
// [ string ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func RegexEscapeNul1(string string, length int32) (result string) {
	iv, err := _I.Get(487, "Regex", "escape_nul", 210, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_regex_escape_string
//
// [ string ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func RegexEscapeString1(string gi.CStrArray, length int32) (result string) {
	iv, err := _I.Get(488, "Regex", "escape_string", 210, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewPointerArgument(string.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_regex_match_simple
//
// [ pattern ] trans: nothing
//
// [ string ] trans: nothing
//
// [ compile_options ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: nothing
//
func RegexMatchSimple1(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result bool) {
	iv, err := _I.Get(489, "Regex", "match_simple", 210, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_regex_split_simple
//
// [ pattern ] trans: nothing
//
// [ string ] trans: nothing
//
// [ compile_options ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: everything
//
func RegexSplitSimple1(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result gi.CStrArray) {
	iv, err := _I.Get(490, "Regex", "split_simple", 210, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// Flags RegexCompileFlags
type RegexCompileFlags int

const (
	RegexCompileFlagsCaseless         RegexCompileFlags = 1
	RegexCompileFlagsMultiline        RegexCompileFlags = 2
	RegexCompileFlagsDotall           RegexCompileFlags = 4
	RegexCompileFlagsExtended         RegexCompileFlags = 8
	RegexCompileFlagsAnchored         RegexCompileFlags = 16
	RegexCompileFlagsDollarEndonly    RegexCompileFlags = 32
	RegexCompileFlagsUngreedy         RegexCompileFlags = 512
	RegexCompileFlagsRaw              RegexCompileFlags = 2048
	RegexCompileFlagsNoAutoCapture    RegexCompileFlags = 4096
	RegexCompileFlagsOptimize         RegexCompileFlags = 8192
	RegexCompileFlagsFirstline        RegexCompileFlags = 262144
	RegexCompileFlagsDupnames         RegexCompileFlags = 524288
	RegexCompileFlagsNewlineCr        RegexCompileFlags = 1048576
	RegexCompileFlagsNewlineLf        RegexCompileFlags = 2097152
	RegexCompileFlagsNewlineCrlf      RegexCompileFlags = 3145728
	RegexCompileFlagsNewlineAnycrlf   RegexCompileFlags = 5242880
	RegexCompileFlagsBsrAnycrlf       RegexCompileFlags = 8388608
	RegexCompileFlagsJavascriptCompat RegexCompileFlags = 33554432
)

func RegexCompileFlagsGetType() gi.GType {
	ret := _I.GetGType(77, "RegexCompileFlags")
	return ret
}

// Enum RegexError
type RegexErrorEnum int

const (
	RegexErrorCompile                                  RegexErrorEnum = 0
	RegexErrorOptimize                                 RegexErrorEnum = 1
	RegexErrorReplace                                  RegexErrorEnum = 2
	RegexErrorMatch                                    RegexErrorEnum = 3
	RegexErrorInternal                                 RegexErrorEnum = 4
	RegexErrorStrayBackslash                           RegexErrorEnum = 101
	RegexErrorMissingControlChar                       RegexErrorEnum = 102
	RegexErrorUnrecognizedEscape                       RegexErrorEnum = 103
	RegexErrorQuantifiersOutOfOrder                    RegexErrorEnum = 104
	RegexErrorQuantifierTooBig                         RegexErrorEnum = 105
	RegexErrorUnterminatedCharacterClass               RegexErrorEnum = 106
	RegexErrorInvalidEscapeInCharacterClass            RegexErrorEnum = 107
	RegexErrorRangeOutOfOrder                          RegexErrorEnum = 108
	RegexErrorNothingToRepeat                          RegexErrorEnum = 109
	RegexErrorUnrecognizedCharacter                    RegexErrorEnum = 112
	RegexErrorPosixNamedClassOutsideClass              RegexErrorEnum = 113
	RegexErrorUnmatchedParenthesis                     RegexErrorEnum = 114
	RegexErrorInexistentSubpatternReference            RegexErrorEnum = 115
	RegexErrorUnterminatedComment                      RegexErrorEnum = 118
	RegexErrorExpressionTooLarge                       RegexErrorEnum = 120
	RegexErrorMemoryError                              RegexErrorEnum = 121
	RegexErrorVariableLengthLookbehind                 RegexErrorEnum = 125
	RegexErrorMalformedCondition                       RegexErrorEnum = 126
	RegexErrorTooManyConditionalBranches               RegexErrorEnum = 127
	RegexErrorAssertionExpected                        RegexErrorEnum = 128
	RegexErrorUnknownPosixClassName                    RegexErrorEnum = 130
	RegexErrorPosixCollatingElementsNotSupported       RegexErrorEnum = 131
	RegexErrorHexCodeTooLarge                          RegexErrorEnum = 134
	RegexErrorInvalidCondition                         RegexErrorEnum = 135
	RegexErrorSingleByteMatchInLookbehind              RegexErrorEnum = 136
	RegexErrorInfiniteLoop                             RegexErrorEnum = 140
	RegexErrorMissingSubpatternNameTerminator          RegexErrorEnum = 142
	RegexErrorDuplicateSubpatternName                  RegexErrorEnum = 143
	RegexErrorMalformedProperty                        RegexErrorEnum = 146
	RegexErrorUnknownProperty                          RegexErrorEnum = 147
	RegexErrorSubpatternNameTooLong                    RegexErrorEnum = 148
	RegexErrorTooManySubpatterns                       RegexErrorEnum = 149
	RegexErrorInvalidOctalValue                        RegexErrorEnum = 151
	RegexErrorTooManyBranchesInDefine                  RegexErrorEnum = 154
	RegexErrorDefineRepetion                           RegexErrorEnum = 155
	RegexErrorInconsistentNewlineOptions               RegexErrorEnum = 156
	RegexErrorMissingBackReference                     RegexErrorEnum = 157
	RegexErrorInvalidRelativeReference                 RegexErrorEnum = 158
	RegexErrorBacktrackingControlVerbArgumentForbidden RegexErrorEnum = 159
	RegexErrorUnknownBacktrackingControlVerb           RegexErrorEnum = 160
	RegexErrorNumberTooBig                             RegexErrorEnum = 161
	RegexErrorMissingSubpatternName                    RegexErrorEnum = 162
	RegexErrorMissingDigit                             RegexErrorEnum = 163
	RegexErrorInvalidDataCharacter                     RegexErrorEnum = 164
	RegexErrorExtraSubpatternName                      RegexErrorEnum = 165
	RegexErrorBacktrackingControlVerbArgumentRequired  RegexErrorEnum = 166
	RegexErrorInvalidControlChar                       RegexErrorEnum = 168
	RegexErrorMissingName                              RegexErrorEnum = 169
	RegexErrorNotSupportedInClass                      RegexErrorEnum = 171
	RegexErrorTooManyForwardReferences                 RegexErrorEnum = 172
	RegexErrorNameTooLong                              RegexErrorEnum = 175
	RegexErrorCharacterValueTooLarge                   RegexErrorEnum = 176
)

func RegexErrorGetType() gi.GType {
	ret := _I.GetGType(78, "RegexError")
	return ret
}

type RegexEvalCallbackStruct struct {
	F_match_info MatchInfo
	F_result     String
}

func GetPointer_myRegexEvalCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibRegexEvalCallback())
}

//export myGLibRegexEvalCallback
func myGLibRegexEvalCallback(match_info *C.GMatchInfo, result *C.GString, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &RegexEvalCallbackStruct{
		F_match_info: MatchInfo{P: unsafe.Pointer(match_info)},
		F_result:     String{P: unsafe.Pointer(result)},
	}
	fn(args)
}

// Flags RegexMatchFlags
type RegexMatchFlags int

const (
	RegexMatchFlagsAnchored        RegexMatchFlags = 16
	RegexMatchFlagsNotbol          RegexMatchFlags = 128
	RegexMatchFlagsNoteol          RegexMatchFlags = 256
	RegexMatchFlagsNotempty        RegexMatchFlags = 1024
	RegexMatchFlagsPartial         RegexMatchFlags = 32768
	RegexMatchFlagsNewlineCr       RegexMatchFlags = 1048576
	RegexMatchFlagsNewlineLf       RegexMatchFlags = 2097152
	RegexMatchFlagsNewlineCrlf     RegexMatchFlags = 3145728
	RegexMatchFlagsNewlineAny      RegexMatchFlags = 4194304
	RegexMatchFlagsNewlineAnycrlf  RegexMatchFlags = 5242880
	RegexMatchFlagsBsrAnycrlf      RegexMatchFlags = 8388608
	RegexMatchFlagsBsrAny          RegexMatchFlags = 16777216
	RegexMatchFlagsPartialSoft     RegexMatchFlags = 32768
	RegexMatchFlagsPartialHard     RegexMatchFlags = 134217728
	RegexMatchFlagsNotemptyAtstart RegexMatchFlags = 268435456
)

func RegexMatchFlagsGetType() gi.GType {
	ret := _I.GetGType(79, "RegexMatchFlags")
	return ret
}

// Struct SList
type SList struct {
	P unsafe.Pointer
}

const SizeOfStructSList = 16

func SListGetType() gi.GType {
	ret := _I.GetGType(80, "SList")
	return ret
}

// Struct Scanner
type Scanner struct {
	P unsafe.Pointer
}

const SizeOfStructScanner = 144

func ScannerGetType() gi.GType {
	ret := _I.GetGType(81, "Scanner")
	return ret
}

// g_scanner_cur_line
//
// [ result ] trans: nothing
//
func (v Scanner) CurLine() (result uint32) {
	iv, err := _I.Get(491, "Scanner", "cur_line", 232, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_scanner_cur_position
//
// [ result ] trans: nothing
//
func (v Scanner) CurPosition() (result uint32) {
	iv, err := _I.Get(492, "Scanner", "cur_position", 232, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_scanner_cur_token
//
// [ result ] trans: nothing
//
func (v Scanner) CurToken() (result TokenTypeEnum) {
	iv, err := _I.Get(493, "Scanner", "cur_token", 232, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TokenTypeEnum(ret.Int())
	return
}

// g_scanner_destroy
//
func (v Scanner) Destroy() {
	iv, err := _I.Get(494, "Scanner", "destroy", 232, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_scanner_eof
//
// [ result ] trans: nothing
//
func (v Scanner) Eof() (result bool) {
	iv, err := _I.Get(495, "Scanner", "eof", 232, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_scanner_get_next_token
//
// [ result ] trans: nothing
//
func (v Scanner) GetNextToken() (result TokenTypeEnum) {
	iv, err := _I.Get(496, "Scanner", "get_next_token", 232, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TokenTypeEnum(ret.Int())
	return
}

// g_scanner_input_file
//
// [ input_fd ] trans: nothing
//
func (v Scanner) InputFile(input_fd int32) {
	iv, err := _I.Get(497, "Scanner", "input_file", 232, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_input_fd := gi.NewInt32Argument(input_fd)
	args := []gi.Argument{arg_v, arg_input_fd}
	iv.Call(args, nil, nil)
}

// g_scanner_input_text
//
// [ text ] trans: nothing
//
// [ text_len ] trans: nothing
//
func (v Scanner) InputText(text string, text_len uint32) {
	iv, err := _I.Get(498, "Scanner", "input_text", 232, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_text_len := gi.NewUint32Argument(text_len)
	args := []gi.Argument{arg_v, arg_text, arg_text_len}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// g_scanner_lookup_symbol
//
// [ symbol ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Scanner) LookupSymbol(symbol string) (result unsafe.Pointer) {
	iv, err := _I.Get(499, "Scanner", "lookup_symbol", 232, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_symbol := gi.NewStringArgument(c_symbol)
	args := []gi.Argument{arg_v, arg_symbol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_symbol)
	result = ret.Pointer()
	return
}

// g_scanner_peek_next_token
//
// [ result ] trans: nothing
//
func (v Scanner) PeekNextToken() (result TokenTypeEnum) {
	iv, err := _I.Get(500, "Scanner", "peek_next_token", 232, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TokenTypeEnum(ret.Int())
	return
}

// g_scanner_scope_add_symbol
//
// [ scope_id ] trans: nothing
//
// [ symbol ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Scanner) ScopeAddSymbol(scope_id uint32, symbol string, value unsafe.Pointer) {
	iv, err := _I.Get(501, "Scanner", "scope_add_symbol", 232, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	arg_symbol := gi.NewStringArgument(c_symbol)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_scope_id, arg_symbol, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_symbol)
}

// g_scanner_scope_lookup_symbol
//
// [ scope_id ] trans: nothing
//
// [ symbol ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Scanner) ScopeLookupSymbol(scope_id uint32, symbol string) (result unsafe.Pointer) {
	iv, err := _I.Get(502, "Scanner", "scope_lookup_symbol", 232, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	arg_symbol := gi.NewStringArgument(c_symbol)
	args := []gi.Argument{arg_v, arg_scope_id, arg_symbol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_symbol)
	result = ret.Pointer()
	return
}

// g_scanner_scope_remove_symbol
//
// [ scope_id ] trans: nothing
//
// [ symbol ] trans: nothing
//
func (v Scanner) ScopeRemoveSymbol(scope_id uint32, symbol string) {
	iv, err := _I.Get(503, "Scanner", "scope_remove_symbol", 232, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	arg_symbol := gi.NewStringArgument(c_symbol)
	args := []gi.Argument{arg_v, arg_scope_id, arg_symbol}
	iv.Call(args, nil, nil)
	gi.Free(c_symbol)
}

// g_scanner_set_scope
//
// [ scope_id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Scanner) SetScope(scope_id uint32) (result uint32) {
	iv, err := _I.Get(504, "Scanner", "set_scope", 232, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	args := []gi.Argument{arg_v, arg_scope_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_scanner_sync_file_offset
//
func (v Scanner) SyncFileOffset() {
	iv, err := _I.Get(505, "Scanner", "sync_file_offset", 232, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_scanner_unexp_token
//
// [ expected_token ] trans: nothing
//
// [ identifier_spec ] trans: nothing
//
// [ symbol_spec ] trans: nothing
//
// [ symbol_name ] trans: nothing
//
// [ message ] trans: nothing
//
// [ is_error ] trans: nothing
//
func (v Scanner) UnexpToken(expected_token TokenTypeEnum, identifier_spec string, symbol_spec string, symbol_name string, message string, is_error int32) {
	iv, err := _I.Get(506, "Scanner", "unexp_token", 232, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_identifier_spec := gi.CString(identifier_spec)
	c_symbol_spec := gi.CString(symbol_spec)
	c_symbol_name := gi.CString(symbol_name)
	c_message := gi.CString(message)
	arg_v := gi.NewPointerArgument(v.P)
	arg_expected_token := gi.NewIntArgument(int(expected_token))
	arg_identifier_spec := gi.NewStringArgument(c_identifier_spec)
	arg_symbol_spec := gi.NewStringArgument(c_symbol_spec)
	arg_symbol_name := gi.NewStringArgument(c_symbol_name)
	arg_message := gi.NewStringArgument(c_message)
	arg_is_error := gi.NewInt32Argument(is_error)
	args := []gi.Argument{arg_v, arg_expected_token, arg_identifier_spec, arg_symbol_spec, arg_symbol_name, arg_message, arg_is_error}
	iv.Call(args, nil, nil)
	gi.Free(c_identifier_spec)
	gi.Free(c_symbol_spec)
	gi.Free(c_symbol_name)
	gi.Free(c_message)
}

// Struct ScannerConfig
type ScannerConfig struct {
	P unsafe.Pointer
}

const SizeOfStructScannerConfig = 128

func ScannerConfigGetType() gi.GType {
	ret := _I.GetGType(82, "ScannerConfig")
	return ret
}

type ScannerMsgFuncStruct struct {
	F_scanner Scanner
	F_message string
	F_error   bool
}

func GetPointer_myScannerMsgFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibScannerMsgFunc())
}

//export myGLibScannerMsgFunc
func myGLibScannerMsgFunc(scanner *C.GScanner, message *C.gchar, error C.gboolean) {
	// TODO: not found user_data
}

// Enum SeekType
type SeekTypeEnum int

const (
	SeekTypeCur SeekTypeEnum = 0
	SeekTypeSet SeekTypeEnum = 1
	SeekTypeEnd SeekTypeEnum = 2
)

func SeekTypeGetType() gi.GType {
	ret := _I.GetGType(83, "SeekType")
	return ret
}

// Struct Sequence
type Sequence struct {
	P unsafe.Pointer
}

func SequenceGetType() gi.GType {
	ret := _I.GetGType(84, "Sequence")
	return ret
}

// g_sequence_append
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Sequence) Append(data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(507, "Sequence", "append", 236, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_free
//
func (v Sequence) Free() {
	iv, err := _I.Get(508, "Sequence", "free", 236, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_sequence_get_begin_iter
//
// [ result ] trans: nothing
//
func (v Sequence) GetBeginIter() (result SequenceIter) {
	iv, err := _I.Get(509, "Sequence", "get_begin_iter", 236, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_get_end_iter
//
// [ result ] trans: nothing
//
func (v Sequence) GetEndIter() (result SequenceIter) {
	iv, err := _I.Get(510, "Sequence", "get_end_iter", 236, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_get_iter_at_pos
//
// [ pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Sequence) GetIterAtPos(pos int32) (result SequenceIter) {
	iv, err := _I.Get(511, "Sequence", "get_iter_at_pos", 236, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt32Argument(pos)
	args := []gi.Argument{arg_v, arg_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_get_length
//
// [ result ] trans: nothing
//
func (v Sequence) GetLength() (result int32) {
	iv, err := _I.Get(512, "Sequence", "get_length", 236, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_sequence_is_empty
//
// [ result ] trans: nothing
//
func (v Sequence) IsEmpty() (result bool) {
	iv, err := _I.Get(513, "Sequence", "is_empty", 236, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_sequence_prepend
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Sequence) Prepend(data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(514, "Sequence", "prepend", 236, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_get
//
// [ iter ] trans: nothing
//
// [ result ] trans: nothing
//
func SequenceGet1(iter SequenceIter) (result unsafe.Pointer) {
	iv, err := _I.Get(515, "Sequence", "get", 236, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_sequence_insert_before
//
// [ iter ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func SequenceInsertBefore1(iter SequenceIter, data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(516, "Sequence", "insert_before", 236, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_move
//
// [ src ] trans: nothing
//
// [ dest ] trans: nothing
//
func SequenceMove1(src SequenceIter, dest SequenceIter) {
	iv, err := _I.Get(517, "Sequence", "move", 236, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src := gi.NewPointerArgument(src.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_src, arg_dest}
	iv.Call(args, nil, nil)
}

// g_sequence_move_range
//
// [ dest ] trans: nothing
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
func SequenceMoveRange1(dest SequenceIter, begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(518, "Sequence", "move_range", 236, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_dest, arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_range_get_midpoint
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func SequenceRangeGetMidpoint1(begin SequenceIter, end SequenceIter) (result SequenceIter) {
	iv, err := _I.Get(519, "Sequence", "range_get_midpoint", 236, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_remove
//
// [ iter ] trans: nothing
//
func SequenceRemove1(iter SequenceIter) {
	iv, err := _I.Get(520, "Sequence", "remove", 236, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	iv.Call(args, nil, nil)
}

// g_sequence_remove_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
func SequenceRemoveRange1(begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(521, "Sequence", "remove_range", 236, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_set
//
// [ iter ] trans: nothing
//
// [ data ] trans: nothing
//
func SequenceSet1(iter SequenceIter, data unsafe.Pointer) {
	iv, err := _I.Get(522, "Sequence", "set", 236, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	iv.Call(args, nil, nil)
}

// g_sequence_swap
//
// [ a ] trans: nothing
//
// [ b ] trans: nothing
//
func SequenceSwap1(a SequenceIter, b SequenceIter) {
	iv, err := _I.Get(523, "Sequence", "swap", 236, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewPointerArgument(a.P)
	arg_b := gi.NewPointerArgument(b.P)
	args := []gi.Argument{arg_a, arg_b}
	iv.Call(args, nil, nil)
}

// Struct SequenceIter
type SequenceIter struct {
	P unsafe.Pointer
}

func SequenceIterGetType() gi.GType {
	ret := _I.GetGType(85, "SequenceIter")
	return ret
}

// g_sequence_iter_compare
//
// [ b ] trans: nothing
//
// [ result ] trans: nothing
//
func (v SequenceIter) Compare(b SequenceIter) (result int32) {
	iv, err := _I.Get(524, "SequenceIter", "compare", 237, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_b := gi.NewPointerArgument(b.P)
	args := []gi.Argument{arg_v, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_sequence_iter_get_position
//
// [ result ] trans: nothing
//
func (v SequenceIter) GetPosition() (result int32) {
	iv, err := _I.Get(525, "SequenceIter", "get_position", 237, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_sequence_iter_get_sequence
//
// [ result ] trans: nothing
//
func (v SequenceIter) GetSequence() (result Sequence) {
	iv, err := _I.Get(526, "SequenceIter", "get_sequence", 237, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_iter_is_begin
//
// [ result ] trans: nothing
//
func (v SequenceIter) IsBegin() (result bool) {
	iv, err := _I.Get(527, "SequenceIter", "is_begin", 237, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_sequence_iter_is_end
//
// [ result ] trans: nothing
//
func (v SequenceIter) IsEnd() (result bool) {
	iv, err := _I.Get(528, "SequenceIter", "is_end", 237, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_sequence_iter_move
//
// [ delta ] trans: nothing
//
// [ result ] trans: nothing
//
func (v SequenceIter) Move(delta int32) (result SequenceIter) {
	iv, err := _I.Get(529, "SequenceIter", "move", 237, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_delta := gi.NewInt32Argument(delta)
	args := []gi.Argument{arg_v, arg_delta}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_iter_next
//
// [ result ] trans: nothing
//
func (v SequenceIter) Next() (result SequenceIter) {
	iv, err := _I.Get(530, "SequenceIter", "next", 237, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_iter_prev
//
// [ result ] trans: nothing
//
func (v SequenceIter) Prev() (result SequenceIter) {
	iv, err := _I.Get(531, "SequenceIter", "prev", 237, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

type SequenceIterCompareFuncStruct struct {
	F_a    SequenceIter
	F_b    SequenceIter
	F_data unsafe.Pointer
}

func GetPointer_mySequenceIterCompareFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibSequenceIterCompareFunc())
}

//export myGLibSequenceIterCompareFunc
func myGLibSequenceIterCompareFunc(a *C.GSequenceIter, b *C.GSequenceIter, data C.gpointer) {
	// TODO: not found user_data
}

// Enum ShellError
type ShellErrorEnum int

const (
	ShellErrorBadQuoting  ShellErrorEnum = 0
	ShellErrorEmptyString ShellErrorEnum = 1
	ShellErrorFailed      ShellErrorEnum = 2
)

func ShellErrorGetType() gi.GType {
	ret := _I.GetGType(86, "ShellError")
	return ret
}

// Enum SliceConfig
type SliceConfigEnum int

const (
	SliceConfigAlwaysMalloc      SliceConfigEnum = 1
	SliceConfigBypassMagazines   SliceConfigEnum = 2
	SliceConfigWorkingSetMsecs   SliceConfigEnum = 3
	SliceConfigColorIncrement    SliceConfigEnum = 4
	SliceConfigChunkSizes        SliceConfigEnum = 5
	SliceConfigContentionCounter SliceConfigEnum = 6
)

func SliceConfigGetType() gi.GType {
	ret := _I.GetGType(87, "SliceConfig")
	return ret
}

// Struct Source
type Source struct {
	P unsafe.Pointer
}

const SizeOfStructSource = 96

func SourceGetType() gi.GType {
	ret := _I.GetGType(88, "Source")
	return ret
}

// g_source_new
//
// [ source_funcs ] trans: nothing
//
// [ struct_size ] trans: nothing
//
// [ result ] trans: everything
//
func NewSource(source_funcs SourceFuncs, struct_size uint32) (result Source) {
	iv, err := _I.Get(532, "Source", "new", 241, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_source_funcs := gi.NewPointerArgument(source_funcs.P)
	arg_struct_size := gi.NewUint32Argument(struct_size)
	args := []gi.Argument{arg_source_funcs, arg_struct_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_source_add_child_source
//
// [ child_source ] trans: nothing
//
func (v Source) AddChildSource(child_source Source) {
	iv, err := _I.Get(533, "Source", "add_child_source", 241, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child_source := gi.NewPointerArgument(child_source.P)
	args := []gi.Argument{arg_v, arg_child_source}
	iv.Call(args, nil, nil)
}

// g_source_add_poll
//
// [ fd ] trans: nothing
//
func (v Source) AddPoll(fd PollFD) {
	iv, err := _I.Get(534, "Source", "add_poll", 241, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// g_source_add_unix_fd
//
// [ fd ] trans: nothing
//
// [ events ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Source) AddUnixFd(fd int32, events IOConditionFlags) (result unsafe.Pointer) {
	iv, err := _I.Get(535, "Source", "add_unix_fd", 241, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewInt32Argument(fd)
	arg_events := gi.NewIntArgument(int(events))
	args := []gi.Argument{arg_v, arg_fd, arg_events}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_source_attach
//
// [ context ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Source) Attach(context MainContext) (result uint32) {
	iv, err := _I.Get(536, "Source", "attach", 241, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(context.P)
	args := []gi.Argument{arg_v, arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_source_destroy
//
func (v Source) Destroy() {
	iv, err := _I.Get(537, "Source", "destroy", 241, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_source_get_can_recurse
//
// [ result ] trans: nothing
//
func (v Source) GetCanRecurse() (result bool) {
	iv, err := _I.Get(538, "Source", "get_can_recurse", 241, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_get_context
//
// [ result ] trans: nothing
//
func (v Source) GetContext() (result MainContext) {
	iv, err := _I.Get(539, "Source", "get_context", 241, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_source_get_current_time
//
// [ timeval ] trans: nothing
//
func (v Source) GetCurrentTime(timeval TimeVal) {
	iv, err := _I.Get(540, "Source", "get_current_time", 241, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeval := gi.NewPointerArgument(timeval.P)
	args := []gi.Argument{arg_v, arg_timeval}
	iv.Call(args, nil, nil)
}

// g_source_get_id
//
// [ result ] trans: nothing
//
func (v Source) GetId() (result uint32) {
	iv, err := _I.Get(541, "Source", "get_id", 241, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_source_get_name
//
// [ result ] trans: nothing
//
func (v Source) GetName() (result string) {
	iv, err := _I.Get(542, "Source", "get_name", 241, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_source_get_priority
//
// [ result ] trans: nothing
//
func (v Source) GetPriority() (result int32) {
	iv, err := _I.Get(543, "Source", "get_priority", 241, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_source_get_ready_time
//
// [ result ] trans: nothing
//
func (v Source) GetReadyTime() (result int64) {
	iv, err := _I.Get(544, "Source", "get_ready_time", 241, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_source_get_time
//
// [ result ] trans: nothing
//
func (v Source) GetTime() (result int64) {
	iv, err := _I.Get(545, "Source", "get_time", 241, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_source_is_destroyed
//
// [ result ] trans: nothing
//
func (v Source) IsDestroyed() (result bool) {
	iv, err := _I.Get(546, "Source", "is_destroyed", 241, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_modify_unix_fd
//
// [ tag ] trans: nothing
//
// [ new_events ] trans: nothing
//
func (v Source) ModifyUnixFd(tag unsafe.Pointer, new_events IOConditionFlags) {
	iv, err := _I.Get(547, "Source", "modify_unix_fd", 241, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tag)
	arg_new_events := gi.NewIntArgument(int(new_events))
	args := []gi.Argument{arg_v, arg_tag, arg_new_events}
	iv.Call(args, nil, nil)
}

// g_source_query_unix_fd
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Source) QueryUnixFd(tag unsafe.Pointer) (result IOConditionFlags) {
	iv, err := _I.Get(548, "Source", "query_unix_fd", 241, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tag)
	args := []gi.Argument{arg_v, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOConditionFlags(ret.Int())
	return
}

// g_source_ref
//
// [ result ] trans: everything
//
func (v Source) Ref() (result Source) {
	iv, err := _I.Get(549, "Source", "ref", 241, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_source_remove_child_source
//
// [ child_source ] trans: nothing
//
func (v Source) RemoveChildSource(child_source Source) {
	iv, err := _I.Get(550, "Source", "remove_child_source", 241, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child_source := gi.NewPointerArgument(child_source.P)
	args := []gi.Argument{arg_v, arg_child_source}
	iv.Call(args, nil, nil)
}

// g_source_remove_poll
//
// [ fd ] trans: nothing
//
func (v Source) RemovePoll(fd PollFD) {
	iv, err := _I.Get(551, "Source", "remove_poll", 241, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// g_source_remove_unix_fd
//
// [ tag ] trans: nothing
//
func (v Source) RemoveUnixFd(tag unsafe.Pointer) {
	iv, err := _I.Get(552, "Source", "remove_unix_fd", 241, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tag)
	args := []gi.Argument{arg_v, arg_tag}
	iv.Call(args, nil, nil)
}

// g_source_set_callback
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Source) SetCallback(func1 int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(553, "Source", "set_callback", 241, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_func1, arg_data, arg_notify}
	iv.Call(args, nil, nil)
}

// g_source_set_callback_indirect
//
// [ callback_data ] trans: nothing
//
// [ callback_funcs ] trans: nothing
//
func (v Source) SetCallbackIndirect(callback_data unsafe.Pointer, callback_funcs SourceCallbackFuncs) {
	iv, err := _I.Get(554, "Source", "set_callback_indirect", 241, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_callback_data := gi.NewPointerArgument(callback_data)
	arg_callback_funcs := gi.NewPointerArgument(callback_funcs.P)
	args := []gi.Argument{arg_v, arg_callback_data, arg_callback_funcs}
	iv.Call(args, nil, nil)
}

// g_source_set_can_recurse
//
// [ can_recurse ] trans: nothing
//
func (v Source) SetCanRecurse(can_recurse bool) {
	iv, err := _I.Get(555, "Source", "set_can_recurse", 241, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_can_recurse := gi.NewBoolArgument(can_recurse)
	args := []gi.Argument{arg_v, arg_can_recurse}
	iv.Call(args, nil, nil)
}

// g_source_set_funcs
//
// [ funcs ] trans: nothing
//
func (v Source) SetFuncs(funcs SourceFuncs) {
	iv, err := _I.Get(556, "Source", "set_funcs", 241, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_funcs := gi.NewPointerArgument(funcs.P)
	args := []gi.Argument{arg_v, arg_funcs}
	iv.Call(args, nil, nil)
}

// g_source_set_name
//
// [ name ] trans: nothing
//
func (v Source) SetName(name string) {
	iv, err := _I.Get(557, "Source", "set_name", 241, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// g_source_set_priority
//
// [ priority ] trans: nothing
//
func (v Source) SetPriority(priority int32) {
	iv, err := _I.Get(558, "Source", "set_priority", 241, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_priority}
	iv.Call(args, nil, nil)
}

// g_source_set_ready_time
//
// [ ready_time ] trans: nothing
//
func (v Source) SetReadyTime(ready_time int64) {
	iv, err := _I.Get(559, "Source", "set_ready_time", 241, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ready_time := gi.NewInt64Argument(ready_time)
	args := []gi.Argument{arg_v, arg_ready_time}
	iv.Call(args, nil, nil)
}

// g_source_unref
//
func (v Source) Unref() {
	iv, err := _I.Get(560, "Source", "unref", 241, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_source_remove
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func SourceRemove1(tag uint32) (result bool) {
	iv, err := _I.Get(561, "Source", "remove", 241, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tag := gi.NewUint32Argument(tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_funcs_user_data
//
// [ funcs ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func SourceRemoveByFuncsUserData1(funcs SourceFuncs, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(562, "Source", "remove_by_funcs_user_data", 241, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_funcs := gi.NewPointerArgument(funcs.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_funcs, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_user_data
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func SourceRemoveByUserData1(user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(563, "Source", "remove_by_user_data", 241, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_set_name_by_id
//
// [ tag ] trans: nothing
//
// [ name ] trans: nothing
//
func SourceSetNameById1(tag uint32, name string) {
	iv, err := _I.Get(564, "Source", "set_name_by_id", 241, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_tag := gi.NewUint32Argument(tag)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_tag, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// Struct SourceCallbackFuncs
type SourceCallbackFuncs struct {
	P unsafe.Pointer
}

const SizeOfStructSourceCallbackFuncs = 24

func SourceCallbackFuncsGetType() gi.GType {
	ret := _I.GetGType(89, "SourceCallbackFuncs")
	return ret
}

type SourceDummyMarshalStruct struct {
}

func GetPointer_mySourceDummyMarshal() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibSourceDummyMarshal())
}

//export myGLibSourceDummyMarshal
func myGLibSourceDummyMarshal() {
	// TODO: not found user_data
}

type SourceFuncStruct struct {
}

func GetPointer_mySourceFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibSourceFunc())
}

//export myGLibSourceFunc
func myGLibSourceFunc(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &SourceFuncStruct{}
	fn(args)
}

// Struct SourceFuncs
type SourceFuncs struct {
	P unsafe.Pointer
}

const SizeOfStructSourceFuncs = 48

func SourceFuncsGetType() gi.GType {
	ret := _I.GetGType(90, "SourceFuncs")
	return ret
}

// Struct SourcePrivate
type SourcePrivate struct {
	P unsafe.Pointer
}

func SourcePrivateGetType() gi.GType {
	ret := _I.GetGType(91, "SourcePrivate")
	return ret
}

type SpawnChildSetupFuncStruct struct {
}

func GetPointer_mySpawnChildSetupFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibSpawnChildSetupFunc())
}

//export myGLibSpawnChildSetupFunc
func myGLibSpawnChildSetupFunc(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &SpawnChildSetupFuncStruct{}
	fn(args)
}

// Enum SpawnError
type SpawnErrorEnum int

const (
	SpawnErrorFork        SpawnErrorEnum = 0
	SpawnErrorRead        SpawnErrorEnum = 1
	SpawnErrorChdir       SpawnErrorEnum = 2
	SpawnErrorAcces       SpawnErrorEnum = 3
	SpawnErrorPerm        SpawnErrorEnum = 4
	SpawnErrorTooBig      SpawnErrorEnum = 5
	SpawnError2big        SpawnErrorEnum = 5
	SpawnErrorNoexec      SpawnErrorEnum = 6
	SpawnErrorNametoolong SpawnErrorEnum = 7
	SpawnErrorNoent       SpawnErrorEnum = 8
	SpawnErrorNomem       SpawnErrorEnum = 9
	SpawnErrorNotdir      SpawnErrorEnum = 10
	SpawnErrorLoop        SpawnErrorEnum = 11
	SpawnErrorTxtbusy     SpawnErrorEnum = 12
	SpawnErrorIo          SpawnErrorEnum = 13
	SpawnErrorNfile       SpawnErrorEnum = 14
	SpawnErrorMfile       SpawnErrorEnum = 15
	SpawnErrorInval       SpawnErrorEnum = 16
	SpawnErrorIsdir       SpawnErrorEnum = 17
	SpawnErrorLibbad      SpawnErrorEnum = 18
	SpawnErrorFailed      SpawnErrorEnum = 19
)

func SpawnErrorGetType() gi.GType {
	ret := _I.GetGType(92, "SpawnError")
	return ret
}

// Flags SpawnFlags
type SpawnFlags int

const (
	SpawnFlagsDefault              SpawnFlags = 0
	SpawnFlagsLeaveDescriptorsOpen SpawnFlags = 1
	SpawnFlagsDoNotReapChild       SpawnFlags = 2
	SpawnFlagsSearchPath           SpawnFlags = 4
	SpawnFlagsStdoutToDevNull      SpawnFlags = 8
	SpawnFlagsStderrToDevNull      SpawnFlags = 16
	SpawnFlagsChildInheritsStdin   SpawnFlags = 32
	SpawnFlagsFileAndArgvZero      SpawnFlags = 64
	SpawnFlagsSearchPathFromEnvp   SpawnFlags = 128
	SpawnFlagsCloexecPipes         SpawnFlags = 256
)

func SpawnFlagsGetType() gi.GType {
	ret := _I.GetGType(93, "SpawnFlags")
	return ret
}

// Struct StatBuf
type StatBuf struct {
	P unsafe.Pointer
}

func StatBufGetType() gi.GType {
	ret := _I.GetGType(94, "StatBuf")
	return ret
}

// Struct String
type String struct {
	P unsafe.Pointer
}

const SizeOfStructString = 24

func StringGetType() gi.GType {
	ret := _I.GetGType(95, "String")
	return ret
}

// g_string_append
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Append(val string) (result String) {
	iv, err := _I.Get(565, "String", "append", 251, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_append_c
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) AppendC(c int8) (result String) {
	iv, err := _I.Get(566, "String", "append_c", 251, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_v, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_append_len
//
// [ val ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) AppendLen(val string, len1 int64) (result String) {
	iv, err := _I.Get(567, "String", "append_len", 251, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_append_unichar
//
// [ wc ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) AppendUnichar(wc rune) (result String) {
	iv, err := _I.Get(568, "String", "append_unichar", 251, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wc := gi.NewUint32Argument(uint32(wc))
	args := []gi.Argument{arg_v, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_append_uri_escaped
//
// [ unescaped ] trans: nothing
//
// [ reserved_chars_allowed ] trans: nothing
//
// [ allow_utf8 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) AppendUriEscaped(unescaped string, reserved_chars_allowed string, allow_utf8 bool) (result String) {
	iv, err := _I.Get(569, "String", "append_uri_escaped", 251, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_unescaped := gi.CString(unescaped)
	c_reserved_chars_allowed := gi.CString(reserved_chars_allowed)
	arg_v := gi.NewPointerArgument(v.P)
	arg_unescaped := gi.NewStringArgument(c_unescaped)
	arg_reserved_chars_allowed := gi.NewStringArgument(c_reserved_chars_allowed)
	arg_allow_utf8 := gi.NewBoolArgument(allow_utf8)
	args := []gi.Argument{arg_v, arg_unescaped, arg_reserved_chars_allowed, arg_allow_utf8}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_unescaped)
	gi.Free(c_reserved_chars_allowed)
	result.P = ret.Pointer()
	return
}

// g_string_ascii_down
//
// [ result ] trans: nothing
//
func (v String) AsciiDown() (result String) {
	iv, err := _I.Get(570, "String", "ascii_down", 251, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_ascii_up
//
// [ result ] trans: nothing
//
func (v String) AsciiUp() (result String) {
	iv, err := _I.Get(571, "String", "ascii_up", 251, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_assign
//
// [ rval ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Assign(rval string) (result String) {
	iv, err := _I.Get(572, "String", "assign", 251, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_rval := gi.CString(rval)
	arg_v := gi.NewPointerArgument(v.P)
	arg_rval := gi.NewStringArgument(c_rval)
	args := []gi.Argument{arg_v, arg_rval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_rval)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_string_down
//
// [ result ] trans: nothing
//
func (v String) Down() (result String) {
	iv, err := _I.Get(573, "String", "down", 251, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_equal
//
// [ v2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Equal(v2 String) (result bool) {
	iv, err := _I.Get(574, "String", "equal", 251, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v2 := gi.NewPointerArgument(v2.P)
	args := []gi.Argument{arg_v, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_string_erase
//
// [ pos ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Erase(pos int64, len1 int64) (result String) {
	iv, err := _I.Get(575, "String", "erase", 251, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_pos, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_free
//
// [ free_segment ] trans: nothing
//
// [ result ] trans: everything
//
func (v String) Free(free_segment bool) (result string) {
	iv, err := _I.Get(576, "String", "free", 251, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_free_segment := gi.NewBoolArgument(free_segment)
	args := []gi.Argument{arg_v, arg_free_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_string_free_to_bytes
//
// [ result ] trans: everything
//
func (v String) FreeToBytes() (result Bytes) {
	iv, err := _I.Get(577, "String", "free_to_bytes", 251, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_hash
//
// [ result ] trans: nothing
//
func (v String) Hash() (result uint32) {
	iv, err := _I.Get(578, "String", "hash", 251, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_string_insert
//
// [ pos ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Insert(pos int64, val string) (result String) {
	iv, err := _I.Get(579, "String", "insert", 251, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_pos, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_insert_c
//
// [ pos ] trans: nothing
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) InsertC(pos int64, c int8) (result String) {
	iv, err := _I.Get(580, "String", "insert_c", 251, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_v, arg_pos, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_insert_len
//
// [ pos ] trans: nothing
//
// [ val ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) InsertLen(pos int64, val string, len1 int64) (result String) {
	iv, err := _I.Get(581, "String", "insert_len", 251, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_pos, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_insert_unichar
//
// [ pos ] trans: nothing
//
// [ wc ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) InsertUnichar(pos int64, wc rune) (result String) {
	iv, err := _I.Get(582, "String", "insert_unichar", 251, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_wc := gi.NewUint32Argument(uint32(wc))
	args := []gi.Argument{arg_v, arg_pos, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_overwrite
//
// [ pos ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Overwrite(pos uint64, val string) (result String) {
	iv, err := _I.Get(583, "String", "overwrite", 251, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_pos, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_overwrite_len
//
// [ pos ] trans: nothing
//
// [ val ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) OverwriteLen(pos uint64, val string, len1 int64) (result String) {
	iv, err := _I.Get(584, "String", "overwrite_len", 251, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_pos, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_prepend
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Prepend(val string) (result String) {
	iv, err := _I.Get(585, "String", "prepend", 251, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_prepend_c
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) PrependC(c int8) (result String) {
	iv, err := _I.Get(586, "String", "prepend_c", 251, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_v, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_prepend_len
//
// [ val ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) PrependLen(val string, len1 int64) (result String) {
	iv, err := _I.Get(587, "String", "prepend_len", 251, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_prepend_unichar
//
// [ wc ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) PrependUnichar(wc rune) (result String) {
	iv, err := _I.Get(588, "String", "prepend_unichar", 251, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wc := gi.NewUint32Argument(uint32(wc))
	args := []gi.Argument{arg_v, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_set_size
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) SetSize(len1 uint64) (result String) {
	iv, err := _I.Get(589, "String", "set_size", 251, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_v, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_truncate
//
// [ len1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v String) Truncate(len1 uint64) (result String) {
	iv, err := _I.Get(590, "String", "truncate", 251, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_v, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_string_up
//
// [ result ] trans: nothing
//
func (v String) Up() (result String) {
	iv, err := _I.Get(591, "String", "up", 251, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Struct StringChunk
type StringChunk struct {
	P unsafe.Pointer
}

func StringChunkGetType() gi.GType {
	ret := _I.GetGType(96, "StringChunk")
	return ret
}

// g_string_chunk_clear
//
func (v StringChunk) Clear() {
	iv, err := _I.Get(592, "StringChunk", "clear", 252, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_string_chunk_free
//
func (v StringChunk) Free() {
	iv, err := _I.Get(593, "StringChunk", "free", 252, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_string_chunk_insert
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func (v StringChunk) Insert(string string) (result string) {
	iv, err := _I.Get(594, "StringChunk", "insert", 252, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_string_chunk_insert_const
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func (v StringChunk) InsertConst(string string) (result string) {
	iv, err := _I.Get(595, "StringChunk", "insert_const", 252, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_string_chunk_insert_len
//
// [ string ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v StringChunk) InsertLen(string string, len1 int64) (result string) {
	iv, err := _I.Get(596, "StringChunk", "insert_len", 252, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_string, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// Struct TestCase
type TestCase struct {
	P unsafe.Pointer
}

func TestCaseGetType() gi.GType {
	ret := _I.GetGType(97, "TestCase")
	return ret
}

// Struct TestConfig
type TestConfig struct {
	P unsafe.Pointer
}

const SizeOfStructTestConfig = 24

func TestConfigGetType() gi.GType {
	ret := _I.GetGType(98, "TestConfig")
	return ret
}

type TestDataFuncStruct struct {
}

func GetPointer_myTestDataFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibTestDataFunc())
}

//export myGLibTestDataFunc
func myGLibTestDataFunc(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TestDataFuncStruct{}
	fn(args)
}

// Enum TestFileType
type TestFileTypeEnum int

const (
	TestFileTypeDist  TestFileTypeEnum = 0
	TestFileTypeBuilt TestFileTypeEnum = 1
)

func TestFileTypeGetType() gi.GType {
	ret := _I.GetGType(99, "TestFileType")
	return ret
}

type TestFixtureFuncStruct struct {
	F_fixture unsafe.Pointer
}

func GetPointer_myTestFixtureFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibTestFixtureFunc())
}

//export myGLibTestFixtureFunc
func myGLibTestFixtureFunc(fixture C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TestFixtureFuncStruct{
		F_fixture: unsafe.Pointer(fixture),
	}
	fn(args)
}

type TestFuncStruct struct {
}

func GetPointer_myTestFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibTestFunc())
}

//export myGLibTestFunc
func myGLibTestFunc() {
	// TODO: not found user_data
}

// Struct TestLogBuffer
type TestLogBuffer struct {
	P unsafe.Pointer
}

const SizeOfStructTestLogBuffer = 16

func TestLogBufferGetType() gi.GType {
	ret := _I.GetGType(100, "TestLogBuffer")
	return ret
}

// g_test_log_buffer_free
//
func (v TestLogBuffer) Free() {
	iv, err := _I.Get(597, "TestLogBuffer", "free", 264, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_test_log_buffer_push
//
// [ n_bytes ] trans: nothing
//
// [ bytes ] trans: nothing
//
func (v TestLogBuffer) Push(n_bytes uint32, bytes uint8) {
	iv, err := _I.Get(598, "TestLogBuffer", "push", 264, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_bytes := gi.NewUint32Argument(n_bytes)
	arg_bytes := gi.NewUint8Argument(bytes)
	args := []gi.Argument{arg_v, arg_n_bytes, arg_bytes}
	iv.Call(args, nil, nil)
}

type TestLogFatalFuncStruct struct {
	F_log_domain string
	F_log_level  LogLevelFlags
	F_message    string
}

func GetPointer_myTestLogFatalFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibTestLogFatalFunc())
}

//export myGLibTestLogFatalFunc
func myGLibTestLogFatalFunc(log_domain *C.gchar, log_level C.GLogLevelFlags, message *C.gchar, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TestLogFatalFuncStruct{
		F_log_domain: gi.GoString(unsafe.Pointer(log_domain)),
		F_log_level:  LogLevelFlags(log_level),
		F_message:    gi.GoString(unsafe.Pointer(message)),
	}
	fn(args)
}

// Struct TestLogMsg
type TestLogMsg struct {
	P unsafe.Pointer
}

const SizeOfStructTestLogMsg = 32

func TestLogMsgGetType() gi.GType {
	ret := _I.GetGType(101, "TestLogMsg")
	return ret
}

// g_test_log_msg_free
//
func (v TestLogMsg) Free() {
	iv, err := _I.Get(599, "TestLogMsg", "free", 266, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum TestLogType
type TestLogTypeEnum int

const (
	TestLogTypeNone        TestLogTypeEnum = 0
	TestLogTypeError       TestLogTypeEnum = 1
	TestLogTypeStartBinary TestLogTypeEnum = 2
	TestLogTypeListCase    TestLogTypeEnum = 3
	TestLogTypeSkipCase    TestLogTypeEnum = 4
	TestLogTypeStartCase   TestLogTypeEnum = 5
	TestLogTypeStopCase    TestLogTypeEnum = 6
	TestLogTypeMinResult   TestLogTypeEnum = 7
	TestLogTypeMaxResult   TestLogTypeEnum = 8
	TestLogTypeMessage     TestLogTypeEnum = 9
	TestLogTypeStartSuite  TestLogTypeEnum = 10
	TestLogTypeStopSuite   TestLogTypeEnum = 11
)

func TestLogTypeGetType() gi.GType {
	ret := _I.GetGType(102, "TestLogType")
	return ret
}

// Enum TestResult
type TestResultEnum int

const (
	TestResultSuccess    TestResultEnum = 0
	TestResultSkipped    TestResultEnum = 1
	TestResultFailure    TestResultEnum = 2
	TestResultIncomplete TestResultEnum = 3
)

func TestResultGetType() gi.GType {
	ret := _I.GetGType(103, "TestResult")
	return ret
}

// Flags TestSubprocessFlags
type TestSubprocessFlags int

const (
	TestSubprocessFlagsStdin  TestSubprocessFlags = 1
	TestSubprocessFlagsStdout TestSubprocessFlags = 2
	TestSubprocessFlagsStderr TestSubprocessFlags = 4
)

func TestSubprocessFlagsGetType() gi.GType {
	ret := _I.GetGType(104, "TestSubprocessFlags")
	return ret
}

// Struct TestSuite
type TestSuite struct {
	P unsafe.Pointer
}

func TestSuiteGetType() gi.GType {
	ret := _I.GetGType(105, "TestSuite")
	return ret
}

// g_test_suite_add
//
// [ test_case ] trans: nothing
//
func (v TestSuite) Add(test_case TestCase) {
	iv, err := _I.Get(600, "TestSuite", "add", 270, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_test_case := gi.NewPointerArgument(test_case.P)
	args := []gi.Argument{arg_v, arg_test_case}
	iv.Call(args, nil, nil)
}

// g_test_suite_add_suite
//
// [ nestedsuite ] trans: nothing
//
func (v TestSuite) AddSuite(nestedsuite TestSuite) {
	iv, err := _I.Get(601, "TestSuite", "add_suite", 270, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nestedsuite := gi.NewPointerArgument(nestedsuite.P)
	args := []gi.Argument{arg_v, arg_nestedsuite}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// Flags TestTrapFlags
type TestTrapFlags int

const (
	TestTrapFlagsSilenceStdout TestTrapFlags = 128
	TestTrapFlagsSilenceStderr TestTrapFlags = 256
	TestTrapFlagsInheritStdin  TestTrapFlags = 512
)

func TestTrapFlagsGetType() gi.GType {
	ret := _I.GetGType(106, "TestTrapFlags")
	return ret
}

// Struct Thread
type Thread struct {
	P unsafe.Pointer
}

func ThreadGetType() gi.GType {
	ret := _I.GetGType(107, "Thread")
	return ret
}

// g_thread_join
//
// [ result ] trans: nothing
//
func (v Thread) Join() (result unsafe.Pointer) {
	iv, err := _I.Get(602, "Thread", "join", 272, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_thread_ref
//
// [ result ] trans: everything
//
func (v Thread) Ref() (result Thread) {
	iv, err := _I.Get(603, "Thread", "ref", 272, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_thread_unref
//
func (v Thread) Unref() {
	iv, err := _I.Get(604, "Thread", "unref", 272, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_thread_exit
//
// [ retval ] trans: nothing
//
func ThreadExit1(retval unsafe.Pointer) {
	iv, err := _I.Get(606, "Thread", "exit", 272, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_retval := gi.NewPointerArgument(retval)
	args := []gi.Argument{arg_retval}
	iv.Call(args, nil, nil)
}

// Enum ThreadError
type ThreadErrorEnum int

const (
	ThreadErrorThreadErrorAgain ThreadErrorEnum = 0
)

func ThreadErrorGetType() gi.GType {
	ret := _I.GetGType(108, "ThreadError")
	return ret
}

type ThreadFuncStruct struct {
	F_data unsafe.Pointer
}

func GetPointer_myThreadFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibThreadFunc())
}

//export myGLibThreadFunc
func myGLibThreadFunc(data C.gpointer) {
	// TODO: not found user_data
}

// Struct ThreadPool
type ThreadPool struct {
	P unsafe.Pointer
}

const SizeOfStructThreadPool = 24

func ThreadPoolGetType() gi.GType {
	ret := _I.GetGType(109, "ThreadPool")
	return ret
}

// g_thread_pool_free
//
// [ immediate ] trans: nothing
//
// [ wait_ ] trans: nothing
//
func (v ThreadPool) Free(immediate bool, wait_ bool) {
	iv, err := _I.Get(609, "ThreadPool", "free", 275, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_immediate := gi.NewBoolArgument(immediate)
	arg_wait_ := gi.NewBoolArgument(wait_)
	args := []gi.Argument{arg_v, arg_immediate, arg_wait_}
	iv.Call(args, nil, nil)
}

// g_thread_pool_get_max_threads
//
// [ result ] trans: nothing
//
func (v ThreadPool) GetMaxThreads() (result int32) {
	iv, err := _I.Get(610, "ThreadPool", "get_max_threads", 275, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_thread_pool_get_num_threads
//
// [ result ] trans: nothing
//
func (v ThreadPool) GetNumThreads() (result uint32) {
	iv, err := _I.Get(611, "ThreadPool", "get_num_threads", 275, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_pool_move_to_front
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ThreadPool) MoveToFront(data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(612, "ThreadPool", "move_to_front", 275, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_thread_pool_push
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ThreadPool) Push(data unsafe.Pointer) (result bool, err error) {
	iv, err := _I.Get(613, "ThreadPool", "push", 275, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_thread_pool_set_max_threads
//
// [ max_threads ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ThreadPool) SetMaxThreads(max_threads int32) (result bool, err error) {
	iv, err := _I.Get(614, "ThreadPool", "set_max_threads", 275, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_threads := gi.NewInt32Argument(max_threads)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_max_threads, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_thread_pool_unprocessed
//
// [ result ] trans: nothing
//
func (v ThreadPool) Unprocessed() (result uint32) {
	iv, err := _I.Get(615, "ThreadPool", "unprocessed", 275, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_pool_set_max_idle_time
//
// [ interval ] trans: nothing
//
func ThreadPoolSetMaxIdleTime1(interval uint32) {
	iv, err := _I.Get(619, "ThreadPool", "set_max_idle_time", 275, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	iv.Call(args, nil, nil)
}

// g_thread_pool_set_max_unused_threads
//
// [ max_threads ] trans: nothing
//
func ThreadPoolSetMaxUnusedThreads1(max_threads int32) {
	iv, err := _I.Get(620, "ThreadPool", "set_max_unused_threads", 275, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_max_threads := gi.NewInt32Argument(max_threads)
	args := []gi.Argument{arg_max_threads}
	iv.Call(args, nil, nil)
}

// Enum TimeType
type TimeTypeEnum int

const (
	TimeTypeStandard  TimeTypeEnum = 0
	TimeTypeDaylight  TimeTypeEnum = 1
	TimeTypeUniversal TimeTypeEnum = 2
)

func TimeTypeGetType() gi.GType {
	ret := _I.GetGType(110, "TimeType")
	return ret
}

// Struct TimeVal
type TimeVal struct {
	P unsafe.Pointer
}

const SizeOfStructTimeVal = 16

func TimeValGetType() gi.GType {
	ret := _I.GetGType(111, "TimeVal")
	return ret
}

// g_time_val_add
//
// [ microseconds ] trans: nothing
//
func (v TimeVal) Add(microseconds int64) {
	iv, err := _I.Get(622, "TimeVal", "add", 277, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_microseconds := gi.NewInt64Argument(microseconds)
	args := []gi.Argument{arg_v, arg_microseconds}
	iv.Call(args, nil, nil)
}

// g_time_val_to_iso8601
//
// [ result ] trans: everything
//
func (v TimeVal) ToIso8601() (result string) {
	iv, err := _I.Get(623, "TimeVal", "to_iso8601", 277, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_time_val_from_iso8601
//
// [ iso_date ] trans: nothing
//
// [ time_ ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func TimeValFromIso86011(iso_date string, time_ TimeVal) (result bool) {
	iv, err := _I.Get(624, "TimeVal", "from_iso8601", 277, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_iso_date := gi.CString(iso_date)
	arg_iso_date := gi.NewStringArgument(c_iso_date)
	arg_time_ := gi.NewPointerArgument(time_.P)
	args := []gi.Argument{arg_iso_date, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_iso_date)
	result = ret.Bool()
	return
}

// Struct TimeZone
type TimeZone struct {
	P unsafe.Pointer
}

func TimeZoneGetType() gi.GType {
	ret := _I.GetGType(112, "TimeZone")
	return ret
}

// g_time_zone_new
//
// [ identifier ] trans: nothing
//
// [ result ] trans: everything
//
func NewTimeZone(identifier string) (result TimeZone) {
	iv, err := _I.Get(625, "TimeZone", "new", 278, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_identifier := gi.CString(identifier)
	arg_identifier := gi.NewStringArgument(c_identifier)
	args := []gi.Argument{arg_identifier}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_identifier)
	result.P = ret.Pointer()
	return
}

// g_time_zone_new_local
//
// [ result ] trans: everything
//
func NewTimeZoneLocal() (result TimeZone) {
	iv, err := _I.Get(626, "TimeZone", "new_local", 278, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_new_offset
//
// [ seconds ] trans: nothing
//
// [ result ] trans: everything
//
func NewTimeZoneOffset(seconds int32) (result TimeZone) {
	iv, err := _I.Get(627, "TimeZone", "new_offset", 278, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_seconds := gi.NewInt32Argument(seconds)
	args := []gi.Argument{arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_new_utc
//
// [ result ] trans: everything
//
func NewTimeZoneUtc() (result TimeZone) {
	iv, err := _I.Get(628, "TimeZone", "new_utc", 278, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_adjust_time
//
// [ type1 ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimeZone) AdjustTime(type1 TimeTypeEnum, time_ int64) (result int32) {
	iv, err := _I.Get(629, "TimeZone", "adjust_time", 278, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_time_ := gi.NewInt64Argument(time_)
	args := []gi.Argument{arg_v, arg_type1, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_time_zone_find_interval
//
// [ type1 ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimeZone) FindInterval(type1 TimeTypeEnum, time_ int64) (result int32) {
	iv, err := _I.Get(630, "TimeZone", "find_interval", 278, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_time_ := gi.NewInt64Argument(time_)
	args := []gi.Argument{arg_v, arg_type1, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_time_zone_get_abbreviation
//
// [ interval ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimeZone) GetAbbreviation(interval int32) (result string) {
	iv, err := _I.Get(631, "TimeZone", "get_abbreviation", 278, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_time_zone_get_identifier
//
// [ result ] trans: nothing
//
func (v TimeZone) GetIdentifier() (result string) {
	iv, err := _I.Get(632, "TimeZone", "get_identifier", 278, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_time_zone_get_offset
//
// [ interval ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimeZone) GetOffset(interval int32) (result int32) {
	iv, err := _I.Get(633, "TimeZone", "get_offset", 278, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_time_zone_is_dst
//
// [ interval ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimeZone) IsDst(interval int32) (result bool) {
	iv, err := _I.Get(634, "TimeZone", "is_dst", 278, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_time_zone_ref
//
// [ result ] trans: everything
//
func (v TimeZone) Ref() (result TimeZone) {
	iv, err := _I.Get(635, "TimeZone", "ref", 278, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_unref
//
func (v TimeZone) Unref() {
	iv, err := _I.Get(636, "TimeZone", "unref", 278, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Timer
type Timer struct {
	P unsafe.Pointer
}

func TimerGetType() gi.GType {
	ret := _I.GetGType(113, "Timer")
	return ret
}

// g_timer_continue
//
func (v Timer) Continue() {
	iv, err := _I.Get(637, "Timer", "continue", 279, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_destroy
//
func (v Timer) Destroy() {
	iv, err := _I.Get(638, "Timer", "destroy", 279, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_elapsed
//
// [ microseconds ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Timer) Elapsed(microseconds uint64) (result float64) {
	iv, err := _I.Get(639, "Timer", "elapsed", 279, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_microseconds := gi.NewUint64Argument(microseconds)
	args := []gi.Argument{arg_v, arg_microseconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_timer_reset
//
func (v Timer) Reset() {
	iv, err := _I.Get(640, "Timer", "reset", 279, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_start
//
func (v Timer) Start() {
	iv, err := _I.Get(641, "Timer", "start", 279, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_stop
//
func (v Timer) Stop() {
	iv, err := _I.Get(642, "Timer", "stop", 279, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum TokenType
type TokenTypeEnum int

const (
	TokenTypeEof            TokenTypeEnum = 0
	TokenTypeLeftParen      TokenTypeEnum = 40
	TokenTypeRightParen     TokenTypeEnum = 41
	TokenTypeLeftCurly      TokenTypeEnum = 123
	TokenTypeRightCurly     TokenTypeEnum = 125
	TokenTypeLeftBrace      TokenTypeEnum = 91
	TokenTypeRightBrace     TokenTypeEnum = 93
	TokenTypeEqualSign      TokenTypeEnum = 61
	TokenTypeComma          TokenTypeEnum = 44
	TokenTypeNone           TokenTypeEnum = 256
	TokenTypeError          TokenTypeEnum = 257
	TokenTypeChar           TokenTypeEnum = 258
	TokenTypeBinary         TokenTypeEnum = 259
	TokenTypeOctal          TokenTypeEnum = 260
	TokenTypeInt            TokenTypeEnum = 261
	TokenTypeHex            TokenTypeEnum = 262
	TokenTypeFloat          TokenTypeEnum = 263
	TokenTypeString         TokenTypeEnum = 264
	TokenTypeSymbol         TokenTypeEnum = 265
	TokenTypeIdentifier     TokenTypeEnum = 266
	TokenTypeIdentifierNull TokenTypeEnum = 267
	TokenTypeCommentSingle  TokenTypeEnum = 268
	TokenTypeCommentMulti   TokenTypeEnum = 269
)

func TokenTypeGetType() gi.GType {
	ret := _I.GetGType(114, "TokenType")
	return ret
}

// Union TokenValue
type TokenValue struct {
	P unsafe.Pointer
}

const SizeOfUnionTokenValue = 8

func TokenValueGetType() gi.GType {
	ret := _I.GetGType(115, "TokenValue")
	return ret
}

type TranslateFuncStruct struct {
	F_str  string
	F_data unsafe.Pointer
}

func GetPointer_myTranslateFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibTranslateFunc())
}

//export myGLibTranslateFunc
func myGLibTranslateFunc(str *C.gchar, data C.gpointer) {
	// TODO: not found user_data
}

// Deprecated
//
// Struct TrashStack
type TrashStack struct {
	P unsafe.Pointer
}

const SizeOfStructTrashStack = 8

func TrashStackGetType() gi.GType {
	ret := _I.GetGType(116, "TrashStack")
	return ret
}

// Deprecated
//
// g_trash_stack_height
//
// [ result ] trans: nothing
//
func (v TrashStack) Height() (result uint32) {
	iv, err := _I.Get(643, "TrashStack", "height", 283, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// g_trash_stack_peek
//
// [ result ] trans: nothing
//
func (v TrashStack) Peek() (result unsafe.Pointer) {
	iv, err := _I.Get(644, "TrashStack", "peek", 283, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Deprecated
//
// g_trash_stack_pop
//
// [ result ] trans: nothing
//
func (v TrashStack) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(645, "TrashStack", "pop", 283, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Deprecated
//
// g_trash_stack_push
//
// [ data_p ] trans: nothing
//
func (v TrashStack) Push(data_p unsafe.Pointer) {
	iv, err := _I.Get(646, "TrashStack", "push", 283, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data_p := gi.NewPointerArgument(data_p)
	args := []gi.Argument{arg_v, arg_data_p}
	iv.Call(args, nil, nil)
}

// Flags TraverseFlags
type TraverseFlags int

const (
	TraverseFlagsLeaves    TraverseFlags = 1
	TraverseFlagsNonLeaves TraverseFlags = 2
	TraverseFlagsAll       TraverseFlags = 3
	TraverseFlagsMask      TraverseFlags = 3
	TraverseFlagsLeafs     TraverseFlags = 1
	TraverseFlagsNonLeafs  TraverseFlags = 2
)

func TraverseFlagsGetType() gi.GType {
	ret := _I.GetGType(117, "TraverseFlags")
	return ret
}

type TraverseFuncStruct struct {
	F_key   unsafe.Pointer
	F_value unsafe.Pointer
	F_data  unsafe.Pointer
}

func GetPointer_myTraverseFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibTraverseFunc())
}

//export myGLibTraverseFunc
func myGLibTraverseFunc(key C.gpointer, value C.gpointer, data C.gpointer) {
	// TODO: not found user_data
}

// Enum TraverseType
type TraverseTypeEnum int

const (
	TraverseTypeInOrder    TraverseTypeEnum = 0
	TraverseTypePreOrder   TraverseTypeEnum = 1
	TraverseTypePostOrder  TraverseTypeEnum = 2
	TraverseTypeLevelOrder TraverseTypeEnum = 3
)

func TraverseTypeGetType() gi.GType {
	ret := _I.GetGType(118, "TraverseType")
	return ret
}

// Struct Tree
type Tree struct {
	P unsafe.Pointer
}

func TreeGetType() gi.GType {
	ret := _I.GetGType(119, "Tree")
	return ret
}

// g_tree_destroy
//
func (v Tree) Destroy() {
	iv, err := _I.Get(647, "Tree", "destroy", 287, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_tree_height
//
// [ result ] trans: nothing
//
func (v Tree) Height() (result int32) {
	iv, err := _I.Get(648, "Tree", "height", 287, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_tree_insert
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Tree) Insert(key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(649, "Tree", "insert", 287, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
}

// g_tree_lookup
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Tree) Lookup(key unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(650, "Tree", "lookup", 287, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_tree_lookup_extended
//
// [ lookup_key ] trans: nothing
//
// [ orig_key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Tree) LookupExtended(lookup_key unsafe.Pointer, orig_key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(651, "Tree", "lookup_extended", 287, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_orig_key := gi.NewPointerArgument(orig_key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_lookup_key, arg_orig_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_tree_nnodes
//
// [ result ] trans: nothing
//
func (v Tree) Nnodes() (result int32) {
	iv, err := _I.Get(652, "Tree", "nnodes", 287, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_tree_remove
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Tree) Remove(key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(653, "Tree", "remove", 287, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_tree_replace
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Tree) Replace(key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(654, "Tree", "replace", 287, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
}

// g_tree_steal
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Tree) Steal(key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(655, "Tree", "steal", 287, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_tree_unref
//
func (v Tree) Unref() {
	iv, err := _I.Get(656, "Tree", "unref", 287, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum UnicodeBreakType
type UnicodeBreakTypeEnum int

const (
	UnicodeBreakTypeMandatory                  UnicodeBreakTypeEnum = 0
	UnicodeBreakTypeCarriageReturn             UnicodeBreakTypeEnum = 1
	UnicodeBreakTypeLineFeed                   UnicodeBreakTypeEnum = 2
	UnicodeBreakTypeCombiningMark              UnicodeBreakTypeEnum = 3
	UnicodeBreakTypeSurrogate                  UnicodeBreakTypeEnum = 4
	UnicodeBreakTypeZeroWidthSpace             UnicodeBreakTypeEnum = 5
	UnicodeBreakTypeInseparable                UnicodeBreakTypeEnum = 6
	UnicodeBreakTypeNonBreakingGlue            UnicodeBreakTypeEnum = 7
	UnicodeBreakTypeContingent                 UnicodeBreakTypeEnum = 8
	UnicodeBreakTypeSpace                      UnicodeBreakTypeEnum = 9
	UnicodeBreakTypeAfter                      UnicodeBreakTypeEnum = 10
	UnicodeBreakTypeBefore                     UnicodeBreakTypeEnum = 11
	UnicodeBreakTypeBeforeAndAfter             UnicodeBreakTypeEnum = 12
	UnicodeBreakTypeHyphen                     UnicodeBreakTypeEnum = 13
	UnicodeBreakTypeNonStarter                 UnicodeBreakTypeEnum = 14
	UnicodeBreakTypeOpenPunctuation            UnicodeBreakTypeEnum = 15
	UnicodeBreakTypeClosePunctuation           UnicodeBreakTypeEnum = 16
	UnicodeBreakTypeQuotation                  UnicodeBreakTypeEnum = 17
	UnicodeBreakTypeExclamation                UnicodeBreakTypeEnum = 18
	UnicodeBreakTypeIdeographic                UnicodeBreakTypeEnum = 19
	UnicodeBreakTypeNumeric                    UnicodeBreakTypeEnum = 20
	UnicodeBreakTypeInfixSeparator             UnicodeBreakTypeEnum = 21
	UnicodeBreakTypeSymbol                     UnicodeBreakTypeEnum = 22
	UnicodeBreakTypeAlphabetic                 UnicodeBreakTypeEnum = 23
	UnicodeBreakTypePrefix                     UnicodeBreakTypeEnum = 24
	UnicodeBreakTypePostfix                    UnicodeBreakTypeEnum = 25
	UnicodeBreakTypeComplexContext             UnicodeBreakTypeEnum = 26
	UnicodeBreakTypeAmbiguous                  UnicodeBreakTypeEnum = 27
	UnicodeBreakTypeUnknown                    UnicodeBreakTypeEnum = 28
	UnicodeBreakTypeNextLine                   UnicodeBreakTypeEnum = 29
	UnicodeBreakTypeWordJoiner                 UnicodeBreakTypeEnum = 30
	UnicodeBreakTypeHangulLJamo                UnicodeBreakTypeEnum = 31
	UnicodeBreakTypeHangulVJamo                UnicodeBreakTypeEnum = 32
	UnicodeBreakTypeHangulTJamo                UnicodeBreakTypeEnum = 33
	UnicodeBreakTypeHangulLvSyllable           UnicodeBreakTypeEnum = 34
	UnicodeBreakTypeHangulLvtSyllable          UnicodeBreakTypeEnum = 35
	UnicodeBreakTypeCloseParanthesis           UnicodeBreakTypeEnum = 36
	UnicodeBreakTypeConditionalJapaneseStarter UnicodeBreakTypeEnum = 37
	UnicodeBreakTypeHebrewLetter               UnicodeBreakTypeEnum = 38
	UnicodeBreakTypeRegionalIndicator          UnicodeBreakTypeEnum = 39
	UnicodeBreakTypeEmojiBase                  UnicodeBreakTypeEnum = 40
	UnicodeBreakTypeEmojiModifier              UnicodeBreakTypeEnum = 41
	UnicodeBreakTypeZeroWidthJoiner            UnicodeBreakTypeEnum = 42
)

func UnicodeBreakTypeGetType() gi.GType {
	ret := _I.GetGType(120, "UnicodeBreakType")
	return ret
}

// Enum UnicodeScript
type UnicodeScriptEnum int

const (
	UnicodeScriptInvalidCode           UnicodeScriptEnum = -1
	UnicodeScriptCommon                UnicodeScriptEnum = 0
	UnicodeScriptInherited             UnicodeScriptEnum = 1
	UnicodeScriptArabic                UnicodeScriptEnum = 2
	UnicodeScriptArmenian              UnicodeScriptEnum = 3
	UnicodeScriptBengali               UnicodeScriptEnum = 4
	UnicodeScriptBopomofo              UnicodeScriptEnum = 5
	UnicodeScriptCherokee              UnicodeScriptEnum = 6
	UnicodeScriptCoptic                UnicodeScriptEnum = 7
	UnicodeScriptCyrillic              UnicodeScriptEnum = 8
	UnicodeScriptDeseret               UnicodeScriptEnum = 9
	UnicodeScriptDevanagari            UnicodeScriptEnum = 10
	UnicodeScriptEthiopic              UnicodeScriptEnum = 11
	UnicodeScriptGeorgian              UnicodeScriptEnum = 12
	UnicodeScriptGothic                UnicodeScriptEnum = 13
	UnicodeScriptGreek                 UnicodeScriptEnum = 14
	UnicodeScriptGujarati              UnicodeScriptEnum = 15
	UnicodeScriptGurmukhi              UnicodeScriptEnum = 16
	UnicodeScriptHan                   UnicodeScriptEnum = 17
	UnicodeScriptHangul                UnicodeScriptEnum = 18
	UnicodeScriptHebrew                UnicodeScriptEnum = 19
	UnicodeScriptHiragana              UnicodeScriptEnum = 20
	UnicodeScriptKannada               UnicodeScriptEnum = 21
	UnicodeScriptKatakana              UnicodeScriptEnum = 22
	UnicodeScriptKhmer                 UnicodeScriptEnum = 23
	UnicodeScriptLao                   UnicodeScriptEnum = 24
	UnicodeScriptLatin                 UnicodeScriptEnum = 25
	UnicodeScriptMalayalam             UnicodeScriptEnum = 26
	UnicodeScriptMongolian             UnicodeScriptEnum = 27
	UnicodeScriptMyanmar               UnicodeScriptEnum = 28
	UnicodeScriptOgham                 UnicodeScriptEnum = 29
	UnicodeScriptOldItalic             UnicodeScriptEnum = 30
	UnicodeScriptOriya                 UnicodeScriptEnum = 31
	UnicodeScriptRunic                 UnicodeScriptEnum = 32
	UnicodeScriptSinhala               UnicodeScriptEnum = 33
	UnicodeScriptSyriac                UnicodeScriptEnum = 34
	UnicodeScriptTamil                 UnicodeScriptEnum = 35
	UnicodeScriptTelugu                UnicodeScriptEnum = 36
	UnicodeScriptThaana                UnicodeScriptEnum = 37
	UnicodeScriptThai                  UnicodeScriptEnum = 38
	UnicodeScriptTibetan               UnicodeScriptEnum = 39
	UnicodeScriptCanadianAboriginal    UnicodeScriptEnum = 40
	UnicodeScriptYi                    UnicodeScriptEnum = 41
	UnicodeScriptTagalog               UnicodeScriptEnum = 42
	UnicodeScriptHanunoo               UnicodeScriptEnum = 43
	UnicodeScriptBuhid                 UnicodeScriptEnum = 44
	UnicodeScriptTagbanwa              UnicodeScriptEnum = 45
	UnicodeScriptBraille               UnicodeScriptEnum = 46
	UnicodeScriptCypriot               UnicodeScriptEnum = 47
	UnicodeScriptLimbu                 UnicodeScriptEnum = 48
	UnicodeScriptOsmanya               UnicodeScriptEnum = 49
	UnicodeScriptShavian               UnicodeScriptEnum = 50
	UnicodeScriptLinearB               UnicodeScriptEnum = 51
	UnicodeScriptTaiLe                 UnicodeScriptEnum = 52
	UnicodeScriptUgaritic              UnicodeScriptEnum = 53
	UnicodeScriptNewTaiLue             UnicodeScriptEnum = 54
	UnicodeScriptBuginese              UnicodeScriptEnum = 55
	UnicodeScriptGlagolitic            UnicodeScriptEnum = 56
	UnicodeScriptTifinagh              UnicodeScriptEnum = 57
	UnicodeScriptSylotiNagri           UnicodeScriptEnum = 58
	UnicodeScriptOldPersian            UnicodeScriptEnum = 59
	UnicodeScriptKharoshthi            UnicodeScriptEnum = 60
	UnicodeScriptUnknown               UnicodeScriptEnum = 61
	UnicodeScriptBalinese              UnicodeScriptEnum = 62
	UnicodeScriptCuneiform             UnicodeScriptEnum = 63
	UnicodeScriptPhoenician            UnicodeScriptEnum = 64
	UnicodeScriptPhagsPa               UnicodeScriptEnum = 65
	UnicodeScriptNko                   UnicodeScriptEnum = 66
	UnicodeScriptKayahLi               UnicodeScriptEnum = 67
	UnicodeScriptLepcha                UnicodeScriptEnum = 68
	UnicodeScriptRejang                UnicodeScriptEnum = 69
	UnicodeScriptSundanese             UnicodeScriptEnum = 70
	UnicodeScriptSaurashtra            UnicodeScriptEnum = 71
	UnicodeScriptCham                  UnicodeScriptEnum = 72
	UnicodeScriptOlChiki               UnicodeScriptEnum = 73
	UnicodeScriptVai                   UnicodeScriptEnum = 74
	UnicodeScriptCarian                UnicodeScriptEnum = 75
	UnicodeScriptLycian                UnicodeScriptEnum = 76
	UnicodeScriptLydian                UnicodeScriptEnum = 77
	UnicodeScriptAvestan               UnicodeScriptEnum = 78
	UnicodeScriptBamum                 UnicodeScriptEnum = 79
	UnicodeScriptEgyptianHieroglyphs   UnicodeScriptEnum = 80
	UnicodeScriptImperialAramaic       UnicodeScriptEnum = 81
	UnicodeScriptInscriptionalPahlavi  UnicodeScriptEnum = 82
	UnicodeScriptInscriptionalParthian UnicodeScriptEnum = 83
	UnicodeScriptJavanese              UnicodeScriptEnum = 84
	UnicodeScriptKaithi                UnicodeScriptEnum = 85
	UnicodeScriptLisu                  UnicodeScriptEnum = 86
	UnicodeScriptMeeteiMayek           UnicodeScriptEnum = 87
	UnicodeScriptOldSouthArabian       UnicodeScriptEnum = 88
	UnicodeScriptOldTurkic             UnicodeScriptEnum = 89
	UnicodeScriptSamaritan             UnicodeScriptEnum = 90
	UnicodeScriptTaiTham               UnicodeScriptEnum = 91
	UnicodeScriptTaiViet               UnicodeScriptEnum = 92
	UnicodeScriptBatak                 UnicodeScriptEnum = 93
	UnicodeScriptBrahmi                UnicodeScriptEnum = 94
	UnicodeScriptMandaic               UnicodeScriptEnum = 95
	UnicodeScriptChakma                UnicodeScriptEnum = 96
	UnicodeScriptMeroiticCursive       UnicodeScriptEnum = 97
	UnicodeScriptMeroiticHieroglyphs   UnicodeScriptEnum = 98
	UnicodeScriptMiao                  UnicodeScriptEnum = 99
	UnicodeScriptSharada               UnicodeScriptEnum = 100
	UnicodeScriptSoraSompeng           UnicodeScriptEnum = 101
	UnicodeScriptTakri                 UnicodeScriptEnum = 102
	UnicodeScriptBassaVah              UnicodeScriptEnum = 103
	UnicodeScriptCaucasianAlbanian     UnicodeScriptEnum = 104
	UnicodeScriptDuployan              UnicodeScriptEnum = 105
	UnicodeScriptElbasan               UnicodeScriptEnum = 106
	UnicodeScriptGrantha               UnicodeScriptEnum = 107
	UnicodeScriptKhojki                UnicodeScriptEnum = 108
	UnicodeScriptKhudawadi             UnicodeScriptEnum = 109
	UnicodeScriptLinearA               UnicodeScriptEnum = 110
	UnicodeScriptMahajani              UnicodeScriptEnum = 111
	UnicodeScriptManichaean            UnicodeScriptEnum = 112
	UnicodeScriptMendeKikakui          UnicodeScriptEnum = 113
	UnicodeScriptModi                  UnicodeScriptEnum = 114
	UnicodeScriptMro                   UnicodeScriptEnum = 115
	UnicodeScriptNabataean             UnicodeScriptEnum = 116
	UnicodeScriptOldNorthArabian       UnicodeScriptEnum = 117
	UnicodeScriptOldPermic             UnicodeScriptEnum = 118
	UnicodeScriptPahawhHmong           UnicodeScriptEnum = 119
	UnicodeScriptPalmyrene             UnicodeScriptEnum = 120
	UnicodeScriptPauCinHau             UnicodeScriptEnum = 121
	UnicodeScriptPsalterPahlavi        UnicodeScriptEnum = 122
	UnicodeScriptSiddham               UnicodeScriptEnum = 123
	UnicodeScriptTirhuta               UnicodeScriptEnum = 124
	UnicodeScriptWarangCiti            UnicodeScriptEnum = 125
	UnicodeScriptAhom                  UnicodeScriptEnum = 126
	UnicodeScriptAnatolianHieroglyphs  UnicodeScriptEnum = 127
	UnicodeScriptHatran                UnicodeScriptEnum = 128
	UnicodeScriptMultani               UnicodeScriptEnum = 129
	UnicodeScriptOldHungarian          UnicodeScriptEnum = 130
	UnicodeScriptSignwriting           UnicodeScriptEnum = 131
	UnicodeScriptAdlam                 UnicodeScriptEnum = 132
	UnicodeScriptBhaiksuki             UnicodeScriptEnum = 133
	UnicodeScriptMarchen               UnicodeScriptEnum = 134
	UnicodeScriptNewa                  UnicodeScriptEnum = 135
	UnicodeScriptOsage                 UnicodeScriptEnum = 136
	UnicodeScriptTangut                UnicodeScriptEnum = 137
	UnicodeScriptMasaramGondi          UnicodeScriptEnum = 138
	UnicodeScriptNushu                 UnicodeScriptEnum = 139
	UnicodeScriptSoyombo               UnicodeScriptEnum = 140
	UnicodeScriptZanabazarSquare       UnicodeScriptEnum = 141
	UnicodeScriptDogra                 UnicodeScriptEnum = 142
	UnicodeScriptGunjalaGondi          UnicodeScriptEnum = 143
	UnicodeScriptHanifiRohingya        UnicodeScriptEnum = 144
	UnicodeScriptMakasar               UnicodeScriptEnum = 145
	UnicodeScriptMedefaidrin           UnicodeScriptEnum = 146
	UnicodeScriptOldSogdian            UnicodeScriptEnum = 147
	UnicodeScriptSogdian               UnicodeScriptEnum = 148
)

func UnicodeScriptGetType() gi.GType {
	ret := _I.GetGType(121, "UnicodeScript")
	return ret
}

// Enum UnicodeType
type UnicodeTypeEnum int

const (
	UnicodeTypeControl            UnicodeTypeEnum = 0
	UnicodeTypeFormat             UnicodeTypeEnum = 1
	UnicodeTypeUnassigned         UnicodeTypeEnum = 2
	UnicodeTypePrivateUse         UnicodeTypeEnum = 3
	UnicodeTypeSurrogate          UnicodeTypeEnum = 4
	UnicodeTypeLowercaseLetter    UnicodeTypeEnum = 5
	UnicodeTypeModifierLetter     UnicodeTypeEnum = 6
	UnicodeTypeOtherLetter        UnicodeTypeEnum = 7
	UnicodeTypeTitlecaseLetter    UnicodeTypeEnum = 8
	UnicodeTypeUppercaseLetter    UnicodeTypeEnum = 9
	UnicodeTypeSpacingMark        UnicodeTypeEnum = 10
	UnicodeTypeEnclosingMark      UnicodeTypeEnum = 11
	UnicodeTypeNonSpacingMark     UnicodeTypeEnum = 12
	UnicodeTypeDecimalNumber      UnicodeTypeEnum = 13
	UnicodeTypeLetterNumber       UnicodeTypeEnum = 14
	UnicodeTypeOtherNumber        UnicodeTypeEnum = 15
	UnicodeTypeConnectPunctuation UnicodeTypeEnum = 16
	UnicodeTypeDashPunctuation    UnicodeTypeEnum = 17
	UnicodeTypeClosePunctuation   UnicodeTypeEnum = 18
	UnicodeTypeFinalPunctuation   UnicodeTypeEnum = 19
	UnicodeTypeInitialPunctuation UnicodeTypeEnum = 20
	UnicodeTypeOtherPunctuation   UnicodeTypeEnum = 21
	UnicodeTypeOpenPunctuation    UnicodeTypeEnum = 22
	UnicodeTypeCurrencySymbol     UnicodeTypeEnum = 23
	UnicodeTypeModifierSymbol     UnicodeTypeEnum = 24
	UnicodeTypeMathSymbol         UnicodeTypeEnum = 25
	UnicodeTypeOtherSymbol        UnicodeTypeEnum = 26
	UnicodeTypeLineSeparator      UnicodeTypeEnum = 27
	UnicodeTypeParagraphSeparator UnicodeTypeEnum = 28
	UnicodeTypeSpaceSeparator     UnicodeTypeEnum = 29
)

func UnicodeTypeGetType() gi.GType {
	ret := _I.GetGType(122, "UnicodeType")
	return ret
}

type UnixFDSourceFuncStruct struct {
	F_fd        int32
	F_condition IOConditionFlags
}

func GetPointer_myUnixFDSourceFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibUnixFDSourceFunc())
}

//export myGLibUnixFDSourceFunc
func myGLibUnixFDSourceFunc(fd C.gint32, condition C.GIOCondition, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &UnixFDSourceFuncStruct{
		F_fd:        int32(fd),
		F_condition: IOConditionFlags(condition),
	}
	fn(args)
}

// Enum UserDirectory
type UserDirectoryEnum int

const (
	UserDirectoryDirectoryDesktop     UserDirectoryEnum = 0
	UserDirectoryDirectoryDocuments   UserDirectoryEnum = 1
	UserDirectoryDirectoryDownload    UserDirectoryEnum = 2
	UserDirectoryDirectoryMusic       UserDirectoryEnum = 3
	UserDirectoryDirectoryPictures    UserDirectoryEnum = 4
	UserDirectoryDirectoryPublicShare UserDirectoryEnum = 5
	UserDirectoryDirectoryTemplates   UserDirectoryEnum = 6
	UserDirectoryDirectoryVideos      UserDirectoryEnum = 7
	UserDirectoryNDirectories         UserDirectoryEnum = 8
)

func UserDirectoryGetType() gi.GType {
	ret := _I.GetGType(123, "UserDirectory")
	return ret
}

// Struct Variant
type Variant struct {
	P unsafe.Pointer
}

func VariantGetType() gi.GType {
	ret := _I.GetGType(124, "Variant")
	return ret
}

// g_variant_new_array
//
// [ child_type ] trans: nothing
//
// [ children ] trans: nothing
//
// [ n_children ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantArray(child_type VariantType, children gi.PointerArray, n_children uint64) (result Variant) {
	iv, err := _I.Get(657, "Variant", "new_array", 299, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_child_type := gi.NewPointerArgument(child_type.P)
	arg_children := gi.NewPointerArgument(children.P)
	arg_n_children := gi.NewUint64Argument(n_children)
	args := []gi.Argument{arg_child_type, arg_children, arg_n_children}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_boolean
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantBoolean(value bool) (result Variant) {
	iv, err := _I.Get(658, "Variant", "new_boolean", 299, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewBoolArgument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_byte
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantByte(value uint8) (result Variant) {
	iv, err := _I.Get(659, "Variant", "new_byte", 299, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint8Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_bytestring
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantBytestring(string gi.Uint8Array) (result Variant) {
	iv, err := _I.Get(660, "Variant", "new_bytestring", 299, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewPointerArgument(string.P)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_bytestring_array
//
// [ strv ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantBytestringArray(strv gi.CStrArray, length int64) (result Variant) {
	iv, err := _I.Get(661, "Variant", "new_bytestring_array", 299, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_strv := gi.NewPointerArgument(strv.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_strv, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_dict_entry
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantDictEntry(key Variant, value Variant) (result Variant) {
	iv, err := _I.Get(662, "Variant", "new_dict_entry", 299, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_key := gi.NewPointerArgument(key.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_double
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantDouble(value float64) (result Variant) {
	iv, err := _I.Get(663, "Variant", "new_double", 299, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewDoubleArgument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_fixed_array
//
// [ element_type ] trans: nothing
//
// [ elements ] trans: nothing
//
// [ n_elements ] trans: nothing
//
// [ element_size ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantFixedArray(element_type VariantType, elements unsafe.Pointer, n_elements uint64, element_size uint64) (result Variant) {
	iv, err := _I.Get(664, "Variant", "new_fixed_array", 299, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_element_type := gi.NewPointerArgument(element_type.P)
	arg_elements := gi.NewPointerArgument(elements)
	arg_n_elements := gi.NewUint64Argument(n_elements)
	arg_element_size := gi.NewUint64Argument(element_size)
	args := []gi.Argument{arg_element_type, arg_elements, arg_n_elements, arg_element_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_from_bytes
//
// [ type1 ] trans: nothing
//
// [ bytes ] trans: nothing
//
// [ trusted ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantFromBytes(type1 VariantType, bytes Bytes, trusted bool) (result Variant) {
	iv, err := _I.Get(665, "Variant", "new_from_bytes", 299, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_bytes := gi.NewPointerArgument(bytes.P)
	arg_trusted := gi.NewBoolArgument(trusted)
	args := []gi.Argument{arg_type1, arg_bytes, arg_trusted}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_from_data
//
// [ type1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ trusted ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantFromData(type1 VariantType, data gi.Uint8Array, size uint64, trusted bool, notify int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result Variant) {
	iv, err := _I.Get(666, "Variant", "new_from_data", 299, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	arg_trusted := gi.NewBoolArgument(trusted)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_type1, arg_data, arg_size, arg_trusted, arg_notify, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_handle
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantHandle(value int32) (result Variant) {
	iv, err := _I.Get(667, "Variant", "new_handle", 299, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_int16
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantInt16(value int16) (result Variant) {
	iv, err := _I.Get(668, "Variant", "new_int16", 299, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt16Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_int32
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantInt32(value int32) (result Variant) {
	iv, err := _I.Get(669, "Variant", "new_int32", 299, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_int64
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantInt64(value int64) (result Variant) {
	iv, err := _I.Get(670, "Variant", "new_int64", 299, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt64Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_maybe
//
// [ child_type ] trans: nothing
//
// [ child ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantMaybe(child_type VariantType, child Variant) (result Variant) {
	iv, err := _I.Get(671, "Variant", "new_maybe", 299, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_child_type := gi.NewPointerArgument(child_type.P)
	arg_child := gi.NewPointerArgument(child.P)
	args := []gi.Argument{arg_child_type, arg_child}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_object_path
//
// [ object_path ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantObjectPath(object_path string) (result Variant) {
	iv, err := _I.Get(672, "Variant", "new_object_path", 299, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_object_path := gi.CString(object_path)
	arg_object_path := gi.NewStringArgument(c_object_path)
	args := []gi.Argument{arg_object_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_object_path)
	result.P = ret.Pointer()
	return
}

// g_variant_new_objv
//
// [ strv ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantObjv(strv gi.CStrArray, length int64) (result Variant) {
	iv, err := _I.Get(673, "Variant", "new_objv", 299, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_strv := gi.NewPointerArgument(strv.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_strv, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_signature
//
// [ signature ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantSignature(signature string) (result Variant) {
	iv, err := _I.Get(674, "Variant", "new_signature", 299, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_signature := gi.CString(signature)
	arg_signature := gi.NewStringArgument(c_signature)
	args := []gi.Argument{arg_signature}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_signature)
	result.P = ret.Pointer()
	return
}

// g_variant_new_string
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantString(string string) (result Variant) {
	iv, err := _I.Get(675, "Variant", "new_string", 299, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result.P = ret.Pointer()
	return
}

// g_variant_new_strv
//
// [ strv ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantStrv(strv gi.CStrArray, length int64) (result Variant) {
	iv, err := _I.Get(676, "Variant", "new_strv", 299, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_strv := gi.NewPointerArgument(strv.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_strv, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_tuple
//
// [ children ] trans: nothing
//
// [ n_children ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantTuple(children gi.PointerArray, n_children uint64) (result Variant) {
	iv, err := _I.Get(677, "Variant", "new_tuple", 299, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_children := gi.NewPointerArgument(children.P)
	arg_n_children := gi.NewUint64Argument(n_children)
	args := []gi.Argument{arg_children, arg_n_children}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_uint16
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantUint16(value uint16) (result Variant) {
	iv, err := _I.Get(678, "Variant", "new_uint16", 299, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint16Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_uint32
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantUint32(value uint32) (result Variant) {
	iv, err := _I.Get(679, "Variant", "new_uint32", 299, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint32Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_uint64
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantUint64(value uint64) (result Variant) {
	iv, err := _I.Get(680, "Variant", "new_uint64", 299, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint64Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_variant
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func NewVariantVariant(value Variant) (result Variant) {
	iv, err := _I.Get(681, "Variant", "new_variant", 299, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_byteswap
//
// [ result ] trans: everything
//
func (v Variant) Byteswap() (result Variant) {
	iv, err := _I.Get(682, "Variant", "byteswap", 299, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_check_format_string
//
// [ format_string ] trans: nothing
//
// [ copy_only ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Variant) CheckFormatString(format_string string, copy_only bool) (result bool) {
	iv, err := _I.Get(683, "Variant", "check_format_string", 299, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_format_string := gi.CString(format_string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_format_string := gi.NewStringArgument(c_format_string)
	arg_copy_only := gi.NewBoolArgument(copy_only)
	args := []gi.Argument{arg_v, arg_format_string, arg_copy_only}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_format_string)
	result = ret.Bool()
	return
}

// g_variant_classify
//
// [ result ] trans: nothing
//
func (v Variant) Classify() (result VariantClassEnum) {
	iv, err := _I.Get(684, "Variant", "classify", 299, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = VariantClassEnum(ret.Int())
	return
}

// g_variant_compare
//
// [ two ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Variant) Compare(two Variant) (result int32) {
	iv, err := _I.Get(685, "Variant", "compare", 299, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_two := gi.NewPointerArgument(two.P)
	args := []gi.Argument{arg_v, arg_two}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_variant_dup_bytestring
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Variant) DupBytestring() (result gi.Uint8Array) {
	iv, err := _I.Get(686, "Variant", "dup_bytestring", 299, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(length)}
	return
}

// g_variant_dup_bytestring_array
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Variant) DupBytestringArray() (result gi.CStrArray) {
	iv, err := _I.Get(687, "Variant", "dup_bytestring_array", 299, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_variant_dup_objv
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Variant) DupObjv() (result gi.CStrArray) {
	iv, err := _I.Get(688, "Variant", "dup_objv", 299, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_dup_string
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Variant) DupString() (result string, length uint64) {
	iv, err := _I.Get(689, "Variant", "dup_string", 299, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	length = outArgs[0].Uint64()
	result = ret.String().Take()
	return
}

// g_variant_dup_strv
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Variant) DupStrv() (result gi.CStrArray) {
	iv, err := _I.Get(690, "Variant", "dup_strv", 299, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_equal
//
// [ two ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Variant) Equal(two Variant) (result bool) {
	iv, err := _I.Get(691, "Variant", "equal", 299, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_two := gi.NewPointerArgument(two.P)
	args := []gi.Argument{arg_v, arg_two}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_get_boolean
//
// [ result ] trans: nothing
//
func (v Variant) GetBoolean() (result bool) {
	iv, err := _I.Get(692, "Variant", "get_boolean", 299, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_get_byte
//
// [ result ] trans: nothing
//
func (v Variant) GetByte() (result uint8) {
	iv, err := _I.Get(693, "Variant", "get_byte", 299, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_variant_get_bytestring
//
// [ result ] trans: nothing
//
func (v Variant) GetBytestring() (result gi.Uint8Array) {
	iv, err := _I.Get(694, "Variant", "get_bytestring", 299, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(0)}
	return
}

// g_variant_get_bytestring_array
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v Variant) GetBytestringArray() (result gi.CStrArray) {
	iv, err := _I.Get(695, "Variant", "get_bytestring_array", 299, 38, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_variant_get_child_value
//
// [ index_ ] trans: nothing
//
// [ result ] trans: everything
//
func (v Variant) GetChildValue(index_ uint64) (result Variant) {
	iv, err := _I.Get(696, "Variant", "get_child_value", 299, 39, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint64Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_get_data
//
// [ result ] trans: nothing
//
func (v Variant) GetData() (result unsafe.Pointer) {
	iv, err := _I.Get(697, "Variant", "get_data", 299, 40, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_variant_get_data_as_bytes
//
// [ result ] trans: everything
//
func (v Variant) GetDataAsBytes() (result Bytes) {
	iv, err := _I.Get(698, "Variant", "get_data_as_bytes", 299, 41, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_get_double
//
// [ result ] trans: nothing
//
func (v Variant) GetDouble() (result float64) {
	iv, err := _I.Get(699, "Variant", "get_double", 299, 42, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_variant_get_handle
//
// [ result ] trans: nothing
//
func (v Variant) GetHandle() (result int32) {
	iv, err := _I.Get(700, "Variant", "get_handle", 299, 43, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_variant_get_int16
//
// [ result ] trans: nothing
//
func (v Variant) GetInt16() (result int16) {
	iv, err := _I.Get(701, "Variant", "get_int16", 299, 44, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int16()
	return
}

// g_variant_get_int32
//
// [ result ] trans: nothing
//
func (v Variant) GetInt32() (result int32) {
	iv, err := _I.Get(702, "Variant", "get_int32", 299, 45, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_variant_get_int64
//
// [ result ] trans: nothing
//
func (v Variant) GetInt64() (result int64) {
	iv, err := _I.Get(703, "Variant", "get_int64", 299, 46, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_variant_get_maybe
//
// [ result ] trans: everything
//
func (v Variant) GetMaybe() (result Variant) {
	iv, err := _I.Get(704, "Variant", "get_maybe", 299, 47, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_get_normal_form
//
// [ result ] trans: everything
//
func (v Variant) GetNormalForm() (result Variant) {
	iv, err := _I.Get(705, "Variant", "get_normal_form", 299, 48, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_get_objv
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v Variant) GetObjv() (result gi.CStrArray) {
	iv, err := _I.Get(706, "Variant", "get_objv", 299, 49, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_get_size
//
// [ result ] trans: nothing
//
func (v Variant) GetSize() (result uint64) {
	iv, err := _I.Get(707, "Variant", "get_size", 299, 50, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_variant_get_string
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Variant) GetString() (result string, length uint64) {
	iv, err := _I.Get(708, "Variant", "get_string", 299, 51, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	length = outArgs[0].Uint64()
	result = ret.String().Copy()
	return
}

// g_variant_get_strv
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v Variant) GetStrv() (result gi.CStrArray) {
	iv, err := _I.Get(709, "Variant", "get_strv", 299, 52, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_get_type
//
// [ result ] trans: nothing
//
func (v Variant) GetType() (result VariantType) {
	iv, err := _I.Get(710, "Variant", "get_type", 299, 53, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_get_type_string
//
// [ result ] trans: nothing
//
func (v Variant) GetTypeString() (result string) {
	iv, err := _I.Get(711, "Variant", "get_type_string", 299, 54, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_variant_get_uint16
//
// [ result ] trans: nothing
//
func (v Variant) GetUint16() (result uint16) {
	iv, err := _I.Get(712, "Variant", "get_uint16", 299, 55, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// g_variant_get_uint32
//
// [ result ] trans: nothing
//
func (v Variant) GetUint32() (result uint32) {
	iv, err := _I.Get(713, "Variant", "get_uint32", 299, 56, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_variant_get_uint64
//
// [ result ] trans: nothing
//
func (v Variant) GetUint64() (result uint64) {
	iv, err := _I.Get(714, "Variant", "get_uint64", 299, 57, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_variant_get_variant
//
// [ result ] trans: everything
//
func (v Variant) GetVariant() (result Variant) {
	iv, err := _I.Get(715, "Variant", "get_variant", 299, 58, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_hash
//
// [ result ] trans: nothing
//
func (v Variant) Hash() (result uint32) {
	iv, err := _I.Get(716, "Variant", "hash", 299, 59, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_variant_is_container
//
// [ result ] trans: nothing
//
func (v Variant) IsContainer() (result bool) {
	iv, err := _I.Get(717, "Variant", "is_container", 299, 60, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_is_floating
//
// [ result ] trans: nothing
//
func (v Variant) IsFloating() (result bool) {
	iv, err := _I.Get(718, "Variant", "is_floating", 299, 61, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_is_normal_form
//
// [ result ] trans: nothing
//
func (v Variant) IsNormalForm() (result bool) {
	iv, err := _I.Get(719, "Variant", "is_normal_form", 299, 62, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_is_of_type
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Variant) IsOfType(type1 VariantType) (result bool) {
	iv, err := _I.Get(720, "Variant", "is_of_type", 299, 63, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_lookup_value
//
// [ key ] trans: nothing
//
// [ expected_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v Variant) LookupValue(key string, expected_type VariantType) (result Variant) {
	iv, err := _I.Get(721, "Variant", "lookup_value", 299, 64, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_expected_type := gi.NewPointerArgument(expected_type.P)
	args := []gi.Argument{arg_v, arg_key, arg_expected_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result.P = ret.Pointer()
	return
}

// g_variant_n_children
//
// [ result ] trans: nothing
//
func (v Variant) NChildren() (result uint64) {
	iv, err := _I.Get(722, "Variant", "n_children", 299, 65, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_variant_print
//
// [ type_annotate ] trans: nothing
//
// [ result ] trans: everything
//
func (v Variant) Print(type_annotate bool) (result string) {
	iv, err := _I.Get(723, "Variant", "print", 299, 66, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type_annotate := gi.NewBoolArgument(type_annotate)
	args := []gi.Argument{arg_v, arg_type_annotate}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_variant_ref
//
// [ result ] trans: everything
//
func (v Variant) Ref() (result Variant) {
	iv, err := _I.Get(724, "Variant", "ref", 299, 67, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_ref_sink
//
// [ result ] trans: everything
//
func (v Variant) RefSink() (result Variant) {
	iv, err := _I.Get(725, "Variant", "ref_sink", 299, 68, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_store
//
// [ data ] trans: nothing
//
func (v Variant) Store(data unsafe.Pointer) {
	iv, err := _I.Get(726, "Variant", "store", 299, 69, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_variant_take_ref
//
// [ result ] trans: everything
//
func (v Variant) TakeRef() (result Variant) {
	iv, err := _I.Get(727, "Variant", "take_ref", 299, 70, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_unref
//
func (v Variant) Unref() {
	iv, err := _I.Get(728, "Variant", "unref", 299, 71, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_is_object_path
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantIsObjectPath1(string string) (result bool) {
	iv, err := _I.Get(729, "Variant", "is_object_path", 299, 72, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_is_signature
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantIsSignature1(string string) (result bool) {
	iv, err := _I.Get(730, "Variant", "is_signature", 299, 73, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_parse
//
// [ type1 ] trans: nothing
//
// [ text ] trans: nothing
//
// [ limit ] trans: nothing
//
// [ endptr ] trans: nothing
//
// [ result ] trans: everything
//
func VariantParse1(type1 VariantType, text string, limit string, endptr string) (result Variant, err error) {
	iv, err := _I.Get(731, "Variant", "parse", 299, 74, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	c_limit := gi.CString(limit)
	c_endptr := gi.CString(endptr)
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewStringArgument(c_endptr)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_text, arg_limit, arg_endptr, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	gi.Free(c_limit)
	gi.Free(c_endptr)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_variant_parse_error_print_context
//
// [ error ] trans: nothing
//
// [ source_str ] trans: nothing
//
// [ result ] trans: everything
//
func VariantParseErrorPrintContext1(error Error, source_str string) (result string) {
	iv, err := _I.Get(732, "Variant", "parse_error_print_context", 299, 75, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source_str := gi.CString(source_str)
	arg_error := gi.NewPointerArgument(error.P)
	arg_source_str := gi.NewStringArgument(c_source_str)
	args := []gi.Argument{arg_error, arg_source_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source_str)
	result = ret.String().Take()
	return
}

// Deprecated
//
// Struct VariantBuilder
type VariantBuilder struct {
	P unsafe.Pointer
}

func VariantBuilderGetType() gi.GType {
	ret := _I.GetGType(125, "VariantBuilder")
	return ret
}

// g_variant_builder_new
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantBuilder(type1 VariantType) (result VariantBuilder) {
	iv, err := _I.Get(735, "VariantBuilder", "new", 300, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_builder_add_value
//
// [ value ] trans: nothing
//
func (v VariantBuilder) AddValue(value Variant) {
	iv, err := _I.Get(736, "VariantBuilder", "add_value", 300, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// g_variant_builder_close
//
func (v VariantBuilder) Close() {
	iv, err := _I.Get(737, "VariantBuilder", "close", 300, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_builder_end
//
// [ result ] trans: nothing
//
func (v VariantBuilder) End() (result Variant) {
	iv, err := _I.Get(738, "VariantBuilder", "end", 300, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_builder_open
//
// [ type1 ] trans: nothing
//
func (v VariantBuilder) Open(type1 VariantType) {
	iv, err := _I.Get(739, "VariantBuilder", "open", 300, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_v, arg_type1}
	iv.Call(args, nil, nil)
}

// g_variant_builder_ref
//
// [ result ] trans: everything
//
func (v VariantBuilder) Ref() (result VariantBuilder) {
	iv, err := _I.Get(740, "VariantBuilder", "ref", 300, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_builder_unref
//
func (v VariantBuilder) Unref() {
	iv, err := _I.Get(741, "VariantBuilder", "unref", 300, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum VariantClass
type VariantClassEnum int

const (
	VariantClassBoolean    VariantClassEnum = 98
	VariantClassByte       VariantClassEnum = 121
	VariantClassInt16      VariantClassEnum = 110
	VariantClassUint16     VariantClassEnum = 113
	VariantClassInt32      VariantClassEnum = 105
	VariantClassUint32     VariantClassEnum = 117
	VariantClassInt64      VariantClassEnum = 120
	VariantClassUint64     VariantClassEnum = 116
	VariantClassHandle     VariantClassEnum = 104
	VariantClassDouble     VariantClassEnum = 100
	VariantClassString     VariantClassEnum = 115
	VariantClassObjectPath VariantClassEnum = 111
	VariantClassSignature  VariantClassEnum = 103
	VariantClassVariant    VariantClassEnum = 118
	VariantClassMaybe      VariantClassEnum = 109
	VariantClassArray      VariantClassEnum = 97
	VariantClassTuple      VariantClassEnum = 40
	VariantClassDictEntry  VariantClassEnum = 123
)

func VariantClassGetType() gi.GType {
	ret := _I.GetGType(126, "VariantClass")
	return ret
}

// Struct VariantDict
type VariantDict struct {
	P unsafe.Pointer
}

func VariantDictGetType() gi.GType {
	ret := _I.GetGType(127, "VariantDict")
	return ret
}

// g_variant_dict_new
//
// [ from_asv ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantDict(from_asv Variant) (result VariantDict) {
	iv, err := _I.Get(742, "VariantDict", "new", 302, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_from_asv := gi.NewPointerArgument(from_asv.P)
	args := []gi.Argument{arg_from_asv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_dict_clear
//
func (v VariantDict) Clear() {
	iv, err := _I.Get(743, "VariantDict", "clear", 302, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_dict_contains
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v VariantDict) Contains(key string) (result bool) {
	iv, err := _I.Get(744, "VariantDict", "contains", 302, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result = ret.Bool()
	return
}

// g_variant_dict_end
//
// [ result ] trans: nothing
//
func (v VariantDict) End() (result Variant) {
	iv, err := _I.Get(745, "VariantDict", "end", 302, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_dict_insert_value
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v VariantDict) InsertValue(key string, value Variant) {
	iv, err := _I.Get(746, "VariantDict", "insert_value", 302, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
}

// g_variant_dict_lookup_value
//
// [ key ] trans: nothing
//
// [ expected_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v VariantDict) LookupValue(key string, expected_type VariantType) (result Variant) {
	iv, err := _I.Get(747, "VariantDict", "lookup_value", 302, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_expected_type := gi.NewPointerArgument(expected_type.P)
	args := []gi.Argument{arg_v, arg_key, arg_expected_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result.P = ret.Pointer()
	return
}

// g_variant_dict_ref
//
// [ result ] trans: everything
//
func (v VariantDict) Ref() (result VariantDict) {
	iv, err := _I.Get(748, "VariantDict", "ref", 302, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_dict_remove
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v VariantDict) Remove(key string) (result bool) {
	iv, err := _I.Get(749, "VariantDict", "remove", 302, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result = ret.Bool()
	return
}

// g_variant_dict_unref
//
func (v VariantDict) Unref() {
	iv, err := _I.Get(750, "VariantDict", "unref", 302, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum VariantParseError
type VariantParseErrorEnum int

const (
	VariantParseErrorFailed                     VariantParseErrorEnum = 0
	VariantParseErrorBasicTypeExpected          VariantParseErrorEnum = 1
	VariantParseErrorCannotInferType            VariantParseErrorEnum = 2
	VariantParseErrorDefiniteTypeExpected       VariantParseErrorEnum = 3
	VariantParseErrorInputNotAtEnd              VariantParseErrorEnum = 4
	VariantParseErrorInvalidCharacter           VariantParseErrorEnum = 5
	VariantParseErrorInvalidFormatString        VariantParseErrorEnum = 6
	VariantParseErrorInvalidObjectPath          VariantParseErrorEnum = 7
	VariantParseErrorInvalidSignature           VariantParseErrorEnum = 8
	VariantParseErrorInvalidTypeString          VariantParseErrorEnum = 9
	VariantParseErrorNoCommonType               VariantParseErrorEnum = 10
	VariantParseErrorNumberOutOfRange           VariantParseErrorEnum = 11
	VariantParseErrorNumberTooBig               VariantParseErrorEnum = 12
	VariantParseErrorTypeError                  VariantParseErrorEnum = 13
	VariantParseErrorUnexpectedToken            VariantParseErrorEnum = 14
	VariantParseErrorUnknownKeyword             VariantParseErrorEnum = 15
	VariantParseErrorUnterminatedStringConstant VariantParseErrorEnum = 16
	VariantParseErrorValueExpected              VariantParseErrorEnum = 17
)

func VariantParseErrorGetType() gi.GType {
	ret := _I.GetGType(128, "VariantParseError")
	return ret
}

// Struct VariantType
type VariantType struct {
	P unsafe.Pointer
}

func VariantTypeGetType() gi.GType {
	ret := _I.GetGType(129, "VariantType")
	return ret
}

// g_variant_type_new
//
// [ type_string ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantType(type_string string) (result VariantType) {
	iv, err := _I.Get(751, "VariantType", "new", 304, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_array
//
// [ element ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantTypeArray(element VariantType) (result VariantType) {
	iv, err := _I.Get(752, "VariantType", "new_array", 304, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_element := gi.NewPointerArgument(element.P)
	args := []gi.Argument{arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_dict_entry
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantTypeDictEntry(key VariantType, value VariantType) (result VariantType) {
	iv, err := _I.Get(753, "VariantType", "new_dict_entry", 304, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_key := gi.NewPointerArgument(key.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_maybe
//
// [ element ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantTypeMaybe(element VariantType) (result VariantType) {
	iv, err := _I.Get(754, "VariantType", "new_maybe", 304, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_element := gi.NewPointerArgument(element.P)
	args := []gi.Argument{arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_tuple
//
// [ items ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func NewVariantTypeTuple(items gi.PointerArray, length int32) (result VariantType) {
	iv, err := _I.Get(755, "VariantType", "new_tuple", 304, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_items := gi.NewPointerArgument(items.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_items, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_copy
//
// [ result ] trans: everything
//
func (v VariantType) Copy() (result VariantType) {
	iv, err := _I.Get(756, "VariantType", "copy", 304, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_dup_string
//
// [ result ] trans: everything
//
func (v VariantType) DupString() (result string) {
	iv, err := _I.Get(757, "VariantType", "dup_string", 304, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_variant_type_element
//
// [ result ] trans: nothing
//
func (v VariantType) Element() (result VariantType) {
	iv, err := _I.Get(758, "VariantType", "element", 304, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_equal
//
// [ type2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v VariantType) Equal(type2 VariantType) (result bool) {
	iv, err := _I.Get(759, "VariantType", "equal", 304, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type2 := gi.NewPointerArgument(type2.P)
	args := []gi.Argument{arg_v, arg_type2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_first
//
// [ result ] trans: nothing
//
func (v VariantType) First() (result VariantType) {
	iv, err := _I.Get(760, "VariantType", "first", 304, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_free
//
func (v VariantType) Free() {
	iv, err := _I.Get(761, "VariantType", "free", 304, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_type_get_string_length
//
// [ result ] trans: nothing
//
func (v VariantType) GetStringLength() (result uint64) {
	iv, err := _I.Get(762, "VariantType", "get_string_length", 304, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_variant_type_hash
//
// [ result ] trans: nothing
//
func (v VariantType) Hash() (result uint32) {
	iv, err := _I.Get(763, "VariantType", "hash", 304, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_variant_type_is_array
//
// [ result ] trans: nothing
//
func (v VariantType) IsArray() (result bool) {
	iv, err := _I.Get(764, "VariantType", "is_array", 304, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_basic
//
// [ result ] trans: nothing
//
func (v VariantType) IsBasic() (result bool) {
	iv, err := _I.Get(765, "VariantType", "is_basic", 304, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_container
//
// [ result ] trans: nothing
//
func (v VariantType) IsContainer() (result bool) {
	iv, err := _I.Get(766, "VariantType", "is_container", 304, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_definite
//
// [ result ] trans: nothing
//
func (v VariantType) IsDefinite() (result bool) {
	iv, err := _I.Get(767, "VariantType", "is_definite", 304, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_dict_entry
//
// [ result ] trans: nothing
//
func (v VariantType) IsDictEntry() (result bool) {
	iv, err := _I.Get(768, "VariantType", "is_dict_entry", 304, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_maybe
//
// [ result ] trans: nothing
//
func (v VariantType) IsMaybe() (result bool) {
	iv, err := _I.Get(769, "VariantType", "is_maybe", 304, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_subtype_of
//
// [ supertype ] trans: nothing
//
// [ result ] trans: nothing
//
func (v VariantType) IsSubtypeOf(supertype VariantType) (result bool) {
	iv, err := _I.Get(770, "VariantType", "is_subtype_of", 304, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_supertype := gi.NewPointerArgument(supertype.P)
	args := []gi.Argument{arg_v, arg_supertype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_tuple
//
// [ result ] trans: nothing
//
func (v VariantType) IsTuple() (result bool) {
	iv, err := _I.Get(771, "VariantType", "is_tuple", 304, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_variant
//
// [ result ] trans: nothing
//
func (v VariantType) IsVariant() (result bool) {
	iv, err := _I.Get(772, "VariantType", "is_variant", 304, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_key
//
// [ result ] trans: nothing
//
func (v VariantType) Key() (result VariantType) {
	iv, err := _I.Get(773, "VariantType", "key", 304, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_n_items
//
// [ result ] trans: nothing
//
func (v VariantType) NItems() (result uint64) {
	iv, err := _I.Get(774, "VariantType", "n_items", 304, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_variant_type_next
//
// [ result ] trans: nothing
//
func (v VariantType) Next() (result VariantType) {
	iv, err := _I.Get(775, "VariantType", "next", 304, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_value
//
// [ result ] trans: nothing
//
func (v VariantType) Value() (result VariantType) {
	iv, err := _I.Get(776, "VariantType", "value", 304, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_checked_
//
// [ arg0 ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantTypeChecked_1(arg0 string) (result VariantType) {
	iv, err := _I.Get(777, "VariantType", "checked_", 304, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg0 := gi.CString(arg0)
	arg_arg0 := gi.NewStringArgument(c_arg0)
	args := []gi.Argument{arg_arg0}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_arg0)
	result.P = ret.Pointer()
	return
}

// g_variant_type_string_get_depth_
//
// [ type_string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantTypeStringGetDepth_1(type_string string) (result uint64) {
	iv, err := _I.Get(778, "VariantType", "string_get_depth_", 304, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Uint64()
	return
}

// g_variant_type_string_is_valid
//
// [ type_string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantTypeStringIsValid1(type_string string) (result bool) {
	iv, err := _I.Get(779, "VariantType", "string_is_valid", 304, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Bool()
	return
}

// g_variant_type_string_scan
//
// [ string ] trans: nothing
//
// [ limit ] trans: nothing
//
// [ endptr ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func VariantTypeStringScan1(string string, limit string) (result bool, endptr string) {
	iv, err := _I.Get(780, "VariantType", "string_scan", 304, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	c_limit := gi.CString(limit)
	arg_string := gi.NewStringArgument(c_string)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_limit, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	gi.Free(c_limit)
	endptr = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

type VoidFuncStruct struct {
}

func GetPointer_myVoidFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGLibVoidFunc())
}

//export myGLibVoidFunc
func myGLibVoidFunc() {
	// TODO: not found user_data
}

// g_access
//
// [ filename ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ result ] trans: nothing
//
func Access(filename string, mode int32) (result int32) {
	iv, err := _I.Get(781, "access", "", 307, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_mode := gi.NewInt32Argument(mode)
	args := []gi.Argument{arg_filename, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Int32()
	return
}

// g_ascii_digit_value
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiDigitValue(c int8) (result int32) {
	iv, err := _I.Get(782, "ascii_digit_value", "", 308, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_ascii_dtostr
//
// [ buffer ] trans: nothing
//
// [ buf_len ] trans: nothing
//
// [ d ] trans: nothing
//
// [ result ] trans: everything
//
func AsciiDtostr(buffer string, buf_len int32, d float64) (result string) {
	iv, err := _I.Get(783, "ascii_dtostr", "", 309, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buffer := gi.CString(buffer)
	arg_buffer := gi.NewStringArgument(c_buffer)
	arg_buf_len := gi.NewInt32Argument(buf_len)
	arg_d := gi.NewDoubleArgument(d)
	args := []gi.Argument{arg_buffer, arg_buf_len, arg_d}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buffer)
	result = ret.String().Take()
	return
}

// g_ascii_formatd
//
// [ buffer ] trans: nothing
//
// [ buf_len ] trans: nothing
//
// [ format ] trans: nothing
//
// [ d ] trans: nothing
//
// [ result ] trans: everything
//
func AsciiFormatd(buffer string, buf_len int32, format string, d float64) (result string) {
	iv, err := _I.Get(784, "ascii_formatd", "", 310, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buffer := gi.CString(buffer)
	c_format := gi.CString(format)
	arg_buffer := gi.NewStringArgument(c_buffer)
	arg_buf_len := gi.NewInt32Argument(buf_len)
	arg_format := gi.NewStringArgument(c_format)
	arg_d := gi.NewDoubleArgument(d)
	args := []gi.Argument{arg_buffer, arg_buf_len, arg_format, arg_d}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buffer)
	gi.Free(c_format)
	result = ret.String().Take()
	return
}

// g_ascii_strcasecmp
//
// [ s1 ] trans: nothing
//
// [ s2 ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiStrcasecmp(s1 string, s2 string) (result int32) {
	iv, err := _I.Get(785, "ascii_strcasecmp", "", 311, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	args := []gi.Argument{arg_s1, arg_s2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_ascii_strdown
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func AsciiStrdown(str string, len1 int64) (result string) {
	iv, err := _I.Get(786, "ascii_strdown", "", 312, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ascii_string_to_signed
//
// [ str ] trans: nothing
//
// [ base ] trans: nothing
//
// [ min ] trans: nothing
//
// [ max ] trans: nothing
//
// [ out_num ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func AsciiStringToSigned(str string, base uint32, min int64, max int64) (result bool, out_num int64, err error) {
	iv, err := _I.Get(787, "ascii_string_to_signed", "", 313, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_base := gi.NewUint32Argument(base)
	arg_min := gi.NewInt64Argument(min)
	arg_max := gi.NewInt64Argument(max)
	arg_out_num := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_str, arg_base, arg_min, arg_max, arg_out_num, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[1].Pointer())
	out_num = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// g_ascii_string_to_unsigned
//
// [ str ] trans: nothing
//
// [ base ] trans: nothing
//
// [ min ] trans: nothing
//
// [ max ] trans: nothing
//
// [ out_num ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func AsciiStringToUnsigned(str string, base uint32, min uint64, max uint64) (result bool, out_num uint64, err error) {
	iv, err := _I.Get(788, "ascii_string_to_unsigned", "", 314, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_base := gi.NewUint32Argument(base)
	arg_min := gi.NewUint64Argument(min)
	arg_max := gi.NewUint64Argument(max)
	arg_out_num := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_str, arg_base, arg_min, arg_max, arg_out_num, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[1].Pointer())
	out_num = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// g_ascii_strncasecmp
//
// [ s1 ] trans: nothing
//
// [ s2 ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiStrncasecmp(s1 string, s2 string, n uint64) (result int32) {
	iv, err := _I.Get(789, "ascii_strncasecmp", "", 315, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_s1, arg_s2, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_ascii_strtod
//
// [ nptr ] trans: nothing
//
// [ endptr ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func AsciiStrtod(nptr string) (result float64, endptr string) {
	iv, err := _I.Get(790, "ascii_strtod", "", 316, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_nptr, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Copy()
	result = ret.Double()
	return
}

// g_ascii_strtoll
//
// [ nptr ] trans: nothing
//
// [ endptr ] trans: nothing, dir: out
//
// [ base ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiStrtoll(nptr string, base uint32) (result int64, endptr string) {
	iv, err := _I.Get(791, "ascii_strtoll", "", 317, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_base := gi.NewUint32Argument(base)
	args := []gi.Argument{arg_nptr, arg_endptr, arg_base}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Copy()
	result = ret.Int64()
	return
}

// g_ascii_strtoull
//
// [ nptr ] trans: nothing
//
// [ endptr ] trans: nothing, dir: out
//
// [ base ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiStrtoull(nptr string, base uint32) (result uint64, endptr string) {
	iv, err := _I.Get(792, "ascii_strtoull", "", 318, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_base := gi.NewUint32Argument(base)
	args := []gi.Argument{arg_nptr, arg_endptr, arg_base}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Copy()
	result = ret.Uint64()
	return
}

// g_ascii_strup
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func AsciiStrup(str string, len1 int64) (result string) {
	iv, err := _I.Get(793, "ascii_strup", "", 319, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ascii_tolower
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiTolower(c int8) (result int8) {
	iv, err := _I.Get(794, "ascii_tolower", "", 320, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int8()
	return
}

// g_ascii_toupper
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiToupper(c int8) (result int8) {
	iv, err := _I.Get(795, "ascii_toupper", "", 321, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int8()
	return
}

// g_ascii_xdigit_value
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func AsciiXdigitValue(c int8) (result int32) {
	iv, err := _I.Get(796, "ascii_xdigit_value", "", 322, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_assert_warning
//
// [ log_domain ] trans: nothing
//
// [ file ] trans: nothing
//
// [ line ] trans: nothing
//
// [ pretty_function ] trans: nothing
//
// [ expression ] trans: nothing
//
func AssertWarning(log_domain string, file string, line int32, pretty_function string, expression string) {
	iv, err := _I.Get(797, "assert_warning", "", 323, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	c_file := gi.CString(file)
	c_pretty_function := gi.CString(pretty_function)
	c_expression := gi.CString(expression)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_pretty_function := gi.NewStringArgument(c_pretty_function)
	arg_expression := gi.NewStringArgument(c_expression)
	args := []gi.Argument{arg_log_domain, arg_file, arg_line, arg_pretty_function, arg_expression}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
	gi.Free(c_file)
	gi.Free(c_pretty_function)
	gi.Free(c_expression)
}

// g_assertion_message
//
// [ domain ] trans: nothing
//
// [ file ] trans: nothing
//
// [ line ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ message ] trans: nothing
//
func AssertionMessage(domain string, file string, line int32, func1 string, message string) {
	iv, err := _I.Get(798, "assertion_message", "", 324, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_message := gi.CString(message)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_message := gi.NewStringArgument(c_message)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_message}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_message)
}

// g_assertion_message_cmpstr
//
// [ domain ] trans: nothing
//
// [ file ] trans: nothing
//
// [ line ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ expr ] trans: nothing
//
// [ arg1 ] trans: nothing
//
// [ cmp ] trans: nothing
//
// [ arg2 ] trans: nothing
//
func AssertionMessageCmpstr(domain string, file string, line int32, func1 string, expr string, arg1 string, cmp string, arg2 string) {
	iv, err := _I.Get(799, "assertion_message_cmpstr", "", 325, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_expr := gi.CString(expr)
	c_arg1 := gi.CString(arg1)
	c_cmp := gi.CString(cmp)
	c_arg2 := gi.CString(arg2)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_expr := gi.NewStringArgument(c_expr)
	arg_arg1 := gi.NewStringArgument(c_arg1)
	arg_cmp := gi.NewStringArgument(c_cmp)
	arg_arg2 := gi.NewStringArgument(c_arg2)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_expr, arg_arg1, arg_cmp, arg_arg2}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_expr)
	gi.Free(c_arg1)
	gi.Free(c_cmp)
	gi.Free(c_arg2)
}

// g_assertion_message_error
//
// [ domain ] trans: nothing
//
// [ file ] trans: nothing
//
// [ line ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ expr ] trans: nothing
//
// [ error ] trans: nothing
//
// [ error_domain ] trans: nothing
//
// [ error_code ] trans: nothing
//
func AssertionMessageError(domain string, file string, line int32, func1 string, expr string, error Error, error_domain uint32, error_code int32) {
	iv, err := _I.Get(800, "assertion_message_error", "", 326, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_expr := gi.CString(expr)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_expr := gi.NewStringArgument(c_expr)
	arg_error := gi.NewPointerArgument(error.P)
	arg_error_domain := gi.NewUint32Argument(error_domain)
	arg_error_code := gi.NewInt32Argument(error_code)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_expr, arg_error, arg_error_domain, arg_error_code}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_expr)
}

// Deprecated
//
// g_atexit
//
// [ func1 ] trans: nothing
//
func Atexit(func1 int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(801, "atexit", "", 327, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myVoidFunc()))
	args := []gi.Argument{arg_func1}
	iv.Call(args, nil, nil)
}

// g_atomic_int_add
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntAdd(atomic int32, val int32) (result int32) {
	iv, err := _I.Get(802, "atomic_int_add", "", 328, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_atomic_int_and
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntAnd(atomic uint32, val uint32) (result uint32) {
	iv, err := _I.Get(803, "atomic_int_and", "", 329, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewUint32Argument(atomic)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_atomic_int_compare_and_exchange
//
// [ atomic ] trans: nothing
//
// [ oldval ] trans: nothing
//
// [ newval ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntCompareAndExchange(atomic int32, oldval int32, newval int32) (result bool) {
	iv, err := _I.Get(804, "atomic_int_compare_and_exchange", "", 330, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_oldval := gi.NewInt32Argument(oldval)
	arg_newval := gi.NewInt32Argument(newval)
	args := []gi.Argument{arg_atomic, arg_oldval, arg_newval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_int_dec_and_test
//
// [ atomic ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntDecAndTest(atomic int32) (result bool) {
	iv, err := _I.Get(805, "atomic_int_dec_and_test", "", 331, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	args := []gi.Argument{arg_atomic}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// g_atomic_int_exchange_and_add
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntExchangeAndAdd(atomic int32, val int32) (result int32) {
	iv, err := _I.Get(806, "atomic_int_exchange_and_add", "", 332, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_atomic_int_get
//
// [ atomic ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntGet(atomic int32) (result int32) {
	iv, err := _I.Get(807, "atomic_int_get", "", 333, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	args := []gi.Argument{arg_atomic}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_atomic_int_inc
//
// [ atomic ] trans: nothing
//
func AtomicIntInc(atomic int32) {
	iv, err := _I.Get(808, "atomic_int_inc", "", 334, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	args := []gi.Argument{arg_atomic}
	iv.Call(args, nil, nil)
}

// g_atomic_int_or
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntOr(atomic uint32, val uint32) (result uint32) {
	iv, err := _I.Get(809, "atomic_int_or", "", 335, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewUint32Argument(atomic)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_atomic_int_set
//
// [ atomic ] trans: nothing
//
// [ newval ] trans: nothing
//
func AtomicIntSet(atomic int32, newval int32) {
	iv, err := _I.Get(810, "atomic_int_set", "", 336, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_newval := gi.NewInt32Argument(newval)
	args := []gi.Argument{arg_atomic, arg_newval}
	iv.Call(args, nil, nil)
}

// g_atomic_int_xor
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicIntXor(atomic uint32, val uint32) (result uint32) {
	iv, err := _I.Get(811, "atomic_int_xor", "", 337, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewUint32Argument(atomic)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_atomic_pointer_add
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicPointerAdd(atomic unsafe.Pointer, val int64) (result int64) {
	iv, err := _I.Get(812, "atomic_pointer_add", "", 338, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewInt64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_atomic_pointer_and
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicPointerAnd(atomic unsafe.Pointer, val uint64) (result uint64) {
	iv, err := _I.Get(813, "atomic_pointer_and", "", 339, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_pointer_compare_and_exchange
//
// [ atomic ] trans: nothing
//
// [ oldval ] trans: nothing
//
// [ newval ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicPointerCompareAndExchange(atomic unsafe.Pointer, oldval unsafe.Pointer, newval unsafe.Pointer) (result bool) {
	iv, err := _I.Get(814, "atomic_pointer_compare_and_exchange", "", 340, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_oldval := gi.NewPointerArgument(oldval)
	arg_newval := gi.NewPointerArgument(newval)
	args := []gi.Argument{arg_atomic, arg_oldval, arg_newval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_pointer_get
//
// [ atomic ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicPointerGet(atomic unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(815, "atomic_pointer_get", "", 341, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	args := []gi.Argument{arg_atomic}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_pointer_or
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicPointerOr(atomic unsafe.Pointer, val uint64) (result uint64) {
	iv, err := _I.Get(816, "atomic_pointer_or", "", 342, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_pointer_set
//
// [ atomic ] trans: nothing
//
// [ newval ] trans: nothing
//
func AtomicPointerSet(atomic unsafe.Pointer, newval unsafe.Pointer) {
	iv, err := _I.Get(817, "atomic_pointer_set", "", 343, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_newval := gi.NewPointerArgument(newval)
	args := []gi.Argument{arg_atomic, arg_newval}
	iv.Call(args, nil, nil)
}

// g_atomic_pointer_xor
//
// [ atomic ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicPointerXor(atomic unsafe.Pointer, val uint64) (result uint64) {
	iv, err := _I.Get(818, "atomic_pointer_xor", "", 344, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_rc_box_acquire
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: everything
//
func AtomicRcBoxAcquire(mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(819, "atomic_rc_box_acquire", "", 345, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_alloc
//
// [ block_size ] trans: nothing
//
// [ result ] trans: everything
//
func AtomicRcBoxAlloc(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(820, "atomic_rc_box_alloc", "", 346, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_alloc0
//
// [ block_size ] trans: nothing
//
// [ result ] trans: everything
//
func AtomicRcBoxAlloc0(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(821, "atomic_rc_box_alloc0", "", 347, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_dup
//
// [ block_size ] trans: nothing
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: everything
//
func AtomicRcBoxDup(block_size uint64, mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(822, "atomic_rc_box_dup", "", 348, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_get_size
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicRcBoxGetSize(mem_block unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(823, "atomic_rc_box_get_size", "", 349, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_rc_box_release
//
// [ mem_block ] trans: everything
//
func AtomicRcBoxRelease(mem_block unsafe.Pointer) {
	iv, err := _I.Get(824, "atomic_rc_box_release", "", 350, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	iv.Call(args, nil, nil)
}

// g_atomic_rc_box_release_full
//
// [ mem_block ] trans: everything
//
// [ clear_func ] trans: nothing
//
func AtomicRcBoxReleaseFull(mem_block unsafe.Pointer, clear_func int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(825, "atomic_rc_box_release_full", "", 351, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	arg_clear_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_mem_block, arg_clear_func}
	iv.Call(args, nil, nil)
}

// g_atomic_ref_count_compare
//
// [ arc ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicRefCountCompare(arc int32, val int32) (result bool) {
	iv, err := _I.Get(826, "atomic_ref_count_compare", "", 352, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_arc, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_ref_count_dec
//
// [ arc ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomicRefCountDec(arc int32) (result bool) {
	iv, err := _I.Get(827, "atomic_ref_count_dec", "", 353, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	args := []gi.Argument{arg_arc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_ref_count_inc
//
// [ arc ] trans: nothing
//
func AtomicRefCountInc(arc int32) {
	iv, err := _I.Get(828, "atomic_ref_count_inc", "", 354, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	args := []gi.Argument{arg_arc}
	iv.Call(args, nil, nil)
}

// g_atomic_ref_count_init
//
// [ arc ] trans: nothing
//
func AtomicRefCountInit(arc int32) {
	iv, err := _I.Get(829, "atomic_ref_count_init", "", 355, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	args := []gi.Argument{arg_arc}
	iv.Call(args, nil, nil)
}

// g_base64_decode
//
// [ text ] trans: nothing
//
// [ out_len ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func Base64Decode(text string) (result gi.Uint8Array) {
	iv, err := _I.Get(830, "base64_decode", "", 356, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_out_len := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_text, arg_out_len}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	var out_len uint64
	_ = out_len
	out_len = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(out_len)}
	return
}

// g_base64_decode_inplace
//
// [ text ] trans: everything, dir: inout
//
// [ out_len ] trans: nothing, dir: inout
//
// [ result ] trans: nothing
//
func Base64DecodeInplace(text int /*TODO:TYPE*/, out_len int /*TODO:TYPE*/) (result uint8) {
	iv, err := _I.Get(831, "base64_decode_inplace", "", 357, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	var ret gi.Argument
	iv.Call(nil, &ret, &outArgs[0])
	result = ret.Uint8()
	return
}

// g_base64_encode
//
// [ data ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Base64Encode(data gi.Uint8Array, len1 uint64) (result string) {
	iv, err := _I.Get(832, "base64_encode", "", 358, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_data, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_base64_encode_close
//
// [ break_lines ] trans: nothing
//
// [ out ] trans: everything, dir: out
//
// [ state ] trans: everything, dir: inout
//
// [ save ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func Base64EncodeClose(break_lines bool, state int /*TODO:TYPE*/, save int /*TODO:TYPE*/) (result uint64, out gi.Uint8Array) {
	iv, err := _I.Get(833, "base64_encode_close", "", 359, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_break_lines := gi.NewBoolArgument(break_lines)
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_break_lines, arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	out.P = outArgs[0].Pointer()
	result = ret.Uint64()
	return
}

// g_base64_encode_step
//
// [ in ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ break_lines ] trans: nothing
//
// [ out ] trans: everything, dir: out
//
// [ state ] trans: everything, dir: inout
//
// [ save ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func Base64EncodeStep(in gi.Uint8Array, len1 uint64, break_lines bool, state int /*TODO:TYPE*/, save int /*TODO:TYPE*/) (result uint64, out gi.Uint8Array) {
	iv, err := _I.Get(834, "base64_encode_step", "", 360, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_in := gi.NewPointerArgument(in.P)
	arg_len1 := gi.NewUint64Argument(len1)
	arg_break_lines := gi.NewBoolArgument(break_lines)
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_in, arg_len1, arg_break_lines, arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	out.P = outArgs[0].Pointer()
	result = ret.Uint64()
	return
}

// Deprecated
//
// g_basename
//
// [ file_name ] trans: nothing
//
// [ result ] trans: nothing
//
func Basename(file_name string) (result string) {
	iv, err := _I.Get(835, "basename", "", 361, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Copy()
	return
}

// g_bit_lock
//
// [ address ] trans: nothing
//
// [ lock_bit ] trans: nothing
//
func BitLock(address int32, lock_bit int32) {
	iv, err := _I.Get(836, "bit_lock", "", 362, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewInt32Argument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_bit_nth_lsf
//
// [ mask ] trans: nothing
//
// [ nth_bit ] trans: nothing
//
// [ result ] trans: nothing
//
func BitNthLsf(mask uint64, nth_bit int32) (result int32) {
	iv, err := _I.Get(837, "bit_nth_lsf", "", 363, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mask := gi.NewUint64Argument(mask)
	arg_nth_bit := gi.NewInt32Argument(nth_bit)
	args := []gi.Argument{arg_mask, arg_nth_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bit_nth_msf
//
// [ mask ] trans: nothing
//
// [ nth_bit ] trans: nothing
//
// [ result ] trans: nothing
//
func BitNthMsf(mask uint64, nth_bit int32) (result int32) {
	iv, err := _I.Get(838, "bit_nth_msf", "", 364, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mask := gi.NewUint64Argument(mask)
	arg_nth_bit := gi.NewInt32Argument(nth_bit)
	args := []gi.Argument{arg_mask, arg_nth_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bit_storage
//
// [ number ] trans: nothing
//
// [ result ] trans: nothing
//
func BitStorage(number uint64) (result uint32) {
	iv, err := _I.Get(839, "bit_storage", "", 365, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_number := gi.NewUint64Argument(number)
	args := []gi.Argument{arg_number}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_bit_trylock
//
// [ address ] trans: nothing
//
// [ lock_bit ] trans: nothing
//
// [ result ] trans: nothing
//
func BitTrylock(address int32, lock_bit int32) (result bool) {
	iv, err := _I.Get(840, "bit_trylock", "", 366, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewInt32Argument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_bit_unlock
//
// [ address ] trans: nothing
//
// [ lock_bit ] trans: nothing
//
func BitUnlock(address int32, lock_bit int32) {
	iv, err := _I.Get(841, "bit_unlock", "", 367, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewInt32Argument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_bookmark_file_error_quark
//
// [ result ] trans: nothing
//
func BookmarkFileErrorQuark() (result uint32) {
	iv, err := _I.Get(842, "bookmark_file_error_quark", "", 368, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_build_filenamev
//
// [ args ] trans: nothing
//
// [ result ] trans: everything
//
func BuildFilenamev(args gi.CStrArray) (result string) {
	iv, err := _I.Get(843, "build_filenamev", "", 369, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_args := gi.NewPointerArgument(args.P)
	args1 := []gi.Argument{arg_args}
	var ret gi.Argument
	iv.Call(args1, &ret, nil)
	result = ret.String().Take()
	return
}

// g_build_pathv
//
// [ separator ] trans: nothing
//
// [ args ] trans: nothing
//
// [ result ] trans: everything
//
func BuildPathv(separator string, args gi.CStrArray) (result string) {
	iv, err := _I.Get(844, "build_pathv", "", 370, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_separator := gi.CString(separator)
	arg_separator := gi.NewStringArgument(c_separator)
	arg_args := gi.NewPointerArgument(args.P)
	args1 := []gi.Argument{arg_separator, arg_args}
	var ret gi.Argument
	iv.Call(args1, &ret, nil)
	gi.Free(c_separator)
	result = ret.String().Take()
	return
}

// g_byte_array_free
//
// [ array ] trans: nothing
//
// [ free_segment ] trans: nothing
//
// [ result ] trans: nothing
//
func ByteArrayFree(array ByteArray, free_segment bool) (result uint8) {
	iv, err := _I.Get(845, "byte_array_free", "", 371, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	arg_free_segment := gi.NewBoolArgument(free_segment)
	args := []gi.Argument{arg_array, arg_free_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_byte_array_free_to_bytes
//
// [ array ] trans: everything
//
// [ result ] trans: everything
//
func ByteArrayFreeToBytes(array ByteArray) (result Bytes) {
	iv, err := _I.Get(846, "byte_array_free_to_bytes", "", 372, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_new
//
// [ result ] trans: everything
//
func ByteArrayNew() (result ByteArray) {
	iv, err := _I.Get(847, "byte_array_new", "", 373, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_new_take
//
// [ data ] trans: everything
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func ByteArrayNewTake(data gi.Uint8Array, len1 uint64) (result ByteArray) {
	iv, err := _I.Get(848, "byte_array_new_take", "", 374, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_data, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_unref
//
// [ array ] trans: nothing
//
func ByteArrayUnref(array ByteArray) {
	iv, err := _I.Get(849, "byte_array_unref", "", 375, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	iv.Call(args, nil, nil)
}

// g_canonicalize_filename
//
// [ filename ] trans: nothing
//
// [ relative_to ] trans: nothing
//
// [ result ] trans: everything
//
func CanonicalizeFilename(filename string, relative_to string) (result string) {
	iv, err := _I.Get(850, "canonicalize_filename", "", 376, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	c_relative_to := gi.CString(relative_to)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_relative_to := gi.NewStringArgument(c_relative_to)
	args := []gi.Argument{arg_filename, arg_relative_to}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	gi.Free(c_relative_to)
	result = ret.String().Take()
	return
}

// g_chdir
//
// [ path ] trans: nothing
//
// [ result ] trans: nothing
//
func Chdir(path string) (result int32) {
	iv, err := _I.Get(851, "chdir", "", 377, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_path)
	result = ret.Int32()
	return
}

// glib_check_version
//
// [ required_major ] trans: nothing
//
// [ required_minor ] trans: nothing
//
// [ required_micro ] trans: nothing
//
// [ result ] trans: nothing
//
func CheckVersion(required_major uint32, required_minor uint32, required_micro uint32) (result string) {
	iv, err := _I.Get(852, "check_version", "", 378, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_required_major := gi.NewUint32Argument(required_major)
	arg_required_minor := gi.NewUint32Argument(required_minor)
	arg_required_micro := gi.NewUint32Argument(required_micro)
	args := []gi.Argument{arg_required_major, arg_required_minor, arg_required_micro}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_checksum_type_get_length
//
// [ checksum_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ChecksumTypeGetLength(checksum_type ChecksumTypeEnum) (result int64) {
	iv, err := _I.Get(853, "checksum_type_get_length", "", 379, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	args := []gi.Argument{arg_checksum_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_child_watch_add_full
//
// [ priority ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func ChildWatchAdd(priority int32, pid int32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(854, "child_watch_add", "", 380, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_pid := gi.NewInt32Argument(pid)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myChildWatchFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_pid, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_child_watch_source_new
//
// [ pid ] trans: nothing
//
// [ result ] trans: everything
//
func ChildWatchSourceNew(pid int32) (result Source) {
	iv, err := _I.Get(855, "child_watch_source_new", "", 381, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_clear_error
//
func ClearError() (err error) {
	iv, err := _I.Get(856, "clear_error", "", 382, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_err}
	iv.Call(args, nil, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	return
}

// g_close
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func Close(fd int32) (result bool, err error) {
	iv, err := _I.Get(857, "close", "", 383, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fd := gi.NewInt32Argument(fd)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_compute_checksum_for_bytes
//
// [ checksum_type ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: everything
//
func ComputeChecksumForBytes(checksum_type ChecksumTypeEnum, data Bytes) (result string) {
	iv, err := _I.Get(858, "compute_checksum_for_bytes", "", 384, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_checksum_type, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_checksum_for_data
//
// [ checksum_type ] trans: nothing
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func ComputeChecksumForData(checksum_type ChecksumTypeEnum, data gi.Uint8Array, length uint64) (result string) {
	iv, err := _I.Get(859, "compute_checksum_for_data", "", 385, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_checksum_type, arg_data, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_checksum_for_string
//
// [ checksum_type ] trans: nothing
//
// [ str ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func ComputeChecksumForString(checksum_type ChecksumTypeEnum, str string, length int64) (result string) {
	iv, err := _I.Get(860, "compute_checksum_for_string", "", 386, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	arg_str := gi.NewStringArgument(c_str)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_checksum_type, arg_str, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_compute_hmac_for_bytes
//
// [ digest_type ] trans: nothing
//
// [ key ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: everything
//
func ComputeHmacForBytes(digest_type ChecksumTypeEnum, key Bytes, data Bytes) (result string) {
	iv, err := _I.Get(861, "compute_hmac_for_bytes", "", 387, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_digest_type := gi.NewIntArgument(int(digest_type))
	arg_key := gi.NewPointerArgument(key.P)
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_digest_type, arg_key, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_hmac_for_data
//
// [ digest_type ] trans: nothing
//
// [ key ] trans: nothing
//
// [ key_len ] trans: nothing
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func ComputeHmacForData(digest_type ChecksumTypeEnum, key gi.Uint8Array, key_len uint64, data gi.Uint8Array, length uint64) (result string) {
	iv, err := _I.Get(862, "compute_hmac_for_data", "", 388, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_digest_type := gi.NewIntArgument(int(digest_type))
	arg_key := gi.NewPointerArgument(key.P)
	arg_key_len := gi.NewUint64Argument(key_len)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_digest_type, arg_key, arg_key_len, arg_data, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_hmac_for_string
//
// [ digest_type ] trans: nothing
//
// [ key ] trans: nothing
//
// [ key_len ] trans: nothing
//
// [ str ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func ComputeHmacForString(digest_type ChecksumTypeEnum, key gi.Uint8Array, key_len uint64, str string, length int64) (result string) {
	iv, err := _I.Get(863, "compute_hmac_for_string", "", 389, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_digest_type := gi.NewIntArgument(int(digest_type))
	arg_key := gi.NewPointerArgument(key.P)
	arg_key_len := gi.NewUint64Argument(key_len)
	arg_str := gi.NewStringArgument(c_str)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_digest_type, arg_key, arg_key_len, arg_str, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_convert
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ to_codeset ] trans: nothing
//
// [ from_codeset ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func Convert(str gi.Uint8Array, len1 int64, to_codeset string, from_codeset string) (result gi.Uint8Array, bytes_read uint64, err error) {
	iv, err := _I.Get(864, "convert", "", 390, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_to_codeset := gi.CString(to_codeset)
	c_from_codeset := gi.CString(from_codeset)
	arg_str := gi.NewPointerArgument(str.P)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_to_codeset := gi.NewStringArgument(c_to_codeset)
	arg_from_codeset := gi.NewStringArgument(c_from_codeset)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_to_codeset, arg_from_codeset, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_to_codeset)
	gi.Free(c_from_codeset)
	var bytes_written uint64
	_ = bytes_written
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(bytes_written)}
	return
}

// g_convert_error_quark
//
// [ result ] trans: nothing
//
func ConvertErrorQuark() (result uint32) {
	iv, err := _I.Get(865, "convert_error_quark", "", 391, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_convert_with_fallback
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ to_codeset ] trans: nothing
//
// [ from_codeset ] trans: nothing
//
// [ fallback ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func ConvertWithFallback(str gi.Uint8Array, len1 int64, to_codeset string, from_codeset string, fallback string) (result gi.Uint8Array, bytes_read uint64, err error) {
	iv, err := _I.Get(866, "convert_with_fallback", "", 392, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_to_codeset := gi.CString(to_codeset)
	c_from_codeset := gi.CString(from_codeset)
	c_fallback := gi.CString(fallback)
	arg_str := gi.NewPointerArgument(str.P)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_to_codeset := gi.NewStringArgument(c_to_codeset)
	arg_from_codeset := gi.NewStringArgument(c_from_codeset)
	arg_fallback := gi.NewStringArgument(c_fallback)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_to_codeset, arg_from_codeset, arg_fallback, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_to_codeset)
	gi.Free(c_from_codeset)
	gi.Free(c_fallback)
	var bytes_written uint64
	_ = bytes_written
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(bytes_written)}
	return
}

// g_datalist_foreach
//
// [ datalist ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DatalistForeach(datalist Data, func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(867, "datalist_foreach", "", 393, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDataForeachFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_datalist, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// g_datalist_get_data
//
// [ datalist ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func DatalistGetData(datalist Data, key string) (result unsafe.Pointer) {
	iv, err := _I.Get(868, "datalist_get_data", "", 394, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_datalist, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result = ret.Pointer()
	return
}

// g_datalist_get_flags
//
// [ datalist ] trans: nothing
//
// [ result ] trans: nothing
//
func DatalistGetFlags(datalist Data) (result uint32) {
	iv, err := _I.Get(869, "datalist_get_flags", "", 395, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	args := []gi.Argument{arg_datalist}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_datalist_id_get_data
//
// [ datalist ] trans: nothing
//
// [ key_id ] trans: nothing
//
// [ result ] trans: nothing
//
func DatalistIdGetData(datalist Data, key_id uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(870, "datalist_id_get_data", "", 396, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_key_id := gi.NewUint32Argument(key_id)
	args := []gi.Argument{arg_datalist, arg_key_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_datalist_set_flags
//
// [ datalist ] trans: nothing
//
// [ flags ] trans: nothing
//
func DatalistSetFlags(datalist Data, flags uint32) {
	iv, err := _I.Get(871, "datalist_set_flags", "", 397, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_datalist, arg_flags}
	iv.Call(args, nil, nil)
}

// g_datalist_unset_flags
//
// [ datalist ] trans: nothing
//
// [ flags ] trans: nothing
//
func DatalistUnsetFlags(datalist Data, flags uint32) {
	iv, err := _I.Get(872, "datalist_unset_flags", "", 398, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_datalist, arg_flags}
	iv.Call(args, nil, nil)
}

// g_dataset_destroy
//
// [ dataset_location ] trans: nothing
//
func DatasetDestroy(dataset_location unsafe.Pointer) {
	iv, err := _I.Get(873, "dataset_destroy", "", 399, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dataset_location := gi.NewPointerArgument(dataset_location)
	args := []gi.Argument{arg_dataset_location}
	iv.Call(args, nil, nil)
}

// g_dataset_foreach
//
// [ dataset_location ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DatasetForeach(dataset_location unsafe.Pointer, func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(874, "dataset_foreach", "", 400, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dataset_location := gi.NewPointerArgument(dataset_location)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDataForeachFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_dataset_location, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// g_dataset_id_get_data
//
// [ dataset_location ] trans: nothing
//
// [ key_id ] trans: nothing
//
// [ result ] trans: nothing
//
func DatasetIdGetData(dataset_location unsafe.Pointer, key_id uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(875, "dataset_id_get_data", "", 401, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dataset_location := gi.NewPointerArgument(dataset_location)
	arg_key_id := gi.NewUint32Argument(key_id)
	args := []gi.Argument{arg_dataset_location, arg_key_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_date_get_days_in_month
//
// [ month ] trans: nothing
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateGetDaysInMonth(month DateMonthEnum, year uint16) (result uint8) {
	iv, err := _I.Get(876, "date_get_days_in_month", "", 402, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_monday_weeks_in_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateGetMondayWeeksInYear(year uint16) (result uint8) {
	iv, err := _I.Get(877, "date_get_monday_weeks_in_year", "", 403, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_sunday_weeks_in_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateGetSundayWeeksInYear(year uint16) (result uint8) {
	iv, err := _I.Get(878, "date_get_sunday_weeks_in_year", "", 404, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_is_leap_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateIsLeapYear(year uint16) (result bool) {
	iv, err := _I.Get(879, "date_is_leap_year", "", 405, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_strftime
//
// [ s ] trans: nothing
//
// [ slen ] trans: nothing
//
// [ format ] trans: nothing
//
// [ date ] trans: nothing
//
// [ result ] trans: nothing
//
func DateStrftime(s string, slen uint64, format string, date Date) (result uint64) {
	iv, err := _I.Get(880, "date_strftime", "", 406, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s := gi.CString(s)
	c_format := gi.CString(format)
	arg_s := gi.NewStringArgument(c_s)
	arg_slen := gi.NewUint64Argument(slen)
	arg_format := gi.NewStringArgument(c_format)
	arg_date := gi.NewPointerArgument(date.P)
	args := []gi.Argument{arg_s, arg_slen, arg_format, arg_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s)
	gi.Free(c_format)
	result = ret.Uint64()
	return
}

// g_date_time_compare
//
// [ dt1 ] trans: nothing
//
// [ dt2 ] trans: nothing
//
// [ result ] trans: nothing
//
func DateTimeCompare(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result int32) {
	iv, err := _I.Get(881, "date_time_compare", "", 407, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_equal
//
// [ dt1 ] trans: nothing
//
// [ dt2 ] trans: nothing
//
// [ result ] trans: nothing
//
func DateTimeEqual(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(882, "date_time_equal", "", 408, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_hash
//
// [ datetime ] trans: nothing
//
// [ result ] trans: nothing
//
func DateTimeHash(datetime unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(883, "date_time_hash", "", 409, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datetime := gi.NewPointerArgument(datetime)
	args := []gi.Argument{arg_datetime}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_valid_day
//
// [ day ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidDay(day uint8) (result bool) {
	iv, err := _I.Get(884, "date_valid_day", "", 410, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	args := []gi.Argument{arg_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_dmy
//
// [ day ] trans: nothing
//
// [ month ] trans: nothing
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidDmy(day uint8, month DateMonthEnum, year uint16) (result bool) {
	iv, err := _I.Get(885, "date_valid_dmy", "", 411, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_day, arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_julian
//
// [ julian_date ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidJulian(julian_date uint32) (result bool) {
	iv, err := _I.Get(886, "date_valid_julian", "", 412, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_julian_date := gi.NewUint32Argument(julian_date)
	args := []gi.Argument{arg_julian_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_month
//
// [ month ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidMonth(month DateMonthEnum) (result bool) {
	iv, err := _I.Get(887, "date_valid_month", "", 413, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	args := []gi.Argument{arg_month}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_weekday
//
// [ weekday ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidWeekday(weekday DateWeekdayEnum) (result bool) {
	iv, err := _I.Get(888, "date_valid_weekday", "", 414, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_weekday := gi.NewIntArgument(int(weekday))
	args := []gi.Argument{arg_weekday}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_year
//
// [ year ] trans: nothing
//
// [ result ] trans: nothing
//
func DateValidYear(year uint16) (result bool) {
	iv, err := _I.Get(889, "date_valid_year", "", 415, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_dcgettext
//
// [ domain ] trans: nothing
//
// [ msgid ] trans: nothing
//
// [ category ] trans: nothing
//
// [ result ] trans: nothing
//
func Dcgettext(domain string, msgid string, category int32) (result string) {
	iv, err := _I.Get(890, "dcgettext", "", 416, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgid := gi.CString(msgid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgid := gi.NewStringArgument(c_msgid)
	arg_category := gi.NewInt32Argument(category)
	args := []gi.Argument{arg_domain, arg_msgid, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgid)
	result = ret.String().Copy()
	return
}

// g_dgettext
//
// [ domain ] trans: nothing
//
// [ msgid ] trans: nothing
//
// [ result ] trans: nothing
//
func Dgettext(domain string, msgid string) (result string) {
	iv, err := _I.Get(891, "dgettext", "", 417, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgid := gi.CString(msgid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgid := gi.NewStringArgument(c_msgid)
	args := []gi.Argument{arg_domain, arg_msgid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgid)
	result = ret.String().Copy()
	return
}

// g_dir_make_tmp
//
// [ tmpl ] trans: nothing
//
// [ result ] trans: everything
//
func DirMakeTmp(tmpl string) (result string, err error) {
	iv, err := _I.Get(892, "dir_make_tmp", "", 418, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_tmpl := gi.CString(tmpl)
	arg_tmpl := gi.NewStringArgument(c_tmpl)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_tmpl, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tmpl)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_direct_equal
//
// [ v1 ] trans: nothing
//
// [ v2 ] trans: nothing
//
// [ result ] trans: nothing
//
func DirectEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(893, "direct_equal", "", 419, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_direct_hash
//
// [ v ] trans: nothing
//
// [ result ] trans: nothing
//
func DirectHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(894, "direct_hash", "", 420, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_dngettext
//
// [ domain ] trans: nothing
//
// [ msgid ] trans: nothing
//
// [ msgid_plural ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: nothing
//
func Dngettext(domain string, msgid string, msgid_plural string, n uint64) (result string) {
	iv, err := _I.Get(895, "dngettext", "", 421, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgid := gi.CString(msgid)
	c_msgid_plural := gi.CString(msgid_plural)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgid := gi.NewStringArgument(c_msgid)
	arg_msgid_plural := gi.NewStringArgument(c_msgid_plural)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_domain, arg_msgid, arg_msgid_plural, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgid)
	gi.Free(c_msgid_plural)
	result = ret.String().Copy()
	return
}

// g_double_equal
//
// [ v1 ] trans: nothing
//
// [ v2 ] trans: nothing
//
// [ result ] trans: nothing
//
func DoubleEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(896, "double_equal", "", 422, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_double_hash
//
// [ v ] trans: nothing
//
// [ result ] trans: nothing
//
func DoubleHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(897, "double_hash", "", 423, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_dpgettext
//
// [ domain ] trans: nothing
//
// [ msgctxtid ] trans: nothing
//
// [ msgidoffset ] trans: nothing
//
// [ result ] trans: nothing
//
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) (result string) {
	iv, err := _I.Get(898, "dpgettext", "", 424, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgctxtid := gi.CString(msgctxtid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgctxtid := gi.NewStringArgument(c_msgctxtid)
	arg_msgidoffset := gi.NewUint64Argument(msgidoffset)
	args := []gi.Argument{arg_domain, arg_msgctxtid, arg_msgidoffset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgctxtid)
	result = ret.String().Copy()
	return
}

// g_dpgettext2
//
// [ domain ] trans: nothing
//
// [ context ] trans: nothing
//
// [ msgid ] trans: nothing
//
// [ result ] trans: nothing
//
func Dpgettext2(domain string, context string, msgid string) (result string) {
	iv, err := _I.Get(899, "dpgettext2", "", 425, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_context := gi.CString(context)
	c_msgid := gi.CString(msgid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_context := gi.NewStringArgument(c_context)
	arg_msgid := gi.NewStringArgument(c_msgid)
	args := []gi.Argument{arg_domain, arg_context, arg_msgid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_context)
	gi.Free(c_msgid)
	result = ret.String().Copy()
	return
}

// g_environ_getenv
//
// [ envp ] trans: nothing
//
// [ variable ] trans: nothing
//
// [ result ] trans: nothing
//
func EnvironGetenv(envp gi.CStrArray, variable string) (result string) {
	iv, err := _I.Get(900, "environ_getenv", "", 426, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_envp, arg_variable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	result = ret.String().Copy()
	return
}

// g_environ_setenv
//
// [ envp ] trans: everything
//
// [ variable ] trans: nothing
//
// [ value ] trans: nothing
//
// [ overwrite ] trans: nothing
//
// [ result ] trans: everything
//
func EnvironSetenv(envp gi.CStrArray, variable string, value string, overwrite bool) (result gi.CStrArray) {
	iv, err := _I.Get(901, "environ_setenv", "", 427, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	c_value := gi.CString(value)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_variable := gi.NewStringArgument(c_variable)
	arg_value := gi.NewStringArgument(c_value)
	arg_overwrite := gi.NewBoolArgument(overwrite)
	args := []gi.Argument{arg_envp, arg_variable, arg_value, arg_overwrite}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	gi.Free(c_value)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_environ_unsetenv
//
// [ envp ] trans: everything
//
// [ variable ] trans: nothing
//
// [ result ] trans: everything
//
func EnvironUnsetenv(envp gi.CStrArray, variable string) (result gi.CStrArray) {
	iv, err := _I.Get(902, "environ_unsetenv", "", 428, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_envp, arg_variable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_file_error_from_errno
//
// [ err_no ] trans: nothing
//
// [ result ] trans: nothing
//
func FileErrorFromErrno(err_no int32) (result FileErrorEnum) {
	iv, err := _I.Get(903, "file_error_from_errno", "", 429, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_err_no := gi.NewInt32Argument(err_no)
	args := []gi.Argument{arg_err_no}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FileErrorEnum(ret.Int())
	return
}

// g_file_error_quark
//
// [ result ] trans: nothing
//
func FileErrorQuark() (result uint32) {
	iv, err := _I.Get(904, "file_error_quark", "", 430, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_file_get_contents
//
// [ filename ] trans: nothing
//
// [ contents ] trans: everything, dir: out
//
// [ length ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func FileGetContents(filename string) (result bool, contents gi.Uint8Array, err error) {
	iv, err := _I.Get(905, "file_get_contents", "", 431, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_contents := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_filename, arg_contents, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[2].Pointer())
	contents.P = outArgs[0].Pointer()
	length = outArgs[1].Uint64()
	result = ret.Bool()
	contents.Len = int(length)
	return
}

// g_file_open_tmp
//
// [ tmpl ] trans: nothing
//
// [ name_used ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func FileOpenTmp(tmpl string) (result int32, name_used string, err error) {
	iv, err := _I.Get(906, "file_open_tmp", "", 432, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_tmpl := gi.CString(tmpl)
	arg_tmpl := gi.NewStringArgument(c_tmpl)
	arg_name_used := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_tmpl, arg_name_used, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tmpl)
	err = gi.ToError(outArgs[1].Pointer())
	name_used = outArgs[0].String().Take()
	result = ret.Int32()
	return
}

// g_file_read_link
//
// [ filename ] trans: nothing
//
// [ result ] trans: everything
//
func FileReadLink(filename string) (result string, err error) {
	iv, err := _I.Get(907, "file_read_link", "", 433, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_file_set_contents
//
// [ filename ] trans: nothing
//
// [ contents ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func FileSetContents(filename string, contents gi.Uint8Array, length int64) (result bool, err error) {
	iv, err := _I.Get(908, "file_set_contents", "", 434, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_contents := gi.NewPointerArgument(contents.P)
	arg_length := gi.NewInt64Argument(length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_contents, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_file_test
//
// [ filename ] trans: nothing
//
// [ test ] trans: nothing
//
// [ result ] trans: nothing
//
func FileTest(filename string, test FileTestFlags) (result bool) {
	iv, err := _I.Get(909, "file_test", "", 435, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_test := gi.NewIntArgument(int(test))
	args := []gi.Argument{arg_filename, arg_test}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Bool()
	return
}

// g_filename_display_basename
//
// [ filename ] trans: nothing
//
// [ result ] trans: everything
//
func FilenameDisplayBasename(filename string) (result string) {
	iv, err := _I.Get(910, "filename_display_basename", "", 436, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.String().Take()
	return
}

// g_filename_display_name
//
// [ filename ] trans: nothing
//
// [ result ] trans: everything
//
func FilenameDisplayName(filename string) (result string) {
	iv, err := _I.Get(911, "filename_display_name", "", 437, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.String().Take()
	return
}

// g_filename_from_uri
//
// [ uri ] trans: nothing
//
// [ hostname ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func FilenameFromUri(uri string) (result string, hostname string, err error) {
	iv, err := _I.Get(912, "filename_from_uri", "", 438, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_uri := gi.CString(uri)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_hostname := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_uri, arg_hostname, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[1].Pointer())
	hostname = outArgs[0].String().Take()
	result = ret.String().Take()
	return
}

// g_filename_from_utf8
//
// [ utf8string ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func FilenameFromUtf8(utf8string string, len1 int64) (result string, bytes_read uint64, bytes_written uint64, err error) {
	iv, err := _I.Get(913, "filename_from_utf8", "", 439, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_utf8string := gi.CString(utf8string)
	arg_utf8string := gi.NewStringArgument(c_utf8string)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_utf8string, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_utf8string)
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = ret.String().Take()
	return
}

// g_filename_to_uri
//
// [ filename ] trans: nothing
//
// [ hostname ] trans: nothing
//
// [ result ] trans: everything
//
func FilenameToUri(filename string, hostname string) (result string, err error) {
	iv, err := _I.Get(914, "filename_to_uri", "", 440, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	c_hostname := gi.CString(hostname)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_hostname := gi.NewStringArgument(c_hostname)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_hostname, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	gi.Free(c_hostname)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_filename_to_utf8
//
// [ opsysstring ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func FilenameToUtf8(opsysstring string, len1 int64) (result string, bytes_read uint64, bytes_written uint64, err error) {
	iv, err := _I.Get(915, "filename_to_utf8", "", 441, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_opsysstring := gi.CString(opsysstring)
	arg_opsysstring := gi.NewStringArgument(c_opsysstring)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_opsysstring, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_opsysstring)
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = ret.String().Take()
	return
}

// g_find_program_in_path
//
// [ program ] trans: nothing
//
// [ result ] trans: everything
//
func FindProgramInPath(program string) (result string) {
	iv, err := _I.Get(916, "find_program_in_path", "", 442, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_program := gi.CString(program)
	arg_program := gi.NewStringArgument(c_program)
	args := []gi.Argument{arg_program}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_program)
	result = ret.String().Take()
	return
}

// g_format_size
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func FormatSize(size uint64) (result string) {
	iv, err := _I.Get(917, "format_size", "", 443, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// Deprecated
//
// g_format_size_for_display
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func FormatSizeForDisplay(size int64) (result string) {
	iv, err := _I.Get(918, "format_size_for_display", "", 444, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_format_size_full
//
// [ size ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func FormatSizeFull(size uint64, flags FormatSizeFlags) (result string) {
	iv, err := _I.Get(919, "format_size_full", "", 445, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint64Argument(size)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_size, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_free
//
// [ mem ] trans: nothing
//
func Free(mem unsafe.Pointer) {
	iv, err := _I.Get(920, "free", "", 446, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	args := []gi.Argument{arg_mem}
	iv.Call(args, nil, nil)
}

// g_get_application_name
//
// [ result ] trans: nothing
//
func GetApplicationName() (result string) {
	iv, err := _I.Get(921, "get_application_name", "", 447, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_charset
//
// [ charset ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func GetCharset() (result bool, charset string) {
	iv, err := _I.Get(922, "get_charset", "", 448, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_charset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_charset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	charset = outArgs[0].String().Copy()
	result = ret.Bool()
	return
}

// g_get_codeset
//
// [ result ] trans: everything
//
func GetCodeset() (result string) {
	iv, err := _I.Get(923, "get_codeset", "", 449, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_current_dir
//
// [ result ] trans: everything
//
func GetCurrentDir() (result string) {
	iv, err := _I.Get(924, "get_current_dir", "", 450, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_current_time
//
// [ result ] trans: nothing
//
func GetCurrentTime(result TimeVal) {
	iv, err := _I.Get(925, "get_current_time", "", 451, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_result := gi.NewPointerArgument(result.P)
	args := []gi.Argument{arg_result}
	iv.Call(args, nil, nil)
}

// g_get_environ
//
// [ result ] trans: everything
//
func GetEnviron() (result gi.CStrArray) {
	iv, err := _I.Get(926, "get_environ", "", 452, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_filename_charsets
//
// [ filename_charsets ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func GetFilenameCharsets() (result bool, filename_charsets gi.CStrArray) {
	iv, err := _I.Get(927, "get_filename_charsets", "", 453, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_filename_charsets := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename_charsets}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	filename_charsets.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_get_home_dir
//
// [ result ] trans: nothing
//
func GetHomeDir() (result string) {
	iv, err := _I.Get(928, "get_home_dir", "", 454, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_host_name
//
// [ result ] trans: nothing
//
func GetHostName() (result string) {
	iv, err := _I.Get(929, "get_host_name", "", 455, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_language_names
//
// [ result ] trans: nothing
//
func GetLanguageNames() (result gi.CStrArray) {
	iv, err := _I.Get(930, "get_language_names", "", 456, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_language_names_with_category
//
// [ category_name ] trans: nothing
//
// [ result ] trans: nothing
//
func GetLanguageNamesWithCategory(category_name string) (result gi.CStrArray) {
	iv, err := _I.Get(931, "get_language_names_with_category", "", 457, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category_name := gi.CString(category_name)
	arg_category_name := gi.NewStringArgument(c_category_name)
	args := []gi.Argument{arg_category_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category_name)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_locale_variants
//
// [ locale ] trans: nothing
//
// [ result ] trans: everything
//
func GetLocaleVariants(locale string) (result gi.CStrArray) {
	iv, err := _I.Get(932, "get_locale_variants", "", 458, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_locale := gi.CString(locale)
	arg_locale := gi.NewStringArgument(c_locale)
	args := []gi.Argument{arg_locale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_locale)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_monotonic_time
//
// [ result ] trans: nothing
//
func GetMonotonicTime() (result int64) {
	iv, err := _I.Get(933, "get_monotonic_time", "", 459, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int64()
	return
}

// g_get_num_processors
//
// [ result ] trans: nothing
//
func GetNumProcessors() (result uint32) {
	iv, err := _I.Get(934, "get_num_processors", "", 460, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_get_prgname
//
// [ result ] trans: nothing
//
func GetPrgname() (result string) {
	iv, err := _I.Get(935, "get_prgname", "", 461, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_real_name
//
// [ result ] trans: nothing
//
func GetRealName() (result string) {
	iv, err := _I.Get(936, "get_real_name", "", 462, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_real_time
//
// [ result ] trans: nothing
//
func GetRealTime() (result int64) {
	iv, err := _I.Get(937, "get_real_time", "", 463, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int64()
	return
}

// g_get_system_config_dirs
//
// [ result ] trans: nothing
//
func GetSystemConfigDirs() (result gi.CStrArray) {
	iv, err := _I.Get(938, "get_system_config_dirs", "", 464, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_system_data_dirs
//
// [ result ] trans: nothing
//
func GetSystemDataDirs() (result gi.CStrArray) {
	iv, err := _I.Get(939, "get_system_data_dirs", "", 465, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_tmp_dir
//
// [ result ] trans: nothing
//
func GetTmpDir() (result string) {
	iv, err := _I.Get(940, "get_tmp_dir", "", 466, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_user_cache_dir
//
// [ result ] trans: nothing
//
func GetUserCacheDir() (result string) {
	iv, err := _I.Get(941, "get_user_cache_dir", "", 467, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_user_config_dir
//
// [ result ] trans: nothing
//
func GetUserConfigDir() (result string) {
	iv, err := _I.Get(942, "get_user_config_dir", "", 468, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_user_data_dir
//
// [ result ] trans: nothing
//
func GetUserDataDir() (result string) {
	iv, err := _I.Get(943, "get_user_data_dir", "", 469, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_user_name
//
// [ result ] trans: nothing
//
func GetUserName() (result string) {
	iv, err := _I.Get(944, "get_user_name", "", 470, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_user_runtime_dir
//
// [ result ] trans: nothing
//
func GetUserRuntimeDir() (result string) {
	iv, err := _I.Get(945, "get_user_runtime_dir", "", 471, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_get_user_special_dir
//
// [ directory ] trans: nothing
//
// [ result ] trans: nothing
//
func GetUserSpecialDir(directory UserDirectoryEnum) (result string) {
	iv, err := _I.Get(946, "get_user_special_dir", "", 472, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_directory := gi.NewIntArgument(int(directory))
	args := []gi.Argument{arg_directory}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_getenv
//
// [ variable ] trans: nothing
//
// [ result ] trans: nothing
//
func Getenv(variable string) (result string) {
	iv, err := _I.Get(947, "getenv", "", 473, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_variable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	result = ret.String().Copy()
	return
}

// g_hash_table_add
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableAdd(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(948, "hash_table_add", "", 474, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_contains
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableContains(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(949, "hash_table_contains", "", 475, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_destroy
//
// [ hash_table ] trans: nothing
//
func HashTableDestroy(hash_table HashTable) {
	iv, err := _I.Get(950, "hash_table_destroy", "", 476, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_insert
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableInsert(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(951, "hash_table_insert", "", 477, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_lookup
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableLookup(hash_table HashTable, key unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(952, "hash_table_lookup", "", 478, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_hash_table_lookup_extended
//
// [ hash_table ] trans: nothing
//
// [ lookup_key ] trans: nothing
//
// [ orig_key ] trans: everything, dir: out
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func HashTableLookupExtended(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, orig_key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(953, "hash_table_lookup_extended", "", 479, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_orig_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_orig_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	orig_key = outArgs[0].Pointer()
	value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_remove
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableRemove(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(954, "hash_table_remove", "", 480, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_remove_all
//
// [ hash_table ] trans: nothing
//
func HashTableRemoveAll(hash_table HashTable) {
	iv, err := _I.Get(955, "hash_table_remove_all", "", 481, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_replace
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableReplace(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(956, "hash_table_replace", "", 482, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_size
//
// [ hash_table ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableSize(hash_table HashTable) (result uint32) {
	iv, err := _I.Get(957, "hash_table_size", "", 483, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_hash_table_steal
//
// [ hash_table ] trans: nothing
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func HashTableSteal(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(958, "hash_table_steal", "", 484, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_steal_all
//
// [ hash_table ] trans: nothing
//
func HashTableStealAll(hash_table HashTable) {
	iv, err := _I.Get(959, "hash_table_steal_all", "", 485, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_steal_extended
//
// [ hash_table ] trans: nothing
//
// [ lookup_key ] trans: nothing
//
// [ stolen_key ] trans: everything, dir: out
//
// [ stolen_value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func HashTableStealExtended(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, stolen_key unsafe.Pointer, stolen_value unsafe.Pointer) {
	iv, err := _I.Get(960, "hash_table_steal_extended", "", 486, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_stolen_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stolen_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_stolen_key, arg_stolen_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	stolen_key = outArgs[0].Pointer()
	stolen_value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_unref
//
// [ hash_table ] trans: nothing
//
func HashTableUnref(hash_table HashTable) {
	iv, err := _I.Get(961, "hash_table_unref", "", 487, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hook_destroy
//
// [ hook_list ] trans: nothing
//
// [ hook_id ] trans: nothing
//
// [ result ] trans: nothing
//
func HookDestroy(hook_list HookList, hook_id uint64) (result bool) {
	iv, err := _I.Get(962, "hook_destroy", "", 488, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook_id := gi.NewUint64Argument(hook_id)
	args := []gi.Argument{arg_hook_list, arg_hook_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hook_destroy_link
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookDestroyLink(hook_list HookList, hook Hook) {
	iv, err := _I.Get(963, "hook_destroy_link", "", 489, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_free
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookFree(hook_list HookList, hook Hook) {
	iv, err := _I.Get(964, "hook_free", "", 490, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_insert_before
//
// [ hook_list ] trans: nothing
//
// [ sibling ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookInsertBefore(hook_list HookList, sibling Hook, hook Hook) {
	iv, err := _I.Get(965, "hook_insert_before", "", 491, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_sibling := gi.NewPointerArgument(sibling.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_sibling, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_prepend
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookPrepend(hook_list HookList, hook Hook) {
	iv, err := _I.Get(966, "hook_prepend", "", 492, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_unref
//
// [ hook_list ] trans: nothing
//
// [ hook ] trans: nothing
//
func HookUnref(hook_list HookList, hook Hook) {
	iv, err := _I.Get(967, "hook_unref", "", 493, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hostname_is_ascii_encoded
//
// [ hostname ] trans: nothing
//
// [ result ] trans: nothing
//
func HostnameIsAsciiEncoded(hostname string) (result bool) {
	iv, err := _I.Get(968, "hostname_is_ascii_encoded", "", 494, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.Bool()
	return
}

// g_hostname_is_ip_address
//
// [ hostname ] trans: nothing
//
// [ result ] trans: nothing
//
func HostnameIsIpAddress(hostname string) (result bool) {
	iv, err := _I.Get(969, "hostname_is_ip_address", "", 495, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.Bool()
	return
}

// g_hostname_is_non_ascii
//
// [ hostname ] trans: nothing
//
// [ result ] trans: nothing
//
func HostnameIsNonAscii(hostname string) (result bool) {
	iv, err := _I.Get(970, "hostname_is_non_ascii", "", 496, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.Bool()
	return
}

// g_hostname_to_ascii
//
// [ hostname ] trans: nothing
//
// [ result ] trans: everything
//
func HostnameToAscii(hostname string) (result string) {
	iv, err := _I.Get(971, "hostname_to_ascii", "", 497, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.String().Take()
	return
}

// g_hostname_to_unicode
//
// [ hostname ] trans: nothing
//
// [ result ] trans: everything
//
func HostnameToUnicode(hostname string) (result string) {
	iv, err := _I.Get(972, "hostname_to_unicode", "", 498, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.String().Take()
	return
}

// g_idle_add_full
//
// [ priority ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func IdleAdd(priority int32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(973, "idle_add", "", 499, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_idle_remove_by_data
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func IdleRemoveByData(data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(974, "idle_remove_by_data", "", 500, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_idle_source_new
//
// [ result ] trans: everything
//
func IdleSourceNew() (result Source) {
	iv, err := _I.Get(975, "idle_source_new", "", 501, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_int64_equal
//
// [ v1 ] trans: nothing
//
// [ v2 ] trans: nothing
//
// [ result ] trans: nothing
//
func Int64Equal(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(976, "int64_equal", "", 502, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_int64_hash
//
// [ v ] trans: nothing
//
// [ result ] trans: nothing
//
func Int64Hash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(977, "int64_hash", "", 503, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_int_equal
//
// [ v1 ] trans: nothing
//
// [ v2 ] trans: nothing
//
// [ result ] trans: nothing
//
func IntEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(978, "int_equal", "", 504, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_int_hash
//
// [ v ] trans: nothing
//
// [ result ] trans: nothing
//
func IntHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(979, "int_hash", "", 505, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_intern_static_string
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func InternStaticString(string string) (result string) {
	iv, err := _I.Get(980, "intern_static_string", "", 506, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Copy()
	return
}

// g_intern_string
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func InternString(string string) (result string) {
	iv, err := _I.Get(981, "intern_string", "", 507, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Copy()
	return
}

// g_io_add_watch_full
//
// [ channel ] trans: nothing
//
// [ priority ] trans: nothing
//
// [ condition ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func IoAddWatch(channel IOChannel, priority int32, condition IOConditionFlags, func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(982, "io_add_watch", "", 508, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_channel := gi.NewPointerArgument(channel.P)
	arg_priority := gi.NewInt32Argument(priority)
	arg_condition := gi.NewIntArgument(int(condition))
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myIOFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_channel, arg_priority, arg_condition, arg_func1, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_io_channel_error_from_errno
//
// [ en ] trans: nothing
//
// [ result ] trans: nothing
//
func IoChannelErrorFromErrno(en int32) (result IOChannelErrorEnum) {
	iv, err := _I.Get(983, "io_channel_error_from_errno", "", 509, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_en := gi.NewInt32Argument(en)
	args := []gi.Argument{arg_en}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOChannelErrorEnum(ret.Int())
	return
}

// g_io_channel_error_quark
//
// [ result ] trans: nothing
//
func IoChannelErrorQuark() (result uint32) {
	iv, err := _I.Get(984, "io_channel_error_quark", "", 510, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_io_create_watch
//
// [ channel ] trans: nothing
//
// [ condition ] trans: nothing
//
// [ result ] trans: everything
//
func IoCreateWatch(channel IOChannel, condition IOConditionFlags) (result Source) {
	iv, err := _I.Get(985, "io_create_watch", "", 511, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_channel := gi.NewPointerArgument(channel.P)
	arg_condition := gi.NewIntArgument(int(condition))
	args := []gi.Argument{arg_channel, arg_condition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_key_file_error_quark
//
// [ result ] trans: nothing
//
func KeyFileErrorQuark() (result uint32) {
	iv, err := _I.Get(986, "key_file_error_quark", "", 512, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_listenv
//
// [ result ] trans: everything
//
func Listenv() (result gi.CStrArray) {
	iv, err := _I.Get(987, "listenv", "", 513, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_locale_from_utf8
//
// [ utf8string ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func LocaleFromUtf8(utf8string string, len1 int64) (result gi.Uint8Array, bytes_read uint64, err error) {
	iv, err := _I.Get(988, "locale_from_utf8", "", 514, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_utf8string := gi.CString(utf8string)
	arg_utf8string := gi.NewStringArgument(c_utf8string)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_utf8string, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_utf8string)
	var bytes_written uint64
	_ = bytes_written
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(bytes_written)}
	return
}

// g_locale_to_utf8
//
// [ opsysstring ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ bytes_read ] trans: everything, dir: out
//
// [ bytes_written ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func LocaleToUtf8(opsysstring gi.Uint8Array, len1 int64) (result string, bytes_read uint64, bytes_written uint64, err error) {
	iv, err := _I.Get(989, "locale_to_utf8", "", 515, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_opsysstring := gi.NewPointerArgument(opsysstring.P)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_opsysstring, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = ret.String().Take()
	return
}

// g_log_default_handler
//
// [ log_domain ] trans: nothing
//
// [ log_level ] trans: nothing
//
// [ message ] trans: nothing
//
// [ unused_data ] trans: nothing
//
func LogDefaultHandler(log_domain string, log_level LogLevelFlags, message string, unused_data unsafe.Pointer) {
	iv, err := _I.Get(990, "log_default_handler", "", 516, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	c_message := gi.CString(message)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_message := gi.NewStringArgument(c_message)
	arg_unused_data := gi.NewPointerArgument(unused_data)
	args := []gi.Argument{arg_log_domain, arg_log_level, arg_message, arg_unused_data}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
	gi.Free(c_message)
}

// g_log_remove_handler
//
// [ log_domain ] trans: nothing
//
// [ handler_id ] trans: nothing
//
func LogRemoveHandler(log_domain string, handler_id uint32) {
	iv, err := _I.Get(991, "log_remove_handler", "", 517, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_handler_id := gi.NewUint32Argument(handler_id)
	args := []gi.Argument{arg_log_domain, arg_handler_id}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
}

// g_log_set_always_fatal
//
// [ fatal_mask ] trans: nothing
//
// [ result ] trans: nothing
//
func LogSetAlwaysFatal(fatal_mask LogLevelFlags) (result LogLevelFlags) {
	iv, err := _I.Get(992, "log_set_always_fatal", "", 518, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fatal_mask := gi.NewIntArgument(int(fatal_mask))
	args := []gi.Argument{arg_fatal_mask}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogLevelFlags(ret.Int())
	return
}

// g_log_set_fatal_mask
//
// [ log_domain ] trans: nothing
//
// [ fatal_mask ] trans: nothing
//
// [ result ] trans: nothing
//
func LogSetFatalMask(log_domain string, fatal_mask LogLevelFlags) (result LogLevelFlags) {
	iv, err := _I.Get(993, "log_set_fatal_mask", "", 519, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_fatal_mask := gi.NewIntArgument(int(fatal_mask))
	args := []gi.Argument{arg_log_domain, arg_fatal_mask}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_log_domain)
	result = LogLevelFlags(ret.Int())
	return
}

// g_log_set_handler_full
//
// [ log_domain ] trans: nothing
//
// [ log_levels ] trans: nothing
//
// [ log_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ destroy ] trans: nothing
//
// [ result ] trans: nothing
//
func LogSetHandler(log_domain string, log_levels LogLevelFlags, log_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, destroy int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(994, "log_set_handler", "", 520, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_levels := gi.NewIntArgument(int(log_levels))
	arg_log_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myLogFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_log_domain, arg_log_levels, arg_log_func, arg_user_data, arg_destroy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_log_domain)
	result = ret.Uint32()
	return
}

// g_log_set_writer_func
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ user_data_free ] trans: nothing
//
func LogSetWriterFunc(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, user_data_free int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(995, "log_set_writer_func", "", 521, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myLogWriterFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_user_data_free := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_func1, arg_user_data, arg_user_data_free}
	iv.Call(args, nil, nil)
}

// g_log_structured_array
//
// [ log_level ] trans: nothing
//
// [ fields ] trans: nothing
//
// [ n_fields ] trans: nothing
//
func LogStructuredArray(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64) {
	iv, err := _I.Get(996, "log_structured_array", "", 522, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields}
	iv.Call(args, nil, nil)
}

// g_log_variant
//
// [ log_domain ] trans: nothing
//
// [ log_level ] trans: nothing
//
// [ fields ] trans: nothing
//
func LogVariant(log_domain string, log_level LogLevelFlags, fields Variant) {
	iv, err := _I.Get(997, "log_variant", "", 523, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields.P)
	args := []gi.Argument{arg_log_domain, arg_log_level, arg_fields}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
}

// g_log_writer_default
//
// [ log_level ] trans: nothing
//
// [ fields ] trans: nothing
//
// [ n_fields ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func LogWriterDefault(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, user_data unsafe.Pointer) (result LogWriterOutputEnum) {
	iv, err := _I.Get(998, "log_writer_default", "", 524, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogWriterOutputEnum(ret.Int())
	return
}

// g_log_writer_format_fields
//
// [ log_level ] trans: nothing
//
// [ fields ] trans: nothing
//
// [ n_fields ] trans: nothing
//
// [ use_color ] trans: nothing
//
// [ result ] trans: everything
//
func LogWriterFormatFields(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, use_color bool) (result string) {
	iv, err := _I.Get(999, "log_writer_format_fields", "", 525, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_use_color := gi.NewBoolArgument(use_color)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_use_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_log_writer_is_journald
//
// [ output_fd ] trans: nothing
//
// [ result ] trans: nothing
//
func LogWriterIsJournald(output_fd int32) (result bool) {
	iv, err := _I.Get(1000, "log_writer_is_journald", "", 526, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_output_fd := gi.NewInt32Argument(output_fd)
	args := []gi.Argument{arg_output_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_log_writer_journald
//
// [ log_level ] trans: nothing
//
// [ fields ] trans: nothing
//
// [ n_fields ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func LogWriterJournald(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, user_data unsafe.Pointer) (result LogWriterOutputEnum) {
	iv, err := _I.Get(1001, "log_writer_journald", "", 527, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogWriterOutputEnum(ret.Int())
	return
}

// g_log_writer_standard_streams
//
// [ log_level ] trans: nothing
//
// [ fields ] trans: nothing
//
// [ n_fields ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func LogWriterStandardStreams(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, user_data unsafe.Pointer) (result LogWriterOutputEnum) {
	iv, err := _I.Get(1002, "log_writer_standard_streams", "", 528, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogWriterOutputEnum(ret.Int())
	return
}

// g_log_writer_supports_color
//
// [ output_fd ] trans: nothing
//
// [ result ] trans: nothing
//
func LogWriterSupportsColor(output_fd int32) (result bool) {
	iv, err := _I.Get(1003, "log_writer_supports_color", "", 529, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_output_fd := gi.NewInt32Argument(output_fd)
	args := []gi.Argument{arg_output_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_default
//
// [ result ] trans: nothing
//
func MainContextDefault() (result MainContext) {
	iv, err := _I.Get(1004, "main_context_default", "", 530, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_get_thread_default
//
// [ result ] trans: nothing
//
func MainContextGetThreadDefault() (result MainContext) {
	iv, err := _I.Get(1005, "main_context_get_thread_default", "", 531, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_ref_thread_default
//
// [ result ] trans: everything
//
func MainContextRefThreadDefault() (result MainContext) {
	iv, err := _I.Get(1006, "main_context_ref_thread_default", "", 532, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_current_source
//
// [ result ] trans: nothing
//
func MainCurrentSource() (result Source) {
	iv, err := _I.Get(1007, "main_current_source", "", 533, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_depth
//
// [ result ] trans: nothing
//
func MainDepth() (result int32) {
	iv, err := _I.Get(1008, "main_depth", "", 534, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_malloc
//
// [ n_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func Malloc(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1009, "malloc", "", 535, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_malloc0
//
// [ n_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func Malloc0(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1010, "malloc0", "", 536, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_malloc0_n
//
// [ n_blocks ] trans: nothing
//
// [ n_block_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func Malloc0N(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1011, "malloc0_n", "", 537, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_malloc_n
//
// [ n_blocks ] trans: nothing
//
// [ n_block_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func MallocN(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1012, "malloc_n", "", 538, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_markup_error_quark
//
// [ result ] trans: nothing
//
func MarkupErrorQuark() (result uint32) {
	iv, err := _I.Get(1013, "markup_error_quark", "", 539, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_markup_escape_text
//
// [ text ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func MarkupEscapeText(text string, length int64) (result string) {
	iv, err := _I.Get(1014, "markup_escape_text", "", 540, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_text, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result = ret.String().Take()
	return
}

// Deprecated
//
// g_mem_is_system_malloc
//
// [ result ] trans: nothing
//
func MemIsSystemMalloc() (result bool) {
	iv, err := _I.Get(1015, "mem_is_system_malloc", "", 541, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// g_mem_profile
//
func MemProfile() {
	iv, err := _I.Get(1016, "mem_profile", "", 542, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// Deprecated
//
// g_mem_set_vtable
//
// [ vtable ] trans: nothing
//
func MemSetVtable(vtable MemVTable) {
	iv, err := _I.Get(1017, "mem_set_vtable", "", 543, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_vtable := gi.NewPointerArgument(vtable.P)
	args := []gi.Argument{arg_vtable}
	iv.Call(args, nil, nil)
}

// g_memdup
//
// [ mem ] trans: nothing
//
// [ byte_size ] trans: nothing
//
// [ result ] trans: nothing
//
func Memdup(mem unsafe.Pointer, byte_size uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(1018, "memdup", "", 544, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_byte_size := gi.NewUint32Argument(byte_size)
	args := []gi.Argument{arg_mem, arg_byte_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_mkdir_with_parents
//
// [ pathname ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ result ] trans: nothing
//
func MkdirWithParents(pathname string, mode int32) (result int32) {
	iv, err := _I.Get(1019, "mkdir_with_parents", "", 545, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pathname := gi.CString(pathname)
	arg_pathname := gi.NewStringArgument(c_pathname)
	arg_mode := gi.NewInt32Argument(mode)
	args := []gi.Argument{arg_pathname, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pathname)
	result = ret.Int32()
	return
}

// g_nullify_pointer
//
// [ nullify_location ] trans: nothing
//
func NullifyPointer(nullify_location unsafe.Pointer) {
	iv, err := _I.Get(1020, "nullify_pointer", "", 546, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_nullify_location := gi.NewPointerArgument(nullify_location)
	args := []gi.Argument{arg_nullify_location}
	iv.Call(args, nil, nil)
}

// g_number_parser_error_quark
//
// [ result ] trans: nothing
//
func NumberParserErrorQuark() (result uint32) {
	iv, err := _I.Get(1021, "number_parser_error_quark", "", 547, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_on_error_query
//
// [ prg_name ] trans: nothing
//
func OnErrorQuery(prg_name string) {
	iv, err := _I.Get(1022, "on_error_query", "", 548, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_prg_name := gi.CString(prg_name)
	arg_prg_name := gi.NewStringArgument(c_prg_name)
	args := []gi.Argument{arg_prg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_prg_name)
}

// g_on_error_stack_trace
//
// [ prg_name ] trans: nothing
//
func OnErrorStackTrace(prg_name string) {
	iv, err := _I.Get(1023, "on_error_stack_trace", "", 549, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_prg_name := gi.CString(prg_name)
	arg_prg_name := gi.NewStringArgument(c_prg_name)
	args := []gi.Argument{arg_prg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_prg_name)
}

// g_once_init_enter
//
// [ location ] trans: nothing
//
// [ result ] trans: nothing
//
func OnceInitEnter(location unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1024, "once_init_enter", "", 550, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	args := []gi.Argument{arg_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_once_init_leave
//
// [ location ] trans: nothing
//
// [ result ] trans: nothing
//
func OnceInitLeave(location unsafe.Pointer, result uint64) {
	iv, err := _I.Get(1025, "once_init_leave", "", 551, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	arg_result := gi.NewUint64Argument(result)
	args := []gi.Argument{arg_location, arg_result}
	iv.Call(args, nil, nil)
}

// g_option_error_quark
//
// [ result ] trans: nothing
//
func OptionErrorQuark() (result uint32) {
	iv, err := _I.Get(1026, "option_error_quark", "", 552, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_parse_debug_string
//
// [ string ] trans: nothing
//
// [ keys ] trans: nothing
//
// [ nkeys ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseDebugString(string string, keys unsafe.Pointer, nkeys uint32) (result uint32) {
	iv, err := _I.Get(1027, "parse_debug_string", "", 553, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_keys := gi.NewPointerArgument(keys)
	arg_nkeys := gi.NewUint32Argument(nkeys)
	args := []gi.Argument{arg_string, arg_keys, arg_nkeys}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_path_get_basename
//
// [ file_name ] trans: nothing
//
// [ result ] trans: everything
//
func PathGetBasename(file_name string) (result string) {
	iv, err := _I.Get(1028, "path_get_basename", "", 554, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Take()
	return
}

// g_path_get_dirname
//
// [ file_name ] trans: nothing
//
// [ result ] trans: everything
//
func PathGetDirname(file_name string) (result string) {
	iv, err := _I.Get(1029, "path_get_dirname", "", 555, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Take()
	return
}

// g_path_is_absolute
//
// [ file_name ] trans: nothing
//
// [ result ] trans: nothing
//
func PathIsAbsolute(file_name string) (result bool) {
	iv, err := _I.Get(1030, "path_is_absolute", "", 556, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.Bool()
	return
}

// g_path_skip_root
//
// [ file_name ] trans: nothing
//
// [ result ] trans: nothing
//
func PathSkipRoot(file_name string) (result string) {
	iv, err := _I.Get(1031, "path_skip_root", "", 557, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Copy()
	return
}

// g_pattern_match
//
// [ pspec ] trans: nothing
//
// [ string_length ] trans: nothing
//
// [ string ] trans: nothing
//
// [ string_reversed ] trans: nothing
//
// [ result ] trans: nothing
//
func PatternMatch(pspec PatternSpec, string_length uint32, string string, string_reversed string) (result bool) {
	iv, err := _I.Get(1032, "pattern_match", "", 558, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	c_string_reversed := gi.CString(string_reversed)
	arg_pspec := gi.NewPointerArgument(pspec.P)
	arg_string_length := gi.NewUint32Argument(string_length)
	arg_string := gi.NewStringArgument(c_string)
	arg_string_reversed := gi.NewStringArgument(c_string_reversed)
	args := []gi.Argument{arg_pspec, arg_string_length, arg_string, arg_string_reversed}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	gi.Free(c_string_reversed)
	result = ret.Bool()
	return
}

// g_pattern_match_simple
//
// [ pattern ] trans: nothing
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func PatternMatchSimple(pattern string, string string) (result bool) {
	iv, err := _I.Get(1033, "pattern_match_simple", "", 559, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_pattern, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_pattern_match_string
//
// [ pspec ] trans: nothing
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func PatternMatchString(pspec PatternSpec, string string) (result bool) {
	iv, err := _I.Get(1034, "pattern_match_string", "", 560, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_pspec := gi.NewPointerArgument(pspec.P)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_pspec, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_pointer_bit_lock
//
// [ address ] trans: nothing
//
// [ lock_bit ] trans: nothing
//
func PointerBitLock(address unsafe.Pointer, lock_bit int32) {
	iv, err := _I.Get(1035, "pointer_bit_lock", "", 561, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewPointerArgument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_pointer_bit_trylock
//
// [ address ] trans: nothing
//
// [ lock_bit ] trans: nothing
//
// [ result ] trans: nothing
//
func PointerBitTrylock(address unsafe.Pointer, lock_bit int32) (result bool) {
	iv, err := _I.Get(1036, "pointer_bit_trylock", "", 562, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewPointerArgument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_pointer_bit_unlock
//
// [ address ] trans: nothing
//
// [ lock_bit ] trans: nothing
//
func PointerBitUnlock(address unsafe.Pointer, lock_bit int32) {
	iv, err := _I.Get(1037, "pointer_bit_unlock", "", 563, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewPointerArgument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_poll
//
// [ fds ] trans: nothing
//
// [ nfds ] trans: nothing
//
// [ timeout ] trans: nothing
//
// [ result ] trans: nothing
//
func Poll(fds PollFD, nfds uint32, timeout int32) (result int32) {
	iv, err := _I.Get(1038, "poll", "", 564, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fds := gi.NewPointerArgument(fds.P)
	arg_nfds := gi.NewUint32Argument(nfds)
	arg_timeout := gi.NewInt32Argument(timeout)
	args := []gi.Argument{arg_fds, arg_nfds, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_propagate_error
//
// [ dest ] trans: everything, dir: out
//
// [ src ] trans: everything
//
func PropagateError(src Error) (dest Error) {
	iv, err := _I.Get(1039, "propagate_error", "", 565, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_dest := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_src := gi.NewPointerArgument(src.P)
	args := []gi.Argument{arg_dest, arg_src}
	iv.Call(args, nil, &outArgs[0])
	dest.P = outArgs[0].Pointer()
	return
}

// g_quark_from_static_string
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func QuarkFromStaticString(string string) (result uint32) {
	iv, err := _I.Get(1040, "quark_from_static_string", "", 566, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_quark_from_string
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func QuarkFromString(string string) (result uint32) {
	iv, err := _I.Get(1041, "quark_from_string", "", 567, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_quark_to_string
//
// [ quark ] trans: nothing
//
// [ result ] trans: nothing
//
func QuarkToString(quark uint32) (result string) {
	iv, err := _I.Get(1042, "quark_to_string", "", 568, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_quark_try_string
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func QuarkTryString(string string) (result uint32) {
	iv, err := _I.Get(1043, "quark_try_string", "", 569, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_random_double
//
// [ result ] trans: nothing
//
func RandomDouble() (result float64) {
	iv, err := _I.Get(1044, "random_double", "", 570, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_random_double_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func RandomDoubleRange(begin float64, end float64) (result float64) {
	iv, err := _I.Get(1045, "random_double_range", "", 571, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewDoubleArgument(begin)
	arg_end := gi.NewDoubleArgument(end)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_random_int
//
// [ result ] trans: nothing
//
func RandomInt() (result uint32) {
	iv, err := _I.Get(1046, "random_int", "", 572, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_random_int_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func RandomIntRange(begin int32, end int32) (result int32) {
	iv, err := _I.Get(1047, "random_int_range", "", 573, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewInt32Argument(begin)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_random_set_seed
//
// [ seed ] trans: nothing
//
func RandomSetSeed(seed uint32) {
	iv, err := _I.Get(1048, "random_set_seed", "", 574, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_seed := gi.NewUint32Argument(seed)
	args := []gi.Argument{arg_seed}
	iv.Call(args, nil, nil)
}

// g_rc_box_acquire
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: everything
//
func RcBoxAcquire(mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1049, "rc_box_acquire", "", 575, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_alloc
//
// [ block_size ] trans: nothing
//
// [ result ] trans: everything
//
func RcBoxAlloc(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1050, "rc_box_alloc", "", 576, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_alloc0
//
// [ block_size ] trans: nothing
//
// [ result ] trans: everything
//
func RcBoxAlloc0(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1051, "rc_box_alloc0", "", 577, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_dup
//
// [ block_size ] trans: nothing
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: everything
//
func RcBoxDup(block_size uint64, mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1052, "rc_box_dup", "", 578, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_get_size
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: nothing
//
func RcBoxGetSize(mem_block unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(1053, "rc_box_get_size", "", 579, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_rc_box_release
//
// [ mem_block ] trans: everything
//
func RcBoxRelease(mem_block unsafe.Pointer) {
	iv, err := _I.Get(1054, "rc_box_release", "", 580, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	iv.Call(args, nil, nil)
}

// g_rc_box_release_full
//
// [ mem_block ] trans: everything
//
// [ clear_func ] trans: nothing
//
func RcBoxReleaseFull(mem_block unsafe.Pointer, clear_func int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1055, "rc_box_release_full", "", 581, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	arg_clear_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_mem_block, arg_clear_func}
	iv.Call(args, nil, nil)
}

// g_realloc
//
// [ mem ] trans: nothing
//
// [ n_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func Realloc(mem unsafe.Pointer, n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1056, "realloc", "", 582, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_mem, arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_realloc_n
//
// [ mem ] trans: nothing
//
// [ n_blocks ] trans: nothing
//
// [ n_block_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func ReallocN(mem unsafe.Pointer, n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1057, "realloc_n", "", 583, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_mem, arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_ref_count_compare
//
// [ rc ] trans: nothing
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func RefCountCompare(rc int32, val int32) (result bool) {
	iv, err := _I.Get(1058, "ref_count_compare", "", 584, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_rc, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_ref_count_dec
//
// [ rc ] trans: nothing
//
// [ result ] trans: nothing
//
func RefCountDec(rc int32) (result bool) {
	iv, err := _I.Get(1059, "ref_count_dec", "", 585, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	args := []gi.Argument{arg_rc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_ref_count_inc
//
// [ rc ] trans: nothing
//
func RefCountInc(rc int32) {
	iv, err := _I.Get(1060, "ref_count_inc", "", 586, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	args := []gi.Argument{arg_rc}
	iv.Call(args, nil, nil)
}

// g_ref_count_init
//
// [ rc ] trans: nothing
//
func RefCountInit(rc int32) {
	iv, err := _I.Get(1061, "ref_count_init", "", 587, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	args := []gi.Argument{arg_rc}
	iv.Call(args, nil, nil)
}

// g_ref_string_acquire
//
// [ str ] trans: nothing
//
// [ result ] trans: everything
//
func RefStringAcquire(str string) (result string) {
	iv, err := _I.Get(1062, "ref_string_acquire", "", 588, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ref_string_length
//
// [ str ] trans: nothing
//
// [ result ] trans: nothing
//
func RefStringLength(str string) (result uint64) {
	iv, err := _I.Get(1063, "ref_string_length", "", 589, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.Uint64()
	return
}

// g_ref_string_new
//
// [ str ] trans: nothing
//
// [ result ] trans: everything
//
func RefStringNew(str string) (result string) {
	iv, err := _I.Get(1064, "ref_string_new", "", 590, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ref_string_new_intern
//
// [ str ] trans: nothing
//
// [ result ] trans: everything
//
func RefStringNewIntern(str string) (result string) {
	iv, err := _I.Get(1065, "ref_string_new_intern", "", 591, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ref_string_new_len
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func RefStringNewLen(str string, len1 int64) (result string) {
	iv, err := _I.Get(1066, "ref_string_new_len", "", 592, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ref_string_release
//
// [ str ] trans: nothing
//
func RefStringRelease(str string) {
	iv, err := _I.Get(1067, "ref_string_release", "", 593, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	iv.Call(args, nil, nil)
	gi.Free(c_str)
}

// g_regex_check_replacement
//
// [ replacement ] trans: nothing
//
// [ has_references ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func RegexCheckReplacement(replacement string) (result bool, has_references bool, err error) {
	iv, err := _I.Get(1068, "regex_check_replacement", "", 594, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_has_references := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_replacement, arg_has_references, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[1].Pointer())
	has_references = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// g_regex_error_quark
//
// [ result ] trans: nothing
//
func RegexErrorQuark() (result uint32) {
	iv, err := _I.Get(1069, "regex_error_quark", "", 595, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_regex_escape_nul
//
// [ string ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func RegexEscapeNul(string string, length int32) (result string) {
	iv, err := _I.Get(1070, "regex_escape_nul", "", 596, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_regex_escape_string
//
// [ string ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func RegexEscapeString(string gi.CStrArray, length int32) (result string) {
	iv, err := _I.Get(1071, "regex_escape_string", "", 597, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewPointerArgument(string.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_regex_match_simple
//
// [ pattern ] trans: nothing
//
// [ string ] trans: nothing
//
// [ compile_options ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: nothing
//
func RegexMatchSimple(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result bool) {
	iv, err := _I.Get(1072, "regex_match_simple", "", 598, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_regex_split_simple
//
// [ pattern ] trans: nothing
//
// [ string ] trans: nothing
//
// [ compile_options ] trans: nothing
//
// [ match_options ] trans: nothing
//
// [ result ] trans: everything
//
func RegexSplitSimple(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result gi.CStrArray) {
	iv, err := _I.Get(1073, "regex_split_simple", "", 599, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_reload_user_special_dirs_cache
//
func ReloadUserSpecialDirsCache() {
	iv, err := _I.Get(1074, "reload_user_special_dirs_cache", "", 600, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_rmdir
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func Rmdir(filename string) (result int32) {
	iv, err := _I.Get(1075, "rmdir", "", 601, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Int32()
	return
}

// g_sequence_get
//
// [ iter ] trans: nothing
//
// [ result ] trans: nothing
//
func SequenceGet(iter SequenceIter) (result unsafe.Pointer) {
	iv, err := _I.Get(1076, "sequence_get", "", 602, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_sequence_insert_before
//
// [ iter ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func SequenceInsertBefore(iter SequenceIter, data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(1077, "sequence_insert_before", "", 603, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_move
//
// [ src ] trans: nothing
//
// [ dest ] trans: nothing
//
func SequenceMove(src SequenceIter, dest SequenceIter) {
	iv, err := _I.Get(1078, "sequence_move", "", 604, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src := gi.NewPointerArgument(src.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_src, arg_dest}
	iv.Call(args, nil, nil)
}

// g_sequence_move_range
//
// [ dest ] trans: nothing
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
func SequenceMoveRange(dest SequenceIter, begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(1079, "sequence_move_range", "", 605, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_dest, arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_range_get_midpoint
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func SequenceRangeGetMidpoint(begin SequenceIter, end SequenceIter) (result SequenceIter) {
	iv, err := _I.Get(1080, "sequence_range_get_midpoint", "", 606, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_remove
//
// [ iter ] trans: nothing
//
func SequenceRemove(iter SequenceIter) {
	iv, err := _I.Get(1081, "sequence_remove", "", 607, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	iv.Call(args, nil, nil)
}

// g_sequence_remove_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
func SequenceRemoveRange(begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(1082, "sequence_remove_range", "", 608, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_set
//
// [ iter ] trans: nothing
//
// [ data ] trans: nothing
//
func SequenceSet(iter SequenceIter, data unsafe.Pointer) {
	iv, err := _I.Get(1083, "sequence_set", "", 609, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	iv.Call(args, nil, nil)
}

// g_sequence_swap
//
// [ a ] trans: nothing
//
// [ b ] trans: nothing
//
func SequenceSwap(a SequenceIter, b SequenceIter) {
	iv, err := _I.Get(1084, "sequence_swap", "", 610, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewPointerArgument(a.P)
	arg_b := gi.NewPointerArgument(b.P)
	args := []gi.Argument{arg_a, arg_b}
	iv.Call(args, nil, nil)
}

// g_set_application_name
//
// [ application_name ] trans: nothing
//
func SetApplicationName(application_name string) {
	iv, err := _I.Get(1085, "set_application_name", "", 611, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_application_name := gi.CString(application_name)
	arg_application_name := gi.NewStringArgument(c_application_name)
	args := []gi.Argument{arg_application_name}
	iv.Call(args, nil, nil)
	gi.Free(c_application_name)
}

// g_set_error_literal
//
// [ err1 ] trans: everything, dir: out
//
// [ domain ] trans: nothing
//
// [ code ] trans: nothing
//
// [ message ] trans: nothing
//
func SetErrorLiteral(domain uint32, code int32, message string) (err1 Error) {
	iv, err := _I.Get(1086, "set_error_literal", "", 612, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_message := gi.CString(message)
	arg_err1 := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	arg_message := gi.NewStringArgument(c_message)
	args := []gi.Argument{arg_err1, arg_domain, arg_code, arg_message}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_message)
	err1.P = outArgs[0].Pointer()
	return
}

// g_set_prgname
//
// [ prgname ] trans: nothing
//
func SetPrgname(prgname string) {
	iv, err := _I.Get(1087, "set_prgname", "", 613, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_prgname := gi.CString(prgname)
	arg_prgname := gi.NewStringArgument(c_prgname)
	args := []gi.Argument{arg_prgname}
	iv.Call(args, nil, nil)
	gi.Free(c_prgname)
}

// g_setenv
//
// [ variable ] trans: nothing
//
// [ value ] trans: nothing
//
// [ overwrite ] trans: nothing
//
// [ result ] trans: nothing
//
func Setenv(variable string, value string, overwrite bool) (result bool) {
	iv, err := _I.Get(1088, "setenv", "", 614, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	c_value := gi.CString(value)
	arg_variable := gi.NewStringArgument(c_variable)
	arg_value := gi.NewStringArgument(c_value)
	arg_overwrite := gi.NewBoolArgument(overwrite)
	args := []gi.Argument{arg_variable, arg_value, arg_overwrite}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	gi.Free(c_value)
	result = ret.Bool()
	return
}

// g_shell_error_quark
//
// [ result ] trans: nothing
//
func ShellErrorQuark() (result uint32) {
	iv, err := _I.Get(1089, "shell_error_quark", "", 615, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_shell_parse_argv
//
// [ command_line ] trans: nothing
//
// [ argcp ] trans: everything, dir: out
//
// [ argvp ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func ShellParseArgv(command_line string) (result bool, argvp gi.CStrArray, err error) {
	iv, err := _I.Get(1090, "shell_parse_argv", "", 616, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_command_line := gi.CString(command_line)
	arg_command_line := gi.NewStringArgument(c_command_line)
	arg_argcp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_argvp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_command_line, arg_argcp, arg_argvp, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_command_line)
	var argcp int32
	_ = argcp
	err = gi.ToError(outArgs[2].Pointer())
	argcp = outArgs[0].Int32()
	argvp.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_shell_quote
//
// [ unquoted_string ] trans: nothing
//
// [ result ] trans: everything
//
func ShellQuote(unquoted_string string) (result string) {
	iv, err := _I.Get(1091, "shell_quote", "", 617, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_unquoted_string := gi.CString(unquoted_string)
	arg_unquoted_string := gi.NewStringArgument(c_unquoted_string)
	args := []gi.Argument{arg_unquoted_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_unquoted_string)
	result = ret.String().Take()
	return
}

// g_shell_unquote
//
// [ quoted_string ] trans: nothing
//
// [ result ] trans: everything
//
func ShellUnquote(quoted_string string) (result string, err error) {
	iv, err := _I.Get(1092, "shell_unquote", "", 618, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_quoted_string := gi.CString(quoted_string)
	arg_quoted_string := gi.NewStringArgument(c_quoted_string)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_quoted_string, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_quoted_string)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_slice_alloc
//
// [ block_size ] trans: nothing
//
// [ result ] trans: nothing
//
func SliceAlloc(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1093, "slice_alloc", "", 619, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_slice_alloc0
//
// [ block_size ] trans: nothing
//
// [ result ] trans: nothing
//
func SliceAlloc0(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1094, "slice_alloc0", "", 620, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_slice_copy
//
// [ block_size ] trans: nothing
//
// [ mem_block ] trans: nothing
//
// [ result ] trans: nothing
//
func SliceCopy(block_size uint64, mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1095, "slice_copy", "", 621, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_slice_free1
//
// [ block_size ] trans: nothing
//
// [ mem_block ] trans: nothing
//
func SliceFree1(block_size uint64, mem_block unsafe.Pointer) {
	iv, err := _I.Get(1096, "slice_free1", "", 622, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	iv.Call(args, nil, nil)
}

// g_slice_free_chain_with_offset
//
// [ block_size ] trans: nothing
//
// [ mem_chain ] trans: nothing
//
// [ next_offset ] trans: nothing
//
func SliceFreeChainWithOffset(block_size uint64, mem_chain unsafe.Pointer, next_offset uint64) {
	iv, err := _I.Get(1097, "slice_free_chain_with_offset", "", 623, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_chain := gi.NewPointerArgument(mem_chain)
	arg_next_offset := gi.NewUint64Argument(next_offset)
	args := []gi.Argument{arg_block_size, arg_mem_chain, arg_next_offset}
	iv.Call(args, nil, nil)
}

// g_slice_get_config
//
// [ ckey ] trans: nothing
//
// [ result ] trans: nothing
//
func SliceGetConfig(ckey SliceConfigEnum) (result int64) {
	iv, err := _I.Get(1098, "slice_get_config", "", 624, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ckey := gi.NewIntArgument(int(ckey))
	args := []gi.Argument{arg_ckey}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_slice_get_config_state
//
// [ ckey ] trans: nothing
//
// [ address ] trans: nothing
//
// [ n_values ] trans: nothing
//
// [ result ] trans: nothing
//
func SliceGetConfigState(ckey SliceConfigEnum, address int64, n_values uint32) (result int64) {
	iv, err := _I.Get(1099, "slice_get_config_state", "", 625, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ckey := gi.NewIntArgument(int(ckey))
	arg_address := gi.NewInt64Argument(address)
	arg_n_values := gi.NewUint32Argument(n_values)
	args := []gi.Argument{arg_ckey, arg_address, arg_n_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_slice_set_config
//
// [ ckey ] trans: nothing
//
// [ value ] trans: nothing
//
func SliceSetConfig(ckey SliceConfigEnum, value int64) {
	iv, err := _I.Get(1100, "slice_set_config", "", 626, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ckey := gi.NewIntArgument(int(ckey))
	arg_value := gi.NewInt64Argument(value)
	args := []gi.Argument{arg_ckey, arg_value}
	iv.Call(args, nil, nil)
}

// g_source_remove
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func SourceRemove(tag uint32) (result bool) {
	iv, err := _I.Get(1101, "source_remove", "", 627, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tag := gi.NewUint32Argument(tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_funcs_user_data
//
// [ funcs ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func SourceRemoveByFuncsUserData(funcs SourceFuncs, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1102, "source_remove_by_funcs_user_data", "", 628, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_funcs := gi.NewPointerArgument(funcs.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_funcs, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_user_data
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func SourceRemoveByUserData(user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1103, "source_remove_by_user_data", "", 629, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_set_name_by_id
//
// [ tag ] trans: nothing
//
// [ name ] trans: nothing
//
func SourceSetNameById(tag uint32, name string) {
	iv, err := _I.Get(1104, "source_set_name_by_id", "", 630, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_tag := gi.NewUint32Argument(tag)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_tag, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// g_spaced_primes_closest
//
// [ num ] trans: nothing
//
// [ result ] trans: nothing
//
func SpacedPrimesClosest(num uint32) (result uint32) {
	iv, err := _I.Get(1105, "spaced_primes_closest", "", 631, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_num := gi.NewUint32Argument(num)
	args := []gi.Argument{arg_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_spawn_async
//
// [ working_directory ] trans: nothing
//
// [ argv ] trans: nothing
//
// [ envp ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ child_setup ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ child_pid ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func SpawnAsync(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool, child_pid int32, err error) {
	iv, err := _I.Get(1106, "spawn_async", "", 632, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySpawnChildSetupFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_child_pid, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[1].Pointer())
	child_pid = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// g_spawn_async_with_fds
//
// [ working_directory ] trans: nothing
//
// [ argv ] trans: nothing
//
// [ envp ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ child_setup ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ child_pid ] trans: everything, dir: out
//
// [ stdin_fd ] trans: nothing
//
// [ stdout_fd ] trans: nothing
//
// [ stderr_fd ] trans: nothing
//
// [ result ] trans: nothing
//
func SpawnAsyncWithFds(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, stdin_fd int32, stdout_fd int32, stderr_fd int32) (result bool, child_pid int32, err error) {
	iv, err := _I.Get(1107, "spawn_async_with_fds", "", 633, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySpawnChildSetupFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stdin_fd := gi.NewInt32Argument(stdin_fd)
	arg_stdout_fd := gi.NewInt32Argument(stdout_fd)
	arg_stderr_fd := gi.NewInt32Argument(stderr_fd)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_child_pid, arg_stdin_fd, arg_stdout_fd, arg_stderr_fd, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[1].Pointer())
	child_pid = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// g_spawn_async_with_pipes
//
// [ working_directory ] trans: nothing
//
// [ argv ] trans: nothing
//
// [ envp ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ child_setup ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ child_pid ] trans: everything, dir: out
//
// [ standard_input ] trans: everything, dir: out
//
// [ standard_output ] trans: everything, dir: out
//
// [ standard_error ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func SpawnAsyncWithPipes(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool, child_pid int32, standard_input int32, standard_output int32, standard_error int32, err error) {
	iv, err := _I.Get(1108, "spawn_async_with_pipes", "", 634, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [5]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySpawnChildSetupFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_standard_input := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_standard_output := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_standard_error := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_child_pid, arg_standard_input, arg_standard_output, arg_standard_error, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[4].Pointer())
	child_pid = outArgs[0].Int32()
	standard_input = outArgs[1].Int32()
	standard_output = outArgs[2].Int32()
	standard_error = outArgs[3].Int32()
	result = ret.Bool()
	return
}

// g_spawn_check_exit_status
//
// [ exit_status ] trans: nothing
//
// [ result ] trans: nothing
//
func SpawnCheckExitStatus(exit_status int32) (result bool, err error) {
	iv, err := _I.Get(1109, "spawn_check_exit_status", "", 635, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_exit_status := gi.NewInt32Argument(exit_status)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_exit_status, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_spawn_close_pid
//
// [ pid ] trans: nothing
//
func SpawnClosePid(pid int32) {
	iv, err := _I.Get(1110, "spawn_close_pid", "", 636, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_pid}
	iv.Call(args, nil, nil)
}

// g_spawn_command_line_async
//
// [ command_line ] trans: nothing
//
// [ result ] trans: nothing
//
func SpawnCommandLineAsync(command_line string) (result bool, err error) {
	iv, err := _I.Get(1111, "spawn_command_line_async", "", 637, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_command_line := gi.CString(command_line)
	arg_command_line := gi.NewStringArgument(c_command_line)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_command_line, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_command_line)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_spawn_command_line_sync
//
// [ command_line ] trans: nothing
//
// [ standard_output ] trans: everything, dir: out
//
// [ standard_error ] trans: everything, dir: out
//
// [ exit_status ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func SpawnCommandLineSync(command_line string) (result bool, standard_output gi.Uint8Array, standard_error gi.Uint8Array, exit_status int32, err error) {
	iv, err := _I.Get(1112, "spawn_command_line_sync", "", 638, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_command_line := gi.CString(command_line)
	arg_command_line := gi.NewStringArgument(c_command_line)
	arg_standard_output := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_standard_error := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_exit_status := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_command_line, arg_standard_output, arg_standard_error, arg_exit_status, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_command_line)
	err = gi.ToError(outArgs[3].Pointer())
	standard_output.P = outArgs[0].Pointer()
	standard_error.P = outArgs[1].Pointer()
	exit_status = outArgs[2].Int32()
	result = ret.Bool()
	return
}

// g_spawn_error_quark
//
// [ result ] trans: nothing
//
func SpawnErrorQuark() (result uint32) {
	iv, err := _I.Get(1113, "spawn_error_quark", "", 639, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_spawn_exit_error_quark
//
// [ result ] trans: nothing
//
func SpawnExitErrorQuark() (result uint32) {
	iv, err := _I.Get(1114, "spawn_exit_error_quark", "", 640, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_spawn_sync
//
// [ working_directory ] trans: nothing
//
// [ argv ] trans: nothing
//
// [ envp ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ child_setup ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ standard_output ] trans: everything, dir: out
//
// [ standard_error ] trans: everything, dir: out
//
// [ exit_status ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func SpawnSync(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool, standard_output gi.Uint8Array, standard_error gi.Uint8Array, exit_status int32, err error) {
	iv, err := _I.Get(1115, "spawn_sync", "", 641, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySpawnChildSetupFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_standard_output := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_standard_error := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_exit_status := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_standard_output, arg_standard_error, arg_exit_status, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[3].Pointer())
	standard_output.P = outArgs[0].Pointer()
	standard_error.P = outArgs[1].Pointer()
	exit_status = outArgs[2].Int32()
	result = ret.Bool()
	return
}

// g_stpcpy
//
// [ dest ] trans: nothing
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func Stpcpy(dest string, src string) (result string) {
	iv, err := _I.Get(1116, "stpcpy", "", 642, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	args := []gi.Argument{arg_dest, arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.String().Take()
	return
}

// g_str_equal
//
// [ v1 ] trans: nothing
//
// [ v2 ] trans: nothing
//
// [ result ] trans: nothing
//
func StrEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1117, "str_equal", "", 643, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_str_has_prefix
//
// [ str ] trans: nothing
//
// [ prefix ] trans: nothing
//
// [ result ] trans: nothing
//
func StrHasPrefix(str string, prefix string) (result bool) {
	iv, err := _I.Get(1118, "str_has_prefix", "", 644, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_prefix := gi.CString(prefix)
	arg_str := gi.NewStringArgument(c_str)
	arg_prefix := gi.NewStringArgument(c_prefix)
	args := []gi.Argument{arg_str, arg_prefix}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_prefix)
	result = ret.Bool()
	return
}

// g_str_has_suffix
//
// [ str ] trans: nothing
//
// [ suffix ] trans: nothing
//
// [ result ] trans: nothing
//
func StrHasSuffix(str string, suffix string) (result bool) {
	iv, err := _I.Get(1119, "str_has_suffix", "", 645, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_suffix := gi.CString(suffix)
	arg_str := gi.NewStringArgument(c_str)
	arg_suffix := gi.NewStringArgument(c_suffix)
	args := []gi.Argument{arg_str, arg_suffix}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_suffix)
	result = ret.Bool()
	return
}

// g_str_hash
//
// [ v ] trans: nothing
//
// [ result ] trans: nothing
//
func StrHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(1120, "str_hash", "", 646, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_str_is_ascii
//
// [ str ] trans: nothing
//
// [ result ] trans: nothing
//
func StrIsAscii(str string) (result bool) {
	iv, err := _I.Get(1121, "str_is_ascii", "", 647, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.Bool()
	return
}

// g_str_match_string
//
// [ search_term ] trans: nothing
//
// [ potential_hit ] trans: nothing
//
// [ accept_alternates ] trans: nothing
//
// [ result ] trans: nothing
//
func StrMatchString(search_term string, potential_hit string, accept_alternates bool) (result bool) {
	iv, err := _I.Get(1122, "str_match_string", "", 648, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_search_term := gi.CString(search_term)
	c_potential_hit := gi.CString(potential_hit)
	arg_search_term := gi.NewStringArgument(c_search_term)
	arg_potential_hit := gi.NewStringArgument(c_potential_hit)
	arg_accept_alternates := gi.NewBoolArgument(accept_alternates)
	args := []gi.Argument{arg_search_term, arg_potential_hit, arg_accept_alternates}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_search_term)
	gi.Free(c_potential_hit)
	result = ret.Bool()
	return
}

// g_str_to_ascii
//
// [ str ] trans: nothing
//
// [ from_locale ] trans: nothing
//
// [ result ] trans: everything
//
func StrToAscii(str string, from_locale string) (result string) {
	iv, err := _I.Get(1123, "str_to_ascii", "", 649, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_from_locale := gi.CString(from_locale)
	arg_str := gi.NewStringArgument(c_str)
	arg_from_locale := gi.NewStringArgument(c_from_locale)
	args := []gi.Argument{arg_str, arg_from_locale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_from_locale)
	result = ret.String().Take()
	return
}

// g_str_tokenize_and_fold
//
// [ string ] trans: nothing
//
// [ translit_locale ] trans: nothing
//
// [ ascii_alternates ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func StrTokenizeAndFold(string string, translit_locale string) (result gi.CStrArray, ascii_alternates gi.CStrArray) {
	iv, err := _I.Get(1124, "str_tokenize_and_fold", "", 650, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	c_translit_locale := gi.CString(translit_locale)
	arg_string := gi.NewStringArgument(c_string)
	arg_translit_locale := gi.NewStringArgument(c_translit_locale)
	arg_ascii_alternates := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_translit_locale, arg_ascii_alternates}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	gi.Free(c_translit_locale)
	ascii_alternates.P = outArgs[0].Pointer()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_strcanon
//
// [ string ] trans: nothing
//
// [ valid_chars ] trans: nothing
//
// [ substitutor ] trans: nothing
//
// [ result ] trans: everything
//
func Strcanon(string string, valid_chars string, substitutor int8) (result string) {
	iv, err := _I.Get(1125, "strcanon", "", 651, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	c_valid_chars := gi.CString(valid_chars)
	arg_string := gi.NewStringArgument(c_string)
	arg_valid_chars := gi.NewStringArgument(c_valid_chars)
	arg_substitutor := gi.NewInt8Argument(substitutor)
	args := []gi.Argument{arg_string, arg_valid_chars, arg_substitutor}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	gi.Free(c_valid_chars)
	result = ret.String().Take()
	return
}

// Deprecated
//
// g_strcasecmp
//
// [ s1 ] trans: nothing
//
// [ s2 ] trans: nothing
//
// [ result ] trans: nothing
//
func Strcasecmp(s1 string, s2 string) (result int32) {
	iv, err := _I.Get(1126, "strcasecmp", "", 652, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	args := []gi.Argument{arg_s1, arg_s2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_strchomp
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func Strchomp(string string) (result string) {
	iv, err := _I.Get(1127, "strchomp", "", 653, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strchug
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func Strchug(string string) (result string) {
	iv, err := _I.Get(1128, "strchug", "", 654, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strcmp0
//
// [ str1 ] trans: nothing
//
// [ str2 ] trans: nothing
//
// [ result ] trans: nothing
//
func Strcmp0(str1 string, str2 string) (result int32) {
	iv, err := _I.Get(1129, "strcmp0", "", 655, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str1 := gi.CString(str1)
	c_str2 := gi.CString(str2)
	arg_str1 := gi.NewStringArgument(c_str1)
	arg_str2 := gi.NewStringArgument(c_str2)
	args := []gi.Argument{arg_str1, arg_str2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str1)
	gi.Free(c_str2)
	result = ret.Int32()
	return
}

// g_strcompress
//
// [ source ] trans: nothing
//
// [ result ] trans: everything
//
func Strcompress(source string) (result string) {
	iv, err := _I.Get(1130, "strcompress", "", 656, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source := gi.CString(source)
	arg_source := gi.NewStringArgument(c_source)
	args := []gi.Argument{arg_source}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source)
	result = ret.String().Take()
	return
}

// g_strdelimit
//
// [ string ] trans: nothing
//
// [ delimiters ] trans: nothing
//
// [ new_delimiter ] trans: nothing
//
// [ result ] trans: everything
//
func Strdelimit(string string, delimiters string, new_delimiter int8) (result string) {
	iv, err := _I.Get(1131, "strdelimit", "", 657, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	c_delimiters := gi.CString(delimiters)
	arg_string := gi.NewStringArgument(c_string)
	arg_delimiters := gi.NewStringArgument(c_delimiters)
	arg_new_delimiter := gi.NewInt8Argument(new_delimiter)
	args := []gi.Argument{arg_string, arg_delimiters, arg_new_delimiter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	gi.Free(c_delimiters)
	result = ret.String().Take()
	return
}

// Deprecated
//
// g_strdown
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func Strdown(string string) (result string) {
	iv, err := _I.Get(1132, "strdown", "", 658, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strdup
//
// [ str ] trans: nothing
//
// [ result ] trans: everything
//
func Strdup(str string) (result string) {
	iv, err := _I.Get(1133, "strdup", "", 659, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_strerror
//
// [ errnum ] trans: nothing
//
// [ result ] trans: nothing
//
func Strerror(errnum int32) (result string) {
	iv, err := _I.Get(1134, "strerror", "", 660, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_errnum := gi.NewInt32Argument(errnum)
	args := []gi.Argument{arg_errnum}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_strescape
//
// [ source ] trans: nothing
//
// [ exceptions ] trans: nothing
//
// [ result ] trans: everything
//
func Strescape(source string, exceptions string) (result string) {
	iv, err := _I.Get(1135, "strescape", "", 661, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source := gi.CString(source)
	c_exceptions := gi.CString(exceptions)
	arg_source := gi.NewStringArgument(c_source)
	arg_exceptions := gi.NewStringArgument(c_exceptions)
	args := []gi.Argument{arg_source, arg_exceptions}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source)
	gi.Free(c_exceptions)
	result = ret.String().Take()
	return
}

// g_strfreev
//
// [ str_array ] trans: nothing
//
func Strfreev(str_array string) {
	iv, err := _I.Get(1136, "strfreev", "", 662, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str_array := gi.CString(str_array)
	arg_str_array := gi.NewStringArgument(c_str_array)
	args := []gi.Argument{arg_str_array}
	iv.Call(args, nil, nil)
	gi.Free(c_str_array)
}

// g_string_new
//
// [ init ] trans: nothing
//
// [ result ] trans: everything
//
func StringNew(init string) (result String) {
	iv, err := _I.Get(1137, "string_new", "", 663, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_init := gi.CString(init)
	arg_init := gi.NewStringArgument(c_init)
	args := []gi.Argument{arg_init}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_init)
	result.P = ret.Pointer()
	return
}

// g_string_new_len
//
// [ init ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func StringNewLen(init string, len1 int64) (result String) {
	iv, err := _I.Get(1138, "string_new_len", "", 664, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_init := gi.CString(init)
	arg_init := gi.NewStringArgument(c_init)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_init, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_init)
	result.P = ret.Pointer()
	return
}

// g_string_sized_new
//
// [ dfl_size ] trans: nothing
//
// [ result ] trans: everything
//
func StringSizedNew(dfl_size uint64) (result String) {
	iv, err := _I.Get(1139, "string_sized_new", "", 665, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dfl_size := gi.NewUint64Argument(dfl_size)
	args := []gi.Argument{arg_dfl_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_strip_context
//
// [ msgid ] trans: nothing
//
// [ msgval ] trans: nothing
//
// [ result ] trans: nothing
//
func StripContext(msgid string, msgval string) (result string) {
	iv, err := _I.Get(1140, "strip_context", "", 666, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_msgid := gi.CString(msgid)
	c_msgval := gi.CString(msgval)
	arg_msgid := gi.NewStringArgument(c_msgid)
	arg_msgval := gi.NewStringArgument(c_msgval)
	args := []gi.Argument{arg_msgid, arg_msgval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_msgid)
	gi.Free(c_msgval)
	result = ret.String().Copy()
	return
}

// g_strjoinv
//
// [ separator ] trans: nothing
//
// [ str_array ] trans: nothing
//
// [ result ] trans: everything
//
func Strjoinv(separator string, str_array string) (result string) {
	iv, err := _I.Get(1141, "strjoinv", "", 667, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_separator := gi.CString(separator)
	c_str_array := gi.CString(str_array)
	arg_separator := gi.NewStringArgument(c_separator)
	arg_str_array := gi.NewStringArgument(c_str_array)
	args := []gi.Argument{arg_separator, arg_str_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_separator)
	gi.Free(c_str_array)
	result = ret.String().Take()
	return
}

// g_strlcat
//
// [ dest ] trans: nothing
//
// [ src ] trans: nothing
//
// [ dest_size ] trans: nothing
//
// [ result ] trans: nothing
//
func Strlcat(dest string, src string, dest_size uint64) (result uint64) {
	iv, err := _I.Get(1142, "strlcat", "", 668, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	arg_dest_size := gi.NewUint64Argument(dest_size)
	args := []gi.Argument{arg_dest, arg_src, arg_dest_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.Uint64()
	return
}

// g_strlcpy
//
// [ dest ] trans: nothing
//
// [ src ] trans: nothing
//
// [ dest_size ] trans: nothing
//
// [ result ] trans: nothing
//
func Strlcpy(dest string, src string, dest_size uint64) (result uint64) {
	iv, err := _I.Get(1143, "strlcpy", "", 669, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	arg_dest_size := gi.NewUint64Argument(dest_size)
	args := []gi.Argument{arg_dest, arg_src, arg_dest_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.Uint64()
	return
}

// Deprecated
//
// g_strncasecmp
//
// [ s1 ] trans: nothing
//
// [ s2 ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: nothing
//
func Strncasecmp(s1 string, s2 string, n uint32) (result int32) {
	iv, err := _I.Get(1144, "strncasecmp", "", 670, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	arg_n := gi.NewUint32Argument(n)
	args := []gi.Argument{arg_s1, arg_s2, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_strndup
//
// [ str ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func Strndup(str string, n uint64) (result string) {
	iv, err := _I.Get(1145, "strndup", "", 671, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_str, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_strnfill
//
// [ length ] trans: nothing
//
// [ fill_char ] trans: nothing
//
// [ result ] trans: everything
//
func Strnfill(length uint64, fill_char int8) (result string) {
	iv, err := _I.Get(1146, "strnfill", "", 672, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_length := gi.NewUint64Argument(length)
	arg_fill_char := gi.NewInt8Argument(fill_char)
	args := []gi.Argument{arg_length, arg_fill_char}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_strreverse
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func Strreverse(string string) (result string) {
	iv, err := _I.Get(1147, "strreverse", "", 673, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strrstr
//
// [ haystack ] trans: nothing
//
// [ needle ] trans: nothing
//
// [ result ] trans: everything
//
func Strrstr(haystack string, needle string) (result string) {
	iv, err := _I.Get(1148, "strrstr", "", 674, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_haystack := gi.CString(haystack)
	c_needle := gi.CString(needle)
	arg_haystack := gi.NewStringArgument(c_haystack)
	arg_needle := gi.NewStringArgument(c_needle)
	args := []gi.Argument{arg_haystack, arg_needle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_haystack)
	gi.Free(c_needle)
	result = ret.String().Take()
	return
}

// g_strrstr_len
//
// [ haystack ] trans: nothing
//
// [ haystack_len ] trans: nothing
//
// [ needle ] trans: nothing
//
// [ result ] trans: everything
//
func StrrstrLen(haystack string, haystack_len int64, needle string) (result string) {
	iv, err := _I.Get(1149, "strrstr_len", "", 675, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_haystack := gi.CString(haystack)
	c_needle := gi.CString(needle)
	arg_haystack := gi.NewStringArgument(c_haystack)
	arg_haystack_len := gi.NewInt64Argument(haystack_len)
	arg_needle := gi.NewStringArgument(c_needle)
	args := []gi.Argument{arg_haystack, arg_haystack_len, arg_needle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_haystack)
	gi.Free(c_needle)
	result = ret.String().Take()
	return
}

// g_strsignal
//
// [ signum ] trans: nothing
//
// [ result ] trans: nothing
//
func Strsignal(signum int32) (result string) {
	iv, err := _I.Get(1150, "strsignal", "", 676, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signum := gi.NewInt32Argument(signum)
	args := []gi.Argument{arg_signum}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_strstr_len
//
// [ haystack ] trans: nothing
//
// [ haystack_len ] trans: nothing
//
// [ needle ] trans: nothing
//
// [ result ] trans: everything
//
func StrstrLen(haystack string, haystack_len int64, needle string) (result string) {
	iv, err := _I.Get(1151, "strstr_len", "", 677, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_haystack := gi.CString(haystack)
	c_needle := gi.CString(needle)
	arg_haystack := gi.NewStringArgument(c_haystack)
	arg_haystack_len := gi.NewInt64Argument(haystack_len)
	arg_needle := gi.NewStringArgument(c_needle)
	args := []gi.Argument{arg_haystack, arg_haystack_len, arg_needle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_haystack)
	gi.Free(c_needle)
	result = ret.String().Take()
	return
}

// g_strtod
//
// [ nptr ] trans: nothing
//
// [ endptr ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Strtod(nptr string) (result float64, endptr string) {
	iv, err := _I.Get(1152, "strtod", "", 678, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_nptr, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Copy()
	result = ret.Double()
	return
}

// Deprecated
//
// g_strup
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func Strup(string string) (result string) {
	iv, err := _I.Get(1153, "strup", "", 679, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strv_contains
//
// [ strv ] trans: nothing
//
// [ str ] trans: nothing
//
// [ result ] trans: nothing
//
func StrvContains(strv string, str string) (result bool) {
	iv, err := _I.Get(1154, "strv_contains", "", 680, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_strv := gi.CString(strv)
	c_str := gi.CString(str)
	arg_strv := gi.NewStringArgument(c_strv)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_strv, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_strv)
	gi.Free(c_str)
	result = ret.Bool()
	return
}

// g_strv_get_type
//
// [ result ] trans: nothing
//
func StrvGetType() (result gi.GType) {
	iv, err := _I.Get(1155, "strv_get_type", "", 681, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_strv_length
//
// [ str_array ] trans: nothing
//
// [ result ] trans: nothing
//
func StrvLength(str_array string) (result uint32) {
	iv, err := _I.Get(1156, "strv_length", "", 682, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str_array := gi.CString(str_array)
	arg_str_array := gi.NewStringArgument(c_str_array)
	args := []gi.Argument{arg_str_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str_array)
	result = ret.Uint32()
	return
}

// g_test_add_data_func
//
// [ testpath ] trans: nothing
//
// [ test_data ] trans: nothing
//
// [ test_func ] trans: nothing
//
func TestAddDataFunc(testpath string, test_data unsafe.Pointer, test_func int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1157, "test_add_data_func", "", 683, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_testpath := gi.CString(testpath)
	arg_testpath := gi.NewStringArgument(c_testpath)
	arg_test_data := gi.NewPointerArgument(test_data)
	arg_test_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTestDataFunc()))
	args := []gi.Argument{arg_testpath, arg_test_data, arg_test_func}
	iv.Call(args, nil, nil)
	gi.Free(c_testpath)
}

// g_test_add_data_func_full
//
// [ testpath ] trans: nothing
//
// [ test_data ] trans: nothing
//
// [ test_func ] trans: nothing
//
// [ data_free_func ] trans: nothing
//
func TestAddDataFuncFull(testpath string, test_data unsafe.Pointer, test_func int /*TODO_TYPE CALLBACK*/, data_free_func int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1158, "test_add_data_func_full", "", 684, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_testpath := gi.CString(testpath)
	arg_testpath := gi.NewStringArgument(c_testpath)
	arg_test_data := gi.NewPointerArgument(test_data)
	arg_test_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTestDataFunc()))
	arg_data_free_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_testpath, arg_test_data, arg_test_func, arg_data_free_func}
	iv.Call(args, nil, nil)
	gi.Free(c_testpath)
}

// g_test_add_func
//
// [ testpath ] trans: nothing
//
// [ test_func ] trans: nothing
//
func TestAddFunc(testpath string, test_func int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1159, "test_add_func", "", 685, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_testpath := gi.CString(testpath)
	arg_testpath := gi.NewStringArgument(c_testpath)
	arg_test_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTestFunc()))
	args := []gi.Argument{arg_testpath, arg_test_func}
	iv.Call(args, nil, nil)
	gi.Free(c_testpath)
}

// g_test_assert_expected_messages_internal
//
// [ domain ] trans: nothing
//
// [ file ] trans: nothing
//
// [ line ] trans: nothing
//
// [ func1 ] trans: nothing
//
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func1 string) {
	iv, err := _I.Get(1160, "test_assert_expected_messages_internal", "", 686, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
}

// g_test_bug
//
// [ bug_uri_snippet ] trans: nothing
//
func TestBug(bug_uri_snippet string) {
	iv, err := _I.Get(1161, "test_bug", "", 687, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_bug_uri_snippet := gi.CString(bug_uri_snippet)
	arg_bug_uri_snippet := gi.NewStringArgument(c_bug_uri_snippet)
	args := []gi.Argument{arg_bug_uri_snippet}
	iv.Call(args, nil, nil)
	gi.Free(c_bug_uri_snippet)
}

// g_test_bug_base
//
// [ uri_pattern ] trans: nothing
//
func TestBugBase(uri_pattern string) {
	iv, err := _I.Get(1162, "test_bug_base", "", 688, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri_pattern := gi.CString(uri_pattern)
	arg_uri_pattern := gi.NewStringArgument(c_uri_pattern)
	args := []gi.Argument{arg_uri_pattern}
	iv.Call(args, nil, nil)
	gi.Free(c_uri_pattern)
}

// g_test_expect_message
//
// [ log_domain ] trans: nothing
//
// [ log_level ] trans: nothing
//
// [ pattern ] trans: nothing
//
func TestExpectMessage(log_domain string, log_level LogLevelFlags, pattern string) {
	iv, err := _I.Get(1163, "test_expect_message", "", 689, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	c_pattern := gi.CString(pattern)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_pattern := gi.NewStringArgument(c_pattern)
	args := []gi.Argument{arg_log_domain, arg_log_level, arg_pattern}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
	gi.Free(c_pattern)
}

// g_test_fail
//
func TestFail() {
	iv, err := _I.Get(1164, "test_fail", "", 690, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_test_failed
//
// [ result ] trans: nothing
//
func TestFailed() (result bool) {
	iv, err := _I.Get(1165, "test_failed", "", 691, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_get_dir
//
// [ file_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TestGetDir(file_type TestFileTypeEnum) (result string) {
	iv, err := _I.Get(1166, "test_get_dir", "", 692, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_file_type := gi.NewIntArgument(int(file_type))
	args := []gi.Argument{arg_file_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_test_incomplete
//
// [ msg ] trans: nothing
//
func TestIncomplete(msg string) {
	iv, err := _I.Get(1167, "test_incomplete", "", 693, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_msg := gi.CString(msg)
	arg_msg := gi.NewStringArgument(c_msg)
	args := []gi.Argument{arg_msg}
	iv.Call(args, nil, nil)
	gi.Free(c_msg)
}

// g_test_log_type_name
//
// [ log_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TestLogTypeName(log_type TestLogTypeEnum) (result string) {
	iv, err := _I.Get(1168, "test_log_type_name", "", 694, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_type := gi.NewIntArgument(int(log_type))
	args := []gi.Argument{arg_log_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_test_queue_destroy
//
// [ destroy_func ] trans: nothing
//
// [ destroy_data ] trans: nothing
//
func TestQueueDestroy(destroy_func int /*TODO_TYPE CALLBACK*/, destroy_data unsafe.Pointer) {
	iv, err := _I.Get(1169, "test_queue_destroy", "", 695, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_destroy_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	arg_destroy_data := gi.NewPointerArgument(destroy_data)
	args := []gi.Argument{arg_destroy_func, arg_destroy_data}
	iv.Call(args, nil, nil)
}

// g_test_queue_free
//
// [ gfree_pointer ] trans: nothing
//
func TestQueueFree(gfree_pointer unsafe.Pointer) {
	iv, err := _I.Get(1170, "test_queue_free", "", 696, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_gfree_pointer := gi.NewPointerArgument(gfree_pointer)
	args := []gi.Argument{arg_gfree_pointer}
	iv.Call(args, nil, nil)
}

// g_test_rand_double
//
// [ result ] trans: nothing
//
func TestRandDouble() (result float64) {
	iv, err := _I.Get(1171, "test_rand_double", "", 697, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_test_rand_double_range
//
// [ range_start ] trans: nothing
//
// [ range_end ] trans: nothing
//
// [ result ] trans: nothing
//
func TestRandDoubleRange(range_start float64, range_end float64) (result float64) {
	iv, err := _I.Get(1172, "test_rand_double_range", "", 698, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_range_start := gi.NewDoubleArgument(range_start)
	arg_range_end := gi.NewDoubleArgument(range_end)
	args := []gi.Argument{arg_range_start, arg_range_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_test_rand_int
//
// [ result ] trans: nothing
//
func TestRandInt() (result int32) {
	iv, err := _I.Get(1173, "test_rand_int", "", 699, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_rand_int_range
//
// [ begin ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: nothing
//
func TestRandIntRange(begin int32, end int32) (result int32) {
	iv, err := _I.Get(1174, "test_rand_int_range", "", 700, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewInt32Argument(begin)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_run
//
// [ result ] trans: nothing
//
func TestRun() (result int32) {
	iv, err := _I.Get(1175, "test_run", "", 701, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_run_suite
//
// [ suite ] trans: nothing
//
// [ result ] trans: nothing
//
func TestRunSuite(suite TestSuite) (result int32) {
	iv, err := _I.Get(1176, "test_run_suite", "", 702, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_suite := gi.NewPointerArgument(suite.P)
	args := []gi.Argument{arg_suite}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_set_nonfatal_assertions
//
func TestSetNonfatalAssertions() {
	iv, err := _I.Get(1177, "test_set_nonfatal_assertions", "", 703, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_test_skip
//
// [ msg ] trans: nothing
//
func TestSkip(msg string) {
	iv, err := _I.Get(1178, "test_skip", "", 704, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_msg := gi.CString(msg)
	arg_msg := gi.NewStringArgument(c_msg)
	args := []gi.Argument{arg_msg}
	iv.Call(args, nil, nil)
	gi.Free(c_msg)
}

// g_test_subprocess
//
// [ result ] trans: nothing
//
func TestSubprocess() (result bool) {
	iv, err := _I.Get(1179, "test_subprocess", "", 705, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_timer_elapsed
//
// [ result ] trans: nothing
//
func TestTimerElapsed() (result float64) {
	iv, err := _I.Get(1180, "test_timer_elapsed", "", 706, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_test_timer_last
//
// [ result ] trans: nothing
//
func TestTimerLast() (result float64) {
	iv, err := _I.Get(1181, "test_timer_last", "", 707, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_test_timer_start
//
func TestTimerStart() {
	iv, err := _I.Get(1182, "test_timer_start", "", 708, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_test_trap_assertions
//
// [ domain ] trans: nothing
//
// [ file ] trans: nothing
//
// [ line ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ assertion_flags ] trans: nothing
//
// [ pattern ] trans: nothing
//
func TestTrapAssertions(domain string, file string, line int32, func1 string, assertion_flags uint64, pattern string) {
	iv, err := _I.Get(1183, "test_trap_assertions", "", 709, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_pattern := gi.CString(pattern)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_assertion_flags := gi.NewUint64Argument(assertion_flags)
	arg_pattern := gi.NewStringArgument(c_pattern)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_assertion_flags, arg_pattern}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_pattern)
}

// Deprecated
//
// g_test_trap_fork
//
// [ usec_timeout ] trans: nothing
//
// [ test_trap_flags ] trans: nothing
//
// [ result ] trans: nothing
//
func TestTrapFork(usec_timeout uint64, test_trap_flags TestTrapFlags) (result bool) {
	iv, err := _I.Get(1184, "test_trap_fork", "", 710, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_usec_timeout := gi.NewUint64Argument(usec_timeout)
	arg_test_trap_flags := gi.NewIntArgument(int(test_trap_flags))
	args := []gi.Argument{arg_usec_timeout, arg_test_trap_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_trap_has_passed
//
// [ result ] trans: nothing
//
func TestTrapHasPassed() (result bool) {
	iv, err := _I.Get(1185, "test_trap_has_passed", "", 711, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_trap_reached_timeout
//
// [ result ] trans: nothing
//
func TestTrapReachedTimeout() (result bool) {
	iv, err := _I.Get(1186, "test_trap_reached_timeout", "", 712, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_trap_subprocess
//
// [ test_path ] trans: nothing
//
// [ usec_timeout ] trans: nothing
//
// [ test_flags ] trans: nothing
//
func TestTrapSubprocess(test_path string, usec_timeout uint64, test_flags TestSubprocessFlags) {
	iv, err := _I.Get(1187, "test_trap_subprocess", "", 713, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_test_path := gi.CString(test_path)
	arg_test_path := gi.NewStringArgument(c_test_path)
	arg_usec_timeout := gi.NewUint64Argument(usec_timeout)
	arg_test_flags := gi.NewIntArgument(int(test_flags))
	args := []gi.Argument{arg_test_path, arg_usec_timeout, arg_test_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_test_path)
}

// g_thread_error_quark
//
// [ result ] trans: nothing
//
func ThreadErrorQuark() (result uint32) {
	iv, err := _I.Get(1188, "thread_error_quark", "", 714, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_exit
//
// [ retval ] trans: nothing
//
func ThreadExit(retval unsafe.Pointer) {
	iv, err := _I.Get(1189, "thread_exit", "", 715, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_retval := gi.NewPointerArgument(retval)
	args := []gi.Argument{arg_retval}
	iv.Call(args, nil, nil)
}

// g_thread_pool_get_max_idle_time
//
// [ result ] trans: nothing
//
func ThreadPoolGetMaxIdleTime() (result uint32) {
	iv, err := _I.Get(1190, "thread_pool_get_max_idle_time", "", 716, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_pool_get_max_unused_threads
//
// [ result ] trans: nothing
//
func ThreadPoolGetMaxUnusedThreads() (result int32) {
	iv, err := _I.Get(1191, "thread_pool_get_max_unused_threads", "", 717, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_thread_pool_get_num_unused_threads
//
// [ result ] trans: nothing
//
func ThreadPoolGetNumUnusedThreads() (result uint32) {
	iv, err := _I.Get(1192, "thread_pool_get_num_unused_threads", "", 718, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_pool_set_max_idle_time
//
// [ interval ] trans: nothing
//
func ThreadPoolSetMaxIdleTime(interval uint32) {
	iv, err := _I.Get(1193, "thread_pool_set_max_idle_time", "", 719, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	iv.Call(args, nil, nil)
}

// g_thread_pool_set_max_unused_threads
//
// [ max_threads ] trans: nothing
//
func ThreadPoolSetMaxUnusedThreads(max_threads int32) {
	iv, err := _I.Get(1194, "thread_pool_set_max_unused_threads", "", 720, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_max_threads := gi.NewInt32Argument(max_threads)
	args := []gi.Argument{arg_max_threads}
	iv.Call(args, nil, nil)
}

// g_thread_pool_stop_unused_threads
//
func ThreadPoolStopUnusedThreads() {
	iv, err := _I.Get(1195, "thread_pool_stop_unused_threads", "", 721, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_thread_self
//
// [ result ] trans: everything
//
func ThreadSelf() (result Thread) {
	iv, err := _I.Get(1196, "thread_self", "", 722, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_thread_yield
//
func ThreadYield() {
	iv, err := _I.Get(1197, "thread_yield", "", 723, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_time_val_from_iso8601
//
// [ iso_date ] trans: nothing
//
// [ time_ ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func TimeValFromIso8601(iso_date string, time_ TimeVal) (result bool) {
	iv, err := _I.Get(1198, "time_val_from_iso8601", "", 724, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_iso_date := gi.CString(iso_date)
	arg_iso_date := gi.NewStringArgument(c_iso_date)
	arg_time_ := gi.NewPointerArgument(time_.P)
	args := []gi.Argument{arg_iso_date, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_iso_date)
	result = ret.Bool()
	return
}

// g_timeout_add_full
//
// [ priority ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func TimeoutAdd(priority int32, interval uint32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(1199, "timeout_add", "", 725, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_interval := gi.NewUint32Argument(interval)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_interval, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_timeout_add_seconds_full
//
// [ priority ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func TimeoutAddSeconds(priority int32, interval uint32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(1200, "timeout_add_seconds", "", 726, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_interval := gi.NewUint32Argument(interval)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_interval, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_timeout_source_new
//
// [ interval ] trans: nothing
//
// [ result ] trans: everything
//
func TimeoutSourceNew(interval uint32) (result Source) {
	iv, err := _I.Get(1201, "timeout_source_new", "", 727, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_timeout_source_new_seconds
//
// [ interval ] trans: nothing
//
// [ result ] trans: everything
//
func TimeoutSourceNewSeconds(interval uint32) (result Source) {
	iv, err := _I.Get(1202, "timeout_source_new_seconds", "", 728, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// g_trash_stack_height
//
// [ stack_p ] trans: nothing
//
// [ result ] trans: nothing
//
func TrashStackHeight(stack_p TrashStack) (result uint32) {
	iv, err := _I.Get(1203, "trash_stack_height", "", 729, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	args := []gi.Argument{arg_stack_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// g_trash_stack_peek
//
// [ stack_p ] trans: nothing
//
// [ result ] trans: nothing
//
func TrashStackPeek(stack_p TrashStack) (result unsafe.Pointer) {
	iv, err := _I.Get(1204, "trash_stack_peek", "", 730, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	args := []gi.Argument{arg_stack_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Deprecated
//
// g_trash_stack_pop
//
// [ stack_p ] trans: nothing
//
// [ result ] trans: nothing
//
func TrashStackPop(stack_p TrashStack) (result unsafe.Pointer) {
	iv, err := _I.Get(1205, "trash_stack_pop", "", 731, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	args := []gi.Argument{arg_stack_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Deprecated
//
// g_trash_stack_push
//
// [ stack_p ] trans: nothing
//
// [ data_p ] trans: nothing
//
func TrashStackPush(stack_p TrashStack, data_p unsafe.Pointer) {
	iv, err := _I.Get(1206, "trash_stack_push", "", 732, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	arg_data_p := gi.NewPointerArgument(data_p)
	args := []gi.Argument{arg_stack_p, arg_data_p}
	iv.Call(args, nil, nil)
}

// g_try_malloc
//
// [ n_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func TryMalloc(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1207, "try_malloc", "", 733, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_malloc0
//
// [ n_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func TryMalloc0(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1208, "try_malloc0", "", 734, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_malloc0_n
//
// [ n_blocks ] trans: nothing
//
// [ n_block_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func TryMalloc0N(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1209, "try_malloc0_n", "", 735, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_malloc_n
//
// [ n_blocks ] trans: nothing
//
// [ n_block_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func TryMallocN(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1210, "try_malloc_n", "", 736, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_realloc
//
// [ mem ] trans: nothing
//
// [ n_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func TryRealloc(mem unsafe.Pointer, n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1211, "try_realloc", "", 737, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_mem, arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_realloc_n
//
// [ mem ] trans: nothing
//
// [ n_blocks ] trans: nothing
//
// [ n_block_bytes ] trans: nothing
//
// [ result ] trans: nothing
//
func TryReallocN(mem unsafe.Pointer, n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1212, "try_realloc_n", "", 738, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_mem, arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_ucs4_to_utf16
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_read ] trans: nothing, dir: out
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Ucs4ToUtf16(str rune, len1 int64) (result uint16, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1213, "ucs4_to_utf16", "", 739, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint32Argument(uint32(str))
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.Uint16()
	return
}

// g_ucs4_to_utf8
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_read ] trans: nothing, dir: out
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: everything
//
func Ucs4ToUtf8(str rune, len1 int64) (result string, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1214, "ucs4_to_utf8", "", 740, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint32Argument(uint32(str))
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.String().Take()
	return
}

// g_unichar_break_type
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharBreakType(c rune) (result UnicodeBreakTypeEnum) {
	iv, err := _I.Get(1215, "unichar_break_type", "", 741, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeBreakTypeEnum(ret.Int())
	return
}

// g_unichar_combining_class
//
// [ uc ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharCombiningClass(uc rune) (result int32) {
	iv, err := _I.Get(1216, "unichar_combining_class", "", 742, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_uc := gi.NewUint32Argument(uint32(uc))
	args := []gi.Argument{arg_uc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_unichar_compose
//
// [ a ] trans: nothing
//
// [ b ] trans: nothing
//
// [ ch ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharCompose(a rune, b rune, ch rune) (result bool) {
	iv, err := _I.Get(1217, "unichar_compose", "", 743, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewUint32Argument(uint32(a))
	arg_b := gi.NewUint32Argument(uint32(b))
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_a, arg_b, arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_decompose
//
// [ ch ] trans: nothing
//
// [ a ] trans: nothing
//
// [ b ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharDecompose(ch rune, a rune, b rune) (result bool) {
	iv, err := _I.Get(1218, "unichar_decompose", "", 744, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_a := gi.NewUint32Argument(uint32(a))
	arg_b := gi.NewUint32Argument(uint32(b))
	args := []gi.Argument{arg_ch, arg_a, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_digit_value
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharDigitValue(c rune) (result int32) {
	iv, err := _I.Get(1219, "unichar_digit_value", "", 745, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_unichar_fully_decompose
//
// [ ch ] trans: nothing
//
// [ compat ] trans: nothing
//
// [ result ] trans: nothing
//
// [ result_len ] trans: nothing
//
// [ result1 ] trans: nothing
//
func UnicharFullyDecompose(ch rune, compat bool, result rune, result_len uint64) (result1 uint64) {
	iv, err := _I.Get(1220, "unichar_fully_decompose", "", 746, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_compat := gi.NewBoolArgument(compat)
	arg_result := gi.NewUint32Argument(uint32(result))
	arg_result_len := gi.NewUint64Argument(result_len)
	args := []gi.Argument{arg_ch, arg_compat, arg_result, arg_result_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result1 = ret.Uint64()
	return
}

// g_unichar_get_mirror_char
//
// [ ch ] trans: nothing
//
// [ mirrored_ch ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharGetMirrorChar(ch rune, mirrored_ch rune) (result bool) {
	iv, err := _I.Get(1221, "unichar_get_mirror_char", "", 747, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_mirrored_ch := gi.NewUint32Argument(uint32(mirrored_ch))
	args := []gi.Argument{arg_ch, arg_mirrored_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_get_script
//
// [ ch ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharGetScript(ch rune) (result UnicodeScriptEnum) {
	iv, err := _I.Get(1222, "unichar_get_script", "", 748, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeScriptEnum(ret.Int())
	return
}

// g_unichar_isalnum
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsalnum(c rune) (result bool) {
	iv, err := _I.Get(1223, "unichar_isalnum", "", 749, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isalpha
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsalpha(c rune) (result bool) {
	iv, err := _I.Get(1224, "unichar_isalpha", "", 750, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iscntrl
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIscntrl(c rune) (result bool) {
	iv, err := _I.Get(1225, "unichar_iscntrl", "", 751, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isdefined
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsdefined(c rune) (result bool) {
	iv, err := _I.Get(1226, "unichar_isdefined", "", 752, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isdigit
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsdigit(c rune) (result bool) {
	iv, err := _I.Get(1227, "unichar_isdigit", "", 753, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isgraph
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsgraph(c rune) (result bool) {
	iv, err := _I.Get(1228, "unichar_isgraph", "", 754, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_islower
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIslower(c rune) (result bool) {
	iv, err := _I.Get(1229, "unichar_islower", "", 755, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_ismark
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsmark(c rune) (result bool) {
	iv, err := _I.Get(1230, "unichar_ismark", "", 756, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isprint
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsprint(c rune) (result bool) {
	iv, err := _I.Get(1231, "unichar_isprint", "", 757, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_ispunct
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIspunct(c rune) (result bool) {
	iv, err := _I.Get(1232, "unichar_ispunct", "", 758, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isspace
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsspace(c rune) (result bool) {
	iv, err := _I.Get(1233, "unichar_isspace", "", 759, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_istitle
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIstitle(c rune) (result bool) {
	iv, err := _I.Get(1234, "unichar_istitle", "", 760, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isupper
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsupper(c rune) (result bool) {
	iv, err := _I.Get(1235, "unichar_isupper", "", 761, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iswide
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIswide(c rune) (result bool) {
	iv, err := _I.Get(1236, "unichar_iswide", "", 762, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iswide_cjk
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIswideCjk(c rune) (result bool) {
	iv, err := _I.Get(1237, "unichar_iswide_cjk", "", 763, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isxdigit
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIsxdigit(c rune) (result bool) {
	iv, err := _I.Get(1238, "unichar_isxdigit", "", 764, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iszerowidth
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharIszerowidth(c rune) (result bool) {
	iv, err := _I.Get(1239, "unichar_iszerowidth", "", 765, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_to_utf8
//
// [ c ] trans: nothing
//
// [ outbuf ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func UnicharToUtf8(c rune) (result int32, outbuf string) {
	iv, err := _I.Get(1240, "unichar_to_utf8", "", 766, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_c := gi.NewUint32Argument(uint32(c))
	arg_outbuf := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_c, arg_outbuf}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	outbuf = outArgs[0].String().Copy()
	result = ret.Int32()
	return
}

// g_unichar_tolower
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharTolower(c rune) (result rune) {
	iv, err := _I.Get(1241, "unichar_tolower", "", 767, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unichar_totitle
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharTotitle(c rune) (result rune) {
	iv, err := _I.Get(1242, "unichar_totitle", "", 768, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unichar_toupper
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharToupper(c rune) (result rune) {
	iv, err := _I.Get(1243, "unichar_toupper", "", 769, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unichar_type
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharType(c rune) (result UnicodeTypeEnum) {
	iv, err := _I.Get(1244, "unichar_type", "", 770, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeTypeEnum(ret.Int())
	return
}

// g_unichar_validate
//
// [ ch ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharValidate(ch rune) (result bool) {
	iv, err := _I.Get(1245, "unichar_validate", "", 771, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_xdigit_value
//
// [ c ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicharXdigitValue(c rune) (result int32) {
	iv, err := _I.Get(1246, "unichar_xdigit_value", "", 772, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// g_unicode_canonical_decomposition
//
// [ ch ] trans: nothing
//
// [ result_len ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicodeCanonicalDecomposition(ch rune, result_len uint64) (result rune) {
	iv, err := _I.Get(1247, "unicode_canonical_decomposition", "", 773, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_result_len := gi.NewUint64Argument(result_len)
	args := []gi.Argument{arg_ch, arg_result_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unicode_canonical_ordering
//
// [ string ] trans: nothing
//
// [ len1 ] trans: nothing
//
func UnicodeCanonicalOrdering(string rune, len1 uint64) {
	iv, err := _I.Get(1248, "unicode_canonical_ordering", "", 774, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewUint32Argument(uint32(string))
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_string, arg_len1}
	iv.Call(args, nil, nil)
}

// g_unicode_script_from_iso15924
//
// [ iso15924 ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicodeScriptFromIso15924(iso15924 uint32) (result UnicodeScriptEnum) {
	iv, err := _I.Get(1249, "unicode_script_from_iso15924", "", 775, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iso15924 := gi.NewUint32Argument(iso15924)
	args := []gi.Argument{arg_iso15924}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeScriptEnum(ret.Int())
	return
}

// g_unicode_script_to_iso15924
//
// [ script ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicodeScriptToIso15924(script UnicodeScriptEnum) (result uint32) {
	iv, err := _I.Get(1250, "unicode_script_to_iso15924", "", 776, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(int(script))
	args := []gi.Argument{arg_script}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_error_quark
//
// [ result ] trans: nothing
//
func UnixErrorQuark() (result uint32) {
	iv, err := _I.Get(1251, "unix_error_quark", "", 777, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_fd_add_full
//
// [ priority ] trans: nothing
//
// [ fd ] trans: nothing
//
// [ condition ] trans: nothing
//
// [ function ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func UnixFdAddFull(priority int32, fd int32, condition IOConditionFlags, function int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(1252, "unix_fd_add_full", "", 778, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_fd := gi.NewInt32Argument(fd)
	arg_condition := gi.NewIntArgument(int(condition))
	arg_function := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myUnixFDSourceFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_fd, arg_condition, arg_function, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_fd_source_new
//
// [ fd ] trans: nothing
//
// [ condition ] trans: nothing
//
// [ result ] trans: everything
//
func UnixFdSourceNew(fd int32, condition IOConditionFlags) (result Source) {
	iv, err := _I.Get(1253, "unix_fd_source_new", "", 779, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fd := gi.NewInt32Argument(fd)
	arg_condition := gi.NewIntArgument(int(condition))
	args := []gi.Argument{arg_fd, arg_condition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_unix_open_pipe
//
// [ fds ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func UnixOpenPipe(fds int32, flags int32) (result bool, err error) {
	iv, err := _I.Get(1254, "unix_open_pipe", "", 780, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fds := gi.NewInt32Argument(fds)
	arg_flags := gi.NewInt32Argument(flags)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fds, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_unix_set_fd_nonblocking
//
// [ fd ] trans: nothing
//
// [ nonblock ] trans: nothing
//
// [ result ] trans: nothing
//
func UnixSetFdNonblocking(fd int32, nonblock bool) (result bool, err error) {
	iv, err := _I.Get(1255, "unix_set_fd_nonblocking", "", 781, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fd := gi.NewInt32Argument(fd)
	arg_nonblock := gi.NewBoolArgument(nonblock)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_nonblock, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_unix_signal_add_full
//
// [ priority ] trans: nothing
//
// [ signum ] trans: nothing
//
// [ handler ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func UnixSignalAdd(priority int32, signum int32, handler int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(1256, "unix_signal_add", "", 782, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_signum := gi.NewInt32Argument(signum)
	arg_handler := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySourceFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_signum, arg_handler, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_signal_source_new
//
// [ signum ] trans: nothing
//
// [ result ] trans: everything
//
func UnixSignalSourceNew(signum int32) (result Source) {
	iv, err := _I.Get(1257, "unix_signal_source_new", "", 783, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signum := gi.NewInt32Argument(signum)
	args := []gi.Argument{arg_signum}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_unlink
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func Unlink(filename string) (result int32) {
	iv, err := _I.Get(1258, "unlink", "", 784, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Int32()
	return
}

// g_unsetenv
//
// [ variable ] trans: nothing
//
func Unsetenv(variable string) {
	iv, err := _I.Get(1259, "unsetenv", "", 785, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_variable}
	iv.Call(args, nil, nil)
	gi.Free(c_variable)
}

// g_uri_escape_string
//
// [ unescaped ] trans: nothing
//
// [ reserved_chars_allowed ] trans: nothing
//
// [ allow_utf8 ] trans: nothing
//
// [ result ] trans: everything
//
func UriEscapeString(unescaped string, reserved_chars_allowed string, allow_utf8 bool) (result string) {
	iv, err := _I.Get(1260, "uri_escape_string", "", 786, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_unescaped := gi.CString(unescaped)
	c_reserved_chars_allowed := gi.CString(reserved_chars_allowed)
	arg_unescaped := gi.NewStringArgument(c_unescaped)
	arg_reserved_chars_allowed := gi.NewStringArgument(c_reserved_chars_allowed)
	arg_allow_utf8 := gi.NewBoolArgument(allow_utf8)
	args := []gi.Argument{arg_unescaped, arg_reserved_chars_allowed, arg_allow_utf8}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_unescaped)
	gi.Free(c_reserved_chars_allowed)
	result = ret.String().Take()
	return
}

// g_uri_list_extract_uris
//
// [ uri_list ] trans: nothing
//
// [ result ] trans: everything
//
func UriListExtractUris(uri_list string) (result gi.CStrArray) {
	iv, err := _I.Get(1261, "uri_list_extract_uris", "", 787, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri_list := gi.CString(uri_list)
	arg_uri_list := gi.NewStringArgument(c_uri_list)
	args := []gi.Argument{arg_uri_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri_list)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_uri_parse_scheme
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriParseScheme(uri string) (result string) {
	iv, err := _I.Get(1262, "uri_parse_scheme", "", 788, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_uri := gi.NewStringArgument(c_uri)
	args := []gi.Argument{arg_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri)
	result = ret.String().Take()
	return
}

// g_uri_unescape_segment
//
// [ escaped_string ] trans: nothing
//
// [ escaped_string_end ] trans: nothing
//
// [ illegal_characters ] trans: nothing
//
// [ result ] trans: everything
//
func UriUnescapeSegment(escaped_string string, escaped_string_end string, illegal_characters string) (result string) {
	iv, err := _I.Get(1263, "uri_unescape_segment", "", 789, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_escaped_string := gi.CString(escaped_string)
	c_escaped_string_end := gi.CString(escaped_string_end)
	c_illegal_characters := gi.CString(illegal_characters)
	arg_escaped_string := gi.NewStringArgument(c_escaped_string)
	arg_escaped_string_end := gi.NewStringArgument(c_escaped_string_end)
	arg_illegal_characters := gi.NewStringArgument(c_illegal_characters)
	args := []gi.Argument{arg_escaped_string, arg_escaped_string_end, arg_illegal_characters}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_escaped_string)
	gi.Free(c_escaped_string_end)
	gi.Free(c_illegal_characters)
	result = ret.String().Take()
	return
}

// g_uri_unescape_string
//
// [ escaped_string ] trans: nothing
//
// [ illegal_characters ] trans: nothing
//
// [ result ] trans: everything
//
func UriUnescapeString(escaped_string string, illegal_characters string) (result string) {
	iv, err := _I.Get(1264, "uri_unescape_string", "", 790, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_escaped_string := gi.CString(escaped_string)
	c_illegal_characters := gi.CString(illegal_characters)
	arg_escaped_string := gi.NewStringArgument(c_escaped_string)
	arg_illegal_characters := gi.NewStringArgument(c_illegal_characters)
	args := []gi.Argument{arg_escaped_string, arg_illegal_characters}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_escaped_string)
	gi.Free(c_illegal_characters)
	result = ret.String().Take()
	return
}

// g_usleep
//
// [ microseconds ] trans: nothing
//
func Usleep(microseconds uint64) {
	iv, err := _I.Get(1265, "usleep", "", 791, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_microseconds := gi.NewUint64Argument(microseconds)
	args := []gi.Argument{arg_microseconds}
	iv.Call(args, nil, nil)
}

// g_utf16_to_ucs4
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_read ] trans: nothing, dir: out
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Utf16ToUcs4(str uint16, len1 int64) (result rune, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1266, "utf16_to_ucs4", "", 792, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint16Argument(str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = rune(ret.Uint32())
	return
}

// g_utf16_to_utf8
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_read ] trans: nothing, dir: out
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: everything
//
func Utf16ToUtf8(str uint16, len1 int64) (result string, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1267, "utf16_to_utf8", "", 793, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint16Argument(str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.String().Take()
	return
}

// g_utf8_casefold
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Casefold(str string, len1 int64) (result string) {
	iv, err := _I.Get(1268, "utf8_casefold", "", 794, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_collate
//
// [ str1 ] trans: nothing
//
// [ str2 ] trans: nothing
//
// [ result ] trans: nothing
//
func Utf8Collate(str1 string, str2 string) (result int32) {
	iv, err := _I.Get(1269, "utf8_collate", "", 795, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str1 := gi.CString(str1)
	c_str2 := gi.CString(str2)
	arg_str1 := gi.NewStringArgument(c_str1)
	arg_str2 := gi.NewStringArgument(c_str2)
	args := []gi.Argument{arg_str1, arg_str2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str1)
	gi.Free(c_str2)
	result = ret.Int32()
	return
}

// g_utf8_collate_key
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8CollateKey(str string, len1 int64) (result string) {
	iv, err := _I.Get(1270, "utf8_collate_key", "", 796, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_collate_key_for_filename
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8CollateKeyForFilename(str string, len1 int64) (result string) {
	iv, err := _I.Get(1271, "utf8_collate_key_for_filename", "", 797, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_find_next_char
//
// [ p ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8FindNextChar(p string, end string) (result string) {
	iv, err := _I.Get(1272, "utf8_find_next_char", "", 798, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	c_end := gi.CString(end)
	arg_p := gi.NewStringArgument(c_p)
	arg_end := gi.NewStringArgument(c_end)
	args := []gi.Argument{arg_p, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	gi.Free(c_end)
	result = ret.String().Take()
	return
}

// g_utf8_find_prev_char
//
// [ str ] trans: nothing
//
// [ p ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8FindPrevChar(str string, p string) (result string) {
	iv, err := _I.Get(1273, "utf8_find_prev_char", "", 799, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_p := gi.CString(p)
	arg_str := gi.NewStringArgument(c_str)
	arg_p := gi.NewStringArgument(c_p)
	args := []gi.Argument{arg_str, arg_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_get_char
//
// [ p ] trans: nothing
//
// [ result ] trans: nothing
//
func Utf8GetChar(p string) (result rune) {
	iv, err := _I.Get(1274, "utf8_get_char", "", 800, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	args := []gi.Argument{arg_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = rune(ret.Uint32())
	return
}

// g_utf8_get_char_validated
//
// [ p ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: nothing
//
func Utf8GetCharValidated(p string, max_len int64) (result rune) {
	iv, err := _I.Get(1275, "utf8_get_char_validated", "", 801, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_max_len := gi.NewInt64Argument(max_len)
	args := []gi.Argument{arg_p, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = rune(ret.Uint32())
	return
}

// g_utf8_make_valid
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8MakeValid(str string, len1 int64) (result string) {
	iv, err := _I.Get(1276, "utf8_make_valid", "", 802, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_normalize
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Normalize(str string, len1 int64, mode NormalizeModeEnum) (result string) {
	iv, err := _I.Get(1277, "utf8_normalize", "", 803, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_str, arg_len1, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_offset_to_pointer
//
// [ str ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8OffsetToPointer(str string, offset int64) (result string) {
	iv, err := _I.Get(1278, "utf8_offset_to_pointer", "", 804, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_offset := gi.NewInt64Argument(offset)
	args := []gi.Argument{arg_str, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_pointer_to_offset
//
// [ str ] trans: nothing
//
// [ pos ] trans: nothing
//
// [ result ] trans: nothing
//
func Utf8PointerToOffset(str string, pos string) (result int64) {
	iv, err := _I.Get(1279, "utf8_pointer_to_offset", "", 805, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_pos := gi.CString(pos)
	arg_str := gi.NewStringArgument(c_str)
	arg_pos := gi.NewStringArgument(c_pos)
	args := []gi.Argument{arg_str, arg_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_pos)
	result = ret.Int64()
	return
}

// g_utf8_prev_char
//
// [ p ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8PrevChar(p string) (result string) {
	iv, err := _I.Get(1280, "utf8_prev_char", "", 806, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	args := []gi.Argument{arg_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_strchr
//
// [ p ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ c ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Strchr(p string, len1 int64, c rune) (result string) {
	iv, err := _I.Get(1281, "utf8_strchr", "", 807, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_p, arg_len1, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_strdown
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Strdown(str string, len1 int64) (result string) {
	iv, err := _I.Get(1282, "utf8_strdown", "", 808, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_strlen
//
// [ p ] trans: nothing
//
// [ max ] trans: nothing
//
// [ result ] trans: nothing
//
func Utf8Strlen(p string, max int64) (result int64) {
	iv, err := _I.Get(1283, "utf8_strlen", "", 809, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_max := gi.NewInt64Argument(max)
	args := []gi.Argument{arg_p, arg_max}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.Int64()
	return
}

// g_utf8_strncpy
//
// [ dest ] trans: nothing
//
// [ src ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Strncpy(dest string, src string, n uint64) (result string) {
	iv, err := _I.Get(1284, "utf8_strncpy", "", 810, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_dest, arg_src, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.String().Take()
	return
}

// g_utf8_strrchr
//
// [ p ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ c ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Strrchr(p string, len1 int64, c rune) (result string) {
	iv, err := _I.Get(1285, "utf8_strrchr", "", 811, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_p, arg_len1, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_strreverse
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Strreverse(str string, len1 int64) (result string) {
	iv, err := _I.Get(1286, "utf8_strreverse", "", 812, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_strup
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Strup(str string, len1 int64) (result string) {
	iv, err := _I.Get(1287, "utf8_strup", "", 813, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_substring
//
// [ str ] trans: nothing
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8Substring(str string, start_pos int64, end_pos int64) (result string) {
	iv, err := _I.Get(1288, "utf8_substring", "", 814, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_start_pos := gi.NewInt64Argument(start_pos)
	arg_end_pos := gi.NewInt64Argument(end_pos)
	args := []gi.Argument{arg_str, arg_start_pos, arg_end_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_to_ucs4
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_read ] trans: nothing, dir: out
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Utf8ToUcs4(str string, len1 int64) (result rune, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1289, "utf8_to_ucs4", "", 815, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = rune(ret.Uint32())
	return
}

// g_utf8_to_ucs4_fast
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Utf8ToUcs4Fast(str string, len1 int64) (result rune, items_written int64) {
	iv, err := _I.Get(1290, "utf8_to_ucs4_fast", "", 816, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_written}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	items_written = outArgs[0].Int64()
	result = rune(ret.Uint32())
	return
}

// g_utf8_to_utf16
//
// [ str ] trans: nothing
//
// [ len1 ] trans: nothing
//
// [ items_read ] trans: nothing, dir: out
//
// [ items_written ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Utf8ToUtf16(str string, len1 int64) (result uint16, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1291, "utf8_to_utf16", "", 817, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.Uint16()
	return
}

// g_utf8_validate
//
// [ str ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ end ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func Utf8Validate(str gi.Uint8Array, max_len int64) (result bool, end string) {
	iv, err := _I.Get(1292, "utf8_validate", "", 818, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_str := gi.NewPointerArgument(str.P)
	arg_max_len := gi.NewInt64Argument(max_len)
	arg_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_str, arg_max_len, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	end = outArgs[0].String().Copy()
	result = ret.Bool()
	return
}

// g_uuid_string_is_valid
//
// [ str ] trans: nothing
//
// [ result ] trans: nothing
//
func UuidStringIsValid(str string) (result bool) {
	iv, err := _I.Get(1293, "uuid_string_is_valid", "", 819, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.Bool()
	return
}

// g_uuid_string_random
//
// [ result ] trans: everything
//
func UuidStringRandom() (result string) {
	iv, err := _I.Get(1294, "uuid_string_random", "", 820, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_variant_get_gtype
//
// [ result ] trans: nothing
//
func VariantGetGtype() (result gi.GType) {
	iv, err := _I.Get(1295, "variant_get_gtype", "", 821, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_variant_is_object_path
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantIsObjectPath(string string) (result bool) {
	iv, err := _I.Get(1296, "variant_is_object_path", "", 822, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_is_signature
//
// [ string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantIsSignature(string string) (result bool) {
	iv, err := _I.Get(1297, "variant_is_signature", "", 823, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_parse
//
// [ type1 ] trans: nothing
//
// [ text ] trans: nothing
//
// [ limit ] trans: nothing
//
// [ endptr ] trans: nothing
//
// [ result ] trans: everything
//
func VariantParse(type1 VariantType, text string, limit string, endptr string) (result Variant, err error) {
	iv, err := _I.Get(1298, "variant_parse", "", 824, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	c_limit := gi.CString(limit)
	c_endptr := gi.CString(endptr)
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewStringArgument(c_endptr)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_text, arg_limit, arg_endptr, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	gi.Free(c_limit)
	gi.Free(c_endptr)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_variant_parse_error_print_context
//
// [ error ] trans: nothing
//
// [ source_str ] trans: nothing
//
// [ result ] trans: everything
//
func VariantParseErrorPrintContext(error Error, source_str string) (result string) {
	iv, err := _I.Get(1299, "variant_parse_error_print_context", "", 825, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source_str := gi.CString(source_str)
	arg_error := gi.NewPointerArgument(error.P)
	arg_source_str := gi.NewStringArgument(c_source_str)
	args := []gi.Argument{arg_error, arg_source_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source_str)
	result = ret.String().Take()
	return
}

// g_variant_parse_error_quark
//
// [ result ] trans: nothing
//
func VariantParseErrorQuark() (result uint32) {
	iv, err := _I.Get(1300, "variant_parse_error_quark", "", 826, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// g_variant_parser_get_error_quark
//
// [ result ] trans: nothing
//
func VariantParserGetErrorQuark() (result uint32) {
	iv, err := _I.Get(1301, "variant_parser_get_error_quark", "", 827, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_variant_type_checked_
//
// [ arg0 ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantTypeChecked_(arg0 string) (result VariantType) {
	iv, err := _I.Get(1302, "variant_type_checked_", "", 828, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg0 := gi.CString(arg0)
	arg_arg0 := gi.NewStringArgument(c_arg0)
	args := []gi.Argument{arg_arg0}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_arg0)
	result.P = ret.Pointer()
	return
}

// g_variant_type_string_get_depth_
//
// [ type_string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantTypeStringGetDepth_(type_string string) (result uint64) {
	iv, err := _I.Get(1303, "variant_type_string_get_depth_", "", 829, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Uint64()
	return
}

// g_variant_type_string_is_valid
//
// [ type_string ] trans: nothing
//
// [ result ] trans: nothing
//
func VariantTypeStringIsValid(type_string string) (result bool) {
	iv, err := _I.Get(1304, "variant_type_string_is_valid", "", 830, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Bool()
	return
}

// g_variant_type_string_scan
//
// [ string ] trans: nothing
//
// [ limit ] trans: nothing
//
// [ endptr ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func VariantTypeStringScan(string string, limit string) (result bool, endptr string) {
	iv, err := _I.Get(1305, "variant_type_string_scan", "", 831, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	c_limit := gi.CString(limit)
	arg_string := gi.NewStringArgument(c_string)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_limit, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	gi.Free(c_limit)
	endptr = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// constants
const (
	ANALYZER_ANALYZING                         = 1
	ASCII_DTOSTR_BUF_SIZE                      = 39
	BIG_ENDIAN                                 = 4321
	CSET_A_2_Z                                 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CSET_DIGITS                                = "0123456789"
	CSET_a_2_z                                 = "abcdefghijklmnopqrstuvwxyz"
	DATALIST_FLAGS_MASK                        = 3
	DATE_BAD_DAY                               = 0
	DATE_BAD_JULIAN                            = 0
	DATE_BAD_YEAR                              = 0
	DIR_SEPARATOR                              = 47
	DIR_SEPARATOR_S                            = "/"
	E                                          = 2.718282
	GINT16_FORMAT                              = "hi"
	GINT16_MODIFIER                            = "h"
	GINT32_FORMAT                              = "i"
	GINT32_MODIFIER                            = ""
	GINT64_FORMAT                              = "li"
	GINT64_MODIFIER                            = "l"
	GINTPTR_FORMAT                             = "li"
	GINTPTR_MODIFIER                           = "l"
	GNUC_FUNCTION                              = ""
	GNUC_PRETTY_FUNCTION                       = ""
	GSIZE_FORMAT                               = "lu"
	GSIZE_MODIFIER                             = "l"
	GSSIZE_FORMAT                              = "li"
	GSSIZE_MODIFIER                            = "l"
	GUINT16_FORMAT                             = "hu"
	GUINT32_FORMAT                             = "u"
	GUINT64_FORMAT                             = "lu"
	GUINTPTR_FORMAT                            = "lu"
	HAVE_GINT64                                = 1
	HAVE_GNUC_VARARGS                          = 1
	HAVE_GNUC_VISIBILITY                       = 1
	HAVE_GROWING_STACK                         = 0
	HAVE_ISO_VARARGS                           = 1
	HOOK_FLAG_USER_SHIFT                       = 4
	IEEE754_DOUBLE_BIAS                        = 1023
	IEEE754_FLOAT_BIAS                         = 127
	KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX       = "Desktop Action"
	KEY_FILE_DESKTOP_GROUP                     = "Desktop Entry"
	KEY_FILE_DESKTOP_KEY_ACTIONS               = "Actions"
	KEY_FILE_DESKTOP_KEY_CATEGORIES            = "Categories"
	KEY_FILE_DESKTOP_KEY_COMMENT               = "Comment"
	KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE      = "DBusActivatable"
	KEY_FILE_DESKTOP_KEY_EXEC                  = "Exec"
	KEY_FILE_DESKTOP_KEY_FULLNAME              = "X-GNOME-FullName"
	KEY_FILE_DESKTOP_KEY_GENERIC_NAME          = "GenericName"
	KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN        = "X-GNOME-Gettext-Domain"
	KEY_FILE_DESKTOP_KEY_HIDDEN                = "Hidden"
	KEY_FILE_DESKTOP_KEY_ICON                  = "Icon"
	KEY_FILE_DESKTOP_KEY_KEYWORDS              = "Keywords"
	KEY_FILE_DESKTOP_KEY_MIME_TYPE             = "MimeType"
	KEY_FILE_DESKTOP_KEY_NAME                  = "Name"
	KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN           = "NotShowIn"
	KEY_FILE_DESKTOP_KEY_NO_DISPLAY            = "NoDisplay"
	KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN          = "OnlyShowIn"
	KEY_FILE_DESKTOP_KEY_PATH                  = "Path"
	KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY        = "StartupNotify"
	KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS      = "StartupWMClass"
	KEY_FILE_DESKTOP_KEY_TERMINAL              = "Terminal"
	KEY_FILE_DESKTOP_KEY_TRY_EXEC              = "TryExec"
	KEY_FILE_DESKTOP_KEY_TYPE                  = "Type"
	KEY_FILE_DESKTOP_KEY_URL                   = "URL"
	KEY_FILE_DESKTOP_KEY_VERSION               = "Version"
	KEY_FILE_DESKTOP_TYPE_APPLICATION          = "Application"
	KEY_FILE_DESKTOP_TYPE_DIRECTORY            = "Directory"
	KEY_FILE_DESKTOP_TYPE_LINK                 = "Link"
	LITTLE_ENDIAN                              = 1234
	LN10                                       = 2.302585
	LN2                                        = 0.693147
	LOG_2_BASE_10                              = 0.30103
	LOG_DOMAIN                                 = 0
	LOG_FATAL_MASK                             = 5
	LOG_LEVEL_USER_SHIFT                       = 8
	MAJOR_VERSION                              = 2
	MAXINT16                                   = 32767
	MAXINT32                                   = 2147483647
	MAXINT64                                   = 9223372036854775807
	MAXINT8                                    = 127
	MAXUINT16                                  = 65535
	MAXUINT32                                  = 4294967295
	MAXUINT64                                  = 18446744073709551615
	MAXUINT8                                   = 255
	MICRO_VERSION                              = 3
	MININT16                                   = -32768
	MININT32                                   = -2147483648
	MININT64                                   = -9223372036854775808
	MININT8                                    = -128
	MINOR_VERSION                              = 58
	MODULE_SUFFIX                              = "so"
	OPTION_REMAINING                           = ""
	PDP_ENDIAN                                 = 3412
	PI                                         = 3.141593
	PID_FORMAT                                 = "i"
	PI_2                                       = 1.570796
	PI_4                                       = 0.785398
	POLLFD_FORMAT                              = "%d"
	PRIORITY_DEFAULT                           = 0
	PRIORITY_DEFAULT_IDLE                      = 200
	PRIORITY_HIGH                              = -100
	PRIORITY_HIGH_IDLE                         = 100
	PRIORITY_LOW                               = 300
	SEARCHPATH_SEPARATOR                       = 58
	SEARCHPATH_SEPARATOR_S                     = ":"
	SIZEOF_LONG                                = 8
	SIZEOF_SIZE_T                              = 8
	SIZEOF_SSIZE_T                             = 8
	SIZEOF_VOID_P                              = 8
	SOURCE_CONTINUE                            = true
	SOURCE_REMOVE                              = false
	SQRT2                                      = 1.414214
	STR_DELIMITERS                             = "_-|> <."
	SYSDEF_AF_INET                             = 2
	SYSDEF_AF_INET6                            = 10
	SYSDEF_AF_UNIX                             = 1
	SYSDEF_MSG_DONTROUTE                       = 4
	SYSDEF_MSG_OOB                             = 1
	SYSDEF_MSG_PEEK                            = 2
	TIME_SPAN_DAY                              = 86400000000
	TIME_SPAN_HOUR                             = 3600000000
	TIME_SPAN_MILLISECOND                      = 1000
	TIME_SPAN_MINUTE                           = 60000000
	TIME_SPAN_SECOND                           = 1000000
	UNICHAR_MAX_DECOMPOSITION_LENGTH           = 18
	URI_RESERVED_CHARS_GENERIC_DELIMITERS      = ":/?#[]@"
	URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS = "!$&'()*+,;="
	USEC_PER_SEC                               = 1000000
	VA_COPY_AS_ARRAY                           = 1
	VERSION_MIN_REQUIRED                       = 2
	WIN32_MSG_HANDLE                           = 19981206
)
