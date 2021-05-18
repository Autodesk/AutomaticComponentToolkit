/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of CReptile

*/


#ifndef __RTTI_REPTILE
#define __RTTI_REPTILE

#include "rtti_interfaces.hpp"

// Parent classes
#include "rtti_animal.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace RTTI {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CReptile 
**************************************************************************************************************************/

class CReptile : public virtual IReptile, public virtual CAnimal {
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
#endif // __RTTI_REPTILE
