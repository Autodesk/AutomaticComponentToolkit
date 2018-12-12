/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is the class declaration of CInternalLibNumbersNumber
Interface version: 1.0.0

**/


#ifndef __LIBNUMBERS_LIBNUMBERSNUMBER
#define __LIBNUMBERS_LIBNUMBERSNUMBER

#include "libnumbers_interfaces.hpp"


// Include custom headers here.


namespace LibNumbers {


/*************************************************************************************************************************
 Class declaration of CInternalLibNumbersNumber 
**************************************************************************************************************************/

class CInternalLibNumbersNumber : public virtual IInternalLibNumbersNumber {
private:

	double m_value;

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

};

}

#endif // __LIBNUMBERS_LIBNUMBERSNUMBER
