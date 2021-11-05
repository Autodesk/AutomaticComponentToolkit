/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of CAnimalIterator

*/


#ifndef __RTTI_ANIMALITERATOR
#define __RTTI_ANIMALITERATOR

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
 Class declaration of CAnimalIterator 
**************************************************************************************************************************/

class CAnimalIterator : public virtual IAnimalIterator, public virtual CBase {
private:

	/**
	* Put private members here.
	*/
        std::vector<CAnimal *>::iterator m_Current;
	std::vector<CAnimal *>::iterator m_End;
protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/
        CAnimalIterator(std::vector<CAnimal *>::iterator begin, std::vector<CAnimal *>::iterator end);

	/**
	* Public member functions to implement.
	*/

	IAnimal * GetNextAnimal() override;
	bool GetNextOptinalAnimal(IAnimal*& pAnimal) override;
	bool GetNextMandatoryAnimal(IAnimal*& pAnimal) override;

};

} // namespace Impl
} // namespace RTTI

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __RTTI_ANIMALITERATOR
