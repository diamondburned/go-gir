package udisks

/*
#cgo pkg-config: udisks2
#include <udisks/udisks.h>
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("UDisks")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("UDisks", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Interface Block
type Block struct {
	BlockIfc
	P unsafe.Pointer
}
type BlockIfc struct{}
type IBlock interface{ P_Block() unsafe.Pointer }

func (v Block) P_Block() unsafe.Pointer { return v.P }
func BlockGetType() gi.GType {
	ret := _I.GetGType(0, "Block")
	return ret
}

// udisks_block_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func BlockOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(1, "Block", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_block_call_add_configuration_item
//
// [ arg_item ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallAddConfigurationItem(arg_item g.Variant, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(2, "Block", "call_add_configuration_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_item := gi.NewPointerArgument(arg_item.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_item, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_add_configuration_item_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallAddConfigurationItemFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(3, "Block", "call_add_configuration_item_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_add_configuration_item_sync
//
// [ arg_item ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallAddConfigurationItemSync(arg_item g.Variant, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(4, "Block", "call_add_configuration_item_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_item := gi.NewPointerArgument(arg_item.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_item, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_format
//
// [ arg_type ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallFormat(arg_type string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(5, "Block", "call_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_block_call_format_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallFormatFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(6, "Block", "call_format_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_format_sync
//
// [ arg_type ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallFormatSync(arg_type string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(7, "Block", "call_format_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_get_secret_configuration
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallGetSecretConfiguration(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(8, "Block", "call_get_secret_configuration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_get_secret_configuration_finish
//
// [ out_configuration ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallGetSecretConfigurationFinish(res g.IAsyncResult) (result bool, out_configuration g.Variant, err error) {
	iv, err := _I.Get(9, "Block", "call_get_secret_configuration_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_configuration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_configuration, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_configuration.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_get_secret_configuration_sync
//
// [ arg_options ] trans: nothing
//
// [ out_configuration ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallGetSecretConfigurationSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_configuration g.Variant, err error) {
	iv, err := _I.Get(10, "Block", "call_get_secret_configuration_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_configuration := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_configuration, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_configuration.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_device
//
// [ arg_mode ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallOpenDevice(arg_mode string, arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(11, "Block", "call_open_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_mode := gi.CString(arg_mode)
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_mode := gi.NewStringArgument(c_arg_mode)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_mode, arg_arg_options, arg_fd_list, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_mode)
}

// udisks_block_call_open_device_finish
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenDeviceFinish(res g.IAsyncResult) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(12, "Block", "call_open_device_finish")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_out_fd, arg_out_fd_list, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_device_sync
//
// [ arg_mode ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenDeviceSync(arg_mode string, arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(13, "Block", "call_open_device_sync")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_arg_mode := gi.CString(arg_mode)
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_mode := gi.NewStringArgument(c_arg_mode)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_arg_mode, arg_arg_options, arg_fd_list, arg_out_fd, arg_out_fd_list, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_mode)
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_for_backup
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallOpenForBackup(arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(14, "Block", "call_open_for_backup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_fd_list, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_open_for_backup_finish
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenForBackupFinish(res g.IAsyncResult) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(15, "Block", "call_open_for_backup_finish")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_out_fd, arg_out_fd_list, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_for_backup_sync
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenForBackupSync(arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(16, "Block", "call_open_for_backup_sync")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_fd_list, arg_out_fd, arg_out_fd_list, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_for_benchmark
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallOpenForBenchmark(arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(17, "Block", "call_open_for_benchmark")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_fd_list, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_open_for_benchmark_finish
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenForBenchmarkFinish(res g.IAsyncResult) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(18, "Block", "call_open_for_benchmark_finish")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_out_fd, arg_out_fd_list, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_for_benchmark_sync
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenForBenchmarkSync(arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(19, "Block", "call_open_for_benchmark_sync")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_fd_list, arg_out_fd, arg_out_fd_list, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_for_restore
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallOpenForRestore(arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(20, "Block", "call_open_for_restore")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_fd_list, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_open_for_restore_finish
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenForRestoreFinish(res g.IAsyncResult) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(21, "Block", "call_open_for_restore_finish")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_out_fd, arg_out_fd_list, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_open_for_restore_sync
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ out_fd ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallOpenForRestoreSync(arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable) (result bool, out_fd g.Variant, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(22, "Block", "call_open_for_restore_sync")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_out_fd := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_fd_list, arg_out_fd, arg_out_fd_list, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_fd.P = outArgs[0].Pointer()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_block_call_remove_configuration_item
//
// [ arg_item ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallRemoveConfigurationItem(arg_item g.Variant, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(23, "Block", "call_remove_configuration_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_item := gi.NewPointerArgument(arg_item.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_item, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_remove_configuration_item_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallRemoveConfigurationItemFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(24, "Block", "call_remove_configuration_item_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_remove_configuration_item_sync
//
// [ arg_item ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallRemoveConfigurationItemSync(arg_item g.Variant, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(25, "Block", "call_remove_configuration_item_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_item := gi.NewPointerArgument(arg_item.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_item, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_rescan
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallRescan(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(26, "Block", "call_rescan")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_rescan_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallRescanFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(27, "Block", "call_rescan_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_rescan_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallRescanSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(28, "Block", "call_rescan_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_update_configuration_item
//
// [ arg_old_item ] trans: nothing
//
// [ arg_new_item ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *BlockIfc) CallUpdateConfigurationItem(arg_old_item g.Variant, arg_new_item g.Variant, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(29, "Block", "call_update_configuration_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_old_item := gi.NewPointerArgument(arg_old_item.P)
	arg_arg_new_item := gi.NewPointerArgument(arg_new_item.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_old_item, arg_arg_new_item, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_block_call_update_configuration_item_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallUpdateConfigurationItemFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(30, "Block", "call_update_configuration_item_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_call_update_configuration_item_sync
//
// [ arg_old_item ] trans: nothing
//
// [ arg_new_item ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *BlockIfc) CallUpdateConfigurationItemSync(arg_old_item g.Variant, arg_new_item g.Variant, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(31, "Block", "call_update_configuration_item_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_old_item := gi.NewPointerArgument(arg_old_item.P)
	arg_arg_new_item := gi.NewPointerArgument(arg_new_item.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_old_item, arg_arg_new_item, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_block_complete_add_configuration_item
//
// [ invocation ] trans: everything
//
func (v *BlockIfc) CompleteAddConfigurationItem(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(32, "Block", "complete_add_configuration_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_format
//
// [ invocation ] trans: everything
//
func (v *BlockIfc) CompleteFormat(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(33, "Block", "complete_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_get_secret_configuration
//
// [ invocation ] trans: everything
//
// [ configuration ] trans: nothing
//
func (v *BlockIfc) CompleteGetSecretConfiguration(invocation g.IDBusMethodInvocation, configuration g.Variant) {
	iv, err := _I.Get(34, "Block", "complete_get_secret_configuration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_configuration := gi.NewPointerArgument(configuration.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_configuration}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_open_device
//
// [ invocation ] trans: everything
//
// [ fd_list ] trans: nothing
//
// [ fd ] trans: nothing
//
func (v *BlockIfc) CompleteOpenDevice(invocation g.IDBusMethodInvocation, fd_list g.IUnixFDList, fd g.Variant) {
	iv, err := _I.Get(35, "Block", "complete_open_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	var tmp1 unsafe.Pointer
	if fd_list != nil {
		tmp1 = fd_list.P_UnixFDList()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_fd_list := gi.NewPointerArgument(tmp1)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_fd_list, arg_fd}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_open_for_backup
//
// [ invocation ] trans: everything
//
// [ fd_list ] trans: nothing
//
// [ fd ] trans: nothing
//
func (v *BlockIfc) CompleteOpenForBackup(invocation g.IDBusMethodInvocation, fd_list g.IUnixFDList, fd g.Variant) {
	iv, err := _I.Get(36, "Block", "complete_open_for_backup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	var tmp1 unsafe.Pointer
	if fd_list != nil {
		tmp1 = fd_list.P_UnixFDList()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_fd_list := gi.NewPointerArgument(tmp1)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_fd_list, arg_fd}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_open_for_benchmark
//
// [ invocation ] trans: everything
//
// [ fd_list ] trans: nothing
//
// [ fd ] trans: nothing
//
func (v *BlockIfc) CompleteOpenForBenchmark(invocation g.IDBusMethodInvocation, fd_list g.IUnixFDList, fd g.Variant) {
	iv, err := _I.Get(37, "Block", "complete_open_for_benchmark")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	var tmp1 unsafe.Pointer
	if fd_list != nil {
		tmp1 = fd_list.P_UnixFDList()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_fd_list := gi.NewPointerArgument(tmp1)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_fd_list, arg_fd}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_open_for_restore
//
// [ invocation ] trans: everything
//
// [ fd_list ] trans: nothing
//
// [ fd ] trans: nothing
//
func (v *BlockIfc) CompleteOpenForRestore(invocation g.IDBusMethodInvocation, fd_list g.IUnixFDList, fd g.Variant) {
	iv, err := _I.Get(38, "Block", "complete_open_for_restore")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	var tmp1 unsafe.Pointer
	if fd_list != nil {
		tmp1 = fd_list.P_UnixFDList()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_fd_list := gi.NewPointerArgument(tmp1)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_fd_list, arg_fd}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_remove_configuration_item
//
// [ invocation ] trans: everything
//
func (v *BlockIfc) CompleteRemoveConfigurationItem(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(39, "Block", "complete_remove_configuration_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_rescan
//
// [ invocation ] trans: everything
//
func (v *BlockIfc) CompleteRescan(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(40, "Block", "complete_rescan")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_block_complete_update_configuration_item
//
// [ invocation ] trans: everything
//
func (v *BlockIfc) CompleteUpdateConfigurationItem(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(41, "Block", "complete_update_configuration_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct BlockIface

// Object BlockProxy
type BlockProxy struct {
	BlockIfc
	g.DBusProxy
}

func WrapBlockProxy(p unsafe.Pointer) (r BlockProxy) { r.P = p; return }

type IBlockProxy interface{ P_BlockProxy() unsafe.Pointer }

func (v BlockProxy) P_BlockProxy() unsafe.Pointer { return v.P }
func (v BlockProxy) P_Block() unsafe.Pointer      { return v.P }
func BlockProxyGetType() gi.GType {
	ret := _I.GetGType(1, "BlockProxy")
	return ret
}

// udisks_block_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewBlockProxyFinish(res g.IAsyncResult) (result BlockProxy, err error) {
	iv, err := _I.Get(42, "BlockProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_block_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewBlockProxyForBusFinish(res g.IAsyncResult) (result BlockProxy, err error) {
	iv, err := _I.Get(43, "BlockProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_block_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewBlockProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result BlockProxy, err error) {
	iv, err := _I.Get(44, "BlockProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_block_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewBlockProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result BlockProxy, err error) {
	iv, err := _I.Get(45, "BlockProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_block_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func BlockProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(46, "BlockProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_block_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func BlockProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(47, "BlockProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct BlockProxyClass

// Struct BlockProxyPrivate
type BlockProxyPrivate struct {
	P unsafe.Pointer
}

func BlockProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(2, "BlockProxyPrivate")
	return ret
}

// Object BlockSkeleton
type BlockSkeleton struct {
	BlockIfc
	g.DBusInterfaceSkeleton
}

func WrapBlockSkeleton(p unsafe.Pointer) (r BlockSkeleton) { r.P = p; return }

type IBlockSkeleton interface{ P_BlockSkeleton() unsafe.Pointer }

func (v BlockSkeleton) P_BlockSkeleton() unsafe.Pointer { return v.P }
func (v BlockSkeleton) P_Block() unsafe.Pointer         { return v.P }
func BlockSkeletonGetType() gi.GType {
	ret := _I.GetGType(3, "BlockSkeleton")
	return ret
}

// udisks_block_skeleton_new
//
// [ result ] trans: everything
//
func NewBlockSkeleton() (result BlockSkeleton) {
	iv, err := _I.Get(48, "BlockSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct BlockSkeletonClass

// Struct BlockSkeletonPrivate
type BlockSkeletonPrivate struct {
	P unsafe.Pointer
}

func BlockSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(4, "BlockSkeletonPrivate")
	return ret
}

// Object Client
type Client struct {
	g.AsyncInitableIfc
	g.InitableIfc
	g.Object
}

func WrapClient(p unsafe.Pointer) (r Client) { r.P = p; return }

type IClient interface{ P_Client() unsafe.Pointer }

func (v Client) P_Client() unsafe.Pointer        { return v.P }
func (v Client) P_AsyncInitable() unsafe.Pointer { return v.P }
func (v Client) P_Initable() unsafe.Pointer      { return v.P }
func ClientGetType() gi.GType {
	ret := _I.GetGType(5, "Client")
	return ret
}

// udisks_client_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewClientFinish(res g.IAsyncResult) (result Client, err error) {
	iv, err := _I.Get(49, "Client", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_client_new_sync
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewClientSync(cancellable g.ICancellable) (result Client, err error) {
	iv, err := _I.Get(50, "Client", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_client_get_job_description_from_operation
//
// [ operation ] trans: nothing
//
// [ result ] trans: everything
//
func ClientGetJobDescriptionFromOperation1(operation string) (result string) {
	iv, err := _I.Get(51, "Client", "get_job_description_from_operation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_operation := gi.CString(operation)
	arg_operation := gi.NewStringArgument(c_operation)
	args := []gi.Argument{arg_operation}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_operation)
	result = ret.String().Take()
	return
}

// udisks_client_new
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func ClientNew1(cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(52, "Client", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_client_get_all_blocks_for_mdraid
//
// [ raid ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetAllBlocksForMdraid(raid IMDRaid) (result g.List) {
	iv, err := _I.Get(53, "Client", "get_all_blocks_for_mdraid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if raid != nil {
		tmp = raid.P_MDRaid()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_raid := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_raid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_block_for_dev
//
// [ block_device_number ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetBlockForDev(block_device_number uint64) (result Block) {
	iv, err := _I.Get(54, "Client", "get_block_for_dev")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_block_device_number := gi.NewUint64Argument(block_device_number)
	args := []gi.Argument{arg_v, arg_block_device_number}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_block_for_drive
//
// [ drive ] trans: nothing
//
// [ get_physical ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetBlockForDrive(drive IDrive, get_physical bool) (result Block) {
	iv, err := _I.Get(55, "Client", "get_block_for_drive")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if drive != nil {
		tmp = drive.P_Drive()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_drive := gi.NewPointerArgument(tmp)
	arg_get_physical := gi.NewBoolArgument(get_physical)
	args := []gi.Argument{arg_v, arg_drive, arg_get_physical}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_block_for_label
//
// [ label ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetBlockForLabel(label string) (result g.List) {
	iv, err := _I.Get(56, "Client", "get_block_for_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_label := gi.CString(label)
	arg_v := gi.NewPointerArgument(v.P)
	arg_label := gi.NewStringArgument(c_label)
	args := []gi.Argument{arg_v, arg_label}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_label)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_block_for_mdraid
//
// [ raid ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetBlockForMdraid(raid IMDRaid) (result Block) {
	iv, err := _I.Get(57, "Client", "get_block_for_mdraid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if raid != nil {
		tmp = raid.P_MDRaid()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_raid := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_raid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_block_for_uuid
//
// [ uuid ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetBlockForUuid(uuid string) (result g.List) {
	iv, err := _I.Get(58, "Client", "get_block_for_uuid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uuid := gi.CString(uuid)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uuid := gi.NewStringArgument(c_uuid)
	args := []gi.Argument{arg_v, arg_uuid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uuid)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_cleartext_block
//
// [ block ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetCleartextBlock(block IBlock) (result Block) {
	iv, err := _I.Get(59, "Client", "get_cleartext_block")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if block != nil {
		tmp = block.P_Block()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_block := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_drive_for_block
//
// [ block ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetDriveForBlock(block IBlock) (result Drive) {
	iv, err := _I.Get(60, "Client", "get_drive_for_block")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if block != nil {
		tmp = block.P_Block()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_block := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// udisks_client_get_drive_info
//
// [ drive ] trans: nothing
//
// [ out_name ] trans: everything, dir: out
//
// [ out_description ] trans: everything, dir: out
//
// [ out_drive_icon ] trans: everything, dir: out
//
// [ out_media_description ] trans: everything, dir: out
//
// [ out_media_icon ] trans: everything, dir: out
//
func (v Client) GetDriveInfo(drive IDrive) (out_name string, out_description string, out_drive_icon g.Icon, out_media_description string, out_media_icon g.Icon) {
	iv, err := _I.Get(61, "Client", "get_drive_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [5]gi.Argument
	var tmp unsafe.Pointer
	if drive != nil {
		tmp = drive.P_Drive()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_drive := gi.NewPointerArgument(tmp)
	arg_out_name := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_description := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_out_drive_icon := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_out_media_description := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_out_media_icon := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_v, arg_drive, arg_out_name, arg_out_description, arg_out_drive_icon, arg_out_media_description, arg_out_media_icon}
	iv.Call(args, nil, &outArgs[0])
	out_name = outArgs[0].String().Take()
	out_description = outArgs[1].String().Take()
	out_drive_icon.P = outArgs[2].Pointer()
	out_media_description = outArgs[3].String().Take()
	out_media_icon.P = outArgs[4].Pointer()
	return
}

// udisks_client_get_drive_siblings
//
// [ drive ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetDriveSiblings(drive IDrive) (result g.List) {
	iv, err := _I.Get(62, "Client", "get_drive_siblings")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if drive != nil {
		tmp = drive.P_Drive()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_drive := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_drive}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_id_for_display
//
// [ usage ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ version ] trans: nothing
//
// [ long_string ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetIdForDisplay(usage string, type1 string, version string, long_string bool) (result string) {
	iv, err := _I.Get(63, "Client", "get_id_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_usage := gi.CString(usage)
	c_type1 := gi.CString(type1)
	c_version := gi.CString(version)
	arg_v := gi.NewPointerArgument(v.P)
	arg_usage := gi.NewStringArgument(c_usage)
	arg_type1 := gi.NewStringArgument(c_type1)
	arg_version := gi.NewStringArgument(c_version)
	arg_long_string := gi.NewBoolArgument(long_string)
	args := []gi.Argument{arg_v, arg_usage, arg_type1, arg_version, arg_long_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_usage)
	gi.Free(c_type1)
	gi.Free(c_version)
	result = ret.String().Take()
	return
}

// udisks_client_get_job_description
//
// [ job ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetJobDescription(job IJob) (result string) {
	iv, err := _I.Get(64, "Client", "get_job_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if job != nil {
		tmp = job.P_Job()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_job := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_job}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// udisks_client_get_jobs_for_object
//
// [ object ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetJobsForObject(object IObject) (result g.List) {
	iv, err := _I.Get(65, "Client", "get_jobs_for_object")
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
	args := []gi.Argument{arg_v, arg_object}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_loop_for_block
//
// [ block ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetLoopForBlock(block IBlock) (result Loop) {
	iv, err := _I.Get(66, "Client", "get_loop_for_block")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if block != nil {
		tmp = block.P_Block()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_block := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_manager
//
// [ result ] trans: nothing
//
func (v Client) GetManager() (result Manager) {
	iv, err := _I.Get(67, "Client", "get_manager")
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

// udisks_client_get_mdraid_for_block
//
// [ block ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetMdraidForBlock(block IBlock) (result MDRaid) {
	iv, err := _I.Get(68, "Client", "get_mdraid_for_block")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if block != nil {
		tmp = block.P_Block()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_block := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_media_compat_for_display
//
// [ media_compat ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetMediaCompatForDisplay(media_compat string) (result string) {
	iv, err := _I.Get(69, "Client", "get_media_compat_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_media_compat := gi.CString(media_compat)
	arg_v := gi.NewPointerArgument(v.P)
	arg_media_compat := gi.NewStringArgument(c_media_compat)
	args := []gi.Argument{arg_v, arg_media_compat}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_media_compat)
	result = ret.String().Take()
	return
}

// udisks_client_get_members_for_mdraid
//
// [ raid ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetMembersForMdraid(raid IMDRaid) (result g.List) {
	iv, err := _I.Get(70, "Client", "get_members_for_mdraid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if raid != nil {
		tmp = raid.P_MDRaid()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_raid := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_raid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_object
//
// [ object_path ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetObject(object_path string) (result Object) {
	iv, err := _I.Get(71, "Client", "get_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_object_path := gi.CString(object_path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_object_path := gi.NewStringArgument(c_object_path)
	args := []gi.Argument{arg_v, arg_object_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_object_path)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_object_info
//
// [ object ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetObjectInfo(object IObject) (result ObjectInfo) {
	iv, err := _I.Get(72, "Client", "get_object_info")
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
	args := []gi.Argument{arg_v, arg_object}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_object_manager
//
// [ result ] trans: nothing
//
func (v Client) GetObjectManager() (result g.DBusObjectManager) {
	iv, err := _I.Get(73, "Client", "get_object_manager")
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

// udisks_client_get_partition_info
//
// [ partition ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetPartitionInfo(partition IPartition) (result string) {
	iv, err := _I.Get(74, "Client", "get_partition_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if partition != nil {
		tmp = partition.P_Partition()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_partition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// udisks_client_get_partition_table
//
// [ partition ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetPartitionTable(partition IPartition) (result PartitionTable) {
	iv, err := _I.Get(75, "Client", "get_partition_table")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if partition != nil {
		tmp = partition.P_Partition()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_partition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_partition_table_subtype_for_display
//
// [ partition_table_type ] trans: nothing
//
// [ partition_table_subtype ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Client) GetPartitionTableSubtypeForDisplay(partition_table_type string, partition_table_subtype string) (result string) {
	iv, err := _I.Get(76, "Client", "get_partition_table_subtype_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_partition_table_type := gi.CString(partition_table_type)
	c_partition_table_subtype := gi.CString(partition_table_subtype)
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition_table_type := gi.NewStringArgument(c_partition_table_type)
	arg_partition_table_subtype := gi.NewStringArgument(c_partition_table_subtype)
	args := []gi.Argument{arg_v, arg_partition_table_type, arg_partition_table_subtype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_partition_table_type)
	gi.Free(c_partition_table_subtype)
	result = ret.String().Copy()
	return
}

// udisks_client_get_partition_table_subtypes
//
// [ partition_table_type ] trans: nothing
//
// [ result ] trans: container
//
func (v Client) GetPartitionTableSubtypes(partition_table_type string) (result gi.CStrArray) {
	iv, err := _I.Get(77, "Client", "get_partition_table_subtypes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_partition_table_type := gi.CString(partition_table_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition_table_type := gi.NewStringArgument(c_partition_table_type)
	args := []gi.Argument{arg_v, arg_partition_table_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_partition_table_type)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// udisks_client_get_partition_table_type_for_display
//
// [ partition_table_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Client) GetPartitionTableTypeForDisplay(partition_table_type string) (result string) {
	iv, err := _I.Get(78, "Client", "get_partition_table_type_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_partition_table_type := gi.CString(partition_table_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition_table_type := gi.NewStringArgument(c_partition_table_type)
	args := []gi.Argument{arg_v, arg_partition_table_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_partition_table_type)
	result = ret.String().Copy()
	return
}

// udisks_client_get_partition_type_and_subtype_for_display
//
// [ partition_table_type ] trans: nothing
//
// [ partition_table_subtype ] trans: nothing
//
// [ partition_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Client) GetPartitionTypeAndSubtypeForDisplay(partition_table_type string, partition_table_subtype string, partition_type string) (result string) {
	iv, err := _I.Get(79, "Client", "get_partition_type_and_subtype_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_partition_table_type := gi.CString(partition_table_type)
	c_partition_table_subtype := gi.CString(partition_table_subtype)
	c_partition_type := gi.CString(partition_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition_table_type := gi.NewStringArgument(c_partition_table_type)
	arg_partition_table_subtype := gi.NewStringArgument(c_partition_table_subtype)
	arg_partition_type := gi.NewStringArgument(c_partition_type)
	args := []gi.Argument{arg_v, arg_partition_table_type, arg_partition_table_subtype, arg_partition_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_partition_table_type)
	gi.Free(c_partition_table_subtype)
	gi.Free(c_partition_type)
	result = ret.String().Copy()
	return
}

// udisks_client_get_partition_type_for_display
//
// [ partition_table_type ] trans: nothing
//
// [ partition_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Client) GetPartitionTypeForDisplay(partition_table_type string, partition_type string) (result string) {
	iv, err := _I.Get(80, "Client", "get_partition_type_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_partition_table_type := gi.CString(partition_table_type)
	c_partition_type := gi.CString(partition_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition_table_type := gi.NewStringArgument(c_partition_table_type)
	arg_partition_type := gi.NewStringArgument(c_partition_type)
	args := []gi.Argument{arg_v, arg_partition_table_type, arg_partition_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_partition_table_type)
	gi.Free(c_partition_type)
	result = ret.String().Copy()
	return
}

// udisks_client_get_partition_type_infos
//
// [ partition_table_type ] trans: nothing
//
// [ partition_table_subtype ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetPartitionTypeInfos(partition_table_type string, partition_table_subtype string) (result g.List) {
	iv, err := _I.Get(81, "Client", "get_partition_type_infos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_partition_table_type := gi.CString(partition_table_type)
	c_partition_table_subtype := gi.CString(partition_table_subtype)
	arg_v := gi.NewPointerArgument(v.P)
	arg_partition_table_type := gi.NewStringArgument(c_partition_table_type)
	arg_partition_table_subtype := gi.NewStringArgument(c_partition_table_subtype)
	args := []gi.Argument{arg_v, arg_partition_table_type, arg_partition_table_subtype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_partition_table_type)
	gi.Free(c_partition_table_subtype)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_partitions
//
// [ table ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetPartitions(table IPartitionTable) (result g.List) {
	iv, err := _I.Get(82, "Client", "get_partitions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if table != nil {
		tmp = table.P_PartitionTable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_table := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_client_get_size_for_display
//
// [ size ] trans: nothing
//
// [ use_pow2 ] trans: nothing
//
// [ long_string ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) GetSizeForDisplay(size uint64, use_pow2 bool, long_string bool) (result string) {
	iv, err := _I.Get(83, "Client", "get_size_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewUint64Argument(size)
	arg_use_pow2 := gi.NewBoolArgument(use_pow2)
	arg_long_string := gi.NewBoolArgument(long_string)
	args := []gi.Argument{arg_v, arg_size, arg_use_pow2, arg_long_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// udisks_client_peek_object
//
// [ object_path ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Client) PeekObject(object_path string) (result Object) {
	iv, err := _I.Get(84, "Client", "peek_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_object_path := gi.CString(object_path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_object_path := gi.NewStringArgument(c_object_path)
	args := []gi.Argument{arg_v, arg_object_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_object_path)
	result.P = ret.Pointer()
	return
}

// udisks_client_queue_changed
//
func (v Client) QueueChanged() {
	iv, err := _I.Get(85, "Client", "queue_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// udisks_client_settle
//
func (v Client) Settle() {
	iv, err := _I.Get(86, "Client", "settle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Interface Drive
type Drive struct {
	DriveIfc
	P unsafe.Pointer
}
type DriveIfc struct{}
type IDrive interface{ P_Drive() unsafe.Pointer }

func (v Drive) P_Drive() unsafe.Pointer { return v.P }
func DriveGetType() gi.GType {
	ret := _I.GetGType(6, "Drive")
	return ret
}

// udisks_drive_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func DriveOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(88, "Drive", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_drive_call_eject
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveIfc) CallEject(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(89, "Drive", "call_eject")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_call_eject_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveIfc) CallEjectFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(90, "Drive", "call_eject_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_call_eject_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveIfc) CallEjectSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(91, "Drive", "call_eject_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_call_power_off
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveIfc) CallPowerOff(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(92, "Drive", "call_power_off")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_call_power_off_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveIfc) CallPowerOffFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(93, "Drive", "call_power_off_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_call_power_off_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveIfc) CallPowerOffSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(94, "Drive", "call_power_off_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_call_set_configuration
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveIfc) CallSetConfiguration(arg_value g.Variant, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(95, "Drive", "call_set_configuration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewPointerArgument(arg_value.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_call_set_configuration_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveIfc) CallSetConfigurationFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(96, "Drive", "call_set_configuration_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_call_set_configuration_sync
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveIfc) CallSetConfigurationSync(arg_value g.Variant, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(97, "Drive", "call_set_configuration_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewPointerArgument(arg_value.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_complete_eject
//
// [ invocation ] trans: everything
//
func (v *DriveIfc) CompleteEject(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(98, "Drive", "complete_eject")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_complete_power_off
//
// [ invocation ] trans: everything
//
func (v *DriveIfc) CompletePowerOff(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(99, "Drive", "complete_power_off")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_complete_set_configuration
//
// [ invocation ] trans: everything
//
func (v *DriveIfc) CompleteSetConfiguration(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(100, "Drive", "complete_set_configuration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// Interface DriveAta
type DriveAta struct {
	DriveAtaIfc
	P unsafe.Pointer
}
type DriveAtaIfc struct{}
type IDriveAta interface{ P_DriveAta() unsafe.Pointer }

func (v DriveAta) P_DriveAta() unsafe.Pointer { return v.P }
func DriveAtaGetType() gi.GType {
	ret := _I.GetGType(7, "DriveAta")
	return ret
}

// udisks_drive_ata_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func DriveAtaOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(102, "DriveAta", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_drive_ata_call_pm_get_state
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallPmGetState(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(103, "DriveAta", "call_pm_get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_pm_get_state_finish
//
// [ out_state ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallPmGetStateFinish(res g.IAsyncResult) (result bool, out_state uint8, err error) {
	iv, err := _I.Get(104, "DriveAta", "call_pm_get_state_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_state := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_state, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_state = outArgs[0].Uint8()
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_pm_get_state_sync
//
// [ arg_options ] trans: nothing
//
// [ out_state ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallPmGetStateSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_state uint8, err error) {
	iv, err := _I.Get(105, "DriveAta", "call_pm_get_state_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_state := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_state, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_state = outArgs[0].Uint8()
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_pm_standby
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallPmStandby(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(106, "DriveAta", "call_pm_standby")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_pm_standby_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallPmStandbyFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(107, "DriveAta", "call_pm_standby_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_pm_standby_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallPmStandbySync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(108, "DriveAta", "call_pm_standby_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_pm_wakeup
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallPmWakeup(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(109, "DriveAta", "call_pm_wakeup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_pm_wakeup_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallPmWakeupFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(110, "DriveAta", "call_pm_wakeup_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_pm_wakeup_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallPmWakeupSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(111, "DriveAta", "call_pm_wakeup_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_security_erase_unit
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallSecurityEraseUnit(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(112, "DriveAta", "call_security_erase_unit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_security_erase_unit_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSecurityEraseUnitFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(113, "DriveAta", "call_security_erase_unit_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_security_erase_unit_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSecurityEraseUnitSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(114, "DriveAta", "call_security_erase_unit_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_get_attributes
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartGetAttributes(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(115, "DriveAta", "call_smart_get_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_smart_get_attributes_finish
//
// [ out_attributes ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartGetAttributesFinish(res g.IAsyncResult) (result bool, out_attributes g.Variant, err error) {
	iv, err := _I.Get(116, "DriveAta", "call_smart_get_attributes_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_attributes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_attributes, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_attributes.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_get_attributes_sync
//
// [ arg_options ] trans: nothing
//
// [ out_attributes ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartGetAttributesSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_attributes g.Variant, err error) {
	iv, err := _I.Get(117, "DriveAta", "call_smart_get_attributes_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_attributes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_attributes, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_attributes.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_selftest_abort
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSelftestAbort(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(118, "DriveAta", "call_smart_selftest_abort")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_smart_selftest_abort_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSelftestAbortFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(119, "DriveAta", "call_smart_selftest_abort_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_selftest_abort_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSelftestAbortSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(120, "DriveAta", "call_smart_selftest_abort_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_selftest_start
//
// [ arg_type ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSelftestStart(arg_type string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(121, "DriveAta", "call_smart_selftest_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_drive_ata_call_smart_selftest_start_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSelftestStartFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(122, "DriveAta", "call_smart_selftest_start_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_selftest_start_sync
//
// [ arg_type ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSelftestStartSync(arg_type string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(123, "DriveAta", "call_smart_selftest_start_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_set_enabled
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSetEnabled(arg_value bool, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(124, "DriveAta", "call_smart_set_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewBoolArgument(arg_value)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_smart_set_enabled_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSetEnabledFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(125, "DriveAta", "call_smart_set_enabled_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_set_enabled_sync
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartSetEnabledSync(arg_value bool, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(126, "DriveAta", "call_smart_set_enabled_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewBoolArgument(arg_value)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_update
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartUpdate(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(127, "DriveAta", "call_smart_update")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_call_smart_update_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartUpdateFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(128, "DriveAta", "call_smart_update_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_call_smart_update_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DriveAtaIfc) CallSmartUpdateSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(129, "DriveAta", "call_smart_update_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_drive_ata_complete_pm_get_state
//
// [ invocation ] trans: everything
//
// [ state ] trans: nothing
//
func (v *DriveAtaIfc) CompletePmGetState(invocation g.IDBusMethodInvocation, state uint8) {
	iv, err := _I.Get(130, "DriveAta", "complete_pm_get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_state := gi.NewUint8Argument(state)
	args := []gi.Argument{arg_v, arg_invocation, arg_state}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_pm_standby
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompletePmStandby(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(131, "DriveAta", "complete_pm_standby")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_pm_wakeup
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompletePmWakeup(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(132, "DriveAta", "complete_pm_wakeup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_security_erase_unit
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompleteSecurityEraseUnit(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(133, "DriveAta", "complete_security_erase_unit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_smart_get_attributes
//
// [ invocation ] trans: everything
//
// [ attributes ] trans: nothing
//
func (v *DriveAtaIfc) CompleteSmartGetAttributes(invocation g.IDBusMethodInvocation, attributes g.Variant) {
	iv, err := _I.Get(134, "DriveAta", "complete_smart_get_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_attributes := gi.NewPointerArgument(attributes.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_attributes}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_smart_selftest_abort
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompleteSmartSelftestAbort(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(135, "DriveAta", "complete_smart_selftest_abort")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_smart_selftest_start
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompleteSmartSelftestStart(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(136, "DriveAta", "complete_smart_selftest_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_smart_set_enabled
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompleteSmartSetEnabled(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(137, "DriveAta", "complete_smart_set_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_drive_ata_complete_smart_update
//
// [ invocation ] trans: everything
//
func (v *DriveAtaIfc) CompleteSmartUpdate(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(138, "DriveAta", "complete_smart_update")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct DriveAtaIface

// Object DriveAtaProxy
type DriveAtaProxy struct {
	DriveAtaIfc
	g.DBusProxy
}

func WrapDriveAtaProxy(p unsafe.Pointer) (r DriveAtaProxy) { r.P = p; return }

type IDriveAtaProxy interface{ P_DriveAtaProxy() unsafe.Pointer }

func (v DriveAtaProxy) P_DriveAtaProxy() unsafe.Pointer { return v.P }
func (v DriveAtaProxy) P_DriveAta() unsafe.Pointer      { return v.P }
func DriveAtaProxyGetType() gi.GType {
	ret := _I.GetGType(8, "DriveAtaProxy")
	return ret
}

// udisks_drive_ata_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveAtaProxyFinish(res g.IAsyncResult) (result DriveAtaProxy, err error) {
	iv, err := _I.Get(139, "DriveAtaProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_ata_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveAtaProxyForBusFinish(res g.IAsyncResult) (result DriveAtaProxy, err error) {
	iv, err := _I.Get(140, "DriveAtaProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_ata_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveAtaProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result DriveAtaProxy, err error) {
	iv, err := _I.Get(141, "DriveAtaProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_ata_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveAtaProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result DriveAtaProxy, err error) {
	iv, err := _I.Get(142, "DriveAtaProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_ata_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DriveAtaProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(143, "DriveAtaProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_drive_ata_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DriveAtaProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(144, "DriveAtaProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct DriveAtaProxyClass

// Struct DriveAtaProxyPrivate
type DriveAtaProxyPrivate struct {
	P unsafe.Pointer
}

func DriveAtaProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(9, "DriveAtaProxyPrivate")
	return ret
}

// Object DriveAtaSkeleton
type DriveAtaSkeleton struct {
	DriveAtaIfc
	g.DBusInterfaceSkeleton
}

func WrapDriveAtaSkeleton(p unsafe.Pointer) (r DriveAtaSkeleton) { r.P = p; return }

type IDriveAtaSkeleton interface{ P_DriveAtaSkeleton() unsafe.Pointer }

func (v DriveAtaSkeleton) P_DriveAtaSkeleton() unsafe.Pointer { return v.P }
func (v DriveAtaSkeleton) P_DriveAta() unsafe.Pointer         { return v.P }
func DriveAtaSkeletonGetType() gi.GType {
	ret := _I.GetGType(10, "DriveAtaSkeleton")
	return ret
}

// udisks_drive_ata_skeleton_new
//
// [ result ] trans: everything
//
func NewDriveAtaSkeleton() (result DriveAtaSkeleton) {
	iv, err := _I.Get(145, "DriveAtaSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct DriveAtaSkeletonClass

// Struct DriveAtaSkeletonPrivate
type DriveAtaSkeletonPrivate struct {
	P unsafe.Pointer
}

func DriveAtaSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(11, "DriveAtaSkeletonPrivate")
	return ret
}

// ignore GType struct DriveIface

// Object DriveProxy
type DriveProxy struct {
	DriveIfc
	g.DBusProxy
}

func WrapDriveProxy(p unsafe.Pointer) (r DriveProxy) { r.P = p; return }

type IDriveProxy interface{ P_DriveProxy() unsafe.Pointer }

func (v DriveProxy) P_DriveProxy() unsafe.Pointer { return v.P }
func (v DriveProxy) P_Drive() unsafe.Pointer      { return v.P }
func DriveProxyGetType() gi.GType {
	ret := _I.GetGType(12, "DriveProxy")
	return ret
}

// udisks_drive_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveProxyFinish(res g.IAsyncResult) (result DriveProxy, err error) {
	iv, err := _I.Get(146, "DriveProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveProxyForBusFinish(res g.IAsyncResult) (result DriveProxy, err error) {
	iv, err := _I.Get(147, "DriveProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result DriveProxy, err error) {
	iv, err := _I.Get(148, "DriveProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewDriveProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result DriveProxy, err error) {
	iv, err := _I.Get(149, "DriveProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_drive_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DriveProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(150, "DriveProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_drive_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func DriveProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(151, "DriveProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct DriveProxyClass

// Struct DriveProxyPrivate
type DriveProxyPrivate struct {
	P unsafe.Pointer
}

func DriveProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(13, "DriveProxyPrivate")
	return ret
}

// Object DriveSkeleton
type DriveSkeleton struct {
	DriveIfc
	g.DBusInterfaceSkeleton
}

func WrapDriveSkeleton(p unsafe.Pointer) (r DriveSkeleton) { r.P = p; return }

type IDriveSkeleton interface{ P_DriveSkeleton() unsafe.Pointer }

func (v DriveSkeleton) P_DriveSkeleton() unsafe.Pointer { return v.P }
func (v DriveSkeleton) P_Drive() unsafe.Pointer         { return v.P }
func DriveSkeletonGetType() gi.GType {
	ret := _I.GetGType(14, "DriveSkeleton")
	return ret
}

// udisks_drive_skeleton_new
//
// [ result ] trans: everything
//
func NewDriveSkeleton() (result DriveSkeleton) {
	iv, err := _I.Get(152, "DriveSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct DriveSkeletonClass

// Struct DriveSkeletonPrivate
type DriveSkeletonPrivate struct {
	P unsafe.Pointer
}

func DriveSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(15, "DriveSkeletonPrivate")
	return ret
}

// Interface Encrypted
type Encrypted struct {
	EncryptedIfc
	P unsafe.Pointer
}
type EncryptedIfc struct{}
type IEncrypted interface{ P_Encrypted() unsafe.Pointer }

func (v Encrypted) P_Encrypted() unsafe.Pointer { return v.P }
func EncryptedGetType() gi.GType {
	ret := _I.GetGType(16, "Encrypted")
	return ret
}

// udisks_encrypted_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func EncryptedOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(154, "Encrypted", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_encrypted_call_change_passphrase
//
// [ arg_passphrase ] trans: nothing
//
// [ arg_new_passphrase ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *EncryptedIfc) CallChangePassphrase(arg_passphrase string, arg_new_passphrase string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(155, "Encrypted", "call_change_passphrase")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_passphrase := gi.CString(arg_passphrase)
	c_arg_new_passphrase := gi.CString(arg_new_passphrase)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_passphrase := gi.NewStringArgument(c_arg_passphrase)
	arg_arg_new_passphrase := gi.NewStringArgument(c_arg_new_passphrase)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_passphrase, arg_arg_new_passphrase, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_passphrase)
	gi.Free(c_arg_new_passphrase)
}

// udisks_encrypted_call_change_passphrase_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallChangePassphraseFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(156, "Encrypted", "call_change_passphrase_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_encrypted_call_change_passphrase_sync
//
// [ arg_passphrase ] trans: nothing
//
// [ arg_new_passphrase ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallChangePassphraseSync(arg_passphrase string, arg_new_passphrase string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(157, "Encrypted", "call_change_passphrase_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_passphrase := gi.CString(arg_passphrase)
	c_arg_new_passphrase := gi.CString(arg_new_passphrase)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_passphrase := gi.NewStringArgument(c_arg_passphrase)
	arg_arg_new_passphrase := gi.NewStringArgument(c_arg_new_passphrase)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_passphrase, arg_arg_new_passphrase, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_passphrase)
	gi.Free(c_arg_new_passphrase)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_encrypted_call_lock
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *EncryptedIfc) CallLock(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(158, "Encrypted", "call_lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_encrypted_call_lock_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallLockFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(159, "Encrypted", "call_lock_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_encrypted_call_lock_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallLockSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(160, "Encrypted", "call_lock_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_encrypted_call_resize
//
// [ arg_size ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *EncryptedIfc) CallResize(arg_size uint64, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(161, "Encrypted", "call_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_size, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_encrypted_call_resize_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallResizeFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(162, "Encrypted", "call_resize_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_encrypted_call_resize_sync
//
// [ arg_size ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallResizeSync(arg_size uint64, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(163, "Encrypted", "call_resize_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_size, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_encrypted_call_unlock
//
// [ arg_passphrase ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *EncryptedIfc) CallUnlock(arg_passphrase string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(164, "Encrypted", "call_unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_passphrase := gi.CString(arg_passphrase)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_passphrase := gi.NewStringArgument(c_arg_passphrase)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_passphrase, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_passphrase)
}

// udisks_encrypted_call_unlock_finish
//
// [ out_cleartext_device ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallUnlockFinish(res g.IAsyncResult) (result bool, out_cleartext_device string, err error) {
	iv, err := _I.Get(165, "Encrypted", "call_unlock_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_cleartext_device := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_cleartext_device, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_cleartext_device = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_encrypted_call_unlock_sync
//
// [ arg_passphrase ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ out_cleartext_device ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EncryptedIfc) CallUnlockSync(arg_passphrase string, arg_options g.Variant, cancellable g.ICancellable) (result bool, out_cleartext_device string, err error) {
	iv, err := _I.Get(166, "Encrypted", "call_unlock_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_passphrase := gi.CString(arg_passphrase)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_passphrase := gi.NewStringArgument(c_arg_passphrase)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_cleartext_device := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_passphrase, arg_arg_options, arg_out_cleartext_device, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_passphrase)
	err = gi.ToError(outArgs[1].Pointer())
	out_cleartext_device = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_encrypted_complete_change_passphrase
//
// [ invocation ] trans: everything
//
func (v *EncryptedIfc) CompleteChangePassphrase(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(167, "Encrypted", "complete_change_passphrase")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_encrypted_complete_lock
//
// [ invocation ] trans: everything
//
func (v *EncryptedIfc) CompleteLock(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(168, "Encrypted", "complete_lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_encrypted_complete_resize
//
// [ invocation ] trans: everything
//
func (v *EncryptedIfc) CompleteResize(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(169, "Encrypted", "complete_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_encrypted_complete_unlock
//
// [ invocation ] trans: everything
//
// [ cleartext_device ] trans: nothing
//
func (v *EncryptedIfc) CompleteUnlock(invocation g.IDBusMethodInvocation, cleartext_device string) {
	iv, err := _I.Get(170, "Encrypted", "complete_unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_cleartext_device := gi.CString(cleartext_device)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_cleartext_device := gi.NewStringArgument(c_cleartext_device)
	args := []gi.Argument{arg_v, arg_invocation, arg_cleartext_device}
	iv.Call(args, nil, nil)
	gi.Free(c_cleartext_device)
}

// ignore GType struct EncryptedIface

// Object EncryptedProxy
type EncryptedProxy struct {
	EncryptedIfc
	g.DBusProxy
}

func WrapEncryptedProxy(p unsafe.Pointer) (r EncryptedProxy) { r.P = p; return }

type IEncryptedProxy interface{ P_EncryptedProxy() unsafe.Pointer }

func (v EncryptedProxy) P_EncryptedProxy() unsafe.Pointer { return v.P }
func (v EncryptedProxy) P_Encrypted() unsafe.Pointer      { return v.P }
func EncryptedProxyGetType() gi.GType {
	ret := _I.GetGType(17, "EncryptedProxy")
	return ret
}

// udisks_encrypted_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewEncryptedProxyFinish(res g.IAsyncResult) (result EncryptedProxy, err error) {
	iv, err := _I.Get(171, "EncryptedProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_encrypted_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewEncryptedProxyForBusFinish(res g.IAsyncResult) (result EncryptedProxy, err error) {
	iv, err := _I.Get(172, "EncryptedProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_encrypted_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewEncryptedProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result EncryptedProxy, err error) {
	iv, err := _I.Get(173, "EncryptedProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_encrypted_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewEncryptedProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result EncryptedProxy, err error) {
	iv, err := _I.Get(174, "EncryptedProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_encrypted_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func EncryptedProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(175, "EncryptedProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_encrypted_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func EncryptedProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(176, "EncryptedProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct EncryptedProxyClass

// Struct EncryptedProxyPrivate
type EncryptedProxyPrivate struct {
	P unsafe.Pointer
}

func EncryptedProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(18, "EncryptedProxyPrivate")
	return ret
}

// Object EncryptedSkeleton
type EncryptedSkeleton struct {
	EncryptedIfc
	g.DBusInterfaceSkeleton
}

func WrapEncryptedSkeleton(p unsafe.Pointer) (r EncryptedSkeleton) { r.P = p; return }

type IEncryptedSkeleton interface{ P_EncryptedSkeleton() unsafe.Pointer }

func (v EncryptedSkeleton) P_EncryptedSkeleton() unsafe.Pointer { return v.P }
func (v EncryptedSkeleton) P_Encrypted() unsafe.Pointer         { return v.P }
func EncryptedSkeletonGetType() gi.GType {
	ret := _I.GetGType(19, "EncryptedSkeleton")
	return ret
}

// udisks_encrypted_skeleton_new
//
// [ result ] trans: everything
//
func NewEncryptedSkeleton() (result EncryptedSkeleton) {
	iv, err := _I.Get(177, "EncryptedSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct EncryptedSkeletonClass

// Struct EncryptedSkeletonPrivate
type EncryptedSkeletonPrivate struct {
	P unsafe.Pointer
}

func EncryptedSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(20, "EncryptedSkeletonPrivate")
	return ret
}

// Enum Error
type ErrorEnum int

const (
	ErrorFailed                     ErrorEnum = 0
	ErrorCancelled                  ErrorEnum = 1
	ErrorAlreadyCancelled           ErrorEnum = 2
	ErrorNotAuthorized              ErrorEnum = 3
	ErrorNotAuthorizedCanObtain     ErrorEnum = 4
	ErrorNotAuthorizedDismissed     ErrorEnum = 5
	ErrorAlreadyMounted             ErrorEnum = 6
	ErrorNotMounted                 ErrorEnum = 7
	ErrorOptionNotPermitted         ErrorEnum = 8
	ErrorMountedByOtherUser         ErrorEnum = 9
	ErrorAlreadyUnmounting          ErrorEnum = 10
	ErrorNotSupported               ErrorEnum = 11
	ErrorTimedOut                   ErrorEnum = 12
	ErrorWouldWakeup                ErrorEnum = 13
	ErrorDeviceBusy                 ErrorEnum = 14
	ErrorIscsiDaemonTransportFailed ErrorEnum = 15
	ErrorIscsiHostNotFound          ErrorEnum = 16
	ErrorIscsiIdmb                  ErrorEnum = 17
	ErrorIscsiLoginFailed           ErrorEnum = 18
	ErrorIscsiLoginAuthFailed       ErrorEnum = 19
	ErrorIscsiLoginFatal            ErrorEnum = 20
	ErrorIscsiLogoutFailed          ErrorEnum = 21
	ErrorIscsiNoFirmware            ErrorEnum = 22
	ErrorIscsiNoObjectsFound        ErrorEnum = 23
	ErrorIscsiNotConnected          ErrorEnum = 24
	ErrorIscsiTransportFailed       ErrorEnum = 25
	ErrorIscsiUnknownDiscoveryType  ErrorEnum = 26
)

func ErrorGetType() gi.GType {
	ret := _I.GetGType(21, "Error")
	return ret
}

// Interface Filesystem
type Filesystem struct {
	FilesystemIfc
	P unsafe.Pointer
}
type FilesystemIfc struct{}
type IFilesystem interface{ P_Filesystem() unsafe.Pointer }

func (v Filesystem) P_Filesystem() unsafe.Pointer { return v.P }
func FilesystemGetType() gi.GType {
	ret := _I.GetGType(22, "Filesystem")
	return ret
}

// udisks_filesystem_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func FilesystemOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(179, "Filesystem", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_filesystem_call_check
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallCheck(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(180, "Filesystem", "call_check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_call_check_finish
//
// [ out_consistent ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallCheckFinish(res g.IAsyncResult) (result bool, out_consistent bool, err error) {
	iv, err := _I.Get(181, "Filesystem", "call_check_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_consistent := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_consistent, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_consistent = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// udisks_filesystem_call_check_sync
//
// [ arg_options ] trans: nothing
//
// [ out_consistent ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallCheckSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_consistent bool, err error) {
	iv, err := _I.Get(182, "Filesystem", "call_check_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_consistent := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_consistent, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_consistent = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// udisks_filesystem_call_mount
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallMount(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(183, "Filesystem", "call_mount")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_call_mount_finish
//
// [ out_mount_path ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallMountFinish(res g.IAsyncResult) (result bool, out_mount_path string, err error) {
	iv, err := _I.Get(184, "Filesystem", "call_mount_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_mount_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_mount_path, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_mount_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_filesystem_call_mount_sync
//
// [ arg_options ] trans: nothing
//
// [ out_mount_path ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallMountSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_mount_path string, err error) {
	iv, err := _I.Get(185, "Filesystem", "call_mount_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_mount_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_mount_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_mount_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_filesystem_call_repair
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallRepair(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(186, "Filesystem", "call_repair")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_call_repair_finish
//
// [ out_repaired ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallRepairFinish(res g.IAsyncResult) (result bool, out_repaired bool, err error) {
	iv, err := _I.Get(187, "Filesystem", "call_repair_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_repaired := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_repaired, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_repaired = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// udisks_filesystem_call_repair_sync
//
// [ arg_options ] trans: nothing
//
// [ out_repaired ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallRepairSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_repaired bool, err error) {
	iv, err := _I.Get(188, "Filesystem", "call_repair_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_repaired := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_repaired, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_repaired = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// udisks_filesystem_call_resize
//
// [ arg_size ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallResize(arg_size uint64, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(189, "Filesystem", "call_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_size, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_call_resize_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallResizeFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(190, "Filesystem", "call_resize_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_resize_sync
//
// [ arg_size ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallResizeSync(arg_size uint64, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(191, "Filesystem", "call_resize_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_size, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_set_label
//
// [ arg_label ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallSetLabel(arg_label string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(192, "Filesystem", "call_set_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_label := gi.CString(arg_label)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_label := gi.NewStringArgument(c_arg_label)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_label, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_label)
}

// udisks_filesystem_call_set_label_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallSetLabelFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(193, "Filesystem", "call_set_label_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_set_label_sync
//
// [ arg_label ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallSetLabelSync(arg_label string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(194, "Filesystem", "call_set_label_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_label := gi.CString(arg_label)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_label := gi.NewStringArgument(c_arg_label)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_label, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_label)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_take_ownership
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallTakeOwnership(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(195, "Filesystem", "call_take_ownership")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_call_take_ownership_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallTakeOwnershipFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(196, "Filesystem", "call_take_ownership_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_take_ownership_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallTakeOwnershipSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(197, "Filesystem", "call_take_ownership_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_unmount
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *FilesystemIfc) CallUnmount(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(198, "Filesystem", "call_unmount")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_call_unmount_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallUnmountFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(199, "Filesystem", "call_unmount_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_call_unmount_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *FilesystemIfc) CallUnmountSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(200, "Filesystem", "call_unmount_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_filesystem_complete_check
//
// [ invocation ] trans: everything
//
// [ consistent ] trans: nothing
//
func (v *FilesystemIfc) CompleteCheck(invocation g.IDBusMethodInvocation, consistent bool) {
	iv, err := _I.Get(201, "Filesystem", "complete_check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_consistent := gi.NewBoolArgument(consistent)
	args := []gi.Argument{arg_v, arg_invocation, arg_consistent}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_complete_mount
//
// [ invocation ] trans: everything
//
// [ mount_path ] trans: nothing
//
func (v *FilesystemIfc) CompleteMount(invocation g.IDBusMethodInvocation, mount_path string) {
	iv, err := _I.Get(202, "Filesystem", "complete_mount")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_mount_path := gi.CString(mount_path)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_mount_path := gi.NewStringArgument(c_mount_path)
	args := []gi.Argument{arg_v, arg_invocation, arg_mount_path}
	iv.Call(args, nil, nil)
	gi.Free(c_mount_path)
}

// udisks_filesystem_complete_repair
//
// [ invocation ] trans: everything
//
// [ repaired ] trans: nothing
//
func (v *FilesystemIfc) CompleteRepair(invocation g.IDBusMethodInvocation, repaired bool) {
	iv, err := _I.Get(203, "Filesystem", "complete_repair")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_repaired := gi.NewBoolArgument(repaired)
	args := []gi.Argument{arg_v, arg_invocation, arg_repaired}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_complete_resize
//
// [ invocation ] trans: everything
//
func (v *FilesystemIfc) CompleteResize(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(204, "Filesystem", "complete_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_complete_set_label
//
// [ invocation ] trans: everything
//
func (v *FilesystemIfc) CompleteSetLabel(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(205, "Filesystem", "complete_set_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_complete_take_ownership
//
// [ invocation ] trans: everything
//
func (v *FilesystemIfc) CompleteTakeOwnership(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(206, "Filesystem", "complete_take_ownership")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_filesystem_complete_unmount
//
// [ invocation ] trans: everything
//
func (v *FilesystemIfc) CompleteUnmount(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(207, "Filesystem", "complete_unmount")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct FilesystemIface

// Object FilesystemProxy
type FilesystemProxy struct {
	FilesystemIfc
	g.DBusProxy
}

func WrapFilesystemProxy(p unsafe.Pointer) (r FilesystemProxy) { r.P = p; return }

type IFilesystemProxy interface{ P_FilesystemProxy() unsafe.Pointer }

func (v FilesystemProxy) P_FilesystemProxy() unsafe.Pointer { return v.P }
func (v FilesystemProxy) P_Filesystem() unsafe.Pointer      { return v.P }
func FilesystemProxyGetType() gi.GType {
	ret := _I.GetGType(23, "FilesystemProxy")
	return ret
}

// udisks_filesystem_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewFilesystemProxyFinish(res g.IAsyncResult) (result FilesystemProxy, err error) {
	iv, err := _I.Get(208, "FilesystemProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_filesystem_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewFilesystemProxyForBusFinish(res g.IAsyncResult) (result FilesystemProxy, err error) {
	iv, err := _I.Get(209, "FilesystemProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_filesystem_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewFilesystemProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result FilesystemProxy, err error) {
	iv, err := _I.Get(210, "FilesystemProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_filesystem_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewFilesystemProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result FilesystemProxy, err error) {
	iv, err := _I.Get(211, "FilesystemProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_filesystem_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func FilesystemProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(212, "FilesystemProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_filesystem_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func FilesystemProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(213, "FilesystemProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct FilesystemProxyClass

// Struct FilesystemProxyPrivate
type FilesystemProxyPrivate struct {
	P unsafe.Pointer
}

func FilesystemProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(24, "FilesystemProxyPrivate")
	return ret
}

// Object FilesystemSkeleton
type FilesystemSkeleton struct {
	FilesystemIfc
	g.DBusInterfaceSkeleton
}

func WrapFilesystemSkeleton(p unsafe.Pointer) (r FilesystemSkeleton) { r.P = p; return }

type IFilesystemSkeleton interface{ P_FilesystemSkeleton() unsafe.Pointer }

func (v FilesystemSkeleton) P_FilesystemSkeleton() unsafe.Pointer { return v.P }
func (v FilesystemSkeleton) P_Filesystem() unsafe.Pointer         { return v.P }
func FilesystemSkeletonGetType() gi.GType {
	ret := _I.GetGType(25, "FilesystemSkeleton")
	return ret
}

// udisks_filesystem_skeleton_new
//
// [ result ] trans: everything
//
func NewFilesystemSkeleton() (result FilesystemSkeleton) {
	iv, err := _I.Get(214, "FilesystemSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct FilesystemSkeletonClass

// Struct FilesystemSkeletonPrivate
type FilesystemSkeletonPrivate struct {
	P unsafe.Pointer
}

func FilesystemSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(26, "FilesystemSkeletonPrivate")
	return ret
}

// Interface Job
type Job struct {
	JobIfc
	P unsafe.Pointer
}
type JobIfc struct{}
type IJob interface{ P_Job() unsafe.Pointer }

func (v Job) P_Job() unsafe.Pointer { return v.P }
func JobGetType() gi.GType {
	ret := _I.GetGType(27, "Job")
	return ret
}

// udisks_job_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func JobOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(216, "Job", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_job_call_cancel
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *JobIfc) CallCancel(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(217, "Job", "call_cancel")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_job_call_cancel_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *JobIfc) CallCancelFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(218, "Job", "call_cancel_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_job_call_cancel_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *JobIfc) CallCancelSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(219, "Job", "call_cancel_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_job_complete_cancel
//
// [ invocation ] trans: everything
//
func (v *JobIfc) CompleteCancel(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(220, "Job", "complete_cancel")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_job_emit_completed
//
// [ arg_success ] trans: nothing
//
// [ arg_message ] trans: nothing
//
func (v *JobIfc) EmitCompleted(arg_success bool, arg_message string) {
	iv, err := _I.Get(221, "Job", "emit_completed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_message := gi.CString(arg_message)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_success := gi.NewBoolArgument(arg_success)
	arg_arg_message := gi.NewStringArgument(c_arg_message)
	args := []gi.Argument{arg_v, arg_arg_success, arg_arg_message}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_message)
}

// ignore GType struct JobIface

// Object JobProxy
type JobProxy struct {
	JobIfc
	g.DBusProxy
}

func WrapJobProxy(p unsafe.Pointer) (r JobProxy) { r.P = p; return }

type IJobProxy interface{ P_JobProxy() unsafe.Pointer }

func (v JobProxy) P_JobProxy() unsafe.Pointer { return v.P }
func (v JobProxy) P_Job() unsafe.Pointer      { return v.P }
func JobProxyGetType() gi.GType {
	ret := _I.GetGType(28, "JobProxy")
	return ret
}

// udisks_job_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewJobProxyFinish(res g.IAsyncResult) (result JobProxy, err error) {
	iv, err := _I.Get(222, "JobProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_job_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewJobProxyForBusFinish(res g.IAsyncResult) (result JobProxy, err error) {
	iv, err := _I.Get(223, "JobProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_job_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewJobProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result JobProxy, err error) {
	iv, err := _I.Get(224, "JobProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_job_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewJobProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result JobProxy, err error) {
	iv, err := _I.Get(225, "JobProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_job_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func JobProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(226, "JobProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_job_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func JobProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(227, "JobProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct JobProxyClass

// Struct JobProxyPrivate
type JobProxyPrivate struct {
	P unsafe.Pointer
}

func JobProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(29, "JobProxyPrivate")
	return ret
}

// Object JobSkeleton
type JobSkeleton struct {
	JobIfc
	g.DBusInterfaceSkeleton
}

func WrapJobSkeleton(p unsafe.Pointer) (r JobSkeleton) { r.P = p; return }

type IJobSkeleton interface{ P_JobSkeleton() unsafe.Pointer }

func (v JobSkeleton) P_JobSkeleton() unsafe.Pointer { return v.P }
func (v JobSkeleton) P_Job() unsafe.Pointer         { return v.P }
func JobSkeletonGetType() gi.GType {
	ret := _I.GetGType(30, "JobSkeleton")
	return ret
}

// udisks_job_skeleton_new
//
// [ result ] trans: everything
//
func NewJobSkeleton() (result JobSkeleton) {
	iv, err := _I.Get(228, "JobSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct JobSkeletonClass

// Struct JobSkeletonPrivate
type JobSkeletonPrivate struct {
	P unsafe.Pointer
}

func JobSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(31, "JobSkeletonPrivate")
	return ret
}

// Interface Loop
type Loop struct {
	LoopIfc
	P unsafe.Pointer
}
type LoopIfc struct{}
type ILoop interface{ P_Loop() unsafe.Pointer }

func (v Loop) P_Loop() unsafe.Pointer { return v.P }
func LoopGetType() gi.GType {
	ret := _I.GetGType(32, "Loop")
	return ret
}

// udisks_loop_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func LoopOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(230, "Loop", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_loop_call_delete
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *LoopIfc) CallDelete(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(231, "Loop", "call_delete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_loop_call_delete_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *LoopIfc) CallDeleteFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(232, "Loop", "call_delete_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_loop_call_delete_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *LoopIfc) CallDeleteSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(233, "Loop", "call_delete_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_loop_call_set_autoclear
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *LoopIfc) CallSetAutoclear(arg_value bool, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(234, "Loop", "call_set_autoclear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewBoolArgument(arg_value)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_loop_call_set_autoclear_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *LoopIfc) CallSetAutoclearFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(235, "Loop", "call_set_autoclear_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_loop_call_set_autoclear_sync
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *LoopIfc) CallSetAutoclearSync(arg_value bool, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(236, "Loop", "call_set_autoclear_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewBoolArgument(arg_value)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_loop_complete_delete
//
// [ invocation ] trans: everything
//
func (v *LoopIfc) CompleteDelete(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(237, "Loop", "complete_delete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_loop_complete_set_autoclear
//
// [ invocation ] trans: everything
//
func (v *LoopIfc) CompleteSetAutoclear(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(238, "Loop", "complete_set_autoclear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct LoopIface

// Object LoopProxy
type LoopProxy struct {
	LoopIfc
	g.DBusProxy
}

func WrapLoopProxy(p unsafe.Pointer) (r LoopProxy) { r.P = p; return }

type ILoopProxy interface{ P_LoopProxy() unsafe.Pointer }

func (v LoopProxy) P_LoopProxy() unsafe.Pointer { return v.P }
func (v LoopProxy) P_Loop() unsafe.Pointer      { return v.P }
func LoopProxyGetType() gi.GType {
	ret := _I.GetGType(33, "LoopProxy")
	return ret
}

// udisks_loop_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewLoopProxyFinish(res g.IAsyncResult) (result LoopProxy, err error) {
	iv, err := _I.Get(239, "LoopProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_loop_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewLoopProxyForBusFinish(res g.IAsyncResult) (result LoopProxy, err error) {
	iv, err := _I.Get(240, "LoopProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_loop_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewLoopProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result LoopProxy, err error) {
	iv, err := _I.Get(241, "LoopProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_loop_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewLoopProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result LoopProxy, err error) {
	iv, err := _I.Get(242, "LoopProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_loop_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func LoopProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(243, "LoopProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_loop_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func LoopProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(244, "LoopProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct LoopProxyClass

// Struct LoopProxyPrivate
type LoopProxyPrivate struct {
	P unsafe.Pointer
}

func LoopProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(34, "LoopProxyPrivate")
	return ret
}

// Object LoopSkeleton
type LoopSkeleton struct {
	LoopIfc
	g.DBusInterfaceSkeleton
}

func WrapLoopSkeleton(p unsafe.Pointer) (r LoopSkeleton) { r.P = p; return }

type ILoopSkeleton interface{ P_LoopSkeleton() unsafe.Pointer }

func (v LoopSkeleton) P_LoopSkeleton() unsafe.Pointer { return v.P }
func (v LoopSkeleton) P_Loop() unsafe.Pointer         { return v.P }
func LoopSkeletonGetType() gi.GType {
	ret := _I.GetGType(35, "LoopSkeleton")
	return ret
}

// udisks_loop_skeleton_new
//
// [ result ] trans: everything
//
func NewLoopSkeleton() (result LoopSkeleton) {
	iv, err := _I.Get(245, "LoopSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct LoopSkeletonClass

// Struct LoopSkeletonPrivate
type LoopSkeletonPrivate struct {
	P unsafe.Pointer
}

func LoopSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(36, "LoopSkeletonPrivate")
	return ret
}

// Interface MDRaid
type MDRaid struct {
	MDRaidIfc
	P unsafe.Pointer
}
type MDRaidIfc struct{}
type IMDRaid interface{ P_MDRaid() unsafe.Pointer }

func (v MDRaid) P_MDRaid() unsafe.Pointer { return v.P }
func MDRaidGetType() gi.GType {
	ret := _I.GetGType(37, "MDRaid")
	return ret
}

// udisks_mdraid_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func MDRaidOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(247, "MDRaid", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_mdraid_call_add_device
//
// [ arg_device ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallAddDevice(arg_device string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(248, "MDRaid", "call_add_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_device := gi.CString(arg_device)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_device := gi.NewStringArgument(c_arg_device)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_device, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_device)
}

// udisks_mdraid_call_add_device_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallAddDeviceFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(249, "MDRaid", "call_add_device_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_add_device_sync
//
// [ arg_device ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallAddDeviceSync(arg_device string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(250, "MDRaid", "call_add_device_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_device := gi.CString(arg_device)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_device := gi.NewStringArgument(c_arg_device)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_device, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_device)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_delete
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallDelete(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(251, "MDRaid", "call_delete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_call_delete_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallDeleteFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(252, "MDRaid", "call_delete_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_delete_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallDeleteSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(253, "MDRaid", "call_delete_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_remove_device
//
// [ arg_device ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallRemoveDevice(arg_device string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(254, "MDRaid", "call_remove_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_device := gi.CString(arg_device)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_device := gi.NewStringArgument(c_arg_device)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_device, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_device)
}

// udisks_mdraid_call_remove_device_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallRemoveDeviceFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(255, "MDRaid", "call_remove_device_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_remove_device_sync
//
// [ arg_device ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallRemoveDeviceSync(arg_device string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(256, "MDRaid", "call_remove_device_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_device := gi.CString(arg_device)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_device := gi.NewStringArgument(c_arg_device)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_device, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_device)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_request_sync_action
//
// [ arg_sync_action ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallRequestSyncAction(arg_sync_action string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(257, "MDRaid", "call_request_sync_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_sync_action := gi.CString(arg_sync_action)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_sync_action := gi.NewStringArgument(c_arg_sync_action)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_sync_action, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_sync_action)
}

// udisks_mdraid_call_request_sync_action_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallRequestSyncActionFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(258, "MDRaid", "call_request_sync_action_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_request_sync_action_sync
//
// [ arg_sync_action ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallRequestSyncActionSync(arg_sync_action string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(259, "MDRaid", "call_request_sync_action_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_sync_action := gi.CString(arg_sync_action)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_sync_action := gi.NewStringArgument(c_arg_sync_action)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_sync_action, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_sync_action)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_set_bitmap_location
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallSetBitmapLocation(arg_value string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(260, "MDRaid", "call_set_bitmap_location")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_value := gi.CString(arg_value)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewStringArgument(c_arg_value)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_value)
}

// udisks_mdraid_call_set_bitmap_location_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallSetBitmapLocationFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(261, "MDRaid", "call_set_bitmap_location_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_set_bitmap_location_sync
//
// [ arg_value ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallSetBitmapLocationSync(arg_value string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(262, "MDRaid", "call_set_bitmap_location_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_value := gi.CString(arg_value)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_value := gi.NewStringArgument(c_arg_value)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_value, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_value)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_start
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallStart(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(263, "MDRaid", "call_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_call_start_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallStartFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(264, "MDRaid", "call_start_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_start_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallStartSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(265, "MDRaid", "call_start_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_stop
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *MDRaidIfc) CallStop(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(266, "MDRaid", "call_stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_call_stop_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallStopFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(267, "MDRaid", "call_stop_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_call_stop_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *MDRaidIfc) CallStopSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(268, "MDRaid", "call_stop_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_mdraid_complete_add_device
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteAddDevice(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(269, "MDRaid", "complete_add_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_complete_delete
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteDelete(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(270, "MDRaid", "complete_delete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_complete_remove_device
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteRemoveDevice(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(271, "MDRaid", "complete_remove_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_complete_request_sync_action
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteRequestSyncAction(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(272, "MDRaid", "complete_request_sync_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_complete_set_bitmap_location
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteSetBitmapLocation(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(273, "MDRaid", "complete_set_bitmap_location")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_complete_start
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteStart(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(274, "MDRaid", "complete_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_mdraid_complete_stop
//
// [ invocation ] trans: everything
//
func (v *MDRaidIfc) CompleteStop(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(275, "MDRaid", "complete_stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct MDRaidIface

// Object MDRaidProxy
type MDRaidProxy struct {
	MDRaidIfc
	g.DBusProxy
}

func WrapMDRaidProxy(p unsafe.Pointer) (r MDRaidProxy) { r.P = p; return }

type IMDRaidProxy interface{ P_MDRaidProxy() unsafe.Pointer }

func (v MDRaidProxy) P_MDRaidProxy() unsafe.Pointer { return v.P }
func (v MDRaidProxy) P_MDRaid() unsafe.Pointer      { return v.P }
func MDRaidProxyGetType() gi.GType {
	ret := _I.GetGType(38, "MDRaidProxy")
	return ret
}

// udisks_mdraid_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewMDRaidProxyFinish(res g.IAsyncResult) (result MDRaidProxy, err error) {
	iv, err := _I.Get(276, "MDRaidProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_mdraid_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewMDRaidProxyForBusFinish(res g.IAsyncResult) (result MDRaidProxy, err error) {
	iv, err := _I.Get(277, "MDRaidProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_mdraid_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewMDRaidProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result MDRaidProxy, err error) {
	iv, err := _I.Get(278, "MDRaidProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_mdraid_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewMDRaidProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result MDRaidProxy, err error) {
	iv, err := _I.Get(279, "MDRaidProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_mdraid_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func MDRaidProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(280, "MDRaidProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_mdraid_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func MDRaidProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(281, "MDRaidProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct MDRaidProxyClass

// Struct MDRaidProxyPrivate
type MDRaidProxyPrivate struct {
	P unsafe.Pointer
}

func MDRaidProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(39, "MDRaidProxyPrivate")
	return ret
}

// Object MDRaidSkeleton
type MDRaidSkeleton struct {
	MDRaidIfc
	g.DBusInterfaceSkeleton
}

func WrapMDRaidSkeleton(p unsafe.Pointer) (r MDRaidSkeleton) { r.P = p; return }

type IMDRaidSkeleton interface{ P_MDRaidSkeleton() unsafe.Pointer }

func (v MDRaidSkeleton) P_MDRaidSkeleton() unsafe.Pointer { return v.P }
func (v MDRaidSkeleton) P_MDRaid() unsafe.Pointer         { return v.P }
func MDRaidSkeletonGetType() gi.GType {
	ret := _I.GetGType(40, "MDRaidSkeleton")
	return ret
}

// udisks_mdraid_skeleton_new
//
// [ result ] trans: everything
//
func NewMDRaidSkeleton() (result MDRaidSkeleton) {
	iv, err := _I.Get(282, "MDRaidSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct MDRaidSkeletonClass

// Struct MDRaidSkeletonPrivate
type MDRaidSkeletonPrivate struct {
	P unsafe.Pointer
}

func MDRaidSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(41, "MDRaidSkeletonPrivate")
	return ret
}

// Interface Manager
type Manager struct {
	ManagerIfc
	P unsafe.Pointer
}
type ManagerIfc struct{}
type IManager interface{ P_Manager() unsafe.Pointer }

func (v Manager) P_Manager() unsafe.Pointer { return v.P }
func ManagerGetType() gi.GType {
	ret := _I.GetGType(42, "Manager")
	return ret
}

// udisks_manager_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func ManagerOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(284, "Manager", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_manager_call_can_check
//
// [ arg_type ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallCanCheck(arg_type string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(285, "Manager", "call_can_check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_manager_call_can_check_finish
//
// [ out_available ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanCheckFinish(res g.IAsyncResult) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(286, "Manager", "call_can_check_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_available, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_check_sync
//
// [ arg_type ] trans: nothing
//
// [ out_available ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanCheckSync(arg_type string, cancellable g.ICancellable) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(287, "Manager", "call_can_check_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_out_available, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_format
//
// [ arg_type ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallCanFormat(arg_type string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(288, "Manager", "call_can_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_manager_call_can_format_finish
//
// [ out_available ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanFormatFinish(res g.IAsyncResult) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(289, "Manager", "call_can_format_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_available, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_format_sync
//
// [ arg_type ] trans: nothing
//
// [ out_available ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanFormatSync(arg_type string, cancellable g.ICancellable) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(290, "Manager", "call_can_format_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_out_available, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_repair
//
// [ arg_type ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallCanRepair(arg_type string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(291, "Manager", "call_can_repair")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_manager_call_can_repair_finish
//
// [ out_available ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanRepairFinish(res g.IAsyncResult) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(292, "Manager", "call_can_repair_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_available, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_repair_sync
//
// [ arg_type ] trans: nothing
//
// [ out_available ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanRepairSync(arg_type string, cancellable g.ICancellable) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(293, "Manager", "call_can_repair_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_out_available, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_resize
//
// [ arg_type ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallCanResize(arg_type string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(294, "Manager", "call_can_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_manager_call_can_resize_finish
//
// [ out_available ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanResizeFinish(res g.IAsyncResult) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(295, "Manager", "call_can_resize_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_available, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_can_resize_sync
//
// [ arg_type ] trans: nothing
//
// [ out_available ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallCanResizeSync(arg_type string, cancellable g.ICancellable) (result bool, out_available g.Variant, err error) {
	iv, err := _I.Get(296, "Manager", "call_can_resize_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_out_available := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_out_available, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[1].Pointer())
	out_available.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_enable_modules
//
// [ arg_enable ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallEnableModules(arg_enable bool, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(297, "Manager", "call_enable_modules")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_enable := gi.NewBoolArgument(arg_enable)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_enable, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_manager_call_enable_modules_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallEnableModulesFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(298, "Manager", "call_enable_modules_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_manager_call_enable_modules_sync
//
// [ arg_enable ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallEnableModulesSync(arg_enable bool, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(299, "Manager", "call_enable_modules_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_enable := gi.NewBoolArgument(arg_enable)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_enable, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_manager_call_get_block_devices
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallGetBlockDevices(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(300, "Manager", "call_get_block_devices")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_manager_call_get_block_devices_finish
//
// [ out_block_objects ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallGetBlockDevicesFinish(res g.IAsyncResult) (result bool, out_block_objects gi.CStrArray, err error) {
	iv, err := _I.Get(301, "Manager", "call_get_block_devices_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_block_objects := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_block_objects, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_block_objects.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_get_block_devices_sync
//
// [ arg_options ] trans: nothing
//
// [ out_block_objects ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallGetBlockDevicesSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, out_block_objects gi.CStrArray, err error) {
	iv, err := _I.Get(302, "Manager", "call_get_block_devices_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_block_objects := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_out_block_objects, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_block_objects.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_loop_setup
//
// [ arg_fd ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallLoopSetup(arg_fd g.Variant, arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(303, "Manager", "call_loop_setup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_fd := gi.NewPointerArgument(arg_fd.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_fd, arg_arg_options, arg_fd_list, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_manager_call_loop_setup_finish
//
// [ out_resulting_device ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallLoopSetupFinish(res g.IAsyncResult) (result bool, out_resulting_device string, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(304, "Manager", "call_loop_setup_finish")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_resulting_device := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_out_resulting_device, arg_out_fd_list, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_resulting_device = outArgs[0].String().Take()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_loop_setup_sync
//
// [ arg_fd ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ fd_list ] trans: nothing
//
// [ out_resulting_device ] trans: everything, dir: out
//
// [ out_fd_list ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallLoopSetupSync(arg_fd g.Variant, arg_options g.Variant, fd_list g.IUnixFDList, cancellable g.ICancellable) (result bool, out_resulting_device string, out_fd_list g.UnixFDList, err error) {
	iv, err := _I.Get(305, "Manager", "call_loop_setup_sync")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if fd_list != nil {
		tmp = fd_list.P_UnixFDList()
	}
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_fd := gi.NewPointerArgument(arg_fd.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_fd_list := gi.NewPointerArgument(tmp)
	arg_out_resulting_device := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_out_fd_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_arg_fd, arg_arg_options, arg_fd_list, arg_out_resulting_device, arg_out_fd_list, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	out_resulting_device = outArgs[0].String().Take()
	out_fd_list.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_mdraid_create
//
// [ arg_blocks ] trans: nothing
//
// [ arg_level ] trans: nothing
//
// [ arg_name ] trans: nothing
//
// [ arg_chunk ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallMdraidCreate(arg_blocks string, arg_level string, arg_name string, arg_chunk uint64, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(306, "Manager", "call_mdraid_create")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_blocks := gi.CString(arg_blocks)
	c_arg_level := gi.CString(arg_level)
	c_arg_name := gi.CString(arg_name)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_blocks := gi.NewStringArgument(c_arg_blocks)
	arg_arg_level := gi.NewStringArgument(c_arg_level)
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_chunk := gi.NewUint64Argument(arg_chunk)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_blocks, arg_arg_level, arg_arg_name, arg_arg_chunk, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_blocks)
	gi.Free(c_arg_level)
	gi.Free(c_arg_name)
}

// udisks_manager_call_mdraid_create_finish
//
// [ out_resulting_array ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallMdraidCreateFinish(res g.IAsyncResult) (result bool, out_resulting_array string, err error) {
	iv, err := _I.Get(307, "Manager", "call_mdraid_create_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_resulting_array := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_resulting_array, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_resulting_array = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_manager_call_mdraid_create_sync
//
// [ arg_blocks ] trans: nothing
//
// [ arg_level ] trans: nothing
//
// [ arg_name ] trans: nothing
//
// [ arg_chunk ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ out_resulting_array ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallMdraidCreateSync(arg_blocks string, arg_level string, arg_name string, arg_chunk uint64, arg_options g.Variant, cancellable g.ICancellable) (result bool, out_resulting_array string, err error) {
	iv, err := _I.Get(308, "Manager", "call_mdraid_create_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_blocks := gi.CString(arg_blocks)
	c_arg_level := gi.CString(arg_level)
	c_arg_name := gi.CString(arg_name)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_blocks := gi.NewStringArgument(c_arg_blocks)
	arg_arg_level := gi.NewStringArgument(c_arg_level)
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_chunk := gi.NewUint64Argument(arg_chunk)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_resulting_array := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_blocks, arg_arg_level, arg_arg_name, arg_arg_chunk, arg_arg_options, arg_out_resulting_array, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_blocks)
	gi.Free(c_arg_level)
	gi.Free(c_arg_name)
	err = gi.ToError(outArgs[1].Pointer())
	out_resulting_array = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_manager_call_resolve_device
//
// [ arg_devspec ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *ManagerIfc) CallResolveDevice(arg_devspec g.Variant, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(309, "Manager", "call_resolve_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_devspec := gi.NewPointerArgument(arg_devspec.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_devspec, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_manager_call_resolve_device_finish
//
// [ out_devices ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallResolveDeviceFinish(res g.IAsyncResult) (result bool, out_devices gi.CStrArray, err error) {
	iv, err := _I.Get(310, "Manager", "call_resolve_device_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_devices := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_devices, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_devices.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_call_resolve_device_sync
//
// [ arg_devspec ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ out_devices ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ManagerIfc) CallResolveDeviceSync(arg_devspec g.Variant, arg_options g.Variant, cancellable g.ICancellable) (result bool, out_devices gi.CStrArray, err error) {
	iv, err := _I.Get(311, "Manager", "call_resolve_device_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_devspec := gi.NewPointerArgument(arg_devspec.P)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_devices := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_devspec, arg_arg_options, arg_out_devices, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_devices.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// udisks_manager_complete_can_check
//
// [ invocation ] trans: everything
//
// [ available ] trans: nothing
//
func (v *ManagerIfc) CompleteCanCheck(invocation g.IDBusMethodInvocation, available g.Variant) {
	iv, err := _I.Get(312, "Manager", "complete_can_check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_available := gi.NewPointerArgument(available.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_available}
	iv.Call(args, nil, nil)
}

// udisks_manager_complete_can_format
//
// [ invocation ] trans: everything
//
// [ available ] trans: nothing
//
func (v *ManagerIfc) CompleteCanFormat(invocation g.IDBusMethodInvocation, available g.Variant) {
	iv, err := _I.Get(313, "Manager", "complete_can_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_available := gi.NewPointerArgument(available.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_available}
	iv.Call(args, nil, nil)
}

// udisks_manager_complete_can_repair
//
// [ invocation ] trans: everything
//
// [ available ] trans: nothing
//
func (v *ManagerIfc) CompleteCanRepair(invocation g.IDBusMethodInvocation, available g.Variant) {
	iv, err := _I.Get(314, "Manager", "complete_can_repair")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_available := gi.NewPointerArgument(available.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_available}
	iv.Call(args, nil, nil)
}

// udisks_manager_complete_can_resize
//
// [ invocation ] trans: everything
//
// [ available ] trans: nothing
//
func (v *ManagerIfc) CompleteCanResize(invocation g.IDBusMethodInvocation, available g.Variant) {
	iv, err := _I.Get(315, "Manager", "complete_can_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_available := gi.NewPointerArgument(available.P)
	args := []gi.Argument{arg_v, arg_invocation, arg_available}
	iv.Call(args, nil, nil)
}

// udisks_manager_complete_enable_modules
//
// [ invocation ] trans: everything
//
func (v *ManagerIfc) CompleteEnableModules(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(316, "Manager", "complete_enable_modules")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_manager_complete_get_block_devices
//
// [ invocation ] trans: everything
//
// [ block_objects ] trans: nothing
//
func (v *ManagerIfc) CompleteGetBlockDevices(invocation g.IDBusMethodInvocation, block_objects string) {
	iv, err := _I.Get(317, "Manager", "complete_get_block_devices")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_block_objects := gi.CString(block_objects)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_block_objects := gi.NewStringArgument(c_block_objects)
	args := []gi.Argument{arg_v, arg_invocation, arg_block_objects}
	iv.Call(args, nil, nil)
	gi.Free(c_block_objects)
}

// udisks_manager_complete_loop_setup
//
// [ invocation ] trans: everything
//
// [ fd_list ] trans: nothing
//
// [ resulting_device ] trans: nothing
//
func (v *ManagerIfc) CompleteLoopSetup(invocation g.IDBusMethodInvocation, fd_list g.IUnixFDList, resulting_device string) {
	iv, err := _I.Get(318, "Manager", "complete_loop_setup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	var tmp1 unsafe.Pointer
	if fd_list != nil {
		tmp1 = fd_list.P_UnixFDList()
	}
	c_resulting_device := gi.CString(resulting_device)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_fd_list := gi.NewPointerArgument(tmp1)
	arg_resulting_device := gi.NewStringArgument(c_resulting_device)
	args := []gi.Argument{arg_v, arg_invocation, arg_fd_list, arg_resulting_device}
	iv.Call(args, nil, nil)
	gi.Free(c_resulting_device)
}

// udisks_manager_complete_mdraid_create
//
// [ invocation ] trans: everything
//
// [ resulting_array ] trans: nothing
//
func (v *ManagerIfc) CompleteMdraidCreate(invocation g.IDBusMethodInvocation, resulting_array string) {
	iv, err := _I.Get(319, "Manager", "complete_mdraid_create")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_resulting_array := gi.CString(resulting_array)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_resulting_array := gi.NewStringArgument(c_resulting_array)
	args := []gi.Argument{arg_v, arg_invocation, arg_resulting_array}
	iv.Call(args, nil, nil)
	gi.Free(c_resulting_array)
}

// udisks_manager_complete_resolve_device
//
// [ invocation ] trans: everything
//
// [ devices ] trans: nothing
//
func (v *ManagerIfc) CompleteResolveDevice(invocation g.IDBusMethodInvocation, devices string) {
	iv, err := _I.Get(320, "Manager", "complete_resolve_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_devices := gi.CString(devices)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_devices := gi.NewStringArgument(c_devices)
	args := []gi.Argument{arg_v, arg_invocation, arg_devices}
	iv.Call(args, nil, nil)
	gi.Free(c_devices)
}

// ignore GType struct ManagerIface

// Object ManagerProxy
type ManagerProxy struct {
	ManagerIfc
	g.DBusProxy
}

func WrapManagerProxy(p unsafe.Pointer) (r ManagerProxy) { r.P = p; return }

type IManagerProxy interface{ P_ManagerProxy() unsafe.Pointer }

func (v ManagerProxy) P_ManagerProxy() unsafe.Pointer { return v.P }
func (v ManagerProxy) P_Manager() unsafe.Pointer      { return v.P }
func ManagerProxyGetType() gi.GType {
	ret := _I.GetGType(43, "ManagerProxy")
	return ret
}

// udisks_manager_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewManagerProxyFinish(res g.IAsyncResult) (result ManagerProxy, err error) {
	iv, err := _I.Get(321, "ManagerProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_manager_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewManagerProxyForBusFinish(res g.IAsyncResult) (result ManagerProxy, err error) {
	iv, err := _I.Get(322, "ManagerProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_manager_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewManagerProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result ManagerProxy, err error) {
	iv, err := _I.Get(323, "ManagerProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_manager_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewManagerProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result ManagerProxy, err error) {
	iv, err := _I.Get(324, "ManagerProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_manager_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func ManagerProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(325, "ManagerProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_manager_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func ManagerProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(326, "ManagerProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct ManagerProxyClass

// Struct ManagerProxyPrivate
type ManagerProxyPrivate struct {
	P unsafe.Pointer
}

func ManagerProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(44, "ManagerProxyPrivate")
	return ret
}

// Object ManagerSkeleton
type ManagerSkeleton struct {
	ManagerIfc
	g.DBusInterfaceSkeleton
}

func WrapManagerSkeleton(p unsafe.Pointer) (r ManagerSkeleton) { r.P = p; return }

type IManagerSkeleton interface{ P_ManagerSkeleton() unsafe.Pointer }

func (v ManagerSkeleton) P_ManagerSkeleton() unsafe.Pointer { return v.P }
func (v ManagerSkeleton) P_Manager() unsafe.Pointer         { return v.P }
func ManagerSkeletonGetType() gi.GType {
	ret := _I.GetGType(45, "ManagerSkeleton")
	return ret
}

// udisks_manager_skeleton_new
//
// [ result ] trans: everything
//
func NewManagerSkeleton() (result ManagerSkeleton) {
	iv, err := _I.Get(327, "ManagerSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct ManagerSkeletonClass

// Struct ManagerSkeletonPrivate
type ManagerSkeletonPrivate struct {
	P unsafe.Pointer
}

func ManagerSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(46, "ManagerSkeletonPrivate")
	return ret
}

// Interface Object
type Object struct {
	ObjectIfc
	P unsafe.Pointer
}
type ObjectIfc struct{}
type IObject interface{ P_Object() unsafe.Pointer }

func (v Object) P_Object() unsafe.Pointer { return v.P }
func ObjectGetType() gi.GType {
	ret := _I.GetGType(47, "Object")
	return ret
}

// udisks_object_get_block
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetBlock() (result Block) {
	iv, err := _I.Get(328, "Object", "get_block")
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

// udisks_object_get_drive
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetDrive() (result Drive) {
	iv, err := _I.Get(329, "Object", "get_drive")
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

// udisks_object_get_drive_ata
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetDriveAta() (result DriveAta) {
	iv, err := _I.Get(330, "Object", "get_drive_ata")
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

// udisks_object_get_encrypted
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetEncrypted() (result Encrypted) {
	iv, err := _I.Get(331, "Object", "get_encrypted")
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

// udisks_object_get_filesystem
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetFilesystem() (result Filesystem) {
	iv, err := _I.Get(332, "Object", "get_filesystem")
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

// udisks_object_get_job
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetJob() (result Job) {
	iv, err := _I.Get(333, "Object", "get_job")
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

// udisks_object_get_loop
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetLoop() (result Loop) {
	iv, err := _I.Get(334, "Object", "get_loop")
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

// udisks_object_get_manager
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetManager() (result Manager) {
	iv, err := _I.Get(335, "Object", "get_manager")
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

// udisks_object_get_mdraid
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetMdraid() (result MDRaid) {
	iv, err := _I.Get(336, "Object", "get_mdraid")
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

// udisks_object_get_partition
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetPartition() (result Partition) {
	iv, err := _I.Get(337, "Object", "get_partition")
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

// udisks_object_get_partition_table
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetPartitionTable() (result PartitionTable) {
	iv, err := _I.Get(338, "Object", "get_partition_table")
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

// udisks_object_get_swapspace
//
// [ result ] trans: everything
//
func (v *ObjectIfc) GetSwapspace() (result Swapspace) {
	iv, err := _I.Get(339, "Object", "get_swapspace")
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

// ignore GType struct ObjectIface

// Object ObjectInfo
type ObjectInfo struct {
	g.Object
}

func WrapObjectInfo(p unsafe.Pointer) (r ObjectInfo) { r.P = p; return }

type IObjectInfo interface{ P_ObjectInfo() unsafe.Pointer }

func (v ObjectInfo) P_ObjectInfo() unsafe.Pointer { return v.P }
func ObjectInfoGetType() gi.GType {
	ret := _I.GetGType(48, "ObjectInfo")
	return ret
}

// udisks_object_info_get_description
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetDescription() (result string) {
	iv, err := _I.Get(340, "ObjectInfo", "get_description")
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

// udisks_object_info_get_icon
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetIcon() (result g.Icon) {
	iv, err := _I.Get(341, "ObjectInfo", "get_icon")
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

// udisks_object_info_get_icon_symbolic
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetIconSymbolic() (result g.Icon) {
	iv, err := _I.Get(342, "ObjectInfo", "get_icon_symbolic")
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

// udisks_object_info_get_media_description
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetMediaDescription() (result string) {
	iv, err := _I.Get(343, "ObjectInfo", "get_media_description")
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

// udisks_object_info_get_media_icon
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetMediaIcon() (result g.Icon) {
	iv, err := _I.Get(344, "ObjectInfo", "get_media_icon")
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

// udisks_object_info_get_media_icon_symbolic
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetMediaIconSymbolic() (result g.Icon) {
	iv, err := _I.Get(345, "ObjectInfo", "get_media_icon_symbolic")
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

// udisks_object_info_get_name
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetName() (result string) {
	iv, err := _I.Get(346, "ObjectInfo", "get_name")
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

// udisks_object_info_get_object
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetObject() (result Object) {
	iv, err := _I.Get(347, "ObjectInfo", "get_object")
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

// udisks_object_info_get_one_liner
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetOneLiner() (result string) {
	iv, err := _I.Get(348, "ObjectInfo", "get_one_liner")
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

// udisks_object_info_get_sort_key
//
// [ result ] trans: nothing
//
func (v ObjectInfo) GetSortKey() (result string) {
	iv, err := _I.Get(349, "ObjectInfo", "get_sort_key")
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

// Object ObjectManagerClient
type ObjectManagerClient struct {
	g.DBusObjectManagerClient
}

func WrapObjectManagerClient(p unsafe.Pointer) (r ObjectManagerClient) { r.P = p; return }

type IObjectManagerClient interface{ P_ObjectManagerClient() unsafe.Pointer }

func (v ObjectManagerClient) P_ObjectManagerClient() unsafe.Pointer { return v.P }
func ObjectManagerClientGetType() gi.GType {
	ret := _I.GetGType(49, "ObjectManagerClient")
	return ret
}

// udisks_object_manager_client_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectManagerClientFinish(res g.IAsyncResult) (result ObjectManagerClient, err error) {
	iv, err := _I.Get(350, "ObjectManagerClient", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_object_manager_client_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectManagerClientForBusFinish(res g.IAsyncResult) (result ObjectManagerClient, err error) {
	iv, err := _I.Get(351, "ObjectManagerClient", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_object_manager_client_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectManagerClientForBusSync(bus_type g.BusTypeEnum, flags g.DBusObjectManagerClientFlags, name string, object_path string, cancellable g.ICancellable) (result ObjectManagerClient, err error) {
	iv, err := _I.Get(352, "ObjectManagerClient", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_object_manager_client_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectManagerClientSync(connection g.IDBusConnection, flags g.DBusObjectManagerClientFlags, name string, object_path string, cancellable g.ICancellable) (result ObjectManagerClient, err error) {
	iv, err := _I.Get(353, "ObjectManagerClient", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_object_manager_client_get_proxy_type
//
// [ manager ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ interface_name ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectManagerClientGetProxyType1(manager g.IDBusObjectManagerClient, object_path string, interface_name string, user_data unsafe.Pointer) (result gi.GType) {
	iv, err := _I.Get(354, "ObjectManagerClient", "get_proxy_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if manager != nil {
		tmp = manager.P_DBusObjectManagerClient()
	}
	c_object_path := gi.CString(object_path)
	c_interface_name := gi.CString(interface_name)
	arg_manager := gi.NewPointerArgument(tmp)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_interface_name := gi.NewStringArgument(c_interface_name)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_manager, arg_object_path, arg_interface_name, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_object_path)
	gi.Free(c_interface_name)
	result = gi.GType(ret.Uint())
	return
}

// udisks_object_manager_client_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func ObjectManagerClientNew1(connection g.IDBusConnection, flags g.DBusObjectManagerClientFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(355, "ObjectManagerClient", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_object_manager_client_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func ObjectManagerClientNewForBus1(bus_type g.BusTypeEnum, flags g.DBusObjectManagerClientFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(356, "ObjectManagerClient", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct ObjectManagerClientClass

// Struct ObjectManagerClientPrivate
type ObjectManagerClientPrivate struct {
	P unsafe.Pointer
}

func ObjectManagerClientPrivateGetType() gi.GType {
	ret := _I.GetGType(50, "ObjectManagerClientPrivate")
	return ret
}

// Object ObjectProxy
type ObjectProxy struct {
	ObjectIfc
	g.DBusObjectProxy
}

func WrapObjectProxy(p unsafe.Pointer) (r ObjectProxy) { r.P = p; return }

type IObjectProxy interface{ P_ObjectProxy() unsafe.Pointer }

func (v ObjectProxy) P_ObjectProxy() unsafe.Pointer { return v.P }
func (v ObjectProxy) P_Object() unsafe.Pointer      { return v.P }
func ObjectProxyGetType() gi.GType {
	ret := _I.GetGType(51, "ObjectProxy")
	return ret
}

// udisks_object_proxy_new
//
// [ connection ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectProxy(connection g.IDBusConnection, object_path string) (result ObjectProxy) {
	iv, err := _I.Get(357, "ObjectProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_object_path := gi.CString(object_path)
	arg_connection := gi.NewPointerArgument(tmp)
	arg_object_path := gi.NewStringArgument(c_object_path)
	args := []gi.Argument{arg_connection, arg_object_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_object_path)
	result.P = ret.Pointer()
	return
}

// ignore GType struct ObjectProxyClass

// Struct ObjectProxyPrivate
type ObjectProxyPrivate struct {
	P unsafe.Pointer
}

func ObjectProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(52, "ObjectProxyPrivate")
	return ret
}

// Object ObjectSkeleton
type ObjectSkeleton struct {
	ObjectIfc
	g.DBusObjectSkeleton
}

func WrapObjectSkeleton(p unsafe.Pointer) (r ObjectSkeleton) { r.P = p; return }

type IObjectSkeleton interface{ P_ObjectSkeleton() unsafe.Pointer }

func (v ObjectSkeleton) P_ObjectSkeleton() unsafe.Pointer { return v.P }
func (v ObjectSkeleton) P_Object() unsafe.Pointer         { return v.P }
func ObjectSkeletonGetType() gi.GType {
	ret := _I.GetGType(53, "ObjectSkeleton")
	return ret
}

// udisks_object_skeleton_new
//
// [ object_path ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectSkeleton(object_path string) (result ObjectSkeleton) {
	iv, err := _I.Get(358, "ObjectSkeleton", "new")
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

// udisks_object_skeleton_set_block
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetBlock(interface_ IBlock) {
	iv, err := _I.Get(359, "ObjectSkeleton", "set_block")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Block()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_drive
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetDrive(interface_ IDrive) {
	iv, err := _I.Get(360, "ObjectSkeleton", "set_drive")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Drive()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_drive_ata
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetDriveAta(interface_ IDriveAta) {
	iv, err := _I.Get(361, "ObjectSkeleton", "set_drive_ata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_DriveAta()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_encrypted
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetEncrypted(interface_ IEncrypted) {
	iv, err := _I.Get(362, "ObjectSkeleton", "set_encrypted")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Encrypted()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_filesystem
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetFilesystem(interface_ IFilesystem) {
	iv, err := _I.Get(363, "ObjectSkeleton", "set_filesystem")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Filesystem()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_job
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetJob(interface_ IJob) {
	iv, err := _I.Get(364, "ObjectSkeleton", "set_job")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Job()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_loop
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetLoop(interface_ ILoop) {
	iv, err := _I.Get(365, "ObjectSkeleton", "set_loop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Loop()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_manager
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetManager(interface_ IManager) {
	iv, err := _I.Get(366, "ObjectSkeleton", "set_manager")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Manager()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_mdraid
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetMdraid(interface_ IMDRaid) {
	iv, err := _I.Get(367, "ObjectSkeleton", "set_mdraid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_MDRaid()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_partition
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetPartition(interface_ IPartition) {
	iv, err := _I.Get(368, "ObjectSkeleton", "set_partition")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Partition()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_partition_table
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetPartitionTable(interface_ IPartitionTable) {
	iv, err := _I.Get(369, "ObjectSkeleton", "set_partition_table")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_PartitionTable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// udisks_object_skeleton_set_swapspace
//
// [ interface_ ] trans: nothing
//
func (v ObjectSkeleton) SetSwapspace(interface_ ISwapspace) {
	iv, err := _I.Get(370, "ObjectSkeleton", "set_swapspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if interface_ != nil {
		tmp = interface_.P_Swapspace()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interface_ := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_interface_}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectSkeletonClass

// Struct ObjectSkeletonPrivate
type ObjectSkeletonPrivate struct {
	P unsafe.Pointer
}

func ObjectSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(54, "ObjectSkeletonPrivate")
	return ret
}

// Interface Partition
type Partition struct {
	PartitionIfc
	P unsafe.Pointer
}
type PartitionIfc struct{}
type IPartition interface{ P_Partition() unsafe.Pointer }

func (v Partition) P_Partition() unsafe.Pointer { return v.P }
func PartitionGetType() gi.GType {
	ret := _I.GetGType(55, "Partition")
	return ret
}

// udisks_partition_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func PartitionOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(372, "Partition", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_partition_call_delete
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionIfc) CallDelete(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(373, "Partition", "call_delete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_partition_call_delete_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallDeleteFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(374, "Partition", "call_delete_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_delete_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallDeleteSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(375, "Partition", "call_delete_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_resize
//
// [ arg_size ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionIfc) CallResize(arg_size uint64, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(376, "Partition", "call_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_size, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_partition_call_resize_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallResizeFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(377, "Partition", "call_resize_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_resize_sync
//
// [ arg_size ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallResizeSync(arg_size uint64, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(378, "Partition", "call_resize_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_size, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_set_flags
//
// [ arg_flags ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionIfc) CallSetFlags(arg_flags uint64, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(379, "Partition", "call_set_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_flags := gi.NewUint64Argument(arg_flags)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_flags, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_partition_call_set_flags_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallSetFlagsFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(380, "Partition", "call_set_flags_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_set_flags_sync
//
// [ arg_flags ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallSetFlagsSync(arg_flags uint64, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(381, "Partition", "call_set_flags_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_flags := gi.NewUint64Argument(arg_flags)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_flags, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_set_name
//
// [ arg_name ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionIfc) CallSetName(arg_name string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(382, "Partition", "call_set_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_name := gi.CString(arg_name)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_name, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_name)
}

// udisks_partition_call_set_name_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallSetNameFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(383, "Partition", "call_set_name_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_set_name_sync
//
// [ arg_name ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallSetNameSync(arg_name string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(384, "Partition", "call_set_name_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_name := gi.CString(arg_name)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_name, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_set_type
//
// [ arg_type ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionIfc) CallSetType(arg_type string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(385, "Partition", "call_set_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_type, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
}

// udisks_partition_call_set_type_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallSetTypeFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(386, "Partition", "call_set_type_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_call_set_type_sync
//
// [ arg_type ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionIfc) CallSetTypeSync(arg_type string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(387, "Partition", "call_set_type_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_type := gi.CString(arg_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_type, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_partition_complete_delete
//
// [ invocation ] trans: everything
//
func (v *PartitionIfc) CompleteDelete(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(388, "Partition", "complete_delete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_partition_complete_resize
//
// [ invocation ] trans: everything
//
func (v *PartitionIfc) CompleteResize(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(389, "Partition", "complete_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_partition_complete_set_flags
//
// [ invocation ] trans: everything
//
func (v *PartitionIfc) CompleteSetFlags(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(390, "Partition", "complete_set_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_partition_complete_set_name
//
// [ invocation ] trans: everything
//
func (v *PartitionIfc) CompleteSetName(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(391, "Partition", "complete_set_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_partition_complete_set_type
//
// [ invocation ] trans: everything
//
func (v *PartitionIfc) CompleteSetType(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(392, "Partition", "complete_set_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct PartitionIface

// Object PartitionProxy
type PartitionProxy struct {
	PartitionIfc
	g.DBusProxy
}

func WrapPartitionProxy(p unsafe.Pointer) (r PartitionProxy) { r.P = p; return }

type IPartitionProxy interface{ P_PartitionProxy() unsafe.Pointer }

func (v PartitionProxy) P_PartitionProxy() unsafe.Pointer { return v.P }
func (v PartitionProxy) P_Partition() unsafe.Pointer      { return v.P }
func PartitionProxyGetType() gi.GType {
	ret := _I.GetGType(56, "PartitionProxy")
	return ret
}

// udisks_partition_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionProxyFinish(res g.IAsyncResult) (result PartitionProxy, err error) {
	iv, err := _I.Get(393, "PartitionProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionProxyForBusFinish(res g.IAsyncResult) (result PartitionProxy, err error) {
	iv, err := _I.Get(394, "PartitionProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result PartitionProxy, err error) {
	iv, err := _I.Get(395, "PartitionProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result PartitionProxy, err error) {
	iv, err := _I.Get(396, "PartitionProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func PartitionProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(397, "PartitionProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_partition_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func PartitionProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(398, "PartitionProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct PartitionProxyClass

// Struct PartitionProxyPrivate
type PartitionProxyPrivate struct {
	P unsafe.Pointer
}

func PartitionProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(57, "PartitionProxyPrivate")
	return ret
}

// Object PartitionSkeleton
type PartitionSkeleton struct {
	PartitionIfc
	g.DBusInterfaceSkeleton
}

func WrapPartitionSkeleton(p unsafe.Pointer) (r PartitionSkeleton) { r.P = p; return }

type IPartitionSkeleton interface{ P_PartitionSkeleton() unsafe.Pointer }

func (v PartitionSkeleton) P_PartitionSkeleton() unsafe.Pointer { return v.P }
func (v PartitionSkeleton) P_Partition() unsafe.Pointer         { return v.P }
func PartitionSkeletonGetType() gi.GType {
	ret := _I.GetGType(58, "PartitionSkeleton")
	return ret
}

// udisks_partition_skeleton_new
//
// [ result ] trans: everything
//
func NewPartitionSkeleton() (result PartitionSkeleton) {
	iv, err := _I.Get(399, "PartitionSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct PartitionSkeletonClass

// Struct PartitionSkeletonPrivate
type PartitionSkeletonPrivate struct {
	P unsafe.Pointer
}

func PartitionSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(59, "PartitionSkeletonPrivate")
	return ret
}

// Interface PartitionTable
type PartitionTable struct {
	PartitionTableIfc
	P unsafe.Pointer
}
type PartitionTableIfc struct{}
type IPartitionTable interface{ P_PartitionTable() unsafe.Pointer }

func (v PartitionTable) P_PartitionTable() unsafe.Pointer { return v.P }
func PartitionTableGetType() gi.GType {
	ret := _I.GetGType(60, "PartitionTable")
	return ret
}

// udisks_partition_table_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func PartitionTableOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(401, "PartitionTable", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_partition_table_call_create_partition
//
// [ arg_offset ] trans: nothing
//
// [ arg_size ] trans: nothing
//
// [ arg_type ] trans: nothing
//
// [ arg_name ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionTableIfc) CallCreatePartition(arg_offset uint64, arg_size uint64, arg_type string, arg_name string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(402, "PartitionTable", "call_create_partition")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	c_arg_name := gi.CString(arg_name)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_offset := gi.NewUint64Argument(arg_offset)
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_offset, arg_arg_size, arg_arg_type, arg_arg_name, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
	gi.Free(c_arg_name)
}

// udisks_partition_table_call_create_partition_and_format
//
// [ arg_offset ] trans: nothing
//
// [ arg_size ] trans: nothing
//
// [ arg_type ] trans: nothing
//
// [ arg_name ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ arg_format_type ] trans: nothing
//
// [ arg_format_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *PartitionTableIfc) CallCreatePartitionAndFormat(arg_offset uint64, arg_size uint64, arg_type string, arg_name string, arg_options g.Variant, arg_format_type string, arg_format_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(403, "PartitionTable", "call_create_partition_and_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_type := gi.CString(arg_type)
	c_arg_name := gi.CString(arg_name)
	c_arg_format_type := gi.CString(arg_format_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_offset := gi.NewUint64Argument(arg_offset)
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_arg_format_type := gi.NewStringArgument(c_arg_format_type)
	arg_arg_format_options := gi.NewPointerArgument(arg_format_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_offset, arg_arg_size, arg_arg_type, arg_arg_name, arg_arg_options, arg_arg_format_type, arg_arg_format_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_type)
	gi.Free(c_arg_name)
	gi.Free(c_arg_format_type)
}

// udisks_partition_table_call_create_partition_and_format_finish
//
// [ out_created_partition ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionTableIfc) CallCreatePartitionAndFormatFinish(res g.IAsyncResult) (result bool, out_created_partition string, err error) {
	iv, err := _I.Get(404, "PartitionTable", "call_create_partition_and_format_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_created_partition := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_created_partition, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_created_partition = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_partition_table_call_create_partition_and_format_sync
//
// [ arg_offset ] trans: nothing
//
// [ arg_size ] trans: nothing
//
// [ arg_type ] trans: nothing
//
// [ arg_name ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ arg_format_type ] trans: nothing
//
// [ arg_format_options ] trans: nothing
//
// [ out_created_partition ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionTableIfc) CallCreatePartitionAndFormatSync(arg_offset uint64, arg_size uint64, arg_type string, arg_name string, arg_options g.Variant, arg_format_type string, arg_format_options g.Variant, cancellable g.ICancellable) (result bool, out_created_partition string, err error) {
	iv, err := _I.Get(405, "PartitionTable", "call_create_partition_and_format_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_type := gi.CString(arg_type)
	c_arg_name := gi.CString(arg_name)
	c_arg_format_type := gi.CString(arg_format_type)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_offset := gi.NewUint64Argument(arg_offset)
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_arg_format_type := gi.NewStringArgument(c_arg_format_type)
	arg_arg_format_options := gi.NewPointerArgument(arg_format_options.P)
	arg_out_created_partition := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_offset, arg_arg_size, arg_arg_type, arg_arg_name, arg_arg_options, arg_arg_format_type, arg_arg_format_options, arg_out_created_partition, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	gi.Free(c_arg_name)
	gi.Free(c_arg_format_type)
	err = gi.ToError(outArgs[1].Pointer())
	out_created_partition = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_partition_table_call_create_partition_finish
//
// [ out_created_partition ] trans: everything, dir: out
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionTableIfc) CallCreatePartitionFinish(res g.IAsyncResult) (result bool, out_created_partition string, err error) {
	iv, err := _I.Get(406, "PartitionTable", "call_create_partition_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_out_created_partition := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_out_created_partition, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	out_created_partition = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_partition_table_call_create_partition_sync
//
// [ arg_offset ] trans: nothing
//
// [ arg_size ] trans: nothing
//
// [ arg_type ] trans: nothing
//
// [ arg_name ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ out_created_partition ] trans: everything, dir: out
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *PartitionTableIfc) CallCreatePartitionSync(arg_offset uint64, arg_size uint64, arg_type string, arg_name string, arg_options g.Variant, cancellable g.ICancellable) (result bool, out_created_partition string, err error) {
	iv, err := _I.Get(407, "PartitionTable", "call_create_partition_sync")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_arg_type := gi.CString(arg_type)
	c_arg_name := gi.CString(arg_name)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_offset := gi.NewUint64Argument(arg_offset)
	arg_arg_size := gi.NewUint64Argument(arg_size)
	arg_arg_type := gi.NewStringArgument(c_arg_type)
	arg_arg_name := gi.NewStringArgument(c_arg_name)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_out_created_partition := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_arg_offset, arg_arg_size, arg_arg_type, arg_arg_name, arg_arg_options, arg_out_created_partition, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_type)
	gi.Free(c_arg_name)
	err = gi.ToError(outArgs[1].Pointer())
	out_created_partition = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// udisks_partition_table_complete_create_partition
//
// [ invocation ] trans: everything
//
// [ created_partition ] trans: nothing
//
func (v *PartitionTableIfc) CompleteCreatePartition(invocation g.IDBusMethodInvocation, created_partition string) {
	iv, err := _I.Get(408, "PartitionTable", "complete_create_partition")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_created_partition := gi.CString(created_partition)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_created_partition := gi.NewStringArgument(c_created_partition)
	args := []gi.Argument{arg_v, arg_invocation, arg_created_partition}
	iv.Call(args, nil, nil)
	gi.Free(c_created_partition)
}

// udisks_partition_table_complete_create_partition_and_format
//
// [ invocation ] trans: everything
//
// [ created_partition ] trans: nothing
//
func (v *PartitionTableIfc) CompleteCreatePartitionAndFormat(invocation g.IDBusMethodInvocation, created_partition string) {
	iv, err := _I.Get(409, "PartitionTable", "complete_create_partition_and_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	c_created_partition := gi.CString(created_partition)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	arg_created_partition := gi.NewStringArgument(c_created_partition)
	args := []gi.Argument{arg_v, arg_invocation, arg_created_partition}
	iv.Call(args, nil, nil)
	gi.Free(c_created_partition)
}

// ignore GType struct PartitionTableIface

// Object PartitionTableProxy
type PartitionTableProxy struct {
	PartitionTableIfc
	g.DBusProxy
}

func WrapPartitionTableProxy(p unsafe.Pointer) (r PartitionTableProxy) { r.P = p; return }

type IPartitionTableProxy interface{ P_PartitionTableProxy() unsafe.Pointer }

func (v PartitionTableProxy) P_PartitionTableProxy() unsafe.Pointer { return v.P }
func (v PartitionTableProxy) P_PartitionTable() unsafe.Pointer      { return v.P }
func PartitionTableProxyGetType() gi.GType {
	ret := _I.GetGType(61, "PartitionTableProxy")
	return ret
}

// udisks_partition_table_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionTableProxyFinish(res g.IAsyncResult) (result PartitionTableProxy, err error) {
	iv, err := _I.Get(410, "PartitionTableProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_table_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionTableProxyForBusFinish(res g.IAsyncResult) (result PartitionTableProxy, err error) {
	iv, err := _I.Get(411, "PartitionTableProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_table_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionTableProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result PartitionTableProxy, err error) {
	iv, err := _I.Get(412, "PartitionTableProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_table_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewPartitionTableProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result PartitionTableProxy, err error) {
	iv, err := _I.Get(413, "PartitionTableProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_partition_table_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func PartitionTableProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(414, "PartitionTableProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_partition_table_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func PartitionTableProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(415, "PartitionTableProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct PartitionTableProxyClass

// Struct PartitionTableProxyPrivate
type PartitionTableProxyPrivate struct {
	P unsafe.Pointer
}

func PartitionTableProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(62, "PartitionTableProxyPrivate")
	return ret
}

// Object PartitionTableSkeleton
type PartitionTableSkeleton struct {
	PartitionTableIfc
	g.DBusInterfaceSkeleton
}

func WrapPartitionTableSkeleton(p unsafe.Pointer) (r PartitionTableSkeleton) { r.P = p; return }

type IPartitionTableSkeleton interface{ P_PartitionTableSkeleton() unsafe.Pointer }

func (v PartitionTableSkeleton) P_PartitionTableSkeleton() unsafe.Pointer { return v.P }
func (v PartitionTableSkeleton) P_PartitionTable() unsafe.Pointer         { return v.P }
func PartitionTableSkeletonGetType() gi.GType {
	ret := _I.GetGType(63, "PartitionTableSkeleton")
	return ret
}

// udisks_partition_table_skeleton_new
//
// [ result ] trans: everything
//
func NewPartitionTableSkeleton() (result PartitionTableSkeleton) {
	iv, err := _I.Get(416, "PartitionTableSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct PartitionTableSkeletonClass

// Struct PartitionTableSkeletonPrivate
type PartitionTableSkeletonPrivate struct {
	P unsafe.Pointer
}

func PartitionTableSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(64, "PartitionTableSkeletonPrivate")
	return ret
}

// Struct PartitionTypeInfo
type PartitionTypeInfo struct {
	P unsafe.Pointer
}

const SizeOfStructPartitionTypeInfo = 32

func PartitionTypeInfoGetType() gi.GType {
	ret := _I.GetGType(65, "PartitionTypeInfo")
	return ret
}

// udisks_partition_type_info_free
//
func (v PartitionTypeInfo) Free() {
	iv, err := _I.Get(417, "PartitionTypeInfo", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags PartitionTypeInfoFlags
type PartitionTypeInfoFlags int

const (
	PartitionTypeInfoFlagsNone       PartitionTypeInfoFlags = 0
	PartitionTypeInfoFlagsSwap       PartitionTypeInfoFlags = 1
	PartitionTypeInfoFlagsRaid       PartitionTypeInfoFlags = 2
	PartitionTypeInfoFlagsHidden     PartitionTypeInfoFlags = 4
	PartitionTypeInfoFlagsCreateOnly PartitionTypeInfoFlags = 8
	PartitionTypeInfoFlagsSystem     PartitionTypeInfoFlags = 16
)

func PartitionTypeInfoFlagsGetType() gi.GType {
	ret := _I.GetGType(66, "PartitionTypeInfoFlags")
	return ret
}

// Interface Swapspace
type Swapspace struct {
	SwapspaceIfc
	P unsafe.Pointer
}
type SwapspaceIfc struct{}
type ISwapspace interface{ P_Swapspace() unsafe.Pointer }

func (v Swapspace) P_Swapspace() unsafe.Pointer { return v.P }
func SwapspaceGetType() gi.GType {
	ret := _I.GetGType(67, "Swapspace")
	return ret
}

// udisks_swapspace_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func SwapspaceOverrideProperties1(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(419, "Swapspace", "override_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_swapspace_call_set_label
//
// [ arg_label ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *SwapspaceIfc) CallSetLabel(arg_label string, arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(420, "Swapspace", "call_set_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg_label := gi.CString(arg_label)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_label := gi.NewStringArgument(c_arg_label)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_label, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_arg_label)
}

// udisks_swapspace_call_set_label_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SwapspaceIfc) CallSetLabelFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(421, "Swapspace", "call_set_label_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_swapspace_call_set_label_sync
//
// [ arg_label ] trans: nothing
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SwapspaceIfc) CallSetLabelSync(arg_label string, arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(422, "Swapspace", "call_set_label_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg_label := gi.CString(arg_label)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_label := gi.NewStringArgument(c_arg_label)
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_label, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg_label)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_swapspace_call_start
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *SwapspaceIfc) CallStart(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(423, "Swapspace", "call_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_swapspace_call_start_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SwapspaceIfc) CallStartFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(424, "Swapspace", "call_start_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_swapspace_call_start_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SwapspaceIfc) CallStartSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(425, "Swapspace", "call_start_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_swapspace_call_stop
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v *SwapspaceIfc) CallStop(arg_options g.Variant, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(426, "Swapspace", "call_stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// udisks_swapspace_call_stop_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SwapspaceIfc) CallStopFinish(res g.IAsyncResult) (result bool, err error) {
	iv, err := _I.Get(427, "Swapspace", "call_stop_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_swapspace_call_stop_sync
//
// [ arg_options ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SwapspaceIfc) CallStopSync(arg_options g.Variant, cancellable g.ICancellable) (result bool, err error) {
	iv, err := _I.Get(428, "Swapspace", "call_stop_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_arg_options := gi.NewPointerArgument(arg_options.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_arg_options, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// udisks_swapspace_complete_set_label
//
// [ invocation ] trans: everything
//
func (v *SwapspaceIfc) CompleteSetLabel(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(429, "Swapspace", "complete_set_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_swapspace_complete_start
//
// [ invocation ] trans: everything
//
func (v *SwapspaceIfc) CompleteStart(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(430, "Swapspace", "complete_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// udisks_swapspace_complete_stop
//
// [ invocation ] trans: everything
//
func (v *SwapspaceIfc) CompleteStop(invocation g.IDBusMethodInvocation) {
	iv, err := _I.Get(431, "Swapspace", "complete_stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if invocation != nil {
		tmp = invocation.P_DBusMethodInvocation()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_invocation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_invocation}
	iv.Call(args, nil, nil)
}

// ignore GType struct SwapspaceIface

// Object SwapspaceProxy
type SwapspaceProxy struct {
	SwapspaceIfc
	g.DBusProxy
}

func WrapSwapspaceProxy(p unsafe.Pointer) (r SwapspaceProxy) { r.P = p; return }

type ISwapspaceProxy interface{ P_SwapspaceProxy() unsafe.Pointer }

func (v SwapspaceProxy) P_SwapspaceProxy() unsafe.Pointer { return v.P }
func (v SwapspaceProxy) P_Swapspace() unsafe.Pointer      { return v.P }
func SwapspaceProxyGetType() gi.GType {
	ret := _I.GetGType(68, "SwapspaceProxy")
	return ret
}

// udisks_swapspace_proxy_new_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewSwapspaceProxyFinish(res g.IAsyncResult) (result SwapspaceProxy, err error) {
	iv, err := _I.Get(432, "SwapspaceProxy", "new_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_swapspace_proxy_new_for_bus_finish
//
// [ res ] trans: nothing
//
// [ result ] trans: everything
//
func NewSwapspaceProxyForBusFinish(res g.IAsyncResult) (result SwapspaceProxy, err error) {
	iv, err := _I.Get(433, "SwapspaceProxy", "new_for_bus_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if res != nil {
		tmp = res.P_AsyncResult()
	}
	arg_res := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_res, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_swapspace_proxy_new_for_bus_sync
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewSwapspaceProxyForBusSync(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result SwapspaceProxy, err error) {
	iv, err := _I.Get(434, "SwapspaceProxy", "new_for_bus_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_swapspace_proxy_new_sync
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewSwapspaceProxySync(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable) (result SwapspaceProxy, err error) {
	iv, err := _I.Get(435, "SwapspaceProxy", "new_sync")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	gi.Free(c_object_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// udisks_swapspace_proxy_new
//
// [ connection ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func SwapspaceProxyNew1(connection g.IDBusConnection, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(436, "SwapspaceProxy", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if connection != nil {
		tmp = connection.P_DBusConnection()
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_connection := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_connection, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// udisks_swapspace_proxy_new_for_bus
//
// [ bus_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ name ] trans: nothing
//
// [ object_path ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
func SwapspaceProxyNewForBus1(bus_type g.BusTypeEnum, flags g.DBusProxyFlags, name string, object_path string, cancellable g.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(437, "SwapspaceProxy", "new_for_bus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_object_path := gi.CString(object_path)
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_bus_type := gi.NewIntArgument(int(bus_type))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_name := gi.NewStringArgument(c_name)
	arg_object_path := gi.NewStringArgument(c_object_path)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_bus_type, arg_flags, arg_name, arg_object_path, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
	gi.Free(c_object_path)
}

// ignore GType struct SwapspaceProxyClass

// Struct SwapspaceProxyPrivate
type SwapspaceProxyPrivate struct {
	P unsafe.Pointer
}

func SwapspaceProxyPrivateGetType() gi.GType {
	ret := _I.GetGType(69, "SwapspaceProxyPrivate")
	return ret
}

// Object SwapspaceSkeleton
type SwapspaceSkeleton struct {
	SwapspaceIfc
	g.DBusInterfaceSkeleton
}

func WrapSwapspaceSkeleton(p unsafe.Pointer) (r SwapspaceSkeleton) { r.P = p; return }

type ISwapspaceSkeleton interface{ P_SwapspaceSkeleton() unsafe.Pointer }

func (v SwapspaceSkeleton) P_SwapspaceSkeleton() unsafe.Pointer { return v.P }
func (v SwapspaceSkeleton) P_Swapspace() unsafe.Pointer         { return v.P }
func SwapspaceSkeletonGetType() gi.GType {
	ret := _I.GetGType(70, "SwapspaceSkeleton")
	return ret
}

// udisks_swapspace_skeleton_new
//
// [ result ] trans: everything
//
func NewSwapspaceSkeleton() (result SwapspaceSkeleton) {
	iv, err := _I.Get(438, "SwapspaceSkeleton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct SwapspaceSkeletonClass

// Struct SwapspaceSkeletonPrivate
type SwapspaceSkeletonPrivate struct {
	P unsafe.Pointer
}

func SwapspaceSkeletonPrivateGetType() gi.GType {
	ret := _I.GetGType(71, "SwapspaceSkeletonPrivate")
	return ret
}

// udisks_block_interface_info
//
// [ result ] trans: nothing
//
func BlockInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(439, "block_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_block_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func BlockOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(440, "block_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_drive_ata_interface_info
//
// [ result ] trans: nothing
//
func DriveAtaInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(441, "drive_ata_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_drive_ata_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func DriveAtaOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(442, "drive_ata_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_drive_interface_info
//
// [ result ] trans: nothing
//
func DriveInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(443, "drive_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_drive_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func DriveOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(444, "drive_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_encrypted_interface_info
//
// [ result ] trans: nothing
//
func EncryptedInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(445, "encrypted_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_encrypted_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func EncryptedOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(446, "encrypted_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_error_quark
//
// [ result ] trans: nothing
//
func ErrorQuark() (result uint32) {
	iv, err := _I.Get(447, "error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_filesystem_interface_info
//
// [ result ] trans: nothing
//
func FilesystemInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(448, "filesystem_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_filesystem_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func FilesystemOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(449, "filesystem_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_job_interface_info
//
// [ result ] trans: nothing
//
func JobInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(450, "job_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_job_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func JobOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(451, "job_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_loop_interface_info
//
// [ result ] trans: nothing
//
func LoopInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(452, "loop_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_loop_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func LoopOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(453, "loop_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_manager_interface_info
//
// [ result ] trans: nothing
//
func ManagerInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(454, "manager_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_manager_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func ManagerOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(455, "manager_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_mdraid_interface_info
//
// [ result ] trans: nothing
//
func MdraidInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(456, "mdraid_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_mdraid_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func MdraidOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(457, "mdraid_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_partition_interface_info
//
// [ result ] trans: nothing
//
func PartitionInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(458, "partition_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_partition_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func PartitionOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(459, "partition_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_partition_table_interface_info
//
// [ result ] trans: nothing
//
func PartitionTableInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(460, "partition_table_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_partition_table_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func PartitionTableOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(461, "partition_table_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// udisks_swapspace_interface_info
//
// [ result ] trans: nothing
//
func SwapspaceInterfaceInfo() (result g.DBusInterfaceInfo) {
	iv, err := _I.Get(462, "swapspace_interface_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// udisks_swapspace_override_properties
//
// [ klass ] trans: nothing
//
// [ property_id_begin ] trans: nothing
//
// [ result ] trans: nothing
//
func SwapspaceOverrideProperties(klass g.ObjectClass, property_id_begin uint32) (result uint32) {
	iv, err := _I.Get(463, "swapspace_override_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_klass := gi.NewPointerArgument(klass.P)
	arg_property_id_begin := gi.NewUint32Argument(property_id_begin)
	args := []gi.Argument{arg_klass, arg_property_id_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// constants
const (
	ERROR_NUM_ENTRIES = 27
	MAJOR_VERSION     = 2
	MICRO_VERSION     = 1
	MINOR_VERSION     = 8
)
