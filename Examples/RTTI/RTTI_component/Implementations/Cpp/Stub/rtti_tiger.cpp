/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CTiger

*/

#include "rtti_tiger.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.
#include <iostream>

using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CTiger 
**************************************************************************************************************************/
CTiger::CTiger(std::string sName)
    : CAnimal(sName)
{
}

void CTiger::Roar()
{
	std::cout << "ROAAAAARRRRR!!" << std::endl;
}

