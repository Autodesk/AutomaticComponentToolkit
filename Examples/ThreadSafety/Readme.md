# ![ACT logo](../../Documentation/images/ACT_logo_50px.png) Automatic Component Toolkit

## Tutorial: Usage of the `threadsafetyoption`class parameter


## Table of Contents

- [1. Problem](#1-problem)
- [2. Solution](#2-solution)
- [3. Example](#3-example)
- [4. Support](#4-support)

# 1. Problem
The Automatic Component Toolkit (ACT) has an issue with thread safety when returning strings (arrays) from
functions or methods, especially in the C++ Binding/Implementation. The problem arises from the caching mechanism
used in ACT. When API functions or methods that return strings are executed concurrently from multiple
threads in the library consumer code, it can result in crashes.

# 2. Solution
### 2.1. Overview
To address the thread safety issue described above, a new parameter called `threadsafetyoption` has been introduced.
This parameter ensures thread safety by adding mutex locking mechanisms on the library implementation side.

### 2.2. Details
The `ThreadSafetyOption` parameter can be set to `none`, `soft` or `strict` value:
- `none`: Does nothing.
- `soft`: The binding side will call the lock mechanism only when a string is returned from an API function.
- `strict`: The binding side will call the lock mechanism every time, regardless of what is returned from the function.

When the `ThreadSafetyOption` attribute in the component class is set to `soft` or `strict`,
the implementation will derive from a base class that includes mutex locking mechanisms.
This solution ensures thread safety both when a single pointer on the binding side is shared across threads,
and when different pointers that point to the same object on the implementation side are used in different threads.
Additionally, error handling mechanisms are in place to unlock the mutex in case of exception propagation, preventing deadlocks.

# 3. Example
### 3.1. Component definition
In the `threadSafeLibrary.xml` file, there are three classes defined with the same methods but different `threadsafetyoption` values:
- `StringReturner` with `threadsafetyoption` set to `none`.
- `SoftStringReturner` with `threadsafetyoption` set to `soft`.
- `StrictStringReturner`: with `threadsafetyoption` set to `strict`.

```xml
<class name="StringReturner" threadsafetyoption="none">
	<method name="GetString" description="Returns a string">
		<param name="Value" type="string" pass="return"/>
	</method>
	<method name="ThreadSafetyCheck" description="Function that may crash when called from different threads at the same time"></method>
</class>
<class name="SoftStringReturner" threadsafetyoption="soft">
    <method name="GetString" description="Returns a string">
		<param name="Value" type="string" pass="return"/>
	</method>
	<method name="ThreadSafetyCheck" description="Function that may crash when called from different threads at the same time"></method>
</class>
<class name="StrictStringReturner" threadsafetyoption="strict">
	<method name="GetString" description="Returns a string">
		<param name="Value" type="string" pass="return"/>
	</method>
	<method name="ThreadSafetyCheck" description="Function that shouldn't crash when called from different threads at the same time"></method>
</class>
```

Each class has the following methods:
- `GetString` which simply returns a random string.
- `ThreadSafetyCheck` which returns nothing but may crash when executed from different threads simultaneously, as no thread safety mechanism is provided in the library implementation.

### 3.2. Usage
Let's consider below cpp example from `LibThreadSafe_example.cpp`.

```cpp
#include <thread>
#include <vector>
#include <iostream>
#include "libthreadsafe_dynamic.hpp"

template<typename ArrayReturner>
void testStringReturn(ArrayReturner& returner)
{
	std::vector<std::jthread> threads;
	for (int i = 0; i < 5; ++i) {
		threads.emplace_back([&returner]() {
			for (int j = 0; j < 10000; ++j) {
				returner->GetString();
			}
		});
	}
}

template<typename ArrayReturner>
void testThreadSafetyCheck(ArrayReturner& returner)
{
	std::vector<std::jthread> threads;
	for (int i = 0; i < 5; ++i) {
		threads.emplace_back([&returner]() {
			for (int j = 0; j < 10000; ++j) {
				returner->ThreadSafetyCheck();
			}
		});
	}
}

int main()
{
	std::string libpath = ("Path to library");
	auto wrapper = LibThreadSafe::CWrapper::loadLibrary(libpath + ".dll");

	LibThreadSafe::PStringReturner stringReturner = wrapper->CreateStringReturner();
	LibThreadSafe::PSoftStringReturner softStringReturner = wrapper->CreateSoftStringReturner();
	LibThreadSafe::PStrictStringReturner strictStringReturner = wrapper->CreateStrictStringReturner();
}
```

Executing `testStringReturn` and `testThreadSafetyCheck` with `stringReturner` will crash because there is no thread safety mechanism enabled for this instance.

Executing `testStringReturn` with `softStringReturner` will work as expected. However, it will crash on `testThreadSafetyCheck` because with the `soft`
value of the `threadsafetyoption` parameter, the thread safety mechanism will lock the instance on the library side only when executing functions that return strings.

Executing `testStringReturn` or `testThreadSafetyCheck` on `strictStringReturner` will work as expected in both cases because with the `strict` value of
the `threadsafetyoption`, the instance on the implementation side will be locked during each function execution.

# 4. Support
Currently, the `threadsafetyoption` parameter is supported only for Cpp/CppDynamic bindings and Cpp implementation.
