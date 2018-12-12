# ![ACT logo](../../Documentation/images/ACT_logo_50px.png) Automatic Component Toolkit

## Tutorial: Generating a C++ shared library and consumers in different Languages


## Table of Contents

- [1. Introduction](#1-introduction)
- [2. Requirements](#2-requirements)
- [3. The Library's Implementation](#3-the-librarys-implementation)
- [4. The Consumers](#4-the-consumers)
- [5. Extending the interface: with a callback](#5-extending-the-interface-with-a-callback)
- [6. Enable journaling](#6-enable-journaling)
- [7. Conclusion](#7-conclusion)

# 1. Introduction
This tutorial demonstrates how to set up a shared library using ACT. You can follow this tutorial step-by-step (more or less). You can find the final project in the folder [LibPrimes_component](LibPrimes_component).
It shows the basic ACT worklfow. All parameters, elements, options and switches are of the interface definition language are documented in the [Documentation](../../Documentation/IDL.md).

Consider you want to implement and use a library that provides datatypes and algorithms related to prime numbers. Such a library will be implemented in this tutorial.
In particular, the library will be able to perform the prime-factorization of positive integers and generate prime numbers.

#### Prime factor decomposition
Given a positive integer n, find distinct prime numbers p_1, ..., p_k, and positive integers e_1, ..., e_k
such that
```
n = p_1 ^ e_1 * p_2 ^ e_2 * ... * p_k ^ e_k
```
The values e_j are called multiplicity.

__NOTE__
_There are much more efficient algorithms and more suitable software packages to perfom these tasks._

# 2. Requirements
 - CMake
 - A C++ compiler / dev environment. This tutorial was tested with Visual Studio 2017, but should also work with GCC and make
 - ACT: This tutorial is tested to work with release 1.0.0 of ACT. You can get it from [the release page](https://github.com/Autodesk/AutomaticComponentToolkit/releases).
Decide on a location for your tutorial's component to live in, and download the binary for your platform into this folder.


# 3. The Library's Implementation
## 3.1. The definition of the component
An ACT component's interface is fully described by its IDL file.
This section sets up a IDL-file for LibPrimes.

First, copy the snippet, a bare-bone IDL-file, and save it into libPrimes.xml in your component's folder.
```xml
<?xml version="1.0" encoding="UTF-8"?>
<component xmlns="http://schemas.autodesk.com/netfabb/automaticcomponenttoolkit/2018" 
	libraryname="Prime numbers Library" namespace="LibPrimes" copyright="PrimeDevelopers" year="2018" basename="libprimes"
	version="1.0.0">
	<license>
		<line value="All rights reserved." />
	</license>
	
	<bindings>
		<binding language="Cpp" indentation="tabs" />
		<binding language="CppDynamic" indentation="tabs" />
		<binding language="Pascal" indentation="4spaces" />
		<binding language="Python" indentation="tabs" />
	</bindings>
	<implementations>
		<implementation language="Cpp" indentation="tabs" claasidentifier="" stubidentifier=""/>
		<implementation language="Pascal" indentation="tabs" claasidentifier="" stubidentifier=""/>
	</implementations>
	
	<errors>
		<error name="NOTIMPLEMENTED" code="1" description="functionality not implemented" />
		<error name="INVALIDPARAM" code="2" description="an invalid parameter was passed" />
		<error name="INVALIDCAST" code="3" description="a type cast failed" />
		<error name="BUFFERTOOSMALL" code="4" description="a provided buffer is too small" />
		<error name="GENERICEXCEPTION" code="5" description="a generic exception occurred" />
		<error name="COULDNOTLOADLIBRARY" code="6" description="the library could not be loaded" />
		<error name="COULDNOTFINDLIBRARYEXPORT" code="7" description="a required exported symbol could not be found in the library" />
	</errors>
	
	<global releasemethod="ReleaseInstance" versionmethod="GetLibraryVersion">
		<method name="ReleaseInstance" description="Releases the memory of an Instance">
			<param name="Instance" type="handle" class="BaseClass" pass="in" description="Instance Handle" />
		</method>

		<method name="GetLibraryVersion" description = "retrieves the current version of the library.">
			<param name="Major" type="uint32" pass="out" description="returns the major version of the library" />
			<param name="Minor" type="uint32" pass="out" description="returns the minor version of the library" />
			<param name="Micro" type="uint32" pass="out" description="returns the micro version of the library" />
		</method>
	</global>
</component>
```

It's elements define the following:
- The attributes of \<component> define essential properties and naming conventions of the component, and contains all other info about the component.
- \<error> define the error codes to be used in the library that will be exposed for it's consumers.
The errors listed in this snippet are required.
-\<global> defines the global functions that can be used as entry points into the component.
It must contain a versionmethod and a releasemethod with the sigantures in this snippet. They will be explained in []().
The syntax for methods will be explained when we add new classes and functions to the IDL-file now.


### 3.1.1 A struct for prime factors
In the \<component> element, add a struct "PrimeFactor" that encodes a prime number with a multiplicity.
```xml
<struct name="PrimeFactor">
	<member name="Prime" type="uint64" />
	<member name="Multiplicity" type="uint32" />
</struct>
```
### 3.1.1 Calculator
In the \<component> element, add a class "Calculator" which implements the base class for
the different calculators we will expose in our API.
```xml
<class name="Calculator">
	<method name="GetValue" description="Returns the current value of this Calculator">
		<param name="Value" type="uint64" pass="return" description="The current value of this Calculator" />
	</method>
	<method name="SetValue" description="Sets the value to be factorized">
		<param name="Value" type="uint64" pass="in" description="The value to be factorized" />
	</method>
	<method name="Calculate" description="Performs the specific calculation of this Calculator">
	</method>
</class>
```
The "GetValue" method returns the current value on which the calculator operates. It has one parameter of type "uint64", that will be returned.
The "SetValue" method sets the value on which the calculator operates. It has one input parameter of type "uint64".
The method "Calculate" performs the specific calculation of this calculator.


### 3.1.2 The FactorizationCalculator
Add another class "FacrtoziationCalculator" as a child class of "Calculator".
```xml
<class name="FactorizationCalculator" parent="Calculator">
	<method name="GetPrimeFactors" description="Returns the prime factors of this number (without multiplicity)">
		<param name="PrimeFactors" type="structarray" class="PrimeFactor" pass="out" description="The prime factors of this number" />
	</method>
</class>
```
We will implement the actual prime-factor decomposition by overwriting the the "Calculate" method of its parent class.
"FactorizationCalculator" introduces one additional method that outputs an array of structs.

### 3.1.3 The SieveCalculator
Add another class "SieveCalculator" as a child class of "Calculator".
```xml
<class name="SieveCalculator" parent="Calculator">	
	<method name="GetPrimes" description="Returns all prime numbers lower or equal to the sieve's value">
		<param name="Primes" type="basicarray" class="uint64" pass="out" description="The primes lower or equal to the sieve's value" />
	</method>
</class>
```
Again, we will implement the actual calculation of primes by overwriting the the "Calculate" method of its parent class.
"SieveCalculator" introduces one additional method that will return the array of prime numbers of type "uint64".

### 3.1.4 Error: no result available
We want to use the error "LIBPRIMES_ERROR_NORESULTAVAILABLE" when a user tries to retrieve
results from a calculator without having performed a calculation before. Thus add a new error
```xml
<error name="NORESULTAVAILABLE" code="8" description="no result is available" />
```


### 3.1.5 Additions to the global section
The global section requires two more methods, that are used as entry points to the component's functionality:
```xml
<method name="CreateFactorizationCalculator" description="Creates a new FactorizationCalculator instance">
	<param name="Instance" type="handle" class="FactorizationCalculator" pass="return" description="New FactorizationCalculator instance" />
</method>
```
"CreateFactorizationCalculator" specifies a global function that will create a new instance of the "FactorizationCalculator".

```xml
<method name="CreateSieveCalculator" description="Creates a new SieveCalculator instance">
	<param name="Instance" type="handle" class="SieveCalculator" pass="return" description="New SieveCalculator instance" />
</method>
```
"CreateSieveCalculator" specifies a global function that will create a new instance of the "SieveCalculator".

This concludes the complete specification of LibPrimes's interface.

__NOTE__
_You can download the initial IDL-file for libPrimes [here](ressources/315/libPrimes.xml)._

## 3.2. Automatic Source Code Generation
With the complete interface definition in libPrimes.xml, we can now turn on ACT:
```shell
act.exe libPrimes.xml
```
This generates a folder "LibPrimes_component" with two subfolder, "Bindings" and "Implementation".

TODO: image

Let's focus on Implementation/CPP for now, which contains a folder "Interfaces" and "Stubs"

TODO: image

### 3.2.1 Interfaces
Consider all files in the "Interfaces" folder read only for your developement.
They will be regenerated/overwritten if you run ACT again, and you should never have to modify them.
Usually, you will not include them in the source code control system of your component.

The "libprimes_interfaces.hpp" file contains all classes from the IDL as pure abstract C++ classes. E.g. have a look at 
the interfaces "ILibPrimesCalculator" and "ILibPrimesFactorizationCalculator":
```cpp
/*...*/
class ILibPrimesCalculator : public virtual ILibPrimesBaseClass {
public:
	virtual unsigned long long GetValue () = 0;
	virtual void SetValue (const unsigned long long nValue) = 0;
	virtual void Calculate () = 0;
};

class ILibPrimesFactorizationCalculator : public virtual ILibPrimesBaseClass, public virtual ILibPrimesCalculator {
public:
	virtual void GetPrimeFactors (unsigned int nPrimeFactorsBufferSize, unsigned int * pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer) = 0;
};
/*...*/
```

The "libprimes_interfaceexception.hpp" and "libprimes_interfaceexception.cpp" files contain the definition of the component's exception class.

The "libprimes_interfacewrapper.cpp" file implements the forwarding of the C89-interface functions to the classes you will implemenent.
It also translates all exceptions into error codes.
```cpp
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_calculator_getvalue (LibPrimes_Calculator pCalculator, unsigned long long * pValue)
{
	try {
		if (pValue == nullptr)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		ILibPrimesBaseClass* pIBaseClass = (ILibPrimesBaseClass *)pCalculator;
		ILibPrimesCalculator* pICalculator = dynamic_cast<ILibPrimesCalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);

		*pValue = pICalculator->GetValue();

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & E) {
		return E.getErrorCode();
	}
	catch (...) {
		return LIBPRIMES_ERROR_GENERICEXCEPTION;
	}
}
```

### 3.2.2 Implementation Stubs
The files in the "Stubs" folder are the "actual" source code, which you will modify. This will contain your domain logic.

For each class in the IDL, a pair of header and source files has been generated.
They contain a concrete class definition derived from the corresponding interface in interfaces.hpp.

```cpp
class CLibPrimesFactorizationCalculator : public virtual ILibPrimesFactorizationCalculator, public virtual CLibPrimesCalculator {
public:
	void Calculate();
	void GetPrimeFactors (unsigned int nPrimeFactorsBufferSize, unsigned int * pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer);
};
```

The autogenerated implementation of each of a class's methods throws a "non-implemented" exception:
```cpp
void CLibPrimesFactorizationCalculator::GetPrimeFactors (unsigned int nPrimeFactorsBufferSize,
	unsigned int * pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer)
{
	throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NOTIMPLEMENTED);
}
```

## 3.3. Implement and build the library
The "Implementations/CPP" folder already contains a "CMakeLists.txt" that allows you to build the generated sources into a shared library.

The following code snipped sets up a Visual Studio solution:
```bash
cd LibPrimes_component/Implementation/
mkdir _build
cd _build
cmake .. -G "Visual Studio 14 Win64"
cmake --build .
```
Adjust the CMake-Generator for your development environment, if required.

Now we can start actually implementing the library.

### 3.3.1. Required steps for every ACT component
#### GetVersion function
```cpp
void CLibPrimesWrapper::GetLibraryVersion (unsigned int & nMajor, unsigned int & nMinor, unsigned int & nMicro)
{
	nMajor = LIBPRIMES_VERSION_MAJOR;
	nMinor = LIBPRIMES_VERSION_MINOR;
	nMicro = LIBPRIMES_VERSION_MICRO;
}
```

#### CreateFunctions
For all methods in the IDL that are used to return a new instance of a class, code similar to this needs to be implemented.
```cpp
#include "libprimes_factorizationcalculator.hpp"
/*...*/
ILibPrimesFactorizationCalculator * CLibPrimesWrapper::CreateFactorizationCalculator ()
{
	return new CLibPrimesFactorizationCalculator();
}
```

#### Release Function
```cpp
void CLibPrimesWrapper::ReleaseInstance (ILibPrimesBaseClass* pInstance)
{
	delete pInstance;
}
```

__NOTE__:
_Obvioulsy, you can do something more clever/robust, than simply handing out "new"-ed chunks of memory at creation
and deleting them if asked for it via "ReleaseInstance"
(e.g. store them in a set-datastructre, so that you can "free" them if the consumer does not do so).
However, this is solution is fine in the context of ACT, as all automatically generated bindings
(except C) handle the lifetime of all generated ILibPrimesBaseClass instances._

### 3.3.2. Domain code implementation: Steps for LibPrimes
#### CreateFunctions
Implement the missing "CreateSieveCalculator"
```cpp
#include "libprimes_sievecalculator.hpp"
/*...*/
ILibPrimesSieveCalculator * CLibPrimesWrapper::CreateSieveCalculator ()
{
	return new CLibPrimesSieveCalculator();
}
```

#### Calculator
Add a protected member "value" to the CLibPrimesCalculator
```cpp
class CLibPrimesCalculator : public virtual ILibPrimesCalculator {
protected:
	unsined long long m_value;
/*...*/
}
```
The Get/Set-Value methods of the Calculator are straight-forward:
```cpp
unsigned long long CLibPrimesCalculator::GetValue()
{
	return m_value;
}

void CLibPrimesCalculator::SetValue(const unsigned long long nValue)
{
	m_value = nValue;
}
```

We can safely leave the "Calculate"-method untouched, or alternatively, declare it pure virtual,
and remove its implementation.

#### FactorizationCalculator
Add an array that holds the calculated prime factors as private member in "CLibPrimesFactorizationCalculator"
and a public "Calculate" method:
```cpp
class CLibPrimesFactorizationCalculator : public virtual ILibPrimesFactorizationCalculator, public virtual CLibPrimesCalculator
{
private:
	std::vector<sLibPrimesPrimeFactor> primeFactors;
public:
	void Calculate();
/*...*/
}
```

A valid implementation of "GetPrimes" is the following
```cpp
void CLibPrimesFactorizationCalculator::GetPrimeFactors (unsigned int nPrimeFactorsBufferSize,
	unsigned int * pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer)
{
	if (primeFactors.size() == 0)
		throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_NORESULTAVAILABLE);

	if (pPrimeFactorsNeededCount)
		*pPrimeFactorsNeededCount = (unsigned int)primeFactors.size();

	if (nPrimeFactorsBufferSize >= primeFactors.size() && pPrimeFactorsBuffer)
	{
		for (int i = 0; i < primeFactors.size(); i++)
		{
			pPrimeFactorsBuffer[i] = primeFactors[i];
		}
	}
}
```

The following snippet calculates the prime factor decomposition of the calculator's member "m_value".
```cpp
void CLibPrimesFactorizationCalculator::Calculate()
{
	primeFactors.clear();

	unsigned long long nValue = m_value;
	for (unsigned long long i = 2; i <= nValue; i++) {
		sLibPrimesPrimeFactor primeFactor;
		primeFactor.m_Prime = i;
		primeFactor.m_Multiplicity = 0;
		while (nValue % i == 0) {
			primeFactor.m_Multiplicity++;
			nValue = nValue / i;
		}
		if (primeFactor.m_Multiplicity > 0) {
			primeFactors.push_back(primeFactor);
		}
	}
}
```

#### SieveCalculator


# 4. The Consumers

# 4.1. Cpp Dynamic
# 4.2. Python
# 4.2.2. Debug the C++ DLL from a Python Host Application

# 4.2. Pascal

# 5. Extending the interface: with a callback
# 5.1 IDL
# 5.2 Library
# 5.3 Consumer

# 6. Enable journaling

# 7. Conclusion


