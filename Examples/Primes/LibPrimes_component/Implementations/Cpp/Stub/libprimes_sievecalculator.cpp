/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is a stub class definition of CLibPrimesSieveCalculator

*/

#include "libprimes_sievecalculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes::Impl;

/*************************************************************************************************************************
 Class definition of CLibPrimesSieveCalculator 
**************************************************************************************************************************/

void CLibPrimesSieveCalculator::Calculate()
{
	primes.clear();

	std::vector<bool> strikenOut(m_value + 1);
	for (LibPrimes_uint64 i = 0; i <= m_value; i++) {
		strikenOut[i] = i < 2;
	}

	LibPrimes_uint64 sqrtValue = (LibPrimes_uint64)(std::sqrt(m_value));

	int progressStep = (int)std::ceil(sqrtValue / 20.0f);

	for (LibPrimes_uint64 i = 2; i <= sqrtValue; i++) {

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
			for (LibPrimes_uint64 j = i * i; j < m_value; j += i) {
				strikenOut[j] = true;
			}
		}
	}

	for (LibPrimes_uint64 i = sqrtValue; i <= m_value; i++) {
		if (!strikenOut[i]) {
			primes.push_back(i);
		}
	}
}

void CLibPrimesSieveCalculator::GetPrimes (LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer)
{
	if (primes.size() == 0)
		throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NORESULTAVAILABLE);

	if (pPrimesNeededCount)
		*pPrimesNeededCount = (unsigned int)primes.size();

	if (nPrimesBufferSize >= primes.size() && pPrimesBuffer)
	{
		for (size_t i = 0; i < primes.size(); i++)
		{
			pPrimesBuffer[i] = primes[i];
		}
	}
}

