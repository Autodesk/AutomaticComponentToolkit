/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CSieveCalculator

*/


#ifndef __LIBPRIMES_SIEVECALCULATOR
#define __LIBPRIMES_SIEVECALCULATOR

#include "libprimes_interfaces.hpp"

// Parent classes
#include "libprimes_calculator.hpp"
#pragma warning( push)
#pragma warning( disable : 4250)

// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CSieveCalculator 
**************************************************************************************************************************/

class CSieveCalculator : public virtual ISieveCalculator, public virtual CCalculator {
private:

	/**
	* Put private members here.
	*/

	std::vector<LibPrimes_uint64> primes;

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

	void Calculate();

	void GetPrimes(LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer);

};

} // namespace Impl
} // namespace LibPrimes

#pragma warning( pop )
#endif // __LIBPRIMES_SIEVECALCULATOR
