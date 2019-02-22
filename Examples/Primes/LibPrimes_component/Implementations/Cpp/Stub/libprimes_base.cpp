/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is a stub class definition of CLibPrimesBase

*/

#include "libprimes_base.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes::Impl;

/*************************************************************************************************************************
 Class definition of CLibPrimesBase 
**************************************************************************************************************************/

bool CLibPrimesBase::GetLastErrorMessage (std::string & sErrorMessage)
{
	auto iIterator = m_errors.rbegin();
	if (iIterator != m_errors.rend()) {
		sErrorMessage = *iIterator;
		return true;
	}else {
		sErrorMessage = "";
		return false;
	}
}

void CLibPrimesBase::ClearErrorMessages ()
{
	m_errors.clear();
}

void CLibPrimesBase::RegisterErrorMessage (const std::string & sErrorMessage)
{
	m_errors.push_back(sErrorMessage);
}

