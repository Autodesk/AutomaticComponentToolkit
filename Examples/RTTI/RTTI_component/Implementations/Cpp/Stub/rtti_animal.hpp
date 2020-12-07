/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of CAnimal

*/


#ifndef __RTTI_ANIMAL
#define __RTTI_ANIMAL

#include "rtti_interfaces.hpp"

// Parent classes
#include "rtti_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace RTTI {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CAnimal 
**************************************************************************************************************************/

class CAnimal : public virtual IAnimal, public virtual CBase {
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

};

} // namespace Impl
} // namespace RTTI

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __RTTI_ANIMAL
