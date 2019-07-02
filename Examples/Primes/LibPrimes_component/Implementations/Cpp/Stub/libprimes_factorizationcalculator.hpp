/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CFactorizationCalculator

*/


#ifndef __LIBPRIMES_FACTORIZATIONCALCULATOR
#define __LIBPRIMES_FACTORIZATIONCALCULATOR

#include "libprimes_interfaces.hpp"

// Parent classes
#include "libprimes_calculator.hpp"
#pragma warning( push)
#pragma warning( disable : 4250)

// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CFactorizationCalculator 
**************************************************************************************************************************/

class CFactorizationCalculator : public virtual IFactorizationCalculator, public virtual CCalculator {
private:

	std::vector<sPrimeFactor> primeFactors;

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/

	void Calculate() override;

	/**
	* Public member functions to implement.
	*/

	void GetPrimeFactors(LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, LibPrimes::sPrimeFactor * pPrimeFactorsBuffer) override;

};

} // namespace Impl
} // namespace LibPrimes

#pragma warning( pop )
#endif // __LIBPRIMES_FACTORIZATIONCALCULATOR
