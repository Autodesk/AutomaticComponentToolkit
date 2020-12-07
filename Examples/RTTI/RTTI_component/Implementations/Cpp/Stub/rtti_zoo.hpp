/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of CZoo

*/


#ifndef __RTTI_ZOO
#define __RTTI_ZOO

#include "rtti_interfaces.hpp"

// Parent classes
#include "rtti_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.
#include <vector>

namespace RTTI {
namespace Impl {

class CAnimal;
/*************************************************************************************************************************
 Class declaration of CZoo 
**************************************************************************************************************************/

class CZoo : public virtual IZoo, public virtual CBase {
private:

	/**
	* Put private members here.
	*/
        std::vector<CAnimal*> m_Animals;

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/
        std::vector<CAnimal*>& Animals();


	/**
	* Public member functions to implement.
	*/

	IAnimalIterator * Iterator() override;

};

} // namespace Impl
} // namespace RTTI

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __RTTI_ZOO
