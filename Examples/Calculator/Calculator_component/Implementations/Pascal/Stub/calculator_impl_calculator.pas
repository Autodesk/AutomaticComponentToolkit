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
  sysutils;

type
  TCalculatorCalculator = class(TCalculatorBase, ICalculatorCalculator)
    private
      FVariables: TList;

    protected

    public
      constructor Create();
      destructor Destroy(); override;
      function ClassTypeId(): QWord; Override;
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
  FVariables := TList.Create();
end;

destructor TCalculatorCalculator.Destroy();
begin
  if Assigned(FVariables) then
  begin
    ClearVariables(); // This will properly decrement reference counts
    FVariables.Free();
  end;
  inherited Destroy();
end;

function TCalculatorCalculator.ClassTypeId(): QWord;
begin
  Result := QWord($5FA5D809D8A728B0); // First 64 bits of SHA1 of a string: "Calculator::Calculator"
end;

procedure TCalculatorCalculator.EnlistVariable(AVariable: TObject);
begin
  if not Assigned(AVariable) then
    raise ECalculatorException.Create(CALCULATOR_ERROR_INVALIDPARAM);
  if not (AVariable is TCalculatorVariable) then
    raise ECalculatorException.Create(CALCULATOR_ERROR_INVALIDCAST);
  
  // Increment reference count since we're holding a reference
  (AVariable as ICalculatorBase).IncRefCount();
  FVariables.Add(AVariable);
end;

function TCalculatorCalculator.GetEnlistedVariable(const AIndex: Cardinal): TObject;
begin
  if AIndex >= Cardinal(FVariables.Count) then
    raise ECalculatorException.Create(CALCULATOR_ERROR_INVALIDPARAM);
  Result := TObject(FVariables[AIndex]);
  
  // Increment reference count since C++ will manage this returned object
  if Assigned(Result) then
    (Result as ICalculatorBase).IncRefCount();
end;

procedure TCalculatorCalculator.ClearVariables();
var
  I: Integer;
  Variable: TObject;
begin
  // Decrement reference count for all variables before clearing
  for I := 0 to FVariables.Count - 1 do
  begin
    Variable := TObject(FVariables[I]);
    if Assigned(Variable) and (Variable is TCalculatorBase) then
    begin
      try
        (Variable as ICalculatorBase).DecRefCount();
      except
        // Ignore exceptions during cleanup to prevent cascade failures
      end;
    end;
  end;
  FVariables.Clear();
end;

function TCalculatorCalculator.Multiply(): TObject;
var
  I: Integer;
  ResultValue: Double;
  Variable: TCalculatorVariable;
begin
  if FVariables.Count = 0 then
    raise ECalculatorException.Create(CALCULATOR_ERROR_INVALIDPARAM);
  
  ResultValue := 1.0;
  for I := 0 to FVariables.Count - 1 do
  begin
    Variable := TCalculatorVariable(FVariables[I]);
    ResultValue := ResultValue * Variable.GetValue();
  end;
  
  Result := TCalculatorVariable.Create(ResultValue);
end;

function TCalculatorCalculator.Add(): TObject;
var
  I: Integer;
  ResultValue: Double;
  Variable: TCalculatorVariable;
begin
  if FVariables.Count = 0 then
    raise ECalculatorException.Create(CALCULATOR_ERROR_INVALIDPARAM);
  
  ResultValue := 0.0;
  for I := 0 to FVariables.Count - 1 do
  begin
    Variable := TCalculatorVariable(FVariables[I]);
    ResultValue := ResultValue + Variable.GetValue();
  end;
  
  Result := TCalculatorVariable.Create(ResultValue);
end;

end.
