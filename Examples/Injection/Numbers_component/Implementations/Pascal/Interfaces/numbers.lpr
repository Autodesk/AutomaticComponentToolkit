(*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Pascal project file in order to allow easy
development of Numbers library.

Interface version: 1.0.0

*)

{$MODE DELPHI}
library numbers;

uses
{$IFDEF UNIX}
  cthreads,
{$ENDIF UNIX}
  syncobjs,
  numbers_types,
  numbers_exports,
  Classes,
  sysutils;

exports
  numbers_base_classtypeid,
  numbers_variable_getvalue,
  numbers_variable_setvalue,
  numbers_createvariable,
  numbers_getversion,
  numbers_getlasterror,
  numbers_releaseinstance,
  numbers_acquireinstance,
  numbers_getsymbollookupmethod;

{$IFDEF NUMBERS_INCLUDE_RES_FILE}
{$R *.res}
{$ENDIF NUMBERS_INCLUDE_RES_FILE}

begin

end.
