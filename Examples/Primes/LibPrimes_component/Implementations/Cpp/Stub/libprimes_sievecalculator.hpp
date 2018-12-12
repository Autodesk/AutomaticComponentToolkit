/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

Abstract: This is the class declaration of CLibPrimesSieveCalculator

*/


#ifndef __LIBPRIMES_LIBPRIMESSIEVECALCULATOR
#define __LIBPRIMES_LIBPRIMESSIEVECALCULATOR

#include "libprimes_interfaces.hpp"

// Parent classes
#include "libprimes_calculator.hpp"
#pragma warning( push)
#pragma warning( disable : 4250)

// Include custom headers here.

#include <vector>

namespace LibPrimes {


/*************************************************************************************************************************
 Class declaration of CLibPrimesSieveCalculator 
**************************************************************************************************************************/

class CLibPrimesSieveCalculator : public virtual ILibPrimesSieveCalculator, public virtual CLibPrimesCalculator {
private:

	std::vector<unsigned long long> primes;

protected:

	/**
	* Put protected members here.
	*/

public:

	void Calculate();

	/**
	* Public member functions to implement.
	*/

	void GetPrimes (unsigned int nPrimesBufferSize, unsigned int * pPrimesNeededCount, unsigned long long * pPrimesBuffer);

};

}

#pragma warning( pop )
#endif // __LIBPRIMES_LIBPRIMESSIEVECALCULATOR
