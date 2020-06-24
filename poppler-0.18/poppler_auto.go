package poppler

/*
#cgo pkg-config: poppler-glib
#include <poppler.h>
extern void myPopplerAttachmentSaveFunc(gpointer buf, guint64 count, gpointer data);
static void* getPointer_myPopplerAttachmentSaveFunc() {
return (void*)(myPopplerAttachmentSaveFunc);
}
extern void myPopplerMediaSaveFunc(gpointer buf, guint64 count, gpointer data);
static void* getPointer_myPopplerMediaSaveFunc() {
return (void*)(myPopplerMediaSaveFunc);
}
*/
import "C"
import "github.com/electricface/go-gir/cairo-1.0"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Poppler")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Poppler", "0.18", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Union Action
type Action struct {
	P unsafe.Pointer
}

const SizeOfUnionAction = 32

func ActionGetType() gi.GType {
	ret := _I.GetGType(0, "Action")
	return ret
}

// poppler_action_copy
//
// [ result ] trans: everything
//
func (v Action) Copy() (result Action) {
	iv, err := _I.Get(0, "Action", "copy")
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

// poppler_action_free
//
func (v Action) Free() {
	iv, err := _I.Get(1, "Action", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct ActionAny
type ActionAny struct {
	P unsafe.Pointer
}

const SizeOfStructActionAny = 16

func ActionAnyGetType() gi.GType {
	ret := _I.GetGType(1, "ActionAny")
	return ret
}

// Struct ActionGotoDest
type ActionGotoDest struct {
	P unsafe.Pointer
}

const SizeOfStructActionGotoDest = 24

func ActionGotoDestGetType() gi.GType {
	ret := _I.GetGType(2, "ActionGotoDest")
	return ret
}

// Struct ActionGotoRemote
type ActionGotoRemote struct {
	P unsafe.Pointer
}

const SizeOfStructActionGotoRemote = 32

func ActionGotoRemoteGetType() gi.GType {
	ret := _I.GetGType(3, "ActionGotoRemote")
	return ret
}

// Struct ActionJavascript
type ActionJavascript struct {
	P unsafe.Pointer
}

const SizeOfStructActionJavascript = 24

func ActionJavascriptGetType() gi.GType {
	ret := _I.GetGType(4, "ActionJavascript")
	return ret
}

// Struct ActionLaunch
type ActionLaunch struct {
	P unsafe.Pointer
}

const SizeOfStructActionLaunch = 32

func ActionLaunchGetType() gi.GType {
	ret := _I.GetGType(5, "ActionLaunch")
	return ret
}

// Struct ActionLayer
type ActionLayer struct {
	P unsafe.Pointer
}

const SizeOfStructActionLayer = 16

func ActionLayerGetType() gi.GType {
	ret := _I.GetGType(6, "ActionLayer")
	return ret
}

// Enum ActionLayerAction
type ActionLayerActionEnum int

const (
	ActionLayerActionOn     ActionLayerActionEnum = 0
	ActionLayerActionOff    ActionLayerActionEnum = 1
	ActionLayerActionToggle ActionLayerActionEnum = 2
)

func ActionLayerActionGetType() gi.GType {
	ret := _I.GetGType(7, "ActionLayerAction")
	return ret
}

// Struct ActionMovie
type ActionMovie struct {
	P unsafe.Pointer
}

const SizeOfStructActionMovie = 32

func ActionMovieGetType() gi.GType {
	ret := _I.GetGType(8, "ActionMovie")
	return ret
}

// Enum ActionMovieOperation
type ActionMovieOperationEnum int

const (
	ActionMovieOperationPlay   ActionMovieOperationEnum = 0
	ActionMovieOperationPause  ActionMovieOperationEnum = 1
	ActionMovieOperationResume ActionMovieOperationEnum = 2
	ActionMovieOperationStop   ActionMovieOperationEnum = 3
)

func ActionMovieOperationGetType() gi.GType {
	ret := _I.GetGType(9, "ActionMovieOperation")
	return ret
}

// Struct ActionNamed
type ActionNamed struct {
	P unsafe.Pointer
}

const SizeOfStructActionNamed = 24

func ActionNamedGetType() gi.GType {
	ret := _I.GetGType(10, "ActionNamed")
	return ret
}

// Struct ActionOCGState
type ActionOCGState struct {
	P unsafe.Pointer
}

const SizeOfStructActionOCGState = 24

func ActionOCGStateGetType() gi.GType {
	ret := _I.GetGType(11, "ActionOCGState")
	return ret
}

// Struct ActionRendition
type ActionRendition struct {
	P unsafe.Pointer
}

const SizeOfStructActionRendition = 32

func ActionRenditionGetType() gi.GType {
	ret := _I.GetGType(12, "ActionRendition")
	return ret
}

// Enum ActionType
type ActionTypeEnum int

const (
	ActionTypeUnknown    ActionTypeEnum = 0
	ActionTypeNone       ActionTypeEnum = 1
	ActionTypeGotoDest   ActionTypeEnum = 2
	ActionTypeGotoRemote ActionTypeEnum = 3
	ActionTypeLaunch     ActionTypeEnum = 4
	ActionTypeUri        ActionTypeEnum = 5
	ActionTypeNamed      ActionTypeEnum = 6
	ActionTypeMovie      ActionTypeEnum = 7
	ActionTypeRendition  ActionTypeEnum = 8
	ActionTypeOcgState   ActionTypeEnum = 9
	ActionTypeJavascript ActionTypeEnum = 10
)

func ActionTypeGetType() gi.GType {
	ret := _I.GetGType(13, "ActionType")
	return ret
}

// Struct ActionUri
type ActionUri struct {
	P unsafe.Pointer
}

const SizeOfStructActionUri = 24

func ActionUriGetType() gi.GType {
	ret := _I.GetGType(14, "ActionUri")
	return ret
}

// Object Annot
type Annot struct {
	g.Object
}

func WrapAnnot(p unsafe.Pointer) (r Annot) { r.P = p; return }

type IAnnot interface{ P_Annot() unsafe.Pointer }

func (v Annot) P_Annot() unsafe.Pointer { return v.P }
func AnnotGetType() gi.GType {
	ret := _I.GetGType(15, "Annot")
	return ret
}

// poppler_annot_get_annot_type
//
// [ result ] trans: nothing
//
func (v Annot) GetAnnotType() (result AnnotTypeEnum) {
	iv, err := _I.Get(2, "Annot", "get_annot_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AnnotTypeEnum(ret.Int())
	return
}

// poppler_annot_get_color
//
// [ result ] trans: everything
//
func (v Annot) GetColor() (result Color) {
	iv, err := _I.Get(3, "Annot", "get_color")
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

// poppler_annot_get_contents
//
// [ result ] trans: everything
//
func (v Annot) GetContents() (result string) {
	iv, err := _I.Get(4, "Annot", "get_contents")
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

// poppler_annot_get_flags
//
// [ result ] trans: nothing
//
func (v Annot) GetFlags() (result AnnotFlagFlags) {
	iv, err := _I.Get(5, "Annot", "get_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AnnotFlagFlags(ret.Int())
	return
}

// poppler_annot_get_modified
//
// [ result ] trans: everything
//
func (v Annot) GetModified() (result string) {
	iv, err := _I.Get(6, "Annot", "get_modified")
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

// poppler_annot_get_name
//
// [ result ] trans: everything
//
func (v Annot) GetName() (result string) {
	iv, err := _I.Get(7, "Annot", "get_name")
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

// poppler_annot_get_page_index
//
// [ result ] trans: nothing
//
func (v Annot) GetPageIndex() (result int32) {
	iv, err := _I.Get(8, "Annot", "get_page_index")
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

// poppler_annot_get_rectangle
//
// [ poppler_rect ] trans: nothing, dir: out
//
func (v Annot) GetRectangle(poppler_rect Rectangle) {
	iv, err := _I.Get(9, "Annot", "get_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_rect := gi.NewPointerArgument(poppler_rect.P)
	args := []gi.Argument{arg_v, arg_poppler_rect}
	iv.Call(args, nil, nil)
}

// poppler_annot_set_color
//
// [ poppler_color ] trans: nothing
//
func (v Annot) SetColor(poppler_color Color) {
	iv, err := _I.Get(10, "Annot", "set_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_color := gi.NewPointerArgument(poppler_color.P)
	args := []gi.Argument{arg_v, arg_poppler_color}
	iv.Call(args, nil, nil)
}

// poppler_annot_set_contents
//
// [ contents ] trans: nothing
//
func (v Annot) SetContents(contents string) {
	iv, err := _I.Get(11, "Annot", "set_contents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_contents := gi.CString(contents)
	arg_v := gi.NewPointerArgument(v.P)
	arg_contents := gi.NewStringArgument(c_contents)
	args := []gi.Argument{arg_v, arg_contents}
	iv.Call(args, nil, nil)
	gi.Free(c_contents)
}

// poppler_annot_set_flags
//
// [ flags ] trans: nothing
//
func (v Annot) SetFlags(flags AnnotFlagFlags) {
	iv, err := _I.Get(12, "Annot", "set_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, nil)
}

// poppler_annot_set_rectangle
//
// [ poppler_rect ] trans: nothing
//
func (v Annot) SetRectangle(poppler_rect Rectangle) {
	iv, err := _I.Get(13, "Annot", "set_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_rect := gi.NewPointerArgument(poppler_rect.P)
	args := []gi.Argument{arg_v, arg_poppler_rect}
	iv.Call(args, nil, nil)
}

// Struct AnnotCalloutLine
type AnnotCalloutLine struct {
	P unsafe.Pointer
}

const SizeOfStructAnnotCalloutLine = 56

func AnnotCalloutLineGetType() gi.GType {
	ret := _I.GetGType(16, "AnnotCalloutLine")
	return ret
}

// poppler_annot_callout_line_new
//
// [ result ] trans: everything
//
func NewAnnotCalloutLine() (result AnnotCalloutLine) {
	iv, err := _I.Get(14, "AnnotCalloutLine", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_callout_line_copy
//
// [ result ] trans: everything
//
func (v AnnotCalloutLine) Copy() (result AnnotCalloutLine) {
	iv, err := _I.Get(15, "AnnotCalloutLine", "copy")
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

// poppler_annot_callout_line_free
//
func (v AnnotCalloutLine) Free() {
	iv, err := _I.Get(16, "AnnotCalloutLine", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object AnnotCircle
type AnnotCircle struct {
	AnnotMarkup
}

func WrapAnnotCircle(p unsafe.Pointer) (r AnnotCircle) { r.P = p; return }

type IAnnotCircle interface{ P_AnnotCircle() unsafe.Pointer }

func (v AnnotCircle) P_AnnotCircle() unsafe.Pointer { return v.P }
func AnnotCircleGetType() gi.GType {
	ret := _I.GetGType(17, "AnnotCircle")
	return ret
}

// poppler_annot_circle_new
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotCircle(doc IDocument, rect Rectangle) (result AnnotCircle) {
	iv, err := _I.Get(17, "AnnotCircle", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_doc, arg_rect}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_circle_get_interior_color
//
// [ result ] trans: everything
//
func (v AnnotCircle) GetInteriorColor() (result Color) {
	iv, err := _I.Get(18, "AnnotCircle", "get_interior_color")
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

// poppler_annot_circle_set_interior_color
//
// [ poppler_color ] trans: nothing
//
func (v AnnotCircle) SetInteriorColor(poppler_color Color) {
	iv, err := _I.Get(19, "AnnotCircle", "set_interior_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_color := gi.NewPointerArgument(poppler_color.P)
	args := []gi.Argument{arg_v, arg_poppler_color}
	iv.Call(args, nil, nil)
}

// Enum AnnotExternalDataType
type AnnotExternalDataTypeEnum int

const (
	AnnotExternalDataType3d      AnnotExternalDataTypeEnum = 0
	AnnotExternalDataTypeUnknown AnnotExternalDataTypeEnum = 1
)

func AnnotExternalDataTypeGetType() gi.GType {
	ret := _I.GetGType(18, "AnnotExternalDataType")
	return ret
}

// Object AnnotFileAttachment
type AnnotFileAttachment struct {
	AnnotMarkup
}

func WrapAnnotFileAttachment(p unsafe.Pointer) (r AnnotFileAttachment) { r.P = p; return }

type IAnnotFileAttachment interface{ P_AnnotFileAttachment() unsafe.Pointer }

func (v AnnotFileAttachment) P_AnnotFileAttachment() unsafe.Pointer { return v.P }
func AnnotFileAttachmentGetType() gi.GType {
	ret := _I.GetGType(19, "AnnotFileAttachment")
	return ret
}

// poppler_annot_file_attachment_get_attachment
//
// [ result ] trans: everything
//
func (v AnnotFileAttachment) GetAttachment() (result Attachment) {
	iv, err := _I.Get(20, "AnnotFileAttachment", "get_attachment")
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

// poppler_annot_file_attachment_get_name
//
// [ result ] trans: everything
//
func (v AnnotFileAttachment) GetName() (result string) {
	iv, err := _I.Get(21, "AnnotFileAttachment", "get_name")
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

// Flags AnnotFlag
type AnnotFlagFlags int

const (
	AnnotFlagUnknown        AnnotFlagFlags = 0
	AnnotFlagInvisible      AnnotFlagFlags = 1
	AnnotFlagHidden         AnnotFlagFlags = 2
	AnnotFlagPrint          AnnotFlagFlags = 4
	AnnotFlagNoZoom         AnnotFlagFlags = 8
	AnnotFlagNoRotate       AnnotFlagFlags = 16
	AnnotFlagNoView         AnnotFlagFlags = 32
	AnnotFlagReadOnly       AnnotFlagFlags = 64
	AnnotFlagLocked         AnnotFlagFlags = 128
	AnnotFlagToggleNoView   AnnotFlagFlags = 256
	AnnotFlagLockedContents AnnotFlagFlags = 512
)

func AnnotFlagGetType() gi.GType {
	ret := _I.GetGType(20, "AnnotFlag")
	return ret
}

// Object AnnotFreeText
type AnnotFreeText struct {
	AnnotMarkup
}

func WrapAnnotFreeText(p unsafe.Pointer) (r AnnotFreeText) { r.P = p; return }

type IAnnotFreeText interface{ P_AnnotFreeText() unsafe.Pointer }

func (v AnnotFreeText) P_AnnotFreeText() unsafe.Pointer { return v.P }
func AnnotFreeTextGetType() gi.GType {
	ret := _I.GetGType(21, "AnnotFreeText")
	return ret
}

// poppler_annot_free_text_get_callout_line
//
// [ result ] trans: everything
//
func (v AnnotFreeText) GetCalloutLine() (result AnnotCalloutLine) {
	iv, err := _I.Get(22, "AnnotFreeText", "get_callout_line")
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

// poppler_annot_free_text_get_quadding
//
// [ result ] trans: nothing
//
func (v AnnotFreeText) GetQuadding() (result AnnotFreeTextQuaddingEnum) {
	iv, err := _I.Get(23, "AnnotFreeText", "get_quadding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AnnotFreeTextQuaddingEnum(ret.Int())
	return
}

// Enum AnnotFreeTextQuadding
type AnnotFreeTextQuaddingEnum int

const (
	AnnotFreeTextQuaddingLeftJustified  AnnotFreeTextQuaddingEnum = 0
	AnnotFreeTextQuaddingCentered       AnnotFreeTextQuaddingEnum = 1
	AnnotFreeTextQuaddingRightJustified AnnotFreeTextQuaddingEnum = 2
)

func AnnotFreeTextQuaddingGetType() gi.GType {
	ret := _I.GetGType(22, "AnnotFreeTextQuadding")
	return ret
}

// Object AnnotLine
type AnnotLine struct {
	AnnotMarkup
}

func WrapAnnotLine(p unsafe.Pointer) (r AnnotLine) { r.P = p; return }

type IAnnotLine interface{ P_AnnotLine() unsafe.Pointer }

func (v AnnotLine) P_AnnotLine() unsafe.Pointer { return v.P }
func AnnotLineGetType() gi.GType {
	ret := _I.GetGType(23, "AnnotLine")
	return ret
}

// poppler_annot_line_new
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotLine(doc IDocument, rect Rectangle, start Point, end Point) (result AnnotLine) {
	iv, err := _I.Get(24, "AnnotLine", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_doc, arg_rect, arg_start, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_line_set_vertices
//
// [ start ] trans: nothing
//
// [ end ] trans: nothing
//
func (v AnnotLine) SetVertices(start Point, end Point) {
	iv, err := _I.Get(25, "AnnotLine", "set_vertices")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// Struct AnnotMapping
type AnnotMapping struct {
	P unsafe.Pointer
}

const SizeOfStructAnnotMapping = 40

func AnnotMappingGetType() gi.GType {
	ret := _I.GetGType(24, "AnnotMapping")
	return ret
}

// poppler_annot_mapping_new
//
// [ result ] trans: everything
//
func NewAnnotMapping() (result AnnotMapping) {
	iv, err := _I.Get(26, "AnnotMapping", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_mapping_copy
//
// [ result ] trans: everything
//
func (v AnnotMapping) Copy() (result AnnotMapping) {
	iv, err := _I.Get(27, "AnnotMapping", "copy")
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

// poppler_annot_mapping_free
//
func (v AnnotMapping) Free() {
	iv, err := _I.Get(28, "AnnotMapping", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object AnnotMarkup
type AnnotMarkup struct {
	Annot
}

func WrapAnnotMarkup(p unsafe.Pointer) (r AnnotMarkup) { r.P = p; return }

type IAnnotMarkup interface{ P_AnnotMarkup() unsafe.Pointer }

func (v AnnotMarkup) P_AnnotMarkup() unsafe.Pointer { return v.P }
func AnnotMarkupGetType() gi.GType {
	ret := _I.GetGType(25, "AnnotMarkup")
	return ret
}

// poppler_annot_markup_get_date
//
// [ result ] trans: everything
//
func (v AnnotMarkup) GetDate() (result g.Date) {
	iv, err := _I.Get(29, "AnnotMarkup", "get_date")
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

// poppler_annot_markup_get_external_data
//
// [ result ] trans: nothing
//
func (v AnnotMarkup) GetExternalData() (result AnnotExternalDataTypeEnum) {
	iv, err := _I.Get(30, "AnnotMarkup", "get_external_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AnnotExternalDataTypeEnum(ret.Int())
	return
}

// poppler_annot_markup_get_label
//
// [ result ] trans: everything
//
func (v AnnotMarkup) GetLabel() (result string) {
	iv, err := _I.Get(31, "AnnotMarkup", "get_label")
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

// poppler_annot_markup_get_opacity
//
// [ result ] trans: nothing
//
func (v AnnotMarkup) GetOpacity() (result float64) {
	iv, err := _I.Get(32, "AnnotMarkup", "get_opacity")
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

// poppler_annot_markup_get_popup_is_open
//
// [ result ] trans: nothing
//
func (v AnnotMarkup) GetPopupIsOpen() (result bool) {
	iv, err := _I.Get(33, "AnnotMarkup", "get_popup_is_open")
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

// poppler_annot_markup_get_popup_rectangle
//
// [ poppler_rect ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v AnnotMarkup) GetPopupRectangle(poppler_rect Rectangle) (result bool) {
	iv, err := _I.Get(34, "AnnotMarkup", "get_popup_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_rect := gi.NewPointerArgument(poppler_rect.P)
	args := []gi.Argument{arg_v, arg_poppler_rect}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_annot_markup_get_reply_to
//
// [ result ] trans: nothing
//
func (v AnnotMarkup) GetReplyTo() (result AnnotMarkupReplyTypeEnum) {
	iv, err := _I.Get(35, "AnnotMarkup", "get_reply_to")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AnnotMarkupReplyTypeEnum(ret.Int())
	return
}

// poppler_annot_markup_get_subject
//
// [ result ] trans: everything
//
func (v AnnotMarkup) GetSubject() (result string) {
	iv, err := _I.Get(36, "AnnotMarkup", "get_subject")
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

// poppler_annot_markup_has_popup
//
// [ result ] trans: nothing
//
func (v AnnotMarkup) HasPopup() (result bool) {
	iv, err := _I.Get(37, "AnnotMarkup", "has_popup")
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

// poppler_annot_markup_set_label
//
// [ label ] trans: nothing
//
func (v AnnotMarkup) SetLabel(label string) {
	iv, err := _I.Get(38, "AnnotMarkup", "set_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_label := gi.CString(label)
	arg_v := gi.NewPointerArgument(v.P)
	arg_label := gi.NewStringArgument(c_label)
	args := []gi.Argument{arg_v, arg_label}
	iv.Call(args, nil, nil)
	gi.Free(c_label)
}

// poppler_annot_markup_set_opacity
//
// [ opacity ] trans: nothing
//
func (v AnnotMarkup) SetOpacity(opacity float64) {
	iv, err := _I.Get(39, "AnnotMarkup", "set_opacity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_opacity := gi.NewDoubleArgument(opacity)
	args := []gi.Argument{arg_v, arg_opacity}
	iv.Call(args, nil, nil)
}

// poppler_annot_markup_set_popup
//
// [ popup_rect ] trans: nothing
//
func (v AnnotMarkup) SetPopup(popup_rect Rectangle) {
	iv, err := _I.Get(40, "AnnotMarkup", "set_popup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_popup_rect := gi.NewPointerArgument(popup_rect.P)
	args := []gi.Argument{arg_v, arg_popup_rect}
	iv.Call(args, nil, nil)
}

// poppler_annot_markup_set_popup_is_open
//
// [ is_open ] trans: nothing
//
func (v AnnotMarkup) SetPopupIsOpen(is_open bool) {
	iv, err := _I.Get(41, "AnnotMarkup", "set_popup_is_open")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_is_open := gi.NewBoolArgument(is_open)
	args := []gi.Argument{arg_v, arg_is_open}
	iv.Call(args, nil, nil)
}

// poppler_annot_markup_set_popup_rectangle
//
// [ poppler_rect ] trans: nothing
//
func (v AnnotMarkup) SetPopupRectangle(poppler_rect Rectangle) {
	iv, err := _I.Get(42, "AnnotMarkup", "set_popup_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_rect := gi.NewPointerArgument(poppler_rect.P)
	args := []gi.Argument{arg_v, arg_poppler_rect}
	iv.Call(args, nil, nil)
}

// Enum AnnotMarkupReplyType
type AnnotMarkupReplyTypeEnum int

const (
	AnnotMarkupReplyTypeR     AnnotMarkupReplyTypeEnum = 0
	AnnotMarkupReplyTypeGroup AnnotMarkupReplyTypeEnum = 1
)

func AnnotMarkupReplyTypeGetType() gi.GType {
	ret := _I.GetGType(26, "AnnotMarkupReplyType")
	return ret
}

// Object AnnotMovie
type AnnotMovie struct {
	Annot
}

func WrapAnnotMovie(p unsafe.Pointer) (r AnnotMovie) { r.P = p; return }

type IAnnotMovie interface{ P_AnnotMovie() unsafe.Pointer }

func (v AnnotMovie) P_AnnotMovie() unsafe.Pointer { return v.P }
func AnnotMovieGetType() gi.GType {
	ret := _I.GetGType(27, "AnnotMovie")
	return ret
}

// poppler_annot_movie_get_movie
//
// [ result ] trans: nothing
//
func (v AnnotMovie) GetMovie() (result Movie) {
	iv, err := _I.Get(43, "AnnotMovie", "get_movie")
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

// poppler_annot_movie_get_title
//
// [ result ] trans: everything
//
func (v AnnotMovie) GetTitle() (result string) {
	iv, err := _I.Get(44, "AnnotMovie", "get_title")
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

// Object AnnotScreen
type AnnotScreen struct {
	Annot
}

func WrapAnnotScreen(p unsafe.Pointer) (r AnnotScreen) { r.P = p; return }

type IAnnotScreen interface{ P_AnnotScreen() unsafe.Pointer }

func (v AnnotScreen) P_AnnotScreen() unsafe.Pointer { return v.P }
func AnnotScreenGetType() gi.GType {
	ret := _I.GetGType(28, "AnnotScreen")
	return ret
}

// poppler_annot_screen_get_action
//
// [ result ] trans: nothing
//
func (v AnnotScreen) GetAction() (result Action) {
	iv, err := _I.Get(45, "AnnotScreen", "get_action")
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

// Object AnnotSquare
type AnnotSquare struct {
	AnnotMarkup
}

func WrapAnnotSquare(p unsafe.Pointer) (r AnnotSquare) { r.P = p; return }

type IAnnotSquare interface{ P_AnnotSquare() unsafe.Pointer }

func (v AnnotSquare) P_AnnotSquare() unsafe.Pointer { return v.P }
func AnnotSquareGetType() gi.GType {
	ret := _I.GetGType(29, "AnnotSquare")
	return ret
}

// poppler_annot_square_new
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotSquare(doc IDocument, rect Rectangle) (result AnnotSquare) {
	iv, err := _I.Get(46, "AnnotSquare", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_doc, arg_rect}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_square_get_interior_color
//
// [ result ] trans: everything
//
func (v AnnotSquare) GetInteriorColor() (result Color) {
	iv, err := _I.Get(47, "AnnotSquare", "get_interior_color")
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

// poppler_annot_square_set_interior_color
//
// [ poppler_color ] trans: nothing
//
func (v AnnotSquare) SetInteriorColor(poppler_color Color) {
	iv, err := _I.Get(48, "AnnotSquare", "set_interior_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_poppler_color := gi.NewPointerArgument(poppler_color.P)
	args := []gi.Argument{arg_v, arg_poppler_color}
	iv.Call(args, nil, nil)
}

// Object AnnotText
type AnnotText struct {
	AnnotMarkup
}

func WrapAnnotText(p unsafe.Pointer) (r AnnotText) { r.P = p; return }

type IAnnotText interface{ P_AnnotText() unsafe.Pointer }

func (v AnnotText) P_AnnotText() unsafe.Pointer { return v.P }
func AnnotTextGetType() gi.GType {
	ret := _I.GetGType(30, "AnnotText")
	return ret
}

// poppler_annot_text_new
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotText(doc IDocument, rect Rectangle) (result AnnotText) {
	iv, err := _I.Get(49, "AnnotText", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_doc, arg_rect}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_text_get_icon
//
// [ result ] trans: everything
//
func (v AnnotText) GetIcon() (result string) {
	iv, err := _I.Get(50, "AnnotText", "get_icon")
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

// poppler_annot_text_get_is_open
//
// [ result ] trans: nothing
//
func (v AnnotText) GetIsOpen() (result bool) {
	iv, err := _I.Get(51, "AnnotText", "get_is_open")
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

// poppler_annot_text_get_state
//
// [ result ] trans: nothing
//
func (v AnnotText) GetState() (result AnnotTextStateEnum) {
	iv, err := _I.Get(52, "AnnotText", "get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AnnotTextStateEnum(ret.Int())
	return
}

// poppler_annot_text_set_icon
//
// [ icon ] trans: nothing
//
func (v AnnotText) SetIcon(icon string) {
	iv, err := _I.Get(53, "AnnotText", "set_icon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_icon := gi.CString(icon)
	arg_v := gi.NewPointerArgument(v.P)
	arg_icon := gi.NewStringArgument(c_icon)
	args := []gi.Argument{arg_v, arg_icon}
	iv.Call(args, nil, nil)
	gi.Free(c_icon)
}

// poppler_annot_text_set_is_open
//
// [ is_open ] trans: nothing
//
func (v AnnotText) SetIsOpen(is_open bool) {
	iv, err := _I.Get(54, "AnnotText", "set_is_open")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_is_open := gi.NewBoolArgument(is_open)
	args := []gi.Argument{arg_v, arg_is_open}
	iv.Call(args, nil, nil)
}

// Object AnnotTextMarkup
type AnnotTextMarkup struct {
	AnnotMarkup
}

func WrapAnnotTextMarkup(p unsafe.Pointer) (r AnnotTextMarkup) { r.P = p; return }

type IAnnotTextMarkup interface{ P_AnnotTextMarkup() unsafe.Pointer }

func (v AnnotTextMarkup) P_AnnotTextMarkup() unsafe.Pointer { return v.P }
func AnnotTextMarkupGetType() gi.GType {
	ret := _I.GetGType(31, "AnnotTextMarkup")
	return ret
}

// poppler_annot_text_markup_new_highlight
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ quadrilaterals ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotTextMarkupHighlight(doc IDocument, rect Rectangle, quadrilaterals int /*TODO_TYPE isPtr: true, tag: array*/) (result AnnotTextMarkup) {
	iv, err := _I.Get(55, "AnnotTextMarkup", "new_highlight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_quadrilaterals := gi.NewIntArgument(quadrilaterals) /*TODO*/
	args := []gi.Argument{arg_doc, arg_rect, arg_quadrilaterals}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_text_markup_new_squiggly
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ quadrilaterals ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotTextMarkupSquiggly(doc IDocument, rect Rectangle, quadrilaterals int /*TODO_TYPE isPtr: true, tag: array*/) (result AnnotTextMarkup) {
	iv, err := _I.Get(56, "AnnotTextMarkup", "new_squiggly")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_quadrilaterals := gi.NewIntArgument(quadrilaterals) /*TODO*/
	args := []gi.Argument{arg_doc, arg_rect, arg_quadrilaterals}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_text_markup_new_strikeout
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ quadrilaterals ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotTextMarkupStrikeout(doc IDocument, rect Rectangle, quadrilaterals int /*TODO_TYPE isPtr: true, tag: array*/) (result AnnotTextMarkup) {
	iv, err := _I.Get(57, "AnnotTextMarkup", "new_strikeout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_quadrilaterals := gi.NewIntArgument(quadrilaterals) /*TODO*/
	args := []gi.Argument{arg_doc, arg_rect, arg_quadrilaterals}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_text_markup_new_underline
//
// [ doc ] trans: nothing
//
// [ rect ] trans: nothing
//
// [ quadrilaterals ] trans: nothing
//
// [ result ] trans: everything
//
func NewAnnotTextMarkupUnderline(doc IDocument, rect Rectangle, quadrilaterals int /*TODO_TYPE isPtr: true, tag: array*/) (result AnnotTextMarkup) {
	iv, err := _I.Get(58, "AnnotTextMarkup", "new_underline")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if doc != nil {
		tmp = doc.P_Document()
	}
	arg_doc := gi.NewPointerArgument(tmp)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_quadrilaterals := gi.NewIntArgument(quadrilaterals) /*TODO*/
	args := []gi.Argument{arg_doc, arg_rect, arg_quadrilaterals}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_annot_text_markup_get_quadrilaterals
//
// [ result ] trans: everything
//
func (v AnnotTextMarkup) GetQuadrilaterals() (result int /*TODO_TYPE array type: 1, isZeroTerm: false*/) {
	iv, err := _I.Get(59, "AnnotTextMarkup", "get_quadrilaterals")
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

// poppler_annot_text_markup_set_quadrilaterals
//
// [ quadrilaterals ] trans: nothing
//
func (v AnnotTextMarkup) SetQuadrilaterals(quadrilaterals int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(60, "AnnotTextMarkup", "set_quadrilaterals")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quadrilaterals := gi.NewIntArgument(quadrilaterals) /*TODO*/
	args := []gi.Argument{arg_v, arg_quadrilaterals}
	iv.Call(args, nil, nil)
}

// Enum AnnotTextState
type AnnotTextStateEnum int

const (
	AnnotTextStateMarked    AnnotTextStateEnum = 0
	AnnotTextStateUnmarked  AnnotTextStateEnum = 1
	AnnotTextStateAccepted  AnnotTextStateEnum = 2
	AnnotTextStateRejected  AnnotTextStateEnum = 3
	AnnotTextStateCancelled AnnotTextStateEnum = 4
	AnnotTextStateCompleted AnnotTextStateEnum = 5
	AnnotTextStateNone      AnnotTextStateEnum = 6
	AnnotTextStateUnknown   AnnotTextStateEnum = 7
)

func AnnotTextStateGetType() gi.GType {
	ret := _I.GetGType(32, "AnnotTextState")
	return ret
}

// Enum AnnotType
type AnnotTypeEnum int

const (
	AnnotTypeUnknown        AnnotTypeEnum = 0
	AnnotTypeText           AnnotTypeEnum = 1
	AnnotTypeLink           AnnotTypeEnum = 2
	AnnotTypeFreeText       AnnotTypeEnum = 3
	AnnotTypeLine           AnnotTypeEnum = 4
	AnnotTypeSquare         AnnotTypeEnum = 5
	AnnotTypeCircle         AnnotTypeEnum = 6
	AnnotTypePolygon        AnnotTypeEnum = 7
	AnnotTypePolyLine       AnnotTypeEnum = 8
	AnnotTypeHighlight      AnnotTypeEnum = 9
	AnnotTypeUnderline      AnnotTypeEnum = 10
	AnnotTypeSquiggly       AnnotTypeEnum = 11
	AnnotTypeStrikeOut      AnnotTypeEnum = 12
	AnnotTypeStamp          AnnotTypeEnum = 13
	AnnotTypeCaret          AnnotTypeEnum = 14
	AnnotTypeInk            AnnotTypeEnum = 15
	AnnotTypePopup          AnnotTypeEnum = 16
	AnnotTypeFileAttachment AnnotTypeEnum = 17
	AnnotTypeSound          AnnotTypeEnum = 18
	AnnotTypeMovie          AnnotTypeEnum = 19
	AnnotTypeWidget         AnnotTypeEnum = 20
	AnnotTypeScreen         AnnotTypeEnum = 21
	AnnotTypePrinterMark    AnnotTypeEnum = 22
	AnnotTypeTrapNet        AnnotTypeEnum = 23
	AnnotTypeWatermark      AnnotTypeEnum = 24
	AnnotType3d             AnnotTypeEnum = 25
)

func AnnotTypeGetType() gi.GType {
	ret := _I.GetGType(33, "AnnotType")
	return ret
}

// Object Attachment
type Attachment struct {
	g.Object
}

func WrapAttachment(p unsafe.Pointer) (r Attachment) { r.P = p; return }

type IAttachment interface{ P_Attachment() unsafe.Pointer }

func (v Attachment) P_Attachment() unsafe.Pointer { return v.P }
func AttachmentGetType() gi.GType {
	ret := _I.GetGType(34, "Attachment")
	return ret
}

// poppler_attachment_save
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Attachment) Save(filename string) (result bool, err error) {
	iv, err := _I.Get(61, "Attachment", "save")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// poppler_attachment_save_to_callback
//
// [ save_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Attachment) SaveToCallback(save_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool, err error) {
	iv, err := _I.Get(62, "Attachment", "save_to_callback")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_save_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myAttachmentSaveFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_save_func, arg_user_data, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// ignore GType struct AttachmentClass

type AttachmentSaveFuncStruct struct {
	F_buf   gi.Uint8Array
	F_count uint64
	F_data  unsafe.Pointer
}

func GetPointer_myAttachmentSaveFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myPopplerAttachmentSaveFunc())
}

//export myPopplerAttachmentSaveFunc
func myPopplerAttachmentSaveFunc(buf C.gpointer, count C.guint64, data C.gpointer) {
	// TODO: not found user_data
}

// Enum Backend
type BackendEnum int

const (
	BackendUnknown BackendEnum = 0
	BackendSplash  BackendEnum = 1
	BackendCairo   BackendEnum = 2
)

func BackendGetType() gi.GType {
	ret := _I.GetGType(35, "Backend")
	return ret
}

// Struct Color
type Color struct {
	P unsafe.Pointer
}

const SizeOfStructColor = 6

func ColorGetType() gi.GType {
	ret := _I.GetGType(36, "Color")
	return ret
}

// poppler_color_new
//
// [ result ] trans: everything
//
func NewColor() (result Color) {
	iv, err := _I.Get(63, "Color", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_color_copy
//
// [ result ] trans: everything
//
func (v Color) Copy() (result Color) {
	iv, err := _I.Get(64, "Color", "copy")
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

// poppler_color_free
//
func (v Color) Free() {
	iv, err := _I.Get(65, "Color", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Dest
type Dest struct {
	P unsafe.Pointer
}

const SizeOfStructDest = 72

func DestGetType() gi.GType {
	ret := _I.GetGType(37, "Dest")
	return ret
}

// poppler_dest_copy
//
// [ result ] trans: everything
//
func (v Dest) Copy() (result Dest) {
	iv, err := _I.Get(66, "Dest", "copy")
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

// poppler_dest_free
//
func (v Dest) Free() {
	iv, err := _I.Get(67, "Dest", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum DestType
type DestTypeEnum int

const (
	DestTypeUnknown DestTypeEnum = 0
	DestTypeXyz     DestTypeEnum = 1
	DestTypeFit     DestTypeEnum = 2
	DestTypeFith    DestTypeEnum = 3
	DestTypeFitv    DestTypeEnum = 4
	DestTypeFitr    DestTypeEnum = 5
	DestTypeFitb    DestTypeEnum = 6
	DestTypeFitbh   DestTypeEnum = 7
	DestTypeFitbv   DestTypeEnum = 8
	DestTypeNamed   DestTypeEnum = 9
)

func DestTypeGetType() gi.GType {
	ret := _I.GetGType(38, "DestType")
	return ret
}

// Object Document
type Document struct {
	g.Object
}

func WrapDocument(p unsafe.Pointer) (r Document) { r.P = p; return }

type IDocument interface{ P_Document() unsafe.Pointer }

func (v Document) P_Document() unsafe.Pointer { return v.P }
func DocumentGetType() gi.GType {
	ret := _I.GetGType(39, "Document")
	return ret
}

// poppler_document_new_from_data
//
// [ data ] trans: nothing
//
// [ length ] trans: nothing
//
// [ password ] trans: nothing
//
// [ result ] trans: everything
//
func NewDocumentFromData(data string, length int32, password string) (result Document, err error) {
	iv, err := _I.Get(68, "Document", "new_from_data")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_data := gi.CString(data)
	c_password := gi.CString(password)
	arg_data := gi.NewStringArgument(c_data)
	arg_length := gi.NewInt32Argument(length)
	arg_password := gi.NewStringArgument(c_password)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_data, arg_length, arg_password, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_data)
	gi.Free(c_password)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// poppler_document_new_from_file
//
// [ uri ] trans: nothing
//
// [ password ] trans: nothing
//
// [ result ] trans: everything
//
func NewDocumentFromFile(uri string, password string) (result Document, err error) {
	iv, err := _I.Get(69, "Document", "new_from_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_password := gi.CString(password)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_password := gi.NewStringArgument(c_password)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_uri, arg_password, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_password)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// poppler_document_new_from_gfile
//
// [ file ] trans: nothing
//
// [ password ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewDocumentFromGfile(file g.IFile, password string, cancellable g.ICancellable) (result Document, err error) {
	iv, err := _I.Get(70, "Document", "new_from_gfile")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if file != nil {
		tmp = file.P_File()
	}
	c_password := gi.CString(password)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_file := gi.NewPointerArgument(tmp)
	arg_password := gi.NewStringArgument(c_password)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_file, arg_password, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_password)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// poppler_document_new_from_stream
//
// [ stream ] trans: nothing
//
// [ length ] trans: nothing
//
// [ password ] trans: nothing
//
// [ cancellable ] trans: nothing
//
// [ result ] trans: everything
//
func NewDocumentFromStream(stream g.IInputStream, length int64, password string, cancellable g.ICancellable) (result Document, err error) {
	iv, err := _I.Get(71, "Document", "new_from_stream")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if stream != nil {
		tmp = stream.P_InputStream()
	}
	c_password := gi.CString(password)
	var tmp1 unsafe.Pointer
	if cancellable != nil {
		tmp1 = cancellable.P_Cancellable()
	}
	arg_stream := gi.NewPointerArgument(tmp)
	arg_length := gi.NewInt64Argument(length)
	arg_password := gi.NewStringArgument(c_password)
	arg_cancellable := gi.NewPointerArgument(tmp1)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_stream, arg_length, arg_password, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_password)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// poppler_document_find_dest
//
// [ link_name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Document) FindDest(link_name string) (result Dest) {
	iv, err := _I.Get(72, "Document", "find_dest")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_link_name := gi.CString(link_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_link_name := gi.NewStringArgument(c_link_name)
	args := []gi.Argument{arg_v, arg_link_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_link_name)
	result.P = ret.Pointer()
	return
}

// poppler_document_get_attachments
//
// [ result ] trans: everything
//
func (v Document) GetAttachments() (result g.List) {
	iv, err := _I.Get(73, "Document", "get_attachments")
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

// poppler_document_get_author
//
// [ result ] trans: everything
//
func (v Document) GetAuthor() (result string) {
	iv, err := _I.Get(74, "Document", "get_author")
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

// poppler_document_get_creation_date
//
// [ result ] trans: nothing
//
func (v Document) GetCreationDate() (result int64) {
	iv, err := _I.Get(75, "Document", "get_creation_date")
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

// poppler_document_get_creator
//
// [ result ] trans: everything
//
func (v Document) GetCreator() (result string) {
	iv, err := _I.Get(76, "Document", "get_creator")
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

// poppler_document_get_form_field
//
// [ id ] trans: nothing
//
// [ result ] trans: everything
//
func (v Document) GetFormField(id int32) (result FormField) {
	iv, err := _I.Get(77, "Document", "get_form_field")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewInt32Argument(id)
	args := []gi.Argument{arg_v, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_document_get_id
//
// [ permanent_id ] trans: everything, dir: out
//
// [ update_id ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Document) GetId() (result bool, permanent_id string, update_id string) {
	iv, err := _I.Get(78, "Document", "get_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_permanent_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_update_id := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_permanent_id, arg_update_id}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	permanent_id = outArgs[0].String().Take()
	update_id = outArgs[1].String().Take()
	result = ret.Bool()
	return
}

// poppler_document_get_keywords
//
// [ result ] trans: everything
//
func (v Document) GetKeywords() (result string) {
	iv, err := _I.Get(79, "Document", "get_keywords")
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

// poppler_document_get_metadata
//
// [ result ] trans: everything
//
func (v Document) GetMetadata() (result string) {
	iv, err := _I.Get(80, "Document", "get_metadata")
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

// poppler_document_get_modification_date
//
// [ result ] trans: nothing
//
func (v Document) GetModificationDate() (result int64) {
	iv, err := _I.Get(81, "Document", "get_modification_date")
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

// poppler_document_get_n_attachments
//
// [ result ] trans: nothing
//
func (v Document) GetNAttachments() (result uint32) {
	iv, err := _I.Get(82, "Document", "get_n_attachments")
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

// poppler_document_get_n_pages
//
// [ result ] trans: nothing
//
func (v Document) GetNPages() (result int32) {
	iv, err := _I.Get(83, "Document", "get_n_pages")
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

// poppler_document_get_page
//
// [ index ] trans: nothing
//
// [ result ] trans: everything
//
func (v Document) GetPage(index int32) (result Page) {
	iv, err := _I.Get(84, "Document", "get_page")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_document_get_page_by_label
//
// [ label ] trans: nothing
//
// [ result ] trans: everything
//
func (v Document) GetPageByLabel(label string) (result Page) {
	iv, err := _I.Get(85, "Document", "get_page_by_label")
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

// poppler_document_get_page_layout
//
// [ result ] trans: nothing
//
func (v Document) GetPageLayout() (result PageLayoutEnum) {
	iv, err := _I.Get(86, "Document", "get_page_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PageLayoutEnum(ret.Int())
	return
}

// poppler_document_get_page_mode
//
// [ result ] trans: nothing
//
func (v Document) GetPageMode() (result PageModeEnum) {
	iv, err := _I.Get(87, "Document", "get_page_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PageModeEnum(ret.Int())
	return
}

// poppler_document_get_pdf_conformance
//
// [ result ] trans: nothing
//
func (v Document) GetPdfConformance() (result PDFConformanceEnum) {
	iv, err := _I.Get(88, "Document", "get_pdf_conformance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PDFConformanceEnum(ret.Int())
	return
}

// poppler_document_get_pdf_part
//
// [ result ] trans: nothing
//
func (v Document) GetPdfPart() (result PDFPartEnum) {
	iv, err := _I.Get(89, "Document", "get_pdf_part")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PDFPartEnum(ret.Int())
	return
}

// poppler_document_get_pdf_subtype
//
// [ result ] trans: nothing
//
func (v Document) GetPdfSubtype() (result PDFSubtypeEnum) {
	iv, err := _I.Get(90, "Document", "get_pdf_subtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PDFSubtypeEnum(ret.Int())
	return
}

// poppler_document_get_pdf_subtype_string
//
// [ result ] trans: everything
//
func (v Document) GetPdfSubtypeString() (result string) {
	iv, err := _I.Get(91, "Document", "get_pdf_subtype_string")
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

// poppler_document_get_pdf_version
//
// [ major_version ] trans: everything, dir: out
//
// [ minor_version ] trans: everything, dir: out
//
func (v Document) GetPdfVersion() (major_version uint32, minor_version uint32) {
	iv, err := _I.Get(92, "Document", "get_pdf_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_major_version := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_minor_version := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_major_version, arg_minor_version}
	iv.Call(args, nil, &outArgs[0])
	major_version = outArgs[0].Uint32()
	minor_version = outArgs[1].Uint32()
	return
}

// poppler_document_get_pdf_version_string
//
// [ result ] trans: everything
//
func (v Document) GetPdfVersionString() (result string) {
	iv, err := _I.Get(93, "Document", "get_pdf_version_string")
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

// poppler_document_get_permissions
//
// [ result ] trans: nothing
//
func (v Document) GetPermissions() (result PermissionsFlags) {
	iv, err := _I.Get(94, "Document", "get_permissions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = PermissionsFlags(ret.Int())
	return
}

// poppler_document_get_producer
//
// [ result ] trans: everything
//
func (v Document) GetProducer() (result string) {
	iv, err := _I.Get(95, "Document", "get_producer")
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

// poppler_document_get_subject
//
// [ result ] trans: everything
//
func (v Document) GetSubject() (result string) {
	iv, err := _I.Get(96, "Document", "get_subject")
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

// poppler_document_get_title
//
// [ result ] trans: everything
//
func (v Document) GetTitle() (result string) {
	iv, err := _I.Get(97, "Document", "get_title")
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

// poppler_document_has_attachments
//
// [ result ] trans: nothing
//
func (v Document) HasAttachments() (result bool) {
	iv, err := _I.Get(98, "Document", "has_attachments")
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

// poppler_document_is_linearized
//
// [ result ] trans: nothing
//
func (v Document) IsLinearized() (result bool) {
	iv, err := _I.Get(99, "Document", "is_linearized")
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

// poppler_document_save
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Document) Save(uri string) (result bool, err error) {
	iv, err := _I.Get(100, "Document", "save")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
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

// poppler_document_save_a_copy
//
// [ uri ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Document) SaveACopy(uri string) (result bool, err error) {
	iv, err := _I.Get(101, "Document", "save_a_copy")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
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

// poppler_document_set_author
//
// [ author ] trans: nothing
//
func (v Document) SetAuthor(author string) {
	iv, err := _I.Get(102, "Document", "set_author")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_author := gi.CString(author)
	arg_v := gi.NewPointerArgument(v.P)
	arg_author := gi.NewStringArgument(c_author)
	args := []gi.Argument{arg_v, arg_author}
	iv.Call(args, nil, nil)
	gi.Free(c_author)
}

// poppler_document_set_creation_date
//
// [ creation_date ] trans: nothing
//
func (v Document) SetCreationDate(creation_date int64) {
	iv, err := _I.Get(103, "Document", "set_creation_date")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_creation_date := gi.NewInt64Argument(creation_date)
	args := []gi.Argument{arg_v, arg_creation_date}
	iv.Call(args, nil, nil)
}

// poppler_document_set_creator
//
// [ creator ] trans: nothing
//
func (v Document) SetCreator(creator string) {
	iv, err := _I.Get(104, "Document", "set_creator")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_creator := gi.CString(creator)
	arg_v := gi.NewPointerArgument(v.P)
	arg_creator := gi.NewStringArgument(c_creator)
	args := []gi.Argument{arg_v, arg_creator}
	iv.Call(args, nil, nil)
	gi.Free(c_creator)
}

// poppler_document_set_keywords
//
// [ keywords ] trans: nothing
//
func (v Document) SetKeywords(keywords string) {
	iv, err := _I.Get(105, "Document", "set_keywords")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_keywords := gi.CString(keywords)
	arg_v := gi.NewPointerArgument(v.P)
	arg_keywords := gi.NewStringArgument(c_keywords)
	args := []gi.Argument{arg_v, arg_keywords}
	iv.Call(args, nil, nil)
	gi.Free(c_keywords)
}

// poppler_document_set_modification_date
//
// [ modification_date ] trans: nothing
//
func (v Document) SetModificationDate(modification_date int64) {
	iv, err := _I.Get(106, "Document", "set_modification_date")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_modification_date := gi.NewInt64Argument(modification_date)
	args := []gi.Argument{arg_v, arg_modification_date}
	iv.Call(args, nil, nil)
}

// poppler_document_set_producer
//
// [ producer ] trans: nothing
//
func (v Document) SetProducer(producer string) {
	iv, err := _I.Get(107, "Document", "set_producer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_producer := gi.CString(producer)
	arg_v := gi.NewPointerArgument(v.P)
	arg_producer := gi.NewStringArgument(c_producer)
	args := []gi.Argument{arg_v, arg_producer}
	iv.Call(args, nil, nil)
	gi.Free(c_producer)
}

// poppler_document_set_subject
//
// [ subject ] trans: nothing
//
func (v Document) SetSubject(subject string) {
	iv, err := _I.Get(108, "Document", "set_subject")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_subject := gi.CString(subject)
	arg_v := gi.NewPointerArgument(v.P)
	arg_subject := gi.NewStringArgument(c_subject)
	args := []gi.Argument{arg_v, arg_subject}
	iv.Call(args, nil, nil)
	gi.Free(c_subject)
}

// poppler_document_set_title
//
// [ title ] trans: nothing
//
func (v Document) SetTitle(title string) {
	iv, err := _I.Get(109, "Document", "set_title")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_title := gi.CString(title)
	arg_v := gi.NewPointerArgument(v.P)
	arg_title := gi.NewStringArgument(c_title)
	args := []gi.Argument{arg_v, arg_title}
	iv.Call(args, nil, nil)
	gi.Free(c_title)
}

// Enum Error
type ErrorEnum int

const (
	ErrorInvalid    ErrorEnum = 0
	ErrorEncrypted  ErrorEnum = 1
	ErrorOpenFile   ErrorEnum = 2
	ErrorBadCatalog ErrorEnum = 3
	ErrorDamaged    ErrorEnum = 4
)

func ErrorGetType() gi.GType {
	ret := _I.GetGType(40, "Error")
	return ret
}

// Flags FindFlags
type FindFlags int

const (
	FindFlagsDefault        FindFlags = 0
	FindFlagsCaseSensitive  FindFlags = 1
	FindFlagsBackwards      FindFlags = 2
	FindFlagsWholeWordsOnly FindFlags = 4
)

func FindFlagsGetType() gi.GType {
	ret := _I.GetGType(41, "FindFlags")
	return ret
}

// Object FontInfo
type FontInfo struct {
	g.Object
}

func WrapFontInfo(p unsafe.Pointer) (r FontInfo) { r.P = p; return }

type IFontInfo interface{ P_FontInfo() unsafe.Pointer }

func (v FontInfo) P_FontInfo() unsafe.Pointer { return v.P }
func FontInfoGetType() gi.GType {
	ret := _I.GetGType(42, "FontInfo")
	return ret
}

// poppler_font_info_new
//
// [ document ] trans: nothing
//
// [ result ] trans: everything
//
func NewFontInfo(document IDocument) (result FontInfo) {
	iv, err := _I.Get(110, "FontInfo", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if document != nil {
		tmp = document.P_Document()
	}
	arg_document := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_document}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_font_info_free
//
func (v FontInfo) Free() {
	iv, err := _I.Get(111, "FontInfo", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_font_info_scan
//
// [ n_pages ] trans: nothing
//
// [ iter ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v FontInfo) Scan(n_pages int32) (result bool, iter FontsIter) {
	iv, err := _I.Get(112, "FontInfo", "scan")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_pages := gi.NewInt32Argument(n_pages)
	arg_iter := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_n_pages, arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	iter.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// Enum FontType
type FontTypeEnum int

const (
	FontTypeUnknown     FontTypeEnum = 0
	FontTypeType1       FontTypeEnum = 1
	FontTypeType1c      FontTypeEnum = 2
	FontTypeType1cot    FontTypeEnum = 3
	FontTypeType3       FontTypeEnum = 4
	FontTypeTruetype    FontTypeEnum = 5
	FontTypeTruetypeot  FontTypeEnum = 6
	FontTypeCidType0    FontTypeEnum = 7
	FontTypeCidType0c   FontTypeEnum = 8
	FontTypeCidType0cot FontTypeEnum = 9
	FontTypeCidType2    FontTypeEnum = 10
	FontTypeCidType2ot  FontTypeEnum = 11
)

func FontTypeGetType() gi.GType {
	ret := _I.GetGType(43, "FontType")
	return ret
}

// Struct FontsIter
type FontsIter struct {
	P unsafe.Pointer
}

func FontsIterGetType() gi.GType {
	ret := _I.GetGType(44, "FontsIter")
	return ret
}

// poppler_fonts_iter_copy
//
// [ result ] trans: everything
//
func (v FontsIter) Copy() (result FontsIter) {
	iv, err := _I.Get(113, "FontsIter", "copy")
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

// poppler_fonts_iter_free
//
func (v FontsIter) Free() {
	iv, err := _I.Get(114, "FontsIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_fonts_iter_get_encoding
//
// [ result ] trans: nothing
//
func (v FontsIter) GetEncoding() (result string) {
	iv, err := _I.Get(115, "FontsIter", "get_encoding")
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

// poppler_fonts_iter_get_file_name
//
// [ result ] trans: nothing
//
func (v FontsIter) GetFileName() (result string) {
	iv, err := _I.Get(116, "FontsIter", "get_file_name")
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

// poppler_fonts_iter_get_font_type
//
// [ result ] trans: nothing
//
func (v FontsIter) GetFontType() (result FontTypeEnum) {
	iv, err := _I.Get(117, "FontsIter", "get_font_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FontTypeEnum(ret.Int())
	return
}

// poppler_fonts_iter_get_full_name
//
// [ result ] trans: nothing
//
func (v FontsIter) GetFullName() (result string) {
	iv, err := _I.Get(118, "FontsIter", "get_full_name")
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

// poppler_fonts_iter_get_name
//
// [ result ] trans: nothing
//
func (v FontsIter) GetName() (result string) {
	iv, err := _I.Get(119, "FontsIter", "get_name")
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

// poppler_fonts_iter_get_substitute_name
//
// [ result ] trans: nothing
//
func (v FontsIter) GetSubstituteName() (result string) {
	iv, err := _I.Get(120, "FontsIter", "get_substitute_name")
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

// poppler_fonts_iter_is_embedded
//
// [ result ] trans: nothing
//
func (v FontsIter) IsEmbedded() (result bool) {
	iv, err := _I.Get(121, "FontsIter", "is_embedded")
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

// poppler_fonts_iter_is_subset
//
// [ result ] trans: nothing
//
func (v FontsIter) IsSubset() (result bool) {
	iv, err := _I.Get(122, "FontsIter", "is_subset")
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

// poppler_fonts_iter_next
//
// [ result ] trans: nothing
//
func (v FontsIter) Next() (result bool) {
	iv, err := _I.Get(123, "FontsIter", "next")
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

// Enum FormButtonType
type FormButtonTypeEnum int

const (
	FormButtonTypePush  FormButtonTypeEnum = 0
	FormButtonTypeCheck FormButtonTypeEnum = 1
	FormButtonTypeRadio FormButtonTypeEnum = 2
)

func FormButtonTypeGetType() gi.GType {
	ret := _I.GetGType(45, "FormButtonType")
	return ret
}

// Enum FormChoiceType
type FormChoiceTypeEnum int

const (
	FormChoiceTypeCombo FormChoiceTypeEnum = 0
	FormChoiceTypeList  FormChoiceTypeEnum = 1
)

func FormChoiceTypeGetType() gi.GType {
	ret := _I.GetGType(46, "FormChoiceType")
	return ret
}

// Object FormField
type FormField struct {
	g.Object
}

func WrapFormField(p unsafe.Pointer) (r FormField) { r.P = p; return }

type IFormField interface{ P_FormField() unsafe.Pointer }

func (v FormField) P_FormField() unsafe.Pointer { return v.P }
func FormFieldGetType() gi.GType {
	ret := _I.GetGType(47, "FormField")
	return ret
}

// poppler_form_field_button_get_button_type
//
// [ result ] trans: nothing
//
func (v FormField) ButtonGetButtonType() (result FormButtonTypeEnum) {
	iv, err := _I.Get(124, "FormField", "button_get_button_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FormButtonTypeEnum(ret.Int())
	return
}

// poppler_form_field_button_get_state
//
// [ result ] trans: nothing
//
func (v FormField) ButtonGetState() (result bool) {
	iv, err := _I.Get(125, "FormField", "button_get_state")
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

// poppler_form_field_button_set_state
//
// [ state ] trans: nothing
//
func (v FormField) ButtonSetState(state bool) {
	iv, err := _I.Get(126, "FormField", "button_set_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewBoolArgument(state)
	args := []gi.Argument{arg_v, arg_state}
	iv.Call(args, nil, nil)
}

// poppler_form_field_choice_can_select_multiple
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceCanSelectMultiple() (result bool) {
	iv, err := _I.Get(127, "FormField", "choice_can_select_multiple")
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

// poppler_form_field_choice_commit_on_change
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceCommitOnChange() (result bool) {
	iv, err := _I.Get(128, "FormField", "choice_commit_on_change")
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

// poppler_form_field_choice_do_spell_check
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceDoSpellCheck() (result bool) {
	iv, err := _I.Get(129, "FormField", "choice_do_spell_check")
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

// poppler_form_field_choice_get_choice_type
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceGetChoiceType() (result FormChoiceTypeEnum) {
	iv, err := _I.Get(130, "FormField", "choice_get_choice_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FormChoiceTypeEnum(ret.Int())
	return
}

// poppler_form_field_choice_get_item
//
// [ index ] trans: nothing
//
// [ result ] trans: everything
//
func (v FormField) ChoiceGetItem(index int32) (result string) {
	iv, err := _I.Get(131, "FormField", "choice_get_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// poppler_form_field_choice_get_n_items
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceGetNItems() (result int32) {
	iv, err := _I.Get(132, "FormField", "choice_get_n_items")
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

// poppler_form_field_choice_get_text
//
// [ result ] trans: everything
//
func (v FormField) ChoiceGetText() (result string) {
	iv, err := _I.Get(133, "FormField", "choice_get_text")
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

// poppler_form_field_choice_is_editable
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceIsEditable() (result bool) {
	iv, err := _I.Get(134, "FormField", "choice_is_editable")
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

// poppler_form_field_choice_is_item_selected
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func (v FormField) ChoiceIsItemSelected(index int32) (result bool) {
	iv, err := _I.Get(135, "FormField", "choice_is_item_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_form_field_choice_select_item
//
// [ index ] trans: nothing
//
func (v FormField) ChoiceSelectItem(index int32) {
	iv, err := _I.Get(136, "FormField", "choice_select_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	iv.Call(args, nil, nil)
}

// poppler_form_field_choice_set_text
//
// [ text ] trans: nothing
//
func (v FormField) ChoiceSetText(text string) {
	iv, err := _I.Get(137, "FormField", "choice_set_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_v, arg_text}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// poppler_form_field_choice_toggle_item
//
// [ index ] trans: nothing
//
func (v FormField) ChoiceToggleItem(index int32) {
	iv, err := _I.Get(138, "FormField", "choice_toggle_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_v, arg_index}
	iv.Call(args, nil, nil)
}

// poppler_form_field_choice_unselect_all
//
func (v FormField) ChoiceUnselectAll() {
	iv, err := _I.Get(139, "FormField", "choice_unselect_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_form_field_get_action
//
// [ result ] trans: nothing
//
func (v FormField) GetAction() (result Action) {
	iv, err := _I.Get(140, "FormField", "get_action")
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

// poppler_form_field_get_field_type
//
// [ result ] trans: nothing
//
func (v FormField) GetFieldType() (result FormFieldTypeEnum) {
	iv, err := _I.Get(141, "FormField", "get_field_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FormFieldTypeEnum(ret.Int())
	return
}

// poppler_form_field_get_font_size
//
// [ result ] trans: nothing
//
func (v FormField) GetFontSize() (result float64) {
	iv, err := _I.Get(142, "FormField", "get_font_size")
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

// poppler_form_field_get_id
//
// [ result ] trans: nothing
//
func (v FormField) GetId() (result int32) {
	iv, err := _I.Get(143, "FormField", "get_id")
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

// poppler_form_field_get_mapping_name
//
// [ result ] trans: everything
//
func (v FormField) GetMappingName() (result string) {
	iv, err := _I.Get(144, "FormField", "get_mapping_name")
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

// poppler_form_field_get_name
//
// [ result ] trans: everything
//
func (v FormField) GetName() (result string) {
	iv, err := _I.Get(145, "FormField", "get_name")
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

// poppler_form_field_get_partial_name
//
// [ result ] trans: everything
//
func (v FormField) GetPartialName() (result string) {
	iv, err := _I.Get(146, "FormField", "get_partial_name")
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

// poppler_form_field_is_read_only
//
// [ result ] trans: nothing
//
func (v FormField) IsReadOnly() (result bool) {
	iv, err := _I.Get(147, "FormField", "is_read_only")
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

// poppler_form_field_text_do_scroll
//
// [ result ] trans: nothing
//
func (v FormField) TextDoScroll() (result bool) {
	iv, err := _I.Get(148, "FormField", "text_do_scroll")
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

// poppler_form_field_text_do_spell_check
//
// [ result ] trans: nothing
//
func (v FormField) TextDoSpellCheck() (result bool) {
	iv, err := _I.Get(149, "FormField", "text_do_spell_check")
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

// poppler_form_field_text_get_max_len
//
// [ result ] trans: nothing
//
func (v FormField) TextGetMaxLen() (result int32) {
	iv, err := _I.Get(150, "FormField", "text_get_max_len")
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

// poppler_form_field_text_get_text
//
// [ result ] trans: everything
//
func (v FormField) TextGetText() (result string) {
	iv, err := _I.Get(151, "FormField", "text_get_text")
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

// poppler_form_field_text_get_text_type
//
// [ result ] trans: nothing
//
func (v FormField) TextGetTextType() (result FormTextTypeEnum) {
	iv, err := _I.Get(152, "FormField", "text_get_text_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FormTextTypeEnum(ret.Int())
	return
}

// poppler_form_field_text_is_password
//
// [ result ] trans: nothing
//
func (v FormField) TextIsPassword() (result bool) {
	iv, err := _I.Get(153, "FormField", "text_is_password")
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

// poppler_form_field_text_is_rich_text
//
// [ result ] trans: nothing
//
func (v FormField) TextIsRichText() (result bool) {
	iv, err := _I.Get(154, "FormField", "text_is_rich_text")
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

// poppler_form_field_text_set_text
//
// [ text ] trans: nothing
//
func (v FormField) TextSetText(text string) {
	iv, err := _I.Get(155, "FormField", "text_set_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_v, arg_text}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// Struct FormFieldMapping
type FormFieldMapping struct {
	P unsafe.Pointer
}

const SizeOfStructFormFieldMapping = 40

func FormFieldMappingGetType() gi.GType {
	ret := _I.GetGType(48, "FormFieldMapping")
	return ret
}

// poppler_form_field_mapping_new
//
// [ result ] trans: everything
//
func NewFormFieldMapping() (result FormFieldMapping) {
	iv, err := _I.Get(156, "FormFieldMapping", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_form_field_mapping_copy
//
// [ result ] trans: everything
//
func (v FormFieldMapping) Copy() (result FormFieldMapping) {
	iv, err := _I.Get(157, "FormFieldMapping", "copy")
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

// poppler_form_field_mapping_free
//
func (v FormFieldMapping) Free() {
	iv, err := _I.Get(158, "FormFieldMapping", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum FormFieldType
type FormFieldTypeEnum int

const (
	FormFieldTypeUnknown   FormFieldTypeEnum = 0
	FormFieldTypeButton    FormFieldTypeEnum = 1
	FormFieldTypeText      FormFieldTypeEnum = 2
	FormFieldTypeChoice    FormFieldTypeEnum = 3
	FormFieldTypeSignature FormFieldTypeEnum = 4
)

func FormFieldTypeGetType() gi.GType {
	ret := _I.GetGType(49, "FormFieldType")
	return ret
}

// Enum FormTextType
type FormTextTypeEnum int

const (
	FormTextTypeNormal     FormTextTypeEnum = 0
	FormTextTypeMultiline  FormTextTypeEnum = 1
	FormTextTypeFileSelect FormTextTypeEnum = 2
)

func FormTextTypeGetType() gi.GType {
	ret := _I.GetGType(50, "FormTextType")
	return ret
}

// Struct ImageMapping
type ImageMapping struct {
	P unsafe.Pointer
}

const SizeOfStructImageMapping = 40

func ImageMappingGetType() gi.GType {
	ret := _I.GetGType(51, "ImageMapping")
	return ret
}

// poppler_image_mapping_new
//
// [ result ] trans: everything
//
func NewImageMapping() (result ImageMapping) {
	iv, err := _I.Get(159, "ImageMapping", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_image_mapping_copy
//
// [ result ] trans: everything
//
func (v ImageMapping) Copy() (result ImageMapping) {
	iv, err := _I.Get(160, "ImageMapping", "copy")
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

// poppler_image_mapping_free
//
func (v ImageMapping) Free() {
	iv, err := _I.Get(161, "ImageMapping", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct IndexIter
type IndexIter struct {
	P unsafe.Pointer
}

func IndexIterGetType() gi.GType {
	ret := _I.GetGType(52, "IndexIter")
	return ret
}

// poppler_index_iter_new
//
// [ document ] trans: nothing
//
// [ result ] trans: everything
//
func NewIndexIter(document IDocument) (result IndexIter) {
	iv, err := _I.Get(162, "IndexIter", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if document != nil {
		tmp = document.P_Document()
	}
	arg_document := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_document}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_index_iter_copy
//
// [ result ] trans: everything
//
func (v IndexIter) Copy() (result IndexIter) {
	iv, err := _I.Get(163, "IndexIter", "copy")
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

// poppler_index_iter_free
//
func (v IndexIter) Free() {
	iv, err := _I.Get(164, "IndexIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_index_iter_get_action
//
// [ result ] trans: everything
//
func (v IndexIter) GetAction() (result Action) {
	iv, err := _I.Get(165, "IndexIter", "get_action")
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

// poppler_index_iter_get_child
//
// [ result ] trans: everything
//
func (v IndexIter) GetChild() (result IndexIter) {
	iv, err := _I.Get(166, "IndexIter", "get_child")
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

// poppler_index_iter_is_open
//
// [ result ] trans: nothing
//
func (v IndexIter) IsOpen() (result bool) {
	iv, err := _I.Get(167, "IndexIter", "is_open")
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

// poppler_index_iter_next
//
// [ result ] trans: nothing
//
func (v IndexIter) Next() (result bool) {
	iv, err := _I.Get(168, "IndexIter", "next")
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

// Object Layer
type Layer struct {
	g.Object
}

func WrapLayer(p unsafe.Pointer) (r Layer) { r.P = p; return }

type ILayer interface{ P_Layer() unsafe.Pointer }

func (v Layer) P_Layer() unsafe.Pointer { return v.P }
func LayerGetType() gi.GType {
	ret := _I.GetGType(53, "Layer")
	return ret
}

// poppler_layer_get_radio_button_group_id
//
// [ result ] trans: nothing
//
func (v Layer) GetRadioButtonGroupId() (result int32) {
	iv, err := _I.Get(169, "Layer", "get_radio_button_group_id")
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

// poppler_layer_get_title
//
// [ result ] trans: nothing
//
func (v Layer) GetTitle() (result string) {
	iv, err := _I.Get(170, "Layer", "get_title")
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

// poppler_layer_hide
//
func (v Layer) Hide() {
	iv, err := _I.Get(171, "Layer", "hide")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_layer_is_parent
//
// [ result ] trans: nothing
//
func (v Layer) IsParent() (result bool) {
	iv, err := _I.Get(172, "Layer", "is_parent")
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

// poppler_layer_is_visible
//
// [ result ] trans: nothing
//
func (v Layer) IsVisible() (result bool) {
	iv, err := _I.Get(173, "Layer", "is_visible")
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

// poppler_layer_show
//
func (v Layer) Show() {
	iv, err := _I.Get(174, "Layer", "show")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct LayersIter
type LayersIter struct {
	P unsafe.Pointer
}

func LayersIterGetType() gi.GType {
	ret := _I.GetGType(54, "LayersIter")
	return ret
}

// poppler_layers_iter_new
//
// [ document ] trans: nothing
//
// [ result ] trans: everything
//
func NewLayersIter(document IDocument) (result LayersIter) {
	iv, err := _I.Get(175, "LayersIter", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if document != nil {
		tmp = document.P_Document()
	}
	arg_document := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_document}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_layers_iter_copy
//
// [ result ] trans: everything
//
func (v LayersIter) Copy() (result LayersIter) {
	iv, err := _I.Get(176, "LayersIter", "copy")
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

// poppler_layers_iter_free
//
func (v LayersIter) Free() {
	iv, err := _I.Get(177, "LayersIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_layers_iter_get_child
//
// [ result ] trans: everything
//
func (v LayersIter) GetChild() (result LayersIter) {
	iv, err := _I.Get(178, "LayersIter", "get_child")
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

// poppler_layers_iter_get_layer
//
// [ result ] trans: everything
//
func (v LayersIter) GetLayer() (result Layer) {
	iv, err := _I.Get(179, "LayersIter", "get_layer")
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

// poppler_layers_iter_get_title
//
// [ result ] trans: everything
//
func (v LayersIter) GetTitle() (result string) {
	iv, err := _I.Get(180, "LayersIter", "get_title")
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

// poppler_layers_iter_next
//
// [ result ] trans: nothing
//
func (v LayersIter) Next() (result bool) {
	iv, err := _I.Get(181, "LayersIter", "next")
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

// Struct LinkMapping
type LinkMapping struct {
	P unsafe.Pointer
}

const SizeOfStructLinkMapping = 40

func LinkMappingGetType() gi.GType {
	ret := _I.GetGType(55, "LinkMapping")
	return ret
}

// poppler_link_mapping_new
//
// [ result ] trans: everything
//
func NewLinkMapping() (result LinkMapping) {
	iv, err := _I.Get(182, "LinkMapping", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_link_mapping_copy
//
// [ result ] trans: everything
//
func (v LinkMapping) Copy() (result LinkMapping) {
	iv, err := _I.Get(183, "LinkMapping", "copy")
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

// poppler_link_mapping_free
//
func (v LinkMapping) Free() {
	iv, err := _I.Get(184, "LinkMapping", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object Media
type Media struct {
	g.Object
}

func WrapMedia(p unsafe.Pointer) (r Media) { r.P = p; return }

type IMedia interface{ P_Media() unsafe.Pointer }

func (v Media) P_Media() unsafe.Pointer { return v.P }
func MediaGetType() gi.GType {
	ret := _I.GetGType(56, "Media")
	return ret
}

// poppler_media_get_filename
//
// [ result ] trans: nothing
//
func (v Media) GetFilename() (result string) {
	iv, err := _I.Get(185, "Media", "get_filename")
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

// poppler_media_get_mime_type
//
// [ result ] trans: nothing
//
func (v Media) GetMimeType() (result string) {
	iv, err := _I.Get(186, "Media", "get_mime_type")
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

// poppler_media_is_embedded
//
// [ result ] trans: nothing
//
func (v Media) IsEmbedded() (result bool) {
	iv, err := _I.Get(187, "Media", "is_embedded")
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

// poppler_media_save
//
// [ filename ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Media) Save(filename string) (result bool, err error) {
	iv, err := _I.Get(188, "Media", "save")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// poppler_media_save_to_callback
//
// [ save_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Media) SaveToCallback(save_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result bool, err error) {
	iv, err := _I.Get(189, "Media", "save_to_callback")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_save_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myMediaSaveFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_save_func, arg_user_data, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

type MediaSaveFuncStruct struct {
	F_buf   gi.Uint8Array
	F_count uint64
	F_data  unsafe.Pointer
}

func GetPointer_myMediaSaveFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myPopplerMediaSaveFunc())
}

//export myPopplerMediaSaveFunc
func myPopplerMediaSaveFunc(buf C.gpointer, count C.guint64, data C.gpointer) {
	// TODO: not found user_data
}

// Object Movie
type Movie struct {
	g.Object
}

func WrapMovie(p unsafe.Pointer) (r Movie) { r.P = p; return }

type IMovie interface{ P_Movie() unsafe.Pointer }

func (v Movie) P_Movie() unsafe.Pointer { return v.P }
func MovieGetType() gi.GType {
	ret := _I.GetGType(57, "Movie")
	return ret
}

// poppler_movie_get_filename
//
// [ result ] trans: nothing
//
func (v Movie) GetFilename() (result string) {
	iv, err := _I.Get(190, "Movie", "get_filename")
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

// poppler_movie_get_play_mode
//
// [ result ] trans: nothing
//
func (v Movie) GetPlayMode() (result MoviePlayModeEnum) {
	iv, err := _I.Get(191, "Movie", "get_play_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = MoviePlayModeEnum(ret.Int())
	return
}

// poppler_movie_need_poster
//
// [ result ] trans: nothing
//
func (v Movie) NeedPoster() (result bool) {
	iv, err := _I.Get(192, "Movie", "need_poster")
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

// poppler_movie_show_controls
//
// [ result ] trans: nothing
//
func (v Movie) ShowControls() (result bool) {
	iv, err := _I.Get(193, "Movie", "show_controls")
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

// Enum MoviePlayMode
type MoviePlayModeEnum int

const (
	MoviePlayModeOnce       MoviePlayModeEnum = 0
	MoviePlayModeOpen       MoviePlayModeEnum = 1
	MoviePlayModeRepeat     MoviePlayModeEnum = 2
	MoviePlayModePalindrome MoviePlayModeEnum = 3
)

func MoviePlayModeGetType() gi.GType {
	ret := _I.GetGType(58, "MoviePlayMode")
	return ret
}

// Enum PDFConformance
type PDFConformanceEnum int

const (
	PDFConformanceUnset PDFConformanceEnum = 0
	PDFConformanceA     PDFConformanceEnum = 1
	PDFConformanceB     PDFConformanceEnum = 2
	PDFConformanceG     PDFConformanceEnum = 3
	PDFConformanceN     PDFConformanceEnum = 4
	PDFConformanceP     PDFConformanceEnum = 5
	PDFConformancePg    PDFConformanceEnum = 6
	PDFConformanceU     PDFConformanceEnum = 7
	PDFConformanceNone  PDFConformanceEnum = 8
)

func PDFConformanceGetType() gi.GType {
	ret := _I.GetGType(59, "PDFConformance")
	return ret
}

// Enum PDFPart
type PDFPartEnum int

const (
	PDFPartUnset PDFPartEnum = 0
	PDFPart1     PDFPartEnum = 1
	PDFPart2     PDFPartEnum = 2
	PDFPart3     PDFPartEnum = 3
	PDFPart4     PDFPartEnum = 4
	PDFPart5     PDFPartEnum = 5
	PDFPart6     PDFPartEnum = 6
	PDFPart7     PDFPartEnum = 7
	PDFPart8     PDFPartEnum = 8
	PDFPartNone  PDFPartEnum = 9
)

func PDFPartGetType() gi.GType {
	ret := _I.GetGType(60, "PDFPart")
	return ret
}

// Enum PDFSubtype
type PDFSubtypeEnum int

const (
	PDFSubtypeUnset PDFSubtypeEnum = 0
	PDFSubtypePdfA  PDFSubtypeEnum = 1
	PDFSubtypePdfE  PDFSubtypeEnum = 2
	PDFSubtypePdfUa PDFSubtypeEnum = 3
	PDFSubtypePdfVt PDFSubtypeEnum = 4
	PDFSubtypePdfX  PDFSubtypeEnum = 5
	PDFSubtypeNone  PDFSubtypeEnum = 6
)

func PDFSubtypeGetType() gi.GType {
	ret := _I.GetGType(61, "PDFSubtype")
	return ret
}

// Object PSFile
type PSFile struct {
	g.Object
}

func WrapPSFile(p unsafe.Pointer) (r PSFile) { r.P = p; return }

type IPSFile interface{ P_PSFile() unsafe.Pointer }

func (v PSFile) P_PSFile() unsafe.Pointer { return v.P }
func PSFileGetType() gi.GType {
	ret := _I.GetGType(62, "PSFile")
	return ret
}

// poppler_ps_file_new
//
// [ document ] trans: nothing
//
// [ filename ] trans: nothing
//
// [ first_page ] trans: nothing
//
// [ n_pages ] trans: nothing
//
// [ result ] trans: everything
//
func NewPSFile(document IDocument, filename string, first_page int32, n_pages int32) (result PSFile) {
	iv, err := _I.Get(194, "PSFile", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if document != nil {
		tmp = document.P_Document()
	}
	c_filename := gi.CString(filename)
	arg_document := gi.NewPointerArgument(tmp)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_first_page := gi.NewInt32Argument(first_page)
	arg_n_pages := gi.NewInt32Argument(n_pages)
	args := []gi.Argument{arg_document, arg_filename, arg_first_page, arg_n_pages}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result.P = ret.Pointer()
	return
}

// poppler_ps_file_free
//
func (v PSFile) Free() {
	iv, err := _I.Get(195, "PSFile", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_ps_file_set_duplex
//
// [ duplex ] trans: nothing
//
func (v PSFile) SetDuplex(duplex bool) {
	iv, err := _I.Get(196, "PSFile", "set_duplex")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_duplex := gi.NewBoolArgument(duplex)
	args := []gi.Argument{arg_v, arg_duplex}
	iv.Call(args, nil, nil)
}

// poppler_ps_file_set_paper_size
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func (v PSFile) SetPaperSize(width float64, height float64) {
	iv, err := _I.Get(197, "PSFile", "set_paper_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewDoubleArgument(width)
	arg_height := gi.NewDoubleArgument(height)
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// Object Page
type Page struct {
	g.Object
}

func WrapPage(p unsafe.Pointer) (r Page) { r.P = p; return }

type IPage interface{ P_Page() unsafe.Pointer }

func (v Page) P_Page() unsafe.Pointer { return v.P }
func PageGetType() gi.GType {
	ret := _I.GetGType(63, "Page")
	return ret
}

// poppler_page_free_annot_mapping
//
// [ list ] trans: nothing
//
func PageFreeAnnotMapping1(list g.List) {
	iv, err := _I.Get(198, "Page", "free_annot_mapping")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// poppler_page_free_form_field_mapping
//
// [ list ] trans: nothing
//
func PageFreeFormFieldMapping1(list g.List) {
	iv, err := _I.Get(199, "Page", "free_form_field_mapping")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// poppler_page_free_image_mapping
//
// [ list ] trans: nothing
//
func PageFreeImageMapping1(list g.List) {
	iv, err := _I.Get(200, "Page", "free_image_mapping")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// poppler_page_free_link_mapping
//
// [ list ] trans: nothing
//
func PageFreeLinkMapping1(list g.List) {
	iv, err := _I.Get(201, "Page", "free_link_mapping")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// poppler_page_free_text_attributes
//
// [ list ] trans: nothing
//
func PageFreeTextAttributes1(list g.List) {
	iv, err := _I.Get(202, "Page", "free_text_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_list}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// poppler_page_selection_region_free
//
// [ region ] trans: nothing
//
func PageSelectionRegionFree1(region g.List) {
	iv, err := _I.Get(203, "Page", "selection_region_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_region := gi.NewPointerArgument(region.P)
	args := []gi.Argument{arg_region}
	iv.Call(args, nil, nil)
}

// poppler_page_add_annot
//
// [ annot ] trans: nothing
//
func (v Page) AddAnnot(annot IAnnot) {
	iv, err := _I.Get(204, "Page", "add_annot")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if annot != nil {
		tmp = annot.P_Annot()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_annot := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_annot}
	iv.Call(args, nil, nil)
}

// poppler_page_find_text
//
// [ text ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) FindText(text string) (result g.List) {
	iv, err := _I.Get(205, "Page", "find_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_v, arg_text}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// poppler_page_find_text_with_options
//
// [ text ] trans: nothing
//
// [ options ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) FindTextWithOptions(text string, options FindFlags) (result g.List) {
	iv, err := _I.Get(206, "Page", "find_text_with_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_options := gi.NewIntArgument(int(options))
	args := []gi.Argument{arg_v, arg_text, arg_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// poppler_page_get_annot_mapping
//
// [ result ] trans: everything
//
func (v Page) GetAnnotMapping() (result g.List) {
	iv, err := _I.Get(207, "Page", "get_annot_mapping")
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

// poppler_page_get_crop_box
//
// [ rect ] trans: nothing, dir: out
//
func (v Page) GetCropBox(rect Rectangle) {
	iv, err := _I.Get(208, "Page", "get_crop_box")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_v, arg_rect}
	iv.Call(args, nil, nil)
}

// poppler_page_get_duration
//
// [ result ] trans: nothing
//
func (v Page) GetDuration() (result float64) {
	iv, err := _I.Get(209, "Page", "get_duration")
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

// poppler_page_get_form_field_mapping
//
// [ result ] trans: everything
//
func (v Page) GetFormFieldMapping() (result g.List) {
	iv, err := _I.Get(210, "Page", "get_form_field_mapping")
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

// poppler_page_get_image
//
// [ image_id ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) GetImage(image_id int32) (result cairo.Surface) {
	iv, err := _I.Get(211, "Page", "get_image")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_image_id := gi.NewInt32Argument(image_id)
	args := []gi.Argument{arg_v, arg_image_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_page_get_image_mapping
//
// [ result ] trans: everything
//
func (v Page) GetImageMapping() (result g.List) {
	iv, err := _I.Get(212, "Page", "get_image_mapping")
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

// poppler_page_get_index
//
// [ result ] trans: nothing
//
func (v Page) GetIndex() (result int32) {
	iv, err := _I.Get(213, "Page", "get_index")
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

// poppler_page_get_label
//
// [ result ] trans: everything
//
func (v Page) GetLabel() (result string) {
	iv, err := _I.Get(214, "Page", "get_label")
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

// poppler_page_get_link_mapping
//
// [ result ] trans: everything
//
func (v Page) GetLinkMapping() (result g.List) {
	iv, err := _I.Get(215, "Page", "get_link_mapping")
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

// poppler_page_get_selected_region
//
// [ scale ] trans: nothing
//
// [ style ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) GetSelectedRegion(scale float64, style SelectionStyleEnum, selection Rectangle) (result cairo.Region) {
	iv, err := _I.Get(216, "Page", "get_selected_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scale := gi.NewDoubleArgument(scale)
	arg_style := gi.NewIntArgument(int(style))
	arg_selection := gi.NewPointerArgument(selection.P)
	args := []gi.Argument{arg_v, arg_scale, arg_style, arg_selection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_page_get_selected_text
//
// [ style ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) GetSelectedText(style SelectionStyleEnum, selection Rectangle) (result string) {
	iv, err := _I.Get(217, "Page", "get_selected_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_style := gi.NewIntArgument(int(style))
	arg_selection := gi.NewPointerArgument(selection.P)
	args := []gi.Argument{arg_v, arg_style, arg_selection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// Deprecated
//
// poppler_page_get_selection_region
//
// [ scale ] trans: nothing
//
// [ style ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) GetSelectionRegion(scale float64, style SelectionStyleEnum, selection Rectangle) (result g.List) {
	iv, err := _I.Get(218, "Page", "get_selection_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scale := gi.NewDoubleArgument(scale)
	arg_style := gi.NewIntArgument(int(style))
	arg_selection := gi.NewPointerArgument(selection.P)
	args := []gi.Argument{arg_v, arg_scale, arg_style, arg_selection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_page_get_size
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
func (v Page) GetSize() (width float64, height float64) {
	iv, err := _I.Get(219, "Page", "get_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	width = outArgs[0].Double()
	height = outArgs[1].Double()
	return
}

// poppler_page_get_text
//
// [ result ] trans: everything
//
func (v Page) GetText() (result string) {
	iv, err := _I.Get(220, "Page", "get_text")
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

// poppler_page_get_text_attributes
//
// [ result ] trans: everything
//
func (v Page) GetTextAttributes() (result g.List) {
	iv, err := _I.Get(221, "Page", "get_text_attributes")
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

// poppler_page_get_text_attributes_for_area
//
// [ area ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) GetTextAttributesForArea(area Rectangle) (result g.List) {
	iv, err := _I.Get(222, "Page", "get_text_attributes_for_area")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_area := gi.NewPointerArgument(area.P)
	args := []gi.Argument{arg_v, arg_area}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_page_get_text_for_area
//
// [ area ] trans: nothing
//
// [ result ] trans: everything
//
func (v Page) GetTextForArea(area Rectangle) (result string) {
	iv, err := _I.Get(223, "Page", "get_text_for_area")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_area := gi.NewPointerArgument(area.P)
	args := []gi.Argument{arg_v, arg_area}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// poppler_page_get_text_layout
//
// [ rectangles ] trans: container, dir: out
//
// [ n_rectangles ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Page) GetTextLayout() (result bool, rectangles unsafe.Pointer) {
	iv, err := _I.Get(224, "Page", "get_text_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_rectangles := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_rectangles := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_rectangles, arg_n_rectangles}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_rectangles uint32
	_ = n_rectangles
	rectangles = outArgs[0].Pointer()
	n_rectangles = outArgs[1].Uint32()
	result = ret.Bool()
	return
}

// poppler_page_get_text_layout_for_area
//
// [ area ] trans: nothing
//
// [ rectangles ] trans: container, dir: out
//
// [ n_rectangles ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Page) GetTextLayoutForArea(area Rectangle) (result bool, rectangles unsafe.Pointer) {
	iv, err := _I.Get(225, "Page", "get_text_layout_for_area")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_area := gi.NewPointerArgument(area.P)
	arg_rectangles := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_rectangles := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_area, arg_rectangles, arg_n_rectangles}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_rectangles uint32
	_ = n_rectangles
	rectangles = outArgs[0].Pointer()
	n_rectangles = outArgs[1].Uint32()
	result = ret.Bool()
	return
}

// poppler_page_get_thumbnail
//
// [ result ] trans: everything
//
func (v Page) GetThumbnail() (result cairo.Surface) {
	iv, err := _I.Get(226, "Page", "get_thumbnail")
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

// poppler_page_get_thumbnail_size
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Page) GetThumbnailSize() (result bool, width int32, height int32) {
	iv, err := _I.Get(227, "Page", "get_thumbnail_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// poppler_page_get_transition
//
// [ result ] trans: everything
//
func (v Page) GetTransition() (result PageTransition) {
	iv, err := _I.Get(228, "Page", "get_transition")
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

// poppler_page_remove_annot
//
// [ annot ] trans: nothing
//
func (v Page) RemoveAnnot(annot IAnnot) {
	iv, err := _I.Get(229, "Page", "remove_annot")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if annot != nil {
		tmp = annot.P_Annot()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_annot := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_annot}
	iv.Call(args, nil, nil)
}

// poppler_page_render
//
// [ cairo ] trans: nothing
//
func (v Page) Render(cairo cairo.Context) {
	iv, err := _I.Get(230, "Page", "render")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cairo := gi.NewPointerArgument(cairo.P)
	args := []gi.Argument{arg_v, arg_cairo}
	iv.Call(args, nil, nil)
}

// poppler_page_render_for_printing
//
// [ cairo ] trans: nothing
//
func (v Page) RenderForPrinting(cairo cairo.Context) {
	iv, err := _I.Get(231, "Page", "render_for_printing")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cairo := gi.NewPointerArgument(cairo.P)
	args := []gi.Argument{arg_v, arg_cairo}
	iv.Call(args, nil, nil)
}

// poppler_page_render_for_printing_with_options
//
// [ cairo ] trans: nothing
//
// [ options ] trans: nothing
//
func (v Page) RenderForPrintingWithOptions(cairo cairo.Context, options PrintFlags) {
	iv, err := _I.Get(232, "Page", "render_for_printing_with_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cairo := gi.NewPointerArgument(cairo.P)
	arg_options := gi.NewIntArgument(int(options))
	args := []gi.Argument{arg_v, arg_cairo, arg_options}
	iv.Call(args, nil, nil)
}

// poppler_page_render_selection
//
// [ cairo ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ old_selection ] trans: nothing
//
// [ style ] trans: nothing
//
// [ glyph_color ] trans: nothing
//
// [ background_color ] trans: nothing
//
func (v Page) RenderSelection(cairo cairo.Context, selection Rectangle, old_selection Rectangle, style SelectionStyleEnum, glyph_color Color, background_color Color) {
	iv, err := _I.Get(233, "Page", "render_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cairo := gi.NewPointerArgument(cairo.P)
	arg_selection := gi.NewPointerArgument(selection.P)
	arg_old_selection := gi.NewPointerArgument(old_selection.P)
	arg_style := gi.NewIntArgument(int(style))
	arg_glyph_color := gi.NewPointerArgument(glyph_color.P)
	arg_background_color := gi.NewPointerArgument(background_color.P)
	args := []gi.Argument{arg_v, arg_cairo, arg_selection, arg_old_selection, arg_style, arg_glyph_color, arg_background_color}
	iv.Call(args, nil, nil)
}

// poppler_page_render_to_ps
//
// [ ps_file ] trans: nothing
//
func (v Page) RenderToPs(ps_file IPSFile) {
	iv, err := _I.Get(234, "Page", "render_to_ps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if ps_file != nil {
		tmp = ps_file.P_PSFile()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ps_file := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_ps_file}
	iv.Call(args, nil, nil)
}

// Enum PageLayout
type PageLayoutEnum int

const (
	PageLayoutUnset          PageLayoutEnum = 0
	PageLayoutSinglePage     PageLayoutEnum = 1
	PageLayoutOneColumn      PageLayoutEnum = 2
	PageLayoutTwoColumnLeft  PageLayoutEnum = 3
	PageLayoutTwoColumnRight PageLayoutEnum = 4
	PageLayoutTwoPageLeft    PageLayoutEnum = 5
	PageLayoutTwoPageRight   PageLayoutEnum = 6
)

func PageLayoutGetType() gi.GType {
	ret := _I.GetGType(64, "PageLayout")
	return ret
}

// Enum PageMode
type PageModeEnum int

const (
	PageModeUnset          PageModeEnum = 0
	PageModeNone           PageModeEnum = 1
	PageModeUseOutlines    PageModeEnum = 2
	PageModeUseThumbs      PageModeEnum = 3
	PageModeFullScreen     PageModeEnum = 4
	PageModeUseOc          PageModeEnum = 5
	PageModeUseAttachments PageModeEnum = 6
)

func PageModeGetType() gi.GType {
	ret := _I.GetGType(65, "PageMode")
	return ret
}

// Struct PageTransition
type PageTransition struct {
	P unsafe.Pointer
}

const SizeOfStructPageTransition = 48

func PageTransitionGetType() gi.GType {
	ret := _I.GetGType(66, "PageTransition")
	return ret
}

// poppler_page_transition_new
//
// [ result ] trans: everything
//
func NewPageTransition() (result PageTransition) {
	iv, err := _I.Get(235, "PageTransition", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_page_transition_copy
//
// [ result ] trans: everything
//
func (v PageTransition) Copy() (result PageTransition) {
	iv, err := _I.Get(236, "PageTransition", "copy")
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

// poppler_page_transition_free
//
func (v PageTransition) Free() {
	iv, err := _I.Get(237, "PageTransition", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum PageTransitionAlignment
type PageTransitionAlignmentEnum int

const (
	PageTransitionAlignmentHorizontal PageTransitionAlignmentEnum = 0
	PageTransitionAlignmentVertical   PageTransitionAlignmentEnum = 1
)

func PageTransitionAlignmentGetType() gi.GType {
	ret := _I.GetGType(67, "PageTransitionAlignment")
	return ret
}

// Enum PageTransitionDirection
type PageTransitionDirectionEnum int

const (
	PageTransitionDirectionInward  PageTransitionDirectionEnum = 0
	PageTransitionDirectionOutward PageTransitionDirectionEnum = 1
)

func PageTransitionDirectionGetType() gi.GType {
	ret := _I.GetGType(68, "PageTransitionDirection")
	return ret
}

// Enum PageTransitionType
type PageTransitionTypeEnum int

const (
	PageTransitionTypeReplace  PageTransitionTypeEnum = 0
	PageTransitionTypeSplit    PageTransitionTypeEnum = 1
	PageTransitionTypeBlinds   PageTransitionTypeEnum = 2
	PageTransitionTypeBox      PageTransitionTypeEnum = 3
	PageTransitionTypeWipe     PageTransitionTypeEnum = 4
	PageTransitionTypeDissolve PageTransitionTypeEnum = 5
	PageTransitionTypeGlitter  PageTransitionTypeEnum = 6
	PageTransitionTypeFly      PageTransitionTypeEnum = 7
	PageTransitionTypePush     PageTransitionTypeEnum = 8
	PageTransitionTypeCover    PageTransitionTypeEnum = 9
	PageTransitionTypeUncover  PageTransitionTypeEnum = 10
	PageTransitionTypeFade     PageTransitionTypeEnum = 11
)

func PageTransitionTypeGetType() gi.GType {
	ret := _I.GetGType(69, "PageTransitionType")
	return ret
}

// Flags Permissions
type PermissionsFlags int

const (
	PermissionsOkToPrint               PermissionsFlags = 1
	PermissionsOkToModify              PermissionsFlags = 2
	PermissionsOkToCopy                PermissionsFlags = 4
	PermissionsOkToAddNotes            PermissionsFlags = 8
	PermissionsOkToFillForm            PermissionsFlags = 16
	PermissionsOkToExtractContents     PermissionsFlags = 32
	PermissionsOkToAssemble            PermissionsFlags = 64
	PermissionsOkToPrintHighResolution PermissionsFlags = 128
	PermissionsFull                    PermissionsFlags = 255
)

func PermissionsGetType() gi.GType {
	ret := _I.GetGType(70, "Permissions")
	return ret
}

// Struct Point
type Point struct {
	P unsafe.Pointer
}

const SizeOfStructPoint = 16

func PointGetType() gi.GType {
	ret := _I.GetGType(71, "Point")
	return ret
}

// poppler_point_new
//
// [ result ] trans: everything
//
func NewPoint() (result Point) {
	iv, err := _I.Get(238, "Point", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_point_copy
//
// [ result ] trans: everything
//
func (v Point) Copy() (result Point) {
	iv, err := _I.Get(239, "Point", "copy")
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

// poppler_point_free
//
func (v Point) Free() {
	iv, err := _I.Get(240, "Point", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags PrintFlags
type PrintFlags int

const (
	PrintFlagsDocument        PrintFlags = 0
	PrintFlagsMarkupAnnots    PrintFlags = 1
	PrintFlagsStampAnnotsOnly PrintFlags = 2
	PrintFlagsAll             PrintFlags = 1
)

func PrintFlagsGetType() gi.GType {
	ret := _I.GetGType(72, "PrintFlags")
	return ret
}

// Struct Quadrilateral
type Quadrilateral struct {
	P unsafe.Pointer
}

const SizeOfStructQuadrilateral = 64

func QuadrilateralGetType() gi.GType {
	ret := _I.GetGType(73, "Quadrilateral")
	return ret
}

// poppler_quadrilateral_new
//
// [ result ] trans: everything
//
func NewQuadrilateral() (result Quadrilateral) {
	iv, err := _I.Get(241, "Quadrilateral", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_quadrilateral_copy
//
// [ result ] trans: everything
//
func (v Quadrilateral) Copy() (result Quadrilateral) {
	iv, err := _I.Get(242, "Quadrilateral", "copy")
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

// poppler_quadrilateral_free
//
func (v Quadrilateral) Free() {
	iv, err := _I.Get(243, "Quadrilateral", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Rectangle
type Rectangle struct {
	P unsafe.Pointer
}

const SizeOfStructRectangle = 32

func RectangleGetType() gi.GType {
	ret := _I.GetGType(74, "Rectangle")
	return ret
}

// poppler_rectangle_new
//
// [ result ] trans: everything
//
func NewRectangle() (result Rectangle) {
	iv, err := _I.Get(244, "Rectangle", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_rectangle_copy
//
// [ result ] trans: everything
//
func (v Rectangle) Copy() (result Rectangle) {
	iv, err := _I.Get(245, "Rectangle", "copy")
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

// poppler_rectangle_free
//
func (v Rectangle) Free() {
	iv, err := _I.Get(246, "Rectangle", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum SelectionStyle
type SelectionStyleEnum int

const (
	SelectionStyleGlyph SelectionStyleEnum = 0
	SelectionStyleWord  SelectionStyleEnum = 1
	SelectionStyleLine  SelectionStyleEnum = 2
)

func SelectionStyleGetType() gi.GType {
	ret := _I.GetGType(75, "SelectionStyle")
	return ret
}

// Enum StructureBlockAlign
type StructureBlockAlignEnum int

const (
	StructureBlockAlignBefore  StructureBlockAlignEnum = 0
	StructureBlockAlignMiddle  StructureBlockAlignEnum = 1
	StructureBlockAlignAfter   StructureBlockAlignEnum = 2
	StructureBlockAlignJustify StructureBlockAlignEnum = 3
)

func StructureBlockAlignGetType() gi.GType {
	ret := _I.GetGType(76, "StructureBlockAlign")
	return ret
}

// Enum StructureBorderStyle
type StructureBorderStyleEnum int

const (
	StructureBorderStyleNone   StructureBorderStyleEnum = 0
	StructureBorderStyleHidden StructureBorderStyleEnum = 1
	StructureBorderStyleDotted StructureBorderStyleEnum = 2
	StructureBorderStyleDashed StructureBorderStyleEnum = 3
	StructureBorderStyleSolid  StructureBorderStyleEnum = 4
	StructureBorderStyleDouble StructureBorderStyleEnum = 5
	StructureBorderStyleGroove StructureBorderStyleEnum = 6
	StructureBorderStyleInset  StructureBorderStyleEnum = 7
	StructureBorderStyleOutset StructureBorderStyleEnum = 8
)

func StructureBorderStyleGetType() gi.GType {
	ret := _I.GetGType(77, "StructureBorderStyle")
	return ret
}

// Object StructureElement
type StructureElement struct {
	g.Object
}

func WrapStructureElement(p unsafe.Pointer) (r StructureElement) { r.P = p; return }

type IStructureElement interface{ P_StructureElement() unsafe.Pointer }

func (v StructureElement) P_StructureElement() unsafe.Pointer { return v.P }
func StructureElementGetType() gi.GType {
	ret := _I.GetGType(78, "StructureElement")
	return ret
}

// poppler_structure_element_get_abbreviation
//
// [ result ] trans: everything
//
func (v StructureElement) GetAbbreviation() (result string) {
	iv, err := _I.Get(247, "StructureElement", "get_abbreviation")
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

// poppler_structure_element_get_actual_text
//
// [ result ] trans: everything
//
func (v StructureElement) GetActualText() (result string) {
	iv, err := _I.Get(248, "StructureElement", "get_actual_text")
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

// poppler_structure_element_get_alt_text
//
// [ result ] trans: everything
//
func (v StructureElement) GetAltText() (result string) {
	iv, err := _I.Get(249, "StructureElement", "get_alt_text")
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

// poppler_structure_element_get_background_color
//
// [ color ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v StructureElement) GetBackgroundColor(color Color) (result bool) {
	iv, err := _I.Get(250, "StructureElement", "get_background_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_structure_element_get_baseline_shift
//
// [ result ] trans: nothing
//
func (v StructureElement) GetBaselineShift() (result float64) {
	iv, err := _I.Get(251, "StructureElement", "get_baseline_shift")
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

// poppler_structure_element_get_block_align
//
// [ result ] trans: nothing
//
func (v StructureElement) GetBlockAlign() (result StructureBlockAlignEnum) {
	iv, err := _I.Get(252, "StructureElement", "get_block_align")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureBlockAlignEnum(ret.Int())
	return
}

// poppler_structure_element_get_border_color
//
// [ colors ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v StructureElement) GetBorderColor(colors unsafe.Pointer) (result bool) {
	iv, err := _I.Get(253, "StructureElement", "get_border_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_colors := gi.NewPointerArgument(colors)
	args := []gi.Argument{arg_v, arg_colors}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_structure_element_get_border_style
//
// [ border_styles ] trans: everything, dir: out
//
func (v StructureElement) GetBorderStyle() (border_styles unsafe.Pointer) {
	iv, err := _I.Get(254, "StructureElement", "get_border_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_border_styles := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_border_styles}
	iv.Call(args, nil, &outArgs[0])
	border_styles = outArgs[0].Pointer()
	return
}

// poppler_structure_element_get_border_thickness
//
// [ border_thicknesses ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v StructureElement) GetBorderThickness() (result bool, border_thicknesses gi.DoubleArray) {
	iv, err := _I.Get(255, "StructureElement", "get_border_thickness")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_border_thicknesses := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_border_thicknesses}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	border_thicknesses.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// poppler_structure_element_get_bounding_box
//
// [ bounding_box ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v StructureElement) GetBoundingBox(bounding_box Rectangle) (result bool) {
	iv, err := _I.Get(256, "StructureElement", "get_bounding_box")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bounding_box := gi.NewPointerArgument(bounding_box.P)
	args := []gi.Argument{arg_v, arg_bounding_box}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_structure_element_get_color
//
// [ color ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v StructureElement) GetColor(color Color) (result bool) {
	iv, err := _I.Get(257, "StructureElement", "get_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_structure_element_get_column_count
//
// [ result ] trans: nothing
//
func (v StructureElement) GetColumnCount() (result uint32) {
	iv, err := _I.Get(258, "StructureElement", "get_column_count")
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

// poppler_structure_element_get_column_gaps
//
// [ n_values ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v StructureElement) GetColumnGaps() (result gi.DoubleArray) {
	iv, err := _I.Get(259, "StructureElement", "get_column_gaps")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_values := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_n_values}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_values uint32
	_ = n_values
	n_values = outArgs[0].Uint32()
	result = gi.DoubleArray{P: ret.Pointer(), Len: int(n_values)}
	return
}

// poppler_structure_element_get_column_widths
//
// [ n_values ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v StructureElement) GetColumnWidths() (result gi.DoubleArray) {
	iv, err := _I.Get(260, "StructureElement", "get_column_widths")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_values := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_n_values}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_values uint32
	_ = n_values
	n_values = outArgs[0].Uint32()
	result = gi.DoubleArray{P: ret.Pointer(), Len: int(n_values)}
	return
}

// poppler_structure_element_get_end_indent
//
// [ result ] trans: nothing
//
func (v StructureElement) GetEndIndent() (result float64) {
	iv, err := _I.Get(261, "StructureElement", "get_end_indent")
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

// poppler_structure_element_get_form_description
//
// [ result ] trans: everything
//
func (v StructureElement) GetFormDescription() (result string) {
	iv, err := _I.Get(262, "StructureElement", "get_form_description")
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

// poppler_structure_element_get_form_role
//
// [ result ] trans: nothing
//
func (v StructureElement) GetFormRole() (result StructureFormRoleEnum) {
	iv, err := _I.Get(263, "StructureElement", "get_form_role")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureFormRoleEnum(ret.Int())
	return
}

// poppler_structure_element_get_form_state
//
// [ result ] trans: nothing
//
func (v StructureElement) GetFormState() (result StructureFormStateEnum) {
	iv, err := _I.Get(264, "StructureElement", "get_form_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureFormStateEnum(ret.Int())
	return
}

// poppler_structure_element_get_glyph_orientation
//
// [ result ] trans: nothing
//
func (v StructureElement) GetGlyphOrientation() (result StructureGlyphOrientationEnum) {
	iv, err := _I.Get(265, "StructureElement", "get_glyph_orientation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureGlyphOrientationEnum(ret.Int())
	return
}

// poppler_structure_element_get_height
//
// [ result ] trans: nothing
//
func (v StructureElement) GetHeight() (result float64) {
	iv, err := _I.Get(266, "StructureElement", "get_height")
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

// poppler_structure_element_get_id
//
// [ result ] trans: everything
//
func (v StructureElement) GetId() (result string) {
	iv, err := _I.Get(267, "StructureElement", "get_id")
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

// poppler_structure_element_get_inline_align
//
// [ result ] trans: nothing
//
func (v StructureElement) GetInlineAlign() (result StructureInlineAlignEnum) {
	iv, err := _I.Get(268, "StructureElement", "get_inline_align")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureInlineAlignEnum(ret.Int())
	return
}

// poppler_structure_element_get_kind
//
// [ result ] trans: nothing
//
func (v StructureElement) GetKind() (result StructureElementKindEnum) {
	iv, err := _I.Get(269, "StructureElement", "get_kind")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureElementKindEnum(ret.Int())
	return
}

// poppler_structure_element_get_language
//
// [ result ] trans: everything
//
func (v StructureElement) GetLanguage() (result string) {
	iv, err := _I.Get(270, "StructureElement", "get_language")
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

// poppler_structure_element_get_line_height
//
// [ result ] trans: nothing
//
func (v StructureElement) GetLineHeight() (result float64) {
	iv, err := _I.Get(271, "StructureElement", "get_line_height")
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

// poppler_structure_element_get_list_numbering
//
// [ result ] trans: nothing
//
func (v StructureElement) GetListNumbering() (result StructureListNumberingEnum) {
	iv, err := _I.Get(272, "StructureElement", "get_list_numbering")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureListNumberingEnum(ret.Int())
	return
}

// poppler_structure_element_get_padding
//
// [ paddings ] trans: everything, dir: out
//
func (v StructureElement) GetPadding() (paddings gi.DoubleArray) {
	iv, err := _I.Get(273, "StructureElement", "get_padding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_paddings := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_paddings}
	iv.Call(args, nil, &outArgs[0])
	paddings.P = outArgs[0].Pointer()
	return
}

// poppler_structure_element_get_page
//
// [ result ] trans: nothing
//
func (v StructureElement) GetPage() (result int32) {
	iv, err := _I.Get(274, "StructureElement", "get_page")
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

// poppler_structure_element_get_placement
//
// [ result ] trans: nothing
//
func (v StructureElement) GetPlacement() (result StructurePlacementEnum) {
	iv, err := _I.Get(275, "StructureElement", "get_placement")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructurePlacementEnum(ret.Int())
	return
}

// poppler_structure_element_get_ruby_align
//
// [ result ] trans: nothing
//
func (v StructureElement) GetRubyAlign() (result StructureRubyAlignEnum) {
	iv, err := _I.Get(276, "StructureElement", "get_ruby_align")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureRubyAlignEnum(ret.Int())
	return
}

// poppler_structure_element_get_ruby_position
//
// [ result ] trans: nothing
//
func (v StructureElement) GetRubyPosition() (result StructureRubyPositionEnum) {
	iv, err := _I.Get(277, "StructureElement", "get_ruby_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureRubyPositionEnum(ret.Int())
	return
}

// poppler_structure_element_get_space_after
//
// [ result ] trans: nothing
//
func (v StructureElement) GetSpaceAfter() (result float64) {
	iv, err := _I.Get(278, "StructureElement", "get_space_after")
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

// poppler_structure_element_get_space_before
//
// [ result ] trans: nothing
//
func (v StructureElement) GetSpaceBefore() (result float64) {
	iv, err := _I.Get(279, "StructureElement", "get_space_before")
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

// poppler_structure_element_get_start_indent
//
// [ result ] trans: nothing
//
func (v StructureElement) GetStartIndent() (result float64) {
	iv, err := _I.Get(280, "StructureElement", "get_start_indent")
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

// poppler_structure_element_get_table_border_style
//
// [ border_styles ] trans: everything, dir: out
//
func (v StructureElement) GetTableBorderStyle() (border_styles unsafe.Pointer) {
	iv, err := _I.Get(281, "StructureElement", "get_table_border_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_border_styles := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_border_styles}
	iv.Call(args, nil, &outArgs[0])
	border_styles = outArgs[0].Pointer()
	return
}

// poppler_structure_element_get_table_column_span
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTableColumnSpan() (result uint32) {
	iv, err := _I.Get(282, "StructureElement", "get_table_column_span")
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

// poppler_structure_element_get_table_headers
//
// [ result ] trans: everything
//
func (v StructureElement) GetTableHeaders() (result gi.CStrArray) {
	iv, err := _I.Get(283, "StructureElement", "get_table_headers")
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

// poppler_structure_element_get_table_padding
//
// [ paddings ] trans: everything, dir: out
//
func (v StructureElement) GetTablePadding() (paddings gi.DoubleArray) {
	iv, err := _I.Get(284, "StructureElement", "get_table_padding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_paddings := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_paddings}
	iv.Call(args, nil, &outArgs[0])
	paddings.P = outArgs[0].Pointer()
	return
}

// poppler_structure_element_get_table_row_span
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTableRowSpan() (result uint32) {
	iv, err := _I.Get(285, "StructureElement", "get_table_row_span")
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

// poppler_structure_element_get_table_scope
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTableScope() (result StructureTableScopeEnum) {
	iv, err := _I.Get(286, "StructureElement", "get_table_scope")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureTableScopeEnum(ret.Int())
	return
}

// poppler_structure_element_get_table_summary
//
// [ result ] trans: everything
//
func (v StructureElement) GetTableSummary() (result string) {
	iv, err := _I.Get(287, "StructureElement", "get_table_summary")
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

// poppler_structure_element_get_text
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func (v StructureElement) GetText(flags StructureGetTextFlags) (result string) {
	iv, err := _I.Get(288, "StructureElement", "get_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// poppler_structure_element_get_text_align
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTextAlign() (result StructureTextAlignEnum) {
	iv, err := _I.Get(289, "StructureElement", "get_text_align")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureTextAlignEnum(ret.Int())
	return
}

// poppler_structure_element_get_text_decoration_color
//
// [ color ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTextDecorationColor(color Color) (result bool) {
	iv, err := _I.Get(290, "StructureElement", "get_text_decoration_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// poppler_structure_element_get_text_decoration_thickness
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTextDecorationThickness() (result float64) {
	iv, err := _I.Get(291, "StructureElement", "get_text_decoration_thickness")
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

// poppler_structure_element_get_text_decoration_type
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTextDecorationType() (result StructureTextDecorationEnum) {
	iv, err := _I.Get(292, "StructureElement", "get_text_decoration_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureTextDecorationEnum(ret.Int())
	return
}

// poppler_structure_element_get_text_indent
//
// [ result ] trans: nothing
//
func (v StructureElement) GetTextIndent() (result float64) {
	iv, err := _I.Get(293, "StructureElement", "get_text_indent")
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

// poppler_structure_element_get_text_spans
//
// [ n_text_spans ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v StructureElement) GetTextSpans() (result unsafe.Pointer) {
	iv, err := _I.Get(294, "StructureElement", "get_text_spans")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_text_spans := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_n_text_spans}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_text_spans uint32
	_ = n_text_spans
	n_text_spans = outArgs[0].Uint32()
	result = ret.Pointer()
	return
}

// poppler_structure_element_get_title
//
// [ result ] trans: everything
//
func (v StructureElement) GetTitle() (result string) {
	iv, err := _I.Get(295, "StructureElement", "get_title")
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

// poppler_structure_element_get_width
//
// [ result ] trans: nothing
//
func (v StructureElement) GetWidth() (result float64) {
	iv, err := _I.Get(296, "StructureElement", "get_width")
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

// poppler_structure_element_get_writing_mode
//
// [ result ] trans: nothing
//
func (v StructureElement) GetWritingMode() (result StructureWritingModeEnum) {
	iv, err := _I.Get(297, "StructureElement", "get_writing_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StructureWritingModeEnum(ret.Int())
	return
}

// poppler_structure_element_is_block
//
// [ result ] trans: nothing
//
func (v StructureElement) IsBlock() (result bool) {
	iv, err := _I.Get(298, "StructureElement", "is_block")
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

// poppler_structure_element_is_content
//
// [ result ] trans: nothing
//
func (v StructureElement) IsContent() (result bool) {
	iv, err := _I.Get(299, "StructureElement", "is_content")
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

// poppler_structure_element_is_grouping
//
// [ result ] trans: nothing
//
func (v StructureElement) IsGrouping() (result bool) {
	iv, err := _I.Get(300, "StructureElement", "is_grouping")
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

// poppler_structure_element_is_inline
//
// [ result ] trans: nothing
//
func (v StructureElement) IsInline() (result bool) {
	iv, err := _I.Get(301, "StructureElement", "is_inline")
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

// Struct StructureElementIter
type StructureElementIter struct {
	P unsafe.Pointer
}

func StructureElementIterGetType() gi.GType {
	ret := _I.GetGType(79, "StructureElementIter")
	return ret
}

// poppler_structure_element_iter_new
//
// [ poppler_document ] trans: nothing
//
// [ result ] trans: everything
//
func NewStructureElementIter(poppler_document IDocument) (result StructureElementIter) {
	iv, err := _I.Get(302, "StructureElementIter", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if poppler_document != nil {
		tmp = poppler_document.P_Document()
	}
	arg_poppler_document := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_poppler_document}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_structure_element_iter_copy
//
// [ result ] trans: everything
//
func (v StructureElementIter) Copy() (result StructureElementIter) {
	iv, err := _I.Get(303, "StructureElementIter", "copy")
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

// poppler_structure_element_iter_free
//
func (v StructureElementIter) Free() {
	iv, err := _I.Get(304, "StructureElementIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_structure_element_iter_get_child
//
// [ result ] trans: everything
//
func (v StructureElementIter) GetChild() (result StructureElementIter) {
	iv, err := _I.Get(305, "StructureElementIter", "get_child")
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

// poppler_structure_element_iter_get_element
//
// [ result ] trans: everything
//
func (v StructureElementIter) GetElement() (result StructureElement) {
	iv, err := _I.Get(306, "StructureElementIter", "get_element")
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

// poppler_structure_element_iter_next
//
// [ result ] trans: nothing
//
func (v StructureElementIter) Next() (result bool) {
	iv, err := _I.Get(307, "StructureElementIter", "next")
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

// Enum StructureElementKind
type StructureElementKindEnum int

const (
	StructureElementKindContent            StructureElementKindEnum = 0
	StructureElementKindObjectReference    StructureElementKindEnum = 1
	StructureElementKindDocument           StructureElementKindEnum = 2
	StructureElementKindPart               StructureElementKindEnum = 3
	StructureElementKindArticle            StructureElementKindEnum = 4
	StructureElementKindSection            StructureElementKindEnum = 5
	StructureElementKindDiv                StructureElementKindEnum = 6
	StructureElementKindSpan               StructureElementKindEnum = 7
	StructureElementKindQuote              StructureElementKindEnum = 8
	StructureElementKindNote               StructureElementKindEnum = 9
	StructureElementKindReference          StructureElementKindEnum = 10
	StructureElementKindBibentry           StructureElementKindEnum = 11
	StructureElementKindCode               StructureElementKindEnum = 12
	StructureElementKindLink               StructureElementKindEnum = 13
	StructureElementKindAnnot              StructureElementKindEnum = 14
	StructureElementKindBlockquote         StructureElementKindEnum = 15
	StructureElementKindCaption            StructureElementKindEnum = 16
	StructureElementKindNonstruct          StructureElementKindEnum = 17
	StructureElementKindToc                StructureElementKindEnum = 18
	StructureElementKindTocItem            StructureElementKindEnum = 19
	StructureElementKindIndex              StructureElementKindEnum = 20
	StructureElementKindPrivate            StructureElementKindEnum = 21
	StructureElementKindParagraph          StructureElementKindEnum = 22
	StructureElementKindHeading            StructureElementKindEnum = 23
	StructureElementKindHeading1           StructureElementKindEnum = 24
	StructureElementKindHeading2           StructureElementKindEnum = 25
	StructureElementKindHeading3           StructureElementKindEnum = 26
	StructureElementKindHeading4           StructureElementKindEnum = 27
	StructureElementKindHeading5           StructureElementKindEnum = 28
	StructureElementKindHeading6           StructureElementKindEnum = 29
	StructureElementKindList               StructureElementKindEnum = 30
	StructureElementKindListItem           StructureElementKindEnum = 31
	StructureElementKindListLabel          StructureElementKindEnum = 32
	StructureElementKindListBody           StructureElementKindEnum = 33
	StructureElementKindTable              StructureElementKindEnum = 34
	StructureElementKindTableRow           StructureElementKindEnum = 35
	StructureElementKindTableHeading       StructureElementKindEnum = 36
	StructureElementKindTableData          StructureElementKindEnum = 37
	StructureElementKindTableHeader        StructureElementKindEnum = 38
	StructureElementKindTableFooter        StructureElementKindEnum = 39
	StructureElementKindTableBody          StructureElementKindEnum = 40
	StructureElementKindRuby               StructureElementKindEnum = 41
	StructureElementKindRubyBaseText       StructureElementKindEnum = 42
	StructureElementKindRubyAnnotText      StructureElementKindEnum = 43
	StructureElementKindRubyPunctuation    StructureElementKindEnum = 44
	StructureElementKindWarichu            StructureElementKindEnum = 45
	StructureElementKindWarichuText        StructureElementKindEnum = 46
	StructureElementKindWarichuPunctuation StructureElementKindEnum = 47
	StructureElementKindFigure             StructureElementKindEnum = 48
	StructureElementKindFormula            StructureElementKindEnum = 49
	StructureElementKindForm               StructureElementKindEnum = 50
)

func StructureElementKindGetType() gi.GType {
	ret := _I.GetGType(80, "StructureElementKind")
	return ret
}

// Enum StructureFormRole
type StructureFormRoleEnum int

const (
	StructureFormRoleUndefined   StructureFormRoleEnum = 0
	StructureFormRoleRadioButton StructureFormRoleEnum = 1
	StructureFormRolePushButton  StructureFormRoleEnum = 2
	StructureFormRoleTextValue   StructureFormRoleEnum = 3
	StructureFormRoleCheckbox    StructureFormRoleEnum = 4
)

func StructureFormRoleGetType() gi.GType {
	ret := _I.GetGType(81, "StructureFormRole")
	return ret
}

// Enum StructureFormState
type StructureFormStateEnum int

const (
	StructureFormStateOn      StructureFormStateEnum = 0
	StructureFormStateOff     StructureFormStateEnum = 1
	StructureFormStateNeutral StructureFormStateEnum = 2
)

func StructureFormStateGetType() gi.GType {
	ret := _I.GetGType(82, "StructureFormState")
	return ret
}

// Flags StructureGetTextFlags
type StructureGetTextFlags int

const (
	StructureGetTextFlagsNone      StructureGetTextFlags = 0
	StructureGetTextFlagsRecursive StructureGetTextFlags = 1
)

func StructureGetTextFlagsGetType() gi.GType {
	ret := _I.GetGType(83, "StructureGetTextFlags")
	return ret
}

// Enum StructureGlyphOrientation
type StructureGlyphOrientationEnum int

const (
	StructureGlyphOrientationAuto StructureGlyphOrientationEnum = 0
	StructureGlyphOrientation0    StructureGlyphOrientationEnum = 0
	StructureGlyphOrientation90   StructureGlyphOrientationEnum = 1
	StructureGlyphOrientation180  StructureGlyphOrientationEnum = 2
	StructureGlyphOrientation270  StructureGlyphOrientationEnum = 3
)

func StructureGlyphOrientationGetType() gi.GType {
	ret := _I.GetGType(84, "StructureGlyphOrientation")
	return ret
}

// Enum StructureInlineAlign
type StructureInlineAlignEnum int

const (
	StructureInlineAlignStart  StructureInlineAlignEnum = 0
	StructureInlineAlignCenter StructureInlineAlignEnum = 1
	StructureInlineAlignEnd    StructureInlineAlignEnum = 2
)

func StructureInlineAlignGetType() gi.GType {
	ret := _I.GetGType(85, "StructureInlineAlign")
	return ret
}

// Enum StructureListNumbering
type StructureListNumberingEnum int

const (
	StructureListNumberingNone       StructureListNumberingEnum = 0
	StructureListNumberingDisc       StructureListNumberingEnum = 1
	StructureListNumberingCircle     StructureListNumberingEnum = 2
	StructureListNumberingSquare     StructureListNumberingEnum = 3
	StructureListNumberingDecimal    StructureListNumberingEnum = 4
	StructureListNumberingUpperRoman StructureListNumberingEnum = 5
	StructureListNumberingLowerRoman StructureListNumberingEnum = 6
	StructureListNumberingUpperAlpha StructureListNumberingEnum = 7
	StructureListNumberingLowerAlpha StructureListNumberingEnum = 8
)

func StructureListNumberingGetType() gi.GType {
	ret := _I.GetGType(86, "StructureListNumbering")
	return ret
}

// Enum StructurePlacement
type StructurePlacementEnum int

const (
	StructurePlacementBlock  StructurePlacementEnum = 0
	StructurePlacementInline StructurePlacementEnum = 1
	StructurePlacementBefore StructurePlacementEnum = 2
	StructurePlacementStart  StructurePlacementEnum = 3
	StructurePlacementEnd    StructurePlacementEnum = 4
)

func StructurePlacementGetType() gi.GType {
	ret := _I.GetGType(87, "StructurePlacement")
	return ret
}

// Enum StructureRubyAlign
type StructureRubyAlignEnum int

const (
	StructureRubyAlignStart      StructureRubyAlignEnum = 0
	StructureRubyAlignCenter     StructureRubyAlignEnum = 1
	StructureRubyAlignEnd        StructureRubyAlignEnum = 2
	StructureRubyAlignJustify    StructureRubyAlignEnum = 3
	StructureRubyAlignDistribute StructureRubyAlignEnum = 4
)

func StructureRubyAlignGetType() gi.GType {
	ret := _I.GetGType(88, "StructureRubyAlign")
	return ret
}

// Enum StructureRubyPosition
type StructureRubyPositionEnum int

const (
	StructureRubyPositionBefore  StructureRubyPositionEnum = 0
	StructureRubyPositionAfter   StructureRubyPositionEnum = 1
	StructureRubyPositionWarichu StructureRubyPositionEnum = 2
	StructureRubyPositionInline  StructureRubyPositionEnum = 3
)

func StructureRubyPositionGetType() gi.GType {
	ret := _I.GetGType(89, "StructureRubyPosition")
	return ret
}

// Enum StructureTableScope
type StructureTableScopeEnum int

const (
	StructureTableScopeRow    StructureTableScopeEnum = 0
	StructureTableScopeColumn StructureTableScopeEnum = 1
	StructureTableScopeBoth   StructureTableScopeEnum = 2
)

func StructureTableScopeGetType() gi.GType {
	ret := _I.GetGType(90, "StructureTableScope")
	return ret
}

// Enum StructureTextAlign
type StructureTextAlignEnum int

const (
	StructureTextAlignStart   StructureTextAlignEnum = 0
	StructureTextAlignCenter  StructureTextAlignEnum = 1
	StructureTextAlignEnd     StructureTextAlignEnum = 2
	StructureTextAlignJustify StructureTextAlignEnum = 3
)

func StructureTextAlignGetType() gi.GType {
	ret := _I.GetGType(91, "StructureTextAlign")
	return ret
}

// Enum StructureTextDecoration
type StructureTextDecorationEnum int

const (
	StructureTextDecorationNone        StructureTextDecorationEnum = 0
	StructureTextDecorationUnderline   StructureTextDecorationEnum = 1
	StructureTextDecorationOverline    StructureTextDecorationEnum = 2
	StructureTextDecorationLinethrough StructureTextDecorationEnum = 3
)

func StructureTextDecorationGetType() gi.GType {
	ret := _I.GetGType(92, "StructureTextDecoration")
	return ret
}

// Enum StructureWritingMode
type StructureWritingModeEnum int

const (
	StructureWritingModeLrTb StructureWritingModeEnum = 0
	StructureWritingModeRlTb StructureWritingModeEnum = 1
	StructureWritingModeTbRl StructureWritingModeEnum = 2
)

func StructureWritingModeGetType() gi.GType {
	ret := _I.GetGType(93, "StructureWritingMode")
	return ret
}

// Struct TextAttributes
type TextAttributes struct {
	P unsafe.Pointer
}

const SizeOfStructTextAttributes = 40

func TextAttributesGetType() gi.GType {
	ret := _I.GetGType(94, "TextAttributes")
	return ret
}

// poppler_text_attributes_new
//
// [ result ] trans: everything
//
func NewTextAttributes() (result TextAttributes) {
	iv, err := _I.Get(308, "TextAttributes", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// poppler_text_attributes_copy
//
// [ result ] trans: everything
//
func (v TextAttributes) Copy() (result TextAttributes) {
	iv, err := _I.Get(309, "TextAttributes", "copy")
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

// poppler_text_attributes_free
//
func (v TextAttributes) Free() {
	iv, err := _I.Get(310, "TextAttributes", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct TextSpan
type TextSpan struct {
	P unsafe.Pointer
}

func TextSpanGetType() gi.GType {
	ret := _I.GetGType(95, "TextSpan")
	return ret
}

// poppler_text_span_copy
//
// [ result ] trans: everything
//
func (v TextSpan) Copy() (result TextSpan) {
	iv, err := _I.Get(311, "TextSpan", "copy")
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

// poppler_text_span_free
//
func (v TextSpan) Free() {
	iv, err := _I.Get(312, "TextSpan", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// poppler_text_span_get_color
//
// [ color ] trans: nothing, dir: out
//
func (v TextSpan) GetColor(color Color) {
	iv, err := _I.Get(313, "TextSpan", "get_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	iv.Call(args, nil, nil)
}

// poppler_text_span_get_font_name
//
// [ result ] trans: nothing
//
func (v TextSpan) GetFontName() (result string) {
	iv, err := _I.Get(314, "TextSpan", "get_font_name")
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

// poppler_text_span_get_text
//
// [ result ] trans: nothing
//
func (v TextSpan) GetText() (result string) {
	iv, err := _I.Get(315, "TextSpan", "get_text")
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

// poppler_text_span_is_bold_font
//
// [ result ] trans: nothing
//
func (v TextSpan) IsBoldFont() (result bool) {
	iv, err := _I.Get(316, "TextSpan", "is_bold_font")
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

// poppler_text_span_is_fixed_width_font
//
// [ result ] trans: nothing
//
func (v TextSpan) IsFixedWidthFont() (result bool) {
	iv, err := _I.Get(317, "TextSpan", "is_fixed_width_font")
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

// poppler_text_span_is_serif_font
//
// [ result ] trans: nothing
//
func (v TextSpan) IsSerifFont() (result bool) {
	iv, err := _I.Get(318, "TextSpan", "is_serif_font")
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

// Flags ViewerPreferences
type ViewerPreferencesFlags int

const (
	ViewerPreferencesUnset           ViewerPreferencesFlags = 0
	ViewerPreferencesHideToolbar     ViewerPreferencesFlags = 1
	ViewerPreferencesHideMenubar     ViewerPreferencesFlags = 2
	ViewerPreferencesHideWindowui    ViewerPreferencesFlags = 4
	ViewerPreferencesFitWindow       ViewerPreferencesFlags = 8
	ViewerPreferencesCenterWindow    ViewerPreferencesFlags = 16
	ViewerPreferencesDisplayDocTitle ViewerPreferencesFlags = 32
	ViewerPreferencesDirectionRtl    ViewerPreferencesFlags = 64
)

func ViewerPreferencesGetType() gi.GType {
	ret := _I.GetGType(96, "ViewerPreferences")
	return ret
}

// poppler_date_parse
//
// [ date ] trans: nothing
//
// [ timet ] trans: nothing
//
// [ result ] trans: nothing
//
func DateParse(date string, timet int64) (result bool) {
	iv, err := _I.Get(319, "date_parse", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_date := gi.CString(date)
	arg_date := gi.NewStringArgument(c_date)
	arg_timet := gi.NewInt64Argument(timet)
	args := []gi.Argument{arg_date, arg_timet}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_date)
	result = ret.Bool()
	return
}

// poppler_error_quark
//
// [ result ] trans: nothing
//
func ErrorQuark() (result uint32) {
	iv, err := _I.Get(320, "error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// poppler_get_backend
//
// [ result ] trans: nothing
//
func GetBackend() (result BackendEnum) {
	iv, err := _I.Get(321, "get_backend", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = BackendEnum(ret.Int())
	return
}

// poppler_get_version
//
// [ result ] trans: nothing
//
func GetVersion() (result string) {
	iv, err := _I.Get(322, "get_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// constants
const (
	ANNOT_TEXT_ICON_CIRCLE        = "Circle"
	ANNOT_TEXT_ICON_COMMENT       = "Comment"
	ANNOT_TEXT_ICON_CROSS         = "Cross"
	ANNOT_TEXT_ICON_HELP          = "Help"
	ANNOT_TEXT_ICON_INSERT        = "Insert"
	ANNOT_TEXT_ICON_KEY           = "Key"
	ANNOT_TEXT_ICON_NEW_PARAGRAPH = "NewParagraph"
	ANNOT_TEXT_ICON_NOTE          = "Note"
	ANNOT_TEXT_ICON_PARAGRAPH     = "Paragraph"
	HAS_CAIRO                     = 1
	MAJOR_VERSION                 = 0
	MICRO_VERSION                 = 0
	MINOR_VERSION                 = 71
)
