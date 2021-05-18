/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CAnimal

*/

#include "rtti_animal.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.
#include <iostream>

using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CAnimal 
**************************************************************************************************************************/
CAnimal::CAnimal() = default;

CAnimal::CAnimal(std::string sName)
    : m_sName(sName)
{}

CAnimal::~CAnimal()
{
    std::cout << "Delete " << m_sName << std::endl;
}

std::string CAnimal::Name() 
{
    return m_sName;
}
