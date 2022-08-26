(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIZoo

*)

{$MODE DELPHI}
unit rtti_impl_zoo;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_base,
	rtti_impl_animal,
	rtti_impl_animaliterator,
	Classes,
	sysutils;

type
	TRTTIZoo = class(TRTTIBase, IRTTIZoo)
		private

		protected

		public
			Animals: TList;
			constructor Create();
			function ClassTypeId(): QWord; Override;
			function Iterator(): TObject;
	end;

implementation

function TRTTIZoo.ClassTypeId(): QWord;
begin
	Result := $2262ABE80A5E7878; // First 64 bits of SHA1 of a string: "RTTI::Zoo"
end;

constructor TRTTIZoo.Create();
begin
        Animals := TList.Create();
end;

function TRTTIZoo.Iterator(): TObject;
begin
	Result := TRTTIAnimalIterator.Create(Animals);
end;

end.
