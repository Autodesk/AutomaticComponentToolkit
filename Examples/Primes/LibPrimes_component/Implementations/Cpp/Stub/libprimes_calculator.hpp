/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CCalculator

*/


#ifndef __LIBPRIMES_CALCULATOR
#define __LIBPRIMES_CALCULATOR

#include "libprimes_interfaces.hpp"

// Parent classes
#include "libprimes_base.hpp"
#pragma warning( push)
#pragma warning( disable : 4250)

// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CCalculator 
**************************************************************************************************************************/

class CCalculator : public virtual ICalculator, public virtual CBase {
private:

	/**
	* Put private members here.
	*/

protected:

	/**
	* Put protected members here.
	*/
	LibPrimes_uint64 m_value;

	ProgressCallback m_Callback;
public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/


	/**
	* Public member functions to implement.
	*/

	void SetProgressCallback(const LibPrimes::ProgressCallback pProgressCallback);

	LibPrimes_uint64 GetValue();

	void SetValue(const LibPrimes_uint64 nValue);

	void Calculate();

};

} // namespace Impl
} // namespace LibPrimes

#pragma warning( pop )
#endif // __LIBPRIMES_CALCULATOR
