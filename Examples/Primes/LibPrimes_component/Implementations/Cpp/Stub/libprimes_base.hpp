/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of CBase

*/


#ifndef __LIBPRIMES_BASE
#define __LIBPRIMES_BASE

#include "libprimes_interfaces.hpp"
#include <vector>


// Include custom headers here.


namespace LibPrimes {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CBase 
**************************************************************************************************************************/

class CBase : public virtual IBase {
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

	bool GetLastErrorMessage(std::string & sErrorMessage);

	void ClearErrorMessages();

	void RegisterErrorMessage(const std::string & sErrorMessage);


	/**
	* Public member functions to implement.
	*/

};

} // namespace Impl
} // namespace LibPrimes

#endif // __LIBPRIMES_BASE
