(*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is the class declaration of TLibPrimesSieveCalculator

*)

{$MODE DELPHI}
unit libprimes_impl_sievecalculator;

interface

uses
	libprimes_types,
	libprimes_interfaces,
	libprimes_exception,
	libprimes_impl_calculator,
	Classes,
	sysutils;

type
	TLibPrimesSieveCalculator = class (TLibPrimesCalculator, ILibPrimesSieveCalculator)
		private
			FPrimes: array of QWord;
		protected

		public
			procedure GetPrimes(const APrimesCount: QWord; PPrimesNeededCount: PQWord; APrimes: PQWord);
			procedure Calculate(); override;
	end;

implementation

procedure TLibPrimesSieveCalculator.Calculate();
var
  AStrikenOut : array of Boolean;
  I, J : QWord;
  ASqrtValue : QWord;
  ANumPrimes: QWord;
begin
  SetLength(FPrimes, 0);
  ANumPrimes := 0;

  SetLength(AStrikenOut, FValue + 1);
  for I := 0 to FValue do begin
    AStrikenOut[I] := I < 2;
  end;

  ASqrtValue := round(sqrt(FValue));

  for I := 2 to ASqrtValue do begin
    if not AStrikenOut[I] then begin
      inc(ANumPrimes);
      SetLength(FPrimes, ANumPrimes);
      FPrimes[ANumPrimes - 1] := I;
      J := I*I;
      while (J <= FValue) do begin
        AStrikenOut[j] := true;
        inc(J, I);
      end;

    end;

  end;

  for I:= ASqrtValue to FValue do begin
    if not AStrikenOut[i] then begin
      inc(ANumPrimes);
      SetLength(FPrimes, ANumPrimes);
      FPrimes[ANumPrimes - 1] := I;
    end;
  end;

end;

procedure TLibPrimesSieveCalculator.GetPrimes(const APrimesCount: QWord; PPrimesNeededCount: PQWord; APrimes: PQWord);
var
  i : QWord;
begin
  if (Length(FPrimes) = 0) then
    raise ELibPrimesException.Create(LIBPRIMES_ERROR_NORESULTAVAILABLE);

  if (assigned(PPrimesNeededCount)) then
     PPrimesNeededCount^ := Length(FPrimes);

  if (APrimesCount >= Length(FPrimes)) then
  begin
    for i:=0 to Length(FPrimes) -1 do begin
      APrimes^ := FPrimes[i];
      inc(APrimes);
    end;
  end;
end;

end.
