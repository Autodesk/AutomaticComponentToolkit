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
      FValue : double;
    protected

    public
      function GetValue(): Double;
      procedure SetValue(const AValue: Double);
  end;

implementation

function TCalculatorVariable.GetValue(): Double;
begin
  result := FValue;
end;

procedure TCalculatorVariable.SetValue(const AValue: Double);
begin
  FValue := AValue;
end;

end.
