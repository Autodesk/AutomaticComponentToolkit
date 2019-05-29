# Injection

These are the things required to get a really SEAMLESS Component Injection working
(i.e. Implementer AND binding-users don't even see the handles :) )

# CPP-world

## Needed Cpp Implementation:
```c++
class ICalculator : public virtual IBase {
public:
	/**
	* ICalculator::EnlistVariable - Adds a Variable to the list of Variables this calculator works on
	* @param[in] pVariable - The new variable in this calculator
	*/
	virtual void EnlistVariable(Numbers::PVariable pVariable) = 0;

	/**
	* ICalculator::GetEnlistedVariable - Returns an instance of a enlisted variable
	* @param[in] nIndex - The index of the variable to query
	* @return The Index-th variable in this calculator
	*/
	virtual Numbers::PVariable GetEnlistedVariable(const Calculation_uint32 nIndex) = 0;

	/**
	* ICalculator::ClearVariables - Clears all variables in enlisted in this calculator
	*/
	virtual void ClearVariables() = 0;

	/**
	* ICalculator::Multiply - Multiplies all enlisted variables
	* @return Variable that holds the product of all enlisted Variables
	*/
	virtual Numbers::PVariable Multiply() = 0;

	/**
	* ICalculator::Add - Sums all enlisted variables
	* @return Variable that holds the sum of all enlisted Variables
	*/
	virtual Numbers::PVariable Add() = 0;
};
```

```c++
// NEW:
#include "..\..\..\..\Numbers_component\Bindings\CppDynamic\numbers_dynamic.hpp"

namespace Calculation {
namespace Impl {

static Numbers::CWrapper* gPWrapper;
```

```c++
Numbers::PVariable CCalculator::Add()
{
	auto pVariable = gPWrapper->CreateVariable(0);
	return pVariable;
}
```

```c++
CalculationResult calculation_calculator_enlistvariable(Calculation_Calculator pCalculator, Numbers::HandleVariable pVariable)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	try {

		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ECalculationInterfaceException(CALCULATION_ERROR_INVALIDCAST);
		// NEW
		// TODO: the next line ALSO needs to increase the refcount on pVariable
		Numbers::PVariable pNumbersVariable = std::make_shared<Numbers::CVariable>(gPWrapper, pVariable);
		pICalculator->EnlistVariable(pNumbersVariable);
		// \NEW
		return CALCULATION_SUCCESS;
	}
```

```c++
CalculationResult calculation_calculator_multiply(Calculation_Calculator pCalculator, Numbers::HandleVariable * pInstance)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	try {
		if (pInstance == nullptr)
			throw ECalculationInterfaceException (CALCULATION_ERROR_INVALIDPARAM);

		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ECalculationInterfaceException(CALCULATION_ERROR_INVALIDCAST);

		// NEW
		Numbers::PVariable pNumbersInstance;

		pNumbersInstance = pICalculator->Multiply();
		// TODO: the refcount on pVariable needs to be increased here!
		*pInstance = pNumbersInstance->GetHandle();
		// \NEW

		return CALCULATION_SUCCESS;
	}
```

```c++
CalculationResult calculation_setnumberssymboladdressmethod(Calculation_uint64 nSymbolAddressMethod)
{
	IBase* pIBaseClass = nullptr;

	try {

		// NEW: setup global wrapper
		CWrapper::SetNumbersSymbolAddressMethod(nSymbolAddressMethod);
```

AND:
- SymbolAddressMethod


## Needed CppBinding:
```c++
// NEW:
#include "..\..\..\Numbers_component\Bindings\CppDynamic\numbers_dynamic.hpp"
```
```c++
private:
	sCalculationDynamicWrapperTable m_WrapperTable;
	Numbers::CWrapper* m_pNumbersWrapper; // or something along those lines...
```
```c++
class CCalculator : public CBase {
public:
	
	/**
	* CCalculator::CCalculator - Constructor for Calculator class.
	*/
	CCalculator(CWrapper* pWrapper, CalculationHandle pHandle)
		: CBase(pWrapper, pHandle)
	{
	}
	// NEW
	inline void EnlistVariable(Numbers::CVariable* pVariable);
	inline Numbers::PVariable GetEnlistedVariable(const Calculation_uint32 nIndex);
	inline void ClearVariables();
	inline Numbers::PVariable Multiply();
	inline Numbers::PVariable Add();
	// \NEW
};

```c++
// NEW
	void CCalculator::EnlistVariable(Numbers::CVariable * pVariable)
	{
		Numbers::HandleVariable hVariable = nullptr;
		if (pVariable != nullptr) {
			hVariable = pVariable->GetHandle();
		};
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_EnlistVariable(m_pHandle, hVariable));
	}
	// \NEW
```

```c++
// NEW
	Numbers::PVariable CCalculator::GetEnlistedVariable(const Calculation_uint32 nIndex)
	{
		Numbers::HandleVariable hVariable = nullptr;
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_GetEnlistedVariable(m_pHandle, nIndex, &hVariable));
		// no need to increase refcount here, the implementation has already done that
		return std::make_shared<Numbers::CVariable>(m_pWrapper->m_pNumbersWrapper, hVariable);
	}
	// \NEW
```

```c++
inline void CWrapper::SetNumbersSymbolAddressMethod(const Calculation_uint64 nSymbolAddressMethod)
	{
		// NEW: create binding side wrapper here
		CheckError(nullptr,m_WrapperTable.m_SetNumbersSymbolAddressMethod(nSymbolAddressMethod));
	}
```

AND:
- create wrapper from "SymbolAddressMethod" (Cpp + Pascal Only)



# Pascal-world
## Pascal Implementation
## Pascal Binding
- create wrapper from "SymbolAddressMethod" (Cpp + Pascal Only)


# Other Bindings:
## Python
## CSharp
... 


