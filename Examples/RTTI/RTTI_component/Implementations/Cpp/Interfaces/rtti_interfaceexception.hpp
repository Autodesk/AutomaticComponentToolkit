/*++

Copyright (C) 2021 ADSK

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated C++ Header file with the basic internal
 exception type in order to allow an easy use of RTTI

Interface version: 1.0.0

*/

#ifndef __RTTI_INTERFACEEXCEPTION_HEADER
#define __RTTI_INTERFACEEXCEPTION_HEADER

#include <exception>
#include <stdexcept>
#include "rtti_types.hpp"

/*************************************************************************************************************************
 Class ERTTIInterfaceException
**************************************************************************************************************************/


class ERTTIInterfaceException : public std::exception {
protected:
	/**
	* Error code for the Exception.
	*/
	RTTIResult m_errorCode;
	/**
	* Error message for the Exception.
	*/
	std::string m_errorMessage;

public:
	/**
	* Exception Constructor.
	*/
	ERTTIInterfaceException(RTTIResult errorCode);

	/**
	* Custom Exception Constructor.
	*/
	ERTTIInterfaceException(RTTIResult errorCode, std::string errorMessage);

	/**
	* Returns error code
	*/
	RTTIResult getErrorCode();
	/**
	* Returns error message
	*/
	const char* what() const noexcept override;
};

#endif // __RTTI_INTERFACEEXCEPTION_HEADER
