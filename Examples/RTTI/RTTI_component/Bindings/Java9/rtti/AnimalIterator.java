/*++

Copyright (C) 2021 ADSK

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Java file in order to allow an easy
 use of RTTI

Interface version: 1.0.0

*/

package rtti;

import com.sun.jna.Library;
import com.sun.jna.Memory;
import com.sun.jna.Native;
import com.sun.jna.Pointer;
import java.lang.ref.Cleaner;


import java.nio.charset.StandardCharsets;
import java.util.Arrays;
import java.util.List;

public class AnimalIterator extends Base {

	public AnimalIterator(RTTIWrapper wrapper, RTTIHandle handle) {
		super(wrapper, handle);
	}

	/**
	 * Return next animal
	 *
	 * @return 
	 * @throws RTTIException
	 */
	public Animal getNextAnimal() throws RTTIException {
		RTTIHandle handleAnimal = new RTTIHandle();
		mWrapper.checkError(this, mWrapper.rtti_animaliterator_getnextanimal.invokeInt(new java.lang.Object[]{mHandle.Value(), handleAnimal}));
		Animal animal = null;
		if (handleAnimal.Handle != Pointer.NULL) {
		  animal = mWrapper.PolymorphicFactory(handleAnimal, Animal.class);
		}
		return animal;
	}


}

