/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated C++-Header file in order to allow an easy
 use of Numbers library

Interface version: 1.0.0

*/

#ifndef __NUMBERS_CPPHEADER_DYNAMIC_CPP
#define __NUMBERS_CPPHEADER_DYNAMIC_CPP

#include "numbers_types.hpp"
#include "numbers_dynamic.h"


#ifdef _WIN32
#include <windows.h>
#else // _WIN32
#include <dlfcn.h>
#endif // _WIN32
#include <array>
#include <string>
#include <memory>
#include <vector>
#include <exception>

namespace Numbers {

/*************************************************************************************************************************
 Forward Declaration of all classes
**************************************************************************************************************************/
class CWrapper;
class CBase;
class CVariable;

/*************************************************************************************************************************
 Declaration of deprecated class types
**************************************************************************************************************************/
typedef CWrapper CNumbersWrapper;
typedef CBase CNumbersBase;
typedef CVariable CNumbersVariable;

/*************************************************************************************************************************
 Declaration of shared pointer types
**************************************************************************************************************************/
typedef std::shared_ptr<CWrapper> PWrapper;
typedef std::shared_ptr<CBase> PBase;
typedef std::shared_ptr<CVariable> PVariable;

/*************************************************************************************************************************
 Declaration of deprecated shared pointer types
**************************************************************************************************************************/
typedef PWrapper PNumbersWrapper;
typedef PBase PNumbersBase;
typedef PVariable PNumbersVariable;


/*************************************************************************************************************************
 classParam Definition
**************************************************************************************************************************/

template<class T> class classParam {
private:
	const T* m_ptr;

public:
	classParam(const T* ptr)
		: m_ptr (ptr)
	{
	}

	classParam(std::shared_ptr <T> sharedPtr)
		: m_ptr (sharedPtr.get())
	{
	}

	NumbersHandle GetHandle()
	{
		if (m_ptr != nullptr)
			return m_ptr->handle();
		return nullptr;
	}
};

/*************************************************************************************************************************
 Class ENumbersException 
**************************************************************************************************************************/
class ENumbersException : public std::exception {
protected:
	/**
	* Error code for the Exception.
	*/
	NumbersResult m_errorCode;
	/**
	* Error message for the Exception.
	*/
	std::string m_errorMessage;
	std::string m_originalErrorMessage;

public:
	/**
	* Exception Constructor.
	*/
	ENumbersException(NumbersResult errorCode, const std::string & sErrorMessage)
		: m_errorCode(errorCode), m_originalErrorMessage(sErrorMessage)
	{
		m_errorMessage = buildErrorMessage();
	}

	/**
	* Returns error code
	*/
	NumbersResult getErrorCode() const noexcept
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

	const char* getErrorMessage() const noexcept
	{
		return m_originalErrorMessage.c_str();
	}

	const char* getErrorName() const noexcept
	{
		switch(getErrorCode()) {
			case NUMBERS_SUCCESS: return "SUCCESS";
			case NUMBERS_ERROR_NOTIMPLEMENTED: return "NOTIMPLEMENTED";
			case NUMBERS_ERROR_INVALIDPARAM: return "INVALIDPARAM";
			case NUMBERS_ERROR_INVALIDCAST: return "INVALIDCAST";
			case NUMBERS_ERROR_BUFFERTOOSMALL: return "BUFFERTOOSMALL";
			case NUMBERS_ERROR_GENERICEXCEPTION: return "GENERICEXCEPTION";
			case NUMBERS_ERROR_COULDNOTLOADLIBRARY: return "COULDNOTLOADLIBRARY";
			case NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT: return "COULDNOTFINDLIBRARYEXPORT";
			case NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION: return "INCOMPATIBLEBINARYVERSION";
		}
		return "UNKNOWN";
	}

	const char* getErrorDescription() const noexcept
	{
		switch(getErrorCode()) {
			case NUMBERS_SUCCESS: return "success";
			case NUMBERS_ERROR_NOTIMPLEMENTED: return "functionality not implemented";
			case NUMBERS_ERROR_INVALIDPARAM: return "an invalid parameter was passed";
			case NUMBERS_ERROR_INVALIDCAST: return "a type cast failed";
			case NUMBERS_ERROR_BUFFERTOOSMALL: return "a provided buffer is too small";
			case NUMBERS_ERROR_GENERICEXCEPTION: return "a generic exception occurred";
			case NUMBERS_ERROR_COULDNOTLOADLIBRARY: return "the library could not be loaded";
			case NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT: return "a required exported symbol could not be found in the library";
			case NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION: return "the version of the binary interface does not match the bindings interface";
		}
		return "unknown error";
	}

private:

	std::string buildErrorMessage() const noexcept
	{
		std::string msg = m_originalErrorMessage;
		if (msg.empty()) {
			msg = getErrorDescription();
		}
		return std::string("Error: ") + getErrorName() + ": " + msg;
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
	
	explicit CInputVector( const std::vector<T>& vec)
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
using CNumbersInputVector = CInputVector<T>;

/*************************************************************************************************************************
 Class CWrapper 
**************************************************************************************************************************/
class CWrapper {
public:
	
	explicit CWrapper(void* pSymbolLookupMethod)
	{
		CheckError(nullptr, initWrapperTable(&m_WrapperTable));
		CheckError(nullptr, loadWrapperTableFromSymbolLookupMethod(&m_WrapperTable, pSymbolLookupMethod));
		
		CheckError(nullptr, checkBinaryVersion());
	}
	
	explicit CWrapper(const std::string &sFileName)
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
	
	inline void CheckError(CBase * pBaseClass, NumbersResult nResult);

	inline PVariable CreateVariable(const Numbers_double dInitialValue);
	inline void GetVersion(Numbers_uint32 & nMajor, Numbers_uint32 & nMinor, Numbers_uint32 & nMicro);
	inline bool GetLastError(classParam<CBase> pInstance, std::string & sErrorMessage);
	inline void ReleaseInstance(classParam<CBase> pInstance);
	inline void AcquireInstance(classParam<CBase> pInstance);
	inline Numbers_pvoid GetSymbolLookupMethod();

	inline CBase* polymorphicFactory(NumbersHandle);

private:
	sNumbersDynamicWrapperTable m_WrapperTable;
	
	NumbersResult checkBinaryVersion()
	{
		Numbers_uint32 nMajor, nMinor, nMicro;
		GetVersion(nMajor, nMinor, nMicro);
		if (nMajor != NUMBERS_VERSION_MAJOR) {
			return NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION;
		}
		return NUMBERS_SUCCESS;
	}
	NumbersResult initWrapperTable(sNumbersDynamicWrapperTable * pWrapperTable);
	NumbersResult releaseWrapperTable(sNumbersDynamicWrapperTable * pWrapperTable);
	NumbersResult loadWrapperTable(sNumbersDynamicWrapperTable * pWrapperTable, const char * pLibraryFileName);
	NumbersResult loadWrapperTableFromSymbolLookupMethod(sNumbersDynamicWrapperTable * pWrapperTable, void* pSymbolLookupMethod);

	friend class CBase;
	friend class CVariable;

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
	NumbersHandle m_pHandle;

	/* Checks for an Error code and raises Exceptions */
	void CheckError(NumbersResult nResult)
	{
		if (m_pWrapper != nullptr)
			m_pWrapper->CheckError(this, nResult);
	}
public:
	/**
	* CBase::CBase - Constructor for Base class.
	*/
	CBase(CWrapper * pWrapper, NumbersHandle pHandle)
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
	* CBase::handle - Returns handle to instance.
	*/
	NumbersHandle handle() const
	{
		return m_pHandle;
	}

	/**
	* CBase::wrapper - Returns wrapper instance.
	*/
	CWrapper * wrapper() const
	{
		return m_pWrapper;
	}

	friend class CWrapper;
	inline Numbers_uint64 ClassTypeId();
};
	
/*************************************************************************************************************************
 Class CVariable 
**************************************************************************************************************************/
class CVariable : public CBase {
public:
	
	/**
	* CVariable::CVariable - Constructor for Variable class.
	*/
	CVariable(CWrapper* pWrapper, NumbersHandle pHandle)
		: CBase(pWrapper, pHandle)
	{
	}
	
	inline Numbers_double GetValue();
	inline void SetValue(const Numbers_double dValue);
};

/*************************************************************************************************************************
 RTTI: Polymorphic Factory implementation
**************************************************************************************************************************/

/**
* IMPORTANT: PolymorphicFactory method should not be used by application directly.
*            It's designed to be used on NumbersHandle object only once.
*            If it's used on any existing object as a form of dynamic cast then
*            CWrapper::AcquireInstance(CBase object) must be called after instantiating new object.
*            This is important to keep reference count matching between application and library sides.
*/
inline CBase* CWrapper::polymorphicFactory(NumbersHandle pHandle)
{
	Numbers_uint64 resultClassTypeId = 0;
	CheckError(nullptr, m_WrapperTable.m_Base_ClassTypeId(pHandle, &resultClassTypeId));
	switch(resultClassTypeId) {
		case 0x27799F69B3FD1C9EUL: return new CBase(this, pHandle); break; // First 64 bits of SHA1 of a string: "Numbers::Base"
		case 0x23934EDF762423EAUL: return new CVariable(this, pHandle); break; // First 64 bits of SHA1 of a string: "Numbers::Variable"
	}
	return new CBase(this, pHandle);
}
	
	/**
	* CWrapper::CreateVariable - Creates a new Variable instance
	* @param[in] dInitialValue - Initial value of the new Variable
	* @return New Variable instance
	*/
	inline PVariable CWrapper::CreateVariable(const Numbers_double dInitialValue)
	{
		NumbersHandle hInstance = nullptr;
		CheckError(nullptr,m_WrapperTable.m_CreateVariable(dInitialValue, &hInstance));
		
		if (!hInstance) {
			CheckError(nullptr,NUMBERS_ERROR_INVALIDPARAM);
		}
		return std::shared_ptr<CVariable>(dynamic_cast<CVariable*>(this->polymorphicFactory(hInstance)));
	}
	
	/**
	* CWrapper::GetVersion - retrieves the binary version of this library.
	* @param[out] nMajor - returns the major version of this library
	* @param[out] nMinor - returns the minor version of this library
	* @param[out] nMicro - returns the micro version of this library
	*/
	inline void CWrapper::GetVersion(Numbers_uint32 & nMajor, Numbers_uint32 & nMinor, Numbers_uint32 & nMicro)
	{
		CheckError(nullptr,m_WrapperTable.m_GetVersion(&nMajor, &nMinor, &nMicro));
	}
	
	/**
	* CWrapper::GetLastError - Returns the last error recorded on this object
	* @param[in] pInstance - Instance Handle
	* @param[out] sErrorMessage - Message of the last error
	* @return Is there a last error to query
	*/
	inline bool CWrapper::GetLastError(classParam<CBase> pInstance, std::string & sErrorMessage)
	{
		NumbersHandle hInstance = pInstance.GetHandle();
		Numbers_uint32 bytesNeededErrorMessage = 0;
		Numbers_uint32 bytesWrittenErrorMessage = 0;
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
	inline void CWrapper::ReleaseInstance(classParam<CBase> pInstance)
	{
		NumbersHandle hInstance = pInstance.GetHandle();
		CheckError(nullptr,m_WrapperTable.m_ReleaseInstance(hInstance));
	}
	
	/**
	* CWrapper::AcquireInstance - Acquires shared ownership of an Instance
	* @param[in] pInstance - Instance Handle
	*/
	inline void CWrapper::AcquireInstance(classParam<CBase> pInstance)
	{
		NumbersHandle hInstance = pInstance.GetHandle();
		CheckError(nullptr,m_WrapperTable.m_AcquireInstance(hInstance));
	}
	
	/**
	* CWrapper::GetSymbolLookupMethod - Returns the address of the SymbolLookupMethod
	* @return Address of the SymbolAddressMethod
	*/
	inline Numbers_pvoid CWrapper::GetSymbolLookupMethod()
	{
		Numbers_pvoid resultSymbolLookupMethod = 0;
		CheckError(nullptr,m_WrapperTable.m_GetSymbolLookupMethod(&resultSymbolLookupMethod));
		
		return resultSymbolLookupMethod;
	}
	
	inline void CWrapper::CheckError(CBase * pBaseClass, NumbersResult nResult)
	{
		if (nResult != 0) {
			std::string sErrorMessage;
			if (pBaseClass != nullptr) {
				GetLastError(pBaseClass, sErrorMessage);
			}
			throw ENumbersException(nResult, sErrorMessage);
		}
	}
	

	inline NumbersResult CWrapper::initWrapperTable(sNumbersDynamicWrapperTable * pWrapperTable)
	{
		if (pWrapperTable == nullptr)
			return NUMBERS_ERROR_INVALIDPARAM;
		
		pWrapperTable->m_LibraryHandle = nullptr;
		pWrapperTable->m_Base_ClassTypeId = nullptr;
		pWrapperTable->m_Variable_GetValue = nullptr;
		pWrapperTable->m_Variable_SetValue = nullptr;
		pWrapperTable->m_CreateVariable = nullptr;
		pWrapperTable->m_GetVersion = nullptr;
		pWrapperTable->m_GetLastError = nullptr;
		pWrapperTable->m_ReleaseInstance = nullptr;
		pWrapperTable->m_AcquireInstance = nullptr;
		pWrapperTable->m_GetSymbolLookupMethod = nullptr;
		
		return NUMBERS_SUCCESS;
	}

	inline NumbersResult CWrapper::releaseWrapperTable(sNumbersDynamicWrapperTable * pWrapperTable)
	{
		if (pWrapperTable == nullptr)
			return NUMBERS_ERROR_INVALIDPARAM;
		
		if (pWrapperTable->m_LibraryHandle != nullptr) {
		#ifdef _WIN32
			HMODULE hModule = (HMODULE) pWrapperTable->m_LibraryHandle;
			FreeLibrary(hModule);
		#else // _WIN32
			dlclose(pWrapperTable->m_LibraryHandle);
		#endif // _WIN32
			return initWrapperTable(pWrapperTable);
		}
		
		return NUMBERS_SUCCESS;
	}

	inline NumbersResult CWrapper::loadWrapperTable(sNumbersDynamicWrapperTable * pWrapperTable, const char * pLibraryFileName)
	{
		if (pWrapperTable == nullptr)
			return NUMBERS_ERROR_INVALIDPARAM;
		if (pLibraryFileName == nullptr)
			return NUMBERS_ERROR_INVALIDPARAM;
		
		#ifdef _WIN32
		// Convert filename to UTF16-string
		int nLength = static_cast<int>(strnlen_s(pLibraryFileName, MAX_PATH));
		int nBufferSize = nLength * 2 + 2;
		std::vector<wchar_t> wsLibraryFileName(nBufferSize);
		int nResult = MultiByteToWideChar(CP_UTF8, 0, pLibraryFileName, nLength, &wsLibraryFileName[0], nBufferSize);
		if (nResult == 0)
			return NUMBERS_ERROR_COULDNOTLOADLIBRARY;
		
		HMODULE hLibrary = LoadLibraryW(wsLibraryFileName.data());
		if (hLibrary == 0) 
			return NUMBERS_ERROR_COULDNOTLOADLIBRARY;
		#else // _WIN32
		void* hLibrary = dlopen(pLibraryFileName, RTLD_LAZY);
		if (hLibrary == 0) 
			return NUMBERS_ERROR_COULDNOTLOADLIBRARY;
		dlerror();
		#endif // _WIN32
		
		#ifdef _WIN32
		pWrapperTable->m_Base_ClassTypeId = (PNumbersBase_ClassTypeIdPtr) GetProcAddress(hLibrary, "numbers_base_classtypeid");
		#else // _WIN32
		pWrapperTable->m_Base_ClassTypeId = (PNumbersBase_ClassTypeIdPtr) dlsym(hLibrary, "numbers_base_classtypeid");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Base_ClassTypeId == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Variable_GetValue = (PNumbersVariable_GetValuePtr) GetProcAddress(hLibrary, "numbers_variable_getvalue");
		#else // _WIN32
		pWrapperTable->m_Variable_GetValue = (PNumbersVariable_GetValuePtr) dlsym(hLibrary, "numbers_variable_getvalue");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Variable_GetValue == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_Variable_SetValue = (PNumbersVariable_SetValuePtr) GetProcAddress(hLibrary, "numbers_variable_setvalue");
		#else // _WIN32
		pWrapperTable->m_Variable_SetValue = (PNumbersVariable_SetValuePtr) dlsym(hLibrary, "numbers_variable_setvalue");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_Variable_SetValue == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_CreateVariable = (PNumbersCreateVariablePtr) GetProcAddress(hLibrary, "numbers_createvariable");
		#else // _WIN32
		pWrapperTable->m_CreateVariable = (PNumbersCreateVariablePtr) dlsym(hLibrary, "numbers_createvariable");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_CreateVariable == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_GetVersion = (PNumbersGetVersionPtr) GetProcAddress(hLibrary, "numbers_getversion");
		#else // _WIN32
		pWrapperTable->m_GetVersion = (PNumbersGetVersionPtr) dlsym(hLibrary, "numbers_getversion");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_GetVersion == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_GetLastError = (PNumbersGetLastErrorPtr) GetProcAddress(hLibrary, "numbers_getlasterror");
		#else // _WIN32
		pWrapperTable->m_GetLastError = (PNumbersGetLastErrorPtr) dlsym(hLibrary, "numbers_getlasterror");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_GetLastError == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_ReleaseInstance = (PNumbersReleaseInstancePtr) GetProcAddress(hLibrary, "numbers_releaseinstance");
		#else // _WIN32
		pWrapperTable->m_ReleaseInstance = (PNumbersReleaseInstancePtr) dlsym(hLibrary, "numbers_releaseinstance");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_ReleaseInstance == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_AcquireInstance = (PNumbersAcquireInstancePtr) GetProcAddress(hLibrary, "numbers_acquireinstance");
		#else // _WIN32
		pWrapperTable->m_AcquireInstance = (PNumbersAcquireInstancePtr) dlsym(hLibrary, "numbers_acquireinstance");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_AcquireInstance == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef _WIN32
		pWrapperTable->m_GetSymbolLookupMethod = (PNumbersGetSymbolLookupMethodPtr) GetProcAddress(hLibrary, "numbers_getsymbollookupmethod");
		#else // _WIN32
		pWrapperTable->m_GetSymbolLookupMethod = (PNumbersGetSymbolLookupMethodPtr) dlsym(hLibrary, "numbers_getsymbollookupmethod");
		dlerror();
		#endif // _WIN32
		if (pWrapperTable->m_GetSymbolLookupMethod == nullptr)
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		pWrapperTable->m_LibraryHandle = hLibrary;
		return NUMBERS_SUCCESS;
	}

	inline NumbersResult CWrapper::loadWrapperTableFromSymbolLookupMethod(sNumbersDynamicWrapperTable * pWrapperTable, void* pSymbolLookupMethod)
{
		if (pWrapperTable == nullptr)
			return NUMBERS_ERROR_INVALIDPARAM;
		if (pSymbolLookupMethod == nullptr)
			return NUMBERS_ERROR_INVALIDPARAM;
		
		typedef NumbersResult(*SymbolLookupType)(const char*, void**);
		
		SymbolLookupType pLookup = (SymbolLookupType)pSymbolLookupMethod;
		
		NumbersResult eLookupError = NUMBERS_SUCCESS;
		eLookupError = (*pLookup)("numbers_base_classtypeid", (void**)&(pWrapperTable->m_Base_ClassTypeId));
		if ( (eLookupError != 0) || (pWrapperTable->m_Base_ClassTypeId == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_variable_getvalue", (void**)&(pWrapperTable->m_Variable_GetValue));
		if ( (eLookupError != 0) || (pWrapperTable->m_Variable_GetValue == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_variable_setvalue", (void**)&(pWrapperTable->m_Variable_SetValue));
		if ( (eLookupError != 0) || (pWrapperTable->m_Variable_SetValue == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_createvariable", (void**)&(pWrapperTable->m_CreateVariable));
		if ( (eLookupError != 0) || (pWrapperTable->m_CreateVariable == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_getversion", (void**)&(pWrapperTable->m_GetVersion));
		if ( (eLookupError != 0) || (pWrapperTable->m_GetVersion == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_getlasterror", (void**)&(pWrapperTable->m_GetLastError));
		if ( (eLookupError != 0) || (pWrapperTable->m_GetLastError == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_releaseinstance", (void**)&(pWrapperTable->m_ReleaseInstance));
		if ( (eLookupError != 0) || (pWrapperTable->m_ReleaseInstance == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_acquireinstance", (void**)&(pWrapperTable->m_AcquireInstance));
		if ( (eLookupError != 0) || (pWrapperTable->m_AcquireInstance == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		eLookupError = (*pLookup)("numbers_getsymbollookupmethod", (void**)&(pWrapperTable->m_GetSymbolLookupMethod));
		if ( (eLookupError != 0) || (pWrapperTable->m_GetSymbolLookupMethod == nullptr) )
			return NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		return NUMBERS_SUCCESS;
}

	
	
	/**
	 * Method definitions for class CBase
	 */
	
	/**
	* CBase::ClassTypeId - Get Class Type Id
	* @return Class type as a 64 bits integer
	*/
	Numbers_uint64 CBase::ClassTypeId()
	{
		Numbers_uint64 resultClassTypeId = 0;
		CheckError(m_pWrapper->m_WrapperTable.m_Base_ClassTypeId(m_pHandle, &resultClassTypeId));
		
		return resultClassTypeId;
	}
	
	/**
	 * Method definitions for class CVariable
	 */
	
	/**
	* CVariable::GetValue - Returns the current value of this Variable
	* @return The current value of this Variable
	*/
	Numbers_double CVariable::GetValue()
	{
		Numbers_double resultValue = 0;
		CheckError(m_pWrapper->m_WrapperTable.m_Variable_GetValue(m_pHandle, &resultValue));
		
		return resultValue;
	}
	
	/**
	* CVariable::SetValue - Set the numerical value of this Variable
	* @param[in] dValue - The new value of this Variable
	*/
	void CVariable::SetValue(const Numbers_double dValue)
	{
		CheckError(m_pWrapper->m_WrapperTable.m_Variable_SetValue(m_pHandle, dValue));
	}

} // namespace Numbers

#endif // __NUMBERS_CPPHEADER_DYNAMIC_CPP

