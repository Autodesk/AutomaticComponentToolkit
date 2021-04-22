package main

import (
    "bytes"
    "strings"
    "text/template"
)

// WriteClientImpl writes client implementation binding code for the given component.
// This does nothing if there are no abstract classes in the component definition.
//
// The generated code goes into namespace <component>::Binding::ClientImpl. What
// you get is a set of abstract C++ classes, one for each abstract class in the
// component definition, plus a base class from which they all derive. The base
// class corresponds to the base class defined in the component.
//
// A instance of a client impl class is an object that can be 'wrapped' in an
// instance of the corresponding class in <component>::Binding, and passed
// into binding functions as parameters. The boilerplate for doing this is
// automatically generated, so from the user's perspective they are simply
// deriving from an ordinary C++ class.
//
// The way this works is as follows:
//
// In ACT, an object is represented by an 'extended handle': a handle, plus a
// pointer to a 'symbol lookup' function. These are passed across the ABI
// together, and the lookup function is used by the binding or implementation
// to obtain pointers to functions that call the object's methods.
//
// To 'derive' an implementation of such an object in client code, we need to
// provide those two things. The client impl classes provide all the boilerplate
// necessary to do this.
//
// Each class has a set of public virtual 'API' methods corresponding to the
// methods declared on the class in the XML. The user needs to override these
// to provide their implementation.   Each class also has a set of generated
// 'ABI' versions of those functions, static functions that cast a handle to
// a pointer to an instance and call the appropriate method on it.  Finally,
// each class provides a symbol lookup function that can resolve those ABI
// functions. An 'extended handle' comprising a pointer to an instance of
// one of these classes, plus the symbol lookup function for that class, can
// thus be used to construct an instance of the corresponding 'binding' class
// and passed across the ABI.
//
// Functions are provided to deal with wrapping and unwrapping of instances.
//
// Client impl classes derive from a common base class. This provides an
// implementation of reference counting and other methods needed by all
// classes.
func WriteClientImpl(component ComponentDefinition, w LanguageWriter) error {

    // Find abstract derived & base classes
    baseClasses := make([]ComponentDefinitionClass, 0)
    derivedClasses := make([]ComponentDefinitionClass, 0)
    for _, class := range component.Classes {
        if class.IsAbstract() {
            derivedClasses = append(derivedClasses, class)
        } else if component.isBaseClass(class) {
            baseClasses = append(baseClasses, class)
        }
    }

    // Don't write any client implementation code if there are no abstract classes
    // to derive from.
    if len(derivedClasses) == 0 {
        return nil
    }

    // Begin the ClientImpl namespace etc.
    if err := buildCppClientImplPreamble(component, w); err != nil {
        return err
    }

    // Declare base classes
    for _, class := range baseClasses {
        if err := buildCppClientImplClassDecl(component, class, w); err != nil {
            return err
        }
    }

    // Declare derived classes
    for _, class := range derivedClasses {
        if err := buildCppClientImplClassDecl(component, class, w); err != nil {
            return err
        }
    }

    // Implement base classes
    for _, class := range baseClasses {
        if err := buildCppClientImplClassImpl(component, class, w); err != nil {
            return err
        }
    }

    // Implement derived classes
    for _, class := range derivedClasses {
        if err := buildCppClientImplClassImpl(component, class, w); err != nil {
            return err
        }
    }

    // Finish - close the namespace etc.
    if err := buildCppClientImplEnd(component, w); err != nil {
        return err
    }

    return nil
}

// Note: structs can be passed directly into the template execution and
// the template can reference struct fields, even doing things like looping
// over slices.   I didn't realise this to begin with, hence the conversion
// to a map.  It might be possible to simplify this code by passing in the
// ComponentDefinition* structs directly.

// getComponentPropertyMap returns a map containing properties pertaining to
// a component, useful for writeSubstitution
func getComponentPropertyMap(component ComponentDefinition) map[string]interface{} {
    ret := map[string]interface{}{}
    ret["NameSpace"] = component.NameSpace
    ret["NameSpaceUpper"] = strings.ToUpper(component.NameSpace)
    ret["NameSpaceLower"] = strings.ToLower(component.NameSpace)
    return ret
}

// getClassPropertyMap returns a map containing properties pertaining to
// a method, useful for writeSubstitution. The result contains component
// properties too.
func getClassPropertyMap(component ComponentDefinition, class ComponentDefinitionClass) map[string]interface{} {
    ret := getComponentPropertyMap(component)
    ret["BaseClassName"] = class.ParentClass
    ret["IsBaseClass"] = component.isBaseClass(class)
    ret["ClassName"] = class.ClassName
    ret["ClassNameUpper"] = strings.ToUpper(class.ClassName)
    ret["ClassNameLower"] = strings.ToLower(class.ClassName)
    return ret
}

// getMethodPropertyMap returns a map containing properties pertaining to
// a class, useful for writeSubstitution. The result contains class and
// component properties too.
func getMethodPropertyMap(component ComponentDefinition, class ComponentDefinitionClass, method ComponentDefinitionMethod) map[string]interface{} {
    ret := getClassPropertyMap(component, class)
    ret["MethodName"] = method.MethodName
    ret["MethodNameUpper"] = strings.ToUpper(method.MethodName)
    ret["MethodNameLower"] = strings.ToLower(method.MethodName)
    return ret
}

// writeSubstitution writes the templateString to the LanguageWriter w, substituting in
// fields from the properties map. This is done via text/template.
func writeSubstitution(templateString string, properties map[string]interface{}, w LanguageWriter) error {
    tmpl, err := template.New("msg").Parse(templateString)
    if err != nil { 
        return err 
    }
    buf := &bytes.Buffer{}
    err = tmpl.Execute(buf, properties)
    if err != nil {
        return err
    }
    s := buf.String()
    w.Writeln(s)
    return nil
}

// buildCppClientImplClassDecl writes the declaration for a client impl class.
func buildCppClientImplClassDecl(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    err := error(nil)
    err = buildCppClientImplClassDeclPublic(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplAPIMethodDecls(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplClassDeclProtected(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplABIMethodDecls(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplClassDeclPrivate(component, class, w)
    if err != nil {
        return err
    }
    return nil;
}

// buildCppclientImplClassImpl will output the 'implementation' code for a client
// impl class.   This assumes the declaration code has already been written.
func buildCppClientImplClassImpl(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    err := error(nil)
    err = buildCppClientImplClassImplCode(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplSymbolLookupFunctionABI(component, class, false, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplAPIMethodImpls(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplABIMethodImpls(component, class, w)
    if (err != nil) {
        return err
    }
    return err
}

// buildCppClientImplPreamble writes the initial section of the client implementation,
// which starts the client impl namespace, defines necessary global functions and types
// etc.
func buildCppClientImplPreamble(component ComponentDefinition, w LanguageWriter) error {
    code := `

/*************************************************************************************************************************
Client implementation
**************************************************************************************************************************/

namespace {{.NameSpace}} {
namespace Binding {
namespace ClientImpl {

// Create a wrapped instance of a client implementation class. This means
// constructing a client implementation class instance, forwarding arguments
// passed to the function, which is then wrapped in a new binding instance of
// the appropriate type. The result can thus be passed into functions that deal
// in binding types.
template <typename tCLASS, typename... tARGS>                         
static typename tCLASS::tBINDING_PTR CreateWrappedInstance(tARGS&&... args)
{          
  auto ptr = std::make_unique<tCLASS>(std::forward<tARGS>(args)...);  
  ptr->AcquireInstance(); // CBase ctor doesn't acquire         
  return typename tCLASS::tBINDING_PTR(
    new typename tCLASS::tBINDING_CLASS(ptr.release()->GetExtendedHandle())
  );
}

// Given a pointer to a binding object, cast the wrapped handle to a client
// implementation instance. The caller is responsible for ensuring that the
// binding object really does wrap a client implementation.
template <typename tCLASS>
static tCLASS* UnsafeGetWrappedInstance(typename tCLASS::tBINDING_PTR pBindingPtr)
{
  return UnsafeGetWrappedInstance<tCLASS>(pBindingPtr->GetExtendedHandle())
}

// Cast a handle to a client implementation instance. The caller is responsible for
// ensuring that the binding object really does wrap a client implementation.
template <typename tCLASS>
static tCLASS* UnsafeGetWrappedInstance({{.NameSpace}}ExtendedHandle extendedHandle)
{
  return UnsafeGetWrappedInstance<tCLASS>(extendedHandle.m_hHandle);
}

// Cast a handle to a client implementation instance. The caller is responsible for
// ensuring that the binding object really does wrap a client implementation.
template <typename tCLASS>
static tCLASS* UnsafeGetWrappedInstance({{.NameSpace}}Handle handle)
{
  return (tCLASS*) handle;
}

// Static handler for E{{.NameSpace}}Exception from client impl abi wrapper function.
// TODO: propagate error code / info
template <typename tCLASS>
{{.NameSpace}}Result handle{{.NameSpace}}Exception(tCLASS* object, const E{{.NameSpace}}Exception& e)
{
  return e.getErrorCode();
}

// Static handler for std::exception from client impl abi wrapper function.
// TODO: propagate error code / info
template <typename tCLASS>
{{.NameSpace}}Result handleStdException(tCLASS* object, const std::exception& e)
{
  return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
}

// Static handler for generic exception from client impl abi wrapper function.
// TODO: propagate error code / info
template <typename tCLASS>
{{.NameSpace}}Result handleUnhandledException(tCLASS* object)
{
  return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
}

// Utility method for SymbolLookupFunction_ABI. Attempt to find the symbol in 
// the map, outputting it if present, and returning an appropriate error code if
// not.
inline static {{.NameSpace}}Result LookupSymbolInMap(
  const char* pProcName, 
  std::map<std::string, void*>& procAddressMap, 
  void** ppProcAddress
)
{
    try {
        if (pProcName == nullptr)
            return {{.NameSpaceUpper}}_ERROR_INVALIDPARAM;
        if (ppProcAddress == nullptr)
            return {{.NameSpaceUpper}}_ERROR_INVALIDPARAM;
        auto it = procAddressMap.find(pProcName);
        *ppProcAddress = it != end(procAddressMap) ? it->second : nullptr;
        if (!*ppProcAddress) {
            return {{.NameSpaceUpper}}_ERROR_COULDNOTFINDLIBRARYEXPORT;
        }
    } catch ({{.NameSpace}}::Binding::E{{.NameSpace}}Exception&) {
        return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
    } catch (std::exception&) {
        return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
    } catch (...) {
        return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
    }
    return {{.NameSpaceUpper}}_SUCCESS;
}
`
    return writeSubstitution(code, getComponentPropertyMap(component), w)
}

// buildCppClientImplEnd writes the end of the client implementation section,
// closing namespaces etc initiated by the 'preamble'.
func buildCppClientImplEnd(component ComponentDefinition, w LanguageWriter) error {
    code := `
} // namespace ClientImpl
} // namespace Binding
} // namespace {{.NameSpace}}
`
    return writeSubstitution(code, getComponentPropertyMap(component), w)
}

// buildCppClientImplClassDeclPublic writes the public section of a client 
// impl class declaration.
func buildCppClientImplClassDeclPublic(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
{{- /**/ -}}
/*************************************************************************************************************************
 Class C{{.ClassName}}
**************************************************************************************************************************/
class C{{.ClassName}} {{ if not .IsBaseClass }}: public C{{.BaseClassName}} {{ end }}{
public:

  // Associated types.  These are used by certain template functions.
  using tBINDING_PTR = {{.NameSpace}}::Binding::P{{.ClassName}};
  using tBINDING_CLASS = {{.NameSpace}}::Binding::C{{.ClassName}};
{{- if not .IsBaseClass }}
  using tBASE = {{.NameSpace}}::Binding::ClientImpl::C{{.BaseClassName}};
{{- end }}

  // Default constructor.
  inline C{{.ClassName}}();

  // Copying is prohibited
  C{{.ClassName}}(const C{{.ClassName}}& that) = delete;

  // Assignment is prohibited
  C{{.ClassName}}& operator=(const C{{.ClassName}}& that) = delete;

  // Destructor
  inline virtual ~C{{.ClassName}}() {{- if not .IsBaseClass }} override{{ end }};

{{- if .IsBaseClass }}

  // Get the extended handle for this instance. The extended handle can
  // be wrapped in a tBINDING_CLASS instance and passed into binding
  // methods as a parameter.
  inline {{.NameSpace}}ExtendedHandle GetExtendedHandle() const;
{{- end }}

{{- if not .IsBaseClass }}

  // Return the symbol lookup method for this class.    In a derived
  // class, this is overridden to return a function that exposes the
  // symbols of that class, calling down to the base class function
  // if a symbol is not found.
  inline virtual {{.NameSpace}}_pvoid GetSymbolLookupMethod() override;
{{- end }}

  // API methods
`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

// buildCppClientImplAPIMethodDecls writes a declaration for each method on a
// client impl class declaration.
func buildCppClientImplAPIMethodDecls(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    for _, method := range class.Methods {
        returnType, parameters, err := buildDynamicCPPMethodDeclaration(method, component.NameSpace, "", "C"+class.ClassName, true)
        if err != nil {
            return err
        }
        w.Writeln("  inline virtual %s %s(%s);", returnType, method.MethodName, parameters)
        w.Writeln("")
    }
    return nil
}

// buildCppClientImplClassDeclProtected writes the 'protected' section for a client
// impl class declation.
func buildCppClientImplClassDeclProtected(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
{{- /**/ -}}
protected:
    
  // Symbol lookup function for the methods on this class. Looks up the _ABI
  // functions of this class by name.    Derived classes must add their own
  // symbol lookup function exposing their own functions, which should call
  // down to this function when a symbol cannot be found.
  inline static {{.NameSpace}}Result SymbolLookupFunction_ABI(
    const char* pProcName, 
    void** ppProcAddress
  );`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

// buildCppClientImplABIMethodDecls outputs a declaration for each 'ABI' method for a client
// impl class.
func buildCppClientImplABIMethodDecls(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    for _, method := range class.Methods {
        sComments, _, sParameters, err := WriteCCPPAbiMethod(method, component.NameSpace, class.ClassName, false, false, true)
        if (err != nil) {
            return err
        }
        w.Writeln("  ")
        w.Writelns("  ", sComments)
        w.Writeln("  inline static %sResult %s_ABI(%s);", component.NameSpace, method.MethodName, sParameters)
    }
    return nil
}

// buildCppClientImplClassDeclPrivate generates the private section of a client impl
// class declaration.
func buildCppClientImplClassDeclPrivate(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
private:
{{- if .IsBaseClass }}
  // Reference count
  {{.NameSpace}}_uint64 m_refcount;
{{- end }}
};
`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

// buildCppClientImplClassImplCode will output the code for the 'special' methods
// on a ClientImpl class, including the constructor and destructor.
func buildCppClientImplClassImplCode(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
/*************************************************************************************************************************
C{{.ClassName}} Implementation
**************************************************************************************************************************/

inline C{{.ClassName}}::C{{.ClassName}}()
{{ if .IsBaseClass -}}
   : m_refcount(0)
{{ end -}}
{
}

inline C{{.ClassName}}::~C{{.ClassName}}()
{
}

inline {{.NameSpace}}_pvoid C{{.ClassName}}::GetSymbolLookupMethod()
{
  return ({{.NameSpace}}_pvoid) &GetSymbolLookupMethod_ABI;
}

{{- if .IsBaseClass }}

inline bool C{{.ClassName}}::GetLastError(std::string & sErrorMessage)
{
  return false;
}

inline void C{{.ClassName}}::ReleaseInstance()
{
  --m_refcount;
  if (m_refcount == 0) {
    delete this;
  }
}

inline void C{{.ClassName}}::AcquireInstance()
{
  ++m_refcount;
}

inline void C{{.ClassName}}::GetVersion(
  {{.NameSpace}}_uint32 & nMajor, 
  {{.NameSpace}}_uint32 & nMinor,
  {{.NameSpace}}_uint32 & nMicro
)
{
  nMajor = {{.NameSpaceUpper}}_VERSION_MAJOR;
  nMinor = {{.NameSpaceUpper}}_VERSION_MINOR;
  nMicro = {{.NameSpaceUpper}}_VERSION_MICRO;
}
{{- end -}}
`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

// buildCppClientImplSymbolLookupFunctionABI generates a static function implementation
// exposing the ABI functions for a class. The generated function returns a 'symbol
// lookup' function pointer which can go in an Extended Handle alongside a pointer to
// an instance of the corresponding class. On a derived class, the generated function
// will call the base class version of the function if a symbol cannot be found. Mappings
// are cached in a static std::map.
func buildCppClientImplSymbolLookupFunctionABI(component ComponentDefinition, class ComponentDefinitionClass, isBaseClass bool, w LanguageWriter) error {
    err := error(nil)
    properties := getClassPropertyMap(component, class)
    beginCode := `
inline {{.NameSpace}}Result C{{.ClassName}}::SymbolLookupFunction_ABI(
  const char* pProcName, 
  void** ppProcAddress
)
{
  static std::map<std::string, void*> sProcAddressMap;
  if (sProcAddressMap.empty()) {`
    err = writeSubstitution(beginCode, properties, w)
    if err != nil {
        return err
    }

    methodCode := `    sProcAddressMap["{{.NameSpaceLower}}_{{.ClassNameLower}}_{{.MethodNameLower}}"] = (void*) &{{.MethodName}}_ABI;`
    for _, method := range class.Methods {
        methodProperties := getMethodPropertyMap(component, class, method)
        err = writeSubstitution(methodCode, methodProperties, w)
        if err != nil {
            return err
        }
    }

    endCode := `  }
{{- if .IsBaseClass }}
  return LookupSymbolInMap(pProcName, sProcAddressMap, ppProcAddress);
{{- else }}
  {{.NameSpace}}Result ret = LookupSymbolInMap(pProcName, sProcAddressMap, ppProcAddress);
  if (ret == {{.NameSpaceUpper}}_ERROR_COULDNOTFINDLIBRARYEXPORT) {
      ret = tBASE::SymbolLookupFunction_ABI(pProcName, ppProcAddress);
  }
  return ret;
{{- end }}
}`
    err = writeSubstitution(endCode, properties, w)
    if err != nil {
        return err
    }

    return err
}

// isSpecialCppClientImplMethod returns true if a method is one of the
// 'special' methods required for the client implementation class scheme
// to work, which have special implementations generated (and so do not
// have generic 'not implemented' dummy implementations generated for
// them.)
func isSpecialCppClientImplMethod(method ComponentDefinitionMethod) bool {
    if method.MethodName == "AcquireInstance" {
        return true
    }
    if method.MethodName == "ReleaseInstance" {
        return true
    }
    if method.MethodName == "GetSymbolLookupMethod" {
        return true
    }
    if method.MethodName == "GetLastError" {
        return true
    }
    if method.MethodName == "GetVersion" {
        return true
    }
    return false
}

// buildCppClientImplAPIMethodImpls writes an implementation for each method on the class
// which is not a 'special' method - which will have been implemented elsewhere. The generated
// implementation simply throws a 'not implemented' error. The user of the library must
// override these methods when deriving from the class.
func buildCppClientImplAPIMethodImpls(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    beginComment := `
/*************************************************************************************************************************
 C{{.ClassName}} API-methods
**************************************************************************************************************************/
`
    if err := writeSubstitution(beginComment, getClassPropertyMap(component, class), w); err != nil {
        return err
    }

    // This is a dummy implementation to be overridden by the user.
    methodCode := `
inline {{.ReturnType}} C{{.ClassName}}::{{.MethodName}}({{.Parameters}})
{
  throw E{{.NameSpace}}Exception({{.NameSpaceUpper}}_ERROR_NOTIMPLEMENTED, "");
}
`

    for _, method := range class.Methods {

        // Some methods don't need to be implemented by the caller; we generate implementations
        // for them ourselves, so skip them.
        if isSpecialCppClientImplMethod(method) {
            continue
        }

        // Otherwise output a dummy implementation
        properties := getMethodPropertyMap(component, class, method)
        classIdentifier := ""
        returnType, parameters, err := buildDynamicCPPMethodDeclaration(method, component.NameSpace, classIdentifier, class.ClassName, true)
        if err != nil {
            return err
        }
        properties["ReturnType"] = returnType
        properties["Parameters"] = parameters
        if err := writeSubstitution(methodCode, properties, w); err != nil {
            return err
        }
    }
    return nil
}

// buildCppClientImplABIMethodImpls generates an 'ABI' wrapper method implementation
// for each method on the class. The wrapper calls the corresponding method on the
// clientimpl class instance.
func buildCppClientImplABIMethodImpls(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    beginComment := `
/*************************************************************************************************************************
 C{{.ClassName}} ABI-methods
**************************************************************************************************************************/
`
    if err := writeSubstitution(beginComment, getClassPropertyMap(component, class), w); err != nil {
        return err
    }

    for _, method := range class.Methods {
        baseName := ""
        classIdentifier := "C"
        isGlobal := false
        doJournal := false
        isSpecialFunction := eSpecialMethodNone
        isClientImpl := true
        if err := writeCImplementationMethod(component, method, w, baseName, component.NameSpace, classIdentifier, class.ClassName, component.Global.BaseClassName, isGlobal, doJournal, isSpecialFunction, isClientImpl); err != nil {
            return err
        }
    }

    return nil
}