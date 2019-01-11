/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CLibPrimesCalculator

*/


#ifndef __LIBPRIMES_LIBPRIMESCALCULATOR
#define __LIBPRIMES_LIBPRIMESCALCULATOR

#include "libprimes_interfaces.hpp"


// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CLibPrimesCalculator 
**************************************************************************************************************************/

class CLibPrimesCalculator : public virtual ILibPrimesCalculator {
private:

protected:

	LibPrimes_uint64  m_value;
	LibPrimesProgressCallback m_Callback;

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/

	CLibPrimesCalculator();
	/**
	* Public member functions to implement.
	*/
	LibPrimes_uint64 GetValue();

	void SetValue(const LibPrimes_uint64 nValue);

	void SetProgressCallback (const LibPrimesProgressCallback pProgressCallback);

	virtual void Calculate () = 0;
};

} // namespace Impl
} // namespace LibPrimes

#endif // __LIBPRIMES_LIBPRIMESCALCULATOR
