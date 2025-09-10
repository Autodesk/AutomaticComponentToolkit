#include <iostream>
#include "libprimes_dynamic.hpp"


int main()
{
  try
  {
    std::string libpath = (""); // TODO: put the location of the LibPrimes-library file here.
    auto wrapper = LibPrimes::CWrapper::loadLibrary(libpath + "/libprimes."); // TODO: add correct suffix of the library
    LibPrimes_uint32 nMajor, nMinor, nMicro;
    wrapper->GetVersion(nMajor, nMinor, nMicro);
    std::cout << "LibPrimes.Version = " << nMajor << "." << nMinor << "." << nMicro;
    std::cout << std::endl;
  }
  catch (std::exception &e)
  {
    std::cout << e.what() << std::endl;
    return 1;
  }
  return 0;
}

