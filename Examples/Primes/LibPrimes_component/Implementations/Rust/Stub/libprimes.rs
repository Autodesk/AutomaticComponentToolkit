/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.8.0-develop.

Abstract: This is an autogenerated Rust implementation file in order to allow easy
development of Prime Numbers Library. It needs to be generated only once.

Interface version: 1.2.0

*/


use libprimes_interfaces::*;
use libprimes_sieve_calculator::CSieveCalculator;
use libprimes_factorization_calculator::CFactorizationCalculator;

// Wrapper struct to implement the wrapper trait for global methods
pub struct CWrapper;

impl Wrapper for CWrapper {

  
  // get_version
  //
  // retrieves the binary version of this library.
  // * @param[out] major - returns the major version of this library
  // * @param[out] minor - returns the minor version of this library
  // * @param[out] micro - returns the micro version of this library
  //
  fn get_version(_major : &mut u32, _minor : &mut u32, _micro : &mut u32) {
    *_major = 1;
    *_minor = 1;
    *_micro = 1;
  }
  
  // get_last_error
  //
  // Returns the last error recorded on this object
  // * @param[in] instance - Instance Handle
  // * @param[out] error_message - Message of the last error
  // * @param[return] has_error - Is there a last error to query
  //
  fn get_last_error(_instance : &mut dyn Base, _error_message : &mut String) -> bool {
    _instance.get_last_error_message(_error_message)
  }
  
  // create_factorization_calculator
  //
  // Creates a new FactorizationCalculator instance
  // * @param[return] instance - New FactorizationCalculator instance
  //
  fn create_factorization_calculator() -> Box<dyn FactorizationCalculator> {
    Box::new(CFactorizationCalculator::new())
  }
  
  // create_sieve_calculator
  //
  // Creates a new SieveCalculator instance
  // * @param[return] instance - New SieveCalculator instance
  //
  fn create_sieve_calculator() -> Box<dyn SieveCalculator> {
    Box::new(CSieveCalculator::new())
  }
  
  // set_journal
  //
  // Handles Library Journaling
  // * @param[in] file_name - Journal FileName
  //
  fn set_journal(_file_name : &str) {
    unimplemented!();
  }
}

