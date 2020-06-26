package gtop

import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GTop")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GTop", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct glibtop
type glibtop struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop = 720

func glibtopGetType() gi.GType {
	ret := _I.GetGType(0, "glibtop")
	return ret
}

// glibtop_call_l
//
// [ command ] trans: nothing
//
// [ send_size ] trans: nothing
//
// [ send_buf ] trans: nothing
//
// [ recv_size ] trans: nothing
//
// [ recv_buf ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) CallL(command uint32, send_size uint64, send_buf unsafe.Pointer, recv_size uint64, recv_buf unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(0, "glibtop", "call_l", 360, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_command := gi.NewUint32Argument(command)
	arg_send_size := gi.NewUint64Argument(send_size)
	arg_send_buf := gi.NewPointerArgument(send_buf)
	arg_recv_size := gi.NewUint64Argument(recv_size)
	arg_recv_buf := gi.NewPointerArgument(recv_buf)
	args := []gi.Argument{arg_v, arg_command, arg_send_size, arg_send_buf, arg_recv_size, arg_recv_buf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// glibtop_call_s
//
// [ command ] trans: nothing
//
// [ send_size ] trans: nothing
//
// [ send_buf ] trans: nothing
//
// [ recv_size ] trans: nothing
//
// [ recv_buf ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) CallS(command uint32, send_size uint64, send_buf unsafe.Pointer, recv_size uint64, recv_buf unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1, "glibtop", "call_s", 360, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_command := gi.NewUint32Argument(command)
	arg_send_size := gi.NewUint64Argument(send_size)
	arg_send_buf := gi.NewPointerArgument(send_buf)
	arg_recv_size := gi.NewUint64Argument(recv_size)
	arg_recv_buf := gi.NewPointerArgument(recv_buf)
	args := []gi.Argument{arg_v, arg_command, arg_send_size, arg_send_buf, arg_recv_size, arg_recv_buf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// glibtop_close_p
//
func (v glibtop) CloseP() {
	iv, err := _I.Get(2, "glibtop", "close_p", 360, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// glibtop_close_r
//
func (v glibtop) CloseR() {
	iv, err := _I.Get(3, "glibtop", "close_r", 360, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// glibtop_close_s
//
func (v glibtop) CloseS() {
	iv, err := _I.Get(4, "glibtop", "close_s", 360, 4, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// glibtop_get_cpu_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetCpuL(buf glibtop_cpu) {
	iv, err := _I.Get(5, "glibtop", "get_cpu_l", 360, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_cpu_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetCpuS(buf glibtop_cpu) {
	iv, err := _I.Get(6, "glibtop", "get_cpu_s", 360, 6, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_fsusage_l
//
// [ buf ] trans: nothing
//
// [ mount_dir ] trans: nothing
//
func (v glibtop) GetFsusageL(buf glibtop_fsusage, mount_dir string) {
	iv, err := _I.Get(7, "glibtop", "get_fsusage_l", 360, 7, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mount_dir := gi.CString(mount_dir)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_mount_dir := gi.NewStringArgument(c_mount_dir)
	args := []gi.Argument{arg_v, arg_buf, arg_mount_dir}
	iv.Call(args, nil, nil)
	gi.Free(c_mount_dir)
}

// glibtop_get_fsusage_s
//
// [ buf ] trans: nothing
//
// [ mount_dir ] trans: nothing
//
func (v glibtop) GetFsusageS(buf glibtop_fsusage, mount_dir string) {
	iv, err := _I.Get(8, "glibtop", "get_fsusage_s", 360, 8, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mount_dir := gi.CString(mount_dir)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_mount_dir := gi.NewStringArgument(c_mount_dir)
	args := []gi.Argument{arg_v, arg_buf, arg_mount_dir}
	iv.Call(args, nil, nil)
	gi.Free(c_mount_dir)
}

// glibtop_get_loadavg_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetLoadavgL(buf glibtop_loadavg) {
	iv, err := _I.Get(9, "glibtop", "get_loadavg_l", 360, 9, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_loadavg_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetLoadavgS(buf glibtop_loadavg) {
	iv, err := _I.Get(10, "glibtop", "get_loadavg_s", 360, 10, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mem_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetMemL(buf glibtop_mem) {
	iv, err := _I.Get(11, "glibtop", "get_mem_l", 360, 11, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mem_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetMemS(buf glibtop_mem) {
	iv, err := _I.Get(12, "glibtop", "get_mem_s", 360, 12, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mountlist_l
//
// [ buf ] trans: nothing, dir: out
//
// [ all_fs ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetMountlistL(buf glibtop_mountlist, all_fs int32) (result glibtop_mountentry) {
	iv, err := _I.Get(13, "glibtop", "get_mountlist_l", 360, 13, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_all_fs := gi.NewInt32Argument(all_fs)
	args := []gi.Argument{arg_v, arg_buf, arg_all_fs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_mountlist_s
//
// [ buf ] trans: nothing, dir: out
//
// [ all_fs ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetMountlistS(buf glibtop_mountlist, all_fs int32) (result glibtop_mountentry) {
	iv, err := _I.Get(14, "glibtop", "get_mountlist_s", 360, 14, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_all_fs := gi.NewInt32Argument(all_fs)
	args := []gi.Argument{arg_v, arg_buf, arg_all_fs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_msg_limits_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetMsgLimitsL(buf glibtop_msg_limits) {
	iv, err := _I.Get(15, "glibtop", "get_msg_limits_l", 360, 15, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_msg_limits_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetMsgLimitsS(buf glibtop_msg_limits) {
	iv, err := _I.Get(16, "glibtop", "get_msg_limits_s", 360, 16, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_netlist_l
//
// [ buf ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetNetlistL(buf glibtop_netlist) (result gi.CStrArray) {
	iv, err := _I.Get(17, "glibtop", "get_netlist_l", 360, 17, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_netlist_s
//
// [ buf ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetNetlistS(buf glibtop_netlist) (result gi.CStrArray) {
	iv, err := _I.Get(18, "glibtop", "get_netlist_s", 360, 18, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_netload_l
//
// [ buf ] trans: nothing
//
// [ interface1 ] trans: nothing
//
func (v glibtop) GetNetloadL(buf glibtop_netload, interface1 string) {
	iv, err := _I.Get(19, "glibtop", "get_netload_l", 360, 19, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_interface1 := gi.CString(interface1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_interface1 := gi.NewStringArgument(c_interface1)
	args := []gi.Argument{arg_v, arg_buf, arg_interface1}
	iv.Call(args, nil, nil)
	gi.Free(c_interface1)
}

// glibtop_get_netload_s
//
// [ buf ] trans: nothing
//
// [ interface1 ] trans: nothing
//
func (v glibtop) GetNetloadS(buf glibtop_netload, interface1 string) {
	iv, err := _I.Get(20, "glibtop", "get_netload_s", 360, 20, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_interface1 := gi.CString(interface1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_interface1 := gi.NewStringArgument(c_interface1)
	args := []gi.Argument{arg_v, arg_buf, arg_interface1}
	iv.Call(args, nil, nil)
	gi.Free(c_interface1)
}

// glibtop_get_parameter_l
//
// [ parameter ] trans: nothing
//
// [ data_ptr ] trans: nothing
//
// [ data_size ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetParameterL(parameter uint32, data_ptr unsafe.Pointer, data_size uint64) (result uint64) {
	iv, err := _I.Get(21, "glibtop", "get_parameter_l", 360, 21, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parameter := gi.NewUint32Argument(parameter)
	arg_data_ptr := gi.NewPointerArgument(data_ptr)
	arg_data_size := gi.NewUint64Argument(data_size)
	args := []gi.Argument{arg_v, arg_parameter, arg_data_ptr, arg_data_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// glibtop_get_ppp_l
//
// [ buf ] trans: nothing
//
// [ device ] trans: nothing
//
func (v glibtop) GetPppL(buf glibtop_ppp, device uint16) {
	iv, err := _I.Get(22, "glibtop", "get_ppp_l", 360, 22, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_device := gi.NewUint16Argument(device)
	args := []gi.Argument{arg_v, arg_buf, arg_device}
	iv.Call(args, nil, nil)
}

// glibtop_get_ppp_s
//
// [ buf ] trans: nothing
//
// [ device ] trans: nothing
//
func (v glibtop) GetPppS(buf glibtop_ppp, device uint16) {
	iv, err := _I.Get(23, "glibtop", "get_ppp_s", 360, 23, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_device := gi.NewUint16Argument(device)
	args := []gi.Argument{arg_v, arg_buf, arg_device}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_affinity_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProcAffinityL(buf glibtop_proc_affinity, pid int32) (result uint16) {
	iv, err := _I.Get(24, "glibtop", "get_proc_affinity_l", 360, 24, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// glibtop_get_proc_affinity_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProcAffinityS(buf glibtop_proc_affinity, pid int32) (result uint16) {
	iv, err := _I.Get(25, "glibtop", "get_proc_affinity_s", 360, 25, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// glibtop_get_proc_args_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: everything
//
func (v glibtop) GetProcArgsL(buf glibtop_proc_args, pid int32, max_len uint32) (result string) {
	iv, err := _I.Get(26, "glibtop", "get_proc_args_l", 360, 26, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	arg_max_len := gi.NewUint32Argument(max_len)
	args := []gi.Argument{arg_v, arg_buf, arg_pid, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// glibtop_get_proc_args_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: everything
//
func (v glibtop) GetProcArgsS(buf glibtop_proc_args, pid int32, max_len uint32) (result string) {
	iv, err := _I.Get(27, "glibtop", "get_proc_args_s", 360, 27, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	arg_max_len := gi.NewUint32Argument(max_len)
	args := []gi.Argument{arg_v, arg_buf, arg_pid, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// glibtop_get_proc_io_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcIoL(buf glibtop_proc_io, pid int32) {
	iv, err := _I.Get(28, "glibtop", "get_proc_io_l", 360, 28, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_io_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcIoS(buf glibtop_proc_io, pid int32) {
	iv, err := _I.Get(29, "glibtop", "get_proc_io_s", 360, 29, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_kernel_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcKernelL(buf glibtop_proc_kernel, pid int32) {
	iv, err := _I.Get(30, "glibtop", "get_proc_kernel_l", 360, 30, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_kernel_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcKernelS(buf glibtop_proc_kernel, pid int32) {
	iv, err := _I.Get(31, "glibtop", "get_proc_kernel_s", 360, 31, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_map_l
//
// [ buf ] trans: nothing, dir: out
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProcMapL(buf glibtop_proc_map, pid int32) (result glibtop_map_entry) {
	iv, err := _I.Get(32, "glibtop", "get_proc_map_l", 360, 32, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_map_s
//
// [ buf ] trans: nothing, dir: out
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProcMapS(buf glibtop_proc_map, pid int32) (result glibtop_map_entry) {
	iv, err := _I.Get(33, "glibtop", "get_proc_map_s", 360, 33, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_mem_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcMemL(buf glibtop_proc_mem, pid int32) {
	iv, err := _I.Get(34, "glibtop", "get_proc_mem_l", 360, 34, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_mem_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcMemS(buf glibtop_proc_mem, pid int32) {
	iv, err := _I.Get(35, "glibtop", "get_proc_mem_s", 360, 35, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_open_files_l
//
// [ buf ] trans: nothing, dir: out
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProcOpenFilesL(buf glibtop_proc_open_files, pid int32) (result glibtop_open_files_entry) {
	iv, err := _I.Get(36, "glibtop", "get_proc_open_files_l", 360, 36, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_open_files_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: everything
//
func (v glibtop) GetProcOpenFilesS(buf glibtop_proc_open_files, pid int32) (result glibtop_open_files_entry) {
	iv, err := _I.Get(37, "glibtop", "get_proc_open_files_s", 360, 37, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_segment_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcSegmentL(buf glibtop_proc_segment, pid int32) {
	iv, err := _I.Get(38, "glibtop", "get_proc_segment_l", 360, 38, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_segment_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcSegmentS(buf glibtop_proc_segment, pid int32) {
	iv, err := _I.Get(39, "glibtop", "get_proc_segment_s", 360, 39, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_signal_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcSignalL(buf glibtop_proc_signal, pid int32) {
	iv, err := _I.Get(40, "glibtop", "get_proc_signal_l", 360, 40, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_signal_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcSignalS(buf glibtop_proc_signal, pid int32) {
	iv, err := _I.Get(41, "glibtop", "get_proc_signal_s", 360, 41, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_state_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcStateL(buf glibtop_proc_state, pid int32) {
	iv, err := _I.Get(42, "glibtop", "get_proc_state_l", 360, 42, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_state_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcStateS(buf glibtop_proc_state, pid int32) {
	iv, err := _I.Get(43, "glibtop", "get_proc_state_s", 360, 43, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_time_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcTimeL(buf glibtop_proc_time, pid int32) {
	iv, err := _I.Get(44, "glibtop", "get_proc_time_l", 360, 44, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_time_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcTimeS(buf glibtop_proc_time, pid int32) {
	iv, err := _I.Get(45, "glibtop", "get_proc_time_s", 360, 45, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_uid_l
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcUidL(buf glibtop_proc_uid, pid int32) {
	iv, err := _I.Get(46, "glibtop", "get_proc_uid_l", 360, 46, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_uid_s
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func (v glibtop) GetProcUidS(buf glibtop_proc_uid, pid int32) {
	iv, err := _I.Get(47, "glibtop", "get_proc_uid_s", 360, 47, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_v, arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proclist_l
//
// [ buf ] trans: nothing
//
// [ which ] trans: nothing
//
// [ arg ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProclistL(buf glibtop_proclist, which int64, arg int64) (result gi.Int32Array) {
	iv, err := _I.Get(48, "glibtop", "get_proclist_l", 360, 48, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_which := gi.NewInt64Argument(which)
	arg_arg := gi.NewInt64Argument(arg)
	args := []gi.Argument{arg_v, arg_buf, arg_which, arg_arg}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.Int32Array{P: ret.Pointer(), Len: int(0)}
	return
}

// glibtop_get_proclist_s
//
// [ buf ] trans: nothing
//
// [ which ] trans: nothing
//
// [ arg ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) GetProclistS(buf glibtop_proclist, which int64, arg int64) (result gi.Int32Array) {
	iv, err := _I.Get(49, "glibtop", "get_proclist_s", 360, 49, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_which := gi.NewInt64Argument(which)
	arg_arg := gi.NewInt64Argument(arg)
	args := []gi.Argument{arg_v, arg_buf, arg_which, arg_arg}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.Int32Array{P: ret.Pointer(), Len: int(0)}
	return
}

// glibtop_get_sem_limits_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetSemLimitsL(buf glibtop_sem_limits) {
	iv, err := _I.Get(50, "glibtop", "get_sem_limits_l", 360, 50, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_sem_limits_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetSemLimitsS(buf glibtop_sem_limits) {
	iv, err := _I.Get(51, "glibtop", "get_sem_limits_s", 360, 51, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_shm_limits_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetShmLimitsL(buf glibtop_shm_limits) {
	iv, err := _I.Get(52, "glibtop", "get_shm_limits_l", 360, 52, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_shm_limits_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetShmLimitsS(buf glibtop_shm_limits) {
	iv, err := _I.Get(53, "glibtop", "get_shm_limits_s", 360, 53, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_swap_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetSwapL(buf glibtop_swap) {
	iv, err := _I.Get(54, "glibtop", "get_swap_l", 360, 54, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_swap_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetSwapS(buf glibtop_swap) {
	iv, err := _I.Get(55, "glibtop", "get_swap_s", 360, 55, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_sysdeps_r
//
// [ buf ] trans: nothing
//
func (v glibtop) GetSysdepsR(buf glibtop_sysdeps) {
	iv, err := _I.Get(56, "glibtop", "get_sysdeps_r", 360, 56, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_sysinfo_s
//
// [ result ] trans: nothing
//
func (v glibtop) GetSysinfoS() (result glibtop_sysinfo) {
	iv, err := _I.Get(57, "glibtop", "get_sysinfo_s", 360, 57, gi.INFO_TYPE_STRUCT, 0)
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

// glibtop_get_uptime_l
//
// [ buf ] trans: nothing
//
func (v glibtop) GetUptimeL(buf glibtop_uptime) {
	iv, err := _I.Get(58, "glibtop", "get_uptime_l", 360, 58, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_uptime_s
//
// [ buf ] trans: nothing
//
func (v glibtop) GetUptimeS(buf glibtop_uptime) {
	iv, err := _I.Get(59, "glibtop", "get_uptime_s", 360, 59, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_v, arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_init_p
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v glibtop) InitP(features uint64, flags uint32) {
	iv, err := _I.Get(60, "glibtop", "init_p", 360, 60, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_features, arg_flags}
	iv.Call(args, nil, nil)
}

// glibtop_open_l
//
// [ program_name ] trans: nothing
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v glibtop) OpenL(program_name string, features uint64, flags uint32) {
	iv, err := _I.Get(61, "glibtop", "open_l", 360, 61, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_program_name := gi.CString(program_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_program_name := gi.NewStringArgument(c_program_name)
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_program_name, arg_features, arg_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_program_name)
}

// glibtop_open_p
//
// [ program_name ] trans: nothing
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v glibtop) OpenP(program_name string, features uint64, flags uint32) {
	iv, err := _I.Get(62, "glibtop", "open_p", 360, 62, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_program_name := gi.CString(program_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_program_name := gi.NewStringArgument(c_program_name)
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_program_name, arg_features, arg_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_program_name)
}

// glibtop_open_s
//
// [ program_name ] trans: nothing
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
func (v glibtop) OpenS(program_name string, features uint64, flags uint32) {
	iv, err := _I.Get(63, "glibtop", "open_s", 360, 63, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_program_name := gi.CString(program_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_program_name := gi.NewStringArgument(c_program_name)
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_program_name, arg_features, arg_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_program_name)
}

// glibtop_set_parameter_l
//
// [ parameter ] trans: nothing
//
// [ data_ptr ] trans: nothing
//
// [ data_size ] trans: nothing
//
func (v glibtop) SetParameterL(parameter uint32, data_ptr unsafe.Pointer, data_size uint64) {
	iv, err := _I.Get(64, "glibtop", "set_parameter_l", 360, 64, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parameter := gi.NewUint32Argument(parameter)
	arg_data_ptr := gi.NewPointerArgument(data_ptr)
	arg_data_size := gi.NewUint64Argument(data_size)
	args := []gi.Argument{arg_v, arg_parameter, arg_data_ptr, arg_data_size}
	iv.Call(args, nil, nil)
}

// glibtop_get_cpu
//
// [ buf ] trans: nothing
//
func glibtopGetCpu1(buf glibtop_cpu) {
	iv, err := _I.Get(66, "glibtop", "get_cpu", 360, 66, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_fsusage
//
// [ buf ] trans: nothing
//
// [ mount_dir ] trans: nothing
//
func glibtopGetFsusage1(buf glibtop_fsusage, mount_dir string) {
	iv, err := _I.Get(67, "glibtop", "get_fsusage", 360, 67, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mount_dir := gi.CString(mount_dir)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_mount_dir := gi.NewStringArgument(c_mount_dir)
	args := []gi.Argument{arg_buf, arg_mount_dir}
	iv.Call(args, nil, nil)
	gi.Free(c_mount_dir)
}

// glibtop_get_loadavg
//
// [ buf ] trans: nothing
//
func glibtopGetLoadavg1(buf glibtop_loadavg) {
	iv, err := _I.Get(68, "glibtop", "get_loadavg", 360, 68, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mem
//
// [ buf ] trans: nothing
//
func glibtopGetMem1(buf glibtop_mem) {
	iv, err := _I.Get(69, "glibtop", "get_mem", 360, 69, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mountlist
//
// [ buf ] trans: nothing
//
// [ all_fs ] trans: nothing
//
// [ result ] trans: everything
//
func glibtopGetMountlist1(buf glibtop_mountlist, all_fs int32) (result glibtop_mountentry) {
	iv, err := _I.Get(70, "glibtop", "get_mountlist", 360, 70, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_all_fs := gi.NewInt32Argument(all_fs)
	args := []gi.Argument{arg_buf, arg_all_fs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_msg_limits
//
// [ buf ] trans: nothing
//
func glibtopGetMsgLimits1(buf glibtop_msg_limits) {
	iv, err := _I.Get(71, "glibtop", "get_msg_limits", 360, 71, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_netlist
//
// [ buf ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopGetNetlist1(buf glibtop_netlist) (result gi.CStrArray) {
	iv, err := _I.Get(72, "glibtop", "get_netlist", 360, 72, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_netload
//
// [ buf ] trans: nothing
//
// [ interface1 ] trans: nothing
//
func glibtopGetNetload1(buf glibtop_netload, interface1 string) {
	iv, err := _I.Get(73, "glibtop", "get_netload", 360, 73, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_interface1 := gi.CString(interface1)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_interface1 := gi.NewStringArgument(c_interface1)
	args := []gi.Argument{arg_buf, arg_interface1}
	iv.Call(args, nil, nil)
	gi.Free(c_interface1)
}

// glibtop_get_ppp
//
// [ buf ] trans: nothing
//
// [ device ] trans: nothing
//
func glibtopGetPpp1(buf glibtop_ppp, device uint16) {
	iv, err := _I.Get(74, "glibtop", "get_ppp", 360, 74, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_device := gi.NewUint16Argument(device)
	args := []gi.Argument{arg_buf, arg_device}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_affinity
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopGetProcAffinity1(buf glibtop_proc_affinity, pid int32) (result uint16) {
	iv, err := _I.Get(75, "glibtop", "get_proc_affinity", 360, 75, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// glibtop_get_proc_args
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: everything
//
func glibtopGetProcArgs1(buf glibtop_proc_args, pid int32, max_len uint32) (result string) {
	iv, err := _I.Get(76, "glibtop", "get_proc_args", 360, 76, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	arg_max_len := gi.NewUint32Argument(max_len)
	args := []gi.Argument{arg_buf, arg_pid, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// glibtop_get_proc_argv
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: everything
//
func glibtopGetProcArgv1(buf glibtop_proc_args, pid int32, max_len uint32) (result gi.CStrArray) {
	iv, err := _I.Get(77, "glibtop", "get_proc_argv", 360, 77, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	arg_max_len := gi.NewUint32Argument(max_len)
	args := []gi.Argument{arg_buf, arg_pid, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_proc_io
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcIo1(buf glibtop_proc_io, pid int32) {
	iv, err := _I.Get(78, "glibtop", "get_proc_io", 360, 78, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_kernel
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcKernel1(buf glibtop_proc_kernel, pid int32) {
	iv, err := _I.Get(79, "glibtop", "get_proc_kernel", 360, 79, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_map
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopGetProcMap1(buf glibtop_proc_map, pid int32) (result glibtop_map_entry) {
	iv, err := _I.Get(80, "glibtop", "get_proc_map", 360, 80, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_mem
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcMem1(buf glibtop_proc_mem, pid int32) {
	iv, err := _I.Get(81, "glibtop", "get_proc_mem", 360, 81, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_open_files
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopGetProcOpenFiles1(buf glibtop_proc_open_files, pid int32) (result glibtop_open_files_entry) {
	iv, err := _I.Get(82, "glibtop", "get_proc_open_files", 360, 82, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_segment
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcSegment1(buf glibtop_proc_segment, pid int32) {
	iv, err := _I.Get(83, "glibtop", "get_proc_segment", 360, 83, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_signal
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcSignal1(buf glibtop_proc_signal, pid int32) {
	iv, err := _I.Get(84, "glibtop", "get_proc_signal", 360, 84, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_state
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcState1(buf glibtop_proc_state, pid int32) {
	iv, err := _I.Get(85, "glibtop", "get_proc_state", 360, 85, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_time
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcTime1(buf glibtop_proc_time, pid int32) {
	iv, err := _I.Get(86, "glibtop", "get_proc_time", 360, 86, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_uid
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func glibtopGetProcUid1(buf glibtop_proc_uid, pid int32) {
	iv, err := _I.Get(87, "glibtop", "get_proc_uid", 360, 87, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_wd
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: everything
//
func glibtopGetProcWd1(buf glibtop_proc_wd, pid int32) (result gi.CStrArray) {
	iv, err := _I.Get(88, "glibtop", "get_proc_wd", 360, 88, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_proclist
//
// [ buf ] trans: nothing
//
// [ which ] trans: nothing
//
// [ arg ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopGetProclist1(buf glibtop_proclist, which int64, arg int64) (result gi.Int32Array) {
	iv, err := _I.Get(89, "glibtop", "get_proclist", 360, 89, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_which := gi.NewInt64Argument(which)
	arg_arg := gi.NewInt64Argument(arg)
	args := []gi.Argument{arg_buf, arg_which, arg_arg}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.Int32Array{P: ret.Pointer(), Len: int(0)}
	return
}

// glibtop_get_sem_limits
//
// [ buf ] trans: nothing
//
func glibtopGetSemLimits1(buf glibtop_sem_limits) {
	iv, err := _I.Get(90, "glibtop", "get_sem_limits", 360, 90, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_shm_limits
//
// [ buf ] trans: nothing
//
func glibtopGetShmLimits1(buf glibtop_shm_limits) {
	iv, err := _I.Get(91, "glibtop", "get_shm_limits", 360, 91, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_swap
//
// [ buf ] trans: nothing
//
func glibtopGetSwap1(buf glibtop_swap) {
	iv, err := _I.Get(92, "glibtop", "get_swap", 360, 92, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_sysdeps
//
// [ buf ] trans: nothing
//
func glibtopGetSysdeps1(buf glibtop_sysdeps) {
	iv, err := _I.Get(93, "glibtop", "get_sysdeps", 360, 93, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_uptime
//
// [ buf ] trans: nothing
//
func glibtopGetUptime1(buf glibtop_uptime) {
	iv, err := _I.Get(95, "glibtop", "get_uptime", 360, 95, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_init_r
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) InitR(features uint64, flags uint32) (result glibtop) {
	iv, err := _I.Get(97, "glibtop", "init_r", 360, 97, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_features, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_init_s
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v glibtop) InitS(features uint64, flags uint32) (result glibtop) {
	iv, err := _I.Get(98, "glibtop", "init_s", 360, 98, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_v, arg_features, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_internet_addr
//
// [ host ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopInternetAddr1(host string) (result int64) {
	iv, err := _I.Get(99, "glibtop", "internet_addr", 360, 99, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_host := gi.CString(host)
	arg_host := gi.NewStringArgument(c_host)
	args := []gi.Argument{arg_host}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_host)
	result = ret.Int64()
	return
}

// glibtop_make_connection
//
// [ hostarg ] trans: nothing
//
// [ portarg ] trans: nothing
//
// [ s ] trans: nothing
//
// [ result ] trans: nothing
//
func glibtopMakeConnection1(hostarg string, portarg int32, s int32) (result int32) {
	iv, err := _I.Get(100, "glibtop", "make_connection", 360, 100, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostarg := gi.CString(hostarg)
	arg_hostarg := gi.NewStringArgument(c_hostarg)
	arg_portarg := gi.NewInt32Argument(portarg)
	arg_s := gi.NewInt32Argument(s)
	args := []gi.Argument{arg_hostarg, arg_portarg, arg_s}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostarg)
	result = ret.Int32()
	return
}

// glibtop_close
//
func GlibtopClose() {
	iv, err := _I.Get(101, "glibtop_close", "", 361, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// Struct glibtop_command
type glibtop_command struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_command = 40

func glibtop_commandGetType() gi.GType {
	ret := _I.GetGType(1, "glibtop_command")
	return ret
}

// Struct glibtop_cpu
type glibtop_cpu struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_cpu = 65624

func glibtop_cpuGetType() gi.GType {
	ret := _I.GetGType(2, "glibtop_cpu")
	return ret
}

// Struct glibtop_entry
type glibtop_entry struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_entry = 24

func glibtop_entryGetType() gi.GType {
	ret := _I.GetGType(3, "glibtop_entry")
	return ret
}

// Struct glibtop_fsusage
type glibtop_fsusage struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_fsusage = 72

func glibtop_fsusageGetType() gi.GType {
	ret := _I.GetGType(4, "glibtop_fsusage")
	return ret
}

// glibtop_get_cpu
//
// [ buf ] trans: nothing
//
func GlibtopGetCpu(buf glibtop_cpu) {
	iv, err := _I.Get(102, "glibtop_get_cpu", "", 366, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_fsusage
//
// [ buf ] trans: nothing
//
// [ mount_dir ] trans: nothing
//
func GlibtopGetFsusage(buf glibtop_fsusage, mount_dir string) {
	iv, err := _I.Get(103, "glibtop_get_fsusage", "", 367, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mount_dir := gi.CString(mount_dir)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_mount_dir := gi.NewStringArgument(c_mount_dir)
	args := []gi.Argument{arg_buf, arg_mount_dir}
	iv.Call(args, nil, nil)
	gi.Free(c_mount_dir)
}

// glibtop_get_loadavg
//
// [ buf ] trans: nothing
//
func GlibtopGetLoadavg(buf glibtop_loadavg) {
	iv, err := _I.Get(104, "glibtop_get_loadavg", "", 368, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mem
//
// [ buf ] trans: nothing
//
func GlibtopGetMem(buf glibtop_mem) {
	iv, err := _I.Get(105, "glibtop_get_mem", "", 369, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_mountlist
//
// [ buf ] trans: nothing
//
// [ all_fs ] trans: nothing
//
// [ result ] trans: everything
//
func GlibtopGetMountlist(buf glibtop_mountlist, all_fs int32) (result glibtop_mountentry) {
	iv, err := _I.Get(106, "glibtop_get_mountlist", "", 370, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_all_fs := gi.NewInt32Argument(all_fs)
	args := []gi.Argument{arg_buf, arg_all_fs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_msg_limits
//
// [ buf ] trans: nothing
//
func GlibtopGetMsgLimits(buf glibtop_msg_limits) {
	iv, err := _I.Get(107, "glibtop_get_msg_limits", "", 371, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_netlist
//
// [ buf ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopGetNetlist(buf glibtop_netlist) (result gi.CStrArray) {
	iv, err := _I.Get(108, "glibtop_get_netlist", "", 372, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_netload
//
// [ buf ] trans: nothing
//
// [ interface1 ] trans: nothing
//
func GlibtopGetNetload(buf glibtop_netload, interface1 string) {
	iv, err := _I.Get(109, "glibtop_get_netload", "", 373, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_interface1 := gi.CString(interface1)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_interface1 := gi.NewStringArgument(c_interface1)
	args := []gi.Argument{arg_buf, arg_interface1}
	iv.Call(args, nil, nil)
	gi.Free(c_interface1)
}

// glibtop_get_ppp
//
// [ buf ] trans: nothing
//
// [ device ] trans: nothing
//
func GlibtopGetPpp(buf glibtop_ppp, device uint16) {
	iv, err := _I.Get(110, "glibtop_get_ppp", "", 374, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_device := gi.NewUint16Argument(device)
	args := []gi.Argument{arg_buf, arg_device}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_affinity
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopGetProcAffinity(buf glibtop_proc_affinity, pid int32) (result uint16) {
	iv, err := _I.Get(111, "glibtop_get_proc_affinity", "", 375, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// glibtop_get_proc_args
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: everything
//
func GlibtopGetProcArgs(buf glibtop_proc_args, pid int32, max_len uint32) (result string) {
	iv, err := _I.Get(112, "glibtop_get_proc_args", "", 376, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	arg_max_len := gi.NewUint32Argument(max_len)
	args := []gi.Argument{arg_buf, arg_pid, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// glibtop_get_proc_argv
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ max_len ] trans: nothing
//
// [ result ] trans: everything
//
func GlibtopGetProcArgv(buf glibtop_proc_args, pid int32, max_len uint32) (result gi.CStrArray) {
	iv, err := _I.Get(113, "glibtop_get_proc_argv", "", 377, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	arg_max_len := gi.NewUint32Argument(max_len)
	args := []gi.Argument{arg_buf, arg_pid, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_proc_io
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcIo(buf glibtop_proc_io, pid int32) {
	iv, err := _I.Get(114, "glibtop_get_proc_io", "", 378, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_kernel
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcKernel(buf glibtop_proc_kernel, pid int32) {
	iv, err := _I.Get(115, "glibtop_get_proc_kernel", "", 379, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_map
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopGetProcMap(buf glibtop_proc_map, pid int32) (result glibtop_map_entry) {
	iv, err := _I.Get(116, "glibtop_get_proc_map", "", 380, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_mem
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcMem(buf glibtop_proc_mem, pid int32) {
	iv, err := _I.Get(117, "glibtop_get_proc_mem", "", 381, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_open_files
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopGetProcOpenFiles(buf glibtop_proc_open_files, pid int32) (result glibtop_open_files_entry) {
	iv, err := _I.Get(118, "glibtop_get_proc_open_files", "", 382, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_proc_segment
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcSegment(buf glibtop_proc_segment, pid int32) {
	iv, err := _I.Get(119, "glibtop_get_proc_segment", "", 383, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_signal
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcSignal(buf glibtop_proc_signal, pid int32) {
	iv, err := _I.Get(120, "glibtop_get_proc_signal", "", 384, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_state
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcState(buf glibtop_proc_state, pid int32) {
	iv, err := _I.Get(121, "glibtop_get_proc_state", "", 385, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_time
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcTime(buf glibtop_proc_time, pid int32) {
	iv, err := _I.Get(122, "glibtop_get_proc_time", "", 386, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_uid
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
func GlibtopGetProcUid(buf glibtop_proc_uid, pid int32) {
	iv, err := _I.Get(123, "glibtop_get_proc_uid", "", 387, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	iv.Call(args, nil, nil)
}

// glibtop_get_proc_wd
//
// [ buf ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ result ] trans: everything
//
func GlibtopGetProcWd(buf glibtop_proc_wd, pid int32) (result gi.CStrArray) {
	iv, err := _I.Get(124, "glibtop_get_proc_wd", "", 388, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_buf, arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// glibtop_get_proclist
//
// [ buf ] trans: nothing
//
// [ which ] trans: nothing
//
// [ arg ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopGetProclist(buf glibtop_proclist, which int64, arg int64) (result gi.Int32Array) {
	iv, err := _I.Get(125, "glibtop_get_proclist", "", 389, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_which := gi.NewInt64Argument(which)
	arg_arg := gi.NewInt64Argument(arg)
	args := []gi.Argument{arg_buf, arg_which, arg_arg}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.Int32Array{P: ret.Pointer(), Len: int(0)}
	return
}

// glibtop_get_sem_limits
//
// [ buf ] trans: nothing
//
func GlibtopGetSemLimits(buf glibtop_sem_limits) {
	iv, err := _I.Get(126, "glibtop_get_sem_limits", "", 390, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_shm_limits
//
// [ buf ] trans: nothing
//
func GlibtopGetShmLimits(buf glibtop_shm_limits) {
	iv, err := _I.Get(127, "glibtop_get_shm_limits", "", 391, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_swap
//
// [ buf ] trans: nothing
//
func GlibtopGetSwap(buf glibtop_swap) {
	iv, err := _I.Get(128, "glibtop_get_swap", "", 392, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_sysdeps
//
// [ buf ] trans: nothing
//
func GlibtopGetSysdeps(buf glibtop_sysdeps) {
	iv, err := _I.Get(129, "glibtop_get_sysdeps", "", 393, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_get_sysinfo
//
// [ result ] trans: nothing
//
func GlibtopGetSysinfo() (result glibtop_sysinfo) {
	iv, err := _I.Get(130, "glibtop_get_sysinfo", "", 394, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_get_uptime
//
// [ buf ] trans: nothing
//
func GlibtopGetUptime(buf glibtop_uptime) {
	iv, err := _I.Get(131, "glibtop_get_uptime", "", 395, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buf := gi.NewPointerArgument(buf.P)
	args := []gi.Argument{arg_buf}
	iv.Call(args, nil, nil)
}

// glibtop_init
//
// [ result ] trans: nothing
//
func GlibtopInit() (result glibtop) {
	iv, err := _I.Get(132, "glibtop_init", "", 396, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// glibtop_init_r
//
// [ server_ptr ] trans: everything, dir: out
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopInitR(features uint64, flags uint32) (result glibtop, server_ptr glibtop) {
	iv, err := _I.Get(133, "glibtop_init_r", "", 397, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_server_ptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_server_ptr, arg_features, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	server_ptr.P = outArgs[0].Pointer()
	result.P = ret.Pointer()
	return
}

// glibtop_init_s
//
// [ server_ptr ] trans: everything, dir: out
//
// [ features ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopInitS(features uint64, flags uint32) (result glibtop, server_ptr glibtop) {
	iv, err := _I.Get(134, "glibtop_init_s", "", 398, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_server_ptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_features := gi.NewUint64Argument(features)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_server_ptr, arg_features, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	server_ptr.P = outArgs[0].Pointer()
	result.P = ret.Pointer()
	return
}

// glibtop_internet_addr
//
// [ host ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopInternetAddr(host string) (result int64) {
	iv, err := _I.Get(135, "glibtop_internet_addr", "", 399, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_host := gi.CString(host)
	arg_host := gi.NewStringArgument(c_host)
	args := []gi.Argument{arg_host}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_host)
	result = ret.Int64()
	return
}

// Struct glibtop_loadavg
type glibtop_loadavg struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_loadavg = 56

func glibtop_loadavgGetType() gi.GType {
	ret := _I.GetGType(5, "glibtop_loadavg")
	return ret
}

// Struct glibtop_machine
type glibtop_machine struct {
	P unsafe.Pointer
}

func glibtop_machineGetType() gi.GType {
	ret := _I.GetGType(6, "glibtop_machine")
	return ret
}

// glibtop_make_connection
//
// [ hostarg ] trans: nothing
//
// [ portarg ] trans: nothing
//
// [ s ] trans: nothing
//
// [ result ] trans: nothing
//
func GlibtopMakeConnection(hostarg string, portarg int32, s int32) (result int32) {
	iv, err := _I.Get(136, "glibtop_make_connection", "", 402, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostarg := gi.CString(hostarg)
	arg_hostarg := gi.NewStringArgument(c_hostarg)
	arg_portarg := gi.NewInt32Argument(portarg)
	arg_s := gi.NewInt32Argument(s)
	args := []gi.Argument{arg_hostarg, arg_portarg, arg_s}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostarg)
	result = ret.Int32()
	return
}

// Struct glibtop_map_entry
type glibtop_map_entry struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_map_entry = 336

func glibtop_map_entryGetType() gi.GType {
	ret := _I.GetGType(7, "glibtop_map_entry")
	return ret
}

// Struct glibtop_mem
type glibtop_mem struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_mem = 72

func glibtop_memGetType() gi.GType {
	ret := _I.GetGType(8, "glibtop_mem")
	return ret
}

// Struct glibtop_mountentry
type glibtop_mountentry struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_mountentry = 248

func glibtop_mountentryGetType() gi.GType {
	ret := _I.GetGType(9, "glibtop_mountentry")
	return ret
}

// Struct glibtop_mountlist
type glibtop_mountlist struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_mountlist = 32

func glibtop_mountlistGetType() gi.GType {
	ret := _I.GetGType(10, "glibtop_mountlist")
	return ret
}

// Struct glibtop_msg_limits
type glibtop_msg_limits struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_msg_limits = 64

func glibtop_msg_limitsGetType() gi.GType {
	ret := _I.GetGType(11, "glibtop_msg_limits")
	return ret
}

// Struct glibtop_netlist
type glibtop_netlist struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_netlist = 16

func glibtop_netlistGetType() gi.GType {
	ret := _I.GetGType(12, "glibtop_netlist")
	return ret
}

// Struct glibtop_netload
type glibtop_netload struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_netload = 160

func glibtop_netloadGetType() gi.GType {
	ret := _I.GetGType(13, "glibtop_netload")
	return ret
}

// Struct glibtop_open_files_entry
type glibtop_open_files_entry struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_open_files_entry = 8

func glibtop_open_files_entryGetType() gi.GType {
	ret := _I.GetGType(14, "glibtop_open_files_entry")
	return ret
}

// Struct glibtop_ppp
type glibtop_ppp struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_ppp = 32

func glibtop_pppGetType() gi.GType {
	ret := _I.GetGType(15, "glibtop_ppp")
	return ret
}

// Struct glibtop_proc_affinity
type glibtop_proc_affinity struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_affinity = 16

func glibtop_proc_affinityGetType() gi.GType {
	ret := _I.GetGType(16, "glibtop_proc_affinity")
	return ret
}

// Struct glibtop_proc_args
type glibtop_proc_args struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_args = 16

func glibtop_proc_argsGetType() gi.GType {
	ret := _I.GetGType(17, "glibtop_proc_args")
	return ret
}

// Struct glibtop_proc_io
type glibtop_proc_io struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_io = 120

func glibtop_proc_ioGetType() gi.GType {
	ret := _I.GetGType(18, "glibtop_proc_io")
	return ret
}

// Struct glibtop_proc_kernel
type glibtop_proc_kernel struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_kernel = 112

func glibtop_proc_kernelGetType() gi.GType {
	ret := _I.GetGType(19, "glibtop_proc_kernel")
	return ret
}

// Struct glibtop_proc_map
type glibtop_proc_map struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_map = 32

func glibtop_proc_mapGetType() gi.GType {
	ret := _I.GetGType(20, "glibtop_proc_map")
	return ret
}

// Struct glibtop_proc_mem
type glibtop_proc_mem struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_mem = 56

func glibtop_proc_memGetType() gi.GType {
	ret := _I.GetGType(21, "glibtop_proc_mem")
	return ret
}

// Struct glibtop_proc_open_files
type glibtop_proc_open_files struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_open_files = 32

func glibtop_proc_open_filesGetType() gi.GType {
	ret := _I.GetGType(22, "glibtop_proc_open_files")
	return ret
}

// Struct glibtop_proc_segment
type glibtop_proc_segment struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_segment = 72

func glibtop_proc_segmentGetType() gi.GType {
	ret := _I.GetGType(23, "glibtop_proc_segment")
	return ret
}

// Struct glibtop_proc_signal
type glibtop_proc_signal struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_signal = 72

func glibtop_proc_signalGetType() gi.GType {
	ret := _I.GetGType(24, "glibtop_proc_signal")
	return ret
}

// Struct glibtop_proc_state
type glibtop_proc_state struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_state = 80

func glibtop_proc_stateGetType() gi.GType {
	ret := _I.GetGType(25, "glibtop_proc_state")
	return ret
}

// Struct glibtop_proc_time
type glibtop_proc_time struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_time = 16464

func glibtop_proc_timeGetType() gi.GType {
	ret := _I.GetGType(26, "glibtop_proc_time")
	return ret
}

// Struct glibtop_proc_uid
type glibtop_proc_uid struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_uid = 336

func glibtop_proc_uidGetType() gi.GType {
	ret := _I.GetGType(27, "glibtop_proc_uid")
	return ret
}

// Struct glibtop_proc_wd
type glibtop_proc_wd struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proc_wd = 448

func glibtop_proc_wdGetType() gi.GType {
	ret := _I.GetGType(28, "glibtop_proc_wd")
	return ret
}

// Struct glibtop_proclist
type glibtop_proclist struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_proclist = 32

func glibtop_proclistGetType() gi.GType {
	ret := _I.GetGType(29, "glibtop_proclist")
	return ret
}

// Struct glibtop_response
type glibtop_response struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_response = 65648

func glibtop_responseGetType() gi.GType {
	ret := _I.GetGType(30, "glibtop_response")
	return ret
}

// Union glibtop_response_union
type glibtop_response_union struct {
	P unsafe.Pointer
}

const SizeOfUnionglibtop_response_union = 65624

func glibtop_response_unionGetType() gi.GType {
	ret := _I.GetGType(31, "glibtop_response_union")
	return ret
}

// Struct glibtop_sem_limits
type glibtop_sem_limits struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_sem_limits = 88

func glibtop_sem_limitsGetType() gi.GType {
	ret := _I.GetGType(32, "glibtop_sem_limits")
	return ret
}

// Struct glibtop_shm_limits
type glibtop_shm_limits struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_shm_limits = 48

func glibtop_shm_limitsGetType() gi.GType {
	ret := _I.GetGType(33, "glibtop_shm_limits")
	return ret
}

// Struct glibtop_signame
type glibtop_signame struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_signame = 24

func glibtop_signameGetType() gi.GType {
	ret := _I.GetGType(34, "glibtop_signame")
	return ret
}

// Struct glibtop_swap
type glibtop_swap struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_swap = 48

func glibtop_swapGetType() gi.GType {
	ret := _I.GetGType(35, "glibtop_swap")
	return ret
}

// Struct glibtop_sysdeps
type glibtop_sysdeps struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_sysdeps = 296

func glibtop_sysdepsGetType() gi.GType {
	ret := _I.GetGType(36, "glibtop_sysdeps")
	return ret
}

// Struct glibtop_sysinfo
type glibtop_sysinfo struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_sysinfo = 24592

func glibtop_sysinfoGetType() gi.GType {
	ret := _I.GetGType(37, "glibtop_sysinfo")
	return ret
}

// Union glibtop_union
type glibtop_union struct {
	P unsafe.Pointer
}

const SizeOfUnionglibtop_union = 65624

func glibtop_unionGetType() gi.GType {
	ret := _I.GetGType(38, "glibtop_union")
	return ret
}

// Struct glibtop_uptime
type glibtop_uptime struct {
	P unsafe.Pointer
}

const SizeOfStructglibtop_uptime = 32

func glibtop_uptimeGetType() gi.GType {
	ret := _I.GetGType(39, "glibtop_uptime")
	return ret
}

// constants
const (
	AUTH_NAMESZ                            = 15
	AUTH_TIMEOUT                           = 15
	CONN_INTERNET                          = 1
	CONN_IPC                               = 2
	CONN_UNIX                              = 0
	DEFAULT_PORT                           = 21490
	DEFAUTH_NAME                           = "GNU-SECURE"
	EOT_CHR                                = 92
	EOT_STR                                = "\x04"
	FALSE                                  = 0
	GLIBTOP_CMND_CPU                       = 2
	GLIBTOP_CMND_FSUSAGE                   = 21
	GLIBTOP_CMND_LOADAVG                   = 6
	GLIBTOP_CMND_MEM                       = 3
	GLIBTOP_CMND_MOUNTLIST                 = 20
	GLIBTOP_CMND_MSG_LIMITS                = 8
	GLIBTOP_CMND_NETLIST                   = 24
	GLIBTOP_CMND_NETLOAD                   = 22
	GLIBTOP_CMND_PPP                       = 23
	GLIBTOP_CMND_PROCLIST                  = 10
	GLIBTOP_CMND_PROC_AFFINITY             = 27
	GLIBTOP_CMND_PROC_ARGS                 = 18
	GLIBTOP_CMND_PROC_IO                   = 28
	GLIBTOP_CMND_PROC_KERNEL               = 16
	GLIBTOP_CMND_PROC_MAP                  = 19
	GLIBTOP_CMND_PROC_MEM                  = 13
	GLIBTOP_CMND_PROC_OPEN_FILES           = 25
	GLIBTOP_CMND_PROC_SEGMENT              = 17
	GLIBTOP_CMND_PROC_SIGNAL               = 15
	GLIBTOP_CMND_PROC_STATE                = 11
	GLIBTOP_CMND_PROC_TIME                 = 14
	GLIBTOP_CMND_PROC_UID                  = 12
	GLIBTOP_CMND_PROC_WD                   = 26
	GLIBTOP_CMND_QUIT                      = 0
	GLIBTOP_CMND_SEM_LIMITS                = 9
	GLIBTOP_CMND_SHM_LIMITS                = 7
	GLIBTOP_CMND_SWAP                      = 4
	GLIBTOP_CMND_SYSDEPS                   = 1
	GLIBTOP_CMND_UPTIME                    = 5
	GLIBTOP_CPU_FREQUENCY                  = 5
	GLIBTOP_CPU_IDLE                       = 4
	GLIBTOP_CPU_IOWAIT                     = 12
	GLIBTOP_CPU_IRQ                        = 13
	GLIBTOP_CPU_NICE                       = 2
	GLIBTOP_CPU_SOFTIRQ                    = 14
	GLIBTOP_CPU_SYS                        = 3
	GLIBTOP_CPU_TOTAL                      = 0
	GLIBTOP_CPU_USER                       = 1
	GLIBTOP_ERROR_METHOD_ABORT             = 3
	GLIBTOP_ERROR_METHOD_IGNORE            = 0
	GLIBTOP_ERROR_METHOD_WARN              = 2
	GLIBTOP_ERROR_METHOD_WARN_ONCE         = 1
	GLIBTOP_EXCLUDE_IDLE                   = 4096
	GLIBTOP_EXCLUDE_NOTTY                  = 16384
	GLIBTOP_EXCLUDE_SYSTEM                 = 8192
	GLIBTOP_FEATURES_EXCEPT                = 8
	GLIBTOP_FEATURES_NO_SERVER             = 4
	GLIBTOP_FILE_ENTRY_FD                  = 0
	GLIBTOP_FILE_ENTRY_INETSOCKET_DST_HOST = 3
	GLIBTOP_FILE_ENTRY_INETSOCKET_DST_PORT = 4
	GLIBTOP_FILE_ENTRY_NAME                = 1
	GLIBTOP_FILE_ENTRY_TYPE                = 2
	GLIBTOP_FSUSAGE_BAVAIL                 = 2
	GLIBTOP_FSUSAGE_BFREE                  = 1
	GLIBTOP_FSUSAGE_BLOCKS                 = 0
	GLIBTOP_FSUSAGE_BLOCK_SIZE             = 5
	GLIBTOP_FSUSAGE_FFREE                  = 4
	GLIBTOP_FSUSAGE_FILES                  = 3
	GLIBTOP_FSUSAGE_READ                   = 6
	GLIBTOP_FSUSAGE_WRITE                  = 7
	GLIBTOP_INIT_NO_INIT                   = 2
	GLIBTOP_INIT_NO_OPEN                   = 1
	GLIBTOP_IPC_MSGMAP                     = 1
	GLIBTOP_IPC_MSGMAX                     = 2
	GLIBTOP_IPC_MSGMNB                     = 3
	GLIBTOP_IPC_MSGMNI                     = 4
	GLIBTOP_IPC_MSGPOOL                    = 0
	GLIBTOP_IPC_MSGSSZ                     = 5
	GLIBTOP_IPC_MSGTQL                     = 6
	GLIBTOP_IPC_SEMAEM                     = 9
	GLIBTOP_IPC_SEMMAP                     = 0
	GLIBTOP_IPC_SEMMNI                     = 1
	GLIBTOP_IPC_SEMMNS                     = 2
	GLIBTOP_IPC_SEMMNU                     = 3
	GLIBTOP_IPC_SEMMSL                     = 4
	GLIBTOP_IPC_SEMOPM                     = 5
	GLIBTOP_IPC_SEMUME                     = 6
	GLIBTOP_IPC_SEMUSZ                     = 7
	GLIBTOP_IPC_SEMVMX                     = 8
	GLIBTOP_IPC_SHMALL                     = 4
	GLIBTOP_IPC_SHMMAX                     = 0
	GLIBTOP_IPC_SHMMIN                     = 1
	GLIBTOP_IPC_SHMMNI                     = 2
	GLIBTOP_IPC_SHMSEG                     = 3
	GLIBTOP_KERN_PROC_ALL                  = 0
	GLIBTOP_KERN_PROC_MASK                 = 15
	GLIBTOP_KERN_PROC_PGRP                 = 2
	GLIBTOP_KERN_PROC_PID                  = 1
	GLIBTOP_KERN_PROC_RUID                 = 6
	GLIBTOP_KERN_PROC_SESSION              = 3
	GLIBTOP_KERN_PROC_TTY                  = 4
	GLIBTOP_KERN_PROC_UID                  = 5
	GLIBTOP_LOADAVG_LAST_PID               = 3
	GLIBTOP_LOADAVG_LOADAVG                = 0
	GLIBTOP_LOADAVG_NR_RUNNING             = 1
	GLIBTOP_LOADAVG_NR_TASKS               = 2
	GLIBTOP_MAP_ENTRY_DEVICE               = 5
	GLIBTOP_MAP_ENTRY_END                  = 1
	GLIBTOP_MAP_ENTRY_FILENAME             = 6
	GLIBTOP_MAP_ENTRY_INODE                = 4
	GLIBTOP_MAP_ENTRY_OFFSET               = 2
	GLIBTOP_MAP_ENTRY_PERM                 = 3
	GLIBTOP_MAP_ENTRY_PRIVATE_CLEAN        = 11
	GLIBTOP_MAP_ENTRY_PRIVATE_DIRTY        = 12
	GLIBTOP_MAP_ENTRY_PSS                  = 13
	GLIBTOP_MAP_ENTRY_RSS                  = 8
	GLIBTOP_MAP_ENTRY_SHARED_CLEAN         = 9
	GLIBTOP_MAP_ENTRY_SHARED_DIRTY         = 10
	GLIBTOP_MAP_ENTRY_SIZE                 = 7
	GLIBTOP_MAP_ENTRY_START                = 0
	GLIBTOP_MAP_ENTRY_SWAP                 = 14
	GLIBTOP_MAP_FILENAME_LEN               = 215
	GLIBTOP_MAP_PERM_EXECUTE               = 4
	GLIBTOP_MAP_PERM_PRIVATE               = 16
	GLIBTOP_MAP_PERM_READ                  = 1
	GLIBTOP_MAP_PERM_SHARED                = 8
	GLIBTOP_MAP_PERM_WRITE                 = 2
	GLIBTOP_MAX_CMND                       = 29
	GLIBTOP_MAX_CPU                        = 18
	GLIBTOP_MAX_FSUSAGE                    = 8
	GLIBTOP_MAX_GROUPS                     = 64
	GLIBTOP_MAX_LOADAVG                    = 4
	GLIBTOP_MAX_MAP_ENTRY                  = 15
	GLIBTOP_MAX_MEM                        = 8
	GLIBTOP_MAX_MOUNTLIST                  = 3
	GLIBTOP_MAX_MSG_LIMITS                 = 7
	GLIBTOP_MAX_NETLIST                    = 1
	GLIBTOP_MAX_NETLOAD                    = 18
	GLIBTOP_MAX_OPEN_FILE_ENTRY            = 5
	GLIBTOP_MAX_PPP                        = 3
	GLIBTOP_MAX_PROCLIST                   = 3
	GLIBTOP_MAX_PROC_AFFINITY              = 2
	GLIBTOP_MAX_PROC_ARGS                  = 1
	GLIBTOP_MAX_PROC_IO                    = 3
	GLIBTOP_MAX_PROC_KERNEL                = 9
	GLIBTOP_MAX_PROC_MAP                   = 3
	GLIBTOP_MAX_PROC_MEM                   = 6
	GLIBTOP_MAX_PROC_OPEN_FILES            = 3
	GLIBTOP_MAX_PROC_SEGMENT               = 8
	GLIBTOP_MAX_PROC_SIGNAL                = 4
	GLIBTOP_MAX_PROC_STATE                 = 9
	GLIBTOP_MAX_PROC_TIME                  = 11
	GLIBTOP_MAX_PROC_UID                   = 18
	GLIBTOP_MAX_PROC_WD                    = 3
	GLIBTOP_MAX_SEM_LIMITS                 = 10
	GLIBTOP_MAX_SHM_LIMITS                 = 5
	GLIBTOP_MAX_SWAP                       = 5
	GLIBTOP_MAX_SYSDEPS                    = 28
	GLIBTOP_MAX_SYSINFO                    = 2
	GLIBTOP_MAX_UPTIME                     = 3
	GLIBTOP_MEM_BUFFER                     = 4
	GLIBTOP_MEM_CACHED                     = 5
	GLIBTOP_MEM_FREE                       = 2
	GLIBTOP_MEM_LOCKED                     = 7
	GLIBTOP_MEM_SHARED                     = 3
	GLIBTOP_MEM_TOTAL                      = 0
	GLIBTOP_MEM_USED                       = 1
	GLIBTOP_MEM_USER                       = 6
	GLIBTOP_METHOD_DIRECT                  = 1
	GLIBTOP_METHOD_INET                    = 3
	GLIBTOP_METHOD_PIPE                    = 2
	GLIBTOP_METHOD_UNIX                    = 4
	GLIBTOP_MOUNTENTRY_LEN                 = 79
	GLIBTOP_MOUNTLIST_NUMBER               = 0
	GLIBTOP_MOUNTLIST_SIZE                 = 2
	GLIBTOP_MOUNTLIST_TOTAL                = 1
	GLIBTOP_NCPU                           = 1024
	GLIBTOP_NETLIST_NUMBER                 = 0
	GLIBTOP_NETLOAD_ADDRESS                = 3
	GLIBTOP_NETLOAD_ADDRESS6               = 14
	GLIBTOP_NETLOAD_BYTES_IN               = 7
	GLIBTOP_NETLOAD_BYTES_OUT              = 8
	GLIBTOP_NETLOAD_BYTES_TOTAL            = 9
	GLIBTOP_NETLOAD_COLLISIONS             = 13
	GLIBTOP_NETLOAD_ERRORS_IN              = 10
	GLIBTOP_NETLOAD_ERRORS_OUT             = 11
	GLIBTOP_NETLOAD_ERRORS_TOTAL           = 12
	GLIBTOP_NETLOAD_HWADDRESS              = 17
	GLIBTOP_NETLOAD_IF_FLAGS               = 0
	GLIBTOP_NETLOAD_MTU                    = 1
	GLIBTOP_NETLOAD_PACKETS_IN             = 4
	GLIBTOP_NETLOAD_PACKETS_OUT            = 5
	GLIBTOP_NETLOAD_PACKETS_TOTAL          = 6
	GLIBTOP_NETLOAD_PREFIX6                = 15
	GLIBTOP_NETLOAD_SCOPE6                 = 16
	GLIBTOP_NETLOAD_SUBNET                 = 2
	GLIBTOP_OPEN_DEST_HOST_LEN             = 46
	GLIBTOP_OPEN_FILENAME_LEN              = 215
	GLIBTOP_PARAM_COMMAND                  = 3
	GLIBTOP_PARAM_ERROR_METHOD             = 6
	GLIBTOP_PARAM_FEATURES                 = 2
	GLIBTOP_PARAM_HOST                     = 4
	GLIBTOP_PARAM_METHOD                   = 1
	GLIBTOP_PARAM_PORT                     = 5
	GLIBTOP_PARAM_REQUIRED                 = 7
	GLIBTOP_PPP_BYTES_IN                   = 1
	GLIBTOP_PPP_BYTES_OUT                  = 2
	GLIBTOP_PPP_STATE                      = 0
	GLIBTOP_PROCESS_DEAD                   = 64
	GLIBTOP_PROCESS_INTERRUPTIBLE          = 2
	GLIBTOP_PROCESS_RUNNING                = 1
	GLIBTOP_PROCESS_STOPPED                = 16
	GLIBTOP_PROCESS_SWAPPING               = 32
	GLIBTOP_PROCESS_UNINTERRUPTIBLE        = 4
	GLIBTOP_PROCESS_ZOMBIE                 = 8
	GLIBTOP_PROCLIST_NUMBER                = 0
	GLIBTOP_PROCLIST_SIZE                  = 2
	GLIBTOP_PROCLIST_TOTAL                 = 1
	GLIBTOP_PROC_AFFINITY_ALL              = 1
	GLIBTOP_PROC_AFFINITY_NUMBER           = 0
	GLIBTOP_PROC_ARGS_SIZE                 = 0
	GLIBTOP_PROC_IO_DISK_RBYTES            = 2
	GLIBTOP_PROC_IO_DISK_RCHAR             = 0
	GLIBTOP_PROC_IO_DISK_WBYTES            = 3
	GLIBTOP_PROC_IO_DISK_WCHAR             = 1
	GLIBTOP_PROC_KERNEL_CMAJ_FLT           = 4
	GLIBTOP_PROC_KERNEL_CMIN_FLT           = 3
	GLIBTOP_PROC_KERNEL_KSTK_EIP           = 6
	GLIBTOP_PROC_KERNEL_KSTK_ESP           = 5
	GLIBTOP_PROC_KERNEL_K_FLAGS            = 0
	GLIBTOP_PROC_KERNEL_MAJ_FLT            = 2
	GLIBTOP_PROC_KERNEL_MIN_FLT            = 1
	GLIBTOP_PROC_KERNEL_NWCHAN             = 7
	GLIBTOP_PROC_KERNEL_WCHAN              = 8
	GLIBTOP_PROC_MAP_NUMBER                = 0
	GLIBTOP_PROC_MAP_SIZE                  = 2
	GLIBTOP_PROC_MAP_TOTAL                 = 1
	GLIBTOP_PROC_MEM_RESIDENT              = 2
	GLIBTOP_PROC_MEM_RSS                   = 4
	GLIBTOP_PROC_MEM_RSS_RLIM              = 5
	GLIBTOP_PROC_MEM_SHARE                 = 3
	GLIBTOP_PROC_MEM_SIZE                  = 0
	GLIBTOP_PROC_MEM_VSIZE                 = 1
	GLIBTOP_PROC_OPEN_FILES_NUMBER         = 0
	GLIBTOP_PROC_OPEN_FILES_SIZE           = 2
	GLIBTOP_PROC_OPEN_FILES_TOTAL          = 1
	GLIBTOP_PROC_SEGMENT_DATA_RSS          = 2
	GLIBTOP_PROC_SEGMENT_DIRTY_SIZE        = 4
	GLIBTOP_PROC_SEGMENT_END_CODE          = 6
	GLIBTOP_PROC_SEGMENT_SHLIB_RSS         = 1
	GLIBTOP_PROC_SEGMENT_STACK_RSS         = 3
	GLIBTOP_PROC_SEGMENT_START_CODE        = 5
	GLIBTOP_PROC_SEGMENT_START_STACK       = 7
	GLIBTOP_PROC_SEGMENT_TEXT_RSS          = 0
	GLIBTOP_PROC_SIGNAL_BLOCKED            = 1
	GLIBTOP_PROC_SIGNAL_SIGCATCH           = 3
	GLIBTOP_PROC_SIGNAL_SIGIGNORE          = 2
	GLIBTOP_PROC_SIGNAL_SIGNAL             = 0
	GLIBTOP_PROC_STATE_CMD                 = 0
	GLIBTOP_PROC_STATE_GID                 = 3
	GLIBTOP_PROC_STATE_HAS_CPU             = 6
	GLIBTOP_PROC_STATE_LAST_PROCESSOR      = 8
	GLIBTOP_PROC_STATE_PROCESSOR           = 7
	GLIBTOP_PROC_STATE_RGID                = 5
	GLIBTOP_PROC_STATE_RUID                = 4
	GLIBTOP_PROC_STATE_STATE               = 1
	GLIBTOP_PROC_STATE_UID                 = 2
	GLIBTOP_PROC_TIME_CSTIME               = 5
	GLIBTOP_PROC_TIME_CUTIME               = 4
	GLIBTOP_PROC_TIME_FREQUENCY            = 8
	GLIBTOP_PROC_TIME_IT_REAL_VALUE        = 7
	GLIBTOP_PROC_TIME_RTIME                = 1
	GLIBTOP_PROC_TIME_START_TIME           = 0
	GLIBTOP_PROC_TIME_STIME                = 3
	GLIBTOP_PROC_TIME_TIMEOUT              = 6
	GLIBTOP_PROC_TIME_UTIME                = 2
	GLIBTOP_PROC_TIME_XCPU_STIME           = 10
	GLIBTOP_PROC_TIME_XCPU_UTIME           = 9
	GLIBTOP_PROC_UID_EGID                  = 3
	GLIBTOP_PROC_UID_EUID                  = 1
	GLIBTOP_PROC_UID_FSGID                 = 7
	GLIBTOP_PROC_UID_FSUID                 = 6
	GLIBTOP_PROC_UID_GID                   = 2
	GLIBTOP_PROC_UID_GROUPS                = 17
	GLIBTOP_PROC_UID_NGROUPS               = 16
	GLIBTOP_PROC_UID_NICE                  = 15
	GLIBTOP_PROC_UID_PGRP                  = 10
	GLIBTOP_PROC_UID_PID                   = 8
	GLIBTOP_PROC_UID_PPID                  = 9
	GLIBTOP_PROC_UID_PRIORITY              = 14
	GLIBTOP_PROC_UID_SESSION               = 11
	GLIBTOP_PROC_UID_SGID                  = 5
	GLIBTOP_PROC_UID_SUID                  = 4
	GLIBTOP_PROC_UID_TPGID                 = 13
	GLIBTOP_PROC_UID_TTY                   = 12
	GLIBTOP_PROC_UID_UID                   = 0
	GLIBTOP_PROC_WD_EXE                    = 2
	GLIBTOP_PROC_WD_EXE_LEN                = 215
	GLIBTOP_PROC_WD_NUMBER                 = 0
	GLIBTOP_PROC_WD_ROOT                   = 1
	GLIBTOP_PROC_WD_ROOT_LEN               = 215
	GLIBTOP_SWAP_FREE                      = 2
	GLIBTOP_SWAP_PAGEIN                    = 3
	GLIBTOP_SWAP_PAGEOUT                   = 4
	GLIBTOP_SWAP_TOTAL                     = 0
	GLIBTOP_SWAP_USED                      = 1
	GLIBTOP_SYSDEPS_ALL                    = 0
	GLIBTOP_SYSDEPS_CPU                    = 1
	GLIBTOP_SYSDEPS_FEATURES               = 0
	GLIBTOP_SYSDEPS_FSUSAGE                = 20
	GLIBTOP_SYSDEPS_LOADAVG                = 5
	GLIBTOP_SYSDEPS_MEM                    = 2
	GLIBTOP_SYSDEPS_MOUNTLIST              = 19
	GLIBTOP_SYSDEPS_MSG_LIMITS             = 7
	GLIBTOP_SYSDEPS_NETLIST                = 23
	GLIBTOP_SYSDEPS_NETLOAD                = 21
	GLIBTOP_SYSDEPS_PPP                    = 22
	GLIBTOP_SYSDEPS_PROCLIST               = 9
	GLIBTOP_SYSDEPS_PROC_AFFINITY          = 26
	GLIBTOP_SYSDEPS_PROC_ARGS              = 17
	GLIBTOP_SYSDEPS_PROC_IO                = 27
	GLIBTOP_SYSDEPS_PROC_KERNEL            = 15
	GLIBTOP_SYSDEPS_PROC_MAP               = 18
	GLIBTOP_SYSDEPS_PROC_MEM               = 12
	GLIBTOP_SYSDEPS_PROC_OPEN_FILES        = 24
	GLIBTOP_SYSDEPS_PROC_SEGMENT           = 16
	GLIBTOP_SYSDEPS_PROC_SIGNAL            = 14
	GLIBTOP_SYSDEPS_PROC_STATE             = 10
	GLIBTOP_SYSDEPS_PROC_TIME              = 13
	GLIBTOP_SYSDEPS_PROC_UID               = 11
	GLIBTOP_SYSDEPS_PROC_WD                = 25
	GLIBTOP_SYSDEPS_SEM_LIMITS             = 8
	GLIBTOP_SYSDEPS_SHM_LIMITS             = 6
	GLIBTOP_SYSDEPS_SWAP                   = 3
	GLIBTOP_SYSDEPS_UPTIME                 = 4
	GLIBTOP_SYSINFO_CPUINFO                = 1
	GLIBTOP_SYSINFO_NCPU                   = 0
	GLIBTOP_UPTIME_BOOT_TIME               = 2
	GLIBTOP_UPTIME_IDLETIME                = 1
	GLIBTOP_UPTIME_UPTIME                  = 0
	GLIBTOP_XCPU_FLAGS                     = 11
	GLIBTOP_XCPU_IDLE                      = 10
	GLIBTOP_XCPU_IOWAIT                    = 15
	GLIBTOP_XCPU_IRQ                       = 16
	GLIBTOP_XCPU_NICE                      = 8
	GLIBTOP_XCPU_SOFTIRQ                   = 17
	GLIBTOP_XCPU_SYS                       = 9
	GLIBTOP_XCPU_TOTAL                     = 6
	GLIBTOP_XCPU_USER                      = 7
	HOSTNAMSZ                              = 255
	LIBGTOP_MAJOR_VERSION                  = 2
	LIBGTOP_MICRO_VERSION                  = 0
	LIBGTOP_MINOR_VERSION                  = 38
	MCOOKIE_NAME                           = "MAGIC-1"
	MCOOKIE_SCREEN                         = "42980"
	MCOOKIE_X_NAME                         = "MIT-MAGIC-COOKIE-1"
	PATCHLEVEL                             = 2
	REPLYSIZ                               = 300
	TABLE_SIZE                             = 101
	TRUE                                   = 1
)
