#include <iostream>
#include "libprimes_implicit.hpp"


int main()
{
  try
  {
    auto wrapper = LibPrimes::CWrapper::loadLibrary();
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

