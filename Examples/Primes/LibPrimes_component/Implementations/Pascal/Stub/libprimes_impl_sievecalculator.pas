(*++

Copyright (C) 2019 PrimeDevelopers

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
  TLibPrimesSieveCalculator = class(TLibPrimesCalculator, ILibPrimesSieveCalculator)
    private

    protected

    public
      function ClassTypeId(): QWord; Override;
      procedure GetPrimes(const APrimesCount: QWord; PPrimesNeededCount: PQWord; APrimes: PQWord);
  end;

implementation

function TLibPrimesSieveCalculator.ClassTypeId(): QWord;
begin
  Result := QWord($58F1CCC31D38375B); // First 64 bits of SHA1 of a string: "LibPrimes::SieveCalculator"
end;

procedure TLibPrimesSieveCalculator.GetPrimes(const APrimesCount: QWord; PPrimesNeededCount: PQWord; APrimes: PQWord);
begin
  raise ELibPrimesException.Create(LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

end.
