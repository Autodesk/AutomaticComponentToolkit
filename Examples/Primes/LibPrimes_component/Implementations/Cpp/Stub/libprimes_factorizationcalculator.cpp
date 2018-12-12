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

	unsigned long long nValue = m_value;
	for (unsigned long long i = 2; i <= nValue; i++) {
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


void CLibPrimesFactorizationCalculator::GetPrimeFactors (unsigned int nPrimeFactorsBufferSize, unsigned int * pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer)
{
	if (primeFactors.size() == 0)
		throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NORESULTAVAILABLE);

	if (pPrimeFactorsNeededCount)
		*pPrimeFactorsNeededCount = (unsigned int)primeFactors.size();

	if (nPrimeFactorsBufferSize >= primeFactors.size() && pPrimeFactorsBuffer)
	{
		for (int i = 0; i < primeFactors.size(); i++)
		{
			pPrimeFactorsBuffer[i] = primeFactors[i];
		}
	}
	
}

