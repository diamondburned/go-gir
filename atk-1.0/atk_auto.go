package atk

/*
#cgo pkg-config: atk
#include <atk/atk.h>
extern void myAtkEventListener(AtkObject* obj);
static void* getPointer_myAtkEventListener() {
return (void*)(myAtkEventListener);
}
extern void myAtkEventListenerInit();
static void* getPointer_myAtkEventListenerInit() {
return (void*)(myAtkEventListenerInit);
}
extern void myAtkFocusHandler(AtkObject* object, gboolean focus_in);
static void* getPointer_myAtkFocusHandler() {
return (void*)(myAtkFocusHandler);
}
extern void myAtkFunction(gpointer user_data);
static void* getPointer_myAtkFunction() {
return (void*)(myAtkFunction);
}
extern void myAtkKeySnoopFunc(AtkKeyEventStruct* event, gpointer user_data);
static void* getPointer_myAtkKeySnoopFunc() {
return (void*)(myAtkKeySnoopFunc);
}
extern void myAtkPropertyChangeHandler(AtkObject* obj, AtkPropertyValues* vals);
static void* getPointer_myAtkPropertyChangeHandler() {
return (void*)(myAtkPropertyChangeHandler);
}
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Atk")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Atk", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
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
	ret := _I.GetGType(0, "Action")
	return ret
}

// atk_action_do_action
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) DoAction(i int32) (result bool) {
	iv, err := _I.Get(0, "Action", "do_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_action_get_description
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) GetDescription(i int32) (result string) {
	iv, err := _I.Get(1, "Action", "get_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_action_get_keybinding
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) GetKeybinding(i int32) (result string) {
	iv, err := _I.Get(2, "Action", "get_keybinding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_action_get_localized_name
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) GetLocalizedName(i int32) (result string) {
	iv, err := _I.Get(3, "Action", "get_localized_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_action_get_n_actions
//
// [ result ] trans: nothing
//
func (v *ActionIfc) GetNActions() (result int32) {
	iv, err := _I.Get(4, "Action", "get_n_actions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_action_get_name
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) GetName(i int32) (result string) {
	iv, err := _I.Get(5, "Action", "get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_action_set_description
//
// [ i ] trans: nothing
//
// [ desc ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ActionIfc) SetDescription(i int32, desc string) (result bool) {
	iv, err := _I.Get(6, "Action", "set_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_desc := gi.CString(desc)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_desc := gi.NewStringArgument(c_desc)
	args := []gi.Argument{arg_v, arg_i, arg_desc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_desc)
	result = ret.Bool()
	return
}

// ignore GType struct ActionIface

// Struct Attribute
type Attribute struct {
	P unsafe.Pointer
}

const SizeOfStructAttribute = 16

func AttributeGetType() gi.GType {
	ret := _I.GetGType(1, "Attribute")
	return ret
}

// atk_attribute_set_free
//
// [ attrib_set ] trans: nothing
//
func AttributeSetFree1(attrib_set g.SList) {
	iv, err := _I.Get(7, "Attribute", "set_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attrib_set := gi.NewPointerArgument(attrib_set.P)
	args := []gi.Argument{arg_attrib_set}
	iv.Call(args, nil, nil)
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
	ret := _I.GetGType(2, "Component")
	return ret
}

// atk_component_contains
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ coord_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) Contains(x int32, y int32, coord_type CoordTypeEnum) (result bool) {
	iv, err := _I.Get(8, "Component", "contains")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_get_alpha
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GetAlpha() (result float64) {
	iv, err := _I.Get(9, "Component", "get_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// atk_component_get_extents
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
// [ coord_type ] trans: nothing
//
func (v *ComponentIfc) GetExtents(coord_type CoordTypeEnum) (x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(10, "Component", "get_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height, arg_coord_type}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	width = outArgs[2].Int32()
	height = outArgs[3].Int32()
	return
}

// atk_component_get_layer
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GetLayer() (result LayerEnum) {
	iv, err := _I.Get(11, "Component", "get_layer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LayerEnum(ret.Int())
	return
}

// atk_component_get_mdi_zorder
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GetMdiZorder() (result int32) {
	iv, err := _I.Get(12, "Component", "get_mdi_zorder")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// atk_component_get_position
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ coord_type ] trans: nothing
//
func (v *ComponentIfc) GetPosition(coord_type CoordTypeEnum) (x int32, y int32) {
	iv, err := _I.Get(13, "Component", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	return
}

// Deprecated
//
// atk_component_get_size
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
func (v *ComponentIfc) GetSize() (width int32, height int32) {
	iv, err := _I.Get(14, "Component", "get_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// atk_component_grab_focus
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) GrabFocus() (result bool) {
	iv, err := _I.Get(15, "Component", "grab_focus")
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

// atk_component_ref_accessible_at_point
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ coord_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v *ComponentIfc) RefAccessibleAtPoint(x int32, y int32, coord_type CoordTypeEnum) (result Object) {
	iv, err := _I.Get(16, "Component", "ref_accessible_at_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// atk_component_remove_focus_handler
//
// [ handler_id ] trans: nothing
//
func (v *ComponentIfc) RemoveFocusHandler(handler_id uint32) {
	iv, err := _I.Get(17, "Component", "remove_focus_handler")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_handler_id := gi.NewUint32Argument(handler_id)
	args := []gi.Argument{arg_v, arg_handler_id}
	iv.Call(args, nil, nil)
}

// atk_component_scroll_to
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) ScrollTo(type1 ScrollTypeEnum) (result bool) {
	iv, err := _I.Get(18, "Component", "scroll_to")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_scroll_to_point
//
// [ coords ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) ScrollToPoint(coords CoordTypeEnum, x int32, y int32) (result bool) {
	iv, err := _I.Get(19, "Component", "scroll_to_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_coords := gi.NewIntArgument(int(coords))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_coords, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_set_extents
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ coord_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) SetExtents(x int32, y int32, width int32, height int32, coord_type CoordTypeEnum) (result bool) {
	iv, err := _I.Get(20, "Component", "set_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_set_position
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ coord_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) SetPosition(x int32, y int32, coord_type CoordTypeEnum) (result bool) {
	iv, err := _I.Get(21, "Component", "set_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_set_size
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ComponentIfc) SetSize(width int32, height int32) (result bool) {
	iv, err := _I.Get(22, "Component", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct ComponentIface

// Enum CoordType
type CoordTypeEnum int

const (
	CoordTypeScreen CoordTypeEnum = 0
	CoordTypeWindow CoordTypeEnum = 1
	CoordTypeParent CoordTypeEnum = 2
)

func CoordTypeGetType() gi.GType {
	ret := _I.GetGType(3, "CoordType")
	return ret
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
	ret := _I.GetGType(4, "Document")
	return ret
}

// atk_document_get_attribute_value
//
// [ attribute_name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetAttributeValue(attribute_name string) (result string) {
	iv, err := _I.Get(23, "Document", "get_attribute_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_attribute_name := gi.CString(attribute_name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attribute_name := gi.NewStringArgument(c_attribute_name)
	args := []gi.Argument{arg_v, arg_attribute_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_attribute_name)
	result = ret.String().Copy()
	return
}

// atk_document_get_attributes
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetAttributes() (result g.SList) {
	iv, err := _I.Get(24, "Document", "get_attributes")
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

// atk_document_get_current_page_number
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetCurrentPageNumber() (result int32) {
	iv, err := _I.Get(25, "Document", "get_current_page_number")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// atk_document_get_document
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetDocument() (result unsafe.Pointer) {
	iv, err := _I.Get(26, "Document", "get_document")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Deprecated
//
// atk_document_get_document_type
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetDocumentType() (result string) {
	iv, err := _I.Get(27, "Document", "get_document_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// Deprecated
//
// atk_document_get_locale
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetLocale() (result string) {
	iv, err := _I.Get(28, "Document", "get_locale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_document_get_page_count
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) GetPageCount() (result int32) {
	iv, err := _I.Get(29, "Document", "get_page_count")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_document_set_attribute_value
//
// [ attribute_name ] trans: nothing
//
// [ attribute_value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DocumentIfc) SetAttributeValue(attribute_name string, attribute_value string) (result bool) {
	iv, err := _I.Get(30, "Document", "set_attribute_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_attribute_name := gi.CString(attribute_name)
	c_attribute_value := gi.CString(attribute_value)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attribute_name := gi.NewStringArgument(c_attribute_name)
	arg_attribute_value := gi.NewStringArgument(c_attribute_value)
	args := []gi.Argument{arg_v, arg_attribute_name, arg_attribute_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_attribute_name)
	gi.Free(c_attribute_value)
	result = ret.Bool()
	return
}

// ignore GType struct DocumentIface

// Interface EditableText
type EditableText struct {
	EditableTextIfc
	P unsafe.Pointer
}
type EditableTextIfc struct{}
type IEditableText interface{ P_EditableText() unsafe.Pointer }

func (v EditableText) P_EditableText() unsafe.Pointer { return v.P }
func EditableTextGetType() gi.GType {
	ret := _I.GetGType(5, "EditableText")
	return ret
}

// atk_editable_text_copy_text
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
func (v *EditableTextIfc) CopyText(start_pos int32, end_pos int32) {
	iv, err := _I.Get(31, "EditableText", "copy_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos}
	iv.Call(args, nil, nil)
}

// atk_editable_text_cut_text
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
func (v *EditableTextIfc) CutText(start_pos int32, end_pos int32) {
	iv, err := _I.Get(32, "EditableText", "cut_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos}
	iv.Call(args, nil, nil)
}

// atk_editable_text_delete_text
//
// [ start_pos ] trans: nothing
//
// [ end_pos ] trans: nothing
//
func (v *EditableTextIfc) DeleteText(start_pos int32, end_pos int32) {
	iv, err := _I.Get(33, "EditableText", "delete_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos}
	iv.Call(args, nil, nil)
}

// atk_editable_text_insert_text
//
// [ string ] trans: nothing
//
// [ length ] trans: nothing
//
// [ position ] trans: nothing
//
func (v *EditableTextIfc) InsertText(string string, length int32, position int32) {
	iv, err := _I.Get(34, "EditableText", "insert_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_string := gi.NewStringArgument(c_string)
	arg_length := gi.NewInt32Argument(length)
	arg_position := gi.NewInt32Argument(position)
	args := []gi.Argument{arg_v, arg_string, arg_length, arg_position}
	iv.Call(args, nil, nil)
	gi.Free(c_string)
}

// atk_editable_text_paste_text
//
// [ position ] trans: nothing
//
func (v *EditableTextIfc) PasteText(position int32) {
	iv, err := _I.Get(35, "EditableText", "paste_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_position := gi.NewInt32Argument(position)
	args := []gi.Argument{arg_v, arg_position}
	iv.Call(args, nil, nil)
}

// atk_editable_text_set_run_attributes
//
// [ attrib_set ] trans: nothing
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *EditableTextIfc) SetRunAttributes(attrib_set g.SList, start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(36, "EditableText", "set_run_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attrib_set := gi.NewPointerArgument(attrib_set.P)
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_attrib_set, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_editable_text_set_text_contents
//
// [ string ] trans: nothing
//
func (v *EditableTextIfc) SetTextContents(string string) {
	iv, err := _I.Get(37, "EditableText", "set_text_contents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_string}
	iv.Call(args, nil, nil)
	gi.Free(c_string)
}

// ignore GType struct EditableTextIface

type EventListenerStruct struct {
	F_obj Object
}

func GetPointer_myEventListener() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtkEventListener())
}

//export myAtkEventListener
func myAtkEventListener(obj *C.AtkObject) {
	// TODO: not found user_data
}

type EventListenerInitStruct struct {
}

func GetPointer_myEventListenerInit() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtkEventListenerInit())
}

//export myAtkEventListenerInit
func myAtkEventListenerInit() {
	// TODO: not found user_data
}

type FocusHandlerStruct struct {
	F_object   Object
	F_focus_in bool
}

func GetPointer_myFocusHandler() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtkFocusHandler())
}

//export myAtkFocusHandler
func myAtkFocusHandler(object *C.AtkObject, focus_in C.gboolean) {
	// TODO: not found user_data
}

type FunctionStruct struct {
}

func GetPointer_myFunction() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtkFunction())
}

//export myAtkFunction
func myAtkFunction(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &FunctionStruct{}
	fn(args)
}

// Object GObjectAccessible
type GObjectAccessible struct {
	Object
}

func WrapGObjectAccessible(p unsafe.Pointer) (r GObjectAccessible) { r.P = p; return }

type IGObjectAccessible interface{ P_GObjectAccessible() unsafe.Pointer }

func (v GObjectAccessible) P_GObjectAccessible() unsafe.Pointer { return v.P }
func GObjectAccessibleGetType() gi.GType {
	ret := _I.GetGType(6, "GObjectAccessible")
	return ret
}

// atk_gobject_accessible_for_object
//
// [ obj ] trans: nothing
//
// [ result ] trans: nothing
//
func GObjectAccessibleForObject1(obj g.IObject) (result Object) {
	iv, err := _I.Get(38, "GObjectAccessible", "for_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	arg_obj := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_gobject_accessible_get_object
//
// [ result ] trans: nothing
//
func (v GObjectAccessible) GetObject() (result g.Object) {
	iv, err := _I.Get(39, "GObjectAccessible", "get_object")
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

// ignore GType struct GObjectAccessibleClass

// Object Hyperlink
type Hyperlink struct {
	ActionIfc
	g.Object
}

func WrapHyperlink(p unsafe.Pointer) (r Hyperlink) { r.P = p; return }

type IHyperlink interface{ P_Hyperlink() unsafe.Pointer }

func (v Hyperlink) P_Hyperlink() unsafe.Pointer { return v.P }
func (v Hyperlink) P_Action() unsafe.Pointer    { return v.P }
func HyperlinkGetType() gi.GType {
	ret := _I.GetGType(7, "Hyperlink")
	return ret
}

// atk_hyperlink_get_end_index
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetEndIndex() (result int32) {
	iv, err := _I.Get(40, "Hyperlink", "get_end_index")
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

// atk_hyperlink_get_n_anchors
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetNAnchors() (result int32) {
	iv, err := _I.Get(41, "Hyperlink", "get_n_anchors")
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

// atk_hyperlink_get_object
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetObject(i int32) (result Object) {
	iv, err := _I.Get(42, "Hyperlink", "get_object")
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

// atk_hyperlink_get_start_index
//
// [ result ] trans: nothing
//
func (v Hyperlink) GetStartIndex() (result int32) {
	iv, err := _I.Get(43, "Hyperlink", "get_start_index")
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

// atk_hyperlink_get_uri
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v Hyperlink) GetUri(i int32) (result string) {
	iv, err := _I.Get(44, "Hyperlink", "get_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_hyperlink_is_inline
//
// [ result ] trans: nothing
//
func (v Hyperlink) IsInline() (result bool) {
	iv, err := _I.Get(45, "Hyperlink", "is_inline")
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
// atk_hyperlink_is_selected_link
//
// [ result ] trans: nothing
//
func (v Hyperlink) IsSelectedLink() (result bool) {
	iv, err := _I.Get(46, "Hyperlink", "is_selected_link")
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

// atk_hyperlink_is_valid
//
// [ result ] trans: nothing
//
func (v Hyperlink) IsValid() (result bool) {
	iv, err := _I.Get(47, "Hyperlink", "is_valid")
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

// ignore GType struct HyperlinkClass

// Interface HyperlinkImpl
type HyperlinkImpl struct {
	HyperlinkImplIfc
	P unsafe.Pointer
}
type HyperlinkImplIfc struct{}
type IHyperlinkImpl interface{ P_HyperlinkImpl() unsafe.Pointer }

func (v HyperlinkImpl) P_HyperlinkImpl() unsafe.Pointer { return v.P }
func HyperlinkImplGetType() gi.GType {
	ret := _I.GetGType(8, "HyperlinkImpl")
	return ret
}

// atk_hyperlink_impl_get_hyperlink
//
// [ result ] trans: everything
//
func (v *HyperlinkImplIfc) GetHyperlink() (result Hyperlink) {
	iv, err := _I.Get(48, "HyperlinkImpl", "get_hyperlink")
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

// ignore GType struct HyperlinkImplIface

// Flags HyperlinkStateFlags
type HyperlinkStateFlags int

const (
	HyperlinkStateFlagsInline HyperlinkStateFlags = 1
)

func HyperlinkStateFlagsGetType() gi.GType {
	ret := _I.GetGType(9, "HyperlinkStateFlags")
	return ret
}

// Interface Hypertext
type Hypertext struct {
	HypertextIfc
	P unsafe.Pointer
}
type HypertextIfc struct{}
type IHypertext interface{ P_Hypertext() unsafe.Pointer }

func (v Hypertext) P_Hypertext() unsafe.Pointer { return v.P }
func HypertextGetType() gi.GType {
	ret := _I.GetGType(10, "Hypertext")
	return ret
}

// atk_hypertext_get_link
//
// [ link_index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *HypertextIfc) GetLink(link_index int32) (result Hyperlink) {
	iv, err := _I.Get(49, "Hypertext", "get_link")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_link_index := gi.NewInt32Argument(link_index)
	args := []gi.Argument{arg_v, arg_link_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_hypertext_get_link_index
//
// [ char_index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *HypertextIfc) GetLinkIndex(char_index int32) (result int32) {
	iv, err := _I.Get(50, "Hypertext", "get_link_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_char_index := gi.NewInt32Argument(char_index)
	args := []gi.Argument{arg_v, arg_char_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_hypertext_get_n_links
//
// [ result ] trans: nothing
//
func (v *HypertextIfc) GetNLinks() (result int32) {
	iv, err := _I.Get(51, "Hypertext", "get_n_links")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// ignore GType struct HypertextIface

// Interface Image
type Image struct {
	ImageIfc
	P unsafe.Pointer
}
type ImageIfc struct{}
type IImage interface{ P_Image() unsafe.Pointer }

func (v Image) P_Image() unsafe.Pointer { return v.P }
func ImageGetType() gi.GType {
	ret := _I.GetGType(11, "Image")
	return ret
}

// atk_image_get_image_description
//
// [ result ] trans: nothing
//
func (v *ImageIfc) GetImageDescription() (result string) {
	iv, err := _I.Get(52, "Image", "get_image_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_image_get_image_locale
//
// [ result ] trans: nothing
//
func (v *ImageIfc) GetImageLocale() (result string) {
	iv, err := _I.Get(53, "Image", "get_image_locale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_image_get_image_position
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ coord_type ] trans: nothing
//
func (v *ImageIfc) GetImagePosition(coord_type CoordTypeEnum) (x int32, y int32) {
	iv, err := _I.Get(54, "Image", "get_image_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	return
}

// atk_image_get_image_size
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
func (v *ImageIfc) GetImageSize() (width int32, height int32) {
	iv, err := _I.Get(55, "Image", "get_image_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// atk_image_set_image_description
//
// [ description ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ImageIfc) SetImageDescription(description string) (result bool) {
	iv, err := _I.Get(56, "Image", "set_image_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_description}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_description)
	result = ret.Bool()
	return
}

// ignore GType struct ImageIface

// Struct Implementor
type Implementor struct {
	P unsafe.Pointer
}

func ImplementorGetType() gi.GType {
	ret := _I.GetGType(12, "Implementor")
	return ret
}

// atk_implementor_ref_accessible
//
// [ result ] trans: everything
//
func (v Implementor) RefAccessible() (result Object) {
	iv, err := _I.Get(57, "Implementor", "ref_accessible")
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

// Interface ImplementorIface
type ImplementorIface struct {
	ImplementorIfaceIfc
	P unsafe.Pointer
}
type ImplementorIfaceIfc struct{}
type IImplementorIface interface{ P_ImplementorIface() unsafe.Pointer }

func (v ImplementorIface) P_ImplementorIface() unsafe.Pointer { return v.P }
func ImplementorIfaceGetType() gi.GType {
	ret := _I.GetGType(13, "ImplementorIface")
	return ret
}

// Struct KeyEventStruct
type KeyEventStruct struct {
	P unsafe.Pointer
}

const SizeOfStructKeyEventStruct = 32

func KeyEventStructGetType() gi.GType {
	ret := _I.GetGType(14, "KeyEventStruct")
	return ret
}

// Enum KeyEventType
type KeyEventTypeEnum int

const (
	KeyEventTypePress       KeyEventTypeEnum = 0
	KeyEventTypeRelease     KeyEventTypeEnum = 1
	KeyEventTypeLastDefined KeyEventTypeEnum = 2
)

func KeyEventTypeGetType() gi.GType {
	ret := _I.GetGType(15, "KeyEventType")
	return ret
}

type KeySnoopFuncStruct struct {
	F_event KeyEventStruct
}

func GetPointer_myKeySnoopFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtkKeySnoopFunc())
}

//export myAtkKeySnoopFunc
func myAtkKeySnoopFunc(event *C.AtkKeyEventStruct, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &KeySnoopFuncStruct{
		F_event: KeyEventStruct{P: unsafe.Pointer(event)},
	}
	fn(args)
}

// Enum Layer
type LayerEnum int

const (
	LayerInvalid    LayerEnum = 0
	LayerBackground LayerEnum = 1
	LayerCanvas     LayerEnum = 2
	LayerWidget     LayerEnum = 3
	LayerMdi        LayerEnum = 4
	LayerPopup      LayerEnum = 5
	LayerOverlay    LayerEnum = 6
	LayerWindow     LayerEnum = 7
)

func LayerGetType() gi.GType {
	ret := _I.GetGType(16, "Layer")
	return ret
}

// Object Misc
type Misc struct {
	g.Object
}

func WrapMisc(p unsafe.Pointer) (r Misc) { r.P = p; return }

type IMisc interface{ P_Misc() unsafe.Pointer }

func (v Misc) P_Misc() unsafe.Pointer { return v.P }
func MiscGetType() gi.GType {
	ret := _I.GetGType(17, "Misc")
	return ret
}

// Deprecated
//
// Deprecated
//
// atk_misc_threads_enter
//
func (v Misc) ThreadsEnter() {
	iv, err := _I.Get(59, "Misc", "threads_enter")
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
// atk_misc_threads_leave
//
func (v Misc) ThreadsLeave() {
	iv, err := _I.Get(60, "Misc", "threads_leave")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct MiscClass

// Object NoOpObject
type NoOpObject struct {
	ActionIfc
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
	WindowIfc
	Object
}

func WrapNoOpObject(p unsafe.Pointer) (r NoOpObject) { r.P = p; return }

type INoOpObject interface{ P_NoOpObject() unsafe.Pointer }

func (v NoOpObject) P_NoOpObject() unsafe.Pointer   { return v.P }
func (v NoOpObject) P_Action() unsafe.Pointer       { return v.P }
func (v NoOpObject) P_Component() unsafe.Pointer    { return v.P }
func (v NoOpObject) P_Document() unsafe.Pointer     { return v.P }
func (v NoOpObject) P_EditableText() unsafe.Pointer { return v.P }
func (v NoOpObject) P_Hypertext() unsafe.Pointer    { return v.P }
func (v NoOpObject) P_Image() unsafe.Pointer        { return v.P }
func (v NoOpObject) P_Selection() unsafe.Pointer    { return v.P }
func (v NoOpObject) P_Table() unsafe.Pointer        { return v.P }
func (v NoOpObject) P_TableCell() unsafe.Pointer    { return v.P }
func (v NoOpObject) P_Text() unsafe.Pointer         { return v.P }
func (v NoOpObject) P_Value() unsafe.Pointer        { return v.P }
func (v NoOpObject) P_Window() unsafe.Pointer       { return v.P }
func NoOpObjectGetType() gi.GType {
	ret := _I.GetGType(18, "NoOpObject")
	return ret
}

// atk_no_op_object_new
//
// [ obj ] trans: nothing
//
// [ result ] trans: everything
//
func NewNoOpObject(obj g.IObject) (result NoOpObject) {
	iv, err := _I.Get(61, "NoOpObject", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	arg_obj := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NoOpObjectClass

// Object NoOpObjectFactory
type NoOpObjectFactory struct {
	ObjectFactory
}

func WrapNoOpObjectFactory(p unsafe.Pointer) (r NoOpObjectFactory) { r.P = p; return }

type INoOpObjectFactory interface{ P_NoOpObjectFactory() unsafe.Pointer }

func (v NoOpObjectFactory) P_NoOpObjectFactory() unsafe.Pointer { return v.P }
func NoOpObjectFactoryGetType() gi.GType {
	ret := _I.GetGType(19, "NoOpObjectFactory")
	return ret
}

// atk_no_op_object_factory_new
//
// [ result ] trans: everything
//
func NewNoOpObjectFactory() (result NoOpObjectFactory) {
	iv, err := _I.Get(62, "NoOpObjectFactory", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NoOpObjectFactoryClass

// Object Object
type Object struct {
	g.Object
}

func WrapObject(p unsafe.Pointer) (r Object) { r.P = p; return }

type IObject interface{ P_Object() unsafe.Pointer }

func (v Object) P_Object() unsafe.Pointer { return v.P }
func ObjectGetType() gi.GType {
	ret := _I.GetGType(20, "Object")
	return ret
}

// atk_object_add_relationship
//
// [ relationship ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) AddRelationship(relationship RelationTypeEnum, target IObject) (result bool) {
	iv, err := _I.Get(63, "Object", "add_relationship")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_object_get_attributes
//
// [ result ] trans: everything
//
func (v Object) GetAttributes() (result g.SList) {
	iv, err := _I.Get(64, "Object", "get_attributes")
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

// atk_object_get_description
//
// [ result ] trans: nothing
//
func (v Object) GetDescription() (result string) {
	iv, err := _I.Get(65, "Object", "get_description")
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

// atk_object_get_index_in_parent
//
// [ result ] trans: nothing
//
func (v Object) GetIndexInParent() (result int32) {
	iv, err := _I.Get(66, "Object", "get_index_in_parent")
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

// Deprecated
//
// atk_object_get_layer
//
// [ result ] trans: nothing
//
func (v Object) GetLayer() (result LayerEnum) {
	iv, err := _I.Get(67, "Object", "get_layer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LayerEnum(ret.Int())
	return
}

// Deprecated
//
// atk_object_get_mdi_zorder
//
// [ result ] trans: nothing
//
func (v Object) GetMdiZorder() (result int32) {
	iv, err := _I.Get(68, "Object", "get_mdi_zorder")
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

// atk_object_get_n_accessible_children
//
// [ result ] trans: nothing
//
func (v Object) GetNAccessibleChildren() (result int32) {
	iv, err := _I.Get(69, "Object", "get_n_accessible_children")
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

// atk_object_get_name
//
// [ result ] trans: nothing
//
func (v Object) GetName() (result string) {
	iv, err := _I.Get(70, "Object", "get_name")
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

// atk_object_get_object_locale
//
// [ result ] trans: nothing
//
func (v Object) GetObjectLocale() (result string) {
	iv, err := _I.Get(71, "Object", "get_object_locale")
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

// atk_object_get_parent
//
// [ result ] trans: nothing
//
func (v Object) GetParent() (result Object) {
	iv, err := _I.Get(72, "Object", "get_parent")
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

// atk_object_get_role
//
// [ result ] trans: nothing
//
func (v Object) GetRole() (result RoleEnum) {
	iv, err := _I.Get(73, "Object", "get_role")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = RoleEnum(ret.Int())
	return
}

// atk_object_initialize
//
// [ data ] trans: nothing
//
func (v Object) Initialize(data unsafe.Pointer) {
	iv, err := _I.Get(74, "Object", "initialize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// atk_object_notify_state_change
//
// [ state ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Object) NotifyStateChange(state uint64, value bool) {
	iv, err := _I.Get(75, "Object", "notify_state_change")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewUint64Argument(state)
	arg_value := gi.NewBoolArgument(value)
	args := []gi.Argument{arg_v, arg_state, arg_value}
	iv.Call(args, nil, nil)
}

// atk_object_peek_parent
//
// [ result ] trans: nothing
//
func (v Object) PeekParent() (result Object) {
	iv, err := _I.Get(76, "Object", "peek_parent")
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

// atk_object_ref_accessible_child
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v Object) RefAccessibleChild(i int32) (result Object) {
	iv, err := _I.Get(77, "Object", "ref_accessible_child")
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

// atk_object_ref_relation_set
//
// [ result ] trans: everything
//
func (v Object) RefRelationSet() (result RelationSet) {
	iv, err := _I.Get(78, "Object", "ref_relation_set")
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

// atk_object_ref_state_set
//
// [ result ] trans: everything
//
func (v Object) RefStateSet() (result StateSet) {
	iv, err := _I.Get(79, "Object", "ref_state_set")
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
// atk_object_remove_property_change_handler
//
// [ handler_id ] trans: nothing
//
func (v Object) RemovePropertyChangeHandler(handler_id uint32) {
	iv, err := _I.Get(80, "Object", "remove_property_change_handler")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_handler_id := gi.NewUint32Argument(handler_id)
	args := []gi.Argument{arg_v, arg_handler_id}
	iv.Call(args, nil, nil)
}

// atk_object_remove_relationship
//
// [ relationship ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) RemoveRelationship(relationship RelationTypeEnum, target IObject) (result bool) {
	iv, err := _I.Get(81, "Object", "remove_relationship")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_object_set_description
//
// [ description ] trans: nothing
//
func (v Object) SetDescription(description string) {
	iv, err := _I.Get(82, "Object", "set_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(v.P)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// atk_object_set_name
//
// [ name ] trans: nothing
//
func (v Object) SetName(name string) {
	iv, err := _I.Get(83, "Object", "set_name")
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

// atk_object_set_parent
//
// [ parent ] trans: nothing
//
func (v Object) SetParent(parent IObject) {
	iv, err := _I.Get(84, "Object", "set_parent")
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
	iv.Call(args, nil, nil)
}

// atk_object_set_role
//
// [ role ] trans: nothing
//
func (v Object) SetRole(role RoleEnum) {
	iv, err := _I.Get(85, "Object", "set_role")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_role := gi.NewIntArgument(int(role))
	args := []gi.Argument{arg_v, arg_role}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectClass

// Object ObjectFactory
type ObjectFactory struct {
	g.Object
}

func WrapObjectFactory(p unsafe.Pointer) (r ObjectFactory) { r.P = p; return }

type IObjectFactory interface{ P_ObjectFactory() unsafe.Pointer }

func (v ObjectFactory) P_ObjectFactory() unsafe.Pointer { return v.P }
func ObjectFactoryGetType() gi.GType {
	ret := _I.GetGType(21, "ObjectFactory")
	return ret
}

// atk_object_factory_create_accessible
//
// [ obj ] trans: nothing
//
// [ result ] trans: everything
//
func (v ObjectFactory) CreateAccessible(obj g.IObject) (result Object) {
	iv, err := _I.Get(86, "ObjectFactory", "create_accessible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if obj != nil {
		tmp = obj.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_obj := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_object_factory_get_accessible_type
//
// [ result ] trans: nothing
//
func (v ObjectFactory) GetAccessibleType() (result gi.GType) {
	iv, err := _I.Get(87, "ObjectFactory", "get_accessible_type")
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

// atk_object_factory_invalidate
//
func (v ObjectFactory) Invalidate() {
	iv, err := _I.Get(88, "ObjectFactory", "invalidate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectFactoryClass

// Object Plug
type Plug struct {
	ComponentIfc
	Object
}

func WrapPlug(p unsafe.Pointer) (r Plug) { r.P = p; return }

type IPlug interface{ P_Plug() unsafe.Pointer }

func (v Plug) P_Plug() unsafe.Pointer      { return v.P }
func (v Plug) P_Component() unsafe.Pointer { return v.P }
func PlugGetType() gi.GType {
	ret := _I.GetGType(22, "Plug")
	return ret
}

// atk_plug_new
//
// [ result ] trans: everything
//
func NewPlug() (result Plug) {
	iv, err := _I.Get(89, "Plug", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_plug_get_id
//
// [ result ] trans: everything
//
func (v Plug) GetId() (result string) {
	iv, err := _I.Get(90, "Plug", "get_id")
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

// ignore GType struct PlugClass

type PropertyChangeHandlerStruct struct {
	F_obj  Object
	F_vals PropertyValues
}

func GetPointer_myPropertyChangeHandler() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myAtkPropertyChangeHandler())
}

//export myAtkPropertyChangeHandler
func myAtkPropertyChangeHandler(obj *C.AtkObject, vals *C.AtkPropertyValues) {
	// TODO: not found user_data
}

// Struct PropertyValues
type PropertyValues struct {
	P unsafe.Pointer
}

const SizeOfStructPropertyValues = 56

func PropertyValuesGetType() gi.GType {
	ret := _I.GetGType(23, "PropertyValues")
	return ret
}

// Struct Range
type Range struct {
	P unsafe.Pointer
}

func RangeGetType() gi.GType {
	ret := _I.GetGType(24, "Range")
	return ret
}

// atk_range_new
//
// [ lower_limit ] trans: nothing
//
// [ upper_limit ] trans: nothing
//
// [ description ] trans: nothing
//
// [ result ] trans: everything
//
func NewRange(lower_limit float64, upper_limit float64, description string) (result Range) {
	iv, err := _I.Get(91, "Range", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_lower_limit := gi.NewDoubleArgument(lower_limit)
	arg_upper_limit := gi.NewDoubleArgument(upper_limit)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_lower_limit, arg_upper_limit, arg_description}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_description)
	result.P = ret.Pointer()
	return
}

// atk_range_copy
//
// [ result ] trans: everything
//
func (v Range) Copy() (result Range) {
	iv, err := _I.Get(92, "Range", "copy")
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

// atk_range_free
//
func (v Range) Free() {
	iv, err := _I.Get(93, "Range", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_range_get_description
//
// [ result ] trans: nothing
//
func (v Range) GetDescription() (result string) {
	iv, err := _I.Get(94, "Range", "get_description")
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

// atk_range_get_lower_limit
//
// [ result ] trans: nothing
//
func (v Range) GetLowerLimit() (result float64) {
	iv, err := _I.Get(95, "Range", "get_lower_limit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// atk_range_get_upper_limit
//
// [ result ] trans: nothing
//
func (v Range) GetUpperLimit() (result float64) {
	iv, err := _I.Get(96, "Range", "get_upper_limit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// Struct Rectangle
type Rectangle struct {
	P unsafe.Pointer
}

const SizeOfStructRectangle = 16

func RectangleGetType() gi.GType {
	ret := _I.GetGType(25, "Rectangle")
	return ret
}

// Object Registry
type Registry struct {
	g.Object
}

func WrapRegistry(p unsafe.Pointer) (r Registry) { r.P = p; return }

type IRegistry interface{ P_Registry() unsafe.Pointer }

func (v Registry) P_Registry() unsafe.Pointer { return v.P }
func RegistryGetType() gi.GType {
	ret := _I.GetGType(26, "Registry")
	return ret
}

// atk_registry_get_factory
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Registry) GetFactory(type1 gi.GType) (result ObjectFactory) {
	iv, err := _I.Get(97, "Registry", "get_factory")
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

// atk_registry_get_factory_type
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Registry) GetFactoryType(type1 gi.GType) (result gi.GType) {
	iv, err := _I.Get(98, "Registry", "get_factory_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// atk_registry_set_factory_type
//
// [ type1 ] trans: nothing
//
// [ factory_type ] trans: nothing
//
func (v Registry) SetFactoryType(type1 gi.GType, factory_type gi.GType) {
	iv, err := _I.Get(99, "Registry", "set_factory_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_factory_type := gi.NewUintArgument(uint(factory_type))
	args := []gi.Argument{arg_v, arg_type1, arg_factory_type}
	iv.Call(args, nil, nil)
}

// ignore GType struct RegistryClass

// Object Relation
type Relation struct {
	g.Object
}

func WrapRelation(p unsafe.Pointer) (r Relation) { r.P = p; return }

type IRelation interface{ P_Relation() unsafe.Pointer }

func (v Relation) P_Relation() unsafe.Pointer { return v.P }
func RelationGetType() gi.GType {
	ret := _I.GetGType(27, "Relation")
	return ret
}

// atk_relation_new
//
// [ targets ] trans: nothing
//
// [ n_targets ] trans: nothing
//
// [ relationship ] trans: nothing
//
// [ result ] trans: everything
//
func NewRelation(targets gi.PointerArray, n_targets int32, relationship RelationTypeEnum) (result Relation) {
	iv, err := _I.Get(100, "Relation", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_targets := gi.NewPointerArgument(targets.P)
	arg_n_targets := gi.NewInt32Argument(n_targets)
	arg_relationship := gi.NewIntArgument(int(relationship))
	args := []gi.Argument{arg_targets, arg_n_targets, arg_relationship}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_add_target
//
// [ target ] trans: nothing
//
func (v Relation) AddTarget(target IObject) {
	iv, err := _I.Get(101, "Relation", "add_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_target}
	iv.Call(args, nil, nil)
}

// atk_relation_get_relation_type
//
// [ result ] trans: nothing
//
func (v Relation) GetRelationType() (result RelationTypeEnum) {
	iv, err := _I.Get(102, "Relation", "get_relation_type")
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

// atk_relation_get_target
//
// [ result ] trans: nothing
//
func (v Relation) GetTarget() (result int /*TODO_TYPE array type: 2, isZeroTerm: false*/) {
	iv, err := _I.Get(103, "Relation", "get_target")
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

// atk_relation_remove_target
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Relation) RemoveTarget(target IObject) (result bool) {
	iv, err := _I.Get(104, "Relation", "remove_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct RelationClass

// Object RelationSet
type RelationSet struct {
	g.Object
}

func WrapRelationSet(p unsafe.Pointer) (r RelationSet) { r.P = p; return }

type IRelationSet interface{ P_RelationSet() unsafe.Pointer }

func (v RelationSet) P_RelationSet() unsafe.Pointer { return v.P }
func RelationSetGetType() gi.GType {
	ret := _I.GetGType(28, "RelationSet")
	return ret
}

// atk_relation_set_new
//
// [ result ] trans: everything
//
func NewRelationSet() (result RelationSet) {
	iv, err := _I.Get(105, "RelationSet", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_set_add
//
// [ relation ] trans: nothing
//
func (v RelationSet) Add(relation IRelation) {
	iv, err := _I.Get(106, "RelationSet", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if relation != nil {
		tmp = relation.P_Relation()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_relation}
	iv.Call(args, nil, nil)
}

// atk_relation_set_add_relation_by_type
//
// [ relationship ] trans: nothing
//
// [ target ] trans: nothing
//
func (v RelationSet) AddRelationByType(relationship RelationTypeEnum, target IObject) {
	iv, err := _I.Get(107, "RelationSet", "add_relation_by_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	iv.Call(args, nil, nil)
}

// atk_relation_set_contains
//
// [ relationship ] trans: nothing
//
// [ result ] trans: nothing
//
func (v RelationSet) Contains(relationship RelationTypeEnum) (result bool) {
	iv, err := _I.Get(108, "RelationSet", "contains")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	args := []gi.Argument{arg_v, arg_relationship}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_relation_set_contains_target
//
// [ relationship ] trans: nothing
//
// [ target ] trans: nothing
//
// [ result ] trans: nothing
//
func (v RelationSet) ContainsTarget(relationship RelationTypeEnum, target IObject) (result bool) {
	iv, err := _I.Get(109, "RelationSet", "contains_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_relation_set_get_n_relations
//
// [ result ] trans: nothing
//
func (v RelationSet) GetNRelations() (result int32) {
	iv, err := _I.Get(110, "RelationSet", "get_n_relations")
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

// atk_relation_set_get_relation
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v RelationSet) GetRelation(i int32) (result Relation) {
	iv, err := _I.Get(111, "RelationSet", "get_relation")
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

// atk_relation_set_get_relation_by_type
//
// [ relationship ] trans: nothing
//
// [ result ] trans: nothing
//
func (v RelationSet) GetRelationByType(relationship RelationTypeEnum) (result Relation) {
	iv, err := _I.Get(112, "RelationSet", "get_relation_by_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	args := []gi.Argument{arg_v, arg_relationship}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_set_remove
//
// [ relation ] trans: nothing
//
func (v RelationSet) Remove(relation IRelation) {
	iv, err := _I.Get(113, "RelationSet", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if relation != nil {
		tmp = relation.P_Relation()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relation := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_relation}
	iv.Call(args, nil, nil)
}

// ignore GType struct RelationSetClass

// Enum RelationType
type RelationTypeEnum int

const (
	RelationTypeNull           RelationTypeEnum = 0
	RelationTypeControlledBy   RelationTypeEnum = 1
	RelationTypeControllerFor  RelationTypeEnum = 2
	RelationTypeLabelFor       RelationTypeEnum = 3
	RelationTypeLabelledBy     RelationTypeEnum = 4
	RelationTypeMemberOf       RelationTypeEnum = 5
	RelationTypeNodeChildOf    RelationTypeEnum = 6
	RelationTypeFlowsTo        RelationTypeEnum = 7
	RelationTypeFlowsFrom      RelationTypeEnum = 8
	RelationTypeSubwindowOf    RelationTypeEnum = 9
	RelationTypeEmbeds         RelationTypeEnum = 10
	RelationTypeEmbeddedBy     RelationTypeEnum = 11
	RelationTypePopupFor       RelationTypeEnum = 12
	RelationTypeParentWindowOf RelationTypeEnum = 13
	RelationTypeDescribedBy    RelationTypeEnum = 14
	RelationTypeDescriptionFor RelationTypeEnum = 15
	RelationTypeNodeParentOf   RelationTypeEnum = 16
	RelationTypeDetails        RelationTypeEnum = 17
	RelationTypeDetailsFor     RelationTypeEnum = 18
	RelationTypeErrorMessage   RelationTypeEnum = 19
	RelationTypeErrorFor       RelationTypeEnum = 20
	RelationTypeLastDefined    RelationTypeEnum = 21
)

func RelationTypeGetType() gi.GType {
	ret := _I.GetGType(29, "RelationType")
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
	RoleFontChooser          RoleEnum = 21
	RoleFrame                RoleEnum = 22
	RoleGlassPane            RoleEnum = 23
	RoleHtmlContainer        RoleEnum = 24
	RoleIcon                 RoleEnum = 25
	RoleImage                RoleEnum = 26
	RoleInternalFrame        RoleEnum = 27
	RoleLabel                RoleEnum = 28
	RoleLayeredPane          RoleEnum = 29
	RoleList                 RoleEnum = 30
	RoleListItem             RoleEnum = 31
	RoleMenu                 RoleEnum = 32
	RoleMenuBar              RoleEnum = 33
	RoleMenuItem             RoleEnum = 34
	RoleOptionPane           RoleEnum = 35
	RolePageTab              RoleEnum = 36
	RolePageTabList          RoleEnum = 37
	RolePanel                RoleEnum = 38
	RolePasswordText         RoleEnum = 39
	RolePopupMenu            RoleEnum = 40
	RoleProgressBar          RoleEnum = 41
	RolePushButton           RoleEnum = 42
	RoleRadioButton          RoleEnum = 43
	RoleRadioMenuItem        RoleEnum = 44
	RoleRootPane             RoleEnum = 45
	RoleRowHeader            RoleEnum = 46
	RoleScrollBar            RoleEnum = 47
	RoleScrollPane           RoleEnum = 48
	RoleSeparator            RoleEnum = 49
	RoleSlider               RoleEnum = 50
	RoleSplitPane            RoleEnum = 51
	RoleSpinButton           RoleEnum = 52
	RoleStatusbar            RoleEnum = 53
	RoleTable                RoleEnum = 54
	RoleTableCell            RoleEnum = 55
	RoleTableColumnHeader    RoleEnum = 56
	RoleTableRowHeader       RoleEnum = 57
	RoleTearOffMenuItem      RoleEnum = 58
	RoleTerminal             RoleEnum = 59
	RoleText                 RoleEnum = 60
	RoleToggleButton         RoleEnum = 61
	RoleToolBar              RoleEnum = 62
	RoleToolTip              RoleEnum = 63
	RoleTree                 RoleEnum = 64
	RoleTreeTable            RoleEnum = 65
	RoleUnknown              RoleEnum = 66
	RoleViewport             RoleEnum = 67
	RoleWindow               RoleEnum = 68
	RoleHeader               RoleEnum = 69
	RoleFooter               RoleEnum = 70
	RoleParagraph            RoleEnum = 71
	RoleRuler                RoleEnum = 72
	RoleApplication          RoleEnum = 73
	RoleAutocomplete         RoleEnum = 74
	RoleEditBar              RoleEnum = 75
	RoleEmbedded             RoleEnum = 76
	RoleEntry                RoleEnum = 77
	RoleChart                RoleEnum = 78
	RoleCaption              RoleEnum = 79
	RoleDocumentFrame        RoleEnum = 80
	RoleHeading              RoleEnum = 81
	RolePage                 RoleEnum = 82
	RoleSection              RoleEnum = 83
	RoleRedundantObject      RoleEnum = 84
	RoleForm                 RoleEnum = 85
	RoleLink                 RoleEnum = 86
	RoleInputMethodWindow    RoleEnum = 87
	RoleTableRow             RoleEnum = 88
	RoleTreeItem             RoleEnum = 89
	RoleDocumentSpreadsheet  RoleEnum = 90
	RoleDocumentPresentation RoleEnum = 91
	RoleDocumentText         RoleEnum = 92
	RoleDocumentWeb          RoleEnum = 93
	RoleDocumentEmail        RoleEnum = 94
	RoleComment              RoleEnum = 95
	RoleListBox              RoleEnum = 96
	RoleGrouping             RoleEnum = 97
	RoleImageMap             RoleEnum = 98
	RoleNotification         RoleEnum = 99
	RoleInfoBar              RoleEnum = 100
	RoleLevelBar             RoleEnum = 101
	RoleTitleBar             RoleEnum = 102
	RoleBlockQuote           RoleEnum = 103
	RoleAudio                RoleEnum = 104
	RoleVideo                RoleEnum = 105
	RoleDefinition           RoleEnum = 106
	RoleArticle              RoleEnum = 107
	RoleLandmark             RoleEnum = 108
	RoleLog                  RoleEnum = 109
	RoleMarquee              RoleEnum = 110
	RoleMath                 RoleEnum = 111
	RoleRating               RoleEnum = 112
	RoleTimer                RoleEnum = 113
	RoleDescriptionList      RoleEnum = 114
	RoleDescriptionTerm      RoleEnum = 115
	RoleDescriptionValue     RoleEnum = 116
	RoleStatic               RoleEnum = 117
	RoleMathFraction         RoleEnum = 118
	RoleMathRoot             RoleEnum = 119
	RoleSubscript            RoleEnum = 120
	RoleSuperscript          RoleEnum = 121
	RoleFootnote             RoleEnum = 122
	RoleLastDefined          RoleEnum = 123
)

func RoleGetType() gi.GType {
	ret := _I.GetGType(30, "Role")
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
	ret := _I.GetGType(31, "ScrollType")
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
	ret := _I.GetGType(32, "Selection")
	return ret
}

// atk_selection_add_selection
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) AddSelection(i int32) (result bool) {
	iv, err := _I.Get(114, "Selection", "add_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_clear_selection
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) ClearSelection() (result bool) {
	iv, err := _I.Get(115, "Selection", "clear_selection")
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

// atk_selection_get_selection_count
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) GetSelectionCount() (result int32) {
	iv, err := _I.Get(116, "Selection", "get_selection_count")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_selection_is_child_selected
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) IsChildSelected(i int32) (result bool) {
	iv, err := _I.Get(117, "Selection", "is_child_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_ref_selection
//
// [ i ] trans: nothing
//
// [ result ] trans: everything
//
func (v *SelectionIfc) RefSelection(i int32) (result Object) {
	iv, err := _I.Get(118, "Selection", "ref_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_selection_remove_selection
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) RemoveSelection(i int32) (result bool) {
	iv, err := _I.Get(119, "Selection", "remove_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_select_all_selection
//
// [ result ] trans: nothing
//
func (v *SelectionIfc) SelectAllSelection() (result bool) {
	iv, err := _I.Get(120, "Selection", "select_all_selection")
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

// ignore GType struct SelectionIface

// Object Socket
type Socket struct {
	ComponentIfc
	Object
}

func WrapSocket(p unsafe.Pointer) (r Socket) { r.P = p; return }

type ISocket interface{ P_Socket() unsafe.Pointer }

func (v Socket) P_Socket() unsafe.Pointer    { return v.P }
func (v Socket) P_Component() unsafe.Pointer { return v.P }
func SocketGetType() gi.GType {
	ret := _I.GetGType(33, "Socket")
	return ret
}

// atk_socket_new
//
// [ result ] trans: everything
//
func NewSocket() (result Socket) {
	iv, err := _I.Get(121, "Socket", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_socket_embed
//
// [ plug_id ] trans: nothing
//
func (v Socket) Embed(plug_id string) {
	iv, err := _I.Get(122, "Socket", "embed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_plug_id := gi.CString(plug_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_plug_id := gi.NewStringArgument(c_plug_id)
	args := []gi.Argument{arg_v, arg_plug_id}
	iv.Call(args, nil, nil)
	gi.Free(c_plug_id)
}

// atk_socket_is_occupied
//
// [ result ] trans: nothing
//
func (v Socket) IsOccupied() (result bool) {
	iv, err := _I.Get(123, "Socket", "is_occupied")
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

// ignore GType struct SocketClass

// Object StateSet
type StateSet struct {
	g.Object
}

func WrapStateSet(p unsafe.Pointer) (r StateSet) { r.P = p; return }

type IStateSet interface{ P_StateSet() unsafe.Pointer }

func (v StateSet) P_StateSet() unsafe.Pointer { return v.P }
func StateSetGetType() gi.GType {
	ret := _I.GetGType(34, "StateSet")
	return ret
}

// atk_state_set_new
//
// [ result ] trans: everything
//
func NewStateSet() (result StateSet) {
	iv, err := _I.Get(124, "StateSet", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_add_state
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StateSet) AddState(type1 StateTypeEnum) (result bool) {
	iv, err := _I.Get(125, "StateSet", "add_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_add_states
//
// [ types ] trans: nothing
//
// [ n_types ] trans: nothing
//
func (v StateSet) AddStates(types unsafe.Pointer, n_types int32) {
	iv, err := _I.Get(126, "StateSet", "add_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_types := gi.NewPointerArgument(types)
	arg_n_types := gi.NewInt32Argument(n_types)
	args := []gi.Argument{arg_v, arg_types, arg_n_types}
	iv.Call(args, nil, nil)
}

// atk_state_set_and_sets
//
// [ compare_set ] trans: nothing
//
// [ result ] trans: everything
//
func (v StateSet) AndSets(compare_set IStateSet) (result StateSet) {
	iv, err := _I.Get(127, "StateSet", "and_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if compare_set != nil {
		tmp = compare_set.P_StateSet()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_clear_states
//
func (v StateSet) ClearStates() {
	iv, err := _I.Get(128, "StateSet", "clear_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_state_set_contains_state
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StateSet) ContainsState(type1 StateTypeEnum) (result bool) {
	iv, err := _I.Get(129, "StateSet", "contains_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_contains_states
//
// [ types ] trans: nothing
//
// [ n_types ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StateSet) ContainsStates(types unsafe.Pointer, n_types int32) (result bool) {
	iv, err := _I.Get(130, "StateSet", "contains_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_types := gi.NewPointerArgument(types)
	arg_n_types := gi.NewInt32Argument(n_types)
	args := []gi.Argument{arg_v, arg_types, arg_n_types}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_is_empty
//
// [ result ] trans: nothing
//
func (v StateSet) IsEmpty() (result bool) {
	iv, err := _I.Get(131, "StateSet", "is_empty")
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

// atk_state_set_or_sets
//
// [ compare_set ] trans: nothing
//
// [ result ] trans: everything
//
func (v StateSet) OrSets(compare_set IStateSet) (result StateSet) {
	iv, err := _I.Get(132, "StateSet", "or_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if compare_set != nil {
		tmp = compare_set.P_StateSet()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_remove_state
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v StateSet) RemoveState(type1 StateTypeEnum) (result bool) {
	iv, err := _I.Get(133, "StateSet", "remove_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_xor_sets
//
// [ compare_set ] trans: nothing
//
// [ result ] trans: everything
//
func (v StateSet) XorSets(compare_set IStateSet) (result StateSet) {
	iv, err := _I.Get(134, "StateSet", "xor_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if compare_set != nil {
		tmp = compare_set.P_StateSet()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
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
	StateTypeDefunct                StateTypeEnum = 5
	StateTypeEditable               StateTypeEnum = 6
	StateTypeEnabled                StateTypeEnum = 7
	StateTypeExpandable             StateTypeEnum = 8
	StateTypeExpanded               StateTypeEnum = 9
	StateTypeFocusable              StateTypeEnum = 10
	StateTypeFocused                StateTypeEnum = 11
	StateTypeHorizontal             StateTypeEnum = 12
	StateTypeIconified              StateTypeEnum = 13
	StateTypeModal                  StateTypeEnum = 14
	StateTypeMultiLine              StateTypeEnum = 15
	StateTypeMultiselectable        StateTypeEnum = 16
	StateTypeOpaque                 StateTypeEnum = 17
	StateTypePressed                StateTypeEnum = 18
	StateTypeResizable              StateTypeEnum = 19
	StateTypeSelectable             StateTypeEnum = 20
	StateTypeSelected               StateTypeEnum = 21
	StateTypeSensitive              StateTypeEnum = 22
	StateTypeShowing                StateTypeEnum = 23
	StateTypeSingleLine             StateTypeEnum = 24
	StateTypeStale                  StateTypeEnum = 25
	StateTypeTransient              StateTypeEnum = 26
	StateTypeVertical               StateTypeEnum = 27
	StateTypeVisible                StateTypeEnum = 28
	StateTypeManagesDescendants     StateTypeEnum = 29
	StateTypeIndeterminate          StateTypeEnum = 30
	StateTypeTruncated              StateTypeEnum = 31
	StateTypeRequired               StateTypeEnum = 32
	StateTypeInvalidEntry           StateTypeEnum = 33
	StateTypeSupportsAutocompletion StateTypeEnum = 34
	StateTypeSelectableText         StateTypeEnum = 35
	StateTypeDefault                StateTypeEnum = 36
	StateTypeAnimated               StateTypeEnum = 37
	StateTypeVisited                StateTypeEnum = 38
	StateTypeCheckable              StateTypeEnum = 39
	StateTypeHasPopup               StateTypeEnum = 40
	StateTypeHasTooltip             StateTypeEnum = 41
	StateTypeReadOnly               StateTypeEnum = 42
	StateTypeLastDefined            StateTypeEnum = 43
)

func StateTypeGetType() gi.GType {
	ret := _I.GetGType(35, "StateType")
	return ret
}

// Interface StreamableContent
type StreamableContent struct {
	StreamableContentIfc
	P unsafe.Pointer
}
type StreamableContentIfc struct{}
type IStreamableContent interface{ P_StreamableContent() unsafe.Pointer }

func (v StreamableContent) P_StreamableContent() unsafe.Pointer { return v.P }
func StreamableContentGetType() gi.GType {
	ret := _I.GetGType(36, "StreamableContent")
	return ret
}

// atk_streamable_content_get_mime_type
//
// [ i ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *StreamableContentIfc) GetMimeType(i int32) (result string) {
	iv, err := _I.Get(135, "StreamableContent", "get_mime_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_streamable_content_get_n_mime_types
//
// [ result ] trans: nothing
//
func (v *StreamableContentIfc) GetNMimeTypes() (result int32) {
	iv, err := _I.Get(136, "StreamableContent", "get_n_mime_types")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_streamable_content_get_stream
//
// [ mime_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v *StreamableContentIfc) GetStream(mime_type string) (result g.IOChannel) {
	iv, err := _I.Get(137, "StreamableContent", "get_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_mime_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_mime_type)
	result.P = ret.Pointer()
	return
}

// atk_streamable_content_get_uri
//
// [ mime_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *StreamableContentIfc) GetUri(mime_type string) (result string) {
	iv, err := _I.Get(138, "StreamableContent", "get_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_mime_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_mime_type)
	result = ret.String().Copy()
	return
}

// ignore GType struct StreamableContentIface

// Interface Table
type Table struct {
	TableIfc
	P unsafe.Pointer
}
type TableIfc struct{}
type ITable interface{ P_Table() unsafe.Pointer }

func (v Table) P_Table() unsafe.Pointer { return v.P }
func TableGetType() gi.GType {
	ret := _I.GetGType(37, "Table")
	return ret
}

// atk_table_add_column_selection
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) AddColumnSelection(column int32) (result bool) {
	iv, err := _I.Get(139, "Table", "add_column_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_add_row_selection
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) AddRowSelection(row int32) (result bool) {
	iv, err := _I.Get(140, "Table", "add_row_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_get_caption
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetCaption() (result Object) {
	iv, err := _I.Get(141, "Table", "get_caption")
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

// Deprecated
//
// atk_table_get_column_at_index
//
// [ index_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetColumnAtIndex(index_ int32) (result int32) {
	iv, err := _I.Get(142, "Table", "get_column_at_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_column_description
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetColumnDescription(column int32) (result string) {
	iv, err := _I.Get(143, "Table", "get_column_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_table_get_column_extent_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetColumnExtentAt(row int32, column int32) (result int32) {
	iv, err := _I.Get(144, "Table", "get_column_extent_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_column_header
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetColumnHeader(column int32) (result Object) {
	iv, err := _I.Get(145, "Table", "get_column_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// atk_table_get_index_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetIndexAt(row int32, column int32) (result int32) {
	iv, err := _I.Get(146, "Table", "get_index_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_n_columns
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetNColumns() (result int32) {
	iv, err := _I.Get(147, "Table", "get_n_columns")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_n_rows
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetNRows() (result int32) {
	iv, err := _I.Get(148, "Table", "get_n_rows")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// atk_table_get_row_at_index
//
// [ index_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowAtIndex(index_ int32) (result int32) {
	iv, err := _I.Get(149, "Table", "get_row_at_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_row_description
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowDescription(row int32) (result string) {
	iv, err := _I.Get(150, "Table", "get_row_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_table_get_row_extent_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowExtentAt(row int32, column int32) (result int32) {
	iv, err := _I.Get(151, "Table", "get_row_extent_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_row_header
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetRowHeader(row int32) (result Object) {
	iv, err := _I.Get(152, "Table", "get_row_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_table_get_selected_columns
//
// [ selected ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetSelectedColumns(selected int32) (result int32) {
	iv, err := _I.Get(153, "Table", "get_selected_columns")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selected := gi.NewInt32Argument(selected)
	args := []gi.Argument{arg_v, arg_selected}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_selected_rows
//
// [ selected ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) GetSelectedRows(selected int32) (result int32) {
	iv, err := _I.Get(154, "Table", "get_selected_rows")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selected := gi.NewInt32Argument(selected)
	args := []gi.Argument{arg_v, arg_selected}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_summary
//
// [ result ] trans: everything
//
func (v *TableIfc) GetSummary() (result Object) {
	iv, err := _I.Get(155, "Table", "get_summary")
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

// atk_table_is_column_selected
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) IsColumnSelected(column int32) (result bool) {
	iv, err := _I.Get(156, "Table", "is_column_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_is_row_selected
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) IsRowSelected(row int32) (result bool) {
	iv, err := _I.Get(157, "Table", "is_row_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_is_selected
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) IsSelected(row int32, column int32) (result bool) {
	iv, err := _I.Get(158, "Table", "is_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_ref_at
//
// [ row ] trans: nothing
//
// [ column ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TableIfc) RefAt(row int32, column int32) (result Object) {
	iv, err := _I.Get(159, "Table", "ref_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_table_remove_column_selection
//
// [ column ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) RemoveColumnSelection(column int32) (result bool) {
	iv, err := _I.Get(160, "Table", "remove_column_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_remove_row_selection
//
// [ row ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TableIfc) RemoveRowSelection(row int32) (result bool) {
	iv, err := _I.Get(161, "Table", "remove_row_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_set_caption
//
// [ caption ] trans: nothing
//
func (v *TableIfc) SetCaption(caption IObject) {
	iv, err := _I.Get(162, "Table", "set_caption")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if caption != nil {
		tmp = caption.P_Object()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_caption := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_caption}
	iv.Call(args, nil, nil)
}

// atk_table_set_column_description
//
// [ column ] trans: nothing
//
// [ description ] trans: nothing
//
func (v *TableIfc) SetColumnDescription(column int32, description string) {
	iv, err := _I.Get(163, "Table", "set_column_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_column, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// atk_table_set_column_header
//
// [ column ] trans: nothing
//
// [ header ] trans: nothing
//
func (v *TableIfc) SetColumnHeader(column int32, header IObject) {
	iv, err := _I.Get(164, "Table", "set_column_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if header != nil {
		tmp = header.P_Object()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_header := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_column, arg_header}
	iv.Call(args, nil, nil)
}

// atk_table_set_row_description
//
// [ row ] trans: nothing
//
// [ description ] trans: nothing
//
func (v *TableIfc) SetRowDescription(row int32, description string) {
	iv, err := _I.Get(165, "Table", "set_row_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_row, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// atk_table_set_row_header
//
// [ row ] trans: nothing
//
// [ header ] trans: nothing
//
func (v *TableIfc) SetRowHeader(row int32, header IObject) {
	iv, err := _I.Get(166, "Table", "set_row_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if header != nil {
		tmp = header.P_Object()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_header := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_row, arg_header}
	iv.Call(args, nil, nil)
}

// atk_table_set_summary
//
// [ accessible ] trans: nothing
//
func (v *TableIfc) SetSummary(accessible IObject) {
	iv, err := _I.Get(167, "Table", "set_summary")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if accessible != nil {
		tmp = accessible.P_Object()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_accessible := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_accessible}
	iv.Call(args, nil, nil)
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
	ret := _I.GetGType(38, "TableCell")
	return ret
}

// atk_table_cell_get_column_header_cells
//
// [ result ] trans: everything
//
func (v *TableCellIfc) GetColumnHeaderCells() (result int /*TODO_TYPE array type: 2, isZeroTerm: false*/) {
	iv, err := _I.Get(168, "TableCell", "get_column_header_cells")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_table_cell_get_column_span
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetColumnSpan() (result int32) {
	iv, err := _I.Get(169, "TableCell", "get_column_span")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_cell_get_position
//
// [ row ] trans: everything, dir: out
//
// [ column ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetPosition() (result bool, row int32, column int32) {
	iv, err := _I.Get(170, "TableCell", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	row = outArgs[0].Int32()
	column = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// atk_table_cell_get_row_column_span
//
// [ row ] trans: everything, dir: out
//
// [ column ] trans: everything, dir: out
//
// [ row_span ] trans: everything, dir: out
//
// [ column_span ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetRowColumnSpan() (result bool, row int32, column int32, row_span int32, column_span int32) {
	iv, err := _I.Get(171, "TableCell", "get_row_column_span")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_row_span := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_column_span := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_row_span, arg_column_span}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	row = outArgs[0].Int32()
	column = outArgs[1].Int32()
	row_span = outArgs[2].Int32()
	column_span = outArgs[3].Int32()
	result = ret.Bool()
	return
}

// atk_table_cell_get_row_header_cells
//
// [ result ] trans: everything
//
func (v *TableCellIfc) GetRowHeaderCells() (result int /*TODO_TYPE array type: 2, isZeroTerm: false*/) {
	iv, err := _I.Get(172, "TableCell", "get_row_header_cells")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_table_cell_get_row_span
//
// [ result ] trans: nothing
//
func (v *TableCellIfc) GetRowSpan() (result int32) {
	iv, err := _I.Get(173, "TableCell", "get_row_span")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_cell_get_table
//
// [ result ] trans: everything
//
func (v *TableCellIfc) GetTable() (result Object) {
	iv, err := _I.Get(174, "TableCell", "get_table")
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

// ignore GType struct TableCellIface

// ignore GType struct TableIface

// Interface Text
type Text struct {
	TextIfc
	P unsafe.Pointer
}
type TextIfc struct{}
type IText interface{ P_Text() unsafe.Pointer }

func (v Text) P_Text() unsafe.Pointer { return v.P }
func TextGetType() gi.GType {
	ret := _I.GetGType(39, "Text")
	return ret
}

// atk_text_free_ranges
//
// [ ranges ] trans: nothing
//
func TextFreeRanges1(ranges gi.PointerArray) {
	iv, err := _I.Get(175, "Text", "free_ranges")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ranges := gi.NewPointerArgument(ranges.P)
	args := []gi.Argument{arg_ranges}
	iv.Call(args, nil, nil)
}

// atk_text_add_selection
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) AddSelection(start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(176, "Text", "add_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_text_get_bounded_ranges
//
// [ rect ] trans: nothing
//
// [ coord_type ] trans: nothing
//
// [ x_clip_type ] trans: nothing
//
// [ y_clip_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetBoundedRanges(rect TextRectangle, coord_type CoordTypeEnum, x_clip_type TextClipTypeEnum, y_clip_type TextClipTypeEnum) (result gi.PointerArray) {
	iv, err := _I.Get(177, "Text", "get_bounded_ranges")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	arg_x_clip_type := gi.NewIntArgument(int(x_clip_type))
	arg_y_clip_type := gi.NewIntArgument(int(y_clip_type))
	args := []gi.Argument{arg_v, arg_rect, arg_coord_type, arg_x_clip_type, arg_y_clip_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.PointerArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// atk_text_get_caret_offset
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetCaretOffset() (result int32) {
	iv, err := _I.Get(178, "Text", "get_caret_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_text_get_character_at_offset
//
// [ offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetCharacterAtOffset(offset int32) (result rune) {
	iv, err := _I.Get(179, "Text", "get_character_at_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// atk_text_get_character_count
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetCharacterCount() (result int32) {
	iv, err := _I.Get(180, "Text", "get_character_count")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_text_get_character_extents
//
// [ offset ] trans: nothing
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
// [ coords ] trans: nothing
//
func (v *TextIfc) GetCharacterExtents(offset int32, coords CoordTypeEnum) (x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(181, "Text", "get_character_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_coords := gi.NewIntArgument(int(coords))
	args := []gi.Argument{arg_v, arg_offset, arg_x, arg_y, arg_width, arg_height, arg_coords}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	width = outArgs[2].Int32()
	height = outArgs[3].Int32()
	return
}

// atk_text_get_default_attributes
//
// [ result ] trans: everything
//
func (v *TextIfc) GetDefaultAttributes() (result g.SList) {
	iv, err := _I.Get(182, "Text", "get_default_attributes")
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

// atk_text_get_n_selections
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetNSelections() (result int32) {
	iv, err := _I.Get(183, "Text", "get_n_selections")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_text_get_offset_at_point
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ coords ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) GetOffsetAtPoint(x int32, y int32, coords CoordTypeEnum) (result int32) {
	iv, err := _I.Get(184, "Text", "get_offset_at_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coords := gi.NewIntArgument(int(coords))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coords}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_text_get_range_extents
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ coord_type ] trans: nothing
//
// [ rect ] trans: nothing, dir: out
//
func (v *TextIfc) GetRangeExtents(start_offset int32, end_offset int32, coord_type CoordTypeEnum, rect TextRectangle) {
	iv, err := _I.Get(185, "Text", "get_range_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset, arg_coord_type, arg_rect}
	iv.Call(args, nil, nil)
}

// atk_text_get_run_attributes
//
// [ offset ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetRunAttributes(offset int32) (result g.SList, start_offset int32, end_offset int32) {
	iv, err := _I.Get(186, "Text", "get_run_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// atk_text_get_selection
//
// [ selection_num ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetSelection(selection_num int32) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(187, "Text", "get_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_selection_num, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result = ret.String().Take()
	return
}

// atk_text_get_string_at_offset
//
// [ offset ] trans: nothing
//
// [ granularity ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetStringAtOffset(offset int32, granularity TextGranularityEnum) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(188, "Text", "get_string_at_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_granularity := gi.NewIntArgument(int(granularity))
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_granularity, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result = ret.String().Take()
	return
}

// atk_text_get_text
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: everything
//
func (v *TextIfc) GetText(start_offset int32, end_offset int32) (result string) {
	iv, err := _I.Get(189, "Text", "get_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// Deprecated
//
// atk_text_get_text_after_offset
//
// [ offset ] trans: nothing
//
// [ boundary_type ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextAfterOffset(offset int32, boundary_type TextBoundaryEnum) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(190, "Text", "get_text_after_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_boundary_type := gi.NewIntArgument(int(boundary_type))
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_boundary_type, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result = ret.String().Take()
	return
}

// Deprecated
//
// atk_text_get_text_at_offset
//
// [ offset ] trans: nothing
//
// [ boundary_type ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextAtOffset(offset int32, boundary_type TextBoundaryEnum) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(191, "Text", "get_text_at_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_boundary_type := gi.NewIntArgument(int(boundary_type))
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_boundary_type, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result = ret.String().Take()
	return
}

// Deprecated
//
// atk_text_get_text_before_offset
//
// [ offset ] trans: nothing
//
// [ boundary_type ] trans: nothing
//
// [ start_offset ] trans: everything, dir: out
//
// [ end_offset ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v *TextIfc) GetTextBeforeOffset(offset int32, boundary_type TextBoundaryEnum) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(192, "Text", "get_text_before_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_boundary_type := gi.NewIntArgument(int(boundary_type))
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_boundary_type, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	result = ret.String().Take()
	return
}

// atk_text_remove_selection
//
// [ selection_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) RemoveSelection(selection_num int32) (result bool) {
	iv, err := _I.Get(193, "Text", "remove_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	args := []gi.Argument{arg_v, arg_selection_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_text_set_caret_offset
//
// [ offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) SetCaretOffset(offset int32) (result bool) {
	iv, err := _I.Get(194, "Text", "set_caret_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_text_set_selection
//
// [ selection_num ] trans: nothing
//
// [ start_offset ] trans: nothing
//
// [ end_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *TextIfc) SetSelection(selection_num int32, start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(195, "Text", "set_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_selection_num, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum TextAttribute
type TextAttributeEnum int

const (
	TextAttributeInvalid          TextAttributeEnum = 0
	TextAttributeLeftMargin       TextAttributeEnum = 1
	TextAttributeRightMargin      TextAttributeEnum = 2
	TextAttributeIndent           TextAttributeEnum = 3
	TextAttributeInvisible        TextAttributeEnum = 4
	TextAttributeEditable         TextAttributeEnum = 5
	TextAttributePixelsAboveLines TextAttributeEnum = 6
	TextAttributePixelsBelowLines TextAttributeEnum = 7
	TextAttributePixelsInsideWrap TextAttributeEnum = 8
	TextAttributeBgFullHeight     TextAttributeEnum = 9
	TextAttributeRise             TextAttributeEnum = 10
	TextAttributeUnderline        TextAttributeEnum = 11
	TextAttributeStrikethrough    TextAttributeEnum = 12
	TextAttributeSize             TextAttributeEnum = 13
	TextAttributeScale            TextAttributeEnum = 14
	TextAttributeWeight           TextAttributeEnum = 15
	TextAttributeLanguage         TextAttributeEnum = 16
	TextAttributeFamilyName       TextAttributeEnum = 17
	TextAttributeBgColor          TextAttributeEnum = 18
	TextAttributeFgColor          TextAttributeEnum = 19
	TextAttributeBgStipple        TextAttributeEnum = 20
	TextAttributeFgStipple        TextAttributeEnum = 21
	TextAttributeWrapMode         TextAttributeEnum = 22
	TextAttributeDirection        TextAttributeEnum = 23
	TextAttributeJustification    TextAttributeEnum = 24
	TextAttributeStretch          TextAttributeEnum = 25
	TextAttributeVariant          TextAttributeEnum = 26
	TextAttributeStyle            TextAttributeEnum = 27
	TextAttributeLastDefined      TextAttributeEnum = 28
)

func TextAttributeGetType() gi.GType {
	ret := _I.GetGType(40, "TextAttribute")
	return ret
}

// Enum TextBoundary
type TextBoundaryEnum int

const (
	TextBoundaryChar          TextBoundaryEnum = 0
	TextBoundaryWordStart     TextBoundaryEnum = 1
	TextBoundaryWordEnd       TextBoundaryEnum = 2
	TextBoundarySentenceStart TextBoundaryEnum = 3
	TextBoundarySentenceEnd   TextBoundaryEnum = 4
	TextBoundaryLineStart     TextBoundaryEnum = 5
	TextBoundaryLineEnd       TextBoundaryEnum = 6
)

func TextBoundaryGetType() gi.GType {
	ret := _I.GetGType(41, "TextBoundary")
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
	ret := _I.GetGType(42, "TextClipType")
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
	ret := _I.GetGType(43, "TextGranularity")
	return ret
}

// ignore GType struct TextIface

// Struct TextRange
type TextRange struct {
	P unsafe.Pointer
}

const SizeOfStructTextRange = 32

func TextRangeGetType() gi.GType {
	ret := _I.GetGType(44, "TextRange")
	return ret
}

// Struct TextRectangle
type TextRectangle struct {
	P unsafe.Pointer
}

const SizeOfStructTextRectangle = 16

func TextRectangleGetType() gi.GType {
	ret := _I.GetGType(45, "TextRectangle")
	return ret
}

// Object Util
type Util struct {
	g.Object
}

func WrapUtil(p unsafe.Pointer) (r Util) { r.P = p; return }

type IUtil interface{ P_Util() unsafe.Pointer }

func (v Util) P_Util() unsafe.Pointer { return v.P }
func UtilGetType() gi.GType {
	ret := _I.GetGType(46, "Util")
	return ret
}

// ignore GType struct UtilClass

// Interface Value
type Value struct {
	ValueIfc
	P unsafe.Pointer
}
type ValueIfc struct{}
type IValue interface{ P_Value() unsafe.Pointer }

func (v Value) P_Value() unsafe.Pointer { return v.P }
func ValueGetType() gi.GType {
	ret := _I.GetGType(47, "Value")
	return ret
}

// Deprecated
//
// atk_value_get_current_value
//
// [ value ] trans: nothing, dir: out
//
func (v *ValueIfc) GetCurrentValue(value g.Value) {
	iv, err := _I.Get(196, "Value", "get_current_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// atk_value_get_increment
//
// [ result ] trans: nothing
//
func (v *ValueIfc) GetIncrement() (result float64) {
	iv, err := _I.Get(197, "Value", "get_increment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// Deprecated
//
// atk_value_get_maximum_value
//
// [ value ] trans: nothing, dir: out
//
func (v *ValueIfc) GetMaximumValue(value g.Value) {
	iv, err := _I.Get(198, "Value", "get_maximum_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// atk_value_get_minimum_increment
//
// [ value ] trans: nothing, dir: out
//
func (v *ValueIfc) GetMinimumIncrement(value g.Value) {
	iv, err := _I.Get(199, "Value", "get_minimum_increment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// atk_value_get_minimum_value
//
// [ value ] trans: nothing, dir: out
//
func (v *ValueIfc) GetMinimumValue(value g.Value) {
	iv, err := _I.Get(200, "Value", "get_minimum_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// atk_value_get_range
//
// [ result ] trans: everything
//
func (v *ValueIfc) GetRange() (result Range) {
	iv, err := _I.Get(201, "Value", "get_range")
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

// atk_value_get_sub_ranges
//
// [ result ] trans: everything
//
func (v *ValueIfc) GetSubRanges() (result g.SList) {
	iv, err := _I.Get(202, "Value", "get_sub_ranges")
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

// atk_value_get_value_and_text
//
// [ value ] trans: everything, dir: out
//
// [ text ] trans: everything, dir: out
//
func (v *ValueIfc) GetValueAndText() (value float64, text string) {
	iv, err := _I.Get(203, "Value", "get_value_and_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_text := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_value, arg_text}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Double()
	text = outArgs[1].String().Take()
	return
}

// Deprecated
//
// atk_value_set_current_value
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *ValueIfc) SetCurrentValue(value g.Value) (result bool) {
	iv, err := _I.Get(204, "Value", "set_current_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_value_set_value
//
// [ new_value ] trans: nothing
//
func (v *ValueIfc) SetValue(new_value float64) {
	iv, err := _I.Get(205, "Value", "set_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_new_value := gi.NewDoubleArgument(new_value)
	args := []gi.Argument{arg_v, arg_new_value}
	iv.Call(args, nil, nil)
}

// ignore GType struct ValueIface

// Enum ValueType
type ValueTypeEnum int

const (
	ValueTypeVeryWeak    ValueTypeEnum = 0
	ValueTypeWeak        ValueTypeEnum = 1
	ValueTypeAcceptable  ValueTypeEnum = 2
	ValueTypeStrong      ValueTypeEnum = 3
	ValueTypeVeryStrong  ValueTypeEnum = 4
	ValueTypeVeryLow     ValueTypeEnum = 5
	ValueTypeLow         ValueTypeEnum = 6
	ValueTypeMedium      ValueTypeEnum = 7
	ValueTypeHigh        ValueTypeEnum = 8
	ValueTypeVeryHigh    ValueTypeEnum = 9
	ValueTypeVeryBad     ValueTypeEnum = 10
	ValueTypeBad         ValueTypeEnum = 11
	ValueTypeGood        ValueTypeEnum = 12
	ValueTypeVeryGood    ValueTypeEnum = 13
	ValueTypeBest        ValueTypeEnum = 14
	ValueTypeLastDefined ValueTypeEnum = 15
)

func ValueTypeGetType() gi.GType {
	ret := _I.GetGType(48, "ValueType")
	return ret
}

// Interface Window
type Window struct {
	WindowIfc
	P unsafe.Pointer
}
type WindowIfc struct{}
type IWindow interface{ P_Window() unsafe.Pointer }

func (v Window) P_Window() unsafe.Pointer { return v.P }
func WindowGetType() gi.GType {
	ret := _I.GetGType(49, "Window")
	return ret
}

// ignore GType struct WindowIface

// atk_attribute_set_free
//
// [ attrib_set ] trans: nothing
//
func AttributeSetFree(attrib_set g.SList) {
	iv, err := _I.Get(206, "attribute_set_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attrib_set := gi.NewPointerArgument(attrib_set.P)
	args := []gi.Argument{arg_attrib_set}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// atk_focus_tracker_notify
//
// [ object ] trans: nothing
//
func FocusTrackerNotify(object IObject) {
	iv, err := _I.Get(207, "focus_tracker_notify", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	arg_object := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_object}
	iv.Call(args, nil, nil)
}

// atk_get_binary_age
//
// [ result ] trans: nothing
//
func GetBinaryAge() (result uint32) {
	iv, err := _I.Get(208, "get_binary_age", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_default_registry
//
// [ result ] trans: everything
//
func GetDefaultRegistry() (result Registry) {
	iv, err := _I.Get(209, "get_default_registry", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_get_focus_object
//
// [ result ] trans: nothing
//
func GetFocusObject() (result Object) {
	iv, err := _I.Get(210, "get_focus_object", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_get_interface_age
//
// [ result ] trans: nothing
//
func GetInterfaceAge() (result uint32) {
	iv, err := _I.Get(211, "get_interface_age", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_major_version
//
// [ result ] trans: nothing
//
func GetMajorVersion() (result uint32) {
	iv, err := _I.Get(212, "get_major_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_micro_version
//
// [ result ] trans: nothing
//
func GetMicroVersion() (result uint32) {
	iv, err := _I.Get(213, "get_micro_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_minor_version
//
// [ result ] trans: nothing
//
func GetMinorVersion() (result uint32) {
	iv, err := _I.Get(214, "get_minor_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_root
//
// [ result ] trans: nothing
//
func GetRoot() (result Object) {
	iv, err := _I.Get(215, "get_root", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_get_toolkit_name
//
// [ result ] trans: nothing
//
func GetToolkitName() (result string) {
	iv, err := _I.Get(216, "get_toolkit_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_get_toolkit_version
//
// [ result ] trans: nothing
//
func GetToolkitVersion() (result string) {
	iv, err := _I.Get(217, "get_toolkit_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_get_version
//
// [ result ] trans: nothing
//
func GetVersion() (result string) {
	iv, err := _I.Get(218, "get_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_relation_type_for_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func RelationTypeForName(name string) (result RelationTypeEnum) {
	iv, err := _I.Get(219, "relation_type_for_name", "")
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
	result = RelationTypeEnum(ret.Int())
	return
}

// atk_relation_type_get_name
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func RelationTypeGetName(type1 RelationTypeEnum) (result string) {
	iv, err := _I.Get(220, "relation_type_get_name", "")
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

// atk_relation_type_register
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func RelationTypeRegister(name string) (result RelationTypeEnum) {
	iv, err := _I.Get(221, "relation_type_register", "")
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
	result = RelationTypeEnum(ret.Int())
	return
}

// Deprecated
//
// atk_remove_focus_tracker
//
// [ tracker_id ] trans: nothing
//
func RemoveFocusTracker(tracker_id uint32) {
	iv, err := _I.Get(222, "remove_focus_tracker", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tracker_id := gi.NewUint32Argument(tracker_id)
	args := []gi.Argument{arg_tracker_id}
	iv.Call(args, nil, nil)
}

// atk_remove_global_event_listener
//
// [ listener_id ] trans: nothing
//
func RemoveGlobalEventListener(listener_id uint32) {
	iv, err := _I.Get(223, "remove_global_event_listener", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_listener_id := gi.NewUint32Argument(listener_id)
	args := []gi.Argument{arg_listener_id}
	iv.Call(args, nil, nil)
}

// atk_remove_key_event_listener
//
// [ listener_id ] trans: nothing
//
func RemoveKeyEventListener(listener_id uint32) {
	iv, err := _I.Get(224, "remove_key_event_listener", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_listener_id := gi.NewUint32Argument(listener_id)
	args := []gi.Argument{arg_listener_id}
	iv.Call(args, nil, nil)
}

// atk_role_for_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func RoleForName(name string) (result RoleEnum) {
	iv, err := _I.Get(225, "role_for_name", "")
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
	result = RoleEnum(ret.Int())
	return
}

// atk_role_get_localized_name
//
// [ role ] trans: nothing
//
// [ result ] trans: nothing
//
func RoleGetLocalizedName(role RoleEnum) (result string) {
	iv, err := _I.Get(226, "role_get_localized_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_role := gi.NewIntArgument(int(role))
	args := []gi.Argument{arg_role}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_role_get_name
//
// [ role ] trans: nothing
//
// [ result ] trans: nothing
//
func RoleGetName(role RoleEnum) (result string) {
	iv, err := _I.Get(227, "role_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_role := gi.NewIntArgument(int(role))
	args := []gi.Argument{arg_role}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// Deprecated
//
// atk_role_register
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func RoleRegister(name string) (result RoleEnum) {
	iv, err := _I.Get(228, "role_register", "")
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
	result = RoleEnum(ret.Int())
	return
}

// atk_state_type_for_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func StateTypeForName(name string) (result StateTypeEnum) {
	iv, err := _I.Get(229, "state_type_for_name", "")
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
	result = StateTypeEnum(ret.Int())
	return
}

// atk_state_type_get_name
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func StateTypeGetName(type1 StateTypeEnum) (result string) {
	iv, err := _I.Get(230, "state_type_get_name", "")
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

// atk_state_type_register
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func StateTypeRegister(name string) (result StateTypeEnum) {
	iv, err := _I.Get(231, "state_type_register", "")
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
	result = StateTypeEnum(ret.Int())
	return
}

// atk_text_attribute_for_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func TextAttributeForName(name string) (result TextAttributeEnum) {
	iv, err := _I.Get(232, "text_attribute_for_name", "")
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
	result = TextAttributeEnum(ret.Int())
	return
}

// atk_text_attribute_get_name
//
// [ attr ] trans: nothing
//
// [ result ] trans: nothing
//
func TextAttributeGetName(attr TextAttributeEnum) (result string) {
	iv, err := _I.Get(233, "text_attribute_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attr := gi.NewIntArgument(int(attr))
	args := []gi.Argument{arg_attr}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_text_attribute_get_value
//
// [ attr ] trans: nothing
//
// [ index_ ] trans: nothing
//
// [ result ] trans: nothing
//
func TextAttributeGetValue(attr TextAttributeEnum, index_ int32) (result string) {
	iv, err := _I.Get(234, "text_attribute_get_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attr := gi.NewIntArgument(int(attr))
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_attr, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_text_attribute_register
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func TextAttributeRegister(name string) (result TextAttributeEnum) {
	iv, err := _I.Get(235, "text_attribute_register", "")
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
	result = TextAttributeEnum(ret.Int())
	return
}

// atk_text_free_ranges
//
// [ ranges ] trans: nothing
//
func TextFreeRanges(ranges gi.PointerArray) {
	iv, err := _I.Get(236, "text_free_ranges", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ranges := gi.NewPointerArgument(ranges.P)
	args := []gi.Argument{arg_ranges}
	iv.Call(args, nil, nil)
}

// atk_value_type_get_localized_name
//
// [ value_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueTypeGetLocalizedName(value_type ValueTypeEnum) (result string) {
	iv, err := _I.Get(237, "value_type_get_localized_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value_type := gi.NewIntArgument(int(value_type))
	args := []gi.Argument{arg_value_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// atk_value_type_get_name
//
// [ value_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueTypeGetName(value_type ValueTypeEnum) (result string) {
	iv, err := _I.Get(238, "value_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value_type := gi.NewIntArgument(int(value_type))
	args := []gi.Argument{arg_value_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// constants
const (
	BINARY_AGE           = 23010
	INTERFACE_AGE        = 1
	MAJOR_VERSION        = 2
	MICRO_VERSION        = 0
	MINOR_VERSION        = 30
	VERSION_MIN_REQUIRED = 2
)
