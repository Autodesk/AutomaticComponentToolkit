/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated C++ header file in order to allow easy
development of Prime Numbers Library. It provides an automatic Journaling mechanism for the library implementation.

Interface version: 1.2.0

*/

#ifndef __LIBPRIMES_INTERFACEJOURNAL_HEADER
#define __LIBPRIMES_INTERFACEJOURNAL_HEADER

#include <iostream>
#include <fstream>
#include <string>
#include <memory>
#include <list>
#include <mutex>
#include <chrono>
#include "libprimes_types.hpp"

/*************************************************************************************************************************
 Class CLibPrimesInterfaceJournal 
**************************************************************************************************************************/

class CLibPrimesInterfaceJournal;

class CLibPrimesInterfaceJournalEntry {

	protected:

		CLibPrimesInterfaceJournal * m_pJournal;
		LibPrimesResult m_ErrorCode;
		std::string m_sClassName;
		std::string m_sMethodName;
		std::string m_sInstanceHandle;
		LibPrimes_uint64 m_nInitTimeStamp;
		LibPrimes_uint64 m_nFinishTimeStamp;
		std::list<std::pair<std::pair<std::string, std::string>, std::string>> m_sParameters;
		std::list<std::pair<std::pair<std::string, std::string>, std::string>> m_sResultValues;

		std::string getXMLString();
		void addParameter (const std::string & sName, const std::string & sParameterType, const std::string & sParameterValue);
		void addResult (const std::string & sName, const std::string & sResultType, const std::string & sResultValue);

	public:
		CLibPrimesInterfaceJournalEntry(CLibPrimesInterfaceJournal * pJournal, std::string sClassName, std::string sMethodName, LibPrimesHandle pInstanceHandle);
		~CLibPrimesInterfaceJournalEntry();

		void writeSuccess ();
		void writeError (LibPrimesResult nErrorCode);

		void addBooleanParameter(const std::string & sName, const bool bValue);
		void addUInt8Parameter(const std::string & sName, const LibPrimes_uint8 nValue);
		void addUInt16Parameter(const std::string & sName, const LibPrimes_uint16 nValue);
		void addUInt32Parameter(const std::string & sName, const LibPrimes_uint32 nValue);
		void addUInt64Parameter(const std::string & sName, const LibPrimes_uint64 nValue);
		void addInt8Parameter(const std::string & sName, const LibPrimes_int8 nValue);
		void addInt16Parameter(const std::string & sName, const LibPrimes_int16 nValue);
		void addInt32Parameter(const std::string & sName, const LibPrimes_int32 nValue);
		void addInt64Parameter(const std::string & sName, const LibPrimes_int64 nValue);
		void addSingleParameter(const std::string & sName, const LibPrimes_single fValue);
		void addDoubleParameter(const std::string & sName, const LibPrimes_double dValue);
		void addPointerParameter(const std::string & sName, const LibPrimes_pvoid pValue);
		void addStringParameter(const std::string & sName, const char * pValue);
		void addHandleParameter(const std::string & sName, const LibPrimesHandle pHandle);
		void addEnumParameter(const std::string & sName, const std::string & sEnumType, const LibPrimes_int32 nValue);

		void addBooleanResult(const std::string & sName, const bool bValue);
		void addUInt8Result(const std::string & sName, const LibPrimes_uint8 nValue);
		void addUInt16Result(const std::string & sName, const LibPrimes_uint16 nValue);
		void addUInt32Result(const std::string & sName, const LibPrimes_uint32 nValue);
		void addUInt64Result(const std::string & sName, const LibPrimes_uint64 nValue);
		void addInt8Result(const std::string & sName, const LibPrimes_int8 nValue);
		void addInt16Result(const std::string & sName, const LibPrimes_int16 nValue);
		void addInt32Result(const std::string & sName, const LibPrimes_int32 nValue);
		void addInt64Result(const std::string & sName, const LibPrimes_int64 nValue);
		void addSingleResult(const std::string & sName, const LibPrimes_single fValue);
		void addDoubleResult(const std::string & sName, const LibPrimes_double dValue);
		void addPointerResult(const std::string & sName, const LibPrimes_pvoid pValue);
		void addStringResult(const std::string & sName, const char * pValue);
		void addHandleResult(const std::string & sName, const LibPrimesHandle pHandle);
		void addEnumResult(const std::string & sName, const std::string & sEnumType, const LibPrimes_int32 nValue);

friend class CLibPrimesInterfaceJournal;

};

typedef std::shared_ptr<CLibPrimesInterfaceJournalEntry> PLibPrimesInterfaceJournalEntry;



class CLibPrimesInterfaceJournal {

	protected:

		std::string m_sFileName;
		std::mutex m_Mutex;
		std::ofstream m_Stream;
		std::chrono::time_point<std::chrono::high_resolution_clock> m_StartTime;
		void writeEntry (CLibPrimesInterfaceJournalEntry * pEntry);
		LibPrimes_uint64 getTimeStamp ();

	public:

		CLibPrimesInterfaceJournal (const std::string & sFileName);
		~CLibPrimesInterfaceJournal ();
		PLibPrimesInterfaceJournalEntry beginClassMethod (const LibPrimesHandle pHandle, const std::string & sClassName, const std::string & sMethodName);
		PLibPrimesInterfaceJournalEntry beginStaticFunction (const std::string & sMethodName);
		friend class CLibPrimesInterfaceJournalEntry;
};

typedef std::shared_ptr<CLibPrimesInterfaceJournal> PLibPrimesInterfaceJournal;

#endif // __LIBPRIMES_INTERFACEJOURNAL_HEADER

