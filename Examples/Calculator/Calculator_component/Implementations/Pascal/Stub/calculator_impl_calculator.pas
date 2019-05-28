(*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is the class declaration of TCalculatorCalculator

*)

{$MODE DELPHI}
unit calculator_impl_calculator;

interface

uses
  calculator_types,
  calculator_interfaces,
  calculator_exception,
  calculator_impl_base,
  calculator_impl_variable,
  Classes,
  sysutils,
  contnrs;

type
  TCalculatorCalculator = class(TCalculatorBase, ICalculatorCalculator)
    private
      FVariableList : TObjectList;
    protected

    public
      constructor Create();
      destructor Destroy(); override;
      procedure EnlistVariable(AVariable: TObject);
      function GetEnlistedVariable(const AIndex: Cardinal): TObject;
      procedure ClearVariables();
      function Multiply(): TObject;
      function Add(): TObject;
  end;

implementation

constructor TCalculatorCalculator.Create();
begin
  inherited Create();
  FVariableList := TObjectList.Create(False);
end;

destructor TCalculatorCalculator.Destroy();
begin
  ClearVariables();
  FreeAndNil(FVariableList);
  inherited Destroy();
end;

procedure TCalculatorCalculator.EnlistVariable(AVariable: TObject);
var
  AVar: TCalculatorVariable;
begin
  AVar := (AVariable as TCalculatorVariable);
  AVar.IncRefCount();
  FVariableList.Add(AVar);
end;

function TCalculatorCalculator.GetEnlistedVariable(const AIndex: Cardinal): TObject;
begin
  if AIndex >= FVariableList.Count then begin
    raise ECalculatorException.Create(CALCULATOR_ERROR_INVALIDPARAM);
  end;
  result := FVariableList[AIndex];
  (result as TCalculatorVariable).IncRefCount();
end;

procedure TCalculatorCalculator.ClearVariables();
var
  AVar: TCalculatorVariable;
  I: integer;
begin
  For I := 0 to FVariableList.Count - 1 do begin
    AVar := (FVariableList[I] as TCalculatorVariable);
    AVar.DecRefCount();
  end;
  FVariableList.Clear;
end;

function TCalculatorCalculator.Multiply(): TObject;
var
  AVar, AResVar: TCalculatorVariable;
  I: integer;
  ResVal : double;
begin
  ResVal := 1.0;
  For I := 0 to FVariableList.Count - 1 do begin
    AVar := (FVariableList[I] as TCalculatorVariable);
    ResVal := ResVal * AVar.GetValue();
  end;
  AResVar := TCalculatorVariable.Create();
  AResVar.SetValue(ResVal);
  result := AResVar;
end;

function TCalculatorCalculator.Add(): TObject;
var
  AVar, AResVar: TCalculatorVariable;
  I: integer;
  ResVal : double;
begin
  ResVal := 0.0;
  For I := 0 to FVariableList.Count - 1 do begin
    AVar := (FVariableList[I] as TCalculatorVariable);
    ResVal := ResVal + AVar.GetValue();
  end;
  AResVar := TCalculatorVariable.Create();
  AResVar.SetValue(ResVal);
  result := AResVar;
end;

end.
