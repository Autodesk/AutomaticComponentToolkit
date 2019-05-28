/*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is the class declaration of CCalculator

*/


#ifndef __CALCULATOR_CALCULATOR
#define __CALCULATOR_CALCULATOR

#include "calculator_interfaces.hpp"

// Parent classes
#include "calculator_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.

#include <vector>

namespace Calculator {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CCalculator 
**************************************************************************************************************************/

class CCalculator : public virtual ICalculator, public virtual CBase {
private:

	/**
	* Put private members here.
	*/

	std::vector<PIVariable> m_pVariables;

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

	void EnlistVariable(IVariable* pVariable);

	IVariable * GetEnlistedVariable(const Calculator_uint32 nIndex);

	void ClearVariables();

	IVariable * Multiply();

	IVariable * Add();

};

} // namespace Impl
} // namespace Calculator

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __CALCULATOR_CALCULATOR
