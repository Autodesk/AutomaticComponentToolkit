(*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is the class declaration of TCalculatorVariable

*)

{$MODE DELPHI}
unit calculator_impl_variable;

interface

uses
  calculator_types,
  calculator_interfaces,
  calculator_exception,
  calculator_impl_base,
  Classes,
  sysutils;

type
  TCalculatorVariable = class(TCalculatorBase, ICalculatorVariable)
    private
      FValue: Double;

    protected

    public
      constructor Create(const AInitialValue: Double);
      function ClassTypeId(): QWord; Override;
      function GetValue(): Double;
      procedure SetValue(const AValue: Double);
  end;

implementation

constructor TCalculatorVariable.Create(const AInitialValue: Double);
begin
  inherited Create();
  FValue := AInitialValue;
end;

function TCalculatorVariable.ClassTypeId(): QWord;
begin
  Result := QWord($C228BBEB32E22A34); // First 64 bits of SHA1 of a string: "Calculator::Variable"
end;

function TCalculatorVariable.GetValue(): Double;
begin
  Result := FValue;
end;

procedure TCalculatorVariable.SetValue(const AValue: Double);
begin
  FValue := AValue;
end;

end.
