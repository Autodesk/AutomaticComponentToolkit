#include "libprimes_dynamic.hpp"
#include <iostream>


void progressCallback(float progress, bool *shouldAbort)
{
	std::cout << "Progress = " << round(progress * 100) << std::endl;
	if (shouldAbort) {
		*shouldAbort = false;
	}
}

void progressCallbackCancel(float progress, bool *shouldAbort)
{
	std::cout << "Progress = " << round(progress * 100) << std::endl;
	if (shouldAbort) {
		*shouldAbort = progress > 0.5f;
	}
}


void calculatePrimes(LibPrimes::CLibPrimesWrapper& wrapper, const unsigned long long number) {
	auto sieve = wrapper.CreateSieveCalculator();

	sieve->SetValue(number);

	std::cout << "Calculate using a cancelling callback" << std::endl;
	try {
		sieve->SetProgressCallback(progressCallbackCancel);
		sieve->Calculate();
	}
	catch (LibPrimes::ELibPrimesException &e) {
		if (e.getErrorCode() == LIBPRIMES_ERROR_CALCULATIONABORTED) {
			std::cout << "Calculation aborted" << std::endl;
		}
		else
			throw e;
	}

	std::cout << "Calculate using a noncancelling callback" << std::endl;
	sieve->SetProgressCallback(progressCallback);
	sieve->Calculate();

	std::vector<unsigned long long> primes(0);
	sieve->GetPrimes(primes);

	std::cout << "Primes <= " << number << ":" << std::endl;
	for (size_t i = 0; i < primes.size(); i++) {
		std::cout << primes[i] << " ";
	}
	std::cout << std::endl;
}

void factorize(LibPrimes::CLibPrimesWrapper& wrapper, const unsigned long long number) {
	auto factorisation = wrapper.CreateFactorizationCalculator();

	factorisation->SetValue(number);
	factorisation->Calculate();

	std::vector<sLibPrimesPrimeFactor> primeFactors(0);
	factorisation->GetPrimeFactors(primeFactors);

	std::cout << factorisation->GetValue() << " = ";
	for (size_t i = 0; i < primeFactors.size(); i++) {
		auto pF = primeFactors[i];
		std::cout << pF.m_Prime << "^" << pF.m_Multiplicity << ((i < (primeFactors.size() - 1)) ? " * " : "");
	}
	std::cout << std::endl;
}

int main()
{
	try {
		auto wrapper = LibPrimes::CLibPrimesWrapper::loadLibrary("LibPrimes.dll");
		factorize(*wrapper.get(), 3 * 3 * 17 * 17);
		calculatePrimes(*wrapper.get(), 100);
	}
	catch (std::exception &e) {
		std::cout << e.what() << std::endl;
		return -1;
	}

	return 0;
}