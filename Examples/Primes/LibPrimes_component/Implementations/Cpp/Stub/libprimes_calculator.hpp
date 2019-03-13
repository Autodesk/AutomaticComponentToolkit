/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CLibPrimesCalculator

*/


#ifndef __LIBPRIMES_LIBPRIMESCALCULATOR
#define __LIBPRIMES_LIBPRIMESCALCULATOR

#include "libprimes_interfaces.hpp"

// Parent classes
#include "libprimes_base.hpp"
#pragma warning( push)
#pragma warning( disable : 4250)

// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CLibPrimesCalculator 
**************************************************************************************************************************/

class CLibPrimesCalculator : public virtual ILibPrimesCalculator, public virtual CLibPrimesBase {
private:

	/**
	* Put private members here.
	*/

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

	LibPrimes_uint64 GetValue ();

	void SetValue (const LibPrimes_uint64 nValue);

	virtual void Calculate ();

	void SetProgressCallback (const LibPrimesProgressCallback pProgressCallback);

};

} // namespace Impl
} // namespace LibPrimes

#pragma warning( pop )
#endif // __LIBPRIMES_LIBPRIMESCALCULATOR
