/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is a stub class definition of CTestClass

*/

#include "libunittest_testclass.hpp"
#include "libunittest_interfaceexception.hpp"

// Include custom headers here.


using namespace LibUnitTest::Impl;

/*************************************************************************************************************************
 Class definition of CTestClass 
**************************************************************************************************************************/

LibUnitTest_uint64 CTestClass::ClassTypeId()
{
	return 1; // TestClass type ID
}

LibUnitTest_double CTestClass::Value()
{
	return m_Value;
}

void CTestClass::SetValue(const LibUnitTest_double dValue)
{
	m_Value = dValue;
}

void CTestClass::SetValueInt(const LibUnitTest_int64 nValue)
{
	m_IntValue = nValue;
	m_Value = static_cast<LibUnitTest_double>(nValue);
}

void CTestClass::SetValueString(const std::string & sValue)
{
	m_StringValue = sValue;
	// Try to convert to double if possible
	try {
		m_Value = std::stod(sValue);
	} catch (...) {
		m_Value = 0.0;
	}
}

void CTestClass::UnitTest1(const LibUnitTest_uint8 nValue1, const LibUnitTest_uint16 nValue2, const LibUnitTest_uint32 nValue3, const LibUnitTest_uint64 nValue4, LibUnitTest_uint8 & nOutValue1, LibUnitTest_uint16 & nOutValue2, LibUnitTest_uint32 & nOutValue3, LibUnitTest_uint64 & nOutValue4)
{
	// Return the same values for unit testing
	nOutValue1 = nValue1;
	nOutValue2 = nValue2;
	nOutValue3 = nValue3;
	nOutValue4 = nValue4;
}

void CTestClass::UnitTest2(const LibUnitTest_int8 nValue1, const LibUnitTest_int16 nValue2, const LibUnitTest_int32 nValue3, const LibUnitTest_int64 nValue4, LibUnitTest_int8 & nOutValue1, LibUnitTest_int16 & nOutValue2, LibUnitTest_int32 & nOutValue3, LibUnitTest_int64 & nOutValue4)
{
	// Return the same values for unit testing
	nOutValue1 = nValue1;
	nOutValue2 = nValue2;
	nOutValue3 = nValue3;
	nOutValue4 = nValue4;
}

void CTestClass::UnitTest3(const bool bValue1, const LibUnitTest_single fValue2, const LibUnitTest_double dValue3, const LibUnitTest::eTestEnum eValue4, bool & bOutValue1, LibUnitTest_single & fOutValue2, LibUnitTest_double & dOutValue3, LibUnitTest::eTestEnum & eOutValue4)
{
	// Return the same values for unit testing
	bOutValue1 = bValue1;
	fOutValue2 = fValue2;
	dOutValue3 = dValue3;
	eOutValue4 = eValue4;
}

std::string CTestClass::UnitTest4(const std::string & sValue, std::string & sOutValue)
{
	// Return the same values for unit testing
	sOutValue = sValue;
	return sValue;
}

