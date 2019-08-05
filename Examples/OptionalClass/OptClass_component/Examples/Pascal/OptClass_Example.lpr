(*++

Copyright (C) 2019 ACT Developers


This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0.

Abstract: This is an autogenerated Pascal application that demonstrates the
 usage of the Pascal bindings of Optional Class Library

Interface version: 1.0.0

*)

program OptClassPascalTest;

uses
  {$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
  {$ENDIF}{$ENDIF}
  Classes, SysUtils, CustApp,
  Unit_OptClass
  { you can add units after this };

type

TOptClass_Example = class(TCustomApplication)
protected
  procedure DoRun; override;
  procedure TestOptClass();
public
  constructor Create(TheOwner: TComponent); override;
  destructor Destroy; override;
end;


procedure TOptClass_Example.TestOptClass();
var
  AOptClassWrapper: TOptClassWrapper;
  AMajor, AMinor, AMicro: Cardinal;
  AVersionString: string;
  ALibPath: string;
  ABaseA, ABaseB: TOptClassBase;
begin
  writeln('loading DLL');
  ALibPath := 'D:/PUBLIC/AutomaticComponentToolkit/Examples/OptionalClass/OptClass_component/Implementations/Cpp/build/Debug/'; // TODO add the location of the shared library binary here
  AOptClassWrapper := TOptClassWrapper.Create(ALibPath + '/' + 'optclass.dll'); // TODO add the extension of the shared library file here
  try
    writeln('loading DLL Done');
    AOptClassWrapper.GetVersion(AMajor, AMinor, AMicro);
    AVersionString := Format('OptClass.version = %d.%d.%d', [AMajor, AMinor, AMicro]);
    writeln(AVersionString);

    AOptClassWrapper.CreateInstanceWithName('A');
    ABaseA := AOptClassWrapper.FindInstanceA('A');
    if not assigned(ABaseA) then begin
      WriteLn('Error: Expected to find Instance "A".');
      exit;
    end;
    AOptClassWrapper.FindInstanceB('DoesNotExist', ABaseB);
    if assigned(ABaseB) then begin
      WriteLn('Error: Did not expect to find Instance "DoesNotExist".');
      exit;
    end;
    if not AOptClassWrapper.UseInstanceMaybe(ABaseA) then begin
      WriteLn('Error: Expected to use Instance "A".');
      exit;
    end;
     if AOptClassWrapper.UseInstanceMaybe(ABaseB) then begin
      WriteLn('Error: Expected to not use Instance "DoesNotExist".');
      exit;
    end;
    WriteLn('Passed');
  finally
    FreeAndNil(AOptClassWrapper);
  end;
end;

procedure TOptClass_Example.DoRun;
begin
  try
    TestOptClass();
  except
    On E: Exception do
      writeln('Fatal error: ', E.Message);
  end;
  Terminate
end;

constructor TOptClass_Example.Create(TheOwner: TComponent);
begin
  inherited Create(TheOwner);
  StopOnException:=True;
end;

destructor TOptClass_Example.Destroy;
begin
  inherited Destroy;
end;


var
  Application: TOptClass_Example;
begin
  Application:=TOptClass_Example.Create(nil);
  Application.Run;
  Application.Free;
end.