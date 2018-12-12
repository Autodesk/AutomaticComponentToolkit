/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

Abstract: This is a stub class definition of CLibPrimesSieveCalculator

*/

#include "libprimes_sievecalculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.
#include <cmath>


using namespace LibPrimes;

/*************************************************************************************************************************
 Class definition of CLibPrimesSieveCalculator 
**************************************************************************************************************************/

void CLibPrimesSieveCalculator::Calculate()
{
	primes.clear();

	std::vector<bool> strikenOut(m_value + 1);
	for (unsigned long long i = 0; i <= m_value; i++) {
		strikenOut[i] = i < 2;
	}

	unsigned long sqrtValue = (unsigned long)(std::sqrt(m_value));

	int progressStep = (int)std::ceil(sqrtValue / 20.0f);
	
	for (unsigned long long i = 2; i <= sqrtValue; i++) {

		if (m_Callback) {
			if (i % progressStep == 0) {
				bool shouldAbort = false;
				(*m_Callback)(float(i) / sqrtValue, &shouldAbort);
				if (shouldAbort) {
					throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_CALCULATIONABORTED);
				}
			}
		}

		if (!strikenOut[i]) {
			primes.push_back(i);
			for (unsigned long long j = i * i; j < m_value; j += i) {
				strikenOut[j] = true;
			}
		}
	}

	for (unsigned long long i = sqrtValue; i <= m_value; i++) {
		if (!strikenOut[i]) {
			primes.push_back(i);
		}
	}
}

void CLibPrimesSieveCalculator::GetPrimes (unsigned int nPrimesBufferSize, unsigned int * pPrimesNeededCount, unsigned long long * pPrimesBuffer)
{
	if (primes.size() == 0)
		throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NORESULTAVAILABLE);

	if (pPrimesNeededCount)
		*pPrimesNeededCount = (unsigned int)primes.size();

	if (nPrimesBufferSize >= primes.size() && pPrimesBuffer)
	{
		for (int i = 0; i < primes.size(); i++)
		{
			pPrimesBuffer[i] = primes[i];
		}
	}

}

