(*++

Copyright (C) 2019 Calculation developers

All rights reserved.

Abstract: This is the class declaration of TCalculationCalculator

*)

{$MODE DELPHI}
unit calculation_impl_calculator;

interface

uses
  calculation_types,
  calculation_interfaces,
  calculation_exception,
  calculation_impl_base,
  Unit_Numbers,
  Classes,
  sysutils,
  contnrs;

type
  TCalculationCalculator = class(TCalculationBase, ICalculationCalculator)
    private
      FVariableList : TObjectList;
    protected

    public
      constructor Create();
      destructor Destroy(); override;
      function ClassTypeId(): QWord; Override;
      procedure EnlistVariable(AVariable: TNumbersVariable);
      function GetEnlistedVariable(const AIndex: Cardinal): TNumbersVariable;
      procedure ClearVariables();
      function Multiply(): TNumbersVariable;
      function Add(): TNumbersVariable;
  end;

implementation

uses calculation_impl;
function TCalculationCalculator.ClassTypeId(): QWord;
begin
  Result := QWord($B23F514353D0C606); // First 64 bits of SHA1 of a string: "Calculation::Calculator"
end;

constructor TCalculationCalculator.Create();
begin
  inherited Create();

  FVariableList := TObjectList.Create(True);
end;

destructor TCalculationCalculator.Destroy();
begin
  ClearVariables();

  FreeAndNil(FVariableList);
  inherited Destroy();
end;

procedure TCalculationCalculator.EnlistVariable(AVariable: TNumbersVariable);
begin
  FVariableList.Add(AVariable);
end;

function TCalculationCalculator.GetEnlistedVariable(const AIndex: Cardinal): TNumbersVariable;
begin
  if AIndex >= FVariableList.Count then begin
    raise ECalculationException.Create(CALCULATION_ERROR_INVALIDPARAM);
  end;
  result := FVariableList[AIndex] as TNumbersVariable;
end;

procedure TCalculationCalculator.ClearVariables();
begin
  FVariableList.Clear;
end;

function TCalculationCalculator.Multiply(): TNumbersVariable;
var
  AVar: TNumbersVariable;
  I: integer;
  ResVal : double;
begin
  ResVal := 1.0;
  For I := 0 to FVariableList.Count - 1 do begin
    AVar := (FVariableList[I] as TNumbersVariable);
    ResVal := ResVal * AVar.GetValue();
  end;
  result := TCalculationWrapper.NumbersWrapper.CreateVariable(ResVal);
end;

function TCalculationCalculator.Add(): TNumbersVariable;
var
  AVar: TNumbersVariable;
  I: integer;
  ResVal : double;
begin
  ResVal := 0.0;
  For I := 0 to FVariableList.Count - 1 do begin
    AVar := (FVariableList[I] as TNumbersVariable);
    ResVal := ResVal + AVar.GetValue();
  end;
  result := TCalculationWrapper.NumbersWrapper.CreateVariable(ResVal);
end;

end.
