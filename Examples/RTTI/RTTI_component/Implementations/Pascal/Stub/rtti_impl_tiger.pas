(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTITiger

*)

{$MODE DELPHI}
unit rtti_impl_tiger;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_mammal,
	Classes,
	sysutils;

type
	TRTTITiger = class(TRTTIMammal, IRTTITiger)
		private

		protected

		public
			function ClassTypeId(): QWord; Override;
			procedure Roar();
	end;

implementation

function TRTTITiger.ClassTypeId(): QWord;
begin
	Result := $08D007E7B5F7BAF4; // First 64 bits of SHA1 of a string: "RTTI::Tiger"
end;

procedure TRTTITiger.Roar();
begin
	writeln('ROAAAAARRRRR!!');
end;

end.
