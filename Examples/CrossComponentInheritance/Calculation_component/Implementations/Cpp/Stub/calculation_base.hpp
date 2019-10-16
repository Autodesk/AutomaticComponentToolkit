/*++

Copyright (C) 2019 Calculation developers

All rights reserved.

Abstract: This is the class declaration of CBase

*/


#ifndef __CALCULATION_BASE
#define __CALCULATION_BASE

#include "calculation_interfaces.hpp"
#include <vector>
#include <list>
#include <memory>


// Include custom headers here.


namespace Calculation {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CBase 
**************************************************************************************************************************/

class CBase : public virtual IBase {
private:

	std::unique_ptr<std::list<std::string>> m_pErrors;
	Calculation_uint32 m_nReferenceCount = 1;

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

	void ClearErrorMessages() override;

	void RegisterErrorMessage(const std::string & sErrorMessage) override;

	bool GetLastError(std::string & sErrorMessage) override;

	void ReleaseInstance() override;

	void AcquireInstance() override;


	/**
	* Public member functions to implement.
	*/

};

} // namespace Impl
} // namespace Calculation

#endif // __CALCULATION_BASE
