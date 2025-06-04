/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is the class declaration of CTestClass

*/


#ifndef __LIBUNITTEST_TESTCLASS
#define __LIBUNITTEST_TESTCLASS

#include "libunittest_interfaces.hpp"

// Parent classes
#include "libunittest_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace LibUnitTest {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CTestClass 
**************************************************************************************************************************/

class CTestClass : public virtual ITestClass, public virtual CBase {
private:

	/**
	* Put private members here.
	*/

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/


	/**
	* Public member functions to implement.
	*/

	LibUnitTest_double Value() override;

	void SetValue(const LibUnitTest_double dValue) override;

	void SetValueInt(const LibUnitTest_int64 nValue) override;

	void SetValueString(const std::string & sValue) override;

	void UnitTest1(const LibUnitTest_uint8 nValue1, const LibUnitTest_uint16 nValue2, const LibUnitTest_uint32 nValue3, const LibUnitTest_uint64 nValue4, LibUnitTest_uint8 & nOutValue1, LibUnitTest_uint16 & nOutValue2, LibUnitTest_uint32 & nOutValue3, LibUnitTest_uint64 & nOutValue4) override;

	void UnitTest2(const LibUnitTest_int8 nValue1, const LibUnitTest_int16 nValue2, const LibUnitTest_int32 nValue3, const LibUnitTest_int64 nValue4, LibUnitTest_int8 & nOutValue1, LibUnitTest_int16 & nOutValue2, LibUnitTest_int32 & nOutValue3, LibUnitTest_int64 & nOutValue4) override;

	void UnitTest3(const bool bValue1, const LibUnitTest_single fValue2, const LibUnitTest_double dValue3, const LibUnitTest::eTestEnum eValue4, bool & bOutValue1, LibUnitTest_single & fOutValue2, LibUnitTest_double & dOutValue3, LibUnitTest::eTestEnum & eOutValue4) override;

	std::string UnitTest4(const std::string & sValue, std::string & sOutValue) override;

};

} // namespace Impl
} // namespace LibUnitTest

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __LIBUNITTEST_TESTCLASS
