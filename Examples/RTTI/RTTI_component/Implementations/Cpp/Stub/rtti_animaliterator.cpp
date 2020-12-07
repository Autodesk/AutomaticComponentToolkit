/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CAnimalIterator

*/

#include "rtti_animaliterator.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.


using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CAnimalIterator 
**************************************************************************************************************************/

IAnimal * CAnimalIterator::GetNextAnimal()
{
	throw ERTTIInterfaceException(RTTI_ERROR_NOTIMPLEMENTED);
}

