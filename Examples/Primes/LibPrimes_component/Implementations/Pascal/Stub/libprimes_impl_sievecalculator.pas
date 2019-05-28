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
      procedure GetPrimes(const APrimesCount: QWord; PPrimesNeededCount: PQWord; APrimes: PQWord);
  end;

implementation

procedure TLibPrimesSieveCalculator.GetPrimes(const APrimesCount: QWord; PPrimesNeededCount: PQWord; APrimes: PQWord);
begin
  raise ELibPrimesException.Create (LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

end.
