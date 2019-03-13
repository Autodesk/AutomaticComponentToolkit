/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0-develop.

Abstract: This is an autogenerated plain C Header file with basic types in
order to allow an easy use of Prime Numbers Library

Interface version: 1.2.0

*/

#ifndef __LIBPRIMES_TYPES_HEADER
#define __LIBPRIMES_TYPES_HEADER

/*************************************************************************************************************************
 Scalar types definition
**************************************************************************************************************************/

#ifdef LIBPRIMES_USELEGACYINTEGERTYPES

typedef unsigned char LibPrimes_uint8;
typedef unsigned short LibPrimes_uint16 ;
typedef unsigned int LibPrimes_uint32;
typedef unsigned long long LibPrimes_uint64;
typedef char LibPrimes_int8;
typedef short LibPrimes_int16;
typedef int LibPrimes_int32;
typedef long long LibPrimes_int64;

#else // LIBPRIMES_USELEGACYINTEGERTYPES

#include <stdint.h>

typedef uint8_t LibPrimes_uint8;
typedef uint16_t LibPrimes_uint16;
typedef uint32_t LibPrimes_uint32;
typedef uint64_t LibPrimes_uint64;
typedef int8_t LibPrimes_int8;
typedef int16_t LibPrimes_int16;
typedef int32_t LibPrimes_int32;
typedef int64_t LibPrimes_int64 ;

#endif // LIBPRIMES_USELEGACYINTEGERTYPES

typedef float LibPrimes_single;
typedef double LibPrimes_double;

/*************************************************************************************************************************
 General type definitions
**************************************************************************************************************************/

typedef LibPrimes_int32 LibPrimesResult;
typedef void * LibPrimesHandle;
typedef void * LibPrimes_pvoid;

/*************************************************************************************************************************
 Version for LibPrimes
**************************************************************************************************************************/

#define LIBPRIMES_VERSION_MAJOR 1
#define LIBPRIMES_VERSION_MINOR 2
#define LIBPRIMES_VERSION_MICRO 0
#define LIBPRIMES_VERSION_PRERELEASEINFO "alpha"
#define LIBPRIMES_VERSION_BUILDINFO "23"

/*************************************************************************************************************************
 Error constants for LibPrimes
**************************************************************************************************************************/

#define LIBPRIMES_SUCCESS 0
#define LIBPRIMES_ERROR_NOTIMPLEMENTED 1
#define LIBPRIMES_ERROR_INVALIDPARAM 2
#define LIBPRIMES_ERROR_INVALIDCAST 3
#define LIBPRIMES_ERROR_BUFFERTOOSMALL 4
#define LIBPRIMES_ERROR_GENERICEXCEPTION 5
#define LIBPRIMES_ERROR_COULDNOTLOADLIBRARY 6
#define LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT 7
#define LIBPRIMES_ERROR_INCOMPATIBLEBINARYVERSION 8
#define LIBPRIMES_ERROR_NORESULTAVAILABLE 9
#define LIBPRIMES_ERROR_CALCULATIONABORTED 10

/*************************************************************************************************************************
 Declaration of handle classes 
**************************************************************************************************************************/

typedef LibPrimesHandle LibPrimes_Base;
typedef LibPrimesHandle LibPrimes_Calculator;
typedef LibPrimesHandle LibPrimes_FactorizationCalculator;
typedef LibPrimesHandle LibPrimes_SieveCalculator;

/*************************************************************************************************************************
 Declaration of structs
**************************************************************************************************************************/

#pragma pack (1)

typedef struct {
    LibPrimes_uint64 m_Prime;
    LibPrimes_uint32 m_Multiplicity;
} sLibPrimesPrimeFactor;

#pragma pack ()

/*************************************************************************************************************************
 Declaration of function pointers 
**************************************************************************************************************************/

/**
* LibPrimesProgressCallback - Callback to report calculation progress and query whether it should be aborted
*
* @param[in] fProgressPercentage - How far has the calculation progressed?
* @param[out] pShouldAbort - Should the calculation be aborted?
*/
typedef void(*LibPrimesProgressCallback)(LibPrimes_single, bool*);

#endif // __LIBPRIMES_TYPES_HEADER
