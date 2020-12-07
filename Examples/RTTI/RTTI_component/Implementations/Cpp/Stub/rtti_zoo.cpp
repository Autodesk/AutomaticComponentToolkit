/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CZoo

*/

#include "rtti_zoo.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.
#include "rtti_animaliterator.hpp"

using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CZoo 
**************************************************************************************************************************/

std::vector<CAnimal *> &CZoo::Animals()
{
	return m_Animals;
}

IAnimalIterator * CZoo::Iterator()
{
	return new CAnimalIterator(m_Animals.begin(), m_Animals.end());
}
