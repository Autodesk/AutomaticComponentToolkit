package main

import (
    "bytes"
    "strings"
    "text/template"
)

func buildClientImplementationSection(component ComponentDefinition, w LanguageWriter) error {

    // Begin the ClientImpl namespace etc.
    if err := buildCppClientImplementationPreamble(component, w); err != nil {
        return err
    }

    // Find abstract derived & base classes
    baseClasses := make([]ComponentDefinitionClass, 0)
    derivedClasses := make([]ComponentDefinitionClass, 0)
    for _, class := range component.Classes {
        if class.IsAbstract() {
            isBase := class.ParentClass == class.ClassName
            if isBase {
                baseClasses = append(baseClasses, class)
            } else {
                derivedClasses = append(derivedClasses, class)
            }
        }
    }

    // Declare base classes
    for _, class := range baseClasses {
        if err := buildCppClientImplementationBaseClassDeclaration(component, class, w); err != nil {
            return err
        }
    }

    // Declare derived classes
    for _, class := range derivedClasses {
        if err := buildCppClientImplementationDerivedClassDeclaration(component, class, w); err != nil {
            return err
        }
    }

    // Implement base classes
    for _, class := range baseClasses {
        if err := buildCppClientImplementationBaseClassImplementation(component, class, w); err != nil {
            return err
        }
    }

    // Implement derived classes
    for _, class := range derivedClasses {
        if err := buildCppClientImplementationDerivedClassImplementation(component, class, w); err != nil {
            return err
        }
    }

    // Finish - close the namespace etc.
    if err := buildCppClientImplementationEnd(component, w); err != nil {
        return err
    }

    return nil
}

func getComponentPropertyMap(component ComponentDefinition) map[string]interface{} {
    ret := map[string]interface{}{}
    ret["NameSpace"] = component.NameSpace
    ret["NameSpaceUpper"] = strings.ToUpper(component.NameSpace)
    ret["NameSpaceLower"] = strings.ToLower(component.NameSpace)
    return ret
}

func getClassPropertyMap(component ComponentDefinition, class ComponentDefinitionClass) map[string]interface{} {
    ret := getComponentPropertyMap(component)
    ret["ClassName"] = class.ClassName
    ret["ClassNameUpper"] = strings.ToUpper(class.ClassName)
    ret["ClassNameLower"] = strings.ToLower(class.ClassName)
    return ret
}

func getMethodPropertyMap(component ComponentDefinition, class ComponentDefinitionClass, method ComponentDefinitionMethod) map[string]interface{} {
    ret := getClassPropertyMap(component, class)
    ret["MethodName"] = method.MethodName
    ret["MethodNameUpper"] = strings.ToUpper(method.MethodName)
    ret["MethodNameLower"] = strings.ToLower(method.MethodName)
    return ret
}

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

func buildCppClientImplementationPreamble(component ComponentDefinition, w LanguageWriter) error {
    code := `

namespace {{.NameSpace}} {
namespace Binding {
namespace ClientImpl {

// Create a wrapped instance of a client implementation class. This means
// constructing a client implementation class instance, forwarding arguments
// passed to the function, which is then wrapped in a new binding instance of
// the appropriate type. The result can thus be passed into functions that deal
// in binding types.
template <typename tCLASS, typename... tARGS>                         
static tCLASS::tBINDING_PTR CreateWrappedInstance(tARGS&&... args)
{          
  auto ptr = std::make_unique<tCLASS>(std::forward<tARGS>(args)...);  
  ptr->AcquireInstance(); // CBase ctor doesn't acquire         
  return tBINDING_PTR(
    new tCLASS::tBINDING_CLASS(ptr.release()->GetExtendedHandle())
  );
}

// Given a pointer to a binding object, cast the wrapped handle to a client
// implementation instance. The caller is responsible for ensuring that the
// binding object really does wrap a client implementation.
template <typename tCLASS>
static tCLASS* UnsafeGetWrappedInstance(tCLASS::tBINDING_PTR pBindingPtr)
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
static tCLASS* UnsafeGetWrappedInstance<tCLASS>({{.NameSpace}}Handle handle)
{
    return (tCLASS*) handle;
}
`
    return writeSubstitution(code, getComponentPropertyMap(component), w)
}

func buildCppClientImplementationEnd(component ComponentDefinition, w LanguageWriter) error {
    code := `
} // namespace ClientImpl
} // namespace Binding
} // namespace {{.NameSpace}}
`
    return writeSubstitution(code, getComponentPropertyMap(component), w)
}

func buildCppClientImplementationBaseClassDeclaration(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    err := error(nil)
    err = buildCppClientImplementationBaseClassDeclarationBegin(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationClassDeclarationAPIMethods(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationBaseClassDeclarationMiddle(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationClassDeclarationABIMethods(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationClassDeclarationEnd(component, class, w)
    if err != nil {
        return err
    }
    return nil;
}

func buildCppClientImplementationDerivedClassDeclaration(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    err := error(nil)
    err = buildCppClientImplementationDerivedClassDeclarationBegin(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationClassDeclarationAPIMethods(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationDerivedClassDeclarationMiddle(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationClassDeclarationABIMethods(component, class, w)
    if err != nil {
        return err
    }
    err = buildCppClientImplementationClassDeclarationEnd(component, class, w)
    if err != nil {
        return err
    }
    return nil;
}

func buildCppClientImplementationBaseClassDeclarationBegin(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `

class C{{.ClassName}} {
public:

    using tBINDING_PTR = {{.NameSpace}}::Binding::P{{.ClassName}};
    using tBINDING_CLASS = {{.NameSpace}}::Binding::C{{.ClassName}};

    C{{.ClassName}}();
    C{{.ClassName}}(const C{{.ClassName}}& that) = delete;
    C{{.ClassName}}& operator=(const C{{.ClassName}}& that) = delete;
    virtual C{{.ClassName}}();

    {{.NameSpace}}ExtendedHandle GetExtendedHandle() const;

    virtual {{.NameSpace}}_pvoid GetSymbolLookupMethod() const;

    // API methods

`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

func buildCppClientImplementationDerivedClassDeclarationBegin(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `

class C{{.ClassName}} : public C{{.BaseClassName}} {
public:

    using tBINDING_PTR = {{.NameSpace}}::Binding::P{{.ClassName}};
    using tBINDING_CLASS = {{.NameSpace}}::Binding::C{{.ClassName}};
    using tBASE = {{.NameSpace}}::Binding::C{{.BaseClassName}};

    C{{.ClassName}}();
    C{{.ClassName}}(const C{{.ClassName}}& that) = delete;
    C{{.ClassName}}& operator=(const C{{.ClassName}}& that) = delete;
    virtual C{{.ClassName}}() override;

    virtual {{.NameSpace}}_pvoid GetSymbolLookupMethod() const override;

    // API methods

`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

func buildCppClientImplementationClassDeclarationAPIMethods(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    for j := 0; j < len(class.Methods); j++ {
        method := class.Methods[j]
        returnType, parameters, err := buildDynamicCPPMethodDeclaration(method, component.NameSpace, "", "C"+class.ClassName)
        if err != nil {
            return err
        }
        w.Writeln("  virtual %s %s(%s);", returnType, method.MethodName, parameters)
        w.Writeln("")
    }
    return nil
}

func buildCppClientImplementationBaseClassDeclarationMiddle(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
    protected:
    
      // NOTE: helper for derived classes.
      static {{.NameSpace}}Result LookupSymbolInMap(
          const char* pProcName, 
          std::map<std::string, void*>& procAddressMap, 
          void** ppProcAddress
      );
    
      // Lookup method for this class.
      // NOTE: derived class must wrap this and call it
      static {{.NameSpace}}Result SymbolLookupFunction_ABI(
          const char* pProcName, 
          void** ppProcAddress
      );
    
    private:
      {{.NameSpace}}_uint64 m_refcount;`
      return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

func buildCppClientImplementationDerivedClassDeclarationMiddle(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
    protected:
    
      // Lookup method for this class.
      static {{.NameSpace}}Result SymbolLookupFunction_ABI(
          const char* pProcName, 
          void** ppProcAddress
      );
    
    private:
`
      return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

func buildCppClientImplementationClassDeclarationABIMethods(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    for j := 0; j < len(class.Methods); j++ {
        method := class.Methods[j]
        sComments, _, sParameters, err := WriteCCPPAbiMethod(method, component.NameSpace, class.ClassName, false, false, true)
        if (err != nil) {
            return err
        }
        w.Writeln("  ")
        w.Writelns("  ", sComments)
        w.Writeln("  static %sResult %s_ABI(%s);", component.NameSpace, method.MethodName, sParameters)
    }
    return nil
}

func buildCppClientImplementationClassDeclarationEnd(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `};
    `
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}


func buildCppClientImplementationBaseClassImplementation(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    err := error(nil)
    err = buildCppClientImplementationBaseClassImplementationBegin(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplementationClassImplementationAPIMethods(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppclientImplementationClassImplementationABIMethods(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplementationSymbolLookupFunctionABI(component, class, true, w)
    if (err != nil) {
        return err
    }
    return err
}

func buildCppClientImplementationDerivedClassImplementation(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    err := error(nil)
    err = buildCppClientImplementationDerivedClassImplementationBegin(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplementationClassImplementationAPIMethods(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppclientImplementationClassImplementationABIMethods(component, class, w)
    if (err != nil) {
        return err
    }
    err = buildCppClientImplementationSymbolLookupFunctionABI(component, class, false, w)
    if (err != nil) {
        return err
    }
    return err
}

func buildCppClientImplementationDerivedClassImplementationBegin(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `
`
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

func buildCppClientImplementationBaseClassImplementationBegin(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    code := `

C{{.ClassName}}::C{{.ClassName}}()
   : m_refcount(0)
{
}

C{{.ClassName}}::~C{{.ClassName}}()
{
}

{{.NameSpace}}Result C{{.ClassName}}::LookupSymbolInMap(
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
    return writeSubstitution(code, getClassPropertyMap(component, class), w)
}

func buildCppClientImplementationSymbolLookupFunctionABI(component ComponentDefinition, class ComponentDefinitionClass, isBaseClass bool, w LanguageWriter) error {
    err := error(nil)
    properties := getClassPropertyMap(component, class)
    beginCode := `
{{.NameSpace}}Result C{{.ClassName}}::SymbolLookupFunction_ABI(
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
    for j := 0; j < len(class.Methods); j++ {
        method := class.Methods[j]
        methodProperties := getMethodPropertyMap(component, class, method)
        err = writeSubstitution(methodCode, methodProperties, w)
        if err != nil {
            return err
        }
    }

    endCode := ""
    if isBaseClass {
        endCode = `
}
return LookupSymbolInMap(pProcName, sProcAddressMap, ppProcAddress);
}`
    } else {
        endCode = `
}
{{.NameSpace}}Result ret = LookupSymbolInMap(pProcName, sProcAddressMap, ppProcAddress);
if (ret == {{.NameSpaceUpper}}_ERROR_COULDNOTFINDLIBRARYEXPORT) {
    ret = tBASE::SymbolLookupFunction_ABI(pProcName, sProcAddressMap, ppProcAddress);
}
return ret;
}`
    }
    err = writeSubstitution(endCode, properties, w)
    if err != nil {
        return err
    }

    return err
}

func buildCppClientImplementationClassImplementationAPIMethods(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    beginComment := `
/*************************************************************************************************************************
 API-methods")
**************************************************************************************************************************/
`
    w.Writeln(beginComment)
    return nil
}

func buildCppclientImplementationClassImplementationABIMethods(component ComponentDefinition, class ComponentDefinitionClass, w LanguageWriter) error {
    beginComment := `
/*************************************************************************************************************************
 ABI-methods")
**************************************************************************************************************************/
`
    w.Writeln(beginComment)

    // TODO: need to generate the code here to forward the call.   This is the main obstacle to having this
    // be part of ACT.

    methodCode := `
{{.NameSpace}}Result C{{.ClassName}}::{{.MethodName}}_ABI({{.Parameters}})
{
  try {
    // map parameters
    // forward call to c++ function
    (C{{.ClassName}}*)p{{.ClassName}}.m_hHandle->{{.MethodName}}();
    // post-process parameters
  } catch ({{.NameSpace}}::Binding::E{{.NameSpace}}Exception&) {
    return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
  } catch (std::exception&) {
    return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
  } catch (...) {
    return {{.NameSpaceUpper}}_ERROR_GENERICEXCEPTION;
  }
  return {{.NameSpaceUpper}}_SUCCESS;
}`

    for j := 0; j < len(class.Methods); j++ {
        method := class.Methods[j]
        properties := getMethodPropertyMap(component, class, method)
        _, _, sParameters, err := WriteCCPPAbiMethod(method, component.NameSpace, class.ClassName, false, false, true)
        if (err != nil) {
            return err
        }
        properties["Parameters"] = sParameters
        err = writeSubstitution(methodCode, properties, w)
        if err != nil {
            return err
        }
    }

    return nil
}