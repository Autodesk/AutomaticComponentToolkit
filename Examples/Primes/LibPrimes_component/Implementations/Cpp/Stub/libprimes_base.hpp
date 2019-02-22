/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CLibPrimesBase

*/


#ifndef __LIBPRIMES_LIBPRIMESBASE
#define __LIBPRIMES_LIBPRIMESBASE

#include "libprimes_interfaces.hpp"
#include <vector>


// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CLibPrimesBase 
**************************************************************************************************************************/

class CLibPrimesBase : public virtual ILibPrimesBase {
private:

	std::vector<std::string> m_errors;

	/**
	* Put private members here.
	*/

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/

	bool GetLastErrorMessage (std::string & sErrorMessage);

	void ClearErrorMessages ();

	void RegisterErrorMessage (const std::string & sErrorMessage);


	/**
	* Public member functions to implement.
	*/

};

} // namespace Impl
} // namespace LibPrimes

#endif // __LIBPRIMES_LIBPRIMESBASE
