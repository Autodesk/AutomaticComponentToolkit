(*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.3.2.

Abstract: This is an autogenerated Pascal export implementation file in order to allow easy
development of Prime Numbers Library. The functions in this file need to be implemented. It needs to be generated only once.

Interface version: 1.2.0

*)

{$MODE DELPHI}
unit libprimes_exports;

interface

uses
	libprimes_impl,
	libprimes_types,
	libprimes_interfaces,
	libprimes_exception,
	Classes,
	sysutils;

(*************************************************************************************************************************
 Class export definition of Calculator 
**************************************************************************************************************************)

(**
* Returns the current value of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @param[out] pValue - The current value of this Calculator
* @return error code or 0 (success)
*)
function libprimes_calculator_getvalue (pCalculator: TLibPrimesHandle; pValue: PQWord): TLibPrimesResult; cdecl;

(**
* Sets the value to be factorized
*
* @param[in] pCalculator - Calculator instance.
* @param[in] nValue - The value to be factorized
* @return error code or 0 (success)
*)
function libprimes_calculator_setvalue (pCalculator: TLibPrimesHandle; nValue: QWord): TLibPrimesResult; cdecl;

(**
* Performs the specific calculation of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @return error code or 0 (success)
*)
function libprimes_calculator_calculate (pCalculator: TLibPrimesHandle): TLibPrimesResult; cdecl;

(**
* Sets the progress callback function
*
* @param[in] pCalculator - Calculator instance.
* @param[in] pProgressCallback - The progress callback
* @return error code or 0 (success)
*)
function libprimes_calculator_setprogresscallback (pCalculator: TLibPrimesHandle; pProgressCallback: PLibPrimes_ProgressCallback): TLibPrimesResult; cdecl;

(*************************************************************************************************************************
 Class export definition of FactorizationCalculator 
**************************************************************************************************************************)

(**
* Returns the prime factors of this number (without multiplicity)
*
* @param[in] pFactorizationCalculator - FactorizationCalculator instance.
* @param[in] nPrimeFactorsCount - Number of elements in buffer
* @param[out] pPrimeFactorsNeededCount - will be filled with the count of the written elements, or needed buffer size.
* @param[out] pPrimeFactorsBuffer - PrimeFactor buffer of The prime factors of this number
* @return error code or 0 (success)
*)
function libprimes_factorizationcalculator_getprimefactors (pFactorizationCalculator: TLibPrimesHandle; nPrimeFactorsCount: QWord; pPrimeFactorsNeededCount: PQWord; pPrimeFactorsBuffer: PLibPrimesPrimeFactor): TLibPrimesResult; cdecl;

(*************************************************************************************************************************
 Class export definition of SieveCalculator 
**************************************************************************************************************************)

(**
* Returns all prime numbers lower or equal to the sieve's value
*
* @param[in] pSieveCalculator - SieveCalculator instance.
* @param[in] nPrimesCount - Number of elements in buffer
* @param[out] pPrimesNeededCount - will be filled with the count of the written elements, or needed buffer size.
* @param[out] pPrimesBuffer - uint64 buffer of The primes lower or equal to the sieve's value
* @return error code or 0 (success)
*)
function libprimes_sievecalculator_getprimes (pSieveCalculator: TLibPrimesHandle; nPrimesCount: QWord; pPrimesNeededCount: PQWord; pPrimesBuffer: PQWord): TLibPrimesResult; cdecl;

(*************************************************************************************************************************
 Global function export definition
**************************************************************************************************************************)

(**
* Releases the memory of an Instance
*
* @param[in] pInstance - Instance Handle
* @return error code or 0 (success)
*)
function libprimes_releaseinstance (pInstance: TLibPrimesHandle): TLibPrimesResult; cdecl;

(**
* retrieves the current version of the library.
*
* @param[out] pMajor - returns the major version of the library
* @param[out] pMinor - returns the minor version of the library
* @param[out] pMicro - returns the micro version of the library
* @return error code or 0 (success)
*)
function libprimes_getlibraryversion (pMajor: PCardinal; pMinor: PCardinal; pMicro: PCardinal): TLibPrimesResult; cdecl;

(**
* Creates a new FactorizationCalculator instance
*
* @param[out] pInstance - New FactorizationCalculator instance
* @return error code or 0 (success)
*)
function libprimes_createfactorizationcalculator (pInstance: PLibPrimesHandle): TLibPrimesResult; cdecl;

(**
* Creates a new SieveCalculator instance
*
* @param[out] pInstance - New SieveCalculator instance
* @return error code or 0 (success)
*)
function libprimes_createsievecalculator (pInstance: PLibPrimesHandle): TLibPrimesResult; cdecl;

(**
* Handles Library Journaling
*
* @param[in] pFileName - Journal FileName
* @return error code or 0 (success)
*)
function libprimes_setjournal (pFileName: PAnsiChar): TLibPrimesResult; cdecl;

implementation

function libprimes_calculator_getvalue (pCalculator: TLibPrimesHandle; pValue: PQWord): TLibPrimesResult; cdecl;
var
	ResultValue: QWord;
	ObjectCalculator: TObject;
	IntfCalculator: ILibPrimesCalculator;
begin
	try
		if not Assigned (pValue) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);
		if not Assigned (pCalculator) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ObjectCalculator := TObject (pCalculator);
		if Supports (ObjectCalculator, ILibPrimesCalculator) then begin
			IntfCalculator := ObjectCalculator as ILibPrimesCalculator;
			ResultValue := IntfCalculator.GetValue();

			pValue^ := ResultValue;
		end else
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_calculator_setvalue (pCalculator: TLibPrimesHandle; nValue: QWord): TLibPrimesResult; cdecl;
var
	ObjectCalculator: TObject;
	IntfCalculator: ILibPrimesCalculator;
begin
	try
		if not Assigned (pCalculator) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ObjectCalculator := TObject (pCalculator);
		if Supports (ObjectCalculator, ILibPrimesCalculator) then begin
			IntfCalculator := ObjectCalculator as ILibPrimesCalculator;
			IntfCalculator.SetValue(nValue);

		end else
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_calculator_calculate (pCalculator: TLibPrimesHandle): TLibPrimesResult; cdecl;
var
	ObjectCalculator: TObject;
	IntfCalculator: ILibPrimesCalculator;
begin
	try
		if not Assigned (pCalculator) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ObjectCalculator := TObject (pCalculator);
		if Supports (ObjectCalculator, ILibPrimesCalculator) then begin
			IntfCalculator := ObjectCalculator as ILibPrimesCalculator;
			IntfCalculator.Calculate();

		end else
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_calculator_setprogresscallback (pCalculator: TLibPrimesHandle; pProgressCallback: PLibPrimes_ProgressCallback): TLibPrimesResult; cdecl;
var
	ObjectCalculator: TObject;
	IntfCalculator: ILibPrimesCalculator;
begin
	try
		if not Assigned (pCalculator) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ObjectCalculator := TObject (pCalculator);
		if Supports (ObjectCalculator, ILibPrimesCalculator) then begin
			IntfCalculator := ObjectCalculator as ILibPrimesCalculator;
			IntfCalculator.SetProgressCallback(pProgressCallback);

		end else
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_factorizationcalculator_getprimefactors (pFactorizationCalculator: TLibPrimesHandle; nPrimeFactorsCount: QWord; pPrimeFactorsNeededCount: PQWord; pPrimeFactorsBuffer: PLibPrimesPrimeFactor): TLibPrimesResult; cdecl;
var
	ObjectFactorizationCalculator: TObject;
	IntfFactorizationCalculator: ILibPrimesFactorizationCalculator;
begin
	try
		if ((not Assigned (pPrimeFactorsNeededCount)) and (not Assigned(pPrimeFactorsBuffer))) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);
		if not Assigned (pFactorizationCalculator) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ObjectFactorizationCalculator := TObject (pFactorizationCalculator);
		if Supports (ObjectFactorizationCalculator, ILibPrimesFactorizationCalculator) then begin
			IntfFactorizationCalculator := ObjectFactorizationCalculator as ILibPrimesFactorizationCalculator;
			IntfFactorizationCalculator.GetPrimeFactors(nPrimeFactorsCount, pPrimeFactorsNeededCount, pPrimeFactorsBuffer);

		end else
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_sievecalculator_getprimes (pSieveCalculator: TLibPrimesHandle; nPrimesCount: QWord; pPrimesNeededCount: PQWord; pPrimesBuffer: PQWord): TLibPrimesResult; cdecl;
var
	ObjectSieveCalculator: TObject;
	IntfSieveCalculator: ILibPrimesSieveCalculator;
begin
	try
		if ((not Assigned (pPrimesNeededCount)) and (not Assigned(pPrimesBuffer))) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);
		if not Assigned (pSieveCalculator) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ObjectSieveCalculator := TObject (pSieveCalculator);
		if Supports (ObjectSieveCalculator, ILibPrimesSieveCalculator) then begin
			IntfSieveCalculator := ObjectSieveCalculator as ILibPrimesSieveCalculator;
			IntfSieveCalculator.GetPrimes(nPrimesCount, pPrimesNeededCount, pPrimesBuffer);

		end else
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_releaseinstance (pInstance: TLibPrimesHandle): TLibPrimesResult; cdecl;
var
	ObjectInstance: TObject;
begin
	try
		ObjectInstance := TObject (pInstance);
		if (not Supports (ObjectInstance, ILibPrimesBaseClass)) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDCAST);
		

		TLibPrimesWrapper.ReleaseInstance(ObjectInstance);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_getlibraryversion (pMajor: PCardinal; pMinor: PCardinal; pMicro: PCardinal): TLibPrimesResult; cdecl;
begin
	try
		if (not Assigned (pMajor)) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		if (not Assigned (pMinor)) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		if (not Assigned (pMicro)) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);


		TLibPrimesWrapper.GetLibraryVersion(pMajor^, pMinor^, pMicro^);

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_createfactorizationcalculator (pInstance: PLibPrimesHandle): TLibPrimesResult; cdecl;
var
	ResultInstance: TObject;
begin
	try
		if not Assigned(pInstance) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ResultInstance := TLibPrimesWrapper.CreateFactorizationCalculator();

		pInstance^ := ResultInstance;
		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_createsievecalculator (pInstance: PLibPrimesHandle): TLibPrimesResult; cdecl;
var
	ResultInstance: TObject;
begin
	try
		if not Assigned(pInstance) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		ResultInstance := TLibPrimesWrapper.CreateSieveCalculator();

		pInstance^ := ResultInstance;
		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;

function libprimes_setjournal (pFileName: PAnsiChar): TLibPrimesResult; cdecl;
begin
	try
		if (not Assigned (pFileName)) then
			raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

		TLibPrimesWrapper.SetJournal(StrPas (pFileName));

		Result := LIBPRIMES_SUCCESS;
	except
		On E: ELibPrimesException do begin
			Result := E.ErrorCode;
		end;
		On E: Exception do begin
			Result := LIBPRIMES_ERROR_GENERICEXCEPTION;
		end;
	end;
end;


end.

