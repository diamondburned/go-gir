package gstbase

/*
#cgo pkg-config: gstreamer-base-1.0
#include <gst/base/base.h>
extern void myGstBaseCollectDataDestroyNotify(GstCollectData* data);
static void* getPointer_myGstBaseCollectDataDestroyNotify() {
return (void*)(myGstBaseCollectDataDestroyNotify);
}
extern void myGstBaseCollectPadsBufferFunction(GstCollectPads* pads, GstCollectData* data, GstBuffer* buffer, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsBufferFunction() {
return (void*)(myGstBaseCollectPadsBufferFunction);
}
extern void myGstBaseCollectPadsClipFunction(GstCollectPads* pads, GstCollectData* data, GstBuffer* inbuffer, gpointer outbuffer, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsClipFunction() {
return (void*)(myGstBaseCollectPadsClipFunction);
}
extern void myGstBaseCollectPadsCompareFunction(GstCollectPads* pads, GstCollectData* data1, guint64 timestamp1, GstCollectData* data2, guint64 timestamp2, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsCompareFunction() {
return (void*)(myGstBaseCollectPadsCompareFunction);
}
extern void myGstBaseCollectPadsEventFunction(GstCollectPads* pads, GstCollectData* pad, GstEvent* event, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsEventFunction() {
return (void*)(myGstBaseCollectPadsEventFunction);
}
extern void myGstBaseCollectPadsFlushFunction(GstCollectPads* pads, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsFlushFunction() {
return (void*)(myGstBaseCollectPadsFlushFunction);
}
extern void myGstBaseCollectPadsFunction(GstCollectPads* pads, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsFunction() {
return (void*)(myGstBaseCollectPadsFunction);
}
extern void myGstBaseCollectPadsQueryFunction(GstCollectPads* pads, GstCollectData* pad, GstQuery* query, gpointer user_data);
static void* getPointer_myGstBaseCollectPadsQueryFunction() {
return (void*)(myGstBaseCollectPadsQueryFunction);
}
extern void myGstBaseDataQueueEmptyCallback(GstDataQueue* queue, gpointer checkdata);
static void* getPointer_myGstBaseDataQueueEmptyCallback() {
return (void*)(myGstBaseDataQueueEmptyCallback);
}
extern void myGstBaseDataQueueFullCallback(GstDataQueue* queue, gpointer checkdata);
static void* getPointer_myGstBaseDataQueueFullCallback() {
return (void*)(myGstBaseDataQueueFullCallback);
}
extern void myGstBaseTypeFindHelperGetRangeFunction(GstObject* obj, GstObject* parent, guint64 offset, guint32 length, gpointer buffer);
static void* getPointer_myGstBaseTypeFindHelperGetRangeFunction() {
return (void*)(myGstBaseTypeFindHelperGetRangeFunction);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gst-1.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GstBase")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GstBase", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Object Adapter
type Adapter struct {
	g.Object
}

func WrapAdapter(p unsafe.Pointer) (r Adapter) { r.P = p; return }

type IAdapter interface{ P_Adapter() unsafe.Pointer }

func (v Adapter) P_Adapter() unsafe.Pointer { return v.P }
func AdapterGetType() gi.GType {
	ret := _I.GetGType(0, "Adapter")
	return ret
}

// gst_adapter_new
//
// [ result ] trans: everything
//
func NewAdapter() (result Adapter) {
	iv, err := _I.Get(0, "Adapter", "new", 0, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_available
//
// [ result ] trans: nothing
//
func (v Adapter) Available() (result uint64) {
	iv, err := _I.Get(1, "Adapter", "available", 0, 1, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_available_fast
//
// [ result ] trans: nothing
//
func (v Adapter) AvailableFast() (result uint64) {
	iv, err := _I.Get(2, "Adapter", "available_fast", 0, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_clear
//
func (v Adapter) Clear() {
	iv, err := _I.Get(3, "Adapter", "clear", 0, 3, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_adapter_copy_bytes
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) Copy(offset uint64, size uint64) (result g.Bytes) {
	iv, err := _I.Get(4, "Adapter", "copy", 0, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_distance_from_discont
//
// [ result ] trans: nothing
//
func (v Adapter) DistanceFromDiscont() (result uint64) {
	iv, err := _I.Get(5, "Adapter", "distance_from_discont", 0, 5, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_dts_at_discont
//
// [ result ] trans: nothing
//
func (v Adapter) DtsAtDiscont() (result uint64) {
	iv, err := _I.Get(6, "Adapter", "dts_at_discont", 0, 6, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_flush
//
// [ flush ] trans: nothing
//
func (v Adapter) Flush(flush uint64) {
	iv, err := _I.Get(7, "Adapter", "flush", 0, 7, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flush := gi.NewUint64Argument(flush)
	args := []gi.Argument{arg_v, arg_flush}
	iv.Call(args, nil, nil)
}

// gst_adapter_get_buffer
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) GetBuffer(nbytes uint64) (result gst.Buffer) {
	iv, err := _I.Get(8, "Adapter", "get_buffer", 0, 8, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_get_buffer_fast
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) GetBufferFast(nbytes uint64) (result gst.Buffer) {
	iv, err := _I.Get(9, "Adapter", "get_buffer_fast", 0, 9, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_get_buffer_list
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) GetBufferList(nbytes uint64) (result gst.BufferList) {
	iv, err := _I.Get(10, "Adapter", "get_buffer_list", 0, 10, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_get_list
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) GetList(nbytes uint64) (result g.List) {
	iv, err := _I.Get(11, "Adapter", "get_list", 0, 11, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_map
//
// [ size ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) Map() (result gi.Uint8Array) {
	iv, err := _I.Get(12, "Adapter", "map", 0, 12, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_masked_scan_uint32
//
// [ mask ] trans: nothing
//
// [ pattern ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Adapter) MaskedScanUint32(mask uint32, pattern uint32, offset uint64, size uint64) (result int64) {
	iv, err := _I.Get(13, "Adapter", "masked_scan_uint32", 0, 13, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewUint32Argument(mask)
	arg_pattern := gi.NewUint32Argument(pattern)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_mask, arg_pattern, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gst_adapter_masked_scan_uint32_peek
//
// [ mask ] trans: nothing
//
// [ pattern ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) MaskedScanUint32Peek(mask uint32, pattern uint32, offset uint64, size uint64) (result int64, value uint32) {
	iv, err := _I.Get(14, "Adapter", "masked_scan_uint32_peek", 0, 14, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewUint32Argument(mask)
	arg_pattern := gi.NewUint32Argument(pattern)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_mask, arg_pattern, arg_offset, arg_size, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	value = outArgs[0].Uint32()
	result = ret.Int64()
	return
}

// gst_adapter_offset_at_discont
//
// [ result ] trans: nothing
//
func (v Adapter) OffsetAtDiscont() (result uint64) {
	iv, err := _I.Get(15, "Adapter", "offset_at_discont", 0, 15, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_prev_dts
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) PrevDts() (result uint64, distance uint64) {
	iv, err := _I.Get(16, "Adapter", "prev_dts", 0, 16, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Uint64()
	result = ret.Uint64()
	return
}

// gst_adapter_prev_dts_at_offset
//
// [ offset ] trans: nothing
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) PrevDtsAtOffset(offset uint64) (result uint64, distance uint64) {
	iv, err := _I.Get(17, "Adapter", "prev_dts_at_offset", 0, 17, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Uint64()
	result = ret.Uint64()
	return
}

// gst_adapter_prev_offset
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) PrevOffset() (result uint64, distance uint64) {
	iv, err := _I.Get(18, "Adapter", "prev_offset", 0, 18, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Uint64()
	result = ret.Uint64()
	return
}

// gst_adapter_prev_pts
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) PrevPts() (result uint64, distance uint64) {
	iv, err := _I.Get(19, "Adapter", "prev_pts", 0, 19, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Uint64()
	result = ret.Uint64()
	return
}

// gst_adapter_prev_pts_at_offset
//
// [ offset ] trans: nothing
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Adapter) PrevPtsAtOffset(offset uint64) (result uint64, distance uint64) {
	iv, err := _I.Get(20, "Adapter", "prev_pts_at_offset", 0, 20, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Uint64()
	result = ret.Uint64()
	return
}

// gst_adapter_pts_at_discont
//
// [ result ] trans: nothing
//
func (v Adapter) PtsAtDiscont() (result uint64) {
	iv, err := _I.Get(21, "Adapter", "pts_at_discont", 0, 21, gi.INFO_TYPE_OBJECT, 0)
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

// gst_adapter_push
//
// [ buf ] trans: everything
//
func (v Adapter) Push(buf gst.Buffer) {
	iv, err := _I.Get(22, "Adapter", "push", 0, 22, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// gst_adapter_take
//
// [ nbytes ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Adapter) Take() (result gi.Uint8Array) {
	iv, err := _I.Get(23, "Adapter", "take", 0, 23, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var nbytes uint64
	_ = nbytes
	nbytes = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(nbytes)}
	return
}

// gst_adapter_take_buffer
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) TakeBuffer(nbytes uint64) (result gst.Buffer) {
	iv, err := _I.Get(24, "Adapter", "take_buffer", 0, 24, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_take_buffer_fast
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) TakeBufferFast(nbytes uint64) (result gst.Buffer) {
	iv, err := _I.Get(25, "Adapter", "take_buffer_fast", 0, 25, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_take_buffer_list
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) TakeBufferList(nbytes uint64) (result gst.BufferList) {
	iv, err := _I.Get(26, "Adapter", "take_buffer_list", 0, 26, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_take_list
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: everything
//
func (v Adapter) TakeList(nbytes uint64) (result g.List) {
	iv, err := _I.Get(27, "Adapter", "take_list", 0, 27, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint64Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_adapter_unmap
//
func (v Adapter) Unmap() {
	iv, err := _I.Get(28, "Adapter", "unmap", 0, 28, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct AdapterClass

// Object Aggregator
type Aggregator struct {
	gst.Element
}

func WrapAggregator(p unsafe.Pointer) (r Aggregator) { r.P = p; return }

type IAggregator interface{ P_Aggregator() unsafe.Pointer }

func (v Aggregator) P_Aggregator() unsafe.Pointer { return v.P }
func AggregatorGetType() gi.GType {
	ret := _I.GetGType(1, "Aggregator")
	return ret
}

// gst_aggregator_finish_buffer
//
// [ buffer ] trans: everything
//
// [ result ] trans: nothing
//
func (v Aggregator) FinishBuffer(buffer gst.Buffer) (result gst.FlowReturnEnum) {
	iv, err := _I.Get(29, "Aggregator", "finish_buffer", 2, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_aggregator_get_allocator
//
// [ allocator ] trans: everything, dir: out
//
// [ params ] trans: everything, dir: out
//
func (v Aggregator) GetAllocator(params gst.AllocationParams) (allocator gst.Allocator) {
	iv, err := _I.Get(30, "Aggregator", "get_allocator", 2, 1, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_allocator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_allocator, arg_params}
	iv.Call(args, nil, &outArgs[0])
	allocator.P = outArgs[0].Pointer()
	return
}

// gst_aggregator_get_buffer_pool
//
// [ result ] trans: everything
//
func (v Aggregator) GetBufferPool() (result gst.BufferPool) {
	iv, err := _I.Get(31, "Aggregator", "get_buffer_pool", 2, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_aggregator_get_latency
//
// [ result ] trans: nothing
//
func (v Aggregator) GetLatency() (result uint64) {
	iv, err := _I.Get(32, "Aggregator", "get_latency", 2, 3, gi.INFO_TYPE_OBJECT, 0)
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

// gst_aggregator_set_latency
//
// [ min_latency ] trans: nothing
//
// [ max_latency ] trans: nothing
//
func (v Aggregator) SetLatency(min_latency uint64, max_latency uint64) {
	iv, err := _I.Get(33, "Aggregator", "set_latency", 2, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_min_latency := gi.NewUint64Argument(min_latency)
	arg_max_latency := gi.NewUint64Argument(max_latency)
	args := []gi.Argument{arg_v, arg_min_latency, arg_max_latency}
	iv.Call(args, nil, nil)
}

// gst_aggregator_set_src_caps
//
// [ caps ] trans: nothing
//
func (v Aggregator) SetSrcCaps(caps gst.Caps) {
	iv, err := _I.Get(34, "Aggregator", "set_src_caps", 2, 5, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_caps}
	iv.Call(args, nil, nil)
}

// ignore GType struct AggregatorClass

// Object AggregatorPad
type AggregatorPad struct {
	gst.Pad
}

func WrapAggregatorPad(p unsafe.Pointer) (r AggregatorPad) { r.P = p; return }

type IAggregatorPad interface{ P_AggregatorPad() unsafe.Pointer }

func (v AggregatorPad) P_AggregatorPad() unsafe.Pointer { return v.P }
func AggregatorPadGetType() gi.GType {
	ret := _I.GetGType(2, "AggregatorPad")
	return ret
}

// gst_aggregator_pad_drop_buffer
//
// [ result ] trans: nothing
//
func (v AggregatorPad) DropBuffer() (result bool) {
	iv, err := _I.Get(35, "AggregatorPad", "drop_buffer", 4, 0, gi.INFO_TYPE_OBJECT, 0)
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

// gst_aggregator_pad_has_buffer
//
// [ result ] trans: nothing
//
func (v AggregatorPad) HasBuffer() (result bool) {
	iv, err := _I.Get(36, "AggregatorPad", "has_buffer", 4, 1, gi.INFO_TYPE_OBJECT, 0)
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

// gst_aggregator_pad_is_eos
//
// [ result ] trans: nothing
//
func (v AggregatorPad) IsEos() (result bool) {
	iv, err := _I.Get(37, "AggregatorPad", "is_eos", 4, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_aggregator_pad_peek_buffer
//
// [ result ] trans: everything
//
func (v AggregatorPad) PeekBuffer() (result gst.Buffer) {
	iv, err := _I.Get(38, "AggregatorPad", "peek_buffer", 4, 3, gi.INFO_TYPE_OBJECT, 0)
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

// gst_aggregator_pad_pop_buffer
//
// [ result ] trans: everything
//
func (v AggregatorPad) PopBuffer() (result gst.Buffer) {
	iv, err := _I.Get(39, "AggregatorPad", "pop_buffer", 4, 4, gi.INFO_TYPE_OBJECT, 0)
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

// ignore GType struct AggregatorPadClass

// Struct AggregatorPadPrivate
type AggregatorPadPrivate struct {
	P unsafe.Pointer
}

func AggregatorPadPrivateGetType() gi.GType {
	ret := _I.GetGType(3, "AggregatorPadPrivate")
	return ret
}

// Struct AggregatorPrivate
type AggregatorPrivate struct {
	P unsafe.Pointer
}

func AggregatorPrivateGetType() gi.GType {
	ret := _I.GetGType(4, "AggregatorPrivate")
	return ret
}

// Object BaseParse
type BaseParse struct {
	gst.Element
}

func WrapBaseParse(p unsafe.Pointer) (r BaseParse) { r.P = p; return }

type IBaseParse interface{ P_BaseParse() unsafe.Pointer }

func (v BaseParse) P_BaseParse() unsafe.Pointer { return v.P }
func BaseParseGetType() gi.GType {
	ret := _I.GetGType(5, "BaseParse")
	return ret
}

// gst_base_parse_add_index_entry
//
// [ offset ] trans: nothing
//
// [ ts ] trans: nothing
//
// [ key ] trans: nothing
//
// [ force ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseParse) AddIndexEntry(offset uint64, ts uint64, key bool, force bool) (result bool) {
	iv, err := _I.Get(40, "BaseParse", "add_index_entry", 12, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_ts := gi.NewUint64Argument(ts)
	arg_key := gi.NewBoolArgument(key)
	arg_force := gi.NewBoolArgument(force)
	args := []gi.Argument{arg_v, arg_offset, arg_ts, arg_key, arg_force}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_base_parse_convert_default
//
// [ src_format ] trans: nothing
//
// [ src_value ] trans: nothing
//
// [ dest_format ] trans: nothing
//
// [ dest_value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BaseParse) ConvertDefault(src_format gst.FormatEnum, src_value int64, dest_format gst.FormatEnum) (result bool, dest_value int64) {
	iv, err := _I.Get(41, "BaseParse", "convert_default", 12, 1, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_format := gi.NewIntArgument(int(src_format))
	arg_src_value := gi.NewInt64Argument(src_value)
	arg_dest_format := gi.NewIntArgument(int(dest_format))
	arg_dest_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_src_format, arg_src_value, arg_dest_format, arg_dest_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	dest_value = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_base_parse_drain
//
func (v BaseParse) Drain() {
	iv, err := _I.Get(42, "BaseParse", "drain", 12, 2, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_base_parse_finish_frame
//
// [ frame ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseParse) FinishFrame(frame BaseParseFrame, size int32) (result gst.FlowReturnEnum) {
	iv, err := _I.Get(43, "BaseParse", "finish_frame", 12, 3, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_frame := gi.NewPointerArgument(frame.P)
	arg_size := gi.NewInt32Argument(size)
	args := []gi.Argument{arg_v, arg_frame, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_base_parse_merge_tags
//
// [ tags ] trans: nothing
//
// [ mode ] trans: nothing
//
func (v BaseParse) MergeTags(tags gst.TagList, mode gst.TagMergeModeEnum) {
	iv, err := _I.Get(44, "BaseParse", "merge_tags", 12, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tags := gi.NewPointerArgument(tags.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_tags, arg_mode}
	iv.Call(args, nil, nil)
}

// gst_base_parse_push_frame
//
// [ frame ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseParse) PushFrame(frame BaseParseFrame) (result gst.FlowReturnEnum) {
	iv, err := _I.Get(45, "BaseParse", "push_frame", 12, 5, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_frame := gi.NewPointerArgument(frame.P)
	args := []gi.Argument{arg_v, arg_frame}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_base_parse_set_average_bitrate
//
// [ bitrate ] trans: nothing
//
func (v BaseParse) SetAverageBitrate(bitrate uint32) {
	iv, err := _I.Get(46, "BaseParse", "set_average_bitrate", 12, 6, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bitrate := gi.NewUint32Argument(bitrate)
	args := []gi.Argument{arg_v, arg_bitrate}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_duration
//
// [ fmt ] trans: nothing
//
// [ duration ] trans: nothing
//
// [ interval ] trans: nothing
//
func (v BaseParse) SetDuration(fmt gst.FormatEnum, duration int64, interval int32) {
	iv, err := _I.Get(47, "BaseParse", "set_duration", 12, 7, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fmt := gi.NewIntArgument(int(fmt))
	arg_duration := gi.NewInt64Argument(duration)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_fmt, arg_duration, arg_interval}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_frame_rate
//
// [ fps_num ] trans: nothing
//
// [ fps_den ] trans: nothing
//
// [ lead_in ] trans: nothing
//
// [ lead_out ] trans: nothing
//
func (v BaseParse) SetFrameRate(fps_num uint32, fps_den uint32, lead_in uint32, lead_out uint32) {
	iv, err := _I.Get(48, "BaseParse", "set_frame_rate", 12, 8, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fps_num := gi.NewUint32Argument(fps_num)
	arg_fps_den := gi.NewUint32Argument(fps_den)
	arg_lead_in := gi.NewUint32Argument(lead_in)
	arg_lead_out := gi.NewUint32Argument(lead_out)
	args := []gi.Argument{arg_v, arg_fps_num, arg_fps_den, arg_lead_in, arg_lead_out}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_has_timing_info
//
// [ has_timing ] trans: nothing
//
func (v BaseParse) SetHasTimingInfo(has_timing bool) {
	iv, err := _I.Get(49, "BaseParse", "set_has_timing_info", 12, 9, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_has_timing := gi.NewBoolArgument(has_timing)
	args := []gi.Argument{arg_v, arg_has_timing}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_infer_ts
//
// [ infer_ts ] trans: nothing
//
func (v BaseParse) SetInferTs(infer_ts bool) {
	iv, err := _I.Get(50, "BaseParse", "set_infer_ts", 12, 10, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_infer_ts := gi.NewBoolArgument(infer_ts)
	args := []gi.Argument{arg_v, arg_infer_ts}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_latency
//
// [ min_latency ] trans: nothing
//
// [ max_latency ] trans: nothing
//
func (v BaseParse) SetLatency(min_latency uint64, max_latency uint64) {
	iv, err := _I.Get(51, "BaseParse", "set_latency", 12, 11, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_min_latency := gi.NewUint64Argument(min_latency)
	arg_max_latency := gi.NewUint64Argument(max_latency)
	args := []gi.Argument{arg_v, arg_min_latency, arg_max_latency}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_min_frame_size
//
// [ min_size ] trans: nothing
//
func (v BaseParse) SetMinFrameSize(min_size uint32) {
	iv, err := _I.Get(52, "BaseParse", "set_min_frame_size", 12, 12, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_min_size := gi.NewUint32Argument(min_size)
	args := []gi.Argument{arg_v, arg_min_size}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_passthrough
//
// [ passthrough ] trans: nothing
//
func (v BaseParse) SetPassthrough(passthrough bool) {
	iv, err := _I.Get(53, "BaseParse", "set_passthrough", 12, 13, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_passthrough := gi.NewBoolArgument(passthrough)
	args := []gi.Argument{arg_v, arg_passthrough}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_pts_interpolation
//
// [ pts_interpolate ] trans: nothing
//
func (v BaseParse) SetPtsInterpolation(pts_interpolate bool) {
	iv, err := _I.Get(54, "BaseParse", "set_pts_interpolation", 12, 14, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pts_interpolate := gi.NewBoolArgument(pts_interpolate)
	args := []gi.Argument{arg_v, arg_pts_interpolate}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_syncable
//
// [ syncable ] trans: nothing
//
func (v BaseParse) SetSyncable(syncable bool) {
	iv, err := _I.Get(55, "BaseParse", "set_syncable", 12, 15, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_syncable := gi.NewBoolArgument(syncable)
	args := []gi.Argument{arg_v, arg_syncable}
	iv.Call(args, nil, nil)
}

// gst_base_parse_set_ts_at_offset
//
// [ offset ] trans: nothing
//
func (v BaseParse) SetTsAtOffset(offset uint64) {
	iv, err := _I.Get(56, "BaseParse", "set_ts_at_offset", 12, 16, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	iv.Call(args, nil, nil)
}

// ignore GType struct BaseParseClass

// Struct BaseParseFrame
type BaseParseFrame struct {
	P unsafe.Pointer
}

const SizeOfStructBaseParseFrame = 72

func BaseParseFrameGetType() gi.GType {
	ret := _I.GetGType(6, "BaseParseFrame")
	return ret
}

// gst_base_parse_frame_new
//
// [ buffer ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ overhead ] trans: nothing
//
// [ result ] trans: everything
//
func NewBaseParseFrame(buffer gst.Buffer, flags BaseParseFrameFlags, overhead int32) (result BaseParseFrame) {
	iv, err := _I.Get(57, "BaseParseFrame", "new", 14, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_overhead := gi.NewInt32Argument(overhead)
	args := []gi.Argument{arg_buffer, arg_flags, arg_overhead}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_base_parse_frame_copy
//
// [ result ] trans: everything
//
func (v BaseParseFrame) Copy() (result BaseParseFrame) {
	iv, err := _I.Get(58, "BaseParseFrame", "copy", 14, 1, gi.INFO_TYPE_STRUCT, 0)
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

// gst_base_parse_frame_free
//
func (v BaseParseFrame) Free() {
	iv, err := _I.Get(59, "BaseParseFrame", "free", 14, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_base_parse_frame_init
//
func (v BaseParseFrame) Init() {
	iv, err := _I.Get(60, "BaseParseFrame", "init", 14, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags BaseParseFrameFlags
type BaseParseFrameFlags int

const (
	BaseParseFrameFlagsNone     BaseParseFrameFlags = 0
	BaseParseFrameFlagsNewFrame BaseParseFrameFlags = 1
	BaseParseFrameFlagsNoFrame  BaseParseFrameFlags = 2
	BaseParseFrameFlagsClip     BaseParseFrameFlags = 4
	BaseParseFrameFlagsDrop     BaseParseFrameFlags = 8
	BaseParseFrameFlagsQueue    BaseParseFrameFlags = 16
)

func BaseParseFrameFlagsGetType() gi.GType {
	ret := _I.GetGType(7, "BaseParseFrameFlags")
	return ret
}

// Struct BaseParsePrivate
type BaseParsePrivate struct {
	P unsafe.Pointer
}

func BaseParsePrivateGetType() gi.GType {
	ret := _I.GetGType(8, "BaseParsePrivate")
	return ret
}

// Object BaseSink
type BaseSink struct {
	gst.Element
}

func WrapBaseSink(p unsafe.Pointer) (r BaseSink) { r.P = p; return }

type IBaseSink interface{ P_BaseSink() unsafe.Pointer }

func (v BaseSink) P_BaseSink() unsafe.Pointer { return v.P }
func BaseSinkGetType() gi.GType {
	ret := _I.GetGType(9, "BaseSink")
	return ret
}

// gst_base_sink_do_preroll
//
// [ obj ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseSink) DoPreroll(obj gst.MiniObject) (result gst.FlowReturnEnum) {
	iv, err := _I.Get(61, "BaseSink", "do_preroll", 17, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_obj := gi.NewPointerArgument(obj.P)
	args := []gi.Argument{arg_v, arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_base_sink_get_blocksize
//
// [ result ] trans: nothing
//
func (v BaseSink) GetBlocksize() (result uint32) {
	iv, err := _I.Get(62, "BaseSink", "get_blocksize", 17, 1, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_drop_out_of_segment
//
// [ result ] trans: nothing
//
func (v BaseSink) GetDropOutOfSegment() (result bool) {
	iv, err := _I.Get(63, "BaseSink", "get_drop_out_of_segment", 17, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_last_sample
//
// [ result ] trans: everything
//
func (v BaseSink) GetLastSample() (result gst.Sample) {
	iv, err := _I.Get(64, "BaseSink", "get_last_sample", 17, 3, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_latency
//
// [ result ] trans: nothing
//
func (v BaseSink) GetLatency() (result uint64) {
	iv, err := _I.Get(65, "BaseSink", "get_latency", 17, 4, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_max_bitrate
//
// [ result ] trans: nothing
//
func (v BaseSink) GetMaxBitrate() (result uint64) {
	iv, err := _I.Get(66, "BaseSink", "get_max_bitrate", 17, 5, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_max_lateness
//
// [ result ] trans: nothing
//
func (v BaseSink) GetMaxLateness() (result int64) {
	iv, err := _I.Get(67, "BaseSink", "get_max_lateness", 17, 6, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_render_delay
//
// [ result ] trans: nothing
//
func (v BaseSink) GetRenderDelay() (result uint64) {
	iv, err := _I.Get(68, "BaseSink", "get_render_delay", 17, 7, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_sync
//
// [ result ] trans: nothing
//
func (v BaseSink) GetSync() (result bool) {
	iv, err := _I.Get(69, "BaseSink", "get_sync", 17, 8, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_throttle_time
//
// [ result ] trans: nothing
//
func (v BaseSink) GetThrottleTime() (result uint64) {
	iv, err := _I.Get(70, "BaseSink", "get_throttle_time", 17, 9, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_get_ts_offset
//
// [ result ] trans: nothing
//
func (v BaseSink) GetTsOffset() (result int64) {
	iv, err := _I.Get(71, "BaseSink", "get_ts_offset", 17, 10, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_is_async_enabled
//
// [ result ] trans: nothing
//
func (v BaseSink) IsAsyncEnabled() (result bool) {
	iv, err := _I.Get(72, "BaseSink", "is_async_enabled", 17, 11, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_is_last_sample_enabled
//
// [ result ] trans: nothing
//
func (v BaseSink) IsLastSampleEnabled() (result bool) {
	iv, err := _I.Get(73, "BaseSink", "is_last_sample_enabled", 17, 12, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_is_qos_enabled
//
// [ result ] trans: nothing
//
func (v BaseSink) IsQosEnabled() (result bool) {
	iv, err := _I.Get(74, "BaseSink", "is_qos_enabled", 17, 13, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_sink_query_latency
//
// [ live ] trans: everything, dir: out
//
// [ upstream_live ] trans: everything, dir: out
//
// [ min_latency ] trans: everything, dir: out
//
// [ max_latency ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BaseSink) QueryLatency() (result bool, live bool, upstream_live bool, min_latency uint64, max_latency uint64) {
	iv, err := _I.Get(75, "BaseSink", "query_latency", 17, 14, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_live := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_upstream_live := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_min_latency := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_max_latency := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_live, arg_upstream_live, arg_min_latency, arg_max_latency}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	live = outArgs[0].Bool()
	upstream_live = outArgs[1].Bool()
	min_latency = outArgs[2].Uint64()
	max_latency = outArgs[3].Uint64()
	result = ret.Bool()
	return
}

// gst_base_sink_set_async_enabled
//
// [ enabled ] trans: nothing
//
func (v BaseSink) SetAsyncEnabled(enabled bool) {
	iv, err := _I.Get(76, "BaseSink", "set_async_enabled", 17, 15, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_enabled}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_blocksize
//
// [ blocksize ] trans: nothing
//
func (v BaseSink) SetBlocksize(blocksize uint32) {
	iv, err := _I.Get(77, "BaseSink", "set_blocksize", 17, 16, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_blocksize := gi.NewUint32Argument(blocksize)
	args := []gi.Argument{arg_v, arg_blocksize}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_drop_out_of_segment
//
// [ drop_out_of_segment ] trans: nothing
//
func (v BaseSink) SetDropOutOfSegment(drop_out_of_segment bool) {
	iv, err := _I.Get(78, "BaseSink", "set_drop_out_of_segment", 17, 17, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_drop_out_of_segment := gi.NewBoolArgument(drop_out_of_segment)
	args := []gi.Argument{arg_v, arg_drop_out_of_segment}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_last_sample_enabled
//
// [ enabled ] trans: nothing
//
func (v BaseSink) SetLastSampleEnabled(enabled bool) {
	iv, err := _I.Get(79, "BaseSink", "set_last_sample_enabled", 17, 18, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_enabled}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_max_bitrate
//
// [ max_bitrate ] trans: nothing
//
func (v BaseSink) SetMaxBitrate(max_bitrate uint64) {
	iv, err := _I.Get(80, "BaseSink", "set_max_bitrate", 17, 19, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_bitrate := gi.NewUint64Argument(max_bitrate)
	args := []gi.Argument{arg_v, arg_max_bitrate}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_max_lateness
//
// [ max_lateness ] trans: nothing
//
func (v BaseSink) SetMaxLateness(max_lateness int64) {
	iv, err := _I.Get(81, "BaseSink", "set_max_lateness", 17, 20, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_lateness := gi.NewInt64Argument(max_lateness)
	args := []gi.Argument{arg_v, arg_max_lateness}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_qos_enabled
//
// [ enabled ] trans: nothing
//
func (v BaseSink) SetQosEnabled(enabled bool) {
	iv, err := _I.Get(82, "BaseSink", "set_qos_enabled", 17, 21, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_enabled}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_render_delay
//
// [ delay ] trans: nothing
//
func (v BaseSink) SetRenderDelay(delay uint64) {
	iv, err := _I.Get(83, "BaseSink", "set_render_delay", 17, 22, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_delay := gi.NewUint64Argument(delay)
	args := []gi.Argument{arg_v, arg_delay}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_sync
//
// [ sync ] trans: nothing
//
func (v BaseSink) SetSync(sync bool) {
	iv, err := _I.Get(84, "BaseSink", "set_sync", 17, 23, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sync := gi.NewBoolArgument(sync)
	args := []gi.Argument{arg_v, arg_sync}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_throttle_time
//
// [ throttle ] trans: nothing
//
func (v BaseSink) SetThrottleTime(throttle uint64) {
	iv, err := _I.Get(85, "BaseSink", "set_throttle_time", 17, 24, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_throttle := gi.NewUint64Argument(throttle)
	args := []gi.Argument{arg_v, arg_throttle}
	iv.Call(args, nil, nil)
}

// gst_base_sink_set_ts_offset
//
// [ offset ] trans: nothing
//
func (v BaseSink) SetTsOffset(offset int64) {
	iv, err := _I.Get(86, "BaseSink", "set_ts_offset", 17, 25, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	iv.Call(args, nil, nil)
}

// gst_base_sink_wait
//
// [ time ] trans: nothing
//
// [ jitter ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BaseSink) Wait(time uint64) (result gst.FlowReturnEnum, jitter int64) {
	iv, err := _I.Get(87, "BaseSink", "wait", 17, 26, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_time := gi.NewUint64Argument(time)
	arg_jitter := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_time, arg_jitter}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	jitter = outArgs[0].Int64()
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_base_sink_wait_clock
//
// [ time ] trans: nothing
//
// [ jitter ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BaseSink) WaitClock(time uint64) (result gst.ClockReturnEnum, jitter int64) {
	iv, err := _I.Get(88, "BaseSink", "wait_clock", 17, 27, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_time := gi.NewUint64Argument(time)
	arg_jitter := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_time, arg_jitter}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	jitter = outArgs[0].Int64()
	result = gst.ClockReturnEnum(ret.Int())
	return
}

// gst_base_sink_wait_preroll
//
// [ result ] trans: nothing
//
func (v BaseSink) WaitPreroll() (result gst.FlowReturnEnum) {
	iv, err := _I.Get(89, "BaseSink", "wait_preroll", 17, 28, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// ignore GType struct BaseSinkClass

// Struct BaseSinkPrivate
type BaseSinkPrivate struct {
	P unsafe.Pointer
}

func BaseSinkPrivateGetType() gi.GType {
	ret := _I.GetGType(10, "BaseSinkPrivate")
	return ret
}

// Object BaseSrc
type BaseSrc struct {
	gst.Element
}

func WrapBaseSrc(p unsafe.Pointer) (r BaseSrc) { r.P = p; return }

type IBaseSrc interface{ P_BaseSrc() unsafe.Pointer }

func (v BaseSrc) P_BaseSrc() unsafe.Pointer { return v.P }
func BaseSrcGetType() gi.GType {
	ret := _I.GetGType(11, "BaseSrc")
	return ret
}

// gst_base_src_get_allocator
//
// [ allocator ] trans: everything, dir: out
//
// [ params ] trans: everything, dir: out
//
func (v BaseSrc) GetAllocator(params gst.AllocationParams) (allocator gst.Allocator) {
	iv, err := _I.Get(90, "BaseSrc", "get_allocator", 20, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_allocator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_allocator, arg_params}
	iv.Call(args, nil, &outArgs[0])
	allocator.P = outArgs[0].Pointer()
	return
}

// gst_base_src_get_blocksize
//
// [ result ] trans: nothing
//
func (v BaseSrc) GetBlocksize() (result uint32) {
	iv, err := _I.Get(91, "BaseSrc", "get_blocksize", 20, 1, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_src_get_buffer_pool
//
// [ result ] trans: everything
//
func (v BaseSrc) GetBufferPool() (result gst.BufferPool) {
	iv, err := _I.Get(92, "BaseSrc", "get_buffer_pool", 20, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_src_get_do_timestamp
//
// [ result ] trans: nothing
//
func (v BaseSrc) GetDoTimestamp() (result bool) {
	iv, err := _I.Get(93, "BaseSrc", "get_do_timestamp", 20, 3, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_src_is_async
//
// [ result ] trans: nothing
//
func (v BaseSrc) IsAsync() (result bool) {
	iv, err := _I.Get(94, "BaseSrc", "is_async", 20, 4, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_src_is_live
//
// [ result ] trans: nothing
//
func (v BaseSrc) IsLive() (result bool) {
	iv, err := _I.Get(95, "BaseSrc", "is_live", 20, 5, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_src_new_seamless_segment
//
// [ start ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseSrc) NewSeamlessSegment(start int64, stop int64, time int64) (result bool) {
	iv, err := _I.Get(96, "BaseSrc", "new_seamless_segment", 20, 6, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewInt64Argument(start)
	arg_stop := gi.NewInt64Argument(stop)
	arg_time := gi.NewInt64Argument(time)
	args := []gi.Argument{arg_v, arg_start, arg_stop, arg_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_base_src_query_latency
//
// [ live ] trans: everything, dir: out
//
// [ min_latency ] trans: everything, dir: out
//
// [ max_latency ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v BaseSrc) QueryLatency() (result bool, live bool, min_latency uint64, max_latency uint64) {
	iv, err := _I.Get(97, "BaseSrc", "query_latency", 20, 7, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_live := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_min_latency := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_max_latency := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_live, arg_min_latency, arg_max_latency}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	live = outArgs[0].Bool()
	min_latency = outArgs[1].Uint64()
	max_latency = outArgs[2].Uint64()
	result = ret.Bool()
	return
}

// gst_base_src_set_async
//
// [ async ] trans: nothing
//
func (v BaseSrc) SetAsync(async bool) {
	iv, err := _I.Get(98, "BaseSrc", "set_async", 20, 8, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_async := gi.NewBoolArgument(async)
	args := []gi.Argument{arg_v, arg_async}
	iv.Call(args, nil, nil)
}

// gst_base_src_set_automatic_eos
//
// [ automatic_eos ] trans: nothing
//
func (v BaseSrc) SetAutomaticEos(automatic_eos bool) {
	iv, err := _I.Get(99, "BaseSrc", "set_automatic_eos", 20, 9, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_automatic_eos := gi.NewBoolArgument(automatic_eos)
	args := []gi.Argument{arg_v, arg_automatic_eos}
	iv.Call(args, nil, nil)
}

// gst_base_src_set_blocksize
//
// [ blocksize ] trans: nothing
//
func (v BaseSrc) SetBlocksize(blocksize uint32) {
	iv, err := _I.Get(100, "BaseSrc", "set_blocksize", 20, 10, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_blocksize := gi.NewUint32Argument(blocksize)
	args := []gi.Argument{arg_v, arg_blocksize}
	iv.Call(args, nil, nil)
}

// gst_base_src_set_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseSrc) SetCaps(caps gst.Caps) (result bool) {
	iv, err := _I.Get(101, "BaseSrc", "set_caps", 20, 11, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_base_src_set_do_timestamp
//
// [ timestamp ] trans: nothing
//
func (v BaseSrc) SetDoTimestamp(timestamp bool) {
	iv, err := _I.Get(102, "BaseSrc", "set_do_timestamp", 20, 12, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewBoolArgument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// gst_base_src_set_dynamic_size
//
// [ dynamic ] trans: nothing
//
func (v BaseSrc) SetDynamicSize(dynamic bool) {
	iv, err := _I.Get(103, "BaseSrc", "set_dynamic_size", 20, 13, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dynamic := gi.NewBoolArgument(dynamic)
	args := []gi.Argument{arg_v, arg_dynamic}
	iv.Call(args, nil, nil)
}

// gst_base_src_set_format
//
// [ format ] trans: nothing
//
func (v BaseSrc) SetFormat(format gst.FormatEnum) {
	iv, err := _I.Get(104, "BaseSrc", "set_format", 20, 14, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_v, arg_format}
	iv.Call(args, nil, nil)
}

// gst_base_src_set_live
//
// [ live ] trans: nothing
//
func (v BaseSrc) SetLive(live bool) {
	iv, err := _I.Get(105, "BaseSrc", "set_live", 20, 15, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_live := gi.NewBoolArgument(live)
	args := []gi.Argument{arg_v, arg_live}
	iv.Call(args, nil, nil)
}

// gst_base_src_start_complete
//
// [ ret ] trans: nothing
//
func (v BaseSrc) StartComplete(ret gst.FlowReturnEnum) {
	iv, err := _I.Get(106, "BaseSrc", "start_complete", 20, 16, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ret := gi.NewIntArgument(int(ret))
	args := []gi.Argument{arg_v, arg_ret}
	iv.Call(args, nil, nil)
}

// gst_base_src_start_wait
//
// [ result ] trans: nothing
//
func (v BaseSrc) StartWait() (result gst.FlowReturnEnum) {
	iv, err := _I.Get(107, "BaseSrc", "start_wait", 20, 17, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_base_src_submit_buffer_list
//
// [ buffer_list ] trans: everything
//
func (v BaseSrc) SubmitBufferList(buffer_list gst.BufferList) {
	iv, err := _I.Get(108, "BaseSrc", "submit_buffer_list", 20, 18, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer_list := gi.NewPointerArgument(buffer_list.P)
	args := []gi.Argument{arg_v, arg_buffer_list}
	iv.Call(args, nil, nil)
}

// gst_base_src_wait_playing
//
// [ result ] trans: nothing
//
func (v BaseSrc) WaitPlaying() (result gst.FlowReturnEnum) {
	iv, err := _I.Get(109, "BaseSrc", "wait_playing", 20, 19, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// ignore GType struct BaseSrcClass

// Flags BaseSrcFlags
type BaseSrcFlags int

const (
	BaseSrcFlagsStarting BaseSrcFlags = 16384
	BaseSrcFlagsStarted  BaseSrcFlags = 32768
	BaseSrcFlagsLast     BaseSrcFlags = 1048576
)

func BaseSrcFlagsGetType() gi.GType {
	ret := _I.GetGType(12, "BaseSrcFlags")
	return ret
}

// Struct BaseSrcPrivate
type BaseSrcPrivate struct {
	P unsafe.Pointer
}

func BaseSrcPrivateGetType() gi.GType {
	ret := _I.GetGType(13, "BaseSrcPrivate")
	return ret
}

// Object BaseTransform
type BaseTransform struct {
	gst.Element
}

func WrapBaseTransform(p unsafe.Pointer) (r BaseTransform) { r.P = p; return }

type IBaseTransform interface{ P_BaseTransform() unsafe.Pointer }

func (v BaseTransform) P_BaseTransform() unsafe.Pointer { return v.P }
func BaseTransformGetType() gi.GType {
	ret := _I.GetGType(14, "BaseTransform")
	return ret
}

// gst_base_transform_get_allocator
//
// [ allocator ] trans: everything, dir: out
//
// [ params ] trans: everything, dir: out
//
func (v BaseTransform) GetAllocator(params gst.AllocationParams) (allocator gst.Allocator) {
	iv, err := _I.Get(110, "BaseTransform", "get_allocator", 24, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_allocator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_allocator, arg_params}
	iv.Call(args, nil, &outArgs[0])
	allocator.P = outArgs[0].Pointer()
	return
}

// gst_base_transform_get_buffer_pool
//
// [ result ] trans: everything
//
func (v BaseTransform) GetBufferPool() (result gst.BufferPool) {
	iv, err := _I.Get(111, "BaseTransform", "get_buffer_pool", 24, 1, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_transform_is_in_place
//
// [ result ] trans: nothing
//
func (v BaseTransform) IsInPlace() (result bool) {
	iv, err := _I.Get(112, "BaseTransform", "is_in_place", 24, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_transform_is_passthrough
//
// [ result ] trans: nothing
//
func (v BaseTransform) IsPassthrough() (result bool) {
	iv, err := _I.Get(113, "BaseTransform", "is_passthrough", 24, 3, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_transform_is_qos_enabled
//
// [ result ] trans: nothing
//
func (v BaseTransform) IsQosEnabled() (result bool) {
	iv, err := _I.Get(114, "BaseTransform", "is_qos_enabled", 24, 4, gi.INFO_TYPE_OBJECT, 0)
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

// gst_base_transform_reconfigure_sink
//
func (v BaseTransform) ReconfigureSink() {
	iv, err := _I.Get(115, "BaseTransform", "reconfigure_sink", 24, 5, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_base_transform_reconfigure_src
//
func (v BaseTransform) ReconfigureSrc() {
	iv, err := _I.Get(116, "BaseTransform", "reconfigure_src", 24, 6, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_base_transform_set_gap_aware
//
// [ gap_aware ] trans: nothing
//
func (v BaseTransform) SetGapAware(gap_aware bool) {
	iv, err := _I.Get(117, "BaseTransform", "set_gap_aware", 24, 7, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gap_aware := gi.NewBoolArgument(gap_aware)
	args := []gi.Argument{arg_v, arg_gap_aware}
	iv.Call(args, nil, nil)
}

// gst_base_transform_set_in_place
//
// [ in_place ] trans: nothing
//
func (v BaseTransform) SetInPlace(in_place bool) {
	iv, err := _I.Get(118, "BaseTransform", "set_in_place", 24, 8, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_in_place := gi.NewBoolArgument(in_place)
	args := []gi.Argument{arg_v, arg_in_place}
	iv.Call(args, nil, nil)
}

// gst_base_transform_set_passthrough
//
// [ passthrough ] trans: nothing
//
func (v BaseTransform) SetPassthrough(passthrough bool) {
	iv, err := _I.Get(119, "BaseTransform", "set_passthrough", 24, 9, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_passthrough := gi.NewBoolArgument(passthrough)
	args := []gi.Argument{arg_v, arg_passthrough}
	iv.Call(args, nil, nil)
}

// gst_base_transform_set_prefer_passthrough
//
// [ prefer_passthrough ] trans: nothing
//
func (v BaseTransform) SetPreferPassthrough(prefer_passthrough bool) {
	iv, err := _I.Get(120, "BaseTransform", "set_prefer_passthrough", 24, 10, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_prefer_passthrough := gi.NewBoolArgument(prefer_passthrough)
	args := []gi.Argument{arg_v, arg_prefer_passthrough}
	iv.Call(args, nil, nil)
}

// gst_base_transform_set_qos_enabled
//
// [ enabled ] trans: nothing
//
func (v BaseTransform) SetQosEnabled(enabled bool) {
	iv, err := _I.Get(121, "BaseTransform", "set_qos_enabled", 24, 11, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_enabled}
	iv.Call(args, nil, nil)
}

// gst_base_transform_update_qos
//
// [ proportion ] trans: nothing
//
// [ diff ] trans: nothing
//
// [ timestamp ] trans: nothing
//
func (v BaseTransform) UpdateQos(proportion float64, diff int64, timestamp uint64) {
	iv, err := _I.Get(122, "BaseTransform", "update_qos", 24, 12, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_proportion := gi.NewDoubleArgument(proportion)
	arg_diff := gi.NewInt64Argument(diff)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	args := []gi.Argument{arg_v, arg_proportion, arg_diff, arg_timestamp}
	iv.Call(args, nil, nil)
}

// gst_base_transform_update_src_caps
//
// [ updated_caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseTransform) UpdateSrcCaps(updated_caps gst.Caps) (result bool) {
	iv, err := _I.Get(123, "BaseTransform", "update_src_caps", 24, 13, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_updated_caps := gi.NewPointerArgument(updated_caps.P)
	args := []gi.Argument{arg_v, arg_updated_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct BaseTransformClass

// Struct BaseTransformPrivate
type BaseTransformPrivate struct {
	P unsafe.Pointer
}

func BaseTransformPrivateGetType() gi.GType {
	ret := _I.GetGType(15, "BaseTransformPrivate")
	return ret
}

// Struct BitReader
type BitReader struct {
	P unsafe.Pointer
}

const SizeOfStructBitReader = 56

func BitReaderGetType() gi.GType {
	ret := _I.GetGType(16, "BitReader")
	return ret
}

// gst_bit_reader_free
//
func (v BitReader) Free() {
	iv, err := _I.Get(124, "BitReader", "free", 27, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_bit_reader_get_bits_uint16
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) GetBitsUint16(nbits uint32) (result bool, val uint16) {
	iv, err := _I.Get(125, "BitReader", "get_bits_uint16", 27, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gst_bit_reader_get_bits_uint32
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) GetBitsUint32(nbits uint32) (result bool, val uint32) {
	iv, err := _I.Get(126, "BitReader", "get_bits_uint32", 27, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_bit_reader_get_bits_uint64
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) GetBitsUint64(nbits uint32) (result bool, val uint64) {
	iv, err := _I.Get(127, "BitReader", "get_bits_uint64", 27, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_bit_reader_get_bits_uint8
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) GetBitsUint8(nbits uint32) (result bool, val uint8) {
	iv, err := _I.Get(128, "BitReader", "get_bits_uint8", 27, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint8()
	result = ret.Bool()
	return
}

// gst_bit_reader_get_pos
//
// [ result ] trans: nothing
//
func (v BitReader) GetPos() (result uint32) {
	iv, err := _I.Get(129, "BitReader", "get_pos", 27, 5, gi.INFO_TYPE_STRUCT, 0)
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

// gst_bit_reader_get_remaining
//
// [ result ] trans: nothing
//
func (v BitReader) GetRemaining() (result uint32) {
	iv, err := _I.Get(130, "BitReader", "get_remaining", 27, 6, gi.INFO_TYPE_STRUCT, 0)
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

// gst_bit_reader_get_size
//
// [ result ] trans: nothing
//
func (v BitReader) GetSize() (result uint32) {
	iv, err := _I.Get(131, "BitReader", "get_size", 27, 7, gi.INFO_TYPE_STRUCT, 0)
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

// gst_bit_reader_init
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
func (v BitReader) Init(data gi.Uint8Array, size uint32) {
	iv, err := _I.Get(132, "BitReader", "init", 27, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_data, arg_size}
	iv.Call(args, nil, nil)
}

// gst_bit_reader_peek_bits_uint16
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) PeekBitsUint16(nbits uint32) (result bool, val uint16) {
	iv, err := _I.Get(133, "BitReader", "peek_bits_uint16", 27, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gst_bit_reader_peek_bits_uint32
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) PeekBitsUint32(nbits uint32) (result bool, val uint32) {
	iv, err := _I.Get(134, "BitReader", "peek_bits_uint32", 27, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_bit_reader_peek_bits_uint64
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) PeekBitsUint64(nbits uint32) (result bool, val uint64) {
	iv, err := _I.Get(135, "BitReader", "peek_bits_uint64", 27, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_bit_reader_peek_bits_uint8
//
// [ val ] trans: everything, dir: out
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) PeekBitsUint8(nbits uint32) (result bool, val uint8) {
	iv, err := _I.Get(136, "BitReader", "peek_bits_uint8", 27, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_val, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint8()
	result = ret.Bool()
	return
}

// gst_bit_reader_set_pos
//
// [ pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) SetPos(pos uint32) (result bool) {
	iv, err := _I.Get(137, "BitReader", "set_pos", 27, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint32Argument(pos)
	args := []gi.Argument{arg_v, arg_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_bit_reader_skip
//
// [ nbits ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BitReader) Skip(nbits uint32) (result bool) {
	iv, err := _I.Get(138, "BitReader", "skip", 27, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbits := gi.NewUint32Argument(nbits)
	args := []gi.Argument{arg_v, arg_nbits}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_bit_reader_skip_to_byte
//
// [ result ] trans: nothing
//
func (v BitReader) SkipToByte() (result bool) {
	iv, err := _I.Get(139, "BitReader", "skip_to_byte", 27, 15, gi.INFO_TYPE_STRUCT, 0)
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

// Struct ByteReader
type ByteReader struct {
	P unsafe.Pointer
}

const SizeOfStructByteReader = 48

func ByteReaderGetType() gi.GType {
	ret := _I.GetGType(17, "ByteReader")
	return ret
}

// gst_byte_reader_dup_data
//
// [ size ] trans: everything, dir: out
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) DupData() (result bool, val gi.Uint8Array) {
	iv, err := _I.Get(140, "ByteReader", "dup_data", 28, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_size, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint32
	_ = size
	size = outArgs[0].Uint32()
	val.P = outArgs[1].Pointer()
	result = ret.Bool()
	val.Len = int(size)
	return
}

// gst_byte_reader_dup_string_utf16
//
// [ str ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) DupStringUtf16() (result bool, str gi.Uint16Array) {
	iv, err := _I.Get(141, "ByteReader", "dup_string_utf16", 28, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	str.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_byte_reader_dup_string_utf32
//
// [ str ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) DupStringUtf32() (result bool, str gi.Uint32Array) {
	iv, err := _I.Get(142, "ByteReader", "dup_string_utf32", 28, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	str.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_byte_reader_dup_string_utf8
//
// [ str ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) DupStringUtf8() (result bool, str gi.CStrArray) {
	iv, err := _I.Get(143, "ByteReader", "dup_string_utf8", 28, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	str.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_byte_reader_free
//
func (v ByteReader) Free() {
	iv, err := _I.Get(144, "ByteReader", "free", 28, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_byte_reader_get_data
//
// [ size ] trans: everything, dir: out
//
// [ val ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetData() (result bool, val gi.Uint8Array) {
	iv, err := _I.Get(145, "ByteReader", "get_data", 28, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_size, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint32
	_ = size
	size = outArgs[0].Uint32()
	val.P = outArgs[1].Pointer()
	result = ret.Bool()
	val.Len = int(size)
	return
}

// gst_byte_reader_get_float32_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetFloat32Be() (result bool, val float32) {
	iv, err := _I.Get(146, "ByteReader", "get_float32_be", 28, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Float()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_float32_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetFloat32Le() (result bool, val float32) {
	iv, err := _I.Get(147, "ByteReader", "get_float32_le", 28, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Float()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_float64_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetFloat64Be() (result bool, val float64) {
	iv, err := _I.Get(148, "ByteReader", "get_float64_be", 28, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_float64_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetFloat64Le() (result bool, val float64) {
	iv, err := _I.Get(149, "ByteReader", "get_float64_le", 28, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int16_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt16Be() (result bool, val int16) {
	iv, err := _I.Get(150, "ByteReader", "get_int16_be", 28, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int16()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int16_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt16Le() (result bool, val int16) {
	iv, err := _I.Get(151, "ByteReader", "get_int16_le", 28, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int16()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int24_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt24Be() (result bool, val int32) {
	iv, err := _I.Get(152, "ByteReader", "get_int24_be", 28, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int24_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt24Le() (result bool, val int32) {
	iv, err := _I.Get(153, "ByteReader", "get_int24_le", 28, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int32_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt32Be() (result bool, val int32) {
	iv, err := _I.Get(154, "ByteReader", "get_int32_be", 28, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int32_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt32Le() (result bool, val int32) {
	iv, err := _I.Get(155, "ByteReader", "get_int32_le", 28, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int64_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt64Be() (result bool, val int64) {
	iv, err := _I.Get(156, "ByteReader", "get_int64_be", 28, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int64_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt64Le() (result bool, val int64) {
	iv, err := _I.Get(157, "ByteReader", "get_int64_le", 28, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_int8
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetInt8() (result bool, val int8) {
	iv, err := _I.Get(158, "ByteReader", "get_int8", 28, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int8()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_pos
//
// [ result ] trans: nothing
//
func (v ByteReader) GetPos() (result uint32) {
	iv, err := _I.Get(159, "ByteReader", "get_pos", 28, 19, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_reader_get_remaining
//
// [ result ] trans: nothing
//
func (v ByteReader) GetRemaining() (result uint32) {
	iv, err := _I.Get(160, "ByteReader", "get_remaining", 28, 20, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_reader_get_size
//
// [ result ] trans: nothing
//
func (v ByteReader) GetSize() (result uint32) {
	iv, err := _I.Get(161, "ByteReader", "get_size", 28, 21, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_reader_get_string_utf8
//
// [ str ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetStringUtf8() (result bool, str gi.CStrArray) {
	iv, err := _I.Get(162, "ByteReader", "get_string_utf8", 28, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	str.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint16_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint16Be() (result bool, val uint16) {
	iv, err := _I.Get(163, "ByteReader", "get_uint16_be", 28, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint16_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint16Le() (result bool, val uint16) {
	iv, err := _I.Get(164, "ByteReader", "get_uint16_le", 28, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint24_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint24Be() (result bool, val uint32) {
	iv, err := _I.Get(165, "ByteReader", "get_uint24_be", 28, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint24_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint24Le() (result bool, val uint32) {
	iv, err := _I.Get(166, "ByteReader", "get_uint24_le", 28, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint32_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint32Be() (result bool, val uint32) {
	iv, err := _I.Get(167, "ByteReader", "get_uint32_be", 28, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint32_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint32Le() (result bool, val uint32) {
	iv, err := _I.Get(168, "ByteReader", "get_uint32_le", 28, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint64_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint64Be() (result bool, val uint64) {
	iv, err := _I.Get(169, "ByteReader", "get_uint64_be", 28, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint64_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint64Le() (result bool, val uint64) {
	iv, err := _I.Get(170, "ByteReader", "get_uint64_le", 28, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_byte_reader_get_uint8
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) GetUint8() (result bool, val uint8) {
	iv, err := _I.Get(171, "ByteReader", "get_uint8", 28, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint8()
	result = ret.Bool()
	return
}

// gst_byte_reader_init
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
func (v ByteReader) Init(data gi.Uint8Array, size uint32) {
	iv, err := _I.Get(172, "ByteReader", "init", 28, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_data, arg_size}
	iv.Call(args, nil, nil)
}

// gst_byte_reader_masked_scan_uint32
//
// [ mask ] trans: nothing
//
// [ pattern ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteReader) MaskedScanUint32(mask uint32, pattern uint32, offset uint32, size uint32) (result uint32) {
	iv, err := _I.Get(173, "ByteReader", "masked_scan_uint32", 28, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewUint32Argument(mask)
	arg_pattern := gi.NewUint32Argument(pattern)
	arg_offset := gi.NewUint32Argument(offset)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_mask, arg_pattern, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_byte_reader_masked_scan_uint32_peek
//
// [ mask ] trans: nothing
//
// [ pattern ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) MaskedScanUint32Peek(mask uint32, pattern uint32, offset uint32, size uint32) (result uint32, value uint32) {
	iv, err := _I.Get(174, "ByteReader", "masked_scan_uint32_peek", 28, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewUint32Argument(mask)
	arg_pattern := gi.NewUint32Argument(pattern)
	arg_offset := gi.NewUint32Argument(offset)
	arg_size := gi.NewUint32Argument(size)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_mask, arg_pattern, arg_offset, arg_size, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	value = outArgs[0].Uint32()
	result = ret.Uint32()
	return
}

// gst_byte_reader_peek_data
//
// [ size ] trans: everything, dir: out
//
// [ val ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekData() (result bool, val gi.Uint8Array) {
	iv, err := _I.Get(175, "ByteReader", "peek_data", 28, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_size, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint32
	_ = size
	size = outArgs[0].Uint32()
	val.P = outArgs[1].Pointer()
	result = ret.Bool()
	val.Len = int(size)
	return
}

// gst_byte_reader_peek_float32_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekFloat32Be() (result bool, val float32) {
	iv, err := _I.Get(176, "ByteReader", "peek_float32_be", 28, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Float()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_float32_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekFloat32Le() (result bool, val float32) {
	iv, err := _I.Get(177, "ByteReader", "peek_float32_le", 28, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Float()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_float64_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekFloat64Be() (result bool, val float64) {
	iv, err := _I.Get(178, "ByteReader", "peek_float64_be", 28, 38, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_float64_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekFloat64Le() (result bool, val float64) {
	iv, err := _I.Get(179, "ByteReader", "peek_float64_le", 28, 39, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int16_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt16Be() (result bool, val int16) {
	iv, err := _I.Get(180, "ByteReader", "peek_int16_be", 28, 40, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int16()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int16_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt16Le() (result bool, val int16) {
	iv, err := _I.Get(181, "ByteReader", "peek_int16_le", 28, 41, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int16()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int24_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt24Be() (result bool, val int32) {
	iv, err := _I.Get(182, "ByteReader", "peek_int24_be", 28, 42, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int24_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt24Le() (result bool, val int32) {
	iv, err := _I.Get(183, "ByteReader", "peek_int24_le", 28, 43, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int32_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt32Be() (result bool, val int32) {
	iv, err := _I.Get(184, "ByteReader", "peek_int32_be", 28, 44, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int32_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt32Le() (result bool, val int32) {
	iv, err := _I.Get(185, "ByteReader", "peek_int32_le", 28, 45, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int64_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt64Be() (result bool, val int64) {
	iv, err := _I.Get(186, "ByteReader", "peek_int64_be", 28, 46, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int64_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt64Le() (result bool, val int64) {
	iv, err := _I.Get(187, "ByteReader", "peek_int64_le", 28, 47, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_int8
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekInt8() (result bool, val int8) {
	iv, err := _I.Get(188, "ByteReader", "peek_int8", 28, 48, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Int8()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_string_utf8
//
// [ str ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekStringUtf8() (result bool, str gi.CStrArray) {
	iv, err := _I.Get(189, "ByteReader", "peek_string_utf8", 28, 49, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	str.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint16_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint16Be() (result bool, val uint16) {
	iv, err := _I.Get(190, "ByteReader", "peek_uint16_be", 28, 50, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint16_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint16Le() (result bool, val uint16) {
	iv, err := _I.Get(191, "ByteReader", "peek_uint16_le", 28, 51, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint24_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint24Be() (result bool, val uint32) {
	iv, err := _I.Get(192, "ByteReader", "peek_uint24_be", 28, 52, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint24_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint24Le() (result bool, val uint32) {
	iv, err := _I.Get(193, "ByteReader", "peek_uint24_le", 28, 53, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint32_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint32Be() (result bool, val uint32) {
	iv, err := _I.Get(194, "ByteReader", "peek_uint32_be", 28, 54, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint32_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint32Le() (result bool, val uint32) {
	iv, err := _I.Get(195, "ByteReader", "peek_uint32_le", 28, 55, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint64_be
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint64Be() (result bool, val uint64) {
	iv, err := _I.Get(196, "ByteReader", "peek_uint64_be", 28, 56, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint64_le
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint64Le() (result bool, val uint64) {
	iv, err := _I.Get(197, "ByteReader", "peek_uint64_le", 28, 57, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_byte_reader_peek_uint8
//
// [ val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ByteReader) PeekUint8() (result bool, val uint8) {
	iv, err := _I.Get(198, "ByteReader", "peek_uint8", 28, 58, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	val = outArgs[0].Uint8()
	result = ret.Bool()
	return
}

// gst_byte_reader_set_pos
//
// [ pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteReader) SetPos(pos uint32) (result bool) {
	iv, err := _I.Get(199, "ByteReader", "set_pos", 28, 59, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint32Argument(pos)
	args := []gi.Argument{arg_v, arg_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_reader_skip
//
// [ nbytes ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteReader) Skip(nbytes uint32) (result bool) {
	iv, err := _I.Get(200, "ByteReader", "skip", 28, 60, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nbytes := gi.NewUint32Argument(nbytes)
	args := []gi.Argument{arg_v, arg_nbytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_reader_skip_string_utf16
//
// [ result ] trans: nothing
//
func (v ByteReader) SkipStringUtf16() (result bool) {
	iv, err := _I.Get(201, "ByteReader", "skip_string_utf16", 28, 61, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_reader_skip_string_utf32
//
// [ result ] trans: nothing
//
func (v ByteReader) SkipStringUtf32() (result bool) {
	iv, err := _I.Get(202, "ByteReader", "skip_string_utf32", 28, 62, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_reader_skip_string_utf8
//
// [ result ] trans: nothing
//
func (v ByteReader) SkipStringUtf8() (result bool) {
	iv, err := _I.Get(203, "ByteReader", "skip_string_utf8", 28, 63, gi.INFO_TYPE_STRUCT, 0)
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

// Struct ByteWriter
type ByteWriter struct {
	P unsafe.Pointer
}

const SizeOfStructByteWriter = 96

func ByteWriterGetType() gi.GType {
	ret := _I.GetGType(18, "ByteWriter")
	return ret
}

// gst_byte_writer_ensure_free_space
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) EnsureFreeSpace(size uint32) (result bool) {
	iv, err := _I.Get(204, "ByteWriter", "ensure_free_space", 29, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_fill
//
// [ value ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) Fill(value uint8, size uint32) (result bool) {
	iv, err := _I.Get(205, "ByteWriter", "fill", 29, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewUint8Argument(value)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_value, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_free
//
func (v ByteWriter) Free() {
	iv, err := _I.Get(206, "ByteWriter", "free", 29, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_byte_writer_free_and_get_buffer
//
// [ result ] trans: everything
//
func (v ByteWriter) FreeAndGetBuffer() (result gst.Buffer) {
	iv, err := _I.Get(207, "ByteWriter", "free_and_get_buffer", 29, 3, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_writer_free_and_get_data
//
// [ result ] trans: everything
//
func (v ByteWriter) FreeAndGetData() (result uint8) {
	iv, err := _I.Get(208, "ByteWriter", "free_and_get_data", 29, 4, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_writer_get_remaining
//
// [ result ] trans: nothing
//
func (v ByteWriter) GetRemaining() (result uint32) {
	iv, err := _I.Get(209, "ByteWriter", "get_remaining", 29, 5, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_writer_init
//
func (v ByteWriter) Init() {
	iv, err := _I.Get(210, "ByteWriter", "init", 29, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_byte_writer_init_with_data
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ initialized ] trans: nothing
//
func (v ByteWriter) InitWithData(data gi.Uint8Array, size uint32, initialized bool) {
	iv, err := _I.Get(211, "ByteWriter", "init_with_data", 29, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	arg_initialized := gi.NewBoolArgument(initialized)
	args := []gi.Argument{arg_v, arg_data, arg_size, arg_initialized}
	iv.Call(args, nil, nil)
}

// gst_byte_writer_init_with_size
//
// [ size ] trans: nothing
//
// [ fixed ] trans: nothing
//
func (v ByteWriter) InitWithSize(size uint32, fixed bool) {
	iv, err := _I.Get(212, "ByteWriter", "init_with_size", 29, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewUint32Argument(size)
	arg_fixed := gi.NewBoolArgument(fixed)
	args := []gi.Argument{arg_v, arg_size, arg_fixed}
	iv.Call(args, nil, nil)
}

// gst_byte_writer_put_buffer
//
// [ buffer ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutBuffer(buffer gst.Buffer, offset uint64, size int64) (result bool) {
	iv, err := _I.Get(213, "ByteWriter", "put_buffer", 29, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_buffer, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_data
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutData(data gi.Uint8Array, size uint32) (result bool) {
	iv, err := _I.Get(214, "ByteWriter", "put_data", 29, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_float32_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutFloat32Be(val float32) (result bool) {
	iv, err := _I.Get(215, "ByteWriter", "put_float32_be", 29, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewFloatArgument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_float32_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutFloat32Le(val float32) (result bool) {
	iv, err := _I.Get(216, "ByteWriter", "put_float32_le", 29, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewFloatArgument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_float64_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutFloat64Be(val float64) (result bool) {
	iv, err := _I.Get(217, "ByteWriter", "put_float64_be", 29, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewDoubleArgument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_float64_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutFloat64Le(val float64) (result bool) {
	iv, err := _I.Get(218, "ByteWriter", "put_float64_le", 29, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewDoubleArgument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int16_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt16Be(val int16) (result bool) {
	iv, err := _I.Get(219, "ByteWriter", "put_int16_be", 29, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt16Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int16_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt16Le(val int16) (result bool) {
	iv, err := _I.Get(220, "ByteWriter", "put_int16_le", 29, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt16Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int24_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt24Be(val int32) (result bool) {
	iv, err := _I.Get(221, "ByteWriter", "put_int24_be", 29, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int24_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt24Le(val int32) (result bool) {
	iv, err := _I.Get(222, "ByteWriter", "put_int24_le", 29, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int32_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt32Be(val int32) (result bool) {
	iv, err := _I.Get(223, "ByteWriter", "put_int32_be", 29, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int32_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt32Le(val int32) (result bool) {
	iv, err := _I.Get(224, "ByteWriter", "put_int32_le", 29, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int64_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt64Be(val int64) (result bool) {
	iv, err := _I.Get(225, "ByteWriter", "put_int64_be", 29, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt64Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int64_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt64Le(val int64) (result bool) {
	iv, err := _I.Get(226, "ByteWriter", "put_int64_le", 29, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt64Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_int8
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutInt8(val int8) (result bool) {
	iv, err := _I.Get(227, "ByteWriter", "put_int8", 29, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewInt8Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_string_utf16
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutStringUtf16(data gi.Uint16Array) (result bool) {
	iv, err := _I.Get(228, "ByteWriter", "put_string_utf16", 29, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_string_utf32
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutStringUtf32(data gi.Uint32Array) (result bool) {
	iv, err := _I.Get(229, "ByteWriter", "put_string_utf32", 29, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_string_utf8
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutStringUtf8(data string) (result bool) {
	iv, err := _I.Get(230, "ByteWriter", "put_string_utf8", 29, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_data := gi.CString(data)
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewStringArgument(c_data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_data)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint16_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint16Be(val uint16) (result bool) {
	iv, err := _I.Get(231, "ByteWriter", "put_uint16_be", 29, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint16Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint16_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint16Le(val uint16) (result bool) {
	iv, err := _I.Get(232, "ByteWriter", "put_uint16_le", 29, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint16Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint24_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint24Be(val uint32) (result bool) {
	iv, err := _I.Get(233, "ByteWriter", "put_uint24_be", 29, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint24_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint24Le(val uint32) (result bool) {
	iv, err := _I.Get(234, "ByteWriter", "put_uint24_le", 29, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint32_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint32Be(val uint32) (result bool) {
	iv, err := _I.Get(235, "ByteWriter", "put_uint32_be", 29, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint32_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint32Le(val uint32) (result bool) {
	iv, err := _I.Get(236, "ByteWriter", "put_uint32_le", 29, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint64_be
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint64Be(val uint64) (result bool) {
	iv, err := _I.Get(237, "ByteWriter", "put_uint64_be", 29, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint64_le
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint64Le(val uint64) (result bool) {
	iv, err := _I.Get(238, "ByteWriter", "put_uint64_le", 29, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_put_uint8
//
// [ val ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ByteWriter) PutUint8(val uint8) (result bool) {
	iv, err := _I.Get(239, "ByteWriter", "put_uint8", 29, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewUint8Argument(val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_byte_writer_reset
//
func (v ByteWriter) Reset() {
	iv, err := _I.Get(240, "ByteWriter", "reset", 29, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_byte_writer_reset_and_get_buffer
//
// [ result ] trans: everything
//
func (v ByteWriter) ResetAndGetBuffer() (result gst.Buffer) {
	iv, err := _I.Get(241, "ByteWriter", "reset_and_get_buffer", 29, 37, gi.INFO_TYPE_STRUCT, 0)
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

// gst_byte_writer_reset_and_get_data
//
// [ result ] trans: everything
//
func (v ByteWriter) ResetAndGetData() (result gi.Uint8Array) {
	iv, err := _I.Get(242, "ByteWriter", "reset_and_get_data", 29, 38, gi.INFO_TYPE_STRUCT, 0)
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

// Struct CollectData
type CollectData struct {
	P unsafe.Pointer
}

const SizeOfStructCollectData = 168

func CollectDataGetType() gi.GType {
	ret := _I.GetGType(19, "CollectData")
	return ret
}

type CollectDataDestroyNotifyStruct struct {
	F_data CollectData
}

func GetPointer_myCollectDataDestroyNotify() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectDataDestroyNotify())
}

//export myGstBaseCollectDataDestroyNotify
func myGstBaseCollectDataDestroyNotify(data *C.GstCollectData) {
	// TODO: not found user_data
}

// Struct CollectDataPrivate
type CollectDataPrivate struct {
	P unsafe.Pointer
}

func CollectDataPrivateGetType() gi.GType {
	ret := _I.GetGType(20, "CollectDataPrivate")
	return ret
}

// Object CollectPads
type CollectPads struct {
	gst.Object
}

func WrapCollectPads(p unsafe.Pointer) (r CollectPads) { r.P = p; return }

type ICollectPads interface{ P_CollectPads() unsafe.Pointer }

func (v CollectPads) P_CollectPads() unsafe.Pointer { return v.P }
func CollectPadsGetType() gi.GType {
	ret := _I.GetGType(21, "CollectPads")
	return ret
}

// gst_collect_pads_new
//
// [ result ] trans: everything
//
func NewCollectPads() (result CollectPads) {
	iv, err := _I.Get(243, "CollectPads", "new", 33, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_collect_pads_add_pad
//
// [ pad ] trans: nothing
//
// [ size ] trans: nothing
//
// [ destroy_notify ] trans: nothing
//
// [ lock ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) AddPad(pad gst.IPad, size uint32, destroy_notify int /*TODO_TYPE CALLBACK*/, lock bool) (result CollectData) {
	iv, err := _I.Get(244, "CollectPads", "add_pad", 33, 1, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pad := gi.NewPointerArgument(tmp)
	arg_size := gi.NewUint32Argument(size)
	arg_destroy_notify := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectDataDestroyNotify()))
	arg_lock := gi.NewBoolArgument(lock)
	args := []gi.Argument{arg_v, arg_pad, arg_size, arg_destroy_notify, arg_lock}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_collect_pads_available
//
// [ result ] trans: nothing
//
func (v CollectPads) Available() (result uint32) {
	iv, err := _I.Get(245, "CollectPads", "available", 33, 2, gi.INFO_TYPE_OBJECT, 0)
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

// gst_collect_pads_clip_running_time
//
// [ cdata ] trans: nothing
//
// [ buf ] trans: nothing
//
// [ outbuf ] trans: everything, dir: out
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) ClipRunningTime(cdata CollectData, buf gst.Buffer, user_data unsafe.Pointer) (result gst.FlowReturnEnum, outbuf gst.Buffer) {
	iv, err := _I.Get(246, "CollectPads", "clip_running_time", 33, 3, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_cdata := gi.NewPointerArgument(cdata.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_outbuf := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_cdata, arg_buf, arg_outbuf, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	outbuf.P = outArgs[0].Pointer()
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_collect_pads_event_default
//
// [ data ] trans: nothing
//
// [ event ] trans: nothing
//
// [ discard ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) EventDefault(data CollectData, event gst.Event, discard bool) (result bool) {
	iv, err := _I.Get(247, "CollectPads", "event_default", 33, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_event := gi.NewPointerArgument(event.P)
	arg_discard := gi.NewBoolArgument(discard)
	args := []gi.Argument{arg_v, arg_data, arg_event, arg_discard}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_collect_pads_flush
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) Flush(data CollectData, size uint32) (result uint32) {
	iv, err := _I.Get(248, "CollectPads", "flush", 33, 5, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_collect_pads_peek
//
// [ data ] trans: nothing
//
// [ result ] trans: everything
//
func (v CollectPads) Peek(data CollectData) (result gst.Buffer) {
	iv, err := _I.Get(249, "CollectPads", "peek", 33, 6, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_collect_pads_pop
//
// [ data ] trans: nothing
//
// [ result ] trans: everything
//
func (v CollectPads) Pop(data CollectData) (result gst.Buffer) {
	iv, err := _I.Get(250, "CollectPads", "pop", 33, 7, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_collect_pads_query_default
//
// [ data ] trans: nothing
//
// [ query ] trans: nothing
//
// [ discard ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) QueryDefault(data CollectData, query gst.Query, discard bool) (result bool) {
	iv, err := _I.Get(251, "CollectPads", "query_default", 33, 8, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_query := gi.NewPointerArgument(query.P)
	arg_discard := gi.NewBoolArgument(discard)
	args := []gi.Argument{arg_v, arg_data, arg_query, arg_discard}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_collect_pads_read_buffer
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v CollectPads) ReadBuffer(data CollectData, size uint32) (result gst.Buffer) {
	iv, err := _I.Get(252, "CollectPads", "read_buffer", 33, 9, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_collect_pads_remove_pad
//
// [ pad ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) RemovePad(pad gst.IPad) (result bool) {
	iv, err := _I.Get(253, "CollectPads", "remove_pad", 33, 10, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pad := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pad}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_collect_pads_set_buffer_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetBufferFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(254, "CollectPads", "set_buffer_function", 33, 11, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsBufferFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_clip_function
//
// [ clipfunc ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetClipFunction(clipfunc int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(255, "CollectPads", "set_clip_function", 33, 12, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_clipfunc := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsClipFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_clipfunc, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_compare_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetCompareFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(256, "CollectPads", "set_compare_function", 33, 13, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsCompareFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_event_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetEventFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(257, "CollectPads", "set_event_function", 33, 14, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsEventFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_flush_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetFlushFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(258, "CollectPads", "set_flush_function", 33, 15, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsFlushFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_flushing
//
// [ flushing ] trans: nothing
//
func (v CollectPads) SetFlushing(flushing bool) {
	iv, err := _I.Get(259, "CollectPads", "set_flushing", 33, 16, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flushing := gi.NewBoolArgument(flushing)
	args := []gi.Argument{arg_v, arg_flushing}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(260, "CollectPads", "set_function", 33, 17, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_query_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v CollectPads) SetQueryFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(261, "CollectPads", "set_query_function", 33, 18, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCollectPadsQueryFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_set_waiting
//
// [ data ] trans: nothing
//
// [ waiting ] trans: nothing
//
func (v CollectPads) SetWaiting(data CollectData, waiting bool) {
	iv, err := _I.Get(262, "CollectPads", "set_waiting", 33, 19, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_waiting := gi.NewBoolArgument(waiting)
	args := []gi.Argument{arg_v, arg_data, arg_waiting}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_src_event_default
//
// [ pad ] trans: nothing
//
// [ event ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CollectPads) SrcEventDefault(pad gst.IPad, event gst.Event) (result bool) {
	iv, err := _I.Get(263, "CollectPads", "src_event_default", 33, 20, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pad := gi.NewPointerArgument(tmp)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_pad, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_collect_pads_start
//
func (v CollectPads) Start() {
	iv, err := _I.Get(264, "CollectPads", "start", 33, 21, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_stop
//
func (v CollectPads) Stop() {
	iv, err := _I.Get(265, "CollectPads", "stop", 33, 22, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_collect_pads_take_buffer
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v CollectPads) TakeBuffer(data CollectData, size uint32) (result gst.Buffer) {
	iv, err := _I.Get(266, "CollectPads", "take_buffer", 33, 23, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_v, arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

type CollectPadsBufferFunctionStruct struct {
	F_pads   CollectPads
	F_data   CollectData
	F_buffer gst.Buffer
}

func GetPointer_myCollectPadsBufferFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsBufferFunction())
}

//export myGstBaseCollectPadsBufferFunction
func myGstBaseCollectPadsBufferFunction(pads *C.GstCollectPads, data *C.GstCollectData, buffer *C.GstBuffer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsBufferFunctionStruct{
		F_pads:   WrapCollectPads(unsafe.Pointer(pads)),
		F_data:   CollectData{P: unsafe.Pointer(data)},
		F_buffer: gst.Buffer{P: unsafe.Pointer(buffer)},
	}
	fn(args)
}

// ignore GType struct CollectPadsClass

type CollectPadsClipFunctionStruct struct {
	F_pads      CollectPads
	F_data      CollectData
	F_inbuffer  gst.Buffer
	F_outbuffer unsafe.Pointer /*TODO_CB tag: interface, isPtr: true*/
}

func GetPointer_myCollectPadsClipFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsClipFunction())
}

//export myGstBaseCollectPadsClipFunction
func myGstBaseCollectPadsClipFunction(pads *C.GstCollectPads, data *C.GstCollectData, inbuffer *C.GstBuffer, outbuffer C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsClipFunctionStruct{
		F_pads:      WrapCollectPads(unsafe.Pointer(pads)),
		F_data:      CollectData{P: unsafe.Pointer(data)},
		F_inbuffer:  gst.Buffer{P: unsafe.Pointer(inbuffer)},
		F_outbuffer: unsafe.Pointer(outbuffer),
	}
	fn(args)
}

type CollectPadsCompareFunctionStruct struct {
	F_pads       CollectPads
	F_data1      CollectData
	F_timestamp1 uint64
	F_data2      CollectData
	F_timestamp2 uint64
}

func GetPointer_myCollectPadsCompareFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsCompareFunction())
}

//export myGstBaseCollectPadsCompareFunction
func myGstBaseCollectPadsCompareFunction(pads *C.GstCollectPads, data1 *C.GstCollectData, timestamp1 C.guint64, data2 *C.GstCollectData, timestamp2 C.guint64, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsCompareFunctionStruct{
		F_pads:       WrapCollectPads(unsafe.Pointer(pads)),
		F_data1:      CollectData{P: unsafe.Pointer(data1)},
		F_timestamp1: uint64(timestamp1),
		F_data2:      CollectData{P: unsafe.Pointer(data2)},
		F_timestamp2: uint64(timestamp2),
	}
	fn(args)
}

type CollectPadsEventFunctionStruct struct {
	F_pads  CollectPads
	F_pad   CollectData
	F_event gst.Event
}

func GetPointer_myCollectPadsEventFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsEventFunction())
}

//export myGstBaseCollectPadsEventFunction
func myGstBaseCollectPadsEventFunction(pads *C.GstCollectPads, pad *C.GstCollectData, event *C.GstEvent, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsEventFunctionStruct{
		F_pads:  WrapCollectPads(unsafe.Pointer(pads)),
		F_pad:   CollectData{P: unsafe.Pointer(pad)},
		F_event: gst.Event{P: unsafe.Pointer(event)},
	}
	fn(args)
}

type CollectPadsFlushFunctionStruct struct {
	F_pads CollectPads
}

func GetPointer_myCollectPadsFlushFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsFlushFunction())
}

//export myGstBaseCollectPadsFlushFunction
func myGstBaseCollectPadsFlushFunction(pads *C.GstCollectPads, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsFlushFunctionStruct{
		F_pads: WrapCollectPads(unsafe.Pointer(pads)),
	}
	fn(args)
}

type CollectPadsFunctionStruct struct {
	F_pads CollectPads
}

func GetPointer_myCollectPadsFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsFunction())
}

//export myGstBaseCollectPadsFunction
func myGstBaseCollectPadsFunction(pads *C.GstCollectPads, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsFunctionStruct{
		F_pads: WrapCollectPads(unsafe.Pointer(pads)),
	}
	fn(args)
}

// Struct CollectPadsPrivate
type CollectPadsPrivate struct {
	P unsafe.Pointer
}

func CollectPadsPrivateGetType() gi.GType {
	ret := _I.GetGType(22, "CollectPadsPrivate")
	return ret
}

type CollectPadsQueryFunctionStruct struct {
	F_pads  CollectPads
	F_pad   CollectData
	F_query gst.Query
}

func GetPointer_myCollectPadsQueryFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseCollectPadsQueryFunction())
}

//export myGstBaseCollectPadsQueryFunction
func myGstBaseCollectPadsQueryFunction(pads *C.GstCollectPads, pad *C.GstCollectData, query *C.GstQuery, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CollectPadsQueryFunctionStruct{
		F_pads:  WrapCollectPads(unsafe.Pointer(pads)),
		F_pad:   CollectData{P: unsafe.Pointer(pad)},
		F_query: gst.Query{P: unsafe.Pointer(query)},
	}
	fn(args)
}

// Flags CollectPadsStateFlags
type CollectPadsStateFlags int

const (
	CollectPadsStateFlagsEos        CollectPadsStateFlags = 1
	CollectPadsStateFlagsFlushing   CollectPadsStateFlags = 2
	CollectPadsStateFlagsNewSegment CollectPadsStateFlags = 4
	CollectPadsStateFlagsWaiting    CollectPadsStateFlags = 8
	CollectPadsStateFlagsLocked     CollectPadsStateFlags = 16
)

func CollectPadsStateFlagsGetType() gi.GType {
	ret := _I.GetGType(23, "CollectPadsStateFlags")
	return ret
}

// Object DataQueue
type DataQueue struct {
	g.Object
}

func WrapDataQueue(p unsafe.Pointer) (r DataQueue) { r.P = p; return }

type IDataQueue interface{ P_DataQueue() unsafe.Pointer }

func (v DataQueue) P_DataQueue() unsafe.Pointer { return v.P }
func DataQueueGetType() gi.GType {
	ret := _I.GetGType(24, "DataQueue")
	return ret
}

// ignore GType struct DataQueueClass

type DataQueueEmptyCallbackStruct struct {
	F_queue     DataQueue
	F_checkdata unsafe.Pointer
}

func GetPointer_myDataQueueEmptyCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseDataQueueEmptyCallback())
}

//export myGstBaseDataQueueEmptyCallback
func myGstBaseDataQueueEmptyCallback(queue *C.GstDataQueue, checkdata C.gpointer) {
	// TODO: not found user_data
}

type DataQueueFullCallbackStruct struct {
	F_queue     DataQueue
	F_checkdata unsafe.Pointer
}

func GetPointer_myDataQueueFullCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseDataQueueFullCallback())
}

//export myGstBaseDataQueueFullCallback
func myGstBaseDataQueueFullCallback(queue *C.GstDataQueue, checkdata C.gpointer) {
	// TODO: not found user_data
}

// Struct DataQueuePrivate
type DataQueuePrivate struct {
	P unsafe.Pointer
}

func DataQueuePrivateGetType() gi.GType {
	ret := _I.GetGType(25, "DataQueuePrivate")
	return ret
}

// Struct FlowCombiner
type FlowCombiner struct {
	P unsafe.Pointer
}

func FlowCombinerGetType() gi.GType {
	ret := _I.GetGType(26, "FlowCombiner")
	return ret
}

// gst_flow_combiner_new
//
// [ result ] trans: everything
//
func NewFlowCombiner() (result FlowCombiner) {
	iv, err := _I.Get(267, "FlowCombiner", "new", 49, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_flow_combiner_add_pad
//
// [ pad ] trans: nothing
//
func (v FlowCombiner) AddPad(pad gst.IPad) {
	iv, err := _I.Get(268, "FlowCombiner", "add_pad", 49, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pad := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pad}
	iv.Call(args, nil, nil)
}

// gst_flow_combiner_clear
//
func (v FlowCombiner) Clear() {
	iv, err := _I.Get(269, "FlowCombiner", "clear", 49, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_flow_combiner_free
//
func (v FlowCombiner) Free() {
	iv, err := _I.Get(270, "FlowCombiner", "free", 49, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_flow_combiner_ref
//
// [ result ] trans: everything
//
func (v FlowCombiner) Ref() (result FlowCombiner) {
	iv, err := _I.Get(271, "FlowCombiner", "ref", 49, 4, gi.INFO_TYPE_STRUCT, 0)
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

// gst_flow_combiner_remove_pad
//
// [ pad ] trans: nothing
//
func (v FlowCombiner) RemovePad(pad gst.IPad) {
	iv, err := _I.Get(272, "FlowCombiner", "remove_pad", 49, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pad := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pad}
	iv.Call(args, nil, nil)
}

// gst_flow_combiner_reset
//
func (v FlowCombiner) Reset() {
	iv, err := _I.Get(273, "FlowCombiner", "reset", 49, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_flow_combiner_unref
//
func (v FlowCombiner) Unref() {
	iv, err := _I.Get(274, "FlowCombiner", "unref", 49, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_flow_combiner_update_flow
//
// [ fret ] trans: nothing
//
// [ result ] trans: nothing
//
func (v FlowCombiner) UpdateFlow(fret gst.FlowReturnEnum) (result gst.FlowReturnEnum) {
	iv, err := _I.Get(275, "FlowCombiner", "update_flow", 49, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fret := gi.NewIntArgument(int(fret))
	args := []gi.Argument{arg_v, arg_fret}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// gst_flow_combiner_update_pad_flow
//
// [ pad ] trans: nothing
//
// [ fret ] trans: nothing
//
// [ result ] trans: nothing
//
func (v FlowCombiner) UpdatePadFlow(pad gst.IPad, fret gst.FlowReturnEnum) (result gst.FlowReturnEnum) {
	iv, err := _I.Get(276, "FlowCombiner", "update_pad_flow", 49, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pad := gi.NewPointerArgument(tmp)
	arg_fret := gi.NewIntArgument(int(fret))
	args := []gi.Argument{arg_v, arg_pad, arg_fret}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// Object PushSrc
type PushSrc struct {
	BaseSrc
}

func WrapPushSrc(p unsafe.Pointer) (r PushSrc) { r.P = p; return }

type IPushSrc interface{ P_PushSrc() unsafe.Pointer }

func (v PushSrc) P_PushSrc() unsafe.Pointer { return v.P }
func PushSrcGetType() gi.GType {
	ret := _I.GetGType(27, "PushSrc")
	return ret
}

// ignore GType struct PushSrcClass

type TypeFindHelperGetRangeFunctionStruct struct {
	F_obj    gst.Object
	F_parent gst.Object
	F_offset uint64
	F_length uint32
	F_buffer unsafe.Pointer /*TODO_CB tag: interface, isPtr: true*/
}

func GetPointer_myTypeFindHelperGetRangeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBaseTypeFindHelperGetRangeFunction())
}

//export myGstBaseTypeFindHelperGetRangeFunction
func myGstBaseTypeFindHelperGetRangeFunction(obj *C.GstObject, parent *C.GstObject, offset C.guint64, length C.guint32, buffer C.gpointer) {
	// TODO: not found user_data
}

// gst_type_find_helper
//
// [ src ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func TypeFindHelper(src gst.IPad, size uint64) (result gst.Caps) {
	iv, err := _I.Get(277, "type_find_helper", "", 53, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Pad()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_src, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_type_find_helper_for_buffer
//
// [ obj ] trans: nothing
//
// [ buf ] trans: nothing
//
// [ prob ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeFindHelperForBuffer(obj gst.IObject, buf gst.Buffer) (result gst.Caps, prob gst.TypeFindProbabilityEnum) {
	iv, err := _I.Get(278, "type_find_helper_for_buffer", "", 54, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	arg_obj := gi.NewPointerArgument(tmp)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_prob := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_obj, arg_buf, arg_prob}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	prob = gst.TypeFindProbabilityEnum(outArgs[0].Int())
	result.P = ret.Pointer()
	return
}

// gst_type_find_helper_for_data
//
// [ obj ] trans: nothing
//
// [ data ] trans: nothing
//
// [ size ] trans: nothing
//
// [ prob ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeFindHelperForData(obj gst.IObject, data gi.Uint8Array, size uint64) (result gst.Caps, prob gst.TypeFindProbabilityEnum) {
	iv, err := _I.Get(279, "type_find_helper_for_data", "", 55, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	arg_obj := gi.NewPointerArgument(tmp)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	arg_prob := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_obj, arg_data, arg_size, arg_prob}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	prob = gst.TypeFindProbabilityEnum(outArgs[0].Int())
	result.P = ret.Pointer()
	return
}

// gst_type_find_helper_for_extension
//
// [ obj ] trans: nothing
//
// [ extension ] trans: nothing
//
// [ result ] trans: everything
//
func TypeFindHelperForExtension(obj gst.IObject, extension string) (result gst.Caps) {
	iv, err := _I.Get(280, "type_find_helper_for_extension", "", 56, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	c_extension := gi.CString(extension)
	arg_obj := gi.NewPointerArgument(tmp)
	arg_extension := gi.NewStringArgument(c_extension)
	args := []gi.Argument{arg_obj, arg_extension}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_extension)
	result.P = ret.Pointer()
	return
}

// gst_type_find_helper_get_range
//
// [ obj ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ size ] trans: nothing
//
// [ extension ] trans: nothing
//
// [ prob ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeFindHelperGetRange(obj gst.IObject, parent gst.IObject, func1 int /*TODO_TYPE CALLBACK*/, size uint64, extension string) (result gst.Caps, prob gst.TypeFindProbabilityEnum) {
	iv, err := _I.Get(281, "type_find_helper_get_range", "", 57, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	c_extension := gi.CString(extension)
	arg_obj := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTypeFindHelperGetRangeFunction()))
	arg_size := gi.NewUint64Argument(size)
	arg_extension := gi.NewStringArgument(c_extension)
	arg_prob := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_obj, arg_parent, arg_func1, arg_size, arg_extension, arg_prob}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_extension)
	prob = gst.TypeFindProbabilityEnum(outArgs[0].Int())
	result.P = ret.Pointer()
	return
}

// gst_type_find_helper_get_range_full
//
// [ obj ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ size ] trans: nothing
//
// [ extension ] trans: nothing
//
// [ caps ] trans: everything, dir: out
//
// [ prob ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func TypeFindHelperGetRangeFull(obj gst.IObject, parent gst.IObject, func1 int /*TODO_TYPE CALLBACK*/, size uint64, extension string) (result gst.FlowReturnEnum, caps gst.Caps, prob gst.TypeFindProbabilityEnum) {
	iv, err := _I.Get(282, "type_find_helper_get_range_full", "", 58, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	c_extension := gi.CString(extension)
	arg_obj := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTypeFindHelperGetRangeFunction()))
	arg_size := gi.NewUint64Argument(size)
	arg_extension := gi.NewStringArgument(c_extension)
	arg_caps := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_prob := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_obj, arg_parent, arg_func1, arg_size, arg_extension, arg_caps, arg_prob}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_extension)
	caps.P = outArgs[0].Pointer()
	prob = gst.TypeFindProbabilityEnum(outArgs[1].Int())
	result = gst.FlowReturnEnum(ret.Int())
	return
}

// constants
const (
	BASE_PARSE_FLAG_DRAINING  = 2
	BASE_PARSE_FLAG_LOST_SYNC = 1
	BASE_TRANSFORM_SINK_NAME  = "sink"
	BASE_TRANSFORM_SRC_NAME   = "src"
)
