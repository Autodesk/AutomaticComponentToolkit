
Types used in Prime Numbers Library
==========================================================================================================


Simple types
--------------

	.. cpp:type:: uint8_t LibPrimes_uint8
	
	.. cpp:type:: uint16_t LibPrimes_uint16
	
	.. cpp:type:: uint32_t LibPrimes_uint32
	
	.. cpp:type:: uint64_t LibPrimes_uint64
	
	.. cpp:type:: int8_t LibPrimes_int8
	
	.. cpp:type:: int16_t LibPrimes_int16
	
	.. cpp:type:: int32_t LibPrimes_int32
	
	.. cpp:type:: int64_t LibPrimes_int64
	
	.. cpp:type:: float LibPrimes_single
	
	.. cpp:type:: double LibPrimes_double
	
	.. cpp:type:: LibPrimes_pvoid = void*
	
	.. cpp:type:: LibPrimesResult = LibPrimes_int32
	
	

Structs
--------------

	All structs are defined as `packed`, i.e. with the
	
	.. code-block:: c
		
		#pragma pack (1)

	.. cpp:struct:: sPrimeFactor
	
		.. cpp:member:: LibPrimes_uint64 m_Prime
	
		.. cpp:member:: LibPrimes_uint32 m_Multiplicity
	


Function types
---------------


	.. cpp:type:: ProgressCallback = void(*)(LibPrimes_single, bool*)
		
		Callback to report calculation progress and query whether it should be aborted
		
		:param fProgressPercentage: How far has the calculation progressed?
		:param pShouldAbort: Should the calculation be aborted?
		

	
ELibPrimesException: The standard exception class of Prime Numbers Library
============================================================================================================================================================================================================
	
	Errors in Prime Numbers Library are reported as Exceptions. It is recommended to not throw these exceptions in your client code.
	
	
	.. cpp:class:: LibPrimes::ELibPrimesException
	
		.. cpp:function:: void ELibPrimesException::what() const noexcept
		
			 Returns error message
		
			 :return: the error message of this exception
		
	
		.. cpp:function:: LibPrimesResult ELibPrimesException::getErrorCode() const noexcept
		
			 Returns error code
		
			 :return: the error code of this exception
		
	
CInputVector: Adapter for passing arrays as input for functions
===============================================================================================================================================================
	
	Several functions of Prime Numbers Library expect arrays of integral types or structs as input parameters.
	To not restrict the interface to, say, std::vector<type>,
	and to have a more abstract interface than a location in memory and the number of elements to input to a function
	Prime Numbers Library provides a templated adapter class to pass arrays as input for functions.
	
	Usually, instances of CInputVector are generated anonymously (or even implicitly) in the call to a function that expects an input array.
	
	
	.. cpp:class:: template<typename T> LibPrimes::CInputVector
	
		.. cpp:function:: CInputVector(const std::vector<T>& vec)
	
			Constructs of a CInputVector from a std::vector<T>
	
		.. cpp:function:: CInputVector(const T* in_data, size_t in_size)
	
			Constructs of a CInputVector from a memory address and a given number of elements
	
		.. cpp:function:: const T* CInputVector::data() const
	
			returns the start address of the data captured by this CInputVector
	
		.. cpp:function:: size_t CInputVector::size() const
	
			returns the number of elements captured by this CInputVector
	
 
