package gst

/*
#cgo pkg-config: gstreamer-1.0
#include <gst/gst.h>
extern void myGstBufferForeachMetaFunc(GstBuffer* buffer, gpointer meta, gpointer user_data);
static void* getPointer_myGstBufferForeachMetaFunc() {
return (void*)(myGstBufferForeachMetaFunc);
}
extern void myGstBufferListFunc(gpointer buffer, guint32 idx, gpointer user_data);
static void* getPointer_myGstBufferListFunc() {
return (void*)(myGstBufferListFunc);
}
extern void myGstBusFunc(GstBus* bus, GstMessage* message, gpointer user_data);
static void* getPointer_myGstBusFunc() {
return (void*)(myGstBusFunc);
}
extern void myGstBusSyncHandler(GstBus* bus, GstMessage* message, gpointer user_data);
static void* getPointer_myGstBusSyncHandler() {
return (void*)(myGstBusSyncHandler);
}
extern void myGstCapsFilterMapFunc(GstCapsFeatures* features, GstStructure* structure, gpointer user_data);
static void* getPointer_myGstCapsFilterMapFunc() {
return (void*)(myGstCapsFilterMapFunc);
}
extern void myGstCapsForeachFunc(GstCapsFeatures* features, GstStructure* structure, gpointer user_data);
static void* getPointer_myGstCapsForeachFunc() {
return (void*)(myGstCapsForeachFunc);
}
extern void myGstCapsMapFunc(GstCapsFeatures* features, GstStructure* structure, gpointer user_data);
static void* getPointer_myGstCapsMapFunc() {
return (void*)(myGstCapsMapFunc);
}
extern void myGstClockCallback(GstClock* clock, guint64 time, gpointer id, gpointer user_data);
static void* getPointer_myGstClockCallback() {
return (void*)(myGstClockCallback);
}
extern void myGstControlBindingConvert(GstControlBinding* binding, gdouble src_value, GValue* dest_value);
static void* getPointer_myGstControlBindingConvert() {
return (void*)(myGstControlBindingConvert);
}
extern void myGstControlSourceGetValue(GstControlSource* self, guint64 timestamp, gdouble* value);
static void* getPointer_myGstControlSourceGetValue() {
return (void*)(myGstControlSourceGetValue);
}
extern void myGstControlSourceGetValueArray(GstControlSource* self, guint64 timestamp, guint64 interval, guint32 n_values, gdouble* values);
static void* getPointer_myGstControlSourceGetValueArray() {
return (void*)(myGstControlSourceGetValueArray);
}
extern void myGstDebugFuncPtr();
static void* getPointer_myGstDebugFuncPtr() {
return (void*)(myGstDebugFuncPtr);
}
extern void myGstElementCallAsyncFunc(GstElement* element, gpointer user_data);
static void* getPointer_myGstElementCallAsyncFunc() {
return (void*)(myGstElementCallAsyncFunc);
}
extern void myGstElementForeachPadFunc(GstElement* element, GstPad* pad, gpointer user_data);
static void* getPointer_myGstElementForeachPadFunc() {
return (void*)(myGstElementForeachPadFunc);
}
extern void myGstIteratorCopyFunction(GstIterator* it, GstIterator* copy1);
static void* getPointer_myGstIteratorCopyFunction() {
return (void*)(myGstIteratorCopyFunction);
}
extern void myGstIteratorFoldFunction(GValue* item, GValue* ret, gpointer user_data);
static void* getPointer_myGstIteratorFoldFunction() {
return (void*)(myGstIteratorFoldFunction);
}
extern void myGstIteratorForeachFunction(GValue* item, gpointer user_data);
static void* getPointer_myGstIteratorForeachFunction() {
return (void*)(myGstIteratorForeachFunction);
}
extern void myGstIteratorFreeFunction(GstIterator* it);
static void* getPointer_myGstIteratorFreeFunction() {
return (void*)(myGstIteratorFreeFunction);
}
extern void myGstIteratorItemFunction(GstIterator* it, GValue* item);
static void* getPointer_myGstIteratorItemFunction() {
return (void*)(myGstIteratorItemFunction);
}
extern void myGstIteratorNextFunction(GstIterator* it, GValue* result);
static void* getPointer_myGstIteratorNextFunction() {
return (void*)(myGstIteratorNextFunction);
}
extern void myGstIteratorResyncFunction(GstIterator* it);
static void* getPointer_myGstIteratorResyncFunction() {
return (void*)(myGstIteratorResyncFunction);
}
extern void myGstLogFunction(GstDebugCategory* category, GstDebugLevel level, gchar* file, gchar* function, gint32 line, GObject* object, GstDebugMessage* message, gpointer user_data);
static void* getPointer_myGstLogFunction() {
return (void*)(myGstLogFunction);
}
extern void myGstMemoryCopyFunction(GstMemory* mem, gint64 offset, gint64 size);
static void* getPointer_myGstMemoryCopyFunction() {
return (void*)(myGstMemoryCopyFunction);
}
extern void myGstMemoryIsSpanFunction(GstMemory* mem1, GstMemory* mem2, guint64* offset);
static void* getPointer_myGstMemoryIsSpanFunction() {
return (void*)(myGstMemoryIsSpanFunction);
}
extern void myGstMemoryMapFullFunction(GstMemory* mem, GstMapInfo* info, guint64 maxsize);
static void* getPointer_myGstMemoryMapFullFunction() {
return (void*)(myGstMemoryMapFullFunction);
}
extern void myGstMemoryMapFunction(GstMemory* mem, guint64 maxsize, GstMapFlags flags);
static void* getPointer_myGstMemoryMapFunction() {
return (void*)(myGstMemoryMapFunction);
}
extern void myGstMemoryShareFunction(GstMemory* mem, gint64 offset, gint64 size);
static void* getPointer_myGstMemoryShareFunction() {
return (void*)(myGstMemoryShareFunction);
}
extern void myGstMemoryUnmapFullFunction(GstMemory* mem, GstMapInfo* info);
static void* getPointer_myGstMemoryUnmapFullFunction() {
return (void*)(myGstMemoryUnmapFullFunction);
}
extern void myGstMemoryUnmapFunction(GstMemory* mem);
static void* getPointer_myGstMemoryUnmapFunction() {
return (void*)(myGstMemoryUnmapFunction);
}
extern void myGstMetaFreeFunction(GstMeta* meta, GstBuffer* buffer);
static void* getPointer_myGstMetaFreeFunction() {
return (void*)(myGstMetaFreeFunction);
}
extern void myGstMetaInitFunction(GstMeta* meta, gpointer params, GstBuffer* buffer);
static void* getPointer_myGstMetaInitFunction() {
return (void*)(myGstMetaInitFunction);
}
extern void myGstMetaTransformFunction(GstBuffer* transbuf, GstMeta* meta, GstBuffer* buffer, guint32 type1, gpointer data);
static void* getPointer_myGstMetaTransformFunction() {
return (void*)(myGstMetaTransformFunction);
}
extern void myGstMiniObjectDisposeFunction(GstMiniObject* obj);
static void* getPointer_myGstMiniObjectDisposeFunction() {
return (void*)(myGstMiniObjectDisposeFunction);
}
extern void myGstMiniObjectFreeFunction(GstMiniObject* obj);
static void* getPointer_myGstMiniObjectFreeFunction() {
return (void*)(myGstMiniObjectFreeFunction);
}
extern void myGstMiniObjectNotify(gpointer user_data, GstMiniObject* obj);
static void* getPointer_myGstMiniObjectNotify() {
return (void*)(myGstMiniObjectNotify);
}
extern void myGstPadActivateFunction(GstPad* pad, GstObject* parent);
static void* getPointer_myGstPadActivateFunction() {
return (void*)(myGstPadActivateFunction);
}
extern void myGstPadActivateModeFunction(GstPad* pad, GstObject* parent, GstPadMode mode, gboolean active);
static void* getPointer_myGstPadActivateModeFunction() {
return (void*)(myGstPadActivateModeFunction);
}
extern void myGstPadChainFunction(GstPad* pad, GstObject* parent, GstBuffer* buffer);
static void* getPointer_myGstPadChainFunction() {
return (void*)(myGstPadChainFunction);
}
extern void myGstPadChainListFunction(GstPad* pad, GstObject* parent, GstBufferList* list);
static void* getPointer_myGstPadChainListFunction() {
return (void*)(myGstPadChainListFunction);
}
extern void myGstPadEventFullFunction(GstPad* pad, GstObject* parent, GstEvent* event);
static void* getPointer_myGstPadEventFullFunction() {
return (void*)(myGstPadEventFullFunction);
}
extern void myGstPadEventFunction(GstPad* pad, GstObject* parent, GstEvent* event);
static void* getPointer_myGstPadEventFunction() {
return (void*)(myGstPadEventFunction);
}
extern void myGstPadForwardFunction(GstPad* pad, gpointer user_data);
static void* getPointer_myGstPadForwardFunction() {
return (void*)(myGstPadForwardFunction);
}
extern void myGstPadGetRangeFunction(GstPad* pad, GstObject* parent, guint64 offset, guint32 length, GstBuffer* buffer);
static void* getPointer_myGstPadGetRangeFunction() {
return (void*)(myGstPadGetRangeFunction);
}
extern void myGstPadIterIntLinkFunction(GstPad* pad, GstObject* parent);
static void* getPointer_myGstPadIterIntLinkFunction() {
return (void*)(myGstPadIterIntLinkFunction);
}
extern void myGstPadLinkFunction(GstPad* pad, GstObject* parent, GstPad* peer);
static void* getPointer_myGstPadLinkFunction() {
return (void*)(myGstPadLinkFunction);
}
extern void myGstPadProbeCallback(GstPad* pad, GstPadProbeInfo* info, gpointer user_data);
static void* getPointer_myGstPadProbeCallback() {
return (void*)(myGstPadProbeCallback);
}
extern void myGstPadQueryFunction(GstPad* pad, GstObject* parent, GstQuery* query);
static void* getPointer_myGstPadQueryFunction() {
return (void*)(myGstPadQueryFunction);
}
extern void myGstPadStickyEventsForeachFunction(GstPad* pad, GstEvent* event, gpointer user_data);
static void* getPointer_myGstPadStickyEventsForeachFunction() {
return (void*)(myGstPadStickyEventsForeachFunction);
}
extern void myGstPadUnlinkFunction(GstPad* pad, GstObject* parent);
static void* getPointer_myGstPadUnlinkFunction() {
return (void*)(myGstPadUnlinkFunction);
}
extern void myGstPluginFeatureFilter(GstPluginFeature* feature, gpointer user_data);
static void* getPointer_myGstPluginFeatureFilter() {
return (void*)(myGstPluginFeatureFilter);
}
extern void myGstPluginFilter(GstPlugin* plugin, gpointer user_data);
static void* getPointer_myGstPluginFilter() {
return (void*)(myGstPluginFilter);
}
extern void myGstPluginInitFullFunc(GstPlugin* plugin, gpointer user_data);
static void* getPointer_myGstPluginInitFullFunc() {
return (void*)(myGstPluginInitFullFunc);
}
extern void myGstPluginInitFunc(GstPlugin* plugin);
static void* getPointer_myGstPluginInitFunc() {
return (void*)(myGstPluginInitFunc);
}
extern void myGstPromiseChangeFunc(GstPromise* promise, gpointer user_data);
static void* getPointer_myGstPromiseChangeFunc() {
return (void*)(myGstPromiseChangeFunc);
}
extern void myGstStructureFilterMapFunc(guint32 field_id, GValue* value, gpointer user_data);
static void* getPointer_myGstStructureFilterMapFunc() {
return (void*)(myGstStructureFilterMapFunc);
}
extern void myGstStructureForeachFunc(guint32 field_id, GValue* value, gpointer user_data);
static void* getPointer_myGstStructureForeachFunc() {
return (void*)(myGstStructureForeachFunc);
}
extern void myGstStructureMapFunc(guint32 field_id, GValue* value, gpointer user_data);
static void* getPointer_myGstStructureMapFunc() {
return (void*)(myGstStructureMapFunc);
}
extern void myGstTagForeachFunc(GstTagList* list, gchar* tag, gpointer user_data);
static void* getPointer_myGstTagForeachFunc() {
return (void*)(myGstTagForeachFunc);
}
extern void myGstTagMergeFunc(GValue* dest, GValue* src);
static void* getPointer_myGstTagMergeFunc() {
return (void*)(myGstTagMergeFunc);
}
extern void myGstTaskFunction(gpointer user_data);
static void* getPointer_myGstTaskFunction() {
return (void*)(myGstTaskFunction);
}
extern void myGstTaskPoolFunction(gpointer user_data);
static void* getPointer_myGstTaskPoolFunction() {
return (void*)(myGstTaskPoolFunction);
}
extern void myGstTaskThreadFunc(GstTask* task, GThread* thread, gpointer user_data);
static void* getPointer_myGstTaskThreadFunc() {
return (void*)(myGstTaskThreadFunc);
}
extern void myGstTypeFindFunction(GstTypeFind* find, gpointer user_data);
static void* getPointer_myGstTypeFindFunction() {
return (void*)(myGstTypeFindFunction);
}
extern void myGstValueCompareFunc(GValue* value1, GValue* value2);
static void* getPointer_myGstValueCompareFunc() {
return (void*)(myGstValueCompareFunc);
}
extern void myGstValueDeserializeFunc(GValue* dest, gchar* s);
static void* getPointer_myGstValueDeserializeFunc() {
return (void*)(myGstValueDeserializeFunc);
}
extern void myGstValueSerializeFunc(GValue* value1);
static void* getPointer_myGstValueSerializeFunc() {
return (void*)(myGstValueSerializeFunc);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Gst")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Gst", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct AllocationParams
type AllocationParams struct {
	P unsafe.Pointer
}

const SizeOfStructAllocationParams = 64

func AllocationParamsGetType() gi.GType {
	ret := _I.GetGType(0, "AllocationParams")
	return ret
}

// gst_allocation_params_copy
//
// [ result ] trans: everything
//
func (v AllocationParams) Copy() (result AllocationParams) {
	iv, err := _I.Get(0, "AllocationParams", "copy")
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

// gst_allocation_params_free
//
func (v AllocationParams) Free() {
	iv, err := _I.Get(1, "AllocationParams", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_allocation_params_init
//
func (v AllocationParams) Init() {
	iv, err := _I.Get(2, "AllocationParams", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object Allocator
type Allocator struct {
	Object
}

func WrapAllocator(p unsafe.Pointer) (r Allocator) { r.P = p; return }

type IAllocator interface{ P_Allocator() unsafe.Pointer }

func (v Allocator) P_Allocator() unsafe.Pointer { return v.P }
func AllocatorGetType() gi.GType {
	ret := _I.GetGType(1, "Allocator")
	return ret
}

// gst_allocator_find
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func AllocatorFind1(name string) (result Allocator) {
	iv, err := _I.Get(3, "Allocator", "find")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_allocator_register
//
// [ name ] trans: nothing
//
// [ allocator ] trans: everything
//
func AllocatorRegister1(name string, allocator IAllocator) {
	iv, err := _I.Get(4, "Allocator", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	var tmp unsafe.Pointer
	if allocator != nil {
		tmp = allocator.P_Allocator()
	}
	arg_name := gi.NewStringArgument(c_name)
	arg_allocator := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_name, arg_allocator}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gst_allocator_alloc
//
// [ size ] trans: nothing
//
// [ params ] trans: nothing
//
// [ result ] trans: everything
//
func (v Allocator) Alloc(size uint64, params AllocationParams) (result Memory) {
	iv, err := _I.Get(5, "Allocator", "alloc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewUint64Argument(size)
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_size, arg_params}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_allocator_free
//
// [ memory ] trans: everything
//
func (v Allocator) Free(memory Memory) {
	iv, err := _I.Get(6, "Allocator", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_memory := gi.NewPointerArgument(memory.P)
	args := []gi.Argument{arg_v, arg_memory}
	iv.Call(args, nil, nil)
}

// gst_allocator_set_default
//
func (v Allocator) SetDefault() {
	iv, err := _I.Get(7, "Allocator", "set_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct AllocatorClass

// Flags AllocatorFlags
type AllocatorFlags int

const (
	AllocatorFlagsCustomAlloc AllocatorFlags = 16
	AllocatorFlagsLast        AllocatorFlags = 1048576
)

func AllocatorFlagsGetType() gi.GType {
	ret := _I.GetGType(2, "AllocatorFlags")
	return ret
}

// Struct AllocatorPrivate
type AllocatorPrivate struct {
	P unsafe.Pointer
}

func AllocatorPrivateGetType() gi.GType {
	ret := _I.GetGType(3, "AllocatorPrivate")
	return ret
}

// Struct AtomicQueue
type AtomicQueue struct {
	P unsafe.Pointer
}

func AtomicQueueGetType() gi.GType {
	ret := _I.GetGType(4, "AtomicQueue")
	return ret
}

// gst_atomic_queue_new
//
// [ initial_size ] trans: nothing
//
// [ result ] trans: everything
//
func NewAtomicQueue(initial_size uint32) (result AtomicQueue) {
	iv, err := _I.Get(8, "AtomicQueue", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_initial_size := gi.NewUint32Argument(initial_size)
	args := []gi.Argument{arg_initial_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_atomic_queue_length
//
// [ result ] trans: nothing
//
func (v AtomicQueue) Length() (result uint32) {
	iv, err := _I.Get(9, "AtomicQueue", "length")
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

// gst_atomic_queue_peek
//
// [ result ] trans: nothing
//
func (v AtomicQueue) Peek() (result unsafe.Pointer) {
	iv, err := _I.Get(10, "AtomicQueue", "peek")
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

// gst_atomic_queue_pop
//
// [ result ] trans: everything
//
func (v AtomicQueue) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(11, "AtomicQueue", "pop")
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

// gst_atomic_queue_push
//
// [ data ] trans: nothing
//
func (v AtomicQueue) Push(data unsafe.Pointer) {
	iv, err := _I.Get(12, "AtomicQueue", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// gst_atomic_queue_ref
//
func (v AtomicQueue) Ref() {
	iv, err := _I.Get(13, "AtomicQueue", "ref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_atomic_queue_unref
//
func (v AtomicQueue) Unref() {
	iv, err := _I.Get(14, "AtomicQueue", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object Bin
type Bin struct {
	ChildProxyIfc
	Element
}

func WrapBin(p unsafe.Pointer) (r Bin) { r.P = p; return }

type IBin interface{ P_Bin() unsafe.Pointer }

func (v Bin) P_Bin() unsafe.Pointer        { return v.P }
func (v Bin) P_ChildProxy() unsafe.Pointer { return v.P }
func BinGetType() gi.GType {
	ret := _I.GetGType(5, "Bin")
	return ret
}

// gst_bin_new
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func NewBin(name string) (result Bin) {
	iv, err := _I.Get(15, "Bin", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_bin_add
//
// [ element ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bin) Add(element IElement) (result bool) {
	iv, err := _I.Get(16, "Bin", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if element != nil {
		tmp = element.P_Element()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_element := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_bin_find_unlinked_pad
//
// [ direction ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bin) FindUnlinkedPad(direction PadDirectionEnum) (result Pad) {
	iv, err := _I.Get(17, "Bin", "find_unlinked_pad")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_direction := gi.NewIntArgument(int(direction))
	args := []gi.Argument{arg_v, arg_direction}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bin_get_by_interface
//
// [ iface ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bin) GetByInterface(iface gi.GType) (result Element) {
	iv, err := _I.Get(18, "Bin", "get_by_interface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iface := gi.NewUintArgument(uint(iface))
	args := []gi.Argument{arg_v, arg_iface}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bin_get_by_name
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bin) GetByName(name string) (result Element) {
	iv, err := _I.Get(19, "Bin", "get_by_name")
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
	result.P = ret.Pointer()
	return
}

// gst_bin_get_by_name_recurse_up
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bin) GetByNameRecurseUp(name string) (result Element) {
	iv, err := _I.Get(20, "Bin", "get_by_name_recurse_up")
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
	result.P = ret.Pointer()
	return
}

// gst_bin_get_suppressed_flags
//
// [ result ] trans: nothing
//
func (v Bin) GetSuppressedFlags() (result ElementFlags) {
	iv, err := _I.Get(21, "Bin", "get_suppressed_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ElementFlags(ret.Int())
	return
}

// gst_bin_iterate_all_by_interface
//
// [ iface ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bin) IterateAllByInterface(iface gi.GType) (result Iterator) {
	iv, err := _I.Get(22, "Bin", "iterate_all_by_interface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iface := gi.NewUintArgument(uint(iface))
	args := []gi.Argument{arg_v, arg_iface}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bin_iterate_elements
//
// [ result ] trans: everything
//
func (v Bin) IterateElements() (result Iterator) {
	iv, err := _I.Get(23, "Bin", "iterate_elements")
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

// gst_bin_iterate_recurse
//
// [ result ] trans: everything
//
func (v Bin) IterateRecurse() (result Iterator) {
	iv, err := _I.Get(24, "Bin", "iterate_recurse")
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

// gst_bin_iterate_sinks
//
// [ result ] trans: everything
//
func (v Bin) IterateSinks() (result Iterator) {
	iv, err := _I.Get(25, "Bin", "iterate_sinks")
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

// gst_bin_iterate_sorted
//
// [ result ] trans: everything
//
func (v Bin) IterateSorted() (result Iterator) {
	iv, err := _I.Get(26, "Bin", "iterate_sorted")
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

// gst_bin_iterate_sources
//
// [ result ] trans: everything
//
func (v Bin) IterateSources() (result Iterator) {
	iv, err := _I.Get(27, "Bin", "iterate_sources")
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

// gst_bin_recalculate_latency
//
// [ result ] trans: nothing
//
func (v Bin) RecalculateLatency() (result bool) {
	iv, err := _I.Get(28, "Bin", "recalculate_latency")
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

// gst_bin_remove
//
// [ element ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bin) Remove(element IElement) (result bool) {
	iv, err := _I.Get(29, "Bin", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if element != nil {
		tmp = element.P_Element()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_element := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_bin_set_suppressed_flags
//
// [ flags ] trans: nothing
//
func (v Bin) SetSuppressedFlags(flags ElementFlags) {
	iv, err := _I.Get(30, "Bin", "set_suppressed_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, nil)
}

// gst_bin_sync_children_states
//
// [ result ] trans: nothing
//
func (v Bin) SyncChildrenStates() (result bool) {
	iv, err := _I.Get(31, "Bin", "sync_children_states")
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

// ignore GType struct BinClass

// Flags BinFlags
type BinFlags int

const (
	BinFlagsNoResync     BinFlags = 16384
	BinFlagsStreamsAware BinFlags = 32768
	BinFlagsLast         BinFlags = 524288
)

func BinFlagsGetType() gi.GType {
	ret := _I.GetGType(6, "BinFlags")
	return ret
}

// Struct BinPrivate
type BinPrivate struct {
	P unsafe.Pointer
}

func BinPrivateGetType() gi.GType {
	ret := _I.GetGType(7, "BinPrivate")
	return ret
}

// Object Bitmask
type Bitmask struct {
	P unsafe.Pointer
}

func WrapBitmask(p unsafe.Pointer) (r Bitmask) { r.P = p; return }

type IBitmask interface{ P_Bitmask() unsafe.Pointer }

func (v Bitmask) P_Bitmask() unsafe.Pointer { return v.P }
func BitmaskGetType() gi.GType {
	ret := _I.GetGType(8, "Bitmask")
	return ret
}

// Struct Buffer
type Buffer struct {
	P unsafe.Pointer
}

const SizeOfStructBuffer = 112

func BufferGetType() gi.GType {
	ret := _I.GetGType(9, "Buffer")
	return ret
}

// gst_buffer_new
//
// [ result ] trans: everything
//
func NewBuffer() (result Buffer) {
	iv, err := _I.Get(32, "Buffer", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_new_allocate
//
// [ allocator ] trans: nothing
//
// [ size ] trans: nothing
//
// [ params ] trans: nothing
//
// [ result ] trans: everything
//
func NewBufferAllocate(allocator IAllocator, size uint64, params AllocationParams) (result Buffer) {
	iv, err := _I.Get(33, "Buffer", "new_allocate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if allocator != nil {
		tmp = allocator.P_Allocator()
	}
	arg_allocator := gi.NewPointerArgument(tmp)
	arg_size := gi.NewUint64Argument(size)
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_allocator, arg_size, arg_params}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_new_wrapped
//
// [ data ] trans: everything
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func NewBufferWrapped(data gi.Uint8Array, size uint64) (result Buffer) {
	iv, err := _I.Get(34, "Buffer", "new_wrapped")
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

// gst_buffer_new_wrapped_full
//
// [ flags ] trans: nothing
//
// [ data ] trans: nothing
//
// [ maxsize ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: everything
//
func NewBufferWrappedFull(flags MemoryFlags, data gi.Uint8Array, maxsize uint64, offset uint64, size uint64, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result Buffer) {
	iv, err := _I.Get(35, "Buffer", "new_wrapped_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags := gi.NewIntArgument(int(flags))
	arg_data := gi.NewPointerArgument(data.P)
	arg_maxsize := gi.NewUint64Argument(maxsize)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_flags, arg_data, arg_maxsize, arg_offset, arg_size, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_add_meta
//
// [ info ] trans: nothing
//
// [ params ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) AddMeta(info MetaInfo, params unsafe.Pointer) (result Meta) {
	iv, err := _I.Get(36, "Buffer", "add_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	arg_params := gi.NewPointerArgument(params)
	args := []gi.Argument{arg_v, arg_info, arg_params}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_add_parent_buffer_meta
//
// [ ref ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) AddParentBufferMeta(ref Buffer) (result ParentBufferMeta) {
	iv, err := _I.Get(37, "Buffer", "add_parent_buffer_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ref := gi.NewPointerArgument(ref.P)
	args := []gi.Argument{arg_v, arg_ref}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_add_protection_meta
//
// [ info ] trans: everything
//
// [ result ] trans: nothing
//
func (v Buffer) AddProtectionMeta(info Structure) (result ProtectionMeta) {
	iv, err := _I.Get(38, "Buffer", "add_protection_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_v, arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_add_reference_timestamp_meta
//
// [ reference ] trans: nothing
//
// [ timestamp ] trans: nothing
//
// [ duration ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) AddReferenceTimestampMeta(reference Caps, timestamp uint64, duration uint64) (result ReferenceTimestampMeta) {
	iv, err := _I.Get(39, "Buffer", "add_reference_timestamp_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_reference := gi.NewPointerArgument(reference.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_duration := gi.NewUint64Argument(duration)
	args := []gi.Argument{arg_v, arg_reference, arg_timestamp, arg_duration}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_append
//
// [ buf2 ] trans: everything
//
// [ result ] trans: everything
//
func (v Buffer) Append(buf2 Buffer) (result Buffer) {
	iv, err := _I.Get(40, "Buffer", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf2 := gi.NewPointerArgument(buf2.P)
	args := []gi.Argument{arg_v, arg_buf2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_append_memory
//
// [ mem ] trans: everything
//
func (v Buffer) AppendMemory(mem Memory) {
	iv, err := _I.Get(41, "Buffer", "append_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mem := gi.NewPointerArgument(mem.P)
	args := []gi.Argument{arg_v, arg_mem}
	iv.Call(args, nil, nil)
}

// gst_buffer_append_region
//
// [ buf2 ] trans: everything
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v Buffer) AppendRegion(buf2 Buffer, offset int64, size int64) (result Buffer) {
	iv, err := _I.Get(42, "Buffer", "append_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf2 := gi.NewPointerArgument(buf2.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_buf2, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_copy_deep
//
// [ result ] trans: everything
//
func (v Buffer) CopyDeep() (result Buffer) {
	iv, err := _I.Get(43, "Buffer", "copy_deep")
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

// gst_buffer_copy_into
//
// [ src ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) CopyInto(src Buffer, flags BufferCopyFlags, offset uint64, size uint64) (result bool) {
	iv, err := _I.Get(44, "Buffer", "copy_into")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_src := gi.NewPointerArgument(src.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_src, arg_flags, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_copy_region
//
// [ flags ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v Buffer) CopyRegion(flags BufferCopyFlags, offset uint64, size uint64) (result Buffer) {
	iv, err := _I.Get(45, "Buffer", "copy_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_flags, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_extract
//
// [ offset ] trans: nothing
//
// [ dest ] trans: nothing, dir: out
//
// [ size ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Buffer) Extract(offset uint64, dest gi.Uint8Array) (result uint64) {
	iv, err := _I.Get(46, "Buffer", "extract")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_dest, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint64
	_ = size
	size = outArgs[0].Uint64()
	result = ret.Uint64()
	return
}

// gst_buffer_extract_dup
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ dest ] trans: everything, dir: out
//
// [ dest_size ] trans: everything, dir: out
//
func (v Buffer) ExtractDup(offset uint64, size uint64) (dest gi.Uint8Array) {
	iv, err := _I.Get(47, "Buffer", "extract_dup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	arg_dest := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_dest_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_size, arg_dest, arg_dest_size}
	iv.Call(args, nil, &outArgs[0])
	var dest_size uint64
	_ = dest_size
	dest.P = outArgs[0].Pointer()
	dest_size = outArgs[1].Uint64()
	dest.Len = int(dest_size)
	return
}

// gst_buffer_fill
//
// [ offset ] trans: nothing
//
// [ src ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) Fill(offset uint64, src gi.Uint8Array, size uint64) (result uint64) {
	iv, err := _I.Get(48, "Buffer", "fill")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_src := gi.NewPointerArgument(src.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_src, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_buffer_find_memory
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ idx ] trans: everything, dir: out
//
// [ length ] trans: everything, dir: out
//
// [ skip ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Buffer) FindMemory(offset uint64, size uint64) (result bool, idx uint32, length uint32, skip uint64) {
	iv, err := _I.Get(49, "Buffer", "find_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	arg_idx := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_skip := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_offset, arg_size, arg_idx, arg_length, arg_skip}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	idx = outArgs[0].Uint32()
	length = outArgs[1].Uint32()
	skip = outArgs[2].Uint64()
	result = ret.Bool()
	return
}

// gst_buffer_foreach_meta
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) ForeachMeta(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(50, "Buffer", "foreach_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myBufferForeachMetaFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_get_all_memory
//
// [ result ] trans: everything
//
func (v Buffer) GetAllMemory() (result Memory) {
	iv, err := _I.Get(51, "Buffer", "get_all_memory")
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

// gst_buffer_get_flags
//
// [ result ] trans: nothing
//
func (v Buffer) GetFlags() (result BufferFlags) {
	iv, err := _I.Get(52, "Buffer", "get_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = BufferFlags(ret.Int())
	return
}

// gst_buffer_get_memory
//
// [ idx ] trans: nothing
//
// [ result ] trans: everything
//
func (v Buffer) GetMemory(idx uint32) (result Memory) {
	iv, err := _I.Get(53, "Buffer", "get_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_get_memory_range
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: everything
//
func (v Buffer) GetMemoryRange(idx uint32, length int32) (result Memory) {
	iv, err := _I.Get(54, "Buffer", "get_memory_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_idx, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_get_meta
//
// [ api ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) GetMeta(api gi.GType) (result Meta) {
	iv, err := _I.Get(55, "Buffer", "get_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_api := gi.NewUintArgument(uint(api))
	args := []gi.Argument{arg_v, arg_api}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_get_n_meta
//
// [ api_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) GetNMeta(api_type gi.GType) (result uint32) {
	iv, err := _I.Get(56, "Buffer", "get_n_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_api_type := gi.NewUintArgument(uint(api_type))
	args := []gi.Argument{arg_v, arg_api_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_buffer_get_reference_timestamp_meta
//
// [ reference ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) GetReferenceTimestampMeta(reference Caps) (result ReferenceTimestampMeta) {
	iv, err := _I.Get(57, "Buffer", "get_reference_timestamp_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_reference := gi.NewPointerArgument(reference.P)
	args := []gi.Argument{arg_v, arg_reference}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_get_size
//
// [ result ] trans: nothing
//
func (v Buffer) GetSize() (result uint64) {
	iv, err := _I.Get(58, "Buffer", "get_size")
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

// gst_buffer_get_sizes
//
// [ offset ] trans: everything, dir: out
//
// [ maxsize ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Buffer) GetSizes() (result uint64, offset uint64, maxsize uint64) {
	iv, err := _I.Get(59, "Buffer", "get_sizes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_maxsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_maxsize}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	offset = outArgs[0].Uint64()
	maxsize = outArgs[1].Uint64()
	result = ret.Uint64()
	return
}

// gst_buffer_get_sizes_range
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
// [ offset ] trans: everything, dir: out
//
// [ maxsize ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Buffer) GetSizesRange(idx uint32, length int32) (result uint64, offset uint64, maxsize uint64) {
	iv, err := _I.Get(60, "Buffer", "get_sizes_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	arg_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_maxsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_idx, arg_length, arg_offset, arg_maxsize}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	offset = outArgs[0].Uint64()
	maxsize = outArgs[1].Uint64()
	result = ret.Uint64()
	return
}

// gst_buffer_has_flags
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) HasFlags(flags BufferFlags) (result bool) {
	iv, err := _I.Get(61, "Buffer", "has_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_insert_memory
//
// [ idx ] trans: nothing
//
// [ mem ] trans: everything
//
func (v Buffer) InsertMemory(idx int32, mem Memory) {
	iv, err := _I.Get(62, "Buffer", "insert_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewInt32Argument(idx)
	arg_mem := gi.NewPointerArgument(mem.P)
	args := []gi.Argument{arg_v, arg_idx, arg_mem}
	iv.Call(args, nil, nil)
}

// gst_buffer_is_all_memory_writable
//
// [ result ] trans: nothing
//
func (v Buffer) IsAllMemoryWritable() (result bool) {
	iv, err := _I.Get(63, "Buffer", "is_all_memory_writable")
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

// gst_buffer_is_memory_range_writable
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) IsMemoryRangeWritable(idx uint32, length int32) (result bool) {
	iv, err := _I.Get(64, "Buffer", "is_memory_range_writable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_idx, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_map
//
// [ info ] trans: nothing, dir: out
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) Map(info MapInfo, flags MapFlags) (result bool) {
	iv, err := _I.Get(65, "Buffer", "map")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_info, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_map_range
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
// [ info ] trans: nothing, dir: out
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) MapRange(idx uint32, length int32, info MapInfo, flags MapFlags) (result bool) {
	iv, err := _I.Get(66, "Buffer", "map_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	arg_info := gi.NewPointerArgument(info.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_idx, arg_length, arg_info, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_memcmp
//
// [ offset ] trans: nothing
//
// [ mem ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) Memcmp(offset uint64, mem gi.Uint8Array, size uint64) (result int32) {
	iv, err := _I.Get(67, "Buffer", "memcmp")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_mem := gi.NewPointerArgument(mem.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_mem, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_buffer_memset
//
// [ offset ] trans: nothing
//
// [ val ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) Memset(offset uint64, val uint8, size uint64) (result uint64) {
	iv, err := _I.Get(68, "Buffer", "memset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_val := gi.NewUint8Argument(val)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_val, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_buffer_n_memory
//
// [ result ] trans: nothing
//
func (v Buffer) NMemory() (result uint32) {
	iv, err := _I.Get(69, "Buffer", "n_memory")
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

// gst_buffer_peek_memory
//
// [ idx ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) PeekMemory(idx uint32) (result Memory) {
	iv, err := _I.Get(70, "Buffer", "peek_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_prepend_memory
//
// [ mem ] trans: everything
//
func (v Buffer) PrependMemory(mem Memory) {
	iv, err := _I.Get(71, "Buffer", "prepend_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mem := gi.NewPointerArgument(mem.P)
	args := []gi.Argument{arg_v, arg_mem}
	iv.Call(args, nil, nil)
}

// gst_buffer_remove_all_memory
//
func (v Buffer) RemoveAllMemory() {
	iv, err := _I.Get(72, "Buffer", "remove_all_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_buffer_remove_memory
//
// [ idx ] trans: nothing
//
func (v Buffer) RemoveMemory(idx uint32) {
	iv, err := _I.Get(73, "Buffer", "remove_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	iv.Call(args, nil, nil)
}

// gst_buffer_remove_memory_range
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
func (v Buffer) RemoveMemoryRange(idx uint32, length int32) {
	iv, err := _I.Get(74, "Buffer", "remove_memory_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_idx, arg_length}
	iv.Call(args, nil, nil)
}

// gst_buffer_remove_meta
//
// [ meta ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) RemoveMeta(meta Meta) (result bool) {
	iv, err := _I.Get(75, "Buffer", "remove_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_meta := gi.NewPointerArgument(meta.P)
	args := []gi.Argument{arg_v, arg_meta}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_replace_all_memory
//
// [ mem ] trans: everything
//
func (v Buffer) ReplaceAllMemory(mem Memory) {
	iv, err := _I.Get(76, "Buffer", "replace_all_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mem := gi.NewPointerArgument(mem.P)
	args := []gi.Argument{arg_v, arg_mem}
	iv.Call(args, nil, nil)
}

// gst_buffer_replace_memory
//
// [ idx ] trans: nothing
//
// [ mem ] trans: everything
//
func (v Buffer) ReplaceMemory(idx uint32, mem Memory) {
	iv, err := _I.Get(77, "Buffer", "replace_memory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_mem := gi.NewPointerArgument(mem.P)
	args := []gi.Argument{arg_v, arg_idx, arg_mem}
	iv.Call(args, nil, nil)
}

// gst_buffer_replace_memory_range
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
// [ mem ] trans: everything
//
func (v Buffer) ReplaceMemoryRange(idx uint32, length int32, mem Memory) {
	iv, err := _I.Get(78, "Buffer", "replace_memory_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	arg_mem := gi.NewPointerArgument(mem.P)
	args := []gi.Argument{arg_v, arg_idx, arg_length, arg_mem}
	iv.Call(args, nil, nil)
}

// gst_buffer_resize
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
func (v Buffer) Resize(offset int64, size int64) {
	iv, err := _I.Get(79, "Buffer", "resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_size}
	iv.Call(args, nil, nil)
}

// gst_buffer_resize_range
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) ResizeRange(idx uint32, length int32, offset int64, size int64) (result bool) {
	iv, err := _I.Get(80, "Buffer", "resize_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewInt32Argument(length)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_idx, arg_length, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_set_flags
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) SetFlags(flags BufferFlags) (result bool) {
	iv, err := _I.Get(81, "Buffer", "set_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_set_size
//
// [ size ] trans: nothing
//
func (v Buffer) SetSize(size int64) {
	iv, err := _I.Get(82, "Buffer", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_size}
	iv.Call(args, nil, nil)
}

// gst_buffer_unmap
//
// [ info ] trans: nothing
//
func (v Buffer) Unmap(info MapInfo) {
	iv, err := _I.Get(83, "Buffer", "unmap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_v, arg_info}
	iv.Call(args, nil, nil)
}

// gst_buffer_unset_flags
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Buffer) UnsetFlags(flags BufferFlags) (result bool) {
	iv, err := _I.Get(84, "Buffer", "unset_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Flags BufferCopyFlags
type BufferCopyFlags int

const (
	BufferCopyFlagsNone       BufferCopyFlags = 0
	BufferCopyFlagsFlags      BufferCopyFlags = 1
	BufferCopyFlagsTimestamps BufferCopyFlags = 2
	BufferCopyFlagsMeta       BufferCopyFlags = 4
	BufferCopyFlagsMemory     BufferCopyFlags = 8
	BufferCopyFlagsMerge      BufferCopyFlags = 16
	BufferCopyFlagsDeep       BufferCopyFlags = 32
)

func BufferCopyFlagsGetType() gi.GType {
	ret := _I.GetGType(10, "BufferCopyFlags")
	return ret
}

// Flags BufferFlags
type BufferFlags int

const (
	BufferFlagsLive         BufferFlags = 16
	BufferFlagsDecodeOnly   BufferFlags = 32
	BufferFlagsDiscont      BufferFlags = 64
	BufferFlagsResync       BufferFlags = 128
	BufferFlagsCorrupted    BufferFlags = 256
	BufferFlagsMarker       BufferFlags = 512
	BufferFlagsHeader       BufferFlags = 1024
	BufferFlagsGap          BufferFlags = 2048
	BufferFlagsDroppable    BufferFlags = 4096
	BufferFlagsDeltaUnit    BufferFlags = 8192
	BufferFlagsTagMemory    BufferFlags = 16384
	BufferFlagsSyncAfter    BufferFlags = 32768
	BufferFlagsNonDroppable BufferFlags = 65536
	BufferFlagsLast         BufferFlags = 1048576
)

func BufferFlagsGetType() gi.GType {
	ret := _I.GetGType(11, "BufferFlags")
	return ret
}

type BufferForeachMetaFuncStruct struct {
	F_buffer Buffer
	F_meta   unsafe.Pointer /*TODO_CB tag: interface, isPtr: true*/
}

func GetPointer_myBufferForeachMetaFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBufferForeachMetaFunc())
}

//export myGstBufferForeachMetaFunc
func myGstBufferForeachMetaFunc(buffer *C.GstBuffer, meta C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &BufferForeachMetaFuncStruct{
		F_buffer: Buffer{P: unsafe.Pointer(buffer)},
		F_meta:   unsafe.Pointer(meta),
	}
	fn(args)
}

// Struct BufferList
type BufferList struct {
	P unsafe.Pointer
}

func BufferListGetType() gi.GType {
	ret := _I.GetGType(12, "BufferList")
	return ret
}

// gst_buffer_list_new
//
// [ result ] trans: everything
//
func NewBufferList() (result BufferList) {
	iv, err := _I.Get(86, "BufferList", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_list_new_sized
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func NewBufferListSized(size uint32) (result BufferList) {
	iv, err := _I.Get(87, "BufferList", "new_sized")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_list_calculate_size
//
// [ result ] trans: nothing
//
func (v BufferList) CalculateSize() (result uint64) {
	iv, err := _I.Get(88, "BufferList", "calculate_size")
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

// gst_buffer_list_copy_deep
//
// [ result ] trans: everything
//
func (v BufferList) CopyDeep() (result BufferList) {
	iv, err := _I.Get(89, "BufferList", "copy_deep")
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

// gst_buffer_list_foreach
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BufferList) Foreach(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(90, "BufferList", "foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myBufferListFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_list_get
//
// [ idx ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BufferList) Get(idx uint32) (result Buffer) {
	iv, err := _I.Get(91, "BufferList", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_list_get_writable
//
// [ idx ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BufferList) GetWritable(idx uint32) (result Buffer) {
	iv, err := _I.Get(92, "BufferList", "get_writable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_list_insert
//
// [ idx ] trans: nothing
//
// [ buffer ] trans: everything
//
func (v BufferList) Insert(idx int32, buffer Buffer) {
	iv, err := _I.Get(93, "BufferList", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewInt32Argument(idx)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_idx, arg_buffer}
	iv.Call(args, nil, nil)
}

// gst_buffer_list_length
//
// [ result ] trans: nothing
//
func (v BufferList) Length() (result uint32) {
	iv, err := _I.Get(94, "BufferList", "length")
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

// gst_buffer_list_remove
//
// [ idx ] trans: nothing
//
// [ length ] trans: nothing
//
func (v BufferList) Remove(idx uint32, length uint32) {
	iv, err := _I.Get(95, "BufferList", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	arg_length := gi.NewUint32Argument(length)
	args := []gi.Argument{arg_v, arg_idx, arg_length}
	iv.Call(args, nil, nil)
}

type BufferListFuncStruct struct {
	F_buffer unsafe.Pointer /*TODO_CB tag: interface, isPtr: true*/
	F_idx    uint32
}

func GetPointer_myBufferListFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBufferListFunc())
}

//export myGstBufferListFunc
func myGstBufferListFunc(buffer C.gpointer, idx C.guint32, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &BufferListFuncStruct{
		F_buffer: unsafe.Pointer(buffer),
		F_idx:    uint32(idx),
	}
	fn(args)
}

// Object BufferPool
type BufferPool struct {
	Object
}

func WrapBufferPool(p unsafe.Pointer) (r BufferPool) { r.P = p; return }

type IBufferPool interface{ P_BufferPool() unsafe.Pointer }

func (v BufferPool) P_BufferPool() unsafe.Pointer { return v.P }
func BufferPoolGetType() gi.GType {
	ret := _I.GetGType(13, "BufferPool")
	return ret
}

// gst_buffer_pool_new
//
// [ result ] trans: everything
//
func NewBufferPool() (result BufferPool) {
	iv, err := _I.Get(96, "BufferPool", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_pool_config_add_option
//
// [ config ] trans: nothing
//
// [ option ] trans: nothing
//
func BufferPoolConfigAddOption1(config Structure, option string) {
	iv, err := _I.Get(97, "BufferPool", "config_add_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_option := gi.CString(option)
	arg_config := gi.NewPointerArgument(config.P)
	arg_option := gi.NewStringArgument(c_option)
	args := []gi.Argument{arg_config, arg_option}
	iv.Call(args, nil, nil)
	gi.Free(c_option)
}

// gst_buffer_pool_config_get_allocator
//
// [ config ] trans: nothing
//
// [ allocator ] trans: nothing, dir: out
//
// [ params ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func BufferPoolConfigGetAllocator1(config Structure, params AllocationParams) (result bool, allocator Allocator) {
	iv, err := _I.Get(98, "BufferPool", "config_get_allocator")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_config := gi.NewPointerArgument(config.P)
	arg_allocator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_config, arg_allocator, arg_params}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	allocator.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_buffer_pool_config_get_option
//
// [ config ] trans: nothing
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferPoolConfigGetOption1(config Structure, index uint32) (result string) {
	iv, err := _I.Get(99, "BufferPool", "config_get_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_config := gi.NewPointerArgument(config.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_config, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_buffer_pool_config_get_params
//
// [ config ] trans: nothing
//
// [ caps ] trans: nothing, dir: out
//
// [ size ] trans: everything, dir: out
//
// [ min_buffers ] trans: everything, dir: out
//
// [ max_buffers ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func BufferPoolConfigGetParams1(config Structure) (result bool, caps Caps, size uint32, min_buffers uint32, max_buffers uint32) {
	iv, err := _I.Get(100, "BufferPool", "config_get_params")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_config := gi.NewPointerArgument(config.P)
	arg_caps := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_min_buffers := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_max_buffers := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_config, arg_caps, arg_size, arg_min_buffers, arg_max_buffers}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	caps.P = outArgs[0].Pointer()
	size = outArgs[1].Uint32()
	min_buffers = outArgs[2].Uint32()
	max_buffers = outArgs[3].Uint32()
	result = ret.Bool()
	return
}

// gst_buffer_pool_config_has_option
//
// [ config ] trans: nothing
//
// [ option ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferPoolConfigHasOption1(config Structure, option string) (result bool) {
	iv, err := _I.Get(101, "BufferPool", "config_has_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_option := gi.CString(option)
	arg_config := gi.NewPointerArgument(config.P)
	arg_option := gi.NewStringArgument(c_option)
	args := []gi.Argument{arg_config, arg_option}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_option)
	result = ret.Bool()
	return
}

// gst_buffer_pool_config_n_options
//
// [ config ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferPoolConfigNOptions1(config Structure) (result uint32) {
	iv, err := _I.Get(102, "BufferPool", "config_n_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_config := gi.NewPointerArgument(config.P)
	args := []gi.Argument{arg_config}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_buffer_pool_config_set_allocator
//
// [ config ] trans: nothing
//
// [ allocator ] trans: nothing
//
// [ params ] trans: nothing
//
func BufferPoolConfigSetAllocator1(config Structure, allocator IAllocator, params AllocationParams) {
	iv, err := _I.Get(103, "BufferPool", "config_set_allocator")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if allocator != nil {
		tmp = allocator.P_Allocator()
	}
	arg_config := gi.NewPointerArgument(config.P)
	arg_allocator := gi.NewPointerArgument(tmp)
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_config, arg_allocator, arg_params}
	iv.Call(args, nil, nil)
}

// gst_buffer_pool_config_set_params
//
// [ config ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ size ] trans: nothing
//
// [ min_buffers ] trans: nothing
//
// [ max_buffers ] trans: nothing
//
func BufferPoolConfigSetParams1(config Structure, caps Caps, size uint32, min_buffers uint32, max_buffers uint32) {
	iv, err := _I.Get(104, "BufferPool", "config_set_params")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_config := gi.NewPointerArgument(config.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_size := gi.NewUint32Argument(size)
	arg_min_buffers := gi.NewUint32Argument(min_buffers)
	arg_max_buffers := gi.NewUint32Argument(max_buffers)
	args := []gi.Argument{arg_config, arg_caps, arg_size, arg_min_buffers, arg_max_buffers}
	iv.Call(args, nil, nil)
}

// gst_buffer_pool_config_validate_params
//
// [ config ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ size ] trans: nothing
//
// [ min_buffers ] trans: nothing
//
// [ max_buffers ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferPoolConfigValidateParams1(config Structure, caps Caps, size uint32, min_buffers uint32, max_buffers uint32) (result bool) {
	iv, err := _I.Get(105, "BufferPool", "config_validate_params")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_config := gi.NewPointerArgument(config.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_size := gi.NewUint32Argument(size)
	arg_min_buffers := gi.NewUint32Argument(min_buffers)
	arg_max_buffers := gi.NewUint32Argument(max_buffers)
	args := []gi.Argument{arg_config, arg_caps, arg_size, arg_min_buffers, arg_max_buffers}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_pool_acquire_buffer
//
// [ buffer ] trans: everything, dir: out
//
// [ params ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BufferPool) AcquireBuffer(params BufferPoolAcquireParams) (result FlowReturnEnum, buffer Buffer) {
	iv, err := _I.Get(106, "BufferPool", "acquire_buffer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_buffer, arg_params}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	buffer.P = outArgs[0].Pointer()
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_buffer_pool_get_config
//
// [ result ] trans: everything
//
func (v BufferPool) GetConfig() (result Structure) {
	iv, err := _I.Get(107, "BufferPool", "get_config")
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

// gst_buffer_pool_get_options
//
// [ result ] trans: nothing
//
func (v BufferPool) GetOptions() (result gi.CStrArray) {
	iv, err := _I.Get(108, "BufferPool", "get_options")
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

// gst_buffer_pool_has_option
//
// [ option ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BufferPool) HasOption(option string) (result bool) {
	iv, err := _I.Get(109, "BufferPool", "has_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_option := gi.CString(option)
	arg_v := gi.NewPointerArgument(v.P)
	arg_option := gi.NewStringArgument(c_option)
	args := []gi.Argument{arg_v, arg_option}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_option)
	result = ret.Bool()
	return
}

// gst_buffer_pool_is_active
//
// [ result ] trans: nothing
//
func (v BufferPool) IsActive() (result bool) {
	iv, err := _I.Get(110, "BufferPool", "is_active")
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

// gst_buffer_pool_release_buffer
//
// [ buffer ] trans: everything
//
func (v BufferPool) ReleaseBuffer(buffer Buffer) {
	iv, err := _I.Get(111, "BufferPool", "release_buffer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_buffer}
	iv.Call(args, nil, nil)
}

// gst_buffer_pool_set_active
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BufferPool) SetActive(active bool) (result bool) {
	iv, err := _I.Get(112, "BufferPool", "set_active")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_v, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_pool_set_config
//
// [ config ] trans: everything
//
// [ result ] trans: nothing
//
func (v BufferPool) SetConfig(config Structure) (result bool) {
	iv, err := _I.Get(113, "BufferPool", "set_config")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_config := gi.NewPointerArgument(config.P)
	args := []gi.Argument{arg_v, arg_config}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_buffer_pool_set_flushing
//
// [ flushing ] trans: nothing
//
func (v BufferPool) SetFlushing(flushing bool) {
	iv, err := _I.Get(114, "BufferPool", "set_flushing")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flushing := gi.NewBoolArgument(flushing)
	args := []gi.Argument{arg_v, arg_flushing}
	iv.Call(args, nil, nil)
}

// Flags BufferPoolAcquireFlags
type BufferPoolAcquireFlags int

const (
	BufferPoolAcquireFlagsNone     BufferPoolAcquireFlags = 0
	BufferPoolAcquireFlagsKeyUnit  BufferPoolAcquireFlags = 1
	BufferPoolAcquireFlagsDontwait BufferPoolAcquireFlags = 2
	BufferPoolAcquireFlagsDiscont  BufferPoolAcquireFlags = 4
	BufferPoolAcquireFlagsLast     BufferPoolAcquireFlags = 65536
)

func BufferPoolAcquireFlagsGetType() gi.GType {
	ret := _I.GetGType(14, "BufferPoolAcquireFlags")
	return ret
}

// Struct BufferPoolAcquireParams
type BufferPoolAcquireParams struct {
	P unsafe.Pointer
}

const SizeOfStructBufferPoolAcquireParams = 64

func BufferPoolAcquireParamsGetType() gi.GType {
	ret := _I.GetGType(15, "BufferPoolAcquireParams")
	return ret
}

// ignore GType struct BufferPoolClass

// Struct BufferPoolPrivate
type BufferPoolPrivate struct {
	P unsafe.Pointer
}

func BufferPoolPrivateGetType() gi.GType {
	ret := _I.GetGType(16, "BufferPoolPrivate")
	return ret
}

// Enum BufferingMode
type BufferingModeEnum int

const (
	BufferingModeStream    BufferingModeEnum = 0
	BufferingModeDownload  BufferingModeEnum = 1
	BufferingModeTimeshift BufferingModeEnum = 2
	BufferingModeLive      BufferingModeEnum = 3
)

func BufferingModeGetType() gi.GType {
	ret := _I.GetGType(17, "BufferingMode")
	return ret
}

// Object Bus
type Bus struct {
	Object
}

func WrapBus(p unsafe.Pointer) (r Bus) { r.P = p; return }

type IBus interface{ P_Bus() unsafe.Pointer }

func (v Bus) P_Bus() unsafe.Pointer { return v.P }
func BusGetType() gi.GType {
	ret := _I.GetGType(18, "Bus")
	return ret
}

// gst_bus_new
//
// [ result ] trans: everything
//
func NewBus() (result Bus) {
	iv, err := _I.Get(115, "Bus", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bus_add_signal_watch
//
func (v Bus) AddSignalWatch() {
	iv, err := _I.Get(116, "Bus", "add_signal_watch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_bus_add_signal_watch_full
//
// [ priority ] trans: nothing
//
func (v Bus) AddSignalWatchFull(priority int32) {
	iv, err := _I.Get(117, "Bus", "add_signal_watch_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_priority}
	iv.Call(args, nil, nil)
}

// gst_bus_add_watch_full
//
// [ priority ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bus) AddWatch(priority int32, func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(118, "Bus", "add_watch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myBusFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_priority, arg_func1, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_bus_async_signal_func
//
// [ message ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bus) AsyncSignalFunc(message Message, data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(119, "Bus", "async_signal_func")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_message := gi.NewPointerArgument(message.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_message, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_bus_create_watch
//
// [ result ] trans: everything
//
func (v Bus) CreateWatch() (result g.Source) {
	iv, err := _I.Get(120, "Bus", "create_watch")
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

// gst_bus_disable_sync_message_emission
//
func (v Bus) DisableSyncMessageEmission() {
	iv, err := _I.Get(121, "Bus", "disable_sync_message_emission")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_bus_enable_sync_message_emission
//
func (v Bus) EnableSyncMessageEmission() {
	iv, err := _I.Get(122, "Bus", "enable_sync_message_emission")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_bus_get_pollfd
//
// [ fd ] trans: nothing
//
func (v Bus) GetPollfd(fd g.PollFD) {
	iv, err := _I.Get(123, "Bus", "get_pollfd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// gst_bus_have_pending
//
// [ result ] trans: nothing
//
func (v Bus) HavePending() (result bool) {
	iv, err := _I.Get(124, "Bus", "have_pending")
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

// gst_bus_peek
//
// [ result ] trans: everything
//
func (v Bus) Peek() (result Message) {
	iv, err := _I.Get(125, "Bus", "peek")
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

// gst_bus_poll
//
// [ events ] trans: nothing
//
// [ timeout ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bus) PollF(events MessageTypeFlags, timeout uint64) (result Message) {
	iv, err := _I.Get(126, "Bus", "poll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_events := gi.NewIntArgument(int(events))
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_events, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bus_pop
//
// [ result ] trans: everything
//
func (v Bus) Pop() (result Message) {
	iv, err := _I.Get(127, "Bus", "pop")
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

// gst_bus_pop_filtered
//
// [ types ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bus) PopFiltered(types MessageTypeFlags) (result Message) {
	iv, err := _I.Get(128, "Bus", "pop_filtered")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_types := gi.NewIntArgument(int(types))
	args := []gi.Argument{arg_v, arg_types}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bus_post
//
// [ message ] trans: everything
//
// [ result ] trans: nothing
//
func (v Bus) Post(message Message) (result bool) {
	iv, err := _I.Get(129, "Bus", "post")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_message := gi.NewPointerArgument(message.P)
	args := []gi.Argument{arg_v, arg_message}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_bus_remove_signal_watch
//
func (v Bus) RemoveSignalWatch() {
	iv, err := _I.Get(130, "Bus", "remove_signal_watch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_bus_remove_watch
//
// [ result ] trans: nothing
//
func (v Bus) RemoveWatch() (result bool) {
	iv, err := _I.Get(131, "Bus", "remove_watch")
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

// gst_bus_set_flushing
//
// [ flushing ] trans: nothing
//
func (v Bus) SetFlushing(flushing bool) {
	iv, err := _I.Get(132, "Bus", "set_flushing")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flushing := gi.NewBoolArgument(flushing)
	args := []gi.Argument{arg_v, arg_flushing}
	iv.Call(args, nil, nil)
}

// gst_bus_set_sync_handler
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Bus) SetSyncHandler(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(133, "Bus", "set_sync_handler")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myBusSyncHandler()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_func1, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_bus_sync_signal_handler
//
// [ message ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Bus) SyncSignalHandler(message Message, data unsafe.Pointer) (result BusSyncReplyEnum) {
	iv, err := _I.Get(134, "Bus", "sync_signal_handler")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_message := gi.NewPointerArgument(message.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_message, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = BusSyncReplyEnum(ret.Int())
	return
}

// gst_bus_timed_pop
//
// [ timeout ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bus) TimedPop(timeout uint64) (result Message) {
	iv, err := _I.Get(135, "Bus", "timed_pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_bus_timed_pop_filtered
//
// [ timeout ] trans: nothing
//
// [ types ] trans: nothing
//
// [ result ] trans: everything
//
func (v Bus) TimedPopFiltered(timeout uint64, types MessageTypeFlags) (result Message) {
	iv, err := _I.Get(136, "Bus", "timed_pop_filtered")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	arg_types := gi.NewIntArgument(int(types))
	args := []gi.Argument{arg_v, arg_timeout, arg_types}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct BusClass

// Flags BusFlags
type BusFlags int

const (
	BusFlagsFlushing BusFlags = 16
	BusFlagsFlagLast BusFlags = 32
)

func BusFlagsGetType() gi.GType {
	ret := _I.GetGType(19, "BusFlags")
	return ret
}

type BusFuncStruct struct {
	F_bus     Bus
	F_message Message
}

func GetPointer_myBusFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBusFunc())
}

//export myGstBusFunc
func myGstBusFunc(bus *C.GstBus, message *C.GstMessage, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &BusFuncStruct{
		F_bus:     WrapBus(unsafe.Pointer(bus)),
		F_message: Message{P: unsafe.Pointer(message)},
	}
	fn(args)
}

// Struct BusPrivate
type BusPrivate struct {
	P unsafe.Pointer
}

func BusPrivateGetType() gi.GType {
	ret := _I.GetGType(20, "BusPrivate")
	return ret
}

type BusSyncHandlerStruct struct {
	F_bus     Bus
	F_message Message
}

func GetPointer_myBusSyncHandler() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstBusSyncHandler())
}

//export myGstBusSyncHandler
func myGstBusSyncHandler(bus *C.GstBus, message *C.GstMessage, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &BusSyncHandlerStruct{
		F_bus:     WrapBus(unsafe.Pointer(bus)),
		F_message: Message{P: unsafe.Pointer(message)},
	}
	fn(args)
}

// Enum BusSyncReply
type BusSyncReplyEnum int

const (
	BusSyncReplyDrop  BusSyncReplyEnum = 0
	BusSyncReplyPass  BusSyncReplyEnum = 1
	BusSyncReplyAsync BusSyncReplyEnum = 2
)

func BusSyncReplyGetType() gi.GType {
	ret := _I.GetGType(21, "BusSyncReply")
	return ret
}

// Struct Caps
type Caps struct {
	P unsafe.Pointer
}

const SizeOfStructCaps = 64

func CapsGetType() gi.GType {
	ret := _I.GetGType(22, "Caps")
	return ret
}

// gst_caps_new_any
//
// [ result ] trans: everything
//
func NewCapsAny() (result Caps) {
	iv, err := _I.Get(137, "Caps", "new_any")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_new_empty
//
// [ result ] trans: everything
//
func NewCapsEmpty() (result Caps) {
	iv, err := _I.Get(138, "Caps", "new_empty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_new_empty_simple
//
// [ media_type ] trans: nothing
//
// [ result ] trans: everything
//
func NewCapsEmptySimple(media_type string) (result Caps) {
	iv, err := _I.Get(139, "Caps", "new_empty_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_media_type := gi.CString(media_type)
	arg_media_type := gi.NewStringArgument(c_media_type)
	args := []gi.Argument{arg_media_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_media_type)
	result.P = ret.Pointer()
	return
}

// gst_caps_append
//
// [ caps2 ] trans: everything
//
func (v Caps) Append(caps2 Caps) {
	iv, err := _I.Get(140, "Caps", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	iv.Call(args, nil, nil)
}

// gst_caps_append_structure
//
// [ structure ] trans: everything
//
func (v Caps) AppendStructure(structure Structure) {
	iv, err := _I.Get(141, "Caps", "append_structure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_v, arg_structure}
	iv.Call(args, nil, nil)
}

// gst_caps_append_structure_full
//
// [ structure ] trans: everything
//
// [ features ] trans: everything
//
func (v Caps) AppendStructureFull(structure Structure, features CapsFeatures) {
	iv, err := _I.Get(142, "Caps", "append_structure_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	arg_features := gi.NewPointerArgument(features.P)
	args := []gi.Argument{arg_v, arg_structure, arg_features}
	iv.Call(args, nil, nil)
}

// gst_caps_can_intersect
//
// [ caps2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) CanIntersect(caps2 Caps) (result bool) {
	iv, err := _I.Get(143, "Caps", "can_intersect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_copy_nth
//
// [ nth ] trans: nothing
//
// [ result ] trans: everything
//
func (v Caps) CopyNth(nth uint32) (result Caps) {
	iv, err := _I.Get(144, "Caps", "copy_nth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nth := gi.NewUint32Argument(nth)
	args := []gi.Argument{arg_v, arg_nth}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_filter_and_map_in_place
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v Caps) FilterAndMapInPlace(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(145, "Caps", "filter_and_map_in_place")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCapsFilterMapFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_caps_fixate
//
// [ result ] trans: everything
//
func (v Caps) Fixate() (result Caps) {
	iv, err := _I.Get(146, "Caps", "fixate")
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

// gst_caps_foreach
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) Foreach(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(147, "Caps", "foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCapsForeachFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_get_features
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) GetFeatures(index uint32) (result CapsFeatures) {
	iv, err := _I.Get(148, "Caps", "get_features")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_get_size
//
// [ result ] trans: nothing
//
func (v Caps) GetSize() (result uint32) {
	iv, err := _I.Get(149, "Caps", "get_size")
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

// gst_caps_get_structure
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) GetStructure(index uint32) (result Structure) {
	iv, err := _I.Get(150, "Caps", "get_structure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_intersect
//
// [ caps2 ] trans: nothing
//
// [ result ] trans: everything
//
func (v Caps) Intersect(caps2 Caps) (result Caps) {
	iv, err := _I.Get(151, "Caps", "intersect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_intersect_full
//
// [ caps2 ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ result ] trans: everything
//
func (v Caps) IntersectFull(caps2 Caps, mode CapsIntersectModeEnum) (result Caps) {
	iv, err := _I.Get(152, "Caps", "intersect_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_caps2, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_is_always_compatible
//
// [ caps2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsAlwaysCompatible(caps2 Caps) (result bool) {
	iv, err := _I.Get(153, "Caps", "is_always_compatible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_is_any
//
// [ result ] trans: nothing
//
func (v Caps) IsAny() (result bool) {
	iv, err := _I.Get(154, "Caps", "is_any")
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

// gst_caps_is_empty
//
// [ result ] trans: nothing
//
func (v Caps) IsEmpty() (result bool) {
	iv, err := _I.Get(155, "Caps", "is_empty")
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

// gst_caps_is_equal
//
// [ caps2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsEqual(caps2 Caps) (result bool) {
	iv, err := _I.Get(156, "Caps", "is_equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_is_equal_fixed
//
// [ caps2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsEqualFixed(caps2 Caps) (result bool) {
	iv, err := _I.Get(157, "Caps", "is_equal_fixed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_is_fixed
//
// [ result ] trans: nothing
//
func (v Caps) IsFixed() (result bool) {
	iv, err := _I.Get(158, "Caps", "is_fixed")
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

// gst_caps_is_strictly_equal
//
// [ caps2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsStrictlyEqual(caps2 Caps) (result bool) {
	iv, err := _I.Get(159, "Caps", "is_strictly_equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_is_subset
//
// [ superset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsSubset(superset Caps) (result bool) {
	iv, err := _I.Get(160, "Caps", "is_subset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_superset := gi.NewPointerArgument(superset.P)
	args := []gi.Argument{arg_v, arg_superset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_is_subset_structure
//
// [ structure ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsSubsetStructure(structure Structure) (result bool) {
	iv, err := _I.Get(161, "Caps", "is_subset_structure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_v, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_is_subset_structure_full
//
// [ structure ] trans: nothing
//
// [ features ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) IsSubsetStructureFull(structure Structure, features CapsFeatures) (result bool) {
	iv, err := _I.Get(162, "Caps", "is_subset_structure_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	arg_features := gi.NewPointerArgument(features.P)
	args := []gi.Argument{arg_v, arg_structure, arg_features}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_map_in_place
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Caps) MapInPlace(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(163, "Caps", "map_in_place")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCapsMapFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_merge
//
// [ caps2 ] trans: everything
//
// [ result ] trans: everything
//
func (v Caps) Merge(caps2 Caps) (result Caps) {
	iv, err := _I.Get(164, "Caps", "merge")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps2 := gi.NewPointerArgument(caps2.P)
	args := []gi.Argument{arg_v, arg_caps2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_merge_structure
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func (v Caps) MergeStructure(structure Structure) (result Caps) {
	iv, err := _I.Get(165, "Caps", "merge_structure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_v, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_merge_structure_full
//
// [ structure ] trans: everything
//
// [ features ] trans: everything
//
// [ result ] trans: everything
//
func (v Caps) MergeStructureFull(structure Structure, features CapsFeatures) (result Caps) {
	iv, err := _I.Get(166, "Caps", "merge_structure_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	arg_features := gi.NewPointerArgument(features.P)
	args := []gi.Argument{arg_v, arg_structure, arg_features}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_normalize
//
// [ result ] trans: everything
//
func (v Caps) Normalize() (result Caps) {
	iv, err := _I.Get(167, "Caps", "normalize")
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

// gst_caps_remove_structure
//
// [ idx ] trans: nothing
//
func (v Caps) RemoveStructure(idx uint32) {
	iv, err := _I.Get(168, "Caps", "remove_structure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	iv.Call(args, nil, nil)
}

// gst_caps_set_features
//
// [ index ] trans: nothing
//
// [ features ] trans: everything
//
func (v Caps) SetFeatures(index uint32, features CapsFeatures) {
	iv, err := _I.Get(169, "Caps", "set_features")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_features := gi.NewPointerArgument(features.P)
	args := []gi.Argument{arg_v, arg_index, arg_features}
	iv.Call(args, nil, nil)
}

// gst_caps_set_value
//
// [ field ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Caps) SetValue(field string, value g.Value) {
	iv, err := _I.Get(170, "Caps", "set_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field := gi.CString(field)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field := gi.NewStringArgument(c_field)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_field, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_field)
}

// gst_caps_simplify
//
// [ result ] trans: everything
//
func (v Caps) Simplify() (result Caps) {
	iv, err := _I.Get(171, "Caps", "simplify")
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

// gst_caps_steal_structure
//
// [ index ] trans: nothing
//
// [ result ] trans: everything
//
func (v Caps) StealStructure(index uint32) (result Structure) {
	iv, err := _I.Get(172, "Caps", "steal_structure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_subtract
//
// [ subtrahend ] trans: nothing
//
// [ result ] trans: everything
//
func (v Caps) Subtract(subtrahend Caps) (result Caps) {
	iv, err := _I.Get(173, "Caps", "subtract")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_subtrahend := gi.NewPointerArgument(subtrahend.P)
	args := []gi.Argument{arg_v, arg_subtrahend}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_to_string
//
// [ result ] trans: everything
//
func (v Caps) ToString() (result string) {
	iv, err := _I.Get(174, "Caps", "to_string")
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

// gst_caps_truncate
//
// [ result ] trans: everything
//
func (v Caps) Truncate() (result Caps) {
	iv, err := _I.Get(175, "Caps", "truncate")
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

// gst_caps_from_string
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func CapsFromString1(string string) (result Caps) {
	iv, err := _I.Get(176, "Caps", "from_string")
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

// Struct CapsFeatures
type CapsFeatures struct {
	P unsafe.Pointer
}

func CapsFeaturesGetType() gi.GType {
	ret := _I.GetGType(23, "CapsFeatures")
	return ret
}

// gst_caps_features_new_any
//
// [ result ] trans: everything
//
func NewCapsFeaturesAny() (result CapsFeatures) {
	iv, err := _I.Get(177, "CapsFeatures", "new_any")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_features_new_empty
//
// [ result ] trans: everything
//
func NewCapsFeaturesEmpty() (result CapsFeatures) {
	iv, err := _I.Get(178, "CapsFeatures", "new_empty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_caps_features_add
//
// [ feature ] trans: nothing
//
func (v CapsFeatures) Add(feature string) {
	iv, err := _I.Get(179, "CapsFeatures", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_feature := gi.CString(feature)
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewStringArgument(c_feature)
	args := []gi.Argument{arg_v, arg_feature}
	iv.Call(args, nil, nil)
	gi.Free(c_feature)
}

// gst_caps_features_add_id
//
// [ feature ] trans: nothing
//
func (v CapsFeatures) AddId(feature uint32) {
	iv, err := _I.Get(180, "CapsFeatures", "add_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewUint32Argument(feature)
	args := []gi.Argument{arg_v, arg_feature}
	iv.Call(args, nil, nil)
}

// gst_caps_features_contains
//
// [ feature ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CapsFeatures) Contains(feature string) (result bool) {
	iv, err := _I.Get(181, "CapsFeatures", "contains")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_feature := gi.CString(feature)
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewStringArgument(c_feature)
	args := []gi.Argument{arg_v, arg_feature}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_feature)
	result = ret.Bool()
	return
}

// gst_caps_features_contains_id
//
// [ feature ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CapsFeatures) ContainsId(feature uint32) (result bool) {
	iv, err := _I.Get(182, "CapsFeatures", "contains_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewUint32Argument(feature)
	args := []gi.Argument{arg_v, arg_feature}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_features_copy
//
// [ result ] trans: everything
//
func (v CapsFeatures) Copy() (result CapsFeatures) {
	iv, err := _I.Get(183, "CapsFeatures", "copy")
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

// gst_caps_features_free
//
func (v CapsFeatures) Free() {
	iv, err := _I.Get(184, "CapsFeatures", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_caps_features_get_nth
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CapsFeatures) GetNth(i uint32) (result string) {
	iv, err := _I.Get(185, "CapsFeatures", "get_nth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewUint32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_caps_features_get_nth_id
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CapsFeatures) GetNthId(i uint32) (result uint32) {
	iv, err := _I.Get(186, "CapsFeatures", "get_nth_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewUint32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_caps_features_get_size
//
// [ result ] trans: nothing
//
func (v CapsFeatures) GetSize() (result uint32) {
	iv, err := _I.Get(187, "CapsFeatures", "get_size")
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

// gst_caps_features_is_any
//
// [ result ] trans: nothing
//
func (v CapsFeatures) IsAny() (result bool) {
	iv, err := _I.Get(188, "CapsFeatures", "is_any")
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

// gst_caps_features_is_equal
//
// [ features2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CapsFeatures) IsEqual(features2 CapsFeatures) (result bool) {
	iv, err := _I.Get(189, "CapsFeatures", "is_equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_features2 := gi.NewPointerArgument(features2.P)
	args := []gi.Argument{arg_v, arg_features2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_features_remove
//
// [ feature ] trans: nothing
//
func (v CapsFeatures) Remove(feature string) {
	iv, err := _I.Get(190, "CapsFeatures", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_feature := gi.CString(feature)
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewStringArgument(c_feature)
	args := []gi.Argument{arg_v, arg_feature}
	iv.Call(args, nil, nil)
	gi.Free(c_feature)
}

// gst_caps_features_remove_id
//
// [ feature ] trans: nothing
//
func (v CapsFeatures) RemoveId(feature uint32) {
	iv, err := _I.Get(191, "CapsFeatures", "remove_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewUint32Argument(feature)
	args := []gi.Argument{arg_v, arg_feature}
	iv.Call(args, nil, nil)
}

// gst_caps_features_set_parent_refcount
//
// [ refcount ] trans: nothing
//
// [ result ] trans: nothing
//
func (v CapsFeatures) SetParentRefcount(refcount int32) (result bool) {
	iv, err := _I.Get(192, "CapsFeatures", "set_parent_refcount")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_refcount := gi.NewInt32Argument(refcount)
	args := []gi.Argument{arg_v, arg_refcount}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_caps_features_to_string
//
// [ result ] trans: everything
//
func (v CapsFeatures) ToString() (result string) {
	iv, err := _I.Get(193, "CapsFeatures", "to_string")
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

// gst_caps_features_from_string
//
// [ features ] trans: nothing
//
// [ result ] trans: everything
//
func CapsFeaturesFromString1(features string) (result CapsFeatures) {
	iv, err := _I.Get(194, "CapsFeatures", "from_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_features := gi.CString(features)
	arg_features := gi.NewStringArgument(c_features)
	args := []gi.Argument{arg_features}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_features)
	result.P = ret.Pointer()
	return
}

type CapsFilterMapFuncStruct struct {
	F_features  CapsFeatures
	F_structure Structure
}

func GetPointer_myCapsFilterMapFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstCapsFilterMapFunc())
}

//export myGstCapsFilterMapFunc
func myGstCapsFilterMapFunc(features *C.GstCapsFeatures, structure *C.GstStructure, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CapsFilterMapFuncStruct{
		F_features:  CapsFeatures{P: unsafe.Pointer(features)},
		F_structure: Structure{P: unsafe.Pointer(structure)},
	}
	fn(args)
}

// Flags CapsFlags
type CapsFlags int

const (
	CapsFlagsAny CapsFlags = 16
)

func CapsFlagsGetType() gi.GType {
	ret := _I.GetGType(24, "CapsFlags")
	return ret
}

type CapsForeachFuncStruct struct {
	F_features  CapsFeatures
	F_structure Structure
}

func GetPointer_myCapsForeachFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstCapsForeachFunc())
}

//export myGstCapsForeachFunc
func myGstCapsForeachFunc(features *C.GstCapsFeatures, structure *C.GstStructure, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CapsForeachFuncStruct{
		F_features:  CapsFeatures{P: unsafe.Pointer(features)},
		F_structure: Structure{P: unsafe.Pointer(structure)},
	}
	fn(args)
}

// Enum CapsIntersectMode
type CapsIntersectModeEnum int

const (
	CapsIntersectModeZigZag CapsIntersectModeEnum = 0
	CapsIntersectModeFirst  CapsIntersectModeEnum = 1
)

func CapsIntersectModeGetType() gi.GType {
	ret := _I.GetGType(25, "CapsIntersectMode")
	return ret
}

type CapsMapFuncStruct struct {
	F_features  CapsFeatures
	F_structure Structure
}

func GetPointer_myCapsMapFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstCapsMapFunc())
}

//export myGstCapsMapFunc
func myGstCapsMapFunc(features *C.GstCapsFeatures, structure *C.GstStructure, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &CapsMapFuncStruct{
		F_features:  CapsFeatures{P: unsafe.Pointer(features)},
		F_structure: Structure{P: unsafe.Pointer(structure)},
	}
	fn(args)
}

// Interface ChildProxy
type ChildProxy struct {
	ChildProxyIfc
	P unsafe.Pointer
}
type ChildProxyIfc struct{}
type IChildProxy interface{ P_ChildProxy() unsafe.Pointer }

func (v ChildProxy) P_ChildProxy() unsafe.Pointer { return v.P }
func ChildProxyGetType() gi.GType {
	ret := _I.GetGType(26, "ChildProxy")
	return ret
}

// gst_child_proxy_child_added
//
// [ child ] trans: nothing
//
// [ name ] trans: nothing
//
func (v *ChildProxyIfc) ChildAdded(child g.IObject, name string) {
	iv, err := _I.Get(195, "ChildProxy", "child_added")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if child != nil {
		tmp = child.P_Object()
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_child := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_child, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gst_child_proxy_child_removed
//
// [ child ] trans: nothing
//
// [ name ] trans: nothing
//
func (v *ChildProxyIfc) ChildRemoved(child g.IObject, name string) {
	iv, err := _I.Get(196, "ChildProxy", "child_removed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if child != nil {
		tmp = child.P_Object()
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_child := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_child, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gst_child_proxy_get_child_by_index
//
// [ index ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ChildProxyIfc) GetChildByIndex(index uint32) (result g.Object) {
	iv, err := _I.Get(197, "ChildProxy", "get_child_by_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_child_proxy_get_child_by_name
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ChildProxyIfc) GetChildByName(name string) (result g.Object) {
	iv, err := _I.Get(198, "ChildProxy", "get_child_by_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_child_proxy_get_children_count
//
// [ result ] trans: nothing
//
func (v *ChildProxyIfc) GetChildrenCount() (result uint32) {
	iv, err := _I.Get(199, "ChildProxy", "get_children_count")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_child_proxy_get_property
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing, dir: out
//
func (v *ChildProxyIfc) GetProperty(name string, value g.Value) {
	iv, err := _I.Get(200, "ChildProxy", "get_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_name, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gst_child_proxy_lookup
//
// [ name ] trans: nothing
//
// [ target ] trans: everything, dir: out
//
// [ pspec ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v *ChildProxyIfc) Lookup(name string) (result bool, target g.Object, pspec g.ParamSpec) {
	iv, err := _I.Get(201, "ChildProxy", "lookup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	arg_target := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_pspec := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_name, arg_target, arg_pspec}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	target.P = outArgs[0].Pointer()
	pspec.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// gst_child_proxy_set_property
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
func (v *ChildProxyIfc) SetProperty(name string, value g.Value) {
	iv, err := _I.Get(202, "ChildProxy", "set_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_name, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// ignore GType struct ChildProxyInterface

// Object Clock
type Clock struct {
	Object
}

func WrapClock(p unsafe.Pointer) (r Clock) { r.P = p; return }

type IClock interface{ P_Clock() unsafe.Pointer }

func (v Clock) P_Clock() unsafe.Pointer { return v.P }
func ClockGetType() gi.GType {
	ret := _I.GetGType(27, "Clock")
	return ret
}

// gst_clock_id_compare_func
//
// [ id1 ] trans: nothing
//
// [ id2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ClockIdCompareFunc1(id1 unsafe.Pointer, id2 unsafe.Pointer) (result int32) {
	iv, err := _I.Get(203, "Clock", "id_compare_func")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id1 := gi.NewPointerArgument(id1)
	arg_id2 := gi.NewPointerArgument(id2)
	args := []gi.Argument{arg_id1, arg_id2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_clock_id_get_time
//
// [ id ] trans: nothing
//
// [ result ] trans: nothing
//
func ClockIdGetTime1(id unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(204, "Clock", "id_get_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id := gi.NewPointerArgument(id)
	args := []gi.Argument{arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_clock_id_ref
//
// [ id ] trans: nothing
//
// [ result ] trans: everything
//
func ClockIdRef1(id unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(205, "Clock", "id_ref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id := gi.NewPointerArgument(id)
	args := []gi.Argument{arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// gst_clock_id_unref
//
// [ id ] trans: everything
//
func ClockIdUnref1(id unsafe.Pointer) {
	iv, err := _I.Get(206, "Clock", "id_unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id := gi.NewPointerArgument(id)
	args := []gi.Argument{arg_id}
	iv.Call(args, nil, nil)
}

// gst_clock_id_unschedule
//
// [ id ] trans: nothing
//
func ClockIdUnschedule1(id unsafe.Pointer) {
	iv, err := _I.Get(207, "Clock", "id_unschedule")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id := gi.NewPointerArgument(id)
	args := []gi.Argument{arg_id}
	iv.Call(args, nil, nil)
}

// gst_clock_id_wait
//
// [ id ] trans: nothing
//
// [ jitter ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func ClockIdWait1(id unsafe.Pointer) (result ClockReturnEnum, jitter int64) {
	iv, err := _I.Get(208, "Clock", "id_wait")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_id := gi.NewPointerArgument(id)
	arg_jitter := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_id, arg_jitter}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	jitter = outArgs[0].Int64()
	result = ClockReturnEnum(ret.Int())
	return
}

// gst_clock_id_wait_async
//
// [ id ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ destroy_data ] trans: nothing
//
// [ result ] trans: nothing
//
func ClockIdWaitAsync1(id unsafe.Pointer, func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, destroy_data int /*TODO_TYPE CALLBACK*/) (result ClockReturnEnum) {
	iv, err := _I.Get(209, "Clock", "id_wait_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id := gi.NewPointerArgument(id)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myClockCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy_data := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_id, arg_func1, arg_user_data, arg_destroy_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ClockReturnEnum(ret.Int())
	return
}

// gst_clock_add_observation
//
// [ slave ] trans: nothing
//
// [ master ] trans: nothing
//
// [ r_squared ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Clock) AddObservation(slave uint64, master uint64) (result bool, r_squared float64) {
	iv, err := _I.Get(210, "Clock", "add_observation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_slave := gi.NewUint64Argument(slave)
	arg_master := gi.NewUint64Argument(master)
	arg_r_squared := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_slave, arg_master, arg_r_squared}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	r_squared = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_clock_add_observation_unapplied
//
// [ slave ] trans: nothing
//
// [ master ] trans: nothing
//
// [ r_squared ] trans: everything, dir: out
//
// [ internal ] trans: everything, dir: out
//
// [ external ] trans: everything, dir: out
//
// [ rate_num ] trans: everything, dir: out
//
// [ rate_denom ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Clock) AddObservationUnapplied(slave uint64, master uint64) (result bool, r_squared float64, internal uint64, external uint64, rate_num uint64, rate_denom uint64) {
	iv, err := _I.Get(211, "Clock", "add_observation_unapplied")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [5]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_slave := gi.NewUint64Argument(slave)
	arg_master := gi.NewUint64Argument(master)
	arg_r_squared := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_internal := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_external := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_rate_num := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_rate_denom := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_v, arg_slave, arg_master, arg_r_squared, arg_internal, arg_external, arg_rate_num, arg_rate_denom}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	r_squared = outArgs[0].Double()
	internal = outArgs[1].Uint64()
	external = outArgs[2].Uint64()
	rate_num = outArgs[3].Uint64()
	rate_denom = outArgs[4].Uint64()
	result = ret.Bool()
	return
}

// gst_clock_adjust_unlocked
//
// [ internal ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) AdjustUnlocked(internal uint64) (result uint64) {
	iv, err := _I.Get(212, "Clock", "adjust_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_internal := gi.NewUint64Argument(internal)
	args := []gi.Argument{arg_v, arg_internal}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_clock_adjust_with_calibration
//
// [ internal_target ] trans: nothing
//
// [ cinternal ] trans: nothing
//
// [ cexternal ] trans: nothing
//
// [ cnum ] trans: nothing
//
// [ cdenom ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) AdjustWithCalibration(internal_target uint64, cinternal uint64, cexternal uint64, cnum uint64, cdenom uint64) (result uint64) {
	iv, err := _I.Get(213, "Clock", "adjust_with_calibration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_internal_target := gi.NewUint64Argument(internal_target)
	arg_cinternal := gi.NewUint64Argument(cinternal)
	arg_cexternal := gi.NewUint64Argument(cexternal)
	arg_cnum := gi.NewUint64Argument(cnum)
	arg_cdenom := gi.NewUint64Argument(cdenom)
	args := []gi.Argument{arg_v, arg_internal_target, arg_cinternal, arg_cexternal, arg_cnum, arg_cdenom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_clock_get_calibration
//
// [ internal ] trans: everything, dir: out
//
// [ external ] trans: everything, dir: out
//
// [ rate_num ] trans: everything, dir: out
//
// [ rate_denom ] trans: everything, dir: out
//
func (v Clock) GetCalibration() (internal uint64, external uint64, rate_num uint64, rate_denom uint64) {
	iv, err := _I.Get(214, "Clock", "get_calibration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_internal := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_external := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_rate_num := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_rate_denom := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_internal, arg_external, arg_rate_num, arg_rate_denom}
	iv.Call(args, nil, &outArgs[0])
	internal = outArgs[0].Uint64()
	external = outArgs[1].Uint64()
	rate_num = outArgs[2].Uint64()
	rate_denom = outArgs[3].Uint64()
	return
}

// gst_clock_get_internal_time
//
// [ result ] trans: nothing
//
func (v Clock) GetInternalTime() (result uint64) {
	iv, err := _I.Get(215, "Clock", "get_internal_time")
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

// gst_clock_get_master
//
// [ result ] trans: everything
//
func (v Clock) GetMaster() (result Clock) {
	iv, err := _I.Get(216, "Clock", "get_master")
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

// gst_clock_get_resolution
//
// [ result ] trans: nothing
//
func (v Clock) GetResolution() (result uint64) {
	iv, err := _I.Get(217, "Clock", "get_resolution")
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

// gst_clock_get_time
//
// [ result ] trans: nothing
//
func (v Clock) GetTime() (result uint64) {
	iv, err := _I.Get(218, "Clock", "get_time")
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

// gst_clock_get_timeout
//
// [ result ] trans: nothing
//
func (v Clock) GetTimeout() (result uint64) {
	iv, err := _I.Get(219, "Clock", "get_timeout")
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

// gst_clock_is_synced
//
// [ result ] trans: nothing
//
func (v Clock) IsSynced() (result bool) {
	iv, err := _I.Get(220, "Clock", "is_synced")
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

// gst_clock_new_periodic_id
//
// [ start_time ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ result ] trans: everything
//
func (v Clock) NewPeriodicId(start_time uint64, interval uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(221, "Clock", "new_periodic_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start_time := gi.NewUint64Argument(start_time)
	arg_interval := gi.NewUint64Argument(interval)
	args := []gi.Argument{arg_v, arg_start_time, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// gst_clock_new_single_shot_id
//
// [ time ] trans: nothing
//
// [ result ] trans: everything
//
func (v Clock) NewSingleShotId(time uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(222, "Clock", "new_single_shot_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time := gi.NewUint64Argument(time)
	args := []gi.Argument{arg_v, arg_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// gst_clock_periodic_id_reinit
//
// [ id ] trans: nothing
//
// [ start_time ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) PeriodicIdReinit(id unsafe.Pointer, start_time uint64, interval uint64) (result bool) {
	iv, err := _I.Get(223, "Clock", "periodic_id_reinit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewPointerArgument(id)
	arg_start_time := gi.NewUint64Argument(start_time)
	arg_interval := gi.NewUint64Argument(interval)
	args := []gi.Argument{arg_v, arg_id, arg_start_time, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_clock_set_calibration
//
// [ internal ] trans: nothing
//
// [ external ] trans: nothing
//
// [ rate_num ] trans: nothing
//
// [ rate_denom ] trans: nothing
//
func (v Clock) SetCalibration(internal uint64, external uint64, rate_num uint64, rate_denom uint64) {
	iv, err := _I.Get(224, "Clock", "set_calibration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_internal := gi.NewUint64Argument(internal)
	arg_external := gi.NewUint64Argument(external)
	arg_rate_num := gi.NewUint64Argument(rate_num)
	arg_rate_denom := gi.NewUint64Argument(rate_denom)
	args := []gi.Argument{arg_v, arg_internal, arg_external, arg_rate_num, arg_rate_denom}
	iv.Call(args, nil, nil)
}

// gst_clock_set_master
//
// [ master ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) SetMaster(master IClock) (result bool) {
	iv, err := _I.Get(225, "Clock", "set_master")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if master != nil {
		tmp = master.P_Clock()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_master := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_master}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_clock_set_resolution
//
// [ resolution ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) SetResolution(resolution uint64) (result uint64) {
	iv, err := _I.Get(226, "Clock", "set_resolution")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_resolution := gi.NewUint64Argument(resolution)
	args := []gi.Argument{arg_v, arg_resolution}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_clock_set_synced
//
// [ synced ] trans: nothing
//
func (v Clock) SetSynced(synced bool) {
	iv, err := _I.Get(227, "Clock", "set_synced")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_synced := gi.NewBoolArgument(synced)
	args := []gi.Argument{arg_v, arg_synced}
	iv.Call(args, nil, nil)
}

// gst_clock_set_timeout
//
// [ timeout ] trans: nothing
//
func (v Clock) SetTimeout(timeout uint64) {
	iv, err := _I.Get(228, "Clock", "set_timeout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	iv.Call(args, nil, nil)
}

// gst_clock_single_shot_id_reinit
//
// [ id ] trans: nothing
//
// [ time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) SingleShotIdReinit(id unsafe.Pointer, time uint64) (result bool) {
	iv, err := _I.Get(229, "Clock", "single_shot_id_reinit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewPointerArgument(id)
	arg_time := gi.NewUint64Argument(time)
	args := []gi.Argument{arg_v, arg_id, arg_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_clock_unadjust_unlocked
//
// [ external ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) UnadjustUnlocked(external uint64) (result uint64) {
	iv, err := _I.Get(230, "Clock", "unadjust_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_external := gi.NewUint64Argument(external)
	args := []gi.Argument{arg_v, arg_external}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_clock_unadjust_with_calibration
//
// [ external_target ] trans: nothing
//
// [ cinternal ] trans: nothing
//
// [ cexternal ] trans: nothing
//
// [ cnum ] trans: nothing
//
// [ cdenom ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) UnadjustWithCalibration(external_target uint64, cinternal uint64, cexternal uint64, cnum uint64, cdenom uint64) (result uint64) {
	iv, err := _I.Get(231, "Clock", "unadjust_with_calibration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_external_target := gi.NewUint64Argument(external_target)
	arg_cinternal := gi.NewUint64Argument(cinternal)
	arg_cexternal := gi.NewUint64Argument(cexternal)
	arg_cnum := gi.NewUint64Argument(cnum)
	arg_cdenom := gi.NewUint64Argument(cdenom)
	args := []gi.Argument{arg_v, arg_external_target, arg_cinternal, arg_cexternal, arg_cnum, arg_cdenom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_clock_wait_for_sync
//
// [ timeout ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Clock) WaitForSync(timeout uint64) (result bool) {
	iv, err := _I.Get(232, "Clock", "wait_for_sync")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

type ClockCallbackStruct struct {
	F_clock Clock
	F_time  uint64
	F_id    unsafe.Pointer
}

func GetPointer_myClockCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstClockCallback())
}

//export myGstClockCallback
func myGstClockCallback(clock *C.GstClock, time C.guint64, id C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &ClockCallbackStruct{
		F_clock: WrapClock(unsafe.Pointer(clock)),
		F_time:  uint64(time),
		F_id:    unsafe.Pointer(id),
	}
	fn(args)
}

// ignore GType struct ClockClass

// Struct ClockEntry
type ClockEntry struct {
	P unsafe.Pointer
}

const SizeOfStructClockEntry = 112

func ClockEntryGetType() gi.GType {
	ret := _I.GetGType(28, "ClockEntry")
	return ret
}

// Enum ClockEntryType
type ClockEntryTypeEnum int

const (
	ClockEntryTypeSingle   ClockEntryTypeEnum = 0
	ClockEntryTypePeriodic ClockEntryTypeEnum = 1
)

func ClockEntryTypeGetType() gi.GType {
	ret := _I.GetGType(29, "ClockEntryType")
	return ret
}

// Flags ClockFlags
type ClockFlags int

const (
	ClockFlagsCanDoSingleSync    ClockFlags = 16
	ClockFlagsCanDoSingleAsync   ClockFlags = 32
	ClockFlagsCanDoPeriodicSync  ClockFlags = 64
	ClockFlagsCanDoPeriodicAsync ClockFlags = 128
	ClockFlagsCanSetResolution   ClockFlags = 256
	ClockFlagsCanSetMaster       ClockFlags = 512
	ClockFlagsNeedsStartupSync   ClockFlags = 1024
	ClockFlagsLast               ClockFlags = 4096
)

func ClockFlagsGetType() gi.GType {
	ret := _I.GetGType(30, "ClockFlags")
	return ret
}

// Struct ClockPrivate
type ClockPrivate struct {
	P unsafe.Pointer
}

func ClockPrivateGetType() gi.GType {
	ret := _I.GetGType(31, "ClockPrivate")
	return ret
}

// Enum ClockReturn
type ClockReturnEnum int

const (
	ClockReturnOk          ClockReturnEnum = 0
	ClockReturnEarly       ClockReturnEnum = 1
	ClockReturnUnscheduled ClockReturnEnum = 2
	ClockReturnBusy        ClockReturnEnum = 3
	ClockReturnBadtime     ClockReturnEnum = 4
	ClockReturnError       ClockReturnEnum = 5
	ClockReturnUnsupported ClockReturnEnum = 6
	ClockReturnDone        ClockReturnEnum = 7
)

func ClockReturnGetType() gi.GType {
	ret := _I.GetGType(32, "ClockReturn")
	return ret
}

// Enum ClockType
type ClockTypeEnum int

const (
	ClockTypeRealtime  ClockTypeEnum = 0
	ClockTypeMonotonic ClockTypeEnum = 1
	ClockTypeOther     ClockTypeEnum = 2
)

func ClockTypeGetType() gi.GType {
	ret := _I.GetGType(33, "ClockType")
	return ret
}

// Struct Context
type Context struct {
	P unsafe.Pointer
}

func ContextGetType() gi.GType {
	ret := _I.GetGType(34, "Context")
	return ret
}

// gst_context_new
//
// [ context_type ] trans: nothing
//
// [ persistent ] trans: nothing
//
// [ result ] trans: everything
//
func NewContext(context_type string, persistent bool) (result Context) {
	iv, err := _I.Get(233, "Context", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_type := gi.CString(context_type)
	arg_context_type := gi.NewStringArgument(c_context_type)
	arg_persistent := gi.NewBoolArgument(persistent)
	args := []gi.Argument{arg_context_type, arg_persistent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_type)
	result.P = ret.Pointer()
	return
}

// gst_context_get_context_type
//
// [ result ] trans: nothing
//
func (v Context) GetContextType() (result string) {
	iv, err := _I.Get(234, "Context", "get_context_type")
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

// gst_context_get_structure
//
// [ result ] trans: nothing
//
func (v Context) GetStructure() (result Structure) {
	iv, err := _I.Get(235, "Context", "get_structure")
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

// gst_context_has_context_type
//
// [ context_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Context) HasContextType(context_type string) (result bool) {
	iv, err := _I.Get(236, "Context", "has_context_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_type := gi.CString(context_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_context_type := gi.NewStringArgument(c_context_type)
	args := []gi.Argument{arg_v, arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_type)
	result = ret.Bool()
	return
}

// gst_context_is_persistent
//
// [ result ] trans: nothing
//
func (v Context) IsPersistent() (result bool) {
	iv, err := _I.Get(237, "Context", "is_persistent")
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

// gst_context_writable_structure
//
// [ result ] trans: everything
//
func (v Context) WritableStructure() (result Structure) {
	iv, err := _I.Get(238, "Context", "writable_structure")
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

// Object ControlBinding
type ControlBinding struct {
	Object
}

func WrapControlBinding(p unsafe.Pointer) (r ControlBinding) { r.P = p; return }

type IControlBinding interface{ P_ControlBinding() unsafe.Pointer }

func (v ControlBinding) P_ControlBinding() unsafe.Pointer { return v.P }
func ControlBindingGetType() gi.GType {
	ret := _I.GetGType(35, "ControlBinding")
	return ret
}

// gst_control_binding_get_g_value_array
//
// [ timestamp ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ n_values ] trans: nothing
//
// [ values ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ControlBinding) GetGValueArray(timestamp uint64, interval uint64, n_values uint32, values unsafe.Pointer) (result bool) {
	iv, err := _I.Get(239, "ControlBinding", "get_g_value_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_interval := gi.NewUint64Argument(interval)
	arg_n_values := gi.NewUint32Argument(n_values)
	arg_values := gi.NewPointerArgument(values)
	args := []gi.Argument{arg_v, arg_timestamp, arg_interval, arg_n_values, arg_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_control_binding_get_value
//
// [ timestamp ] trans: nothing
//
// [ result ] trans: everything
//
func (v ControlBinding) GetValue(timestamp uint64) (result g.Value) {
	iv, err := _I.Get(240, "ControlBinding", "get_value")
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

// gst_control_binding_is_disabled
//
// [ result ] trans: nothing
//
func (v ControlBinding) IsDisabled() (result bool) {
	iv, err := _I.Get(241, "ControlBinding", "is_disabled")
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

// gst_control_binding_set_disabled
//
// [ disabled ] trans: nothing
//
func (v ControlBinding) SetDisabled(disabled bool) {
	iv, err := _I.Get(242, "ControlBinding", "set_disabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_disabled := gi.NewBoolArgument(disabled)
	args := []gi.Argument{arg_v, arg_disabled}
	iv.Call(args, nil, nil)
}

// gst_control_binding_sync_values
//
// [ object ] trans: nothing
//
// [ timestamp ] trans: nothing
//
// [ last_sync ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ControlBinding) SyncValues(object IObject, timestamp uint64, last_sync uint64) (result bool) {
	iv, err := _I.Get(243, "ControlBinding", "sync_values")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_object := gi.NewPointerArgument(tmp)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_last_sync := gi.NewUint64Argument(last_sync)
	args := []gi.Argument{arg_v, arg_object, arg_timestamp, arg_last_sync}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct ControlBindingClass

type ControlBindingConvertStruct struct {
	F_binding    ControlBinding
	F_src_value  float64
	F_dest_value g.Value
}

func GetPointer_myControlBindingConvert() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstControlBindingConvert())
}

//export myGstControlBindingConvert
func myGstControlBindingConvert(binding *C.GstControlBinding, src_value C.gdouble, dest_value *C.GValue) {
	// TODO: not found user_data
}

// Struct ControlBindingPrivate
type ControlBindingPrivate struct {
	P unsafe.Pointer
}

func ControlBindingPrivateGetType() gi.GType {
	ret := _I.GetGType(36, "ControlBindingPrivate")
	return ret
}

// Object ControlSource
type ControlSource struct {
	Object
}

func WrapControlSource(p unsafe.Pointer) (r ControlSource) { r.P = p; return }

type IControlSource interface{ P_ControlSource() unsafe.Pointer }

func (v ControlSource) P_ControlSource() unsafe.Pointer { return v.P }
func ControlSourceGetType() gi.GType {
	ret := _I.GetGType(37, "ControlSource")
	return ret
}

// gst_control_source_get_value
//
// [ timestamp ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v ControlSource) ControlSourceGetValue(timestamp uint64) (result bool, value float64) {
	iv, err := _I.Get(244, "ControlSource", "control_source_get_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_timestamp, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	value = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_control_source_get_value_array
//
// [ timestamp ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ n_values ] trans: nothing
//
// [ values ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ControlSource) ControlSourceGetValueArray(timestamp uint64, interval uint64, n_values uint32, values gi.DoubleArray) (result bool) {
	iv, err := _I.Get(245, "ControlSource", "control_source_get_value_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_interval := gi.NewUint64Argument(interval)
	arg_n_values := gi.NewUint32Argument(n_values)
	arg_values := gi.NewPointerArgument(values.P)
	args := []gi.Argument{arg_v, arg_timestamp, arg_interval, arg_n_values, arg_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct ControlSourceClass

type ControlSourceGetValueStruct struct {
	F_self      ControlSource
	F_timestamp uint64
	F_value     *float64
}

func GetPointer_myControlSourceGetValue() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstControlSourceGetValue())
}

//export myGstControlSourceGetValue
func myGstControlSourceGetValue(self *C.GstControlSource, timestamp C.guint64, value *C.gdouble) {
	// TODO: not found user_data
}

type ControlSourceGetValueArrayStruct struct {
	F_self      ControlSource
	F_timestamp uint64
	F_interval  uint64
	F_n_values  uint32
	F_values    *float64
}

func GetPointer_myControlSourceGetValueArray() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstControlSourceGetValueArray())
}

//export myGstControlSourceGetValueArray
func myGstControlSourceGetValueArray(self *C.GstControlSource, timestamp C.guint64, interval C.guint64, n_values C.guint32, values *C.gdouble) {
	// TODO: not found user_data
}

// Enum CoreError
type CoreErrorEnum int

const (
	CoreErrorFailed         CoreErrorEnum = 1
	CoreErrorTooLazy        CoreErrorEnum = 2
	CoreErrorNotImplemented CoreErrorEnum = 3
	CoreErrorStateChange    CoreErrorEnum = 4
	CoreErrorPad            CoreErrorEnum = 5
	CoreErrorThread         CoreErrorEnum = 6
	CoreErrorNegotiation    CoreErrorEnum = 7
	CoreErrorEvent          CoreErrorEnum = 8
	CoreErrorSeek           CoreErrorEnum = 9
	CoreErrorCaps           CoreErrorEnum = 10
	CoreErrorTag            CoreErrorEnum = 11
	CoreErrorMissingPlugin  CoreErrorEnum = 12
	CoreErrorClock          CoreErrorEnum = 13
	CoreErrorDisabled       CoreErrorEnum = 14
	CoreErrorNumErrors      CoreErrorEnum = 15
)

func CoreErrorGetType() gi.GType {
	ret := _I.GetGType(38, "CoreError")
	return ret
}

// Struct DateTime
type DateTime struct {
	P unsafe.Pointer
}

func DateTimeGetType() gi.GType {
	ret := _I.GetGType(39, "DateTime")
	return ret
}

// gst_date_time_new
//
// [ tzoffset ] trans: nothing
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
func NewDateTime(tzoffset float32, year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(246, "DateTime", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tzoffset := gi.NewFloatArgument(tzoffset)
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_tzoffset, arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_from_g_date_time
//
// [ dt ] trans: everything
//
// [ result ] trans: everything
//
func NewDateTimeFromGDateTime(dt g.DateTime) (result DateTime) {
	iv, err := _I.Get(247, "DateTime", "new_from_g_date_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt := gi.NewPointerArgument(dt.P)
	args := []gi.Argument{arg_dt}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_from_iso8601_string
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromIso8601String(string string) (result DateTime) {
	iv, err := _I.Get(248, "DateTime", "new_from_iso8601_string")
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

// gst_date_time_new_from_unix_epoch_local_time
//
// [ secs ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromUnixEpochLocalTime(secs int64) (result DateTime) {
	iv, err := _I.Get(249, "DateTime", "new_from_unix_epoch_local_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_secs := gi.NewInt64Argument(secs)
	args := []gi.Argument{arg_secs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_from_unix_epoch_utc
//
// [ secs ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeFromUnixEpochUtc(secs int64) (result DateTime) {
	iv, err := _I.Get(250, "DateTime", "new_from_unix_epoch_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_secs := gi.NewInt64Argument(secs)
	args := []gi.Argument{arg_secs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_local_time
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
func NewDateTimeLocalTime(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(251, "DateTime", "new_local_time")
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

// gst_date_time_new_now_local_time
//
// [ result ] trans: everything
//
func NewDateTimeNowLocalTime() (result DateTime) {
	iv, err := _I.Get(252, "DateTime", "new_now_local_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_now_utc
//
// [ result ] trans: everything
//
func NewDateTimeNowUtc() (result DateTime) {
	iv, err := _I.Get(253, "DateTime", "new_now_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_y
//
// [ year ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeY(year int32) (result DateTime) {
	iv, err := _I.Get(254, "DateTime", "new_y")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_ym
//
// [ year ] trans: nothing
//
// [ month ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeYm(year int32, month int32) (result DateTime) {
	iv, err := _I.Get(255, "DateTime", "new_ym")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	args := []gi.Argument{arg_year, arg_month}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_new_ymd
//
// [ year ] trans: nothing
//
// [ month ] trans: nothing
//
// [ day ] trans: nothing
//
// [ result ] trans: everything
//
func NewDateTimeYmd(year int32, month int32, day int32) (result DateTime) {
	iv, err := _I.Get(256, "DateTime", "new_ymd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	args := []gi.Argument{arg_year, arg_month, arg_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_date_time_get_day
//
// [ result ] trans: nothing
//
func (v DateTime) GetDay() (result int32) {
	iv, err := _I.Get(257, "DateTime", "get_day")
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

// gst_date_time_get_hour
//
// [ result ] trans: nothing
//
func (v DateTime) GetHour() (result int32) {
	iv, err := _I.Get(258, "DateTime", "get_hour")
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

// gst_date_time_get_microsecond
//
// [ result ] trans: nothing
//
func (v DateTime) GetMicrosecond() (result int32) {
	iv, err := _I.Get(259, "DateTime", "get_microsecond")
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

// gst_date_time_get_minute
//
// [ result ] trans: nothing
//
func (v DateTime) GetMinute() (result int32) {
	iv, err := _I.Get(260, "DateTime", "get_minute")
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

// gst_date_time_get_month
//
// [ result ] trans: nothing
//
func (v DateTime) GetMonth() (result int32) {
	iv, err := _I.Get(261, "DateTime", "get_month")
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

// gst_date_time_get_second
//
// [ result ] trans: nothing
//
func (v DateTime) GetSecond() (result int32) {
	iv, err := _I.Get(262, "DateTime", "get_second")
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

// gst_date_time_get_time_zone_offset
//
// [ result ] trans: nothing
//
func (v DateTime) GetTimeZoneOffset() (result float32) {
	iv, err := _I.Get(263, "DateTime", "get_time_zone_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Float()
	return
}

// gst_date_time_get_year
//
// [ result ] trans: nothing
//
func (v DateTime) GetYear() (result int32) {
	iv, err := _I.Get(264, "DateTime", "get_year")
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

// gst_date_time_has_day
//
// [ result ] trans: nothing
//
func (v DateTime) HasDay() (result bool) {
	iv, err := _I.Get(265, "DateTime", "has_day")
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

// gst_date_time_has_month
//
// [ result ] trans: nothing
//
func (v DateTime) HasMonth() (result bool) {
	iv, err := _I.Get(266, "DateTime", "has_month")
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

// gst_date_time_has_second
//
// [ result ] trans: nothing
//
func (v DateTime) HasSecond() (result bool) {
	iv, err := _I.Get(267, "DateTime", "has_second")
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

// gst_date_time_has_time
//
// [ result ] trans: nothing
//
func (v DateTime) HasTime() (result bool) {
	iv, err := _I.Get(268, "DateTime", "has_time")
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

// gst_date_time_has_year
//
// [ result ] trans: nothing
//
func (v DateTime) HasYear() (result bool) {
	iv, err := _I.Get(269, "DateTime", "has_year")
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

// gst_date_time_ref
//
// [ result ] trans: everything
//
func (v DateTime) Ref() (result DateTime) {
	iv, err := _I.Get(270, "DateTime", "ref")
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

// gst_date_time_to_g_date_time
//
// [ result ] trans: everything
//
func (v DateTime) ToGDateTime() (result g.DateTime) {
	iv, err := _I.Get(271, "DateTime", "to_g_date_time")
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

// gst_date_time_to_iso8601_string
//
// [ result ] trans: everything
//
func (v DateTime) ToIso8601String() (result string) {
	iv, err := _I.Get(272, "DateTime", "to_iso8601_string")
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

// gst_date_time_unref
//
func (v DateTime) Unref() {
	iv, err := _I.Get(273, "DateTime", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct DebugCategory
type DebugCategory struct {
	P unsafe.Pointer
}

const SizeOfStructDebugCategory = 24

func DebugCategoryGetType() gi.GType {
	ret := _I.GetGType(40, "DebugCategory")
	return ret
}

// gst_debug_category_free
//
func (v DebugCategory) Free() {
	iv, err := _I.Get(274, "DebugCategory", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_debug_category_get_color
//
// [ result ] trans: nothing
//
func (v DebugCategory) GetColor() (result uint32) {
	iv, err := _I.Get(275, "DebugCategory", "get_color")
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

// gst_debug_category_get_description
//
// [ result ] trans: nothing
//
func (v DebugCategory) GetDescription() (result string) {
	iv, err := _I.Get(276, "DebugCategory", "get_description")
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

// gst_debug_category_get_name
//
// [ result ] trans: nothing
//
func (v DebugCategory) GetName() (result string) {
	iv, err := _I.Get(277, "DebugCategory", "get_name")
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

// gst_debug_category_get_threshold
//
// [ result ] trans: nothing
//
func (v DebugCategory) GetThreshold() (result DebugLevelEnum) {
	iv, err := _I.Get(278, "DebugCategory", "get_threshold")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DebugLevelEnum(ret.Int())
	return
}

// gst_debug_category_reset_threshold
//
func (v DebugCategory) ResetThreshold() {
	iv, err := _I.Get(279, "DebugCategory", "reset_threshold")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_debug_category_set_threshold
//
// [ level ] trans: nothing
//
func (v DebugCategory) SetThreshold(level DebugLevelEnum) {
	iv, err := _I.Get(280, "DebugCategory", "set_threshold")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_level := gi.NewIntArgument(int(level))
	args := []gi.Argument{arg_v, arg_level}
	iv.Call(args, nil, nil)
}

// Flags DebugColorFlags
type DebugColorFlags int

const (
	DebugColorFlagsFgBlack   DebugColorFlags = 0
	DebugColorFlagsFgRed     DebugColorFlags = 1
	DebugColorFlagsFgGreen   DebugColorFlags = 2
	DebugColorFlagsFgYellow  DebugColorFlags = 3
	DebugColorFlagsFgBlue    DebugColorFlags = 4
	DebugColorFlagsFgMagenta DebugColorFlags = 5
	DebugColorFlagsFgCyan    DebugColorFlags = 6
	DebugColorFlagsFgWhite   DebugColorFlags = 7
	DebugColorFlagsBgBlack   DebugColorFlags = 0
	DebugColorFlagsBgRed     DebugColorFlags = 16
	DebugColorFlagsBgGreen   DebugColorFlags = 32
	DebugColorFlagsBgYellow  DebugColorFlags = 48
	DebugColorFlagsBgBlue    DebugColorFlags = 64
	DebugColorFlagsBgMagenta DebugColorFlags = 80
	DebugColorFlagsBgCyan    DebugColorFlags = 96
	DebugColorFlagsBgWhite   DebugColorFlags = 112
	DebugColorFlagsBold      DebugColorFlags = 256
	DebugColorFlagsUnderline DebugColorFlags = 512
)

func DebugColorFlagsGetType() gi.GType {
	ret := _I.GetGType(41, "DebugColorFlags")
	return ret
}

// Enum DebugColorMode
type DebugColorModeEnum int

const (
	DebugColorModeOff  DebugColorModeEnum = 0
	DebugColorModeOn   DebugColorModeEnum = 1
	DebugColorModeUnix DebugColorModeEnum = 2
)

func DebugColorModeGetType() gi.GType {
	ret := _I.GetGType(42, "DebugColorMode")
	return ret
}

type DebugFuncPtrStruct struct {
}

func GetPointer_myDebugFuncPtr() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstDebugFuncPtr())
}

//export myGstDebugFuncPtr
func myGstDebugFuncPtr() {
	// TODO: not found user_data
}

// Flags DebugGraphDetails
type DebugGraphDetailsFlags int

const (
	DebugGraphDetailsMediaType        DebugGraphDetailsFlags = 1
	DebugGraphDetailsCapsDetails      DebugGraphDetailsFlags = 2
	DebugGraphDetailsNonDefaultParams DebugGraphDetailsFlags = 4
	DebugGraphDetailsStates           DebugGraphDetailsFlags = 8
	DebugGraphDetailsFullParams       DebugGraphDetailsFlags = 16
	DebugGraphDetailsAll              DebugGraphDetailsFlags = 15
	DebugGraphDetailsVerbose          DebugGraphDetailsFlags = -1
)

func DebugGraphDetailsGetType() gi.GType {
	ret := _I.GetGType(43, "DebugGraphDetails")
	return ret
}

// Enum DebugLevel
type DebugLevelEnum int

const (
	DebugLevelNone    DebugLevelEnum = 0
	DebugLevelError   DebugLevelEnum = 1
	DebugLevelWarning DebugLevelEnum = 2
	DebugLevelFixme   DebugLevelEnum = 3
	DebugLevelInfo    DebugLevelEnum = 4
	DebugLevelDebug   DebugLevelEnum = 5
	DebugLevelLog     DebugLevelEnum = 6
	DebugLevelTrace   DebugLevelEnum = 7
	DebugLevelMemdump DebugLevelEnum = 9
	DebugLevelCount   DebugLevelEnum = 10
)

func DebugLevelGetType() gi.GType {
	ret := _I.GetGType(44, "DebugLevel")
	return ret
}

// Struct DebugMessage
type DebugMessage struct {
	P unsafe.Pointer
}

func DebugMessageGetType() gi.GType {
	ret := _I.GetGType(45, "DebugMessage")
	return ret
}

// gst_debug_message_get
//
// [ result ] trans: nothing
//
func (v DebugMessage) Get() (result string) {
	iv, err := _I.Get(281, "DebugMessage", "get")
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

// Object Device
type Device struct {
	Object
}

func WrapDevice(p unsafe.Pointer) (r Device) { r.P = p; return }

type IDevice interface{ P_Device() unsafe.Pointer }

func (v Device) P_Device() unsafe.Pointer { return v.P }
func DeviceGetType() gi.GType {
	ret := _I.GetGType(46, "Device")
	return ret
}

// gst_device_create_element
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Device) CreateElement(name string) (result Element) {
	iv, err := _I.Get(282, "Device", "create_element")
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
	result.P = ret.Pointer()
	return
}

// gst_device_get_caps
//
// [ result ] trans: everything
//
func (v Device) GetCaps() (result Caps) {
	iv, err := _I.Get(283, "Device", "get_caps")
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

// gst_device_get_device_class
//
// [ result ] trans: everything
//
func (v Device) GetDeviceClass() (result string) {
	iv, err := _I.Get(284, "Device", "get_device_class")
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

// gst_device_get_display_name
//
// [ result ] trans: everything
//
func (v Device) GetDisplayName() (result string) {
	iv, err := _I.Get(285, "Device", "get_display_name")
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

// gst_device_get_properties
//
// [ result ] trans: everything
//
func (v Device) GetProperties() (result Structure) {
	iv, err := _I.Get(286, "Device", "get_properties")
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

// gst_device_has_classes
//
// [ classes ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) HasClasses(classes string) (result bool) {
	iv, err := _I.Get(287, "Device", "has_classes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_classes := gi.CString(classes)
	arg_v := gi.NewPointerArgument(v.P)
	arg_classes := gi.NewStringArgument(c_classes)
	args := []gi.Argument{arg_v, arg_classes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_classes)
	result = ret.Bool()
	return
}

// gst_device_has_classesv
//
// [ classes ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) HasClassesv(classes gi.CStrArray) (result bool) {
	iv, err := _I.Get(288, "Device", "has_classesv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_classes := gi.NewPointerArgument(classes.P)
	args := []gi.Argument{arg_v, arg_classes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_device_reconfigure_element
//
// [ element ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) ReconfigureElement(element IElement) (result bool) {
	iv, err := _I.Get(289, "Device", "reconfigure_element")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if element != nil {
		tmp = element.P_Element()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_element := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct DeviceClass

// Object DeviceMonitor
type DeviceMonitor struct {
	Object
}

func WrapDeviceMonitor(p unsafe.Pointer) (r DeviceMonitor) { r.P = p; return }

type IDeviceMonitor interface{ P_DeviceMonitor() unsafe.Pointer }

func (v DeviceMonitor) P_DeviceMonitor() unsafe.Pointer { return v.P }
func DeviceMonitorGetType() gi.GType {
	ret := _I.GetGType(47, "DeviceMonitor")
	return ret
}

// gst_device_monitor_new
//
// [ result ] trans: everything
//
func NewDeviceMonitor() (result DeviceMonitor) {
	iv, err := _I.Get(290, "DeviceMonitor", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_device_monitor_add_filter
//
// [ classes ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceMonitor) AddFilter(classes string, caps Caps) (result uint32) {
	iv, err := _I.Get(291, "DeviceMonitor", "add_filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_classes := gi.CString(classes)
	arg_v := gi.NewPointerArgument(v.P)
	arg_classes := gi.NewStringArgument(c_classes)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_classes, arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_classes)
	result = ret.Uint32()
	return
}

// gst_device_monitor_get_bus
//
// [ result ] trans: everything
//
func (v DeviceMonitor) GetBus() (result Bus) {
	iv, err := _I.Get(292, "DeviceMonitor", "get_bus")
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

// gst_device_monitor_get_devices
//
// [ result ] trans: everything
//
func (v DeviceMonitor) GetDevices() (result g.List) {
	iv, err := _I.Get(293, "DeviceMonitor", "get_devices")
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

// gst_device_monitor_get_providers
//
// [ result ] trans: everything
//
func (v DeviceMonitor) GetProviders() (result gi.CStrArray) {
	iv, err := _I.Get(294, "DeviceMonitor", "get_providers")
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

// gst_device_monitor_get_show_all_devices
//
// [ result ] trans: nothing
//
func (v DeviceMonitor) GetShowAllDevices() (result bool) {
	iv, err := _I.Get(295, "DeviceMonitor", "get_show_all_devices")
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

// gst_device_monitor_remove_filter
//
// [ filter_id ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceMonitor) RemoveFilter(filter_id uint32) (result bool) {
	iv, err := _I.Get(296, "DeviceMonitor", "remove_filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_filter_id := gi.NewUint32Argument(filter_id)
	args := []gi.Argument{arg_v, arg_filter_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_device_monitor_set_show_all_devices
//
// [ show_all ] trans: nothing
//
func (v DeviceMonitor) SetShowAllDevices(show_all bool) {
	iv, err := _I.Get(297, "DeviceMonitor", "set_show_all_devices")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_show_all := gi.NewBoolArgument(show_all)
	args := []gi.Argument{arg_v, arg_show_all}
	iv.Call(args, nil, nil)
}

// gst_device_monitor_start
//
// [ result ] trans: nothing
//
func (v DeviceMonitor) Start() (result bool) {
	iv, err := _I.Get(298, "DeviceMonitor", "start")
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

// gst_device_monitor_stop
//
func (v DeviceMonitor) Stop() {
	iv, err := _I.Get(299, "DeviceMonitor", "stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct DeviceMonitorClass

// Struct DeviceMonitorPrivate
type DeviceMonitorPrivate struct {
	P unsafe.Pointer
}

func DeviceMonitorPrivateGetType() gi.GType {
	ret := _I.GetGType(48, "DeviceMonitorPrivate")
	return ret
}

// Struct DevicePrivate
type DevicePrivate struct {
	P unsafe.Pointer
}

func DevicePrivateGetType() gi.GType {
	ret := _I.GetGType(49, "DevicePrivate")
	return ret
}

// Object DeviceProvider
type DeviceProvider struct {
	Object
}

func WrapDeviceProvider(p unsafe.Pointer) (r DeviceProvider) { r.P = p; return }

type IDeviceProvider interface{ P_DeviceProvider() unsafe.Pointer }

func (v DeviceProvider) P_DeviceProvider() unsafe.Pointer { return v.P }
func DeviceProviderGetType() gi.GType {
	ret := _I.GetGType(50, "DeviceProvider")
	return ret
}

// gst_device_provider_register
//
// [ plugin ] trans: nothing
//
// [ name ] trans: nothing
//
// [ rank ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func DeviceProviderRegister1(plugin IPlugin, name string, rank uint32, type1 gi.GType) (result bool) {
	iv, err := _I.Get(300, "DeviceProvider", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	c_name := gi.CString(name)
	arg_plugin := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_rank := gi.NewUint32Argument(rank)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_plugin, arg_name, arg_rank, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_device_provider_can_monitor
//
// [ result ] trans: nothing
//
func (v DeviceProvider) CanMonitor() (result bool) {
	iv, err := _I.Get(301, "DeviceProvider", "can_monitor")
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

// gst_device_provider_device_add
//
// [ device ] trans: nothing
//
func (v DeviceProvider) DeviceAdd(device IDevice) {
	iv, err := _I.Get(302, "DeviceProvider", "device_add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, nil)
}

// gst_device_provider_device_remove
//
// [ device ] trans: nothing
//
func (v DeviceProvider) DeviceRemove(device IDevice) {
	iv, err := _I.Get(303, "DeviceProvider", "device_remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, nil)
}

// gst_device_provider_get_bus
//
// [ result ] trans: everything
//
func (v DeviceProvider) GetBus() (result Bus) {
	iv, err := _I.Get(304, "DeviceProvider", "get_bus")
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

// gst_device_provider_get_devices
//
// [ result ] trans: everything
//
func (v DeviceProvider) GetDevices() (result g.List) {
	iv, err := _I.Get(305, "DeviceProvider", "get_devices")
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

// gst_device_provider_get_factory
//
// [ result ] trans: nothing
//
func (v DeviceProvider) GetFactory() (result DeviceProviderFactory) {
	iv, err := _I.Get(306, "DeviceProvider", "get_factory")
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

// gst_device_provider_get_hidden_providers
//
// [ result ] trans: everything
//
func (v DeviceProvider) GetHiddenProviders() (result gi.CStrArray) {
	iv, err := _I.Get(307, "DeviceProvider", "get_hidden_providers")
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

// gst_device_provider_get_metadata
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceProvider) GetMetadata(key string) (result string) {
	iv, err := _I.Get(308, "DeviceProvider", "get_metadata")
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
	result = ret.String().Copy()
	return
}

// gst_device_provider_hide_provider
//
// [ name ] trans: nothing
//
func (v DeviceProvider) HideProvider(name string) {
	iv, err := _I.Get(309, "DeviceProvider", "hide_provider")
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

// gst_device_provider_start
//
// [ result ] trans: nothing
//
func (v DeviceProvider) Start() (result bool) {
	iv, err := _I.Get(310, "DeviceProvider", "start")
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

// gst_device_provider_stop
//
func (v DeviceProvider) Stop() {
	iv, err := _I.Get(311, "DeviceProvider", "stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_device_provider_unhide_provider
//
// [ name ] trans: nothing
//
func (v DeviceProvider) UnhideProvider(name string) {
	iv, err := _I.Get(312, "DeviceProvider", "unhide_provider")
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

// Struct DeviceProviderClass
type DeviceProviderClass struct {
	P unsafe.Pointer
}

const SizeOfStructDeviceProviderClass = 256

func DeviceProviderClassGetType() gi.GType {
	ret := _I.GetGType(51, "DeviceProviderClass")
	return ret
}

// gst_device_provider_class_add_metadata
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v DeviceProviderClass) AddMetadata(key string, value string) {
	iv, err := _I.Get(313, "DeviceProviderClass", "add_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
	gi.Free(c_value)
}

// gst_device_provider_class_add_static_metadata
//
// [ key ] trans: nothing
//
// [ value ] trans: everything
//
func (v DeviceProviderClass) AddStaticMetadata(key string, value string) {
	iv, err := _I.Get(314, "DeviceProviderClass", "add_static_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
	gi.Free(c_value)
}

// gst_device_provider_class_get_metadata
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceProviderClass) GetMetadata(key string) (result string) {
	iv, err := _I.Get(315, "DeviceProviderClass", "get_metadata")
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
	result = ret.String().Copy()
	return
}

// gst_device_provider_class_set_metadata
//
// [ longname ] trans: nothing
//
// [ classification ] trans: nothing
//
// [ description ] trans: nothing
//
// [ author ] trans: nothing
//
func (v DeviceProviderClass) SetMetadata(longname string, classification string, description string, author string) {
	iv, err := _I.Get(316, "DeviceProviderClass", "set_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_longname := gi.CString(longname)
	c_classification := gi.CString(classification)
	c_description := gi.CString(description)
	c_author := gi.CString(author)
	arg_v := gi.NewPointerArgument(v.P)
	arg_longname := gi.NewStringArgument(c_longname)
	arg_classification := gi.NewStringArgument(c_classification)
	arg_description := gi.NewStringArgument(c_description)
	arg_author := gi.NewStringArgument(c_author)
	args := []gi.Argument{arg_v, arg_longname, arg_classification, arg_description, arg_author}
	iv.Call(args, nil, nil)
	gi.Free(c_longname)
	gi.Free(c_classification)
	gi.Free(c_description)
	gi.Free(c_author)
}

// gst_device_provider_class_set_static_metadata
//
// [ longname ] trans: everything
//
// [ classification ] trans: everything
//
// [ description ] trans: everything
//
// [ author ] trans: everything
//
func (v DeviceProviderClass) SetStaticMetadata(longname string, classification string, description string, author string) {
	iv, err := _I.Get(317, "DeviceProviderClass", "set_static_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_longname := gi.CString(longname)
	c_classification := gi.CString(classification)
	c_description := gi.CString(description)
	c_author := gi.CString(author)
	arg_v := gi.NewPointerArgument(v.P)
	arg_longname := gi.NewStringArgument(c_longname)
	arg_classification := gi.NewStringArgument(c_classification)
	arg_description := gi.NewStringArgument(c_description)
	arg_author := gi.NewStringArgument(c_author)
	args := []gi.Argument{arg_v, arg_longname, arg_classification, arg_description, arg_author}
	iv.Call(args, nil, nil)
	gi.Free(c_longname)
	gi.Free(c_classification)
	gi.Free(c_description)
	gi.Free(c_author)
}

// Object DeviceProviderFactory
type DeviceProviderFactory struct {
	PluginFeature
}

func WrapDeviceProviderFactory(p unsafe.Pointer) (r DeviceProviderFactory) { r.P = p; return }

type IDeviceProviderFactory interface{ P_DeviceProviderFactory() unsafe.Pointer }

func (v DeviceProviderFactory) P_DeviceProviderFactory() unsafe.Pointer { return v.P }
func DeviceProviderFactoryGetType() gi.GType {
	ret := _I.GetGType(52, "DeviceProviderFactory")
	return ret
}

// gst_device_provider_factory_find
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func DeviceProviderFactoryFind1(name string) (result DeviceProviderFactory) {
	iv, err := _I.Get(318, "DeviceProviderFactory", "find")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_device_provider_factory_get_by_name
//
// [ factoryname ] trans: nothing
//
// [ result ] trans: everything
//
func DeviceProviderFactoryGetByName1(factoryname string) (result DeviceProvider) {
	iv, err := _I.Get(319, "DeviceProviderFactory", "get_by_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_factoryname := gi.CString(factoryname)
	arg_factoryname := gi.NewStringArgument(c_factoryname)
	args := []gi.Argument{arg_factoryname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_factoryname)
	result.P = ret.Pointer()
	return
}

// gst_device_provider_factory_list_get_device_providers
//
// [ minrank ] trans: nothing
//
// [ result ] trans: everything
//
func DeviceProviderFactoryListGetDeviceProviders1(minrank RankEnum) (result g.List) {
	iv, err := _I.Get(320, "DeviceProviderFactory", "list_get_device_providers")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_minrank := gi.NewIntArgument(int(minrank))
	args := []gi.Argument{arg_minrank}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_device_provider_factory_get
//
// [ result ] trans: everything
//
func (v DeviceProviderFactory) Get() (result DeviceProvider) {
	iv, err := _I.Get(321, "DeviceProviderFactory", "get")
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

// gst_device_provider_factory_get_device_provider_type
//
// [ result ] trans: nothing
//
func (v DeviceProviderFactory) GetDeviceProviderType() (result gi.GType) {
	iv, err := _I.Get(322, "DeviceProviderFactory", "get_device_provider_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_device_provider_factory_get_metadata
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceProviderFactory) GetMetadata(key string) (result string) {
	iv, err := _I.Get(323, "DeviceProviderFactory", "get_metadata")
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
	result = ret.String().Copy()
	return
}

// gst_device_provider_factory_get_metadata_keys
//
// [ result ] trans: everything
//
func (v DeviceProviderFactory) GetMetadataKeys() (result gi.CStrArray) {
	iv, err := _I.Get(324, "DeviceProviderFactory", "get_metadata_keys")
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

// gst_device_provider_factory_has_classes
//
// [ classes ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceProviderFactory) HasClasses(classes string) (result bool) {
	iv, err := _I.Get(325, "DeviceProviderFactory", "has_classes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_classes := gi.CString(classes)
	arg_v := gi.NewPointerArgument(v.P)
	arg_classes := gi.NewStringArgument(c_classes)
	args := []gi.Argument{arg_v, arg_classes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_classes)
	result = ret.Bool()
	return
}

// gst_device_provider_factory_has_classesv
//
// [ classes ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DeviceProviderFactory) HasClassesv(classes gi.CStrArray) (result bool) {
	iv, err := _I.Get(326, "DeviceProviderFactory", "has_classesv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_classes := gi.NewPointerArgument(classes.P)
	args := []gi.Argument{arg_v, arg_classes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct DeviceProviderFactoryClass

// Struct DeviceProviderPrivate
type DeviceProviderPrivate struct {
	P unsafe.Pointer
}

func DeviceProviderPrivateGetType() gi.GType {
	ret := _I.GetGType(53, "DeviceProviderPrivate")
	return ret
}

// Object DoubleRange
type DoubleRange struct {
	P unsafe.Pointer
}

func WrapDoubleRange(p unsafe.Pointer) (r DoubleRange) { r.P = p; return }

type IDoubleRange interface{ P_DoubleRange() unsafe.Pointer }

func (v DoubleRange) P_DoubleRange() unsafe.Pointer { return v.P }
func DoubleRangeGetType() gi.GType {
	ret := _I.GetGType(54, "DoubleRange")
	return ret
}

// Object DynamicTypeFactory
type DynamicTypeFactory struct {
	PluginFeature
}

func WrapDynamicTypeFactory(p unsafe.Pointer) (r DynamicTypeFactory) { r.P = p; return }

type IDynamicTypeFactory interface{ P_DynamicTypeFactory() unsafe.Pointer }

func (v DynamicTypeFactory) P_DynamicTypeFactory() unsafe.Pointer { return v.P }
func DynamicTypeFactoryGetType() gi.GType {
	ret := _I.GetGType(55, "DynamicTypeFactory")
	return ret
}

// gst_dynamic_type_factory_load
//
// [ factoryname ] trans: nothing
//
// [ result ] trans: nothing
//
func DynamicTypeFactoryLoad1(factoryname string) (result gi.GType) {
	iv, err := _I.Get(327, "DynamicTypeFactory", "load")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_factoryname := gi.CString(factoryname)
	arg_factoryname := gi.NewStringArgument(c_factoryname)
	args := []gi.Argument{arg_factoryname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_factoryname)
	result = gi.GType(ret.Uint())
	return
}

// ignore GType struct DynamicTypeFactoryClass

// Object Element
type Element struct {
	Object
}

func WrapElement(p unsafe.Pointer) (r Element) { r.P = p; return }

type IElement interface{ P_Element() unsafe.Pointer }

func (v Element) P_Element() unsafe.Pointer { return v.P }
func ElementGetType() gi.GType {
	ret := _I.GetGType(56, "Element")
	return ret
}

// gst_element_make_from_uri
//
// [ type1 ] trans: nothing
//
// [ uri ] trans: nothing
//
// [ elementname ] trans: nothing
//
// [ result ] trans: nothing
//
func ElementMakeFromUri1(type1 URITypeEnum, uri string, elementname string) (result Element, err error) {
	iv, err := _I.Get(328, "Element", "make_from_uri")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_elementname := gi.CString(elementname)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_uri := gi.NewStringArgument(c_uri)
	arg_elementname := gi.NewStringArgument(c_elementname)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_uri, arg_elementname, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_elementname)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_element_register
//
// [ plugin ] trans: nothing
//
// [ name ] trans: nothing
//
// [ rank ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func ElementRegister1(plugin IPlugin, name string, rank uint32, type1 gi.GType) (result bool) {
	iv, err := _I.Get(329, "Element", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	c_name := gi.CString(name)
	arg_plugin := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_rank := gi.NewUint32Argument(rank)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_plugin, arg_name, arg_rank, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_element_state_change_return_get_name
//
// [ state_ret ] trans: nothing
//
// [ result ] trans: nothing
//
func ElementStateChangeReturnGetName1(state_ret StateChangeReturnEnum) (result string) {
	iv, err := _I.Get(330, "Element", "state_change_return_get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_state_ret := gi.NewIntArgument(int(state_ret))
	args := []gi.Argument{arg_state_ret}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_element_state_get_name
//
// [ state ] trans: nothing
//
// [ result ] trans: nothing
//
func ElementStateGetName1(state StateEnum) (result string) {
	iv, err := _I.Get(331, "Element", "state_get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_state}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_element_abort_state
//
func (v Element) AbortState() {
	iv, err := _I.Get(332, "Element", "abort_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_element_add_pad
//
// [ pad ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) AddPad(pad IPad) (result bool) {
	iv, err := _I.Get(333, "Element", "add_pad")
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

// gst_element_add_property_deep_notify_watch
//
// [ property_name ] trans: nothing
//
// [ include_value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) AddPropertyDeepNotifyWatch(property_name string, include_value bool) (result uint64) {
	iv, err := _I.Get(334, "Element", "add_property_deep_notify_watch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_include_value := gi.NewBoolArgument(include_value)
	args := []gi.Argument{arg_v, arg_property_name, arg_include_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result = ret.Uint64()
	return
}

// gst_element_add_property_notify_watch
//
// [ property_name ] trans: nothing
//
// [ include_value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) AddPropertyNotifyWatch(property_name string, include_value bool) (result uint64) {
	iv, err := _I.Get(335, "Element", "add_property_notify_watch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_include_value := gi.NewBoolArgument(include_value)
	args := []gi.Argument{arg_v, arg_property_name, arg_include_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result = ret.Uint64()
	return
}

// gst_element_call_async
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ destroy_notify ] trans: nothing
//
func (v Element) CallAsync(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, destroy_notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(336, "Element", "call_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myElementCallAsyncFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_func1, arg_user_data, arg_destroy_notify}
	iv.Call(args, nil, nil)
}

// gst_element_change_state
//
// [ transition ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) ChangeState(transition StateChangeEnum) (result StateChangeReturnEnum) {
	iv, err := _I.Get(337, "Element", "change_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_transition := gi.NewIntArgument(int(transition))
	args := []gi.Argument{arg_v, arg_transition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StateChangeReturnEnum(ret.Int())
	return
}

// gst_element_continue_state
//
// [ ret ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) ContinueState(ret StateChangeReturnEnum) (result StateChangeReturnEnum) {
	iv, err := _I.Get(338, "Element", "continue_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ret := gi.NewIntArgument(int(ret))
	args := []gi.Argument{arg_v, arg_ret}
	var ret1 gi.Argument
	iv.Call(args, &ret1, nil)
	result = StateChangeReturnEnum(ret1.Int())
	return
}

// gst_element_create_all_pads
//
func (v Element) CreateAllPads() {
	iv, err := _I.Get(339, "Element", "create_all_pads")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_element_foreach_pad
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) ForeachPad(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(340, "Element", "foreach_pad")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myElementForeachPadFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_foreach_sink_pad
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) ForeachSinkPad(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(341, "Element", "foreach_sink_pad")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myElementForeachPadFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_foreach_src_pad
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) ForeachSrcPad(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(342, "Element", "foreach_src_pad")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myElementForeachPadFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_get_base_time
//
// [ result ] trans: nothing
//
func (v Element) GetBaseTime() (result uint64) {
	iv, err := _I.Get(343, "Element", "get_base_time")
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

// gst_element_get_bus
//
// [ result ] trans: everything
//
func (v Element) GetBus() (result Bus) {
	iv, err := _I.Get(344, "Element", "get_bus")
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

// gst_element_get_clock
//
// [ result ] trans: everything
//
func (v Element) GetClock() (result Clock) {
	iv, err := _I.Get(345, "Element", "get_clock")
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

// gst_element_get_compatible_pad
//
// [ pad ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ result ] trans: everything
//
func (v Element) GetCompatiblePad(pad IPad, caps Caps) (result Pad) {
	iv, err := _I.Get(346, "Element", "get_compatible_pad")
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
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_pad, arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_element_get_compatible_pad_template
//
// [ compattempl ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) GetCompatiblePadTemplate(compattempl IPadTemplate) (result PadTemplate) {
	iv, err := _I.Get(347, "Element", "get_compatible_pad_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if compattempl != nil {
		tmp = compattempl.P_PadTemplate()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compattempl := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_compattempl}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_element_get_context
//
// [ context_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v Element) GetContext(context_type string) (result Context) {
	iv, err := _I.Get(348, "Element", "get_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_type := gi.CString(context_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_context_type := gi.NewStringArgument(c_context_type)
	args := []gi.Argument{arg_v, arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_type)
	result.P = ret.Pointer()
	return
}

// gst_element_get_context_unlocked
//
// [ context_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v Element) GetContextUnlocked(context_type string) (result Context) {
	iv, err := _I.Get(349, "Element", "get_context_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_type := gi.CString(context_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_context_type := gi.NewStringArgument(c_context_type)
	args := []gi.Argument{arg_v, arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_type)
	result.P = ret.Pointer()
	return
}

// gst_element_get_contexts
//
// [ result ] trans: everything
//
func (v Element) GetContexts() (result g.List) {
	iv, err := _I.Get(350, "Element", "get_contexts")
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

// gst_element_get_factory
//
// [ result ] trans: nothing
//
func (v Element) GetFactory() (result ElementFactory) {
	iv, err := _I.Get(351, "Element", "get_factory")
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

// gst_element_get_metadata
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) GetMetadata(key string) (result string) {
	iv, err := _I.Get(352, "Element", "get_metadata")
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
	result = ret.String().Copy()
	return
}

// gst_element_get_pad_template
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) GetPadTemplate(name string) (result PadTemplate) {
	iv, err := _I.Get(353, "Element", "get_pad_template")
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
	result.P = ret.Pointer()
	return
}

// gst_element_get_pad_template_list
//
// [ result ] trans: nothing
//
func (v Element) GetPadTemplateList() (result g.List) {
	iv, err := _I.Get(354, "Element", "get_pad_template_list")
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

// gst_element_get_request_pad
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Element) GetRequestPad(name string) (result Pad) {
	iv, err := _I.Get(355, "Element", "get_request_pad")
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
	result.P = ret.Pointer()
	return
}

// gst_element_get_start_time
//
// [ result ] trans: nothing
//
func (v Element) GetStartTime() (result uint64) {
	iv, err := _I.Get(356, "Element", "get_start_time")
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

// gst_element_get_state
//
// [ state ] trans: everything, dir: out
//
// [ pending ] trans: everything, dir: out
//
// [ timeout ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) GetState(timeout uint64) (result StateChangeReturnEnum, state StateEnum, pending StateEnum) {
	iv, err := _I.Get(357, "Element", "get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_pending := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_state, arg_pending, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	state = StateEnum(outArgs[0].Int())
	pending = StateEnum(outArgs[1].Int())
	result = StateChangeReturnEnum(ret.Int())
	return
}

// gst_element_get_static_pad
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Element) GetStaticPad(name string) (result Pad) {
	iv, err := _I.Get(358, "Element", "get_static_pad")
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
	result.P = ret.Pointer()
	return
}

// gst_element_is_locked_state
//
// [ result ] trans: nothing
//
func (v Element) IsLockedState() (result bool) {
	iv, err := _I.Get(359, "Element", "is_locked_state")
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

// gst_element_iterate_pads
//
// [ result ] trans: everything
//
func (v Element) IteratePads() (result Iterator) {
	iv, err := _I.Get(360, "Element", "iterate_pads")
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

// gst_element_iterate_sink_pads
//
// [ result ] trans: everything
//
func (v Element) IterateSinkPads() (result Iterator) {
	iv, err := _I.Get(361, "Element", "iterate_sink_pads")
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

// gst_element_iterate_src_pads
//
// [ result ] trans: everything
//
func (v Element) IterateSrcPads() (result Iterator) {
	iv, err := _I.Get(362, "Element", "iterate_src_pads")
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

// gst_element_link
//
// [ dest ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) Link(dest IElement) (result bool) {
	iv, err := _I.Get(363, "Element", "link")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_dest}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_link_filtered
//
// [ dest ] trans: nothing
//
// [ filter ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) LinkFiltered(dest IElement, filter Caps) (result bool) {
	iv, err := _I.Get(364, "Element", "link_filtered")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(tmp)
	arg_filter := gi.NewPointerArgument(filter.P)
	args := []gi.Argument{arg_v, arg_dest, arg_filter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_link_pads
//
// [ srcpadname ] trans: nothing
//
// [ dest ] trans: nothing
//
// [ destpadname ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) LinkPads(srcpadname string, dest IElement, destpadname string) (result bool) {
	iv, err := _I.Get(365, "Element", "link_pads")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_srcpadname := gi.CString(srcpadname)
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	c_destpadname := gi.CString(destpadname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_srcpadname := gi.NewStringArgument(c_srcpadname)
	arg_dest := gi.NewPointerArgument(tmp)
	arg_destpadname := gi.NewStringArgument(c_destpadname)
	args := []gi.Argument{arg_v, arg_srcpadname, arg_dest, arg_destpadname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_srcpadname)
	gi.Free(c_destpadname)
	result = ret.Bool()
	return
}

// gst_element_link_pads_filtered
//
// [ srcpadname ] trans: nothing
//
// [ dest ] trans: nothing
//
// [ destpadname ] trans: nothing
//
// [ filter ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) LinkPadsFiltered(srcpadname string, dest IElement, destpadname string, filter Caps) (result bool) {
	iv, err := _I.Get(366, "Element", "link_pads_filtered")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_srcpadname := gi.CString(srcpadname)
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	c_destpadname := gi.CString(destpadname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_srcpadname := gi.NewStringArgument(c_srcpadname)
	arg_dest := gi.NewPointerArgument(tmp)
	arg_destpadname := gi.NewStringArgument(c_destpadname)
	arg_filter := gi.NewPointerArgument(filter.P)
	args := []gi.Argument{arg_v, arg_srcpadname, arg_dest, arg_destpadname, arg_filter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_srcpadname)
	gi.Free(c_destpadname)
	result = ret.Bool()
	return
}

// gst_element_link_pads_full
//
// [ srcpadname ] trans: nothing
//
// [ dest ] trans: nothing
//
// [ destpadname ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) LinkPadsFull(srcpadname string, dest IElement, destpadname string, flags PadLinkCheckFlags) (result bool) {
	iv, err := _I.Get(367, "Element", "link_pads_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_srcpadname := gi.CString(srcpadname)
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	c_destpadname := gi.CString(destpadname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_srcpadname := gi.NewStringArgument(c_srcpadname)
	arg_dest := gi.NewPointerArgument(tmp)
	arg_destpadname := gi.NewStringArgument(c_destpadname)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_srcpadname, arg_dest, arg_destpadname, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_srcpadname)
	gi.Free(c_destpadname)
	result = ret.Bool()
	return
}

// gst_element_lost_state
//
func (v Element) LostState() {
	iv, err := _I.Get(368, "Element", "lost_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_element_message_full
//
// [ type1 ] trans: nothing
//
// [ domain ] trans: nothing
//
// [ code ] trans: nothing
//
// [ text ] trans: everything
//
// [ debug ] trans: everything
//
// [ file ] trans: nothing
//
// [ function ] trans: nothing
//
// [ line ] trans: nothing
//
func (v Element) MessageFull(type1 MessageTypeFlags, domain uint32, code int32, text string, debug string, file string, function string, line int32) {
	iv, err := _I.Get(369, "Element", "message_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	c_debug := gi.CString(debug)
	c_file := gi.CString(file)
	c_function := gi.CString(function)
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	arg_text := gi.NewStringArgument(c_text)
	arg_debug := gi.NewStringArgument(c_debug)
	arg_file := gi.NewStringArgument(c_file)
	arg_function := gi.NewStringArgument(c_function)
	arg_line := gi.NewInt32Argument(line)
	args := []gi.Argument{arg_v, arg_type1, arg_domain, arg_code, arg_text, arg_debug, arg_file, arg_function, arg_line}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
	gi.Free(c_debug)
	gi.Free(c_file)
	gi.Free(c_function)
}

// gst_element_message_full_with_details
//
// [ type1 ] trans: nothing
//
// [ domain ] trans: nothing
//
// [ code ] trans: nothing
//
// [ text ] trans: everything
//
// [ debug ] trans: everything
//
// [ file ] trans: nothing
//
// [ function ] trans: nothing
//
// [ line ] trans: nothing
//
// [ structure ] trans: everything
//
func (v Element) MessageFullWithDetails(type1 MessageTypeFlags, domain uint32, code int32, text string, debug string, file string, function string, line int32, structure Structure) {
	iv, err := _I.Get(370, "Element", "message_full_with_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	c_debug := gi.CString(debug)
	c_file := gi.CString(file)
	c_function := gi.CString(function)
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	arg_text := gi.NewStringArgument(c_text)
	arg_debug := gi.NewStringArgument(c_debug)
	arg_file := gi.NewStringArgument(c_file)
	arg_function := gi.NewStringArgument(c_function)
	arg_line := gi.NewInt32Argument(line)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_v, arg_type1, arg_domain, arg_code, arg_text, arg_debug, arg_file, arg_function, arg_line, arg_structure}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
	gi.Free(c_debug)
	gi.Free(c_file)
	gi.Free(c_function)
}

// gst_element_no_more_pads
//
func (v Element) NoMorePads() {
	iv, err := _I.Get(371, "Element", "no_more_pads")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_element_post_message
//
// [ message ] trans: everything
//
// [ result ] trans: nothing
//
func (v Element) PostMessage(message Message) (result bool) {
	iv, err := _I.Get(372, "Element", "post_message")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_message := gi.NewPointerArgument(message.P)
	args := []gi.Argument{arg_v, arg_message}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_provide_clock
//
// [ result ] trans: everything
//
func (v Element) ProvideClock() (result Clock) {
	iv, err := _I.Get(373, "Element", "provide_clock")
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

// gst_element_query
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) QueryF(query Query) (result bool) {
	iv, err := _I.Get(374, "Element", "query")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_v, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_query_convert
//
// [ src_format ] trans: nothing
//
// [ src_val ] trans: nothing
//
// [ dest_format ] trans: nothing
//
// [ dest_val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Element) QueryConvert(src_format FormatEnum, src_val int64, dest_format FormatEnum) (result bool, dest_val int64) {
	iv, err := _I.Get(375, "Element", "query_convert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_format := gi.NewIntArgument(int(src_format))
	arg_src_val := gi.NewInt64Argument(src_val)
	arg_dest_format := gi.NewIntArgument(int(dest_format))
	arg_dest_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_src_format, arg_src_val, arg_dest_format, arg_dest_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	dest_val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_element_query_duration
//
// [ format ] trans: nothing
//
// [ duration ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Element) QueryDuration(format FormatEnum) (result bool, duration int64) {
	iv, err := _I.Get(376, "Element", "query_duration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_duration}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	duration = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_element_query_position
//
// [ format ] trans: nothing
//
// [ cur ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Element) QueryPosition(format FormatEnum) (result bool, cur int64) {
	iv, err := _I.Get(377, "Element", "query_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_cur := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_cur}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	cur = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_element_release_request_pad
//
// [ pad ] trans: nothing
//
func (v Element) ReleaseRequestPad(pad IPad) {
	iv, err := _I.Get(378, "Element", "release_request_pad")
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

// gst_element_remove_pad
//
// [ pad ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) RemovePad(pad IPad) (result bool) {
	iv, err := _I.Get(379, "Element", "remove_pad")
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

// gst_element_remove_property_notify_watch
//
// [ watch_id ] trans: nothing
//
func (v Element) RemovePropertyNotifyWatch(watch_id uint64) {
	iv, err := _I.Get(380, "Element", "remove_property_notify_watch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_watch_id := gi.NewUint64Argument(watch_id)
	args := []gi.Argument{arg_v, arg_watch_id}
	iv.Call(args, nil, nil)
}

// gst_element_request_pad
//
// [ templ ] trans: nothing
//
// [ name ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ result ] trans: everything
//
func (v Element) RequestPad(templ IPadTemplate, name string, caps Caps) (result Pad) {
	iv, err := _I.Get(381, "Element", "request_pad")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if templ != nil {
		tmp = templ.P_PadTemplate()
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_templ := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_templ, arg_name, arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_element_seek
//
// [ rate ] trans: nothing
//
// [ format ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ start_type ] trans: nothing
//
// [ start ] trans: nothing
//
// [ stop_type ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) Seek(rate float64, format FormatEnum, flags SeekFlags, start_type SeekTypeEnum, start int64, stop_type SeekTypeEnum, stop int64) (result bool) {
	iv, err := _I.Get(382, "Element", "seek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rate := gi.NewDoubleArgument(rate)
	arg_format := gi.NewIntArgument(int(format))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_start_type := gi.NewIntArgument(int(start_type))
	arg_start := gi.NewInt64Argument(start)
	arg_stop_type := gi.NewIntArgument(int(stop_type))
	arg_stop := gi.NewInt64Argument(stop)
	args := []gi.Argument{arg_v, arg_rate, arg_format, arg_flags, arg_start_type, arg_start, arg_stop_type, arg_stop}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_seek_simple
//
// [ format ] trans: nothing
//
// [ seek_flags ] trans: nothing
//
// [ seek_pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) SeekSimple(format FormatEnum, seek_flags SeekFlags, seek_pos int64) (result bool) {
	iv, err := _I.Get(383, "Element", "seek_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_seek_flags := gi.NewIntArgument(int(seek_flags))
	arg_seek_pos := gi.NewInt64Argument(seek_pos)
	args := []gi.Argument{arg_v, arg_format, arg_seek_flags, arg_seek_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_send_event
//
// [ event ] trans: everything
//
// [ result ] trans: nothing
//
func (v Element) SendEvent(event Event) (result bool) {
	iv, err := _I.Get(384, "Element", "send_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_set_base_time
//
// [ time ] trans: nothing
//
func (v Element) SetBaseTime(time uint64) {
	iv, err := _I.Get(385, "Element", "set_base_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time := gi.NewUint64Argument(time)
	args := []gi.Argument{arg_v, arg_time}
	iv.Call(args, nil, nil)
}

// gst_element_set_bus
//
// [ bus ] trans: nothing
//
func (v Element) SetBus(bus IBus) {
	iv, err := _I.Get(386, "Element", "set_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if bus != nil {
		tmp = bus.P_Bus()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bus := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_bus}
	iv.Call(args, nil, nil)
}

// gst_element_set_clock
//
// [ clock ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) SetClock(clock IClock) (result bool) {
	iv, err := _I.Get(387, "Element", "set_clock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if clock != nil {
		tmp = clock.P_Clock()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_clock := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_clock}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_set_context
//
// [ context ] trans: nothing
//
func (v Element) SetContext(context Context) {
	iv, err := _I.Get(388, "Element", "set_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(context.P)
	args := []gi.Argument{arg_v, arg_context}
	iv.Call(args, nil, nil)
}

// gst_element_set_locked_state
//
// [ locked_state ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) SetLockedState(locked_state bool) (result bool) {
	iv, err := _I.Get(389, "Element", "set_locked_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_locked_state := gi.NewBoolArgument(locked_state)
	args := []gi.Argument{arg_v, arg_locked_state}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_element_set_start_time
//
// [ time ] trans: nothing
//
func (v Element) SetStartTime(time uint64) {
	iv, err := _I.Get(390, "Element", "set_start_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time := gi.NewUint64Argument(time)
	args := []gi.Argument{arg_v, arg_time}
	iv.Call(args, nil, nil)
}

// gst_element_set_state
//
// [ state ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Element) SetState(state StateEnum) (result StateChangeReturnEnum) {
	iv, err := _I.Get(391, "Element", "set_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_v, arg_state}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StateChangeReturnEnum(ret.Int())
	return
}

// gst_element_sync_state_with_parent
//
// [ result ] trans: nothing
//
func (v Element) SyncStateWithParent() (result bool) {
	iv, err := _I.Get(392, "Element", "sync_state_with_parent")
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

// gst_element_unlink
//
// [ dest ] trans: nothing
//
func (v Element) Unlink(dest IElement) {
	iv, err := _I.Get(393, "Element", "unlink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_dest}
	iv.Call(args, nil, nil)
}

// gst_element_unlink_pads
//
// [ srcpadname ] trans: nothing
//
// [ dest ] trans: nothing
//
// [ destpadname ] trans: nothing
//
func (v Element) UnlinkPads(srcpadname string, dest IElement, destpadname string) {
	iv, err := _I.Get(394, "Element", "unlink_pads")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_srcpadname := gi.CString(srcpadname)
	var tmp unsafe.Pointer
	if dest != nil {
		tmp = dest.P_Element()
	}
	c_destpadname := gi.CString(destpadname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_srcpadname := gi.NewStringArgument(c_srcpadname)
	arg_dest := gi.NewPointerArgument(tmp)
	arg_destpadname := gi.NewStringArgument(c_destpadname)
	args := []gi.Argument{arg_v, arg_srcpadname, arg_dest, arg_destpadname}
	iv.Call(args, nil, nil)
	gi.Free(c_srcpadname)
	gi.Free(c_destpadname)
}

type ElementCallAsyncFuncStruct struct {
	F_element Element
}

func GetPointer_myElementCallAsyncFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstElementCallAsyncFunc())
}

//export myGstElementCallAsyncFunc
func myGstElementCallAsyncFunc(element *C.GstElement, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &ElementCallAsyncFuncStruct{
		F_element: WrapElement(unsafe.Pointer(element)),
	}
	fn(args)
}

// Struct ElementClass
type ElementClass struct {
	P unsafe.Pointer
}

const SizeOfStructElementClass = 488

func ElementClassGetType() gi.GType {
	ret := _I.GetGType(57, "ElementClass")
	return ret
}

// gst_element_class_add_metadata
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v ElementClass) AddMetadata(key string, value string) {
	iv, err := _I.Get(395, "ElementClass", "add_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
	gi.Free(c_value)
}

// gst_element_class_add_pad_template
//
// [ templ ] trans: nothing
//
func (v ElementClass) AddPadTemplate(templ IPadTemplate) {
	iv, err := _I.Get(396, "ElementClass", "add_pad_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if templ != nil {
		tmp = templ.P_PadTemplate()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_templ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_templ}
	iv.Call(args, nil, nil)
}

// gst_element_class_add_static_metadata
//
// [ key ] trans: nothing
//
// [ value ] trans: nothing
//
func (v ElementClass) AddStaticMetadata(key string, value string) {
	iv, err := _I.Get(397, "ElementClass", "add_static_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
	gi.Free(c_value)
}

// gst_element_class_add_static_pad_template
//
// [ static_templ ] trans: nothing
//
func (v ElementClass) AddStaticPadTemplate(static_templ StaticPadTemplate) {
	iv, err := _I.Get(398, "ElementClass", "add_static_pad_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_static_templ := gi.NewPointerArgument(static_templ.P)
	args := []gi.Argument{arg_v, arg_static_templ}
	iv.Call(args, nil, nil)
}

// gst_element_class_add_static_pad_template_with_gtype
//
// [ static_templ ] trans: nothing
//
// [ pad_type ] trans: nothing
//
func (v ElementClass) AddStaticPadTemplateWithGtype(static_templ StaticPadTemplate, pad_type gi.GType) {
	iv, err := _I.Get(399, "ElementClass", "add_static_pad_template_with_gtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_static_templ := gi.NewPointerArgument(static_templ.P)
	arg_pad_type := gi.NewUintArgument(uint(pad_type))
	args := []gi.Argument{arg_v, arg_static_templ, arg_pad_type}
	iv.Call(args, nil, nil)
}

// gst_element_class_get_metadata
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementClass) GetMetadata(key string) (result string) {
	iv, err := _I.Get(400, "ElementClass", "get_metadata")
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
	result = ret.String().Copy()
	return
}

// gst_element_class_get_pad_template
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementClass) GetPadTemplate(name string) (result PadTemplate) {
	iv, err := _I.Get(401, "ElementClass", "get_pad_template")
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
	result.P = ret.Pointer()
	return
}

// gst_element_class_get_pad_template_list
//
// [ result ] trans: nothing
//
func (v ElementClass) GetPadTemplateList() (result g.List) {
	iv, err := _I.Get(402, "ElementClass", "get_pad_template_list")
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

// gst_element_class_set_metadata
//
// [ longname ] trans: nothing
//
// [ classification ] trans: nothing
//
// [ description ] trans: nothing
//
// [ author ] trans: nothing
//
func (v ElementClass) SetMetadata(longname string, classification string, description string, author string) {
	iv, err := _I.Get(403, "ElementClass", "set_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_longname := gi.CString(longname)
	c_classification := gi.CString(classification)
	c_description := gi.CString(description)
	c_author := gi.CString(author)
	arg_v := gi.NewPointerArgument(v.P)
	arg_longname := gi.NewStringArgument(c_longname)
	arg_classification := gi.NewStringArgument(c_classification)
	arg_description := gi.NewStringArgument(c_description)
	arg_author := gi.NewStringArgument(c_author)
	args := []gi.Argument{arg_v, arg_longname, arg_classification, arg_description, arg_author}
	iv.Call(args, nil, nil)
	gi.Free(c_longname)
	gi.Free(c_classification)
	gi.Free(c_description)
	gi.Free(c_author)
}

// gst_element_class_set_static_metadata
//
// [ longname ] trans: nothing
//
// [ classification ] trans: nothing
//
// [ description ] trans: nothing
//
// [ author ] trans: nothing
//
func (v ElementClass) SetStaticMetadata(longname string, classification string, description string, author string) {
	iv, err := _I.Get(404, "ElementClass", "set_static_metadata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_longname := gi.CString(longname)
	c_classification := gi.CString(classification)
	c_description := gi.CString(description)
	c_author := gi.CString(author)
	arg_v := gi.NewPointerArgument(v.P)
	arg_longname := gi.NewStringArgument(c_longname)
	arg_classification := gi.NewStringArgument(c_classification)
	arg_description := gi.NewStringArgument(c_description)
	arg_author := gi.NewStringArgument(c_author)
	args := []gi.Argument{arg_v, arg_longname, arg_classification, arg_description, arg_author}
	iv.Call(args, nil, nil)
	gi.Free(c_longname)
	gi.Free(c_classification)
	gi.Free(c_description)
	gi.Free(c_author)
}

// Object ElementFactory
type ElementFactory struct {
	PluginFeature
}

func WrapElementFactory(p unsafe.Pointer) (r ElementFactory) { r.P = p; return }

type IElementFactory interface{ P_ElementFactory() unsafe.Pointer }

func (v ElementFactory) P_ElementFactory() unsafe.Pointer { return v.P }
func ElementFactoryGetType() gi.GType {
	ret := _I.GetGType(58, "ElementFactory")
	return ret
}

// gst_element_factory_find
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func ElementFactoryFind1(name string) (result ElementFactory) {
	iv, err := _I.Get(405, "ElementFactory", "find")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_element_factory_list_filter
//
// [ list ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ direction ] trans: nothing
//
// [ subsetonly ] trans: nothing
//
// [ result ] trans: everything
//
func ElementFactoryListFilter1(list g.List, caps Caps, direction PadDirectionEnum, subsetonly bool) (result g.List) {
	iv, err := _I.Get(406, "ElementFactory", "list_filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_direction := gi.NewIntArgument(int(direction))
	arg_subsetonly := gi.NewBoolArgument(subsetonly)
	args := []gi.Argument{arg_list, arg_caps, arg_direction, arg_subsetonly}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_element_factory_list_get_elements
//
// [ type1 ] trans: nothing
//
// [ minrank ] trans: nothing
//
// [ result ] trans: everything
//
func ElementFactoryListGetElements1(type1 uint64, minrank RankEnum) (result g.List) {
	iv, err := _I.Get(407, "ElementFactory", "list_get_elements")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUint64Argument(type1)
	arg_minrank := gi.NewIntArgument(int(minrank))
	args := []gi.Argument{arg_type1, arg_minrank}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_element_factory_make
//
// [ factoryname ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func ElementFactoryMake1(factoryname string, name string) (result Element) {
	iv, err := _I.Get(408, "ElementFactory", "make")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_factoryname := gi.CString(factoryname)
	c_name := gi.CString(name)
	arg_factoryname := gi.NewStringArgument(c_factoryname)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_factoryname, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_factoryname)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_element_factory_can_sink_all_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) CanSinkAllCaps(caps Caps) (result bool) {
	iv, err := _I.Get(409, "ElementFactory", "can_sink_all_caps")
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

// gst_element_factory_can_sink_any_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) CanSinkAnyCaps(caps Caps) (result bool) {
	iv, err := _I.Get(410, "ElementFactory", "can_sink_any_caps")
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

// gst_element_factory_can_src_all_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) CanSrcAllCaps(caps Caps) (result bool) {
	iv, err := _I.Get(411, "ElementFactory", "can_src_all_caps")
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

// gst_element_factory_can_src_any_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) CanSrcAnyCaps(caps Caps) (result bool) {
	iv, err := _I.Get(412, "ElementFactory", "can_src_any_caps")
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

// gst_element_factory_create
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) Create(name string) (result Element) {
	iv, err := _I.Get(413, "ElementFactory", "create")
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
	result.P = ret.Pointer()
	return
}

// gst_element_factory_get_element_type
//
// [ result ] trans: nothing
//
func (v ElementFactory) GetElementType() (result gi.GType) {
	iv, err := _I.Get(414, "ElementFactory", "get_element_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_element_factory_get_metadata
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) GetMetadata(key string) (result string) {
	iv, err := _I.Get(415, "ElementFactory", "get_metadata")
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
	result = ret.String().Copy()
	return
}

// gst_element_factory_get_metadata_keys
//
// [ result ] trans: everything
//
func (v ElementFactory) GetMetadataKeys() (result gi.CStrArray) {
	iv, err := _I.Get(416, "ElementFactory", "get_metadata_keys")
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

// gst_element_factory_get_num_pad_templates
//
// [ result ] trans: nothing
//
func (v ElementFactory) GetNumPadTemplates() (result uint32) {
	iv, err := _I.Get(417, "ElementFactory", "get_num_pad_templates")
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

// gst_element_factory_get_static_pad_templates
//
// [ result ] trans: nothing
//
func (v ElementFactory) GetStaticPadTemplates() (result g.List) {
	iv, err := _I.Get(418, "ElementFactory", "get_static_pad_templates")
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

// gst_element_factory_get_uri_protocols
//
// [ result ] trans: nothing
//
func (v ElementFactory) GetUriProtocols() (result gi.CStrArray) {
	iv, err := _I.Get(419, "ElementFactory", "get_uri_protocols")
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

// gst_element_factory_get_uri_type
//
// [ result ] trans: nothing
//
func (v ElementFactory) GetUriType() (result URITypeEnum) {
	iv, err := _I.Get(420, "ElementFactory", "get_uri_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = URITypeEnum(ret.Int())
	return
}

// gst_element_factory_has_interface
//
// [ interfacename ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) HasInterface(interfacename string) (result bool) {
	iv, err := _I.Get(421, "ElementFactory", "has_interface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_interfacename := gi.CString(interfacename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_interfacename := gi.NewStringArgument(c_interfacename)
	args := []gi.Argument{arg_v, arg_interfacename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_interfacename)
	result = ret.Bool()
	return
}

// gst_element_factory_list_is_type
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ElementFactory) ListIsType(type1 uint64) (result bool) {
	iv, err := _I.Get(422, "ElementFactory", "list_is_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewUint64Argument(type1)
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct ElementFactoryClass

// Flags ElementFlags
type ElementFlags int

const (
	ElementFlagsLockedState  ElementFlags = 16
	ElementFlagsSink         ElementFlags = 32
	ElementFlagsSource       ElementFlags = 64
	ElementFlagsProvideClock ElementFlags = 128
	ElementFlagsRequireClock ElementFlags = 256
	ElementFlagsIndexable    ElementFlags = 512
	ElementFlagsLast         ElementFlags = 16384
)

func ElementFlagsGetType() gi.GType {
	ret := _I.GetGType(59, "ElementFlags")
	return ret
}

type ElementForeachPadFuncStruct struct {
	F_element Element
	F_pad     Pad
}

func GetPointer_myElementForeachPadFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstElementForeachPadFunc())
}

//export myGstElementForeachPadFunc
func myGstElementForeachPadFunc(element *C.GstElement, pad *C.GstPad, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &ElementForeachPadFuncStruct{
		F_element: WrapElement(unsafe.Pointer(element)),
		F_pad:     WrapPad(unsafe.Pointer(pad)),
	}
	fn(args)
}

// Struct Event
type Event struct {
	P unsafe.Pointer
}

const SizeOfStructEvent = 88

func EventGetType() gi.GType {
	ret := _I.GetGType(60, "Event")
	return ret
}

// gst_event_new_buffer_size
//
// [ format ] trans: nothing
//
// [ minsize ] trans: nothing
//
// [ maxsize ] trans: nothing
//
// [ async ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventBufferSize(format FormatEnum, minsize int64, maxsize int64, async bool) (result Event) {
	iv, err := _I.Get(423, "Event", "new_buffer_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	arg_minsize := gi.NewInt64Argument(minsize)
	arg_maxsize := gi.NewInt64Argument(maxsize)
	arg_async := gi.NewBoolArgument(async)
	args := []gi.Argument{arg_format, arg_minsize, arg_maxsize, arg_async}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventCaps(caps Caps) (result Event) {
	iv, err := _I.Get(424, "Event", "new_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_custom
//
// [ type1 ] trans: nothing
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func NewEventCustom(type1 EventTypeEnum, structure Structure) (result Event) {
	iv, err := _I.Get(425, "Event", "new_custom")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_type1, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_eos
//
// [ result ] trans: everything
//
func NewEventEos() (result Event) {
	iv, err := _I.Get(426, "Event", "new_eos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_flush_start
//
// [ result ] trans: everything
//
func NewEventFlushStart() (result Event) {
	iv, err := _I.Get(427, "Event", "new_flush_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_flush_stop
//
// [ reset_time ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventFlushStop(reset_time bool) (result Event) {
	iv, err := _I.Get(428, "Event", "new_flush_stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_reset_time := gi.NewBoolArgument(reset_time)
	args := []gi.Argument{arg_reset_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_gap
//
// [ timestamp ] trans: nothing
//
// [ duration ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventGap(timestamp uint64, duration uint64) (result Event) {
	iv, err := _I.Get(429, "Event", "new_gap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_duration := gi.NewUint64Argument(duration)
	args := []gi.Argument{arg_timestamp, arg_duration}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_latency
//
// [ latency ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventLatency(latency uint64) (result Event) {
	iv, err := _I.Get(430, "Event", "new_latency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_latency := gi.NewUint64Argument(latency)
	args := []gi.Argument{arg_latency}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_navigation
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func NewEventNavigation(structure Structure) (result Event) {
	iv, err := _I.Get(431, "Event", "new_navigation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_protection
//
// [ system_id ] trans: nothing
//
// [ data ] trans: nothing
//
// [ origin ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventProtection(system_id string, data Buffer, origin string) (result Event) {
	iv, err := _I.Get(432, "Event", "new_protection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_system_id := gi.CString(system_id)
	c_origin := gi.CString(origin)
	arg_system_id := gi.NewStringArgument(c_system_id)
	arg_data := gi.NewPointerArgument(data.P)
	arg_origin := gi.NewStringArgument(c_origin)
	args := []gi.Argument{arg_system_id, arg_data, arg_origin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_system_id)
	gi.Free(c_origin)
	result.P = ret.Pointer()
	return
}

// gst_event_new_qos
//
// [ type1 ] trans: nothing
//
// [ proportion ] trans: nothing
//
// [ diff ] trans: nothing
//
// [ timestamp ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventQos(type1 QOSTypeEnum, proportion float64, diff int64, timestamp uint64) (result Event) {
	iv, err := _I.Get(433, "Event", "new_qos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_proportion := gi.NewDoubleArgument(proportion)
	arg_diff := gi.NewInt64Argument(diff)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	args := []gi.Argument{arg_type1, arg_proportion, arg_diff, arg_timestamp}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_reconfigure
//
// [ result ] trans: everything
//
func NewEventReconfigure() (result Event) {
	iv, err := _I.Get(434, "Event", "new_reconfigure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_seek
//
// [ rate ] trans: nothing
//
// [ format ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ start_type ] trans: nothing
//
// [ start ] trans: nothing
//
// [ stop_type ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventSeek(rate float64, format FormatEnum, flags SeekFlags, start_type SeekTypeEnum, start int64, stop_type SeekTypeEnum, stop int64) (result Event) {
	iv, err := _I.Get(435, "Event", "new_seek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rate := gi.NewDoubleArgument(rate)
	arg_format := gi.NewIntArgument(int(format))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_start_type := gi.NewIntArgument(int(start_type))
	arg_start := gi.NewInt64Argument(start)
	arg_stop_type := gi.NewIntArgument(int(stop_type))
	arg_stop := gi.NewInt64Argument(stop)
	args := []gi.Argument{arg_rate, arg_format, arg_flags, arg_start_type, arg_start, arg_stop_type, arg_stop}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_segment
//
// [ segment ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventSegment(segment Segment) (result Event) {
	iv, err := _I.Get(436, "Event", "new_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_segment := gi.NewPointerArgument(segment.P)
	args := []gi.Argument{arg_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_segment_done
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventSegmentDone(format FormatEnum, position int64) (result Event) {
	iv, err := _I.Get(437, "Event", "new_segment_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewInt64Argument(position)
	args := []gi.Argument{arg_format, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_select_streams
//
// [ streams ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventSelectStreams(streams g.List) (result Event) {
	iv, err := _I.Get(438, "Event", "new_select_streams")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_streams := gi.NewPointerArgument(streams.P)
	args := []gi.Argument{arg_streams}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_sink_message
//
// [ name ] trans: nothing
//
// [ msg ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventSinkMessage(name string, msg Message) (result Event) {
	iv, err := _I.Get(439, "Event", "new_sink_message")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_msg := gi.NewPointerArgument(msg.P)
	args := []gi.Argument{arg_name, arg_msg}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_event_new_step
//
// [ format ] trans: nothing
//
// [ amount ] trans: nothing
//
// [ rate ] trans: nothing
//
// [ flush ] trans: nothing
//
// [ intermediate ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventStep(format FormatEnum, amount uint64, rate float64, flush bool, intermediate bool) (result Event) {
	iv, err := _I.Get(440, "Event", "new_step")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	arg_amount := gi.NewUint64Argument(amount)
	arg_rate := gi.NewDoubleArgument(rate)
	arg_flush := gi.NewBoolArgument(flush)
	arg_intermediate := gi.NewBoolArgument(intermediate)
	args := []gi.Argument{arg_format, arg_amount, arg_rate, arg_flush, arg_intermediate}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_stream_collection
//
// [ collection ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventStreamCollection(collection IStreamCollection) (result Event) {
	iv, err := _I.Get(441, "Event", "new_stream_collection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if collection != nil {
		tmp = collection.P_StreamCollection()
	}
	arg_collection := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_collection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_stream_group_done
//
// [ group_id ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventStreamGroupDone(group_id uint32) (result Event) {
	iv, err := _I.Get(442, "Event", "new_stream_group_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_group_id := gi.NewUint32Argument(group_id)
	args := []gi.Argument{arg_group_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_stream_start
//
// [ stream_id ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventStreamStart(stream_id string) (result Event) {
	iv, err := _I.Get(443, "Event", "new_stream_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_stream_id := gi.CString(stream_id)
	arg_stream_id := gi.NewStringArgument(c_stream_id)
	args := []gi.Argument{arg_stream_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_stream_id)
	result.P = ret.Pointer()
	return
}

// gst_event_new_tag
//
// [ taglist ] trans: everything
//
// [ result ] trans: everything
//
func NewEventTag(taglist TagList) (result Event) {
	iv, err := _I.Get(444, "Event", "new_tag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_taglist := gi.NewPointerArgument(taglist.P)
	args := []gi.Argument{arg_taglist}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_toc
//
// [ toc ] trans: nothing
//
// [ updated ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventToc(toc Toc, updated bool) (result Event) {
	iv, err := _I.Get(445, "Event", "new_toc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_toc := gi.NewPointerArgument(toc.P)
	arg_updated := gi.NewBoolArgument(updated)
	args := []gi.Argument{arg_toc, arg_updated}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_event_new_toc_select
//
// [ uid ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventTocSelect(uid string) (result Event) {
	iv, err := _I.Get(446, "Event", "new_toc_select")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uid := gi.CString(uid)
	arg_uid := gi.NewStringArgument(c_uid)
	args := []gi.Argument{arg_uid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uid)
	result.P = ret.Pointer()
	return
}

// gst_event_copy_segment
//
// [ segment ] trans: nothing
//
func (v Event) CopySegment(segment Segment) {
	iv, err := _I.Get(447, "Event", "copy_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_segment := gi.NewPointerArgument(segment.P)
	args := []gi.Argument{arg_v, arg_segment}
	iv.Call(args, nil, nil)
}

// gst_event_get_running_time_offset
//
// [ result ] trans: nothing
//
func (v Event) GetRunningTimeOffset() (result int64) {
	iv, err := _I.Get(448, "Event", "get_running_time_offset")
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

// gst_event_get_seqnum
//
// [ result ] trans: nothing
//
func (v Event) GetSeqnum() (result uint32) {
	iv, err := _I.Get(449, "Event", "get_seqnum")
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

// gst_event_get_structure
//
// [ result ] trans: nothing
//
func (v Event) GetStructure() (result Structure) {
	iv, err := _I.Get(450, "Event", "get_structure")
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

// gst_event_has_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Event) HasName(name string) (result bool) {
	iv, err := _I.Get(451, "Event", "has_name")
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
	result = ret.Bool()
	return
}

// gst_event_parse_buffer_size
//
// [ format ] trans: everything, dir: out
//
// [ minsize ] trans: everything, dir: out
//
// [ maxsize ] trans: everything, dir: out
//
// [ async ] trans: everything, dir: out
//
func (v Event) ParseBufferSize() (format FormatEnum, minsize int64, maxsize int64, async bool) {
	iv, err := _I.Get(452, "Event", "parse_buffer_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_minsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_maxsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_async := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_format, arg_minsize, arg_maxsize, arg_async}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	minsize = outArgs[1].Int64()
	maxsize = outArgs[2].Int64()
	async = outArgs[3].Bool()
	return
}

// gst_event_parse_caps
//
// [ caps ] trans: nothing, dir: out
//
func (v Event) ParseCaps() (caps Caps) {
	iv, err := _I.Get(453, "Event", "parse_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_caps}
	iv.Call(args, nil, &outArgs[0])
	caps.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_flush_stop
//
// [ reset_time ] trans: everything, dir: out
//
func (v Event) ParseFlushStop() (reset_time bool) {
	iv, err := _I.Get(454, "Event", "parse_flush_stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_reset_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_reset_time}
	iv.Call(args, nil, &outArgs[0])
	reset_time = outArgs[0].Bool()
	return
}

// gst_event_parse_gap
//
// [ timestamp ] trans: everything, dir: out
//
// [ duration ] trans: everything, dir: out
//
func (v Event) ParseGap() (timestamp uint64, duration uint64) {
	iv, err := _I.Get(455, "Event", "parse_gap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_timestamp, arg_duration}
	iv.Call(args, nil, &outArgs[0])
	timestamp = outArgs[0].Uint64()
	duration = outArgs[1].Uint64()
	return
}

// gst_event_parse_group_id
//
// [ group_id ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) ParseGroupId() (result bool, group_id uint32) {
	iv, err := _I.Get(456, "Event", "parse_group_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_id}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	group_id = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_event_parse_latency
//
// [ latency ] trans: everything, dir: out
//
func (v Event) ParseLatency() (latency uint64) {
	iv, err := _I.Get(457, "Event", "parse_latency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_latency := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_latency}
	iv.Call(args, nil, &outArgs[0])
	latency = outArgs[0].Uint64()
	return
}

// gst_event_parse_protection
//
// [ system_id ] trans: nothing, dir: out
//
// [ data ] trans: nothing, dir: out
//
// [ origin ] trans: nothing, dir: out
//
func (v Event) ParseProtection() (system_id string, data Buffer, origin string) {
	iv, err := _I.Get(458, "Event", "parse_protection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_system_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_data := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_origin := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_system_id, arg_data, arg_origin}
	iv.Call(args, nil, &outArgs[0])
	system_id = outArgs[0].String().Copy()
	data.P = outArgs[1].Pointer()
	origin = outArgs[2].String().Copy()
	return
}

// gst_event_parse_qos
//
// [ type1 ] trans: everything, dir: out
//
// [ proportion ] trans: everything, dir: out
//
// [ diff ] trans: everything, dir: out
//
// [ timestamp ] trans: everything, dir: out
//
func (v Event) ParseQos() (type1 QOSTypeEnum, proportion float64, diff int64, timestamp uint64) {
	iv, err := _I.Get(459, "Event", "parse_qos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_proportion := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_diff := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_timestamp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_type1, arg_proportion, arg_diff, arg_timestamp}
	iv.Call(args, nil, &outArgs[0])
	type1 = QOSTypeEnum(outArgs[0].Int())
	proportion = outArgs[1].Double()
	diff = outArgs[2].Int64()
	timestamp = outArgs[3].Uint64()
	return
}

// gst_event_parse_seek
//
// [ rate ] trans: everything, dir: out
//
// [ format ] trans: everything, dir: out
//
// [ flags ] trans: everything, dir: out
//
// [ start_type ] trans: everything, dir: out
//
// [ start ] trans: everything, dir: out
//
// [ stop_type ] trans: everything, dir: out
//
// [ stop ] trans: everything, dir: out
//
func (v Event) ParseSeek() (rate float64, format FormatEnum, flags SeekFlags, start_type SeekTypeEnum, start int64, stop_type SeekTypeEnum, stop int64) {
	iv, err := _I.Get(460, "Event", "parse_seek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [7]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_rate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_flags := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_start_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	arg_stop_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[5]))
	arg_stop := gi.NewPointerArgument(unsafe.Pointer(&outArgs[6]))
	args := []gi.Argument{arg_v, arg_rate, arg_format, arg_flags, arg_start_type, arg_start, arg_stop_type, arg_stop}
	iv.Call(args, nil, &outArgs[0])
	rate = outArgs[0].Double()
	format = FormatEnum(outArgs[1].Int())
	flags = SeekFlags(outArgs[2].Int())
	start_type = SeekTypeEnum(outArgs[3].Int())
	start = outArgs[4].Int64()
	stop_type = SeekTypeEnum(outArgs[5].Int())
	stop = outArgs[6].Int64()
	return
}

// gst_event_parse_segment
//
// [ segment ] trans: nothing, dir: out
//
func (v Event) ParseSegment() (segment Segment) {
	iv, err := _I.Get(461, "Event", "parse_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_segment := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_segment}
	iv.Call(args, nil, &outArgs[0])
	segment.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_segment_done
//
// [ format ] trans: everything, dir: out
//
// [ position ] trans: everything, dir: out
//
func (v Event) ParseSegmentDone() (format FormatEnum, position int64) {
	iv, err := _I.Get(462, "Event", "parse_segment_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_position := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_format, arg_position}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	position = outArgs[1].Int64()
	return
}

// gst_event_parse_select_streams
//
// [ streams ] trans: everything, dir: out
//
func (v Event) ParseSelectStreams() (streams g.List) {
	iv, err := _I.Get(463, "Event", "parse_select_streams")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_streams := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_streams}
	iv.Call(args, nil, &outArgs[0])
	streams.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_sink_message
//
// [ msg ] trans: everything, dir: out
//
func (v Event) ParseSinkMessage() (msg Message) {
	iv, err := _I.Get(464, "Event", "parse_sink_message")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_msg := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_msg}
	iv.Call(args, nil, &outArgs[0])
	msg.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_step
//
// [ format ] trans: everything, dir: out
//
// [ amount ] trans: everything, dir: out
//
// [ rate ] trans: everything, dir: out
//
// [ flush ] trans: everything, dir: out
//
// [ intermediate ] trans: everything, dir: out
//
func (v Event) ParseStep() (format FormatEnum, amount uint64, rate float64, flush bool, intermediate bool) {
	iv, err := _I.Get(465, "Event", "parse_step")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [5]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_amount := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_rate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_flush := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_intermediate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_v, arg_format, arg_amount, arg_rate, arg_flush, arg_intermediate}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	amount = outArgs[1].Uint64()
	rate = outArgs[2].Double()
	flush = outArgs[3].Bool()
	intermediate = outArgs[4].Bool()
	return
}

// gst_event_parse_stream
//
// [ stream ] trans: everything, dir: out
//
func (v Event) ParseStream() (stream Stream) {
	iv, err := _I.Get(466, "Event", "parse_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream}
	iv.Call(args, nil, &outArgs[0])
	stream.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_stream_collection
//
// [ collection ] trans: everything, dir: out
//
func (v Event) ParseStreamCollection() (collection StreamCollection) {
	iv, err := _I.Get(467, "Event", "parse_stream_collection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_collection := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_collection}
	iv.Call(args, nil, &outArgs[0])
	collection.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_stream_flags
//
// [ flags ] trans: everything, dir: out
//
func (v Event) ParseStreamFlags() (flags StreamFlags) {
	iv, err := _I.Get(468, "Event", "parse_stream_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, &outArgs[0])
	flags = StreamFlags(outArgs[0].Int())
	return
}

// gst_event_parse_stream_group_done
//
// [ group_id ] trans: everything, dir: out
//
func (v Event) ParseStreamGroupDone() (group_id uint32) {
	iv, err := _I.Get(469, "Event", "parse_stream_group_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_id}
	iv.Call(args, nil, &outArgs[0])
	group_id = outArgs[0].Uint32()
	return
}

// gst_event_parse_stream_start
//
// [ stream_id ] trans: nothing, dir: out
//
func (v Event) ParseStreamStart() (stream_id string) {
	iv, err := _I.Get(470, "Event", "parse_stream_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream_id}
	iv.Call(args, nil, &outArgs[0])
	stream_id = outArgs[0].String().Copy()
	return
}

// gst_event_parse_tag
//
// [ taglist ] trans: nothing, dir: out
//
func (v Event) ParseTag() (taglist TagList) {
	iv, err := _I.Get(471, "Event", "parse_tag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_taglist := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_taglist}
	iv.Call(args, nil, &outArgs[0])
	taglist.P = outArgs[0].Pointer()
	return
}

// gst_event_parse_toc
//
// [ toc ] trans: everything, dir: out
//
// [ updated ] trans: everything, dir: out
//
func (v Event) ParseToc() (toc Toc, updated bool) {
	iv, err := _I.Get(472, "Event", "parse_toc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_toc := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_updated := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_toc, arg_updated}
	iv.Call(args, nil, &outArgs[0])
	toc.P = outArgs[0].Pointer()
	updated = outArgs[1].Bool()
	return
}

// gst_event_parse_toc_select
//
// [ uid ] trans: everything, dir: out
//
func (v Event) ParseTocSelect() (uid string) {
	iv, err := _I.Get(473, "Event", "parse_toc_select")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_uid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uid}
	iv.Call(args, nil, &outArgs[0])
	uid = outArgs[0].String().Take()
	return
}

// gst_event_set_group_id
//
// [ group_id ] trans: nothing
//
func (v Event) SetGroupId(group_id uint32) {
	iv, err := _I.Get(474, "Event", "set_group_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_id := gi.NewUint32Argument(group_id)
	args := []gi.Argument{arg_v, arg_group_id}
	iv.Call(args, nil, nil)
}

// gst_event_set_running_time_offset
//
// [ offset ] trans: nothing
//
func (v Event) SetRunningTimeOffset(offset int64) {
	iv, err := _I.Get(475, "Event", "set_running_time_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	iv.Call(args, nil, nil)
}

// gst_event_set_seqnum
//
// [ seqnum ] trans: nothing
//
func (v Event) SetSeqnum(seqnum uint32) {
	iv, err := _I.Get(476, "Event", "set_seqnum")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seqnum := gi.NewUint32Argument(seqnum)
	args := []gi.Argument{arg_v, arg_seqnum}
	iv.Call(args, nil, nil)
}

// gst_event_set_stream
//
// [ stream ] trans: nothing
//
func (v Event) SetStream(stream IStream) {
	iv, err := _I.Get(477, "Event", "set_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if stream != nil {
		tmp = stream.P_Stream()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_stream}
	iv.Call(args, nil, nil)
}

// gst_event_set_stream_flags
//
// [ flags ] trans: nothing
//
func (v Event) SetStreamFlags(flags StreamFlags) {
	iv, err := _I.Get(478, "Event", "set_stream_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, nil)
}

// gst_event_writable_structure
//
// [ result ] trans: nothing
//
func (v Event) WritableStructure() (result Structure) {
	iv, err := _I.Get(479, "Event", "writable_structure")
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

// Enum EventType
type EventTypeEnum int

const (
	EventTypeUnknown                EventTypeEnum = 0
	EventTypeFlushStart             EventTypeEnum = 2563
	EventTypeFlushStop              EventTypeEnum = 5127
	EventTypeStreamStart            EventTypeEnum = 10254
	EventTypeCaps                   EventTypeEnum = 12814
	EventTypeSegment                EventTypeEnum = 17934
	EventTypeStreamCollection       EventTypeEnum = 19230
	EventTypeTag                    EventTypeEnum = 20510
	EventTypeBuffersize             EventTypeEnum = 23054
	EventTypeSinkMessage            EventTypeEnum = 25630
	EventTypeStreamGroupDone        EventTypeEnum = 26894
	EventTypeEos                    EventTypeEnum = 28174
	EventTypeToc                    EventTypeEnum = 30750
	EventTypeProtection             EventTypeEnum = 33310
	EventTypeSegmentDone            EventTypeEnum = 38406
	EventTypeGap                    EventTypeEnum = 40966
	EventTypeQos                    EventTypeEnum = 48641
	EventTypeSeek                   EventTypeEnum = 51201
	EventTypeNavigation             EventTypeEnum = 53761
	EventTypeLatency                EventTypeEnum = 56321
	EventTypeStep                   EventTypeEnum = 58881
	EventTypeReconfigure            EventTypeEnum = 61441
	EventTypeTocSelect              EventTypeEnum = 64001
	EventTypeSelectStreams          EventTypeEnum = 66561
	EventTypeCustomUpstream         EventTypeEnum = 69121
	EventTypeCustomDownstream       EventTypeEnum = 71686
	EventTypeCustomDownstreamOob    EventTypeEnum = 74242
	EventTypeCustomDownstreamSticky EventTypeEnum = 76830
	EventTypeCustomBoth             EventTypeEnum = 79367
	EventTypeCustomBothOob          EventTypeEnum = 81923
)

func EventTypeGetType() gi.GType {
	ret := _I.GetGType(61, "EventType")
	return ret
}

// Flags EventTypeFlags
type EventTypeFlags int

const (
	EventTypeFlagsUpstream    EventTypeFlags = 1
	EventTypeFlagsDownstream  EventTypeFlags = 2
	EventTypeFlagsSerialized  EventTypeFlags = 4
	EventTypeFlagsSticky      EventTypeFlags = 8
	EventTypeFlagsStickyMulti EventTypeFlags = 16
)

func EventTypeFlagsGetType() gi.GType {
	ret := _I.GetGType(62, "EventTypeFlags")
	return ret
}

// Object FlagSet
type FlagSet struct {
	P unsafe.Pointer
}

func WrapFlagSet(p unsafe.Pointer) (r FlagSet) { r.P = p; return }

type IFlagSet interface{ P_FlagSet() unsafe.Pointer }

func (v FlagSet) P_FlagSet() unsafe.Pointer { return v.P }
func FlagSetGetType() gi.GType {
	ret := _I.GetGType(63, "FlagSet")
	return ret
}

// gst_flagset_register
//
// [ flags_type ] trans: nothing
//
// [ result ] trans: nothing
//
func FlagSetRegister1(flags_type gi.GType) (result gi.GType) {
	iv, err := _I.Get(480, "FlagSet", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags_type := gi.NewUintArgument(uint(flags_type))
	args := []gi.Argument{arg_flags_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// Enum FlowReturn
type FlowReturnEnum int

const (
	FlowReturnCustomSuccess2 FlowReturnEnum = 102
	FlowReturnCustomSuccess1 FlowReturnEnum = 101
	FlowReturnCustomSuccess  FlowReturnEnum = 100
	FlowReturnOk             FlowReturnEnum = 0
	FlowReturnNotLinked      FlowReturnEnum = -1
	FlowReturnFlushing       FlowReturnEnum = -2
	FlowReturnEos            FlowReturnEnum = -3
	FlowReturnNotNegotiated  FlowReturnEnum = -4
	FlowReturnError          FlowReturnEnum = -5
	FlowReturnNotSupported   FlowReturnEnum = -6
	FlowReturnCustomError    FlowReturnEnum = -100
	FlowReturnCustomError1   FlowReturnEnum = -101
	FlowReturnCustomError2   FlowReturnEnum = -102
)

func FlowReturnGetType() gi.GType {
	ret := _I.GetGType(64, "FlowReturn")
	return ret
}

// Enum Format
type FormatEnum int

const (
	FormatUndefined FormatEnum = 0
	FormatDefault   FormatEnum = 1
	FormatBytes     FormatEnum = 2
	FormatTime      FormatEnum = 3
	FormatBuffers   FormatEnum = 4
	FormatPercent   FormatEnum = 5
)

func FormatGetType() gi.GType {
	ret := _I.GetGType(65, "Format")
	return ret
}

// Struct FormatDefinition
type FormatDefinition struct {
	P unsafe.Pointer
}

const SizeOfStructFormatDefinition = 32

func FormatDefinitionGetType() gi.GType {
	ret := _I.GetGType(66, "FormatDefinition")
	return ret
}

// Object Fraction
type Fraction struct {
	P unsafe.Pointer
}

func WrapFraction(p unsafe.Pointer) (r Fraction) { r.P = p; return }

type IFraction interface{ P_Fraction() unsafe.Pointer }

func (v Fraction) P_Fraction() unsafe.Pointer { return v.P }
func FractionGetType() gi.GType {
	ret := _I.GetGType(67, "Fraction")
	return ret
}

// Object FractionRange
type FractionRange struct {
	P unsafe.Pointer
}

func WrapFractionRange(p unsafe.Pointer) (r FractionRange) { r.P = p; return }

type IFractionRange interface{ P_FractionRange() unsafe.Pointer }

func (v FractionRange) P_FractionRange() unsafe.Pointer { return v.P }
func FractionRangeGetType() gi.GType {
	ret := _I.GetGType(68, "FractionRange")
	return ret
}

// Object GhostPad
type GhostPad struct {
	ProxyPad
}

func WrapGhostPad(p unsafe.Pointer) (r GhostPad) { r.P = p; return }

type IGhostPad interface{ P_GhostPad() unsafe.Pointer }

func (v GhostPad) P_GhostPad() unsafe.Pointer { return v.P }
func GhostPadGetType() gi.GType {
	ret := _I.GetGType(69, "GhostPad")
	return ret
}

// gst_ghost_pad_new
//
// [ name ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func NewGhostPad(name string, target IPad) (result GhostPad) {
	iv, err := _I.Get(481, "GhostPad", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Pad()
	}
	arg_name := gi.NewStringArgument(c_name)
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_name, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_ghost_pad_new_from_template
//
// [ name ] trans: nothing
//
// [ target ] trans: nothing
//
// [ templ ] trans: nothing
//
// [ result ] trans: nothing
//
func NewGhostPadFromTemplate(name string, target IPad, templ IPadTemplate) (result GhostPad) {
	iv, err := _I.Get(482, "GhostPad", "new_from_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if templ != nil {
		tmp1 = templ.P_PadTemplate()
	}
	arg_name := gi.NewStringArgument(c_name)
	arg_target := gi.NewPointerArgument(tmp)
	arg_templ := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_name, arg_target, arg_templ}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_ghost_pad_new_no_target
//
// [ name ] trans: nothing
//
// [ dir ] trans: nothing
//
// [ result ] trans: nothing
//
func NewGhostPadNoTarget(name string, dir PadDirectionEnum) (result GhostPad) {
	iv, err := _I.Get(483, "GhostPad", "new_no_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_dir := gi.NewIntArgument(int(dir))
	args := []gi.Argument{arg_name, arg_dir}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_ghost_pad_new_no_target_from_template
//
// [ name ] trans: nothing
//
// [ templ ] trans: nothing
//
// [ result ] trans: nothing
//
func NewGhostPadNoTargetFromTemplate(name string, templ IPadTemplate) (result GhostPad) {
	iv, err := _I.Get(484, "GhostPad", "new_no_target_from_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	var tmp unsafe.Pointer
	if templ != nil {
		tmp = templ.P_PadTemplate()
	}
	arg_name := gi.NewStringArgument(c_name)
	arg_templ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_name, arg_templ}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_ghost_pad_activate_mode_default
//
// [ pad ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func GhostPadActivateModeDefault1(pad IPad, parent IObject, mode PadModeEnum, active bool) (result bool) {
	iv, err := _I.Get(485, "GhostPad", "activate_mode_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	arg_pad := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_pad, arg_parent, arg_mode, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_ghost_pad_internal_activate_mode_default
//
// [ pad ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func GhostPadInternalActivateModeDefault1(pad IPad, parent IObject, mode PadModeEnum, active bool) (result bool) {
	iv, err := _I.Get(486, "GhostPad", "internal_activate_mode_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	arg_pad := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_pad, arg_parent, arg_mode, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_ghost_pad_construct
//
// [ result ] trans: nothing
//
func (v GhostPad) Construct() (result bool) {
	iv, err := _I.Get(487, "GhostPad", "construct")
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

// gst_ghost_pad_get_target
//
// [ result ] trans: everything
//
func (v GhostPad) GetTarget() (result Pad) {
	iv, err := _I.Get(488, "GhostPad", "get_target")
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

// gst_ghost_pad_set_target
//
// [ newtarget ] trans: nothing
//
// [ result ] trans: nothing
//
func (v GhostPad) SetTarget(newtarget IPad) (result bool) {
	iv, err := _I.Get(489, "GhostPad", "set_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if newtarget != nil {
		tmp = newtarget.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_newtarget := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_newtarget}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct GhostPadClass

// Struct GhostPadPrivate
type GhostPadPrivate struct {
	P unsafe.Pointer
}

func GhostPadPrivateGetType() gi.GType {
	ret := _I.GetGType(70, "GhostPadPrivate")
	return ret
}

// Object Int64Range
type Int64Range struct {
	P unsafe.Pointer
}

func WrapInt64Range(p unsafe.Pointer) (r Int64Range) { r.P = p; return }

type IInt64Range interface{ P_Int64Range() unsafe.Pointer }

func (v Int64Range) P_Int64Range() unsafe.Pointer { return v.P }
func Int64RangeGetType() gi.GType {
	ret := _I.GetGType(71, "Int64Range")
	return ret
}

// Object IntRange
type IntRange struct {
	P unsafe.Pointer
}

func WrapIntRange(p unsafe.Pointer) (r IntRange) { r.P = p; return }

type IIntRange interface{ P_IntRange() unsafe.Pointer }

func (v IntRange) P_IntRange() unsafe.Pointer { return v.P }
func IntRangeGetType() gi.GType {
	ret := _I.GetGType(72, "IntRange")
	return ret
}

// Struct Iterator
type Iterator struct {
	P unsafe.Pointer
}

const SizeOfStructIterator = 120

func IteratorGetType() gi.GType {
	ret := _I.GetGType(73, "Iterator")
	return ret
}

// gst_iterator_new_single
//
// [ type1 ] trans: nothing
//
// [ object ] trans: nothing
//
// [ result ] trans: everything
//
func NewIteratorSingle(type1 gi.GType, object g.Value) (result Iterator) {
	iv, err := _I.Get(490, "Iterator", "new_single")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_object := gi.NewPointerArgument(object.P)
	args := []gi.Argument{arg_type1, arg_object}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_iterator_copy
//
// [ result ] trans: everything
//
func (v Iterator) Copy() (result Iterator) {
	iv, err := _I.Get(491, "Iterator", "copy")
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

// gst_iterator_filter
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: everything
//
func (v Iterator) Filter(func1 int /*TODO_TYPE CALLBACK*/, user_data g.Value) (result Iterator) {
	iv, err := _I.Get(492, "Iterator", "filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myCompareFunc()))
	arg_user_data := gi.NewPointerArgument(user_data.P)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_iterator_find_custom
//
// [ func1 ] trans: nothing
//
// [ elem ] trans: nothing, dir: out
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Iterator) FindCustom(func1 int /*TODO_TYPE CALLBACK*/, elem g.Value, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(493, "Iterator", "find_custom")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myCompareFunc()))
	arg_elem := gi.NewPointerArgument(elem.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_elem, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_iterator_fold
//
// [ func1 ] trans: nothing
//
// [ ret ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Iterator) Fold(func1 int /*TODO_TYPE CALLBACK*/, ret g.Value, user_data unsafe.Pointer) (result IteratorResultEnum) {
	iv, err := _I.Get(494, "Iterator", "fold")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myIteratorFoldFunction()))
	arg_ret := gi.NewPointerArgument(ret.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_ret, arg_user_data}
	var ret1 gi.Argument
	iv.Call(args, &ret1, nil)
	result = IteratorResultEnum(ret1.Int())
	return
}

// gst_iterator_foreach
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Iterator) Foreach(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result IteratorResultEnum) {
	iv, err := _I.Get(495, "Iterator", "foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myIteratorForeachFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IteratorResultEnum(ret.Int())
	return
}

// gst_iterator_free
//
func (v Iterator) Free() {
	iv, err := _I.Get(496, "Iterator", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_iterator_next
//
// [ elem ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Iterator) Next(elem g.Value) (result IteratorResultEnum) {
	iv, err := _I.Get(497, "Iterator", "next")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_elem := gi.NewPointerArgument(elem.P)
	args := []gi.Argument{arg_v, arg_elem}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IteratorResultEnum(ret.Int())
	return
}

// gst_iterator_push
//
// [ other ] trans: nothing
//
func (v Iterator) Push(other Iterator) {
	iv, err := _I.Get(498, "Iterator", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_other := gi.NewPointerArgument(other.P)
	args := []gi.Argument{arg_v, arg_other}
	iv.Call(args, nil, nil)
}

// gst_iterator_resync
//
func (v Iterator) Resync() {
	iv, err := _I.Get(499, "Iterator", "resync")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

type IteratorCopyFunctionStruct struct {
	F_it    Iterator
	F_copy1 Iterator
}

func GetPointer_myIteratorCopyFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorCopyFunction())
}

//export myGstIteratorCopyFunction
func myGstIteratorCopyFunction(it *C.GstIterator, copy1 *C.GstIterator) {
	// TODO: not found user_data
}

type IteratorFoldFunctionStruct struct {
	F_item g.Value
	F_ret  g.Value
}

func GetPointer_myIteratorFoldFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorFoldFunction())
}

//export myGstIteratorFoldFunction
func myGstIteratorFoldFunction(item *C.GValue, ret *C.GValue, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &IteratorFoldFunctionStruct{
		F_item: g.Value{P: unsafe.Pointer(item)},
		F_ret:  g.Value{P: unsafe.Pointer(ret)},
	}
	fn(args)
}

type IteratorForeachFunctionStruct struct {
	F_item g.Value
}

func GetPointer_myIteratorForeachFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorForeachFunction())
}

//export myGstIteratorForeachFunction
func myGstIteratorForeachFunction(item *C.GValue, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &IteratorForeachFunctionStruct{
		F_item: g.Value{P: unsafe.Pointer(item)},
	}
	fn(args)
}

type IteratorFreeFunctionStruct struct {
	F_it Iterator
}

func GetPointer_myIteratorFreeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorFreeFunction())
}

//export myGstIteratorFreeFunction
func myGstIteratorFreeFunction(it *C.GstIterator) {
	// TODO: not found user_data
}

// Enum IteratorItem
type IteratorItemEnum int

const (
	IteratorItemSkip IteratorItemEnum = 0
	IteratorItemPass IteratorItemEnum = 1
	IteratorItemEnd  IteratorItemEnum = 2
)

func IteratorItemGetType() gi.GType {
	ret := _I.GetGType(74, "IteratorItem")
	return ret
}

type IteratorItemFunctionStruct struct {
	F_it   Iterator
	F_item g.Value
}

func GetPointer_myIteratorItemFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorItemFunction())
}

//export myGstIteratorItemFunction
func myGstIteratorItemFunction(it *C.GstIterator, item *C.GValue) {
	// TODO: not found user_data
}

type IteratorNextFunctionStruct struct {
	F_it     Iterator
	F_result g.Value
}

func GetPointer_myIteratorNextFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorNextFunction())
}

//export myGstIteratorNextFunction
func myGstIteratorNextFunction(it *C.GstIterator, result *C.GValue) {
	// TODO: not found user_data
}

// Enum IteratorResult
type IteratorResultEnum int

const (
	IteratorResultDone   IteratorResultEnum = 0
	IteratorResultOk     IteratorResultEnum = 1
	IteratorResultResync IteratorResultEnum = 2
	IteratorResultError  IteratorResultEnum = 3
)

func IteratorResultGetType() gi.GType {
	ret := _I.GetGType(75, "IteratorResult")
	return ret
}

type IteratorResyncFunctionStruct struct {
	F_it Iterator
}

func GetPointer_myIteratorResyncFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstIteratorResyncFunction())
}

//export myGstIteratorResyncFunction
func myGstIteratorResyncFunction(it *C.GstIterator) {
	// TODO: not found user_data
}

// Enum LibraryError
type LibraryErrorEnum int

const (
	LibraryErrorFailed    LibraryErrorEnum = 1
	LibraryErrorTooLazy   LibraryErrorEnum = 2
	LibraryErrorInit      LibraryErrorEnum = 3
	LibraryErrorShutdown  LibraryErrorEnum = 4
	LibraryErrorSettings  LibraryErrorEnum = 5
	LibraryErrorEncode    LibraryErrorEnum = 6
	LibraryErrorNumErrors LibraryErrorEnum = 7
)

func LibraryErrorGetType() gi.GType {
	ret := _I.GetGType(76, "LibraryError")
	return ret
}

// Flags LockFlags
type LockFlags int

const (
	LockFlagsRead      LockFlags = 1
	LockFlagsWrite     LockFlags = 2
	LockFlagsExclusive LockFlags = 4
	LockFlagsLast      LockFlags = 256
)

func LockFlagsGetType() gi.GType {
	ret := _I.GetGType(77, "LockFlags")
	return ret
}

type LogFunctionStruct struct {
	F_category DebugCategory
	F_level    DebugLevelEnum
	F_file     string
	F_function string
	F_line     int32
	F_object   g.Object
	F_message  DebugMessage
}

func GetPointer_myLogFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstLogFunction())
}

//export myGstLogFunction
func myGstLogFunction(category *C.GstDebugCategory, level C.GstDebugLevel, file *C.gchar, function *C.gchar, line C.gint32, object *C.GObject, message *C.GstDebugMessage, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &LogFunctionStruct{
		F_category: DebugCategory{P: unsafe.Pointer(category)},
		F_level:    DebugLevelEnum(level),
		F_file:     gi.GoString(unsafe.Pointer(file)),
		F_function: gi.GoString(unsafe.Pointer(function)),
		F_line:     int32(line),
		F_object:   g.WrapObject(unsafe.Pointer(object)),
		F_message:  DebugMessage{P: unsafe.Pointer(message)},
	}
	fn(args)
}

// Flags MapFlags
type MapFlags int

const (
	MapFlagsRead     MapFlags = 1
	MapFlagsWrite    MapFlags = 2
	MapFlagsFlagLast MapFlags = 65536
)

func MapFlagsGetType() gi.GType {
	ret := _I.GetGType(78, "MapFlags")
	return ret
}

// Struct MapInfo
type MapInfo struct {
	P unsafe.Pointer
}

const SizeOfStructMapInfo = 104

func MapInfoGetType() gi.GType {
	ret := _I.GetGType(79, "MapInfo")
	return ret
}

// Struct Memory
type Memory struct {
	P unsafe.Pointer
}

const SizeOfStructMemory = 112

func MemoryGetType() gi.GType {
	ret := _I.GetGType(80, "Memory")
	return ret
}

// gst_memory_new_wrapped
//
// [ flags ] trans: nothing
//
// [ data ] trans: nothing
//
// [ maxsize ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: everything
//
func NewMemoryWrapped(flags MemoryFlags, data gi.Uint8Array, maxsize uint64, offset uint64, size uint64, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result Memory) {
	iv, err := _I.Get(500, "Memory", "new_wrapped")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags := gi.NewIntArgument(int(flags))
	arg_data := gi.NewPointerArgument(data.P)
	arg_maxsize := gi.NewUint64Argument(maxsize)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_flags, arg_data, arg_maxsize, arg_offset, arg_size, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_memory_copy
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v Memory) Copy(offset int64, size int64) (result Memory) {
	iv, err := _I.Get(501, "Memory", "copy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_memory_get_sizes
//
// [ offset ] trans: everything, dir: out
//
// [ maxsize ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Memory) GetSizes() (result uint64, offset uint64, maxsize uint64) {
	iv, err := _I.Get(502, "Memory", "get_sizes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_maxsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_maxsize}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	offset = outArgs[0].Uint64()
	maxsize = outArgs[1].Uint64()
	result = ret.Uint64()
	return
}

// gst_memory_is_span
//
// [ mem2 ] trans: nothing
//
// [ offset ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Memory) IsSpan(mem2 Memory) (result bool, offset uint64) {
	iv, err := _I.Get(503, "Memory", "is_span")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mem2 := gi.NewPointerArgument(mem2.P)
	arg_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_mem2, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	offset = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_memory_is_type
//
// [ mem_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Memory) IsType(mem_type string) (result bool) {
	iv, err := _I.Get(504, "Memory", "is_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mem_type := gi.CString(mem_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_mem_type := gi.NewStringArgument(c_mem_type)
	args := []gi.Argument{arg_v, arg_mem_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_mem_type)
	result = ret.Bool()
	return
}

// gst_memory_make_mapped
//
// [ info ] trans: nothing, dir: out
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func (v Memory) MakeMapped(info MapInfo, flags MapFlags) (result Memory) {
	iv, err := _I.Get(505, "Memory", "make_mapped")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_info, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_memory_map
//
// [ info ] trans: nothing, dir: out
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Memory) Map(info MapInfo, flags MapFlags) (result bool) {
	iv, err := _I.Get(506, "Memory", "map")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_info, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_memory_resize
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
func (v Memory) Resize(offset int64, size uint64) {
	iv, err := _I.Get(507, "Memory", "resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_size}
	iv.Call(args, nil, nil)
}

// gst_memory_share
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ result ] trans: everything
//
func (v Memory) Share(offset int64, size int64) (result Memory) {
	iv, err := _I.Get(508, "Memory", "share")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_v, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_memory_unmap
//
// [ info ] trans: nothing
//
func (v Memory) Unmap(info MapInfo) {
	iv, err := _I.Get(509, "Memory", "unmap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_v, arg_info}
	iv.Call(args, nil, nil)
}

type MemoryCopyFunctionStruct struct {
	F_mem    Memory
	F_offset int64
	F_size   int64
}

func GetPointer_myMemoryCopyFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryCopyFunction())
}

//export myGstMemoryCopyFunction
func myGstMemoryCopyFunction(mem *C.GstMemory, offset C.gint64, size C.gint64) {
	// TODO: not found user_data
}

// Flags MemoryFlags
type MemoryFlags int

const (
	MemoryFlagsReadonly             MemoryFlags = 2
	MemoryFlagsNoShare              MemoryFlags = 16
	MemoryFlagsZeroPrefixed         MemoryFlags = 32
	MemoryFlagsZeroPadded           MemoryFlags = 64
	MemoryFlagsPhysicallyContiguous MemoryFlags = 128
	MemoryFlagsNotMappable          MemoryFlags = 256
	MemoryFlagsLast                 MemoryFlags = 1048576
)

func MemoryFlagsGetType() gi.GType {
	ret := _I.GetGType(81, "MemoryFlags")
	return ret
}

type MemoryIsSpanFunctionStruct struct {
	F_mem1   Memory
	F_mem2   Memory
	F_offset *uint64
}

func GetPointer_myMemoryIsSpanFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryIsSpanFunction())
}

//export myGstMemoryIsSpanFunction
func myGstMemoryIsSpanFunction(mem1 *C.GstMemory, mem2 *C.GstMemory, offset *C.guint64) {
	// TODO: not found user_data
}

type MemoryMapFullFunctionStruct struct {
	F_mem     Memory
	F_info    MapInfo
	F_maxsize uint64
}

func GetPointer_myMemoryMapFullFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryMapFullFunction())
}

//export myGstMemoryMapFullFunction
func myGstMemoryMapFullFunction(mem *C.GstMemory, info *C.GstMapInfo, maxsize C.guint64) {
	// TODO: not found user_data
}

type MemoryMapFunctionStruct struct {
	F_mem     Memory
	F_maxsize uint64
	F_flags   MapFlags
}

func GetPointer_myMemoryMapFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryMapFunction())
}

//export myGstMemoryMapFunction
func myGstMemoryMapFunction(mem *C.GstMemory, maxsize C.guint64, flags C.GstMapFlags) {
	// TODO: not found user_data
}

type MemoryShareFunctionStruct struct {
	F_mem    Memory
	F_offset int64
	F_size   int64
}

func GetPointer_myMemoryShareFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryShareFunction())
}

//export myGstMemoryShareFunction
func myGstMemoryShareFunction(mem *C.GstMemory, offset C.gint64, size C.gint64) {
	// TODO: not found user_data
}

type MemoryUnmapFullFunctionStruct struct {
	F_mem  Memory
	F_info MapInfo
}

func GetPointer_myMemoryUnmapFullFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryUnmapFullFunction())
}

//export myGstMemoryUnmapFullFunction
func myGstMemoryUnmapFullFunction(mem *C.GstMemory, info *C.GstMapInfo) {
	// TODO: not found user_data
}

type MemoryUnmapFunctionStruct struct {
	F_mem Memory
}

func GetPointer_myMemoryUnmapFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMemoryUnmapFunction())
}

//export myGstMemoryUnmapFunction
func myGstMemoryUnmapFunction(mem *C.GstMemory) {
	// TODO: not found user_data
}

// Struct Message
type Message struct {
	P unsafe.Pointer
}

const SizeOfStructMessage = 120

func MessageGetType() gi.GType {
	ret := _I.GetGType(82, "Message")
	return ret
}

// gst_message_new_application
//
// [ src ] trans: nothing
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageApplication(src IObject, structure Structure) (result Message) {
	iv, err := _I.Get(510, "Message", "new_application")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_src, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_async_done
//
// [ src ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageAsyncDone(src IObject, running_time uint64) (result Message) {
	iv, err := _I.Get(511, "Message", "new_async_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_running_time := gi.NewUint64Argument(running_time)
	args := []gi.Argument{arg_src, arg_running_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_async_start
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageAsyncStart(src IObject) (result Message) {
	iv, err := _I.Get(512, "Message", "new_async_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_buffering
//
// [ src ] trans: nothing
//
// [ percent ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageBuffering(src IObject, percent int32) (result Message) {
	iv, err := _I.Get(513, "Message", "new_buffering")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_percent := gi.NewInt32Argument(percent)
	args := []gi.Argument{arg_src, arg_percent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_clock_lost
//
// [ src ] trans: nothing
//
// [ clock ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageClockLost(src IObject, clock IClock) (result Message) {
	iv, err := _I.Get(514, "Message", "new_clock_lost")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if clock != nil {
		tmp1 = clock.P_Clock()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_clock := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_clock}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_clock_provide
//
// [ src ] trans: nothing
//
// [ clock ] trans: nothing
//
// [ ready ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageClockProvide(src IObject, clock IClock, ready bool) (result Message) {
	iv, err := _I.Get(515, "Message", "new_clock_provide")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if clock != nil {
		tmp1 = clock.P_Clock()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_clock := gi.NewPointerArgument(tmp1)
	arg_ready := gi.NewBoolArgument(ready)
	args := []gi.Argument{arg_src, arg_clock, arg_ready}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_custom
//
// [ type1 ] trans: nothing
//
// [ src ] trans: nothing
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageCustom(type1 MessageTypeFlags, src IObject, structure Structure) (result Message) {
	iv, err := _I.Get(516, "Message", "new_custom")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_src := gi.NewPointerArgument(tmp)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_type1, arg_src, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_device_added
//
// [ src ] trans: nothing
//
// [ device ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageDeviceAdded(src IObject, device IDevice) (result Message) {
	iv, err := _I.Get(517, "Message", "new_device_added")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if device != nil {
		tmp1 = device.P_Device()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_device := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_device}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_device_removed
//
// [ src ] trans: nothing
//
// [ device ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageDeviceRemoved(src IObject, device IDevice) (result Message) {
	iv, err := _I.Get(518, "Message", "new_device_removed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if device != nil {
		tmp1 = device.P_Device()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_device := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_device}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_duration_changed
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageDurationChanged(src IObject) (result Message) {
	iv, err := _I.Get(519, "Message", "new_duration_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_element
//
// [ src ] trans: nothing
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageElement(src IObject, structure Structure) (result Message) {
	iv, err := _I.Get(520, "Message", "new_element")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_src, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_eos
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageEos(src IObject) (result Message) {
	iv, err := _I.Get(521, "Message", "new_eos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_error
//
// [ src ] trans: nothing
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageError(src IObject, error g.Error, debug string) (result Message) {
	iv, err := _I.Get(522, "Message", "new_error")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_debug := gi.CString(debug)
	arg_src := gi.NewPointerArgument(tmp)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	args := []gi.Argument{arg_src, arg_error, arg_debug}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_debug)
	result.P = ret.Pointer()
	return
}

// gst_message_new_error_with_details
//
// [ src ] trans: nothing
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
// [ details ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageErrorWithDetails(src IObject, error g.Error, debug string, details Structure) (result Message) {
	iv, err := _I.Get(523, "Message", "new_error_with_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_debug := gi.CString(debug)
	arg_src := gi.NewPointerArgument(tmp)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	arg_details := gi.NewPointerArgument(details.P)
	args := []gi.Argument{arg_src, arg_error, arg_debug, arg_details}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_debug)
	result.P = ret.Pointer()
	return
}

// gst_message_new_have_context
//
// [ src ] trans: nothing
//
// [ context ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageHaveContext(src IObject, context Context) (result Message) {
	iv, err := _I.Get(524, "Message", "new_have_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_context := gi.NewPointerArgument(context.P)
	args := []gi.Argument{arg_src, arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_info
//
// [ src ] trans: nothing
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageInfo(src IObject, error g.Error, debug string) (result Message) {
	iv, err := _I.Get(525, "Message", "new_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_debug := gi.CString(debug)
	arg_src := gi.NewPointerArgument(tmp)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	args := []gi.Argument{arg_src, arg_error, arg_debug}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_debug)
	result.P = ret.Pointer()
	return
}

// gst_message_new_info_with_details
//
// [ src ] trans: nothing
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
// [ details ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageInfoWithDetails(src IObject, error g.Error, debug string, details Structure) (result Message) {
	iv, err := _I.Get(526, "Message", "new_info_with_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_debug := gi.CString(debug)
	arg_src := gi.NewPointerArgument(tmp)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	arg_details := gi.NewPointerArgument(details.P)
	args := []gi.Argument{arg_src, arg_error, arg_debug, arg_details}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_debug)
	result.P = ret.Pointer()
	return
}

// gst_message_new_latency
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageLatency(src IObject) (result Message) {
	iv, err := _I.Get(527, "Message", "new_latency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_need_context
//
// [ src ] trans: nothing
//
// [ context_type ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageNeedContext(src IObject, context_type string) (result Message) {
	iv, err := _I.Get(528, "Message", "new_need_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_context_type := gi.CString(context_type)
	arg_src := gi.NewPointerArgument(tmp)
	arg_context_type := gi.NewStringArgument(c_context_type)
	args := []gi.Argument{arg_src, arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_type)
	result.P = ret.Pointer()
	return
}

// gst_message_new_new_clock
//
// [ src ] trans: nothing
//
// [ clock ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageNewClock(src IObject, clock IClock) (result Message) {
	iv, err := _I.Get(529, "Message", "new_new_clock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if clock != nil {
		tmp1 = clock.P_Clock()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_clock := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_clock}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_progress
//
// [ src ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ code ] trans: nothing
//
// [ text ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageProgress(src IObject, type1 ProgressTypeEnum, code string, text string) (result Message) {
	iv, err := _I.Get(530, "Message", "new_progress")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_code := gi.CString(code)
	c_text := gi.CString(text)
	arg_src := gi.NewPointerArgument(tmp)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_code := gi.NewStringArgument(c_code)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_src, arg_type1, arg_code, arg_text}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_code)
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// gst_message_new_property_notify
//
// [ src ] trans: nothing
//
// [ property_name ] trans: nothing
//
// [ val ] trans: everything
//
// [ result ] trans: everything
//
func NewMessagePropertyNotify(src IObject, property_name string, val g.Value) (result Message) {
	iv, err := _I.Get(531, "Message", "new_property_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_property_name := gi.CString(property_name)
	arg_src := gi.NewPointerArgument(tmp)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_val := gi.NewPointerArgument(val.P)
	args := []gi.Argument{arg_src, arg_property_name, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// gst_message_new_qos
//
// [ src ] trans: nothing
//
// [ live ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ stream_time ] trans: nothing
//
// [ timestamp ] trans: nothing
//
// [ duration ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageQos(src IObject, live bool, running_time uint64, stream_time uint64, timestamp uint64, duration uint64) (result Message) {
	iv, err := _I.Get(532, "Message", "new_qos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_live := gi.NewBoolArgument(live)
	arg_running_time := gi.NewUint64Argument(running_time)
	arg_stream_time := gi.NewUint64Argument(stream_time)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_duration := gi.NewUint64Argument(duration)
	args := []gi.Argument{arg_src, arg_live, arg_running_time, arg_stream_time, arg_timestamp, arg_duration}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_redirect
//
// [ src ] trans: nothing
//
// [ location ] trans: nothing
//
// [ tag_list ] trans: everything
//
// [ entry_struct ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageRedirect(src IObject, location string, tag_list TagList, entry_struct Structure) (result Message) {
	iv, err := _I.Get(533, "Message", "new_redirect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_location := gi.CString(location)
	arg_src := gi.NewPointerArgument(tmp)
	arg_location := gi.NewStringArgument(c_location)
	arg_tag_list := gi.NewPointerArgument(tag_list.P)
	arg_entry_struct := gi.NewPointerArgument(entry_struct.P)
	args := []gi.Argument{arg_src, arg_location, arg_tag_list, arg_entry_struct}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_location)
	result.P = ret.Pointer()
	return
}

// gst_message_new_request_state
//
// [ src ] trans: nothing
//
// [ state ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageRequestState(src IObject, state StateEnum) (result Message) {
	iv, err := _I.Get(534, "Message", "new_request_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_src, arg_state}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_reset_time
//
// [ src ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageResetTime(src IObject, running_time uint64) (result Message) {
	iv, err := _I.Get(535, "Message", "new_reset_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_running_time := gi.NewUint64Argument(running_time)
	args := []gi.Argument{arg_src, arg_running_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_segment_done
//
// [ src ] trans: nothing
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageSegmentDone(src IObject, format FormatEnum, position int64) (result Message) {
	iv, err := _I.Get(536, "Message", "new_segment_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewInt64Argument(position)
	args := []gi.Argument{arg_src, arg_format, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_segment_start
//
// [ src ] trans: nothing
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageSegmentStart(src IObject, format FormatEnum, position int64) (result Message) {
	iv, err := _I.Get(537, "Message", "new_segment_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewInt64Argument(position)
	args := []gi.Argument{arg_src, arg_format, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_state_changed
//
// [ src ] trans: nothing
//
// [ oldstate ] trans: nothing
//
// [ newstate ] trans: nothing
//
// [ pending ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStateChanged(src IObject, oldstate StateEnum, newstate StateEnum, pending StateEnum) (result Message) {
	iv, err := _I.Get(538, "Message", "new_state_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_oldstate := gi.NewIntArgument(int(oldstate))
	arg_newstate := gi.NewIntArgument(int(newstate))
	arg_pending := gi.NewIntArgument(int(pending))
	args := []gi.Argument{arg_src, arg_oldstate, arg_newstate, arg_pending}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_state_dirty
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStateDirty(src IObject) (result Message) {
	iv, err := _I.Get(539, "Message", "new_state_dirty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_step_done
//
// [ src ] trans: nothing
//
// [ format ] trans: nothing
//
// [ amount ] trans: nothing
//
// [ rate ] trans: nothing
//
// [ flush ] trans: nothing
//
// [ intermediate ] trans: nothing
//
// [ duration ] trans: nothing
//
// [ eos ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStepDone(src IObject, format FormatEnum, amount uint64, rate float64, flush bool, intermediate bool, duration uint64, eos bool) (result Message) {
	iv, err := _I.Get(540, "Message", "new_step_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_format := gi.NewIntArgument(int(format))
	arg_amount := gi.NewUint64Argument(amount)
	arg_rate := gi.NewDoubleArgument(rate)
	arg_flush := gi.NewBoolArgument(flush)
	arg_intermediate := gi.NewBoolArgument(intermediate)
	arg_duration := gi.NewUint64Argument(duration)
	arg_eos := gi.NewBoolArgument(eos)
	args := []gi.Argument{arg_src, arg_format, arg_amount, arg_rate, arg_flush, arg_intermediate, arg_duration, arg_eos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_step_start
//
// [ src ] trans: nothing
//
// [ active ] trans: nothing
//
// [ format ] trans: nothing
//
// [ amount ] trans: nothing
//
// [ rate ] trans: nothing
//
// [ flush ] trans: nothing
//
// [ intermediate ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStepStart(src IObject, active bool, format FormatEnum, amount uint64, rate float64, flush bool, intermediate bool) (result Message) {
	iv, err := _I.Get(541, "Message", "new_step_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_active := gi.NewBoolArgument(active)
	arg_format := gi.NewIntArgument(int(format))
	arg_amount := gi.NewUint64Argument(amount)
	arg_rate := gi.NewDoubleArgument(rate)
	arg_flush := gi.NewBoolArgument(flush)
	arg_intermediate := gi.NewBoolArgument(intermediate)
	args := []gi.Argument{arg_src, arg_active, arg_format, arg_amount, arg_rate, arg_flush, arg_intermediate}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_stream_collection
//
// [ src ] trans: nothing
//
// [ collection ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStreamCollection(src IObject, collection IStreamCollection) (result Message) {
	iv, err := _I.Get(542, "Message", "new_stream_collection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if collection != nil {
		tmp1 = collection.P_StreamCollection()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_collection := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_collection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_stream_start
//
// [ src ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStreamStart(src IObject) (result Message) {
	iv, err := _I.Get(543, "Message", "new_stream_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_stream_status
//
// [ src ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ owner ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStreamStatus(src IObject, type1 StreamStatusTypeEnum, owner IElement) (result Message) {
	iv, err := _I.Get(544, "Message", "new_stream_status")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if owner != nil {
		tmp1 = owner.P_Element()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_owner := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_type1, arg_owner}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_streams_selected
//
// [ src ] trans: nothing
//
// [ collection ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStreamsSelected(src IObject, collection IStreamCollection) (result Message) {
	iv, err := _I.Get(545, "Message", "new_streams_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if collection != nil {
		tmp1 = collection.P_StreamCollection()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_collection := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_src, arg_collection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_structure_change
//
// [ src ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ owner ] trans: nothing
//
// [ busy ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageStructureChange(src IObject, type1 StructureChangeTypeEnum, owner IElement, busy bool) (result Message) {
	iv, err := _I.Get(546, "Message", "new_structure_change")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	var tmp1 unsafe.Pointer
	if owner != nil {
		tmp1 = owner.P_Element()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_owner := gi.NewPointerArgument(tmp1)
	arg_busy := gi.NewBoolArgument(busy)
	args := []gi.Argument{arg_src, arg_type1, arg_owner, arg_busy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_tag
//
// [ src ] trans: nothing
//
// [ tag_list ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageTag(src IObject, tag_list TagList) (result Message) {
	iv, err := _I.Get(547, "Message", "new_tag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_tag_list := gi.NewPointerArgument(tag_list.P)
	args := []gi.Argument{arg_src, arg_tag_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_toc
//
// [ src ] trans: nothing
//
// [ toc ] trans: nothing
//
// [ updated ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageToc(src IObject, toc Toc, updated bool) (result Message) {
	iv, err := _I.Get(548, "Message", "new_toc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	arg_src := gi.NewPointerArgument(tmp)
	arg_toc := gi.NewPointerArgument(toc.P)
	arg_updated := gi.NewBoolArgument(updated)
	args := []gi.Argument{arg_src, arg_toc, arg_updated}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_new_warning
//
// [ src ] trans: nothing
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
// [ result ] trans: everything
//
func NewMessageWarning(src IObject, error g.Error, debug string) (result Message) {
	iv, err := _I.Get(549, "Message", "new_warning")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_debug := gi.CString(debug)
	arg_src := gi.NewPointerArgument(tmp)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	args := []gi.Argument{arg_src, arg_error, arg_debug}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_debug)
	result.P = ret.Pointer()
	return
}

// gst_message_new_warning_with_details
//
// [ src ] trans: nothing
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
// [ details ] trans: everything
//
// [ result ] trans: everything
//
func NewMessageWarningWithDetails(src IObject, error g.Error, debug string, details Structure) (result Message) {
	iv, err := _I.Get(550, "Message", "new_warning_with_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if src != nil {
		tmp = src.P_Object()
	}
	c_debug := gi.CString(debug)
	arg_src := gi.NewPointerArgument(tmp)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	arg_details := gi.NewPointerArgument(details.P)
	args := []gi.Argument{arg_src, arg_error, arg_debug, arg_details}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_debug)
	result.P = ret.Pointer()
	return
}

// gst_message_add_redirect_entry
//
// [ location ] trans: nothing
//
// [ tag_list ] trans: everything
//
// [ entry_struct ] trans: everything
//
func (v Message) AddRedirectEntry(location string, tag_list TagList, entry_struct Structure) {
	iv, err := _I.Get(551, "Message", "add_redirect_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_location := gi.CString(location)
	arg_v := gi.NewPointerArgument(v.P)
	arg_location := gi.NewStringArgument(c_location)
	arg_tag_list := gi.NewPointerArgument(tag_list.P)
	arg_entry_struct := gi.NewPointerArgument(entry_struct.P)
	args := []gi.Argument{arg_v, arg_location, arg_tag_list, arg_entry_struct}
	iv.Call(args, nil, nil)
	gi.Free(c_location)
}

// gst_message_get_num_redirect_entries
//
// [ result ] trans: nothing
//
func (v Message) GetNumRedirectEntries() (result uint64) {
	iv, err := _I.Get(552, "Message", "get_num_redirect_entries")
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

// gst_message_get_seqnum
//
// [ result ] trans: nothing
//
func (v Message) GetSeqnum() (result uint32) {
	iv, err := _I.Get(553, "Message", "get_seqnum")
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

// gst_message_get_stream_status_object
//
// [ result ] trans: nothing
//
func (v Message) GetStreamStatusObject() (result g.Value) {
	iv, err := _I.Get(554, "Message", "get_stream_status_object")
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

// gst_message_get_structure
//
// [ result ] trans: nothing
//
func (v Message) GetStructure() (result Structure) {
	iv, err := _I.Get(555, "Message", "get_structure")
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

// gst_message_has_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Message) HasName(name string) (result bool) {
	iv, err := _I.Get(556, "Message", "has_name")
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
	result = ret.Bool()
	return
}

// gst_message_parse_async_done
//
// [ running_time ] trans: everything, dir: out
//
func (v Message) ParseAsyncDone() (running_time uint64) {
	iv, err := _I.Get(557, "Message", "parse_async_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_running_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_running_time}
	iv.Call(args, nil, &outArgs[0])
	running_time = outArgs[0].Uint64()
	return
}

// gst_message_parse_buffering
//
// [ percent ] trans: everything, dir: out
//
func (v Message) ParseBuffering() (percent int32) {
	iv, err := _I.Get(558, "Message", "parse_buffering")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_percent := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_percent}
	iv.Call(args, nil, &outArgs[0])
	percent = outArgs[0].Int32()
	return
}

// gst_message_parse_buffering_stats
//
// [ mode ] trans: everything, dir: out
//
// [ avg_in ] trans: everything, dir: out
//
// [ avg_out ] trans: everything, dir: out
//
// [ buffering_left ] trans: everything, dir: out
//
func (v Message) ParseBufferingStats() (mode BufferingModeEnum, avg_in int32, avg_out int32, buffering_left int64) {
	iv, err := _I.Get(559, "Message", "parse_buffering_stats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_avg_in := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_avg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_buffering_left := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_mode, arg_avg_in, arg_avg_out, arg_buffering_left}
	iv.Call(args, nil, &outArgs[0])
	mode = BufferingModeEnum(outArgs[0].Int())
	avg_in = outArgs[1].Int32()
	avg_out = outArgs[2].Int32()
	buffering_left = outArgs[3].Int64()
	return
}

// gst_message_parse_clock_lost
//
// [ clock ] trans: nothing, dir: out
//
func (v Message) ParseClockLost() (clock Clock) {
	iv, err := _I.Get(560, "Message", "parse_clock_lost")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_clock := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_clock}
	iv.Call(args, nil, &outArgs[0])
	clock.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_clock_provide
//
// [ clock ] trans: nothing, dir: out
//
// [ ready ] trans: everything, dir: out
//
func (v Message) ParseClockProvide() (clock Clock, ready bool) {
	iv, err := _I.Get(561, "Message", "parse_clock_provide")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_clock := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_ready := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_clock, arg_ready}
	iv.Call(args, nil, &outArgs[0])
	clock.P = outArgs[0].Pointer()
	ready = outArgs[1].Bool()
	return
}

// gst_message_parse_context_type
//
// [ context_type ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Message) ParseContextType() (result bool, context_type string) {
	iv, err := _I.Get(562, "Message", "parse_context_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_context_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	context_type = outArgs[0].String().Copy()
	result = ret.Bool()
	return
}

// gst_message_parse_device_added
//
// [ device ] trans: everything, dir: out
//
func (v Message) ParseDeviceAdded() (device Device) {
	iv, err := _I.Get(563, "Message", "parse_device_added")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, &outArgs[0])
	device.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_device_removed
//
// [ device ] trans: everything, dir: out
//
func (v Message) ParseDeviceRemoved() (device Device) {
	iv, err := _I.Get(564, "Message", "parse_device_removed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, &outArgs[0])
	device.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_error
//
// [ gerror ] trans: everything, dir: out
//
// [ debug ] trans: everything, dir: out
//
func (v Message) ParseError() (gerror g.Error, debug string) {
	iv, err := _I.Get(565, "Message", "parse_error")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_gerror := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_debug := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_gerror, arg_debug}
	iv.Call(args, nil, &outArgs[0])
	gerror.P = outArgs[0].Pointer()
	debug = outArgs[1].String().Take()
	return
}

// gst_message_parse_error_details
//
// [ structure ] trans: nothing, dir: out
//
func (v Message) ParseErrorDetails() (structure Structure) {
	iv, err := _I.Get(566, "Message", "parse_error_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_structure}
	iv.Call(args, nil, &outArgs[0])
	structure.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_group_id
//
// [ group_id ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Message) ParseGroupId() (result bool, group_id uint32) {
	iv, err := _I.Get(567, "Message", "parse_group_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_id}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	group_id = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_message_parse_have_context
//
// [ context ] trans: everything, dir: out
//
func (v Message) ParseHaveContext() (context Context) {
	iv, err := _I.Get(568, "Message", "parse_have_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_context}
	iv.Call(args, nil, &outArgs[0])
	context.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_info
//
// [ gerror ] trans: everything, dir: out
//
// [ debug ] trans: everything, dir: out
//
func (v Message) ParseInfo() (gerror g.Error, debug string) {
	iv, err := _I.Get(569, "Message", "parse_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_gerror := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_debug := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_gerror, arg_debug}
	iv.Call(args, nil, &outArgs[0])
	gerror.P = outArgs[0].Pointer()
	debug = outArgs[1].String().Take()
	return
}

// gst_message_parse_info_details
//
// [ structure ] trans: nothing, dir: out
//
func (v Message) ParseInfoDetails() (structure Structure) {
	iv, err := _I.Get(570, "Message", "parse_info_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_structure}
	iv.Call(args, nil, &outArgs[0])
	structure.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_new_clock
//
// [ clock ] trans: nothing, dir: out
//
func (v Message) ParseNewClock() (clock Clock) {
	iv, err := _I.Get(571, "Message", "parse_new_clock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_clock := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_clock}
	iv.Call(args, nil, &outArgs[0])
	clock.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_progress
//
// [ type1 ] trans: everything, dir: out
//
// [ code ] trans: everything, dir: out
//
// [ text ] trans: everything, dir: out
//
func (v Message) ParseProgress() (type1 ProgressTypeEnum, code string, text string) {
	iv, err := _I.Get(572, "Message", "parse_progress")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_code := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_text := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_type1, arg_code, arg_text}
	iv.Call(args, nil, &outArgs[0])
	type1 = ProgressTypeEnum(outArgs[0].Int())
	code = outArgs[1].String().Take()
	text = outArgs[2].String().Take()
	return
}

// gst_message_parse_property_notify
//
// [ object ] trans: nothing, dir: out
//
// [ property_name ] trans: nothing, dir: out
//
// [ property_value ] trans: nothing, dir: out
//
func (v Message) ParsePropertyNotify() (object Object, property_name string, property_value g.Value) {
	iv, err := _I.Get(573, "Message", "parse_property_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_object := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_property_name := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_property_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_object, arg_property_name, arg_property_value}
	iv.Call(args, nil, &outArgs[0])
	object.P = outArgs[0].Pointer()
	property_name = outArgs[1].String().Copy()
	property_value.P = outArgs[2].Pointer()
	return
}

// gst_message_parse_qos
//
// [ live ] trans: everything, dir: out
//
// [ running_time ] trans: everything, dir: out
//
// [ stream_time ] trans: everything, dir: out
//
// [ timestamp ] trans: everything, dir: out
//
// [ duration ] trans: everything, dir: out
//
func (v Message) ParseQos() (live bool, running_time uint64, stream_time uint64, timestamp uint64, duration uint64) {
	iv, err := _I.Get(574, "Message", "parse_qos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [5]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_live := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_running_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_stream_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_timestamp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_v, arg_live, arg_running_time, arg_stream_time, arg_timestamp, arg_duration}
	iv.Call(args, nil, &outArgs[0])
	live = outArgs[0].Bool()
	running_time = outArgs[1].Uint64()
	stream_time = outArgs[2].Uint64()
	timestamp = outArgs[3].Uint64()
	duration = outArgs[4].Uint64()
	return
}

// gst_message_parse_qos_stats
//
// [ format ] trans: everything, dir: out
//
// [ processed ] trans: everything, dir: out
//
// [ dropped ] trans: everything, dir: out
//
func (v Message) ParseQosStats() (format FormatEnum, processed uint64, dropped uint64) {
	iv, err := _I.Get(575, "Message", "parse_qos_stats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_processed := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_dropped := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_format, arg_processed, arg_dropped}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	processed = outArgs[1].Uint64()
	dropped = outArgs[2].Uint64()
	return
}

// gst_message_parse_qos_values
//
// [ jitter ] trans: everything, dir: out
//
// [ proportion ] trans: everything, dir: out
//
// [ quality ] trans: everything, dir: out
//
func (v Message) ParseQosValues() (jitter int64, proportion float64, quality int32) {
	iv, err := _I.Get(576, "Message", "parse_qos_values")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_jitter := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_proportion := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_quality := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_jitter, arg_proportion, arg_quality}
	iv.Call(args, nil, &outArgs[0])
	jitter = outArgs[0].Int64()
	proportion = outArgs[1].Double()
	quality = outArgs[2].Int32()
	return
}

// gst_message_parse_redirect_entry
//
// [ entry_index ] trans: nothing
//
// [ location ] trans: nothing, dir: out
//
// [ tag_list ] trans: nothing, dir: out
//
// [ entry_struct ] trans: nothing, dir: out
//
func (v Message) ParseRedirectEntry(entry_index uint64) (location string, tag_list TagList, entry_struct Structure) {
	iv, err := _I.Get(577, "Message", "parse_redirect_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_entry_index := gi.NewUint64Argument(entry_index)
	arg_location := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_tag_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_entry_struct := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_entry_index, arg_location, arg_tag_list, arg_entry_struct}
	iv.Call(args, nil, &outArgs[0])
	location = outArgs[0].String().Copy()
	tag_list.P = outArgs[1].Pointer()
	entry_struct.P = outArgs[2].Pointer()
	return
}

// gst_message_parse_request_state
//
// [ state ] trans: everything, dir: out
//
func (v Message) ParseRequestState() (state StateEnum) {
	iv, err := _I.Get(578, "Message", "parse_request_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_state}
	iv.Call(args, nil, &outArgs[0])
	state = StateEnum(outArgs[0].Int())
	return
}

// gst_message_parse_reset_time
//
// [ running_time ] trans: everything, dir: out
//
func (v Message) ParseResetTime() (running_time uint64) {
	iv, err := _I.Get(579, "Message", "parse_reset_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_running_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_running_time}
	iv.Call(args, nil, &outArgs[0])
	running_time = outArgs[0].Uint64()
	return
}

// gst_message_parse_segment_done
//
// [ format ] trans: everything, dir: out
//
// [ position ] trans: everything, dir: out
//
func (v Message) ParseSegmentDone() (format FormatEnum, position int64) {
	iv, err := _I.Get(580, "Message", "parse_segment_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_position := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_format, arg_position}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	position = outArgs[1].Int64()
	return
}

// gst_message_parse_segment_start
//
// [ format ] trans: everything, dir: out
//
// [ position ] trans: everything, dir: out
//
func (v Message) ParseSegmentStart() (format FormatEnum, position int64) {
	iv, err := _I.Get(581, "Message", "parse_segment_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_position := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_format, arg_position}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	position = outArgs[1].Int64()
	return
}

// gst_message_parse_state_changed
//
// [ oldstate ] trans: everything, dir: out
//
// [ newstate ] trans: everything, dir: out
//
// [ pending ] trans: everything, dir: out
//
func (v Message) ParseStateChanged() (oldstate StateEnum, newstate StateEnum, pending StateEnum) {
	iv, err := _I.Get(582, "Message", "parse_state_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_oldstate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_newstate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_pending := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_oldstate, arg_newstate, arg_pending}
	iv.Call(args, nil, &outArgs[0])
	oldstate = StateEnum(outArgs[0].Int())
	newstate = StateEnum(outArgs[1].Int())
	pending = StateEnum(outArgs[2].Int())
	return
}

// gst_message_parse_step_done
//
// [ format ] trans: everything, dir: out
//
// [ amount ] trans: everything, dir: out
//
// [ rate ] trans: everything, dir: out
//
// [ flush ] trans: everything, dir: out
//
// [ intermediate ] trans: everything, dir: out
//
// [ duration ] trans: everything, dir: out
//
// [ eos ] trans: everything, dir: out
//
func (v Message) ParseStepDone() (format FormatEnum, amount uint64, rate float64, flush bool, intermediate bool, duration uint64, eos bool) {
	iv, err := _I.Get(583, "Message", "parse_step_done")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [7]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_amount := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_rate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_flush := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_intermediate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[5]))
	arg_eos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[6]))
	args := []gi.Argument{arg_v, arg_format, arg_amount, arg_rate, arg_flush, arg_intermediate, arg_duration, arg_eos}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	amount = outArgs[1].Uint64()
	rate = outArgs[2].Double()
	flush = outArgs[3].Bool()
	intermediate = outArgs[4].Bool()
	duration = outArgs[5].Uint64()
	eos = outArgs[6].Bool()
	return
}

// gst_message_parse_step_start
//
// [ active ] trans: everything, dir: out
//
// [ format ] trans: everything, dir: out
//
// [ amount ] trans: everything, dir: out
//
// [ rate ] trans: everything, dir: out
//
// [ flush ] trans: everything, dir: out
//
// [ intermediate ] trans: everything, dir: out
//
func (v Message) ParseStepStart() (active bool, format FormatEnum, amount uint64, rate float64, flush bool, intermediate bool) {
	iv, err := _I.Get(584, "Message", "parse_step_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [6]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_active := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_amount := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_rate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_flush := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	arg_intermediate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[5]))
	args := []gi.Argument{arg_v, arg_active, arg_format, arg_amount, arg_rate, arg_flush, arg_intermediate}
	iv.Call(args, nil, &outArgs[0])
	active = outArgs[0].Bool()
	format = FormatEnum(outArgs[1].Int())
	amount = outArgs[2].Uint64()
	rate = outArgs[3].Double()
	flush = outArgs[4].Bool()
	intermediate = outArgs[5].Bool()
	return
}

// gst_message_parse_stream_collection
//
// [ collection ] trans: everything, dir: out
//
func (v Message) ParseStreamCollection() (collection StreamCollection) {
	iv, err := _I.Get(585, "Message", "parse_stream_collection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_collection := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_collection}
	iv.Call(args, nil, &outArgs[0])
	collection.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_stream_status
//
// [ type1 ] trans: everything, dir: out
//
// [ owner ] trans: nothing, dir: out
//
func (v Message) ParseStreamStatus() (type1 StreamStatusTypeEnum, owner Element) {
	iv, err := _I.Get(586, "Message", "parse_stream_status")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_owner := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_type1, arg_owner}
	iv.Call(args, nil, &outArgs[0])
	type1 = StreamStatusTypeEnum(outArgs[0].Int())
	owner.P = outArgs[1].Pointer()
	return
}

// gst_message_parse_streams_selected
//
// [ collection ] trans: everything, dir: out
//
func (v Message) ParseStreamsSelected() (collection StreamCollection) {
	iv, err := _I.Get(587, "Message", "parse_streams_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_collection := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_collection}
	iv.Call(args, nil, &outArgs[0])
	collection.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_structure_change
//
// [ type1 ] trans: everything, dir: out
//
// [ owner ] trans: nothing, dir: out
//
// [ busy ] trans: everything, dir: out
//
func (v Message) ParseStructureChange() (type1 StructureChangeTypeEnum, owner Element, busy bool) {
	iv, err := _I.Get(588, "Message", "parse_structure_change")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_owner := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_busy := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_type1, arg_owner, arg_busy}
	iv.Call(args, nil, &outArgs[0])
	type1 = StructureChangeTypeEnum(outArgs[0].Int())
	owner.P = outArgs[1].Pointer()
	busy = outArgs[2].Bool()
	return
}

// gst_message_parse_tag
//
// [ tag_list ] trans: everything, dir: out
//
func (v Message) ParseTag() (tag_list TagList) {
	iv, err := _I.Get(589, "Message", "parse_tag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag_list}
	iv.Call(args, nil, &outArgs[0])
	tag_list.P = outArgs[0].Pointer()
	return
}

// gst_message_parse_toc
//
// [ toc ] trans: everything, dir: out
//
// [ updated ] trans: everything, dir: out
//
func (v Message) ParseToc() (toc Toc, updated bool) {
	iv, err := _I.Get(590, "Message", "parse_toc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_toc := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_updated := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_toc, arg_updated}
	iv.Call(args, nil, &outArgs[0])
	toc.P = outArgs[0].Pointer()
	updated = outArgs[1].Bool()
	return
}

// gst_message_parse_warning
//
// [ gerror ] trans: everything, dir: out
//
// [ debug ] trans: everything, dir: out
//
func (v Message) ParseWarning() (gerror g.Error, debug string) {
	iv, err := _I.Get(591, "Message", "parse_warning")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_gerror := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_debug := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_gerror, arg_debug}
	iv.Call(args, nil, &outArgs[0])
	gerror.P = outArgs[0].Pointer()
	debug = outArgs[1].String().Take()
	return
}

// gst_message_parse_warning_details
//
// [ structure ] trans: nothing, dir: out
//
func (v Message) ParseWarningDetails() (structure Structure) {
	iv, err := _I.Get(592, "Message", "parse_warning_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_structure}
	iv.Call(args, nil, &outArgs[0])
	structure.P = outArgs[0].Pointer()
	return
}

// gst_message_set_buffering_stats
//
// [ mode ] trans: nothing
//
// [ avg_in ] trans: nothing
//
// [ avg_out ] trans: nothing
//
// [ buffering_left ] trans: nothing
//
func (v Message) SetBufferingStats(mode BufferingModeEnum, avg_in int32, avg_out int32, buffering_left int64) {
	iv, err := _I.Get(593, "Message", "set_buffering_stats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_avg_in := gi.NewInt32Argument(avg_in)
	arg_avg_out := gi.NewInt32Argument(avg_out)
	arg_buffering_left := gi.NewInt64Argument(buffering_left)
	args := []gi.Argument{arg_v, arg_mode, arg_avg_in, arg_avg_out, arg_buffering_left}
	iv.Call(args, nil, nil)
}

// gst_message_set_group_id
//
// [ group_id ] trans: nothing
//
func (v Message) SetGroupId(group_id uint32) {
	iv, err := _I.Get(594, "Message", "set_group_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_id := gi.NewUint32Argument(group_id)
	args := []gi.Argument{arg_v, arg_group_id}
	iv.Call(args, nil, nil)
}

// gst_message_set_qos_stats
//
// [ format ] trans: nothing
//
// [ processed ] trans: nothing
//
// [ dropped ] trans: nothing
//
func (v Message) SetQosStats(format FormatEnum, processed uint64, dropped uint64) {
	iv, err := _I.Get(595, "Message", "set_qos_stats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_processed := gi.NewUint64Argument(processed)
	arg_dropped := gi.NewUint64Argument(dropped)
	args := []gi.Argument{arg_v, arg_format, arg_processed, arg_dropped}
	iv.Call(args, nil, nil)
}

// gst_message_set_qos_values
//
// [ jitter ] trans: nothing
//
// [ proportion ] trans: nothing
//
// [ quality ] trans: nothing
//
func (v Message) SetQosValues(jitter int64, proportion float64, quality int32) {
	iv, err := _I.Get(596, "Message", "set_qos_values")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_jitter := gi.NewInt64Argument(jitter)
	arg_proportion := gi.NewDoubleArgument(proportion)
	arg_quality := gi.NewInt32Argument(quality)
	args := []gi.Argument{arg_v, arg_jitter, arg_proportion, arg_quality}
	iv.Call(args, nil, nil)
}

// gst_message_set_seqnum
//
// [ seqnum ] trans: nothing
//
func (v Message) SetSeqnum(seqnum uint32) {
	iv, err := _I.Get(597, "Message", "set_seqnum")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seqnum := gi.NewUint32Argument(seqnum)
	args := []gi.Argument{arg_v, arg_seqnum}
	iv.Call(args, nil, nil)
}

// gst_message_set_stream_status_object
//
// [ object ] trans: nothing
//
func (v Message) SetStreamStatusObject(object g.Value) {
	iv, err := _I.Get(598, "Message", "set_stream_status_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_object := gi.NewPointerArgument(object.P)
	args := []gi.Argument{arg_v, arg_object}
	iv.Call(args, nil, nil)
}

// gst_message_streams_selected_add
//
// [ stream ] trans: nothing
//
func (v Message) StreamsSelectedAdd(stream IStream) {
	iv, err := _I.Get(599, "Message", "streams_selected_add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if stream != nil {
		tmp = stream.P_Stream()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_stream}
	iv.Call(args, nil, nil)
}

// gst_message_streams_selected_get_size
//
// [ result ] trans: nothing
//
func (v Message) StreamsSelectedGetSize() (result uint32) {
	iv, err := _I.Get(600, "Message", "streams_selected_get_size")
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

// gst_message_streams_selected_get_stream
//
// [ idx ] trans: nothing
//
// [ result ] trans: everything
//
func (v Message) StreamsSelectedGetStream(idx uint32) (result Stream) {
	iv, err := _I.Get(601, "Message", "streams_selected_get_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_message_writable_structure
//
// [ result ] trans: nothing
//
func (v Message) WritableStructure() (result Structure) {
	iv, err := _I.Get(602, "Message", "writable_structure")
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

// Flags MessageType
type MessageTypeFlags int

const (
	MessageTypeUnknown          MessageTypeFlags = 0
	MessageTypeEos              MessageTypeFlags = 1
	MessageTypeError            MessageTypeFlags = 2
	MessageTypeWarning          MessageTypeFlags = 4
	MessageTypeInfo             MessageTypeFlags = 8
	MessageTypeTag              MessageTypeFlags = 16
	MessageTypeBuffering        MessageTypeFlags = 32
	MessageTypeStateChanged     MessageTypeFlags = 64
	MessageTypeStateDirty       MessageTypeFlags = 128
	MessageTypeStepDone         MessageTypeFlags = 256
	MessageTypeClockProvide     MessageTypeFlags = 512
	MessageTypeClockLost        MessageTypeFlags = 1024
	MessageTypeNewClock         MessageTypeFlags = 2048
	MessageTypeStructureChange  MessageTypeFlags = 4096
	MessageTypeStreamStatus     MessageTypeFlags = 8192
	MessageTypeApplication      MessageTypeFlags = 16384
	MessageTypeElement          MessageTypeFlags = 32768
	MessageTypeSegmentStart     MessageTypeFlags = 65536
	MessageTypeSegmentDone      MessageTypeFlags = 131072
	MessageTypeDurationChanged  MessageTypeFlags = 262144
	MessageTypeLatency          MessageTypeFlags = 524288
	MessageTypeAsyncStart       MessageTypeFlags = 1048576
	MessageTypeAsyncDone        MessageTypeFlags = 2097152
	MessageTypeRequestState     MessageTypeFlags = 4194304
	MessageTypeStepStart        MessageTypeFlags = 8388608
	MessageTypeQos              MessageTypeFlags = 16777216
	MessageTypeProgress         MessageTypeFlags = 33554432
	MessageTypeToc              MessageTypeFlags = 67108864
	MessageTypeResetTime        MessageTypeFlags = 134217728
	MessageTypeStreamStart      MessageTypeFlags = 268435456
	MessageTypeNeedContext      MessageTypeFlags = 536870912
	MessageTypeHaveContext      MessageTypeFlags = 1073741824
	MessageTypeExtended         MessageTypeFlags = 2147483648
	MessageTypeDeviceAdded      MessageTypeFlags = 2147483649
	MessageTypeDeviceRemoved    MessageTypeFlags = 2147483650
	MessageTypePropertyNotify   MessageTypeFlags = 2147483651
	MessageTypeStreamCollection MessageTypeFlags = 2147483652
	MessageTypeStreamsSelected  MessageTypeFlags = 2147483653
	MessageTypeRedirect         MessageTypeFlags = 2147483654
	MessageTypeAny              MessageTypeFlags = 4294967295
)

func MessageTypeGetType() gi.GType {
	ret := _I.GetGType(83, "MessageType")
	return ret
}

// Struct Meta
type Meta struct {
	P unsafe.Pointer
}

const SizeOfStructMeta = 16

func MetaGetType() gi.GType {
	ret := _I.GetGType(84, "Meta")
	return ret
}

// gst_meta_api_type_get_tags
//
// [ api ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaApiTypeGetTags1(api gi.GType) (result gi.CStrArray) {
	iv, err := _I.Get(603, "Meta", "api_type_get_tags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_api := gi.NewUintArgument(uint(api))
	args := []gi.Argument{arg_api}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gst_meta_api_type_has_tag
//
// [ api ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaApiTypeHasTag1(api gi.GType, tag uint32) (result bool) {
	iv, err := _I.Get(604, "Meta", "api_type_has_tag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_api := gi.NewUintArgument(uint(api))
	arg_tag := gi.NewUint32Argument(tag)
	args := []gi.Argument{arg_api, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_meta_api_type_register
//
// [ api ] trans: nothing
//
// [ tags ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaApiTypeRegister1(api string, tags gi.CStrArray) (result gi.GType) {
	iv, err := _I.Get(605, "Meta", "api_type_register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_api := gi.CString(api)
	arg_api := gi.NewStringArgument(c_api)
	arg_tags := gi.NewPointerArgument(tags.P)
	args := []gi.Argument{arg_api, arg_tags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_api)
	result = gi.GType(ret.Uint())
	return
}

// gst_meta_get_info
//
// [ impl ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaGetInfo1(impl string) (result MetaInfo) {
	iv, err := _I.Get(606, "Meta", "get_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_impl := gi.CString(impl)
	arg_impl := gi.NewStringArgument(c_impl)
	args := []gi.Argument{arg_impl}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_impl)
	result.P = ret.Pointer()
	return
}

// gst_meta_register
//
// [ api ] trans: nothing
//
// [ impl ] trans: nothing
//
// [ size ] trans: nothing
//
// [ init_func ] trans: nothing
//
// [ free_func ] trans: nothing
//
// [ transform_func ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaRegister1(api gi.GType, impl string, size uint64, init_func int /*TODO_TYPE CALLBACK*/, free_func int /*TODO_TYPE CALLBACK*/, transform_func int /*TODO_TYPE CALLBACK*/) (result MetaInfo) {
	iv, err := _I.Get(607, "Meta", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_impl := gi.CString(impl)
	arg_api := gi.NewUintArgument(uint(api))
	arg_impl := gi.NewStringArgument(c_impl)
	arg_size := gi.NewUint64Argument(size)
	arg_init_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMetaInitFunction()))
	arg_free_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMetaFreeFunction()))
	arg_transform_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMetaTransformFunction()))
	args := []gi.Argument{arg_api, arg_impl, arg_size, arg_init_func, arg_free_func, arg_transform_func}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_impl)
	result.P = ret.Pointer()
	return
}

// Flags MetaFlags
type MetaFlags int

const (
	MetaFlagsNone     MetaFlags = 0
	MetaFlagsReadonly MetaFlags = 1
	MetaFlagsPooled   MetaFlags = 2
	MetaFlagsLocked   MetaFlags = 4
	MetaFlagsLast     MetaFlags = 65536
)

func MetaFlagsGetType() gi.GType {
	ret := _I.GetGType(85, "MetaFlags")
	return ret
}

type MetaFreeFunctionStruct struct {
	F_meta   Meta
	F_buffer Buffer
}

func GetPointer_myMetaFreeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMetaFreeFunction())
}

//export myGstMetaFreeFunction
func myGstMetaFreeFunction(meta *C.GstMeta, buffer *C.GstBuffer) {
	// TODO: not found user_data
}

// Struct MetaInfo
type MetaInfo struct {
	P unsafe.Pointer
}

const SizeOfStructMetaInfo = 48

func MetaInfoGetType() gi.GType {
	ret := _I.GetGType(86, "MetaInfo")
	return ret
}

type MetaInitFunctionStruct struct {
	F_meta   Meta
	F_params unsafe.Pointer
	F_buffer Buffer
}

func GetPointer_myMetaInitFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMetaInitFunction())
}

//export myGstMetaInitFunction
func myGstMetaInitFunction(meta *C.GstMeta, params C.gpointer, buffer *C.GstBuffer) {
	// TODO: not found user_data
}

// Struct MetaTransformCopy
type MetaTransformCopy struct {
	P unsafe.Pointer
}

const SizeOfStructMetaTransformCopy = 24

func MetaTransformCopyGetType() gi.GType {
	ret := _I.GetGType(87, "MetaTransformCopy")
	return ret
}

type MetaTransformFunctionStruct struct {
	F_transbuf Buffer
	F_meta     Meta
	F_buffer   Buffer
	F_type1    uint32
	F_data     unsafe.Pointer
}

func GetPointer_myMetaTransformFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMetaTransformFunction())
}

//export myGstMetaTransformFunction
func myGstMetaTransformFunction(transbuf *C.GstBuffer, meta *C.GstMeta, buffer *C.GstBuffer, type1 C.guint32, data C.gpointer) {
	// TODO: not found user_data
}

// Struct MiniObject
type MiniObject struct {
	P unsafe.Pointer
}

const SizeOfStructMiniObject = 64

func MiniObjectGetType() gi.GType {
	ret := _I.GetGType(88, "MiniObject")
	return ret
}

// gst_mini_object_get_qdata
//
// [ quark ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MiniObject) GetQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(608, "MiniObject", "get_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_v, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// gst_mini_object_is_writable
//
// [ result ] trans: nothing
//
func (v MiniObject) IsWritable() (result bool) {
	iv, err := _I.Get(609, "MiniObject", "is_writable")
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

// gst_mini_object_lock
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MiniObject) Lock(flags LockFlags) (result bool) {
	iv, err := _I.Get(610, "MiniObject", "lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_mini_object_set_qdata
//
// [ quark ] trans: nothing
//
// [ data ] trans: nothing
//
// [ destroy ] trans: nothing
//
func (v MiniObject) SetQdata(quark uint32, data unsafe.Pointer, destroy int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(611, "MiniObject", "set_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	arg_data := gi.NewPointerArgument(data)
	arg_destroy := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_quark, arg_data, arg_destroy}
	iv.Call(args, nil, nil)
}

// gst_mini_object_steal_qdata
//
// [ quark ] trans: nothing
//
// [ result ] trans: everything
//
func (v MiniObject) StealQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(612, "MiniObject", "steal_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_v, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// gst_mini_object_unlock
//
// [ flags ] trans: nothing
//
func (v MiniObject) Unlock(flags LockFlags) {
	iv, err := _I.Get(613, "MiniObject", "unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, nil)
}

// gst_mini_object_replace
//
// [ newdata ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MiniObject) Replace(newdata MiniObject) (result bool) {
	iv, err := _I.Get(614, "MiniObject", "replace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_newdata := gi.NewPointerArgument(newdata.P)
	args := []gi.Argument{arg_v, arg_newdata}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_mini_object_take
//
// [ newdata ] trans: nothing
//
// [ result ] trans: nothing
//
func (v MiniObject) Take(newdata MiniObject) (result bool) {
	iv, err := _I.Get(615, "MiniObject", "take")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_newdata := gi.NewPointerArgument(newdata.P)
	args := []gi.Argument{arg_v, arg_newdata}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

type MiniObjectDisposeFunctionStruct struct {
	F_obj MiniObject
}

func GetPointer_myMiniObjectDisposeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMiniObjectDisposeFunction())
}

//export myGstMiniObjectDisposeFunction
func myGstMiniObjectDisposeFunction(obj *C.GstMiniObject) {
	// TODO: not found user_data
}

// Flags MiniObjectFlags
type MiniObjectFlags int

const (
	MiniObjectFlagsLockable     MiniObjectFlags = 1
	MiniObjectFlagsLockReadonly MiniObjectFlags = 2
	MiniObjectFlagsMayBeLeaked  MiniObjectFlags = 4
	MiniObjectFlagsLast         MiniObjectFlags = 16
)

func MiniObjectFlagsGetType() gi.GType {
	ret := _I.GetGType(89, "MiniObjectFlags")
	return ret
}

type MiniObjectFreeFunctionStruct struct {
	F_obj MiniObject
}

func GetPointer_myMiniObjectFreeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMiniObjectFreeFunction())
}

//export myGstMiniObjectFreeFunction
func myGstMiniObjectFreeFunction(obj *C.GstMiniObject) {
	// TODO: not found user_data
}

type MiniObjectNotifyStruct struct {
	F_obj MiniObject
}

func GetPointer_myMiniObjectNotify() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstMiniObjectNotify())
}

//export myGstMiniObjectNotify
func myGstMiniObjectNotify(user_data C.gpointer, obj *C.GstMiniObject) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &MiniObjectNotifyStruct{
		F_obj: MiniObject{P: unsafe.Pointer(obj)},
	}
	fn(args)
}

// Object Object
type Object struct {
	g.InitiallyUnowned
}

func WrapObject(p unsafe.Pointer) (r Object) { r.P = p; return }

type IObject interface{ P_Object() unsafe.Pointer }

func (v Object) P_Object() unsafe.Pointer { return v.P }
func ObjectGetType() gi.GType {
	ret := _I.GetGType(90, "Object")
	return ret
}

// gst_object_check_uniqueness
//
// [ list ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectCheckUniqueness1(list g.List, name string) (result bool) {
	iv, err := _I.Get(616, "Object", "check_uniqueness")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_list := gi.NewPointerArgument(list.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_list, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_object_default_deep_notify
//
// [ orig ] trans: nothing
//
// [ pspec ] trans: nothing
//
// [ excluded_props ] trans: nothing
//
func (v Object) DefaultDeepNotify(orig IObject, pspec g.IParamSpec, excluded_props gi.CStrArray) {
	iv, err := _I.Get(617, "Object", "default_deep_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if orig != nil {
		tmp = orig.P_Object()
	}
	var tmp1 unsafe.Pointer
	if pspec != nil {
		tmp1 = pspec.P_ParamSpec()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_orig := gi.NewPointerArgument(tmp)
	arg_pspec := gi.NewPointerArgument(tmp1)
	arg_excluded_props := gi.NewPointerArgument(excluded_props.P)
	args := []gi.Argument{arg_v, arg_orig, arg_pspec, arg_excluded_props}
	iv.Call(args, nil, nil)
}

// gst_object_replace
//
// [ newobj ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) Replace(newobj IObject) (result bool) {
	iv, err := _I.Get(618, "Object", "replace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if newobj != nil {
		tmp = newobj.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_newobj := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_newobj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_add_control_binding
//
// [ binding ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) AddControlBinding(binding IControlBinding) (result bool) {
	iv, err := _I.Get(619, "Object", "add_control_binding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if binding != nil {
		tmp = binding.P_ControlBinding()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_binding := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_binding}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_default_error
//
// [ error ] trans: nothing
//
// [ debug ] trans: nothing
//
func (v Object) DefaultError(error g.Error, debug string) {
	iv, err := _I.Get(620, "Object", "default_error")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_debug := gi.CString(debug)
	arg_v := gi.NewPointerArgument(v.P)
	arg_error := gi.NewPointerArgument(error.P)
	arg_debug := gi.NewStringArgument(c_debug)
	args := []gi.Argument{arg_v, arg_error, arg_debug}
	iv.Call(args, nil, nil)
	gi.Free(c_debug)
}

// gst_object_get_control_binding
//
// [ property_name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Object) GetControlBinding(property_name string) (result ControlBinding) {
	iv, err := _I.Get(621, "Object", "get_control_binding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	args := []gi.Argument{arg_v, arg_property_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// gst_object_get_control_rate
//
// [ result ] trans: nothing
//
func (v Object) GetControlRate() (result uint64) {
	iv, err := _I.Get(622, "Object", "get_control_rate")
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

// gst_object_get_g_value_array
//
// [ property_name ] trans: nothing
//
// [ timestamp ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ n_values ] trans: nothing
//
// [ values ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) GetGValueArray(property_name string, timestamp uint64, interval uint64, n_values uint32, values unsafe.Pointer) (result bool) {
	iv, err := _I.Get(623, "Object", "get_g_value_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	arg_interval := gi.NewUint64Argument(interval)
	arg_n_values := gi.NewUint32Argument(n_values)
	arg_values := gi.NewPointerArgument(values)
	args := []gi.Argument{arg_v, arg_property_name, arg_timestamp, arg_interval, arg_n_values, arg_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result = ret.Bool()
	return
}

// gst_object_get_name
//
// [ result ] trans: everything
//
func (v Object) GetName() (result string) {
	iv, err := _I.Get(624, "Object", "get_name")
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

// gst_object_get_parent
//
// [ result ] trans: everything
//
func (v Object) GetParent() (result Object) {
	iv, err := _I.Get(625, "Object", "get_parent")
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

// gst_object_get_path_string
//
// [ result ] trans: everything
//
func (v Object) GetPathString() (result string) {
	iv, err := _I.Get(626, "Object", "get_path_string")
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

// gst_object_get_value
//
// [ property_name ] trans: nothing
//
// [ timestamp ] trans: nothing
//
// [ result ] trans: everything
//
func (v Object) GetValue(property_name string, timestamp uint64) (result g.Value) {
	iv, err := _I.Get(627, "Object", "get_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_timestamp := gi.NewUint64Argument(timestamp)
	args := []gi.Argument{arg_v, arg_property_name, arg_timestamp}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// gst_object_has_active_control_bindings
//
// [ result ] trans: nothing
//
func (v Object) HasActiveControlBindings() (result bool) {
	iv, err := _I.Get(628, "Object", "has_active_control_bindings")
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
// gst_object_has_ancestor
//
// [ ancestor ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) HasAncestor(ancestor IObject) (result bool) {
	iv, err := _I.Get(629, "Object", "has_ancestor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if ancestor != nil {
		tmp = ancestor.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ancestor := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_ancestor}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_has_as_ancestor
//
// [ ancestor ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) HasAsAncestor(ancestor IObject) (result bool) {
	iv, err := _I.Get(630, "Object", "has_as_ancestor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if ancestor != nil {
		tmp = ancestor.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ancestor := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_ancestor}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_has_as_parent
//
// [ parent ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) HasAsParent(parent IObject) (result bool) {
	iv, err := _I.Get(631, "Object", "has_as_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_parent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_ref
//
// [ result ] trans: everything
//
func (v Object) Ref() (result Object) {
	iv, err := _I.Get(632, "Object", "ref")
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

// gst_object_remove_control_binding
//
// [ binding ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) RemoveControlBinding(binding IControlBinding) (result bool) {
	iv, err := _I.Get(633, "Object", "remove_control_binding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if binding != nil {
		tmp = binding.P_ControlBinding()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_binding := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_binding}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_set_control_binding_disabled
//
// [ property_name ] trans: nothing
//
// [ disabled ] trans: nothing
//
func (v Object) SetControlBindingDisabled(property_name string, disabled bool) {
	iv, err := _I.Get(634, "Object", "set_control_binding_disabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_disabled := gi.NewBoolArgument(disabled)
	args := []gi.Argument{arg_v, arg_property_name, arg_disabled}
	iv.Call(args, nil, nil)
	gi.Free(c_property_name)
}

// gst_object_set_control_bindings_disabled
//
// [ disabled ] trans: nothing
//
func (v Object) SetControlBindingsDisabled(disabled bool) {
	iv, err := _I.Get(635, "Object", "set_control_bindings_disabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_disabled := gi.NewBoolArgument(disabled)
	args := []gi.Argument{arg_v, arg_disabled}
	iv.Call(args, nil, nil)
}

// gst_object_set_control_rate
//
// [ control_rate ] trans: nothing
//
func (v Object) SetControlRate(control_rate uint64) {
	iv, err := _I.Get(636, "Object", "set_control_rate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_control_rate := gi.NewUint64Argument(control_rate)
	args := []gi.Argument{arg_v, arg_control_rate}
	iv.Call(args, nil, nil)
}

// gst_object_set_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) SetName(name string) (result bool) {
	iv, err := _I.Get(637, "Object", "set_name")
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
	result = ret.Bool()
	return
}

// gst_object_set_parent
//
// [ parent ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) SetParent(parent IObject) (result bool) {
	iv, err := _I.Get(638, "Object", "set_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_parent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_object_suggest_next_sync
//
// [ result ] trans: nothing
//
func (v Object) SuggestNextSync() (result uint64) {
	iv, err := _I.Get(639, "Object", "suggest_next_sync")
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

// gst_object_sync_values
//
// [ timestamp ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) SyncValues(timestamp uint64) (result bool) {
	iv, err := _I.Get(640, "Object", "sync_values")
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

// gst_object_unparent
//
func (v Object) Unparent() {
	iv, err := _I.Get(641, "Object", "unparent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_object_unref
//
func (v Object) Unref() {
	iv, err := _I.Get(642, "Object", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectClass

// Flags ObjectFlags
type ObjectFlags int

const (
	ObjectFlagsMayBeLeaked ObjectFlags = 1
	ObjectFlagsLast        ObjectFlags = 16
)

func ObjectFlagsGetType() gi.GType {
	ret := _I.GetGType(91, "ObjectFlags")
	return ret
}

// Object Pad
type Pad struct {
	Object
}

func WrapPad(p unsafe.Pointer) (r Pad) { r.P = p; return }

type IPad interface{ P_Pad() unsafe.Pointer }

func (v Pad) P_Pad() unsafe.Pointer { return v.P }
func PadGetType() gi.GType {
	ret := _I.GetGType(92, "Pad")
	return ret
}

// gst_pad_new
//
// [ name ] trans: nothing
//
// [ direction ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPad(name string, direction PadDirectionEnum) (result Pad) {
	iv, err := _I.Get(643, "Pad", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_direction := gi.NewIntArgument(int(direction))
	args := []gi.Argument{arg_name, arg_direction}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_pad_new_from_static_template
//
// [ templ ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPadFromStaticTemplate(templ StaticPadTemplate, name string) (result Pad) {
	iv, err := _I.Get(644, "Pad", "new_from_static_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_templ := gi.NewPointerArgument(templ.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_templ, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_pad_new_from_template
//
// [ templ ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPadFromTemplate(templ IPadTemplate, name string) (result Pad) {
	iv, err := _I.Get(645, "Pad", "new_from_template")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if templ != nil {
		tmp = templ.P_PadTemplate()
	}
	c_name := gi.CString(name)
	arg_templ := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_templ, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_pad_link_get_name
//
// [ ret ] trans: nothing
//
// [ result ] trans: nothing
//
func PadLinkGetName1(ret PadLinkReturnEnum) (result string) {
	iv, err := _I.Get(646, "Pad", "link_get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ret := gi.NewIntArgument(int(ret))
	args := []gi.Argument{arg_ret}
	var ret1 gi.Argument
	iv.Call(args, &ret1, nil)
	result = ret1.String().Copy()
	return
}

// gst_pad_activate_mode
//
// [ mode ] trans: nothing
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) ActivateMode(mode PadModeEnum, active bool) (result bool) {
	iv, err := _I.Get(647, "Pad", "activate_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_v, arg_mode, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_add_probe
//
// [ mask ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ destroy_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) AddProbe(mask PadProbeTypeFlags, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, destroy_data int /*TODO_TYPE CALLBACK*/) (result uint64) {
	iv, err := _I.Get(648, "Pad", "add_probe")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewIntArgument(int(mask))
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadProbeCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy_data := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_mask, arg_callback, arg_user_data, arg_destroy_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_pad_can_link
//
// [ sinkpad ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) CanLink(sinkpad IPad) (result bool) {
	iv, err := _I.Get(649, "Pad", "can_link")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sinkpad != nil {
		tmp = sinkpad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sinkpad := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_sinkpad}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_chain
//
// [ buffer ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) Chain(buffer Buffer) (result FlowReturnEnum) {
	iv, err := _I.Get(650, "Pad", "chain")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_chain_list
//
// [ list ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) ChainList(list BufferList) (result FlowReturnEnum) {
	iv, err := _I.Get(651, "Pad", "chain_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_v, arg_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_check_reconfigure
//
// [ result ] trans: nothing
//
func (v Pad) CheckReconfigure() (result bool) {
	iv, err := _I.Get(652, "Pad", "check_reconfigure")
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

// gst_pad_create_stream_id
//
// [ parent ] trans: nothing
//
// [ stream_id ] trans: nothing
//
// [ result ] trans: everything
//
func (v Pad) CreateStreamId(parent IElement, stream_id string) (result string) {
	iv, err := _I.Get(653, "Pad", "create_stream_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Element()
	}
	c_stream_id := gi.CString(stream_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	arg_stream_id := gi.NewStringArgument(c_stream_id)
	args := []gi.Argument{arg_v, arg_parent, arg_stream_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_stream_id)
	result = ret.String().Take()
	return
}

// gst_pad_event_default
//
// [ parent ] trans: nothing
//
// [ event ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) EventDefault(parent IObject, event Event) (result bool) {
	iv, err := _I.Get(654, "Pad", "event_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_parent, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_forward
//
// [ forward ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) Forward(forward int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(655, "Pad", "forward")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_forward := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadForwardFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_forward, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_get_allowed_caps
//
// [ result ] trans: everything
//
func (v Pad) GetAllowedCaps() (result Caps) {
	iv, err := _I.Get(656, "Pad", "get_allowed_caps")
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

// gst_pad_get_current_caps
//
// [ result ] trans: everything
//
func (v Pad) GetCurrentCaps() (result Caps) {
	iv, err := _I.Get(657, "Pad", "get_current_caps")
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

// gst_pad_get_direction
//
// [ result ] trans: nothing
//
func (v Pad) GetDirection() (result PadDirectionEnum) {
	iv, err := _I.Get(658, "Pad", "get_direction")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PadDirectionEnum(ret.Int())
	return
}

// gst_pad_get_element_private
//
// [ result ] trans: nothing
//
func (v Pad) GetElementPrivate() (result unsafe.Pointer) {
	iv, err := _I.Get(659, "Pad", "get_element_private")
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

// gst_pad_get_last_flow_return
//
// [ result ] trans: nothing
//
func (v Pad) GetLastFlowReturn() (result FlowReturnEnum) {
	iv, err := _I.Get(660, "Pad", "get_last_flow_return")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_get_offset
//
// [ result ] trans: nothing
//
func (v Pad) GetOffset() (result int64) {
	iv, err := _I.Get(661, "Pad", "get_offset")
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

// gst_pad_get_pad_template
//
// [ result ] trans: everything
//
func (v Pad) GetPadTemplate() (result PadTemplate) {
	iv, err := _I.Get(662, "Pad", "get_pad_template")
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

// gst_pad_get_pad_template_caps
//
// [ result ] trans: everything
//
func (v Pad) GetPadTemplateCaps() (result Caps) {
	iv, err := _I.Get(663, "Pad", "get_pad_template_caps")
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

// gst_pad_get_parent_element
//
// [ result ] trans: everything
//
func (v Pad) GetParentElement() (result Element) {
	iv, err := _I.Get(664, "Pad", "get_parent_element")
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

// gst_pad_get_peer
//
// [ result ] trans: everything
//
func (v Pad) GetPeer() (result Pad) {
	iv, err := _I.Get(665, "Pad", "get_peer")
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

// gst_pad_get_range
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ buffer ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) GetRange(offset uint64, size uint32) (result FlowReturnEnum, buffer Buffer) {
	iv, err := _I.Get(666, "Pad", "get_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint32Argument(size)
	arg_buffer := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_size, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	buffer.P = outArgs[0].Pointer()
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_get_sticky_event
//
// [ event_type ] trans: nothing
//
// [ idx ] trans: nothing
//
// [ result ] trans: everything
//
func (v Pad) GetStickyEvent(event_type EventTypeEnum, idx uint32) (result Event) {
	iv, err := _I.Get(667, "Pad", "get_sticky_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event_type := gi.NewIntArgument(int(event_type))
	arg_idx := gi.NewUint32Argument(idx)
	args := []gi.Argument{arg_v, arg_event_type, arg_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_pad_get_stream
//
// [ result ] trans: everything
//
func (v Pad) GetStream() (result Stream) {
	iv, err := _I.Get(668, "Pad", "get_stream")
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

// gst_pad_get_stream_id
//
// [ result ] trans: everything
//
func (v Pad) GetStreamId() (result string) {
	iv, err := _I.Get(669, "Pad", "get_stream_id")
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

// gst_pad_get_task_state
//
// [ result ] trans: nothing
//
func (v Pad) GetTaskState() (result TaskStateEnum) {
	iv, err := _I.Get(670, "Pad", "get_task_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TaskStateEnum(ret.Int())
	return
}

// gst_pad_has_current_caps
//
// [ result ] trans: nothing
//
func (v Pad) HasCurrentCaps() (result bool) {
	iv, err := _I.Get(671, "Pad", "has_current_caps")
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

// gst_pad_is_active
//
// [ result ] trans: nothing
//
func (v Pad) IsActive() (result bool) {
	iv, err := _I.Get(672, "Pad", "is_active")
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

// gst_pad_is_blocked
//
// [ result ] trans: nothing
//
func (v Pad) IsBlocked() (result bool) {
	iv, err := _I.Get(673, "Pad", "is_blocked")
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

// gst_pad_is_blocking
//
// [ result ] trans: nothing
//
func (v Pad) IsBlocking() (result bool) {
	iv, err := _I.Get(674, "Pad", "is_blocking")
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

// gst_pad_is_linked
//
// [ result ] trans: nothing
//
func (v Pad) IsLinked() (result bool) {
	iv, err := _I.Get(675, "Pad", "is_linked")
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

// gst_pad_iterate_internal_links
//
// [ result ] trans: everything
//
func (v Pad) IterateInternalLinks() (result Iterator) {
	iv, err := _I.Get(676, "Pad", "iterate_internal_links")
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

// gst_pad_iterate_internal_links_default
//
// [ parent ] trans: nothing
//
// [ result ] trans: everything
//
func (v Pad) IterateInternalLinksDefault(parent IObject) (result Iterator) {
	iv, err := _I.Get(677, "Pad", "iterate_internal_links_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_parent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_pad_link
//
// [ sinkpad ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) Link(sinkpad IPad) (result PadLinkReturnEnum) {
	iv, err := _I.Get(678, "Pad", "link")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sinkpad != nil {
		tmp = sinkpad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sinkpad := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_sinkpad}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PadLinkReturnEnum(ret.Int())
	return
}

// gst_pad_link_full
//
// [ sinkpad ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) LinkFull(sinkpad IPad, flags PadLinkCheckFlags) (result PadLinkReturnEnum) {
	iv, err := _I.Get(679, "Pad", "link_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sinkpad != nil {
		tmp = sinkpad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sinkpad := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_sinkpad, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PadLinkReturnEnum(ret.Int())
	return
}

// gst_pad_link_maybe_ghosting
//
// [ sink ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) LinkMaybeGhosting(sink IPad) (result bool) {
	iv, err := _I.Get(680, "Pad", "link_maybe_ghosting")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sink != nil {
		tmp = sink.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sink := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_sink}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_link_maybe_ghosting_full
//
// [ sink ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) LinkMaybeGhostingFull(sink IPad, flags PadLinkCheckFlags) (result bool) {
	iv, err := _I.Get(681, "Pad", "link_maybe_ghosting_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sink != nil {
		tmp = sink.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sink := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_sink, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_mark_reconfigure
//
func (v Pad) MarkReconfigure() {
	iv, err := _I.Get(682, "Pad", "mark_reconfigure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_pad_needs_reconfigure
//
// [ result ] trans: nothing
//
func (v Pad) NeedsReconfigure() (result bool) {
	iv, err := _I.Get(683, "Pad", "needs_reconfigure")
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

// gst_pad_pause_task
//
// [ result ] trans: nothing
//
func (v Pad) PauseTask() (result bool) {
	iv, err := _I.Get(684, "Pad", "pause_task")
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

// gst_pad_peer_query
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) PeerQuery(query Query) (result bool) {
	iv, err := _I.Get(685, "Pad", "peer_query")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_v, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_peer_query_accept_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) PeerQueryAcceptCaps(caps Caps) (result bool) {
	iv, err := _I.Get(686, "Pad", "peer_query_accept_caps")
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

// gst_pad_peer_query_caps
//
// [ filter ] trans: nothing
//
// [ result ] trans: everything
//
func (v Pad) PeerQueryCaps(filter Caps) (result Caps) {
	iv, err := _I.Get(687, "Pad", "peer_query_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_filter := gi.NewPointerArgument(filter.P)
	args := []gi.Argument{arg_v, arg_filter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_pad_peer_query_convert
//
// [ src_format ] trans: nothing
//
// [ src_val ] trans: nothing
//
// [ dest_format ] trans: nothing
//
// [ dest_val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) PeerQueryConvert(src_format FormatEnum, src_val int64, dest_format FormatEnum) (result bool, dest_val int64) {
	iv, err := _I.Get(688, "Pad", "peer_query_convert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_format := gi.NewIntArgument(int(src_format))
	arg_src_val := gi.NewInt64Argument(src_val)
	arg_dest_format := gi.NewIntArgument(int(dest_format))
	arg_dest_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_src_format, arg_src_val, arg_dest_format, arg_dest_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	dest_val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_pad_peer_query_duration
//
// [ format ] trans: nothing
//
// [ duration ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) PeerQueryDuration(format FormatEnum) (result bool, duration int64) {
	iv, err := _I.Get(689, "Pad", "peer_query_duration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_duration}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	duration = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_pad_peer_query_position
//
// [ format ] trans: nothing
//
// [ cur ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) PeerQueryPosition(format FormatEnum) (result bool, cur int64) {
	iv, err := _I.Get(690, "Pad", "peer_query_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_cur := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_cur}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	cur = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_pad_proxy_query_accept_caps
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) ProxyQueryAcceptCaps(query Query) (result bool) {
	iv, err := _I.Get(691, "Pad", "proxy_query_accept_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_v, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_proxy_query_caps
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) ProxyQueryCaps(query Query) (result bool) {
	iv, err := _I.Get(692, "Pad", "proxy_query_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_v, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_pull_range
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ buffer ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) PullRange(offset uint64, size uint32) (result FlowReturnEnum, buffer Buffer) {
	iv, err := _I.Get(693, "Pad", "pull_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint32Argument(size)
	arg_buffer := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_size, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	buffer.P = outArgs[0].Pointer()
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_push
//
// [ buffer ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) Push(buffer Buffer) (result FlowReturnEnum) {
	iv, err := _I.Get(694, "Pad", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_push_event
//
// [ event ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) PushEvent(event Event) (result bool) {
	iv, err := _I.Get(695, "Pad", "push_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_push_list
//
// [ list ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) PushList(list BufferList) (result FlowReturnEnum) {
	iv, err := _I.Get(696, "Pad", "push_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_v, arg_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_query
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) QueryF(query Query) (result bool) {
	iv, err := _I.Get(697, "Pad", "query")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_v, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_query_accept_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) QueryAcceptCaps(caps Caps) (result bool) {
	iv, err := _I.Get(698, "Pad", "query_accept_caps")
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

// gst_pad_query_caps
//
// [ filter ] trans: nothing
//
// [ result ] trans: everything
//
func (v Pad) QueryCaps(filter Caps) (result Caps) {
	iv, err := _I.Get(699, "Pad", "query_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_filter := gi.NewPointerArgument(filter.P)
	args := []gi.Argument{arg_v, arg_filter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_pad_query_convert
//
// [ src_format ] trans: nothing
//
// [ src_val ] trans: nothing
//
// [ dest_format ] trans: nothing
//
// [ dest_val ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) QueryConvert(src_format FormatEnum, src_val int64, dest_format FormatEnum) (result bool, dest_val int64) {
	iv, err := _I.Get(700, "Pad", "query_convert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_format := gi.NewIntArgument(int(src_format))
	arg_src_val := gi.NewInt64Argument(src_val)
	arg_dest_format := gi.NewIntArgument(int(dest_format))
	arg_dest_val := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_src_format, arg_src_val, arg_dest_format, arg_dest_val}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	dest_val = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_pad_query_default
//
// [ parent ] trans: nothing
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) QueryDefault(parent IObject, query Query) (result bool) {
	iv, err := _I.Get(701, "Pad", "query_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_v, arg_parent, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_query_duration
//
// [ format ] trans: nothing
//
// [ duration ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) QueryDuration(format FormatEnum) (result bool, duration int64) {
	iv, err := _I.Get(702, "Pad", "query_duration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_duration}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	duration = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_pad_query_position
//
// [ format ] trans: nothing
//
// [ cur ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Pad) QueryPosition(format FormatEnum) (result bool, cur int64) {
	iv, err := _I.Get(703, "Pad", "query_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_cur := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_cur}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	cur = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_pad_remove_probe
//
// [ id ] trans: nothing
//
func (v Pad) RemoveProbe(id uint64) {
	iv, err := _I.Get(704, "Pad", "remove_probe")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewUint64Argument(id)
	args := []gi.Argument{arg_v, arg_id}
	iv.Call(args, nil, nil)
}

// gst_pad_send_event
//
// [ event ] trans: everything
//
// [ result ] trans: nothing
//
func (v Pad) SendEvent(event Event) (result bool) {
	iv, err := _I.Get(705, "Pad", "send_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_set_activate_function_full
//
// [ activate ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetActivateFunctionFull(activate int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(706, "Pad", "set_activate_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_activate := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadActivateFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_activate, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_activatemode_function_full
//
// [ activatemode ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetActivatemodeFunctionFull(activatemode int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(707, "Pad", "set_activatemode_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_activatemode := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadActivateModeFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_activatemode, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_active
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) SetActive(active bool) (result bool) {
	iv, err := _I.Get(708, "Pad", "set_active")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_v, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_set_chain_function_full
//
// [ chain ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetChainFunctionFull(chain int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(709, "Pad", "set_chain_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_chain := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadChainFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_chain, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_chain_list_function_full
//
// [ chainlist ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetChainListFunctionFull(chainlist int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(710, "Pad", "set_chain_list_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_chainlist := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadChainListFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_chainlist, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_element_private
//
// [ priv ] trans: nothing
//
func (v Pad) SetElementPrivate(priv unsafe.Pointer) {
	iv, err := _I.Get(711, "Pad", "set_element_private")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priv := gi.NewPointerArgument(priv)
	args := []gi.Argument{arg_v, arg_priv}
	iv.Call(args, nil, nil)
}

// gst_pad_set_event_full_function_full
//
// [ event ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetEventFullFunctionFull(event int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(712, "Pad", "set_event_full_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadEventFullFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_event, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_event_function_full
//
// [ event ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetEventFunctionFull(event int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(713, "Pad", "set_event_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadEventFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_event, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_getrange_function_full
//
// [ get ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetGetrangeFunctionFull(get int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(714, "Pad", "set_getrange_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_get := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadGetRangeFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_get, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_iterate_internal_links_function_full
//
// [ iterintlink ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetIterateInternalLinksFunctionFull(iterintlink int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(715, "Pad", "set_iterate_internal_links_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iterintlink := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadIterIntLinkFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_iterintlink, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_link_function_full
//
// [ link ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetLinkFunctionFull(link int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(716, "Pad", "set_link_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_link := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadLinkFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_link, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_offset
//
// [ offset ] trans: nothing
//
func (v Pad) SetOffset(offset int64) {
	iv, err := _I.Get(717, "Pad", "set_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	iv.Call(args, nil, nil)
}

// gst_pad_set_query_function_full
//
// [ query ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetQueryFunctionFull(query int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(718, "Pad", "set_query_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadQueryFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_query, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_set_unlink_function_full
//
// [ unlink ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Pad) SetUnlinkFunctionFull(unlink int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(719, "Pad", "set_unlink_function_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_unlink := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadUnlinkFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_unlink, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_pad_start_task
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) StartTask(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result bool) {
	iv, err := _I.Get(720, "Pad", "start_task")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTaskFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_func1, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_sticky_events_foreach
//
// [ foreach_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v Pad) StickyEventsForeach(foreach_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(721, "Pad", "sticky_events_foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_foreach_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPadStickyEventsForeachFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_foreach_func, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_pad_stop_task
//
// [ result ] trans: nothing
//
func (v Pad) StopTask() (result bool) {
	iv, err := _I.Get(722, "Pad", "stop_task")
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

// gst_pad_store_sticky_event
//
// [ event ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) StoreStickyEvent(event Event) (result FlowReturnEnum) {
	iv, err := _I.Get(723, "Pad", "store_sticky_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_pad_unlink
//
// [ sinkpad ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pad) Unlink(sinkpad IPad) (result bool) {
	iv, err := _I.Get(724, "Pad", "unlink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sinkpad != nil {
		tmp = sinkpad.P_Pad()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sinkpad := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_sinkpad}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_pad_use_fixed_caps
//
func (v Pad) UseFixedCaps() {
	iv, err := _I.Get(725, "Pad", "use_fixed_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

type PadActivateFunctionStruct struct {
	F_pad    Pad
	F_parent Object
}

func GetPointer_myPadActivateFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadActivateFunction())
}

//export myGstPadActivateFunction
func myGstPadActivateFunction(pad *C.GstPad, parent *C.GstObject) {
	// TODO: not found user_data
}

type PadActivateModeFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_mode   PadModeEnum
	F_active bool
}

func GetPointer_myPadActivateModeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadActivateModeFunction())
}

//export myGstPadActivateModeFunction
func myGstPadActivateModeFunction(pad *C.GstPad, parent *C.GstObject, mode C.GstPadMode, active C.gboolean) {
	// TODO: not found user_data
}

type PadChainFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_buffer Buffer
}

func GetPointer_myPadChainFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadChainFunction())
}

//export myGstPadChainFunction
func myGstPadChainFunction(pad *C.GstPad, parent *C.GstObject, buffer *C.GstBuffer) {
	// TODO: not found user_data
}

type PadChainListFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_list   BufferList
}

func GetPointer_myPadChainListFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadChainListFunction())
}

//export myGstPadChainListFunction
func myGstPadChainListFunction(pad *C.GstPad, parent *C.GstObject, list *C.GstBufferList) {
	// TODO: not found user_data
}

// ignore GType struct PadClass

// Enum PadDirection
type PadDirectionEnum int

const (
	PadDirectionUnknown PadDirectionEnum = 0
	PadDirectionSrc     PadDirectionEnum = 1
	PadDirectionSink    PadDirectionEnum = 2
)

func PadDirectionGetType() gi.GType {
	ret := _I.GetGType(93, "PadDirection")
	return ret
}

type PadEventFullFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_event  Event
}

func GetPointer_myPadEventFullFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadEventFullFunction())
}

//export myGstPadEventFullFunction
func myGstPadEventFullFunction(pad *C.GstPad, parent *C.GstObject, event *C.GstEvent) {
	// TODO: not found user_data
}

type PadEventFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_event  Event
}

func GetPointer_myPadEventFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadEventFunction())
}

//export myGstPadEventFunction
func myGstPadEventFunction(pad *C.GstPad, parent *C.GstObject, event *C.GstEvent) {
	// TODO: not found user_data
}

// Flags PadFlags
type PadFlags int

const (
	PadFlagsBlocked         PadFlags = 16
	PadFlagsFlushing        PadFlags = 32
	PadFlagsEos             PadFlags = 64
	PadFlagsBlocking        PadFlags = 128
	PadFlagsNeedParent      PadFlags = 256
	PadFlagsNeedReconfigure PadFlags = 512
	PadFlagsPendingEvents   PadFlags = 1024
	PadFlagsFixedCaps       PadFlags = 2048
	PadFlagsProxyCaps       PadFlags = 4096
	PadFlagsProxyAllocation PadFlags = 8192
	PadFlagsProxyScheduling PadFlags = 16384
	PadFlagsAcceptIntersect PadFlags = 32768
	PadFlagsAcceptTemplate  PadFlags = 65536
	PadFlagsLast            PadFlags = 1048576
)

func PadFlagsGetType() gi.GType {
	ret := _I.GetGType(94, "PadFlags")
	return ret
}

type PadForwardFunctionStruct struct {
	F_pad Pad
}

func GetPointer_myPadForwardFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadForwardFunction())
}

//export myGstPadForwardFunction
func myGstPadForwardFunction(pad *C.GstPad, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PadForwardFunctionStruct{
		F_pad: WrapPad(unsafe.Pointer(pad)),
	}
	fn(args)
}

type PadGetRangeFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_offset uint64
	F_length uint32
	F_buffer Buffer
}

func GetPointer_myPadGetRangeFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadGetRangeFunction())
}

//export myGstPadGetRangeFunction
func myGstPadGetRangeFunction(pad *C.GstPad, parent *C.GstObject, offset C.guint64, length C.guint32, buffer *C.GstBuffer) {
	// TODO: not found user_data
}

type PadIterIntLinkFunctionStruct struct {
	F_pad    Pad
	F_parent Object
}

func GetPointer_myPadIterIntLinkFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadIterIntLinkFunction())
}

//export myGstPadIterIntLinkFunction
func myGstPadIterIntLinkFunction(pad *C.GstPad, parent *C.GstObject) {
	// TODO: not found user_data
}

// Flags PadLinkCheck
type PadLinkCheckFlags int

const (
	PadLinkCheckNothing       PadLinkCheckFlags = 0
	PadLinkCheckHierarchy     PadLinkCheckFlags = 1
	PadLinkCheckTemplateCaps  PadLinkCheckFlags = 2
	PadLinkCheckCaps          PadLinkCheckFlags = 4
	PadLinkCheckNoReconfigure PadLinkCheckFlags = 8
	PadLinkCheckDefault       PadLinkCheckFlags = 5
)

func PadLinkCheckGetType() gi.GType {
	ret := _I.GetGType(95, "PadLinkCheck")
	return ret
}

type PadLinkFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_peer   Pad
}

func GetPointer_myPadLinkFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadLinkFunction())
}

//export myGstPadLinkFunction
func myGstPadLinkFunction(pad *C.GstPad, parent *C.GstObject, peer *C.GstPad) {
	// TODO: not found user_data
}

// Enum PadLinkReturn
type PadLinkReturnEnum int

const (
	PadLinkReturnOk             PadLinkReturnEnum = 0
	PadLinkReturnWrongHierarchy PadLinkReturnEnum = -1
	PadLinkReturnWasLinked      PadLinkReturnEnum = -2
	PadLinkReturnWrongDirection PadLinkReturnEnum = -3
	PadLinkReturnNoformat       PadLinkReturnEnum = -4
	PadLinkReturnNosched        PadLinkReturnEnum = -5
	PadLinkReturnRefused        PadLinkReturnEnum = -6
)

func PadLinkReturnGetType() gi.GType {
	ret := _I.GetGType(96, "PadLinkReturn")
	return ret
}

// Enum PadMode
type PadModeEnum int

const (
	PadModeNone PadModeEnum = 0
	PadModePush PadModeEnum = 1
	PadModePull PadModeEnum = 2
)

func PadModeGetType() gi.GType {
	ret := _I.GetGType(97, "PadMode")
	return ret
}

// Enum PadPresence
type PadPresenceEnum int

const (
	PadPresenceAlways    PadPresenceEnum = 0
	PadPresenceSometimes PadPresenceEnum = 1
	PadPresenceRequest   PadPresenceEnum = 2
)

func PadPresenceGetType() gi.GType {
	ret := _I.GetGType(98, "PadPresence")
	return ret
}

// Struct PadPrivate
type PadPrivate struct {
	P unsafe.Pointer
}

func PadPrivateGetType() gi.GType {
	ret := _I.GetGType(99, "PadPrivate")
	return ret
}

type PadProbeCallbackStruct struct {
	F_pad  Pad
	F_info PadProbeInfo
}

func GetPointer_myPadProbeCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadProbeCallback())
}

//export myGstPadProbeCallback
func myGstPadProbeCallback(pad *C.GstPad, info *C.GstPadProbeInfo, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PadProbeCallbackStruct{
		F_pad:  WrapPad(unsafe.Pointer(pad)),
		F_info: PadProbeInfo{P: unsafe.Pointer(info)},
	}
	fn(args)
}

// Struct PadProbeInfo
type PadProbeInfo struct {
	P unsafe.Pointer
}

const SizeOfStructPadProbeInfo = 40

func PadProbeInfoGetType() gi.GType {
	ret := _I.GetGType(100, "PadProbeInfo")
	return ret
}

// gst_pad_probe_info_get_buffer
//
// [ result ] trans: nothing
//
func (v PadProbeInfo) GetBuffer() (result Buffer) {
	iv, err := _I.Get(726, "PadProbeInfo", "get_buffer")
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

// gst_pad_probe_info_get_buffer_list
//
// [ result ] trans: nothing
//
func (v PadProbeInfo) GetBufferList() (result BufferList) {
	iv, err := _I.Get(727, "PadProbeInfo", "get_buffer_list")
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

// gst_pad_probe_info_get_event
//
// [ result ] trans: nothing
//
func (v PadProbeInfo) GetEvent() (result Event) {
	iv, err := _I.Get(728, "PadProbeInfo", "get_event")
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

// gst_pad_probe_info_get_query
//
// [ result ] trans: nothing
//
func (v PadProbeInfo) GetQuery() (result Query) {
	iv, err := _I.Get(729, "PadProbeInfo", "get_query")
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

// Enum PadProbeReturn
type PadProbeReturnEnum int

const (
	PadProbeReturnDrop    PadProbeReturnEnum = 0
	PadProbeReturnOk      PadProbeReturnEnum = 1
	PadProbeReturnRemove  PadProbeReturnEnum = 2
	PadProbeReturnPass    PadProbeReturnEnum = 3
	PadProbeReturnHandled PadProbeReturnEnum = 4
)

func PadProbeReturnGetType() gi.GType {
	ret := _I.GetGType(101, "PadProbeReturn")
	return ret
}

// Flags PadProbeType
type PadProbeTypeFlags int

const (
	PadProbeTypeInvalid         PadProbeTypeFlags = 0
	PadProbeTypeIdle            PadProbeTypeFlags = 1
	PadProbeTypeBlock           PadProbeTypeFlags = 2
	PadProbeTypeBuffer          PadProbeTypeFlags = 16
	PadProbeTypeBufferList      PadProbeTypeFlags = 32
	PadProbeTypeEventDownstream PadProbeTypeFlags = 64
	PadProbeTypeEventUpstream   PadProbeTypeFlags = 128
	PadProbeTypeEventFlush      PadProbeTypeFlags = 256
	PadProbeTypeQueryDownstream PadProbeTypeFlags = 512
	PadProbeTypeQueryUpstream   PadProbeTypeFlags = 1024
	PadProbeTypePush            PadProbeTypeFlags = 4096
	PadProbeTypePull            PadProbeTypeFlags = 8192
	PadProbeTypeBlocking        PadProbeTypeFlags = 3
	PadProbeTypeDataDownstream  PadProbeTypeFlags = 112
	PadProbeTypeDataUpstream    PadProbeTypeFlags = 128
	PadProbeTypeDataBoth        PadProbeTypeFlags = 240
	PadProbeTypeBlockDownstream PadProbeTypeFlags = 114
	PadProbeTypeBlockUpstream   PadProbeTypeFlags = 130
	PadProbeTypeEventBoth       PadProbeTypeFlags = 192
	PadProbeTypeQueryBoth       PadProbeTypeFlags = 1536
	PadProbeTypeAllBoth         PadProbeTypeFlags = 1776
	PadProbeTypeScheduling      PadProbeTypeFlags = 12288
)

func PadProbeTypeGetType() gi.GType {
	ret := _I.GetGType(102, "PadProbeType")
	return ret
}

type PadQueryFunctionStruct struct {
	F_pad    Pad
	F_parent Object
	F_query  Query
}

func GetPointer_myPadQueryFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadQueryFunction())
}

//export myGstPadQueryFunction
func myGstPadQueryFunction(pad *C.GstPad, parent *C.GstObject, query *C.GstQuery) {
	// TODO: not found user_data
}

type PadStickyEventsForeachFunctionStruct struct {
	F_pad   Pad
	F_event Event
}

func GetPointer_myPadStickyEventsForeachFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadStickyEventsForeachFunction())
}

//export myGstPadStickyEventsForeachFunction
func myGstPadStickyEventsForeachFunction(pad *C.GstPad, event *C.GstEvent, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PadStickyEventsForeachFunctionStruct{
		F_pad:   WrapPad(unsafe.Pointer(pad)),
		F_event: Event{P: unsafe.Pointer(event)},
	}
	fn(args)
}

// Object PadTemplate
type PadTemplate struct {
	Object
}

func WrapPadTemplate(p unsafe.Pointer) (r PadTemplate) { r.P = p; return }

type IPadTemplate interface{ P_PadTemplate() unsafe.Pointer }

func (v PadTemplate) P_PadTemplate() unsafe.Pointer { return v.P }
func PadTemplateGetType() gi.GType {
	ret := _I.GetGType(103, "PadTemplate")
	return ret
}

// gst_pad_template_new
//
// [ name_template ] trans: nothing
//
// [ direction ] trans: nothing
//
// [ presence ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPadTemplate(name_template string, direction PadDirectionEnum, presence PadPresenceEnum, caps Caps) (result PadTemplate) {
	iv, err := _I.Get(730, "PadTemplate", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name_template := gi.CString(name_template)
	arg_name_template := gi.NewStringArgument(c_name_template)
	arg_direction := gi.NewIntArgument(int(direction))
	arg_presence := gi.NewIntArgument(int(presence))
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_name_template, arg_direction, arg_presence, arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name_template)
	result.P = ret.Pointer()
	return
}

// gst_pad_template_new_from_static_pad_template_with_gtype
//
// [ pad_template ] trans: nothing
//
// [ pad_type ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPadTemplateFromStaticPadTemplateWithGtype(pad_template StaticPadTemplate, pad_type gi.GType) (result PadTemplate) {
	iv, err := _I.Get(731, "PadTemplate", "new_from_static_pad_template_with_gtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_pad_template := gi.NewPointerArgument(pad_template.P)
	arg_pad_type := gi.NewUintArgument(uint(pad_type))
	args := []gi.Argument{arg_pad_template, arg_pad_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_pad_template_new_with_gtype
//
// [ name_template ] trans: nothing
//
// [ direction ] trans: nothing
//
// [ presence ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ pad_type ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPadTemplateWithGtype(name_template string, direction PadDirectionEnum, presence PadPresenceEnum, caps Caps, pad_type gi.GType) (result PadTemplate) {
	iv, err := _I.Get(732, "PadTemplate", "new_with_gtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name_template := gi.CString(name_template)
	arg_name_template := gi.NewStringArgument(c_name_template)
	arg_direction := gi.NewIntArgument(int(direction))
	arg_presence := gi.NewIntArgument(int(presence))
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_pad_type := gi.NewUintArgument(uint(pad_type))
	args := []gi.Argument{arg_name_template, arg_direction, arg_presence, arg_caps, arg_pad_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name_template)
	result.P = ret.Pointer()
	return
}

// gst_pad_template_get_caps
//
// [ result ] trans: everything
//
func (v PadTemplate) GetCaps() (result Caps) {
	iv, err := _I.Get(733, "PadTemplate", "get_caps")
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

// gst_pad_template_pad_created
//
// [ pad ] trans: nothing
//
func (v PadTemplate) PadCreated(pad IPad) {
	iv, err := _I.Get(734, "PadTemplate", "pad_created")
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

// ignore GType struct PadTemplateClass

// Flags PadTemplateFlags
type PadTemplateFlags int

const (
	PadTemplateFlagsLast PadTemplateFlags = 256
)

func PadTemplateFlagsGetType() gi.GType {
	ret := _I.GetGType(104, "PadTemplateFlags")
	return ret
}

type PadUnlinkFunctionStruct struct {
	F_pad    Pad
	F_parent Object
}

func GetPointer_myPadUnlinkFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPadUnlinkFunction())
}

//export myGstPadUnlinkFunction
func myGstPadUnlinkFunction(pad *C.GstPad, parent *C.GstObject) {
	// TODO: not found user_data
}

// Object ParamArray
type ParamArray struct {
	g.ParamSpec
}

func WrapParamArray(p unsafe.Pointer) (r ParamArray) { r.P = p; return }

type IParamArray interface{ P_ParamArray() unsafe.Pointer }

func (v ParamArray) P_ParamArray() unsafe.Pointer { return v.P }
func ParamArrayGetType() gi.GType {
	ret := _I.GetGType(105, "ParamArray")
	return ret
}

// Object ParamFraction
type ParamFraction struct {
	g.ParamSpec
}

func WrapParamFraction(p unsafe.Pointer) (r ParamFraction) { r.P = p; return }

type IParamFraction interface{ P_ParamFraction() unsafe.Pointer }

func (v ParamFraction) P_ParamFraction() unsafe.Pointer { return v.P }
func ParamFractionGetType() gi.GType {
	ret := _I.GetGType(106, "ParamFraction")
	return ret
}

// Struct ParamSpecArray
type ParamSpecArray struct {
	P unsafe.Pointer
}

const SizeOfStructParamSpecArray = 80

func ParamSpecArrayGetType() gi.GType {
	ret := _I.GetGType(107, "ParamSpecArray")
	return ret
}

// Struct ParamSpecFraction
type ParamSpecFraction struct {
	P unsafe.Pointer
}

const SizeOfStructParamSpecFraction = 96

func ParamSpecFractionGetType() gi.GType {
	ret := _I.GetGType(108, "ParamSpecFraction")
	return ret
}

// Struct ParentBufferMeta
type ParentBufferMeta struct {
	P unsafe.Pointer
}

const SizeOfStructParentBufferMeta = 24

func ParentBufferMetaGetType() gi.GType {
	ret := _I.GetGType(109, "ParentBufferMeta")
	return ret
}

// Struct ParseContext
type ParseContext struct {
	P unsafe.Pointer
}

func ParseContextGetType() gi.GType {
	ret := _I.GetGType(110, "ParseContext")
	return ret
}

// gst_parse_context_new
//
// [ result ] trans: everything
//
func NewParseContext() (result ParseContext) {
	iv, err := _I.Get(736, "ParseContext", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_parse_context_copy
//
// [ result ] trans: everything
//
func (v ParseContext) Copy() (result ParseContext) {
	iv, err := _I.Get(737, "ParseContext", "copy")
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

// gst_parse_context_free
//
func (v ParseContext) Free() {
	iv, err := _I.Get(738, "ParseContext", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_parse_context_get_missing_elements
//
// [ result ] trans: everything
//
func (v ParseContext) GetMissingElements() (result gi.CStrArray) {
	iv, err := _I.Get(739, "ParseContext", "get_missing_elements")
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

// Enum ParseError
type ParseErrorEnum int

const (
	ParseErrorSyntax              ParseErrorEnum = 0
	ParseErrorNoSuchElement       ParseErrorEnum = 1
	ParseErrorNoSuchProperty      ParseErrorEnum = 2
	ParseErrorLink                ParseErrorEnum = 3
	ParseErrorCouldNotSetProperty ParseErrorEnum = 4
	ParseErrorEmptyBin            ParseErrorEnum = 5
	ParseErrorEmpty               ParseErrorEnum = 6
	ParseErrorDelayedLink         ParseErrorEnum = 7
)

func ParseErrorGetType() gi.GType {
	ret := _I.GetGType(111, "ParseError")
	return ret
}

// Flags ParseFlags
type ParseFlags int

const (
	ParseFlagsNone                ParseFlags = 0
	ParseFlagsFatalErrors         ParseFlags = 1
	ParseFlagsNoSingleElementBins ParseFlags = 2
	ParseFlagsPlaceInBin          ParseFlags = 4
)

func ParseFlagsGetType() gi.GType {
	ret := _I.GetGType(112, "ParseFlags")
	return ret
}

// Object Pipeline
type Pipeline struct {
	Bin
}

func WrapPipeline(p unsafe.Pointer) (r Pipeline) { r.P = p; return }

type IPipeline interface{ P_Pipeline() unsafe.Pointer }

func (v Pipeline) P_Pipeline() unsafe.Pointer { return v.P }
func PipelineGetType() gi.GType {
	ret := _I.GetGType(113, "Pipeline")
	return ret
}

// gst_pipeline_new
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func NewPipeline(name string) (result Pipeline) {
	iv, err := _I.Get(740, "Pipeline", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_pipeline_auto_clock
//
func (v Pipeline) AutoClock() {
	iv, err := _I.Get(741, "Pipeline", "auto_clock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_pipeline_get_auto_flush_bus
//
// [ result ] trans: nothing
//
func (v Pipeline) GetAutoFlushBus() (result bool) {
	iv, err := _I.Get(742, "Pipeline", "get_auto_flush_bus")
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

// gst_pipeline_get_bus
//
// [ result ] trans: everything
//
func (v Pipeline) GetBus() (result Bus) {
	iv, err := _I.Get(743, "Pipeline", "get_bus")
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

// gst_pipeline_get_delay
//
// [ result ] trans: nothing
//
func (v Pipeline) GetDelay() (result uint64) {
	iv, err := _I.Get(744, "Pipeline", "get_delay")
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

// gst_pipeline_get_latency
//
// [ result ] trans: nothing
//
func (v Pipeline) GetLatency() (result uint64) {
	iv, err := _I.Get(745, "Pipeline", "get_latency")
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

// gst_pipeline_get_pipeline_clock
//
// [ result ] trans: everything
//
func (v Pipeline) GetPipelineClock() (result Clock) {
	iv, err := _I.Get(746, "Pipeline", "get_pipeline_clock")
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

// gst_pipeline_set_auto_flush_bus
//
// [ auto_flush ] trans: nothing
//
func (v Pipeline) SetAutoFlushBus(auto_flush bool) {
	iv, err := _I.Get(747, "Pipeline", "set_auto_flush_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_auto_flush := gi.NewBoolArgument(auto_flush)
	args := []gi.Argument{arg_v, arg_auto_flush}
	iv.Call(args, nil, nil)
}

// gst_pipeline_set_delay
//
// [ delay ] trans: nothing
//
func (v Pipeline) SetDelay(delay uint64) {
	iv, err := _I.Get(748, "Pipeline", "set_delay")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_delay := gi.NewUint64Argument(delay)
	args := []gi.Argument{arg_v, arg_delay}
	iv.Call(args, nil, nil)
}

// gst_pipeline_set_latency
//
// [ latency ] trans: nothing
//
func (v Pipeline) SetLatency(latency uint64) {
	iv, err := _I.Get(749, "Pipeline", "set_latency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_latency := gi.NewUint64Argument(latency)
	args := []gi.Argument{arg_v, arg_latency}
	iv.Call(args, nil, nil)
}

// gst_pipeline_use_clock
//
// [ clock ] trans: nothing
//
func (v Pipeline) UseClock(clock IClock) {
	iv, err := _I.Get(750, "Pipeline", "use_clock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if clock != nil {
		tmp = clock.P_Clock()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_clock := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_clock}
	iv.Call(args, nil, nil)
}

// ignore GType struct PipelineClass

// Flags PipelineFlags
type PipelineFlags int

const (
	PipelineFlagsFixedClock PipelineFlags = 524288
	PipelineFlagsLast       PipelineFlags = 8388608
)

func PipelineFlagsGetType() gi.GType {
	ret := _I.GetGType(114, "PipelineFlags")
	return ret
}

// Struct PipelinePrivate
type PipelinePrivate struct {
	P unsafe.Pointer
}

func PipelinePrivateGetType() gi.GType {
	ret := _I.GetGType(115, "PipelinePrivate")
	return ret
}

// Object Plugin
type Plugin struct {
	Object
}

func WrapPlugin(p unsafe.Pointer) (r Plugin) { r.P = p; return }

type IPlugin interface{ P_Plugin() unsafe.Pointer }

func (v Plugin) P_Plugin() unsafe.Pointer { return v.P }
func PluginGetType() gi.GType {
	ret := _I.GetGType(116, "Plugin")
	return ret
}

// gst_plugin_list_free
//
// [ list ] trans: everything
//
func PluginListFree1(list g.List) {
	iv, err := _I.Get(751, "Plugin", "list_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// gst_plugin_load_by_name
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func PluginLoadByName1(name string) (result Plugin) {
	iv, err := _I.Get(752, "Plugin", "load_by_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_plugin_load_file
//
// [ filename ] trans: nothing
//
// [ result ] trans: everything
//
func PluginLoadFile1(filename string) (result Plugin, err error) {
	iv, err := _I.Get(753, "Plugin", "load_file")
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
	result.P = ret.Pointer()
	return
}

// gst_plugin_register_static
//
// [ major_version ] trans: nothing
//
// [ minor_version ] trans: nothing
//
// [ name ] trans: nothing
//
// [ description ] trans: nothing
//
// [ init_func ] trans: nothing
//
// [ version ] trans: nothing
//
// [ license ] trans: nothing
//
// [ source ] trans: nothing
//
// [ package1 ] trans: nothing
//
// [ origin ] trans: nothing
//
// [ result ] trans: nothing
//
func PluginRegisterStatic1(major_version int32, minor_version int32, name string, description string, init_func int /*TODO_TYPE CALLBACK*/, version string, license string, source string, package1 string, origin string) (result bool) {
	iv, err := _I.Get(754, "Plugin", "register_static")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_description := gi.CString(description)
	c_version := gi.CString(version)
	c_license := gi.CString(license)
	c_source := gi.CString(source)
	c_package1 := gi.CString(package1)
	c_origin := gi.CString(origin)
	arg_major_version := gi.NewInt32Argument(major_version)
	arg_minor_version := gi.NewInt32Argument(minor_version)
	arg_name := gi.NewStringArgument(c_name)
	arg_description := gi.NewStringArgument(c_description)
	arg_init_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPluginInitFunc()))
	arg_version := gi.NewStringArgument(c_version)
	arg_license := gi.NewStringArgument(c_license)
	arg_source := gi.NewStringArgument(c_source)
	arg_package1 := gi.NewStringArgument(c_package1)
	arg_origin := gi.NewStringArgument(c_origin)
	args := []gi.Argument{arg_major_version, arg_minor_version, arg_name, arg_description, arg_init_func, arg_version, arg_license, arg_source, arg_package1, arg_origin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_description)
	gi.Free(c_version)
	gi.Free(c_license)
	gi.Free(c_source)
	gi.Free(c_package1)
	gi.Free(c_origin)
	result = ret.Bool()
	return
}

// gst_plugin_register_static_full
//
// [ major_version ] trans: nothing
//
// [ minor_version ] trans: nothing
//
// [ name ] trans: nothing
//
// [ description ] trans: nothing
//
// [ init_full_func ] trans: nothing
//
// [ version ] trans: nothing
//
// [ license ] trans: nothing
//
// [ source ] trans: nothing
//
// [ package1 ] trans: nothing
//
// [ origin ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func PluginRegisterStaticFull1(major_version int32, minor_version int32, name string, description string, init_full_func int /*TODO_TYPE CALLBACK*/, version string, license string, source string, package1 string, origin string, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(755, "Plugin", "register_static_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_description := gi.CString(description)
	c_version := gi.CString(version)
	c_license := gi.CString(license)
	c_source := gi.CString(source)
	c_package1 := gi.CString(package1)
	c_origin := gi.CString(origin)
	arg_major_version := gi.NewInt32Argument(major_version)
	arg_minor_version := gi.NewInt32Argument(minor_version)
	arg_name := gi.NewStringArgument(c_name)
	arg_description := gi.NewStringArgument(c_description)
	arg_init_full_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPluginInitFullFunc()))
	arg_version := gi.NewStringArgument(c_version)
	arg_license := gi.NewStringArgument(c_license)
	arg_source := gi.NewStringArgument(c_source)
	arg_package1 := gi.NewStringArgument(c_package1)
	arg_origin := gi.NewStringArgument(c_origin)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_major_version, arg_minor_version, arg_name, arg_description, arg_init_full_func, arg_version, arg_license, arg_source, arg_package1, arg_origin, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_description)
	gi.Free(c_version)
	gi.Free(c_license)
	gi.Free(c_source)
	gi.Free(c_package1)
	gi.Free(c_origin)
	result = ret.Bool()
	return
}

// gst_plugin_add_dependency
//
// [ env_vars ] trans: nothing
//
// [ paths ] trans: nothing
//
// [ names ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v Plugin) AddDependency(env_vars gi.CStrArray, paths gi.CStrArray, names gi.CStrArray, flags PluginDependencyFlags) {
	iv, err := _I.Get(756, "Plugin", "add_dependency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_env_vars := gi.NewPointerArgument(env_vars.P)
	arg_paths := gi.NewPointerArgument(paths.P)
	arg_names := gi.NewPointerArgument(names.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_env_vars, arg_paths, arg_names, arg_flags}
	iv.Call(args, nil, nil)
}

// gst_plugin_add_dependency_simple
//
// [ env_vars ] trans: nothing
//
// [ paths ] trans: nothing
//
// [ names ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v Plugin) AddDependencySimple(env_vars string, paths string, names string, flags PluginDependencyFlags) {
	iv, err := _I.Get(757, "Plugin", "add_dependency_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_env_vars := gi.CString(env_vars)
	c_paths := gi.CString(paths)
	c_names := gi.CString(names)
	arg_v := gi.NewPointerArgument(v.P)
	arg_env_vars := gi.NewStringArgument(c_env_vars)
	arg_paths := gi.NewStringArgument(c_paths)
	arg_names := gi.NewStringArgument(c_names)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_env_vars, arg_paths, arg_names, arg_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_env_vars)
	gi.Free(c_paths)
	gi.Free(c_names)
}

// gst_plugin_get_cache_data
//
// [ result ] trans: nothing
//
func (v Plugin) GetCacheData() (result Structure) {
	iv, err := _I.Get(758, "Plugin", "get_cache_data")
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

// gst_plugin_get_description
//
// [ result ] trans: nothing
//
func (v Plugin) GetDescription() (result string) {
	iv, err := _I.Get(759, "Plugin", "get_description")
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

// gst_plugin_get_filename
//
// [ result ] trans: nothing
//
func (v Plugin) GetFilename() (result string) {
	iv, err := _I.Get(760, "Plugin", "get_filename")
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

// gst_plugin_get_license
//
// [ result ] trans: nothing
//
func (v Plugin) GetLicense() (result string) {
	iv, err := _I.Get(761, "Plugin", "get_license")
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

// gst_plugin_get_name
//
// [ result ] trans: nothing
//
func (v Plugin) GetName() (result string) {
	iv, err := _I.Get(762, "Plugin", "get_name")
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

// gst_plugin_get_origin
//
// [ result ] trans: nothing
//
func (v Plugin) GetOrigin() (result string) {
	iv, err := _I.Get(763, "Plugin", "get_origin")
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

// gst_plugin_get_package
//
// [ result ] trans: nothing
//
func (v Plugin) GetPackage() (result string) {
	iv, err := _I.Get(764, "Plugin", "get_package")
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

// gst_plugin_get_release_date_string
//
// [ result ] trans: nothing
//
func (v Plugin) GetReleaseDateString() (result string) {
	iv, err := _I.Get(765, "Plugin", "get_release_date_string")
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

// gst_plugin_get_source
//
// [ result ] trans: nothing
//
func (v Plugin) GetSource() (result string) {
	iv, err := _I.Get(766, "Plugin", "get_source")
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

// gst_plugin_get_version
//
// [ result ] trans: nothing
//
func (v Plugin) GetVersion() (result string) {
	iv, err := _I.Get(767, "Plugin", "get_version")
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

// gst_plugin_is_loaded
//
// [ result ] trans: nothing
//
func (v Plugin) IsLoaded() (result bool) {
	iv, err := _I.Get(768, "Plugin", "is_loaded")
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

// gst_plugin_load
//
// [ result ] trans: everything
//
func (v Plugin) Load() (result Plugin) {
	iv, err := _I.Get(769, "Plugin", "load")
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

// gst_plugin_set_cache_data
//
// [ cache_data ] trans: everything
//
func (v Plugin) SetCacheData(cache_data Structure) {
	iv, err := _I.Get(770, "Plugin", "set_cache_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cache_data := gi.NewPointerArgument(cache_data.P)
	args := []gi.Argument{arg_v, arg_cache_data}
	iv.Call(args, nil, nil)
}

// ignore GType struct PluginClass

// Flags PluginDependencyFlags
type PluginDependencyFlags int

const (
	PluginDependencyFlagsNone                  PluginDependencyFlags = 0
	PluginDependencyFlagsRecurse               PluginDependencyFlags = 1
	PluginDependencyFlagsPathsAreDefaultOnly   PluginDependencyFlags = 2
	PluginDependencyFlagsFileNameIsSuffix      PluginDependencyFlags = 4
	PluginDependencyFlagsFileNameIsPrefix      PluginDependencyFlags = 8
	PluginDependencyFlagsPathsAreRelativeToExe PluginDependencyFlags = 16
)

func PluginDependencyFlagsGetType() gi.GType {
	ret := _I.GetGType(117, "PluginDependencyFlags")
	return ret
}

// Struct PluginDesc
type PluginDesc struct {
	P unsafe.Pointer
}

const SizeOfStructPluginDesc = 112

func PluginDescGetType() gi.GType {
	ret := _I.GetGType(118, "PluginDesc")
	return ret
}

// Enum PluginError
type PluginErrorEnum int

const (
	PluginErrorModule       PluginErrorEnum = 0
	PluginErrorDependencies PluginErrorEnum = 1
	PluginErrorNameMismatch PluginErrorEnum = 2
)

func PluginErrorGetType() gi.GType {
	ret := _I.GetGType(119, "PluginError")
	return ret
}

// Object PluginFeature
type PluginFeature struct {
	Object
}

func WrapPluginFeature(p unsafe.Pointer) (r PluginFeature) { r.P = p; return }

type IPluginFeature interface{ P_PluginFeature() unsafe.Pointer }

func (v PluginFeature) P_PluginFeature() unsafe.Pointer { return v.P }
func PluginFeatureGetType() gi.GType {
	ret := _I.GetGType(120, "PluginFeature")
	return ret
}

// gst_plugin_feature_list_copy
//
// [ list ] trans: nothing
//
// [ result ] trans: everything
//
func PluginFeatureListCopy1(list g.List) (result g.List) {
	iv, err := _I.Get(771, "PluginFeature", "list_copy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_plugin_feature_list_debug
//
// [ list ] trans: nothing
//
func PluginFeatureListDebug1(list g.List) {
	iv, err := _I.Get(772, "PluginFeature", "list_debug")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// gst_plugin_feature_list_free
//
// [ list ] trans: everything
//
func PluginFeatureListFree1(list g.List) {
	iv, err := _I.Get(773, "PluginFeature", "list_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// gst_plugin_feature_rank_compare_func
//
// [ p1 ] trans: nothing
//
// [ p2 ] trans: nothing
//
// [ result ] trans: nothing
//
func PluginFeatureRankCompareFunc1(p1 unsafe.Pointer, p2 unsafe.Pointer) (result int32) {
	iv, err := _I.Get(774, "PluginFeature", "rank_compare_func")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_p1 := gi.NewPointerArgument(p1)
	arg_p2 := gi.NewPointerArgument(p2)
	args := []gi.Argument{arg_p1, arg_p2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_plugin_feature_check_version
//
// [ min_major ] trans: nothing
//
// [ min_minor ] trans: nothing
//
// [ min_micro ] trans: nothing
//
// [ result ] trans: nothing
//
func (v PluginFeature) CheckVersion(min_major uint32, min_minor uint32, min_micro uint32) (result bool) {
	iv, err := _I.Get(775, "PluginFeature", "check_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_min_major := gi.NewUint32Argument(min_major)
	arg_min_minor := gi.NewUint32Argument(min_minor)
	arg_min_micro := gi.NewUint32Argument(min_micro)
	args := []gi.Argument{arg_v, arg_min_major, arg_min_minor, arg_min_micro}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_plugin_feature_get_plugin
//
// [ result ] trans: everything
//
func (v PluginFeature) GetPlugin() (result Plugin) {
	iv, err := _I.Get(776, "PluginFeature", "get_plugin")
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

// gst_plugin_feature_get_plugin_name
//
// [ result ] trans: nothing
//
func (v PluginFeature) GetPluginName() (result string) {
	iv, err := _I.Get(777, "PluginFeature", "get_plugin_name")
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

// gst_plugin_feature_get_rank
//
// [ result ] trans: nothing
//
func (v PluginFeature) GetRank() (result uint32) {
	iv, err := _I.Get(778, "PluginFeature", "get_rank")
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

// gst_plugin_feature_load
//
// [ result ] trans: everything
//
func (v PluginFeature) Load() (result PluginFeature) {
	iv, err := _I.Get(779, "PluginFeature", "load")
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

// gst_plugin_feature_set_rank
//
// [ rank ] trans: nothing
//
func (v PluginFeature) SetRank(rank uint32) {
	iv, err := _I.Get(780, "PluginFeature", "set_rank")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rank := gi.NewUint32Argument(rank)
	args := []gi.Argument{arg_v, arg_rank}
	iv.Call(args, nil, nil)
}

// ignore GType struct PluginFeatureClass

type PluginFeatureFilterStruct struct {
	F_feature PluginFeature
}

func GetPointer_myPluginFeatureFilter() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPluginFeatureFilter())
}

//export myGstPluginFeatureFilter
func myGstPluginFeatureFilter(feature *C.GstPluginFeature, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PluginFeatureFilterStruct{
		F_feature: WrapPluginFeature(unsafe.Pointer(feature)),
	}
	fn(args)
}

type PluginFilterStruct struct {
	F_plugin Plugin
}

func GetPointer_myPluginFilter() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPluginFilter())
}

//export myGstPluginFilter
func myGstPluginFilter(plugin *C.GstPlugin, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PluginFilterStruct{
		F_plugin: WrapPlugin(unsafe.Pointer(plugin)),
	}
	fn(args)
}

// Flags PluginFlags
type PluginFlags int

const (
	PluginFlagsCached      PluginFlags = 16
	PluginFlagsBlacklisted PluginFlags = 32
)

func PluginFlagsGetType() gi.GType {
	ret := _I.GetGType(121, "PluginFlags")
	return ret
}

type PluginInitFullFuncStruct struct {
	F_plugin Plugin
}

func GetPointer_myPluginInitFullFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPluginInitFullFunc())
}

//export myGstPluginInitFullFunc
func myGstPluginInitFullFunc(plugin *C.GstPlugin, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PluginInitFullFuncStruct{
		F_plugin: WrapPlugin(unsafe.Pointer(plugin)),
	}
	fn(args)
}

type PluginInitFuncStruct struct {
	F_plugin Plugin
}

func GetPointer_myPluginInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPluginInitFunc())
}

//export myGstPluginInitFunc
func myGstPluginInitFunc(plugin *C.GstPlugin) {
	// TODO: not found user_data
}

// Struct Poll
type Poll struct {
	P unsafe.Pointer
}

func PollGetType() gi.GType {
	ret := _I.GetGType(122, "Poll")
	return ret
}

// gst_poll_add_fd
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) AddFd(fd PollFD) (result bool) {
	iv, err := _I.Get(781, "Poll", "add_fd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_can_read
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) FdCanRead(fd PollFD) (result bool) {
	iv, err := _I.Get(782, "Poll", "fd_can_read")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_can_write
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) FdCanWrite(fd PollFD) (result bool) {
	iv, err := _I.Get(783, "Poll", "fd_can_write")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_ctl_read
//
// [ fd ] trans: nothing
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) FdCtlRead(fd PollFD, active bool) (result bool) {
	iv, err := _I.Get(784, "Poll", "fd_ctl_read")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_v, arg_fd, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_ctl_write
//
// [ fd ] trans: nothing
//
// [ active ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) FdCtlWrite(fd PollFD, active bool) (result bool) {
	iv, err := _I.Get(785, "Poll", "fd_ctl_write")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_v, arg_fd, arg_active}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_has_closed
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) FdHasClosed(fd PollFD) (result bool) {
	iv, err := _I.Get(786, "Poll", "fd_has_closed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_has_error
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) FdHasError(fd PollFD) (result bool) {
	iv, err := _I.Get(787, "Poll", "fd_has_error")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_fd_ignored
//
// [ fd ] trans: nothing
//
func (v Poll) FdIgnored(fd PollFD) {
	iv, err := _I.Get(788, "Poll", "fd_ignored")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// gst_poll_free
//
func (v Poll) Free() {
	iv, err := _I.Get(789, "Poll", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_poll_get_read_gpollfd
//
// [ fd ] trans: nothing
//
func (v Poll) GetReadGpollfd(fd g.PollFD) {
	iv, err := _I.Get(790, "Poll", "get_read_gpollfd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// gst_poll_read_control
//
// [ result ] trans: nothing
//
func (v Poll) ReadControl() (result bool) {
	iv, err := _I.Get(791, "Poll", "read_control")
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

// gst_poll_remove_fd
//
// [ fd ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) RemoveFd(fd PollFD) (result bool) {
	iv, err := _I.Get(792, "Poll", "remove_fd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_restart
//
func (v Poll) Restart() {
	iv, err := _I.Get(793, "Poll", "restart")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_poll_set_controllable
//
// [ controllable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) SetControllable(controllable bool) (result bool) {
	iv, err := _I.Get(794, "Poll", "set_controllable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_controllable := gi.NewBoolArgument(controllable)
	args := []gi.Argument{arg_v, arg_controllable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_poll_set_flushing
//
// [ flushing ] trans: nothing
//
func (v Poll) SetFlushing(flushing bool) {
	iv, err := _I.Get(795, "Poll", "set_flushing")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flushing := gi.NewBoolArgument(flushing)
	args := []gi.Argument{arg_v, arg_flushing}
	iv.Call(args, nil, nil)
}

// gst_poll_wait
//
// [ timeout ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Poll) Wait(timeout uint64) (result int32) {
	iv, err := _I.Get(796, "Poll", "wait")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_poll_write_control
//
// [ result ] trans: nothing
//
func (v Poll) WriteControl() (result bool) {
	iv, err := _I.Get(797, "Poll", "write_control")
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

// Struct PollFD
type PollFD struct {
	P unsafe.Pointer
}

const SizeOfStructPollFD = 8

func PollFDGetType() gi.GType {
	ret := _I.GetGType(123, "PollFD")
	return ret
}

// gst_poll_fd_init
//
func (v PollFD) Init() {
	iv, err := _I.Get(798, "PollFD", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Interface Preset
type Preset struct {
	PresetIfc
	P unsafe.Pointer
}
type PresetIfc struct{}
type IPreset interface{ P_Preset() unsafe.Pointer }

func (v Preset) P_Preset() unsafe.Pointer { return v.P }
func PresetGetType() gi.GType {
	ret := _I.GetGType(124, "Preset")
	return ret
}

// gst_preset_set_app_dir
//
// [ app_dir ] trans: nothing
//
// [ result ] trans: nothing
//
func PresetSetAppDir1(app_dir string) (result bool) {
	iv, err := _I.Get(800, "Preset", "set_app_dir")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_app_dir := gi.CString(app_dir)
	arg_app_dir := gi.NewStringArgument(c_app_dir)
	args := []gi.Argument{arg_app_dir}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_app_dir)
	result = ret.Bool()
	return
}

// gst_preset_delete_preset
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PresetIfc) DeletePreset(name string) (result bool) {
	iv, err := _I.Get(801, "Preset", "delete_preset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_preset_get_meta
//
// [ name ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v *PresetIfc) GetMeta(name string, tag string) (result bool, value string) {
	iv, err := _I.Get(802, "Preset", "get_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_name, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_tag)
	value = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// gst_preset_get_preset_names
//
// [ result ] trans: everything
//
func (v *PresetIfc) GetPresetNames() (result gi.CStrArray) {
	iv, err := _I.Get(803, "Preset", "get_preset_names")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gst_preset_get_property_names
//
// [ result ] trans: everything
//
func (v *PresetIfc) GetPropertyNames() (result gi.CStrArray) {
	iv, err := _I.Get(804, "Preset", "get_property_names")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gst_preset_is_editable
//
// [ result ] trans: nothing
//
func (v *PresetIfc) IsEditable() (result bool) {
	iv, err := _I.Get(805, "Preset", "is_editable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_preset_load_preset
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PresetIfc) LoadPreset(name string) (result bool) {
	iv, err := _I.Get(806, "Preset", "load_preset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_preset_rename_preset
//
// [ old_name ] trans: nothing
//
// [ new_name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PresetIfc) RenamePreset(old_name string, new_name string) (result bool) {
	iv, err := _I.Get(807, "Preset", "rename_preset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_old_name := gi.CString(old_name)
	c_new_name := gi.CString(new_name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_old_name := gi.NewStringArgument(c_old_name)
	arg_new_name := gi.NewStringArgument(c_new_name)
	args := []gi.Argument{arg_v, arg_old_name, arg_new_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_old_name)
	gi.Free(c_new_name)
	result = ret.Bool()
	return
}

// gst_preset_save_preset
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PresetIfc) SavePreset(name string) (result bool) {
	iv, err := _I.Get(808, "Preset", "save_preset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_preset_set_meta
//
// [ name ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PresetIfc) SetMeta(name string, tag string, value string) (result bool) {
	iv, err := _I.Get(809, "Preset", "set_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_tag := gi.CString(tag)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_name := gi.NewStringArgument(c_name)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_name, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_tag)
	gi.Free(c_value)
	result = ret.Bool()
	return
}

// ignore GType struct PresetInterface

// Enum ProgressType
type ProgressTypeEnum int

const (
	ProgressTypeStart    ProgressTypeEnum = 0
	ProgressTypeContinue ProgressTypeEnum = 1
	ProgressTypeComplete ProgressTypeEnum = 2
	ProgressTypeCanceled ProgressTypeEnum = 3
	ProgressTypeError    ProgressTypeEnum = 4
)

func ProgressTypeGetType() gi.GType {
	ret := _I.GetGType(125, "ProgressType")
	return ret
}

// Struct Promise
type Promise struct {
	P unsafe.Pointer
}

const SizeOfStructPromise = 64

func PromiseGetType() gi.GType {
	ret := _I.GetGType(126, "Promise")
	return ret
}

// gst_promise_new
//
// [ result ] trans: everything
//
func NewPromise() (result Promise) {
	iv, err := _I.Get(810, "Promise", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_promise_new_with_change_func
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: everything
//
func NewPromiseWithChangeFunc(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result Promise) {
	iv, err := _I.Get(811, "Promise", "new_with_change_func")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPromiseChangeFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_func1, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_promise_expire
//
func (v Promise) Expire() {
	iv, err := _I.Get(812, "Promise", "expire")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_promise_get_reply
//
// [ result ] trans: nothing
//
func (v Promise) GetReply() (result Structure) {
	iv, err := _I.Get(813, "Promise", "get_reply")
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

// gst_promise_interrupt
//
func (v Promise) Interrupt() {
	iv, err := _I.Get(814, "Promise", "interrupt")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_promise_reply
//
// [ s ] trans: everything
//
func (v Promise) Reply(s Structure) {
	iv, err := _I.Get(815, "Promise", "reply")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_s := gi.NewPointerArgument(s.P)
	args := []gi.Argument{arg_v, arg_s}
	iv.Call(args, nil, nil)
}

// gst_promise_wait
//
// [ result ] trans: nothing
//
func (v Promise) Wait() (result PromiseResultEnum) {
	iv, err := _I.Get(816, "Promise", "wait")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PromiseResultEnum(ret.Int())
	return
}

type PromiseChangeFuncStruct struct {
	F_promise Promise
}

func GetPointer_myPromiseChangeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstPromiseChangeFunc())
}

//export myGstPromiseChangeFunc
func myGstPromiseChangeFunc(promise *C.GstPromise, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PromiseChangeFuncStruct{
		F_promise: Promise{P: unsafe.Pointer(promise)},
	}
	fn(args)
}

// Enum PromiseResult
type PromiseResultEnum int

const (
	PromiseResultPending     PromiseResultEnum = 0
	PromiseResultInterrupted PromiseResultEnum = 1
	PromiseResultReplied     PromiseResultEnum = 2
	PromiseResultExpired     PromiseResultEnum = 3
)

func PromiseResultGetType() gi.GType {
	ret := _I.GetGType(127, "PromiseResult")
	return ret
}

// Struct ProtectionMeta
type ProtectionMeta struct {
	P unsafe.Pointer
}

const SizeOfStructProtectionMeta = 24

func ProtectionMetaGetType() gi.GType {
	ret := _I.GetGType(128, "ProtectionMeta")
	return ret
}

// Object ProxyPad
type ProxyPad struct {
	Pad
}

func WrapProxyPad(p unsafe.Pointer) (r ProxyPad) { r.P = p; return }

type IProxyPad interface{ P_ProxyPad() unsafe.Pointer }

func (v ProxyPad) P_ProxyPad() unsafe.Pointer { return v.P }
func ProxyPadGetType() gi.GType {
	ret := _I.GetGType(129, "ProxyPad")
	return ret
}

// gst_proxy_pad_chain_default
//
// [ pad ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ buffer ] trans: everything
//
// [ result ] trans: nothing
//
func ProxyPadChainDefault1(pad IPad, parent IObject, buffer Buffer) (result FlowReturnEnum) {
	iv, err := _I.Get(818, "ProxyPad", "chain_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	arg_pad := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_pad, arg_parent, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_proxy_pad_chain_list_default
//
// [ pad ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ list ] trans: everything
//
// [ result ] trans: nothing
//
func ProxyPadChainListDefault1(pad IPad, parent IObject, list BufferList) (result FlowReturnEnum) {
	iv, err := _I.Get(819, "ProxyPad", "chain_list_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	arg_pad := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_pad, arg_parent, arg_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_proxy_pad_getrange_default
//
// [ pad ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ size ] trans: nothing
//
// [ buffer ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func ProxyPadGetrangeDefault1(pad IPad, parent IObject, offset uint64, size uint32) (result FlowReturnEnum, buffer Buffer) {
	iv, err := _I.Get(820, "ProxyPad", "getrange_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	arg_pad := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	arg_offset := gi.NewUint64Argument(offset)
	arg_size := gi.NewUint32Argument(size)
	arg_buffer := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pad, arg_parent, arg_offset, arg_size, arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	buffer.P = outArgs[0].Pointer()
	result = FlowReturnEnum(ret.Int())
	return
}

// gst_proxy_pad_iterate_internal_links_default
//
// [ pad ] trans: nothing
//
// [ parent ] trans: nothing
//
// [ result ] trans: everything
//
func ProxyPadIterateInternalLinksDefault1(pad IPad, parent IObject) (result Iterator) {
	iv, err := _I.Get(821, "ProxyPad", "iterate_internal_links_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pad != nil {
		tmp = pad.P_Pad()
	}
	var tmp1 unsafe.Pointer
	if parent != nil {
		tmp1 = parent.P_Object()
	}
	arg_pad := gi.NewPointerArgument(tmp)
	arg_parent := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_pad, arg_parent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_proxy_pad_get_internal
//
// [ result ] trans: everything
//
func (v ProxyPad) GetInternal() (result ProxyPad) {
	iv, err := _I.Get(822, "ProxyPad", "get_internal")
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

// ignore GType struct ProxyPadClass

// Struct ProxyPadPrivate
type ProxyPadPrivate struct {
	P unsafe.Pointer
}

func ProxyPadPrivateGetType() gi.GType {
	ret := _I.GetGType(130, "ProxyPadPrivate")
	return ret
}

// Enum QOSType
type QOSTypeEnum int

const (
	QOSTypeOverflow  QOSTypeEnum = 0
	QOSTypeUnderflow QOSTypeEnum = 1
	QOSTypeThrottle  QOSTypeEnum = 2
)

func QOSTypeGetType() gi.GType {
	ret := _I.GetGType(131, "QOSType")
	return ret
}

// Struct Query
type Query struct {
	P unsafe.Pointer
}

const SizeOfStructQuery = 72

func QueryGetType() gi.GType {
	ret := _I.GetGType(132, "Query")
	return ret
}

// gst_query_new_accept_caps
//
// [ caps ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryAcceptCaps(caps Caps) (result Query) {
	iv, err := _I.Get(823, "Query", "new_accept_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_caps}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_allocation
//
// [ caps ] trans: nothing
//
// [ need_pool ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryAllocation(caps Caps, need_pool bool) (result Query) {
	iv, err := _I.Get(824, "Query", "new_allocation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_need_pool := gi.NewBoolArgument(need_pool)
	args := []gi.Argument{arg_caps, arg_need_pool}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_buffering
//
// [ format ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryBuffering(format FormatEnum) (result Query) {
	iv, err := _I.Get(825, "Query", "new_buffering")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_caps
//
// [ filter ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryCaps(filter Caps) (result Query) {
	iv, err := _I.Get(826, "Query", "new_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_filter := gi.NewPointerArgument(filter.P)
	args := []gi.Argument{arg_filter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_context
//
// [ context_type ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryContext(context_type string) (result Query) {
	iv, err := _I.Get(827, "Query", "new_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_type := gi.CString(context_type)
	arg_context_type := gi.NewStringArgument(c_context_type)
	args := []gi.Argument{arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_type)
	result.P = ret.Pointer()
	return
}

// gst_query_new_convert
//
// [ src_format ] trans: nothing
//
// [ value ] trans: nothing
//
// [ dest_format ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryConvert(src_format FormatEnum, value int64, dest_format FormatEnum) (result Query) {
	iv, err := _I.Get(828, "Query", "new_convert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src_format := gi.NewIntArgument(int(src_format))
	arg_value := gi.NewInt64Argument(value)
	arg_dest_format := gi.NewIntArgument(int(dest_format))
	args := []gi.Argument{arg_src_format, arg_value, arg_dest_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_custom
//
// [ type1 ] trans: nothing
//
// [ structure ] trans: everything
//
// [ result ] trans: everything
//
func NewQueryCustom(type1 QueryTypeEnum, structure Structure) (result Query) {
	iv, err := _I.Get(829, "Query", "new_custom")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_type1, arg_structure}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_drain
//
// [ result ] trans: everything
//
func NewQueryDrain() (result Query) {
	iv, err := _I.Get(830, "Query", "new_drain")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_duration
//
// [ format ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryDuration(format FormatEnum) (result Query) {
	iv, err := _I.Get(831, "Query", "new_duration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_formats
//
// [ result ] trans: everything
//
func NewQueryFormats() (result Query) {
	iv, err := _I.Get(832, "Query", "new_formats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_latency
//
// [ result ] trans: everything
//
func NewQueryLatency() (result Query) {
	iv, err := _I.Get(833, "Query", "new_latency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_position
//
// [ format ] trans: nothing
//
// [ result ] trans: everything
//
func NewQueryPosition(format FormatEnum) (result Query) {
	iv, err := _I.Get(834, "Query", "new_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_scheduling
//
// [ result ] trans: everything
//
func NewQueryScheduling() (result Query) {
	iv, err := _I.Get(835, "Query", "new_scheduling")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_seeking
//
// [ format ] trans: nothing
//
// [ result ] trans: everything
//
func NewQuerySeeking(format FormatEnum) (result Query) {
	iv, err := _I.Get(836, "Query", "new_seeking")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_segment
//
// [ format ] trans: nothing
//
// [ result ] trans: everything
//
func NewQuerySegment(format FormatEnum) (result Query) {
	iv, err := _I.Get(837, "Query", "new_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_new_uri
//
// [ result ] trans: everything
//
func NewQueryUri() (result Query) {
	iv, err := _I.Get(838, "Query", "new_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_query_add_allocation_meta
//
// [ api ] trans: nothing
//
// [ params ] trans: nothing
//
func (v Query) AddAllocationMeta(api gi.GType, params Structure) {
	iv, err := _I.Get(839, "Query", "add_allocation_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_api := gi.NewUintArgument(uint(api))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_api, arg_params}
	iv.Call(args, nil, nil)
}

// gst_query_add_allocation_param
//
// [ allocator ] trans: nothing
//
// [ params ] trans: nothing
//
func (v Query) AddAllocationParam(allocator IAllocator, params AllocationParams) {
	iv, err := _I.Get(840, "Query", "add_allocation_param")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if allocator != nil {
		tmp = allocator.P_Allocator()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_allocator := gi.NewPointerArgument(tmp)
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_allocator, arg_params}
	iv.Call(args, nil, nil)
}

// gst_query_add_allocation_pool
//
// [ pool ] trans: nothing
//
// [ size ] trans: nothing
//
// [ min_buffers ] trans: nothing
//
// [ max_buffers ] trans: nothing
//
func (v Query) AddAllocationPool(pool IBufferPool, size uint32, min_buffers uint32, max_buffers uint32) {
	iv, err := _I.Get(841, "Query", "add_allocation_pool")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pool != nil {
		tmp = pool.P_BufferPool()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pool := gi.NewPointerArgument(tmp)
	arg_size := gi.NewUint32Argument(size)
	arg_min_buffers := gi.NewUint32Argument(min_buffers)
	arg_max_buffers := gi.NewUint32Argument(max_buffers)
	args := []gi.Argument{arg_v, arg_pool, arg_size, arg_min_buffers, arg_max_buffers}
	iv.Call(args, nil, nil)
}

// gst_query_add_buffering_range
//
// [ start ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Query) AddBufferingRange(start int64, stop int64) (result bool) {
	iv, err := _I.Get(842, "Query", "add_buffering_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewInt64Argument(start)
	arg_stop := gi.NewInt64Argument(stop)
	args := []gi.Argument{arg_v, arg_start, arg_stop}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_query_add_scheduling_mode
//
// [ mode ] trans: nothing
//
func (v Query) AddSchedulingMode(mode PadModeEnum) {
	iv, err := _I.Get(843, "Query", "add_scheduling_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	iv.Call(args, nil, nil)
}

// gst_query_find_allocation_meta
//
// [ api ] trans: nothing
//
// [ index ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Query) FindAllocationMeta(api gi.GType) (result bool, index uint32) {
	iv, err := _I.Get(844, "Query", "find_allocation_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_api := gi.NewUintArgument(uint(api))
	arg_index := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_api, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	index = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_query_get_n_allocation_metas
//
// [ result ] trans: nothing
//
func (v Query) GetNAllocationMetas() (result uint32) {
	iv, err := _I.Get(845, "Query", "get_n_allocation_metas")
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

// gst_query_get_n_allocation_params
//
// [ result ] trans: nothing
//
func (v Query) GetNAllocationParams() (result uint32) {
	iv, err := _I.Get(846, "Query", "get_n_allocation_params")
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

// gst_query_get_n_allocation_pools
//
// [ result ] trans: nothing
//
func (v Query) GetNAllocationPools() (result uint32) {
	iv, err := _I.Get(847, "Query", "get_n_allocation_pools")
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

// gst_query_get_n_buffering_ranges
//
// [ result ] trans: nothing
//
func (v Query) GetNBufferingRanges() (result uint32) {
	iv, err := _I.Get(848, "Query", "get_n_buffering_ranges")
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

// gst_query_get_n_scheduling_modes
//
// [ result ] trans: nothing
//
func (v Query) GetNSchedulingModes() (result uint32) {
	iv, err := _I.Get(849, "Query", "get_n_scheduling_modes")
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

// gst_query_get_structure
//
// [ result ] trans: nothing
//
func (v Query) GetStructure() (result Structure) {
	iv, err := _I.Get(850, "Query", "get_structure")
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

// gst_query_has_scheduling_mode
//
// [ mode ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Query) HasSchedulingMode(mode PadModeEnum) (result bool) {
	iv, err := _I.Get(851, "Query", "has_scheduling_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_query_has_scheduling_mode_with_flags
//
// [ mode ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Query) HasSchedulingModeWithFlags(mode PadModeEnum, flags SchedulingFlags) (result bool) {
	iv, err := _I.Get(852, "Query", "has_scheduling_mode_with_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_mode, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_query_parse_accept_caps
//
// [ caps ] trans: nothing, dir: out
//
func (v Query) ParseAcceptCaps() (caps Caps) {
	iv, err := _I.Get(853, "Query", "parse_accept_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_caps}
	iv.Call(args, nil, &outArgs[0])
	caps.P = outArgs[0].Pointer()
	return
}

// gst_query_parse_accept_caps_result
//
// [ result ] trans: everything, dir: out
//
func (v Query) ParseAcceptCapsResult() (result bool) {
	iv, err := _I.Get(854, "Query", "parse_accept_caps_result")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_result}
	iv.Call(args, nil, &outArgs[0])
	result = outArgs[0].Bool()
	return
}

// gst_query_parse_allocation
//
// [ caps ] trans: nothing, dir: out
//
// [ need_pool ] trans: everything, dir: out
//
func (v Query) ParseAllocation() (caps Caps, need_pool bool) {
	iv, err := _I.Get(855, "Query", "parse_allocation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_need_pool := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_caps, arg_need_pool}
	iv.Call(args, nil, &outArgs[0])
	caps.P = outArgs[0].Pointer()
	need_pool = outArgs[1].Bool()
	return
}

// gst_query_parse_buffering_percent
//
// [ busy ] trans: everything, dir: out
//
// [ percent ] trans: everything, dir: out
//
func (v Query) ParseBufferingPercent() (busy bool, percent int32) {
	iv, err := _I.Get(856, "Query", "parse_buffering_percent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_busy := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_percent := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_busy, arg_percent}
	iv.Call(args, nil, &outArgs[0])
	busy = outArgs[0].Bool()
	percent = outArgs[1].Int32()
	return
}

// gst_query_parse_buffering_range
//
// [ format ] trans: everything, dir: out
//
// [ start ] trans: everything, dir: out
//
// [ stop ] trans: everything, dir: out
//
// [ estimated_total ] trans: everything, dir: out
//
func (v Query) ParseBufferingRange() (format FormatEnum, start int64, stop int64, estimated_total int64) {
	iv, err := _I.Get(857, "Query", "parse_buffering_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_stop := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_estimated_total := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_format, arg_start, arg_stop, arg_estimated_total}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	start = outArgs[1].Int64()
	stop = outArgs[2].Int64()
	estimated_total = outArgs[3].Int64()
	return
}

// gst_query_parse_buffering_stats
//
// [ mode ] trans: everything, dir: out
//
// [ avg_in ] trans: everything, dir: out
//
// [ avg_out ] trans: everything, dir: out
//
// [ buffering_left ] trans: everything, dir: out
//
func (v Query) ParseBufferingStats() (mode BufferingModeEnum, avg_in int32, avg_out int32, buffering_left int64) {
	iv, err := _I.Get(858, "Query", "parse_buffering_stats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_avg_in := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_avg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_buffering_left := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_mode, arg_avg_in, arg_avg_out, arg_buffering_left}
	iv.Call(args, nil, &outArgs[0])
	mode = BufferingModeEnum(outArgs[0].Int())
	avg_in = outArgs[1].Int32()
	avg_out = outArgs[2].Int32()
	buffering_left = outArgs[3].Int64()
	return
}

// gst_query_parse_caps
//
// [ filter ] trans: nothing, dir: out
//
func (v Query) ParseCaps() (filter Caps) {
	iv, err := _I.Get(859, "Query", "parse_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_filter := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filter}
	iv.Call(args, nil, &outArgs[0])
	filter.P = outArgs[0].Pointer()
	return
}

// gst_query_parse_caps_result
//
// [ caps ] trans: nothing, dir: out
//
func (v Query) ParseCapsResult() (caps Caps) {
	iv, err := _I.Get(860, "Query", "parse_caps_result")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_caps}
	iv.Call(args, nil, &outArgs[0])
	caps.P = outArgs[0].Pointer()
	return
}

// gst_query_parse_context
//
// [ context ] trans: nothing, dir: out
//
func (v Query) ParseContextF() (context Context) {
	iv, err := _I.Get(861, "Query", "parse_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_context}
	iv.Call(args, nil, &outArgs[0])
	context.P = outArgs[0].Pointer()
	return
}

// gst_query_parse_context_type
//
// [ context_type ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Query) ParseContextType() (result bool, context_type string) {
	iv, err := _I.Get(862, "Query", "parse_context_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_context_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_context_type}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	context_type = outArgs[0].String().Copy()
	result = ret.Bool()
	return
}

// gst_query_parse_convert
//
// [ src_format ] trans: everything, dir: out
//
// [ src_value ] trans: everything, dir: out
//
// [ dest_format ] trans: everything, dir: out
//
// [ dest_value ] trans: everything, dir: out
//
func (v Query) ParseConvert() (src_format FormatEnum, src_value int64, dest_format FormatEnum, dest_value int64) {
	iv, err := _I.Get(863, "Query", "parse_convert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_src_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_dest_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_dest_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_src_format, arg_src_value, arg_dest_format, arg_dest_value}
	iv.Call(args, nil, &outArgs[0])
	src_format = FormatEnum(outArgs[0].Int())
	src_value = outArgs[1].Int64()
	dest_format = FormatEnum(outArgs[2].Int())
	dest_value = outArgs[3].Int64()
	return
}

// gst_query_parse_duration
//
// [ format ] trans: everything, dir: out
//
// [ duration ] trans: everything, dir: out
//
func (v Query) ParseDuration() (format FormatEnum, duration int64) {
	iv, err := _I.Get(864, "Query", "parse_duration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_duration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_format, arg_duration}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	duration = outArgs[1].Int64()
	return
}

// gst_query_parse_latency
//
// [ live ] trans: everything, dir: out
//
// [ min_latency ] trans: everything, dir: out
//
// [ max_latency ] trans: everything, dir: out
//
func (v Query) ParseLatency() (live bool, min_latency uint64, max_latency uint64) {
	iv, err := _I.Get(865, "Query", "parse_latency")
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
	iv.Call(args, nil, &outArgs[0])
	live = outArgs[0].Bool()
	min_latency = outArgs[1].Uint64()
	max_latency = outArgs[2].Uint64()
	return
}

// gst_query_parse_n_formats
//
// [ n_formats ] trans: everything, dir: out
//
func (v Query) ParseNFormats() (n_formats uint32) {
	iv, err := _I.Get(866, "Query", "parse_n_formats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_formats := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_n_formats}
	iv.Call(args, nil, &outArgs[0])
	n_formats = outArgs[0].Uint32()
	return
}

// gst_query_parse_nth_allocation_meta
//
// [ index ] trans: nothing
//
// [ params ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Query) ParseNthAllocationMeta(index uint32) (result gi.GType, params Structure) {
	iv, err := _I.Get(867, "Query", "parse_nth_allocation_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_params := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_index, arg_params}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	params.P = outArgs[0].Pointer()
	result = gi.GType(ret.Uint())
	return
}

// gst_query_parse_nth_allocation_param
//
// [ index ] trans: nothing
//
// [ allocator ] trans: everything, dir: out
//
// [ params ] trans: nothing, dir: out
//
func (v Query) ParseNthAllocationParam(index uint32, params AllocationParams) (allocator Allocator) {
	iv, err := _I.Get(868, "Query", "parse_nth_allocation_param")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_allocator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_index, arg_allocator, arg_params}
	iv.Call(args, nil, &outArgs[0])
	allocator.P = outArgs[0].Pointer()
	return
}

// gst_query_parse_nth_allocation_pool
//
// [ index ] trans: nothing
//
// [ pool ] trans: everything, dir: out
//
// [ size ] trans: everything, dir: out
//
// [ min_buffers ] trans: everything, dir: out
//
// [ max_buffers ] trans: everything, dir: out
//
func (v Query) ParseNthAllocationPool(index uint32) (pool BufferPool, size uint32, min_buffers uint32, max_buffers uint32) {
	iv, err := _I.Get(869, "Query", "parse_nth_allocation_pool")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_pool := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_min_buffers := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_max_buffers := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_index, arg_pool, arg_size, arg_min_buffers, arg_max_buffers}
	iv.Call(args, nil, &outArgs[0])
	pool.P = outArgs[0].Pointer()
	size = outArgs[1].Uint32()
	min_buffers = outArgs[2].Uint32()
	max_buffers = outArgs[3].Uint32()
	return
}

// gst_query_parse_nth_buffering_range
//
// [ index ] trans: nothing
//
// [ start ] trans: everything, dir: out
//
// [ stop ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Query) ParseNthBufferingRange(index uint32) (result bool, start int64, stop int64) {
	iv, err := _I.Get(870, "Query", "parse_nth_buffering_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stop := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_index, arg_start, arg_stop}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start = outArgs[0].Int64()
	stop = outArgs[1].Int64()
	result = ret.Bool()
	return
}

// gst_query_parse_nth_format
//
// [ nth ] trans: nothing
//
// [ format ] trans: everything, dir: out
//
func (v Query) ParseNthFormat(nth uint32) (format FormatEnum) {
	iv, err := _I.Get(871, "Query", "parse_nth_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_nth := gi.NewUint32Argument(nth)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_nth, arg_format}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	return
}

// gst_query_parse_nth_scheduling_mode
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Query) ParseNthSchedulingMode(index uint32) (result PadModeEnum) {
	iv, err := _I.Get(872, "Query", "parse_nth_scheduling_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PadModeEnum(ret.Int())
	return
}

// gst_query_parse_position
//
// [ format ] trans: everything, dir: out
//
// [ cur ] trans: everything, dir: out
//
func (v Query) ParsePosition() (format FormatEnum, cur int64) {
	iv, err := _I.Get(873, "Query", "parse_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cur := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_format, arg_cur}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	cur = outArgs[1].Int64()
	return
}

// gst_query_parse_scheduling
//
// [ flags ] trans: everything, dir: out
//
// [ minsize ] trans: everything, dir: out
//
// [ maxsize ] trans: everything, dir: out
//
// [ align ] trans: everything, dir: out
//
func (v Query) ParseScheduling() (flags SchedulingFlags, minsize int32, maxsize int32, align int32) {
	iv, err := _I.Get(874, "Query", "parse_scheduling")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_minsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_maxsize := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_align := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_flags, arg_minsize, arg_maxsize, arg_align}
	iv.Call(args, nil, &outArgs[0])
	flags = SchedulingFlags(outArgs[0].Int())
	minsize = outArgs[1].Int32()
	maxsize = outArgs[2].Int32()
	align = outArgs[3].Int32()
	return
}

// gst_query_parse_seeking
//
// [ format ] trans: everything, dir: out
//
// [ seekable ] trans: everything, dir: out
//
// [ segment_start ] trans: everything, dir: out
//
// [ segment_end ] trans: everything, dir: out
//
func (v Query) ParseSeeking() (format FormatEnum, seekable bool, segment_start int64, segment_end int64) {
	iv, err := _I.Get(875, "Query", "parse_seeking")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_seekable := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_segment_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_segment_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_format, arg_seekable, arg_segment_start, arg_segment_end}
	iv.Call(args, nil, &outArgs[0])
	format = FormatEnum(outArgs[0].Int())
	seekable = outArgs[1].Bool()
	segment_start = outArgs[2].Int64()
	segment_end = outArgs[3].Int64()
	return
}

// gst_query_parse_segment
//
// [ rate ] trans: everything, dir: out
//
// [ format ] trans: everything, dir: out
//
// [ start_value ] trans: everything, dir: out
//
// [ stop_value ] trans: everything, dir: out
//
func (v Query) ParseSegment() (rate float64, format FormatEnum, start_value int64, stop_value int64) {
	iv, err := _I.Get(876, "Query", "parse_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_rate := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_start_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_stop_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_rate, arg_format, arg_start_value, arg_stop_value}
	iv.Call(args, nil, &outArgs[0])
	rate = outArgs[0].Double()
	format = FormatEnum(outArgs[1].Int())
	start_value = outArgs[2].Int64()
	stop_value = outArgs[3].Int64()
	return
}

// gst_query_parse_uri
//
// [ uri ] trans: everything, dir: out
//
func (v Query) ParseUri() (uri string) {
	iv, err := _I.Get(877, "Query", "parse_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri}
	iv.Call(args, nil, &outArgs[0])
	uri = outArgs[0].String().Take()
	return
}

// gst_query_parse_uri_redirection
//
// [ uri ] trans: everything, dir: out
//
func (v Query) ParseUriRedirection() (uri string) {
	iv, err := _I.Get(878, "Query", "parse_uri_redirection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri}
	iv.Call(args, nil, &outArgs[0])
	uri = outArgs[0].String().Take()
	return
}

// gst_query_parse_uri_redirection_permanent
//
// [ permanent ] trans: everything, dir: out
//
func (v Query) ParseUriRedirectionPermanent() (permanent bool) {
	iv, err := _I.Get(879, "Query", "parse_uri_redirection_permanent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_permanent := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_permanent}
	iv.Call(args, nil, &outArgs[0])
	permanent = outArgs[0].Bool()
	return
}

// gst_query_remove_nth_allocation_meta
//
// [ index ] trans: nothing
//
func (v Query) RemoveNthAllocationMeta(index uint32) {
	iv, err := _I.Get(880, "Query", "remove_nth_allocation_meta")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	iv.Call(args, nil, nil)
}

// gst_query_remove_nth_allocation_param
//
// [ index ] trans: nothing
//
func (v Query) RemoveNthAllocationParam(index uint32) {
	iv, err := _I.Get(881, "Query", "remove_nth_allocation_param")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	iv.Call(args, nil, nil)
}

// gst_query_remove_nth_allocation_pool
//
// [ index ] trans: nothing
//
func (v Query) RemoveNthAllocationPool(index uint32) {
	iv, err := _I.Get(882, "Query", "remove_nth_allocation_pool")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	iv.Call(args, nil, nil)
}

// gst_query_set_accept_caps_result
//
// [ result ] trans: nothing
//
func (v Query) SetAcceptCapsResult(result bool) {
	iv, err := _I.Get(883, "Query", "set_accept_caps_result")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewBoolArgument(result)
	args := []gi.Argument{arg_v, arg_result}
	iv.Call(args, nil, nil)
}

// gst_query_set_buffering_percent
//
// [ busy ] trans: nothing
//
// [ percent ] trans: nothing
//
func (v Query) SetBufferingPercent(busy bool, percent int32) {
	iv, err := _I.Get(884, "Query", "set_buffering_percent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_busy := gi.NewBoolArgument(busy)
	arg_percent := gi.NewInt32Argument(percent)
	args := []gi.Argument{arg_v, arg_busy, arg_percent}
	iv.Call(args, nil, nil)
}

// gst_query_set_buffering_range
//
// [ format ] trans: nothing
//
// [ start ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ estimated_total ] trans: nothing
//
func (v Query) SetBufferingRange(format FormatEnum, start int64, stop int64, estimated_total int64) {
	iv, err := _I.Get(885, "Query", "set_buffering_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_start := gi.NewInt64Argument(start)
	arg_stop := gi.NewInt64Argument(stop)
	arg_estimated_total := gi.NewInt64Argument(estimated_total)
	args := []gi.Argument{arg_v, arg_format, arg_start, arg_stop, arg_estimated_total}
	iv.Call(args, nil, nil)
}

// gst_query_set_buffering_stats
//
// [ mode ] trans: nothing
//
// [ avg_in ] trans: nothing
//
// [ avg_out ] trans: nothing
//
// [ buffering_left ] trans: nothing
//
func (v Query) SetBufferingStats(mode BufferingModeEnum, avg_in int32, avg_out int32, buffering_left int64) {
	iv, err := _I.Get(886, "Query", "set_buffering_stats")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_avg_in := gi.NewInt32Argument(avg_in)
	arg_avg_out := gi.NewInt32Argument(avg_out)
	arg_buffering_left := gi.NewInt64Argument(buffering_left)
	args := []gi.Argument{arg_v, arg_mode, arg_avg_in, arg_avg_out, arg_buffering_left}
	iv.Call(args, nil, nil)
}

// gst_query_set_caps_result
//
// [ caps ] trans: nothing
//
func (v Query) SetCapsResult(caps Caps) {
	iv, err := _I.Get(887, "Query", "set_caps_result")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_caps}
	iv.Call(args, nil, nil)
}

// gst_query_set_context
//
// [ context ] trans: nothing
//
func (v Query) SetContext(context Context) {
	iv, err := _I.Get(888, "Query", "set_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(context.P)
	args := []gi.Argument{arg_v, arg_context}
	iv.Call(args, nil, nil)
}

// gst_query_set_convert
//
// [ src_format ] trans: nothing
//
// [ src_value ] trans: nothing
//
// [ dest_format ] trans: nothing
//
// [ dest_value ] trans: nothing
//
func (v Query) SetConvert(src_format FormatEnum, src_value int64, dest_format FormatEnum, dest_value int64) {
	iv, err := _I.Get(889, "Query", "set_convert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_format := gi.NewIntArgument(int(src_format))
	arg_src_value := gi.NewInt64Argument(src_value)
	arg_dest_format := gi.NewIntArgument(int(dest_format))
	arg_dest_value := gi.NewInt64Argument(dest_value)
	args := []gi.Argument{arg_v, arg_src_format, arg_src_value, arg_dest_format, arg_dest_value}
	iv.Call(args, nil, nil)
}

// gst_query_set_duration
//
// [ format ] trans: nothing
//
// [ duration ] trans: nothing
//
func (v Query) SetDuration(format FormatEnum, duration int64) {
	iv, err := _I.Get(890, "Query", "set_duration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_duration := gi.NewInt64Argument(duration)
	args := []gi.Argument{arg_v, arg_format, arg_duration}
	iv.Call(args, nil, nil)
}

// gst_query_set_formatsv
//
// [ n_formats ] trans: nothing
//
// [ formats ] trans: nothing
//
func (v Query) SetFormatsv(n_formats int32, formats unsafe.Pointer) {
	iv, err := _I.Get(891, "Query", "set_formatsv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_formats := gi.NewInt32Argument(n_formats)
	arg_formats := gi.NewPointerArgument(formats)
	args := []gi.Argument{arg_v, arg_n_formats, arg_formats}
	iv.Call(args, nil, nil)
}

// gst_query_set_latency
//
// [ live ] trans: nothing
//
// [ min_latency ] trans: nothing
//
// [ max_latency ] trans: nothing
//
func (v Query) SetLatency(live bool, min_latency uint64, max_latency uint64) {
	iv, err := _I.Get(892, "Query", "set_latency")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_live := gi.NewBoolArgument(live)
	arg_min_latency := gi.NewUint64Argument(min_latency)
	arg_max_latency := gi.NewUint64Argument(max_latency)
	args := []gi.Argument{arg_v, arg_live, arg_min_latency, arg_max_latency}
	iv.Call(args, nil, nil)
}

// gst_query_set_nth_allocation_param
//
// [ index ] trans: nothing
//
// [ allocator ] trans: nothing
//
// [ params ] trans: nothing
//
func (v Query) SetNthAllocationParam(index uint32, allocator IAllocator, params AllocationParams) {
	iv, err := _I.Get(893, "Query", "set_nth_allocation_param")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if allocator != nil {
		tmp = allocator.P_Allocator()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_allocator := gi.NewPointerArgument(tmp)
	arg_params := gi.NewPointerArgument(params.P)
	args := []gi.Argument{arg_v, arg_index, arg_allocator, arg_params}
	iv.Call(args, nil, nil)
}

// gst_query_set_nth_allocation_pool
//
// [ index ] trans: nothing
//
// [ pool ] trans: nothing
//
// [ size ] trans: nothing
//
// [ min_buffers ] trans: nothing
//
// [ max_buffers ] trans: nothing
//
func (v Query) SetNthAllocationPool(index uint32, pool IBufferPool, size uint32, min_buffers uint32, max_buffers uint32) {
	iv, err := _I.Get(894, "Query", "set_nth_allocation_pool")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pool != nil {
		tmp = pool.P_BufferPool()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	arg_pool := gi.NewPointerArgument(tmp)
	arg_size := gi.NewUint32Argument(size)
	arg_min_buffers := gi.NewUint32Argument(min_buffers)
	arg_max_buffers := gi.NewUint32Argument(max_buffers)
	args := []gi.Argument{arg_v, arg_index, arg_pool, arg_size, arg_min_buffers, arg_max_buffers}
	iv.Call(args, nil, nil)
}

// gst_query_set_position
//
// [ format ] trans: nothing
//
// [ cur ] trans: nothing
//
func (v Query) SetPosition(format FormatEnum, cur int64) {
	iv, err := _I.Get(895, "Query", "set_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_cur := gi.NewInt64Argument(cur)
	args := []gi.Argument{arg_v, arg_format, arg_cur}
	iv.Call(args, nil, nil)
}

// gst_query_set_scheduling
//
// [ flags ] trans: nothing
//
// [ minsize ] trans: nothing
//
// [ maxsize ] trans: nothing
//
// [ align ] trans: nothing
//
func (v Query) SetScheduling(flags SchedulingFlags, minsize int32, maxsize int32, align int32) {
	iv, err := _I.Get(896, "Query", "set_scheduling")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_minsize := gi.NewInt32Argument(minsize)
	arg_maxsize := gi.NewInt32Argument(maxsize)
	arg_align := gi.NewInt32Argument(align)
	args := []gi.Argument{arg_v, arg_flags, arg_minsize, arg_maxsize, arg_align}
	iv.Call(args, nil, nil)
}

// gst_query_set_seeking
//
// [ format ] trans: nothing
//
// [ seekable ] trans: nothing
//
// [ segment_start ] trans: nothing
//
// [ segment_end ] trans: nothing
//
func (v Query) SetSeeking(format FormatEnum, seekable bool, segment_start int64, segment_end int64) {
	iv, err := _I.Get(897, "Query", "set_seeking")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_seekable := gi.NewBoolArgument(seekable)
	arg_segment_start := gi.NewInt64Argument(segment_start)
	arg_segment_end := gi.NewInt64Argument(segment_end)
	args := []gi.Argument{arg_v, arg_format, arg_seekable, arg_segment_start, arg_segment_end}
	iv.Call(args, nil, nil)
}

// gst_query_set_segment
//
// [ rate ] trans: nothing
//
// [ format ] trans: nothing
//
// [ start_value ] trans: nothing
//
// [ stop_value ] trans: nothing
//
func (v Query) SetSegment(rate float64, format FormatEnum, start_value int64, stop_value int64) {
	iv, err := _I.Get(898, "Query", "set_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rate := gi.NewDoubleArgument(rate)
	arg_format := gi.NewIntArgument(int(format))
	arg_start_value := gi.NewInt64Argument(start_value)
	arg_stop_value := gi.NewInt64Argument(stop_value)
	args := []gi.Argument{arg_v, arg_rate, arg_format, arg_start_value, arg_stop_value}
	iv.Call(args, nil, nil)
}

// gst_query_set_uri
//
// [ uri ] trans: nothing
//
func (v Query) SetUri(uri string) {
	iv, err := _I.Get(899, "Query", "set_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	args := []gi.Argument{arg_v, arg_uri}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// gst_query_set_uri_redirection
//
// [ uri ] trans: nothing
//
func (v Query) SetUriRedirection(uri string) {
	iv, err := _I.Get(900, "Query", "set_uri_redirection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	args := []gi.Argument{arg_v, arg_uri}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// gst_query_set_uri_redirection_permanent
//
// [ permanent ] trans: nothing
//
func (v Query) SetUriRedirectionPermanent(permanent bool) {
	iv, err := _I.Get(901, "Query", "set_uri_redirection_permanent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_permanent := gi.NewBoolArgument(permanent)
	args := []gi.Argument{arg_v, arg_permanent}
	iv.Call(args, nil, nil)
}

// gst_query_writable_structure
//
// [ result ] trans: nothing
//
func (v Query) WritableStructure() (result Structure) {
	iv, err := _I.Get(902, "Query", "writable_structure")
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

// Enum QueryType
type QueryTypeEnum int

const (
	QueryTypeUnknown    QueryTypeEnum = 0
	QueryTypePosition   QueryTypeEnum = 2563
	QueryTypeDuration   QueryTypeEnum = 5123
	QueryTypeLatency    QueryTypeEnum = 7683
	QueryTypeJitter     QueryTypeEnum = 10243
	QueryTypeRate       QueryTypeEnum = 12803
	QueryTypeSeeking    QueryTypeEnum = 15363
	QueryTypeSegment    QueryTypeEnum = 17923
	QueryTypeConvert    QueryTypeEnum = 20483
	QueryTypeFormats    QueryTypeEnum = 23043
	QueryTypeBuffering  QueryTypeEnum = 28163
	QueryTypeCustom     QueryTypeEnum = 30723
	QueryTypeUri        QueryTypeEnum = 33283
	QueryTypeAllocation QueryTypeEnum = 35846
	QueryTypeScheduling QueryTypeEnum = 38401
	QueryTypeAcceptCaps QueryTypeEnum = 40963
	QueryTypeCaps       QueryTypeEnum = 43523
	QueryTypeDrain      QueryTypeEnum = 46086
	QueryTypeContext    QueryTypeEnum = 48643
)

func QueryTypeGetType() gi.GType {
	ret := _I.GetGType(133, "QueryType")
	return ret
}

// Flags QueryTypeFlags
type QueryTypeFlags int

const (
	QueryTypeFlagsUpstream   QueryTypeFlags = 1
	QueryTypeFlagsDownstream QueryTypeFlags = 2
	QueryTypeFlagsSerialized QueryTypeFlags = 4
)

func QueryTypeFlagsGetType() gi.GType {
	ret := _I.GetGType(134, "QueryTypeFlags")
	return ret
}

// Enum Rank
type RankEnum int

const (
	RankNone      RankEnum = 0
	RankMarginal  RankEnum = 64
	RankSecondary RankEnum = 128
	RankPrimary   RankEnum = 256
)

func RankGetType() gi.GType {
	ret := _I.GetGType(135, "Rank")
	return ret
}

// Struct ReferenceTimestampMeta
type ReferenceTimestampMeta struct {
	P unsafe.Pointer
}

const SizeOfStructReferenceTimestampMeta = 40

func ReferenceTimestampMetaGetType() gi.GType {
	ret := _I.GetGType(136, "ReferenceTimestampMeta")
	return ret
}

// Object Registry
type Registry struct {
	Object
}

func WrapRegistry(p unsafe.Pointer) (r Registry) { r.P = p; return }

type IRegistry interface{ P_Registry() unsafe.Pointer }

func (v Registry) P_Registry() unsafe.Pointer { return v.P }
func RegistryGetType() gi.GType {
	ret := _I.GetGType(137, "Registry")
	return ret
}

// gst_registry_fork_set_enabled
//
// [ enabled ] trans: nothing
//
func RegistryForkSetEnabled1(enabled bool) {
	iv, err := _I.Get(905, "Registry", "fork_set_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_enabled}
	iv.Call(args, nil, nil)
}

// gst_registry_add_feature
//
// [ feature ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Registry) AddFeature(feature IPluginFeature) (result bool) {
	iv, err := _I.Get(907, "Registry", "add_feature")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if feature != nil {
		tmp = feature.P_PluginFeature()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_feature}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_registry_add_plugin
//
// [ plugin ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Registry) AddPlugin(plugin IPlugin) (result bool) {
	iv, err := _I.Get(908, "Registry", "add_plugin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_plugin := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_plugin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_registry_check_feature_version
//
// [ feature_name ] trans: nothing
//
// [ min_major ] trans: nothing
//
// [ min_minor ] trans: nothing
//
// [ min_micro ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Registry) CheckFeatureVersion(feature_name string, min_major uint32, min_minor uint32, min_micro uint32) (result bool) {
	iv, err := _I.Get(909, "Registry", "check_feature_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_feature_name := gi.CString(feature_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature_name := gi.NewStringArgument(c_feature_name)
	arg_min_major := gi.NewUint32Argument(min_major)
	arg_min_minor := gi.NewUint32Argument(min_minor)
	arg_min_micro := gi.NewUint32Argument(min_micro)
	args := []gi.Argument{arg_v, arg_feature_name, arg_min_major, arg_min_minor, arg_min_micro}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_feature_name)
	result = ret.Bool()
	return
}

// gst_registry_feature_filter
//
// [ filter ] trans: nothing
//
// [ first ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) FeatureFilter(filter int /*TODO_TYPE CALLBACK*/, first bool, user_data unsafe.Pointer) (result g.List) {
	iv, err := _I.Get(910, "Registry", "feature_filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_filter := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPluginFeatureFilter()))
	arg_first := gi.NewBoolArgument(first)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_filter, arg_first, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_registry_find_feature
//
// [ name ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) FindFeature(name string, type1 gi.GType) (result PluginFeature) {
	iv, err := _I.Get(911, "Registry", "find_feature")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_v, arg_name, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_registry_find_plugin
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) FindPlugin(name string) (result Plugin) {
	iv, err := _I.Get(912, "Registry", "find_plugin")
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
	result.P = ret.Pointer()
	return
}

// gst_registry_get_feature_list
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) GetFeatureList(type1 gi.GType) (result g.List) {
	iv, err := _I.Get(913, "Registry", "get_feature_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_registry_get_feature_list_by_plugin
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) GetFeatureListByPlugin(name string) (result g.List) {
	iv, err := _I.Get(914, "Registry", "get_feature_list_by_plugin")
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
	result.P = ret.Pointer()
	return
}

// gst_registry_get_feature_list_cookie
//
// [ result ] trans: nothing
//
func (v Registry) GetFeatureListCookie() (result uint32) {
	iv, err := _I.Get(915, "Registry", "get_feature_list_cookie")
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

// gst_registry_get_plugin_list
//
// [ result ] trans: everything
//
func (v Registry) GetPluginList() (result g.List) {
	iv, err := _I.Get(916, "Registry", "get_plugin_list")
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

// gst_registry_lookup
//
// [ filename ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) Lookup(filename string) (result Plugin) {
	iv, err := _I.Get(917, "Registry", "lookup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_v, arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result.P = ret.Pointer()
	return
}

// gst_registry_lookup_feature
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) LookupFeature(name string) (result PluginFeature) {
	iv, err := _I.Get(918, "Registry", "lookup_feature")
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
	result.P = ret.Pointer()
	return
}

// gst_registry_plugin_filter
//
// [ filter ] trans: nothing
//
// [ first ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: everything
//
func (v Registry) PluginFilter(filter int /*TODO_TYPE CALLBACK*/, first bool, user_data unsafe.Pointer) (result g.List) {
	iv, err := _I.Get(919, "Registry", "plugin_filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_filter := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPluginFilter()))
	arg_first := gi.NewBoolArgument(first)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_filter, arg_first, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_registry_remove_feature
//
// [ feature ] trans: nothing
//
func (v Registry) RemoveFeature(feature IPluginFeature) {
	iv, err := _I.Get(920, "Registry", "remove_feature")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if feature != nil {
		tmp = feature.P_PluginFeature()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_feature := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_feature}
	iv.Call(args, nil, nil)
}

// gst_registry_remove_plugin
//
// [ plugin ] trans: nothing
//
func (v Registry) RemovePlugin(plugin IPlugin) {
	iv, err := _I.Get(921, "Registry", "remove_plugin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_plugin := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_plugin}
	iv.Call(args, nil, nil)
}

// gst_registry_scan_path
//
// [ path ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Registry) ScanPath(path string) (result bool) {
	iv, err := _I.Get(922, "Registry", "scan_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_v, arg_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_path)
	result = ret.Bool()
	return
}

// ignore GType struct RegistryClass

// Struct RegistryPrivate
type RegistryPrivate struct {
	P unsafe.Pointer
}

func RegistryPrivateGetType() gi.GType {
	ret := _I.GetGType(138, "RegistryPrivate")
	return ret
}

// Enum ResourceError
type ResourceErrorEnum int

const (
	ResourceErrorFailed        ResourceErrorEnum = 1
	ResourceErrorTooLazy       ResourceErrorEnum = 2
	ResourceErrorNotFound      ResourceErrorEnum = 3
	ResourceErrorBusy          ResourceErrorEnum = 4
	ResourceErrorOpenRead      ResourceErrorEnum = 5
	ResourceErrorOpenWrite     ResourceErrorEnum = 6
	ResourceErrorOpenReadWrite ResourceErrorEnum = 7
	ResourceErrorClose         ResourceErrorEnum = 8
	ResourceErrorRead          ResourceErrorEnum = 9
	ResourceErrorWrite         ResourceErrorEnum = 10
	ResourceErrorSeek          ResourceErrorEnum = 11
	ResourceErrorSync          ResourceErrorEnum = 12
	ResourceErrorSettings      ResourceErrorEnum = 13
	ResourceErrorNoSpaceLeft   ResourceErrorEnum = 14
	ResourceErrorNotAuthorized ResourceErrorEnum = 15
	ResourceErrorNumErrors     ResourceErrorEnum = 16
)

func ResourceErrorGetType() gi.GType {
	ret := _I.GetGType(139, "ResourceError")
	return ret
}

// Struct Sample
type Sample struct {
	P unsafe.Pointer
}

func SampleGetType() gi.GType {
	ret := _I.GetGType(140, "Sample")
	return ret
}

// gst_sample_new
//
// [ buffer ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ segment ] trans: nothing
//
// [ info ] trans: everything
//
// [ result ] trans: everything
//
func NewSample(buffer Buffer, caps Caps, segment Segment, info Structure) (result Sample) {
	iv, err := _I.Get(923, "Sample", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_segment := gi.NewPointerArgument(segment.P)
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_buffer, arg_caps, arg_segment, arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_sample_get_buffer
//
// [ result ] trans: nothing
//
func (v Sample) GetBuffer() (result Buffer) {
	iv, err := _I.Get(924, "Sample", "get_buffer")
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

// gst_sample_get_buffer_list
//
// [ result ] trans: nothing
//
func (v Sample) GetBufferList() (result BufferList) {
	iv, err := _I.Get(925, "Sample", "get_buffer_list")
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

// gst_sample_get_caps
//
// [ result ] trans: nothing
//
func (v Sample) GetCaps() (result Caps) {
	iv, err := _I.Get(926, "Sample", "get_caps")
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

// gst_sample_get_info
//
// [ result ] trans: nothing
//
func (v Sample) GetInfo() (result Structure) {
	iv, err := _I.Get(927, "Sample", "get_info")
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

// gst_sample_get_segment
//
// [ result ] trans: nothing
//
func (v Sample) GetSegment() (result Segment) {
	iv, err := _I.Get(928, "Sample", "get_segment")
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

// gst_sample_set_buffer_list
//
// [ buffer_list ] trans: nothing
//
func (v Sample) SetBufferList(buffer_list BufferList) {
	iv, err := _I.Get(929, "Sample", "set_buffer_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer_list := gi.NewPointerArgument(buffer_list.P)
	args := []gi.Argument{arg_v, arg_buffer_list}
	iv.Call(args, nil, nil)
}

// Flags SchedulingFlags
type SchedulingFlags int

const (
	SchedulingFlagsSeekable         SchedulingFlags = 1
	SchedulingFlagsSequential       SchedulingFlags = 2
	SchedulingFlagsBandwidthLimited SchedulingFlags = 4
)

func SchedulingFlagsGetType() gi.GType {
	ret := _I.GetGType(141, "SchedulingFlags")
	return ret
}

// Enum SearchMode
type SearchModeEnum int

const (
	SearchModeExact  SearchModeEnum = 0
	SearchModeBefore SearchModeEnum = 1
	SearchModeAfter  SearchModeEnum = 2
)

func SearchModeGetType() gi.GType {
	ret := _I.GetGType(142, "SearchMode")
	return ret
}

// Flags SeekFlags
type SeekFlags int

const (
	SeekFlagsNone              SeekFlags = 0
	SeekFlagsFlush             SeekFlags = 1
	SeekFlagsAccurate          SeekFlags = 2
	SeekFlagsKeyUnit           SeekFlags = 4
	SeekFlagsSegment           SeekFlags = 8
	SeekFlagsTrickmode         SeekFlags = 16
	SeekFlagsSkip              SeekFlags = 16
	SeekFlagsSnapBefore        SeekFlags = 32
	SeekFlagsSnapAfter         SeekFlags = 64
	SeekFlagsSnapNearest       SeekFlags = 96
	SeekFlagsTrickmodeKeyUnits SeekFlags = 128
	SeekFlagsTrickmodeNoAudio  SeekFlags = 256
)

func SeekFlagsGetType() gi.GType {
	ret := _I.GetGType(143, "SeekFlags")
	return ret
}

// Enum SeekType
type SeekTypeEnum int

const (
	SeekTypeNone SeekTypeEnum = 0
	SeekTypeSet  SeekTypeEnum = 1
	SeekTypeEnd  SeekTypeEnum = 2
)

func SeekTypeGetType() gi.GType {
	ret := _I.GetGType(144, "SeekType")
	return ret
}

// Struct Segment
type Segment struct {
	P unsafe.Pointer
}

const SizeOfStructSegment = 120

func SegmentGetType() gi.GType {
	ret := _I.GetGType(145, "Segment")
	return ret
}

// gst_segment_new
//
// [ result ] trans: everything
//
func NewSegment() (result Segment) {
	iv, err := _I.Get(930, "Segment", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_segment_clip
//
// [ format ] trans: nothing
//
// [ start ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ clip_start ] trans: everything, dir: out
//
// [ clip_stop ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Segment) Clip(format FormatEnum, start uint64, stop uint64) (result bool, clip_start uint64, clip_stop uint64) {
	iv, err := _I.Get(931, "Segment", "clip")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_start := gi.NewUint64Argument(start)
	arg_stop := gi.NewUint64Argument(stop)
	arg_clip_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_clip_stop := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_format, arg_start, arg_stop, arg_clip_start, arg_clip_stop}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	clip_start = outArgs[0].Uint64()
	clip_stop = outArgs[1].Uint64()
	result = ret.Bool()
	return
}

// gst_segment_copy
//
// [ result ] trans: everything
//
func (v Segment) Copy() (result Segment) {
	iv, err := _I.Get(932, "Segment", "copy")
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

// gst_segment_copy_into
//
// [ dest ] trans: nothing
//
func (v Segment) CopyInto(dest Segment) {
	iv, err := _I.Get(933, "Segment", "copy_into")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_v, arg_dest}
	iv.Call(args, nil, nil)
}

// gst_segment_do_seek
//
// [ rate ] trans: nothing
//
// [ format ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ start_type ] trans: nothing
//
// [ start ] trans: nothing
//
// [ stop_type ] trans: nothing
//
// [ stop ] trans: nothing
//
// [ update ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Segment) DoSeek(rate float64, format FormatEnum, flags SeekFlags, start_type SeekTypeEnum, start uint64, stop_type SeekTypeEnum, stop uint64) (result bool, update bool) {
	iv, err := _I.Get(934, "Segment", "do_seek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_rate := gi.NewDoubleArgument(rate)
	arg_format := gi.NewIntArgument(int(format))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_start_type := gi.NewIntArgument(int(start_type))
	arg_start := gi.NewUint64Argument(start)
	arg_stop_type := gi.NewIntArgument(int(stop_type))
	arg_stop := gi.NewUint64Argument(stop)
	arg_update := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_rate, arg_format, arg_flags, arg_start_type, arg_start, arg_stop_type, arg_stop, arg_update}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	update = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// gst_segment_free
//
func (v Segment) Free() {
	iv, err := _I.Get(935, "Segment", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_segment_init
//
// [ format ] trans: nothing
//
func (v Segment) Init(format FormatEnum) {
	iv, err := _I.Get(936, "Segment", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_v, arg_format}
	iv.Call(args, nil, nil)
}

// gst_segment_is_equal
//
// [ s1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) IsEqual(s1 Segment) (result bool) {
	iv, err := _I.Get(937, "Segment", "is_equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_s1 := gi.NewPointerArgument(s1.P)
	args := []gi.Argument{arg_v, arg_s1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_segment_offset_running_time
//
// [ format ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) OffsetRunningTime(format FormatEnum, offset int64) (result bool) {
	iv, err := _I.Get(938, "Segment", "offset_running_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_offset := gi.NewInt64Argument(offset)
	args := []gi.Argument{arg_v, arg_format, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_segment_position_from_running_time
//
// [ format ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) PositionFromRunningTime(format FormatEnum, running_time uint64) (result uint64) {
	iv, err := _I.Get(939, "Segment", "position_from_running_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_running_time := gi.NewUint64Argument(running_time)
	args := []gi.Argument{arg_v, arg_format, arg_running_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_segment_position_from_running_time_full
//
// [ format ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ position ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Segment) PositionFromRunningTimeFull(format FormatEnum, running_time uint64) (result int32, position uint64) {
	iv, err := _I.Get(940, "Segment", "position_from_running_time_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_running_time := gi.NewUint64Argument(running_time)
	arg_position := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_running_time, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	position = outArgs[0].Uint64()
	result = ret.Int32()
	return
}

// gst_segment_position_from_stream_time
//
// [ format ] trans: nothing
//
// [ stream_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) PositionFromStreamTime(format FormatEnum, stream_time uint64) (result uint64) {
	iv, err := _I.Get(941, "Segment", "position_from_stream_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_stream_time := gi.NewUint64Argument(stream_time)
	args := []gi.Argument{arg_v, arg_format, arg_stream_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_segment_position_from_stream_time_full
//
// [ format ] trans: nothing
//
// [ stream_time ] trans: nothing
//
// [ position ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Segment) PositionFromStreamTimeFull(format FormatEnum, stream_time uint64) (result int32, position uint64) {
	iv, err := _I.Get(942, "Segment", "position_from_stream_time_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_stream_time := gi.NewUint64Argument(stream_time)
	arg_position := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_stream_time, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	position = outArgs[0].Uint64()
	result = ret.Int32()
	return
}

// gst_segment_set_running_time
//
// [ format ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) SetRunningTime(format FormatEnum, running_time uint64) (result bool) {
	iv, err := _I.Get(943, "Segment", "set_running_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_running_time := gi.NewUint64Argument(running_time)
	args := []gi.Argument{arg_v, arg_format, arg_running_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// gst_segment_to_position
//
// [ format ] trans: nothing
//
// [ running_time ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) ToPosition(format FormatEnum, running_time uint64) (result uint64) {
	iv, err := _I.Get(944, "Segment", "to_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_running_time := gi.NewUint64Argument(running_time)
	args := []gi.Argument{arg_v, arg_format, arg_running_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_segment_to_running_time
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) ToRunningTime(format FormatEnum, position uint64) (result uint64) {
	iv, err := _I.Get(945, "Segment", "to_running_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewUint64Argument(position)
	args := []gi.Argument{arg_v, arg_format, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_segment_to_running_time_full
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ running_time ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Segment) ToRunningTimeFull(format FormatEnum, position uint64) (result int32, running_time uint64) {
	iv, err := _I.Get(946, "Segment", "to_running_time_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewUint64Argument(position)
	arg_running_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_position, arg_running_time}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	running_time = outArgs[0].Uint64()
	result = ret.Int32()
	return
}

// gst_segment_to_stream_time
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Segment) ToStreamTime(format FormatEnum, position uint64) (result uint64) {
	iv, err := _I.Get(947, "Segment", "to_stream_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewUint64Argument(position)
	args := []gi.Argument{arg_v, arg_format, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_segment_to_stream_time_full
//
// [ format ] trans: nothing
//
// [ position ] trans: nothing
//
// [ stream_time ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Segment) ToStreamTimeFull(format FormatEnum, position uint64) (result int32, stream_time uint64) {
	iv, err := _I.Get(948, "Segment", "to_stream_time_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewIntArgument(int(format))
	arg_position := gi.NewUint64Argument(position)
	arg_stream_time := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_format, arg_position, arg_stream_time}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	stream_time = outArgs[0].Uint64()
	result = ret.Int32()
	return
}

// Flags SegmentFlags
type SegmentFlags int

const (
	SegmentFlagsNone              SegmentFlags = 0
	SegmentFlagsReset             SegmentFlags = 1
	SegmentFlagsTrickmode         SegmentFlags = 16
	SegmentFlagsSkip              SegmentFlags = 16
	SegmentFlagsSegment           SegmentFlags = 8
	SegmentFlagsTrickmodeKeyUnits SegmentFlags = 128
	SegmentFlagsTrickmodeNoAudio  SegmentFlags = 256
)

func SegmentFlagsGetType() gi.GType {
	ret := _I.GetGType(146, "SegmentFlags")
	return ret
}

// Flags StackTraceFlags
type StackTraceFlags int

const (
	StackTraceFlagsFull StackTraceFlags = 1
)

func StackTraceFlagsGetType() gi.GType {
	ret := _I.GetGType(147, "StackTraceFlags")
	return ret
}

// Enum State
type StateEnum int

const (
	StateVoidPending StateEnum = 0
	StateNull        StateEnum = 1
	StateReady       StateEnum = 2
	StatePaused      StateEnum = 3
	StatePlaying     StateEnum = 4
)

func StateGetType() gi.GType {
	ret := _I.GetGType(148, "State")
	return ret
}

// Enum StateChange
type StateChangeEnum int

const (
	StateChangeNullToReady      StateChangeEnum = 10
	StateChangeReadyToPaused    StateChangeEnum = 19
	StateChangePausedToPlaying  StateChangeEnum = 28
	StateChangePlayingToPaused  StateChangeEnum = 35
	StateChangePausedToReady    StateChangeEnum = 26
	StateChangeReadyToNull      StateChangeEnum = 17
	StateChangeNullToNull       StateChangeEnum = 9
	StateChangeReadyToReady     StateChangeEnum = 18
	StateChangePausedToPaused   StateChangeEnum = 27
	StateChangePlayingToPlaying StateChangeEnum = 36
)

func StateChangeGetType() gi.GType {
	ret := _I.GetGType(149, "StateChange")
	return ret
}

// Enum StateChangeReturn
type StateChangeReturnEnum int

const (
	StateChangeReturnFailure   StateChangeReturnEnum = 0
	StateChangeReturnSuccess   StateChangeReturnEnum = 1
	StateChangeReturnAsync     StateChangeReturnEnum = 2
	StateChangeReturnNoPreroll StateChangeReturnEnum = 3
)

func StateChangeReturnGetType() gi.GType {
	ret := _I.GetGType(150, "StateChangeReturn")
	return ret
}

// Struct StaticCaps
type StaticCaps struct {
	P unsafe.Pointer
}

const SizeOfStructStaticCaps = 48

func StaticCapsGetType() gi.GType {
	ret := _I.GetGType(151, "StaticCaps")
	return ret
}

// gst_static_caps_cleanup
//
func (v StaticCaps) Cleanup() {
	iv, err := _I.Get(949, "StaticCaps", "cleanup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_static_caps_get
//
// [ result ] trans: everything
//
func (v StaticCaps) Get() (result Caps) {
	iv, err := _I.Get(950, "StaticCaps", "get")
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

// Struct StaticPadTemplate
type StaticPadTemplate struct {
	P unsafe.Pointer
}

const SizeOfStructStaticPadTemplate = 64

func StaticPadTemplateGetType() gi.GType {
	ret := _I.GetGType(152, "StaticPadTemplate")
	return ret
}

// gst_static_pad_template_get
//
// [ result ] trans: nothing
//
func (v StaticPadTemplate) Get() (result PadTemplate) {
	iv, err := _I.Get(951, "StaticPadTemplate", "get")
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

// gst_static_pad_template_get_caps
//
// [ result ] trans: everything
//
func (v StaticPadTemplate) GetCaps() (result Caps) {
	iv, err := _I.Get(952, "StaticPadTemplate", "get_caps")
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

// Object Stream
type Stream struct {
	Object
}

func WrapStream(p unsafe.Pointer) (r Stream) { r.P = p; return }

type IStream interface{ P_Stream() unsafe.Pointer }

func (v Stream) P_Stream() unsafe.Pointer { return v.P }
func StreamGetType() gi.GType {
	ret := _I.GetGType(153, "Stream")
	return ret
}

// gst_stream_new
//
// [ stream_id ] trans: nothing
//
// [ caps ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func NewStream(stream_id string, caps Caps, type1 StreamTypeFlags, flags StreamFlags) (result Stream) {
	iv, err := _I.Get(953, "Stream", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_stream_id := gi.CString(stream_id)
	arg_stream_id := gi.NewStringArgument(c_stream_id)
	arg_caps := gi.NewPointerArgument(caps.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_stream_id, arg_caps, arg_type1, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_stream_id)
	result.P = ret.Pointer()
	return
}

// gst_stream_get_caps
//
// [ result ] trans: everything
//
func (v Stream) GetCaps() (result Caps) {
	iv, err := _I.Get(954, "Stream", "get_caps")
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

// gst_stream_get_stream_flags
//
// [ result ] trans: nothing
//
func (v Stream) GetStreamFlags() (result StreamFlags) {
	iv, err := _I.Get(955, "Stream", "get_stream_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StreamFlags(ret.Int())
	return
}

// gst_stream_get_stream_id
//
// [ result ] trans: nothing
//
func (v Stream) GetStreamId() (result string) {
	iv, err := _I.Get(956, "Stream", "get_stream_id")
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

// gst_stream_get_stream_type
//
// [ result ] trans: nothing
//
func (v Stream) GetStreamType() (result StreamTypeFlags) {
	iv, err := _I.Get(957, "Stream", "get_stream_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StreamTypeFlags(ret.Int())
	return
}

// gst_stream_get_tags
//
// [ result ] trans: everything
//
func (v Stream) GetTags() (result TagList) {
	iv, err := _I.Get(958, "Stream", "get_tags")
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

// gst_stream_set_caps
//
// [ caps ] trans: nothing
//
func (v Stream) SetCaps(caps Caps) {
	iv, err := _I.Get(959, "Stream", "set_caps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_caps}
	iv.Call(args, nil, nil)
}

// gst_stream_set_stream_flags
//
// [ flags ] trans: nothing
//
func (v Stream) SetStreamFlags(flags StreamFlags) {
	iv, err := _I.Get(960, "Stream", "set_stream_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, nil)
}

// gst_stream_set_stream_type
//
// [ stream_type ] trans: nothing
//
func (v Stream) SetStreamType(stream_type StreamTypeFlags) {
	iv, err := _I.Get(961, "Stream", "set_stream_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream_type := gi.NewIntArgument(int(stream_type))
	args := []gi.Argument{arg_v, arg_stream_type}
	iv.Call(args, nil, nil)
}

// gst_stream_set_tags
//
// [ tags ] trans: nothing
//
func (v Stream) SetTags(tags TagList) {
	iv, err := _I.Get(962, "Stream", "set_tags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tags := gi.NewPointerArgument(tags.P)
	args := []gi.Argument{arg_v, arg_tags}
	iv.Call(args, nil, nil)
}

// ignore GType struct StreamClass

// Object StreamCollection
type StreamCollection struct {
	Object
}

func WrapStreamCollection(p unsafe.Pointer) (r StreamCollection) { r.P = p; return }

type IStreamCollection interface{ P_StreamCollection() unsafe.Pointer }

func (v StreamCollection) P_StreamCollection() unsafe.Pointer { return v.P }
func StreamCollectionGetType() gi.GType {
	ret := _I.GetGType(154, "StreamCollection")
	return ret
}

// gst_stream_collection_new
//
// [ upstream_id ] trans: nothing
//
// [ result ] trans: everything
//
func NewStreamCollection(upstream_id string) (result StreamCollection) {
	iv, err := _I.Get(963, "StreamCollection", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_upstream_id := gi.CString(upstream_id)
	arg_upstream_id := gi.NewStringArgument(c_upstream_id)
	args := []gi.Argument{arg_upstream_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_upstream_id)
	result.P = ret.Pointer()
	return
}

// gst_stream_collection_add_stream
//
// [ stream ] trans: everything
//
// [ result ] trans: nothing
//
func (v StreamCollection) AddStream(stream IStream) (result bool) {
	iv, err := _I.Get(964, "StreamCollection", "add_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if stream != nil {
		tmp = stream.P_Stream()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_stream}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_stream_collection_get_size
//
// [ result ] trans: nothing
//
func (v StreamCollection) GetSize() (result uint32) {
	iv, err := _I.Get(965, "StreamCollection", "get_size")
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

// gst_stream_collection_get_stream
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StreamCollection) GetStream(index uint32) (result Stream) {
	iv, err := _I.Get(966, "StreamCollection", "get_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_stream_collection_get_upstream_id
//
// [ result ] trans: nothing
//
func (v StreamCollection) GetUpstreamId() (result string) {
	iv, err := _I.Get(967, "StreamCollection", "get_upstream_id")
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

// ignore GType struct StreamCollectionClass

// Struct StreamCollectionPrivate
type StreamCollectionPrivate struct {
	P unsafe.Pointer
}

func StreamCollectionPrivateGetType() gi.GType {
	ret := _I.GetGType(155, "StreamCollectionPrivate")
	return ret
}

// Enum StreamError
type StreamErrorEnum int

const (
	StreamErrorFailed         StreamErrorEnum = 1
	StreamErrorTooLazy        StreamErrorEnum = 2
	StreamErrorNotImplemented StreamErrorEnum = 3
	StreamErrorTypeNotFound   StreamErrorEnum = 4
	StreamErrorWrongType      StreamErrorEnum = 5
	StreamErrorCodecNotFound  StreamErrorEnum = 6
	StreamErrorDecode         StreamErrorEnum = 7
	StreamErrorEncode         StreamErrorEnum = 8
	StreamErrorDemux          StreamErrorEnum = 9
	StreamErrorMux            StreamErrorEnum = 10
	StreamErrorFormat         StreamErrorEnum = 11
	StreamErrorDecrypt        StreamErrorEnum = 12
	StreamErrorDecryptNokey   StreamErrorEnum = 13
	StreamErrorNumErrors      StreamErrorEnum = 14
)

func StreamErrorGetType() gi.GType {
	ret := _I.GetGType(156, "StreamError")
	return ret
}

// Flags StreamFlags
type StreamFlags int

const (
	StreamFlagsNone     StreamFlags = 0
	StreamFlagsSparse   StreamFlags = 1
	StreamFlagsSelect   StreamFlags = 2
	StreamFlagsUnselect StreamFlags = 4
)

func StreamFlagsGetType() gi.GType {
	ret := _I.GetGType(157, "StreamFlags")
	return ret
}

// Struct StreamPrivate
type StreamPrivate struct {
	P unsafe.Pointer
}

func StreamPrivateGetType() gi.GType {
	ret := _I.GetGType(158, "StreamPrivate")
	return ret
}

// Enum StreamStatusType
type StreamStatusTypeEnum int

const (
	StreamStatusTypeCreate  StreamStatusTypeEnum = 0
	StreamStatusTypeEnter   StreamStatusTypeEnum = 1
	StreamStatusTypeLeave   StreamStatusTypeEnum = 2
	StreamStatusTypeDestroy StreamStatusTypeEnum = 3
	StreamStatusTypeStart   StreamStatusTypeEnum = 8
	StreamStatusTypePause   StreamStatusTypeEnum = 9
	StreamStatusTypeStop    StreamStatusTypeEnum = 10
)

func StreamStatusTypeGetType() gi.GType {
	ret := _I.GetGType(159, "StreamStatusType")
	return ret
}

// Flags StreamType
type StreamTypeFlags int

const (
	StreamTypeUnknown   StreamTypeFlags = 1
	StreamTypeAudio     StreamTypeFlags = 2
	StreamTypeVideo     StreamTypeFlags = 4
	StreamTypeContainer StreamTypeFlags = 8
	StreamTypeText      StreamTypeFlags = 16
)

func StreamTypeGetType() gi.GType {
	ret := _I.GetGType(160, "StreamType")
	return ret
}

// Struct Structure
type Structure struct {
	P unsafe.Pointer
}

const SizeOfStructStructure = 16

func StructureGetType() gi.GType {
	ret := _I.GetGType(161, "Structure")
	return ret
}

// gst_structure_new_empty
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func NewStructureEmpty(name string) (result Structure) {
	iv, err := _I.Get(968, "Structure", "new_empty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gst_structure_new_from_string
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func NewStructureFromString(string string) (result Structure) {
	iv, err := _I.Get(969, "Structure", "new_from_string")
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

// gst_structure_new_id_empty
//
// [ quark ] trans: nothing
//
// [ result ] trans: everything
//
func NewStructureIdEmpty(quark uint32) (result Structure) {
	iv, err := _I.Get(970, "Structure", "new_id_empty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_structure_can_intersect
//
// [ struct2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) CanIntersect(struct2 Structure) (result bool) {
	iv, err := _I.Get(971, "Structure", "can_intersect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_struct2 := gi.NewPointerArgument(struct2.P)
	args := []gi.Argument{arg_v, arg_struct2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_copy
//
// [ result ] trans: everything
//
func (v Structure) Copy() (result Structure) {
	iv, err := _I.Get(972, "Structure", "copy")
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

// gst_structure_filter_and_map_in_place
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v Structure) FilterAndMapInPlace(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(973, "Structure", "filter_and_map_in_place")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myStructureFilterMapFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_structure_fixate
//
func (v Structure) Fixate() {
	iv, err := _I.Get(974, "Structure", "fixate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_structure_fixate_field
//
// [ field_name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) FixateField(field_name string) (result bool) {
	iv, err := _I.Get(975, "Structure", "fixate_field")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field_name := gi.CString(field_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field_name := gi.NewStringArgument(c_field_name)
	args := []gi.Argument{arg_v, arg_field_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_field_name)
	result = ret.Bool()
	return
}

// gst_structure_fixate_field_boolean
//
// [ field_name ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) FixateFieldBoolean(field_name string, target bool) (result bool) {
	iv, err := _I.Get(976, "Structure", "fixate_field_boolean")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field_name := gi.CString(field_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field_name := gi.NewStringArgument(c_field_name)
	arg_target := gi.NewBoolArgument(target)
	args := []gi.Argument{arg_v, arg_field_name, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_field_name)
	result = ret.Bool()
	return
}

// gst_structure_fixate_field_nearest_double
//
// [ field_name ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) FixateFieldNearestDouble(field_name string, target float64) (result bool) {
	iv, err := _I.Get(977, "Structure", "fixate_field_nearest_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field_name := gi.CString(field_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field_name := gi.NewStringArgument(c_field_name)
	arg_target := gi.NewDoubleArgument(target)
	args := []gi.Argument{arg_v, arg_field_name, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_field_name)
	result = ret.Bool()
	return
}

// gst_structure_fixate_field_nearest_fraction
//
// [ field_name ] trans: nothing
//
// [ target_numerator ] trans: nothing
//
// [ target_denominator ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) FixateFieldNearestFraction(field_name string, target_numerator int32, target_denominator int32) (result bool) {
	iv, err := _I.Get(978, "Structure", "fixate_field_nearest_fraction")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field_name := gi.CString(field_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field_name := gi.NewStringArgument(c_field_name)
	arg_target_numerator := gi.NewInt32Argument(target_numerator)
	arg_target_denominator := gi.NewInt32Argument(target_denominator)
	args := []gi.Argument{arg_v, arg_field_name, arg_target_numerator, arg_target_denominator}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_field_name)
	result = ret.Bool()
	return
}

// gst_structure_fixate_field_nearest_int
//
// [ field_name ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) FixateFieldNearestInt(field_name string, target int32) (result bool) {
	iv, err := _I.Get(979, "Structure", "fixate_field_nearest_int")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field_name := gi.CString(field_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field_name := gi.NewStringArgument(c_field_name)
	arg_target := gi.NewInt32Argument(target)
	args := []gi.Argument{arg_v, arg_field_name, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_field_name)
	result = ret.Bool()
	return
}

// gst_structure_fixate_field_string
//
// [ field_name ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) FixateFieldString(field_name string, target string) (result bool) {
	iv, err := _I.Get(980, "Structure", "fixate_field_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_field_name := gi.CString(field_name)
	c_target := gi.CString(target)
	arg_v := gi.NewPointerArgument(v.P)
	arg_field_name := gi.NewStringArgument(c_field_name)
	arg_target := gi.NewStringArgument(c_target)
	args := []gi.Argument{arg_v, arg_field_name, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_field_name)
	gi.Free(c_target)
	result = ret.Bool()
	return
}

// gst_structure_foreach
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) Foreach(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(981, "Structure", "foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myStructureForeachFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_free
//
func (v Structure) Free() {
	iv, err := _I.Get(982, "Structure", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_structure_get_array
//
// [ fieldname ] trans: nothing
//
// [ array ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetArray(fieldname string) (result bool, array g.ValueArray) {
	iv, err := _I.Get(983, "Structure", "get_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_array := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	array.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_structure_get_boolean
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetBoolean(fieldname string) (result bool, value bool) {
	iv, err := _I.Get(984, "Structure", "get_boolean")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// gst_structure_get_clock_time
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetClockTime(fieldname string) (result bool, value uint64) {
	iv, err := _I.Get(985, "Structure", "get_clock_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_structure_get_date
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetDate(fieldname string) (result bool, value g.Date) {
	iv, err := _I.Get(986, "Structure", "get_date")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_structure_get_date_time
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetDateTime(fieldname string) (result bool, value DateTime) {
	iv, err := _I.Get(987, "Structure", "get_date_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_structure_get_double
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetDouble(fieldname string) (result bool, value float64) {
	iv, err := _I.Get(988, "Structure", "get_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_structure_get_enum
//
// [ fieldname ] trans: nothing
//
// [ enumtype ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetEnum(fieldname string, enumtype gi.GType) (result bool, value int32) {
	iv, err := _I.Get(989, "Structure", "get_enum")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_enumtype := gi.NewUintArgument(uint(enumtype))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_enumtype, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_structure_get_field_type
//
// [ fieldname ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) GetFieldType(fieldname string) (result gi.GType) {
	iv, err := _I.Get(990, "Structure", "get_field_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	args := []gi.Argument{arg_v, arg_fieldname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_fieldname)
	result = gi.GType(ret.Uint())
	return
}

// gst_structure_get_flagset
//
// [ fieldname ] trans: nothing
//
// [ value_flags ] trans: everything, dir: out
//
// [ value_mask ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetFlagset(fieldname string) (result bool, value_flags uint32, value_mask uint32) {
	iv, err := _I.Get(991, "Structure", "get_flagset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value_flags := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value_flags, arg_value_mask}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value_flags = outArgs[0].Uint32()
	value_mask = outArgs[1].Uint32()
	result = ret.Bool()
	return
}

// gst_structure_get_fraction
//
// [ fieldname ] trans: nothing
//
// [ value_numerator ] trans: everything, dir: out
//
// [ value_denominator ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetFraction(fieldname string) (result bool, value_numerator int32, value_denominator int32) {
	iv, err := _I.Get(992, "Structure", "get_fraction")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value_numerator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value_denominator := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value_numerator, arg_value_denominator}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value_numerator = outArgs[0].Int32()
	value_denominator = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// gst_structure_get_int
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetInt(fieldname string) (result bool, value int32) {
	iv, err := _I.Get(993, "Structure", "get_int")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_structure_get_int64
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetInt64(fieldname string) (result bool, value int64) {
	iv, err := _I.Get(994, "Structure", "get_int64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_structure_get_list
//
// [ fieldname ] trans: nothing
//
// [ array ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetList(fieldname string) (result bool, array g.ValueArray) {
	iv, err := _I.Get(995, "Structure", "get_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_array := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	array.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_structure_get_name
//
// [ result ] trans: nothing
//
func (v Structure) GetName() (result string) {
	iv, err := _I.Get(996, "Structure", "get_name")
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

// gst_structure_get_name_id
//
// [ result ] trans: nothing
//
func (v Structure) GetNameId() (result uint32) {
	iv, err := _I.Get(997, "Structure", "get_name_id")
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

// gst_structure_get_string
//
// [ fieldname ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) GetString(fieldname string) (result string) {
	iv, err := _I.Get(998, "Structure", "get_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	args := []gi.Argument{arg_v, arg_fieldname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_fieldname)
	result = ret.String().Copy()
	return
}

// gst_structure_get_uint
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetUint(fieldname string) (result bool, value uint32) {
	iv, err := _I.Get(999, "Structure", "get_uint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_structure_get_uint64
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Structure) GetUint64(fieldname string) (result bool, value uint64) {
	iv, err := _I.Get(1000, "Structure", "get_uint64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_fieldname)
	value = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_structure_get_value
//
// [ fieldname ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) GetValue(fieldname string) (result g.Value) {
	iv, err := _I.Get(1001, "Structure", "get_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	args := []gi.Argument{arg_v, arg_fieldname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_fieldname)
	result.P = ret.Pointer()
	return
}

// gst_structure_has_field
//
// [ fieldname ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) HasField(fieldname string) (result bool) {
	iv, err := _I.Get(1002, "Structure", "has_field")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	args := []gi.Argument{arg_v, arg_fieldname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_fieldname)
	result = ret.Bool()
	return
}

// gst_structure_has_field_typed
//
// [ fieldname ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) HasFieldTyped(fieldname string, type1 gi.GType) (result bool) {
	iv, err := _I.Get(1003, "Structure", "has_field_typed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_v, arg_fieldname, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_fieldname)
	result = ret.Bool()
	return
}

// gst_structure_has_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) HasName(name string) (result bool) {
	iv, err := _I.Get(1004, "Structure", "has_name")
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
	result = ret.Bool()
	return
}

// gst_structure_id_get_value
//
// [ field ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) IdGetValue(field uint32) (result g.Value) {
	iv, err := _I.Get(1005, "Structure", "id_get_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_field := gi.NewUint32Argument(field)
	args := []gi.Argument{arg_v, arg_field}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_structure_id_has_field
//
// [ field ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) IdHasField(field uint32) (result bool) {
	iv, err := _I.Get(1006, "Structure", "id_has_field")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_field := gi.NewUint32Argument(field)
	args := []gi.Argument{arg_v, arg_field}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_id_has_field_typed
//
// [ field ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) IdHasFieldTyped(field uint32, type1 gi.GType) (result bool) {
	iv, err := _I.Get(1007, "Structure", "id_has_field_typed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_field := gi.NewUint32Argument(field)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_v, arg_field, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_id_set_value
//
// [ field ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Structure) IdSetValue(field uint32, value g.Value) {
	iv, err := _I.Get(1008, "Structure", "id_set_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_field := gi.NewUint32Argument(field)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_field, arg_value}
	iv.Call(args, nil, nil)
}

// gst_structure_id_take_value
//
// [ field ] trans: nothing
//
// [ value ] trans: everything
//
func (v Structure) IdTakeValue(field uint32, value g.Value) {
	iv, err := _I.Get(1009, "Structure", "id_take_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_field := gi.NewUint32Argument(field)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_field, arg_value}
	iv.Call(args, nil, nil)
}

// gst_structure_intersect
//
// [ struct2 ] trans: nothing
//
// [ result ] trans: everything
//
func (v Structure) Intersect(struct2 Structure) (result Structure) {
	iv, err := _I.Get(1010, "Structure", "intersect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_struct2 := gi.NewPointerArgument(struct2.P)
	args := []gi.Argument{arg_v, arg_struct2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_structure_is_equal
//
// [ structure2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) IsEqual(structure2 Structure) (result bool) {
	iv, err := _I.Get(1011, "Structure", "is_equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_structure2 := gi.NewPointerArgument(structure2.P)
	args := []gi.Argument{arg_v, arg_structure2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_is_subset
//
// [ superset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) IsSubset(superset Structure) (result bool) {
	iv, err := _I.Get(1012, "Structure", "is_subset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_superset := gi.NewPointerArgument(superset.P)
	args := []gi.Argument{arg_v, arg_superset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_map_in_place
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) MapInPlace(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1013, "Structure", "map_in_place")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myStructureMapFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_n_fields
//
// [ result ] trans: nothing
//
func (v Structure) NFields() (result int32) {
	iv, err := _I.Get(1014, "Structure", "n_fields")
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

// gst_structure_nth_field_name
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) NthFieldName(index uint32) (result string) {
	iv, err := _I.Get(1015, "Structure", "nth_field_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_structure_remove_all_fields
//
func (v Structure) RemoveAllFields() {
	iv, err := _I.Get(1016, "Structure", "remove_all_fields")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_structure_remove_field
//
// [ fieldname ] trans: nothing
//
func (v Structure) RemoveField(fieldname string) {
	iv, err := _I.Get(1017, "Structure", "remove_field")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	args := []gi.Argument{arg_v, arg_fieldname}
	iv.Call(args, nil, nil)
	gi.Free(c_fieldname)
}

// gst_structure_set_array
//
// [ fieldname ] trans: nothing
//
// [ array ] trans: nothing
//
func (v Structure) SetArray(fieldname string, array g.ValueArray) {
	iv, err := _I.Get(1018, "Structure", "set_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_v, arg_fieldname, arg_array}
	iv.Call(args, nil, nil)
	gi.Free(c_fieldname)
}

// gst_structure_set_list
//
// [ fieldname ] trans: nothing
//
// [ array ] trans: nothing
//
func (v Structure) SetList(fieldname string, array g.ValueArray) {
	iv, err := _I.Get(1019, "Structure", "set_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_v, arg_fieldname, arg_array}
	iv.Call(args, nil, nil)
	gi.Free(c_fieldname)
}

// gst_structure_set_name
//
// [ name ] trans: nothing
//
func (v Structure) SetName(name string) {
	iv, err := _I.Get(1020, "Structure", "set_name")
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

// gst_structure_set_parent_refcount
//
// [ refcount ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Structure) SetParentRefcount(refcount int32) (result bool) {
	iv, err := _I.Get(1021, "Structure", "set_parent_refcount")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_refcount := gi.NewInt32Argument(refcount)
	args := []gi.Argument{arg_v, arg_refcount}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_structure_set_value
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Structure) SetValue(fieldname string, value g.Value) {
	iv, err := _I.Get(1022, "Structure", "set_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_fieldname)
}

// gst_structure_take_value
//
// [ fieldname ] trans: nothing
//
// [ value ] trans: everything
//
func (v Structure) TakeValue(fieldname string, value g.Value) {
	iv, err := _I.Get(1023, "Structure", "take_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fieldname := gi.CString(fieldname)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fieldname := gi.NewStringArgument(c_fieldname)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_fieldname, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_fieldname)
}

// gst_structure_to_string
//
// [ result ] trans: everything
//
func (v Structure) ToString() (result string) {
	iv, err := _I.Get(1024, "Structure", "to_string")
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

// gst_structure_from_string
//
// [ string ] trans: nothing
//
// [ end ] trans: nothing, dir: out
//
// [ result ] trans: everything
//
func StructureFromString1(string string) (result Structure, end string) {
	iv, err := _I.Get(1025, "Structure", "from_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	end = outArgs[0].String().Copy()
	result.P = ret.Pointer()
	return
}

// Enum StructureChangeType
type StructureChangeTypeEnum int

const (
	StructureChangeTypeLink   StructureChangeTypeEnum = 0
	StructureChangeTypeUnlink StructureChangeTypeEnum = 1
)

func StructureChangeTypeGetType() gi.GType {
	ret := _I.GetGType(162, "StructureChangeType")
	return ret
}

type StructureFilterMapFuncStruct struct {
	F_field_id uint32
	F_value    g.Value
}

func GetPointer_myStructureFilterMapFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstStructureFilterMapFunc())
}

//export myGstStructureFilterMapFunc
func myGstStructureFilterMapFunc(field_id C.guint32, value *C.GValue, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &StructureFilterMapFuncStruct{
		F_field_id: uint32(field_id),
		F_value:    g.Value{P: unsafe.Pointer(value)},
	}
	fn(args)
}

type StructureForeachFuncStruct struct {
	F_field_id uint32
	F_value    g.Value
}

func GetPointer_myStructureForeachFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstStructureForeachFunc())
}

//export myGstStructureForeachFunc
func myGstStructureForeachFunc(field_id C.guint32, value *C.GValue, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &StructureForeachFuncStruct{
		F_field_id: uint32(field_id),
		F_value:    g.Value{P: unsafe.Pointer(value)},
	}
	fn(args)
}

type StructureMapFuncStruct struct {
	F_field_id uint32
	F_value    g.Value
}

func GetPointer_myStructureMapFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstStructureMapFunc())
}

//export myGstStructureMapFunc
func myGstStructureMapFunc(field_id C.guint32, value *C.GValue, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &StructureMapFuncStruct{
		F_field_id: uint32(field_id),
		F_value:    g.Value{P: unsafe.Pointer(value)},
	}
	fn(args)
}

// Object SystemClock
type SystemClock struct {
	Clock
}

func WrapSystemClock(p unsafe.Pointer) (r SystemClock) { r.P = p; return }

type ISystemClock interface{ P_SystemClock() unsafe.Pointer }

func (v SystemClock) P_SystemClock() unsafe.Pointer { return v.P }
func SystemClockGetType() gi.GType {
	ret := _I.GetGType(163, "SystemClock")
	return ret
}

// gst_system_clock_set_default
//
// [ new_clock ] trans: nothing
//
func SystemClockSetDefault1(new_clock IClock) {
	iv, err := _I.Get(1027, "SystemClock", "set_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if new_clock != nil {
		tmp = new_clock.P_Clock()
	}
	arg_new_clock := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_new_clock}
	iv.Call(args, nil, nil)
}

// ignore GType struct SystemClockClass

// Struct SystemClockPrivate
type SystemClockPrivate struct {
	P unsafe.Pointer
}

func SystemClockPrivateGetType() gi.GType {
	ret := _I.GetGType(164, "SystemClockPrivate")
	return ret
}

// Enum TagFlag
type TagFlagEnum int

const (
	TagFlagUndefined TagFlagEnum = 0
	TagFlagMeta      TagFlagEnum = 1
	TagFlagEncoded   TagFlagEnum = 2
	TagFlagDecoded   TagFlagEnum = 3
	TagFlagCount     TagFlagEnum = 4
)

func TagFlagGetType() gi.GType {
	ret := _I.GetGType(165, "TagFlag")
	return ret
}

type TagForeachFuncStruct struct {
	F_list TagList
	F_tag  string
}

func GetPointer_myTagForeachFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstTagForeachFunc())
}

//export myGstTagForeachFunc
func myGstTagForeachFunc(list *C.GstTagList, tag *C.gchar, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TagForeachFuncStruct{
		F_list: TagList{P: unsafe.Pointer(list)},
		F_tag:  gi.GoString(unsafe.Pointer(tag)),
	}
	fn(args)
}

// Struct TagList
type TagList struct {
	P unsafe.Pointer
}

const SizeOfStructTagList = 64

func TagListGetType() gi.GType {
	ret := _I.GetGType(166, "TagList")
	return ret
}

// gst_tag_list_new_empty
//
// [ result ] trans: everything
//
func NewTagListEmpty() (result TagList) {
	iv, err := _I.Get(1028, "TagList", "new_empty")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_tag_list_new_from_string
//
// [ str ] trans: nothing
//
// [ result ] trans: everything
//
func NewTagListFromString(str string) (result TagList) {
	iv, err := _I.Get(1029, "TagList", "new_from_string")
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
	result.P = ret.Pointer()
	return
}

// gst_tag_list_add_value
//
// [ mode ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ value ] trans: nothing
//
func (v TagList) AddValue(mode TagMergeModeEnum, tag string, value g.Value) {
	iv, err := _I.Get(1030, "TagList", "add_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_mode, arg_tag, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_tag)
}

// gst_tag_list_foreach
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v TagList) Foreach(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(1031, "TagList", "foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTagForeachFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// gst_tag_list_get_boolean
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetBoolean(tag string) (result bool, value bool) {
	iv, err := _I.Get(1032, "TagList", "get_boolean")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// gst_tag_list_get_boolean_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetBooleanIndex(tag string, index uint32) (result bool, value bool) {
	iv, err := _I.Get(1033, "TagList", "get_boolean_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// gst_tag_list_get_date
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetDate(tag string) (result bool, value g.Date) {
	iv, err := _I.Get(1034, "TagList", "get_date")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_date_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetDateIndex(tag string, index uint32) (result bool, value g.Date) {
	iv, err := _I.Get(1035, "TagList", "get_date_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_date_time
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetDateTime(tag string) (result bool, value DateTime) {
	iv, err := _I.Get(1036, "TagList", "get_date_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_date_time_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetDateTimeIndex(tag string, index uint32) (result bool, value DateTime) {
	iv, err := _I.Get(1037, "TagList", "get_date_time_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_double
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetDouble(tag string) (result bool, value float64) {
	iv, err := _I.Get(1038, "TagList", "get_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_tag_list_get_double_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetDoubleIndex(tag string, index uint32) (result bool, value float64) {
	iv, err := _I.Get(1039, "TagList", "get_double_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gst_tag_list_get_float
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetFloat(tag string) (result bool, value float32) {
	iv, err := _I.Get(1040, "TagList", "get_float")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Float()
	result = ret.Bool()
	return
}

// gst_tag_list_get_float_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetFloatIndex(tag string, index uint32) (result bool, value float32) {
	iv, err := _I.Get(1041, "TagList", "get_float_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Float()
	result = ret.Bool()
	return
}

// gst_tag_list_get_int
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetInt(tag string) (result bool, value int32) {
	iv, err := _I.Get(1042, "TagList", "get_int")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_tag_list_get_int64
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetInt64(tag string) (result bool, value int64) {
	iv, err := _I.Get(1043, "TagList", "get_int64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_tag_list_get_int64_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetInt64Index(tag string, index uint32) (result bool, value int64) {
	iv, err := _I.Get(1044, "TagList", "get_int64_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// gst_tag_list_get_int_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetIntIndex(tag string, index uint32) (result bool, value int32) {
	iv, err := _I.Get(1045, "TagList", "get_int_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// gst_tag_list_get_pointer
//
// [ tag ] trans: nothing
//
// [ value ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetPointer(tag string) (result bool, value unsafe.Pointer) {
	iv, err := _I.Get(1046, "TagList", "get_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_pointer_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetPointerIndex(tag string, index uint32) (result bool, value unsafe.Pointer) {
	iv, err := _I.Get(1047, "TagList", "get_pointer_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_sample
//
// [ tag ] trans: nothing
//
// [ sample ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetSample(tag string) (result bool, sample Sample) {
	iv, err := _I.Get(1048, "TagList", "get_sample")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_sample := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_sample}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	sample.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_sample_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ sample ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetSampleIndex(tag string, index uint32) (result bool, sample Sample) {
	iv, err := _I.Get(1049, "TagList", "get_sample_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_sample := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_sample}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	sample.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_tag_list_get_scope
//
// [ result ] trans: nothing
//
func (v TagList) GetScope() (result TagScopeEnum) {
	iv, err := _I.Get(1050, "TagList", "get_scope")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TagScopeEnum(ret.Int())
	return
}

// gst_tag_list_get_string
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetString(tag string) (result bool, value string) {
	iv, err := _I.Get(1051, "TagList", "get_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// gst_tag_list_get_string_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetStringIndex(tag string, index uint32) (result bool, value string) {
	iv, err := _I.Get(1052, "TagList", "get_string_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// gst_tag_list_get_tag_size
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TagList) GetTagSize(tag string) (result uint32) {
	iv, err := _I.Get(1053, "TagList", "get_tag_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_v, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.Uint32()
	return
}

// gst_tag_list_get_uint
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetUint(tag string) (result bool, value uint32) {
	iv, err := _I.Get(1054, "TagList", "get_uint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_tag_list_get_uint64
//
// [ tag ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetUint64(tag string) (result bool, value uint64) {
	iv, err := _I.Get(1055, "TagList", "get_uint64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_tag_list_get_uint64_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetUint64Index(tag string, index uint32) (result bool, value uint64) {
	iv, err := _I.Get(1056, "TagList", "get_uint64_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// gst_tag_list_get_uint_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) GetUintIndex(tag string, index uint32) (result bool, value uint32) {
	iv, err := _I.Get(1057, "TagList", "get_uint_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gst_tag_list_get_value_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TagList) GetValueIndex(tag string, index uint32) (result g.Value) {
	iv, err := _I.Get(1058, "TagList", "get_value_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_tag, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result.P = ret.Pointer()
	return
}

// gst_tag_list_insert
//
// [ from ] trans: nothing
//
// [ mode ] trans: nothing
//
func (v TagList) Insert(from TagList, mode TagMergeModeEnum) {
	iv, err := _I.Get(1059, "TagList", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_from := gi.NewPointerArgument(from.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_from, arg_mode}
	iv.Call(args, nil, nil)
}

// gst_tag_list_is_empty
//
// [ result ] trans: nothing
//
func (v TagList) IsEmpty() (result bool) {
	iv, err := _I.Get(1060, "TagList", "is_empty")
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

// gst_tag_list_is_equal
//
// [ list2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TagList) IsEqual(list2 TagList) (result bool) {
	iv, err := _I.Get(1061, "TagList", "is_equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_list2 := gi.NewPointerArgument(list2.P)
	args := []gi.Argument{arg_v, arg_list2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_tag_list_merge
//
// [ list2 ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ result ] trans: everything
//
func (v TagList) Merge(list2 TagList, mode TagMergeModeEnum) (result TagList) {
	iv, err := _I.Get(1062, "TagList", "merge")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_list2 := gi.NewPointerArgument(list2.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_list2, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_tag_list_n_tags
//
// [ result ] trans: nothing
//
func (v TagList) NTags() (result int32) {
	iv, err := _I.Get(1063, "TagList", "n_tags")
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

// gst_tag_list_nth_tag_name
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TagList) NthTagName(index uint32) (result string) {
	iv, err := _I.Get(1064, "TagList", "nth_tag_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_tag_list_peek_string_index
//
// [ tag ] trans: nothing
//
// [ index ] trans: nothing
//
// [ value ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v TagList) PeekStringIndex(tag string, index uint32) (result bool, value string) {
	iv, err := _I.Get(1065, "TagList", "peek_string_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	arg_index := gi.NewUint32Argument(index)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_tag, arg_index, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tag)
	value = outArgs[0].String().Copy()
	result = ret.Bool()
	return
}

// gst_tag_list_remove_tag
//
// [ tag ] trans: nothing
//
func (v TagList) RemoveTag(tag string) {
	iv, err := _I.Get(1066, "TagList", "remove_tag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_v, arg_tag}
	iv.Call(args, nil, nil)
	gi.Free(c_tag)
}

// gst_tag_list_set_scope
//
// [ scope ] trans: nothing
//
func (v TagList) SetScope(scope TagScopeEnum) {
	iv, err := _I.Get(1067, "TagList", "set_scope")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope := gi.NewIntArgument(int(scope))
	args := []gi.Argument{arg_v, arg_scope}
	iv.Call(args, nil, nil)
}

// gst_tag_list_to_string
//
// [ result ] trans: everything
//
func (v TagList) ToString() (result string) {
	iv, err := _I.Get(1068, "TagList", "to_string")
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

// gst_tag_list_copy_value
//
// [ dest ] trans: nothing, dir: out
//
// [ list ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagListCopyValue1(dest g.Value, list TagList, tag string) (result bool) {
	iv, err := _I.Get(1069, "TagList", "copy_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_list := gi.NewPointerArgument(list.P)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_dest, arg_list, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.Bool()
	return
}

type TagMergeFuncStruct struct {
	F_dest g.Value
	F_src  g.Value
}

func GetPointer_myTagMergeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstTagMergeFunc())
}

//export myGstTagMergeFunc
func myGstTagMergeFunc(dest *C.GValue, src *C.GValue) {
	// TODO: not found user_data
}

// Enum TagMergeMode
type TagMergeModeEnum int

const (
	TagMergeModeUndefined  TagMergeModeEnum = 0
	TagMergeModeReplaceAll TagMergeModeEnum = 1
	TagMergeModeReplace    TagMergeModeEnum = 2
	TagMergeModeAppend     TagMergeModeEnum = 3
	TagMergeModePrepend    TagMergeModeEnum = 4
	TagMergeModeKeep       TagMergeModeEnum = 5
	TagMergeModeKeepAll    TagMergeModeEnum = 6
	TagMergeModeCount      TagMergeModeEnum = 7
)

func TagMergeModeGetType() gi.GType {
	ret := _I.GetGType(167, "TagMergeMode")
	return ret
}

// Enum TagScope
type TagScopeEnum int

const (
	TagScopeStream TagScopeEnum = 0
	TagScopeGlobal TagScopeEnum = 1
)

func TagScopeGetType() gi.GType {
	ret := _I.GetGType(168, "TagScope")
	return ret
}

// Interface TagSetter
type TagSetter struct {
	TagSetterIfc
	P unsafe.Pointer
}
type TagSetterIfc struct{}
type ITagSetter interface{ P_TagSetter() unsafe.Pointer }

func (v TagSetter) P_TagSetter() unsafe.Pointer { return v.P }
func TagSetterGetType() gi.GType {
	ret := _I.GetGType(169, "TagSetter")
	return ret
}

// gst_tag_setter_add_tag_value
//
// [ mode ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ value ] trans: nothing
//
func (v *TagSetterIfc) AddTagValue(mode TagMergeModeEnum, tag string, value g.Value) {
	iv, err := _I.Get(1070, "TagSetter", "add_tag_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_mode := gi.NewIntArgument(int(mode))
	arg_tag := gi.NewStringArgument(c_tag)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_mode, arg_tag, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_tag)
}

// gst_tag_setter_get_tag_list
//
// [ result ] trans: nothing
//
func (v *TagSetterIfc) GetTagList() (result TagList) {
	iv, err := _I.Get(1071, "TagSetter", "get_tag_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_tag_setter_get_tag_merge_mode
//
// [ result ] trans: nothing
//
func (v *TagSetterIfc) GetTagMergeMode() (result TagMergeModeEnum) {
	iv, err := _I.Get(1072, "TagSetter", "get_tag_merge_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TagMergeModeEnum(ret.Int())
	return
}

// gst_tag_setter_merge_tags
//
// [ list ] trans: nothing
//
// [ mode ] trans: nothing
//
func (v *TagSetterIfc) MergeTags(list TagList, mode TagMergeModeEnum) {
	iv, err := _I.Get(1073, "TagSetter", "merge_tags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_list := gi.NewPointerArgument(list.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_list, arg_mode}
	iv.Call(args, nil, nil)
}

// gst_tag_setter_reset_tags
//
func (v *TagSetterIfc) ResetTags() {
	iv, err := _I.Get(1074, "TagSetter", "reset_tags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_tag_setter_set_tag_merge_mode
//
// [ mode ] trans: nothing
//
func (v *TagSetterIfc) SetTagMergeMode(mode TagMergeModeEnum) {
	iv, err := _I.Get(1075, "TagSetter", "set_tag_merge_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	iv.Call(args, nil, nil)
}

// ignore GType struct TagSetterInterface

// Object Task
type Task struct {
	Object
}

func WrapTask(p unsafe.Pointer) (r Task) { r.P = p; return }

type ITask interface{ P_Task() unsafe.Pointer }

func (v Task) P_Task() unsafe.Pointer { return v.P }
func TaskGetType() gi.GType {
	ret := _I.GetGType(170, "Task")
	return ret
}

// gst_task_new
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: everything
//
func NewTask(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result Task) {
	iv, err := _I.Get(1076, "Task", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTaskFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_func1, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_task_get_pool
//
// [ result ] trans: everything
//
func (v Task) GetPool() (result TaskPool) {
	iv, err := _I.Get(1078, "Task", "get_pool")
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

// gst_task_get_state
//
// [ result ] trans: nothing
//
func (v Task) GetState() (result TaskStateEnum) {
	iv, err := _I.Get(1079, "Task", "get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TaskStateEnum(ret.Int())
	return
}

// gst_task_join
//
// [ result ] trans: nothing
//
func (v Task) Join() (result bool) {
	iv, err := _I.Get(1080, "Task", "join")
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

// gst_task_pause
//
// [ result ] trans: nothing
//
func (v Task) Pause() (result bool) {
	iv, err := _I.Get(1081, "Task", "pause")
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

// gst_task_set_enter_callback
//
// [ enter_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Task) SetEnterCallback(enter_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1082, "Task", "set_enter_callback")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enter_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTaskThreadFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_enter_func, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_task_set_leave_callback
//
// [ leave_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func (v Task) SetLeaveCallback(leave_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1083, "Task", "set_leave_callback")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_leave_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTaskThreadFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_v, arg_leave_func, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_task_set_lock
//
// [ mutex ] trans: nothing
//
func (v Task) SetLock(mutex g.RecMutex) {
	iv, err := _I.Get(1084, "Task", "set_lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	args := []gi.Argument{arg_v, arg_mutex}
	iv.Call(args, nil, nil)
}

// gst_task_set_pool
//
// [ pool ] trans: nothing
//
func (v Task) SetPool(pool ITaskPool) {
	iv, err := _I.Get(1085, "Task", "set_pool")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pool != nil {
		tmp = pool.P_TaskPool()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pool := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pool}
	iv.Call(args, nil, nil)
}

// gst_task_set_state
//
// [ state ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Task) SetState(state TaskStateEnum) (result bool) {
	iv, err := _I.Get(1086, "Task", "set_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_v, arg_state}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_task_start
//
// [ result ] trans: nothing
//
func (v Task) Start() (result bool) {
	iv, err := _I.Get(1087, "Task", "start")
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

// gst_task_stop
//
// [ result ] trans: nothing
//
func (v Task) Stop() (result bool) {
	iv, err := _I.Get(1088, "Task", "stop")
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

// ignore GType struct TaskClass

type TaskFunctionStruct struct {
}

func GetPointer_myTaskFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstTaskFunction())
}

//export myGstTaskFunction
func myGstTaskFunction(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TaskFunctionStruct{}
	fn(args)
}

// Object TaskPool
type TaskPool struct {
	Object
}

func WrapTaskPool(p unsafe.Pointer) (r TaskPool) { r.P = p; return }

type ITaskPool interface{ P_TaskPool() unsafe.Pointer }

func (v TaskPool) P_TaskPool() unsafe.Pointer { return v.P }
func TaskPoolGetType() gi.GType {
	ret := _I.GetGType(171, "TaskPool")
	return ret
}

// gst_task_pool_new
//
// [ result ] trans: everything
//
func NewTaskPool() (result TaskPool) {
	iv, err := _I.Get(1089, "TaskPool", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_task_pool_cleanup
//
func (v TaskPool) Cleanup() {
	iv, err := _I.Get(1090, "TaskPool", "cleanup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_task_pool_join
//
// [ id ] trans: nothing
//
func (v TaskPool) Join(id unsafe.Pointer) {
	iv, err := _I.Get(1091, "TaskPool", "join")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewPointerArgument(id)
	args := []gi.Argument{arg_v, arg_id}
	iv.Call(args, nil, nil)
}

// gst_task_pool_prepare
//
func (v TaskPool) Prepare() (err error) {
	iv, err := _I.Get(1092, "TaskPool", "prepare")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	iv.Call(args, nil, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	return
}

// gst_task_pool_push
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TaskPool) Push(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result unsafe.Pointer, err error) {
	iv, err := _I.Get(1093, "TaskPool", "push")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTaskPoolFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_func1, arg_user_data, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Pointer()
	return
}

// ignore GType struct TaskPoolClass

type TaskPoolFunctionStruct struct {
}

func GetPointer_myTaskPoolFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstTaskPoolFunction())
}

//export myGstTaskPoolFunction
func myGstTaskPoolFunction(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TaskPoolFunctionStruct{}
	fn(args)
}

// Struct TaskPrivate
type TaskPrivate struct {
	P unsafe.Pointer
}

func TaskPrivateGetType() gi.GType {
	ret := _I.GetGType(172, "TaskPrivate")
	return ret
}

// Enum TaskState
type TaskStateEnum int

const (
	TaskStateStarted TaskStateEnum = 0
	TaskStateStopped TaskStateEnum = 1
	TaskStatePaused  TaskStateEnum = 2
)

func TaskStateGetType() gi.GType {
	ret := _I.GetGType(173, "TaskState")
	return ret
}

type TaskThreadFuncStruct struct {
	F_task   Task
	F_thread g.Thread
}

func GetPointer_myTaskThreadFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstTaskThreadFunc())
}

//export myGstTaskThreadFunc
func myGstTaskThreadFunc(task *C.GstTask, thread *C.GThread, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TaskThreadFuncStruct{
		F_task:   WrapTask(unsafe.Pointer(task)),
		F_thread: g.Thread{P: unsafe.Pointer(thread)},
	}
	fn(args)
}

// Struct TimedValue
type TimedValue struct {
	P unsafe.Pointer
}

const SizeOfStructTimedValue = 16

func TimedValueGetType() gi.GType {
	ret := _I.GetGType(174, "TimedValue")
	return ret
}

// Struct Toc
type Toc struct {
	P unsafe.Pointer
}

func TocGetType() gi.GType {
	ret := _I.GetGType(175, "Toc")
	return ret
}

// gst_toc_new
//
// [ scope ] trans: nothing
//
// [ result ] trans: everything
//
func NewToc(scope TocScopeEnum) (result Toc) {
	iv, err := _I.Get(1094, "Toc", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_scope := gi.NewIntArgument(int(scope))
	args := []gi.Argument{arg_scope}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_toc_append_entry
//
// [ entry ] trans: everything
//
func (v Toc) AppendEntry(entry TocEntry) {
	iv, err := _I.Get(1095, "Toc", "append_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_entry := gi.NewPointerArgument(entry.P)
	args := []gi.Argument{arg_v, arg_entry}
	iv.Call(args, nil, nil)
}

// gst_toc_dump
//
func (v Toc) Dump() {
	iv, err := _I.Get(1096, "Toc", "dump")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_toc_find_entry
//
// [ uid ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Toc) FindEntry(uid string) (result TocEntry) {
	iv, err := _I.Get(1097, "Toc", "find_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uid := gi.CString(uid)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uid := gi.NewStringArgument(c_uid)
	args := []gi.Argument{arg_v, arg_uid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uid)
	result.P = ret.Pointer()
	return
}

// gst_toc_get_entries
//
// [ result ] trans: nothing
//
func (v Toc) GetEntries() (result g.List) {
	iv, err := _I.Get(1098, "Toc", "get_entries")
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

// gst_toc_get_scope
//
// [ result ] trans: nothing
//
func (v Toc) GetScope() (result TocScopeEnum) {
	iv, err := _I.Get(1099, "Toc", "get_scope")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TocScopeEnum(ret.Int())
	return
}

// gst_toc_get_tags
//
// [ result ] trans: nothing
//
func (v Toc) GetTags() (result TagList) {
	iv, err := _I.Get(1100, "Toc", "get_tags")
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

// gst_toc_merge_tags
//
// [ tags ] trans: nothing
//
// [ mode ] trans: nothing
//
func (v Toc) MergeTags(tags TagList, mode TagMergeModeEnum) {
	iv, err := _I.Get(1101, "Toc", "merge_tags")
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

// gst_toc_set_tags
//
// [ tags ] trans: everything
//
func (v Toc) SetTags(tags TagList) {
	iv, err := _I.Get(1102, "Toc", "set_tags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tags := gi.NewPointerArgument(tags.P)
	args := []gi.Argument{arg_v, arg_tags}
	iv.Call(args, nil, nil)
}

// Struct TocEntry
type TocEntry struct {
	P unsafe.Pointer
}

func TocEntryGetType() gi.GType {
	ret := _I.GetGType(176, "TocEntry")
	return ret
}

// gst_toc_entry_new
//
// [ type1 ] trans: nothing
//
// [ uid ] trans: nothing
//
// [ result ] trans: everything
//
func NewTocEntry(type1 TocEntryTypeEnum, uid string) (result TocEntry) {
	iv, err := _I.Get(1103, "TocEntry", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uid := gi.CString(uid)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_uid := gi.NewStringArgument(c_uid)
	args := []gi.Argument{arg_type1, arg_uid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uid)
	result.P = ret.Pointer()
	return
}

// gst_toc_entry_append_sub_entry
//
// [ subentry ] trans: everything
//
func (v TocEntry) AppendSubEntry(subentry TocEntry) {
	iv, err := _I.Get(1104, "TocEntry", "append_sub_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_subentry := gi.NewPointerArgument(subentry.P)
	args := []gi.Argument{arg_v, arg_subentry}
	iv.Call(args, nil, nil)
}

// gst_toc_entry_get_entry_type
//
// [ result ] trans: nothing
//
func (v TocEntry) GetEntryType() (result TocEntryTypeEnum) {
	iv, err := _I.Get(1105, "TocEntry", "get_entry_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TocEntryTypeEnum(ret.Int())
	return
}

// gst_toc_entry_get_loop
//
// [ loop_type ] trans: everything, dir: out
//
// [ repeat_count ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TocEntry) GetLoop() (result bool, loop_type TocLoopTypeEnum, repeat_count int32) {
	iv, err := _I.Get(1106, "TocEntry", "get_loop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_loop_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_repeat_count := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_loop_type, arg_repeat_count}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	loop_type = TocLoopTypeEnum(outArgs[0].Int())
	repeat_count = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// gst_toc_entry_get_parent
//
// [ result ] trans: nothing
//
func (v TocEntry) GetParent() (result TocEntry) {
	iv, err := _I.Get(1107, "TocEntry", "get_parent")
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

// gst_toc_entry_get_start_stop_times
//
// [ start ] trans: everything, dir: out
//
// [ stop ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TocEntry) GetStartStopTimes() (result bool, start int64, stop int64) {
	iv, err := _I.Get(1108, "TocEntry", "get_start_stop_times")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stop := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_start, arg_stop}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start = outArgs[0].Int64()
	stop = outArgs[1].Int64()
	result = ret.Bool()
	return
}

// gst_toc_entry_get_sub_entries
//
// [ result ] trans: nothing
//
func (v TocEntry) GetSubEntries() (result g.List) {
	iv, err := _I.Get(1109, "TocEntry", "get_sub_entries")
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

// gst_toc_entry_get_tags
//
// [ result ] trans: nothing
//
func (v TocEntry) GetTags() (result TagList) {
	iv, err := _I.Get(1110, "TocEntry", "get_tags")
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

// gst_toc_entry_get_toc
//
// [ result ] trans: nothing
//
func (v TocEntry) GetToc() (result Toc) {
	iv, err := _I.Get(1111, "TocEntry", "get_toc")
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

// gst_toc_entry_get_uid
//
// [ result ] trans: nothing
//
func (v TocEntry) GetUid() (result string) {
	iv, err := _I.Get(1112, "TocEntry", "get_uid")
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

// gst_toc_entry_is_alternative
//
// [ result ] trans: nothing
//
func (v TocEntry) IsAlternative() (result bool) {
	iv, err := _I.Get(1113, "TocEntry", "is_alternative")
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

// gst_toc_entry_is_sequence
//
// [ result ] trans: nothing
//
func (v TocEntry) IsSequence() (result bool) {
	iv, err := _I.Get(1114, "TocEntry", "is_sequence")
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

// gst_toc_entry_merge_tags
//
// [ tags ] trans: nothing
//
// [ mode ] trans: nothing
//
func (v TocEntry) MergeTags(tags TagList, mode TagMergeModeEnum) {
	iv, err := _I.Get(1115, "TocEntry", "merge_tags")
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

// gst_toc_entry_set_loop
//
// [ loop_type ] trans: nothing
//
// [ repeat_count ] trans: nothing
//
func (v TocEntry) SetLoop(loop_type TocLoopTypeEnum, repeat_count int32) {
	iv, err := _I.Get(1116, "TocEntry", "set_loop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_loop_type := gi.NewIntArgument(int(loop_type))
	arg_repeat_count := gi.NewInt32Argument(repeat_count)
	args := []gi.Argument{arg_v, arg_loop_type, arg_repeat_count}
	iv.Call(args, nil, nil)
}

// gst_toc_entry_set_start_stop_times
//
// [ start ] trans: nothing
//
// [ stop ] trans: nothing
//
func (v TocEntry) SetStartStopTimes(start int64, stop int64) {
	iv, err := _I.Get(1117, "TocEntry", "set_start_stop_times")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewInt64Argument(start)
	arg_stop := gi.NewInt64Argument(stop)
	args := []gi.Argument{arg_v, arg_start, arg_stop}
	iv.Call(args, nil, nil)
}

// gst_toc_entry_set_tags
//
// [ tags ] trans: everything
//
func (v TocEntry) SetTags(tags TagList) {
	iv, err := _I.Get(1118, "TocEntry", "set_tags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tags := gi.NewPointerArgument(tags.P)
	args := []gi.Argument{arg_v, arg_tags}
	iv.Call(args, nil, nil)
}

// Enum TocEntryType
type TocEntryTypeEnum int

const (
	TocEntryTypeAngle   TocEntryTypeEnum = -3
	TocEntryTypeVersion TocEntryTypeEnum = -2
	TocEntryTypeEdition TocEntryTypeEnum = -1
	TocEntryTypeInvalid TocEntryTypeEnum = 0
	TocEntryTypeTitle   TocEntryTypeEnum = 1
	TocEntryTypeTrack   TocEntryTypeEnum = 2
	TocEntryTypeChapter TocEntryTypeEnum = 3
)

func TocEntryTypeGetType() gi.GType {
	ret := _I.GetGType(177, "TocEntryType")
	return ret
}

// Enum TocLoopType
type TocLoopTypeEnum int

const (
	TocLoopTypeNone     TocLoopTypeEnum = 0
	TocLoopTypeForward  TocLoopTypeEnum = 1
	TocLoopTypeReverse  TocLoopTypeEnum = 2
	TocLoopTypePingPong TocLoopTypeEnum = 3
)

func TocLoopTypeGetType() gi.GType {
	ret := _I.GetGType(178, "TocLoopType")
	return ret
}

// Enum TocScope
type TocScopeEnum int

const (
	TocScopeGlobal  TocScopeEnum = 1
	TocScopeCurrent TocScopeEnum = 2
)

func TocScopeGetType() gi.GType {
	ret := _I.GetGType(179, "TocScope")
	return ret
}

// Interface TocSetter
type TocSetter struct {
	TocSetterIfc
	P unsafe.Pointer
}
type TocSetterIfc struct{}
type ITocSetter interface{ P_TocSetter() unsafe.Pointer }

func (v TocSetter) P_TocSetter() unsafe.Pointer { return v.P }
func TocSetterGetType() gi.GType {
	ret := _I.GetGType(180, "TocSetter")
	return ret
}

// gst_toc_setter_get_toc
//
// [ result ] trans: everything
//
func (v *TocSetterIfc) GetToc() (result Toc) {
	iv, err := _I.Get(1119, "TocSetter", "get_toc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_toc_setter_reset
//
func (v *TocSetterIfc) Reset() {
	iv, err := _I.Get(1120, "TocSetter", "reset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_toc_setter_set_toc
//
// [ toc ] trans: nothing
//
func (v *TocSetterIfc) SetToc(toc Toc) {
	iv, err := _I.Get(1121, "TocSetter", "set_toc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_toc := gi.NewPointerArgument(toc.P)
	args := []gi.Argument{arg_v, arg_toc}
	iv.Call(args, nil, nil)
}

// ignore GType struct TocSetterInterface

// Object Tracer
type Tracer struct {
	Object
}

func WrapTracer(p unsafe.Pointer) (r Tracer) { r.P = p; return }

type ITracer interface{ P_Tracer() unsafe.Pointer }

func (v Tracer) P_Tracer() unsafe.Pointer { return v.P }
func TracerGetType() gi.GType {
	ret := _I.GetGType(181, "Tracer")
	return ret
}

// ignore GType struct TracerClass

// Object TracerFactory
type TracerFactory struct {
	PluginFeature
}

func WrapTracerFactory(p unsafe.Pointer) (r TracerFactory) { r.P = p; return }

type ITracerFactory interface{ P_TracerFactory() unsafe.Pointer }

func (v TracerFactory) P_TracerFactory() unsafe.Pointer { return v.P }
func TracerFactoryGetType() gi.GType {
	ret := _I.GetGType(182, "TracerFactory")
	return ret
}

// gst_tracer_factory_get_tracer_type
//
// [ result ] trans: nothing
//
func (v TracerFactory) GetTracerType() (result gi.GType) {
	iv, err := _I.Get(1123, "TracerFactory", "get_tracer_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// ignore GType struct TracerFactoryClass

// Struct TracerPrivate
type TracerPrivate struct {
	P unsafe.Pointer
}

func TracerPrivateGetType() gi.GType {
	ret := _I.GetGType(183, "TracerPrivate")
	return ret
}

// Object TracerRecord
type TracerRecord struct {
	Object
}

func WrapTracerRecord(p unsafe.Pointer) (r TracerRecord) { r.P = p; return }

type ITracerRecord interface{ P_TracerRecord() unsafe.Pointer }

func (v TracerRecord) P_TracerRecord() unsafe.Pointer { return v.P }
func TracerRecordGetType() gi.GType {
	ret := _I.GetGType(184, "TracerRecord")
	return ret
}

// ignore GType struct TracerRecordClass

// Flags TracerValueFlags
type TracerValueFlags int

const (
	TracerValueFlagsNone       TracerValueFlags = 0
	TracerValueFlagsOptional   TracerValueFlags = 1
	TracerValueFlagsAggregated TracerValueFlags = 2
)

func TracerValueFlagsGetType() gi.GType {
	ret := _I.GetGType(185, "TracerValueFlags")
	return ret
}

// Enum TracerValueScope
type TracerValueScopeEnum int

const (
	TracerValueScopeProcess TracerValueScopeEnum = 0
	TracerValueScopeThread  TracerValueScopeEnum = 1
	TracerValueScopeElement TracerValueScopeEnum = 2
	TracerValueScopePad     TracerValueScopeEnum = 3
)

func TracerValueScopeGetType() gi.GType {
	ret := _I.GetGType(186, "TracerValueScope")
	return ret
}

// Struct TypeFind
type TypeFind struct {
	P unsafe.Pointer
}

const SizeOfStructTypeFind = 64

func TypeFindGetType() gi.GType {
	ret := _I.GetGType(187, "TypeFind")
	return ret
}

// gst_type_find_get_length
//
// [ result ] trans: nothing
//
func (v TypeFind) GetLength() (result uint64) {
	iv, err := _I.Get(1124, "TypeFind", "get_length")
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

// gst_type_find_peek
//
// [ offset ] trans: nothing
//
// [ size ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v TypeFind) Peek(offset int64) (result gi.Uint8Array) {
	iv, err := _I.Get(1125, "TypeFind", "peek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint32
	_ = size
	size = outArgs[0].Uint32()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(size)}
	return
}

// gst_type_find_suggest
//
// [ probability ] trans: nothing
//
// [ caps ] trans: nothing
//
func (v TypeFind) Suggest(probability uint32, caps Caps) {
	iv, err := _I.Get(1126, "TypeFind", "suggest")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_probability := gi.NewUint32Argument(probability)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_v, arg_probability, arg_caps}
	iv.Call(args, nil, nil)
}

// gst_type_find_register
//
// [ plugin ] trans: nothing
//
// [ name ] trans: nothing
//
// [ rank ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ extensions ] trans: nothing
//
// [ possible_caps ] trans: nothing
//
// [ data ] trans: nothing
//
// [ data_notify ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeFindRegister1(plugin IPlugin, name string, rank uint32, func1 int /*TODO_TYPE CALLBACK*/, extensions string, possible_caps Caps, data unsafe.Pointer, data_notify int /*TODO_TYPE CALLBACK*/) (result bool) {
	iv, err := _I.Get(1127, "TypeFind", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	c_name := gi.CString(name)
	c_extensions := gi.CString(extensions)
	arg_plugin := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_rank := gi.NewUint32Argument(rank)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTypeFindFunction()))
	arg_extensions := gi.NewStringArgument(c_extensions)
	arg_possible_caps := gi.NewPointerArgument(possible_caps.P)
	arg_data := gi.NewPointerArgument(data)
	arg_data_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_plugin, arg_name, arg_rank, arg_func1, arg_extensions, arg_possible_caps, arg_data, arg_data_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_extensions)
	result = ret.Bool()
	return
}

// Object TypeFindFactory
type TypeFindFactory struct {
	PluginFeature
}

func WrapTypeFindFactory(p unsafe.Pointer) (r TypeFindFactory) { r.P = p; return }

type ITypeFindFactory interface{ P_TypeFindFactory() unsafe.Pointer }

func (v TypeFindFactory) P_TypeFindFactory() unsafe.Pointer { return v.P }
func TypeFindFactoryGetType() gi.GType {
	ret := _I.GetGType(188, "TypeFindFactory")
	return ret
}

// gst_type_find_factory_call_function
//
// [ find ] trans: nothing
//
func (v TypeFindFactory) CallFunction(find TypeFind) {
	iv, err := _I.Get(1129, "TypeFindFactory", "call_function")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_find := gi.NewPointerArgument(find.P)
	args := []gi.Argument{arg_v, arg_find}
	iv.Call(args, nil, nil)
}

// gst_type_find_factory_get_caps
//
// [ result ] trans: nothing
//
func (v TypeFindFactory) GetCaps() (result Caps) {
	iv, err := _I.Get(1130, "TypeFindFactory", "get_caps")
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

// gst_type_find_factory_get_extensions
//
// [ result ] trans: nothing
//
func (v TypeFindFactory) GetExtensions() (result gi.CStrArray) {
	iv, err := _I.Get(1131, "TypeFindFactory", "get_extensions")
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

// gst_type_find_factory_has_function
//
// [ result ] trans: nothing
//
func (v TypeFindFactory) HasFunction() (result bool) {
	iv, err := _I.Get(1132, "TypeFindFactory", "has_function")
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

// ignore GType struct TypeFindFactoryClass

type TypeFindFunctionStruct struct {
	F_find TypeFind
}

func GetPointer_myTypeFindFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstTypeFindFunction())
}

//export myGstTypeFindFunction
func myGstTypeFindFunction(find *C.GstTypeFind, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &TypeFindFunctionStruct{
		F_find: TypeFind{P: unsafe.Pointer(find)},
	}
	fn(args)
}

// Enum TypeFindProbability
type TypeFindProbabilityEnum int

const (
	TypeFindProbabilityNone          TypeFindProbabilityEnum = 0
	TypeFindProbabilityMinimum       TypeFindProbabilityEnum = 1
	TypeFindProbabilityPossible      TypeFindProbabilityEnum = 50
	TypeFindProbabilityLikely        TypeFindProbabilityEnum = 80
	TypeFindProbabilityNearlyCertain TypeFindProbabilityEnum = 99
	TypeFindProbabilityMaximum       TypeFindProbabilityEnum = 100
)

func TypeFindProbabilityGetType() gi.GType {
	ret := _I.GetGType(189, "TypeFindProbability")
	return ret
}

// Enum URIError
type URIErrorEnum int

const (
	URIErrorUnsupportedProtocol URIErrorEnum = 0
	URIErrorBadUri              URIErrorEnum = 1
	URIErrorBadState            URIErrorEnum = 2
	URIErrorBadReference        URIErrorEnum = 3
)

func URIErrorGetType() gi.GType {
	ret := _I.GetGType(190, "URIError")
	return ret
}

// Interface URIHandler
type URIHandler struct {
	URIHandlerIfc
	P unsafe.Pointer
}
type URIHandlerIfc struct{}
type IURIHandler interface{ P_URIHandler() unsafe.Pointer }

func (v URIHandler) P_URIHandler() unsafe.Pointer { return v.P }
func URIHandlerGetType() gi.GType {
	ret := _I.GetGType(191, "URIHandler")
	return ret
}

// gst_uri_handler_get_protocols
//
// [ result ] trans: nothing
//
func (v *URIHandlerIfc) GetProtocols() (result gi.CStrArray) {
	iv, err := _I.Get(1133, "URIHandler", "get_protocols")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gst_uri_handler_get_uri
//
// [ result ] trans: everything
//
func (v *URIHandlerIfc) GetUri() (result string) {
	iv, err := _I.Get(1134, "URIHandler", "get_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gst_uri_handler_get_uri_type
//
// [ result ] trans: nothing
//
func (v *URIHandlerIfc) GetUriType() (result URITypeEnum) {
	iv, err := _I.Get(1135, "URIHandler", "get_uri_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = URITypeEnum(ret.Int())
	return
}

// gst_uri_handler_set_uri
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *URIHandlerIfc) SetUri(uri string) (result bool, err error) {
	iv, err := _I.Get(1136, "URIHandler", "set_uri")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
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

// ignore GType struct URIHandlerInterface

// Enum URIType
type URITypeEnum int

const (
	URITypeUnknown URITypeEnum = 0
	URITypeSink    URITypeEnum = 1
	URITypeSrc     URITypeEnum = 2
)

func URITypeGetType() gi.GType {
	ret := _I.GetGType(192, "URIType")
	return ret
}

// Struct Uri
type Uri struct {
	P unsafe.Pointer
}

func UriGetType() gi.GType {
	ret := _I.GetGType(193, "Uri")
	return ret
}

// gst_uri_new
//
// [ scheme ] trans: nothing
//
// [ userinfo ] trans: nothing
//
// [ host ] trans: nothing
//
// [ port ] trans: nothing
//
// [ path ] trans: nothing
//
// [ query ] trans: nothing
//
// [ fragment ] trans: nothing
//
// [ result ] trans: everything
//
func NewUri(scheme string, userinfo string, host string, port uint32, path string, query string, fragment string) (result Uri) {
	iv, err := _I.Get(1137, "Uri", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_scheme := gi.CString(scheme)
	c_userinfo := gi.CString(userinfo)
	c_host := gi.CString(host)
	c_path := gi.CString(path)
	c_query := gi.CString(query)
	c_fragment := gi.CString(fragment)
	arg_scheme := gi.NewStringArgument(c_scheme)
	arg_userinfo := gi.NewStringArgument(c_userinfo)
	arg_host := gi.NewStringArgument(c_host)
	arg_port := gi.NewUint32Argument(port)
	arg_path := gi.NewStringArgument(c_path)
	arg_query := gi.NewStringArgument(c_query)
	arg_fragment := gi.NewStringArgument(c_fragment)
	args := []gi.Argument{arg_scheme, arg_userinfo, arg_host, arg_port, arg_path, arg_query, arg_fragment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_scheme)
	gi.Free(c_userinfo)
	gi.Free(c_host)
	gi.Free(c_path)
	gi.Free(c_query)
	gi.Free(c_fragment)
	result.P = ret.Pointer()
	return
}

// gst_uri_append_path
//
// [ relative_path ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) AppendPath(relative_path string) (result bool) {
	iv, err := _I.Get(1138, "Uri", "append_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_relative_path := gi.CString(relative_path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_relative_path := gi.NewStringArgument(c_relative_path)
	args := []gi.Argument{arg_v, arg_relative_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_relative_path)
	result = ret.Bool()
	return
}

// gst_uri_append_path_segment
//
// [ path_segment ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) AppendPathSegment(path_segment string) (result bool) {
	iv, err := _I.Get(1139, "Uri", "append_path_segment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path_segment := gi.CString(path_segment)
	arg_v := gi.NewPointerArgument(v.P)
	arg_path_segment := gi.NewStringArgument(c_path_segment)
	args := []gi.Argument{arg_v, arg_path_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_path_segment)
	result = ret.Bool()
	return
}

// gst_uri_equal
//
// [ second ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) Equal(second Uri) (result bool) {
	iv, err := _I.Get(1140, "Uri", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_second := gi.NewPointerArgument(second.P)
	args := []gi.Argument{arg_v, arg_second}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_uri_from_string_with_base
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func (v Uri) FromStringWithBase(uri string) (result Uri) {
	iv, err := _I.Get(1141, "Uri", "from_string_with_base")
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
	result.P = ret.Pointer()
	return
}

// gst_uri_get_fragment
//
// [ result ] trans: nothing
//
func (v Uri) GetFragment() (result string) {
	iv, err := _I.Get(1142, "Uri", "get_fragment")
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

// gst_uri_get_host
//
// [ result ] trans: nothing
//
func (v Uri) GetHost() (result string) {
	iv, err := _I.Get(1143, "Uri", "get_host")
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

// gst_uri_get_media_fragment_table
//
// [ result ] trans: everything
//
func (v Uri) GetMediaFragmentTable() (result g.HashTable) {
	iv, err := _I.Get(1144, "Uri", "get_media_fragment_table")
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

// gst_uri_get_path
//
// [ result ] trans: everything
//
func (v Uri) GetPath() (result string) {
	iv, err := _I.Get(1145, "Uri", "get_path")
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

// gst_uri_get_path_segments
//
// [ result ] trans: everything
//
func (v Uri) GetPathSegments() (result g.List) {
	iv, err := _I.Get(1146, "Uri", "get_path_segments")
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

// gst_uri_get_path_string
//
// [ result ] trans: everything
//
func (v Uri) GetPathString() (result string) {
	iv, err := _I.Get(1147, "Uri", "get_path_string")
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

// gst_uri_get_port
//
// [ result ] trans: nothing
//
func (v Uri) GetPort() (result uint32) {
	iv, err := _I.Get(1148, "Uri", "get_port")
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

// gst_uri_get_query_keys
//
// [ result ] trans: container
//
func (v Uri) GetQueryKeys() (result g.List) {
	iv, err := _I.Get(1149, "Uri", "get_query_keys")
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

// gst_uri_get_query_string
//
// [ result ] trans: everything
//
func (v Uri) GetQueryString() (result string) {
	iv, err := _I.Get(1150, "Uri", "get_query_string")
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

// gst_uri_get_query_table
//
// [ result ] trans: everything
//
func (v Uri) GetQueryTable() (result g.HashTable) {
	iv, err := _I.Get(1151, "Uri", "get_query_table")
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

// gst_uri_get_query_value
//
// [ query_key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) GetQueryValue(query_key string) (result string) {
	iv, err := _I.Get(1152, "Uri", "get_query_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_query_key := gi.CString(query_key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_query_key := gi.NewStringArgument(c_query_key)
	args := []gi.Argument{arg_v, arg_query_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_query_key)
	result = ret.String().Copy()
	return
}

// gst_uri_get_scheme
//
// [ result ] trans: nothing
//
func (v Uri) GetScheme() (result string) {
	iv, err := _I.Get(1153, "Uri", "get_scheme")
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

// gst_uri_get_userinfo
//
// [ result ] trans: nothing
//
func (v Uri) GetUserinfo() (result string) {
	iv, err := _I.Get(1154, "Uri", "get_userinfo")
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

// gst_uri_is_normalized
//
// [ result ] trans: nothing
//
func (v Uri) IsNormalized() (result bool) {
	iv, err := _I.Get(1155, "Uri", "is_normalized")
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

// gst_uri_is_writable
//
// [ result ] trans: nothing
//
func (v Uri) IsWritable() (result bool) {
	iv, err := _I.Get(1156, "Uri", "is_writable")
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

// gst_uri_join
//
// [ ref_uri ] trans: nothing
//
// [ result ] trans: everything
//
func (v Uri) Join(ref_uri Uri) (result Uri) {
	iv, err := _I.Get(1157, "Uri", "join")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ref_uri := gi.NewPointerArgument(ref_uri.P)
	args := []gi.Argument{arg_v, arg_ref_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_uri_make_writable
//
// [ result ] trans: everything
//
func (v Uri) MakeWritable() (result Uri) {
	iv, err := _I.Get(1158, "Uri", "make_writable")
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

// gst_uri_new_with_base
//
// [ scheme ] trans: nothing
//
// [ userinfo ] trans: nothing
//
// [ host ] trans: nothing
//
// [ port ] trans: nothing
//
// [ path ] trans: nothing
//
// [ query ] trans: nothing
//
// [ fragment ] trans: nothing
//
// [ result ] trans: everything
//
func (v Uri) NewWithBase(scheme string, userinfo string, host string, port uint32, path string, query string, fragment string) (result Uri) {
	iv, err := _I.Get(1159, "Uri", "new_with_base")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_scheme := gi.CString(scheme)
	c_userinfo := gi.CString(userinfo)
	c_host := gi.CString(host)
	c_path := gi.CString(path)
	c_query := gi.CString(query)
	c_fragment := gi.CString(fragment)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scheme := gi.NewStringArgument(c_scheme)
	arg_userinfo := gi.NewStringArgument(c_userinfo)
	arg_host := gi.NewStringArgument(c_host)
	arg_port := gi.NewUint32Argument(port)
	arg_path := gi.NewStringArgument(c_path)
	arg_query := gi.NewStringArgument(c_query)
	arg_fragment := gi.NewStringArgument(c_fragment)
	args := []gi.Argument{arg_v, arg_scheme, arg_userinfo, arg_host, arg_port, arg_path, arg_query, arg_fragment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_scheme)
	gi.Free(c_userinfo)
	gi.Free(c_host)
	gi.Free(c_path)
	gi.Free(c_query)
	gi.Free(c_fragment)
	result.P = ret.Pointer()
	return
}

// gst_uri_normalize
//
// [ result ] trans: nothing
//
func (v Uri) Normalize() (result bool) {
	iv, err := _I.Get(1160, "Uri", "normalize")
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

// gst_uri_query_has_key
//
// [ query_key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) QueryHasKey(query_key string) (result bool) {
	iv, err := _I.Get(1161, "Uri", "query_has_key")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_query_key := gi.CString(query_key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_query_key := gi.NewStringArgument(c_query_key)
	args := []gi.Argument{arg_v, arg_query_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_query_key)
	result = ret.Bool()
	return
}

// gst_uri_remove_query_key
//
// [ query_key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) RemoveQueryKey(query_key string) (result bool) {
	iv, err := _I.Get(1162, "Uri", "remove_query_key")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_query_key := gi.CString(query_key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_query_key := gi.NewStringArgument(c_query_key)
	args := []gi.Argument{arg_v, arg_query_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_query_key)
	result = ret.Bool()
	return
}

// gst_uri_set_fragment
//
// [ fragment ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetFragment(fragment string) (result bool) {
	iv, err := _I.Get(1163, "Uri", "set_fragment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_fragment := gi.CString(fragment)
	arg_v := gi.NewPointerArgument(v.P)
	arg_fragment := gi.NewStringArgument(c_fragment)
	args := []gi.Argument{arg_v, arg_fragment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_fragment)
	result = ret.Bool()
	return
}

// gst_uri_set_host
//
// [ host ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetHost(host string) (result bool) {
	iv, err := _I.Get(1164, "Uri", "set_host")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_host := gi.CString(host)
	arg_v := gi.NewPointerArgument(v.P)
	arg_host := gi.NewStringArgument(c_host)
	args := []gi.Argument{arg_v, arg_host}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_host)
	result = ret.Bool()
	return
}

// gst_uri_set_path
//
// [ path ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetPath(path string) (result bool) {
	iv, err := _I.Get(1165, "Uri", "set_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_v, arg_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_path)
	result = ret.Bool()
	return
}

// gst_uri_set_path_segments
//
// [ path_segments ] trans: everything
//
// [ result ] trans: nothing
//
func (v Uri) SetPathSegments(path_segments g.List) (result bool) {
	iv, err := _I.Get(1166, "Uri", "set_path_segments")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_path_segments := gi.NewPointerArgument(path_segments.P)
	args := []gi.Argument{arg_v, arg_path_segments}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_uri_set_path_string
//
// [ path ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetPathString(path string) (result bool) {
	iv, err := _I.Get(1167, "Uri", "set_path_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_v, arg_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_path)
	result = ret.Bool()
	return
}

// gst_uri_set_port
//
// [ port ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetPort(port uint32) (result bool) {
	iv, err := _I.Get(1168, "Uri", "set_port")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_port := gi.NewUint32Argument(port)
	args := []gi.Argument{arg_v, arg_port}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_uri_set_query_string
//
// [ query ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetQueryString(query string) (result bool) {
	iv, err := _I.Get(1169, "Uri", "set_query_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_query := gi.CString(query)
	arg_v := gi.NewPointerArgument(v.P)
	arg_query := gi.NewStringArgument(c_query)
	args := []gi.Argument{arg_v, arg_query}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_query)
	result = ret.Bool()
	return
}

// gst_uri_set_query_table
//
// [ query_table ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetQueryTable(query_table g.HashTable) (result bool) {
	iv, err := _I.Get(1170, "Uri", "set_query_table")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_query_table := gi.NewPointerArgument(query_table.P)
	args := []gi.Argument{arg_v, arg_query_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_uri_set_query_value
//
// [ query_key ] trans: nothing
//
// [ query_value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetQueryValue(query_key string, query_value string) (result bool) {
	iv, err := _I.Get(1171, "Uri", "set_query_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_query_key := gi.CString(query_key)
	c_query_value := gi.CString(query_value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_query_key := gi.NewStringArgument(c_query_key)
	arg_query_value := gi.NewStringArgument(c_query_value)
	args := []gi.Argument{arg_v, arg_query_key, arg_query_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_query_key)
	gi.Free(c_query_value)
	result = ret.Bool()
	return
}

// gst_uri_set_scheme
//
// [ scheme ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetScheme(scheme string) (result bool) {
	iv, err := _I.Get(1172, "Uri", "set_scheme")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_scheme := gi.CString(scheme)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scheme := gi.NewStringArgument(c_scheme)
	args := []gi.Argument{arg_v, arg_scheme}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_scheme)
	result = ret.Bool()
	return
}

// gst_uri_set_userinfo
//
// [ userinfo ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Uri) SetUserinfo(userinfo string) (result bool) {
	iv, err := _I.Get(1173, "Uri", "set_userinfo")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_userinfo := gi.CString(userinfo)
	arg_v := gi.NewPointerArgument(v.P)
	arg_userinfo := gi.NewStringArgument(c_userinfo)
	args := []gi.Argument{arg_v, arg_userinfo}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_userinfo)
	result = ret.Bool()
	return
}

// gst_uri_to_string
//
// [ result ] trans: everything
//
func (v Uri) ToString() (result string) {
	iv, err := _I.Get(1174, "Uri", "to_string")
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

// Deprecated
//
// gst_uri_construct
//
// [ protocol ] trans: nothing
//
// [ location ] trans: nothing
//
// [ result ] trans: everything
//
func UriConstruct1(protocol string, location string) (result string) {
	iv, err := _I.Get(1175, "Uri", "construct")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_protocol := gi.CString(protocol)
	c_location := gi.CString(location)
	arg_protocol := gi.NewStringArgument(c_protocol)
	arg_location := gi.NewStringArgument(c_location)
	args := []gi.Argument{arg_protocol, arg_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_protocol)
	gi.Free(c_location)
	result = ret.String().Take()
	return
}

// gst_uri_from_string
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriFromString1(uri string) (result Uri) {
	iv, err := _I.Get(1176, "Uri", "from_string")
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
	result.P = ret.Pointer()
	return
}

// gst_uri_get_location
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriGetLocation1(uri string) (result string) {
	iv, err := _I.Get(1177, "Uri", "get_location")
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

// gst_uri_get_protocol
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriGetProtocol1(uri string) (result string) {
	iv, err := _I.Get(1178, "Uri", "get_protocol")
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

// gst_uri_has_protocol
//
// [ uri ] trans: nothing
//
// [ protocol ] trans: nothing
//
// [ result ] trans: nothing
//
func UriHasProtocol1(uri string, protocol string) (result bool) {
	iv, err := _I.Get(1179, "Uri", "has_protocol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_protocol := gi.CString(protocol)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_protocol := gi.NewStringArgument(c_protocol)
	args := []gi.Argument{arg_uri, arg_protocol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri)
	gi.Free(c_protocol)
	result = ret.Bool()
	return
}

// gst_uri_is_valid
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func UriIsValid1(uri string) (result bool) {
	iv, err := _I.Get(1180, "Uri", "is_valid")
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
	result = ret.Bool()
	return
}

// gst_uri_join_strings
//
// [ base_uri ] trans: nothing
//
// [ ref_uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriJoinStrings1(base_uri string, ref_uri string) (result string) {
	iv, err := _I.Get(1181, "Uri", "join_strings")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_base_uri := gi.CString(base_uri)
	c_ref_uri := gi.CString(ref_uri)
	arg_base_uri := gi.NewStringArgument(c_base_uri)
	arg_ref_uri := gi.NewStringArgument(c_ref_uri)
	args := []gi.Argument{arg_base_uri, arg_ref_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_base_uri)
	gi.Free(c_ref_uri)
	result = ret.String().Take()
	return
}

// gst_uri_protocol_is_supported
//
// [ type1 ] trans: nothing
//
// [ protocol ] trans: nothing
//
// [ result ] trans: nothing
//
func UriProtocolIsSupported1(type1 URITypeEnum, protocol string) (result bool) {
	iv, err := _I.Get(1182, "Uri", "protocol_is_supported")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_protocol := gi.CString(protocol)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_protocol := gi.NewStringArgument(c_protocol)
	args := []gi.Argument{arg_type1, arg_protocol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_protocol)
	result = ret.Bool()
	return
}

// gst_uri_protocol_is_valid
//
// [ protocol ] trans: nothing
//
// [ result ] trans: nothing
//
func UriProtocolIsValid1(protocol string) (result bool) {
	iv, err := _I.Get(1183, "Uri", "protocol_is_valid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_protocol := gi.CString(protocol)
	arg_protocol := gi.NewStringArgument(c_protocol)
	args := []gi.Argument{arg_protocol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_protocol)
	result = ret.Bool()
	return
}

// Object ValueArray
type ValueArray struct {
	P unsafe.Pointer
}

func WrapValueArray(p unsafe.Pointer) (r ValueArray) { r.P = p; return }

type IValueArray interface{ P_ValueArray() unsafe.Pointer }

func (v ValueArray) P_ValueArray() unsafe.Pointer { return v.P }
func ValueArrayGetType() gi.GType {
	ret := _I.GetGType(194, "ValueArray")
	return ret
}

// gst_value_array_append_and_take_value
//
// [ value ] trans: nothing
//
// [ append_value ] trans: everything
//
func ValueArrayAppendAndTakeValue1(value g.Value, append_value g.Value) {
	iv, err := _I.Get(1184, "ValueArray", "append_and_take_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_append_value := gi.NewPointerArgument(append_value.P)
	args := []gi.Argument{arg_value, arg_append_value}
	iv.Call(args, nil, nil)
}

// gst_value_array_append_value
//
// [ value ] trans: nothing
//
// [ append_value ] trans: nothing
//
func ValueArrayAppendValue1(value g.Value, append_value g.Value) {
	iv, err := _I.Get(1185, "ValueArray", "append_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_append_value := gi.NewPointerArgument(append_value.P)
	args := []gi.Argument{arg_value, arg_append_value}
	iv.Call(args, nil, nil)
}

// gst_value_array_get_size
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueArrayGetSize1(value g.Value) (result uint32) {
	iv, err := _I.Get(1186, "ValueArray", "get_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_value_array_get_value
//
// [ value ] trans: nothing
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueArrayGetValue1(value g.Value, index uint32) (result g.Value) {
	iv, err := _I.Get(1187, "ValueArray", "get_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_value, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_value_array_prepend_value
//
// [ value ] trans: nothing
//
// [ prepend_value ] trans: nothing
//
func ValueArrayPrependValue1(value g.Value, prepend_value g.Value) {
	iv, err := _I.Get(1188, "ValueArray", "prepend_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_prepend_value := gi.NewPointerArgument(prepend_value.P)
	args := []gi.Argument{arg_value, arg_prepend_value}
	iv.Call(args, nil, nil)
}

type ValueCompareFuncStruct struct {
	F_value1 g.Value
	F_value2 g.Value
}

func GetPointer_myValueCompareFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstValueCompareFunc())
}

//export myGstValueCompareFunc
func myGstValueCompareFunc(value1 *C.GValue, value2 *C.GValue) {
	// TODO: not found user_data
}

type ValueDeserializeFuncStruct struct {
	F_dest g.Value
	F_s    string
}

func GetPointer_myValueDeserializeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstValueDeserializeFunc())
}

//export myGstValueDeserializeFunc
func myGstValueDeserializeFunc(dest *C.GValue, s *C.gchar) {
	// TODO: not found user_data
}

// Object ValueList
type ValueList struct {
	P unsafe.Pointer
}

func WrapValueList(p unsafe.Pointer) (r ValueList) { r.P = p; return }

type IValueList interface{ P_ValueList() unsafe.Pointer }

func (v ValueList) P_ValueList() unsafe.Pointer { return v.P }
func ValueListGetType() gi.GType {
	ret := _I.GetGType(195, "ValueList")
	return ret
}

// gst_value_list_append_and_take_value
//
// [ value ] trans: nothing
//
// [ append_value ] trans: everything
//
func ValueListAppendAndTakeValue1(value g.Value, append_value g.Value) {
	iv, err := _I.Get(1189, "ValueList", "append_and_take_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_append_value := gi.NewPointerArgument(append_value.P)
	args := []gi.Argument{arg_value, arg_append_value}
	iv.Call(args, nil, nil)
}

// gst_value_list_append_value
//
// [ value ] trans: nothing
//
// [ append_value ] trans: nothing
//
func ValueListAppendValue1(value g.Value, append_value g.Value) {
	iv, err := _I.Get(1190, "ValueList", "append_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_append_value := gi.NewPointerArgument(append_value.P)
	args := []gi.Argument{arg_value, arg_append_value}
	iv.Call(args, nil, nil)
}

// gst_value_list_concat
//
// [ dest ] trans: nothing, dir: out
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
func ValueListConcat1(dest g.Value, value1 g.Value, value2 g.Value) {
	iv, err := _I.Get(1191, "ValueList", "concat")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_dest, arg_value1, arg_value2}
	iv.Call(args, nil, nil)
}

// gst_value_list_get_size
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueListGetSize1(value g.Value) (result uint32) {
	iv, err := _I.Get(1192, "ValueList", "get_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_value_list_get_value
//
// [ value ] trans: nothing
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueListGetValue1(value g.Value, index uint32) (result g.Value) {
	iv, err := _I.Get(1193, "ValueList", "get_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_index := gi.NewUint32Argument(index)
	args := []gi.Argument{arg_value, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_value_list_merge
//
// [ dest ] trans: nothing, dir: out
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
func ValueListMerge1(dest g.Value, value1 g.Value, value2 g.Value) {
	iv, err := _I.Get(1194, "ValueList", "merge")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_dest, arg_value1, arg_value2}
	iv.Call(args, nil, nil)
}

// gst_value_list_prepend_value
//
// [ value ] trans: nothing
//
// [ prepend_value ] trans: nothing
//
func ValueListPrependValue1(value g.Value, prepend_value g.Value) {
	iv, err := _I.Get(1195, "ValueList", "prepend_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_prepend_value := gi.NewPointerArgument(prepend_value.P)
	args := []gi.Argument{arg_value, arg_prepend_value}
	iv.Call(args, nil, nil)
}

type ValueSerializeFuncStruct struct {
	F_value1 g.Value
}

func GetPointer_myValueSerializeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstValueSerializeFunc())
}

//export myGstValueSerializeFunc
func myGstValueSerializeFunc(value1 *C.GValue) {
	// TODO: not found user_data
}

// Struct ValueTable
type ValueTable struct {
	P unsafe.Pointer
}

const SizeOfStructValueTable = 64

func ValueTableGetType() gi.GType {
	ret := _I.GetGType(196, "ValueTable")
	return ret
}

// gst_buffer_get_max_memory
//
// [ result ] trans: nothing
//
func BufferGetMaxMemory() (result uint32) {
	iv, err := _I.Get(1196, "buffer_get_max_memory", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_caps_features_from_string
//
// [ features ] trans: nothing
//
// [ result ] trans: everything
//
func CapsFeaturesFromString(features string) (result CapsFeatures) {
	iv, err := _I.Get(1197, "caps_features_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_features := gi.CString(features)
	arg_features := gi.NewStringArgument(c_features)
	args := []gi.Argument{arg_features}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_features)
	result.P = ret.Pointer()
	return
}

// gst_caps_from_string
//
// [ string ] trans: nothing
//
// [ result ] trans: everything
//
func CapsFromString(string string) (result Caps) {
	iv, err := _I.Get(1198, "caps_from_string", "")
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

// gst_core_error_quark
//
// [ result ] trans: nothing
//
func CoreErrorQuark() (result uint32) {
	iv, err := _I.Get(1199, "core_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_debug_add_log_function
//
// [ func1 ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ notify ] trans: nothing
//
func DebugAddLogFunction(func1 int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1200, "debug_add_log_function", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myLogFunction()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_func1, arg_user_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gst_debug_add_ring_buffer_logger
//
// [ max_size_per_thread ] trans: nothing
//
// [ thread_timeout ] trans: nothing
//
func DebugAddRingBufferLogger(max_size_per_thread uint32, thread_timeout uint32) {
	iv, err := _I.Get(1201, "debug_add_ring_buffer_logger", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_max_size_per_thread := gi.NewUint32Argument(max_size_per_thread)
	arg_thread_timeout := gi.NewUint32Argument(thread_timeout)
	args := []gi.Argument{arg_max_size_per_thread, arg_thread_timeout}
	iv.Call(args, nil, nil)
}

// gst_debug_bin_to_dot_data
//
// [ bin ] trans: nothing
//
// [ details ] trans: nothing
//
// [ result ] trans: everything
//
func DebugBinToDotData(bin IBin, details DebugGraphDetailsFlags) (result string) {
	iv, err := _I.Get(1202, "debug_bin_to_dot_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if bin != nil {
		tmp = bin.P_Bin()
	}
	arg_bin := gi.NewPointerArgument(tmp)
	arg_details := gi.NewIntArgument(int(details))
	args := []gi.Argument{arg_bin, arg_details}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gst_debug_bin_to_dot_file
//
// [ bin ] trans: nothing
//
// [ details ] trans: nothing
//
// [ file_name ] trans: nothing
//
func DebugBinToDotFile(bin IBin, details DebugGraphDetailsFlags, file_name string) {
	iv, err := _I.Get(1203, "debug_bin_to_dot_file", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if bin != nil {
		tmp = bin.P_Bin()
	}
	c_file_name := gi.CString(file_name)
	arg_bin := gi.NewPointerArgument(tmp)
	arg_details := gi.NewIntArgument(int(details))
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_bin, arg_details, arg_file_name}
	iv.Call(args, nil, nil)
	gi.Free(c_file_name)
}

// gst_debug_bin_to_dot_file_with_ts
//
// [ bin ] trans: nothing
//
// [ details ] trans: nothing
//
// [ file_name ] trans: nothing
//
func DebugBinToDotFileWithTs(bin IBin, details DebugGraphDetailsFlags, file_name string) {
	iv, err := _I.Get(1204, "debug_bin_to_dot_file_with_ts", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if bin != nil {
		tmp = bin.P_Bin()
	}
	c_file_name := gi.CString(file_name)
	arg_bin := gi.NewPointerArgument(tmp)
	arg_details := gi.NewIntArgument(int(details))
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_bin, arg_details, arg_file_name}
	iv.Call(args, nil, nil)
	gi.Free(c_file_name)
}

// gst_debug_construct_term_color
//
// [ colorinfo ] trans: nothing
//
// [ result ] trans: everything
//
func DebugConstructTermColor(colorinfo uint32) (result string) {
	iv, err := _I.Get(1205, "debug_construct_term_color", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_colorinfo := gi.NewUint32Argument(colorinfo)
	args := []gi.Argument{arg_colorinfo}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gst_debug_construct_win_color
//
// [ colorinfo ] trans: nothing
//
// [ result ] trans: nothing
//
func DebugConstructWinColor(colorinfo uint32) (result int32) {
	iv, err := _I.Get(1206, "debug_construct_win_color", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_colorinfo := gi.NewUint32Argument(colorinfo)
	args := []gi.Argument{arg_colorinfo}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_debug_get_all_categories
//
// [ result ] trans: container
//
func DebugGetAllCategories() (result g.SList) {
	iv, err := _I.Get(1207, "debug_get_all_categories", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_debug_get_color_mode
//
// [ result ] trans: nothing
//
func DebugGetColorMode() (result DebugColorModeEnum) {
	iv, err := _I.Get(1208, "debug_get_color_mode", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = DebugColorModeEnum(ret.Int())
	return
}

// gst_debug_get_default_threshold
//
// [ result ] trans: nothing
//
func DebugGetDefaultThreshold() (result DebugLevelEnum) {
	iv, err := _I.Get(1209, "debug_get_default_threshold", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = DebugLevelEnum(ret.Int())
	return
}

// gst_debug_get_stack_trace
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func DebugGetStackTrace(flags StackTraceFlags) (result string) {
	iv, err := _I.Get(1210, "debug_get_stack_trace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gst_debug_is_active
//
// [ result ] trans: nothing
//
func DebugIsActive() (result bool) {
	iv, err := _I.Get(1211, "debug_is_active", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gst_debug_is_colored
//
// [ result ] trans: nothing
//
func DebugIsColored() (result bool) {
	iv, err := _I.Get(1212, "debug_is_colored", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gst_debug_level_get_name
//
// [ level ] trans: nothing
//
// [ result ] trans: nothing
//
func DebugLevelGetName(level DebugLevelEnum) (result string) {
	iv, err := _I.Get(1213, "debug_level_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_level := gi.NewIntArgument(int(level))
	args := []gi.Argument{arg_level}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_debug_log_default
//
// [ category ] trans: nothing
//
// [ level ] trans: nothing
//
// [ file ] trans: nothing
//
// [ function ] trans: nothing
//
// [ line ] trans: nothing
//
// [ object ] trans: nothing
//
// [ message ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DebugLogDefault(category DebugCategory, level DebugLevelEnum, file string, function string, line int32, object g.IObject, message DebugMessage, user_data unsafe.Pointer) {
	iv, err := _I.Get(1214, "debug_log_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file := gi.CString(file)
	c_function := gi.CString(function)
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	arg_category := gi.NewPointerArgument(category.P)
	arg_level := gi.NewIntArgument(int(level))
	arg_file := gi.NewStringArgument(c_file)
	arg_function := gi.NewStringArgument(c_function)
	arg_line := gi.NewInt32Argument(line)
	arg_object := gi.NewPointerArgument(tmp)
	arg_message := gi.NewPointerArgument(message.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_category, arg_level, arg_file, arg_function, arg_line, arg_object, arg_message, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_file)
	gi.Free(c_function)
}

// gst_debug_print_stack_trace
//
func DebugPrintStackTrace() {
	iv, err := _I.Get(1215, "debug_print_stack_trace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gst_debug_remove_log_function
//
// [ func1 ] trans: nothing
//
// [ result ] trans: nothing
//
func DebugRemoveLogFunction(func1 int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(1216, "debug_remove_log_function", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myLogFunction()))
	args := []gi.Argument{arg_func1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_debug_remove_log_function_by_data
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func DebugRemoveLogFunctionByData(data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(1217, "debug_remove_log_function_by_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_debug_remove_ring_buffer_logger
//
func DebugRemoveRingBufferLogger() {
	iv, err := _I.Get(1218, "debug_remove_ring_buffer_logger", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gst_debug_ring_buffer_logger_get_logs
//
// [ result ] trans: everything
//
func DebugRingBufferLoggerGetLogs() (result gi.CStrArray) {
	iv, err := _I.Get(1219, "debug_ring_buffer_logger_get_logs", "")
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

// gst_debug_set_active
//
// [ active ] trans: nothing
//
func DebugSetActive(active bool) {
	iv, err := _I.Get(1220, "debug_set_active", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_active := gi.NewBoolArgument(active)
	args := []gi.Argument{arg_active}
	iv.Call(args, nil, nil)
}

// gst_debug_set_color_mode
//
// [ mode ] trans: nothing
//
func DebugSetColorMode(mode DebugColorModeEnum) {
	iv, err := _I.Get(1221, "debug_set_color_mode", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_mode}
	iv.Call(args, nil, nil)
}

// gst_debug_set_color_mode_from_string
//
// [ mode ] trans: nothing
//
func DebugSetColorModeFromString(mode string) {
	iv, err := _I.Get(1222, "debug_set_color_mode_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mode := gi.CString(mode)
	arg_mode := gi.NewStringArgument(c_mode)
	args := []gi.Argument{arg_mode}
	iv.Call(args, nil, nil)
	gi.Free(c_mode)
}

// gst_debug_set_colored
//
// [ colored ] trans: nothing
//
func DebugSetColored(colored bool) {
	iv, err := _I.Get(1223, "debug_set_colored", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_colored := gi.NewBoolArgument(colored)
	args := []gi.Argument{arg_colored}
	iv.Call(args, nil, nil)
}

// gst_debug_set_default_threshold
//
// [ level ] trans: nothing
//
func DebugSetDefaultThreshold(level DebugLevelEnum) {
	iv, err := _I.Get(1224, "debug_set_default_threshold", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_level := gi.NewIntArgument(int(level))
	args := []gi.Argument{arg_level}
	iv.Call(args, nil, nil)
}

// gst_debug_set_threshold_for_name
//
// [ name ] trans: nothing
//
// [ level ] trans: nothing
//
func DebugSetThresholdForName(name string, level DebugLevelEnum) {
	iv, err := _I.Get(1225, "debug_set_threshold_for_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_level := gi.NewIntArgument(int(level))
	args := []gi.Argument{arg_name, arg_level}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gst_debug_set_threshold_from_string
//
// [ list ] trans: nothing
//
// [ reset ] trans: nothing
//
func DebugSetThresholdFromString(list string, reset bool) {
	iv, err := _I.Get(1226, "debug_set_threshold_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_list := gi.CString(list)
	arg_list := gi.NewStringArgument(c_list)
	arg_reset := gi.NewBoolArgument(reset)
	args := []gi.Argument{arg_list, arg_reset}
	iv.Call(args, nil, nil)
	gi.Free(c_list)
}

// gst_debug_unset_threshold_for_name
//
// [ name ] trans: nothing
//
func DebugUnsetThresholdForName(name string) {
	iv, err := _I.Get(1227, "debug_unset_threshold_for_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gst_deinit
//
func Deinit() {
	iv, err := _I.Get(1228, "deinit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gst_dynamic_type_register
//
// [ plugin ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func DynamicTypeRegister(plugin IPlugin, type1 gi.GType) (result bool) {
	iv, err := _I.Get(1229, "dynamic_type_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	arg_plugin := gi.NewPointerArgument(tmp)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_plugin, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_error_get_message
//
// [ domain ] trans: nothing
//
// [ code ] trans: nothing
//
// [ result ] trans: everything
//
func ErrorGetMessage(domain uint32, code int32) (result string) {
	iv, err := _I.Get(1230, "error_get_message", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	args := []gi.Argument{arg_domain, arg_code}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gst_event_type_get_flags
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func EventTypeGetFlags(type1 EventTypeEnum) (result EventTypeFlags) {
	iv, err := _I.Get(1231, "event_type_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = EventTypeFlags(ret.Int())
	return
}

// gst_event_type_get_name
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func EventTypeGetName(type1 EventTypeEnum) (result string) {
	iv, err := _I.Get(1232, "event_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_event_type_to_quark
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func EventTypeToQuark(type1 EventTypeEnum) (result uint32) {
	iv, err := _I.Get(1233, "event_type_to_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_filename_to_uri
//
// [ filename ] trans: nothing
//
// [ result ] trans: everything
//
func FilenameToUri(filename string) (result string, err error) {
	iv, err := _I.Get(1234, "filename_to_uri", "")
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

// gst_flow_get_name
//
// [ ret ] trans: nothing
//
// [ result ] trans: nothing
//
func FlowGetName(ret FlowReturnEnum) (result string) {
	iv, err := _I.Get(1235, "flow_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ret := gi.NewIntArgument(int(ret))
	args := []gi.Argument{arg_ret}
	var ret1 gi.Argument
	iv.Call(args, &ret1, nil)
	result = ret1.String().Copy()
	return
}

// gst_flow_to_quark
//
// [ ret ] trans: nothing
//
// [ result ] trans: nothing
//
func FlowToQuark(ret FlowReturnEnum) (result uint32) {
	iv, err := _I.Get(1236, "flow_to_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ret := gi.NewIntArgument(int(ret))
	args := []gi.Argument{arg_ret}
	var ret1 gi.Argument
	iv.Call(args, &ret1, nil)
	result = ret1.Uint32()
	return
}

// gst_format_get_by_nick
//
// [ nick ] trans: nothing
//
// [ result ] trans: nothing
//
func FormatGetByNick(nick string) (result FormatEnum) {
	iv, err := _I.Get(1237, "format_get_by_nick", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_nick := gi.CString(nick)
	arg_nick := gi.NewStringArgument(c_nick)
	args := []gi.Argument{arg_nick}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_nick)
	result = FormatEnum(ret.Int())
	return
}

// gst_format_get_details
//
// [ format ] trans: nothing
//
// [ result ] trans: nothing
//
func FormatGetDetails(format FormatEnum) (result FormatDefinition) {
	iv, err := _I.Get(1238, "format_get_details", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_format_get_name
//
// [ format ] trans: nothing
//
// [ result ] trans: nothing
//
func FormatGetName(format FormatEnum) (result string) {
	iv, err := _I.Get(1239, "format_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_format_iterate_definitions
//
// [ result ] trans: everything
//
func FormatIterateDefinitions() (result Iterator) {
	iv, err := _I.Get(1240, "format_iterate_definitions", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_format_register
//
// [ nick ] trans: nothing
//
// [ description ] trans: nothing
//
// [ result ] trans: nothing
//
func FormatRegister(nick string, description string) (result FormatEnum) {
	iv, err := _I.Get(1241, "format_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_nick := gi.CString(nick)
	c_description := gi.CString(description)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_nick, arg_description}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_nick)
	gi.Free(c_description)
	result = FormatEnum(ret.Int())
	return
}

// gst_format_to_quark
//
// [ format ] trans: nothing
//
// [ result ] trans: nothing
//
func FormatToQuark(format FormatEnum) (result uint32) {
	iv, err := _I.Get(1242, "format_to_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_formats_contains
//
// [ formats ] trans: nothing
//
// [ format ] trans: nothing
//
// [ result ] trans: nothing
//
func FormatsContains(formats unsafe.Pointer, format FormatEnum) (result bool) {
	iv, err := _I.Get(1243, "formats_contains", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_formats := gi.NewPointerArgument(formats)
	arg_format := gi.NewIntArgument(int(format))
	args := []gi.Argument{arg_formats, arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_get_main_executable_path
//
// [ result ] trans: nothing
//
func GetMainExecutablePath() (result string) {
	iv, err := _I.Get(1244, "get_main_executable_path", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_init
//
// [ argc ] trans: everything, dir: inout
//
// [ argv ] trans: everything, dir: inout
//
func Init(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) {
	iv, err := _I.Get(1245, "init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	iv.Call(nil, nil, &outArgs[0])
}

// gst_init_check
//
// [ argc ] trans: everything, dir: inout
//
// [ argv ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func InitCheck(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) (result bool, err error) {
	iv, err := _I.Get(1246, "init_check", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gst_is_caps_features
//
// [ obj ] trans: nothing
//
// [ result ] trans: nothing
//
func IsCapsFeatures(obj unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1247, "is_caps_features", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_obj := gi.NewPointerArgument(obj)
	args := []gi.Argument{arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_is_initialized
//
// [ result ] trans: nothing
//
func IsInitialized() (result bool) {
	iv, err := _I.Get(1248, "is_initialized", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gst_library_error_quark
//
// [ result ] trans: nothing
//
func LibraryErrorQuark() (result uint32) {
	iv, err := _I.Get(1249, "library_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_message_type_get_name
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func MessageTypeGetName(type1 MessageTypeFlags) (result string) {
	iv, err := _I.Get(1250, "message_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_message_type_to_quark
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func MessageTypeToQuark(type1 MessageTypeFlags) (result uint32) {
	iv, err := _I.Get(1251, "message_type_to_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_meta_api_type_get_tags
//
// [ api ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaApiTypeGetTags(api gi.GType) (result gi.CStrArray) {
	iv, err := _I.Get(1252, "meta_api_type_get_tags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_api := gi.NewUintArgument(uint(api))
	args := []gi.Argument{arg_api}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gst_meta_api_type_has_tag
//
// [ api ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaApiTypeHasTag(api gi.GType, tag uint32) (result bool) {
	iv, err := _I.Get(1253, "meta_api_type_has_tag", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_api := gi.NewUintArgument(uint(api))
	arg_tag := gi.NewUint32Argument(tag)
	args := []gi.Argument{arg_api, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_meta_api_type_register
//
// [ api ] trans: nothing
//
// [ tags ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaApiTypeRegister(api string, tags gi.CStrArray) (result gi.GType) {
	iv, err := _I.Get(1254, "meta_api_type_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_api := gi.CString(api)
	arg_api := gi.NewStringArgument(c_api)
	arg_tags := gi.NewPointerArgument(tags.P)
	args := []gi.Argument{arg_api, arg_tags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_api)
	result = gi.GType(ret.Uint())
	return
}

// gst_meta_get_info
//
// [ impl ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaGetInfo(impl string) (result MetaInfo) {
	iv, err := _I.Get(1255, "meta_get_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_impl := gi.CString(impl)
	arg_impl := gi.NewStringArgument(c_impl)
	args := []gi.Argument{arg_impl}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_impl)
	result.P = ret.Pointer()
	return
}

// gst_meta_register
//
// [ api ] trans: nothing
//
// [ impl ] trans: nothing
//
// [ size ] trans: nothing
//
// [ init_func ] trans: nothing
//
// [ free_func ] trans: nothing
//
// [ transform_func ] trans: nothing
//
// [ result ] trans: nothing
//
func MetaRegister(api gi.GType, impl string, size uint64, init_func int /*TODO_TYPE CALLBACK*/, free_func int /*TODO_TYPE CALLBACK*/, transform_func int /*TODO_TYPE CALLBACK*/) (result MetaInfo) {
	iv, err := _I.Get(1256, "meta_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_impl := gi.CString(impl)
	arg_api := gi.NewUintArgument(uint(api))
	arg_impl := gi.NewStringArgument(c_impl)
	arg_size := gi.NewUint64Argument(size)
	arg_init_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMetaInitFunction()))
	arg_free_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMetaFreeFunction()))
	arg_transform_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMetaTransformFunction()))
	args := []gi.Argument{arg_api, arg_impl, arg_size, arg_init_func, arg_free_func, arg_transform_func}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_impl)
	result.P = ret.Pointer()
	return
}

// gst_mini_object_replace
//
// [ olddata ] trans: everything, dir: inout
//
// [ newdata ] trans: nothing
//
// [ result ] trans: nothing
//
func MiniObjectReplace(olddata int /*TODO:TYPE*/, newdata MiniObject) (result bool) {
	iv, err := _I.Get(1257, "mini_object_replace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_newdata := gi.NewPointerArgument(newdata.P)
	args := []gi.Argument{arg_newdata}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// gst_mini_object_take
//
// [ olddata ] trans: everything, dir: inout
//
// [ newdata ] trans: nothing
//
// [ result ] trans: nothing
//
func MiniObjectTake(olddata int /*TODO:TYPE*/, newdata MiniObject) (result bool) {
	iv, err := _I.Get(1258, "mini_object_take", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_newdata := gi.NewPointerArgument(newdata.P)
	args := []gi.Argument{arg_newdata}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// gst_pad_mode_get_name
//
// [ mode ] trans: nothing
//
// [ result ] trans: nothing
//
func PadModeGetName(mode PadModeEnum) (result string) {
	iv, err := _I.Get(1259, "pad_mode_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_param_spec_array
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ element_spec ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecArrayF(name string, nick string, blurb string, element_spec g.IParamSpec, flags g.ParamFlags) (result g.ParamSpec) {
	iv, err := _I.Get(1260, "param_spec_array", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	var tmp unsafe.Pointer
	if element_spec != nil {
		tmp = element_spec.P_ParamSpec()
	}
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_element_spec := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_element_spec, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// gst_param_spec_fraction
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ min_num ] trans: nothing
//
// [ min_denom ] trans: nothing
//
// [ max_num ] trans: nothing
//
// [ max_denom ] trans: nothing
//
// [ default_num ] trans: nothing
//
// [ default_denom ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecFractionF(name string, nick string, blurb string, min_num int32, min_denom int32, max_num int32, max_denom int32, default_num int32, default_denom int32, flags g.ParamFlags) (result g.ParamSpec) {
	iv, err := _I.Get(1261, "param_spec_fraction", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_min_num := gi.NewInt32Argument(min_num)
	arg_min_denom := gi.NewInt32Argument(min_denom)
	arg_max_num := gi.NewInt32Argument(max_num)
	arg_max_denom := gi.NewInt32Argument(max_denom)
	arg_default_num := gi.NewInt32Argument(default_num)
	arg_default_denom := gi.NewInt32Argument(default_denom)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_min_num, arg_min_denom, arg_max_num, arg_max_denom, arg_default_num, arg_default_denom, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// gst_parent_buffer_meta_api_get_type
//
// [ result ] trans: nothing
//
func ParentBufferMetaApiGetType() (result gi.GType) {
	iv, err := _I.Get(1262, "parent_buffer_meta_api_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_parent_buffer_meta_get_info
//
// [ result ] trans: nothing
//
func ParentBufferMetaGetInfo() (result MetaInfo) {
	iv, err := _I.Get(1263, "parent_buffer_meta_get_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_parse_bin_from_description
//
// [ bin_description ] trans: nothing
//
// [ ghost_unlinked_pads ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseBinFromDescription(bin_description string, ghost_unlinked_pads bool) (result Bin, err error) {
	iv, err := _I.Get(1264, "parse_bin_from_description", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_bin_description := gi.CString(bin_description)
	arg_bin_description := gi.NewStringArgument(c_bin_description)
	arg_ghost_unlinked_pads := gi.NewBoolArgument(ghost_unlinked_pads)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bin_description, arg_ghost_unlinked_pads, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_bin_description)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_parse_bin_from_description_full
//
// [ bin_description ] trans: nothing
//
// [ ghost_unlinked_pads ] trans: nothing
//
// [ context ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseBinFromDescriptionFull(bin_description string, ghost_unlinked_pads bool, context ParseContext, flags ParseFlags) (result Element, err error) {
	iv, err := _I.Get(1265, "parse_bin_from_description_full", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_bin_description := gi.CString(bin_description)
	arg_bin_description := gi.NewStringArgument(c_bin_description)
	arg_ghost_unlinked_pads := gi.NewBoolArgument(ghost_unlinked_pads)
	arg_context := gi.NewPointerArgument(context.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bin_description, arg_ghost_unlinked_pads, arg_context, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_bin_description)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_parse_error_quark
//
// [ result ] trans: nothing
//
func ParseErrorQuark() (result uint32) {
	iv, err := _I.Get(1266, "parse_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_parse_launch
//
// [ pipeline_description ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseLaunch(pipeline_description string) (result Element, err error) {
	iv, err := _I.Get(1267, "parse_launch", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_pipeline_description := gi.CString(pipeline_description)
	arg_pipeline_description := gi.NewStringArgument(c_pipeline_description)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pipeline_description, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_pipeline_description)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_parse_launch_full
//
// [ pipeline_description ] trans: nothing
//
// [ context ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseLaunchFull(pipeline_description string, context ParseContext, flags ParseFlags) (result Element, err error) {
	iv, err := _I.Get(1268, "parse_launch_full", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_pipeline_description := gi.CString(pipeline_description)
	arg_pipeline_description := gi.NewStringArgument(c_pipeline_description)
	arg_context := gi.NewPointerArgument(context.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pipeline_description, arg_context, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_pipeline_description)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_parse_launchv
//
// [ argv ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseLaunchv(argv gi.CStrArray) (result Element, err error) {
	iv, err := _I.Get(1269, "parse_launchv", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_argv, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_parse_launchv_full
//
// [ argv ] trans: nothing
//
// [ context ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func ParseLaunchvFull(argv gi.CStrArray, context ParseContext, flags ParseFlags) (result Element, err error) {
	iv, err := _I.Get(1270, "parse_launchv_full", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_context := gi.NewPointerArgument(context.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_argv, arg_context, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gst_plugin_error_quark
//
// [ result ] trans: nothing
//
func PluginErrorQuark() (result uint32) {
	iv, err := _I.Get(1271, "plugin_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_preset_get_app_dir
//
// [ result ] trans: nothing
//
func PresetGetAppDir() (result string) {
	iv, err := _I.Get(1272, "preset_get_app_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_preset_set_app_dir
//
// [ app_dir ] trans: nothing
//
// [ result ] trans: nothing
//
func PresetSetAppDir(app_dir string) (result bool) {
	iv, err := _I.Get(1273, "preset_set_app_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_app_dir := gi.CString(app_dir)
	arg_app_dir := gi.NewStringArgument(c_app_dir)
	args := []gi.Argument{arg_app_dir}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_app_dir)
	result = ret.Bool()
	return
}

// gst_protection_filter_systems_by_available_decryptors
//
// [ system_identifiers ] trans: nothing
//
// [ result ] trans: everything
//
func ProtectionFilterSystemsByAvailableDecryptors(system_identifiers gi.CStrArray) (result gi.CStrArray) {
	iv, err := _I.Get(1274, "protection_filter_systems_by_available_decryptors", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_system_identifiers := gi.NewPointerArgument(system_identifiers.P)
	args := []gi.Argument{arg_system_identifiers}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gst_protection_meta_api_get_type
//
// [ result ] trans: nothing
//
func ProtectionMetaApiGetType() (result gi.GType) {
	iv, err := _I.Get(1275, "protection_meta_api_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_protection_meta_get_info
//
// [ result ] trans: nothing
//
func ProtectionMetaGetInfo() (result MetaInfo) {
	iv, err := _I.Get(1276, "protection_meta_get_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_protection_select_system
//
// [ system_identifiers ] trans: nothing
//
// [ result ] trans: nothing
//
func ProtectionSelectSystem(system_identifiers gi.CStrArray) (result string) {
	iv, err := _I.Get(1277, "protection_select_system", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_system_identifiers := gi.NewPointerArgument(system_identifiers.P)
	args := []gi.Argument{arg_system_identifiers}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_query_type_get_flags
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func QueryTypeGetFlags(type1 QueryTypeEnum) (result QueryTypeFlags) {
	iv, err := _I.Get(1278, "query_type_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = QueryTypeFlags(ret.Int())
	return
}

// gst_query_type_get_name
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func QueryTypeGetName(type1 QueryTypeEnum) (result string) {
	iv, err := _I.Get(1279, "query_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_query_type_to_quark
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func QueryTypeToQuark(type1 QueryTypeEnum) (result uint32) {
	iv, err := _I.Get(1280, "query_type_to_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_reference_timestamp_meta_api_get_type
//
// [ result ] trans: nothing
//
func ReferenceTimestampMetaApiGetType() (result gi.GType) {
	iv, err := _I.Get(1281, "reference_timestamp_meta_api_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_reference_timestamp_meta_get_info
//
// [ result ] trans: nothing
//
func ReferenceTimestampMetaGetInfo() (result MetaInfo) {
	iv, err := _I.Get(1282, "reference_timestamp_meta_get_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_resource_error_quark
//
// [ result ] trans: nothing
//
func ResourceErrorQuark() (result uint32) {
	iv, err := _I.Get(1283, "resource_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_segtrap_is_enabled
//
// [ result ] trans: nothing
//
func SegtrapIsEnabled() (result bool) {
	iv, err := _I.Get(1284, "segtrap_is_enabled", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gst_segtrap_set_enabled
//
// [ enabled ] trans: nothing
//
func SegtrapSetEnabled(enabled bool) {
	iv, err := _I.Get(1285, "segtrap_set_enabled", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_enabled}
	iv.Call(args, nil, nil)
}

// gst_state_change_get_name
//
// [ transition ] trans: nothing
//
// [ result ] trans: nothing
//
func StateChangeGetName(transition StateChangeEnum) (result string) {
	iv, err := _I.Get(1286, "state_change_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_transition := gi.NewIntArgument(int(transition))
	args := []gi.Argument{arg_transition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_static_caps_get_type
//
// [ result ] trans: nothing
//
func StaticCapsGetType() (result gi.GType) {
	iv, err := _I.Get(1287, "static_caps_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_static_pad_template_get_type
//
// [ result ] trans: nothing
//
func StaticPadTemplateGetType() (result gi.GType) {
	iv, err := _I.Get(1288, "static_pad_template_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_stream_error_quark
//
// [ result ] trans: nothing
//
func StreamErrorQuark() (result uint32) {
	iv, err := _I.Get(1289, "stream_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_stream_type_get_name
//
// [ stype ] trans: nothing
//
// [ result ] trans: nothing
//
func StreamTypeGetName(stype StreamTypeFlags) (result string) {
	iv, err := _I.Get(1290, "stream_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stype := gi.NewIntArgument(int(stype))
	args := []gi.Argument{arg_stype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_structure_from_string
//
// [ string ] trans: nothing
//
// [ end ] trans: nothing, dir: out
//
// [ result ] trans: everything
//
func StructureFromString(string string) (result Structure, end string) {
	iv, err := _I.Get(1291, "structure_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	end = outArgs[0].String().Copy()
	result.P = ret.Pointer()
	return
}

// gst_tag_exists
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagExists(tag string) (result bool) {
	iv, err := _I.Get(1292, "tag_exists", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.Bool()
	return
}

// gst_tag_get_description
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagGetDescription(tag string) (result string) {
	iv, err := _I.Get(1293, "tag_get_description", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.String().Copy()
	return
}

// gst_tag_get_flag
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagGetFlag(tag string) (result TagFlagEnum) {
	iv, err := _I.Get(1294, "tag_get_flag", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = TagFlagEnum(ret.Int())
	return
}

// gst_tag_get_nick
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagGetNick(tag string) (result string) {
	iv, err := _I.Get(1295, "tag_get_nick", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.String().Copy()
	return
}

// gst_tag_get_type
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagGetType(tag string) (result gi.GType) {
	iv, err := _I.Get(1296, "tag_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = gi.GType(ret.Uint())
	return
}

// gst_tag_is_fixed
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagIsFixed(tag string) (result bool) {
	iv, err := _I.Get(1297, "tag_is_fixed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.Bool()
	return
}

// gst_tag_list_copy_value
//
// [ dest ] trans: nothing, dir: out
//
// [ list ] trans: nothing
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func TagListCopyValue(dest g.Value, list TagList, tag string) (result bool) {
	iv, err := _I.Get(1298, "tag_list_copy_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_list := gi.NewPointerArgument(list.P)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_dest, arg_list, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result = ret.Bool()
	return
}

// gst_tag_merge_strings_with_comma
//
// [ dest ] trans: nothing, dir: out
//
// [ src ] trans: nothing
//
func TagMergeStringsWithComma(dest g.Value, src g.Value) {
	iv, err := _I.Get(1299, "tag_merge_strings_with_comma", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_src := gi.NewPointerArgument(src.P)
	args := []gi.Argument{arg_dest, arg_src}
	iv.Call(args, nil, nil)
}

// gst_tag_merge_use_first
//
// [ dest ] trans: nothing, dir: out
//
// [ src ] trans: nothing
//
func TagMergeUseFirst(dest g.Value, src g.Value) {
	iv, err := _I.Get(1300, "tag_merge_use_first", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_src := gi.NewPointerArgument(src.P)
	args := []gi.Argument{arg_dest, arg_src}
	iv.Call(args, nil, nil)
}

// gst_tag_register
//
// [ name ] trans: nothing
//
// [ flag ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ func1 ] trans: nothing
//
func TagRegister(name string, flag TagFlagEnum, type1 gi.GType, nick string, blurb string, func1 int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1301, "tag_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_flag := gi.NewIntArgument(int(flag))
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTagMergeFunc()))
	args := []gi.Argument{arg_name, arg_flag, arg_type1, arg_nick, arg_blurb, arg_func1}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
}

// gst_tag_register_static
//
// [ name ] trans: nothing
//
// [ flag ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ func1 ] trans: nothing
//
func TagRegisterStatic(name string, flag TagFlagEnum, type1 gi.GType, nick string, blurb string, func1 int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(1302, "tag_register_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_flag := gi.NewIntArgument(int(flag))
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTagMergeFunc()))
	args := []gi.Argument{arg_name, arg_flag, arg_type1, arg_nick, arg_blurb, arg_func1}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
}

// gst_toc_entry_type_get_nick
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TocEntryTypeGetNick(type1 TocEntryTypeEnum) (result string) {
	iv, err := _I.Get(1303, "toc_entry_type_get_nick", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gst_type_find_get_type
//
// [ result ] trans: nothing
//
func TypeFindGetType() (result gi.GType) {
	iv, err := _I.Get(1304, "type_find_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_type_find_register
//
// [ plugin ] trans: nothing
//
// [ name ] trans: nothing
//
// [ rank ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ extensions ] trans: nothing
//
// [ possible_caps ] trans: nothing
//
// [ data ] trans: nothing
//
// [ data_notify ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeFindRegister(plugin IPlugin, name string, rank uint32, func1 int /*TODO_TYPE CALLBACK*/, extensions string, possible_caps Caps, data unsafe.Pointer, data_notify int /*TODO_TYPE CALLBACK*/) (result bool) {
	iv, err := _I.Get(1305, "type_find_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_Plugin()
	}
	c_name := gi.CString(name)
	c_extensions := gi.CString(extensions)
	arg_plugin := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_rank := gi.NewUint32Argument(rank)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myTypeFindFunction()))
	arg_extensions := gi.NewStringArgument(c_extensions)
	arg_possible_caps := gi.NewPointerArgument(possible_caps.P)
	arg_data := gi.NewPointerArgument(data)
	arg_data_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_plugin, arg_name, arg_rank, arg_func1, arg_extensions, arg_possible_caps, arg_data, arg_data_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_extensions)
	result = ret.Bool()
	return
}

// gst_update_registry
//
// [ result ] trans: nothing
//
func UpdateRegistry() (result bool) {
	iv, err := _I.Get(1306, "update_registry", "")
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
// gst_uri_construct
//
// [ protocol ] trans: nothing
//
// [ location ] trans: nothing
//
// [ result ] trans: everything
//
func UriConstruct(protocol string, location string) (result string) {
	iv, err := _I.Get(1307, "uri_construct", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_protocol := gi.CString(protocol)
	c_location := gi.CString(location)
	arg_protocol := gi.NewStringArgument(c_protocol)
	arg_location := gi.NewStringArgument(c_location)
	args := []gi.Argument{arg_protocol, arg_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_protocol)
	gi.Free(c_location)
	result = ret.String().Take()
	return
}

// gst_uri_error_quark
//
// [ result ] trans: nothing
//
func UriErrorQuark() (result uint32) {
	iv, err := _I.Get(1308, "uri_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_uri_from_string
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriFromString(uri string) (result Uri) {
	iv, err := _I.Get(1309, "uri_from_string", "")
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
	result.P = ret.Pointer()
	return
}

// gst_uri_get_location
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriGetLocation(uri string) (result string) {
	iv, err := _I.Get(1310, "uri_get_location", "")
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

// gst_uri_get_protocol
//
// [ uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriGetProtocol(uri string) (result string) {
	iv, err := _I.Get(1311, "uri_get_protocol", "")
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

// gst_uri_has_protocol
//
// [ uri ] trans: nothing
//
// [ protocol ] trans: nothing
//
// [ result ] trans: nothing
//
func UriHasProtocol(uri string, protocol string) (result bool) {
	iv, err := _I.Get(1312, "uri_has_protocol", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_protocol := gi.CString(protocol)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_protocol := gi.NewStringArgument(c_protocol)
	args := []gi.Argument{arg_uri, arg_protocol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri)
	gi.Free(c_protocol)
	result = ret.Bool()
	return
}

// gst_uri_is_valid
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func UriIsValid(uri string) (result bool) {
	iv, err := _I.Get(1313, "uri_is_valid", "")
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
	result = ret.Bool()
	return
}

// gst_uri_join_strings
//
// [ base_uri ] trans: nothing
//
// [ ref_uri ] trans: nothing
//
// [ result ] trans: everything
//
func UriJoinStrings(base_uri string, ref_uri string) (result string) {
	iv, err := _I.Get(1314, "uri_join_strings", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_base_uri := gi.CString(base_uri)
	c_ref_uri := gi.CString(ref_uri)
	arg_base_uri := gi.NewStringArgument(c_base_uri)
	arg_ref_uri := gi.NewStringArgument(c_ref_uri)
	args := []gi.Argument{arg_base_uri, arg_ref_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_base_uri)
	gi.Free(c_ref_uri)
	result = ret.String().Take()
	return
}

// gst_uri_protocol_is_supported
//
// [ type1 ] trans: nothing
//
// [ protocol ] trans: nothing
//
// [ result ] trans: nothing
//
func UriProtocolIsSupported(type1 URITypeEnum, protocol string) (result bool) {
	iv, err := _I.Get(1315, "uri_protocol_is_supported", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_protocol := gi.CString(protocol)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_protocol := gi.NewStringArgument(c_protocol)
	args := []gi.Argument{arg_type1, arg_protocol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_protocol)
	result = ret.Bool()
	return
}

// gst_uri_protocol_is_valid
//
// [ protocol ] trans: nothing
//
// [ result ] trans: nothing
//
func UriProtocolIsValid(protocol string) (result bool) {
	iv, err := _I.Get(1316, "uri_protocol_is_valid", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_protocol := gi.CString(protocol)
	arg_protocol := gi.NewStringArgument(c_protocol)
	args := []gi.Argument{arg_protocol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_protocol)
	result = ret.Bool()
	return
}

// gst_util_array_binary_search
//
// [ array ] trans: nothing
//
// [ num_elements ] trans: nothing
//
// [ element_size ] trans: nothing
//
// [ search_func ] trans: nothing
//
// [ mode ] trans: nothing
//
// [ search_data ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilArrayBinarySearch(array unsafe.Pointer, num_elements uint32, element_size uint64, search_func int /*TODO_TYPE CALLBACK*/, mode SearchModeEnum, search_data unsafe.Pointer, user_data unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1317, "util_array_binary_search", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array)
	arg_num_elements := gi.NewUint32Argument(num_elements)
	arg_element_size := gi.NewUint64Argument(element_size)
	arg_search_func := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myCompareDataFunc()))
	arg_mode := gi.NewIntArgument(int(mode))
	arg_search_data := gi.NewPointerArgument(search_data)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_array, arg_num_elements, arg_element_size, arg_search_func, arg_mode, arg_search_data, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// gst_util_double_to_fraction
//
// [ src ] trans: nothing
//
// [ dest_n ] trans: everything, dir: out
//
// [ dest_d ] trans: everything, dir: out
//
func UtilDoubleToFraction(src float64) (dest_n int32, dest_d int32) {
	iv, err := _I.Get(1318, "util_double_to_fraction", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_src := gi.NewDoubleArgument(src)
	arg_dest_n := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_dest_d := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_src, arg_dest_n, arg_dest_d}
	iv.Call(args, nil, &outArgs[0])
	dest_n = outArgs[0].Int32()
	dest_d = outArgs[1].Int32()
	return
}

// gst_util_dump_buffer
//
// [ buf ] trans: nothing
//
func UtilDumpBuffer(buf Buffer) {
	iv, err := _I.Get(1319, "util_dump_buffer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// gst_util_dump_mem
//
// [ mem ] trans: nothing
//
// [ size ] trans: nothing
//
func UtilDumpMem(mem gi.Uint8Array, size uint32) {
	iv, err := _I.Get(1320, "util_dump_mem", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem.P)
	arg_size := gi.NewUint32Argument(size)
	args := []gi.Argument{arg_mem, arg_size}
	iv.Call(args, nil, nil)
}

// gst_util_fraction_add
//
// [ a_n ] trans: nothing
//
// [ a_d ] trans: nothing
//
// [ b_n ] trans: nothing
//
// [ b_d ] trans: nothing
//
// [ res_n ] trans: everything, dir: out
//
// [ res_d ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func UtilFractionAdd(a_n int32, a_d int32, b_n int32, b_d int32) (result bool, res_n int32, res_d int32) {
	iv, err := _I.Get(1321, "util_fraction_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_a_n := gi.NewInt32Argument(a_n)
	arg_a_d := gi.NewInt32Argument(a_d)
	arg_b_n := gi.NewInt32Argument(b_n)
	arg_b_d := gi.NewInt32Argument(b_d)
	arg_res_n := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res_d := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_a_n, arg_a_d, arg_b_n, arg_b_d, arg_res_n, arg_res_d}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	res_n = outArgs[0].Int32()
	res_d = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// gst_util_fraction_compare
//
// [ a_n ] trans: nothing
//
// [ a_d ] trans: nothing
//
// [ b_n ] trans: nothing
//
// [ b_d ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilFractionCompare(a_n int32, a_d int32, b_n int32, b_d int32) (result int32) {
	iv, err := _I.Get(1322, "util_fraction_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a_n := gi.NewInt32Argument(a_n)
	arg_a_d := gi.NewInt32Argument(a_d)
	arg_b_n := gi.NewInt32Argument(b_n)
	arg_b_d := gi.NewInt32Argument(b_d)
	args := []gi.Argument{arg_a_n, arg_a_d, arg_b_n, arg_b_d}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_util_fraction_multiply
//
// [ a_n ] trans: nothing
//
// [ a_d ] trans: nothing
//
// [ b_n ] trans: nothing
//
// [ b_d ] trans: nothing
//
// [ res_n ] trans: everything, dir: out
//
// [ res_d ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func UtilFractionMultiply(a_n int32, a_d int32, b_n int32, b_d int32) (result bool, res_n int32, res_d int32) {
	iv, err := _I.Get(1323, "util_fraction_multiply", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_a_n := gi.NewInt32Argument(a_n)
	arg_a_d := gi.NewInt32Argument(a_d)
	arg_b_n := gi.NewInt32Argument(b_n)
	arg_b_d := gi.NewInt32Argument(b_d)
	arg_res_n := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res_d := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_a_n, arg_a_d, arg_b_n, arg_b_d, arg_res_n, arg_res_d}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	res_n = outArgs[0].Int32()
	res_d = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// gst_util_fraction_to_double
//
// [ src_n ] trans: nothing
//
// [ src_d ] trans: nothing
//
// [ dest ] trans: everything, dir: out
//
func UtilFractionToDouble(src_n int32, src_d int32) (dest float64) {
	iv, err := _I.Get(1324, "util_fraction_to_double", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_src_n := gi.NewInt32Argument(src_n)
	arg_src_d := gi.NewInt32Argument(src_d)
	arg_dest := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_src_n, arg_src_d, arg_dest}
	iv.Call(args, nil, &outArgs[0])
	dest = outArgs[0].Double()
	return
}

// gst_util_gdouble_to_guint64
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilGdoubleToGuint64(value float64) (result uint64) {
	iv, err := _I.Get(1325, "util_gdouble_to_guint64", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewDoubleArgument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_get_object_array
//
// [ object ] trans: nothing
//
// [ name ] trans: nothing
//
// [ array ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func UtilGetObjectArray(object g.IObject, name string) (result bool, array g.ValueArray) {
	iv, err := _I.Get(1326, "util_get_object_array", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_name := gi.CString(name)
	arg_object := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_array := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_object, arg_name, arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	array.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// gst_util_get_timestamp
//
// [ result ] trans: nothing
//
func UtilGetTimestamp() (result uint64) {
	iv, err := _I.Get(1327, "util_get_timestamp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_greatest_common_divisor
//
// [ a ] trans: nothing
//
// [ b ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilGreatestCommonDivisor(a int32, b int32) (result int32) {
	iv, err := _I.Get(1328, "util_greatest_common_divisor", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewInt32Argument(a)
	arg_b := gi.NewInt32Argument(b)
	args := []gi.Argument{arg_a, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_util_greatest_common_divisor_int64
//
// [ a ] trans: nothing
//
// [ b ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilGreatestCommonDivisorInt64(a int64, b int64) (result int64) {
	iv, err := _I.Get(1329, "util_greatest_common_divisor_int64", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewInt64Argument(a)
	arg_b := gi.NewInt64Argument(b)
	args := []gi.Argument{arg_a, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gst_util_group_id_next
//
// [ result ] trans: nothing
//
func UtilGroupIdNext() (result uint32) {
	iv, err := _I.Get(1330, "util_group_id_next", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_util_guint64_to_gdouble
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilGuint64ToGdouble(value uint64) (result float64) {
	iv, err := _I.Get(1331, "util_guint64_to_gdouble", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint64Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gst_util_seqnum_compare
//
// [ s1 ] trans: nothing
//
// [ s2 ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilSeqnumCompare(s1 uint32, s2 uint32) (result int32) {
	iv, err := _I.Get(1332, "util_seqnum_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_s1 := gi.NewUint32Argument(s1)
	arg_s2 := gi.NewUint32Argument(s2)
	args := []gi.Argument{arg_s1, arg_s2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_util_seqnum_next
//
// [ result ] trans: nothing
//
func UtilSeqnumNext() (result uint32) {
	iv, err := _I.Get(1333, "util_seqnum_next", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_util_set_object_arg
//
// [ object ] trans: nothing
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
func UtilSetObjectArg(object g.IObject, name string, value string) {
	iv, err := _I.Get(1334, "util_set_object_arg", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_name := gi.CString(name)
	c_value := gi.CString(value)
	arg_object := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_object, arg_name, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_value)
}

// gst_util_set_object_array
//
// [ object ] trans: nothing
//
// [ name ] trans: nothing
//
// [ array ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilSetObjectArray(object g.IObject, name string, array g.ValueArray) (result bool) {
	iv, err := _I.Get(1335, "util_set_object_array", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	c_name := gi.CString(name)
	arg_object := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_object, arg_name, arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gst_util_set_value_from_string
//
// [ value ] trans: nothing, dir: out
//
// [ value_str ] trans: nothing
//
func UtilSetValueFromString(value g.Value, value_str string) {
	iv, err := _I.Get(1336, "util_set_value_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_value_str := gi.CString(value_str)
	arg_value := gi.NewPointerArgument(value.P)
	arg_value_str := gi.NewStringArgument(c_value_str)
	args := []gi.Argument{arg_value, arg_value_str}
	iv.Call(args, nil, nil)
	gi.Free(c_value_str)
}

// gst_util_uint64_scale
//
// [ val ] trans: nothing
//
// [ num ] trans: nothing
//
// [ denom ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilUint64Scale(val uint64, num uint64, denom uint64) (result uint64) {
	iv, err := _I.Get(1337, "util_uint64_scale", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewUint64Argument(val)
	arg_num := gi.NewUint64Argument(num)
	arg_denom := gi.NewUint64Argument(denom)
	args := []gi.Argument{arg_val, arg_num, arg_denom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_uint64_scale_ceil
//
// [ val ] trans: nothing
//
// [ num ] trans: nothing
//
// [ denom ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilUint64ScaleCeil(val uint64, num uint64, denom uint64) (result uint64) {
	iv, err := _I.Get(1338, "util_uint64_scale_ceil", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewUint64Argument(val)
	arg_num := gi.NewUint64Argument(num)
	arg_denom := gi.NewUint64Argument(denom)
	args := []gi.Argument{arg_val, arg_num, arg_denom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_uint64_scale_int
//
// [ val ] trans: nothing
//
// [ num ] trans: nothing
//
// [ denom ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilUint64ScaleInt(val uint64, num int32, denom int32) (result uint64) {
	iv, err := _I.Get(1339, "util_uint64_scale_int", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewUint64Argument(val)
	arg_num := gi.NewInt32Argument(num)
	arg_denom := gi.NewInt32Argument(denom)
	args := []gi.Argument{arg_val, arg_num, arg_denom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_uint64_scale_int_ceil
//
// [ val ] trans: nothing
//
// [ num ] trans: nothing
//
// [ denom ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilUint64ScaleIntCeil(val uint64, num int32, denom int32) (result uint64) {
	iv, err := _I.Get(1340, "util_uint64_scale_int_ceil", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewUint64Argument(val)
	arg_num := gi.NewInt32Argument(num)
	arg_denom := gi.NewInt32Argument(denom)
	args := []gi.Argument{arg_val, arg_num, arg_denom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_uint64_scale_int_round
//
// [ val ] trans: nothing
//
// [ num ] trans: nothing
//
// [ denom ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilUint64ScaleIntRound(val uint64, num int32, denom int32) (result uint64) {
	iv, err := _I.Get(1341, "util_uint64_scale_int_round", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewUint64Argument(val)
	arg_num := gi.NewInt32Argument(num)
	arg_denom := gi.NewInt32Argument(denom)
	args := []gi.Argument{arg_val, arg_num, arg_denom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_util_uint64_scale_round
//
// [ val ] trans: nothing
//
// [ num ] trans: nothing
//
// [ denom ] trans: nothing
//
// [ result ] trans: nothing
//
func UtilUint64ScaleRound(val uint64, num uint64, denom uint64) (result uint64) {
	iv, err := _I.Get(1342, "util_uint64_scale_round", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewUint64Argument(val)
	arg_num := gi.NewUint64Argument(num)
	arg_denom := gi.NewUint64Argument(denom)
	args := []gi.Argument{arg_val, arg_num, arg_denom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_value_can_compare
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueCanCompare(value1 g.Value, value2 g.Value) (result bool) {
	iv, err := _I.Get(1343, "value_can_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_can_intersect
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueCanIntersect(value1 g.Value, value2 g.Value) (result bool) {
	iv, err := _I.Get(1344, "value_can_intersect", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_can_subtract
//
// [ minuend ] trans: nothing
//
// [ subtrahend ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueCanSubtract(minuend g.Value, subtrahend g.Value) (result bool) {
	iv, err := _I.Get(1345, "value_can_subtract", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_minuend := gi.NewPointerArgument(minuend.P)
	arg_subtrahend := gi.NewPointerArgument(subtrahend.P)
	args := []gi.Argument{arg_minuend, arg_subtrahend}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_can_union
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueCanUnion(value1 g.Value, value2 g.Value) (result bool) {
	iv, err := _I.Get(1346, "value_can_union", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_compare
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueCompare(value1 g.Value, value2 g.Value) (result int32) {
	iv, err := _I.Get(1347, "value_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_value_deserialize
//
// [ dest ] trans: nothing, dir: out
//
// [ src ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueDeserialize(dest g.Value, src string) (result bool) {
	iv, err := _I.Get(1348, "value_deserialize", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_src := gi.CString(src)
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_src := gi.NewStringArgument(c_src)
	args := []gi.Argument{arg_dest, arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_src)
	result = ret.Bool()
	return
}

// gst_value_fixate
//
// [ dest ] trans: nothing
//
// [ src ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueFixate(dest g.Value, src g.Value) (result bool) {
	iv, err := _I.Get(1349, "value_fixate", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_src := gi.NewPointerArgument(src.P)
	args := []gi.Argument{arg_dest, arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_fraction_multiply
//
// [ product ] trans: nothing
//
// [ factor1 ] trans: nothing
//
// [ factor2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueFractionMultiply(product g.Value, factor1 g.Value, factor2 g.Value) (result bool) {
	iv, err := _I.Get(1350, "value_fraction_multiply", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_product := gi.NewPointerArgument(product.P)
	arg_factor1 := gi.NewPointerArgument(factor1.P)
	arg_factor2 := gi.NewPointerArgument(factor2.P)
	args := []gi.Argument{arg_product, arg_factor1, arg_factor2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_fraction_subtract
//
// [ dest ] trans: nothing
//
// [ minuend ] trans: nothing
//
// [ subtrahend ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueFractionSubtract(dest g.Value, minuend g.Value, subtrahend g.Value) (result bool) {
	iv, err := _I.Get(1351, "value_fraction_subtract", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_minuend := gi.NewPointerArgument(minuend.P)
	arg_subtrahend := gi.NewPointerArgument(subtrahend.P)
	args := []gi.Argument{arg_dest, arg_minuend, arg_subtrahend}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_get_bitmask
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetBitmask(value g.Value) (result uint64) {
	iv, err := _I.Get(1352, "value_get_bitmask", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_value_get_caps
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetCaps(value g.Value) (result Caps) {
	iv, err := _I.Get(1353, "value_get_caps", "")
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

// gst_value_get_caps_features
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetCapsFeatures(value g.Value) (result CapsFeatures) {
	iv, err := _I.Get(1354, "value_get_caps_features", "")
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

// gst_value_get_double_range_max
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetDoubleRangeMax(value g.Value) (result float64) {
	iv, err := _I.Get(1355, "value_get_double_range_max", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gst_value_get_double_range_min
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetDoubleRangeMin(value g.Value) (result float64) {
	iv, err := _I.Get(1356, "value_get_double_range_min", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gst_value_get_flagset_flags
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetFlagsetFlags(value g.Value) (result uint32) {
	iv, err := _I.Get(1357, "value_get_flagset_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_value_get_flagset_mask
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetFlagsetMask(value g.Value) (result uint32) {
	iv, err := _I.Get(1358, "value_get_flagset_mask", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gst_value_get_fraction_denominator
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetFractionDenominator(value g.Value) (result int32) {
	iv, err := _I.Get(1359, "value_get_fraction_denominator", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_value_get_fraction_numerator
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetFractionNumerator(value g.Value) (result int32) {
	iv, err := _I.Get(1360, "value_get_fraction_numerator", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_value_get_fraction_range_max
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetFractionRangeMax(value g.Value) (result g.Value) {
	iv, err := _I.Get(1361, "value_get_fraction_range_max", "")
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

// gst_value_get_fraction_range_min
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetFractionRangeMin(value g.Value) (result g.Value) {
	iv, err := _I.Get(1362, "value_get_fraction_range_min", "")
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

// gst_value_get_int64_range_max
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetInt64RangeMax(value g.Value) (result int64) {
	iv, err := _I.Get(1363, "value_get_int64_range_max", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gst_value_get_int64_range_min
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetInt64RangeMin(value g.Value) (result int64) {
	iv, err := _I.Get(1364, "value_get_int64_range_min", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gst_value_get_int64_range_step
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetInt64RangeStep(value g.Value) (result int64) {
	iv, err := _I.Get(1365, "value_get_int64_range_step", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gst_value_get_int_range_max
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetIntRangeMax(value g.Value) (result int32) {
	iv, err := _I.Get(1366, "value_get_int_range_max", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_value_get_int_range_min
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetIntRangeMin(value g.Value) (result int32) {
	iv, err := _I.Get(1367, "value_get_int_range_min", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_value_get_int_range_step
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetIntRangeStep(value g.Value) (result int32) {
	iv, err := _I.Get(1368, "value_get_int_range_step", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gst_value_get_structure
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueGetStructure(value g.Value) (result Structure) {
	iv, err := _I.Get(1369, "value_get_structure", "")
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

// gst_value_init_and_copy
//
// [ dest ] trans: nothing, dir: out
//
// [ src ] trans: nothing
//
func ValueInitAndCopy(dest g.Value, src g.Value) {
	iv, err := _I.Get(1370, "value_init_and_copy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_src := gi.NewPointerArgument(src.P)
	args := []gi.Argument{arg_dest, arg_src}
	iv.Call(args, nil, nil)
}

// gst_value_intersect
//
// [ dest ] trans: everything, dir: out
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueIntersect(dest g.Value, value1 g.Value, value2 g.Value) (result bool) {
	iv, err := _I.Get(1371, "value_intersect", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_dest, arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_is_fixed
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueIsFixed(value g.Value) (result bool) {
	iv, err := _I.Get(1372, "value_is_fixed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_is_subset
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueIsSubset(value1 g.Value, value2 g.Value) (result bool) {
	iv, err := _I.Get(1373, "value_is_subset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_register
//
// [ table ] trans: nothing
//
func ValueRegister(table ValueTable) {
	iv, err := _I.Get(1374, "value_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_table := gi.NewPointerArgument(table.P)
	args := []gi.Argument{arg_table}
	iv.Call(args, nil, nil)
}

// gst_value_serialize
//
// [ value ] trans: nothing
//
// [ result ] trans: everything
//
func ValueSerialize(value g.Value) (result string) {
	iv, err := _I.Get(1375, "value_serialize", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gst_value_set_bitmask
//
// [ value ] trans: nothing
//
// [ bitmask ] trans: nothing
//
func ValueSetBitmask(value g.Value, bitmask uint64) {
	iv, err := _I.Get(1376, "value_set_bitmask", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_bitmask := gi.NewUint64Argument(bitmask)
	args := []gi.Argument{arg_value, arg_bitmask}
	iv.Call(args, nil, nil)
}

// gst_value_set_caps
//
// [ value ] trans: nothing
//
// [ caps ] trans: nothing
//
func ValueSetCaps(value g.Value, caps Caps) {
	iv, err := _I.Get(1377, "value_set_caps", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_caps := gi.NewPointerArgument(caps.P)
	args := []gi.Argument{arg_value, arg_caps}
	iv.Call(args, nil, nil)
}

// gst_value_set_caps_features
//
// [ value ] trans: nothing
//
// [ features ] trans: nothing
//
func ValueSetCapsFeatures(value g.Value, features CapsFeatures) {
	iv, err := _I.Get(1378, "value_set_caps_features", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_features := gi.NewPointerArgument(features.P)
	args := []gi.Argument{arg_value, arg_features}
	iv.Call(args, nil, nil)
}

// gst_value_set_double_range
//
// [ value ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
func ValueSetDoubleRange(value g.Value, start float64, end float64) {
	iv, err := _I.Get(1379, "value_set_double_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_start := gi.NewDoubleArgument(start)
	arg_end := gi.NewDoubleArgument(end)
	args := []gi.Argument{arg_value, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gst_value_set_flagset
//
// [ value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ mask ] trans: nothing
//
func ValueSetFlagset(value g.Value, flags uint32, mask uint32) {
	iv, err := _I.Get(1380, "value_set_flagset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_flags := gi.NewUint32Argument(flags)
	arg_mask := gi.NewUint32Argument(mask)
	args := []gi.Argument{arg_value, arg_flags, arg_mask}
	iv.Call(args, nil, nil)
}

// gst_value_set_fraction
//
// [ value ] trans: nothing
//
// [ numerator ] trans: nothing
//
// [ denominator ] trans: nothing
//
func ValueSetFraction(value g.Value, numerator int32, denominator int32) {
	iv, err := _I.Get(1381, "value_set_fraction", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_numerator := gi.NewInt32Argument(numerator)
	arg_denominator := gi.NewInt32Argument(denominator)
	args := []gi.Argument{arg_value, arg_numerator, arg_denominator}
	iv.Call(args, nil, nil)
}

// gst_value_set_fraction_range
//
// [ value ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
func ValueSetFractionRange(value g.Value, start g.Value, end g.Value) {
	iv, err := _I.Get(1382, "value_set_fraction_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_value, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gst_value_set_fraction_range_full
//
// [ value ] trans: nothing
//
// [ numerator_start ] trans: nothing
//
// [ denominator_start ] trans: nothing
//
// [ numerator_end ] trans: nothing
//
// [ denominator_end ] trans: nothing
//
func ValueSetFractionRangeFull(value g.Value, numerator_start int32, denominator_start int32, numerator_end int32, denominator_end int32) {
	iv, err := _I.Get(1383, "value_set_fraction_range_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_numerator_start := gi.NewInt32Argument(numerator_start)
	arg_denominator_start := gi.NewInt32Argument(denominator_start)
	arg_numerator_end := gi.NewInt32Argument(numerator_end)
	arg_denominator_end := gi.NewInt32Argument(denominator_end)
	args := []gi.Argument{arg_value, arg_numerator_start, arg_denominator_start, arg_numerator_end, arg_denominator_end}
	iv.Call(args, nil, nil)
}

// gst_value_set_int64_range
//
// [ value ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
func ValueSetInt64Range(value g.Value, start int64, end int64) {
	iv, err := _I.Get(1384, "value_set_int64_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_start := gi.NewInt64Argument(start)
	arg_end := gi.NewInt64Argument(end)
	args := []gi.Argument{arg_value, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gst_value_set_int64_range_step
//
// [ value ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
// [ step ] trans: nothing
//
func ValueSetInt64RangeStep(value g.Value, start int64, end int64, step int64) {
	iv, err := _I.Get(1385, "value_set_int64_range_step", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_start := gi.NewInt64Argument(start)
	arg_end := gi.NewInt64Argument(end)
	arg_step := gi.NewInt64Argument(step)
	args := []gi.Argument{arg_value, arg_start, arg_end, arg_step}
	iv.Call(args, nil, nil)
}

// gst_value_set_int_range
//
// [ value ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
func ValueSetIntRange(value g.Value, start int32, end int32) {
	iv, err := _I.Get(1386, "value_set_int_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_start := gi.NewInt32Argument(start)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_value, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gst_value_set_int_range_step
//
// [ value ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
// [ step ] trans: nothing
//
func ValueSetIntRangeStep(value g.Value, start int32, end int32, step int32) {
	iv, err := _I.Get(1387, "value_set_int_range_step", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_start := gi.NewInt32Argument(start)
	arg_end := gi.NewInt32Argument(end)
	arg_step := gi.NewInt32Argument(step)
	args := []gi.Argument{arg_value, arg_start, arg_end, arg_step}
	iv.Call(args, nil, nil)
}

// gst_value_set_structure
//
// [ value ] trans: nothing
//
// [ structure ] trans: nothing
//
func ValueSetStructure(value g.Value, structure Structure) {
	iv, err := _I.Get(1388, "value_set_structure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_structure := gi.NewPointerArgument(structure.P)
	args := []gi.Argument{arg_value, arg_structure}
	iv.Call(args, nil, nil)
}

// gst_value_subtract
//
// [ dest ] trans: nothing, dir: out
//
// [ minuend ] trans: nothing
//
// [ subtrahend ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueSubtract(dest g.Value, minuend g.Value, subtrahend g.Value) (result bool) {
	iv, err := _I.Get(1389, "value_subtract", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_minuend := gi.NewPointerArgument(minuend.P)
	arg_subtrahend := gi.NewPointerArgument(subtrahend.P)
	args := []gi.Argument{arg_dest, arg_minuend, arg_subtrahend}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_value_union
//
// [ dest ] trans: nothing, dir: out
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueUnion(dest g.Value, value1 g.Value, value2 g.Value) (result bool) {
	iv, err := _I.Get(1390, "value_union", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_dest, arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_version
//
// [ major ] trans: everything, dir: out
//
// [ minor ] trans: everything, dir: out
//
// [ micro ] trans: everything, dir: out
//
// [ nano ] trans: everything, dir: out
//
func Version() (major uint32, minor uint32, micro uint32, nano uint32) {
	iv, err := _I.Get(1391, "version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_major := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_minor := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_micro := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_nano := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_major, arg_minor, arg_micro, arg_nano}
	iv.Call(args, nil, &outArgs[0])
	major = outArgs[0].Uint32()
	minor = outArgs[1].Uint32()
	micro = outArgs[2].Uint32()
	nano = outArgs[3].Uint32()
	return
}

// gst_version_string
//
// [ result ] trans: everything
//
func VersionString() (result string) {
	iv, err := _I.Get(1392, "version_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// constants
const (
	ALLOCATOR_SYSMEM                      = "SystemMemory"
	BUFFER_OFFSET_NONE                    = 18446744073709551615
	CAN_INLINE                            = 1
	CAPS_FEATURE_MEMORY_SYSTEM_MEMORY     = "memory:SystemMemory"
	CLOCK_TIME_NONE                       = 18446744073709551615
	DEBUG_BG_MASK                         = 240
	DEBUG_FG_MASK                         = 15
	DEBUG_FORMAT_MASK                     = 65280
	ELEMENT_FACTORY_KLASS_DECODER         = "Decoder"
	ELEMENT_FACTORY_KLASS_DECRYPTOR       = "Decryptor"
	ELEMENT_FACTORY_KLASS_DEMUXER         = "Demuxer"
	ELEMENT_FACTORY_KLASS_DEPAYLOADER     = "Depayloader"
	ELEMENT_FACTORY_KLASS_ENCODER         = "Encoder"
	ELEMENT_FACTORY_KLASS_ENCRYPTOR       = "Encryptor"
	ELEMENT_FACTORY_KLASS_FORMATTER       = "Formatter"
	ELEMENT_FACTORY_KLASS_MEDIA_AUDIO     = "Audio"
	ELEMENT_FACTORY_KLASS_MEDIA_IMAGE     = "Image"
	ELEMENT_FACTORY_KLASS_MEDIA_METADATA  = "Metadata"
	ELEMENT_FACTORY_KLASS_MEDIA_SUBTITLE  = "Subtitle"
	ELEMENT_FACTORY_KLASS_MEDIA_VIDEO     = "Video"
	ELEMENT_FACTORY_KLASS_MUXER           = "Muxer"
	ELEMENT_FACTORY_KLASS_PARSER          = "Parser"
	ELEMENT_FACTORY_KLASS_PAYLOADER       = "Payloader"
	ELEMENT_FACTORY_KLASS_SINK            = "Sink"
	ELEMENT_FACTORY_KLASS_SRC             = "Source"
	ELEMENT_FACTORY_TYPE_ANY              = 562949953421311
	ELEMENT_FACTORY_TYPE_AUDIOVIDEO_SINKS = 3940649673949188
	ELEMENT_FACTORY_TYPE_AUDIO_ENCODER    = 1125899906842626
	ELEMENT_FACTORY_TYPE_DECODABLE        = 1377
	ELEMENT_FACTORY_TYPE_DECODER          = 1
	ELEMENT_FACTORY_TYPE_DECRYPTOR        = 1024
	ELEMENT_FACTORY_TYPE_DEMUXER          = 32
	ELEMENT_FACTORY_TYPE_DEPAYLOADER      = 256
	ELEMENT_FACTORY_TYPE_ENCODER          = 2
	ELEMENT_FACTORY_TYPE_ENCRYPTOR        = 2048
	ELEMENT_FACTORY_TYPE_FORMATTER        = 512
	ELEMENT_FACTORY_TYPE_MAX_ELEMENTS     = 281474976710656
	ELEMENT_FACTORY_TYPE_MEDIA_ANY        = 18446462598732840960
	ELEMENT_FACTORY_TYPE_MEDIA_AUDIO      = 1125899906842624
	ELEMENT_FACTORY_TYPE_MEDIA_IMAGE      = 2251799813685248
	ELEMENT_FACTORY_TYPE_MEDIA_METADATA   = 9007199254740992
	ELEMENT_FACTORY_TYPE_MEDIA_SUBTITLE   = 4503599627370496
	ELEMENT_FACTORY_TYPE_MEDIA_VIDEO      = 562949953421312
	ELEMENT_FACTORY_TYPE_MUXER            = 16
	ELEMENT_FACTORY_TYPE_PARSER           = 64
	ELEMENT_FACTORY_TYPE_PAYLOADER        = 128
	ELEMENT_FACTORY_TYPE_SINK             = 4
	ELEMENT_FACTORY_TYPE_SRC              = 8
	ELEMENT_FACTORY_TYPE_VIDEO_ENCODER    = 2814749767106562
	ELEMENT_METADATA_AUTHOR               = "author"
	ELEMENT_METADATA_DESCRIPTION          = "description"
	ELEMENT_METADATA_DOC_URI              = "doc-uri"
	ELEMENT_METADATA_ICON_NAME            = "icon-name"
	ELEMENT_METADATA_KLASS                = "klass"
	ELEMENT_METADATA_LONGNAME             = "long-name"
	ERROR_SYSTEM                          = "system error: %s"
	EVENT_NUM_SHIFT                       = 8
	FLAG_SET_MASK_EXACT                   = 4294967295
	FORMAT_PERCENT_MAX                    = 1000000
	FORMAT_PERCENT_SCALE                  = 10000
	FOURCC_FORMAT                         = "c%c%c%c"
	GROUP_ID_INVALID                      = 0
	LICENSE_UNKNOWN                       = "unknown"
	META_TAG_MEMORY_STR                   = "memory"
	MSECOND                               = 1000000
	NSECOND                               = 1
	PARAM_CONTROLLABLE                    = 512
	PARAM_MUTABLE_PAUSED                  = 2048
	PARAM_MUTABLE_PLAYING                 = 4096
	PARAM_MUTABLE_READY                   = 1024
	PARAM_USER_SHIFT                      = 65536
	PROTECTION_SYSTEM_ID_CAPS_FIELD       = "protection-system"
	PTR_FORMAT                            = "paA"
	QUERY_NUM_SHIFT                       = 8
	SECOND                                = 1000000000
	SEGMENT_FORMAT                        = "paB"
	SEQNUM_INVALID                        = 0
	STIME_FORMAT                          = "c%"
	TAG_ALBUM                             = "album"
	TAG_ALBUM_ARTIST                      = "album-artist"
	TAG_ALBUM_ARTIST_SORTNAME             = "album-artist-sortname"
	TAG_ALBUM_GAIN                        = "replaygain-album-gain"
	TAG_ALBUM_PEAK                        = "replaygain-album-peak"
	TAG_ALBUM_SORTNAME                    = "album-sortname"
	TAG_ALBUM_VOLUME_COUNT                = "album-disc-count"
	TAG_ALBUM_VOLUME_NUMBER               = "album-disc-number"
	TAG_APPLICATION_DATA                  = "application-data"
	TAG_APPLICATION_NAME                  = "application-name"
	TAG_ARTIST                            = "artist"
	TAG_ARTIST_SORTNAME                   = "artist-sortname"
	TAG_ATTACHMENT                        = "attachment"
	TAG_AUDIO_CODEC                       = "audio-codec"
	TAG_BEATS_PER_MINUTE                  = "beats-per-minute"
	TAG_BITRATE                           = "bitrate"
	TAG_CODEC                             = "codec"
	TAG_COMMENT                           = "comment"
	TAG_COMPOSER                          = "composer"
	TAG_COMPOSER_SORTNAME                 = "composer-sortname"
	TAG_CONDUCTOR                         = "conductor"
	TAG_CONTACT                           = "contact"
	TAG_CONTAINER_FORMAT                  = "container-format"
	TAG_COPYRIGHT                         = "copyright"
	TAG_COPYRIGHT_URI                     = "copyright-uri"
	TAG_DATE                              = "date"
	TAG_DATE_TIME                         = "datetime"
	TAG_DESCRIPTION                       = "description"
	TAG_DEVICE_MANUFACTURER               = "device-manufacturer"
	TAG_DEVICE_MODEL                      = "device-model"
	TAG_DURATION                          = "duration"
	TAG_ENCODED_BY                        = "encoded-by"
	TAG_ENCODER                           = "encoder"
	TAG_ENCODER_VERSION                   = "encoder-version"
	TAG_EXTENDED_COMMENT                  = "extended-comment"
	TAG_GENRE                             = "genre"
	TAG_GEO_LOCATION_CAPTURE_DIRECTION    = "geo-location-capture-direction"
	TAG_GEO_LOCATION_CITY                 = "geo-location-city"
	TAG_GEO_LOCATION_COUNTRY              = "geo-location-country"
	TAG_GEO_LOCATION_ELEVATION            = "geo-location-elevation"
	TAG_GEO_LOCATION_HORIZONTAL_ERROR     = "geo-location-horizontal-error"
	TAG_GEO_LOCATION_LATITUDE             = "geo-location-latitude"
	TAG_GEO_LOCATION_LONGITUDE            = "geo-location-longitude"
	TAG_GEO_LOCATION_MOVEMENT_DIRECTION   = "geo-location-movement-direction"
	TAG_GEO_LOCATION_MOVEMENT_SPEED       = "geo-location-movement-speed"
	TAG_GEO_LOCATION_NAME                 = "geo-location-name"
	TAG_GEO_LOCATION_SUBLOCATION          = "geo-location-sublocation"
	TAG_GROUPING                          = "grouping"
	TAG_HOMEPAGE                          = "homepage"
	TAG_IMAGE                             = "image"
	TAG_IMAGE_ORIENTATION                 = "image-orientation"
	TAG_INTERPRETED_BY                    = "interpreted-by"
	TAG_ISRC                              = "isrc"
	TAG_KEYWORDS                          = "keywords"
	TAG_LANGUAGE_CODE                     = "language-code"
	TAG_LANGUAGE_NAME                     = "language-name"
	TAG_LICENSE                           = "license"
	TAG_LICENSE_URI                       = "license-uri"
	TAG_LOCATION                          = "location"
	TAG_LYRICS                            = "lyrics"
	TAG_MAXIMUM_BITRATE                   = "maximum-bitrate"
	TAG_MIDI_BASE_NOTE                    = "midi-base-note"
	TAG_MINIMUM_BITRATE                   = "minimum-bitrate"
	TAG_NOMINAL_BITRATE                   = "nominal-bitrate"
	TAG_ORGANIZATION                      = "organization"
	TAG_PERFORMER                         = "performer"
	TAG_PREVIEW_IMAGE                     = "preview-image"
	TAG_PRIVATE_DATA                      = "private-data"
	TAG_PUBLISHER                         = "publisher"
	TAG_REFERENCE_LEVEL                   = "replaygain-reference-level"
	TAG_SERIAL                            = "serial"
	TAG_SHOW_EPISODE_NUMBER               = "show-episode-number"
	TAG_SHOW_NAME                         = "show-name"
	TAG_SHOW_SEASON_NUMBER                = "show-season-number"
	TAG_SHOW_SORTNAME                     = "show-sortname"
	TAG_SUBTITLE_CODEC                    = "subtitle-codec"
	TAG_TITLE                             = "title"
	TAG_TITLE_SORTNAME                    = "title-sortname"
	TAG_TRACK_COUNT                       = "track-count"
	TAG_TRACK_GAIN                        = "replaygain-track-gain"
	TAG_TRACK_NUMBER                      = "track-number"
	TAG_TRACK_PEAK                        = "replaygain-track-peak"
	TAG_USER_RATING                       = "user-rating"
	TAG_VERSION                           = "version"
	TAG_VIDEO_CODEC                       = "video-codec"
	TIME_FORMAT                           = "u:%02u:%02u.%09u"
	TOC_REPEAT_COUNT_INFINITE             = -1
	URI_NO_PORT                           = 0
	USECOND                               = 1000
	VALUE_EQUAL                           = 0
	VALUE_GREATER_THAN                    = 1
	VALUE_LESS_THAN                       = -1
	VALUE_UNORDERED                       = 2
	VERSION_MAJOR                         = 1
	VERSION_MICRO                         = 4
	VERSION_MINOR                         = 14
	VERSION_NANO                          = 0
)
