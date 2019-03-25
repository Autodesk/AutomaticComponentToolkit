/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0-develop2.

Abstract: This is an autogenerated C++-Header file in order to allow an easy
 use of Prime Numbers Library

Interface version: 1.2.0

*/

#ifndef __LIBPRIMES_DYNAMICHEADER_CPP
#define __LIBPRIMES_DYNAMICHEADER_CPP

#include "libprimes_types.hpp"


/*************************************************************************************************************************
 Class definition for Base
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class definition for Calculator
**************************************************************************************************************************/

/**
* Returns the current value of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @param[out] pValue - The current value of this Calculator
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCalculator_GetValuePtr) (LibPrimes_Calculator pCalculator, LibPrimes_uint64 * pValue);

/**
* Returns the current value of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @param[out] pValue - The current value of this Calculator
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCalculator_GetSelfPtr) (LibPrimes_Calculator pCalculator, LibPrimes_Calculator * pValue);

/**
* Sets the value to be factorized
*
* @param[in] pCalculator - Calculator instance.
* @param[in] nValue - The value to be factorized
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCalculator_SetValuePtr) (LibPrimes_Calculator pCalculator, LibPrimes_uint64 nValue);

/**
* Performs the specific calculation of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCalculator_CalculatePtr) (LibPrimes_Calculator pCalculator);

/**
* Sets the progress callback function
*
* @param[in] pCalculator - Calculator instance.
* @param[in] pProgressCallback - The progress callback
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCalculator_SetProgressCallbackPtr) (LibPrimes_Calculator pCalculator, LibPrimes::ProgressCallback pProgressCallback);

/*************************************************************************************************************************
 Class definition for FactorizationCalculator
**************************************************************************************************************************/

/**
* Returns the prime factors of this number (without multiplicity)
*
* @param[in] pFactorizationCalculator - FactorizationCalculator instance.
* @param[in] nPrimeFactorsBufferSize - Number of elements in buffer
* @param[out] pPrimeFactorsNeededCount - will be filled with the count of the written elements, or needed buffer size.
* @param[out] pPrimeFactorsBuffer - PrimeFactor buffer of The prime factors of this number
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesFactorizationCalculator_GetPrimeFactorsPtr) (LibPrimes_FactorizationCalculator pFactorizationCalculator, const LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, LibPrimes::sPrimeFactor * pPrimeFactorsBuffer);

/*************************************************************************************************************************
 Class definition for SieveCalculator
**************************************************************************************************************************/

/**
* Returns all prime numbers lower or equal to the sieve's value
*
* @param[in] pSieveCalculator - SieveCalculator instance.
* @param[in] nPrimesBufferSize - Number of elements in buffer
* @param[out] pPrimesNeededCount - will be filled with the count of the written elements, or needed buffer size.
* @param[out] pPrimesBuffer - uint64 buffer of The primes lower or equal to the sieve's value
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesSieveCalculator_GetPrimesPtr) (LibPrimes_SieveCalculator pSieveCalculator, const LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer);

/*************************************************************************************************************************
 Global functions
**************************************************************************************************************************/

/**
* Returns the last error recorded on this object
*
* @param[in] pInstance - Instance Handle
* @param[in] nErrorMessageBufferSize - size of the buffer (including trailing 0)
* @param[out] pErrorMessageNeededChars - will be filled with the count of the written bytes, or needed buffer size.
* @param[out] pErrorMessageBuffer -  buffer of Message of the last error, may be NULL
* @param[out] pHasError - Is there a last error to query
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesGetLastErrorPtr) (LibPrimes_Base pInstance, const LibPrimes_uint32 nErrorMessageBufferSize, LibPrimes_uint32* pErrorMessageNeededChars, char * pErrorMessageBuffer, bool * pHasError);

/**
* Releases the memory of an Instance
*
* @param[in] pInstance - Instance Handle
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesReleaseInstancePtr) (LibPrimes_Base pInstance);

/**
* retrieves the binary version of this library.
*
* @param[out] pMajor - returns the major version of this library
* @param[out] pMinor - returns the minor version of this library
* @param[out] pMicro - returns the micro version of this library
* @param[in] nPreReleaseInfoBufferSize - size of the buffer (including trailing 0)
* @param[out] pPreReleaseInfoNeededChars - will be filled with the count of the written bytes, or needed buffer size.
* @param[out] pPreReleaseInfoBuffer -  buffer of returns pre-release info of this library (if this is a pre-release binary), may be NULL
* @param[in] nBuildInfoBufferSize - size of the buffer (including trailing 0)
* @param[out] pBuildInfoNeededChars - will be filled with the count of the written bytes, or needed buffer size.
* @param[out] pBuildInfoBuffer -  buffer of returns build-information of this library (optional), may be NULL
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesGetLibraryVersionPtr) (LibPrimes_uint32 * pMajor, LibPrimes_uint32 * pMinor, LibPrimes_uint32 * pMicro, const LibPrimes_uint32 nPreReleaseInfoBufferSize, LibPrimes_uint32* pPreReleaseInfoNeededChars, char * pPreReleaseInfoBuffer, const LibPrimes_uint32 nBuildInfoBufferSize, LibPrimes_uint32* pBuildInfoNeededChars, char * pBuildInfoBuffer);

/**
* Creates a new FactorizationCalculator instance
*
* @param[out] pInstance - New FactorizationCalculator instance
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCreateFactorizationCalculatorPtr) (LibPrimes_FactorizationCalculator * pInstance);

/**
* Creates a new SieveCalculator instance
*
* @param[out] pInstance - New SieveCalculator instance
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesCreateSieveCalculatorPtr) (LibPrimes_SieveCalculator * pInstance);

/**
* Handles Library Journaling
*
* @param[in] pFileName - Journal FileName
* @return error code or 0 (success)
*/
typedef LibPrimesResult (*PLibPrimesSetJournalPtr) (const char * pFileName);

/*************************************************************************************************************************
 Function Table Structure
**************************************************************************************************************************/

typedef struct {
	void * m_LibraryHandle;
	PLibPrimesCalculator_GetValuePtr m_Calculator_GetValue;
	PLibPrimesCalculator_GetSelfPtr m_Calculator_GetSelf;
	PLibPrimesCalculator_SetValuePtr m_Calculator_SetValue;
	PLibPrimesCalculator_CalculatePtr m_Calculator_Calculate;
	PLibPrimesCalculator_SetProgressCallbackPtr m_Calculator_SetProgressCallback;
	PLibPrimesFactorizationCalculator_GetPrimeFactorsPtr m_FactorizationCalculator_GetPrimeFactors;
	PLibPrimesSieveCalculator_GetPrimesPtr m_SieveCalculator_GetPrimes;
	PLibPrimesGetLastErrorPtr m_GetLastError;
	PLibPrimesReleaseInstancePtr m_ReleaseInstance;
	PLibPrimesGetLibraryVersionPtr m_GetLibraryVersion;
	PLibPrimesCreateFactorizationCalculatorPtr m_CreateFactorizationCalculator;
	PLibPrimesCreateSieveCalculatorPtr m_CreateSieveCalculator;
	PLibPrimesSetJournalPtr m_SetJournal;
} sLibPrimesDynamicWrapperTable;

#endif // __LIBPRIMES_DYNAMICHEADER_CPP

