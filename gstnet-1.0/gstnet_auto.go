package gstnet

/*
#cgo pkg-config: gstreamer-net-1.0
#include <gst/net/net.h>
extern void myGstNetPtpStatisticsCallback(guint8 domain, GstStructure* stats, gpointer user_data);
static void* getPointer_myGstNetPtpStatisticsCallback() {
return (void*)(myGstNetPtpStatisticsCallback);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gst-1.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GstNet")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GstNet", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct NetAddressMeta
type NetAddressMeta struct {
	P unsafe.Pointer
}

const SizeOfStructNetAddressMeta = 24

func NetAddressMetaGetType() gi.GType {
	ret := _I.GetGType(0, "NetAddressMeta")
	return ret
}

// Object NetClientClock
type NetClientClock struct {
	gst.SystemClock
}

func WrapNetClientClock(p unsafe.Pointer) (r NetClientClock) { r.P = p; return }

type INetClientClock interface{ P_NetClientClock() unsafe.Pointer }

func (v NetClientClock) P_NetClientClock() unsafe.Pointer { return v.P }
func NetClientClockGetType() gi.GType {
	ret := _I.GetGType(1, "NetClientClock")
	return ret
}

// gst_net_client_clock_new
//
// [ name ] trans: nothing
//
// [ remote_address ] trans: nothing
//
// [ remote_port ] trans: nothing
//
// [ base_time ] trans: nothing
//
// [ result ] trans: everything
//
func NewNetClientClock(name string, remote_address string, remote_port int32, base_time uint64) (result NetClientClock) {
	iv, err := _I.Get(1, "NetClientClock", "new", 2, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_remote_address := gi.CString(remote_address)
	arg_name := gi.NewStringArgument(c_name)
	arg_remote_address := gi.NewStringArgument(c_remote_address)
	arg_remote_port := gi.NewInt32Argument(remote_port)
	arg_base_time := gi.NewUint64Argument(base_time)
	args := []gi.Argument{arg_name, arg_remote_address, arg_remote_port, arg_base_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_remote_address)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NetClientClockClass

// Struct NetClientClockPrivate
type NetClientClockPrivate struct {
	P unsafe.Pointer
}

func NetClientClockPrivateGetType() gi.GType {
	ret := _I.GetGType(2, "NetClientClockPrivate")
	return ret
}

// Struct NetControlMessageMeta
type NetControlMessageMeta struct {
	P unsafe.Pointer
}

const SizeOfStructNetControlMessageMeta = 24

func NetControlMessageMetaGetType() gi.GType {
	ret := _I.GetGType(3, "NetControlMessageMeta")
	return ret
}

// Struct NetTimePacket
type NetTimePacket struct {
	P unsafe.Pointer
}

const SizeOfStructNetTimePacket = 16

func NetTimePacketGetType() gi.GType {
	ret := _I.GetGType(4, "NetTimePacket")
	return ret
}

// gst_net_time_packet_new
//
// [ buffer ] trans: nothing
//
// [ result ] trans: everything
//
func NewNetTimePacket(buffer gi.Uint8Array) (result NetTimePacket) {
	iv, err := _I.Get(3, "NetTimePacket", "new", 6, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_net_time_packet_copy
//
// [ result ] trans: everything
//
func (v NetTimePacket) Copy() (result NetTimePacket) {
	iv, err := _I.Get(4, "NetTimePacket", "copy", 6, 1, gi.INFO_TYPE_STRUCT, 0)
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

// gst_net_time_packet_free
//
func (v NetTimePacket) Free() {
	iv, err := _I.Get(5, "NetTimePacket", "free", 6, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gst_net_time_packet_send
//
// [ socket ] trans: nothing
//
// [ dest_address ] trans: nothing
//
// [ result ] trans: nothing
//
func (v NetTimePacket) Send(socket g.ISocket, dest_address g.ISocketAddress) (result bool, err error) {
	iv, err := _I.Get(6, "NetTimePacket", "send", 6, 3, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if socket != nil {
		tmp = socket.P_Socket()
	}
	var tmp1 unsafe.Pointer
	if dest_address != nil {
		tmp1 = dest_address.P_SocketAddress()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_socket := gi.NewPointerArgument(tmp)
	arg_dest_address := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_socket, arg_dest_address, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gst_net_time_packet_serialize
//
// [ result ] trans: nothing
//
func (v NetTimePacket) Serialize() (result uint8) {
	iv, err := _I.Get(7, "NetTimePacket", "serialize", 6, 4, gi.INFO_TYPE_STRUCT, 0)
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

// gst_net_time_packet_receive
//
// [ socket ] trans: nothing
//
// [ src_address ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func NetTimePacketReceive1(socket g.ISocket) (result NetTimePacket, src_address g.SocketAddress, err error) {
	iv, err := _I.Get(8, "NetTimePacket", "receive", 6, 5, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if socket != nil {
		tmp = socket.P_Socket()
	}
	arg_socket := gi.NewPointerArgument(tmp)
	arg_src_address := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_socket, arg_src_address, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	src_address.P = outArgs[0].Pointer()
	result.P = ret.Pointer()
	return
}

// Object NetTimeProvider
type NetTimeProvider struct {
	g.InitableIfc
	gst.Object
}

func WrapNetTimeProvider(p unsafe.Pointer) (r NetTimeProvider) { r.P = p; return }

type INetTimeProvider interface{ P_NetTimeProvider() unsafe.Pointer }

func (v NetTimeProvider) P_NetTimeProvider() unsafe.Pointer { return v.P }
func (v NetTimeProvider) P_Initable() unsafe.Pointer        { return v.P }
func NetTimeProviderGetType() gi.GType {
	ret := _I.GetGType(5, "NetTimeProvider")
	return ret
}

// gst_net_time_provider_new
//
// [ clock ] trans: nothing
//
// [ address ] trans: nothing
//
// [ port ] trans: nothing
//
// [ result ] trans: everything
//
func NewNetTimeProvider(clock gst.IClock, address string, port int32) (result NetTimeProvider) {
	iv, err := _I.Get(9, "NetTimeProvider", "new", 7, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if clock != nil {
		tmp = clock.P_Clock()
	}
	c_address := gi.CString(address)
	arg_clock := gi.NewPointerArgument(tmp)
	arg_address := gi.NewStringArgument(c_address)
	arg_port := gi.NewInt32Argument(port)
	args := []gi.Argument{arg_clock, arg_address, arg_port}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_address)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NetTimeProviderClass

// Struct NetTimeProviderPrivate
type NetTimeProviderPrivate struct {
	P unsafe.Pointer
}

func NetTimeProviderPrivateGetType() gi.GType {
	ret := _I.GetGType(6, "NetTimeProviderPrivate")
	return ret
}

// Object NtpClock
type NtpClock struct {
	NetClientClock
}

func WrapNtpClock(p unsafe.Pointer) (r NtpClock) { r.P = p; return }

type INtpClock interface{ P_NtpClock() unsafe.Pointer }

func (v NtpClock) P_NtpClock() unsafe.Pointer { return v.P }
func NtpClockGetType() gi.GType {
	ret := _I.GetGType(7, "NtpClock")
	return ret
}

// gst_ntp_clock_new
//
// [ name ] trans: nothing
//
// [ remote_address ] trans: nothing
//
// [ remote_port ] trans: nothing
//
// [ base_time ] trans: nothing
//
// [ result ] trans: everything
//
func NewNtpClock(name string, remote_address string, remote_port int32, base_time uint64) (result NtpClock) {
	iv, err := _I.Get(10, "NtpClock", "new", 10, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_remote_address := gi.CString(remote_address)
	arg_name := gi.NewStringArgument(c_name)
	arg_remote_address := gi.NewStringArgument(c_remote_address)
	arg_remote_port := gi.NewInt32Argument(remote_port)
	arg_base_time := gi.NewUint64Argument(base_time)
	args := []gi.Argument{arg_name, arg_remote_address, arg_remote_port, arg_base_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_remote_address)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NtpClockClass

// Object PtpClock
type PtpClock struct {
	gst.SystemClock
}

func WrapPtpClock(p unsafe.Pointer) (r PtpClock) { r.P = p; return }

type IPtpClock interface{ P_PtpClock() unsafe.Pointer }

func (v PtpClock) P_PtpClock() unsafe.Pointer { return v.P }
func PtpClockGetType() gi.GType {
	ret := _I.GetGType(8, "PtpClock")
	return ret
}

// gst_ptp_clock_new
//
// [ name ] trans: nothing
//
// [ domain ] trans: nothing
//
// [ result ] trans: everything
//
func NewPtpClock(name string, domain uint32) (result PtpClock) {
	iv, err := _I.Get(11, "PtpClock", "new", 17, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_domain := gi.NewUint32Argument(domain)
	args := []gi.Argument{arg_name, arg_domain}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// ignore GType struct PtpClockClass

// Struct PtpClockPrivate
type PtpClockPrivate struct {
	P unsafe.Pointer
}

func PtpClockPrivateGetType() gi.GType {
	ret := _I.GetGType(9, "PtpClockPrivate")
	return ret
}

type PtpStatisticsCallbackStruct struct {
	F_domain uint8
	F_stats  gst.Structure
}

func GetPointer_myPtpStatisticsCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGstNetPtpStatisticsCallback())
}

//export myGstNetPtpStatisticsCallback
func myGstNetPtpStatisticsCallback(domain C.guint8, stats *C.GstStructure, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &PtpStatisticsCallbackStruct{
		F_domain: uint8(domain),
		F_stats:  gst.Structure{P: unsafe.Pointer(stats)},
	}
	fn(args)
}

// gst_buffer_add_net_address_meta
//
// [ buffer ] trans: nothing
//
// [ addr ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferAddNetAddressMeta(buffer gst.Buffer, addr g.ISocketAddress) (result NetAddressMeta) {
	iv, err := _I.Get(12, "buffer_add_net_address_meta", "", 21, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if addr != nil {
		tmp = addr.P_SocketAddress()
	}
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_addr := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_buffer, arg_addr}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_add_net_control_message_meta
//
// [ buffer ] trans: nothing
//
// [ message ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferAddNetControlMessageMeta(buffer gst.Buffer, message g.ISocketControlMessage) (result NetControlMessageMeta) {
	iv, err := _I.Get(13, "buffer_add_net_control_message_meta", "", 22, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if message != nil {
		tmp = message.P_SocketControlMessage()
	}
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_message := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_buffer, arg_message}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_buffer_get_net_address_meta
//
// [ buffer ] trans: nothing
//
// [ result ] trans: nothing
//
func BufferGetNetAddressMeta(buffer gst.Buffer) (result NetAddressMeta) {
	iv, err := _I.Get(14, "buffer_get_net_address_meta", "", 23, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_net_address_meta_api_get_type
//
// [ result ] trans: nothing
//
func NetAddressMetaApiGetType() (result gi.GType) {
	iv, err := _I.Get(15, "net_address_meta_api_get_type", "", 24, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_net_address_meta_get_info
//
// [ result ] trans: nothing
//
func NetAddressMetaGetInfo() (result gst.MetaInfo) {
	iv, err := _I.Get(16, "net_address_meta_get_info", "", 25, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_net_control_message_meta_api_get_type
//
// [ result ] trans: nothing
//
func NetControlMessageMetaApiGetType() (result gi.GType) {
	iv, err := _I.Get(17, "net_control_message_meta_api_get_type", "", 26, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// gst_net_control_message_meta_get_info
//
// [ result ] trans: nothing
//
func NetControlMessageMetaGetInfo() (result gst.MetaInfo) {
	iv, err := _I.Get(18, "net_control_message_meta_get_info", "", 27, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gst_net_time_packet_receive
//
// [ socket ] trans: nothing
//
// [ src_address ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func NetTimePacketReceive(socket g.ISocket) (result NetTimePacket, src_address g.SocketAddress, err error) {
	iv, err := _I.Get(19, "net_time_packet_receive", "", 28, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if socket != nil {
		tmp = socket.P_Socket()
	}
	arg_socket := gi.NewPointerArgument(tmp)
	arg_src_address := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_socket, arg_src_address, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	src_address.P = outArgs[0].Pointer()
	result.P = ret.Pointer()
	return
}

// gst_ptp_deinit
//
func PtpDeinit() {
	iv, err := _I.Get(20, "ptp_deinit", "", 29, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gst_ptp_init
//
// [ clock_id ] trans: nothing
//
// [ interfaces ] trans: nothing
//
// [ result ] trans: nothing
//
func PtpInit(clock_id uint64, interfaces gi.CStrArray) (result bool) {
	iv, err := _I.Get(21, "ptp_init", "", 30, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_clock_id := gi.NewUint64Argument(clock_id)
	arg_interfaces := gi.NewPointerArgument(interfaces.P)
	args := []gi.Argument{arg_clock_id, arg_interfaces}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gst_ptp_is_initialized
//
// [ result ] trans: nothing
//
func PtpIsInitialized() (result bool) {
	iv, err := _I.Get(22, "ptp_is_initialized", "", 31, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gst_ptp_is_supported
//
// [ result ] trans: nothing
//
func PtpIsSupported() (result bool) {
	iv, err := _I.Get(23, "ptp_is_supported", "", 32, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gst_ptp_statistics_callback_add
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ destroy_data ] trans: nothing
//
// [ result ] trans: nothing
//
func PtpStatisticsCallbackAdd(callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, destroy_data int /*TODO_TYPE CALLBACK*/) (result uint64) {
	iv, err := _I.Get(24, "ptp_statistics_callback_add", "", 33, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myPtpStatisticsCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy_data := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_callback, arg_user_data, arg_destroy_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gst_ptp_statistics_callback_remove
//
// [ id ] trans: nothing
//
func PtpStatisticsCallbackRemove(id uint64) {
	iv, err := _I.Get(25, "ptp_statistics_callback_remove", "", 34, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_id := gi.NewUint64Argument(id)
	args := []gi.Argument{arg_id}
	iv.Call(args, nil, nil)
}

// constants
const (
	NET_TIME_PACKET_SIZE                      = 16
	PTP_CLOCK_ID_NONE                         = 18446744073709551615
	PTP_STATISTICS_BEST_MASTER_CLOCK_SELECTED = "GstPtpStatisticsBestMasterClockSelected"
	PTP_STATISTICS_NEW_DOMAIN_FOUND           = "GstPtpStatisticsNewDomainFound"
	PTP_STATISTICS_PATH_DELAY_MEASURED        = "GstPtpStatisticsPathDelayMeasured"
	PTP_STATISTICS_TIME_UPDATED               = "GstPtpStatisticsTimeUpdated"
)
