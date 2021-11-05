(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIMammal

*)

{$MODE DELPHI}
unit rtti_impl_mammal;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_animal,
	Classes,
	sysutils;

type
	TRTTIMammal = class(TRTTIAnimal, IRTTIMammal)
		private

		protected

		public
			function ClassTypeId(): QWord; Override;
	end;

implementation

function TRTTIMammal.ClassTypeId(): QWord;
begin
	Result := QWord($BC9D5FA7750C1020); // First 64 bits of SHA1 of a string: "RTTI::Mammal"
end;

end.
