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

	double m_dValue;

protected:

	/**
	* Put protected members here.
	*/

public:

	/**
	* Put additional public members here. They will not be visible in the external API.
	*/
	CVariable(double dInitialValue);

	/**
	* Public member functions to implement.
	*/

	Numbers_double GetValue();

	void SetValue(const Numbers_double dValue);

};

} // namespace Impl
} // namespace Numbers

#ifdef _MSC_VER
#pragma warning(pop)
#endif
#endif // __NUMBERS_VARIABLE
