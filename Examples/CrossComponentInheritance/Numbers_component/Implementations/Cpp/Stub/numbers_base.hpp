/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

Abstract: This is the class declaration of CBase

*/


#ifndef __NUMBERS_BASE
#define __NUMBERS_BASE

#include "numbers_interfaces.hpp"
#include <vector>
#include <list>
#include <memory>


// Include custom headers here.


namespace Numbers {
namespace Impl {


/*************************************************************************************************************************
 Class declaration of CBase 
**************************************************************************************************************************/

class CBase : public virtual IBase {
private:

	std::unique_ptr<std::list<std::string>> m_pErrors;
	Numbers_uint32 m_nReferenceCount = 1;

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

	bool GetLastErrorMessage(std::string & sErrorMessage) override;

	void ClearErrorMessages() override;

	void RegisterErrorMessage(const std::string & sErrorMessage) override;

	void IncRefCount() override;

	bool DecRefCount() override;


	/**
	* Public member functions to implement.
	*/

};

} // namespace Impl
} // namespace Numbers

#endif // __NUMBERS_BASE
