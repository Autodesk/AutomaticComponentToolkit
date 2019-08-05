(*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is the class declaration of TCalculatorBase

*)

{$MODE DELPHI}
unit calculator_impl_base;

interface

uses
  calculator_types,
  calculator_interfaces,
  calculator_exception,
  Classes,
  sysutils;

type
  TCalculatorBase = class(TObject, ICalculatorBase)
    private
      FMessages: TStringList;
      FReferenceCount: integer;

    protected

    public
      constructor Create();
      destructor Destroy(); override;
      function GetLastErrorMessage(out AErrorMessage: String): Boolean;
      procedure ClearErrorMessages();
      procedure RegisterErrorMessage(const AErrorMessage: String);
      procedure IncRefCount();
      function DecRefCount(): Boolean;
  end;

implementation

constructor TCalculatorBase.Create();
begin
  inherited Create();
  FMessages := TStringList.Create();
  FReferenceCount := 1;
end;

destructor TCalculatorBase.Destroy();
begin
  FreeAndNil(FMessages);
  inherited Destroy();
end;

function TCalculatorBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
  result := (FMessages.Count>0);
  if (result) then
    AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TCalculatorBase.ClearErrorMessages();
begin
  FMessages.Clear();
end;

procedure TCalculatorBase.RegisterErrorMessage(const AErrorMessage: String);
begin
  FMessages.Clear();
  FMessages.Add(AErrorMessage);
end;

procedure TCalculatorBase.IncRefCount();
begin
  inc(FReferenceCount);
end;

function TCalculatorBase.DecRefCount(): Boolean;
begin
  dec(FReferenceCount);
  if (FReferenceCount = 0) then begin
    result := true;
    self.Destroy();
  end;
   result := false;
end;

end.
