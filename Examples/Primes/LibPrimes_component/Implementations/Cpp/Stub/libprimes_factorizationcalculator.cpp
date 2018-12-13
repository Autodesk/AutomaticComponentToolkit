/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

Abstract: This is a stub class definition of CLibPrimesFactorizationCalculator

*/

#include "libprimes_factorizationcalculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes;

/*************************************************************************************************************************
 Class definition of CLibPrimesFactorizationCalculator 
**************************************************************************************************************************/

void CLibPrimesFactorizationCalculator::Calculate()
{
	primeFactors.clear();

	LibPrimes_uint64 nValue = m_value;
	for (LibPrimes_uint64 i = 2; i <= nValue; i++) {

		if (m_Callback) {
			bool shouldAbort = false;
			(*m_Callback)(1 - float(nValue) / m_value, &shouldAbort);
			if (shouldAbort) {
				throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_CALCULATIONABORTED);
			}
		}

		sLibPrimesPrimeFactor primeFactor;
		primeFactor.m_Prime = i;
		primeFactor.m_Multiplicity = 0;
		while (nValue % i == 0) {
			primeFactor.m_Multiplicity++;
			nValue = nValue / i;
		}
		if (primeFactor.m_Multiplicity > 0) {
			primeFactors.push_back(primeFactor);
		}
	}
}


void CLibPrimesFactorizationCalculator::GetPrimeFactors (LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64 * pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer)
{
	if (primeFactors.size() == 0)
		throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NORESULTAVAILABLE);

	if (pPrimeFactorsNeededCount)
		*pPrimeFactorsNeededCount = (unsigned int)primeFactors.size();

	if (nPrimeFactorsBufferSize >= primeFactors.size() && pPrimeFactorsBuffer)
	{
		for (size_t i = 0; i < primeFactors.size(); i++)
		{
			pPrimeFactorsBuffer[i] = primeFactors[i];
		}
	}
}

