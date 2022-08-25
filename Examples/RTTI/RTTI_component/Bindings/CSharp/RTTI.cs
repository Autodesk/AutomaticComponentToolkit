using System;
using System.Text;
using System.Runtime.InteropServices;

namespace RTTI {


	namespace Internal {


		public class RTTIWrapper
		{
			[DllImport("rtti.dll", EntryPoint = "rtti_base_classtypeid", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Base_ClassTypeId (IntPtr Handle, out UInt64 AClassTypeId);

			[DllImport("rtti.dll", EntryPoint = "rtti_animal_name", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Animal_Name (IntPtr Handle, UInt32 sizeResult, out UInt32 neededResult, IntPtr dataResult);

			[DllImport("rtti.dll", EntryPoint = "rtti_tiger_roar", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Tiger_Roar (IntPtr Handle);

			[DllImport("rtti.dll", EntryPoint = "rtti_animaliterator_getnextanimal", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 AnimalIterator_GetNextAnimal (IntPtr Handle, out IntPtr AAnimal);

			[DllImport("rtti.dll", EntryPoint = "rtti_animaliterator_getnextoptinalanimal", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 AnimalIterator_GetNextOptinalAnimal (IntPtr Handle, out IntPtr AAnimal, out Byte AError);

			[DllImport("rtti.dll", EntryPoint = "rtti_animaliterator_getnextmandatoryanimal", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 AnimalIterator_GetNextMandatoryAnimal (IntPtr Handle, out IntPtr AAnimal, out Byte AError);

			[DllImport("rtti.dll", EntryPoint = "rtti_zoo_iterator", CallingConvention=CallingConvention.Cdecl)]
			public unsafe extern static Int32 Zoo_Iterator (IntPtr Handle, out IntPtr AIterator);

			[DllImport("rtti.dll", EntryPoint = "rtti_getversion", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro);

			[DllImport("rtti.dll", EntryPoint = "rtti_getlasterror", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetLastError (IntPtr AInstance, UInt32 sizeErrorMessage, out UInt32 neededErrorMessage, IntPtr dataErrorMessage, out Byte AHasError);

			[DllImport("rtti.dll", EntryPoint = "rtti_releaseinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 ReleaseInstance (IntPtr AInstance);

			[DllImport("rtti.dll", EntryPoint = "rtti_acquireinstance", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 AcquireInstance (IntPtr AInstance);

			[DllImport("rtti.dll", EntryPoint = "rtti_injectcomponent", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 InjectComponent (byte[] ANameSpace, UInt64 ASymbolAddressMethod);

			[DllImport("rtti.dll", EntryPoint = "rtti_getsymbollookupmethod", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 GetSymbolLookupMethod (out UInt64 ASymbolLookupMethod);

			[DllImport("rtti.dll", EntryPoint = "rtti_createzoo", CharSet = CharSet.Ansi, CallingConvention=CallingConvention.Cdecl)]
			public extern static Int32 CreateZoo (out IntPtr AInstance);

			public static void ThrowError(IntPtr Handle, Int32 errorCode)
			{
				String sMessage = "RTTI Error";
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
			 *            RTTIWrapper::AcquireInstance(CBase object) must be called after instantiating new object.
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
					case 0x1549AD28813DAE05: Object = new CBase(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Base"
					case 0x8B40467DA6D327AF: Object = new CAnimal(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Animal"
					case 0xBC9D5FA7750C1020: Object = new CMammal(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Mammal"
					case 0x6756AA8EA5802EC3: Object = new CReptile(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Reptile"
					case 0x9751971BD2C2D958: Object = new CGiraffe(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Giraffe"
					case 0x08D007E7B5F7BAF4: Object = new CTiger(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Tiger"
					case 0x5F6826EF909803B2: Object = new CSnake(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Snake"
					case 0x8E551B208A2E8321: Object = new CTurtle(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Turtle"
					case 0xF1917FE6BBE77831: Object = new CAnimalIterator(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::AnimalIterator"
					case 0x2262ABE80A5E7878: Object = new CZoo(Handle) as T; break; // First 64 bits of SHA1 of a string: "RTTI::Zoo"
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
				Internal.RTTIWrapper.ReleaseInstance (Handle);
				Handle = IntPtr.Zero;
			}
		}

		protected void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.RTTIWrapper.ThrowError (Handle, errorCode);
			}
		}

		public IntPtr GetHandle ()
		{
			return Handle;
		}

		public UInt64 ClassTypeId ()
		{
			UInt64 resultClassTypeId = 0;

			CheckError(Internal.RTTIWrapper.Base_ClassTypeId (Handle, out resultClassTypeId));
			return resultClassTypeId;
		}

	}

	public class CAnimal : CBase
	{
		public CAnimal (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public String Name ()
		{
			UInt32 sizeResult = 0;
			UInt32 neededResult = 0;
			CheckError(Internal.RTTIWrapper.Animal_Name (Handle, sizeResult, out neededResult, IntPtr.Zero));
			sizeResult = neededResult;
			byte[] bytesResult = new byte[sizeResult];
			GCHandle dataResult = GCHandle.Alloc(bytesResult, GCHandleType.Pinned);

			CheckError(Internal.RTTIWrapper.Animal_Name (Handle, sizeResult, out neededResult, dataResult.AddrOfPinnedObject()));
			dataResult.Free();
			return Encoding.UTF8.GetString(bytesResult).TrimEnd(char.MinValue);
		}

	}

	public class CMammal : CAnimal
	{
		public CMammal (IntPtr NewHandle) : base (NewHandle)
		{
		}

	}

	public class CReptile : CAnimal
	{
		public CReptile (IntPtr NewHandle) : base (NewHandle)
		{
		}

	}

	public class CGiraffe : CMammal
	{
		public CGiraffe (IntPtr NewHandle) : base (NewHandle)
		{
		}

	}

	public class CTiger : CMammal
	{
		public CTiger (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public void Roar ()
		{

			CheckError(Internal.RTTIWrapper.Tiger_Roar (Handle));
		}

	}

	public class CSnake : CReptile
	{
		public CSnake (IntPtr NewHandle) : base (NewHandle)
		{
		}

	}

	public class CTurtle : CReptile
	{
		public CTurtle (IntPtr NewHandle) : base (NewHandle)
		{
		}

	}

	public class CAnimalIterator : CBase
	{
		public CAnimalIterator (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public CAnimal GetNextAnimal ()
		{
			IntPtr newAnimal = IntPtr.Zero;

			CheckError(Internal.RTTIWrapper.AnimalIterator_GetNextAnimal (Handle, out newAnimal));
			return Internal.RTTIWrapper.PolymorphicFactory<CAnimal>(newAnimal);
		}

		public bool GetNextOptinalAnimal (out CAnimal AAnimal)
		{
			IntPtr newAnimal = IntPtr.Zero;
			Byte resultError = 0;

			CheckError(Internal.RTTIWrapper.AnimalIterator_GetNextOptinalAnimal (Handle, out newAnimal, out resultError));
			AAnimal = Internal.RTTIWrapper.PolymorphicFactory<CAnimal>(newAnimal);
			return (resultError != 0);
		}

		public bool GetNextMandatoryAnimal (out CAnimal AAnimal)
		{
			IntPtr newAnimal = IntPtr.Zero;
			Byte resultError = 0;

			CheckError(Internal.RTTIWrapper.AnimalIterator_GetNextMandatoryAnimal (Handle, out newAnimal, out resultError));
			AAnimal = Internal.RTTIWrapper.PolymorphicFactory<CAnimal>(newAnimal);
			return (resultError != 0);
		}

	}

	public class CZoo : CBase
	{
		public CZoo (IntPtr NewHandle) : base (NewHandle)
		{
		}

		public CAnimalIterator Iterator ()
		{
			IntPtr newIterator = IntPtr.Zero;

			CheckError(Internal.RTTIWrapper.Zoo_Iterator (Handle, out newIterator));
			return Internal.RTTIWrapper.PolymorphicFactory<CAnimalIterator>(newIterator);
		}

	}

	class Wrapper
	{
		private static void CheckError (Int32 errorCode)
		{
			if (errorCode != 0) {
				Internal.RTTIWrapper.ThrowError (IntPtr.Zero, errorCode);
			}
		}

		public static void GetVersion (out UInt32 AMajor, out UInt32 AMinor, out UInt32 AMicro)
		{

			CheckError(Internal.RTTIWrapper.GetVersion (out AMajor, out AMinor, out AMicro));
		}

		public static bool GetLastError (CBase AInstance, out String AErrorMessage)
		{
			Byte resultHasError = 0;
			UInt32 sizeErrorMessage = 0;
			UInt32 neededErrorMessage = 0;
			CheckError(Internal.RTTIWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, IntPtr.Zero, out resultHasError));
			sizeErrorMessage = neededErrorMessage;
			byte[] bytesErrorMessage = new byte[sizeErrorMessage];
			GCHandle dataErrorMessage = GCHandle.Alloc(bytesErrorMessage, GCHandleType.Pinned);

			CheckError(Internal.RTTIWrapper.GetLastError (AInstance.GetHandle(), sizeErrorMessage, out neededErrorMessage, dataErrorMessage.AddrOfPinnedObject(), out resultHasError));
			dataErrorMessage.Free();
			AErrorMessage = Encoding.UTF8.GetString(bytesErrorMessage).TrimEnd(char.MinValue);
			return (resultHasError != 0);
		}

		public static void ReleaseInstance (CBase AInstance)
		{

			CheckError(Internal.RTTIWrapper.ReleaseInstance (AInstance.GetHandle()));
		}

		public static void AcquireInstance (CBase AInstance)
		{

			CheckError(Internal.RTTIWrapper.AcquireInstance (AInstance.GetHandle()));
		}

		public static void InjectComponent (String ANameSpace, UInt64 ASymbolAddressMethod)
		{
		throw new Exception("Component injection is not supported in CSharp.");
		}

		public static UInt64 GetSymbolLookupMethod ()
		{
			UInt64 resultSymbolLookupMethod = 0;

			CheckError(Internal.RTTIWrapper.GetSymbolLookupMethod (out resultSymbolLookupMethod));
			return resultSymbolLookupMethod;
		}

		public static CZoo CreateZoo ()
		{
			IntPtr newInstance = IntPtr.Zero;

			CheckError(Internal.RTTIWrapper.CreateZoo (out newInstance));
			return Internal.RTTIWrapper.PolymorphicFactory<CZoo>(newInstance);
		}

	}

}