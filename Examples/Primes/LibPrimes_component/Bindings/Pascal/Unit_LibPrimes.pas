{$IFDEF FPC}{$MODE DELPHI}{$ENDIF}
(*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.3.2.

Abstract: This is an autogenerated Pascal Header file in order to allow an easy
 use of Prime Numbers Library

Interface version: 1.2.0

*)

unit Unit_LibPrimes;

interface

uses
    {$IFDEF WINDOWS}
        Windows,
    {$ELSE}
        dynlibs,
    {$ENDIF}
    Types,
    Classes,
    SysUtils;

(*************************************************************************************************************************
 Version definition for LibPrimes
**************************************************************************************************************************)

const
    LIBPRIMES_VERSION_MAJOR = 1;
    LIBPRIMES_VERSION_MINOR = 2;
    LIBPRIMES_VERSION_MICRO = 0;


(*************************************************************************************************************************
 General type definitions
**************************************************************************************************************************)

type
    TLibPrimesResult = Cardinal;
    TLibPrimesHandle = Pointer;

    PLibPrimesResult = ^TLibPrimesResult;
    PLibPrimesHandle = ^TLibPrimesHandle;

(*************************************************************************************************************************
 Error Constants for LibPrimes
**************************************************************************************************************************)

const
    LIBPRIMES_SUCCESS = 0;
    LIBPRIMES_ERROR_NOTIMPLEMENTED = 1;
    LIBPRIMES_ERROR_INVALIDPARAM = 2;
    LIBPRIMES_ERROR_INVALIDCAST = 3;
    LIBPRIMES_ERROR_BUFFERTOOSMALL = 4;
    LIBPRIMES_ERROR_GENERICEXCEPTION = 5;
    LIBPRIMES_ERROR_COULDNOTLOADLIBRARY = 6;
    LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT = 7;
    LIBPRIMES_ERROR_INCOMPATIBLEBINARYVERSION = 8;
    LIBPRIMES_ERROR_NORESULTAVAILABLE = 9;
    LIBPRIMES_ERROR_CALCULATIONABORTED = 10;

(*************************************************************************************************************************
 Declaration of structs
**************************************************************************************************************************)

type

    PLibPrimesPrimeFactor = ^TLibPrimesPrimeFactor;
    TLibPrimesPrimeFactor = packed record
        FPrime: QWord;
        FMultiplicity: Cardinal;
    end;


(*************************************************************************************************************************
 Declaration of struct arrays
**************************************************************************************************************************)

    ArrayOfLibPrimesPrimeFactor = array of TLibPrimesPrimeFactor;

(*************************************************************************************************************************
 Declaration of function types
**************************************************************************************************************************)

type

    PLibPrimes_ProgressCallback = function(const fProgressPercentage: Single; out pShouldAbort: Cardinal): Integer; cdecl;

(*************************************************************************************************************************
 Declaration of handle classes 
**************************************************************************************************************************)

type
    TLibPrimesBaseClass = class;
    TLibPrimesWrapper = class;
    TLibPrimesCalculator = class;
    TLibPrimesFactorizationCalculator = class;
    TLibPrimesSieveCalculator = class;


(*************************************************************************************************************************
 Function type definitions for Calculator
**************************************************************************************************************************)

    (**
    * Returns the current value of this Calculator
    *
    * @param[in] pCalculator - Calculator instance.
    * @param[out] pValue - The current value of this Calculator
    * @return error code or 0 (success)
    *)
    TLibPrimesCalculator_GetValueFunc = function (pCalculator: TLibPrimesHandle; out pValue: QWord): TLibPrimesResult; cdecl;
    
    (**
    * Sets the value to be factorized
    *
    * @param[in] pCalculator - Calculator instance.
    * @param[in] nValue - The value to be factorized
    * @return error code or 0 (success)
    *)
    TLibPrimesCalculator_SetValueFunc = function (pCalculator: TLibPrimesHandle; const nValue: QWord): TLibPrimesResult; cdecl;
    
    (**
    * Performs the specific calculation of this Calculator
    *
    * @param[in] pCalculator - Calculator instance.
    * @return error code or 0 (success)
    *)
    TLibPrimesCalculator_CalculateFunc = function (pCalculator: TLibPrimesHandle): TLibPrimesResult; cdecl;
    
    (**
    * Sets the progress callback function
    *
    * @param[in] pCalculator - Calculator instance.
    * @param[in] pProgressCallback - The progress callback
    * @return error code or 0 (success)
    *)
    TLibPrimesCalculator_SetProgressCallbackFunc = function (pCalculator: TLibPrimesHandle; const pProgressCallback: PLibPrimes_ProgressCallback): TLibPrimesResult; cdecl;
    

(*************************************************************************************************************************
 Function type definitions for FactorizationCalculator
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
    TLibPrimesFactorizationCalculator_GetPrimeFactorsFunc = function (pFactorizationCalculator: TLibPrimesHandle; const nPrimeFactorsCount: QWord; out pPrimeFactorsNeededCount: QWord; pPrimeFactorsBuffer: PLibPrimesPrimeFactor): TLibPrimesResult; cdecl;
    

(*************************************************************************************************************************
 Function type definitions for SieveCalculator
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
    TLibPrimesSieveCalculator_GetPrimesFunc = function (pSieveCalculator: TLibPrimesHandle; const nPrimesCount: QWord; out pPrimesNeededCount: QWord; pPrimesBuffer: PQWord): TLibPrimesResult; cdecl;
    
(*************************************************************************************************************************
 Global function definitions 
**************************************************************************************************************************)

    (**
    * Releases the memory of an Instance
    *
    * @param[in] pInstance - Instance Handle
    * @return error code or 0 (success)
    *)
    TLibPrimesReleaseInstanceFunc = function (const pInstance: TLibPrimesHandle): TLibPrimesResult; cdecl;
    
    (**
    * retrieves the current version of the library.
    *
    * @param[out] pMajor - returns the major version of the library
    * @param[out] pMinor - returns the minor version of the library
    * @param[out] pMicro - returns the micro version of the library
    * @return error code or 0 (success)
    *)
    TLibPrimesGetLibraryVersionFunc = function (out pMajor: Cardinal; out pMinor: Cardinal; out pMicro: Cardinal): TLibPrimesResult; cdecl;
    
    (**
    * Creates a new FactorizationCalculator instance
    *
    * @param[out] pInstance - New FactorizationCalculator instance
    * @return error code or 0 (success)
    *)
    TLibPrimesCreateFactorizationCalculatorFunc = function (out pInstance: TLibPrimesHandle): TLibPrimesResult; cdecl;
    
    (**
    * Creates a new SieveCalculator instance
    *
    * @param[out] pInstance - New SieveCalculator instance
    * @return error code or 0 (success)
    *)
    TLibPrimesCreateSieveCalculatorFunc = function (out pInstance: TLibPrimesHandle): TLibPrimesResult; cdecl;
    
    (**
    * Handles Library Journaling
    *
    * @param[in] pFileName - Journal FileName
    * @return error code or 0 (success)
    *)
    TLibPrimesSetJournalFunc = function (const pFileName: PAnsiChar): TLibPrimesResult; cdecl;
    
(*************************************************************************************************************************
 Exception definition
**************************************************************************************************************************)

    ELibPrimesException = class (Exception)
    private
        FErrorCode: TLibPrimesResult;
        FCustomMessage: String;
    public
        property ErrorCode: TLibPrimesResult read FErrorCode;
        property CustomMessage: String read FCustomMessage;
        constructor Create (AErrorCode: TLibPrimesResult);
        constructor CreateCustomMessage (AErrorCode: TLibPrimesResult; AMessage: String);
    end;

(*************************************************************************************************************************
 Base class definition
**************************************************************************************************************************)

    TLibPrimesBaseClass = class (TObject)
    private
        FWrapper: TLibPrimesWrapper;
        FHandle: TLibPrimesHandle;
    public
        constructor Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
        destructor Destroy; override;
    end;


(*************************************************************************************************************************
 Class definition for Calculator
**************************************************************************************************************************)

    TLibPrimesCalculator = class (TLibPrimesBaseClass)
    private
    public
        constructor Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
        destructor Destroy; override;
        function GetValue(): QWord;
        procedure SetValue(const AValue: QWord);
        procedure Calculate();
        procedure SetProgressCallback(const AProgressCallback: PLibPrimes_ProgressCallback);
    end;


(*************************************************************************************************************************
 Class definition for FactorizationCalculator
**************************************************************************************************************************)

    TLibPrimesFactorizationCalculator = class (TLibPrimesCalculator)
    private
    public
        constructor Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
        destructor Destroy; override;
        procedure GetPrimeFactors(out APrimeFactors: ArrayOfLibPrimesPrimeFactor);
    end;


(*************************************************************************************************************************
 Class definition for SieveCalculator
**************************************************************************************************************************)

    TLibPrimesSieveCalculator = class (TLibPrimesCalculator)
    private
    public
        constructor Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
        destructor Destroy; override;
        procedure GetPrimes(out APrimes: TQWordDynArray);
    end;

(*************************************************************************************************************************
 Wrapper definition
**************************************************************************************************************************)

    TLibPrimesWrapper = class (TObject)
    private
        FModule: HMODULE;
        FLibPrimesCalculator_GetValueFunc: TLibPrimesCalculator_GetValueFunc;
        FLibPrimesCalculator_SetValueFunc: TLibPrimesCalculator_SetValueFunc;
        FLibPrimesCalculator_CalculateFunc: TLibPrimesCalculator_CalculateFunc;
        FLibPrimesCalculator_SetProgressCallbackFunc: TLibPrimesCalculator_SetProgressCallbackFunc;
        FLibPrimesFactorizationCalculator_GetPrimeFactorsFunc: TLibPrimesFactorizationCalculator_GetPrimeFactorsFunc;
        FLibPrimesSieveCalculator_GetPrimesFunc: TLibPrimesSieveCalculator_GetPrimesFunc;
        FLibPrimesReleaseInstanceFunc: TLibPrimesReleaseInstanceFunc;
        FLibPrimesGetLibraryVersionFunc: TLibPrimesGetLibraryVersionFunc;
        FLibPrimesCreateFactorizationCalculatorFunc: TLibPrimesCreateFactorizationCalculatorFunc;
        FLibPrimesCreateSieveCalculatorFunc: TLibPrimesCreateSieveCalculatorFunc;
        FLibPrimesSetJournalFunc: TLibPrimesSetJournalFunc;

        {$IFDEF MSWINDOWS}
        function LoadFunction (AFunctionName: AnsiString; FailIfNotExistent: Boolean = True): FARPROC;
        {$ELSE}
        function LoadFunction (AFunctionName: AnsiString; FailIfNotExistent: Boolean = True): Pointer;
        {$ENDIF MSWINDOWS}

        procedure checkBinaryVersion();

    protected
        property LibPrimesCalculator_GetValueFunc: TLibPrimesCalculator_GetValueFunc read FLibPrimesCalculator_GetValueFunc;
        property LibPrimesCalculator_SetValueFunc: TLibPrimesCalculator_SetValueFunc read FLibPrimesCalculator_SetValueFunc;
        property LibPrimesCalculator_CalculateFunc: TLibPrimesCalculator_CalculateFunc read FLibPrimesCalculator_CalculateFunc;
        property LibPrimesCalculator_SetProgressCallbackFunc: TLibPrimesCalculator_SetProgressCallbackFunc read FLibPrimesCalculator_SetProgressCallbackFunc;
        property LibPrimesFactorizationCalculator_GetPrimeFactorsFunc: TLibPrimesFactorizationCalculator_GetPrimeFactorsFunc read FLibPrimesFactorizationCalculator_GetPrimeFactorsFunc;
        property LibPrimesSieveCalculator_GetPrimesFunc: TLibPrimesSieveCalculator_GetPrimesFunc read FLibPrimesSieveCalculator_GetPrimesFunc;
        property LibPrimesReleaseInstanceFunc: TLibPrimesReleaseInstanceFunc read FLibPrimesReleaseInstanceFunc;
        property LibPrimesGetLibraryVersionFunc: TLibPrimesGetLibraryVersionFunc read FLibPrimesGetLibraryVersionFunc;
        property LibPrimesCreateFactorizationCalculatorFunc: TLibPrimesCreateFactorizationCalculatorFunc read FLibPrimesCreateFactorizationCalculatorFunc;
        property LibPrimesCreateSieveCalculatorFunc: TLibPrimesCreateSieveCalculatorFunc read FLibPrimesCreateSieveCalculatorFunc;
        property LibPrimesSetJournalFunc: TLibPrimesSetJournalFunc read FLibPrimesSetJournalFunc;
        procedure CheckError (AInstance: TLibPrimesBaseClass; AErrorCode: TLibPrimesResult);
    public
        constructor Create (ADLLName: String);
        destructor Destroy; override;
        procedure ReleaseInstance(const AInstance: TLibPrimesBaseClass);
        procedure GetLibraryVersion(out AMajor: Cardinal; out AMinor: Cardinal; out AMicro: Cardinal);
        function CreateFactorizationCalculator(): TLibPrimesFactorizationCalculator;
        function CreateSieveCalculator(): TLibPrimesSieveCalculator;
        procedure SetJournal(const AFileName: String);
    end;


implementation


(*************************************************************************************************************************
 Exception implementation
**************************************************************************************************************************)

    constructor ELibPrimesException.Create (AErrorCode: TLibPrimesResult);
    var
        ADescription: String;
    begin
        FErrorCode := AErrorCode;
        case FErrorCode of
            LIBPRIMES_ERROR_NOTIMPLEMENTED: ADescription := 'functionality not implemented';
            LIBPRIMES_ERROR_INVALIDPARAM: ADescription := 'an invalid parameter was passed';
            LIBPRIMES_ERROR_INVALIDCAST: ADescription := 'a type cast failed';
            LIBPRIMES_ERROR_BUFFERTOOSMALL: ADescription := 'a provided buffer is too small';
            LIBPRIMES_ERROR_GENERICEXCEPTION: ADescription := 'a generic exception occurred';
            LIBPRIMES_ERROR_COULDNOTLOADLIBRARY: ADescription := 'the library could not be loaded';
            LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT: ADescription := 'a required exported symbol could not be found in the library';
            LIBPRIMES_ERROR_INCOMPATIBLEBINARYVERSION: ADescription := 'the version of the binary interface does not match the bindings interface';
            LIBPRIMES_ERROR_NORESULTAVAILABLE: ADescription := 'no result is available';
            LIBPRIMES_ERROR_CALCULATIONABORTED: ADescription := 'a calculation has been aborted';
            else
                ADescription := 'unknown';
        end;

        inherited Create (Format ('Prime Numbers Library Error - %s (#%d)', [ ADescription, AErrorCode ]));
    end;

    constructor ELibPrimesException.CreateCustomMessage (AErrorCode: TLibPrimesResult; AMessage: String);
    begin
        FCustomMessage := AMessage;
        FErrorCode := AErrorCode;
        inherited Create (Format ('%s (%d)', [FCustomMessage, AErrorCode]));
    end;

(*************************************************************************************************************************
 Base class implementation
**************************************************************************************************************************)

    constructor TLibPrimesBaseClass.Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
    begin
        if not Assigned (AWrapper) then
            raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);
        if not Assigned (AHandle) then
            raise ELibPrimesException.Create (LIBPRIMES_ERROR_INVALIDPARAM);

        inherited Create ();
        FWrapper := AWrapper;
        FHandle := AHandle;
    end;

    destructor TLibPrimesBaseClass.Destroy;
    begin
        FWrapper.ReleaseInstance(self);
        inherited;
    end;

(*************************************************************************************************************************
 Class implementation for Calculator
**************************************************************************************************************************)

    constructor TLibPrimesCalculator.Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
    begin
        inherited Create (AWrapper, AHandle);
    end;

    destructor TLibPrimesCalculator.Destroy;
    begin
        inherited;
    end;

    function TLibPrimesCalculator.GetValue(): QWord;
    begin
        FWrapper.CheckError (Self, FWrapper.LibPrimesCalculator_GetValueFunc (FHandle, Result));
    end;

    procedure TLibPrimesCalculator.SetValue(const AValue: QWord);
    begin
        FWrapper.CheckError (Self, FWrapper.LibPrimesCalculator_SetValueFunc (FHandle, AValue));
    end;

    procedure TLibPrimesCalculator.Calculate();
    begin
        FWrapper.CheckError (Self, FWrapper.LibPrimesCalculator_CalculateFunc (FHandle));
    end;

    procedure TLibPrimesCalculator.SetProgressCallback(const AProgressCallback: PLibPrimes_ProgressCallback);
    begin
        if not Assigned (AProgressCallback) then
            raise ELibPrimesException.CreateCustomMessage (LIBPRIMES_ERROR_INVALIDPARAM, 'AProgressCallback is a nil value.');
        FWrapper.CheckError (Self, FWrapper.LibPrimesCalculator_SetProgressCallbackFunc (FHandle, AProgressCallback));
    end;

(*************************************************************************************************************************
 Class implementation for FactorizationCalculator
**************************************************************************************************************************)

    constructor TLibPrimesFactorizationCalculator.Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
    begin
        inherited Create (AWrapper, AHandle);
    end;

    destructor TLibPrimesFactorizationCalculator.Destroy;
    begin
        inherited;
    end;

    procedure TLibPrimesFactorizationCalculator.GetPrimeFactors(out APrimeFactors: ArrayOfLibPrimesPrimeFactor);
    var
        countNeededPrimeFactors: QWord;
        countWrittenPrimeFactors: QWord;
    begin
        countNeededPrimeFactors:= 0;
        countWrittenPrimeFactors:= 0;
        FWrapper.CheckError (Self, FWrapper.LibPrimesFactorizationCalculator_GetPrimeFactorsFunc (FHandle, 0, countNeededPrimeFactors, nil));
        SetLength (APrimeFactors, countNeededPrimeFactors);
        FWrapper.CheckError (Self, FWrapper.LibPrimesFactorizationCalculator_GetPrimeFactorsFunc (FHandle, countNeededPrimeFactors, countWrittenPrimeFactors, @APrimeFactors[0]));
    end;

(*************************************************************************************************************************
 Class implementation for SieveCalculator
**************************************************************************************************************************)

    constructor TLibPrimesSieveCalculator.Create (AWrapper: TLibPrimesWrapper; AHandle: TLibPrimesHandle);
    begin
        inherited Create (AWrapper, AHandle);
    end;

    destructor TLibPrimesSieveCalculator.Destroy;
    begin
        inherited;
    end;

    procedure TLibPrimesSieveCalculator.GetPrimes(out APrimes: TQWordDynArray);
    var
        countNeededPrimes: QWord;
        countWrittenPrimes: QWord;
    begin
        countNeededPrimes:= 0;
        countWrittenPrimes:= 0;
        FWrapper.CheckError (Self, FWrapper.LibPrimesSieveCalculator_GetPrimesFunc (FHandle, 0, countNeededPrimes, nil));
        SetLength (APrimes, countNeededPrimes);
        FWrapper.CheckError (Self, FWrapper.LibPrimesSieveCalculator_GetPrimesFunc (FHandle, countNeededPrimes, countWrittenPrimes, @APrimes[0]));
    end;

(*************************************************************************************************************************
 Wrapper class implementation
**************************************************************************************************************************)

    constructor TLibPrimesWrapper.Create (ADLLName: String);
    {$IFDEF MSWINDOWS}
    var
        AWideString: WideString;
    {$ENDIF MSWINDOWS}
    begin
        inherited Create;
        {$IFDEF MSWINDOWS}
            AWideString := UTF8Decode(ADLLName + #0);
            FModule := LoadLibraryW (PWideChar (AWideString));
        {$ELSE}
            FModule := dynlibs.LoadLibrary (ADLLName);
        {$ENDIF MSWINDOWS}
        if FModule = 0 then
            raise ELibPrimesException.Create (LIBPRIMES_ERROR_COULDNOTLOADLIBRARY);

        FLibPrimesCalculator_GetValueFunc := LoadFunction ('libprimes_calculator_getvalue');
        FLibPrimesCalculator_SetValueFunc := LoadFunction ('libprimes_calculator_setvalue');
        FLibPrimesCalculator_CalculateFunc := LoadFunction ('libprimes_calculator_calculate');
        FLibPrimesCalculator_SetProgressCallbackFunc := LoadFunction ('libprimes_calculator_setprogresscallback');
        FLibPrimesFactorizationCalculator_GetPrimeFactorsFunc := LoadFunction ('libprimes_factorizationcalculator_getprimefactors');
        FLibPrimesSieveCalculator_GetPrimesFunc := LoadFunction ('libprimes_sievecalculator_getprimes');
        FLibPrimesReleaseInstanceFunc := LoadFunction ('libprimes_releaseinstance');
        FLibPrimesGetLibraryVersionFunc := LoadFunction ('libprimes_getlibraryversion');
        FLibPrimesCreateFactorizationCalculatorFunc := LoadFunction ('libprimes_createfactorizationcalculator');
        FLibPrimesCreateSieveCalculatorFunc := LoadFunction ('libprimes_createsievecalculator');
        FLibPrimesSetJournalFunc := LoadFunction ('libprimes_setjournal');
        
        checkBinaryVersion();
    end;

    destructor TLibPrimesWrapper.Destroy;
    begin
        {$IFDEF MSWINDOWS}
            if FModule <> 0 then
                FreeLibrary (FModule);
        {$ELSE}
            if FModule <> 0 then
                UnloadLibrary (FModule);
        {$ENDIF MSWINDOWS}
        inherited;
    end;

    procedure TLibPrimesWrapper.CheckError (AInstance: TLibPrimesBaseClass; AErrorCode: TLibPrimesResult);
    begin
        if AInstance <> nil then begin
            if AInstance.FWrapper <> Self then
                raise ELibPrimesException.CreateCustomMessage (LIBPRIMES_ERROR_INVALIDCAST, 'invalid wrapper call');
        end;
        if AErrorCode <> LIBPRIMES_SUCCESS then
            raise ELibPrimesException.Create (AErrorCode);
    end;

    {$IFDEF MSWINDOWS}
    function TLibPrimesWrapper.LoadFunction (AFunctionName: AnsiString; FailIfNotExistent: Boolean): FARPROC;
    begin
        Result := GetProcAddress (FModule, PAnsiChar (AFunctionName));
        if FailIfNotExistent and not Assigned (Result) then
            raise ELibPrimesException.CreateCustomMessage (LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT, 'could not find function ' + AFunctionName);
    end;
    {$ELSE}
    function TLibPrimesWrapper.LoadFunction (AFunctionName: AnsiString; FailIfNotExistent: Boolean): Pointer;
    begin
        Result := dynlibs.GetProcAddress (FModule, AFunctionName);
        if FailIfNotExistent and not Assigned (Result) then
            raise ELibPrimesException.CreateCustomMessage (LIBPRIMES_ERROR_COULDNOTFINDLIBRARYEXPORT, 'could not find function ' + AFunctionName);
    end;
    {$ENDIF MSWINDOWS}

    procedure TLibPrimesWrapper.checkBinaryVersion();
    var
        AMajor, AMinor, AMicro: Cardinal;
    begin
        GetLibraryVersion(AMajor, AMinor, AMicro);
        if (AMajor <> LIBPRIMES_VERSION_MAJOR) or (AMinor < LIBPRIMES_VERSION_MINOR) then
            raise ELibPrimesException.Create(LIBPRIMES_ERROR_INCOMPATIBLEBINARYVERSION);
    end;
    
    procedure TLibPrimesWrapper.ReleaseInstance(const AInstance: TLibPrimesBaseClass);
    begin
        if not Assigned (AInstance) then
            raise ELibPrimesException.CreateCustomMessage (LIBPRIMES_ERROR_INVALIDPARAM, 'AInstance is a nil value.');
        CheckError (nil, LibPrimesReleaseInstanceFunc (AInstance.FHandle));
    end;

    procedure TLibPrimesWrapper.GetLibraryVersion(out AMajor: Cardinal; out AMinor: Cardinal; out AMicro: Cardinal);
    begin
        CheckError (nil, LibPrimesGetLibraryVersionFunc (AMajor, AMinor, AMicro));
    end;

    function TLibPrimesWrapper.CreateFactorizationCalculator(): TLibPrimesFactorizationCalculator;
    var
        HInstance: TLibPrimesHandle;
    begin
        Result := nil;
        HInstance := nil;
        CheckError (nil, LibPrimesCreateFactorizationCalculatorFunc (HInstance));
        if Assigned (HInstance) then
            Result := TLibPrimesFactorizationCalculator.Create (Self, HInstance);
    end;

    function TLibPrimesWrapper.CreateSieveCalculator(): TLibPrimesSieveCalculator;
    var
        HInstance: TLibPrimesHandle;
    begin
        Result := nil;
        HInstance := nil;
        CheckError (nil, LibPrimesCreateSieveCalculatorFunc (HInstance));
        if Assigned (HInstance) then
            Result := TLibPrimesSieveCalculator.Create (Self, HInstance);
    end;

    procedure TLibPrimesWrapper.SetJournal(const AFileName: String);
    begin
        CheckError (nil, LibPrimesSetJournalFunc (PAnsiChar (AFileName)));
    end;


end.