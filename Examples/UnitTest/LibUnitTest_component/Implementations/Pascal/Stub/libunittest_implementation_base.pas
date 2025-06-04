(*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is the class declaration of TLibUnitTestBase

*)

{$MODE DELPHI}
unit libunittest_implementation_base;

interface

uses
  libunittest_types,
  libunittest_interfaces,
  libunittest_exception,
  Classes,
  sysutils;

type
  TLibUnitTestBase = class(TObject, ILibUnitTestBase)
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
      function ClassTypeId(): QWord; Virtual; Abstract;
  end;

implementation

constructor TLibUnitTestBase.Create();
begin
  inherited Create();
  FMessages := TStringList.Create();
  FReferenceCount := 1;
end;

destructor TLibUnitTestBase.Destroy();
begin
  FreeAndNil(FMessages);
  inherited Destroy();
end;

function TLibUnitTestBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
  result := (FMessages.Count>0);
  if (result) then
    AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TLibUnitTestBase.ClearErrorMessages();
begin
  FMessages.Clear();
end;

procedure TLibUnitTestBase.RegisterErrorMessage(const AErrorMessage: String);
begin
  FMessages.Add(AErrorMessage);
end;

procedure TLibUnitTestBase.IncRefCount();
begin
  inc(FReferenceCount);
end;

function TLibUnitTestBase.DecRefCount(): Boolean;
begin
  dec(FReferenceCount);
  if (FReferenceCount = 0) then begin
    self.Destroy();
    result := true;
  end
  else
    result := false;
end;

end.
