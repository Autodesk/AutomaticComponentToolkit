(*++

Copyright (C) 2019 Numbers developers

All rights reserved.

Abstract: This is the class declaration of TNumbersVariable

*)

{$MODE DELPHI}
unit numbers_impl_variable;

interface

uses
  numbers_types,
  numbers_interfaces,
  numbers_exception,
  numbers_impl_base,
  Classes,
  sysutils;

type
  TNumbersVariable = class(TNumbersBase, INumbersVariable)
    private
      FValue : double;
    protected

    public
      function ClassTypeId(): QWord; Override;
	  constructor Create(AInitialValue: double);
      function GetValue(): Double;
      procedure SetValue(const AValue: Double);
  end;

implementation

function TNumbersVariable.ClassTypeId(): QWord;
begin
  Result := QWord($23934EDF762423EA); // First 64 bits of SHA1 of a string: "Numbers::Variable"
end;

constructor TNumbersVariable.Create(AInitialValue: double);
begin
  inherited Create();
  FValue := AInitialValue;
end;

function TNumbersVariable.GetValue(): Double;
begin
  result := FValue;
end;

procedure TNumbersVariable.SetValue(const AValue: Double);
begin
  FValue := AValue;
end;

end.
