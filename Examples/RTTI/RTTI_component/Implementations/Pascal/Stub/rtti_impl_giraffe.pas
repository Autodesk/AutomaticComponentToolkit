(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIGiraffe

*)

{$MODE DELPHI}
unit rtti_impl_giraffe;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_mammal,
	Classes,
	sysutils;

type
	TRTTIGiraffe = class(TRTTIMammal, IRTTIGiraffe)
		private

		protected

		public
			function ClassTypeId(): QWord; Override;
	end;

implementation

function TRTTIGiraffe.ClassTypeId(): QWord;
begin
	Result := $9751971BD2C2D958; // First 64 bits of SHA1 of a string: "RTTI::Giraffe"
end;

end.
