/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0-develop5.

Abstract: This is an autogenerated C++-Header file with basic types in
order to allow an easy use of Prime Numbers Library

Interface version: 1.2.0

*/

#ifndef __LIBPRIMES_TYPES_HEADER_CPP
#define __LIBPRIMES_TYPES_HEADER_CPP

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

namespace LibPrimes {

  /*************************************************************************************************************************
   Declaration of enums
  **************************************************************************************************************************/
  
  enum class eProgressIdentifier : LibPrimes_int32 {
    NoOp = -1,
    Starting = 0,
    Running = 1,
    Done = 2
  };
  
  /*************************************************************************************************************************
   Declaration of structs
  **************************************************************************************************************************/
  
  #pragma pack (1)
  
  typedef struct {
      LibPrimes_uint64 m_Prime;
      LibPrimes_uint32 m_Multiplicity;
  } sPrimeFactor;
  
  #pragma pack ()
  
  /*************************************************************************************************************************
   Declaration of function pointers 
  **************************************************************************************************************************/
  
  /**
  * ProgressCallback - Callback to report calculation progress and query whether it should be aborted
  *
  * @param[in] fProgressPercentage - How far has the calculation progressed?
  * @param[in] eTheProgressIdentifier - What is the current state the calculation is in?
  * @param[out] pShouldAbort - Should the calculation be aborted?
  */
  typedef void(*ProgressCallback)(LibPrimes_single, LibPrimes::eProgressIdentifier, bool*);
  
} // namespace LibPrimes;

// define legacy C-names for enums, structs and function types
typedef LibPrimes::eProgressIdentifier eLibPrimesProgressIdentifier;
typedef LibPrimes::sPrimeFactor sLibPrimesPrimeFactor;
typedef LibPrimes::ProgressCallback LibPrimesProgressCallback;

#endif // __LIBPRIMES_TYPES_HEADER_CPP
