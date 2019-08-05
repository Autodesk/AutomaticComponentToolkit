(*++

Copyright (C) 2019 Numbers developers

All rights reserved.

Abstract: This is the class declaration of TNumbersBase

*)

{$MODE DELPHI}
unit numbers_impl_base;

interface

uses
  numbers_types,
  numbers_interfaces,
  numbers_exception,
  Classes,
  sysutils;

type
  TNumbersBase = class(TObject, INumbersBase)
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

constructor TNumbersBase.Create();
begin
  inherited Create();
  FMessages := TStringList.Create();
  FReferenceCount := 1;
end;

destructor TNumbersBase.Destroy();
begin
  FreeAndNil(FMessages);
  inherited Destroy();
end;

function TNumbersBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
  result := (FMessages.Count>0);
  if (result) then
    AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TNumbersBase.ClearErrorMessages();
begin
  FMessages.Clear();
end;

procedure TNumbersBase.RegisterErrorMessage(const AErrorMessage: String);
begin
  FMessages.Clear();
  FMessages.Add(AErrorMessage);
end;

procedure TNumbersBase.IncRefCount();
begin
  inc(FReferenceCount);
end;

function TNumbersBase.DecRefCount(): Boolean;
begin
  dec(FReferenceCount);
  if (FReferenceCount = 0) then begin
    result := true;
    self.Destroy();
  end;
   result := false;
end;

end.
