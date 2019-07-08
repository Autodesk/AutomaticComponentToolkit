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
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CFactorizationCalculator 
**************************************************************************************************************************/

class CFactorizationCalculator : public virtual IFactorizationCalculator, public virtual CCalculator {
private:

	/**
	* Put private members here.
	*/
	std::vector<sPrimeFactor> primeFactors;

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

	void GetPrimeFactors(LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, LibPrimes::sPrimeFactor * pPrimeFactorsBuffer) override;

	void Calculate() override;
};

} // namespace Impl
} // namespace LibPrimes

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __LIBPRIMES_FACTORIZATIONCALCULATOR
