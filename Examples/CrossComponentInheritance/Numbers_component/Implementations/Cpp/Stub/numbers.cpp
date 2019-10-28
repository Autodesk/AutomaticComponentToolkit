/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated C++ implementation file in order to allow easy
development of Numbers library. It needs to be generated only once.

Interface version: 1.0.0

*/

#include "numbers_abi.hpp"
#include "numbers_interfaces.hpp"
#include "numbers_interfaceexception.hpp"

#include "numbers_variableimpl.hpp"

using namespace Numbers;
using namespace Numbers::Impl;

// TODO: self injected header
Numbers::Binding::PWrapper CWrapper::sPNumbersWrapper;

Numbers::Binding::PVariable CWrapper::CreateVariable(const Numbers_double dInitialValue)
{
	PIVariableImpl pImpl(new CVariableImpl());
	pImpl->SetValue(dInitialValue);
	return std::make_shared<Numbers::Binding::CVariable>(pImpl->GetExtendedHandle());
}

IVariableImpl * CWrapper::CreateVariableImpl(const Numbers_double dInitialValue)
{
	PIVariableImpl pImpl(new CVariableImpl());
	pImpl->SetValue(dInitialValue);
	return pImpl.getCoOwningPtr();
}

bool CWrapper::GetLastError(std::string & sErrorMessage)
{
	throw ENumbersInterfaceException(NUMBERS_ERROR_NOTIMPLEMENTED);
}

void CWrapper::GetVersion(Numbers_uint32 & nMajor, Numbers_uint32 & nMinor, Numbers_uint32 & nMicro)
{
	nMajor = NUMBERS_VERSION_MAJOR;
	nMinor = NUMBERS_VERSION_MINOR;
	nMicro = NUMBERS_VERSION_MICRO;
}

