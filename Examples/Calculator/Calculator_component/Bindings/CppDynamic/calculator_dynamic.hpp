/*++

Copyright (C) 2019 Calculator developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0.

Abstract: This is an autogenerated C++-Header file in order to allow an easy
 use of Calculator library

Interface version: 1.0.0

*/

#ifndef __CALCULATOR_CPPHEADER_DYNAMIC_CPP
#define __CALCULATOR_CPPHEADER_DYNAMIC_CPP

#include "calculator_types.hpp"
#include "calculator_dynamic.h"


#ifdef _WIN32
#include <windows.h>
#else // _WIN32
#include <dlfcn.h>
#endif // _WIN32
#include <string>
#include <memory>
#include <vector>
#include <exception>

namespace Calculator {

/*************************************************************************************************************************
 Forward Declaration of all classes
**************************************************************************************************************************/
class CWrapper;
class CBase;
class CVariable;
class CCalculator;

/*************************************************************************************************************************
 Declaration of deprecated class types
**************************************************************************************************************************/
typedef CWrapper CCalculatorWrapper;
typedef CBase CCalculatorBase;
typedef CVariable CCalculatorVariable;
typedef CCalculator CCalculatorCalculator;

/*************************************************************************************************************************
 Declaration of shared pointer types
**************************************************************************************************************************/
typedef std::shared_ptr<CWrapper> PWrapper;
typedef std::shared_ptr<CBase> PBase;
typedef std::shared_ptr<CVariable> PVariable;
typedef std::shared_ptr<CCalculator> PCalculator;

/*************************************************************************************************************************
 Declaration of deprecated shared pointer types
**************************************************************************************************************************/
typedef PWrapper PCalculatorWrapper;
typedef PBase PCalculatorBase;
typedef PVariable PCalculatorVariable;
typedef PCalculator PCalculatorCalculator;


/*************************************************************************************************************************
 Class ECalculatorException 
**************************************************************************************************************************/
class ECalculatorException : public std::exception {
protected:
	/**
	* Error code for the Exception.
	*/
	CalculatorResult m_errorCode;
	/**
	* Error message for the Exception.
	*/
	std::string m_errorMessage;

public:
	/**
	* Exception Constructor.
	*/
	ECalculatorException(CalculatorResult errorCode, const std::string & sErrorMessage)
		: m_errorMessage("Calculator Error " + std::to_string(errorCode) + " (" + sErrorMessage + ")")
	{
		m_errorCode = errorCode;
	}

	/**
	* Returns error code
	*/
	CalculatorResult getErrorCode() const noexcept
	{
		return m_errorCode;
	}

	/**
	* Returns error message
	*/
	const char* what() const noexcept
	{
		return m_errorMessage.c_str();
	}

};

/*************************************************************************************************************************
 Class CInputVector
**************************************************************************************************************************/
template <typename T>
class CInputVector {
private:
	
	const T* m_data;
	size_t m_size;
	
public:
	
	CInputVector( const std::vector<T>& vec)
		: m_data( vec.data() ), m_size( vec.size() )
	{
	}
	
	CInputVector( const T* in_data, size_t in_size)
		: m_data( in_data ), m_size(in_size )
	{
	}
	
	const T* data() const
	{
		return m_data;
	}
	
	size_t size() const
	{
		return m_size;
	}
	
};

// declare deprecated class name
template<typename T>
using CCalculatorInputVector = CInputVector<T>;

/*************************************************************************************************************************
 Class CWrapper 
**************************************************************************************************************************/
class CWrapper {
public:
	
	CWrapper(void* pSymbolLookupMethod)
	{
		CheckError(nullptr, initWrapperTable(&m_WrapperTable));
		CheckError(nullptr, loadWrapperTableFromSymbolLookupMethod(&m_WrapperTable, pSymbolLookupMethod));
		
		CheckError(nullptr, checkBinaryVersion());
	}
	
	CWrapper(const std::string &sFileName)
	{
		CheckError(nullptr, initWrapperTable(&m_WrapperTable));
		CheckError(nullptr, loadWrapperTable(&m_WrapperTable, sFileName.c_str()));
		
		CheckError(nullptr, checkBinaryVersion());
	}
	
	static PWrapper loadLibrary(const std::string &sFileName)
	{
		return std::make_shared<CWrapper>(sFileName);
	}
	
	static PWrapper loadLibraryFromSymbolLookupMethod(void* pSymbolLookupMethod)
	{
		return std::make_shared<CWrapper>(pSymbolLookupMethod);
	}
	
	~CWrapper()
	{
		releaseWrapperTable(&m_WrapperTable);
	}
	
	inline void CheckError(CBase * pBaseClass, CalculatorResult nResult);

	inline void GetVersion(Calculator_uint32 & nMajor, Calculator_uint32 & nMinor, Calculator_uint32 & nMicro);
	inline bool GetLastError(CBase * pInstance, std::string & sErrorMessage);
	inline void ReleaseInstance(CBase * pInstance);
	inline void AcquireInstance(CBase * pInstance);
	inline PVariable CreateVariable(const Calculator_double dInitialValue);
	inline PCalculator CreateCalculator();

private:
	sCalculatorDynamicWrapperTable m_WrapperTable;
	
	CalculatorResult checkBinaryVersion()
	{
		Calculator_uint32 nMajor, nMinor, nMicro;
		GetVersion(nMajor, nMinor, nMicro);
		if ( (nMajor != CALCULATOR_VERSION_MAJOR) || (nMinor < CALCULATOR_VERSION_MINOR) ) {
			return CALCULATOR_ERROR_INCOMPATIBLEBINARYVERSION;
		}
		return CALCULATOR_SUCCESS;
	}
	CalculatorResult initWrapperTable(sCalculatorDynamicWrapperTable * pWrapperTable);
	CalculatorResult releaseWrapperTable(sCalculatorDynamicWrapperTable * pWrapperTable);
	CalculatorResult loadWrapperTable(sCalculatorDynamicWrapperTable * pWrapperTable, const char * pLibraryFileName);
	CalculatorResult loadWrapperTableFromSymbolLookupMethod(sCalculatorDynamicWrapperTable * pWrapperTable, void* pSymbolLookupMethod);

	friend class CBase;
	friend class CVariable;
	friend class CCalculator;

};

	
/*************************************************************************************************************************
 Class CBase 
**************************************************************************************************************************/
class CBase {
public:
	
protected:
	/* Wrapper Object that created the class. */
	CWrapper * m_pWrapper;
	/* Handle to Instance in library*/
	CalculatorHandle m_pHandle;

	/* Checks for an Error code and raises Exceptions */
	void CheckError(CalculatorResult nResult)
	{
		if (m_pWrapper != nullptr)
			m_pWrapper->CheckError(this, nResult);
	}
public:
	/**
	* CBase::CBase - Constructor for Base class.
	*/
	CBase(CWrapper * pWrapper, CalculatorHandle pHandle)
		: m_pWrapper(pWrapper), m_pHandle(pHandle)
	{
	}

	/**
	* CBase::~CBase - Destructor for Base class.
	*/
	virtual ~CBase()
	{
		if (m_pWrapper != nullptr)
			m_pWrapper->ReleaseInstance(this);
		m_pWrapper = nullptr;
	}

	/**
	* CBase::GetHandle - Returns handle to instance.
	*/
	CalculatorHandle GetHandle()
	{
		return m_pHandle;
	}
	
	friend class CWrapper;
};
	
/*************************************************************************************************************************
 Class CVariable 
**************************************************************************************************************************/
class CVariable : public CBase {
public:
	
	/**
	* CVariable::CVariable - Constructor for Variable class.
	*/
	CVariable(CWrapper* pWrapper, CalculatorHandle pHandle)
		: CBase(pWrapper, pHandle)
	{
	}
	
	inline Calculator_double GetValue();
	inline void SetValue(const Calculator_double dValue);
};
	
/*************************************************************************************************************************
 Class CCalculator 
**************************************************************************************************************************/
class CCalculator : public CBase {
public:
	
	/**
	* CCalculator::CCalculator - Constructor for Calculator class.
	*/
	CCalculator(CWrapper* pWrapper, CalculatorHandle pHandle)
		: CBase(pWrapper, pHandle)
	{
	}
	
	inline void EnlistVariable(CVariable * pVariable);
	inline PVariable GetEnlistedVariable(const Calculator_uint32 nIndex);
	inline void ClearVariables();
	inline PVariable Multiply();
	inline PVariable Add();
};
	
	/**
	* CWrapper::GetVersion - retrieves the binary version of this library.
	* @param[out] nMajor - returns the major version of this library
	* @param[out] nMinor - returns the minor version of this library
	* @param[out] nMicro - returns the micro version of this library
	*/
	inline void CWrapper::GetVersion(Calculator_uint32 & nMajor, Calculator_uint32 & nMinor, Calculator_uint32 & nMicro)
	{
		CheckError(nullptr,m_WrapperTable.m_GetVersion(&nMajor, &nMinor, &nMicro));
	}
	
	/**
	* CWrapper::GetLastError - Returns the last error recorded on this object
	* @param[in] pInstance - Instance Handle
	* @param[out] sErrorMessage - Message of the last error
	* @return Is there a last error to query
	*/
	inline bool CWrapper::GetLastError(CBase * pInstance, std::string & sErrorMessage)
	{
		CalculatorHandle hInstance = nullptr;
		if (pInstance != nullptr) {
			hInstance = pInstance->GetHandle();
		};
		Calculator_uint32 bytesNeededErrorMessage = 0;
		Calculator_uint32 bytesWrittenErrorMessage = 0;
		bool resultHasError = 0;
		CheckError(nullptr,m_WrapperTable.m_GetLastError(hInstance, 0, &bytesNeededErrorMessage, nullptr, &resultHasError));
		std::vector<char> bufferErrorMessage(bytesNeededErrorMessage);
		CheckError(nullptr,m_WrapperTable.m_GetLastError(hInstance, bytesNeededErrorMessage, &bytesWrittenErrorMessage, &bufferErrorMessage[0], &resultHasError));
		sErrorMessage = std::string(&bufferErrorMessage[0]);
		
		return resultHasError;
	}
	
	/**
	* CWrapper::ReleaseInstance - Releases shared ownership of an Instance
	* @param[in] pInstance - Instance Handle
	*/
	inline void CWrapper::ReleaseInstance(CBase * pInstance)
	{
		CalculatorHandle hInstance = nullptr;
		if (pInstance != nullptr) {
			hInstance = pInstance->GetHandle();
		};
		CheckError(nullptr,m_WrapperTable.m_ReleaseInstance(hInstance));
	}
	
	/**
	* CWrapper::AcquireInstance - Acquires shared ownership of an Instance
	* @param[in] pInstance - Instance Handle
	*/
	inline void CWrapper::AcquireInstance(CBase * pInstance)
	{
		CalculatorHandle hInstance = nullptr;
		if (pInstance != nullptr) {
			hInstance = pInstance->GetHandle();
		};
		CheckError(nullptr,m_WrapperTable.m_AcquireInstance(hInstance));
	}
	
	/**
	* CWrapper::CreateVariable - Creates a new Variable instance
	* @param[in] dInitialValue - Initial value of the new Variable
	* @return New Variable instance
	*/
	inline PVariable CWrapper::CreateVariable(const Calculator_double dInitialValue)
	{
		CalculatorHandle hInstance = nullptr;
		CheckError(nullptr,m_WrapperTable.m_CreateVariable(dInitialValue, &hInstance));
		
		if (!hInstance) {
			CheckError(nullptr,CALCULATOR_ERROR_INVALIDPARAM);
		}
		return std::make_shared<CVariable>(this, hInstance);
	}
	
	/**
	* CWrapper::CreateCalculator - Creates a new Calculator instance
	* @return New Calculator instance
	*/
	inline PCalculator CWrapper::CreateCalculator()
	{
		CalculatorHandle hInstance = nullptr;
		CheckError(nullptr,m_WrapperTable.m_CreateCalculator(&hInstance));
		
		if (!hInstance) {
			CheckError(nullptr,CALCULATOR_ERROR_INVALIDPARAM);
		}
		return std::make_shared<CCalculator>(this, hInstance);
	}
	
	inline void CWrapper::CheckError(CBase * pBaseClass, CalculatorResult nResult)
	{
		if (nResult != 0) {
			std::string sErrorMessage;
			if (pBaseClass != nullptr) {
				GetLastError(pBaseClass, sErrorMessage);
			}
			throw ECalculatorException(nResult, sErrorMessage);
		}
	}
	

	inline CalculatorResult CWrapper::initWrapperTable(sCalculatorDynamicWrapperTable * pWrapperTable)
	{
		if (pWrapperTable == nullptr)
			return CALCULATOR_ERROR_INVALIDPARAM;
		
		pWrapperTable->m_LibraryHandle = nullptr;
		pWrapperTable->m_Variable_GetValue = nullptr;
		pWrapperTable->m_Variable_SetValue = nullptr;
		pWrapperTable->m_Calculator_EnlistVariable = nullptr;
		pWrapperTable->m_Calculator_GetEnlistedVariable = nullptr;
		pWrapperTable->m_Calculator_ClearVariables = nullptr;
		pWrapperTable->m_Calculator_Multiply = nullptr;
		pWrapperTable->m_Calculator_Add = nullptr;
		pWrapperTable->m_GetVersion = nullptr;
		pWrapperTable->m_GetLastError = nullptr;
		pWrapperTable->m_ReleaseInstance = nullptr;
		pWrapperTable->m_AcquireInstance = nullptr;
		pWrapperTable->m_CreateVariable = nullptr;
		pWrapperTable->m_CreateCalculator = nullptr;
		
		return CALCULATOR_SUCCESS;
	}

	inline CalculatorResult CWrapper::releaseWrapperTable(sCalculatorDynamicWrapperTable * pWrapperTable)
	{
		if (pWrapperTable == nullptr)
			return CALCULATOR_ERROR_INVALIDPARAM;
		
		if (pWrapperTable->m_LibraryHandle != nullptr) {
		#ifdef _WIN32
			HMODULE hModule = (HMODULE) pWrapperTable->m_LibraryHandle;
			FreeLibrary(hModule);
		#else // _WIN32
			dlclose(pWrapperTable->m_LibraryHandle);
		#endif // _WIN32
			return initWrapperTable(pWrapperTable);
		}
		
		return CALCULATOR_SUCCESS;
	}

	inline CalculatorResult CWrapper::loadWrapperTable(sCalculatorDynamicWrapperTable * pWrapperTable, const char * pLibraryFileName)
	{
		if (pWrapperTable == nullptr)
			return CALCULATOR_ERROR_INVALIDPARAM;
		if (pLibraryFileName == nullptr)
			return CALCULATOR_ERROR_INVALIDPARAM;
		
		#ifdef _WIN32
		// Convert filename to UTF16-string
		int nLength = (int)strlen(pLibraryFileName);
		int nBufferSize = nLength * 2 + 2;
		std::vector<wchar_t> wsLibraryFileName(nBufferSize);
		int nResult = MultiByteToWideChar(CP_UTF8, 0, pLibraryFileName, nLength, &wsLibraryFileName[0], nBufferSize);
		if (nResult == 0)
			return CALCULATOR_ERROR_COULDNOTLOADLIBRARY;
		
		HMODULE hLibrary = LoadLibraryW(wsLibraryFileName.data());
		if (hLibrary == 0) 
			return CALCULATOR_ERROR_COULDNOTLOADLIBRARY;
		#else // _WIN32
		void* hLibrary = dlopen(pLibraryFileName, RTLD_LAZY);
		if (hLibrary == 0) 
			return CALCULATOR_ERROR_COULDNOTLOADLIBRARY;
		dlerror();
		#endif // _WIN32
		
		#ifdef _WIN32
		pWrapperTable->m_Variable_GetValue = (PCalculatorVariable_GetValuePtr) GetProcAddress(hLibrary, "calculator_variable_getvalue");
		#else // _WIN32
		pWrapperTable->m_Variable_GetValue = (PCalculatorVariable_GetValuePtr) dlsym(hLibrary, "calculator_variable_getvalue");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Variable_GetValue == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Variable_SetValue = (PCalculatorVariable_SetValuePtr) GetProcAddress(hLibrary, "calculator_variable_setvalue");
		#else // _WIN32
		pWrapperTable->m_Variable_SetValue = (PCalculatorVariable_SetValuePtr) dlsym(hLibrary, "calculator_variable_setvalue");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Variable_SetValue == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Calculator_EnlistVariable = (PCalculatorCalculator_EnlistVariablePtr) GetProcAddress(hLibrary, "calculator_calculator_enlistvariable");
		#else // _WIN32
		pWrapperTable->m_Calculator_EnlistVariable = (PCalculatorCalculator_EnlistVariablePtr) dlsym(hLibrary, "calculator_calculator_enlistvariable");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Calculator_EnlistVariable == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Calculator_GetEnlistedVariable = (PCalculatorCalculator_GetEnlistedVariablePtr) GetProcAddress(hLibrary, "calculator_calculator_getenlistedvariable");
		#else // _WIN32
		pWrapperTable->m_Calculator_GetEnlistedVariable = (PCalculatorCalculator_GetEnlistedVariablePtr) dlsym(hLibrary, "calculator_calculator_getenlistedvariable");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Calculator_GetEnlistedVariable == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Calculator_ClearVariables = (PCalculatorCalculator_ClearVariablesPtr) GetProcAddress(hLibrary, "calculator_calculator_clearvariables");
		#else // _WIN32
		pWrapperTable->m_Calculator_ClearVariables = (PCalculatorCalculator_ClearVariablesPtr) dlsym(hLibrary, "calculator_calculator_clearvariables");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Calculator_ClearVariables == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Calculator_Multiply = (PCalculatorCalculator_MultiplyPtr) GetProcAddress(hLibrary, "calculator_calculator_multiply");
		#else // _WIN32
		pWrapperTable->m_Calculator_Multiply = (PCalculatorCalculator_MultiplyPtr) dlsym(hLibrary, "calculator_calculator_multiply");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Calculator_Multiply == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Calculator_Add = (PCalculatorCalculator_AddPtr) GetProcAddress(hLibrary, "calculator_calculator_add");
		#else // _WIN32
		pWrapperTable->m_Calculator_Add = (PCalculatorCalculator_AddPtr) dlsym(hLibrary, "calculator_calculator_add");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Calculator_Add == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_GetVersion = (PCalculatorGetVersionPtr) GetProcAddress(hLibrary, "calculator_getversion");
		#else // _WIN32
		pWrapperTable->m_GetVersion = (PCalculatorGetVersionPtr) dlsym(hLibrary, "calculator_getversion");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_GetVersion == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_GetLastError = (PCalculatorGetLastErrorPtr) GetProcAddress(hLibrary, "calculator_getlasterror");
		#else // _WIN32
		pWrapperTable->m_GetLastError = (PCalculatorGetLastErrorPtr) dlsym(hLibrary, "calculator_getlasterror");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_GetLastError == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_ReleaseInstance = (PCalculatorReleaseInstancePtr) GetProcAddress(hLibrary, "calculator_releaseinstance");
		#else // _WIN32
		pWrapperTable->m_ReleaseInstance = (PCalculatorReleaseInstancePtr) dlsym(hLibrary, "calculator_releaseinstance");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_ReleaseInstance == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_AcquireInstance = (PCalculatorAcquireInstancePtr) GetProcAddress(hLibrary, "calculator_acquireinstance");
		#else // _WIN32
		pWrapperTable->m_AcquireInstance = (PCalculatorAcquireInstancePtr) dlsym(hLibrary, "calculator_acquireinstance");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_AcquireInstance == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_CreateVariable = (PCalculatorCreateVariablePtr) GetProcAddress(hLibrary, "calculator_createvariable");
		#else // _WIN32
		pWrapperTable->m_CreateVariable = (PCalculatorCreateVariablePtr) dlsym(hLibrary, "calculator_createvariable");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_CreateVariable == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_CreateCalculator = (PCalculatorCreateCalculatorPtr) GetProcAddress(hLibrary, "calculator_createcalculator");
		#else // _WIN32
		pWrapperTable->m_CreateCalculator = (PCalculatorCreateCalculatorPtr) dlsym(hLibrary, "calculator_createcalculator");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_CreateCalculator == nullptr)
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		pWrapperTable->m_LibraryHandle = hLibrary;
		return CALCULATOR_SUCCESS;
	}

	inline CalculatorResult CWrapper::loadWrapperTableFromSymbolLookupMethod(sCalculatorDynamicWrapperTable * pWrapperTable, void* pSymbolLookupMethod)
{
		if (pWrapperTable == nullptr)
			return CALCULATOR_ERROR_INVALIDPARAM;
		if (pSymbolLookupMethod == nullptr)
			return CALCULATOR_ERROR_INVALIDPARAM;
		
		typedef CalculatorResult(*SymbolLookupType)(const char*, void**);
		
		SymbolLookupType pLookup = (SymbolLookupType)pSymbolLookupMethod;
		
		CalculatorResult eLookupError = CALCULATOR_SUCCESS;
		eLookupError = (*pLookup)("calculator_variable_getvalue", (void**)&(pWrapperTable->m_Variable_GetValue));
		if ( (eLookupError != 0) || (pWrapperTable->m_Variable_GetValue == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_variable_setvalue", (void**)&(pWrapperTable->m_Variable_SetValue));
		if ( (eLookupError != 0) || (pWrapperTable->m_Variable_SetValue == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_calculator_enlistvariable", (void**)&(pWrapperTable->m_Calculator_EnlistVariable));
		if ( (eLookupError != 0) || (pWrapperTable->m_Calculator_EnlistVariable == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_calculator_getenlistedvariable", (void**)&(pWrapperTable->m_Calculator_GetEnlistedVariable));
		if ( (eLookupError != 0) || (pWrapperTable->m_Calculator_GetEnlistedVariable == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_calculator_clearvariables", (void**)&(pWrapperTable->m_Calculator_ClearVariables));
		if ( (eLookupError != 0) || (pWrapperTable->m_Calculator_ClearVariables == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_calculator_multiply", (void**)&(pWrapperTable->m_Calculator_Multiply));
		if ( (eLookupError != 0) || (pWrapperTable->m_Calculator_Multiply == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_calculator_add", (void**)&(pWrapperTable->m_Calculator_Add));
		if ( (eLookupError != 0) || (pWrapperTable->m_Calculator_Add == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_getversion", (void**)&(pWrapperTable->m_GetVersion));
		if ( (eLookupError != 0) || (pWrapperTable->m_GetVersion == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_getlasterror", (void**)&(pWrapperTable->m_GetLastError));
		if ( (eLookupError != 0) || (pWrapperTable->m_GetLastError == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_releaseinstance", (void**)&(pWrapperTable->m_ReleaseInstance));
		if ( (eLookupError != 0) || (pWrapperTable->m_ReleaseInstance == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_acquireinstance", (void**)&(pWrapperTable->m_AcquireInstance));
		if ( (eLookupError != 0) || (pWrapperTable->m_AcquireInstance == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_createvariable", (void**)&(pWrapperTable->m_CreateVariable));
		if ( (eLookupError != 0) || (pWrapperTable->m_CreateVariable == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("calculator_createcalculator", (void**)&(pWrapperTable->m_CreateCalculator));
		if ( (eLookupError != 0) || (pWrapperTable->m_CreateCalculator == nullptr) )
			return CALCULATOR_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		return CALCULATOR_SUCCESS;
}

	
	
	/**
	 * Method definitions for class CBase
	 */
	
	/**
	 * Method definitions for class CVariable
	 */
	
	/**
	* CVariable::GetValue - Returns the current value of this Variable
	* @return The current value of this Variable
	*/
	Calculator_double CVariable::GetValue()
	{
		Calculator_double resultValue = 0;
		CheckError(m_pWrapper->m_WrapperTable.m_Variable_GetValue(m_pHandle, &resultValue));
		
		return resultValue;
	}
	
	/**
	* CVariable::SetValue - Set the numerical value of this Variable
	* @param[in] dValue - The new value of this Variable
	*/
	void CVariable::SetValue(const Calculator_double dValue)
	{
		CheckError(m_pWrapper->m_WrapperTable.m_Variable_SetValue(m_pHandle, dValue));
	}
	
	/**
	 * Method definitions for class CCalculator
	 */
	
	/**
	* CCalculator::EnlistVariable - Adds a Variable to the list of Variables this calculator works on
	* @param[in] pVariable - The new variable in this calculator
	*/
	void CCalculator::EnlistVariable(CVariable * pVariable)
	{
		CalculatorHandle hVariable = nullptr;
		if (pVariable != nullptr) {
			hVariable = pVariable->GetHandle();
		};
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_EnlistVariable(m_pHandle, hVariable));
	}
	
	/**
	* CCalculator::GetEnlistedVariable - Returns an instance of a enlisted variable
	* @param[in] nIndex - The index of the variable to query
	* @return The Index-th variable in this calculator
	*/
	PVariable CCalculator::GetEnlistedVariable(const Calculator_uint32 nIndex)
	{
		CalculatorHandle hVariable = nullptr;
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_GetEnlistedVariable(m_pHandle, nIndex, &hVariable));
		
		if (!hVariable) {
			CheckError(CALCULATOR_ERROR_INVALIDPARAM);
		}
		return std::make_shared<CVariable>(m_pWrapper, hVariable);
	}
	
	/**
	* CCalculator::ClearVariables - Clears all variables in enlisted in this calculator
	*/
	void CCalculator::ClearVariables()
	{
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_ClearVariables(m_pHandle));
	}
	
	/**
	* CCalculator::Multiply - Multiplies all enlisted variables
	* @return Variable that holds the product of all enlisted Variables
	*/
	PVariable CCalculator::Multiply()
	{
		CalculatorHandle hInstance = nullptr;
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_Multiply(m_pHandle, &hInstance));
		
		if (!hInstance) {
			CheckError(CALCULATOR_ERROR_INVALIDPARAM);
		}
		return std::make_shared<CVariable>(m_pWrapper, hInstance);
	}
	
	/**
	* CCalculator::Add - Sums all enlisted variables
	* @return Variable that holds the sum of all enlisted Variables
	*/
	PVariable CCalculator::Add()
	{
		CalculatorHandle hInstance = nullptr;
		CheckError(m_pWrapper->m_WrapperTable.m_Calculator_Add(m_pHandle, &hInstance));
		
		if (!hInstance) {
			CheckError(CALCULATOR_ERROR_INVALIDPARAM);
		}
		return std::make_shared<CVariable>(m_pWrapper, hInstance);
	}

} // namespace Calculator

#endif // __CALCULATOR_CPPHEADER_DYNAMIC_CPP

