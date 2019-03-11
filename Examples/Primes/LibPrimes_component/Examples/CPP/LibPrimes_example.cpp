/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0-develop.

Abstract: This is an autogenerated C++ application that demonstrates the
 usage of the C++ bindings of Prime Numbers Library

Interface version: 1.2.0

*/

#include <iostream>
#include "libprimes.hpp"


int main()
{
  try
  {
    unsigned int nMajor, nMinor, nMicro;
    std::string sPreReleaseInfo, sBuildInfo;
    LibPrimes::CLibPrimesWrapper::GetLibraryVersion(nMajor, nMinor, nMicro, sPreReleaseInfo, sBuildInfo);
    std::cout << "LibPrimes.Version = " << nMajor << "." << nMinor << "." << nMicro;
    if (!sPreReleaseInfo.empty())
      std::cout << "-" << sPreReleaseInfo;
    if (!sBuildInfo.empty())
      std::cout << "+" << sBuildInfo;
    std::cout << std::endl;
  }
  catch (std::exception &e)
  {
    std::cout << e.what() << std::endl;
    return 1;
  }
  return 0;
}

