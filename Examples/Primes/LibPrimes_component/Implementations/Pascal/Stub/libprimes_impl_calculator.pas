(*++

Copyright (C) 2018 PrimeDevelopers

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
  libprimes_impl_baseclass,
	Classes,
	sysutils;

type
	TLibPrimesCalculator = class (TLibPrimesBaseClass, ILibPrimesCalculator)
		private

		protected
      FValue : QWord;
      FProgressCallback: PLibPrimes_ProgressCallback;
		public
			function GetValue(): QWord;
			procedure SetValue(const AValue: QWord);
			procedure SetProgressCallback(const AProgressCallback: PLibPrimes_ProgressCallback);
			procedure Calculate(); virtual;
	end;

implementation

function TLibPrimesCalculator.GetValue(): QWord;
begin
	result := FValue;
end;

procedure TLibPrimesCalculator.SetValue(const AValue: QWord);
begin
	FValue := AValue;
end;

procedure TLibPrimesCalculator.SetProgressCallback(const AProgressCallback: PLibPrimes_ProgressCallback);
begin
  FProgressCallback:=AProgressCallback;
end;

procedure TLibPrimesCalculator.Calculate();
begin
  raise ELibPrimesException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

end.
