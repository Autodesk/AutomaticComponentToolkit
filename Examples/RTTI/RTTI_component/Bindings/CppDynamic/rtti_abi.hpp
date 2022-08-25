/*++

Copyright (C) 2021 ADSK

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0.

Abstract: This is an autogenerated C++-Header file in order to allow an easy
 use of RTTI

Interface version: 1.0.0

*/

#ifndef __RTTI_HEADER_CPP
#define __RTTI_HEADER_CPP

#ifdef __RTTI_EXPORTS
#ifdef _WIN32
#define RTTI_DECLSPEC __declspec (dllexport)
#else // _WIN32
#define RTTI_DECLSPEC __attribute__((visibility("default")))
#endif // _WIN32
#else // __RTTI_EXPORTS
#define RTTI_DECLSPEC
#endif // __RTTI_EXPORTS

#include "rtti_types.hpp"


extern "C" {

/*************************************************************************************************************************
 Class definition for Base
**************************************************************************************************************************/

/**
* Get Class Type Id
*
* @param[in] pBase - Base instance.
* @param[out] pClassTypeId - Class type as a 64 bits integer
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_base_classtypeid(RTTI_Base pBase, RTTI_uint64 * pClassTypeId);

/*************************************************************************************************************************
 Class definition for Animal
**************************************************************************************************************************/

/**
* Get the name of the animal
*
* @param[in] pAnimal - Animal instance.
* @param[in] nResultBufferSize - size of the buffer (including trailing 0)
* @param[out] pResultNeededChars - will be filled with the count of the written bytes, or needed buffer size.
* @param[out] pResultBuffer -  buffer of , may be NULL
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_animal_name(RTTI_Animal pAnimal, const RTTI_uint32 nResultBufferSize, RTTI_uint32* pResultNeededChars, char * pResultBuffer);

/*************************************************************************************************************************
 Class definition for Mammal
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class definition for Reptile
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class definition for Giraffe
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class definition for Tiger
**************************************************************************************************************************/

/**
* Roar like a tiger
*
* @param[in] pTiger - Tiger instance.
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_tiger_roar(RTTI_Tiger pTiger);

/*************************************************************************************************************************
 Class definition for Snake
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class definition for Turtle
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class definition for AnimalIterator
**************************************************************************************************************************/

/**
* Return next animal
*
* @param[in] pAnimalIterator - AnimalIterator instance.
* @param[out] pAnimal - 
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_animaliterator_getnextanimal(RTTI_AnimalIterator pAnimalIterator, RTTI_Animal * pAnimal);

/*************************************************************************************************************************
 Class definition for Zoo
**************************************************************************************************************************/

/**
* Return an iterator over all zoo animals
*
* @param[in] pZoo - Zoo instance.
* @param[out] pIterator - 
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_zoo_iterator(RTTI_Zoo pZoo, RTTI_AnimalIterator * pIterator);

/*************************************************************************************************************************
 Global functions
**************************************************************************************************************************/

/**
* retrieves the binary version of this library.
*
* @param[out] pMajor - returns the major version of this library
* @param[out] pMinor - returns the minor version of this library
* @param[out] pMicro - returns the micro version of this library
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_getversion(RTTI_uint32 * pMajor, RTTI_uint32 * pMinor, RTTI_uint32 * pMicro);

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
RTTI_DECLSPEC RTTIResult rtti_getlasterror(RTTI_Base pInstance, const RTTI_uint32 nErrorMessageBufferSize, RTTI_uint32* pErrorMessageNeededChars, char * pErrorMessageBuffer, bool * pHasError);

/**
* Releases shared ownership of an Instance
*
* @param[in] pInstance - Instance Handle
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_releaseinstance(RTTI_Base pInstance);

/**
* Acquires shared ownership of an Instance
*
* @param[in] pInstance - Instance Handle
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_acquireinstance(RTTI_Base pInstance);

/**
* Injects an imported component for usage within this component
*
* @param[in] pNameSpace - NameSpace of the injected component
* @param[in] pSymbolAddressMethod - Address of the SymbolAddressMethod of the injected component
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_injectcomponent(const char * pNameSpace, RTTI_pvoid pSymbolAddressMethod);

/**
* Returns the address of the SymbolLookupMethod
*
* @param[out] pSymbolLookupMethod - Address of the SymbolAddressMethod
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_getsymbollookupmethod(RTTI_pvoid * pSymbolLookupMethod);

/**
* Create a new zoo with animals
*
* @param[out] pInstance - 
* @return error code or 0 (success)
*/
RTTI_DECLSPEC RTTIResult rtti_createzoo(RTTI_Zoo * pInstance);

}

#endif // __RTTI_HEADER_CPP
