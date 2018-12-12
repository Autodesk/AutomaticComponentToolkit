/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is the class declaration of CInternalLibUnitTestTestClass
Interface version: 1.0.0

**/


#ifndef __LIBUNITTEST_LIBUNITTESTTESTCLASS
#define __LIBUNITTEST_LIBUNITTESTTESTCLASS

#include "libunittest_interfaces.hpp"


// Include custom headers here.


namespace LibUnitTest {


/*************************************************************************************************************************
 Class declaration of CInternalLibUnitTestTestClass 
**************************************************************************************************************************/

class CInternalLibUnitTestTestClass : public virtual IInternalLibUnitTestTestClass {
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

	double Value ();

	void SetValue (const double dValue);

	void SetValueInt (const long long nValue);

	void SetValueString (const std::string sValue);

	void UnitTest1 (const unsigned char nValue1, const unsigned short nValue2, const unsigned int nValue3, const unsigned long long nValue4, unsigned char & nOutValue1, unsigned short & nOutValue2, unsigned int & nOutValue3, unsigned long long & nOutValue4);

	void UnitTest2 (const char nValue1, const short nValue2, const int nValue3, const long long nValue4, char & nOutValue1, short & nOutValue2, int & nOutValue3, long long & nOutValue4);

	void UnitTest3 (const bool bValue1, const float fValue2, const double dValue3, const eLibUnitTestTestEnum eValue4, bool & bOutValue1, float & fOutValue2, double & dOutValue3, eLibUnitTestTestEnum & eOutValue4);

	void UnitTest4 (const unsigned int nValue1BufferSize, const unsigned char * pValue1Buffer, const unsigned int nValue2BufferSize, const unsigned short * pValue2Buffer, const unsigned int nValue3BufferSize, const unsigned int * pValue3Buffer, const unsigned int nValue4BufferSize, const unsigned long long * pValue4Buffer);

};

}

#endif // __LIBUNITTEST_LIBUNITTESTTESTCLASS
