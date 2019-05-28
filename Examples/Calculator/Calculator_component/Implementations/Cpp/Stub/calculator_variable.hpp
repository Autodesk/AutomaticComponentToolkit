/*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is the class declaration of CVariable

*/


#ifndef __CALCULATOR_VARIABLE
#define __CALCULATOR_VARIABLE

#include "calculator_interfaces.hpp"

// Parent classes
#include "calculator_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace Calculator {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CVariable 
**************************************************************************************************************************/

class CVariable : public virtual IVariable, public virtual CBase {
private:

	Calculator_double m_dValue;

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
	CVariable(Calculator_double);

	/**
	* Public member functions to implement.
	*/

	Calculator_double GetValue();

	void SetValue(const Calculator_double dValue);

};

} // namespace Impl
} // namespace Calculator

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __CALCULATOR_VARIABLE
