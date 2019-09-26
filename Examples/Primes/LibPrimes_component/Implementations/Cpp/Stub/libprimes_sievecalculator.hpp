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
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

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

	void GetPrimes(LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer) override;

	void Calculate() override;
};

} // namespace Impl
} // namespace LibPrimes

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __LIBPRIMES_SIEVECALCULATOR
