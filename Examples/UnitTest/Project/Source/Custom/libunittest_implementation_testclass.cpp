/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is a stub class definition of CInternalLibUnitTestTestClass
Interface version: 1.0.0

**/

#include "libunittest_implementation_testclass.hpp"
#include "libunittest_interfaceexception.hpp"

// Include custom headers here.


using namespace LibUnitTest;

/*************************************************************************************************************************
 Class definition of CInternalLibUnitTestTestClass 
**************************************************************************************************************************/

double CInternalLibUnitTestTestClass::Value ()
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::SetValue (const double dValue)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::SetValueInt (const long long nValue)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::SetValueString (const std::string sValue)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::UnitTest1 (const unsigned char nValue1, const unsigned short nValue2, const unsigned int nValue3, const unsigned long long nValue4, unsigned char & nOutValue1, unsigned short & nOutValue2, unsigned int & nOutValue3, unsigned long long & nOutValue4)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::UnitTest2 (const char nValue1, const short nValue2, const int nValue3, const long long nValue4, char & nOutValue1, short & nOutValue2, int & nOutValue3, long long & nOutValue4)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::UnitTest3 (const bool bValue1, const float fValue2, const double dValue3, const eLibUnitTestTestEnum eValue4, bool & bOutValue1, float & fOutValue2, double & dOutValue3, eLibUnitTestTestEnum & eOutValue4)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CInternalLibUnitTestTestClass::UnitTest4 (const unsigned int nValue1BufferSize, const unsigned char * pValue1Buffer, const unsigned int nValue2BufferSize, const unsigned short * pValue2Buffer, const unsigned int nValue3BufferSize, const unsigned int * pValue3Buffer, const unsigned int nValue4BufferSize, const unsigned long long * pValue4Buffer)
{
	throw ELibUnitTestInterfaceException (LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

