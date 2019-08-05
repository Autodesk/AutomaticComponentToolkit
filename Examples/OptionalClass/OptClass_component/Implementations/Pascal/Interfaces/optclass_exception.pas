(*++

Copyright (C) 2019 ACT Developers


This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0.

Abstract: This is an autogenerated Pascal exception class definition file in order to allow easy
development of Optional Class Library. The functions in this file need to be implemented. It needs to be generated only once.

Interface version: 1.0.0

*)

{$MODE DELPHI}
unit optclass_exception;

interface

uses
  optclass_types,
  optclass_interfaces,
  Classes,
  sysutils;

type
  EOptClassException = class(Exception)
  private
    FErrorCode: TOptClassResult;
    FCustomMessage: String;
  public
    property ErrorCode: TOptClassResult read FErrorCode;
    property CustomMessage: String read FCustomMessage;
    constructor Create(AErrorCode: TOptClassResult);
    constructor CreateCustomMessage(AErrorCode: TOptClassResult; AMessage: String);
  end;


(*************************************************************************************************************************
 Definition of exception handling functionality for OptClass
**************************************************************************************************************************)

function HandleOptClassException(AOptClassObject: TObject; E: EOptClassException): TOptClassResult;
function HandleStdException(AOptClassObject: TObject; E: Exception): TOptClassResult;
function HandleUnhandledException(AOptClassObject: TObject): TOptClassResult;


implementation

  constructor EOptClassException.Create(AErrorCode: TOptClassResult);
  var
    ADescription: String;
  begin
    FErrorCode := AErrorCode;
    case FErrorCode of
      OPTCLASS_ERROR_NOTIMPLEMENTED: ADescription := 'functionality not implemented';
      OPTCLASS_ERROR_INVALIDPARAM: ADescription := 'an invalid parameter was passed';
      OPTCLASS_ERROR_INVALIDCAST: ADescription := 'a type cast failed';
      OPTCLASS_ERROR_BUFFERTOOSMALL: ADescription := 'a provided buffer is too small';
      OPTCLASS_ERROR_GENERICEXCEPTION: ADescription := 'a generic exception occurred';
      OPTCLASS_ERROR_COULDNOTLOADLIBRARY: ADescription := 'the library could not be loaded';
      OPTCLASS_ERROR_COULDNOTFINDLIBRARYEXPORT: ADescription := 'a required exported symbol could not be found in the library';
      OPTCLASS_ERROR_INCOMPATIBLEBINARYVERSION: ADescription := 'the version of the binary interface does not match the bindings interface';
      else
        ADescription := 'unknown';
    end;

    inherited Create(Format('Optional Class Library Error - %s (#%d)', [ ADescription, AErrorCode ]));
  end;

  constructor EOptClassException.CreateCustomMessage(AErrorCode: TOptClassResult; AMessage: String);
  begin
    FCustomMessage := AMessage;
    FErrorCode := AErrorCode;
    inherited Create(Format('%s(%d)', [FCustomMessage, AErrorCode]));
  end;

(*************************************************************************************************************************
 Implementation of exception handling functionality for OptClass
**************************************************************************************************************************)

function HandleOptClassException(AOptClassObject: TObject; E: EOptClassException): TOptClassResult;
begin
  result := E.ErrorCode;
  if Supports(AOptClassObject, IOptClassBase) then begin
    (AOptClassObject as IOptClassBase).RegisterErrorMessage(E.CustomMessage)
  end;
end;
function HandleStdException(AOptClassObject: TObject; E: Exception): TOptClassResult;
begin
  Result := OPTCLASS_ERROR_GENERICEXCEPTION;
  if Supports(AOptClassObject, IOptClassBase) then begin
    (AOptClassObject as IOptClassBase).RegisterErrorMessage(E.Message)
  end;
end;
function HandleUnhandledException(AOptClassObject: TObject): TOptClassResult;
begin
  Result := OPTCLASS_ERROR_GENERICEXCEPTION;
  if Supports(AOptClassObject, IOptClassBase) then begin
    (AOptClassObject as IOptClassBase).RegisterErrorMessage('Unhandled Exception')
  end;
end;
end.