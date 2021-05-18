/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CTurtle

*/

#include "rtti_turtle.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.


using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CTurtle 
**************************************************************************************************************************/
CTurtle::CTurtle(std::string sName)
    : CAnimal(sName)
{
}
