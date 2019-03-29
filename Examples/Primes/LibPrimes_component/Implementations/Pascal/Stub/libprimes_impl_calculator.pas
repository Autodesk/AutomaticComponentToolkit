(*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of TLibPrimesCalculator

*)

{$MODE DELPHI}
unit libprimes_impl_calculator;

interface

uses
	libprimes_types,
	libprimes_interfaces,
	libprimes_exception,
	libprimes_impl_base,
	Classes,
	sysutils;

type
	TLibPrimesCalculator = class(TLibPrimesBase, ILibPrimesCalculator)
		private

		protected

		public
			function GetValue(): QWord;
			procedure SetValue(const AValue: QWord);
			procedure Calculate();
	end;

implementation

function TLibPrimesCalculator.GetValue(): QWord;
begin
	raise ELibPrimesException.Create (LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

procedure TLibPrimesCalculator.SetValue(const AValue: QWord);
begin
	raise ELibPrimesException.Create (LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

procedure TLibPrimesCalculator.Calculate();
begin
	raise ELibPrimesException.Create (LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

end.
