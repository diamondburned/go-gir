package rsvg

/*
#cgo pkg-config: librsvg-2.0
#include <librsvg/rsvg.h>
*/
import "C"
import "github.com/electricface/go-gir/cairo-1.0"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gdkpixbuf-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Rsvg")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Rsvg", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct DimensionData
type DimensionData struct {
	P unsafe.Pointer
}

const SizeOfStructDimensionData = 24

func DimensionDataGetType() gi.GType {
	ret := _I.GetGType(0, "DimensionData")
	return ret
}

// Enum Error
type ErrorEnum int

const (
	ErrorFailed ErrorEnum = 0
)

func ErrorGetType() gi.GType {
	ret := _I.GetGType(1, "Error")
	return ret
}

// Object Handle
type Handle struct {
	g.Object
}

func WrapHandle(p unsafe.Pointer) (r Handle) { r.P = p; return }

type IHandle interface{ P_Handle() unsafe.Pointer }

func (v Handle) P_Handle() unsafe.Pointer { return v.P }
func HandleGetType() gi.GType {
	ret := _I.GetGType(2, "Handle")
	return ret
}

// rsvg_handle_new
//
// [ result ] trans: everything
//
func NewHandle() (result Handle) {
	iv, err := _I.Get(0, "Handle", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// rsvg_handle_new_from_data
//
// [ data ] trans: nothing
//
// [ data_len ] trans: nothing
//
// [ result ] trans: everything
//
func NewHandleFromData(data gi.Uint8Array, data_len uint64) (result Handle, err error) {
	iv, err := _I.Get(1, "Handle", "new_from_data")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_data := gi.NewPointerArgument(data.P)
	arg_data_len := gi.NewUint64Argument(data_len)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_data, arg_data_len, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// rsvg_handle_new_from_file
//
// [ file_name ] trans: nothing
//
// [ result ] trans: everything
//
func NewHandleFromFile(file_name string) (result Handle, err error) {
	iv, err := _I.Get(2, "Handle", "new_from_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_file_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file_name)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// rsvg_handle_new_from_gfile_sync
//
// [ file ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewHandleFromGfileSync(file g.IFile, flags HandleFlags, cancellable g.ICancellable) (result Handle, err error) {
	iv, err := _I.Get(3, "Handle", "new_from_gfile_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if file != nil {
		tmp = file.P_File()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_file := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_file, arg_flags, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// rsvg_handle_new_from_stream_sync
//
// [ input_stream ] trans: nothing
//
// [ base_file ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewHandleFromStreamSync(input_stream g.IInputStream, base_file g.IFile, flags HandleFlags, cancellable g.ICancellable) (result Handle, err error) {
	iv, err := _I.Get(4, "Handle", "new_from_stream_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if input_stream != nil {
		tmp = input_stream.P_InputStream()
	}
	var tmp1 unsafe.Pointer
	if base_file != nil {
		tmp1 = base_file.P_File()
	}
	var tmp2 unsafe.Pointer
	if cancellable != nil {
		tmp2 = cancellable.P_Cancellable()
	}
	arg_input_stream := gi.NewPointerArgument(tmp)
	arg_base_file := gi.NewPointerArgument(tmp1)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_cancellable := gi.NewPointerArgument(tmp2)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_input_stream, arg_base_file, arg_flags, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// rsvg_handle_new_with_flags
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func NewHandleWithFlags(flags HandleFlags) (result Handle) {
	iv, err := _I.Get(5, "Handle", "new_with_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// rsvg_handle_close
//
// [ result ] trans: nothing
//
func (v Handle) Close() (result bool, err error) {
	iv, err := _I.Get(6, "Handle", "close")
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

// rsvg_handle_get_base_uri
//
// [ result ] trans: nothing
//
func (v Handle) GetBaseUri() (result string) {
	iv, err := _I.Get(7, "Handle", "get_base_uri")
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

// rsvg_handle_get_dimensions
//
// [ dimension_data ] trans: nothing, dir: out
//
func (v Handle) GetDimensions(dimension_data DimensionData) {
	iv, err := _I.Get(8, "Handle", "get_dimensions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dimension_data := gi.NewPointerArgument(dimension_data.P)
	args := []gi.Argument{arg_v, arg_dimension_data}
	iv.Call(args, nil, nil)
}

// rsvg_handle_get_dimensions_sub
//
// [ dimension_data ] trans: nothing, dir: out
//
// [ id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) GetDimensionsSub(dimension_data DimensionData, id string) (result bool) {
	iv, err := _I.Get(9, "Handle", "get_dimensions_sub")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_dimension_data := gi.NewPointerArgument(dimension_data.P)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_v, arg_dimension_data, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result = ret.Bool()
	return
}

// rsvg_handle_get_pixbuf
//
// [ result ] trans: everything
//
func (v Handle) GetPixbuf() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(10, "Handle", "get_pixbuf")
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

// rsvg_handle_get_pixbuf_sub
//
// [ id ] trans: nothing
//
// [ result ] trans: everything
//
func (v Handle) GetPixbufSub(id string) (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(11, "Handle", "get_pixbuf_sub")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_v, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result.P = ret.Pointer()
	return
}

// rsvg_handle_get_position_sub
//
// [ position_data ] trans: nothing, dir: out
//
// [ id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) GetPositionSub(position_data PositionData, id string) (result bool) {
	iv, err := _I.Get(12, "Handle", "get_position_sub")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_position_data := gi.NewPointerArgument(position_data.P)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_v, arg_position_data, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result = ret.Bool()
	return
}

// rsvg_handle_has_sub
//
// [ id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) HasSub(id string) (result bool) {
	iv, err := _I.Get(13, "Handle", "has_sub")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_v, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result = ret.Bool()
	return
}

// rsvg_handle_internal_set_testing
//
// [ testing ] trans: nothing
//
func (v Handle) InternalSetTesting(testing bool) {
	iv, err := _I.Get(14, "Handle", "internal_set_testing")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_testing := gi.NewBoolArgument(testing)
	args := []gi.Argument{arg_v, arg_testing}
	iv.Call(args, nil, nil)
}

// rsvg_handle_read_stream_sync
//
// [ stream ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) ReadStreamSync(stream g.IInputStream, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(15, "Handle", "read_stream_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if stream != nil {
		tmp = stream.P_InputStream()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(tmp)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// rsvg_handle_render_cairo
//
// [ cr ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) RenderCairo(cr cairo.Context) (result bool) {
	iv, err := _I.Get(16, "Handle", "render_cairo")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cr := gi.NewPointerArgument(cr.P)
	args := []gi.Argument{arg_v, arg_cr}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// rsvg_handle_render_cairo_sub
//
// [ cr ] trans: nothing
//
// [ id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) RenderCairoSub(cr cairo.Context, id string) (result bool) {
	iv, err := _I.Get(17, "Handle", "render_cairo_sub")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_v, arg_cr, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result = ret.Bool()
	return
}

// rsvg_handle_set_base_gfile
//
// [ base_file ] trans: nothing
//
func (v Handle) SetBaseGfile(base_file g.IFile) {
	iv, err := _I.Get(18, "Handle", "set_base_gfile")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if base_file != nil {
		tmp = base_file.P_File()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_base_file := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_base_file}
	iv.Call(args, nil, nil)
}

// rsvg_handle_set_base_uri
//
// [ base_uri ] trans: nothing
//
func (v Handle) SetBaseUri(base_uri string) {
	iv, err := _I.Get(19, "Handle", "set_base_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_base_uri := gi.CString(base_uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_base_uri := gi.NewStringArgument(c_base_uri)
	args := []gi.Argument{arg_v, arg_base_uri}
	iv.Call(args, nil, nil)
	gi.Free(c_base_uri)
}

// rsvg_handle_set_dpi
//
// [ dpi ] trans: nothing
//
func (v Handle) SetDpi(dpi float64) {
	iv, err := _I.Get(20, "Handle", "set_dpi")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dpi := gi.NewDoubleArgument(dpi)
	args := []gi.Argument{arg_v, arg_dpi}
	iv.Call(args, nil, nil)
}

// rsvg_handle_set_dpi_x_y
//
// [ dpi_x ] trans: nothing
//
// [ dpi_y ] trans: nothing
//
func (v Handle) SetDpiXY(dpi_x float64, dpi_y float64) {
	iv, err := _I.Get(21, "Handle", "set_dpi_x_y")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dpi_x := gi.NewDoubleArgument(dpi_x)
	arg_dpi_y := gi.NewDoubleArgument(dpi_y)
	args := []gi.Argument{arg_v, arg_dpi_x, arg_dpi_y}
	iv.Call(args, nil, nil)
}

// rsvg_handle_write
//
// [ buf ] trans: nothing
//
// [ count ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Handle) Write(buf gi.Uint8Array, count uint64) (result bool, err error) {
	iv, err := _I.Get(22, "Handle", "write")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_count := gi.NewUint64Argument(count)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// ignore GType struct HandleClass

// Flags HandleFlags
type HandleFlags int

const (
	HandleFlagsFlagsNone         HandleFlags = 0
	HandleFlagsFlagUnlimited     HandleFlags = 1
	HandleFlagsFlagKeepImageData HandleFlags = 2
)

func HandleFlagsGetType() gi.GType {
	ret := _I.GetGType(3, "HandleFlags")
	return ret
}

// Struct HandlePrivate
type HandlePrivate struct {
	P unsafe.Pointer
}

func HandlePrivateGetType() gi.GType {
	ret := _I.GetGType(4, "HandlePrivate")
	return ret
}

// Struct PositionData
type PositionData struct {
	P unsafe.Pointer
}

const SizeOfStructPositionData = 8

func PositionDataGetType() gi.GType {
	ret := _I.GetGType(5, "PositionData")
	return ret
}

// rsvg_cleanup
//
func Cleanup() {
	iv, err := _I.Get(23, "cleanup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// rsvg_error_quark
//
// [ result ] trans: nothing
//
func ErrorQuark() (result uint32) {
	iv, err := _I.Get(24, "error_quark", "")
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
// rsvg_set_default_dpi
//
// [ dpi ] trans: nothing
//
func SetDefaultDpi(dpi float64) {
	iv, err := _I.Get(25, "set_default_dpi", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dpi := gi.NewDoubleArgument(dpi)
	args := []gi.Argument{arg_dpi}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// rsvg_set_default_dpi_x_y
//
// [ dpi_x ] trans: nothing
//
// [ dpi_y ] trans: nothing
//
func SetDefaultDpiXY(dpi_x float64, dpi_y float64) {
	iv, err := _I.Get(26, "set_default_dpi_x_y", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dpi_x := gi.NewDoubleArgument(dpi_x)
	arg_dpi_y := gi.NewDoubleArgument(dpi_y)
	args := []gi.Argument{arg_dpi_x, arg_dpi_y}
	iv.Call(args, nil, nil)
}

// constants
const (
	MAJOR_VERSION = 2
	MICRO_VERSION = 10
	MINOR_VERSION = 44
	VERSION       = "2.44.10"
)
