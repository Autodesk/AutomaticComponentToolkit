using System;
using System.Text;
using System.Runtime.InteropServices;

namespace Calculation {


	namespace Internal {


		public class CalculationWrapper
		{
			[DllImport("calculation.dll", EntryPoint = "calculation_base_classtypeid", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Base_ClassTypeId (IntPtr Handle, out UInt64 AClassTypeId);

			[DllImport("calculation.dll", EntryPoint = "calculation_calculator_enlistvariable", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_EnlistVariable (IntPtr Handle, IntPtr AVariable);

			[DllImport("calculation.dll", EntryPoint = "calculation_calculator_getenlistedvariable", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_GetEnlistedVariable (IntPtr Handle, UInt32 AIndex, out IntPtr AVariable);

			[DllImport("calculation.dll", EntryPoint = "calculation_calculator_clearvariables", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_ClearVariables (IntPtr Handle);

			[DllImport("calculation.dll", EntryPoint = "calculation_calculator_multiply", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_Multiply (IntPtr Handle, out IntPtr AInstance);

			[DllImport("calculation.dll", EntryPoint = "calculation_calculator_add", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Calculator_Add (IntPtr Handle, out IntPtr AInstance);

			[DllImport("calculation.dll", EntryPoint = "calculation_createcalculator", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 CreateCalculator (out IntPtr AInstance);

			[DllImport("calculation.dll", EntryPoint = "calculation_getversion", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro);

			[DllImport("calculation.dll", EntryPoint = "calculation_getlasterror", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetLastError (IntPtr AInstance, UInt32 sizeErrorMessage, out UInt32 neededErrorMessage, IntPtr dataErrorMessage, out Byte AHasError);

			[DllImport("calculation.dll", EntryPoint = "calculation_releaseinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 ReleaseInstance (IntPtr AInstance);

			[DllImport("calculation.dll", EntryPoint = "calculation_acquireinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 AcquireInstance (IntPtr AInstance);

			[DllImport("calculation.dll", EntryPoint = "calculation_injectcomponent", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 InjectComponent (byte[] ANameSpace, UInt64 ASymbolAddressMethod);

			[DllImport("calculation.dll", EntryPoint = "calculation_getsymbollookupmethod", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetSymbolLookupMethod (out UInt64 ASymbolLookupMethod);

			public static void ThrowError(IntPtr Handle, Int32 errorCode)
			{
				String sMessage = "Calculation Error";
				if (Handle != IntPtr.Zero) {
					UInt32 sizeMessage = 0;
					UInt32 neededMessage = 0;
					Byte hasLastError = 0;
					Int32 resultCode1 = GetLastError (Handle, sizeMessage, out neededMessage, IntPtr.Zero, out hasLastError);
					if ((resultCode1 == 0) && (hasLastError != 0)) {
						sizeMessage = neededMessage;
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

			/**
			 * IMPORTANT: PolymorphicFactory method should not be used by application directly.
			 *            It's designed to be used on Handle object only once.
			 *            If it's used on any existing object as a form of dynamic cast then
			 *            CalculationWrapper::AcquireInstance(CBase object) must be called after instantiating new object.
			 *            This is important to keep reference count matching between application and library sides.
			*/
			public static T PolymorphicFactory<T>(IntPtr Handle) where T : class
			{
				T Object;
				if (Handle == IntPtr.Zero)
					return System.Activator.CreateInstance(typeof(T), Handle) as T;
				
				UInt64 resultClassTypeId = 0;
				Int32 errorCode = Base_ClassTypeId (Handle, out resultClassTypeId);
				if (errorCode != 0)
					ThrowError (IntPtr.Zero, errorCode);
				switch (resultClassTypeId) {
					case 0x3BA5271BAB410E5D: Object = new CBase(Handle) as T; break; // First 64 bits of SHA1 of a string: "Calculation::Base"
					case 0xB23F514353D0C606: Object = new CCalculator(Handle) as T; break; // First 64 bits of SHA1 of a string: "Calculation::Calculator"
					default: Object = System.Activator.CreateInstance(typeof(T), Handle) as T; break;
				}
				return Object;
			}

		}
	}


	public class CBase 
	{
		protected IntPtr Handle;

		public CBase (IntPtr NewHandle)
		{
			Handle = NewHandle;
		}

		~CBase ()
		{
			if (Handle != IntPtr.Zero) {
				Internal.CalculationWrapper.ReleaseInstance (Handle);
				Handle = IntPtr.Zero;
			}
		}

		protected void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.CalculationWrapper.ThrowError (Handle, errorCode);
			}
		}

		public IntPtr GetHandle ()
		{
			return Handle;
		}

		public UInt64 ClassTypeId ()
		{
			UInt64 resultClassTypeId = 0;

			CheckError(Internal.CalculationWrapper.Base_ClassTypeId (Handle, out resultClassTypeId));
			return resultClassTypeId;
		}

	}

	public class CCalculator : CBase
	{
		public CCalculator (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public void EnlistVariable (IntPtr AVariable)
		{

			CheckError(Internal.CalculationWrapper.Calculator_EnlistVariable (Handle, AVariable));
		}

		public IntPtr GetEnlistedVariable (UInt32 AIndex)
		{
			IntPtr newVariable = IntPtr.Zero;

			CheckError(Internal.CalculationWrapper.Calculator_GetEnlistedVariable (Handle, AIndex, out newVariable));
			return newVariable;
		}

		public void ClearVariables ()
		{

			CheckError(Internal.CalculationWrapper.Calculator_ClearVariables (Handle));
		}

		public IntPtr Multiply ()
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError(Internal.CalculationWrapper.Calculator_Multiply (Handle, out newInstance));
			return newInstance;
		}

		public IntPtr Add ()
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError(Internal.CalculationWrapper.Calculator_Add (Handle, out newInstance));
			return newInstance;
		}

	}

	class Wrapper
	{
		private static void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.CalculationWrapper.ThrowError (IntPtr.Zero, errorCode);
			}
		}

		public static CCalculator CreateCalculator ()
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError(Internal.CalculationWrapper.CreateCalculator (out newInstance));
			return Internal.CalculationWrapper.PolymorphicFactory<CCalculator>(newInstance);
		}

		public static void GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro)
		{

			CheckError(Internal.CalculationWrapper.GetVersion (out AMajor, out AMinor, out AMicro));
		}

		public static bool GetLastError (CBase AInstance, out String AErrorMessage)
		{
			Byte resultHasError = 0;
			UInt32 sizeErrorMessage = 0;
			UInt32 neededErrorMessage = 0;
			CheckError(Internal.CalculationWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, IntPtr.Zero, out resultHasError));
			sizeErrorMessage = neededErrorMessage;
			byte[] bytesErrorMessage = new byte[sizeErrorMessage];
			GCHandle dataErrorMessage = GCHandle.Alloc(bytesErrorMessage, GCHandleType.Pinned);

			CheckError(Internal.CalculationWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, dataErrorMessage.AddrOfPinnedObject(), out resultHasError));
			dataErrorMessage.Free();
			AErrorMessage = Encoding.UTF8.GetString(bytesErrorMessage).TrimEnd(char.MinValue);
			return (resultHasError != 0);
		}

		public static void ReleaseInstance (CBase AInstance)
		{

			CheckError(Internal.CalculationWrapper.ReleaseInstance (AInstance.GetHandle()));
		}

		public static void AcquireInstance (CBase AInstance)
		{

			CheckError(Internal.CalculationWrapper.AcquireInstance (AInstance.GetHandle()));
		}

		public static void InjectComponent (String ANameSpace, UInt64 ASymbolAddressMethod)
		{
		throw new Exception("Component injection is not supported in CSharp.");
		}

		public static UInt64 GetSymbolLookupMethod ()
		{
			UInt64 resultSymbolLookupMethod = 0;

			CheckError(Internal.CalculationWrapper.GetSymbolLookupMethod (out resultSymbolLookupMethod));
			return resultSymbolLookupMethod;
		}

	}

}
