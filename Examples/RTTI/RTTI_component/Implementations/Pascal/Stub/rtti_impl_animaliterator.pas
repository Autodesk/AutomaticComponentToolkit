(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIAnimalIterator

*)

{$MODE DELPHI}
unit rtti_impl_animaliterator;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_base,
	rtti_impl_animal,
	Classes,
	sysutils;

type
	TRTTIAnimalIterator = class(TRTTIBase, IRTTIAnimalIterator)
		private
		protected
			Animals: TList;
			Index: integer;
		public
			function ClassTypeId(): QWord; Override;

			constructor Create(Animals: TList);
			function GetNextAnimal(): TObject;
			function GetNextOptinalAnimal(out AAnimal: TObject): Boolean;
			function GetNextMandatoryAnimal(out AAnimal: TObject): Boolean;

		protected
	end;

implementation

function TRTTIAnimalIterator.ClassTypeId(): QWord;
begin
	Result := QWord($F1917FE6BBE77831); // First 64 bits of SHA1 of a string: "RTTI::AnimalIterator"
end;

constructor TRTTIAnimalIterator.Create(Animals: TList);
begin
     self.Animals := Animals;
     Index := 0;
end;

function TRTTIAnimalIterator.GetNextAnimal(): TObject;
begin
	if Index < Animals.Count then
            Result := Animals[Index]
        else
            Result := nil;
        Index := Index + 1;
end;

function TRTTIAnimalIterator.GetNextOptinalAnimal(out AAnimal: TObject): Boolean;
begin
    AAnimal := GetNextAnimal();
	Result := true;
end;

function TRTTIAnimalIterator.GetNextMandatoryAnimal(out AAnimal: TObject): Boolean;
begin
    AAnimal := GetNextAnimal();
	Result := true;
end;

end.
