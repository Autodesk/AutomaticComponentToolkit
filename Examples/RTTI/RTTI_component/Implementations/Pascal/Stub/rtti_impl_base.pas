(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIBase

*)

{$MODE DELPHI}
unit rtti_impl_base;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	Classes,
	sysutils;

type
	TRTTIBase = class(TObject, IRTTIBase)
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

constructor TRTTIBase.Create();
begin
	inherited Create();
	FMessages := TStringList.Create();
	FReferenceCount := 1;
end;

destructor TRTTIBase.Destroy();
begin
	FreeAndNil(FMessages);
	inherited Destroy();
end;

function TRTTIBase.GetLastErrorMessage(out AErrorMessage: String): Boolean;
begin
	result := (FMessages.Count>0);
	if (result) then
		AErrorMessage := FMessages[FMessages.Count-1];
end;

procedure TRTTIBase.ClearErrorMessages();
begin
	FMessages.Clear();
end;

procedure TRTTIBase.RegisterErrorMessage(const AErrorMessage: String);
begin
	FMessages.Add(AErrorMessage);
end;

procedure TRTTIBase.IncRefCount();
begin
	inc(FReferenceCount);
end;

function TRTTIBase.DecRefCount(): Boolean;
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
