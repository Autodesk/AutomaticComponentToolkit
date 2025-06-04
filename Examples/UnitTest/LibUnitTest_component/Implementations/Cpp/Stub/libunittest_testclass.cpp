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

LibUnitTest_double CTestClass::Value()
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CTestClass::SetValue(const LibUnitTest_double dValue)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CTestClass::SetValueInt(const LibUnitTest_int64 nValue)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CTestClass::SetValueString(const std::string & sValue)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CTestClass::UnitTest1(const LibUnitTest_uint8 nValue1, const LibUnitTest_uint16 nValue2, const LibUnitTest_uint32 nValue3, const LibUnitTest_uint64 nValue4, LibUnitTest_uint8 & nOutValue1, LibUnitTest_uint16 & nOutValue2, LibUnitTest_uint32 & nOutValue3, LibUnitTest_uint64 & nOutValue4)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CTestClass::UnitTest2(const LibUnitTest_int8 nValue1, const LibUnitTest_int16 nValue2, const LibUnitTest_int32 nValue3, const LibUnitTest_int64 nValue4, LibUnitTest_int8 & nOutValue1, LibUnitTest_int16 & nOutValue2, LibUnitTest_int32 & nOutValue3, LibUnitTest_int64 & nOutValue4)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

void CTestClass::UnitTest3(const bool bValue1, const LibUnitTest_single fValue2, const LibUnitTest_double dValue3, const LibUnitTest::eTestEnum eValue4, bool & bOutValue1, LibUnitTest_single & fOutValue2, LibUnitTest_double & dOutValue3, LibUnitTest::eTestEnum & eOutValue4)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

std::string CTestClass::UnitTest4(const std::string & sValue, std::string & sOutValue)
{
	throw ELibUnitTestInterfaceException(LIBUNITTEST_ERROR_NOTIMPLEMENTED);
}

