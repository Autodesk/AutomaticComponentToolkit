(*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is the class declaration of TRTTIAnimal

*)

{$MODE DELPHI}
unit rtti_impl_animal;

interface

uses
	rtti_types,
	rtti_interfaces,
	rtti_exception,
	rtti_impl_base,
	Classes,
	sysutils;

type
	TRTTIAnimal = class(TRTTIBase, IRTTIAnimal)
		private

		protected
			m_sName: String;
		public
			function ClassTypeId(): QWord; Override;
			function Name(): String;
		    constructor Create(Name: String); 
			destructor Done;
	end;

implementation

function TRTTIAnimal.ClassTypeId(): QWord;
begin
	Result := $8B40467DA6D327AF; // First 64 bits of SHA1 of a string: "RTTI::Animal"
end;

function TRTTIAnimal.Name(): String;
begin
    Result := m_sName;
end;

constructor TRTTIAnimal.Create(Name: String);
begin
   self.m_sName := Name;
end;

destructor TRTTIAnimal.Done;
begin
   writeln('"Delete ', Name);
end;

end.
