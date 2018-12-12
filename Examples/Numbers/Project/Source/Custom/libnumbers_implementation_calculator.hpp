/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is the class declaration of CInternalLibNumbersCalculator
Interface version: 1.0.0

**/


#ifndef __LIBNUMBERS_LIBNUMBERSCALCULATOR
#define __LIBNUMBERS_LIBNUMBERSCALCULATOR

#include "libnumbers_interfaces.hpp"


// Include custom headers here.


namespace LibNumbers {


/*************************************************************************************************************************
 Class declaration of CInternalLibNumbersCalculator 
**************************************************************************************************************************/

class CInternalLibNumbersCalculator : public virtual IInternalLibNumbersCalculator {
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

	IInternalLibNumbersNumber * AddNumbers (IInternalLibNumbersNumber* pValue1, IInternalLibNumbersNumber* pValue2);

};

}

#endif // __LIBNUMBERS_LIBNUMBERSCALCULATOR
