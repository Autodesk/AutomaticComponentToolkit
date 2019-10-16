/*++

Copyright (C) 2019 Calculation developers

All rights reserved.

Abstract: This is the class declaration of CCalculator

*/


#ifndef __CALCULATION_CALCULATOR
#define __CALCULATION_CALCULATOR

#include "calculation_interfaces.hpp"

// Parent classes
#include "calculation_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace Calculation {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CCalculator 
**************************************************************************************************************************/

class CCalculator : public virtual ICalculator, public virtual CBase {
private:

	std::vector<Numbers::PVariable> m_pVariables;
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
	~CCalculator() {
		m_pVariables.clear();
	}

	/**
	* Public member functions to implement.
	*/

	void EnlistVariable(Numbers::PVariable pVariable) override;

	Numbers::PVariable GetEnlistedVariable(const Calculation_uint32 nIndex) override;

	void ClearVariables() override;

	Numbers::PVariable Multiply() override;

	Numbers::PVariable Add() override;

};

} // namespace Impl
} // namespace Calculation

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __CALCULATION_CALCULATOR
