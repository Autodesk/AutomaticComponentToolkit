(*++

Copyright (C) 2019 Calculation developers

All rights reserved.

Abstract: This is the class declaration of TCalculationBase

*)

{$MODE DELPHI}
unit calculation_impl_base;

interface

uses
  calculation_types,
  calculation_interfaces,
  calculation_exception,
  Unit_Numbers,
  Classes,
  sysutils;

type
  TCalculationBase = class(TObject, ICalculationBase)
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

constructor TCalculationBase.Create();
begin
  inherited Create();
  FMessages := TStringList.Create();
  FReferenceCount := 1;
end;

destructor TCalculationBase.Destroy();
begin
  FreeAndNil(FMessages);
  inherited Destroy();
end;

function TCalculationBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
  result := (FMessages.Count>0);
  if (result) then
    AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TCalculationBase.ClearErrorMessages();
begin
  FMessages.Clear();
end;

procedure TCalculationBase.RegisterErrorMessage(const AErrorMessage: String);
begin
  FMessages.Add(AErrorMessage);
end;

procedure TCalculationBase.IncRefCount();
begin
  inc(FReferenceCount);
end;

function TCalculationBase.DecRefCount(): Boolean;
begin
  dec(FReferenceCount);
  if (FReferenceCount = 0) then begin
    result := true;
    self.Destroy();
  end;
   result := false;
end;

end.
