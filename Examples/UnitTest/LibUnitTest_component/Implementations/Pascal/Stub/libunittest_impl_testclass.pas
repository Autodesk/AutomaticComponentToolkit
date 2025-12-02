(*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is the class declaration of TLibUnitTestTestClass

*)

{$MODE DELPHI}
unit libunittest_impl_testclass;

interface

uses
  libunittest_types,
  libunittest_interfaces,
  libunittest_exception,
  libunittest_impl_base,
  Classes,
  sysutils;

type
  TLibUnitTestTestClass = class(TLibUnitTestBase, ILibUnitTestTestClass)
    private
      FTestValue: Double;

    protected

    public
      constructor Create();
      function ClassTypeId(): QWord; Override;
      function Value(): Double;
      procedure SetValue(const AValue: Double);
      procedure SetValueInt(const AValue: Int64);
      procedure SetValueString(const AValue: String);
      procedure UnitTest1(const AValue1: Byte; const AValue2: Word; const AValue3: Cardinal; const AValue4: QWord; out AOutValue1: Byte; out AOutValue2: Word; out AOutValue3: Cardinal; out AOutValue4: QWord);
      procedure UnitTest2(const AValue1: ShortInt; const AValue2: SmallInt; const AValue3: Integer; const AValue4: Int64; out AOutValue1: ShortInt; out AOutValue2: SmallInt; out AOutValue3: Integer; out AOutValue4: Int64);
      procedure UnitTest3(const AValue1: Boolean; const AValue2: Single; const AValue3: Double; const AValue4: TLibUnitTestTestEnum; out AOutValue1: Boolean; out AOutValue2: Single; out AOutValue3: Double; out AOutValue4: TLibUnitTestTestEnum);
      function UnitTest4(const AValue: String; out AOutValue: String): String;
  end;

implementation

constructor TLibUnitTestTestClass.Create();
begin
  inherited Create();
  FTestValue := 0.0;
end;

function TLibUnitTestTestClass.ClassTypeId(): QWord;
begin
  Result := QWord($7F4DF478E28956D1); // First 64 bits of SHA1 of a string: "LibUnitTest::TestClass"
end;

function TLibUnitTestTestClass.Value(): Double;
begin
  Result := FTestValue;
end;

procedure TLibUnitTestTestClass.SetValue(const AValue: Double);
begin
  FTestValue := AValue;
end;

procedure TLibUnitTestTestClass.SetValueInt(const AValue: Int64);
begin
  FTestValue := Double(AValue);
end;

procedure TLibUnitTestTestClass.SetValueString(const AValue: String);
begin
  try
    FTestValue := StrToFloat(AValue);
  except
    FTestValue := 0.0;
  end;
end;

procedure TLibUnitTestTestClass.UnitTest1(const AValue1: Byte; const AValue2: Word; const AValue3: Cardinal; const AValue4: QWord; out AOutValue1: Byte; out AOutValue2: Word; out AOutValue3: Cardinal; out AOutValue4: QWord);
begin
  // Echo back the input values for unit testing verification
  AOutValue1 := AValue1;
  AOutValue2 := AValue2;
  AOutValue3 := AValue3;
  AOutValue4 := AValue4;
end;

procedure TLibUnitTestTestClass.UnitTest2(const AValue1: ShortInt; const AValue2: SmallInt; const AValue3: Integer; const AValue4: Int64; out AOutValue1: ShortInt; out AOutValue2: SmallInt; out AOutValue3: Integer; out AOutValue4: Int64);
begin
  // Echo back the input values for unit testing verification
  AOutValue1 := AValue1;
  AOutValue2 := AValue2;
  AOutValue3 := AValue3;
  AOutValue4 := AValue4;
end;

procedure TLibUnitTestTestClass.UnitTest3(const AValue1: Boolean; const AValue2: Single; const AValue3: Double; const AValue4: TLibUnitTestTestEnum; out AOutValue1: Boolean; out AOutValue2: Single; out AOutValue3: Double; out AOutValue4: TLibUnitTestTestEnum);
begin
  // Echo back the input values for unit testing verification
  AOutValue1 := AValue1;
  AOutValue2 := AValue2;
  AOutValue3 := AValue3;
  AOutValue4 := AValue4;
end;

function TLibUnitTestTestClass.UnitTest4(const AValue: String; out AOutValue: String): String;
begin
  // Echo back the input string for unit testing verification
  AOutValue := AValue;
  Result := AValue;
end;

end.
