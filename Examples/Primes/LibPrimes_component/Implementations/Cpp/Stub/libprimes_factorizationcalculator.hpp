/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CLibPrimesFactorizationCalculator

*/


#ifndef __LIBPRIMES_LIBPRIMESFACTORIZATIONCALCULATOR
#define __LIBPRIMES_LIBPRIMESFACTORIZATIONCALCULATOR

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
 Class declaration of CLibPrimesFactorizationCalculator 
**************************************************************************************************************************/

class CLibPrimesFactorizationCalculator : public virtual ILibPrimesFactorizationCalculator, public virtual CLibPrimesCalculator {
private:

	std::vector<sLibPrimesPrimeFactor> primeFactors;

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

	void GetPrimeFactors (LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer);

};

} // namespace Impl
} // namespace LibPrimes

#pragma warning( pop )
#endif // __LIBPRIMES_LIBPRIMESFACTORIZATIONCALCULATOR
