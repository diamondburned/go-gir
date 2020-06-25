package gstcontroller

/*
#cgo pkg-config: gstreamer-controller-1.0
#include <gst/controller/controller.h>
extern void myGstControllerDirectControlBindingConvertGValue(GstDirectControlBinding* self, gdouble src_value, GValue* dest_value);
static void* getPointer_myGstControllerDirectControlBindingConvertGValue() {
return (void*)(myGstControllerDirectControlBindingConvertGValue);
}
extern void myGstControllerDirectControlBindingConvertValue(GstDirectControlBinding* self, gdouble src_value, gpointer dest_value);
static void* getPointer_myGstControllerDirectControlBindingConvertValue() {
return (void*)(myGstControllerDirectControlBindingConvertValue);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gst-1.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GstController")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GstController", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Object ARGBControlBinding
type ARGBControlBinding struct {
	gst.ControlBinding
}

func WrapARGBControlBinding(p unsafe.Pointer) (r ARGBControlBinding) { r.P = p; return }

type IARGBControlBinding interface{ P_ARGBControlBinding() unsafe.Pointer }

func (v ARGBControlBinding) P_ARGBControlBinding() unsafe.Pointer { return v.P }
func ARGBControlBindingGetType() gi.GType {
	ret := _I.GetGType(0, "ARGBControlBinding")
	return ret
}

// gst_argb_control_binding_new
//
// [ object ] trans: nothing
//
// [ property_name ] trans: nothing
//
// [ cs_a ] trans: nothing
//
// [ cs_r ] trans: nothing
//
// [ cs_g ] trans: nothing
//
// [ cs_b ] trans: nothing
//
// [ result ] trans: nothing
//
func NewARGBControlBinding(object gst.IObject, property_name string, cs_a gst.IControlSource, cs_r gst.IControlSource, cs_g gst.IControlSource, cs_b gst.IControlSource) (result ARGBControlBinding) {
	iv, err := _I.Get(0, "ARGBControlBinding", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_property_name := gi.CString(property_name)
	var tmp1 unsafe.Pointer
	if cs_a != nil {
		tmp1 = cs_a.P_ControlSource()
	}
	var tmp2 unsafe.Pointer
	if cs_r != nil {
		tmp2 = cs_r.P_ControlSource()
	}
	var tmp3 unsafe.Pointer
	if cs_g != nil {
		tmp3 = cs_g.P_ControlSource()
	}
	var tmp4 unsafe.Pointer
	if cs_b != nil {
		tmp4 = cs_b.P_ControlSource()
	}
	arg_object := gi.NewPointerArgument(tmp)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_cs_a := gi.NewPointerArgument(tmp1)
	arg_cs_r := gi.NewPointerArgument(tmp2)
	arg_cs_g := gi.NewPointerArgument(tmp3)
	arg_cs_b := gi.NewPointerArgument(tmp4)
	args := []gi.Argument{arg_object, arg_property_name, arg_cs_a, arg_cs_r, arg_cs_g, arg_cs_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// ignore GType struct ARGBControlBindingClass

// Struct ControlPoint
type ControlPoint struct {
	P unsafe.Pointer
}

const SizeOfStructControlPoint = 16

func ControlPointGetType() gi.GType {
	ret := _I.GetGType(1, "ControlPoint")
	return ret
}

// gst_control_point_copy
//
// [ result ] trans: everything
//
func (v ControlPoint) Copy() (result ControlPoint) {
	iv, err := _I.Get(1, "ControlPoint", "copy")
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

// gst_control_point_free
//
func (v ControlPoint) Free() {
	iv, err := _I.Get(2, "ControlPoint", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object DirectControlBinding
type DirectControlBinding struct {
	gst.ControlBinding
}

func WrapDirectControlBinding(p unsafe.Pointer) (r DirectControlBinding) { r.P = p; return }

type IDirectControlBinding interface{ P_DirectControlBinding() unsafe.Pointer }

func (v DirectControlBinding) P_DirectControlBinding() unsafe.Pointer { return v.P }
func DirectControlBindingGetType() gi.GType {
	ret := _I.GetGType(2, "DirectControlBinding")
	return ret
}

// gst_direct_control_binding_new
//
// [ object ] trans: nothing
//
// [ property_name ] trans: nothing
//
// [ cs ] trans: nothing
//
// [ result ] trans: nothing
//
func NewDirectControlBinding(object gst.IObject, property_name string, cs gst.IControlSource) (result DirectControlBinding) {
	iv, err := _I.Get(3, "DirectControlBinding", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_property_name := gi.CString(property_name)
	var tmp1 unsafe.Pointer
	if cs != nil {
		tmp1 = cs.P_ControlSource()
	}
	arg_object := gi.NewPointerArgument(tmp)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_cs := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_object, arg_property_name, arg_cs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// gst_direct_control_binding_new_absolute
//
// [ object ] trans: nothing
//
// [ property_name ] trans: nothing
//
// [ cs ] trans: nothing
//
// [ result ] trans: nothing
//
func NewDirectControlBindingAbsolute(object gst.IObject, property_name string, cs gst.IControlSource) (result DirectControlBinding) {
	iv, err := _I.Get(4, "DirectControlBinding", "new_absolute")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_property_name := gi.CString(property_name)
	var tmp1 unsafe.Pointer
	if cs != nil {
		tmp1 = cs.P_ControlSource()
	}
	arg_object := gi.NewPointerArgument(tmp)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_cs := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_object, arg_property_name, arg_cs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// ignore GType struct DirectControlBindingClass

type DirectControlBindingConvertGValueStruct struct {
	F_self       DirectControlBinding
	F_src_value  float64
	F_dest_value g.Value
}

func GetPointer_myDirectControlBindingConvertGValue() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstControllerDirectControlBindingConvertGValue())
}

//export myGstControllerDirectControlBindingConvertGValue
func myGstControllerDirectControlBindingConvertGValue(self *C.GstDirectControlBinding, src_value C.gdouble, dest_value *C.GValue) {
	// TODO: not found user_data
}

type DirectControlBindingConvertValueStruct struct {
	F_self       DirectControlBinding
	F_src_value  float64
	F_dest_value unsafe.Pointer
}

func GetPointer_myDirectControlBindingConvertValue() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstControllerDirectControlBindingConvertValue())
}

//export myGstControllerDirectControlBindingConvertValue
func myGstControllerDirectControlBindingConvertValue(self *C.GstDirectControlBinding, src_value C.gdouble, dest_value C.gpointer) {
	// TODO: not found user_data
}

// Object InterpolationControlSource
type InterpolationControlSource struct {
	TimedValueControlSource
}

func WrapInterpolationControlSource(p unsafe.Pointer) (r InterpolationControlSource) { r.P = p; return }

type IInterpolationControlSource interface{ P_InterpolationControlSource() unsafe.Pointer }

func (v InterpolationControlSource) P_InterpolationControlSource() unsafe.Pointer { return v.P }
func InterpolationControlSourceGetType() gi.GType {
	ret := _I.GetGType(3, "InterpolationControlSource")
	return ret
}

// gst_interpolation_control_source_new
//
// [ result ] trans: everything
//
func NewInterpolationControlSource() (result InterpolationControlSource) {
	iv, err := _I.Get(5, "InterpolationControlSource", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct InterpolationControlSourceClass

// Struct InterpolationControlSourcePrivate
type InterpolationControlSourcePrivate struct {
	P unsafe.Pointer
}

func InterpolationControlSourcePrivateGetType() gi.GType {
	ret := _I.GetGType(4, "InterpolationControlSourcePrivate")
	return ret
}

// Enum InterpolationMode
type InterpolationModeEnum int

const (
	InterpolationModeNone           InterpolationModeEnum = 0
	InterpolationModeLinear         InterpolationModeEnum = 1
	InterpolationModeCubic          InterpolationModeEnum = 2
	InterpolationModeCubicMonotonic InterpolationModeEnum = 3
)

func InterpolationModeGetType() gi.GType {
	ret := _I.GetGType(5, "InterpolationMode")
	return ret
}

// Object LFOControlSource
type LFOControlSource struct {
	gst.ControlSource
}

func WrapLFOControlSource(p unsafe.Pointer) (r LFOControlSource) { r.P = p; return }

type ILFOControlSource interface{ P_LFOControlSource() unsafe.Pointer }

func (v LFOControlSource) P_LFOControlSource() unsafe.Pointer { return v.P }
func LFOControlSourceGetType() gi.GType {
	ret := _I.GetGType(6, "LFOControlSource")
	return ret
}

// gst_lfo_control_source_new
//
// [ result ] trans: everything
//
func NewLFOControlSource() (result LFOControlSource) {
	iv, err := _I.Get(6, "LFOControlSource", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct LFOControlSourceClass

// Struct LFOControlSourcePrivate
type LFOControlSourcePrivate struct {
	P unsafe.Pointer
}

func LFOControlSourcePrivateGetType() gi.GType {
	ret := _I.GetGType(7, "LFOControlSourcePrivate")
	return ret
}

// Enum LFOWaveform
type LFOWaveformEnum int

const (
	LFOWaveformSine       LFOWaveformEnum = 0
	LFOWaveformSquare     LFOWaveformEnum = 1
	LFOWaveformSaw        LFOWaveformEnum = 2
	LFOWaveformReverseSaw LFOWaveformEnum = 3
	LFOWaveformTriangle   LFOWaveformEnum = 4
)

func LFOWaveformGetType() gi.GType {
	ret := _I.GetGType(8, "LFOWaveform")
	return ret
}

// Object ProxyControlBinding
type ProxyControlBinding struct {
	gst.ControlBinding
}

func WrapProxyControlBinding(p unsafe.Pointer) (r ProxyControlBinding) { r.P = p; return }

type IProxyControlBinding interface{ P_ProxyControlBinding() unsafe.Pointer }

func (v ProxyControlBinding) P_ProxyControlBinding() unsafe.Pointer { return v.P }
func ProxyControlBindingGetType() gi.GType {
	ret := _I.GetGType(9, "ProxyControlBinding")
	return ret
}

// gst_proxy_control_binding_new
//
// [ object ] trans: nothing
//
// [ property_name ] trans: nothing
//
// [ ref_object ] trans: nothing
//
// [ ref_property_name ] trans: nothing
//
// [ result ] trans: nothing
//
func NewProxyControlBinding(object gst.IObject, property_name string, ref_object gst.IObject, ref_property_name string) (result ProxyControlBinding) {
	iv, err := _I.Get(7, "ProxyControlBinding", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_property_name := gi.CString(property_name)
	var tmp1 unsafe.Pointer
	if ref_object != nil {
		tmp1 = ref_object.P_Object()
	}
	c_ref_property_name := gi.CString(ref_property_name)
	arg_object := gi.NewPointerArgument(tmp)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_ref_object := gi.NewPointerArgument(tmp1)
	arg_ref_property_name := gi.NewStringArgument(c_ref_property_name)
	args := []gi.Argument{arg_object, arg_property_name, arg_ref_object, arg_ref_property_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	gi.Free(c_ref_property_name)
	result.P = ret.Pointer()
	return
}

// ignore GType struct ProxyControlBindingClass

// Object TimedValueControlSource
type TimedValueControlSource struct {
	gst.ControlSource
}

func WrapTimedValueControlSource(p unsafe.Pointer) (r TimedValueControlSource) { r.P = p; return }

type ITimedValueControlSource interface{ P_TimedValueControlSource() unsafe.Pointer }

func (v TimedValueControlSource) P_TimedValueControlSource() unsafe.Pointer { return v.P }
func TimedValueControlSourceGetType() gi.GType {
	ret := _I.GetGType(10, "TimedValueControlSource")
	return ret
}

// gst_timed_value_control_source_find_control_point_iter
//
// [ timestamp ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimedValueControlSource) FindControlPointIter(timestamp uint64) (result g.SequenceIter) {
	iv, err := _I.Get(8, "TimedValueControlSource", "find_control_point_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_timed_value_control_source_get_all
//
// [ result ] trans: container
//
func (v TimedValueControlSource) GetAll() (result g.List) {
	iv, err := _I.Get(9, "TimedValueControlSource", "get_all")
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

// gst_timed_value_control_source_get_count
//
// [ result ] trans: nothing
//
func (v TimedValueControlSource) GetCount() (result int32) {
	iv, err := _I.Get(10, "TimedValueControlSource", "get_count")
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

// gst_timed_value_control_source_set
//
// [ timestamp ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimedValueControlSource) Set(timestamp uint64, value float64) (result bool) {
	iv, err := _I.Get(11, "TimedValueControlSource", "set")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_value := gi.NewDoubleArgument(value)
	args := []gi.Argument{arg_v, arg_timestamp, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_timed_value_control_source_set_from_list
//
// [ timedvalues ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimedValueControlSource) SetFromList(timedvalues g.SList) (result bool) {
	iv, err := _I.Get(12, "TimedValueControlSource", "set_from_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timedvalues := gi.NewPointerArgument(timedvalues.P)
	args := []gi.Argument{arg_v, arg_timedvalues}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_timed_value_control_source_unset
//
// [ timestamp ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TimedValueControlSource) Unset(timestamp uint64) (result bool) {
	iv, err := _I.Get(13, "TimedValueControlSource", "unset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_timed_value_control_source_unset_all
//
func (v TimedValueControlSource) UnsetAll() {
	iv, err := _I.Get(14, "TimedValueControlSource", "unset_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct TimedValueControlSourceClass

// Struct TimedValueControlSourcePrivate
type TimedValueControlSourcePrivate struct {
	P unsafe.Pointer
}

func TimedValueControlSourcePrivateGetType() gi.GType {
	ret := _I.GetGType(11, "TimedValueControlSourcePrivate")
	return ret
}

// Object TriggerControlSource
type TriggerControlSource struct {
	TimedValueControlSource
}

func WrapTriggerControlSource(p unsafe.Pointer) (r TriggerControlSource) { r.P = p; return }

type ITriggerControlSource interface{ P_TriggerControlSource() unsafe.Pointer }

func (v TriggerControlSource) P_TriggerControlSource() unsafe.Pointer { return v.P }
func TriggerControlSourceGetType() gi.GType {
	ret := _I.GetGType(12, "TriggerControlSource")
	return ret
}

// gst_trigger_control_source_new
//
// [ result ] trans: everything
//
func NewTriggerControlSource() (result TriggerControlSource) {
	iv, err := _I.Get(15, "TriggerControlSource", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct TriggerControlSourceClass

// Struct TriggerControlSourcePrivate
type TriggerControlSourcePrivate struct {
	P unsafe.Pointer
}

func TriggerControlSourcePrivateGetType() gi.GType {
	ret := _I.GetGType(13, "TriggerControlSourcePrivate")
	return ret
}

// gst_timed_value_control_invalidate_cache
//
// [ self ] trans: nothing
//
func TimedValueControlInvalidateCache(self ITimedValueControlSource) {
	iv, err := _I.Get(16, "timed_value_control_invalidate_cache", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if self != nil {
		tmp = self.P_TimedValueControlSource()
	}
	arg_self := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_self}
	iv.Call(args, nil, nil)
}

// constants
const ()
