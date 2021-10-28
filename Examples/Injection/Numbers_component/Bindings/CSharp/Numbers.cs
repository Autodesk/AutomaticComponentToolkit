using System;
using System.Text;
using System.Runtime.InteropServices;

namespace Numbers {


	namespace Internal {


		public class NumbersWrapper
		{
			[DllImport("numbers.dll", EntryPoint = "numbers_base_classtypeid", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Base_ClassTypeId (IntPtr Handle, out UInt64 AClassTypeId);

			[DllImport("numbers.dll", EntryPoint = "numbers_variable_getvalue", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Variable_GetValue (IntPtr Handle, out Double AValue);

			[DllImport("numbers.dll", EntryPoint = "numbers_variable_setvalue", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Variable_SetValue (IntPtr Handle, Double AValue);

			[DllImport("numbers.dll", EntryPoint = "numbers_createvariable", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 CreateVariable (Double AInitialValue, out IntPtr AInstance);

			[DllImport("numbers.dll", EntryPoint = "numbers_getversion", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro);

			[DllImport("numbers.dll", EntryPoint = "numbers_getlasterror", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetLastError (IntPtr AInstance, UInt32 sizeErrorMessage, out UInt32 neededErrorMessage, IntPtr dataErrorMessage, out Byte AHasError);

			[DllImport("numbers.dll", EntryPoint = "numbers_releaseinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 ReleaseInstance (IntPtr AInstance);

			[DllImport("numbers.dll", EntryPoint = "numbers_acquireinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 AcquireInstance (IntPtr AInstance);

			[DllImport("numbers.dll", EntryPoint = "numbers_getsymbollookupmethod", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetSymbolLookupMethod (out UInt64 ASymbolLookupMethod);

			public static void ThrowError(IntPtr Handle, Int32 errorCode)
			{
				String sMessage = "Numbers Error";
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
			 *            NumbersWrapper::AcquireInstance(CBase object) must be called after instantiating new object.
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
					case 0x27799F69B3FD1C9E: Object = new CBase(Handle) as T; break; // First 64 bits of SHA1 of a string: "Numbers::Base"
					case 0x23934EDF762423EA: Object = new CVariable(Handle) as T; break; // First 64 bits of SHA1 of a string: "Numbers::Variable"
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
				Internal.NumbersWrapper.ReleaseInstance (Handle);
				Handle = IntPtr.Zero;
			}
		}

		protected void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.NumbersWrapper.ThrowError (Handle, errorCode);
			}
		}

		public IntPtr GetHandle ()
		{
			return Handle;
		}

		public UInt64 ClassTypeId ()
		{
			UInt64 resultClassTypeId = 0;

			CheckError(Internal.NumbersWrapper.Base_ClassTypeId (Handle, out resultClassTypeId));
			return resultClassTypeId;
		}

	}

	public class CVariable : CBase
	{
		public CVariable (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public Double GetValue ()
		{
			Double resultValue = 0;

			CheckError(Internal.NumbersWrapper.Variable_GetValue (Handle, out resultValue));
			return resultValue;
		}

		public void SetValue (Double AValue)
		{

			CheckError(Internal.NumbersWrapper.Variable_SetValue (Handle, AValue));
		}

	}

	class Wrapper
	{
		private static void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.NumbersWrapper.ThrowError (IntPtr.Zero, errorCode);
			}
		}

		public static CVariable CreateVariable (Double AInitialValue)
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError(Internal.NumbersWrapper.CreateVariable (AInitialValue, out newInstance));
			return Internal.NumbersWrapper.PolymorphicFactory<CVariable>(newInstance);
		}

		public static void GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro)
		{

			CheckError(Internal.NumbersWrapper.GetVersion (out AMajor, out AMinor, out AMicro));
		}

		public static bool GetLastError (CBase AInstance, out String AErrorMessage)
		{
			Byte resultHasError = 0;
			UInt32 sizeErrorMessage = 0;
			UInt32 neededErrorMessage = 0;
			CheckError(Internal.NumbersWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, IntPtr.Zero, out resultHasError));
			sizeErrorMessage = neededErrorMessage;
			byte[] bytesErrorMessage = new byte[sizeErrorMessage];
			GCHandle dataErrorMessage = GCHandle.Alloc(bytesErrorMessage, GCHandleType.Pinned);

			CheckError(Internal.NumbersWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, dataErrorMessage.AddrOfPinnedObject(), out resultHasError));
			dataErrorMessage.Free();
			AErrorMessage = Encoding.UTF8.GetString(bytesErrorMessage).TrimEnd(char.MinValue);
			return (resultHasError != 0);
		}

		public static void ReleaseInstance (CBase AInstance)
		{

			CheckError(Internal.NumbersWrapper.ReleaseInstance (AInstance.GetHandle()));
		}

		public static void AcquireInstance (CBase AInstance)
		{

			CheckError(Internal.NumbersWrapper.AcquireInstance (AInstance.GetHandle()));
		}

		public static UInt64 GetSymbolLookupMethod ()
		{
			UInt64 resultSymbolLookupMethod = 0;

			CheckError(Internal.NumbersWrapper.GetSymbolLookupMethod (out resultSymbolLookupMethod));
			return resultSymbolLookupMethod;
		}

	}

}
