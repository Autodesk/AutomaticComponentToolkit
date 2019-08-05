(*++

Copyright (C) 2019 ACT Developers


Abstract: This is the class declaration of TOptClassBase

*)

{$MODE DELPHI}
unit optclass_impl_base;

interface

uses
  optclass_types,
  optclass_interfaces,
  optclass_exception,
  Classes,
  sysutils;

type
  TOptClassBase = class(TObject, IOptClassBase)
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

constructor TOptClassBase.Create();
begin
  inherited Create();
  FMessages := TStringList.Create();
  FReferenceCount := 1;
end;

destructor TOptClassBase.Destroy();
begin
  FreeAndNil(FMessages);
  inherited Destroy();
end;

function TOptClassBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
  result := (FMessages.Count>0);
  if (result) then
    AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TOptClassBase.ClearErrorMessages();
begin
  FMessages.Clear();
end;

procedure TOptClassBase.RegisterErrorMessage(const AErrorMessage: String);
begin
  FMessages.Clear();
  FMessages.Add(AErrorMessage);
end;

procedure TOptClassBase.IncRefCount();
begin
  inc(FReferenceCount);
end;

function TOptClassBase.DecRefCount(): Boolean;
begin
  dec(FReferenceCount);
  if (FReferenceCount = 0) then begin
    result := true;
    self.Destroy();
  end;
   result := false;
end;

end.
