(*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of TLibPrimesBase

*)

{$MODE DELPHI}
unit libprimes_impl_base;

interface

uses
  libprimes_types,
  libprimes_interfaces,
  libprimes_exception,
  Classes,
  sysutils;

type
  TLibPrimesBase = class(TObject, ILibPrimesBase)
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

constructor TLibPrimesBase.Create();
begin
  inherited Create();
  FMessages := TStringList.Create();
  FReferenceCount := 1;
end;

destructor TLibPrimesBase.Destroy();
begin
  FreeAndNil(FMessages);
  inherited Destroy();
end;

function TLibPrimesBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
  result := (FMessages.Count>0);
  if (result) then
    AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TLibPrimesBase.ClearErrorMessages();
begin
  FMessages.Clear();
end;

procedure TLibPrimesBase.RegisterErrorMessage(const AErrorMessage: String);
begin
  FMessages.Add(AErrorMessage);
end;

procedure TLibPrimesBase.IncRefCount();
begin
  inc(FReferenceCount);
end;

function TLibPrimesBase.DecRefCount(): Boolean;
begin
  dec(FReferenceCount);
  if (FReferenceCount = 0) then begin
    result := true;
    self.Destroy();
  end;
   result := false;
end;

end.
