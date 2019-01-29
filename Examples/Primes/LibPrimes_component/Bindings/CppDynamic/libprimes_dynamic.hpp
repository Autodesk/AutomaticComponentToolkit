/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.4.0.

Abstract: This is an autogenerated C++ Header file in order to allow an easy
 use of Prime Numbers Library

Interface version: 1.2.0

*/

#ifndef __LIBPRIMES_DYNAMICCPPHEADER
#define __LIBPRIMES_DYNAMICCPPHEADER

#include "libprimes_types.h"
#include "libprimes_dynamic.h"

#ifdef WIN32
#include <Windows.h>
#else // WIN32
#include <dlfcn.h>
#endif // WIN32
#include <string>
#include <memory>
#include <vector>
#include <exception>

namespace LibPrimes {

/*************************************************************************************************************************
 Forward Declaration of all classes 
**************************************************************************************************************************/

class CLibPrimesBaseClass;
class CLibPrimesWrapper;
class CLibPrimesCalculator;
class CLibPrimesFactorizationCalculator;
class CLibPrimesSieveCalculator;

/*************************************************************************************************************************
 Declaration of shared pointer types 
**************************************************************************************************************************/

typedef std::shared_ptr<CLibPrimesBaseClass> PLibPrimesBaseClass;
typedef std::shared_ptr<CLibPrimesWrapper> PLibPrimesWrapper;
typedef std::shared_ptr<CLibPrimesCalculator> PLibPrimesCalculator;
typedef std::shared_ptr<CLibPrimesFactorizationCalculator> PLibPrimesFactorizationCalculator;
typedef std::shared_ptr<CLibPrimesSieveCalculator> PLibPrimesSieveCalculator;

/*************************************************************************************************************************
 Class ELibPrimesException 
**************************************************************************************************************************/
class ELibPrimesException : public std::exception {
protected:
	/**
	* Error code for the Exception.
	*/
	LibPrimesResult m_errorCode;
	/**
	* Error message for the Exception.
	*/
	std::string m_errorMessage;

public:
	/**
	* Exception Constructor.
	*/
	ELibPrimesException (LibPrimesResult errorCode)
		: m_errorMessage("LibPrimes Error " + std::to_string (errorCode))
	{
		m_errorCode = errorCode;
	}

	/**
	* Returns error code
	*/
	LibPrimesResult getErrorCode ()
	{
		return m_errorCode;
	}

	/**
	* Returns error message
	*/
	const char* what () const noexcept
	{
		return m_errorMessage.c_str();
	}

};

/*************************************************************************************************************************
 Class CLibPrimesInputVector
**************************************************************************************************************************/
template <typename T>
class CLibPrimesInputVector {
private:
	
	const T* m_data;
	size_t m_size;
	
public:
	
	CLibPrimesInputVector( const std::vector<T>& vec)
		: m_data( vec.data() ), m_size( vec.size() )
	{
	}
	
	CLibPrimesInputVector( const T* in_data, size_t in_size)
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

/*************************************************************************************************************************
 Class CLibPrimesWrapper 
**************************************************************************************************************************/
class CLibPrimesWrapper {
public:
	
	CLibPrimesWrapper (const std::string &sFileName)
	{
		CheckError (nullptr, initWrapperTable (&m_WrapperTable));
		CheckError (nullptr, loadWrapperTable (&m_WrapperTable, sFileName.c_str ()));
		
		CheckError(nullptr, checkBinaryVersion());
	}
	
	static PLibPrimesWrapper loadLibrary (const std::string &sFileName)
	{
		return std::make_shared<CLibPrimesWrapper> (sFileName);
	}
	
	~CLibPrimesWrapper ()
	{
		releaseWrapperTable (&m_WrapperTable);
	}
	
	void CheckError(LibPrimesHandle handle, LibPrimesResult nResult)
	{
		if (nResult != 0) 
			throw ELibPrimesException (nResult);
	}
	

	inline void ReleaseInstance (CLibPrimesBaseClass * pInstance);
	inline void GetLibraryVersion (LibPrimes_uint32 & nMajor, LibPrimes_uint32 & nMinor, LibPrimes_uint32 & nMicro);
	inline PLibPrimesFactorizationCalculator CreateFactorizationCalculator ();
	inline PLibPrimesSieveCalculator CreateSieveCalculator ();
	inline void SetJournal (const std::string & sFileName);

private:
	sLibPrimesDynamicWrapperTable m_WrapperTable;

	LibPrimesResult checkBinaryVersion()
	{
		LibPrimes_uint32 nMajor, nMinor, nMicro;
		GetLibraryVersion(nMajor, nMinor, nMicro);
		if ( (nMajor != LIBPRIMES_VERSION_MAJOR) || (nMinor < LIBPRIMES_VERSION_MINOR) ) {
			return LIBPRIMES_ERROR_INCOMPATIBLEBINARYVERSION;
		}
		return LIBPRIMES_SUCCESS;
	}
	LibPrimesResult initWrapperTable (sLibPrimesDynamicWrapperTable * pWrapperTable);
	LibPrimesResult releaseWrapperTable (sLibPrimesDynamicWrapperTable * pWrapperTable);
	LibPrimesResult loadWrapperTable (sLibPrimesDynamicWrapperTable * pWrapperTable, const char * pLibraryFileName);

	friend class CLibPrimesCalculator;
	friend class CLibPrimesFactorizationCalculator;
	friend class CLibPrimesSieveCalculator;

};

/*************************************************************************************************************************
 Class CLibPrimesBaseClass 
**************************************************************************************************************************/
class CLibPrimesBaseClass {
protected:
	/* Wrapper Object that created the class..*/
	CLibPrimesWrapper * m_pWrapper;
	/* Handle to Instance in library*/
	LibPrimesHandle m_pHandle;

	/* Checks for an Error code and raises Exceptions */
	void CheckError(LibPrimesResult nResult)
	{
		if (m_pWrapper != nullptr)
			m_pWrapper->CheckError (m_pHandle, nResult);
	}

public:

	/**
	* CLibPrimesBaseClass::CLibPrimesBaseClass - Constructor for Base class.
	*/
	CLibPrimesBaseClass(CLibPrimesWrapper * pWrapper, LibPrimesHandle pHandle)
		: m_pWrapper (pWrapper), m_pHandle (pHandle)
	{
	}

	/**
	* CLibPrimesBaseClass::~CLibPrimesBaseClass - Destructor for Base class.
	*/
	virtual ~CLibPrimesBaseClass()
	{
		if (m_pWrapper != nullptr)
			m_pWrapper->ReleaseInstance (this);
		m_pWrapper = nullptr;
	}

	/**
	* CLibPrimesBaseClass::GetHandle - Returns handle to instance.
	*/
	LibPrimesHandle GetHandle()
	{
		return m_pHandle;
	}
};
	
	
/*************************************************************************************************************************
 Class CLibPrimesCalculator 
**************************************************************************************************************************/
class CLibPrimesCalculator : public CLibPrimesBaseClass {
public:
	
	/**
	* CLibPrimesCalculator::CLibPrimesCalculator - Constructor for Calculator class.
	*/
	CLibPrimesCalculator (CLibPrimesWrapper * pWrapper, LibPrimesHandle pHandle)
		: CLibPrimesBaseClass (pWrapper, pHandle)
	{
	}
	
	inline LibPrimes_uint64 GetValue ();
	inline void SetValue (const LibPrimes_uint64 nValue);
	inline void Calculate ();
	inline void SetProgressCallback (const LibPrimesProgressCallback pProgressCallback);
};
	
/*************************************************************************************************************************
 Class CLibPrimesFactorizationCalculator 
**************************************************************************************************************************/
class CLibPrimesFactorizationCalculator : public CLibPrimesCalculator {
public:
	
	/**
	* CLibPrimesFactorizationCalculator::CLibPrimesFactorizationCalculator - Constructor for FactorizationCalculator class.
	*/
	CLibPrimesFactorizationCalculator (CLibPrimesWrapper * pWrapper, LibPrimesHandle pHandle)
		: CLibPrimesCalculator (pWrapper, pHandle)
	{
	}
	
	inline void GetPrimeFactors (std::vector<sLibPrimesPrimeFactor> & PrimeFactorsBuffer);
};
	
/*************************************************************************************************************************
 Class CLibPrimesSieveCalculator 
**************************************************************************************************************************/
class CLibPrimesSieveCalculator : public CLibPrimesCalculator {
public:
	
	/**
	* CLibPrimesSieveCalculator::CLibPrimesSieveCalculator - Constructor for SieveCalculator class.
	*/
	CLibPrimesSieveCalculator (CLibPrimesWrapper * pWrapper, LibPrimesHandle pHandle)
		: CLibPrimesCalculator (pWrapper, pHandle)
	{
	}
	
	inline void GetPrimes (std::vector<LibPrimes_uint64> & PrimesBuffer);
};
	
	/**
	* CLibPrimesWrapper::ReleaseInstance - Releases the memory of an Instance
	* @param[in] pInstance - Instance Handle
	*/
	inline void CLibPrimesWrapper::ReleaseInstance (CLibPrimesBaseClass * pInstance)
	{
		LibPrimesHandle hInstance = nullptr;
		if (pInstance != nullptr) {
			hInstance = pInstance->GetHandle ();
		};
		CheckError (nullptr, m_WrapperTable.m_ReleaseInstance (hInstance) );
	}
	
	/**
	* CLibPrimesWrapper::GetLibraryVersion - retrieves the current version of the library.
	* @param[out] nMajor - returns the major version of the library
	* @param[out] nMinor - returns the minor version of the library
	* @param[out] nMicro - returns the micro version of the library
	*/
	inline void CLibPrimesWrapper::GetLibraryVersion (LibPrimes_uint32 & nMajor, LibPrimes_uint32 & nMinor, LibPrimes_uint32 & nMicro)
	{
		CheckError (nullptr, m_WrapperTable.m_GetLibraryVersion (&nMajor, &nMinor, &nMicro) );
	}
	
	/**
	* CLibPrimesWrapper::CreateFactorizationCalculator - Creates a new FactorizationCalculator instance
	* @return New FactorizationCalculator instance
	*/
	inline PLibPrimesFactorizationCalculator CLibPrimesWrapper::CreateFactorizationCalculator ()
	{
		LibPrimesHandle hInstance = nullptr;
		CheckError (nullptr, m_WrapperTable.m_CreateFactorizationCalculator (&hInstance) );
		return std::make_shared<CLibPrimesFactorizationCalculator> (this, hInstance);
	}
	
	/**
	* CLibPrimesWrapper::CreateSieveCalculator - Creates a new SieveCalculator instance
	* @return New SieveCalculator instance
	*/
	inline PLibPrimesSieveCalculator CLibPrimesWrapper::CreateSieveCalculator ()
	{
		LibPrimesHandle hInstance = nullptr;
		CheckError (nullptr, m_WrapperTable.m_CreateSieveCalculator (&hInstance) );
		return std::make_shared<CLibPrimesSieveCalculator> (this, hInstance);
	}
	
	/**
	* CLibPrimesWrapper::SetJournal - Handles Library Journaling
	* @param[in] sFileName - Journal FileName
	*/
	inline void CLibPrimesWrapper::SetJournal (const std::string & sFileName)
	{
		CheckError (nullptr, m_WrapperTable.m_SetJournal (sFileName.c_str()) );
	}

	inline LibPrimesResult CLibPrimesWrapper::initWrapperTable (sLibPrimesDynamicWrapperTable * pWrapperTable)
	{
		if (pWrapperTable == nullptr)
			return LIBPRIMES_ERROR_INVALIDPARAM;
		
		pWrapperTable->m_LibraryHandle = nullptr;
		pWrapperTable->m_Calculator_GetValue = nullptr;
		pWrapperTable->m_Calculator_SetValue = nullptr;
		pWrapperTable->m_Calculator_Calculate = nullptr;
		pWrapperTable->m_Calculator_SetProgressCallback = nullptr;
		pWrapperTable->m_FactorizationCalculator_GetPrimeFactors = nullptr;
		pWrapperTable->m_SieveCalculator_GetPrimes = nullptr;
		pWrapperTable->m_ReleaseInstance = nullptr;
		pWrapperTable->m_GetLibraryVersion = nullptr;
		pWrapperTable->m_CreateFactorizationCalculator = nullptr;
		pWrapperTable->m_CreateSieveCalculator = nullptr;
		pWrapperTable->m_SetJournal = nullptr;
		
		return LIBPRIMES_SUCCESS;
	}

	inline LibPrimesResult CLibPrimesWrapper::releaseWrapperTable (sLibPrimesDynamicWrapperTable * pWrapperTable)
	{
		if (pWrapperTable == nullptr)
			return LIBPRIMES_ERROR_INVALIDPARAM;
		
		if (pWrapperTable->m_LibraryHandle != nullptr) {
		#ifdef WIN32
			HMODULE hModule = (HMODULE) pWrapperTable->m_LibraryHandle;
			FreeLibrary (hModule);
		#else // WIN32
			dlclose (pWrapperTable->m_LibraryHandle);
		#endif // WIN32
			return initWrapperTable (pWrapperTable);
		}
		
		return LIBPRIMES_SUCCESS;
	}

	inline LibPrimesResult CLibPrimesWrapper::loadWrapperTable (sLibPrimesDynamicWrapperTable * pWrapperTable, const char * pLibraryFileName)
	{
		if (pWrapperTable == nullptr)
			return LIBPRIMES_ERROR_INVALIDPARAM;
		if (pLibraryFileName == nullptr)
			return LIBPRIMES_ERROR_INVALIDPARAM;
		
		#ifdef WIN32
		HMODULE hLibrary = LoadLibraryExA(pLibraryFileName, nullptr, LOAD_LIBRARY_SEARCH_DLL_LOAD_DIR);
		if (hLibrary == 0) 
			return LIBPRIMES_ERROR_COULDNOTLOADLIBRARY;
		#else // WIN32
		void* hLibrary = dlopen (pLibraryFileName, RTLD_LAZY);
		if (hLibrary == 0) 
			return LIBPRIMES_ERROR_COULDNOTLOADLIBRARY;
		dlerror();
		#endif // WIN32
		
		#ifdef WIN32
		pWrapperTable->m_Calculator_GetValue = (PLibPrimesCalculator_GetValuePtr) GetProcAddress (hLibrary, "libprimes_calculator_getvalue");
		#else // WIN32
		pWrapperTable->m_Calculator_GetValue = (PLibPrimesCalculator_GetValuePtr) dlsym (hLibrary, "libprimes_calculator_getvalue");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_Calculator_GetValue == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_Calculator_SetValue = (PLibPrimesCalculator_SetValuePtr) GetProcAddress (hLibrary, "libprimes_calculator_setvalue");
		#else // WIN32
		pWrapperTable->m_Calculator_SetValue = (PLibPrimesCalculator_SetValuePtr) dlsym (hLibrary, "libprimes_calculator_setvalue");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_Calculator_SetValue == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_Calculator_Calculate = (PLibPrimesCalculator_CalculatePtr) GetProcAddress (hLibrary, "libprimes_calculator_calculate");
		#else // WIN32
		pWrapperTable->m_Calculator_Calculate = (PLibPrimesCalculator_CalculatePtr) dlsym (hLibrary, "libprimes_calculator_calculate");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_Calculator_Calculate == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_Calculator_SetProgressCallback = (PLibPrimesCalculator_SetProgressCallbackPtr) GetProcAddress (hLibrary, "libprimes_calculator_setprogresscallback");
		#else // WIN32
		pWrapperTable->m_Calculator_SetProgressCallback = (PLibPrimesCalculator_SetProgressCallbackPtr) dlsym (hLibrary, "libprimes_calculator_setprogresscallback");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_Calculator_SetProgressCallback == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_FactorizationCalculator_GetPrimeFactors = (PLibPrimesFactorizationCalculator_GetPrimeFactorsPtr) GetProcAddress (hLibrary, "libprimes_factorizationcalculator_getprimefactors");
		#else // WIN32
		pWrapperTable->m_FactorizationCalculator_GetPrimeFactors = (PLibPrimesFactorizationCalculator_GetPrimeFactorsPtr) dlsym (hLibrary, "libprimes_factorizationcalculator_getprimefactors");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_FactorizationCalculator_GetPrimeFactors == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_SieveCalculator_GetPrimes = (PLibPrimesSieveCalculator_GetPrimesPtr) GetProcAddress (hLibrary, "libprimes_sievecalculator_getprimes");
		#else // WIN32
		pWrapperTable->m_SieveCalculator_GetPrimes = (PLibPrimesSieveCalculator_GetPrimesPtr) dlsym (hLibrary, "libprimes_sievecalculator_getprimes");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_SieveCalculator_GetPrimes == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_ReleaseInstance = (PLibPrimesReleaseInstancePtr) GetProcAddress (hLibrary, "libprimes_releaseinstance");
		#else // WIN32
		pWrapperTable->m_ReleaseInstance = (PLibPrimesReleaseInstancePtr) dlsym (hLibrary, "libprimes_releaseinstance");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_ReleaseInstance == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_GetLibraryVersion = (PLibPrimesGetLibraryVersionPtr) GetProcAddress (hLibrary, "libprimes_getlibraryversion");
		#else // WIN32
		pWrapperTable->m_GetLibraryVersion = (PLibPrimesGetLibraryVersionPtr) dlsym (hLibrary, "libprimes_getlibraryversion");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_GetLibraryVersion == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_CreateFactorizationCalculator = (PLibPrimesCreateFactorizationCalculatorPtr) GetProcAddress (hLibrary, "libprimes_createfactorizationcalculator");
		#else // WIN32
		pWrapperTable->m_CreateFactorizationCalculator = (PLibPrimesCreateFactorizationCalculatorPtr) dlsym (hLibrary, "libprimes_createfactorizationcalculator");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_CreateFactorizationCalculator == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_CreateSieveCalculator = (PLibPrimesCreateSieveCalculatorPtr) GetProcAddress (hLibrary, "libprimes_createsievecalculator");
		#else // WIN32
		pWrapperTable->m_CreateSieveCalculator = (PLibPrimesCreateSieveCalculatorPtr) dlsym (hLibrary, "libprimes_createsievecalculator");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_CreateSieveCalculator == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		#ifdef WIN32
		pWrapperTable->m_SetJournal = (PLibPrimesSetJournalPtr) GetProcAddress (hLibrary, "libprimes_setjournal");
		#else // WIN32
		pWrapperTable->m_SetJournal = (PLibPrimesSetJournalPtr) dlsym (hLibrary, "libprimes_setjournal");
		dlerror();
		#endif // WIN32
		if (pWrapperTable->m_SetJournal == nullptr)
			return LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT;
		
		pWrapperTable->m_LibraryHandle = hLibrary;
		return LIBPRIMES_SUCCESS;
	}
	
	
	/**
	 * Method definitions for class CLibPrimesCalculator
	 */
	
	LibPrimes_uint64 CLibPrimesCalculator::GetValue ()
	{
		LibPrimes_uint64 resultValue = 0;
		CheckError ( m_pWrapper->m_WrapperTable.m_Calculator_GetValue (m_pHandle, &resultValue) );
		return resultValue;
	}
	
	void CLibPrimesCalculator::SetValue (const LibPrimes_uint64 nValue)
	{
		CheckError ( m_pWrapper->m_WrapperTable.m_Calculator_SetValue (m_pHandle, nValue) );
	}
	
	void CLibPrimesCalculator::Calculate ()
	{
		CheckError ( m_pWrapper->m_WrapperTable.m_Calculator_Calculate (m_pHandle) );
	}
	
	void CLibPrimesCalculator::SetProgressCallback (const LibPrimesProgressCallback pProgressCallback)
	{
		CheckError ( m_pWrapper->m_WrapperTable.m_Calculator_SetProgressCallback (m_pHandle, pProgressCallback) );
	}
	
	/**
	 * Method definitions for class CLibPrimesFactorizationCalculator
	 */
	
	void CLibPrimesFactorizationCalculator::GetPrimeFactors (std::vector<sLibPrimesPrimeFactor> & PrimeFactorsBuffer)
	{
		LibPrimes_uint64 elementsNeededPrimeFactors = 0;
		LibPrimes_uint64 elementsWrittenPrimeFactors = 0;
		CheckError ( m_pWrapper->m_WrapperTable.m_FactorizationCalculator_GetPrimeFactors (m_pHandle, 0, &elementsNeededPrimeFactors, nullptr) );
		PrimeFactorsBuffer.resize(elementsNeededPrimeFactors);
		CheckError ( m_pWrapper->m_WrapperTable.m_FactorizationCalculator_GetPrimeFactors (m_pHandle, elementsNeededPrimeFactors, &elementsWrittenPrimeFactors, PrimeFactorsBuffer.data()) );
	}
	
	/**
	 * Method definitions for class CLibPrimesSieveCalculator
	 */
	
	void CLibPrimesSieveCalculator::GetPrimes (std::vector<LibPrimes_uint64> & PrimesBuffer)
	{
		LibPrimes_uint64 elementsNeededPrimes = 0;
		LibPrimes_uint64 elementsWrittenPrimes = 0;
		CheckError ( m_pWrapper->m_WrapperTable.m_SieveCalculator_GetPrimes (m_pHandle, 0, &elementsNeededPrimes, nullptr) );
		PrimesBuffer.resize(elementsNeededPrimes);
		CheckError ( m_pWrapper->m_WrapperTable.m_SieveCalculator_GetPrimes (m_pHandle, elementsNeededPrimes, &elementsWrittenPrimes, PrimesBuffer.data()) );
	}

} // namespace LibPrimes

#endif // __LIBPRIMES_DYNAMICCPPHEADER

