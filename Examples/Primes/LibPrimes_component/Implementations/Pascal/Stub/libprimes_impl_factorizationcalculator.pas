(*++

Copyright (C) 2019 PrimeDevelopers

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
  TLibPrimesFactorizationCalculator = class(TLibPrimesCalculator, ILibPrimesFactorizationCalculator)
    private
      FPrimeFactors : Array Of TLibPrimesPrimeFactor;
    protected

    public
      constructor Create();
      destructor Destroy(); override;
      function ClassTypeId(): QWord; Override;
      procedure Calculate(); override;
      procedure GetPrimeFactors(const APrimeFactorsCount: QWord; PPrimeFactorsNeededCount: PQWord; APrimeFactors: PLibPrimesPrimeFactor);
  end;

implementation

constructor TLibPrimesFactorizationCalculator.Create();
begin
  inherited Create();
  SetLength(FPrimeFactors, 0);
end;

destructor TLibPrimesFactorizationCalculator.Destroy();
begin
  SetLength(FPrimeFactors, 0);
  inherited Destroy();
end;

function TLibPrimesFactorizationCalculator.ClassTypeId(): QWord;
begin
  Result := QWord($6C7A0FD2ECC65118); // First 64 bits of SHA1 of a string: "LibPrimes::FactorizationCalculator"
end;

procedure TLibPrimesFactorizationCalculator.Calculate();
var
  AValue: QWord;
  I: QWord;
  APFCount: QWord;
  APrimeFactor: TLibPrimesPrimeFactor;
begin
  SetLength(FPrimeFactors, 0);

  APFCount := 0;
  AValue := FValue;
  I := 2;
  while (I * I <= AValue) and (AValue > 1)
  do begin
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
  
  // If AValue is still greater than 1, it's a prime factor itself
  if (AValue > 1) then begin
    APrimeFactor.FMultiplicity := 1;
    APrimeFactor.FPrime := AValue;
    inc(APFCount);
    SetLength(FPrimeFactors, APFCount);
    FPrimeFactors[APFCount-1] := APrimeFactor;
  end;
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

end.
