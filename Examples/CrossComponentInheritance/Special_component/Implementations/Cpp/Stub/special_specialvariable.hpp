/*++

Copyright (C) 2019 Special Numbers developers

All rights reserved.

Abstract: This is the class declaration of CSpecialVariable

*/


#ifndef __SPECIAL_SPECIALVARIABLE
#define __SPECIAL_SPECIALVARIABLE

#include "special_interfaces.hpp"

// Parent classes
#include "special_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace Special {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CSpecialVariable 
**************************************************************************************************************************/

class CSpecialVariable : public virtual ISpecialVariable, public virtual CBase {
private:

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


	/**
	* Public member functions to implement.
	*/

	Special_int64 GetSpecialValue() override;

};

} // namespace Impl
} // namespace Special

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __SPECIAL_SPECIALVARIABLE
