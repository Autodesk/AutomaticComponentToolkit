
The wrapper class CWrapper
===================================================================================


.. cpp:class:: LibPrimes::CWrapper

	All types of Prime Numbers Library reside in the namespace LibPrimes and all
	functionality of Prime Numbers Library resides in LibPrimes::CWrapper.

	A suitable way to use LibPrimes::CWrapper is as a singleton.

	.. cpp:function:: void GetVersion(LibPrimes_uint32 & nMajor, LibPrimes_uint32 & nMinor, LibPrimes_uint32 & nMicro)
	
		retrieves the binary version of this library.
	
		:param nMajor: returns the major version of this library 
		:param nMinor: returns the minor version of this library 
		:param nMicro: returns the micro version of this library 

	
	.. cpp:function:: bool GetLastError(classParam<CBase> pInstance, std::string & sErrorMessage)
	
		Returns the last error recorded on this object
	
		:param pInstance: Instance Handle 
		:param sErrorMessage: Message of the last error 
		:returns: Is there a last error to query

	
	.. cpp:function:: void AcquireInstance(classParam<CBase> pInstance)
	
		Acquire shared ownership of an Instance
	
		:param pInstance: Instance Handle 

	
	.. cpp:function:: void ReleaseInstance(classParam<CBase> pInstance)
	
		Releases shared ownership of an Instance
	
		:param pInstance: Instance Handle 

	
	.. cpp:function:: PFactorizationCalculator CreateFactorizationCalculator()
	
		Creates a new FactorizationCalculator instance
	
		:returns: New FactorizationCalculator instance

	
	.. cpp:function:: PSieveCalculator CreateSieveCalculator()
	
		Creates a new SieveCalculator instance
	
		:returns: New SieveCalculator instance

	
	.. cpp:function:: void SetJournal(const std::string & sFileName)
	
		Handles Library Journaling
	
		:param sFileName: Journal FileName 

	
.. cpp:type:: std::shared_ptr<CWrapper> LibPrimes::PWrapper
	
