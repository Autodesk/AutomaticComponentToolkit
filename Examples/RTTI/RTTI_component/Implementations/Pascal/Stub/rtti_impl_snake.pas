(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTISnake

*)

{$MODE DELPHI}
unit rtti_impl_snake;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_reptile,
	Classes,
	sysutils;

type
	TRTTISnake = class(TRTTIReptile, IRTTISnake)
		private

		protected

		public
			function ClassTypeId(): QWord; Override;
	end;

implementation

function TRTTISnake.ClassTypeId(): QWord;
begin
	Result := $5F6826EF909803B2; // First 64 bits of SHA1 of a string: "RTTI::Snake"
end;

end.
