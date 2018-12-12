#ifdef MACHININGUSECPPDYNAMIC
#include "LibNumbers_dynamic.hpp"
#else
#include "LibNumbers.hpp"
#endif


#include <iostream>
#include <exception>

using namespace LibNumbers;

void tryout()
{
#ifdef MACHININGUSECPPDYNAMIC
	std::cout << "LibNumbers_dynamic" << std::endl;
	auto wrapper = LibOMV::CLibNumbersWrapper::CLibNumbersWrapper("LibNumbers.dll");
	auto number = wrapper->CreateNumber();
#else
	std::cout << "LibNumbers" << std::endl;
	auto number = CLibNumbersWrapper::CreateNumber();
#endif
	number->SetValueString("1e3");
	number->SetValueString("3.14");
	number->SetValueString("3,14");
}

int main()
{
	try {
		tryout();
	}
	catch (std::runtime_error &e) {
		std::cout << "Error: " << e.what() << std::endl;
		return -1;
	}
	catch (...) {
		std::cout << "Generic Error" << std::endl;
		return -1;
	}
	std::cout << "Success" << std::endl;
	return 0;
}
