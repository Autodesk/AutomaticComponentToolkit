/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CAnimalIterator

*/

#include "rtti_animaliterator.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.
#include "rtti_animal.hpp"

using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CAnimalIterator 
**************************************************************************************************************************/
CAnimalIterator::CAnimalIterator(std::vector<CAnimal*>::iterator begin, std::vector<CAnimal*>::iterator end)
    : m_Current(begin)
    , m_End(end)
{
}

IAnimal * CAnimalIterator::GetNextAnimal()
{
	if (m_Current != m_End) {
		auto i = *(m_Current++);
		i->IncRefCount();
		return i;
	} else {
		return nullptr;
	}
}
