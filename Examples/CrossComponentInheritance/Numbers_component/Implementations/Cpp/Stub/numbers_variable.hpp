/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

Abstract: This is the class declaration of CVariable

*/


#ifndef __NUMBERS_VARIABLE
#define __NUMBERS_VARIABLE

#include "numbers_interfaces.hpp"

// Parent classes
#include "numbers_base.hpp"
#ifdef _MSC_VER
#pragma warning(push)
#pragma warning(disable : 4250)
#endif

// Include custom headers here.


namespace Numbers {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CVariable 
**************************************************************************************************************************/

class CVariable : public virtual IVariable, public virtual CBase {
private:

	Numbers_double m_dValue = 0.0;
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
	CVariable(Numbers_double dValue) : m_dValue(dValue) {};

	/**
	* Public member functions to implement.
	*/

	Numbers_double GetValue() override;

	void SetValue(const Numbers_double dValue) override;

};

} // namespace Impl
} // namespace Numbers

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __NUMBERS_VARIABLE
