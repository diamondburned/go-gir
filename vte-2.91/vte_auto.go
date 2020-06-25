package vte

/*
#cgo pkg-config: vte-2.91
#include <vte/vte.h>
extern void myVteSelectionFunc(VteTerminal* terminal, gint64 column, gint64 row, gpointer data);
static void* getPointer_myVteSelectionFunc() {
return (void*)(myVteSelectionFunc);
}
extern void myVteTerminalSpawnAsyncCallback(VteTerminal* terminal, gint32 pid, GError** error, gpointer user_data);
static void* getPointer_myVteTerminalSpawnAsyncCallback() {
return (void*)(myVteTerminalSpawnAsyncCallback);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gdk-3.0"
import "github.com/electricface/go-gir/gtk-3.0"
import "github.com/electricface/go-gir/pango-1.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Vte")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Vte", "2.91", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct CharAttributes
type CharAttributes struct {
	P unsafe.Pointer
}

const SizeOfStructCharAttributes = 40

func CharAttributesGetType() gi.GType {
	ret := _I.GetGType(0, "CharAttributes")
	return ret
}

// Enum CursorBlinkMode
type CursorBlinkModeEnum int

const (
	CursorBlinkModeSystem CursorBlinkModeEnum = 0
	CursorBlinkModeOn     CursorBlinkModeEnum = 1
	CursorBlinkModeOff    CursorBlinkModeEnum = 2
)

func CursorBlinkModeGetType() gi.GType {
	ret := _I.GetGType(1, "CursorBlinkMode")
	return ret
}

// Enum CursorShape
type CursorShapeEnum int

const (
	CursorShapeBlock     CursorShapeEnum = 0
	CursorShapeIbeam     CursorShapeEnum = 1
	CursorShapeUnderline CursorShapeEnum = 2
)

func CursorShapeGetType() gi.GType {
	ret := _I.GetGType(2, "CursorShape")
	return ret
}

// Enum EraseBinding
type EraseBindingEnum int

const (
	EraseBindingAuto           EraseBindingEnum = 0
	EraseBindingAsciiBackspace EraseBindingEnum = 1
	EraseBindingAsciiDelete    EraseBindingEnum = 2
	EraseBindingDeleteSequence EraseBindingEnum = 3
	EraseBindingTty            EraseBindingEnum = 4
)

func EraseBindingGetType() gi.GType {
	ret := _I.GetGType(3, "EraseBinding")
	return ret
}

// Enum Format
type FormatEnum int

const (
	FormatText FormatEnum = 1
	FormatHtml FormatEnum = 2
)

func FormatGetType() gi.GType {
	ret := _I.GetGType(4, "Format")
	return ret
}

// Object Pty
type Pty struct {
	g.InitableIfc
	g.Object
}

func WrapPty(p unsafe.Pointer) (r Pty) { r.P = p; return }

type IPty interface{ P_Pty() unsafe.Pointer }

func (v Pty) P_Pty() unsafe.Pointer      { return v.P }
func (v Pty) P_Initable() unsafe.Pointer { return v.P }
func PtyGetType() gi.GType {
	ret := _I.GetGType(5, "Pty")
	return ret
}

// vte_pty_new_foreign_sync
//
// [ fd ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewPtyForeignSync(fd int32, cancellable g.ICancellable) (result Pty, err error) {
	iv, err := _I.Get(0, "Pty", "new_foreign_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_fd := gi.NewInt32Argument(fd)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// vte_pty_new_sync
//
// [ flags ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewPtySync(flags PtyFlags, cancellable g.ICancellable) (result Pty, err error) {
	iv, err := _I.Get(1, "Pty", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_flags := gi.NewIntArgument(int(flags))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_flags, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// vte_pty_child_setup
//
func (v Pty) ChildSetup() {
	iv, err := _I.Get(2, "Pty", "child_setup")
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
// vte_pty_close
//
func (v Pty) Close() {
	iv, err := _I.Get(3, "Pty", "close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_pty_get_fd
//
// [ result ] trans: nothing
//
func (v Pty) GetFd() (result int32) {
	iv, err := _I.Get(4, "Pty", "get_fd")
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

// vte_pty_get_size
//
// [ rows ] trans: everything, dir: out
//
// [ columns ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pty) GetSize() (result bool, rows int32, columns int32, err error) {
	iv, err := _I.Get(5, "Pty", "get_size")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_rows := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_columns := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_rows, arg_columns, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	rows = outArgs[0].Int32()
	columns = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// vte_pty_set_size
//
// [ rows ] trans: nothing
//
// [ columns ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pty) SetSize(rows int32, columns int32) (result bool, err error) {
	iv, err := _I.Get(6, "Pty", "set_size")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_rows := gi.NewInt32Argument(rows)
	arg_columns := gi.NewInt32Argument(columns)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_rows, arg_columns, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// vte_pty_set_utf8
//
// [ utf8 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pty) SetUtf8(utf8 bool) (result bool, err error) {
	iv, err := _I.Get(7, "Pty", "set_utf8")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_utf8 := gi.NewBoolArgument(utf8)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_utf8, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// vte_pty_spawn_async
//
// [ working_directory ] trans: nothing
//
// [ argv ] trans: nothing
//
// [ envv ] trans: nothing
//
// [ spawn_flags ] trans: nothing
//
// [ child_setup ] trans: nothing
//
// [ child_setup_data ] trans: nothing
//
// [ child_setup_data_destroy ] trans: nothing
//
// [ timeout ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v Pty) SpawnAsync(working_directory string, argv gi.CStrArray, envv gi.CStrArray, spawn_flags g.SpawnFlags, child_setup int /*TODO_TYPE CALLBACK*/, child_setup_data unsafe.Pointer, child_setup_data_destroy int /*TODO_TYPE CALLBACK*/, timeout int32, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(8, "Pty", "spawn_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_working_directory := gi.CString(working_directory)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envv := gi.NewPointerArgument(envv.P)
	arg_spawn_flags := gi.NewIntArgument(int(spawn_flags))
	arg_child_setup := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_mySpawnChildSetupFunc()))
	arg_child_setup_data := gi.NewPointerArgument(child_setup_data)
	arg_child_setup_data_destroy := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	arg_timeout := gi.NewInt32Argument(timeout)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_working_directory, arg_argv, arg_envv, arg_spawn_flags, arg_child_setup, arg_child_setup_data, arg_child_setup_data_destroy, arg_timeout, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_working_directory)
}

// vte_pty_spawn_finish
//
// [ result ] trans: nothing
//
// [ child_pid ] trans: everything, dir: out
//
// [ result1 ] trans: nothing
//
func (v Pty) SpawnFinish(result g.IAsyncResult) (result1 bool, child_pid int32, err error) {
	iv, err := _I.Get(9, "Pty", "spawn_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if result != nil {
		tmp = result.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewPointerArgument(tmp)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_result, arg_child_pid, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	child_pid = outArgs[0].Int32()
	result1 = ret.Bool()
	return
}

// ignore GType struct PtyClass

// Enum PtyError
type PtyErrorEnum int

const (
	PtyErrorPtyHelperFailed PtyErrorEnum = 0
	PtyErrorPty98Failed     PtyErrorEnum = 1
)

func PtyErrorGetType() gi.GType {
	ret := _I.GetGType(6, "PtyError")
	return ret
}

// Flags PtyFlags
type PtyFlags int

const (
	PtyFlagsNoLastlog  PtyFlags = 1
	PtyFlagsNoUtmp     PtyFlags = 2
	PtyFlagsNoWtmp     PtyFlags = 4
	PtyFlagsNoHelper   PtyFlags = 8
	PtyFlagsNoFallback PtyFlags = 16
	PtyFlagsDefault    PtyFlags = 0
)

func PtyFlagsGetType() gi.GType {
	ret := _I.GetGType(7, "PtyFlags")
	return ret
}

// Struct Regex
type Regex struct {
	P unsafe.Pointer
}

func RegexGetType() gi.GType {
	ret := _I.GetGType(8, "Regex")
	return ret
}

// vte_regex_new_for_match
//
// [ pattern ] trans: nothing
//
// [ pattern_length ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func NewRegexForMatch(pattern string, pattern_length int64, flags uint32) (result Regex, err error) {
	iv, err := _I.Get(10, "Regex", "new_for_match")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_pattern := gi.CString(pattern)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_pattern_length := gi.NewInt64Argument(pattern_length)
	arg_flags := gi.NewUint32Argument(flags)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pattern, arg_pattern_length, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_pattern)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// vte_regex_new_for_search
//
// [ pattern ] trans: nothing
//
// [ pattern_length ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func NewRegexForSearch(pattern string, pattern_length int64, flags uint32) (result Regex, err error) {
	iv, err := _I.Get(11, "Regex", "new_for_search")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_pattern := gi.CString(pattern)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_pattern_length := gi.NewInt64Argument(pattern_length)
	arg_flags := gi.NewUint32Argument(flags)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pattern, arg_pattern_length, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_pattern)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// vte_regex_jit
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Regex) Jit(flags uint32) (result bool, err error) {
	iv, err := _I.Get(12, "Regex", "jit")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewUint32Argument(flags)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// vte_regex_ref
//
// [ result ] trans: everything
//
func (v Regex) Ref() (result Regex) {
	iv, err := _I.Get(13, "Regex", "ref")
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

// vte_regex_unref
//
// [ result ] trans: everything
//
func (v Regex) Unref() (result Regex) {
	iv, err := _I.Get(14, "Regex", "unref")
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

// Enum RegexError
type RegexErrorEnum int

const (
	RegexErrorIncompatible RegexErrorEnum = 2147483646
	RegexErrorNotSupported RegexErrorEnum = 2147483647
)

func RegexErrorGetType() gi.GType {
	ret := _I.GetGType(9, "RegexError")
	return ret
}

type SelectionFuncStruct struct {
	F_terminal Terminal
	F_column   int64
	F_row      int64
	F_data     unsafe.Pointer
}

func GetPointer_mySelectionFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myVteSelectionFunc())
}

//export myVteSelectionFunc
func myVteSelectionFunc(terminal *C.VteTerminal, column C.gint64, row C.gint64, data C.gpointer) {
	// TODO: not found user_data
}

// Object Terminal
type Terminal struct {
	gtk.ScrollableIfc
	gtk.Widget
}

func WrapTerminal(p unsafe.Pointer) (r Terminal) { r.P = p; return }

type ITerminal interface{ P_Terminal() unsafe.Pointer }

func (v Terminal) P_Terminal() unsafe.Pointer   { return v.P }
func (v Terminal) P_Scrollable() unsafe.Pointer { return v.P }
func TerminalGetType() gi.GType {
	ret := _I.GetGType(10, "Terminal")
	return ret
}

// vte_terminal_new
//
// [ result ] trans: nothing
//
func NewTerminal() (result Terminal) {
	iv, err := _I.Get(15, "Terminal", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// vte_terminal_copy_clipboard
//
func (v Terminal) CopyClipboard() {
	iv, err := _I.Get(16, "Terminal", "copy_clipboard")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_terminal_copy_clipboard_format
//
// [ format ] trans: nothing
//
func (v Terminal) CopyClipboardFormat(format FormatEnum) {
	iv, err := _I.Get(17, "Terminal", "copy_clipboard_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_v, arg_format}
	iv.Call(args, nil, nil)
}

// vte_terminal_copy_primary
//
func (v Terminal) CopyPrimary() {
	iv, err := _I.Get(18, "Terminal", "copy_primary")
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
// vte_terminal_event_check_gregex_simple
//
// [ event ] trans: nothing
//
// [ regexes ] trans: nothing
//
// [ n_regexes ] trans: everything, dir: out
//
// [ match_flags ] trans: nothing
//
// [ matches ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Terminal) EventCheckGregexSimple(event gdk.Event, regexes gi.PointerArray, match_flags g.RegexMatchFlags, matches gi.CStrArray) (result bool) {
	iv, err := _I.Get(19, "Terminal", "event_check_gregex_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	arg_regexes := gi.NewPointerArgument(regexes.P)
	arg_n_regexes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_match_flags := gi.NewIntArgument(int(match_flags))
	arg_matches := gi.NewPointerArgument(matches.P)
	args := []gi.Argument{arg_v, arg_event, arg_regexes, arg_n_regexes, arg_match_flags, arg_matches}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_regexes uint64
	_ = n_regexes
	n_regexes = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// vte_terminal_event_check_regex_simple
//
// [ event ] trans: nothing
//
// [ regexes ] trans: nothing
//
// [ n_regexes ] trans: everything, dir: out
//
// [ match_flags ] trans: nothing
//
// [ matches ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Terminal) EventCheckRegexSimple(event gdk.Event, regexes gi.PointerArray, match_flags uint32, matches gi.CStrArray) (result bool) {
	iv, err := _I.Get(20, "Terminal", "event_check_regex_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	arg_regexes := gi.NewPointerArgument(regexes.P)
	arg_n_regexes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_match_flags := gi.NewUint32Argument(match_flags)
	arg_matches := gi.NewPointerArgument(matches.P)
	args := []gi.Argument{arg_v, arg_event, arg_regexes, arg_n_regexes, arg_match_flags, arg_matches}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_regexes uint64
	_ = n_regexes
	n_regexes = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// vte_terminal_feed
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
func (v Terminal) Feed(data gi.Uint8Array, length int64) {
	iv, err := _I.Get(21, "Terminal", "feed")
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

// vte_terminal_feed_child
//
// [ text ] trans: nothing
//
// [ length ] trans: nothing
//
func (v Terminal) FeedChild(text gi.Int8Array, length int64) {
	iv, err := _I.Get(22, "Terminal", "feed_child")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewPointerArgument(text.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_v, arg_text, arg_length}
	iv.Call(args, nil, nil)
}

// vte_terminal_feed_child_binary
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
func (v Terminal) FeedChildBinary(data gi.Uint8Array, length uint64) {
	iv, err := _I.Get(23, "Terminal", "feed_child_binary")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_data, arg_length}
	iv.Call(args, nil, nil)
}

// vte_terminal_get_allow_bold
//
// [ result ] trans: nothing
//
func (v Terminal) GetAllowBold() (result bool) {
	iv, err := _I.Get(24, "Terminal", "get_allow_bold")
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

// vte_terminal_get_allow_hyperlink
//
// [ result ] trans: nothing
//
func (v Terminal) GetAllowHyperlink() (result bool) {
	iv, err := _I.Get(25, "Terminal", "get_allow_hyperlink")
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

// vte_terminal_get_audible_bell
//
// [ result ] trans: nothing
//
func (v Terminal) GetAudibleBell() (result bool) {
	iv, err := _I.Get(26, "Terminal", "get_audible_bell")
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

// vte_terminal_get_bold_is_bright
//
// [ result ] trans: nothing
//
func (v Terminal) GetBoldIsBright() (result bool) {
	iv, err := _I.Get(27, "Terminal", "get_bold_is_bright")
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

// vte_terminal_get_cell_height_scale
//
// [ result ] trans: nothing
//
func (v Terminal) GetCellHeightScale() (result float64) {
	iv, err := _I.Get(28, "Terminal", "get_cell_height_scale")
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

// vte_terminal_get_cell_width_scale
//
// [ result ] trans: nothing
//
func (v Terminal) GetCellWidthScale() (result float64) {
	iv, err := _I.Get(29, "Terminal", "get_cell_width_scale")
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

// vte_terminal_get_char_height
//
// [ result ] trans: nothing
//
func (v Terminal) GetCharHeight() (result int64) {
	iv, err := _I.Get(30, "Terminal", "get_char_height")
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

// vte_terminal_get_char_width
//
// [ result ] trans: nothing
//
func (v Terminal) GetCharWidth() (result int64) {
	iv, err := _I.Get(31, "Terminal", "get_char_width")
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

// vte_terminal_get_cjk_ambiguous_width
//
// [ result ] trans: nothing
//
func (v Terminal) GetCjkAmbiguousWidth() (result int32) {
	iv, err := _I.Get(32, "Terminal", "get_cjk_ambiguous_width")
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

// vte_terminal_get_color_background_for_draw
//
// [ color ] trans: nothing, dir: out
//
func (v Terminal) GetColorBackgroundForDraw(color gdk.RGBA) {
	iv, err := _I.Get(33, "Terminal", "get_color_background_for_draw")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	iv.Call(args, nil, nil)
}

// vte_terminal_get_column_count
//
// [ result ] trans: nothing
//
func (v Terminal) GetColumnCount() (result int64) {
	iv, err := _I.Get(34, "Terminal", "get_column_count")
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

// vte_terminal_get_current_directory_uri
//
// [ result ] trans: nothing
//
func (v Terminal) GetCurrentDirectoryUri() (result string) {
	iv, err := _I.Get(35, "Terminal", "get_current_directory_uri")
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

// vte_terminal_get_current_file_uri
//
// [ result ] trans: nothing
//
func (v Terminal) GetCurrentFileUri() (result string) {
	iv, err := _I.Get(36, "Terminal", "get_current_file_uri")
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

// vte_terminal_get_cursor_blink_mode
//
// [ result ] trans: nothing
//
func (v Terminal) GetCursorBlinkMode() (result CursorBlinkModeEnum) {
	iv, err := _I.Get(37, "Terminal", "get_cursor_blink_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CursorBlinkModeEnum(ret.Int())
	return
}

// vte_terminal_get_cursor_position
//
// [ column ] trans: everything, dir: out
//
// [ row ] trans: everything, dir: out
//
func (v Terminal) GetCursorPosition() (column int64, row int64) {
	iv, err := _I.Get(38, "Terminal", "get_cursor_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_column, arg_row}
	iv.Call(args, nil, &outArgs[0])
	column = outArgs[0].Int64()
	row = outArgs[1].Int64()
	return
}

// vte_terminal_get_cursor_shape
//
// [ result ] trans: nothing
//
func (v Terminal) GetCursorShape() (result CursorShapeEnum) {
	iv, err := _I.Get(39, "Terminal", "get_cursor_shape")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CursorShapeEnum(ret.Int())
	return
}

// Deprecated
//
// vte_terminal_get_encoding
//
// [ result ] trans: nothing
//
func (v Terminal) GetEncoding() (result string) {
	iv, err := _I.Get(40, "Terminal", "get_encoding")
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

// vte_terminal_get_font
//
// [ result ] trans: nothing
//
func (v Terminal) GetFont() (result pango.FontDescription) {
	iv, err := _I.Get(41, "Terminal", "get_font")
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

// vte_terminal_get_font_scale
//
// [ result ] trans: nothing
//
func (v Terminal) GetFontScale() (result float64) {
	iv, err := _I.Get(42, "Terminal", "get_font_scale")
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

// Deprecated
//
// vte_terminal_get_geometry_hints
//
// [ hints ] trans: nothing, dir: out
//
// [ min_rows ] trans: nothing
//
// [ min_columns ] trans: nothing
//
func (v Terminal) GetGeometryHints(hints gdk.Geometry, min_rows int32, min_columns int32) {
	iv, err := _I.Get(43, "Terminal", "get_geometry_hints")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hints := gi.NewPointerArgument(hints.P)
	arg_min_rows := gi.NewInt32Argument(min_rows)
	arg_min_columns := gi.NewInt32Argument(min_columns)
	args := []gi.Argument{arg_v, arg_hints, arg_min_rows, arg_min_columns}
	iv.Call(args, nil, nil)
}

// vte_terminal_get_has_selection
//
// [ result ] trans: nothing
//
func (v Terminal) GetHasSelection() (result bool) {
	iv, err := _I.Get(44, "Terminal", "get_has_selection")
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

// Deprecated
//
// vte_terminal_get_icon_title
//
// [ result ] trans: nothing
//
func (v Terminal) GetIconTitle() (result string) {
	iv, err := _I.Get(45, "Terminal", "get_icon_title")
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

// vte_terminal_get_input_enabled
//
// [ result ] trans: nothing
//
func (v Terminal) GetInputEnabled() (result bool) {
	iv, err := _I.Get(46, "Terminal", "get_input_enabled")
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

// vte_terminal_get_mouse_autohide
//
// [ result ] trans: nothing
//
func (v Terminal) GetMouseAutohide() (result bool) {
	iv, err := _I.Get(47, "Terminal", "get_mouse_autohide")
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

// vte_terminal_get_pty
//
// [ result ] trans: nothing
//
func (v Terminal) GetPty() (result Pty) {
	iv, err := _I.Get(48, "Terminal", "get_pty")
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

// vte_terminal_get_rewrap_on_resize
//
// [ result ] trans: nothing
//
func (v Terminal) GetRewrapOnResize() (result bool) {
	iv, err := _I.Get(49, "Terminal", "get_rewrap_on_resize")
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

// vte_terminal_get_row_count
//
// [ result ] trans: nothing
//
func (v Terminal) GetRowCount() (result int64) {
	iv, err := _I.Get(50, "Terminal", "get_row_count")
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

// vte_terminal_get_scroll_on_keystroke
//
// [ result ] trans: nothing
//
func (v Terminal) GetScrollOnKeystroke() (result bool) {
	iv, err := _I.Get(51, "Terminal", "get_scroll_on_keystroke")
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

// vte_terminal_get_scroll_on_output
//
// [ result ] trans: nothing
//
func (v Terminal) GetScrollOnOutput() (result bool) {
	iv, err := _I.Get(52, "Terminal", "get_scroll_on_output")
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

// vte_terminal_get_scrollback_lines
//
// [ result ] trans: nothing
//
func (v Terminal) GetScrollbackLines() (result int64) {
	iv, err := _I.Get(53, "Terminal", "get_scrollback_lines")
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

// vte_terminal_get_text
//
// [ is_selected ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ attributes ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Terminal) GetText(is_selected int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, attributes unsafe.Pointer /*TODO:TYPE*/) (result string) {
	iv, err := _I.Get(54, "Terminal", "get_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_is_selected := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySelectionFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_attributes := gi.NewPointerArgument(attributes /*TODO*/)
	args := []gi.Argument{arg_v, arg_is_selected, arg_user_data, arg_attributes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// vte_terminal_get_text_blink_mode
//
// [ result ] trans: nothing
//
func (v Terminal) GetTextBlinkMode() (result TextBlinkModeEnum) {
	iv, err := _I.Get(55, "Terminal", "get_text_blink_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TextBlinkModeEnum(ret.Int())
	return
}

// vte_terminal_get_text_include_trailing_spaces
//
// [ is_selected ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ attributes ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Terminal) GetTextIncludeTrailingSpaces(is_selected int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, attributes unsafe.Pointer /*TODO:TYPE*/) (result string) {
	iv, err := _I.Get(56, "Terminal", "get_text_include_trailing_spaces")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_is_selected := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySelectionFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_attributes := gi.NewPointerArgument(attributes /*TODO*/)
	args := []gi.Argument{arg_v, arg_is_selected, arg_user_data, arg_attributes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// vte_terminal_get_text_range
//
// [ start_row ] trans: nothing
//
// [ start_col ] trans: nothing
//
// [ end_row ] trans: nothing
//
// [ end_col ] trans: nothing
//
// [ is_selected ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ attributes ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Terminal) GetTextRange(start_row int64, start_col int64, end_row int64, end_col int64, is_selected int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, attributes unsafe.Pointer /*TODO:TYPE*/) (result string) {
	iv, err := _I.Get(57, "Terminal", "get_text_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start_row := gi.NewInt64Argument(start_row)
	arg_start_col := gi.NewInt64Argument(start_col)
	arg_end_row := gi.NewInt64Argument(end_row)
	arg_end_col := gi.NewInt64Argument(end_col)
	arg_is_selected := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySelectionFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_attributes := gi.NewPointerArgument(attributes /*TODO*/)
	args := []gi.Argument{arg_v, arg_start_row, arg_start_col, arg_end_row, arg_end_col, arg_is_selected, arg_user_data, arg_attributes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// vte_terminal_get_window_title
//
// [ result ] trans: nothing
//
func (v Terminal) GetWindowTitle() (result string) {
	iv, err := _I.Get(58, "Terminal", "get_window_title")
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

// vte_terminal_get_word_char_exceptions
//
// [ result ] trans: nothing
//
func (v Terminal) GetWordCharExceptions() (result string) {
	iv, err := _I.Get(59, "Terminal", "get_word_char_exceptions")
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

// vte_terminal_hyperlink_check_event
//
// [ event ] trans: nothing
//
// [ result ] trans: everything
//
func (v Terminal) HyperlinkCheckEvent(event gdk.Event) (result string) {
	iv, err := _I.Get(60, "Terminal", "hyperlink_check_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// Deprecated
//
// vte_terminal_match_add_gregex
//
// [ gregex ] trans: nothing
//
// [ gflags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Terminal) MatchAddGregex(gregex g.Regex, gflags g.RegexMatchFlags) (result int32) {
	iv, err := _I.Get(61, "Terminal", "match_add_gregex")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gregex := gi.NewPointerArgument(gregex.P)
	arg_gflags := gi.NewIntArgument(int(gflags))
	args := []gi.Argument{arg_v, arg_gregex, arg_gflags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// vte_terminal_match_add_regex
//
// [ regex ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Terminal) MatchAddRegex(regex Regex, flags uint32) (result int32) {
	iv, err := _I.Get(62, "Terminal", "match_add_regex")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_regex := gi.NewPointerArgument(regex.P)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_regex, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// vte_terminal_match_check
//
// [ column ] trans: nothing
//
// [ row ] trans: nothing
//
// [ tag ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Terminal) MatchCheck(column int64, row int64) (result string, tag int32) {
	iv, err := _I.Get(63, "Terminal", "match_check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_column := gi.NewInt64Argument(column)
	arg_row := gi.NewInt64Argument(row)
	arg_tag := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_column, arg_row, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	tag = outArgs[0].Int32()
	result = ret.String().Take()
	return
}

// vte_terminal_match_check_event
//
// [ event ] trans: nothing
//
// [ tag ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Terminal) MatchCheckEvent(event gdk.Event) (result string, tag int32) {
	iv, err := _I.Get(64, "Terminal", "match_check_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	arg_tag := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_event, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	tag = outArgs[0].Int32()
	result = ret.String().Take()
	return
}

// vte_terminal_match_remove
//
// [ tag ] trans: nothing
//
func (v Terminal) MatchRemove(tag int32) {
	iv, err := _I.Get(65, "Terminal", "match_remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewInt32Argument(tag)
	args := []gi.Argument{arg_v, arg_tag}
	iv.Call(args, nil, nil)
}

// vte_terminal_match_remove_all
//
func (v Terminal) MatchRemoveAll() {
	iv, err := _I.Get(66, "Terminal", "match_remove_all")
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
// vte_terminal_match_set_cursor
//
// [ tag ] trans: nothing
//
// [ cursor ] trans: nothing
//
func (v Terminal) MatchSetCursor(tag int32, cursor gdk.ICursor) {
	iv, err := _I.Get(67, "Terminal", "match_set_cursor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cursor != nil {
		tmp = cursor.P_Cursor()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewInt32Argument(tag)
	arg_cursor := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_tag, arg_cursor}
	iv.Call(args, nil, nil)
}

// vte_terminal_match_set_cursor_name
//
// [ tag ] trans: nothing
//
// [ cursor_name ] trans: nothing
//
func (v Terminal) MatchSetCursorName(tag int32, cursor_name string) {
	iv, err := _I.Get(68, "Terminal", "match_set_cursor_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_cursor_name := gi.CString(cursor_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewInt32Argument(tag)
	arg_cursor_name := gi.NewStringArgument(c_cursor_name)
	args := []gi.Argument{arg_v, arg_tag, arg_cursor_name}
	iv.Call(args, nil, nil)
	gi.Free(c_cursor_name)
}

// Deprecated
//
// vte_terminal_match_set_cursor_type
//
// [ tag ] trans: nothing
//
// [ cursor_type ] trans: nothing
//
func (v Terminal) MatchSetCursorType(tag int32, cursor_type gdk.CursorTypeEnum) {
	iv, err := _I.Get(69, "Terminal", "match_set_cursor_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewInt32Argument(tag)
	arg_cursor_type := gi.NewIntArgument(int(cursor_type))
	args := []gi.Argument{arg_v, arg_tag, arg_cursor_type}
	iv.Call(args, nil, nil)
}

// vte_terminal_paste_clipboard
//
func (v Terminal) PasteClipboard() {
	iv, err := _I.Get(70, "Terminal", "paste_clipboard")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_terminal_paste_primary
//
func (v Terminal) PastePrimary() {
	iv, err := _I.Get(71, "Terminal", "paste_primary")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_terminal_pty_new_sync
//
// [ flags ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func (v Terminal) PtyNewSync(flags PtyFlags, cancellable g.ICancellable) (result Pty, err error) {
	iv, err := _I.Get(72, "Terminal", "pty_new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flags, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// vte_terminal_reset
//
// [ clear_tabstops ] trans: nothing
//
// [ clear_history ] trans: nothing
//
func (v Terminal) Reset(clear_tabstops bool, clear_history bool) {
	iv, err := _I.Get(73, "Terminal", "reset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_clear_tabstops := gi.NewBoolArgument(clear_tabstops)
	arg_clear_history := gi.NewBoolArgument(clear_history)
	args := []gi.Argument{arg_v, arg_clear_tabstops, arg_clear_history}
	iv.Call(args, nil, nil)
}

// vte_terminal_search_find_next
//
// [ result ] trans: nothing
//
func (v Terminal) SearchFindNext() (result bool) {
	iv, err := _I.Get(74, "Terminal", "search_find_next")
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

// vte_terminal_search_find_previous
//
// [ result ] trans: nothing
//
func (v Terminal) SearchFindPrevious() (result bool) {
	iv, err := _I.Get(75, "Terminal", "search_find_previous")
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

// Deprecated
//
// vte_terminal_search_get_gregex
//
// [ result ] trans: nothing
//
func (v Terminal) SearchGetGregex() (result g.Regex) {
	iv, err := _I.Get(76, "Terminal", "search_get_gregex")
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

// vte_terminal_search_get_regex
//
// [ result ] trans: nothing
//
func (v Terminal) SearchGetRegex() (result Regex) {
	iv, err := _I.Get(77, "Terminal", "search_get_regex")
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

// vte_terminal_search_get_wrap_around
//
// [ result ] trans: nothing
//
func (v Terminal) SearchGetWrapAround() (result bool) {
	iv, err := _I.Get(78, "Terminal", "search_get_wrap_around")
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

// Deprecated
//
// vte_terminal_search_set_gregex
//
// [ gregex ] trans: nothing
//
// [ gflags ] trans: nothing
//
func (v Terminal) SearchSetGregex(gregex g.Regex, gflags g.RegexMatchFlags) {
	iv, err := _I.Get(79, "Terminal", "search_set_gregex")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gregex := gi.NewPointerArgument(gregex.P)
	arg_gflags := gi.NewIntArgument(int(gflags))
	args := []gi.Argument{arg_v, arg_gregex, arg_gflags}
	iv.Call(args, nil, nil)
}

// vte_terminal_search_set_regex
//
// [ regex ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v Terminal) SearchSetRegex(regex Regex, flags uint32) {
	iv, err := _I.Get(80, "Terminal", "search_set_regex")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_regex := gi.NewPointerArgument(regex.P)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_regex, arg_flags}
	iv.Call(args, nil, nil)
}

// vte_terminal_search_set_wrap_around
//
// [ wrap_around ] trans: nothing
//
func (v Terminal) SearchSetWrapAround(wrap_around bool) {
	iv, err := _I.Get(81, "Terminal", "search_set_wrap_around")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wrap_around := gi.NewBoolArgument(wrap_around)
	args := []gi.Argument{arg_v, arg_wrap_around}
	iv.Call(args, nil, nil)
}

// vte_terminal_select_all
//
func (v Terminal) SelectAll() {
	iv, err := _I.Get(82, "Terminal", "select_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_allow_bold
//
// [ allow_bold ] trans: nothing
//
func (v Terminal) SetAllowBold(allow_bold bool) {
	iv, err := _I.Get(83, "Terminal", "set_allow_bold")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_allow_bold := gi.NewBoolArgument(allow_bold)
	args := []gi.Argument{arg_v, arg_allow_bold}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_allow_hyperlink
//
// [ allow_hyperlink ] trans: nothing
//
func (v Terminal) SetAllowHyperlink(allow_hyperlink bool) {
	iv, err := _I.Get(84, "Terminal", "set_allow_hyperlink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_allow_hyperlink := gi.NewBoolArgument(allow_hyperlink)
	args := []gi.Argument{arg_v, arg_allow_hyperlink}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_audible_bell
//
// [ is_audible ] trans: nothing
//
func (v Terminal) SetAudibleBell(is_audible bool) {
	iv, err := _I.Get(85, "Terminal", "set_audible_bell")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_is_audible := gi.NewBoolArgument(is_audible)
	args := []gi.Argument{arg_v, arg_is_audible}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_backspace_binding
//
// [ binding ] trans: nothing
//
func (v Terminal) SetBackspaceBinding(binding EraseBindingEnum) {
	iv, err := _I.Get(86, "Terminal", "set_backspace_binding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_binding := gi.NewIntArgument(int(binding))
	args := []gi.Argument{arg_v, arg_binding}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_bold_is_bright
//
// [ bold_is_bright ] trans: nothing
//
func (v Terminal) SetBoldIsBright(bold_is_bright bool) {
	iv, err := _I.Get(87, "Terminal", "set_bold_is_bright")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bold_is_bright := gi.NewBoolArgument(bold_is_bright)
	args := []gi.Argument{arg_v, arg_bold_is_bright}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_cell_height_scale
//
// [ scale ] trans: nothing
//
func (v Terminal) SetCellHeightScale(scale float64) {
	iv, err := _I.Get(88, "Terminal", "set_cell_height_scale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scale := gi.NewDoubleArgument(scale)
	args := []gi.Argument{arg_v, arg_scale}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_cell_width_scale
//
// [ scale ] trans: nothing
//
func (v Terminal) SetCellWidthScale(scale float64) {
	iv, err := _I.Get(89, "Terminal", "set_cell_width_scale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scale := gi.NewDoubleArgument(scale)
	args := []gi.Argument{arg_v, arg_scale}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_cjk_ambiguous_width
//
// [ width ] trans: nothing
//
func (v Terminal) SetCjkAmbiguousWidth(width int32) {
	iv, err := _I.Get(90, "Terminal", "set_cjk_ambiguous_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewInt32Argument(width)
	args := []gi.Argument{arg_v, arg_width}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_clear_background
//
// [ setting ] trans: nothing
//
func (v Terminal) SetClearBackground(setting bool) {
	iv, err := _I.Get(91, "Terminal", "set_clear_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_setting := gi.NewBoolArgument(setting)
	args := []gi.Argument{arg_v, arg_setting}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_background
//
// [ background ] trans: nothing
//
func (v Terminal) SetColorBackground(background gdk.RGBA) {
	iv, err := _I.Get(92, "Terminal", "set_color_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_background := gi.NewPointerArgument(background.P)
	args := []gi.Argument{arg_v, arg_background}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_bold
//
// [ bold ] trans: nothing
//
func (v Terminal) SetColorBold(bold gdk.RGBA) {
	iv, err := _I.Get(93, "Terminal", "set_color_bold")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bold := gi.NewPointerArgument(bold.P)
	args := []gi.Argument{arg_v, arg_bold}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_cursor
//
// [ cursor_background ] trans: nothing
//
func (v Terminal) SetColorCursor(cursor_background gdk.RGBA) {
	iv, err := _I.Get(94, "Terminal", "set_color_cursor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cursor_background := gi.NewPointerArgument(cursor_background.P)
	args := []gi.Argument{arg_v, arg_cursor_background}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_cursor_foreground
//
// [ cursor_foreground ] trans: nothing
//
func (v Terminal) SetColorCursorForeground(cursor_foreground gdk.RGBA) {
	iv, err := _I.Get(95, "Terminal", "set_color_cursor_foreground")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cursor_foreground := gi.NewPointerArgument(cursor_foreground.P)
	args := []gi.Argument{arg_v, arg_cursor_foreground}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_foreground
//
// [ foreground ] trans: nothing
//
func (v Terminal) SetColorForeground(foreground gdk.RGBA) {
	iv, err := _I.Get(96, "Terminal", "set_color_foreground")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_foreground := gi.NewPointerArgument(foreground.P)
	args := []gi.Argument{arg_v, arg_foreground}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_highlight
//
// [ highlight_background ] trans: nothing
//
func (v Terminal) SetColorHighlight(highlight_background gdk.RGBA) {
	iv, err := _I.Get(97, "Terminal", "set_color_highlight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight_background := gi.NewPointerArgument(highlight_background.P)
	args := []gi.Argument{arg_v, arg_highlight_background}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_color_highlight_foreground
//
// [ highlight_foreground ] trans: nothing
//
func (v Terminal) SetColorHighlightForeground(highlight_foreground gdk.RGBA) {
	iv, err := _I.Get(98, "Terminal", "set_color_highlight_foreground")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight_foreground := gi.NewPointerArgument(highlight_foreground.P)
	args := []gi.Argument{arg_v, arg_highlight_foreground}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_colors
//
// [ foreground ] trans: nothing
//
// [ background ] trans: nothing
//
// [ palette ] trans: nothing
//
// [ palette_size ] trans: nothing
//
func (v Terminal) SetColors(foreground gdk.RGBA, background gdk.RGBA, palette unsafe.Pointer, palette_size uint64) {
	iv, err := _I.Get(99, "Terminal", "set_colors")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_foreground := gi.NewPointerArgument(foreground.P)
	arg_background := gi.NewPointerArgument(background.P)
	arg_palette := gi.NewPointerArgument(palette)
	arg_palette_size := gi.NewUint64Argument(palette_size)
	args := []gi.Argument{arg_v, arg_foreground, arg_background, arg_palette, arg_palette_size}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_cursor_blink_mode
//
// [ mode ] trans: nothing
//
func (v Terminal) SetCursorBlinkMode(mode CursorBlinkModeEnum) {
	iv, err := _I.Get(100, "Terminal", "set_cursor_blink_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_cursor_shape
//
// [ shape ] trans: nothing
//
func (v Terminal) SetCursorShape(shape CursorShapeEnum) {
	iv, err := _I.Get(101, "Terminal", "set_cursor_shape")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_shape := gi.NewIntArgument(int(shape))
	args := []gi.Argument{arg_v, arg_shape}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_default_colors
//
func (v Terminal) SetDefaultColors() {
	iv, err := _I.Get(102, "Terminal", "set_default_colors")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_delete_binding
//
// [ binding ] trans: nothing
//
func (v Terminal) SetDeleteBinding(binding EraseBindingEnum) {
	iv, err := _I.Get(103, "Terminal", "set_delete_binding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_binding := gi.NewIntArgument(int(binding))
	args := []gi.Argument{arg_v, arg_binding}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// vte_terminal_set_encoding
//
// [ codeset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Terminal) SetEncoding(codeset string) (result bool, err error) {
	iv, err := _I.Get(104, "Terminal", "set_encoding")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_codeset := gi.CString(codeset)
	arg_v := gi.NewPointerArgument(v.P)
	arg_codeset := gi.NewStringArgument(c_codeset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_codeset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_codeset)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// vte_terminal_set_font
//
// [ font_desc ] trans: nothing
//
func (v Terminal) SetFont(font_desc pango.FontDescription) {
	iv, err := _I.Get(105, "Terminal", "set_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_desc := gi.NewPointerArgument(font_desc.P)
	args := []gi.Argument{arg_v, arg_font_desc}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_font_scale
//
// [ scale ] trans: nothing
//
func (v Terminal) SetFontScale(scale float64) {
	iv, err := _I.Get(106, "Terminal", "set_font_scale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scale := gi.NewDoubleArgument(scale)
	args := []gi.Argument{arg_v, arg_scale}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// vte_terminal_set_geometry_hints_for_window
//
// [ window ] trans: nothing
//
func (v Terminal) SetGeometryHintsForWindow(window gtk.IWindow) {
	iv, err := _I.Get(107, "Terminal", "set_geometry_hints_for_window")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_window}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_input_enabled
//
// [ enabled ] trans: nothing
//
func (v Terminal) SetInputEnabled(enabled bool) {
	iv, err := _I.Get(108, "Terminal", "set_input_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_enabled}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_mouse_autohide
//
// [ setting ] trans: nothing
//
func (v Terminal) SetMouseAutohide(setting bool) {
	iv, err := _I.Get(109, "Terminal", "set_mouse_autohide")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_setting := gi.NewBoolArgument(setting)
	args := []gi.Argument{arg_v, arg_setting}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_pty
//
// [ pty ] trans: nothing
//
func (v Terminal) SetPty(pty IPty) {
	iv, err := _I.Get(110, "Terminal", "set_pty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pty != nil {
		tmp = pty.P_Pty()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pty := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pty}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_rewrap_on_resize
//
// [ rewrap ] trans: nothing
//
func (v Terminal) SetRewrapOnResize(rewrap bool) {
	iv, err := _I.Get(111, "Terminal", "set_rewrap_on_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rewrap := gi.NewBoolArgument(rewrap)
	args := []gi.Argument{arg_v, arg_rewrap}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_scroll_on_keystroke
//
// [ scroll ] trans: nothing
//
func (v Terminal) SetScrollOnKeystroke(scroll bool) {
	iv, err := _I.Get(112, "Terminal", "set_scroll_on_keystroke")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scroll := gi.NewBoolArgument(scroll)
	args := []gi.Argument{arg_v, arg_scroll}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_scroll_on_output
//
// [ scroll ] trans: nothing
//
func (v Terminal) SetScrollOnOutput(scroll bool) {
	iv, err := _I.Get(113, "Terminal", "set_scroll_on_output")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scroll := gi.NewBoolArgument(scroll)
	args := []gi.Argument{arg_v, arg_scroll}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_scrollback_lines
//
// [ lines ] trans: nothing
//
func (v Terminal) SetScrollbackLines(lines int64) {
	iv, err := _I.Get(114, "Terminal", "set_scrollback_lines")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_lines := gi.NewInt64Argument(lines)
	args := []gi.Argument{arg_v, arg_lines}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_size
//
// [ columns ] trans: nothing
//
// [ rows ] trans: nothing
//
func (v Terminal) SetSize(columns int64, rows int64) {
	iv, err := _I.Get(115, "Terminal", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_columns := gi.NewInt64Argument(columns)
	arg_rows := gi.NewInt64Argument(rows)
	args := []gi.Argument{arg_v, arg_columns, arg_rows}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_text_blink_mode
//
// [ text_blink_mode ] trans: nothing
//
func (v Terminal) SetTextBlinkMode(text_blink_mode TextBlinkModeEnum) {
	iv, err := _I.Get(116, "Terminal", "set_text_blink_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_text_blink_mode := gi.NewIntArgument(int(text_blink_mode))
	args := []gi.Argument{arg_v, arg_text_blink_mode}
	iv.Call(args, nil, nil)
}

// vte_terminal_set_word_char_exceptions
//
// [ exceptions ] trans: nothing
//
func (v Terminal) SetWordCharExceptions(exceptions string) {
	iv, err := _I.Get(117, "Terminal", "set_word_char_exceptions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_exceptions := gi.CString(exceptions)
	arg_v := gi.NewPointerArgument(v.P)
	arg_exceptions := gi.NewStringArgument(c_exceptions)
	args := []gi.Argument{arg_v, arg_exceptions}
	iv.Call(args, nil, nil)
	gi.Free(c_exceptions)
}

// Deprecated
//
// vte_terminal_spawn_sync
//
// [ pty_flags ] trans: nothing
//
// [ working_directory ] trans: nothing
//
// [ argv ] trans: nothing
//
// [ envv ] trans: nothing
//
// [ spawn_flags ] trans: nothing
//
// [ child_setup ] trans: nothing
//
// [ child_setup_data ] trans: nothing
//
// [ child_pid ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Terminal) SpawnSync(pty_flags PtyFlags, working_directory string, argv gi.CStrArray, envv gi.CStrArray, spawn_flags g.SpawnFlags, child_setup int /*TODO_TYPE CALLBACK*/, child_setup_data unsafe.Pointer, cancellable g.ICancellable) (result bool, child_pid int32, err error) {
	iv, err := _I.Get(118, "Terminal", "spawn_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_working_directory := gi.CString(working_directory)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pty_flags := gi.NewIntArgument(int(pty_flags))
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envv := gi.NewPointerArgument(envv.P)
	arg_spawn_flags := gi.NewIntArgument(int(spawn_flags))
	arg_child_setup := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_mySpawnChildSetupFunc()))
	arg_child_setup_data := gi.NewPointerArgument(child_setup_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_pty_flags, arg_working_directory, arg_argv, arg_envv, arg_spawn_flags, arg_child_setup, arg_child_setup_data, arg_child_pid, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[1].Pointer())
	child_pid = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// vte_terminal_unselect_all
//
func (v Terminal) UnselectAll() {
	iv, err := _I.Get(119, "Terminal", "unselect_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// vte_terminal_watch_child
//
// [ child_pid ] trans: nothing
//
func (v Terminal) WatchChild(child_pid int32) {
	iv, err := _I.Get(120, "Terminal", "watch_child")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child_pid := gi.NewInt32Argument(child_pid)
	args := []gi.Argument{arg_v, arg_child_pid}
	iv.Call(args, nil, nil)
}

// vte_terminal_write_contents_sync
//
// [ stream ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Terminal) WriteContentsSync(stream g.IOutputStream, flags WriteFlagsEnum, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(121, "Terminal", "write_contents_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if stream != nil {
		tmp = stream.P_OutputStream()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream, arg_flags, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// ignore GType struct TerminalClass

// Struct TerminalClassPrivate
type TerminalClassPrivate struct {
	P unsafe.Pointer
}

func TerminalClassPrivateGetType() gi.GType {
	ret := _I.GetGType(11, "TerminalClassPrivate")
	return ret
}

type TerminalSpawnAsyncCallbackStruct struct {
	F_terminal Terminal
	F_pid      int32
	F_error    unsafe.Pointer
}

func GetPointer_myTerminalSpawnAsyncCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myVteTerminalSpawnAsyncCallback())
}

//export myVteTerminalSpawnAsyncCallback
func myVteTerminalSpawnAsyncCallback(terminal *C.VteTerminal, pid C.gint32, error **C.GError, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TerminalSpawnAsyncCallbackStruct{
		F_terminal: WrapTerminal(unsafe.Pointer(terminal)),
		F_pid:      int32(pid),
		F_error:    unsafe.Pointer(error),
	}
	fn(args)
}

// Enum TextBlinkMode
type TextBlinkModeEnum int

const (
	TextBlinkModeNever     TextBlinkModeEnum = 0
	TextBlinkModeFocused   TextBlinkModeEnum = 1
	TextBlinkModeUnfocused TextBlinkModeEnum = 2
	TextBlinkModeAlways    TextBlinkModeEnum = 3
)

func TextBlinkModeGetType() gi.GType {
	ret := _I.GetGType(12, "TextBlinkMode")
	return ret
}

// Enum WriteFlags
type WriteFlagsEnum int

const (
	WriteFlagsDefault WriteFlagsEnum = 0
)

func WriteFlagsGetType() gi.GType {
	ret := _I.GetGType(13, "WriteFlags")
	return ret
}

// vte_get_features
//
// [ result ] trans: nothing
//
func GetFeatures() (result string) {
	iv, err := _I.Get(122, "get_features", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// vte_get_major_version
//
// [ result ] trans: nothing
//
func GetMajorVersion() (result uint32) {
	iv, err := _I.Get(123, "get_major_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// vte_get_micro_version
//
// [ result ] trans: nothing
//
func GetMicroVersion() (result uint32) {
	iv, err := _I.Get(124, "get_micro_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// vte_get_minor_version
//
// [ result ] trans: nothing
//
func GetMinorVersion() (result uint32) {
	iv, err := _I.Get(125, "get_minor_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// vte_get_user_shell
//
// [ result ] trans: everything
//
func GetUserShell() (result string) {
	iv, err := _I.Get(126, "get_user_shell", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// vte_pty_error_quark
//
// [ result ] trans: nothing
//
func PtyErrorQuark() (result uint32) {
	iv, err := _I.Get(127, "pty_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// vte_regex_error_quark
//
// [ result ] trans: nothing
//
func RegexErrorQuark() (result uint32) {
	iv, err := _I.Get(128, "regex_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// constants
const (
	MAJOR_VERSION        = 0
	MICRO_VERSION        = 2
	MINOR_VERSION        = 54
	REGEX_FLAGS_DEFAULT  = 1075314688
	SPAWN_NO_PARENT_ENVV = 33554432
	TEST_FLAGS_ALL       = 18446744073709551615
	TEST_FLAGS_NONE      = 0
)
