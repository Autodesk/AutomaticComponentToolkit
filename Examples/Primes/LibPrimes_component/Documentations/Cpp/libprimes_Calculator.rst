
CCalculator
====================================================================================================


.. cpp:class:: LibPrimes::CCalculator : public CBase 

	




	.. cpp:function:: LibPrimes_uint64 GetValue()

		Returns the current value of this Calculator

		:returns: The current value of this Calculator


	.. cpp:function:: void SetValue(const LibPrimes_uint64 nValue)

		Sets the value to be factorized

		:param nValue: The value to be factorized 


	.. cpp:function:: void Calculate()

		Performs the specific calculation of this Calculator



	.. cpp:function:: void SetProgressCallback(const ProgressCallback pProgressCallback)

		Sets the progress callback function

		:param pProgressCallback: The progress callback 


.. cpp:type:: std::shared_ptr<CCalculator> LibPrimes::PCalculator

	Shared pointer to CCalculator to easily allow reference counting.

