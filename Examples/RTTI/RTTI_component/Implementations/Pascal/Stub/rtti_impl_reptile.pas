(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIReptile

*)

{$MODE DELPHI}
unit rtti_impl_reptile;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_animal,
	Classes,
	sysutils;

type
	TRTTIReptile = class(TRTTIAnimal, IRTTIReptile)
		private

		protected

		public
			function ClassTypeId(): QWord; Override;
	end;

implementation

function TRTTIReptile.ClassTypeId(): QWord;
begin
	Result := $6756AA8EA5802EC3; // First 64 bits of SHA1 of a string: "RTTI::Reptile"
end;

end.
