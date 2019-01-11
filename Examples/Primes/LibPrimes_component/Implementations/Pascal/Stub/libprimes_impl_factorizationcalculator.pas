(*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of TLibPrimesFactorizationCalculator

*)

{$MODE DELPHI}
unit libprimes_impl_factorizationcalculator;

interface

uses
	libprimes_types,
	libprimes_interfaces,
	libprimes_exception,
	libprimes_impl_calculator,
	Classes,
	sysutils;

type
	TLibPrimesFactorizationCalculator = class (TLibPrimesCalculator, ILibPrimesFactorizationCalculator)
		private
			FPrimeFactors : Array Of TLibPrimesPrimeFactor;
		protected

		public
			procedure GetPrimeFactors(const APrimeFactorsCount: QWord; PPrimeFactorsNeededCount: PQWord; APrimeFactors: PLibPrimesPrimeFactor);
			procedure Calculate(); override;
			destructor Destroy(); override;
	end;

implementation

destructor TLibPrimesFactorizationCalculator.Destroy();
begin
  SetLength(FPrimeFactors, 0);
end;

procedure TLibPrimesFactorizationCalculator.GetPrimeFactors(const APrimeFactorsCount: QWord; PPrimeFactorsNeededCount: PQWord; APrimeFactors: PLibPrimesPrimeFactor);
var
  i : QWord;
begin
  if (Length(FPrimeFactors) = 0) then
    raise ELibPrimesException.Create(LIBPRIMES_ERROR_NORESULTAVAILABLE);

  if (assigned(PPrimeFactorsNeededCount)) then
     PPrimeFactorsNeededCount^ := Length(FPrimeFactors);

  if (APrimeFactorsCount >= Length(FPrimeFactors)) then
  begin
    for i:=0 to Length(FPrimeFactors) -1 do begin
      APrimeFactors^ := FPrimeFactors[i];
      inc(APrimeFactors);
    end;
  end;
end;

procedure TLibPrimesFactorizationCalculator.Calculate();
var
  AValue: QWord;
  I: QWord;
  APFCount: QWord;
  APrimeFactor: TLibPrimesPrimeFactor;
  AShouldAbort: Cardinal;
begin
  SetLength(FPrimeFactors, 0);

  APFCount := 0;
  AValue := FValue;
  I := 2;
  while I < AValue
  do begin


    if (assigned(FProgressCallback)) then begin
    	AShouldAbort := 0;
			FProgressCallback(1 - 1.0*AValue / FValue, AShouldAbort);
			if (AShouldAbort <> 0) then
				raise ELibPrimesException.Create(LIBPRIMES_ERROR_CALCULATIONABORTED);
    end;


    APrimeFactor.FMultiplicity:=0;
    APrimeFactor.FPrime:=I;
    while (AValue mod i = 0) do begin
      inc(APrimeFactor.FMultiplicity);
      AValue := AValue div I;
    end;
    if (APrimeFactor.FMultiplicity > 0) then begin
      inc(APFCount);
      SetLength(FPrimeFactors, APFCount);
      FPrimeFactors[APFCount-1] := APrimeFactor;
    end;
    inc(I);
  end;
end;

end.
