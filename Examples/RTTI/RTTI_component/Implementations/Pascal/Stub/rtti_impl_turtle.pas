(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTITurtle

*)

{$MODE DELPHI}
unit rtti_impl_turtle;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_reptile,
	Classes,
	sysutils;

type
	TRTTITurtle = class(TRTTIReptile, IRTTITurtle)
		private

		protected

		public
			function ClassTypeId(): QWord; Override;
	end;

implementation

function TRTTITurtle.ClassTypeId(): QWord;
begin
	Result := QWord($8E551B208A2E8321); // First 64 bits of SHA1 of a string: "RTTI::Turtle"
end;

end.
