/*++

Copyright (C) 2025 Thread-safe library developers

All rights reserved.

Abstract: This is the class declaration of CSoftStringReturner

*/


#ifndef __LIBTHREADSAFE_SOFTSTRINGRETURNER
#define __LIBTHREADSAFE_SOFTSTRINGRETURNER

#include "libthreadsafe_interfaces.hpp"

// Parent classes
#include "libthreadsafe_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.
#include <vector>

namespace LibThreadSafe {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CSoftStringReturner 
**************************************************************************************************************************/

class CSoftStringReturner : public virtual ISoftStringReturner, public virtual CBase {
private:

	/**
	* Put private members here.
	*/
	std::vector<int> values;

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/


	/**
	* Public member functions to implement.
	*/

	std::string GetString() override;

	void ThreadSafetyCheck() override;

};

} // namespace Impl
} // namespace LibThreadSafe

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __LIBTHREADSAFE_SOFTSTRINGRETURNER
