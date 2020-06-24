package atspi

/*
#cgo pkg-config: atspi-2
#include <atspi/atspi.h>
extern void myAtspiDeviceListenerCB(AtspiDeviceEvent* stroke, gpointer user_data);
static void* getPointer_myAtspiDeviceListenerCB() {
return (void*)(myAtspiDeviceListenerCB);
}
extern void myAtspiDeviceListenerSimpleCB(AtspiDeviceEvent* stroke);
static void* getPointer_myAtspiDeviceListenerSimpleCB() {
return (void*)(myAtspiDeviceListenerSimpleCB);
}
extern void myAtspiEventListenerCB(AtspiEvent* event, gpointer user_data);
static void* getPointer_myAtspiEventListenerCB() {
return (void*)(myAtspiEventListenerCB);
}
extern void myAtspiEventListenerSimpleCB(AtspiEvent* event);
static void* getPointer_myAtspiEventListenerSimpleCB() {
return (void*)(myAtspiEventListenerSimpleCB);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Atspi")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Atspi", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Object Accessible
type Accessible struct {
	ActionIfc
	CollectionIfc
	ComponentIfc
	DocumentIfc
	EditableTextIfc
	HypertextIfc
	ImageIfc
	SelectionIfc
	TableIfc
	TableCellIfc
	TextIfc
	ValueIfc
	Object
}

func WrapAccessible(p unsafe.Pointer) (r Accessible) { r.P = p; return }

type IAccessible interface{ P_Accessible() unsafe.Pointer }

func (v Accessible) P_Accessible() unsafe.Pointer   { return v.P }
func (v Accessible) P_Action() unsafe.Pointer       { return v.P }
func (v Accessible) P_Collection() unsafe.Pointer   { return v.P }
func (v Accessible) P_Component() unsafe.Pointer    { return v.P }
func (v Accessible) P_Document() unsafe.Pointer     { return v.P }
func (v Accessible) P_EditableText() unsafe.Pointer { return v.P }
func (v Accessible) P_Hypertext() unsafe.Pointer    { return v.P }
func (v Accessible) P_Image() unsafe.Pointer        { return v.P }
func (v Accessible) P_Selection() unsafe.Pointer    { return v.P }
func (v Accessible) P_Table() unsafe.Pointer        { return v.P }
func (v Accessible) P_TableCell() unsafe.Pointer    { return v.P }
func (v Accessible) P_Text() unsafe.Pointer         { return v.P }
func (v Accessible) P_Value() unsafe.Pointer        { return v.P }
func AccessibleGetType() gi.GType {
	ret := _I.GetGType(0, "Accessible")
	return ret
}

// atspi_accessible_clear_cache
//
func (v Accessible) ClearCache() {
	iv, err := _I.Get(0, "Accessible", "clear_cache")
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
// atspi_accessible_get_action
//
// [ result ] trans: everything
//
func (v Accessible) GetActionIface() (result Action) {
	iv, err := _I.Get(1, "Accessible", "get_action_iface")
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

// atspi_accessible_get_application
//
// [ result ] trans: everything
//
func (v Accessible) GetApplication() (result Accessible, err error) {
	iv, err := _I.Get(2, "Accessible", "get_application")
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
	result.P = ret.Pointer()
	return
}

// atspi_accessible_get_atspi_version
//
// [ result ] trans: everything
//
func (v Accessible) GetAtspiVersion() (result string, err error) {
	iv, err := _I.Get(3, "Accessible", "get_atspi_version")
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
	result = ret.String().Take()
	return
}

// atspi_accessible_get_attributes
//
// [ result ] trans: everything
//
func (v Accessible) GetAttributes() (result g.HashTable, err error) {
	iv, err := _I.Get(4, "Accessible", "get_attributes")
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
	result.P = ret.Pointer()
	return
}

// atspi_accessible_get_attributes_as_array
//
// [ result ] trans: everything
//
func (v Accessible) GetAttributesAsArray() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(5, "Accessible", "get_attributes_as_array")
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
	result = ret.Int() /*TODO*/
	return
}

// atspi_accessible_get_child_at_index
//
// [ child_index ] trans: nothing
//
// [ result ] trans: everything
//
func (v Accessible) GetChildAtIndex(child_index int32) (result Accessible, err error) {
	iv, err := _I.Get(6, "Accessible", "get_child_at_index")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_child_index := gi.NewInt32Argument(child_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_child_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_accessible_get_child_count
//
// [ result ] trans: nothing
//
func (v Accessible) GetChildCount() (result int32, err error) {
	iv, err := _I.Get(7, "Accessible", "get_child_count")
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
	result = ret.Int32()
	return
}

// Deprecated
//
// atspi_accessible_get_collection
//
// [ result ] trans: everything
//
func (v Accessible) GetCollectionIface() (result Collection) {
	iv, err := _I.Get(8, "Accessible", "get_collection_iface")
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

// Deprecated
//
// atspi_accessible_get_component
//
// [ result ] trans: everything
//
func (v Accessible) GetComponentIface() (result Component) {
	iv, err := _I.Get(9, "Accessible", "get_component_iface")
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

// atspi_accessible_get_description
//
// [ result ] trans: everything
//
func (v Accessible) GetDescription() (result string, err error) {
	iv, err := _I.Get(10, "Accessible", "get_description")
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
	result = ret.String().Take()
	return
}

// Deprecated
//
// atspi_accessible_get_document
//
// [ result ] trans: everything
//
func (v Accessible) GetDocumentIface() (result Document) {
	iv, err := _I.Get(11, "Accessible", "get_document_iface")
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

// Deprecated
//
// atspi_accessible_get_editable_text
//
// [ result ] trans: everything
//
func (v Accessible) GetEditableTextIface() (result EditableText) {
	iv, err := _I.Get(12, "Accessible", "get_editable_text_iface")
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

// atspi_accessible_get_hyperlink
//
// [ result ] trans: everything
//
func (v Accessible) GetHyperlink() (result Hyperlink) {
	iv, err := _I.Get(13, "Accessible", "get_hyperlink")
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

// Deprecated
//
// atspi_accessible_get_hypertext
//
// [ result ] trans: everything
//
func (v Accessible) GetHypertextIface() (result Hypertext) {
	iv, err := _I.Get(14, "Accessible", "get_hypertext_iface")
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

// atspi_accessible_get_id
//
// [ result ] trans: nothing
//
func (v Accessible) GetId() (result int32, err error) {
	iv, err := _I.Get(15, "Accessible", "get_id")
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
	result = ret.Int32()
	return
}

// Deprecated
//
// atspi_accessible_get_image
//
// [ result ] trans: everything
//
func (v Accessible) GetImageIface() (result Image) {
	iv, err := _I.Get(16, "Accessible", "get_image_iface")
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

// atspi_accessible_get_index_in_parent
//
// [ result ] trans: nothing
//
func (v Accessible) GetIndexInParent() (result int32, err error) {
	iv, err := _I.Get(17, "Accessible", "get_index_in_parent")
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
	result = ret.Int32()
	return
}

// atspi_accessible_get_interfaces
//
// [ result ] trans: everything
//
func (v Accessible) GetInterfaces() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/) {
	iv, err := _I.Get(18, "Accessible", "get_interfaces")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atspi_accessible_get_localized_role_name
//
// [ result ] trans: everything
//
func (v Accessible) GetLocalizedRoleName() (result string, err error) {
	iv, err := _I.Get(19, "Accessible", "get_localized_role_name")
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
	result = ret.String().Take()
	return
}

// atspi_accessible_get_name
//
// [ result ] trans: everything
//
func (v Accessible) GetName() (result string, err error) {
	iv, err := _I.Get(20, "Accessible", "get_name")
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
	result = ret.String().Take()
	return
}

// atspi_accessible_get_object_locale
//
// [ result ] trans: nothing
//
func (v Accessible) GetObjectLocale() (result string, err error) {
	iv, err := _I.Get(21, "Accessible", "get_object_locale")
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
	result = ret.String().Copy()
	return
}

// atspi_accessible_get_parent
//
// [ result ] trans: everything
//
func (v Accessible) GetParent() (result Accessible, err error) {
	iv, err := _I.Get(22, "Accessible", "get_parent")
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
	result.P = ret.Pointer()
	return
}

// atspi_accessible_get_process_id
//
// [ result ] trans: nothing
//
func (v Accessible) GetProcessId() (result uint32, err error) {
	iv, err := _I.Get(23, "Accessible", "get_process_id")
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
	result = ret.Uint32()
	return
}

// atspi_accessible_get_relation_set
//
// [ result ] trans: everything
//
func (v Accessible) GetRelationSet() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(24, "Accessible", "get_relation_set")
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
	result = ret.Int() /*TODO*/
	return
}

// atspi_accessible_get_role
//
// [ result ] trans: nothing
//
func (v Accessible) GetRole() (result RoleEnum, err error) {
	iv, err := _I.Get(25, "Accessible", "get_role")
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
	result = RoleEnum(ret.Int())
	return
}

// atspi_accessible_get_role_name
//
// [ result ] trans: everything
//
func (v Accessible) GetRoleName() (result string, err error) {
	iv, err := _I.Get(26, "Accessible", "get_role_name")
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
	result = ret.String().Take()
	return
}

// Deprecated
//
// atspi_accessible_get_selection
//
// [ result ] trans: everything
//
func (v Accessible) GetSelectionIface() (result Selection) {
	iv, err := _I.Get(27, "Accessible", "get_selection_iface")
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

// atspi_accessible_get_state_set
//
// [ result ] trans: everything
//
func (v Accessible) GetStateSet() (result StateSet) {
	iv, err := _I.Get(28, "Accessible", "get_state_set")
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

// Deprecated
//
// atspi_accessible_get_table
//
// [ result ] trans: everything
//
func (v Accessible) GetTableIface() (result Table) {
	iv, err := _I.Get(29, "Accessible", "get_table_iface")
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

// atspi_accessible_get_table_cell
//
// [ result ] trans: everything
//
func (v Accessible) GetTableCell() (result TableCell) {
	iv, err := _I.Get(30, "Accessible", "get_table_cell")
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

// Deprecated
//
// atspi_accessible_get_text
//
// [ result ] trans: everything
//
func (v Accessible) GetTextIface() (result Text) {
	iv, err := _I.Get(31, "Accessible", "get_text_iface")
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

// atspi_accessible_get_toolkit_name
//
// [ result ] trans: everything
//
func (v Accessible) GetToolkitName() (result string, err error) {
	iv, err := _I.Get(32, "Accessible", "get_toolkit_name")
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
	result = ret.String().Take()
	return
}

// atspi_accessible_get_toolkit_version
//
// [ result ] trans: everything
//
func (v Accessible) GetToolkitVersion() (result string, err error) {
	iv, err := _I.Get(33, "Accessible", "get_toolkit_version")
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
	result = ret.String().Take()
	return
}

// Deprecated
//
// atspi_accessible_get_value
//
// [ result ] trans: everything
//
func (v Accessible) GetValueIface() (result Value) {
	iv, err := _I.Get(34, "Accessible", "get_value_iface")
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

// atspi_accessible_set_cache_mask
//
// [ mask ] trans: nothing
//
func (v Accessible) SetCacheMask(mask CacheFlags) {
	iv, err := _I.Get(35, "Accessible", "set_cache_mask")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewIntArgument(int(mask))
	args := []gi.Argument{arg_v, arg_mask}
	iv.Call(args, nil, nil)
}

// ignore GType struct AccessibleClass

// Struct AccessiblePrivate
type AccessiblePrivate struct {
	P unsafe.Pointer
}

func AccessiblePrivateGetType() gi.GType {
	ret := _I.GetGType(1, "AccessiblePrivate")
	return ret
}

// Interface Action
type Action struct {
	ActionIfc
	P unsafe.Pointer
}
type ActionIfc struct{}
type IAction interface{ P_Action() unsafe.Pointer }

func (v Action) P_Action() unsafe.Pointer { return v.P }
func ActionGetType() gi.GType {
	ret := _I.GetGType(2, "Action")
	return ret
}

// atspi_action_do_action
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) DoAction(i int32) (result bool, err error) {
	iv, err := _I.Get(36, "Action", "do_action")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Deprecated
//
// atspi_action_get_description
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ActionIfc) GetActionDescription(i int32) (result string, err error) {
	iv, err := _I.Get(37, "Action", "get_action_description")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_action_get_key_binding
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ActionIfc) GetKeyBinding(i int32) (result string, err error) {
	iv, err := _I.Get(38, "Action", "get_key_binding")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_action_get_localized_name
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ActionIfc) GetLocalizedName(i int32) (result string, err error) {
	iv, err := _I.Get(39, "Action", "get_localized_name")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_action_get_n_actions
//
// [ result ] trans: nothing
//
func (v *ActionIfc) GetNActions() (result int32, err error) {
	iv, err := _I.Get(40, "Action", "get_n_actions")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// Deprecated
//
// atspi_action_get_name
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ActionIfc) GetActionName(i int32) (result string, err error) {
	iv, err := _I.Get(41, "Action", "get_action_name")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// Struct Application
type Application struct {
	P unsafe.Pointer
}

const SizeOfStructApplication = 96

func ApplicationGetType() gi.GType {
	ret := _I.GetGType(3, "Application")
	return ret
}

// ignore Class struct ApplicationClass, type of Application is struct

// Flags Cache
type CacheFlags int

const (
	CacheNone        CacheFlags = 0
	CacheParent      CacheFlags = 1
	CacheChildren    CacheFlags = 2
	CacheName        CacheFlags = 4
	CacheDescription CacheFlags = 8
	CacheStates      CacheFlags = 16
	CacheRole        CacheFlags = 32
	CacheInterfaces  CacheFlags = 64
	CacheAttributes  CacheFlags = 128
	CacheAll         CacheFlags = 1073741823
	CacheDefault     CacheFlags = 127
	CacheUndefined   CacheFlags = 1073741824
)

func CacheGetType() gi.GType {
	ret := _I.GetGType(4, "Cache")
	return ret
}

// Interface Collection
type Collection struct {
	CollectionIfc
	P unsafe.Pointer
}
type CollectionIfc struct{}
type ICollection interface{ P_Collection() unsafe.Pointer }

func (v Collection) P_Collection() unsafe.Pointer { return v.P }
func CollectionGetType() gi.GType {
	ret := _I.GetGType(5, "Collection")
	return ret
}

// atspi_collection_get_active_descendant
//
// [ result ] trans: everything
//
func (v *CollectionIfc) GetActiveDescendant() (result Accessible, err error) {
	iv, err := _I.Get(42, "Collection", "get_active_descendant")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_collection_get_matches
//
// [ rule ] trans: nothing
//
// [ sortby ] trans: nothing
//
// [ count ] trans: nothing
//
// [ traverse ] trans: nothing
//
// [ result ] trans: everything
//
func (v *CollectionIfc) GetMatches(rule IMatchRule, sortby CollectionSortOrderEnum, count int32, traverse bool) (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(43, "Collection", "get_matches")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if rule != nil {
		tmp = rule.P_MatchRule()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_rule := gi.NewPointerArgument(tmp)
	arg_sortby := gi.NewIntArgument(int(sortby))
	arg_count := gi.NewInt32Argument(count)
	arg_traverse := gi.NewBoolArgument(traverse)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_rule, arg_sortby, arg_count, arg_traverse, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_collection_get_matches_from
//
// [ current_object ] trans: nothing
//
// [ rule ] trans: nothing
//
// [ sortby ] trans: nothing
//
// [ tree ] trans: nothing
//
// [ count ] trans: nothing
//
// [ traverse ] trans: nothing
//
// [ result ] trans: everything
//
func (v *CollectionIfc) GetMatchesFrom(current_object IAccessible, rule IMatchRule, sortby CollectionSortOrderEnum, tree CollectionTreeTraversalTypeEnum, count int32, traverse bool) (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(44, "Collection", "get_matches_from")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if current_object != nil {
		tmp = current_object.P_Accessible()
	}
	var tmp1 unsafe.Pointer
	if rule != nil {
		tmp1 = rule.P_MatchRule()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_current_object := gi.NewPointerArgument(tmp)
	arg_rule := gi.NewPointerArgument(tmp1)
	arg_sortby := gi.NewIntArgument(int(sortby))
	arg_tree := gi.NewIntArgument(int(tree))
	arg_count := gi.NewInt32Argument(count)
	arg_traverse := gi.NewBoolArgument(traverse)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_current_object, arg_rule, arg_sortby, arg_tree, arg_count, arg_traverse, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_collection_get_matches_to
//
// [ current_object ] trans: nothing
//
// [ rule ] trans: nothing
//
// [ sortby ] trans: nothing
//
// [ tree ] trans: nothing
//
// [ limit_scope ] trans: nothing
//
// [ count ] trans: nothing
//
// [ traverse ] trans: nothing
//
// [ result ] trans: everything
//
func (v *CollectionIfc) GetMatchesTo(current_object IAccessible, rule IMatchRule, sortby CollectionSortOrderEnum, tree CollectionTreeTraversalTypeEnum, limit_scope bool, count int32, traverse bool) (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(45, "Collection", "get_matches_to")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if current_object != nil {
		tmp = current_object.P_Accessible()
	}
	var tmp1 unsafe.Pointer
	if rule != nil {
		tmp1 = rule.P_MatchRule()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_current_object := gi.NewPointerArgument(tmp)
	arg_rule := gi.NewPointerArgument(tmp1)
	arg_sortby := gi.NewIntArgument(int(sortby))
	arg_tree := gi.NewIntArgument(int(tree))
	arg_limit_scope := gi.NewBoolArgument(limit_scope)
	arg_count := gi.NewInt32Argument(count)
	arg_traverse := gi.NewBoolArgument(traverse)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_current_object, arg_rule, arg_sortby, arg_tree, arg_limit_scope, arg_count, arg_traverse, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_collection_is_ancestor_of
//
// [ test ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *CollectionIfc) IsAncestorOf(test IAccessible) (result bool, err error) {
	iv, err := _I.Get(46, "Collection", "is_ancestor_of")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if test != nil {
		tmp = test.P_Accessible()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_test := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_test, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Enum CollectionMatchType
type CollectionMatchTypeEnum int

const (
	CollectionMatchTypeInvalid     CollectionMatchTypeEnum = 0
	CollectionMatchTypeAll         CollectionMatchTypeEnum = 1
	CollectionMatchTypeAny         CollectionMatchTypeEnum = 2
	CollectionMatchTypeNone        CollectionMatchTypeEnum = 3
	CollectionMatchTypeEmpty       CollectionMatchTypeEnum = 4
	CollectionMatchTypeLastDefined CollectionMatchTypeEnum = 5
)

func CollectionMatchTypeGetType() gi.GType {
	ret := _I.GetGType(6, "CollectionMatchType")
	return ret
}

// Enum CollectionSortOrder
type CollectionSortOrderEnum int

const (
	CollectionSortOrderInvalid          CollectionSortOrderEnum = 0
	CollectionSortOrderCanonical        CollectionSortOrderEnum = 1
	CollectionSortOrderFlow             CollectionSortOrderEnum = 2
	CollectionSortOrderTab              CollectionSortOrderEnum = 3
	CollectionSortOrderReverseCanonical CollectionSortOrderEnum = 4
	CollectionSortOrderReverseFlow      CollectionSortOrderEnum = 5
	CollectionSortOrderReverseTab       CollectionSortOrderEnum = 6
	CollectionSortOrderLastDefined      CollectionSortOrderEnum = 7
)

func CollectionSortOrderGetType() gi.GType {
	ret := _I.GetGType(7, "CollectionSortOrder")
	return ret
}

// Enum CollectionTreeTraversalType
type CollectionTreeTraversalTypeEnum int

const (
	CollectionTreeTraversalTypeRestrictChildren CollectionTreeTraversalTypeEnum = 0
	CollectionTreeTraversalTypeRestrictSibling  CollectionTreeTraversalTypeEnum = 1
	CollectionTreeTraversalTypeInorder          CollectionTreeTraversalTypeEnum = 2
	CollectionTreeTraversalTypeLastDefined      CollectionTreeTraversalTypeEnum = 3
)

func CollectionTreeTraversalTypeGetType() gi.GType {
	ret := _I.GetGType(8, "CollectionTreeTraversalType")
	return ret
}

// Interface Component
type Component struct {
	ComponentIfc
	P unsafe.Pointer
}
type ComponentIfc struct{}
type IComponent interface{ P_Component() unsafe.Pointer }

func (v Component) P_Component() unsafe.Pointer { return v.P }
func ComponentGetType() gi.GType {
	ret := _I.GetGType(9, "Component")
	return ret
}

// atspi_component_contains
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ ctype ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) Contains(x int32, y int32, ctype CoordTypeEnum) (result bool, err error) {
	iv, err := _I.Get(47, "Component", "contains")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_component_get_accessible_at_point
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ ctype ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ComponentIfc) GetAccessibleAtPoint(x int32, y int32, ctype CoordTypeEnum) (result Accessible, err error) {
	iv, err := _I.Get(48, "Component", "get_accessible_at_point")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_component_get_alpha
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GetAlpha() (result float64, err error) {
	iv, err := _I.Get(49, "Component", "get_alpha")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// atspi_component_get_extents
//
// [ ctype ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ComponentIfc) GetExtents(ctype CoordTypeEnum) (result Rect, err error) {
	iv, err := _I.Get(50, "Component", "get_extents")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_component_get_layer
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GetLayer() (result ComponentLayerEnum, err error) {
	iv, err := _I.Get(51, "Component", "get_layer")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ComponentLayerEnum(ret.Int())
	return
}

// atspi_component_get_mdi_z_order
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GetMdiZOrder() (result int16, err error) {
	iv, err := _I.Get(52, "Component", "get_mdi_z_order")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int16()
	return
}

// atspi_component_get_position
//
// [ ctype ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ComponentIfc) GetPosition(ctype CoordTypeEnum) (result Point, err error) {
	iv, err := _I.Get(53, "Component", "get_position")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_component_get_size
//
// [ result ] trans: everything
//
func (v *ComponentIfc) GetSize() (result Point, err error) {
	iv, err := _I.Get(54, "Component", "get_size")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_component_grab_focus
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GrabFocus() (result bool, err error) {
	iv, err := _I.Get(55, "Component", "grab_focus")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_component_scroll_to
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) ScrollTo(type1 ScrollTypeEnum) (result bool, err error) {
	iv, err := _I.Get(56, "Component", "scroll_to")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_component_scroll_to_point
//
// [ coords ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) ScrollToPoint(coords CoordTypeEnum, x int32, y int32) (result bool, err error) {
	iv, err := _I.Get(57, "Component", "scroll_to_point")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_coords := gi.NewIntArgument(int(coords))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_coords, arg_x, arg_y, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_component_set_extents
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ ctype ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) SetExtents(x int32, y int32, width int32, height int32, ctype CoordTypeEnum) (result bool, err error) {
	iv, err := _I.Get(58, "Component", "set_extents")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_component_set_position
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ ctype ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) SetPosition(x int32, y int32, ctype CoordTypeEnum) (result bool, err error) {
	iv, err := _I.Get(59, "Component", "set_position")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_component_set_size
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) SetSize(width int32, height int32) (result bool, err error) {
	iv, err := _I.Get(60, "Component", "set_size")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_width, arg_height, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Enum ComponentLayer
type ComponentLayerEnum int

const (
	ComponentLayerInvalid     ComponentLayerEnum = 0
	ComponentLayerBackground  ComponentLayerEnum = 1
	ComponentLayerCanvas      ComponentLayerEnum = 2
	ComponentLayerWidget      ComponentLayerEnum = 3
	ComponentLayerMdi         ComponentLayerEnum = 4
	ComponentLayerPopup       ComponentLayerEnum = 5
	ComponentLayerOverlay     ComponentLayerEnum = 6
	ComponentLayerWindow      ComponentLayerEnum = 7
	ComponentLayerLastDefined ComponentLayerEnum = 8
)

func ComponentLayerGetType() gi.GType {
	ret := _I.GetGType(10, "ComponentLayer")
	return ret
}

// Enum CoordType
type CoordTypeEnum int

const (
	CoordTypeScreen CoordTypeEnum = 0
	CoordTypeWindow CoordTypeEnum = 1
	CoordTypeParent CoordTypeEnum = 2
)

func CoordTypeGetType() gi.GType {
	ret := _I.GetGType(11, "CoordType")
	return ret
}

// Struct DeviceEvent
type DeviceEvent struct {
	P unsafe.Pointer
}

const SizeOfStructDeviceEvent = 32

func DeviceEventGetType() gi.GType {
	ret := _I.GetGType(12, "DeviceEvent")
	return ret
}

// Object DeviceListener
type DeviceListener struct {
	g.Object
}

func WrapDeviceListener(p unsafe.Pointer) (r DeviceListener) { r.P = p; return }

type IDeviceListener interface{ P_DeviceListener() unsafe.Pointer }

func (v DeviceListener) P_DeviceListener() unsafe.Pointer { return v.P }
func DeviceListenerGetType() gi.GType {
	ret := _I.GetGType(13, "DeviceListener")
	return ret
}

// atspi_device_listener_new
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ callback_destroyed ] trans: nothing
//
// [ result ] trans: everything
//
func NewDeviceListener(callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, callback_destroyed int /*TODO_TYPE CALLBACK*/) (result DeviceListener) {
	iv, err := _I.Get(61, "DeviceListener", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDeviceListenerCB()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_callback_destroyed := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_callback, arg_user_data, arg_callback_destroyed}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atspi_device_listener_add_callback
//
// [ callback ] trans: nothing
//
// [ callback_destroyed ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v DeviceListener) AddCallback(callback int /*TODO_TYPE CALLBACK*/, callback_destroyed int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(62, "DeviceListener", "add_callback")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDeviceListenerCB()))
	arg_callback_destroyed := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_callback, arg_callback_destroyed, arg_user_data}
	iv.Call(args, nil, nil)
}

// atspi_device_listener_remove_callback
//
// [ callback ] trans: nothing
//
func (v DeviceListener) RemoveCallback(callback int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(63, "DeviceListener", "remove_callback")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDeviceListenerCB()))
	args := []gi.Argument{arg_v, arg_callback}
	iv.Call(args, nil, nil)
}

type DeviceListenerCBStruct struct {
	F_stroke DeviceEvent
}

func GetPointer_myDeviceListenerCB() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtspiDeviceListenerCB())
}

//export myAtspiDeviceListenerCB
func myAtspiDeviceListenerCB(stroke *C.AtspiDeviceEvent, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &DeviceListenerCBStruct{
		F_stroke: DeviceEvent{P: unsafe.Pointer(stroke)},
	}
	fn(args)
}

// ignore GType struct DeviceListenerClass

type DeviceListenerSimpleCBStruct struct {
	F_stroke DeviceEvent
}

func GetPointer_myDeviceListenerSimpleCB() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtspiDeviceListenerSimpleCB())
}

//export myAtspiDeviceListenerSimpleCB
func myAtspiDeviceListenerSimpleCB(stroke *C.AtspiDeviceEvent) {
	// TODO: not found user_data
}

// Interface Document
type Document struct {
	DocumentIfc
	P unsafe.Pointer
}
type DocumentIfc struct{}
type IDocument interface{ P_Document() unsafe.Pointer }

func (v Document) P_Document() unsafe.Pointer { return v.P }
func DocumentGetType() gi.GType {
	ret := _I.GetGType(14, "Document")
	return ret
}

// Deprecated
//
// atspi_document_get_attribute_value
//
// [ attribute ] trans: nothing
//
// [ result ] trans: everything
//
func (v *DocumentIfc) GetDocumentAttributeValue(attribute string) (result string, err error) {
	iv, err := _I.Get(64, "Document", "get_document_attribute_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_attribute := gi.CString(attribute)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attribute := gi.NewStringArgument(c_attribute)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_attribute, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_attribute)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// Deprecated
//
// atspi_document_get_attributes
//
// [ result ] trans: everything
//
func (v *DocumentIfc) GetDocumentAttributes() (result g.HashTable, err error) {
	iv, err := _I.Get(65, "Document", "get_document_attributes")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_document_get_current_page_number
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetCurrentPageNumber() (result int32, err error) {
	iv, err := _I.Get(66, "Document", "get_current_page_number")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_document_get_locale
//
// [ result ] trans: everything
//
func (v *DocumentIfc) GetLocale() (result string, err error) {
	iv, err := _I.Get(67, "Document", "get_locale")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_document_get_page_count
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetPageCount() (result int32, err error) {
	iv, err := _I.Get(68, "Document", "get_page_count")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// Interface EditableText
type EditableText struct {
	EditableTextIfc
	P unsafe.Pointer
}
type EditableTextIfc struct{}
type IEditableText interface{ P_EditableText() unsafe.Pointer }

func (v EditableText) P_EditableText() unsafe.Pointer { return v.P }
func EditableTextGetType() gi.GType {
	ret := _I.GetGType(15, "EditableText")
	return ret
}

// atspi_editable_text_copy_text
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) CopyText(start_pos int32, end_pos int32) (result bool, err error) {
	iv, err := _I.Get(69, "EditableText", "copy_text")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_editable_text_cut_text
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) CutText(start_pos int32, end_pos int32) (result bool, err error) {
	iv, err := _I.Get(70, "EditableText", "cut_text")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_editable_text_delete_text
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) DeleteText(start_pos int32, end_pos int32) (result bool, err error) {
	iv, err := _I.Get(71, "EditableText", "delete_text")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_editable_text_insert_text
//
// [ position ] trans: nothing
//
// [ text ] trans: nothing
//
// [ length ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) InsertText(position int32, text string, length int32) (result bool, err error) {
	iv, err := _I.Get(72, "EditableText", "insert_text")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_position := gi.NewInt32Argument(position)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_position, arg_text, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_editable_text_paste_text
//
// [ position ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) PasteText(position int32) (result bool, err error) {
	iv, err := _I.Get(73, "EditableText", "paste_text")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_position := gi.NewInt32Argument(position)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_position, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_editable_text_set_text_contents
//
// [ new_contents ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) SetTextContents(new_contents string) (result bool, err error) {
	iv, err := _I.Get(74, "EditableText", "set_text_contents")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_new_contents := gi.CString(new_contents)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_new_contents := gi.NewStringArgument(c_new_contents)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_new_contents, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_new_contents)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Struct Event
type Event struct {
	P unsafe.Pointer
}

const SizeOfStructEvent = 48

func EventGetType() gi.GType {
	ret := _I.GetGType(16, "Event")
	return ret
}

// Object EventListener
type EventListener struct {
	g.Object
}

func WrapEventListener(p unsafe.Pointer) (r EventListener) { r.P = p; return }

type IEventListener interface{ P_EventListener() unsafe.Pointer }

func (v EventListener) P_EventListener() unsafe.Pointer { return v.P }
func EventListenerGetType() gi.GType {
	ret := _I.GetGType(17, "EventListener")
	return ret
}

// atspi_event_listener_new
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ callback_destroyed ] trans: nothing
//
// [ result ] trans: everything
//
func NewEventListener(callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, callback_destroyed int /*TODO_TYPE CALLBACK*/) (result EventListener) {
	iv, err := _I.Get(77, "EventListener", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myEventListenerCB()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_callback_destroyed := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_callback, arg_user_data, arg_callback_destroyed}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atspi_event_listener_deregister_from_callback
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ event_type ] trans: nothing
//
// [ result ] trans: nothing
//
func EventListenerDeregisterFromCallback1(callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, event_type string) (result bool, err error) {
	iv, err := _I.Get(78, "EventListener", "deregister_from_callback")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_event_type := gi.CString(event_type)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myEventListenerCB()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_event_type := gi.NewStringArgument(c_event_type)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_callback, arg_user_data, arg_event_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_event_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_event_listener_register_from_callback
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ callback_destroyed ] trans: nothing
//
// [ event_type ] trans: nothing
//
// [ result ] trans: nothing
//
func EventListenerRegisterFromCallback1(callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, callback_destroyed int /*TODO_TYPE CALLBACK*/, event_type string) (result bool, err error) {
	iv, err := _I.Get(79, "EventListener", "register_from_callback")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_event_type := gi.CString(event_type)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myEventListenerCB()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_callback_destroyed := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	arg_event_type := gi.NewStringArgument(c_event_type)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_callback, arg_user_data, arg_callback_destroyed, arg_event_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_event_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_event_listener_register_from_callback_full
//
// [ callback ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ callback_destroyed ] trans: nothing
//
// [ event_type ] trans: nothing
//
// [ properties ] trans: nothing
//
// [ result ] trans: nothing
//
func EventListenerRegisterFromCallbackFull1(callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer, callback_destroyed int /*TODO_TYPE CALLBACK*/, event_type string, properties int /*TODO_TYPE isPtr: true, tag: array*/) (result bool, err error) {
	iv, err := _I.Get(80, "EventListener", "register_from_callback_full")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_event_type := gi.CString(event_type)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myEventListenerCB()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_callback_destroyed := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	arg_event_type := gi.NewStringArgument(c_event_type)
	arg_properties := gi.NewIntArgument(properties) /*TODO*/
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_callback, arg_user_data, arg_callback_destroyed, arg_event_type, arg_properties, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_event_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_event_listener_deregister
//
// [ event_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v EventListener) Deregister(event_type string) (result bool, err error) {
	iv, err := _I.Get(81, "EventListener", "deregister")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_event_type := gi.CString(event_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_event_type := gi.NewStringArgument(c_event_type)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_event_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_event_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_event_listener_register
//
// [ event_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v EventListener) Register(event_type string) (result bool, err error) {
	iv, err := _I.Get(82, "EventListener", "register")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_event_type := gi.CString(event_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_event_type := gi.NewStringArgument(c_event_type)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_event_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_event_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_event_listener_register_full
//
// [ event_type ] trans: nothing
//
// [ properties ] trans: nothing
//
// [ result ] trans: nothing
//
func (v EventListener) RegisterFull(event_type string, properties int /*TODO_TYPE isPtr: true, tag: array*/) (result bool, err error) {
	iv, err := _I.Get(83, "EventListener", "register_full")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_event_type := gi.CString(event_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_event_type := gi.NewStringArgument(c_event_type)
	arg_properties := gi.NewIntArgument(properties) /*TODO*/
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_event_type, arg_properties, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_event_type)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

type EventListenerCBStruct struct {
	F_event Event
}

func GetPointer_myEventListenerCB() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtspiEventListenerCB())
}

//export myAtspiEventListenerCB
func myAtspiEventListenerCB(event *C.AtspiEvent, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &EventListenerCBStruct{
		F_event: Event{P: unsafe.Pointer(event)},
	}
	fn(args)
}

// ignore GType struct EventListenerClass

// Struct EventListenerMode
type EventListenerMode struct {
	P unsafe.Pointer
}

const SizeOfStructEventListenerMode = 12

func EventListenerModeGetType() gi.GType {
	ret := _I.GetGType(18, "EventListenerMode")
	return ret
}

type EventListenerSimpleCBStruct struct {
	F_event Event
}

func GetPointer_myEventListenerSimpleCB() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtspiEventListenerSimpleCB())
}

//export myAtspiEventListenerSimpleCB
func myAtspiEventListenerSimpleCB(event *C.AtspiEvent) {
	// TODO: not found user_data
}

// Enum EventType
type EventTypeEnum int

const (
	EventTypeKeyPressedEvent     EventTypeEnum = 0
	EventTypeKeyReleasedEvent    EventTypeEnum = 1
	EventTypeButtonPressedEvent  EventTypeEnum = 2
	EventTypeButtonReleasedEvent EventTypeEnum = 3
)

func EventTypeGetType() gi.GType {
	ret := _I.GetGType(19, "EventType")
	return ret
}

// Object Hyperlink
type Hyperlink struct {
	Object
}

func WrapHyperlink(p unsafe.Pointer) (r Hyperlink) { r.P = p; return }

type IHyperlink interface{ P_Hyperlink() unsafe.Pointer }

func (v Hyperlink) P_Hyperlink() unsafe.Pointer { return v.P }
func HyperlinkGetType() gi.GType {
	ret := _I.GetGType(20, "Hyperlink")
	return ret
}

// atspi_hyperlink_get_end_index
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetEndIndex() (result int32, err error) {
	iv, err := _I.Get(84, "Hyperlink", "get_end_index")
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
	result = ret.Int32()
	return
}

// atspi_hyperlink_get_index_range
//
// [ result ] trans: everything
//
func (v Hyperlink) GetIndexRange() (result Range, err error) {
	iv, err := _I.Get(85, "Hyperlink", "get_index_range")
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
	result.P = ret.Pointer()
	return
}

// atspi_hyperlink_get_n_anchors
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetNAnchors() (result int32, err error) {
	iv, err := _I.Get(86, "Hyperlink", "get_n_anchors")
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
	result = ret.Int32()
	return
}

// atspi_hyperlink_get_object
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v Hyperlink) GetObject(i int32) (result Accessible, err error) {
	iv, err := _I.Get(87, "Hyperlink", "get_object")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_hyperlink_get_start_index
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetStartIndex() (result int32, err error) {
	iv, err := _I.Get(88, "Hyperlink", "get_start_index")
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
	result = ret.Int32()
	return
}

// atspi_hyperlink_get_uri
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v Hyperlink) GetUri(i int32) (result string, err error) {
	iv, err := _I.Get(89, "Hyperlink", "get_uri")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_i, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_hyperlink_is_valid
//
// [ result ] trans: nothing
//
func (v Hyperlink) IsValid() (result bool, err error) {
	iv, err := _I.Get(90, "Hyperlink", "is_valid")
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

// ignore GType struct HyperlinkClass

// Interface Hypertext
type Hypertext struct {
	HypertextIfc
	P unsafe.Pointer
}
type HypertextIfc struct{}
type IHypertext interface{ P_Hypertext() unsafe.Pointer }

func (v Hypertext) P_Hypertext() unsafe.Pointer { return v.P }
func HypertextGetType() gi.GType {
	ret := _I.GetGType(21, "Hypertext")
	return ret
}

// atspi_hypertext_get_link
//
// [ link_index ] trans: nothing
//
// [ result ] trans: everything
//
func (v *HypertextIfc) GetLink(link_index int32) (result Hyperlink, err error) {
	iv, err := _I.Get(91, "Hypertext", "get_link")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_link_index := gi.NewInt32Argument(link_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_link_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_hypertext_get_link_index
//
// [ character_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *HypertextIfc) GetLinkIndex(character_offset int32) (result int32, err error) {
	iv, err := _I.Get(92, "Hypertext", "get_link_index")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_character_offset := gi.NewInt32Argument(character_offset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_character_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_hypertext_get_n_links
//
// [ result ] trans: nothing
//
func (v *HypertextIfc) GetNLinks() (result int32, err error) {
	iv, err := _I.Get(93, "Hypertext", "get_n_links")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// Interface Image
type Image struct {
	ImageIfc
	P unsafe.Pointer
}
type ImageIfc struct{}
type IImage interface{ P_Image() unsafe.Pointer }

func (v Image) P_Image() unsafe.Pointer { return v.P }
func ImageGetType() gi.GType {
	ret := _I.GetGType(22, "Image")
	return ret
}

// atspi_image_get_image_description
//
// [ result ] trans: everything
//
func (v *ImageIfc) GetImageDescription() (result string, err error) {
	iv, err := _I.Get(94, "Image", "get_image_description")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_image_get_image_extents
//
// [ ctype ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ImageIfc) GetImageExtents(ctype CoordTypeEnum) (result Rect, err error) {
	iv, err := _I.Get(95, "Image", "get_image_extents")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_image_get_image_locale
//
// [ result ] trans: everything
//
func (v *ImageIfc) GetImageLocale() (result string, err error) {
	iv, err := _I.Get(96, "Image", "get_image_locale")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_image_get_image_position
//
// [ ctype ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ImageIfc) GetImagePosition(ctype CoordTypeEnum) (result Point, err error) {
	iv, err := _I.Get(97, "Image", "get_image_position")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_ctype := gi.NewIntArgument(int(ctype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_ctype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_image_get_image_size
//
// [ result ] trans: everything
//
func (v *ImageIfc) GetImageSize() (result Point, err error) {
	iv, err := _I.Get(98, "Image", "get_image_size")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// Struct KeyDefinition
type KeyDefinition struct {
	P unsafe.Pointer
}

const SizeOfStructKeyDefinition = 24

func KeyDefinitionGetType() gi.GType {
	ret := _I.GetGType(23, "KeyDefinition")
	return ret
}

// Enum KeyEventType
type KeyEventTypeEnum int

const (
	KeyEventTypePressed  KeyEventTypeEnum = 0
	KeyEventTypeReleased KeyEventTypeEnum = 1
)

func KeyEventTypeGetType() gi.GType {
	ret := _I.GetGType(24, "KeyEventType")
	return ret
}

// Flags KeyListenerSyncType
type KeyListenerSyncTypeFlags int

const (
	KeyListenerSyncTypeNosync      KeyListenerSyncTypeFlags = 0
	KeyListenerSyncTypeSynchronous KeyListenerSyncTypeFlags = 1
	KeyListenerSyncTypeCanconsume  KeyListenerSyncTypeFlags = 2
	KeyListenerSyncTypeAllWindows  KeyListenerSyncTypeFlags = 4
)

func KeyListenerSyncTypeGetType() gi.GType {
	ret := _I.GetGType(25, "KeyListenerSyncType")
	return ret
}

// Struct KeySet
type KeySet struct {
	P unsafe.Pointer
}

const SizeOfStructKeySet = 32

func KeySetGetType() gi.GType {
	ret := _I.GetGType(26, "KeySet")
	return ret
}

// Enum KeySynthType
type KeySynthTypeEnum int

const (
	KeySynthTypePress           KeySynthTypeEnum = 0
	KeySynthTypeRelease         KeySynthTypeEnum = 1
	KeySynthTypePressrelease    KeySynthTypeEnum = 2
	KeySynthTypeSym             KeySynthTypeEnum = 3
	KeySynthTypeString          KeySynthTypeEnum = 4
	KeySynthTypeLockmodifiers   KeySynthTypeEnum = 5
	KeySynthTypeUnlockmodifiers KeySynthTypeEnum = 6
)

func KeySynthTypeGetType() gi.GType {
	ret := _I.GetGType(27, "KeySynthType")
	return ret
}

// Enum LocaleType
type LocaleTypeEnum int

const (
	LocaleTypeMessages LocaleTypeEnum = 0
	LocaleTypeCollate  LocaleTypeEnum = 1
	LocaleTypeCtype    LocaleTypeEnum = 2
	LocaleTypeMonetary LocaleTypeEnum = 3
	LocaleTypeNumeric  LocaleTypeEnum = 4
	LocaleTypeTime     LocaleTypeEnum = 5
)

func LocaleTypeGetType() gi.GType {
	ret := _I.GetGType(28, "LocaleType")
	return ret
}

// Object MatchRule
type MatchRule struct {
	g.Object
}

func WrapMatchRule(p unsafe.Pointer) (r MatchRule) { r.P = p; return }

type IMatchRule interface{ P_MatchRule() unsafe.Pointer }

func (v MatchRule) P_MatchRule() unsafe.Pointer { return v.P }
func MatchRuleGetType() gi.GType {
	ret := _I.GetGType(29, "MatchRule")
	return ret
}

// atspi_match_rule_new
//
// [ states ] trans: nothing
//
// [ statematchtype ] trans: nothing
//
// [ attributes ] trans: nothing
//
// [ attributematchtype ] trans: nothing
//
// [ roles ] trans: nothing
//
// [ rolematchtype ] trans: nothing
//
// [ interfaces ] trans: nothing
//
// [ interfacematchtype ] trans: nothing
//
// [ invert ] trans: nothing
//
// [ result ] trans: everything
//
func NewMatchRule(states IStateSet, statematchtype CollectionMatchTypeEnum, attributes g.HashTable, attributematchtype CollectionMatchTypeEnum, roles int /*TODO_TYPE isPtr: true, tag: array*/, rolematchtype CollectionMatchTypeEnum, interfaces int /*TODO_TYPE isPtr: true, tag: array*/, interfacematchtype CollectionMatchTypeEnum, invert bool) (result MatchRule) {
	iv, err := _I.Get(99, "MatchRule", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if states != nil {
		tmp = states.P_StateSet()
	}
	arg_states := gi.NewPointerArgument(tmp)
	arg_statematchtype := gi.NewIntArgument(int(statematchtype))
	arg_attributes := gi.NewPointerArgument(attributes.P)
	arg_attributematchtype := gi.NewIntArgument(int(attributematchtype))
	arg_roles := gi.NewIntArgument(roles) /*TODO*/
	arg_rolematchtype := gi.NewIntArgument(int(rolematchtype))
	arg_interfaces := gi.NewIntArgument(interfaces) /*TODO*/
	arg_interfacematchtype := gi.NewIntArgument(int(interfacematchtype))
	arg_invert := gi.NewBoolArgument(invert)
	args := []gi.Argument{arg_states, arg_statematchtype, arg_attributes, arg_attributematchtype, arg_roles, arg_rolematchtype, arg_interfaces, arg_interfacematchtype, arg_invert}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct MatchRuleClass

// Enum ModifierType
type ModifierTypeEnum int

const (
	ModifierTypeShift     ModifierTypeEnum = 0
	ModifierTypeShiftlock ModifierTypeEnum = 1
	ModifierTypeControl   ModifierTypeEnum = 2
	ModifierTypeAlt       ModifierTypeEnum = 3
	ModifierTypeMeta      ModifierTypeEnum = 4
	ModifierTypeMeta2     ModifierTypeEnum = 5
	ModifierTypeMeta3     ModifierTypeEnum = 6
	ModifierTypeNumlock   ModifierTypeEnum = 14
)

func ModifierTypeGetType() gi.GType {
	ret := _I.GetGType(30, "ModifierType")
	return ret
}

// Object Object
type Object struct {
	g.Object
}

func WrapObject(p unsafe.Pointer) (r Object) { r.P = p; return }

type IObject interface{ P_Object() unsafe.Pointer }

func (v Object) P_Object() unsafe.Pointer { return v.P }
func ObjectGetType() gi.GType {
	ret := _I.GetGType(31, "Object")
	return ret
}

// ignore GType struct ObjectClass

// Struct Point
type Point struct {
	P unsafe.Pointer
}

const SizeOfStructPoint = 8

func PointGetType() gi.GType {
	ret := _I.GetGType(32, "Point")
	return ret
}

// atspi_point_copy
//
// [ result ] trans: everything
//
func (v Point) Copy() (result Point) {
	iv, err := _I.Get(100, "Point", "copy")
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

// Struct Range
type Range struct {
	P unsafe.Pointer
}

const SizeOfStructRange = 8

func RangeGetType() gi.GType {
	ret := _I.GetGType(33, "Range")
	return ret
}

// atspi_range_copy
//
// [ result ] trans: everything
//
func (v Range) Copy() (result Range) {
	iv, err := _I.Get(101, "Range", "copy")
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

// Struct Rect
type Rect struct {
	P unsafe.Pointer
}

const SizeOfStructRect = 16

func RectGetType() gi.GType {
	ret := _I.GetGType(34, "Rect")
	return ret
}

// atspi_rect_copy
//
// [ result ] trans: everything
//
func (v Rect) Copy() (result Rect) {
	iv, err := _I.Get(102, "Rect", "copy")
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

// Object Relation
type Relation struct {
	g.Object
}

func WrapRelation(p unsafe.Pointer) (r Relation) { r.P = p; return }

type IRelation interface{ P_Relation() unsafe.Pointer }

func (v Relation) P_Relation() unsafe.Pointer { return v.P }
func RelationGetType() gi.GType {
	ret := _I.GetGType(35, "Relation")
	return ret
}

// atspi_relation_get_n_targets
//
// [ result ] trans: nothing
//
func (v Relation) GetNTargets() (result int32) {
	iv, err := _I.Get(103, "Relation", "get_n_targets")
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

// atspi_relation_get_relation_type
//
// [ result ] trans: nothing
//
func (v Relation) GetRelationType() (result RelationTypeEnum) {
	iv, err := _I.Get(104, "Relation", "get_relation_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = RelationTypeEnum(ret.Int())
	return
}

// atspi_relation_get_target
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v Relation) GetTarget(i int32) (result Accessible) {
	iv, err := _I.Get(105, "Relation", "get_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct RelationClass

// Enum RelationType
type RelationTypeEnum int

const (
	RelationTypeNull           RelationTypeEnum = 0
	RelationTypeLabelFor       RelationTypeEnum = 1
	RelationTypeLabelledBy     RelationTypeEnum = 2
	RelationTypeControllerFor  RelationTypeEnum = 3
	RelationTypeControlledBy   RelationTypeEnum = 4
	RelationTypeMemberOf       RelationTypeEnum = 5
	RelationTypeTooltipFor     RelationTypeEnum = 6
	RelationTypeNodeChildOf    RelationTypeEnum = 7
	RelationTypeNodeParentOf   RelationTypeEnum = 8
	RelationTypeExtended       RelationTypeEnum = 9
	RelationTypeFlowsTo        RelationTypeEnum = 10
	RelationTypeFlowsFrom      RelationTypeEnum = 11
	RelationTypeSubwindowOf    RelationTypeEnum = 12
	RelationTypeEmbeds         RelationTypeEnum = 13
	RelationTypeEmbeddedBy     RelationTypeEnum = 14
	RelationTypePopupFor       RelationTypeEnum = 15
	RelationTypeParentWindowOf RelationTypeEnum = 16
	RelationTypeDescriptionFor RelationTypeEnum = 17
	RelationTypeDescribedBy    RelationTypeEnum = 18
	RelationTypeDetails        RelationTypeEnum = 19
	RelationTypeDetailsFor     RelationTypeEnum = 20
	RelationTypeErrorMessage   RelationTypeEnum = 21
	RelationTypeErrorFor       RelationTypeEnum = 22
	RelationTypeLastDefined    RelationTypeEnum = 23
)

func RelationTypeGetType() gi.GType {
	ret := _I.GetGType(36, "RelationType")
	return ret
}

// Enum Role
type RoleEnum int

const (
	RoleInvalid              RoleEnum = 0
	RoleAcceleratorLabel     RoleEnum = 1
	RoleAlert                RoleEnum = 2
	RoleAnimation            RoleEnum = 3
	RoleArrow                RoleEnum = 4
	RoleCalendar             RoleEnum = 5
	RoleCanvas               RoleEnum = 6
	RoleCheckBox             RoleEnum = 7
	RoleCheckMenuItem        RoleEnum = 8
	RoleColorChooser         RoleEnum = 9
	RoleColumnHeader         RoleEnum = 10
	RoleComboBox             RoleEnum = 11
	RoleDateEditor           RoleEnum = 12
	RoleDesktopIcon          RoleEnum = 13
	RoleDesktopFrame         RoleEnum = 14
	RoleDial                 RoleEnum = 15
	RoleDialog               RoleEnum = 16
	RoleDirectoryPane        RoleEnum = 17
	RoleDrawingArea          RoleEnum = 18
	RoleFileChooser          RoleEnum = 19
	RoleFiller               RoleEnum = 20
	RoleFocusTraversable     RoleEnum = 21
	RoleFontChooser          RoleEnum = 22
	RoleFrame                RoleEnum = 23
	RoleGlassPane            RoleEnum = 24
	RoleHtmlContainer        RoleEnum = 25
	RoleIcon                 RoleEnum = 26
	RoleImage                RoleEnum = 27
	RoleInternalFrame        RoleEnum = 28
	RoleLabel                RoleEnum = 29
	RoleLayeredPane          RoleEnum = 30
	RoleList                 RoleEnum = 31
	RoleListItem             RoleEnum = 32
	RoleMenu                 RoleEnum = 33
	RoleMenuBar              RoleEnum = 34
	RoleMenuItem             RoleEnum = 35
	RoleOptionPane           RoleEnum = 36
	RolePageTab              RoleEnum = 37
	RolePageTabList          RoleEnum = 38
	RolePanel                RoleEnum = 39
	RolePasswordText         RoleEnum = 40
	RolePopupMenu            RoleEnum = 41
	RoleProgressBar          RoleEnum = 42
	RolePushButton           RoleEnum = 43
	RoleRadioButton          RoleEnum = 44
	RoleRadioMenuItem        RoleEnum = 45
	RoleRootPane             RoleEnum = 46
	RoleRowHeader            RoleEnum = 47
	RoleScrollBar            RoleEnum = 48
	RoleScrollPane           RoleEnum = 49
	RoleSeparator            RoleEnum = 50
	RoleSlider               RoleEnum = 51
	RoleSpinButton           RoleEnum = 52
	RoleSplitPane            RoleEnum = 53
	RoleStatusBar            RoleEnum = 54
	RoleTable                RoleEnum = 55
	RoleTableCell            RoleEnum = 56
	RoleTableColumnHeader    RoleEnum = 57
	RoleTableRowHeader       RoleEnum = 58
	RoleTearoffMenuItem      RoleEnum = 59
	RoleTerminal             RoleEnum = 60
	RoleText                 RoleEnum = 61
	RoleToggleButton         RoleEnum = 62
	RoleToolBar              RoleEnum = 63
	RoleToolTip              RoleEnum = 64
	RoleTree                 RoleEnum = 65
	RoleTreeTable            RoleEnum = 66
	RoleUnknown              RoleEnum = 67
	RoleViewport             RoleEnum = 68
	RoleWindow               RoleEnum = 69
	RoleExtended             RoleEnum = 70
	RoleHeader               RoleEnum = 71
	RoleFooter               RoleEnum = 72
	RoleParagraph            RoleEnum = 73
	RoleRuler                RoleEnum = 74
	RoleApplication          RoleEnum = 75
	RoleAutocomplete         RoleEnum = 76
	RoleEditbar              RoleEnum = 77
	RoleEmbedded             RoleEnum = 78
	RoleEntry                RoleEnum = 79
	RoleChart                RoleEnum = 80
	RoleCaption              RoleEnum = 81
	RoleDocumentFrame        RoleEnum = 82
	RoleHeading              RoleEnum = 83
	RolePage                 RoleEnum = 84
	RoleSection              RoleEnum = 85
	RoleRedundantObject      RoleEnum = 86
	RoleForm                 RoleEnum = 87
	RoleLink                 RoleEnum = 88
	RoleInputMethodWindow    RoleEnum = 89
	RoleTableRow             RoleEnum = 90
	RoleTreeItem             RoleEnum = 91
	RoleDocumentSpreadsheet  RoleEnum = 92
	RoleDocumentPresentation RoleEnum = 93
	RoleDocumentText         RoleEnum = 94
	RoleDocumentWeb          RoleEnum = 95
	RoleDocumentEmail        RoleEnum = 96
	RoleComment              RoleEnum = 97
	RoleListBox              RoleEnum = 98
	RoleGrouping             RoleEnum = 99
	RoleImageMap             RoleEnum = 100
	RoleNotification         RoleEnum = 101
	RoleInfoBar              RoleEnum = 102
	RoleLevelBar             RoleEnum = 103
	RoleTitleBar             RoleEnum = 104
	RoleBlockQuote           RoleEnum = 105
	RoleAudio                RoleEnum = 106
	RoleVideo                RoleEnum = 107
	RoleDefinition           RoleEnum = 108
	RoleArticle              RoleEnum = 109
	RoleLandmark             RoleEnum = 110
	RoleLog                  RoleEnum = 111
	RoleMarquee              RoleEnum = 112
	RoleMath                 RoleEnum = 113
	RoleRating               RoleEnum = 114
	RoleTimer                RoleEnum = 115
	RoleStatic               RoleEnum = 116
	RoleMathFraction         RoleEnum = 117
	RoleMathRoot             RoleEnum = 118
	RoleSubscript            RoleEnum = 119
	RoleSuperscript          RoleEnum = 120
	RoleDescriptionList      RoleEnum = 121
	RoleDescriptionTerm      RoleEnum = 122
	RoleDescriptionValue     RoleEnum = 123
	RoleFootnote             RoleEnum = 124
	RoleLastDefined          RoleEnum = 125
)

func RoleGetType() gi.GType {
	ret := _I.GetGType(37, "Role")
	return ret
}

// Enum ScrollType
type ScrollTypeEnum int

const (
	ScrollTypeTopLeft     ScrollTypeEnum = 0
	ScrollTypeBottomRight ScrollTypeEnum = 1
	ScrollTypeTopEdge     ScrollTypeEnum = 2
	ScrollTypeBottomEdge  ScrollTypeEnum = 3
	ScrollTypeLeftEdge    ScrollTypeEnum = 4
	ScrollTypeRightEdge   ScrollTypeEnum = 5
	ScrollTypeAnywhere    ScrollTypeEnum = 6
)

func ScrollTypeGetType() gi.GType {
	ret := _I.GetGType(38, "ScrollType")
	return ret
}

// Interface Selection
type Selection struct {
	SelectionIfc
	P unsafe.Pointer
}
type SelectionIfc struct{}
type ISelection interface{ P_Selection() unsafe.Pointer }

func (v Selection) P_Selection() unsafe.Pointer { return v.P }
func SelectionGetType() gi.GType {
	ret := _I.GetGType(39, "Selection")
	return ret
}

// atspi_selection_clear_selection
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) ClearSelection() (result bool, err error) {
	iv, err := _I.Get(106, "Selection", "clear_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_selection_deselect_child
//
// [ child_index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) DeselectChild(child_index int32) (result bool, err error) {
	iv, err := _I.Get(107, "Selection", "deselect_child")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_child_index := gi.NewInt32Argument(child_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_child_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_selection_deselect_selected_child
//
// [ selected_child_index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) DeselectSelectedChild(selected_child_index int32) (result bool, err error) {
	iv, err := _I.Get(108, "Selection", "deselect_selected_child")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selected_child_index := gi.NewInt32Argument(selected_child_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_selected_child_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_selection_get_n_selected_children
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) GetNSelectedChildren() (result int32, err error) {
	iv, err := _I.Get(109, "Selection", "get_n_selected_children")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_selection_get_selected_child
//
// [ selected_child_index ] trans: nothing
//
// [ result ] trans: everything
//
func (v *SelectionIfc) GetSelectedChild(selected_child_index int32) (result Accessible, err error) {
	iv, err := _I.Get(110, "Selection", "get_selected_child")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selected_child_index := gi.NewInt32Argument(selected_child_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_selected_child_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_selection_is_child_selected
//
// [ child_index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) IsChildSelected(child_index int32) (result bool, err error) {
	iv, err := _I.Get(111, "Selection", "is_child_selected")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_child_index := gi.NewInt32Argument(child_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_child_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_selection_select_all
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) SelectAll() (result bool, err error) {
	iv, err := _I.Get(112, "Selection", "select_all")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_selection_select_child
//
// [ child_index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) SelectChild(child_index int32) (result bool, err error) {
	iv, err := _I.Get(113, "Selection", "select_child")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_child_index := gi.NewInt32Argument(child_index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_child_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Object StateSet
type StateSet struct {
	g.Object
}

func WrapStateSet(p unsafe.Pointer) (r StateSet) { r.P = p; return }

type IStateSet interface{ P_StateSet() unsafe.Pointer }

func (v StateSet) P_StateSet() unsafe.Pointer { return v.P }
func StateSetGetType() gi.GType {
	ret := _I.GetGType(40, "StateSet")
	return ret
}

// atspi_state_set_new
//
// [ states ] trans: nothing
//
// [ result ] trans: everything
//
func NewStateSet(states int /*TODO_TYPE isPtr: true, tag: array*/) (result StateSet) {
	iv, err := _I.Get(114, "StateSet", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_states := gi.NewIntArgument(states) /*TODO*/
	args := []gi.Argument{arg_states}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atspi_state_set_add
//
// [ state ] trans: nothing
//
func (v StateSet) Add(state StateTypeEnum) {
	iv, err := _I.Get(115, "StateSet", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_v, arg_state}
	iv.Call(args, nil, nil)
}

// atspi_state_set_compare
//
// [ set2 ] trans: nothing
//
// [ result ] trans: everything
//
func (v StateSet) Compare(set2 IStateSet) (result StateSet) {
	iv, err := _I.Get(116, "StateSet", "compare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if set2 != nil {
		tmp = set2.P_StateSet()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_set2 := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_set2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atspi_state_set_contains
//
// [ state ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StateSet) Contains(state StateTypeEnum) (result bool) {
	iv, err := _I.Get(117, "StateSet", "contains")
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

// atspi_state_set_equals
//
// [ set2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StateSet) Equals(set2 IStateSet) (result bool) {
	iv, err := _I.Get(118, "StateSet", "equals")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if set2 != nil {
		tmp = set2.P_StateSet()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_set2 := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_set2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atspi_state_set_get_states
//
// [ result ] trans: everything
//
func (v StateSet) GetStates() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/) {
	iv, err := _I.Get(119, "StateSet", "get_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atspi_state_set_is_empty
//
// [ result ] trans: nothing
//
func (v StateSet) IsEmpty() (result bool) {
	iv, err := _I.Get(120, "StateSet", "is_empty")
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

// atspi_state_set_remove
//
// [ state ] trans: nothing
//
func (v StateSet) Remove(state StateTypeEnum) {
	iv, err := _I.Get(121, "StateSet", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_v, arg_state}
	iv.Call(args, nil, nil)
}

// atspi_state_set_set_by_name
//
// [ name ] trans: nothing
//
// [ enabled ] trans: nothing
//
func (v StateSet) SetByName(name string, enabled bool) {
	iv, err := _I.Get(122, "StateSet", "set_by_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_name, arg_enabled}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// ignore GType struct StateSetClass

// Enum StateType
type StateTypeEnum int

const (
	StateTypeInvalid                StateTypeEnum = 0
	StateTypeActive                 StateTypeEnum = 1
	StateTypeArmed                  StateTypeEnum = 2
	StateTypeBusy                   StateTypeEnum = 3
	StateTypeChecked                StateTypeEnum = 4
	StateTypeCollapsed              StateTypeEnum = 5
	StateTypeDefunct                StateTypeEnum = 6
	StateTypeEditable               StateTypeEnum = 7
	StateTypeEnabled                StateTypeEnum = 8
	StateTypeExpandable             StateTypeEnum = 9
	StateTypeExpanded               StateTypeEnum = 10
	StateTypeFocusable              StateTypeEnum = 11
	StateTypeFocused                StateTypeEnum = 12
	StateTypeHasTooltip             StateTypeEnum = 13
	StateTypeHorizontal             StateTypeEnum = 14
	StateTypeIconified              StateTypeEnum = 15
	StateTypeModal                  StateTypeEnum = 16
	StateTypeMultiLine              StateTypeEnum = 17
	StateTypeMultiselectable        StateTypeEnum = 18
	StateTypeOpaque                 StateTypeEnum = 19
	StateTypePressed                StateTypeEnum = 20
	StateTypeResizable              StateTypeEnum = 21
	StateTypeSelectable             StateTypeEnum = 22
	StateTypeSelected               StateTypeEnum = 23
	StateTypeSensitive              StateTypeEnum = 24
	StateTypeShowing                StateTypeEnum = 25
	StateTypeSingleLine             StateTypeEnum = 26
	StateTypeStale                  StateTypeEnum = 27
	StateTypeTransient              StateTypeEnum = 28
	StateTypeVertical               StateTypeEnum = 29
	StateTypeVisible                StateTypeEnum = 30
	StateTypeManagesDescendants     StateTypeEnum = 31
	StateTypeIndeterminate          StateTypeEnum = 32
	StateTypeRequired               StateTypeEnum = 33
	StateTypeTruncated              StateTypeEnum = 34
	StateTypeAnimated               StateTypeEnum = 35
	StateTypeInvalidEntry           StateTypeEnum = 36
	StateTypeSupportsAutocompletion StateTypeEnum = 37
	StateTypeSelectableText         StateTypeEnum = 38
	StateTypeIsDefault              StateTypeEnum = 39
	StateTypeVisited                StateTypeEnum = 40
	StateTypeCheckable              StateTypeEnum = 41
	StateTypeHasPopup               StateTypeEnum = 42
	StateTypeReadOnly               StateTypeEnum = 43
	StateTypeLastDefined            StateTypeEnum = 44
)

func StateTypeGetType() gi.GType {
	ret := _I.GetGType(41, "StateType")
	return ret
}

// Interface Table
type Table struct {
	TableIfc
	P unsafe.Pointer
}
type TableIfc struct{}
type ITable interface{ P_Table() unsafe.Pointer }

func (v Table) P_Table() unsafe.Pointer { return v.P }
func TableGetType() gi.GType {
	ret := _I.GetGType(42, "Table")
	return ret
}

// atspi_table_add_column_selection
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) AddColumnSelection(column int32) (result bool, err error) {
	iv, err := _I.Get(123, "Table", "add_column_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_table_add_row_selection
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) AddRowSelection(row int32) (result bool, err error) {
	iv, err := _I.Get(124, "Table", "add_row_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_table_get_accessible_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TableIfc) GetAccessibleAt(row int32, column int32) (result Accessible, err error) {
	iv, err := _I.Get(125, "Table", "get_accessible_at")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_table_get_caption
//
// [ result ] trans: everything
//
func (v *TableIfc) GetCaption() (result Accessible, err error) {
	iv, err := _I.Get(126, "Table", "get_caption")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_table_get_column_at_index
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetColumnAtIndex(index int32) (result int32, err error) {
	iv, err := _I.Get(127, "Table", "get_column_at_index")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index := gi.NewInt32Argument(index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_column_description
//
// [ column ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TableIfc) GetColumnDescription(column int32) (result string, err error) {
	iv, err := _I.Get(128, "Table", "get_column_description")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_table_get_column_extent_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetColumnExtentAt(row int32, column int32) (result int32, err error) {
	iv, err := _I.Get(129, "Table", "get_column_extent_at")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_column_header
//
// [ column ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TableIfc) GetColumnHeader(column int32) (result Accessible, err error) {
	iv, err := _I.Get(130, "Table", "get_column_header")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_table_get_index_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetIndexAt(row int32, column int32) (result int32, err error) {
	iv, err := _I.Get(131, "Table", "get_index_at")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_n_columns
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetNColumns() (result int32, err error) {
	iv, err := _I.Get(132, "Table", "get_n_columns")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_n_rows
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetNRows() (result int32, err error) {
	iv, err := _I.Get(133, "Table", "get_n_rows")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_n_selected_columns
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetNSelectedColumns() (result int32, err error) {
	iv, err := _I.Get(134, "Table", "get_n_selected_columns")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_n_selected_rows
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetNSelectedRows() (result int32, err error) {
	iv, err := _I.Get(135, "Table", "get_n_selected_rows")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_row_at_index
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowAtIndex(index int32) (result int32, err error) {
	iv, err := _I.Get(136, "Table", "get_row_at_index")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index := gi.NewInt32Argument(index)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_index, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_row_column_extents_at_index
//
// [ index ] trans: nothing
//
// [ row ] trans: everything, dir: out
//
// [ col ] trans: everything, dir: out
//
// [ row_extents ] trans: everything, dir: out
//
// [ col_extents ] trans: everything, dir: out
//
// [ is_selected ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowColumnExtentsAtIndex(index int32) (result bool, row int32, col int32, row_extents int32, col_extents int32, is_selected bool, err error) {
	iv, err := _I.Get(137, "Table", "get_row_column_extents_at_index")
	if err != nil {
		return
	}
	var outArgs [6]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index := gi.NewInt32Argument(index)
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_col := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_row_extents := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_col_extents := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_is_selected := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[5]))
	args := []gi.Argument{arg_v, arg_index, arg_row, arg_col, arg_row_extents, arg_col_extents, arg_is_selected, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[5].Pointer())
	row = outArgs[0].Int32()
	col = outArgs[1].Int32()
	row_extents = outArgs[2].Int32()
	col_extents = outArgs[3].Int32()
	is_selected = outArgs[4].Bool()
	result = ret.Bool()
	return
}

// atspi_table_get_row_description
//
// [ row ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TableIfc) GetRowDescription(row int32) (result string, err error) {
	iv, err := _I.Get(138, "Table", "get_row_description")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_table_get_row_extent_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowExtentAt(row int32, column int32) (result int32, err error) {
	iv, err := _I.Get(139, "Table", "get_row_extent_at")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_get_row_header
//
// [ row ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TableIfc) GetRowHeader(row int32) (result Accessible, err error) {
	iv, err := _I.Get(140, "Table", "get_row_header")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_table_get_selected_columns
//
// [ result ] trans: everything
//
func (v *TableIfc) GetSelectedColumns() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(141, "Table", "get_selected_columns")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_table_get_selected_rows
//
// [ result ] trans: everything
//
func (v *TableIfc) GetSelectedRows() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(142, "Table", "get_selected_rows")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_table_get_summary
//
// [ result ] trans: everything
//
func (v *TableIfc) GetSummary() (result Accessible, err error) {
	iv, err := _I.Get(143, "Table", "get_summary")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_table_is_column_selected
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) IsColumnSelected(column int32) (result bool, err error) {
	iv, err := _I.Get(144, "Table", "is_column_selected")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_table_is_row_selected
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) IsRowSelected(row int32) (result bool, err error) {
	iv, err := _I.Get(145, "Table", "is_row_selected")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_table_is_selected
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) IsSelected(row int32, column int32) (result bool, err error) {
	iv, err := _I.Get(146, "Table", "is_selected")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_table_remove_column_selection
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) RemoveColumnSelection(column int32) (result bool, err error) {
	iv, err := _I.Get(147, "Table", "remove_column_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_table_remove_row_selection
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) RemoveRowSelection(row int32) (result bool, err error) {
	iv, err := _I.Get(148, "Table", "remove_row_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_row, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Interface TableCell
type TableCell struct {
	TableCellIfc
	P unsafe.Pointer
}
type TableCellIfc struct{}
type ITableCell interface{ P_TableCell() unsafe.Pointer }

func (v TableCell) P_TableCell() unsafe.Pointer { return v.P }
func TableCellGetType() gi.GType {
	ret := _I.GetGType(43, "TableCell")
	return ret
}

// atspi_table_cell_get_column_header_cells
//
// [ result ] trans: everything
//
func (v *TableCellIfc) GetColumnHeaderCells() (result int /*TODO_TYPE array type: 2, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(149, "TableCell", "get_column_header_cells")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_table_cell_get_column_index
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetColumnIndex() (result int32, err error) {
	iv, err := _I.Get(150, "TableCell", "get_column_index")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_cell_get_column_span
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetColumnSpan() (result int32, err error) {
	iv, err := _I.Get(151, "TableCell", "get_column_span")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_cell_get_position
//
// [ row ] trans: everything, dir: out
//
// [ column ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetPosition() (result int32, row int32, column int32, err error) {
	iv, err := _I.Get(152, "TableCell", "get_position")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	row = outArgs[0].Int32()
	column = outArgs[1].Int32()
	result = ret.Int32()
	return
}

// atspi_table_cell_get_row_column_span
//
// [ row ] trans: everything, dir: out
//
// [ column ] trans: everything, dir: out
//
// [ row_span ] trans: everything, dir: out
//
// [ column_span ] trans: everything, dir: out
//
func (v *TableCellIfc) GetRowColumnSpan() (row int32, column int32, row_span int32, column_span int32, err error) {
	iv, err := _I.Get(153, "TableCell", "get_row_column_span")
	if err != nil {
		return
	}
	var outArgs [5]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_row_span := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_column_span := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_row_span, arg_column_span, arg_err}
	iv.Call(args, nil, &outArgs[0])
	err = gi.ToError(outArgs[4].Pointer())
	row = outArgs[0].Int32()
	column = outArgs[1].Int32()
	row_span = outArgs[2].Int32()
	column_span = outArgs[3].Int32()
	return
}

// atspi_table_cell_get_row_header_cells
//
// [ result ] trans: everything
//
func (v *TableCellIfc) GetRowHeaderCells() (result int /*TODO_TYPE array type: 2, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(154, "TableCell", "get_row_header_cells")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_table_cell_get_row_span
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetRowSpan() (result int32, err error) {
	iv, err := _I.Get(155, "TableCell", "get_row_span")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_table_cell_get_table
//
// [ result ] trans: everything
//
func (v *TableCellIfc) GetTable() (result Accessible, err error) {
	iv, err := _I.Get(156, "TableCell", "get_table")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// Interface Text
type Text struct {
	TextIfc
	P unsafe.Pointer
}
type TextIfc struct{}
type IText interface{ P_Text() unsafe.Pointer }

func (v Text) P_Text() unsafe.Pointer { return v.P }
func TextGetType() gi.GType {
	ret := _I.GetGType(44, "Text")
	return ret
}

// atspi_text_add_selection
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) AddSelection(start_offset int32, end_offset int32) (result bool, err error) {
	iv, err := _I.Get(157, "Text", "add_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_text_get_attribute_run
//
// [ offset ] trans: nothing
//
// [ include_defaults ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetAttributeRun(offset int32, include_defaults bool) (result g.HashTable, start_offset int32, end_offset int32, err error) {
	iv, err := _I.Get(158, "Text", "get_attribute_run")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_include_defaults := gi.NewBoolArgument(include_defaults)
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_offset, arg_include_defaults, arg_start_offset, arg_end_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// atspi_text_get_attribute_value
//
// [ offset ] trans: nothing
//
// [ attribute_name ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextAttributeValue(offset int32, attribute_name string) (result string, err error) {
	iv, err := _I.Get(159, "Text", "get_text_attribute_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_attribute_name := gi.CString(attribute_name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_attribute_name := gi.NewStringArgument(c_attribute_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_attribute_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_attribute_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// Deprecated
//
// atspi_text_get_attributes
//
// [ offset ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextAttributes(offset int32) (result g.HashTable, start_offset int32, end_offset int32, err error) {
	iv, err := _I.Get(160, "Text", "get_text_attributes")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_offset, arg_start_offset, arg_end_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// atspi_text_get_bounded_ranges
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ clipTypeX ] trans: nothing
//
// [ clipTypeY ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetBoundedRanges(x int32, y int32, width int32, height int32, type1 CoordTypeEnum, clipTypeX TextClipTypeEnum, clipTypeY TextClipTypeEnum) (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/, err error) {
	iv, err := _I.Get(161, "Text", "get_bounded_ranges")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_clipTypeX := gi.NewIntArgument(int(clipTypeX))
	arg_clipTypeY := gi.NewIntArgument(int(clipTypeY))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height, arg_type1, arg_clipTypeX, arg_clipTypeY, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int() /*TODO*/
	return
}

// atspi_text_get_caret_offset
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetCaretOffset() (result int32, err error) {
	iv, err := _I.Get(162, "Text", "get_caret_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_text_get_character_at_offset
//
// [ offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetCharacterAtOffset(offset int32) (result uint32, err error) {
	iv, err := _I.Get(163, "Text", "get_character_at_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Uint32()
	return
}

// atspi_text_get_character_count
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetCharacterCount() (result int32, err error) {
	iv, err := _I.Get(164, "Text", "get_character_count")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_text_get_character_extents
//
// [ offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetCharacterExtents(offset int32, type1 CoordTypeEnum) (result Rect, err error) {
	iv, err := _I.Get(165, "Text", "get_character_extents")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_get_default_attributes
//
// [ result ] trans: everything
//
func (v *TextIfc) GetDefaultAttributes() (result g.HashTable, err error) {
	iv, err := _I.Get(166, "Text", "get_default_attributes")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_get_n_selections
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetNSelections() (result int32, err error) {
	iv, err := _I.Get(167, "Text", "get_n_selections")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_text_get_offset_at_point
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetOffsetAtPoint(x int32, y int32, type1 CoordTypeEnum) (result int32, err error) {
	iv, err := _I.Get(168, "Text", "get_offset_at_point")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// atspi_text_get_range_extents
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetRangeExtents(start_offset int32, end_offset int32, type1 CoordTypeEnum) (result Rect, err error) {
	iv, err := _I.Get(169, "Text", "get_range_extents")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_get_selection
//
// [ selection_num ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetSelection(selection_num int32) (result Range, err error) {
	iv, err := _I.Get(170, "Text", "get_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_selection_num, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_get_string_at_offset
//
// [ offset ] trans: nothing
//
// [ granularity ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetStringAtOffset(offset int32, granularity TextGranularityEnum) (result TextRange, err error) {
	iv, err := _I.Get(171, "Text", "get_string_at_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_granularity := gi.NewIntArgument(int(granularity))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_granularity, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_get_text
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetText(start_offset int32, end_offset int32) (result string, err error) {
	iv, err := _I.Get(172, "Text", "get_text")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// atspi_text_get_text_after_offset
//
// [ offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextAfterOffset(offset int32, type1 TextBoundaryTypeEnum) (result TextRange, err error) {
	iv, err := _I.Get(173, "Text", "get_text_after_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// atspi_text_get_text_at_offset
//
// [ offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextAtOffset(offset int32, type1 TextBoundaryTypeEnum) (result TextRange, err error) {
	iv, err := _I.Get(174, "Text", "get_text_at_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_get_text_before_offset
//
// [ offset ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextBeforeOffset(offset int32, type1 TextBoundaryTypeEnum) (result TextRange, err error) {
	iv, err := _I.Get(175, "Text", "get_text_before_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// atspi_text_remove_selection
//
// [ selection_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) RemoveSelection(selection_num int32) (result bool, err error) {
	iv, err := _I.Get(176, "Text", "remove_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_selection_num, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_text_set_caret_offset
//
// [ new_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) SetCaretOffset(new_offset int32) (result bool, err error) {
	iv, err := _I.Get(177, "Text", "set_caret_offset")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_new_offset := gi.NewInt32Argument(new_offset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_new_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_text_set_selection
//
// [ selection_num ] trans: nothing
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) SetSelection(selection_num int32, start_offset int32, end_offset int32) (result bool, err error) {
	iv, err := _I.Get(178, "Text", "set_selection")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_selection_num, arg_start_offset, arg_end_offset, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Enum TextBoundaryType
type TextBoundaryTypeEnum int

const (
	TextBoundaryTypeChar          TextBoundaryTypeEnum = 0
	TextBoundaryTypeWordStart     TextBoundaryTypeEnum = 1
	TextBoundaryTypeWordEnd       TextBoundaryTypeEnum = 2
	TextBoundaryTypeSentenceStart TextBoundaryTypeEnum = 3
	TextBoundaryTypeSentenceEnd   TextBoundaryTypeEnum = 4
	TextBoundaryTypeLineStart     TextBoundaryTypeEnum = 5
	TextBoundaryTypeLineEnd       TextBoundaryTypeEnum = 6
)

func TextBoundaryTypeGetType() gi.GType {
	ret := _I.GetGType(45, "TextBoundaryType")
	return ret
}

// Enum TextClipType
type TextClipTypeEnum int

const (
	TextClipTypeNone TextClipTypeEnum = 0
	TextClipTypeMin  TextClipTypeEnum = 1
	TextClipTypeMax  TextClipTypeEnum = 2
	TextClipTypeBoth TextClipTypeEnum = 3
)

func TextClipTypeGetType() gi.GType {
	ret := _I.GetGType(46, "TextClipType")
	return ret
}

// Enum TextGranularity
type TextGranularityEnum int

const (
	TextGranularityChar      TextGranularityEnum = 0
	TextGranularityWord      TextGranularityEnum = 1
	TextGranularitySentence  TextGranularityEnum = 2
	TextGranularityLine      TextGranularityEnum = 3
	TextGranularityParagraph TextGranularityEnum = 4
)

func TextGranularityGetType() gi.GType {
	ret := _I.GetGType(47, "TextGranularity")
	return ret
}

// Struct TextRange
type TextRange struct {
	P unsafe.Pointer
}

const SizeOfStructTextRange = 16

func TextRangeGetType() gi.GType {
	ret := _I.GetGType(48, "TextRange")
	return ret
}

// Interface Value
type Value struct {
	ValueIfc
	P unsafe.Pointer
}
type ValueIfc struct{}
type IValue interface{ P_Value() unsafe.Pointer }

func (v Value) P_Value() unsafe.Pointer { return v.P }
func ValueGetType() gi.GType {
	ret := _I.GetGType(49, "Value")
	return ret
}

// atspi_value_get_current_value
//
// [ result ] trans: nothing
//
func (v *ValueIfc) GetCurrentValue() (result float64, err error) {
	iv, err := _I.Get(179, "Value", "get_current_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// atspi_value_get_maximum_value
//
// [ result ] trans: nothing
//
func (v *ValueIfc) GetMaximumValue() (result float64, err error) {
	iv, err := _I.Get(180, "Value", "get_maximum_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// atspi_value_get_minimum_increment
//
// [ result ] trans: nothing
//
func (v *ValueIfc) GetMinimumIncrement() (result float64, err error) {
	iv, err := _I.Get(181, "Value", "get_minimum_increment")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// atspi_value_get_minimum_value
//
// [ result ] trans: nothing
//
func (v *ValueIfc) GetMinimumValue() (result float64, err error) {
	iv, err := _I.Get(182, "Value", "get_minimum_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// atspi_value_set_current_value
//
// [ new_value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ValueIfc) SetCurrentValue(new_value float64) (result bool, err error) {
	iv, err := _I.Get(183, "Value", "set_current_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_new_value := gi.NewDoubleArgument(new_value)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_new_value, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_deregister_device_event_listener
//
// [ listener ] trans: nothing
//
// [ filter ] trans: nothing
//
// [ result ] trans: nothing
//
func DeregisterDeviceEventListener(listener IDeviceListener, filter unsafe.Pointer) (result bool, err error) {
	iv, err := _I.Get(184, "deregister_device_event_listener", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if listener != nil {
		tmp = listener.P_DeviceListener()
	}
	arg_listener := gi.NewPointerArgument(tmp)
	arg_filter := gi.NewPointerArgument(filter)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_listener, arg_filter, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_deregister_keystroke_listener
//
// [ listener ] trans: nothing
//
// [ key_set ] trans: nothing
//
// [ modmask ] trans: nothing
//
// [ event_types ] trans: nothing
//
// [ result ] trans: nothing
//
func DeregisterKeystrokeListener(listener IDeviceListener, key_set int /*TODO_TYPE isPtr: true, tag: array*/, modmask uint32, event_types uint32) (result bool, err error) {
	iv, err := _I.Get(185, "deregister_keystroke_listener", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if listener != nil {
		tmp = listener.P_DeviceListener()
	}
	arg_listener := gi.NewPointerArgument(tmp)
	arg_key_set := gi.NewIntArgument(key_set) /*TODO*/
	arg_modmask := gi.NewUint32Argument(modmask)
	arg_event_types := gi.NewUint32Argument(event_types)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_listener, arg_key_set, arg_modmask, arg_event_types, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_event_main
//
func EventMain() {
	iv, err := _I.Get(186, "event_main", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// atspi_event_quit
//
func EventQuit() {
	iv, err := _I.Get(187, "event_quit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// atspi_exit
//
// [ result ] trans: nothing
//
func Exit() (result int32) {
	iv, err := _I.Get(188, "exit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// atspi_generate_keyboard_event
//
// [ keyval ] trans: nothing
//
// [ keystring ] trans: nothing
//
// [ synth_type ] trans: nothing
//
// [ result ] trans: nothing
//
func GenerateKeyboardEvent(keyval int64, keystring string, synth_type KeySynthTypeEnum) (result bool, err error) {
	iv, err := _I.Get(189, "generate_keyboard_event", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_keystring := gi.CString(keystring)
	arg_keyval := gi.NewInt64Argument(keyval)
	arg_keystring := gi.NewStringArgument(c_keystring)
	arg_synth_type := gi.NewIntArgument(int(synth_type))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_keyval, arg_keystring, arg_synth_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_keystring)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_generate_mouse_event
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func GenerateMouseEvent(x int64, y int64, name string) (result bool, err error) {
	iv, err := _I.Get(190, "generate_mouse_event", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	arg_x := gi.NewInt64Argument(x)
	arg_y := gi.NewInt64Argument(y)
	arg_name := gi.NewStringArgument(c_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_x, arg_y, arg_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_get_desktop
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func GetDesktop(i int32) (result Accessible) {
	iv, err := _I.Get(191, "get_desktop", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atspi_get_desktop_count
//
// [ result ] trans: nothing
//
func GetDesktopCount() (result int32) {
	iv, err := _I.Get(192, "get_desktop_count", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// atspi_get_desktop_list
//
// [ result ] trans: everything
//
func GetDesktopList() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/) {
	iv, err := _I.Get(193, "get_desktop_list", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atspi_init
//
// [ result ] trans: nothing
//
func Init() (result int32) {
	iv, err := _I.Get(194, "init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// atspi_is_initialized
//
// [ result ] trans: nothing
//
func IsInitialized() (result bool) {
	iv, err := _I.Get(195, "is_initialized", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// atspi_register_device_event_listener
//
// [ listener ] trans: nothing
//
// [ event_types ] trans: nothing
//
// [ filter ] trans: nothing
//
// [ result ] trans: nothing
//
func RegisterDeviceEventListener(listener IDeviceListener, event_types uint32, filter unsafe.Pointer) (result bool, err error) {
	iv, err := _I.Get(196, "register_device_event_listener", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if listener != nil {
		tmp = listener.P_DeviceListener()
	}
	arg_listener := gi.NewPointerArgument(tmp)
	arg_event_types := gi.NewUint32Argument(event_types)
	arg_filter := gi.NewPointerArgument(filter)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_listener, arg_event_types, arg_filter, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_register_keystroke_listener
//
// [ listener ] trans: nothing
//
// [ key_set ] trans: nothing
//
// [ modmask ] trans: nothing
//
// [ event_types ] trans: nothing
//
// [ sync_type ] trans: nothing
//
// [ result ] trans: nothing
//
func RegisterKeystrokeListener(listener IDeviceListener, key_set int /*TODO_TYPE isPtr: true, tag: array*/, modmask uint32, event_types uint32, sync_type KeyListenerSyncTypeFlags) (result bool, err error) {
	iv, err := _I.Get(197, "register_keystroke_listener", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if listener != nil {
		tmp = listener.P_DeviceListener()
	}
	arg_listener := gi.NewPointerArgument(tmp)
	arg_key_set := gi.NewIntArgument(key_set) /*TODO*/
	arg_modmask := gi.NewUint32Argument(modmask)
	arg_event_types := gi.NewUint32Argument(event_types)
	arg_sync_type := gi.NewIntArgument(int(sync_type))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_listener, arg_key_set, arg_modmask, arg_event_types, arg_sync_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// atspi_role_get_name
//
// [ role ] trans: nothing
//
// [ result ] trans: everything
//
func RoleGetName(role RoleEnum) (result string) {
	iv, err := _I.Get(198, "role_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_role := gi.NewIntArgument(int(role))
	args := []gi.Argument{arg_role}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atspi_set_main_context
//
// [ cnx ] trans: nothing
//
func SetMainContext(cnx g.MainContext) {
	iv, err := _I.Get(199, "set_main_context", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cnx := gi.NewPointerArgument(cnx.P)
	args := []gi.Argument{arg_cnx}
	iv.Call(args, nil, nil)
}

// atspi_set_timeout
//
// [ val ] trans: nothing
//
// [ startup_time ] trans: nothing
//
func SetTimeout(val int32, startup_time int32) {
	iv, err := _I.Get(200, "set_timeout", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_val := gi.NewInt32Argument(val)
	arg_startup_time := gi.NewInt32Argument(startup_time)
	args := []gi.Argument{arg_val, arg_startup_time}
	iv.Call(args, nil, nil)
}

// constants
const (
	COMPONENTLAYER_COUNT                 = 9
	COORD_TYPE_COUNT                     = 3
	DBUS_INTERFACE_ACCESSIBLE            = "org.a11y.atspi.Accessible"
	DBUS_INTERFACE_ACTION                = "org.a11y.atspi.Action"
	DBUS_INTERFACE_APPLICATION           = "org.a11y.atspi.Application"
	DBUS_INTERFACE_CACHE                 = "org.a11y.atspi.Cache"
	DBUS_INTERFACE_COLLECTION            = "org.a11y.atspi.Collection"
	DBUS_INTERFACE_COMPONENT             = "org.a11y.atspi.Component"
	DBUS_INTERFACE_DEC                   = "org.a11y.atspi.DeviceEventController"
	DBUS_INTERFACE_DEVICE_EVENT_LISTENER = "org.a11y.atspi.DeviceEventListener"
	DBUS_INTERFACE_DOCUMENT              = "org.a11y.atspi.Document"
	DBUS_INTERFACE_EDITABLE_TEXT         = "org.a11y.atspi.EditableText"
	DBUS_INTERFACE_EVENT_KEYBOARD        = "org.a11y.atspi.Event.Keyboard"
	DBUS_INTERFACE_EVENT_MOUSE           = "org.a11y.atspi.Event.Mouse"
	DBUS_INTERFACE_EVENT_OBJECT          = "org.a11y.atspi.Event.Object"
	DBUS_INTERFACE_HYPERLINK             = "org.a11y.atspi.Hyperlink"
	DBUS_INTERFACE_HYPERTEXT             = "org.a11y.atspi.Hypertext"
	DBUS_INTERFACE_IMAGE                 = "org.a11y.atspi.Image"
	DBUS_INTERFACE_REGISTRY              = "org.a11y.atspi.Registry"
	DBUS_INTERFACE_SELECTION             = "org.a11y.atspi.Selection"
	DBUS_INTERFACE_SOCKET                = "org.a11y.atspi.Socket"
	DBUS_INTERFACE_TABLE                 = "org.a11y.atspi.Table"
	DBUS_INTERFACE_TABLE_CELL            = "org.a11y.atspi.TableCell"
	DBUS_INTERFACE_TEXT                  = "org.a11y.atspi.Text"
	DBUS_INTERFACE_VALUE                 = "org.a11y.atspi.Value"
	DBUS_NAME_REGISTRY                   = "org.a11y.atspi.Registry"
	DBUS_PATH_DEC                        = "/org/a11y/atspi/registry/deviceeventcontroller"
	DBUS_PATH_NULL                       = "/org/a11y/atspi/null"
	DBUS_PATH_REGISTRY                   = "/org/a11y/atspi/registry"
	DBUS_PATH_ROOT                       = "/org/a11y/atspi/accessible/root"
	EVENTTYPE_COUNT                      = 4
	KEYEVENTTYPE_COUNT                   = 2
	KEYSYNTHTYPE_COUNT                   = 5
	MATCHTYPES_COUNT                     = 6
	MODIFIERTYPE_COUNT                   = 8
	RELATIONTYPE_COUNT                   = 24
	ROLE_COUNT                           = 126
	SCROLLTYPE_COUNT                     = 7
	SORTORDER_COUNT                      = 8
	STATETYPE_COUNT                      = 42
	TEXT_BOUNDARY_TYPE_COUNT             = 7
	TEXT_CLIP_TYPE_COUNT                 = 4
)
