package girepository

/*
#cgo pkg-config: gobject-introspection-1.0
#include <girepository.h>
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GIRepository")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GIRepository", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Union Argument
type Argument struct {
	P unsafe.Pointer
}

const SizeOfUnionArgument = 8

func ArgumentGetType() gi.GType {
	ret := _I.GetGType(0, "Argument")
	return ret
}

// Enum ArrayType
type ArrayTypeEnum int

const (
	ArrayTypeC         ArrayTypeEnum = 0
	ArrayTypeArray     ArrayTypeEnum = 1
	ArrayTypePtrArray  ArrayTypeEnum = 2
	ArrayTypeByteArray ArrayTypeEnum = 3
)

func ArrayTypeGetType() gi.GType {
	ret := _I.GetGType(1, "ArrayType")
	return ret
}

// Struct AttributeIter
type AttributeIter struct {
	P unsafe.Pointer
}

const SizeOfStructAttributeIter = 32

func AttributeIterGetType() gi.GType {
	ret := _I.GetGType(2, "AttributeIter")
	return ret
}

// Struct BaseInfo
type BaseInfo struct {
	P unsafe.Pointer
}

const SizeOfStructBaseInfo = 72

func BaseInfoGetType() gi.GType {
	ret := _I.GetGType(3, "BaseInfo")
	return ret
}

// g_base_info_equal
//
// [ info2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseInfo) Equal(info2 BaseInfo) (result bool) {
	iv, err := _I.Get(0, "BaseInfo", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_info2 := gi.NewPointerArgument(info2.P)
	args := []gi.Argument{arg_v, arg_info2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_base_info_get_attribute
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v BaseInfo) GetAttribute(name string) (result string) {
	iv, err := _I.Get(1, "BaseInfo", "get_attribute")
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
	result = ret.String().Copy()
	return
}

// g_base_info_get_container
//
// [ result ] trans: nothing
//
func (v BaseInfo) GetContainer() (result BaseInfo) {
	iv, err := _I.Get(2, "BaseInfo", "get_container")
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

// g_base_info_get_name
//
// [ result ] trans: nothing
//
func (v BaseInfo) GetName() (result string) {
	iv, err := _I.Get(3, "BaseInfo", "get_name")
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

// g_base_info_get_namespace
//
// [ result ] trans: nothing
//
func (v BaseInfo) GetNamespace() (result string) {
	iv, err := _I.Get(4, "BaseInfo", "get_namespace")
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

// g_base_info_get_type
//
// [ result ] trans: nothing
//
func (v BaseInfo) GetType() (result InfoTypeEnum) {
	iv, err := _I.Get(5, "BaseInfo", "get_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = InfoTypeEnum(ret.Int())
	return
}

// g_base_info_get_typelib
//
// [ result ] trans: nothing
//
func (v BaseInfo) GetTypelib() (result Typelib) {
	iv, err := _I.Get(6, "BaseInfo", "get_typelib")
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

// g_base_info_is_deprecated
//
// [ result ] trans: nothing
//
func (v BaseInfo) IsDeprecated() (result bool) {
	iv, err := _I.Get(7, "BaseInfo", "is_deprecated")
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

// g_base_info_iterate_attributes
//
// [ iterator ] trans: everything, dir: inout
//
// [ name ] trans: nothing, dir: out
//
// [ value ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v BaseInfo) IterateAttributes(iterator int /*TODO:TYPE*/) (result bool, name string, value string) {
	iv, err := _I.Get(8, "BaseInfo", "iterate_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	name = outArgs[0].String().Copy()
	value = outArgs[1].String().Copy()
	result = ret.Bool()
	return
}

// Enum Direction
type DirectionEnum int

const (
	DirectionIn    DirectionEnum = 0
	DirectionOut   DirectionEnum = 1
	DirectionInout DirectionEnum = 2
)

func DirectionGetType() gi.GType {
	ret := _I.GetGType(4, "Direction")
	return ret
}

// Flags FieldInfoFlags
type FieldInfoFlags int

const (
	FieldInfoFlagsReadable FieldInfoFlags = 1
	FieldInfoFlagsWritable FieldInfoFlags = 2
)

func FieldInfoFlagsGetType() gi.GType {
	ret := _I.GetGType(5, "FieldInfoFlags")
	return ret
}

// Flags FunctionInfoFlags
type FunctionInfoFlags int

const (
	FunctionInfoFlagsIsMethod      FunctionInfoFlags = 1
	FunctionInfoFlagsIsConstructor FunctionInfoFlags = 2
	FunctionInfoFlagsIsGetter      FunctionInfoFlags = 4
	FunctionInfoFlagsIsSetter      FunctionInfoFlags = 8
	FunctionInfoFlagsWrapsVfunc    FunctionInfoFlags = 16
	FunctionInfoFlagsThrows        FunctionInfoFlags = 32
)

func FunctionInfoFlagsGetType() gi.GType {
	ret := _I.GetGType(6, "FunctionInfoFlags")
	return ret
}

// Enum InfoType
type InfoTypeEnum int

const (
	InfoTypeInvalid    InfoTypeEnum = 0
	InfoTypeFunction   InfoTypeEnum = 1
	InfoTypeCallback   InfoTypeEnum = 2
	InfoTypeStruct     InfoTypeEnum = 3
	InfoTypeBoxed      InfoTypeEnum = 4
	InfoTypeEnum       InfoTypeEnum = 5
	InfoTypeFlags      InfoTypeEnum = 6
	InfoTypeObject     InfoTypeEnum = 7
	InfoTypeInterface  InfoTypeEnum = 8
	InfoTypeConstant   InfoTypeEnum = 9
	InfoTypeInvalid0   InfoTypeEnum = 10
	InfoTypeUnion      InfoTypeEnum = 11
	InfoTypeValue      InfoTypeEnum = 12
	InfoTypeSignal     InfoTypeEnum = 13
	InfoTypeVfunc      InfoTypeEnum = 14
	InfoTypeProperty   InfoTypeEnum = 15
	InfoTypeField      InfoTypeEnum = 16
	InfoTypeArg        InfoTypeEnum = 17
	InfoTypeType       InfoTypeEnum = 18
	InfoTypeUnresolved InfoTypeEnum = 19
)

func InfoTypeGetType() gi.GType {
	ret := _I.GetGType(7, "InfoType")
	return ret
}

// Object Repository
type Repository struct {
	g.Object
}

func WrapRepository(p unsafe.Pointer) (r Repository) { r.P = p; return }

type IRepository interface{ P_Repository() unsafe.Pointer }

func (v Repository) P_Repository() unsafe.Pointer { return v.P }
func RepositoryGetType() gi.GType {
	ret := _I.GetGType(8, "Repository")
	return ret
}

// g_irepository_dump
//
// [ arg ] trans: nothing
//
// [ result ] trans: nothing
//
func RepositoryDump1(arg string) (result bool, err error) {
	iv, err := _I.Get(9, "Repository", "dump")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_arg := gi.CString(arg)
	arg_arg := gi.NewStringArgument(c_arg)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_arg, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_arg)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_irepository_prepend_library_path
//
// [ directory ] trans: nothing
//
func RepositoryPrependLibraryPath1(directory string) {
	iv, err := _I.Get(14, "Repository", "prepend_library_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_directory := gi.CString(directory)
	arg_directory := gi.NewStringArgument(c_directory)
	args := []gi.Argument{arg_directory}
	iv.Call(args, nil, nil)
	gi.Free(c_directory)
}

// g_irepository_prepend_search_path
//
// [ directory ] trans: nothing
//
func RepositoryPrependSearchPath1(directory string) {
	iv, err := _I.Get(15, "Repository", "prepend_search_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_directory := gi.CString(directory)
	arg_directory := gi.NewStringArgument(c_directory)
	args := []gi.Argument{arg_directory}
	iv.Call(args, nil, nil)
	gi.Free(c_directory)
}

// g_irepository_enumerate_versions
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) EnumerateVersions(namespace_ string) (result g.List) {
	iv, err := _I.Get(16, "Repository", "enumerate_versions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result.P = ret.Pointer()
	return
}

// g_irepository_find_by_error_domain
//
// [ domain ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) FindByErrorDomain(domain uint32) (result BaseInfo) {
	iv, err := _I.Get(17, "Repository", "find_by_error_domain")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewUint32Argument(domain)
	args := []gi.Argument{arg_v, arg_domain}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_irepository_find_by_gtype
//
// [ gtype ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) FindByGtype(gtype gi.GType) (result BaseInfo) {
	iv, err := _I.Get(18, "Repository", "find_by_gtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gtype := gi.NewUintArgument(uint(gtype))
	args := []gi.Argument{arg_v, arg_gtype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_irepository_find_by_name
//
// [ namespace_ ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) FindByName(namespace_ string, name string) (result BaseInfo) {
	iv, err := _I.Get(19, "Repository", "find_by_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_namespace_, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_irepository_get_c_prefix
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) GetCPrefix(namespace_ string) (result string) {
	iv, err := _I.Get(20, "Repository", "get_c_prefix")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = ret.String().Copy()
	return
}

// g_irepository_get_dependencies
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) GetDependencies(namespace_ string) (result gi.CStrArray) {
	iv, err := _I.Get(21, "Repository", "get_dependencies")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_irepository_get_immediate_dependencies
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) GetImmediateDependencies(namespace_ string) (result gi.CStrArray) {
	iv, err := _I.Get(22, "Repository", "get_immediate_dependencies")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_irepository_get_info
//
// [ namespace_ ] trans: nothing
//
// [ index ] trans: nothing
//
// [ result ] trans: everything
//
func (v Repository) GetInfo(namespace_ string, index int32) (result BaseInfo) {
	iv, err := _I.Get(23, "Repository", "get_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_v, arg_namespace_, arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result.P = ret.Pointer()
	return
}

// g_irepository_get_loaded_namespaces
//
// [ result ] trans: everything
//
func (v Repository) GetLoadedNamespaces() (result gi.CStrArray) {
	iv, err := _I.Get(24, "Repository", "get_loaded_namespaces")
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

// g_irepository_get_n_infos
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) GetNInfos(namespace_ string) (result int32) {
	iv, err := _I.Get(25, "Repository", "get_n_infos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = ret.Int32()
	return
}

// g_irepository_get_shared_library
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) GetSharedLibrary(namespace_ string) (result string) {
	iv, err := _I.Get(26, "Repository", "get_shared_library")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = ret.String().Copy()
	return
}

// g_irepository_get_typelib_path
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) GetTypelibPath(namespace_ string) (result string) {
	iv, err := _I.Get(27, "Repository", "get_typelib_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = ret.String().Copy()
	return
}

// g_irepository_get_version
//
// [ namespace_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) GetVersion(namespace_ string) (result string) {
	iv, err := _I.Get(28, "Repository", "get_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	args := []gi.Argument{arg_v, arg_namespace_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	result = ret.String().Copy()
	return
}

// g_irepository_is_registered
//
// [ namespace_ ] trans: nothing
//
// [ version ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) IsRegistered(namespace_ string, version string) (result bool) {
	iv, err := _I.Get(29, "Repository", "is_registered")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_namespace_ := gi.CString(namespace_)
	c_version := gi.CString(version)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	arg_version := gi.NewStringArgument(c_version)
	args := []gi.Argument{arg_v, arg_namespace_, arg_version}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_namespace_)
	gi.Free(c_version)
	result = ret.Bool()
	return
}

// g_irepository_load_typelib
//
// [ typelib ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) LoadTypelib(typelib Typelib, flags RepositoryLoadFlags) (result string, err error) {
	iv, err := _I.Get(30, "Repository", "load_typelib")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_typelib := gi.NewPointerArgument(typelib.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_typelib, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Copy()
	return
}

// g_irepository_require
//
// [ namespace_ ] trans: nothing
//
// [ version ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) Require(namespace_ string, version string, flags RepositoryLoadFlags) (result Typelib, err error) {
	iv, err := _I.Get(31, "Repository", "require")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_namespace_ := gi.CString(namespace_)
	c_version := gi.CString(version)
	arg_v := gi.NewPointerArgument(v.P)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	arg_version := gi.NewStringArgument(c_version)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_namespace_, arg_version, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_namespace_)
	gi.Free(c_version)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_irepository_require_private
//
// [ typelib_dir ] trans: nothing
//
// [ namespace_ ] trans: nothing
//
// [ version ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Repository) RequirePrivate(typelib_dir string, namespace_ string, version string, flags RepositoryLoadFlags) (result Typelib, err error) {
	iv, err := _I.Get(32, "Repository", "require_private")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_typelib_dir := gi.CString(typelib_dir)
	c_namespace_ := gi.CString(namespace_)
	c_version := gi.CString(version)
	arg_v := gi.NewPointerArgument(v.P)
	arg_typelib_dir := gi.NewStringArgument(c_typelib_dir)
	arg_namespace_ := gi.NewStringArgument(c_namespace_)
	arg_version := gi.NewStringArgument(c_version)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_typelib_dir, arg_namespace_, arg_version, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_typelib_dir)
	gi.Free(c_namespace_)
	gi.Free(c_version)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// ignore GType struct RepositoryClass

// Enum RepositoryError
type RepositoryErrorEnum int

const (
	RepositoryErrorTypelibNotFound          RepositoryErrorEnum = 0
	RepositoryErrorNamespaceMismatch        RepositoryErrorEnum = 1
	RepositoryErrorNamespaceVersionConflict RepositoryErrorEnum = 2
	RepositoryErrorLibraryNotFound          RepositoryErrorEnum = 3
)

func RepositoryErrorGetType() gi.GType {
	ret := _I.GetGType(9, "RepositoryError")
	return ret
}

// Flags RepositoryLoadFlags
type RepositoryLoadFlags int

const (
	RepositoryLoadFlagsIrepositoryLoadFlagLazy RepositoryLoadFlags = 1
)

func RepositoryLoadFlagsGetType() gi.GType {
	ret := _I.GetGType(10, "RepositoryLoadFlags")
	return ret
}

// Struct RepositoryPrivate
type RepositoryPrivate struct {
	P unsafe.Pointer
}

func RepositoryPrivateGetType() gi.GType {
	ret := _I.GetGType(11, "RepositoryPrivate")
	return ret
}

// Enum ScopeType
type ScopeTypeEnum int

const (
	ScopeTypeInvalid  ScopeTypeEnum = 0
	ScopeTypeCall     ScopeTypeEnum = 1
	ScopeTypeAsync    ScopeTypeEnum = 2
	ScopeTypeNotified ScopeTypeEnum = 3
)

func ScopeTypeGetType() gi.GType {
	ret := _I.GetGType(12, "ScopeType")
	return ret
}

// Enum Transfer
type TransferEnum int

const (
	TransferNothing    TransferEnum = 0
	TransferContainer  TransferEnum = 1
	TransferEverything TransferEnum = 2
)

func TransferGetType() gi.GType {
	ret := _I.GetGType(13, "Transfer")
	return ret
}

// Enum TypeTag
type TypeTagEnum int

const (
	TypeTagVoid      TypeTagEnum = 0
	TypeTagBoolean   TypeTagEnum = 1
	TypeTagInt8      TypeTagEnum = 2
	TypeTagUint8     TypeTagEnum = 3
	TypeTagInt16     TypeTagEnum = 4
	TypeTagUint16    TypeTagEnum = 5
	TypeTagInt32     TypeTagEnum = 6
	TypeTagUint32    TypeTagEnum = 7
	TypeTagInt64     TypeTagEnum = 8
	TypeTagUint64    TypeTagEnum = 9
	TypeTagFloat     TypeTagEnum = 10
	TypeTagDouble    TypeTagEnum = 11
	TypeTagGtype     TypeTagEnum = 12
	TypeTagUtf8      TypeTagEnum = 13
	TypeTagFilename  TypeTagEnum = 14
	TypeTagArray     TypeTagEnum = 15
	TypeTagInterface TypeTagEnum = 16
	TypeTagGlist     TypeTagEnum = 17
	TypeTagGslist    TypeTagEnum = 18
	TypeTagGhash     TypeTagEnum = 19
	TypeTagError     TypeTagEnum = 20
	TypeTagUnichar   TypeTagEnum = 21
)

func TypeTagGetType() gi.GType {
	ret := _I.GetGType(14, "TypeTag")
	return ret
}

// Struct Typelib
type Typelib struct {
	P unsafe.Pointer
}

func TypelibGetType() gi.GType {
	ret := _I.GetGType(15, "Typelib")
	return ret
}

// g_typelib_free
//
func (v Typelib) Free() {
	iv, err := _I.Get(33, "Typelib", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_typelib_get_namespace
//
// [ result ] trans: nothing
//
func (v Typelib) GetNamespace() (result string) {
	iv, err := _I.Get(34, "Typelib", "get_namespace")
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

// g_typelib_symbol
//
// [ symbol_name ] trans: nothing
//
// [ symbol ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Typelib) Symbol(symbol_name string, symbol unsafe.Pointer) (result bool) {
	iv, err := _I.Get(35, "Typelib", "symbol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol_name := gi.CString(symbol_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_symbol_name := gi.NewStringArgument(c_symbol_name)
	arg_symbol := gi.NewPointerArgument(symbol)
	args := []gi.Argument{arg_v, arg_symbol_name, arg_symbol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_symbol_name)
	result = ret.Bool()
	return
}

// Struct UnresolvedInfo
type UnresolvedInfo struct {
	P unsafe.Pointer
}

func UnresolvedInfoGetType() gi.GType {
	ret := _I.GetGType(16, "UnresolvedInfo")
	return ret
}

// Flags VFuncInfoFlags
type VFuncInfoFlags int

const (
	VFuncInfoFlagsMustChainUp     VFuncInfoFlags = 1
	VFuncInfoFlagsMustOverride    VFuncInfoFlags = 2
	VFuncInfoFlagsMustNotOverride VFuncInfoFlags = 4
	VFuncInfoFlagsThrows          VFuncInfoFlags = 8
)

func VFuncInfoFlagsGetType() gi.GType {
	ret := _I.GetGType(17, "VFuncInfoFlags")
	return ret
}

// g_arg_info_get_closure
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoGetClosure(info BaseInfo) (result int32) {
	iv, err := _I.Get(36, "arg_info_get_closure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_arg_info_get_destroy
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoGetDestroy(info BaseInfo) (result int32) {
	iv, err := _I.Get(37, "arg_info_get_destroy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_arg_info_get_direction
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoGetDirection(info BaseInfo) (result DirectionEnum) {
	iv, err := _I.Get(38, "arg_info_get_direction", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DirectionEnum(ret.Int())
	return
}

// g_arg_info_get_ownership_transfer
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoGetOwnershipTransfer(info BaseInfo) (result TransferEnum) {
	iv, err := _I.Get(39, "arg_info_get_ownership_transfer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TransferEnum(ret.Int())
	return
}

// g_arg_info_get_scope
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoGetScope(info BaseInfo) (result ScopeTypeEnum) {
	iv, err := _I.Get(40, "arg_info_get_scope", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ScopeTypeEnum(ret.Int())
	return
}

// g_arg_info_get_type
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func ArgInfoGetType(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(41, "arg_info_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_arg_info_is_caller_allocates
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoIsCallerAllocates(info BaseInfo) (result bool) {
	iv, err := _I.Get(42, "arg_info_is_caller_allocates", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_arg_info_is_optional
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoIsOptional(info BaseInfo) (result bool) {
	iv, err := _I.Get(43, "arg_info_is_optional", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_arg_info_is_return_value
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoIsReturnValue(info BaseInfo) (result bool) {
	iv, err := _I.Get(44, "arg_info_is_return_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_arg_info_is_skip
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoIsSkip(info BaseInfo) (result bool) {
	iv, err := _I.Get(45, "arg_info_is_skip", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_arg_info_load_type
//
// [ info ] trans: nothing
//
// [ type1 ] trans: nothing, dir: out
//
func ArgInfoLoadType(info BaseInfo, type1 BaseInfo) {
	iv, err := _I.Get(46, "arg_info_load_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_info, arg_type1}
	iv.Call(args, nil, nil)
}

// g_arg_info_may_be_null
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ArgInfoMayBeNull(info BaseInfo) (result bool) {
	iv, err := _I.Get(47, "arg_info_may_be_null", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_callable_info_can_throw_gerror
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoCanThrowGerror(info BaseInfo) (result bool) {
	iv, err := _I.Get(48, "callable_info_can_throw_gerror", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_callable_info_get_arg
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func CallableInfoGetArg(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(49, "callable_info_get_arg", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_callable_info_get_caller_owns
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoGetCallerOwns(info BaseInfo) (result TransferEnum) {
	iv, err := _I.Get(50, "callable_info_get_caller_owns", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TransferEnum(ret.Int())
	return
}

// g_callable_info_get_instance_ownership_transfer
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoGetInstanceOwnershipTransfer(info BaseInfo) (result TransferEnum) {
	iv, err := _I.Get(51, "callable_info_get_instance_ownership_transfer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TransferEnum(ret.Int())
	return
}

// g_callable_info_get_n_args
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoGetNArgs(info BaseInfo) (result int32) {
	iv, err := _I.Get(52, "callable_info_get_n_args", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_callable_info_get_return_attribute
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoGetReturnAttribute(info BaseInfo, name string) (result string) {
	iv, err := _I.Get(53, "callable_info_get_return_attribute", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.String().Copy()
	return
}

// g_callable_info_get_return_type
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func CallableInfoGetReturnType(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(54, "callable_info_get_return_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_callable_info_invoke
//
// [ info ] trans: nothing
//
// [ function ] trans: nothing
//
// [ in_args ] trans: nothing
//
// [ n_in_args ] trans: nothing
//
// [ out_args ] trans: nothing
//
// [ n_out_args ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ is_method ] trans: nothing
//
// [ throws ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoInvoke(info BaseInfo, function unsafe.Pointer, in_args unsafe.Pointer, n_in_args int32, out_args unsafe.Pointer, n_out_args int32, return_value Argument, is_method bool, throws bool) (result bool, err error) {
	iv, err := _I.Get(55, "callable_info_invoke", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_info := gi.NewPointerArgument(info.P)
	arg_function := gi.NewPointerArgument(function)
	arg_in_args := gi.NewPointerArgument(in_args)
	arg_n_in_args := gi.NewInt32Argument(n_in_args)
	arg_out_args := gi.NewPointerArgument(out_args)
	arg_n_out_args := gi.NewInt32Argument(n_out_args)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_is_method := gi.NewBoolArgument(is_method)
	arg_throws := gi.NewBoolArgument(throws)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_info, arg_function, arg_in_args, arg_n_in_args, arg_out_args, arg_n_out_args, arg_return_value, arg_is_method, arg_throws, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_callable_info_is_method
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoIsMethod(info BaseInfo) (result bool) {
	iv, err := _I.Get(56, "callable_info_is_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_callable_info_iterate_return_attributes
//
// [ info ] trans: nothing
//
// [ iterator ] trans: everything, dir: inout
//
// [ name ] trans: nothing, dir: out
//
// [ value ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func CallableInfoIterateReturnAttributes(info BaseInfo, iterator int /*TODO:TYPE*/) (result bool, name string, value string) {
	iv, err := _I.Get(57, "callable_info_iterate_return_attributes", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_info, arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	name = outArgs[0].String().Copy()
	value = outArgs[1].String().Copy()
	result = ret.Bool()
	return
}

// g_callable_info_load_arg
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ arg ] trans: nothing, dir: out
//
func CallableInfoLoadArg(info BaseInfo, n int32, arg BaseInfo) {
	iv, err := _I.Get(58, "callable_info_load_arg", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	arg_arg := gi.NewPointerArgument(arg.P)
	args := []gi.Argument{arg_info, arg_n, arg_arg}
	iv.Call(args, nil, nil)
}

// g_callable_info_load_return_type
//
// [ info ] trans: nothing
//
// [ type1 ] trans: nothing, dir: out
//
func CallableInfoLoadReturnType(info BaseInfo, type1 BaseInfo) {
	iv, err := _I.Get(59, "callable_info_load_return_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_info, arg_type1}
	iv.Call(args, nil, nil)
}

// g_callable_info_may_return_null
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoMayReturnNull(info BaseInfo) (result bool) {
	iv, err := _I.Get(60, "callable_info_may_return_null", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_callable_info_skip_return
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func CallableInfoSkipReturn(info BaseInfo) (result bool) {
	iv, err := _I.Get(61, "callable_info_skip_return", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_constant_info_get_type
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func ConstantInfoGetType(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(62, "constant_info_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_enum_info_get_error_domain
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumInfoGetErrorDomain(info BaseInfo) (result string) {
	iv, err := _I.Get(63, "enum_info_get_error_domain", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_enum_info_get_method
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func EnumInfoGetMethod(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(64, "enum_info_get_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_enum_info_get_n_methods
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumInfoGetNMethods(info BaseInfo) (result int32) {
	iv, err := _I.Get(65, "enum_info_get_n_methods", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_enum_info_get_n_values
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumInfoGetNValues(info BaseInfo) (result int32) {
	iv, err := _I.Get(66, "enum_info_get_n_values", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_enum_info_get_storage_type
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumInfoGetStorageType(info BaseInfo) (result TypeTagEnum) {
	iv, err := _I.Get(67, "enum_info_get_storage_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TypeTagEnum(ret.Int())
	return
}

// g_enum_info_get_value
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func EnumInfoGetValue(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(68, "enum_info_get_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_field_info_get_flags
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func FieldInfoGetFlags(info BaseInfo) (result FieldInfoFlags) {
	iv, err := _I.Get(69, "field_info_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FieldInfoFlags(ret.Int())
	return
}

// g_field_info_get_offset
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func FieldInfoGetOffset(info BaseInfo) (result int32) {
	iv, err := _I.Get(70, "field_info_get_offset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_field_info_get_size
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func FieldInfoGetSize(info BaseInfo) (result int32) {
	iv, err := _I.Get(71, "field_info_get_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_field_info_get_type
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func FieldInfoGetType(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(72, "field_info_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_function_info_get_flags
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func FunctionInfoGetFlags(info BaseInfo) (result FunctionInfoFlags) {
	iv, err := _I.Get(73, "function_info_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FunctionInfoFlags(ret.Int())
	return
}

// g_function_info_get_property
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func FunctionInfoGetProperty(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(74, "function_info_get_property", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_function_info_get_symbol
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func FunctionInfoGetSymbol(info BaseInfo) (result string) {
	iv, err := _I.Get(75, "function_info_get_symbol", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_function_info_get_vfunc
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func FunctionInfoGetVfunc(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(76, "function_info_get_vfunc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_info_new
//
// [ type1 ] trans: nothing
//
// [ container ] trans: nothing
//
// [ typelib ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ result ] trans: everything
//
func InfoNew(type1 InfoTypeEnum, container BaseInfo, typelib Typelib, offset uint32) (result BaseInfo) {
	iv, err := _I.Get(77, "info_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_container := gi.NewPointerArgument(container.P)
	arg_typelib := gi.NewPointerArgument(typelib.P)
	arg_offset := gi.NewUint32Argument(offset)
	args := []gi.Argument{arg_type1, arg_container, arg_typelib, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_info_type_to_string
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func InfoTypeToString(type1 InfoTypeEnum) (result string) {
	iv, err := _I.Get(78, "info_type_to_string", "")
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

// g_interface_info_find_method
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoFindMethod(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(79, "interface_info_find_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_interface_info_find_signal
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoFindSignal(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(80, "interface_info_find_signal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_interface_info_find_vfunc
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoFindVfunc(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(81, "interface_info_find_vfunc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_constant
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetConstant(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(82, "interface_info_get_constant", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_iface_struct
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetIfaceStruct(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(83, "interface_info_get_iface_struct", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_method
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetMethod(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(84, "interface_info_get_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_n_constants
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func InterfaceInfoGetNConstants(info BaseInfo) (result int32) {
	iv, err := _I.Get(85, "interface_info_get_n_constants", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_interface_info_get_n_methods
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func InterfaceInfoGetNMethods(info BaseInfo) (result int32) {
	iv, err := _I.Get(86, "interface_info_get_n_methods", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_interface_info_get_n_prerequisites
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func InterfaceInfoGetNPrerequisites(info BaseInfo) (result int32) {
	iv, err := _I.Get(87, "interface_info_get_n_prerequisites", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_interface_info_get_n_properties
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func InterfaceInfoGetNProperties(info BaseInfo) (result int32) {
	iv, err := _I.Get(88, "interface_info_get_n_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_interface_info_get_n_signals
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func InterfaceInfoGetNSignals(info BaseInfo) (result int32) {
	iv, err := _I.Get(89, "interface_info_get_n_signals", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_interface_info_get_n_vfuncs
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func InterfaceInfoGetNVfuncs(info BaseInfo) (result int32) {
	iv, err := _I.Get(90, "interface_info_get_n_vfuncs", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_interface_info_get_prerequisite
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetPrerequisite(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(91, "interface_info_get_prerequisite", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_property
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetProperty(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(92, "interface_info_get_property", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_signal
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetSignal(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(93, "interface_info_get_signal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_interface_info_get_vfunc
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func InterfaceInfoGetVfunc(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(94, "interface_info_get_vfunc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_invoke_error_quark
//
// [ result ] trans: nothing
//
func InvokeErrorQuark() (result uint32) {
	iv, err := _I.Get(95, "invoke_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// Enum nvokeError
type nvokeErrorEnum int

const (
	nvokeErrorFailed           nvokeErrorEnum = 0
	nvokeErrorSymbolNotFound   nvokeErrorEnum = 1
	nvokeErrorArgumentMismatch nvokeErrorEnum = 2
)

func nvokeErrorGetType() gi.GType {
	ret := _I.GetGType(18, "nvokeError")
	return ret
}

// g_object_info_find_method
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoFindMethod(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(96, "object_info_find_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_object_info_find_method_using_interfaces
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ implementor ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func ObjectInfoFindMethodUsingInterfaces(info BaseInfo, name string) (result BaseInfo, implementor BaseInfo) {
	iv, err := _I.Get(97, "object_info_find_method_using_interfaces", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_implementor := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_info, arg_name, arg_implementor}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	implementor.P = outArgs[0].Pointer()
	result.P = ret.Pointer()
	return
}

// g_object_info_find_signal
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoFindSignal(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(98, "object_info_find_signal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_object_info_find_vfunc
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoFindVfunc(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(99, "object_info_find_vfunc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_object_info_find_vfunc_using_interfaces
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ implementor ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func ObjectInfoFindVfuncUsingInterfaces(info BaseInfo, name string) (result BaseInfo, implementor BaseInfo) {
	iv, err := _I.Get(100, "object_info_find_vfunc_using_interfaces", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_implementor := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_info, arg_name, arg_implementor}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	implementor.P = outArgs[0].Pointer()
	result.P = ret.Pointer()
	return
}

// g_object_info_get_abstract
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetAbstract(info BaseInfo) (result bool) {
	iv, err := _I.Get(101, "object_info_get_abstract", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_object_info_get_class_struct
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetClassStruct(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(102, "object_info_get_class_struct", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_constant
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetConstant(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(103, "object_info_get_constant", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_field
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetField(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(104, "object_info_get_field", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_fundamental
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetFundamental(info BaseInfo) (result bool) {
	iv, err := _I.Get(105, "object_info_get_fundamental", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_object_info_get_get_value_function
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetGetValueFunction(info BaseInfo) (result string) {
	iv, err := _I.Get(106, "object_info_get_get_value_function", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_object_info_get_interface
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetInterface(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(107, "object_info_get_interface", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_method
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetMethod(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(108, "object_info_get_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_n_constants
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNConstants(info BaseInfo) (result int32) {
	iv, err := _I.Get(109, "object_info_get_n_constants", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_n_fields
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNFields(info BaseInfo) (result int32) {
	iv, err := _I.Get(110, "object_info_get_n_fields", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_n_interfaces
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNInterfaces(info BaseInfo) (result int32) {
	iv, err := _I.Get(111, "object_info_get_n_interfaces", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_n_methods
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNMethods(info BaseInfo) (result int32) {
	iv, err := _I.Get(112, "object_info_get_n_methods", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_n_properties
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNProperties(info BaseInfo) (result int32) {
	iv, err := _I.Get(113, "object_info_get_n_properties", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_n_signals
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNSignals(info BaseInfo) (result int32) {
	iv, err := _I.Get(114, "object_info_get_n_signals", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_n_vfuncs
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetNVfuncs(info BaseInfo) (result int32) {
	iv, err := _I.Get(115, "object_info_get_n_vfuncs", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_object_info_get_parent
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetParent(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(116, "object_info_get_parent", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_property
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetProperty(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(117, "object_info_get_property", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_ref_function
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetRefFunction(info BaseInfo) (result string) {
	iv, err := _I.Get(118, "object_info_get_ref_function", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_object_info_get_set_value_function
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetSetValueFunction(info BaseInfo) (result string) {
	iv, err := _I.Get(119, "object_info_get_set_value_function", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_object_info_get_signal
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetSignal(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(120, "object_info_get_signal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_info_get_type_init
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetTypeInit(info BaseInfo) (result string) {
	iv, err := _I.Get(121, "object_info_get_type_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_object_info_get_type_name
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetTypeName(info BaseInfo) (result string) {
	iv, err := _I.Get(122, "object_info_get_type_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_object_info_get_unref_function
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInfoGetUnrefFunction(info BaseInfo) (result string) {
	iv, err := _I.Get(123, "object_info_get_unref_function", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_object_info_get_vfunc
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func ObjectInfoGetVfunc(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(124, "object_info_get_vfunc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_property_info_get_flags
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func PropertyInfoGetFlags(info BaseInfo) (result g.ParamFlags) {
	iv, err := _I.Get(125, "property_info_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = g.ParamFlags(ret.Int())
	return
}

// g_property_info_get_ownership_transfer
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func PropertyInfoGetOwnershipTransfer(info BaseInfo) (result TransferEnum) {
	iv, err := _I.Get(126, "property_info_get_ownership_transfer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TransferEnum(ret.Int())
	return
}

// g_property_info_get_type
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func PropertyInfoGetType(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(127, "property_info_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_registered_type_info_get_g_type
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func RegisteredTypeInfoGetGType(info BaseInfo) (result gi.GType) {
	iv, err := _I.Get(128, "registered_type_info_get_g_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_registered_type_info_get_type_init
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func RegisteredTypeInfoGetTypeInit(info BaseInfo) (result string) {
	iv, err := _I.Get(129, "registered_type_info_get_type_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_registered_type_info_get_type_name
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func RegisteredTypeInfoGetTypeName(info BaseInfo) (result string) {
	iv, err := _I.Get(130, "registered_type_info_get_type_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_signal_info_get_class_closure
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func SignalInfoGetClassClosure(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(131, "signal_info_get_class_closure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_signal_info_get_flags
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalInfoGetFlags(info BaseInfo) (result g.SignalFlags) {
	iv, err := _I.Get(132, "signal_info_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = g.SignalFlags(ret.Int())
	return
}

// g_signal_info_true_stops_emit
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalInfoTrueStopsEmit(info BaseInfo) (result bool) {
	iv, err := _I.Get(133, "signal_info_true_stops_emit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_struct_info_find_field
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func StructInfoFindField(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(134, "struct_info_find_field", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_struct_info_find_method
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func StructInfoFindMethod(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(135, "struct_info_find_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_struct_info_get_alignment
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func StructInfoGetAlignment(info BaseInfo) (result uint64) {
	iv, err := _I.Get(136, "struct_info_get_alignment", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_struct_info_get_field
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func StructInfoGetField(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(137, "struct_info_get_field", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_struct_info_get_method
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func StructInfoGetMethod(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(138, "struct_info_get_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_struct_info_get_n_fields
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func StructInfoGetNFields(info BaseInfo) (result int32) {
	iv, err := _I.Get(139, "struct_info_get_n_fields", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_struct_info_get_n_methods
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func StructInfoGetNMethods(info BaseInfo) (result int32) {
	iv, err := _I.Get(140, "struct_info_get_n_methods", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_struct_info_get_size
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func StructInfoGetSize(info BaseInfo) (result uint64) {
	iv, err := _I.Get(141, "struct_info_get_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_struct_info_is_foreign
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func StructInfoIsForeign(info BaseInfo) (result bool) {
	iv, err := _I.Get(142, "struct_info_is_foreign", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_struct_info_is_gtype_struct
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func StructInfoIsGtypeStruct(info BaseInfo) (result bool) {
	iv, err := _I.Get(143, "struct_info_is_gtype_struct", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_info_get_array_fixed_size
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInfoGetArrayFixedSize(info BaseInfo) (result int32) {
	iv, err := _I.Get(144, "type_info_get_array_fixed_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_type_info_get_array_length
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInfoGetArrayLength(info BaseInfo) (result int32) {
	iv, err := _I.Get(145, "type_info_get_array_length", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_type_info_get_array_type
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInfoGetArrayType(info BaseInfo) (result ArrayTypeEnum) {
	iv, err := _I.Get(146, "type_info_get_array_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ArrayTypeEnum(ret.Int())
	return
}

// g_type_info_get_interface
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func TypeInfoGetInterface(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(147, "type_info_get_interface", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_info_get_param_type
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func TypeInfoGetParamType(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(148, "type_info_get_param_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_info_get_tag
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInfoGetTag(info BaseInfo) (result TypeTagEnum) {
	iv, err := _I.Get(149, "type_info_get_tag", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TypeTagEnum(ret.Int())
	return
}

// g_type_info_is_pointer
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInfoIsPointer(info BaseInfo) (result bool) {
	iv, err := _I.Get(150, "type_info_is_pointer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_info_is_zero_terminated
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInfoIsZeroTerminated(info BaseInfo) (result bool) {
	iv, err := _I.Get(151, "type_info_is_zero_terminated", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_tag_to_string
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeTagToString(type1 TypeTagEnum) (result string) {
	iv, err := _I.Get(152, "type_tag_to_string", "")
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

// g_union_info_find_method
//
// [ info ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func UnionInfoFindMethod(info BaseInfo, name string) (result BaseInfo) {
	iv, err := _I.Get(153, "union_info_find_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_info, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_union_info_get_alignment
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func UnionInfoGetAlignment(info BaseInfo) (result uint64) {
	iv, err := _I.Get(154, "union_info_get_alignment", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_union_info_get_discriminator
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func UnionInfoGetDiscriminator(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(155, "union_info_get_discriminator", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_union_info_get_discriminator_offset
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func UnionInfoGetDiscriminatorOffset(info BaseInfo) (result int32) {
	iv, err := _I.Get(156, "union_info_get_discriminator_offset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_union_info_get_discriminator_type
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func UnionInfoGetDiscriminatorType(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(157, "union_info_get_discriminator_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_union_info_get_field
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func UnionInfoGetField(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(158, "union_info_get_field", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_union_info_get_method
//
// [ info ] trans: nothing
//
// [ n ] trans: nothing
//
// [ result ] trans: everything
//
func UnionInfoGetMethod(info BaseInfo, n int32) (result BaseInfo) {
	iv, err := _I.Get(159, "union_info_get_method", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_info, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_union_info_get_n_fields
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func UnionInfoGetNFields(info BaseInfo) (result int32) {
	iv, err := _I.Get(160, "union_info_get_n_fields", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_union_info_get_n_methods
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func UnionInfoGetNMethods(info BaseInfo) (result int32) {
	iv, err := _I.Get(161, "union_info_get_n_methods", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_union_info_get_size
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func UnionInfoGetSize(info BaseInfo) (result uint64) {
	iv, err := _I.Get(162, "union_info_get_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_union_info_is_discriminated
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func UnionInfoIsDiscriminated(info BaseInfo) (result bool) {
	iv, err := _I.Get(163, "union_info_is_discriminated", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_value_info_get_value
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueInfoGetValue(info BaseInfo) (result int64) {
	iv, err := _I.Get(164, "value_info_get_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_vfunc_info_get_address
//
// [ info ] trans: nothing
//
// [ implementor_gtype ] trans: nothing
//
// [ result ] trans: nothing
//
func VfuncInfoGetAddress(info BaseInfo, implementor_gtype gi.GType) (result unsafe.Pointer, err error) {
	iv, err := _I.Get(165, "vfunc_info_get_address", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_info := gi.NewPointerArgument(info.P)
	arg_implementor_gtype := gi.NewUintArgument(uint(implementor_gtype))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_info, arg_implementor_gtype, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Pointer()
	return
}

// g_vfunc_info_get_flags
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func VfuncInfoGetFlags(info BaseInfo) (result VFuncInfoFlags) {
	iv, err := _I.Get(166, "vfunc_info_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = VFuncInfoFlags(ret.Int())
	return
}

// g_vfunc_info_get_invoker
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func VfuncInfoGetInvoker(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(167, "vfunc_info_get_invoker", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_vfunc_info_get_offset
//
// [ info ] trans: nothing
//
// [ result ] trans: nothing
//
func VfuncInfoGetOffset(info BaseInfo) (result int32) {
	iv, err := _I.Get(168, "vfunc_info_get_offset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_vfunc_info_get_signal
//
// [ info ] trans: nothing
//
// [ result ] trans: everything
//
func VfuncInfoGetSignal(info BaseInfo) (result BaseInfo) {
	iv, err := _I.Get(169, "vfunc_info_get_signal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// constants
const ()
