/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is a stub class definition of CBase

*/

#include "libprimes_base.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes::Impl;

/*************************************************************************************************************************
 Class definition of CBase 
**************************************************************************************************************************/

bool CBase::GetLastErrorMessage(std::string & sErrorMessage)
{
	if (m_pErrors && !m_pErrors->empty()) {
		sErrorMessage = m_pErrors->back();
		m_pErrors->pop_back();
		return true;
	} else {
		sErrorMessage = "";
		return false;
	}
}

void CBase::ClearErrorMessages()
{
	m_pErrors.reset();
}

void CBase::RegisterErrorMessage(const std::string & sErrorMessage)
{
	if (!m_pErrors) {
		m_pErrors.reset(new std::list<std::string>());
	}
	m_pErrors->push_back(sErrorMessage);
}

void CBase::IncRefCount()
{
	++m_nReferenceCount;
}

bool CBase::DecRefCount()
{
	m_nReferenceCount--;
	if (!m_nReferenceCount) {;
		delete this;
		return true;
	}
	return false;
}

