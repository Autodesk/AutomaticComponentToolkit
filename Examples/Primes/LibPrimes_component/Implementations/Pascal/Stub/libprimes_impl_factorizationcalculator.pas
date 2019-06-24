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

    protected

    public
      procedure GetPrimeFactors(const APrimeFactorsCount: QWord; PPrimeFactorsNeededCount: PQWord; APrimeFactors: PLibPrimesPrimeFactor);
  end;

implementation

procedure TLibPrimesFactorizationCalculator.GetPrimeFactors(const APrimeFactorsCount: QWord; PPrimeFactorsNeededCount: PQWord; APrimeFactors: PLibPrimesPrimeFactor);
begin
  raise ELibPrimesException.Create (LIBPRIMES_ERROR_NOTIMPLEMENTED);
end;

end.
