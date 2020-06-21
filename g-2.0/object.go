package g

import "unsafe"

// Connect is a wrapper around g_signal_connect_closure().  f must be
// a function with a signaure matching the callback signature for
// detailedSignal.  userData must either 0 or 1 elements which can
// be optionally passed to f.  If f takes less arguments than it is
// passed from the GLib runtime, the extra arguments are ignored.
//
// Arguments for f must be a matching Go equivalent type for the
// C callback, or an interface type which the value may be packed in.
// If the type is not suitable, a runtime panic will occur when the
// signal is emitted.
func (v Object) Connect(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(false, detailedSignal, f)
}

// ConnectAfter is a wrapper around g_signal_connect_closure().  f must be
// a function with a signaure matching the callback signature for
// detailedSignal.  userData must either 0 or 1 elements which can
// be optionally passed to f.  If f takes less arguments than it is
// passed from the GLib runtime, the extra arguments are ignored.
//
// Arguments for f must be a matching Go equivalent type for the
// C callback, or an interface type which the value may be packed in.
// If the type is not suitable, a runtime panic will occur when the
// signal is emitted.
//
// The difference between Connect and ConnectAfter is that the latter
// will be invoked after the default handler, not before.
func (v Object) ConnectAfter(detailedSignal string, f interface{}) SignalHandle {
	return v.connectClosure(true, detailedSignal, f)
}

//func Unref(p unsafe.Pointer) {
//	WrapObject(p).Unref()
//}
//
//func Ref(p unsafe.Pointer) {
//	WrapObject(p).Ref()
//}
//
//func RefSink(p unsafe.Pointer) {
//	WrapObject(p).RefSink()
//}

func WrapObject(p unsafe.Pointer) Object {
    return Object{P: p}
}
