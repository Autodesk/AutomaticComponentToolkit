using System;
using System.Text;
using System.Runtime.InteropServices;

namespace LibPrimes {

	public struct sPrimeFactor
	{
		public UInt64 Prime;
		public UInt32 Multiplicity;
	}


	namespace Internal {

		[StructLayout(LayoutKind.Explicit, Size=12)]
		public unsafe struct InternalPrimeFactor
		{
			[FieldOffset(0)] public UInt64 Prime;
			[FieldOffset(8)] public UInt32 Multiplicity;
		}


		public class LibPrimesWrapper
		{
			[DllImport("libprimes.dll", EntryPoint = "libprimes_calculator_getvalue", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_GetValue (IntPtr Handle, out UInt64 AValue);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_calculator_setvalue", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_SetValue (IntPtr Handle, UInt64 AValue);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_calculator_calculate", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_Calculate (IntPtr Handle);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_calculator_setprogresscallback", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_SetProgressCallback (IntPtr Handle, IntPtr AProgressCallback);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_factorizationcalculator_getprimefactors", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 FactorizationCalculator_GetPrimeFactors (IntPtr Handle, UInt64 sizePrimeFactors, out UInt64 neededPrimeFactors, IntPtr dataPrimeFactors);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_sievecalculator_getprimes", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 SieveCalculator_GetPrimes (IntPtr Handle, UInt64 sizePrimes, out UInt64 neededPrimes, IntPtr dataPrimes);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_getversion", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_getlasterror", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetLastError (IntPtr AInstance, UInt32 sizeErrorMessage, out UInt32 neededErrorMessage, IntPtr dataErrorMessage, out Byte AHasError);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_releaseinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 ReleaseInstance (IntPtr AInstance);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_createfactorizationcalculator", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 CreateFactorizationCalculator (out IntPtr AInstance);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_createsievecalculator", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 CreateSieveCalculator (out IntPtr AInstance);

			[DllImport("libprimes.dll", EntryPoint = "libprimes_setjournal", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 SetJournal (byte[] AFileName);

			public unsafe static sPrimeFactor convertInternalToStruct_PrimeFactor (InternalPrimeFactor intPrimeFactor)
			{
				sPrimeFactor PrimeFactor;
				PrimeFactor.Prime = intPrimeFactor.Prime;
				PrimeFactor.Multiplicity = intPrimeFactor.Multiplicity;
				return PrimeFactor;
			}

			public unsafe static InternalPrimeFactor convertStructToInternal_PrimeFactor (sPrimeFactor PrimeFactor)
			{
				InternalPrimeFactor intPrimeFactor;
				intPrimeFactor.Prime = PrimeFactor.Prime;
				intPrimeFactor.Multiplicity = PrimeFactor.Multiplicity;
				return intPrimeFactor;
			}

			public static void ThrowError(IntPtr Handle, Int32 errorCode)
			{
				String sMessage = "LibPrimes Error";
				if (Handle != IntPtr.Zero) {
					UInt32 sizeMessage = 0;
					UInt32 neededMessage = 0;
					Byte hasLastError = 0;
					Int32 resultCode1 = GetLastError (Handle, sizeMessage, out neededMessage, IntPtr.Zero, out hasLastError);
					if ((resultCode1 == 0) && (hasLastError != 0)) {
						sizeMessage = neededMessage + 1;
						byte[] bytesMessage = new byte[sizeMessage];

						GCHandle dataMessage = GCHandle.Alloc(bytesMessage, GCHandleType.Pinned);
						Int32 resultCode2 = GetLastError(Handle, sizeMessage, out neededMessage, dataMessage.AddrOfPinnedObject(), out hasLastError);
						dataMessage.Free();

						if ((resultCode2 == 0) && (hasLastError != 0)) {
							sMessage = sMessage + ": " + Encoding.UTF8.GetString(bytesMessage).TrimEnd(char.MinValue);
						}
					}
				}

				throw new Exception(sMessage + "(# " + errorCode + ")");
			}

		}
	}


	class CBase 
	{
		protected IntPtr Handle;

		public CBase (IntPtr NewHandle)
		{
			Handle = NewHandle;
		}

		~CBase ()
		{
			if (Handle != IntPtr.Zero) {
				Internal.LibPrimesWrapper.ReleaseInstance (Handle);
				Handle = IntPtr.Zero;
			}
		}

		protected void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.LibPrimesWrapper.ThrowError (Handle, errorCode);
			}
		}

		public IntPtr GetHandle ()
		{
			return Handle;
		}

	}

	class CCalculator : CBase
	{
		public CCalculator (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public UInt64 GetValue ()
		{
			UInt64 resultValue = 0;

			CheckError (Internal.LibPrimesWrapper.Calculator_GetValue (Handle, out resultValue));
			return resultValue;
		}

		public void SetValue (UInt64 AValue)
		{

			CheckError (Internal.LibPrimesWrapper.Calculator_SetValue (Handle, AValue));
		}

		public void Calculate ()
		{

			CheckError (Internal.LibPrimesWrapper.Calculator_Calculate (Handle));
		}

		public void SetProgressCallback (IntPtr AProgressCallback)
		{

			CheckError (Internal.LibPrimesWrapper.Calculator_SetProgressCallback (Handle, IntPtr.Zero));
		}

	}

	class CFactorizationCalculator : CCalculator
	{
		public CFactorizationCalculator (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public void GetPrimeFactors (out sPrimeFactor[] APrimeFactors)
		{
			UInt64 sizePrimeFactors = 0;
			UInt64 neededPrimeFactors = 0;
			CheckError (Internal.LibPrimesWrapper.FactorizationCalculator_GetPrimeFactors (Handle, sizePrimeFactors, out neededPrimeFactors, IntPtr.Zero));
			sizePrimeFactors = neededPrimeFactors;
			var arrayPrimeFactors = new Internal.InternalPrimeFactor[sizePrimeFactors];
			GCHandle dataPrimeFactors = GCHandle.Alloc(arrayPrimeFactors, GCHandleType.Pinned);

			CheckError (Internal.LibPrimesWrapper.FactorizationCalculator_GetPrimeFactors (Handle, sizePrimeFactors, out neededPrimeFactors, dataPrimeFactors.AddrOfPinnedObject()));
			dataPrimeFactors.Free();
			APrimeFactors = new sPrimeFactor[sizePrimeFactors];
			for (int index = 0; index < APrimeFactors.Length; index++)
				APrimeFactors[index] = Internal.LibPrimesWrapper.convertInternalToStruct_PrimeFactor(arrayPrimeFactors[index]);
		}

	}

	class CSieveCalculator : CCalculator
	{
		public CSieveCalculator (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public void GetPrimes (out UInt64[] APrimes)
		{
			UInt64 sizePrimes = 0;
			UInt64 neededPrimes = 0;
			CheckError (Internal.LibPrimesWrapper.SieveCalculator_GetPrimes (Handle, sizePrimes, out neededPrimes, IntPtr.Zero));
			sizePrimes = neededPrimes;
			APrimes = new UInt64[sizePrimes];
			GCHandle dataPrimes = GCHandle.Alloc(APrimes, GCHandleType.Pinned);

			CheckError (Internal.LibPrimesWrapper.SieveCalculator_GetPrimes (Handle, sizePrimes, out neededPrimes, dataPrimes.AddrOfPinnedObject()));
			dataPrimes.Free();
		}

	}

	class Wrapper
	{
		private static void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.LibPrimesWrapper.ThrowError (IntPtr.Zero, errorCode);
			}
		}

		public static void GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro)
		{

			CheckError (Internal.LibPrimesWrapper.GetVersion (out AMajor, out AMinor, out AMicro));
		}

		public static bool GetLastError (CBase AInstance, out String AErrorMessage)
		{
			Byte resultHasError = 0;
			UInt32 sizeErrorMessage = 0;
			UInt32 neededErrorMessage = 0;
			CheckError (Internal.LibPrimesWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, IntPtr.Zero, out resultHasError));
			sizeErrorMessage = neededErrorMessage + 1;
			byte[] bytesErrorMessage = new byte[sizeErrorMessage];
			GCHandle dataErrorMessage = GCHandle.Alloc(bytesErrorMessage, GCHandleType.Pinned);

			CheckError (Internal.LibPrimesWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, dataErrorMessage.AddrOfPinnedObject(), out resultHasError));
			dataErrorMessage.Free();
			AErrorMessage = Encoding.UTF8.GetString(bytesErrorMessage).TrimEnd(char.MinValue);
			return (resultHasError != 0);
		}

		public static void ReleaseInstance (CBase AInstance)
		{

			CheckError (Internal.LibPrimesWrapper.ReleaseInstance (AInstance.GetHandle()));
		}

		public static CFactorizationCalculator CreateFactorizationCalculator ()
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError (Internal.LibPrimesWrapper.CreateFactorizationCalculator (out newInstance));
			return new CFactorizationCalculator (newInstance );
		}

		public static CSieveCalculator CreateSieveCalculator ()
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError (Internal.LibPrimesWrapper.CreateSieveCalculator (out newInstance));
			return new CSieveCalculator (newInstance );
		}

		public static void SetJournal (String AFileName)
		{
			byte[] byteFileName = Encoding.UTF8.GetBytes(AFileName + char.MinValue);

			CheckError (Internal.LibPrimesWrapper.SetJournal (byteFileName));
		}

	}

}
