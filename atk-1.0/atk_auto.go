package atk

import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

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

// atk_action_do_action
// container is not nil, container is Action
// is method
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
// container is not nil, container is Action
// is method
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
	result = ret.String().Take()
	return
}

// atk_action_get_keybinding
// container is not nil, container is Action
// is method
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
	result = ret.String().Take()
	return
}

// atk_action_get_localized_name
// container is not nil, container is Action
// is method
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
	result = ret.String().Take()
	return
}

// atk_action_get_n_actions
// container is not nil, container is Action
// is method
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
// container is not nil, container is Action
// is method
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
	result = ret.String().Take()
	return
}

// atk_action_set_description
// container is not nil, container is Action
// is method
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

// atk_attribute_set_free
// container is not nil, container is Attribute
// is method
// arg0Type tag: gslist, isPtr: true
func AttributeSetFree1(attrib_set int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(7, "Attribute", "set_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attrib_set := gi.NewIntArgument(attrib_set) /*TODO*/
	args := []gi.Argument{arg_attrib_set}
	iv.Call(args, nil, nil)
}

// Interface Component
type Component struct {
	ComponentIfc
	P unsafe.Pointer
}
type ComponentIfc struct{}

// atk_component_contains
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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

// atk_component_get_position
// container is not nil, container is Component
// is method
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

// atk_component_get_size
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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

// atk_component_remove_focus_handler
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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
// container is not nil, container is Component
// is method
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

// Interface Document
type Document struct {
	DocumentIfc
	P unsafe.Pointer
}
type DocumentIfc struct{}

// atk_document_get_attribute_value
// container is not nil, container is Document
// is method
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
	result = ret.String().Take()
	return
}

// atk_document_get_attributes
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetAttributes() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(24, "Document", "get_attributes")
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

// atk_document_get_current_page_number
// container is not nil, container is Document
// is method
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

// atk_document_get_document
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetDocument() {
	iv, err := _I.Get(26, "Document", "get_document")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_document_get_document_type
// container is not nil, container is Document
// is method
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
	result = ret.String().Take()
	return
}

// atk_document_get_locale
// container is not nil, container is Document
// is method
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
	result = ret.String().Take()
	return
}

// atk_document_get_page_count
// container is not nil, container is Document
// is method
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
// container is not nil, container is Document
// is method
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

// atk_editable_text_copy_text
// container is not nil, container is EditableText
// is method
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
// container is not nil, container is EditableText
// is method
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
// container is not nil, container is EditableText
// is method
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
// container is not nil, container is EditableText
// is method
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
// container is not nil, container is EditableText
// is method
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
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) SetRunAttributes(attrib_set int /*TODO_TYPE isPtr: true, tag: gslist*/, start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(36, "EditableText", "set_run_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attrib_set := gi.NewIntArgument(attrib_set) /*TODO*/
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_attrib_set, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_editable_text_set_text_contents
// container is not nil, container is EditableText
// is method
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
// Object GObjectAccessible
type GObjectAccessible struct {
	Object
}

func WrapGObjectAccessible(p unsafe.Pointer) (r GObjectAccessible) { r.P = p; return }

type IGObjectAccessible interface{ P_GObjectAccessible() unsafe.Pointer }

func (v GObjectAccessible) P_GObjectAccessible() unsafe.Pointer { return v.P }

// atk_gobject_accessible_for_object
// container is not nil, container is GObjectAccessible
// is method
// arg0Type tag: interface, isPtr: true
func GObjectAccessibleForObject1(obj gobject.IObject) (result Object) {
	iv, err := _I.Get(38, "GObjectAccessible", "for_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_obj := gi.NewPointerArgument(obj.P_Object())
	args := []gi.Argument{arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_gobject_accessible_get_object
// container is not nil, container is GObjectAccessible
// is method
func (v GObjectAccessible) GetObject() (result gobject.Object) {
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
	gobject.Object
}

func WrapHyperlink(p unsafe.Pointer) (r Hyperlink) { r.P = p; return }

type IHyperlink interface{ P_Hyperlink() unsafe.Pointer }

func (v Hyperlink) P_Hyperlink() unsafe.Pointer { return v.P }

// atk_hyperlink_get_end_index
// container is not nil, container is Hyperlink
// is method
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
// container is not nil, container is Hyperlink
// is method
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
// container is not nil, container is Hyperlink
// is method
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
// container is not nil, container is Hyperlink
// is method
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
// container is not nil, container is Hyperlink
// is method
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
// container is not nil, container is Hyperlink
// is method
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

// atk_hyperlink_is_selected_link
// container is not nil, container is Hyperlink
// is method
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
// container is not nil, container is Hyperlink
// is method
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

// atk_hyperlink_impl_get_hyperlink
// container is not nil, container is HyperlinkImpl
// is method
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

// Interface Hypertext
type Hypertext struct {
	HypertextIfc
	P unsafe.Pointer
}
type HypertextIfc struct{}

// atk_hypertext_get_link
// container is not nil, container is Hypertext
// is method
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
// container is not nil, container is Hypertext
// is method
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
// container is not nil, container is Hypertext
// is method
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

// atk_image_get_image_description
// container is not nil, container is Image
// is method
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
	result = ret.String().Take()
	return
}

// atk_image_get_image_locale
// container is not nil, container is Image
// is method
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
	result = ret.String().Take()
	return
}

// atk_image_get_image_position
// container is not nil, container is Image
// is method
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
// container is not nil, container is Image
// is method
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
// container is not nil, container is Image
// is method
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

// atk_implementor_ref_accessible
// container is not nil, container is Implementor
// is method
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

// Struct KeyEventStruct
type KeyEventStruct struct {
	P unsafe.Pointer
}

// Enum KeyEventType
type KeyEventTypeEnum int

const (
	KeyEventTypePress       KeyEventTypeEnum = 0
	KeyEventTypeRelease     KeyEventTypeEnum = 1
	KeyEventTypeLastDefined KeyEventTypeEnum = 2
)

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

// Object Misc
type Misc struct {
	gobject.Object
}

func WrapMisc(p unsafe.Pointer) (r Misc) { r.P = p; return }

type IMisc interface{ P_Misc() unsafe.Pointer }

func (v Misc) P_Misc() unsafe.Pointer { return v.P }

// atk_misc_threads_enter
// container is not nil, container is Misc
// is method
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

// atk_misc_threads_leave
// container is not nil, container is Misc
// is method
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

func (v NoOpObject) P_NoOpObject() unsafe.Pointer { return v.P }

// atk_no_op_object_new
// container is not nil, container is NoOpObject
// is constructor
func NewNoOpObject(obj gobject.IObject) (result Object) {
	iv, err := _I.Get(61, "NoOpObject", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_obj := gi.NewPointerArgument(obj.P_Object())
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

// atk_no_op_object_factory_new
// container is not nil, container is NoOpObjectFactory
// is constructor
func NewNoOpObjectFactory() (result ObjectFactory) {
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
	gobject.Object
}

func WrapObject(p unsafe.Pointer) (r Object) { r.P = p; return }

type IObject interface{ P_Object() unsafe.Pointer }

func (v Object) P_Object() unsafe.Pointer { return v.P }

// atk_object_add_relationship
// container is not nil, container is Object
// is method
func (v Object) AddRelationship(relationship RelationTypeEnum, target IObject) (result bool) {
	iv, err := _I.Get(63, "Object", "add_relationship")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(target.P_Object())
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_object_get_attributes
// container is not nil, container is Object
// is method
func (v Object) GetAttributes() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(64, "Object", "get_attributes")
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

// atk_object_get_description
// container is not nil, container is Object
// is method
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
	result = ret.String().Take()
	return
}

// atk_object_get_index_in_parent
// container is not nil, container is Object
// is method
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

// atk_object_get_layer
// container is not nil, container is Object
// is method
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

// atk_object_get_mdi_zorder
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
	result = ret.String().Take()
	return
}

// atk_object_get_object_locale
// container is not nil, container is Object
// is method
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
	result = ret.String().Take()
	return
}

// atk_object_get_parent
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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

// atk_object_remove_property_change_handler
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
func (v Object) RemoveRelationship(relationship RelationTypeEnum, target IObject) (result bool) {
	iv, err := _I.Get(81, "Object", "remove_relationship")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(target.P_Object())
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_object_set_description
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
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
// container is not nil, container is Object
// is method
func (v Object) SetParent(parent IObject) {
	iv, err := _I.Get(84, "Object", "set_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(parent.P_Object())
	args := []gi.Argument{arg_v, arg_parent}
	iv.Call(args, nil, nil)
}

// atk_object_set_role
// container is not nil, container is Object
// is method
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
	gobject.Object
}

func WrapObjectFactory(p unsafe.Pointer) (r ObjectFactory) { r.P = p; return }

type IObjectFactory interface{ P_ObjectFactory() unsafe.Pointer }

func (v ObjectFactory) P_ObjectFactory() unsafe.Pointer { return v.P }

// atk_object_factory_create_accessible
// container is not nil, container is ObjectFactory
// is method
func (v ObjectFactory) CreateAccessible(obj gobject.IObject) (result Object) {
	iv, err := _I.Get(86, "ObjectFactory", "create_accessible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_obj := gi.NewPointerArgument(obj.P_Object())
	args := []gi.Argument{arg_v, arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_object_factory_get_accessible_type
// container is not nil, container is ObjectFactory
// is method
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
// container is not nil, container is ObjectFactory
// is method
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

func (v Plug) P_Plug() unsafe.Pointer { return v.P }

// atk_plug_new
// container is not nil, container is Plug
// is constructor
func NewPlug() (result Object) {
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
// container is not nil, container is Plug
// is method
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
// Struct PropertyValues
type PropertyValues struct {
	P unsafe.Pointer
}

// Struct Range
type Range struct {
	P unsafe.Pointer
}

// atk_range_new
// container is not nil, container is Range
// is constructor
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
// container is not nil, container is Range
// is method
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
// container is not nil, container is Range
// is method
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
// container is not nil, container is Range
// is method
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
	result = ret.String().Take()
	return
}

// atk_range_get_lower_limit
// container is not nil, container is Range
// is method
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
// container is not nil, container is Range
// is method
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

// Object Registry
type Registry struct {
	gobject.Object
}

func WrapRegistry(p unsafe.Pointer) (r Registry) { r.P = p; return }

type IRegistry interface{ P_Registry() unsafe.Pointer }

func (v Registry) P_Registry() unsafe.Pointer { return v.P }

// atk_registry_get_factory
// container is not nil, container is Registry
// is method
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
// container is not nil, container is Registry
// is method
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
// container is not nil, container is Registry
// is method
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
	gobject.Object
}

func WrapRelation(p unsafe.Pointer) (r Relation) { r.P = p; return }

type IRelation interface{ P_Relation() unsafe.Pointer }

func (v Relation) P_Relation() unsafe.Pointer { return v.P }

// atk_relation_new
// container is not nil, container is Relation
// is constructor
// arg 0 targets lenArgIdx 1
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
// container is not nil, container is Relation
// is method
func (v Relation) AddTarget(target IObject) {
	iv, err := _I.Get(101, "Relation", "add_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(target.P_Object())
	args := []gi.Argument{arg_v, arg_target}
	iv.Call(args, nil, nil)
}

// atk_relation_get_relation_type
// container is not nil, container is Relation
// is method
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
// container is not nil, container is Relation
// is method
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
// container is not nil, container is Relation
// is method
func (v Relation) RemoveTarget(target IObject) (result bool) {
	iv, err := _I.Get(104, "Relation", "remove_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(target.P_Object())
	args := []gi.Argument{arg_v, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct RelationClass
// Object RelationSet
type RelationSet struct {
	gobject.Object
}

func WrapRelationSet(p unsafe.Pointer) (r RelationSet) { r.P = p; return }

type IRelationSet interface{ P_RelationSet() unsafe.Pointer }

func (v RelationSet) P_RelationSet() unsafe.Pointer { return v.P }

// atk_relation_set_new
// container is not nil, container is RelationSet
// is constructor
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
// container is not nil, container is RelationSet
// is method
func (v RelationSet) Add(relation IRelation) {
	iv, err := _I.Get(106, "RelationSet", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relation := gi.NewPointerArgument(relation.P_Relation())
	args := []gi.Argument{arg_v, arg_relation}
	iv.Call(args, nil, nil)
}

// atk_relation_set_add_relation_by_type
// container is not nil, container is RelationSet
// is method
func (v RelationSet) AddRelationByType(relationship RelationTypeEnum, target IObject) {
	iv, err := _I.Get(107, "RelationSet", "add_relation_by_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(target.P_Object())
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	iv.Call(args, nil, nil)
}

// atk_relation_set_contains
// container is not nil, container is RelationSet
// is method
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
// container is not nil, container is RelationSet
// is method
func (v RelationSet) ContainsTarget(relationship RelationTypeEnum, target IObject) (result bool) {
	iv, err := _I.Get(109, "RelationSet", "contains_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(int(relationship))
	arg_target := gi.NewPointerArgument(target.P_Object())
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_relation_set_get_n_relations
// container is not nil, container is RelationSet
// is method
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
// container is not nil, container is RelationSet
// is method
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
// container is not nil, container is RelationSet
// is method
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
// container is not nil, container is RelationSet
// is method
func (v RelationSet) Remove(relation IRelation) {
	iv, err := _I.Get(113, "RelationSet", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relation := gi.NewPointerArgument(relation.P_Relation())
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

// Interface Selection
type Selection struct {
	SelectionIfc
	P unsafe.Pointer
}
type SelectionIfc struct{}

// atk_selection_add_selection
// container is not nil, container is Selection
// is method
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
// container is not nil, container is Selection
// is method
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
// container is not nil, container is Selection
// is method
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
// container is not nil, container is Selection
// is method
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
// container is not nil, container is Selection
// is method
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
// container is not nil, container is Selection
// is method
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
// container is not nil, container is Selection
// is method
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

func (v Socket) P_Socket() unsafe.Pointer { return v.P }

// atk_socket_new
// container is not nil, container is Socket
// is constructor
func NewSocket() (result Object) {
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
// container is not nil, container is Socket
// is method
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
// container is not nil, container is Socket
// is method
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
	gobject.Object
}

func WrapStateSet(p unsafe.Pointer) (r StateSet) { r.P = p; return }

type IStateSet interface{ P_StateSet() unsafe.Pointer }

func (v StateSet) P_StateSet() unsafe.Pointer { return v.P }

// atk_state_set_new
// container is not nil, container is StateSet
// is constructor
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
// container is not nil, container is StateSet
// is method
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
// container is not nil, container is StateSet
// is method
// arg 0 types lenArgIdx 1
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
// container is not nil, container is StateSet
// is method
func (v StateSet) AndSets(compare_set IStateSet) (result StateSet) {
	iv, err := _I.Get(127, "StateSet", "and_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(compare_set.P_StateSet())
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_clear_states
// container is not nil, container is StateSet
// is method
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
// container is not nil, container is StateSet
// is method
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
// container is not nil, container is StateSet
// is method
// arg 0 types lenArgIdx 1
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
// container is not nil, container is StateSet
// is method
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
// container is not nil, container is StateSet
// is method
func (v StateSet) OrSets(compare_set IStateSet) (result StateSet) {
	iv, err := _I.Get(132, "StateSet", "or_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(compare_set.P_StateSet())
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_remove_state
// container is not nil, container is StateSet
// is method
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
// container is not nil, container is StateSet
// is method
func (v StateSet) XorSets(compare_set IStateSet) (result StateSet) {
	iv, err := _I.Get(134, "StateSet", "xor_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(compare_set.P_StateSet())
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

// Interface StreamableContent
type StreamableContent struct {
	StreamableContentIfc
	P unsafe.Pointer
}
type StreamableContentIfc struct{}

// atk_streamable_content_get_mime_type
// container is not nil, container is StreamableContent
// is method
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
	result = ret.String().Take()
	return
}

// atk_streamable_content_get_n_mime_types
// container is not nil, container is StreamableContent
// is method
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
// container is not nil, container is StreamableContent
// is method
func (v *StreamableContentIfc) GetStream(mime_type string) (result glib.IOChannel) {
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
// container is not nil, container is StreamableContent
// is method
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
	result = ret.String().Take()
	return
}

// ignore GType struct StreamableContentIface
// Interface Table
type Table struct {
	TableIfc
	P unsafe.Pointer
}
type TableIfc struct{}

// atk_table_add_column_selection
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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

// atk_table_get_column_at_index
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
	result = ret.String().Take()
	return
}

// atk_table_get_column_extent_at
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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

// atk_table_get_index_at
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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

// atk_table_get_row_at_index
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
	result = ret.String().Take()
	return
}

// atk_table_get_row_extent_at
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
func (v *TableIfc) SetCaption(caption IObject) {
	iv, err := _I.Get(162, "Table", "set_caption")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_caption := gi.NewPointerArgument(caption.P_Object())
	args := []gi.Argument{arg_v, arg_caption}
	iv.Call(args, nil, nil)
}

// atk_table_set_column_description
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
func (v *TableIfc) SetColumnHeader(column int32, header IObject) {
	iv, err := _I.Get(164, "Table", "set_column_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_header := gi.NewPointerArgument(header.P_Object())
	args := []gi.Argument{arg_v, arg_column, arg_header}
	iv.Call(args, nil, nil)
}

// atk_table_set_row_description
// container is not nil, container is Table
// is method
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
// container is not nil, container is Table
// is method
func (v *TableIfc) SetRowHeader(row int32, header IObject) {
	iv, err := _I.Get(166, "Table", "set_row_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_header := gi.NewPointerArgument(header.P_Object())
	args := []gi.Argument{arg_v, arg_row, arg_header}
	iv.Call(args, nil, nil)
}

// atk_table_set_summary
// container is not nil, container is Table
// is method
func (v *TableIfc) SetSummary(accessible IObject) {
	iv, err := _I.Get(167, "Table", "set_summary")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_accessible := gi.NewPointerArgument(accessible.P_Object())
	args := []gi.Argument{arg_v, arg_accessible}
	iv.Call(args, nil, nil)
}

// Interface TableCell
type TableCell struct {
	TableCellIfc
	P unsafe.Pointer
}
type TableCellIfc struct{}

// atk_table_cell_get_column_header_cells
// container is not nil, container is TableCell
// is method
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
// container is not nil, container is TableCell
// is method
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
// container is not nil, container is TableCell
// is method
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
// container is not nil, container is TableCell
// is method
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
// container is not nil, container is TableCell
// is method
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
// container is not nil, container is TableCell
// is method
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
// container is not nil, container is TableCell
// is method
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

// atk_text_free_ranges
// container is not nil, container is Text
// is method
// arg0Type tag: array, isPtr: true
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
func (v *TextIfc) GetDefaultAttributes() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(182, "Text", "get_default_attributes")
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

// atk_text_get_n_selections
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
func (v *TextIfc) GetRangeExtents(start_offset int32, end_offset int32, coord_type CoordTypeEnum) (rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(185, "Text", "get_range_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_coord_type := gi.NewIntArgument(int(coord_type))
	arg_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset, arg_coord_type, arg_rect}
	iv.Call(args, nil, &outArgs[0])
	rect = outArgs[0].Int() /*TODO*/
	return
}

// atk_text_get_run_attributes
// container is not nil, container is Text
// is method
func (v *TextIfc) GetRunAttributes(offset int32) (result int /*TODO_TYPE isPtr: true, tag: gslist*/, start_offset int32, end_offset int32) {
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
	result = ret.Int() /*TODO*/
	return
}

// atk_text_get_selection
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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

// atk_text_get_text_after_offset
// container is not nil, container is Text
// is method
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

// atk_text_get_text_at_offset
// container is not nil, container is Text
// is method
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

// atk_text_get_text_before_offset
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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
// container is not nil, container is Text
// is method
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

// Enum TextClipType
type TextClipTypeEnum int

const (
	TextClipTypeNone TextClipTypeEnum = 0
	TextClipTypeMin  TextClipTypeEnum = 1
	TextClipTypeMax  TextClipTypeEnum = 2
	TextClipTypeBoth TextClipTypeEnum = 3
)

// Enum TextGranularity
type TextGranularityEnum int

const (
	TextGranularityChar      TextGranularityEnum = 0
	TextGranularityWord      TextGranularityEnum = 1
	TextGranularitySentence  TextGranularityEnum = 2
	TextGranularityLine      TextGranularityEnum = 3
	TextGranularityParagraph TextGranularityEnum = 4
)

// ignore GType struct TextIface
// Struct TextRange
type TextRange struct {
	P unsafe.Pointer
}

// Struct TextRectangle
type TextRectangle struct {
	P unsafe.Pointer
}

// Object Util
type Util struct {
	gobject.Object
}

func WrapUtil(p unsafe.Pointer) (r Util) { r.P = p; return }

type IUtil interface{ P_Util() unsafe.Pointer }

func (v Util) P_Util() unsafe.Pointer { return v.P }

// ignore GType struct UtilClass
// Interface Value
type Value struct {
	ValueIfc
	P unsafe.Pointer
}
type ValueIfc struct{}

// atk_value_get_current_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetCurrentValue() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(196, "Value", "get_current_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_increment
// container is not nil, container is Value
// is method
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

// atk_value_get_maximum_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetMaximumValue() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(198, "Value", "get_maximum_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_minimum_increment
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetMinimumIncrement() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(199, "Value", "get_minimum_increment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_minimum_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetMinimumValue() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(200, "Value", "get_minimum_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_range
// container is not nil, container is Value
// is method
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
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetSubRanges() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(202, "Value", "get_sub_ranges")
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

// atk_value_get_value_and_text
// container is not nil, container is Value
// is method
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

// atk_value_set_current_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) SetCurrentValue(value gobject.Value) (result bool) {
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
// container is not nil, container is Value
// is method
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

// Interface Window
type Window struct {
	WindowIfc
	P unsafe.Pointer
}
type WindowIfc struct{}

// ignore GType struct WindowIface
// atk_attribute_set_free
// container is nil
func AttributeSetFree(attrib_set int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(206, "attribute_set_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attrib_set := gi.NewIntArgument(attrib_set) /*TODO*/
	args := []gi.Argument{arg_attrib_set}
	iv.Call(args, nil, nil)
}

// atk_focus_tracker_notify
// container is nil
func FocusTrackerNotify(object IObject) {
	iv, err := _I.Get(207, "focus_tracker_notify", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_object := gi.NewPointerArgument(object.P_Object())
	args := []gi.Argument{arg_object}
	iv.Call(args, nil, nil)
}

// atk_get_binary_age
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
func GetToolkitName() (result string) {
	iv, err := _I.Get(216, "get_toolkit_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_get_toolkit_version
// container is nil
func GetToolkitVersion() (result string) {
	iv, err := _I.Get(217, "get_toolkit_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_get_version
// container is nil
func GetVersion() (result string) {
	iv, err := _I.Get(218, "get_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_relation_type_for_name
// container is nil
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
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_relation_type_register
// container is nil
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

// atk_remove_focus_tracker
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_role_get_name
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_role_register
// container is nil
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
// container is nil
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
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_state_type_register
// container is nil
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
// container is nil
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
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_text_attribute_get_value
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_text_attribute_register
// container is nil
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
// container is nil
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
// container is nil
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
	result = ret.String().Take()
	return
}

// atk_value_type_get_name
// container is nil
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
	result = ret.String().Take()
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
