/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Java file in order to allow an easy
 use of Numbers library

Interface version: 1.0.0

*/

package numbers;

import java.util.HashMap;
import java.util.Map;

public class NumbersException extends Exception {

	// Error Constants for Numbers
	public static final int NUMBERS_SUCCESS = 0;
	public static final int NUMBERS_ERROR_NOTIMPLEMENTED = 1;
	public static final int NUMBERS_ERROR_INVALIDPARAM = 2;
	public static final int NUMBERS_ERROR_INVALIDCAST = 3;
	public static final int NUMBERS_ERROR_BUFFERTOOSMALL = 4;
	public static final int NUMBERS_ERROR_GENERICEXCEPTION = 5;
	public static final int NUMBERS_ERROR_COULDNOTLOADLIBRARY = 6;
	public static final int NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT = 7;
	public static final int NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION = 8;

	public static final Map<Integer, String> ErrorCodeMap = new HashMap<Integer, String>();
	public static final Map<Integer, String> ErrorDescriptionMap = new HashMap<Integer, String>();

	static {
		ErrorCodeMap.put(NUMBERS_ERROR_NOTIMPLEMENTED, "NUMBERS_ERROR_NOTIMPLEMENTED");
		ErrorDescriptionMap.put(NUMBERS_ERROR_NOTIMPLEMENTED, "functionality not implemented");
		ErrorCodeMap.put(NUMBERS_ERROR_INVALIDPARAM, "NUMBERS_ERROR_INVALIDPARAM");
		ErrorDescriptionMap.put(NUMBERS_ERROR_INVALIDPARAM, "an invalid parameter was passed");
		ErrorCodeMap.put(NUMBERS_ERROR_INVALIDCAST, "NUMBERS_ERROR_INVALIDCAST");
		ErrorDescriptionMap.put(NUMBERS_ERROR_INVALIDCAST, "a type cast failed");
		ErrorCodeMap.put(NUMBERS_ERROR_BUFFERTOOSMALL, "NUMBERS_ERROR_BUFFERTOOSMALL");
		ErrorDescriptionMap.put(NUMBERS_ERROR_BUFFERTOOSMALL, "a provided buffer is too small");
		ErrorCodeMap.put(NUMBERS_ERROR_GENERICEXCEPTION, "NUMBERS_ERROR_GENERICEXCEPTION");
		ErrorDescriptionMap.put(NUMBERS_ERROR_GENERICEXCEPTION, "a generic exception occurred");
		ErrorCodeMap.put(NUMBERS_ERROR_COULDNOTLOADLIBRARY, "NUMBERS_ERROR_COULDNOTLOADLIBRARY");
		ErrorDescriptionMap.put(NUMBERS_ERROR_COULDNOTLOADLIBRARY, "the library could not be loaded");
		ErrorCodeMap.put(NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT, "NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT");
		ErrorDescriptionMap.put(NUMBERS_ERROR_COULDNOTFINDLIBRARYEXPORT, "a required exported symbol could not be found in the library");
		ErrorCodeMap.put(NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION, "NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION");
		ErrorDescriptionMap.put(NUMBERS_ERROR_INCOMPATIBLEBINARYVERSION, "the version of the binary interface does not match the bindings interface");
	}

	protected int mErrorCode;

	protected String mErrorString;

	protected String mErrorDescription;

	public NumbersException(int errorCode, String message){
		super(message);
		mErrorCode = errorCode;
		mErrorString = ErrorCodeMap.get(errorCode);
		mErrorString = (mErrorString != null) ? mErrorString : "Unknown error code";
		mErrorDescription = ErrorDescriptionMap.get(errorCode);
		mErrorDescription = (mErrorDescription != null) ? mErrorDescription : "";
	}

	@Override
	public String toString() {
		return mErrorCode + ": " + mErrorString + " (" + mErrorDescription + " - " + getMessage() + ")";
	}
}
