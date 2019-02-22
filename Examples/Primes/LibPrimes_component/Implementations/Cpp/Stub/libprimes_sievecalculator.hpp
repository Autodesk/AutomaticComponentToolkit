/*++

Copyright (C) 2018 PrimeDevelopers

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
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CLibPrimesSieveCalculator 
**************************************************************************************************************************/

class CLibPrimesSieveCalculator : public virtual ILibPrimesSieveCalculator, public virtual CLibPrimesCalculator {
private:

	std::vector<LibPrimes_uint64> primes;

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/

	void Calculate();

	/**
	* Public member functions to implement.
	*/

	void GetPrimes (LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer);

};

} // namespace Impl
} // namespace LibPrimes

#pragma warning( pop )
#endif // __LIBPRIMES_LIBPRIMESSIEVECALCULATOR
