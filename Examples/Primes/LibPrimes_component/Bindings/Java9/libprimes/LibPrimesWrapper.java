/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Java file in order to allow an easy
 use of Prime Numbers Library

Interface version: 1.2.0

*/

package libprimes;

import com.sun.jna.*;

import java.nio.charset.StandardCharsets;


public class LibPrimesWrapper {

	public static class EnumConversion {
	}

	public interface ProgressCallback extends Callback {

		void progressCallback (float progressPercentage, Pointer shouldAbort);

	}

	protected Function libprimes_getversion;
	protected Function libprimes_getlasterror;
	protected Function libprimes_acquireinstance;
	protected Function libprimes_releaseinstance;
	protected Function libprimes_createfactorizationcalculator;
	protected Function libprimes_createsievecalculator;
	protected Function libprimes_setjournal;
	protected Function libprimes_calculator_getvalue;
	protected Function libprimes_calculator_setvalue;
	protected Function libprimes_calculator_calculate;
	protected Function libprimes_calculator_setprogresscallback;
	protected Function libprimes_factorizationcalculator_getprimefactors;
	protected Function libprimes_sievecalculator_getprimes;

	protected NativeLibrary mLibrary;

	public LibPrimesWrapper(String libraryPath) {
		mLibrary = NativeLibrary.getInstance(libraryPath);
		libprimes_getversion = mLibrary.getFunction("libprimes_getversion");
		libprimes_getlasterror = mLibrary.getFunction("libprimes_getlasterror");
		libprimes_acquireinstance = mLibrary.getFunction("libprimes_acquireinstance");
		libprimes_releaseinstance = mLibrary.getFunction("libprimes_releaseinstance");
		libprimes_createfactorizationcalculator = mLibrary.getFunction("libprimes_createfactorizationcalculator");
		libprimes_createsievecalculator = mLibrary.getFunction("libprimes_createsievecalculator");
		libprimes_setjournal = mLibrary.getFunction("libprimes_setjournal");
		libprimes_calculator_getvalue = mLibrary.getFunction("libprimes_calculator_getvalue");
		libprimes_calculator_setvalue = mLibrary.getFunction("libprimes_calculator_setvalue");
		libprimes_calculator_calculate = mLibrary.getFunction("libprimes_calculator_calculate");
		libprimes_calculator_setprogresscallback = mLibrary.getFunction("libprimes_calculator_setprogresscallback");
		libprimes_factorizationcalculator_getprimefactors = mLibrary.getFunction("libprimes_factorizationcalculator_getprimefactors");
		libprimes_sievecalculator_getprimes = mLibrary.getFunction("libprimes_sievecalculator_getprimes");
	}

	public LibPrimesWrapper(Pointer lookupPointer) throws LibPrimesException {
		Function lookupMethod = Function.getFunction(lookupPointer);
		libprimes_getversion = loadFunctionByLookup(lookupMethod, "libprimes_getversion");
		libprimes_getlasterror = loadFunctionByLookup(lookupMethod, "libprimes_getlasterror");
		libprimes_acquireinstance = loadFunctionByLookup(lookupMethod, "libprimes_acquireinstance");
		libprimes_releaseinstance = loadFunctionByLookup(lookupMethod, "libprimes_releaseinstance");
		libprimes_createfactorizationcalculator = loadFunctionByLookup(lookupMethod, "libprimes_createfactorizationcalculator");
		libprimes_createsievecalculator = loadFunctionByLookup(lookupMethod, "libprimes_createsievecalculator");
		libprimes_setjournal = loadFunctionByLookup(lookupMethod, "libprimes_setjournal");
		libprimes_calculator_getvalue = loadFunctionByLookup(lookupMethod, "libprimes_calculator_getvalue");
		libprimes_calculator_setvalue = loadFunctionByLookup(lookupMethod, "libprimes_calculator_setvalue");
		libprimes_calculator_calculate = loadFunctionByLookup(lookupMethod, "libprimes_calculator_calculate");
		libprimes_calculator_setprogresscallback = loadFunctionByLookup(lookupMethod, "libprimes_calculator_setprogresscallback");
		libprimes_factorizationcalculator_getprimefactors = loadFunctionByLookup(lookupMethod, "libprimes_factorizationcalculator_getprimefactors");
		libprimes_sievecalculator_getprimes = loadFunctionByLookup(lookupMethod, "libprimes_sievecalculator_getprimes");
	}

	protected void checkError(Base instance, int errorCode) throws LibPrimesException {
		if (instance != null && instance.mWrapper != this) {
			throw new LibPrimesException(LibPrimesException.LIBPRIMES_ERROR_INVALIDCAST, "invalid wrapper call");
		}
		if (errorCode != LibPrimesException.LIBPRIMES_SUCCESS) {
			if (instance != null) {
				GetLastErrorResult result = getLastError(instance);
				throw new LibPrimesException(errorCode, result.ErrorMessage);
			} else {
				throw new LibPrimesException(errorCode, "");
			}
		}
	}

	private Function loadFunctionByLookup(Function lookupMethod, String functionName) throws LibPrimesException {
		byte[] bytes = functionName.getBytes(StandardCharsets.UTF_8);
		Memory name = new Memory(bytes.length+1);
		name.write(0, bytes, 0, bytes.length);
		name.setByte(bytes.length, (byte)0);
		Pointer address = new Memory(8);
		java.lang.Object[] addressParam = new java.lang.Object[]{name, address};
		checkError(null, lookupMethod.invokeInt(addressParam));
		return Function.getFunction(address.getPointer(0));
	}

	/**
	 * retrieves the binary version of this library.
	 *
	 * @return GetVersion Result Tuple
	 * @throws LibPrimesException
	 */
	public GetVersionResult getVersion() throws LibPrimesException {
		Pointer bufferMajor = new Memory(4);
		Pointer bufferMinor = new Memory(4);
		Pointer bufferMicro = new Memory(4);
		checkError(null, libprimes_getversion.invokeInt(new java.lang.Object[]{bufferMajor, bufferMinor, bufferMicro}));
		GetVersionResult returnTuple = new GetVersionResult();
		returnTuple.Major = bufferMajor.getInt(0);
		returnTuple.Minor = bufferMinor.getInt(0);
		returnTuple.Micro = bufferMicro.getInt(0);
		return returnTuple;
	}

	public static class GetVersionResult {
		/**
		 * returns the major version of this library
		 */
		public int Major;

		/**
		 * returns the minor version of this library
		 */
		public int Minor;

		/**
		 * returns the micro version of this library
		 */
		public int Micro;

	}
	/**
	 * Returns the last error recorded on this object
	 *
	 * @param instance Instance Handle
	 * @return GetLastError Result Tuple
	 * @throws LibPrimesException
	 */
	public GetLastErrorResult getLastError(Base instance) throws LibPrimesException {
		Pointer instanceHandle = null;
		if (instance != null) {
			instanceHandle = instance.getHandle();
		} else {
			throw new LibPrimesException(LibPrimesException.LIBPRIMES_ERROR_INVALIDPARAM, "Instance is a null value.");
		}
		Pointer bytesNeededErrorMessage = new Memory(4);
		Pointer bufferHasError = new Memory(1);
		checkError(null, libprimes_getlasterror.invokeInt(new java.lang.Object[]{instanceHandle, 0, bytesNeededErrorMessage, null, bufferHasError}));
		int sizeErrorMessage = bytesNeededErrorMessage.getInt(0);
		Pointer bufferErrorMessage = new Memory(sizeErrorMessage);
		checkError(null, libprimes_getlasterror.invokeInt(new java.lang.Object[]{instanceHandle, sizeErrorMessage, bytesNeededErrorMessage, bufferErrorMessage, bufferHasError}));
		GetLastErrorResult returnTuple = new GetLastErrorResult();
		returnTuple.ErrorMessage = new String(bufferErrorMessage.getByteArray(0, sizeErrorMessage - 1), StandardCharsets.UTF_8);
		returnTuple.HasError = bufferHasError.getByte(0) != 0;
		return returnTuple;
	}

	public static class GetLastErrorResult {
		/**
		 * Message of the last error
		 */
		public String ErrorMessage;

		/**
		 * Is there a last error to query
		 */
		public boolean HasError;

	}
	/**
	 * Acquire shared ownership of an Instance
	 *
	 * @param instance Instance Handle
	 * @throws LibPrimesException
	 */
	public void acquireInstance(Base instance) throws LibPrimesException {
		Pointer instanceHandle = null;
		if (instance != null) {
			instanceHandle = instance.getHandle();
		} else {
			throw new LibPrimesException(LibPrimesException.LIBPRIMES_ERROR_INVALIDPARAM, "Instance is a null value.");
		}
		checkError(null, libprimes_acquireinstance.invokeInt(new java.lang.Object[]{instanceHandle}));
	}

	/**
	 * Releases shared ownership of an Instance
	 *
	 * @param instance Instance Handle
	 * @throws LibPrimesException
	 */
	public void releaseInstance(Base instance) throws LibPrimesException {
		Pointer instanceHandle = null;
		if (instance != null) {
			instanceHandle = instance.getHandle();
		} else {
			throw new LibPrimesException(LibPrimesException.LIBPRIMES_ERROR_INVALIDPARAM, "Instance is a null value.");
		}
		checkError(null, libprimes_releaseinstance.invokeInt(new java.lang.Object[]{instanceHandle}));
	}

	/**
	 * Creates a new FactorizationCalculator instance
	 *
	 * @return New FactorizationCalculator instance
	 * @throws LibPrimesException
	 */
	public FactorizationCalculator createFactorizationCalculator() throws LibPrimesException {
		Pointer bufferInstance = new Memory(8);
		checkError(null, libprimes_createfactorizationcalculator.invokeInt(new java.lang.Object[]{bufferInstance}));
		Pointer valueInstance = bufferInstance.getPointer(0);
		FactorizationCalculator instance = null;
		if (valueInstance == Pointer.NULL) {
		  throw new LibPrimesException(LibPrimesException.LIBPRIMES_ERROR_NORESULTAVAILABLE, "Instance was a null pointer");
		}
		return (valueInstance == Pointer.NULL) ? null : new FactorizationCalculator(this, valueInstance);
	}

	/**
	 * Creates a new SieveCalculator instance
	 *
	 * @return New SieveCalculator instance
	 * @throws LibPrimesException
	 */
	public SieveCalculator createSieveCalculator() throws LibPrimesException {
		Pointer bufferInstance = new Memory(8);
		checkError(null, libprimes_createsievecalculator.invokeInt(new java.lang.Object[]{bufferInstance}));
		Pointer valueInstance = bufferInstance.getPointer(0);
		SieveCalculator instance = null;
		if (valueInstance == Pointer.NULL) {
		  throw new LibPrimesException(LibPrimesException.LIBPRIMES_ERROR_NORESULTAVAILABLE, "Instance was a null pointer");
		}
		return (valueInstance == Pointer.NULL) ? null : new SieveCalculator(this, valueInstance);
	}

	/**
	 * Handles Library Journaling
	 *
	 * @param fileName Journal FileName
	 * @throws LibPrimesException
	 */
	public void setJournal(String fileName) throws LibPrimesException {
		byte[] bytesFileName = fileName.getBytes(StandardCharsets.UTF_8);
		Memory bufferFileName = new Memory(bytesFileName.length + 1);
		bufferFileName.write(0, bytesFileName, 0, bytesFileName.length);
		bufferFileName.setByte(bytesFileName.length, (byte)0);
		checkError(null, libprimes_setjournal.invokeInt(new java.lang.Object[]{bufferFileName}));
	}

}

