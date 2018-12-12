/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

Abstract: This is the class declaration of CLibPrimesCalculator

*/


#ifndef __LIBPRIMES_LIBPRIMESCALCULATOR
#define __LIBPRIMES_LIBPRIMESCALCULATOR

#include "libprimes_interfaces.hpp"


// Include custom headers here.


namespace LibPrimes {


/*************************************************************************************************************************
 Class declaration of CLibPrimesCalculator 
**************************************************************************************************************************/

class CLibPrimesCalculator : public virtual ILibPrimesCalculator {
private:

protected:

	unsigned long long m_value;
	LibPrimesProgressCallback m_Callback;

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/

	CLibPrimesCalculator();
	/**
	* Public member functions to implement.
	*/
	unsigned long long GetValue();

	void SetValue(const unsigned long long nValue);

	void SetProgressCallback (const LibPrimesProgressCallback pProgressCallback);

	virtual void Calculate () = 0;
};

}

#endif // __LIBPRIMES_LIBPRIMESCALCULATOR
