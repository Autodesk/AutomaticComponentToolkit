/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated Go wrapper file in order to allow an easy
 use of Prime Numbers Library

Interface version: 1.2.0

*/


package libprimes

/*************************************************************************************************************************
 Declaration of structs
**************************************************************************************************************************/

type sLibPrimesPrimeFactor struct {
		Prime uint64;
		Multiplicity uint32;
}


/*************************************************************************************************************************
 Declaration of interfaces
**************************************************************************************************************************/

type LibPrimesHandle interface {
		Close() error
}

type LibPrimesGoInterface interface {

	/**
	* Returns the current value of this Calculator
	*
	* @param[in] Calculator - Calculator instance.
	* @return The current value of this Calculator
	*/
	Calculator_GetValue(Calculator LibPrimesHandle) (uint64, error)


	/**
	* Sets the value to be factorized
	*
	* @param[in] Calculator - Calculator instance.
	* @param[in] nValue - The value to be factorized
	*/
	Calculator_SetValue(Calculator LibPrimesHandle, nValue uint64) (error)


	/**
	* Performs the specific calculation of this Calculator
	*
	* @param[in] Calculator - Calculator instance.
	*/
	Calculator_Calculate(Calculator LibPrimesHandle) (error)


	/**
	* Sets the progress callback function
	*
	* @param[in] Calculator - Calculator instance.
	* @param[in] pProgressCallback - The progress callback
	*/
	Calculator_SetProgressCallback(Calculator LibPrimesHandle, pProgressCallback int64) (error)


	/**
	* Returns the prime factors of this number (without multiplicity)
	*
	* @param[in] FactorizationCalculator - FactorizationCalculator instance.
	* @return The prime factors of this number
	*/
	FactorizationCalculator_GetPrimeFactors(FactorizationCalculator LibPrimesHandle) ([]sLibPrimesPrimeFactor, error)


	/**
	* Returns all prime numbers lower or equal to the sieve's value
	*
	* @param[in] SieveCalculator - SieveCalculator instance.
	* @return The primes lower or equal to the sieve's value
	*/
	SieveCalculator_GetPrimes(SieveCalculator LibPrimesHandle) ([]uint64, error)


	/**
	* retrieves the binary version of this library.
	*
	* @param[in] Wrapper - Wrapper instance.
	* @return returns the major version of this library
	* @return returns the minor version of this library
	* @return returns the micro version of this library
	*/
	GetVersion() (uint32, uint32, uint32, error)


	/**
	* Returns the last error recorded on this object
	*
	* @param[in] Wrapper - Wrapper instance.
	* @param[in] Instance - Instance Handle
	* @return Message of the last error
	* @return Is there a last error to query
	*/
	GetLastError(Instance LibPrimesHandle) (string, bool, error)


	/**
	* Releases the memory of an Instance
	*
	* @param[in] Wrapper - Wrapper instance.
	* @param[in] Instance - Instance Handle
	*/
	ReleaseInstance(Instance LibPrimesHandle) (error)


	/**
	* Creates a new FactorizationCalculator instance
	*
	* @param[in] Wrapper - Wrapper instance.
	* @return New FactorizationCalculator instance
	*/
	CreateFactorizationCalculator() (LibPrimesHandle, error)


	/**
	* Creates a new SieveCalculator instance
	*
	* @param[in] Wrapper - Wrapper instance.
	* @return New SieveCalculator instance
	*/
	CreateSieveCalculator() (LibPrimesHandle, error)


	/**
	* Handles Library Journaling
	*
	* @param[in] Wrapper - Wrapper instance.
	* @param[in] sFileName - Journal FileName
	*/
	SetJournal(sFileName string) (error)


}


/*************************************************************************************************************************
Class definition LibPrimesBase
**************************************************************************************************************************/

type LibPrimesBase struct {
	Interface LibPrimesGoInterface
	Handle LibPrimesHandle
}

func (instance *LibPrimesBase) Close() (error) {
	return instance.Handle.Close()
}


/*************************************************************************************************************************
Class definition LibPrimesCalculator
**************************************************************************************************************************/

type LibPrimesCalculator struct {
	LibPrimesBase
}

func (instance *LibPrimesCalculator) Close() (error) {
	return instance.Handle.Close()
}

func (instance *LibPrimesCalculator) GetValue() (uint64, error) {
	nValue, error := instance.Interface.Calculator_GetValue (instance.Handle)
	return nValue, error
}

func (instance *LibPrimesCalculator) SetValue(nValue uint64) (error) {
	error := instance.Interface.Calculator_SetValue (instance.Handle, nValue)
	return error
}

func (instance *LibPrimesCalculator) Calculate() (error) {
	error := instance.Interface.Calculator_Calculate (instance.Handle)
	return error
}

func (instance *LibPrimesCalculator) SetProgressCallback(pProgressCallback int64) (error) {
	error := instance.Interface.Calculator_SetProgressCallback (instance.Handle, pProgressCallback)
	return error
}


/*************************************************************************************************************************
Class definition LibPrimesFactorizationCalculator
**************************************************************************************************************************/

type LibPrimesFactorizationCalculator struct {
	LibPrimesCalculator
}

func (instance *LibPrimesFactorizationCalculator) Close() (error) {
	return instance.Handle.Close()
}

func (instance *LibPrimesFactorizationCalculator) GetPrimeFactors() ([]sLibPrimesPrimeFactor, error) {
	arrayPrimeFactors, error := instance.Interface.FactorizationCalculator_GetPrimeFactors (instance.Handle)
	return arrayPrimeFactors, error
}


/*************************************************************************************************************************
Class definition LibPrimesSieveCalculator
**************************************************************************************************************************/

type LibPrimesSieveCalculator struct {
	LibPrimesCalculator
}

func (instance *LibPrimesSieveCalculator) Close() (error) {
	return instance.Handle.Close()
}

func (instance *LibPrimesSieveCalculator) GetPrimes() ([]uint64, error) {
	bufferPrimes, error := instance.Interface.SieveCalculator_GetPrimes (instance.Handle)
	return bufferPrimes, error
}

func (instance *LibPrimesWrapper) GetVersion() (uint32, uint32, uint32, error) {
	nMajor, nMinor, nMicro, error := instance.Interface.GetVersion ()
	return nMajor, nMinor, nMicro, error
}

func (instance *LibPrimesWrapper) GetLastError(Instance LibPrimesHandle) (string, bool, error) {
	sErrorMessage, bHasError, error := instance.Interface.GetLastError (Instance)
	return sErrorMessage, bHasError, error
}

func (instance *LibPrimesWrapper) ReleaseInstance(Instance LibPrimesHandle) (error) {
	error := instance.Interface.ReleaseInstance (Instance)
	return error
}

func (instance *LibPrimesWrapper) CreateFactorizationCalculator() (LibPrimesFactorizationCalculator, error) {
	hInstance, error := instance.Interface.CreateFactorizationCalculator ()
	var cInstance LibPrimesFactorizationCalculator
	cInstance.Interface = instance.Interface
	cInstance.Handle = hInstance
	return cInstance, error
}

func (instance *LibPrimesWrapper) CreateSieveCalculator() (LibPrimesSieveCalculator, error) {
	hInstance, error := instance.Interface.CreateSieveCalculator ()
	var cInstance LibPrimesSieveCalculator
	cInstance.Interface = instance.Interface
	cInstance.Handle = hInstance
	return cInstance, error
}

func (instance *LibPrimesWrapper) SetJournal(sFileName string) (error) {
	error := instance.Interface.SetJournal (sFileName)
	return error
}

