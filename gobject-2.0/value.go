package gobject

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>

static GValue *
_g_value_alloc()
{
	return (g_new0(GValue, 1));
}

static GValue *
_g_value_init(GType g_type)
{
	GValue          *value;

	value = g_new0(GValue, 1);
	return (g_value_init(value, g_type));
}

static gboolean
_g_is_value(GValue *val)
{
	return (G_IS_VALUE(val));
}

static GType
_g_value_type(GValue *val)
{
	return (G_VALUE_TYPE(val));
}

static GType
_g_value_fundamental(GType type)
{
	return (G_TYPE_FUNDAMENTAL(type));
}

static GObjectClass *
_g_object_get_class (GObject *object)
{
	return (G_OBJECT_GET_CLASS(object));
}
*/
import "C"

import (
	"errors"
	"github.com/electricface/go-gir/glib-2.0"
	"unsafe"

	"github.com/electricface/go-gir/util"
	"github.com/electricface/go-gir3/gi-lite"
)

/*
 * GValue
 */

// ValueAlloc allocates a Value and sets a runtime finalizer to call
// g_value_unset() on the underlying GValue after leaving scope.
// ValueAlloc() returns a non-nil error if the allocation failed.

func ValueNew() Value {
	return Value{unsafe.Pointer(C._g_value_alloc())}
}

func (v Value) Free() {
	C.g_free(C.gpointer(v.P))
}

func (v Value) native() *C.GValue {
	return (*C.GValue)(v.P)
}

// ValueInit is a wrapper around g_value_init() and allocates and
// initializes a new Value with the Type t.
//func (v Value) Init(gType Type) {
//	C.g_value_init(v.native(), C.GType(gType))
//}

// Type is a wrapper around the G_VALUE_HOLDS_GTYPE() macro and
// the g_value_get_gtype() function.  GetType() returns TYPE_INVALID if v
// does not hold a Type, or otherwise returns the Type of v.
func (v Value) Type() (actual gi.GType, fundamental gi.GType, err error) {
	if !util.Int2Bool(int(C._g_is_value(v.native()))) {
		return actual, fundamental, errors.New("invalid GValue")
	}
	cActual := C._g_value_type(v.native())
	cFundamental := C._g_value_fundamental(cActual)
	return gi.GType(cActual), gi.GType(cFundamental), nil
}

//var gvalueGetters = struct {
//	m map[gi.GType]GValueGetter
//	sync.Mutex
//}{
//	m: make(map[gi.GType]GValueGetter),
//}
//
//type GValueGetter func(unsafe.Pointer) (interface{}, error)

//func registerGValueGetter(typ gi.GType, getter GValueGetter) {
//	gvalueGetters.Lock()
//	gvalueGetters.m[typ] = getter
//	gvalueGetters.Unlock()
//}

func (v Value) Get() (interface{}, error) {
	actualType, fundamentalType, err := v.Type()
	if err != nil {
		return nil, err
	}
	//fmt.Println("actual Type:", actualType)
	//fmt.Println("fund Type:", fundamentalType)

	val, err := v.get(actualType)
	if err == nil {
		// good
		return val, nil
	} else if err == errTypeUnknown {
		// fallback to fundamental type
		return v.get(fundamentalType)
	}
	return nil, err
}

var errTypeUnknown = errors.New("unknown type")

func (v Value) get(typ gi.GType) (ret interface{}, err error) {
	switch typ {
	case TYPE_INVALID:
		err = errors.New("invalid type")
	case TYPE_NONE:
	case TYPE_INTERFACE:
		err = errors.New("interface conversion not yet implemented")
	case TYPE_CHAR:
		ret = v.GetSchar()
	case TYPE_UCHAR:
		ret = v.GetUchar()
	case TYPE_BOOLEAN:
		ret = v.GetBoolean()
	case TYPE_INT:
		ret = v.GetInt()

	case TYPE_UINT:
		ret = v.GetUint()

	case TYPE_LONG:
		ret = v.GetLong()

	case TYPE_ULONG:
		ret = v.GetUlong()

	case TYPE_INT64:
		ret = v.GetInt64()

	case TYPE_UINT64:
		ret = v.GetUint64()

	case TYPE_ENUM:
		ret = v.GetEnum()

	case TYPE_FLAGS:
		ret = v.GetFlags()

	case TYPE_FLOAT:
		ret = v.GetFloat()

	case TYPE_DOUBLE:
		ret = v.GetDouble()

	case TYPE_STRING:
		ret = v.GetString()

	case TYPE_POINTER:
		ret = v.GetPointer()

	case TYPE_BOXED:
		ret = v.GetBoxed()

	case TYPE_PARAM:
		ret = v.GetParam()

	case TYPE_OBJECT:
		ret = v.GetObject()

	case TYPE_VARIANT:
		ret = v.GetVariant()

	default:
		err = errTypeUnknown
	}

	//gvalueGetters.Lock()
	//getter, ok := gvalueGetters.m[typ]
	//gvalueGetters.Unlock()
	//if !ok {
	//	return nil, errTypeUnknown
	//}
	//return getter(v.P)
	return
}

//func (v Value) GetWithType(reflectType reflect.Type) (interface{}, error) {
//	kind := reflectType.Kind()
//	switch kind {
//	case reflect.Bool:
//		val := v.GetBoolean()
//		return val, nil
//
//	case reflect.Int:
//		_, fundType, err := v.Type()
//		if err != nil {
//			return nil, err
//		}
//
//		switch fundType {
//		case TYPE_ENUM:
//			val := v.GetEnum()
//			return val, nil
//
//		default:
//			val := v.GetInt()
//			return val, nil
//		}
//
//	case reflect.Int8:
//		val := v.GetSchar()
//		return val, nil
//
//	case reflect.Int16:
//		val := v.GetInt()
//		return int16(val), nil
//
//	case reflect.Int32:
//		val := v.GetInt()
//		return int32(val), nil
//
//	case reflect.Int64:
//		val := v.GetInt64()
//		return val, nil
//
//	case reflect.Uint:
//		val := v.GetUint()
//		return val, nil
//
//	case reflect.Uint8:
//		val := v.GetUchar()
//		return val, nil
//
//	case reflect.Uint16:
//		val := v.GetUint()
//		return uint16(val), nil
//
//	case reflect.Uint32:
//		val := v.GetUint()
//		return uint32(val), nil
//
//	case reflect.Uint64:
//		val := v.GetUint64()
//		return val, nil
//
//	case reflect.Uintptr:
//		val := v.GetPointer()
//		return uintptr(val), nil
//
//	case reflect.Float32:
//		val := v.GetFloat()
//		return val, nil
//
//	case reflect.Float64:
//		val := v.GetDouble()
//		return val, nil
//
//	case reflect.UnsafePointer:
//		val := v.GetPointer()
//		return val, nil
//
//	case reflect.String:
//		val := v.GetString()
//		return val, nil
//
//	case reflect.Struct:
//		val := unsafe.Pointer(C.g_value_get_object(v.native()))
//
//		newValPtr := reflect.New(reflectType)
//		newVal := newValPtr.Elem()
//		ptrFieldVal := newVal.FieldByName("P")
//		ptrFieldVal.SetPointer(val)
//
//		return newVal.Interface(), nil
//
//	default:
//		// Complex64, Complex128
//		// Array
//		// Chan
//		// Func
//		// Interface
//		// Map
//		// Ptr
//		// Slice
//		return nil, errors.New("unsupported reflect type")
//	}
//}

var errTypeConvert = errors.New("type convert failed")

func (v Value) Set(iVal interface{}) error {
	cType := C._g_value_type(v.native())
	gType := gi.GType(cType)
	switch gType {
	case TYPE_INVALID:
		return errors.New("type is invalid")

	case TYPE_NONE:
		return nil

	case TYPE_INTERFACE:
		return errors.New("unsupported type interface")

	case TYPE_CHAR:
		val, ok := iVal.(int8)
		if !ok {
			return errTypeConvert
		}
		v.SetSchar(val)

	case TYPE_UCHAR:
		val, ok := iVal.(uint8)
		if !ok {
			return errTypeConvert
		}
		v.SetUchar(val)

	case TYPE_BOOLEAN:
		val, ok := iVal.(bool)
		if !ok {
			return errTypeConvert
		}
		v.SetBoolean(val)

	case TYPE_INT:
		val, ok := iVal.(int32)
		if !ok {
			return errTypeConvert
		}
		v.SetInt(val)

	case TYPE_UINT:
		val, ok := iVal.(uint32)
		if !ok {
			return errTypeConvert
		}
		v.SetUint(val)

	case TYPE_LONG:
		val, ok := iVal.(int64)
		if !ok {
			return errTypeConvert
		}
		v.SetLong(val)

	case TYPE_ULONG:
		val, ok := iVal.(uint64)
		if !ok {
			return errTypeConvert
		}
		v.SetUlong(val)

	case TYPE_INT64:
		val, ok := iVal.(int64)
		if !ok {
			return errTypeConvert
		}
		v.SetInt64(val)

	case TYPE_UINT64:
		val, ok := iVal.(uint64)
		if !ok {
			return errTypeConvert
		}
		v.SetUint64(val)

	case TYPE_ENUM:
		val, ok := iVal.(int32)
		if !ok {
			return errTypeConvert
		}
		v.SetEnum(val)

	case TYPE_FLAGS:
		val, ok := iVal.(uint32)
		if !ok {
			return errTypeConvert
		}
		v.SetFlags(val)

	case TYPE_FLOAT:
		val, ok := iVal.(float32)
		if !ok {
			return errTypeConvert
		}
		v.SetFloat(val)

	case TYPE_DOUBLE:
		val, ok := iVal.(float64)
		if !ok {
			return errTypeConvert
		}
		v.SetDouble(val)

	case TYPE_STRING:
		val, ok := iVal.(string)
		if !ok {
			return errTypeConvert
		}
		v.SetString(val)

	case TYPE_POINTER:
		val, ok := iVal.(unsafe.Pointer)
		if ok {
			return errTypeConvert
		}
		v.SetPointer(val)

	case TYPE_BOXED:
		val, ok := iVal.(unsafe.Pointer)
		if !ok {
			return errTypeConvert
		}
		v.SetBoxed(val)

	case TYPE_PARAM:
		val, ok := iVal.(ParamSpec)
		if !ok {
			return errTypeConvert
		}
		v.SetParam(val)

	case TYPE_OBJECT:
		val, ok := iVal.(Object)
		if !ok {
			return errTypeConvert
		}
		v.SetObject(val)

	case TYPE_VARIANT:
		val, ok := iVal.(glib.Variant)
		if !ok {
			return errTypeConvert
		}
		v.SetVariant(val)

	default:
		return errTypeConvert
	}

	return nil
}
